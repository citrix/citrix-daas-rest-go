package citrixclient

import (
	"context"
	"crypto/tls"
	"math"
	"net/http"
	"reflect"
	"time"

	"github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

type CitrixDaasClient struct {
	ApiClient    *citrixorchestration.APIClient
	AuthConfig   *AuthenticationConfiguration
	ClientConfig *ClientConfiguration
	AuthToken    *AuthTokenModel
}

type AuthTokenModel struct {
	Token     string `json:"access_token"`
	ExpiresAt string `json:"expires_in"`
}

// MiddlewareAuthFunction provides way to implement custom middleware which can do auth
type MiddlewareAuthFunction func(authClient *CitrixDaasClient, r *http.Request)

func getMiddlewareWithClient(authClient *CitrixDaasClient, middlewareAuthFunc MiddlewareAuthFunction) citrixorchestration.MiddlewareFunction {
	return func(r *http.Request) {
		middlewareAuthFunc(authClient, r)
	}
}

func NewCitrixDaasClient(authUrl, hostname, customerId, clientId, clientSecret string, onPremise bool, disableSslVerification bool, userAgent *string, middlewareFunc MiddlewareAuthFunction) (*CitrixDaasClient, *http.Response, error) {
	daasClient := &CitrixDaasClient{}

	/* ------ Setup API Client ------ */
	localCfg := citrixorchestration.NewConfiguration()
	localCfg.Host = hostname
	localCfg.Scheme = "https"
	// When running against on-prem, set disableSslVerification to true when the DDC does not have a valid TLS/SSL certificate
	if onPremise {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: disableSslVerification},
		}
		client := &http.Client{Transport: tr}
		localCfg.HTTPClient = client

		localCfg.Servers = citrixorchestration.ServerConfigurations{
			{
				URL: localCfg.Scheme + "://" + hostname + "/citrix/orchestration/api/techpreview",
			},
		}
	}

	daasClient.ApiClient = citrixorchestration.NewAPIClient(localCfg)
	localCfg.Middleware = getMiddlewareWithClient(daasClient, middlewareFunc)

	/* ------ Setup Authentication Configuration ------ */
	localAuthCfg := &AuthenticationConfiguration{}
	localAuthCfg.AuthUrl = authUrl
	localAuthCfg.ClientId = clientId
	localAuthCfg.ClientSecret = clientSecret
	localAuthCfg.OnPremise = onPremise

	daasClient.AuthConfig = localAuthCfg

	/* ------ Setup Client Configuration ------*/
	req := daasClient.ApiClient.MeAPIsDAAS.MeGetMe(context.Background())
	token, httpResp, err := daasClient.SignIn()
	if err != nil {
		return nil, httpResp, err
	}
	req = req.Authorization(token).CitrixCustomerId(customerId)
	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, err
	}

	localClientCfg := &ClientConfiguration{}
	localClientCfg.CustomerId = customerId
	localClientCfg.SiteId = resp.Customers[0].Sites[0].Id

	if userAgent == nil {
		defaultUserAgent := "citrix-daas-rest-go (https://github.com/citrix/citrix-daas-rest-go)"
		userAgent = &defaultUserAgent
	}
	localClientCfg.UserAgent = *userAgent

	daasClient.ClientConfig = localClientCfg

	if onPremise {
		// add CustomerId and SiteId to base path for on-prem. The Me Api will no longer work.
		localCfg.Servers[0].URL += "/" + localClientCfg.CustomerId + "/" + localClientCfg.SiteId
	}

	return daasClient, httpResp, nil
}

func (c *CitrixDaasClient) WaitForJob(ctx context.Context, jobId string, maxWaitTimeInMinutes int) (*citrixorchestration.JobResponseModel, error) {

	maxWaitTime := time.Now().UTC().Add(time.Minute * time.Duration(maxWaitTimeInMinutes))
	getJobIdRequest := c.ApiClient.JobsAPIsDAAS.JobsGetJob(ctx, jobId)
	var jobResponseModel *citrixorchestration.JobResponseModel
	var err error

	for {
		if time.Now().UTC().After(maxWaitTime) {
			break
		}

		jobResponseModel, _, err = AddRequestData(getJobIdRequest, c).Execute()
		if err != nil {
			return jobResponseModel, err
		}

		if jobResponseModel.Status == citrixorchestration.JOBSTATUS_UNKNOWN || jobResponseModel.Status == citrixorchestration.JOBSTATUS_NOT_STARTED || jobResponseModel.Status == citrixorchestration.JOBSTATUS_IN_PROGRESS {
			time.Sleep(time.Second * time.Duration(30))
			continue
		}

		return jobResponseModel, err
	}

	return jobResponseModel, err
}

func AddRequestData[T any](request T, c *CitrixDaasClient) T {
	// Set the customerId, siteId, and user agent on the request using reflection (request models don't share a struct)
	requestValue := reflect.ValueOf(request)
	requestType := reflect.TypeOf(request)

	method, ok := requestType.MethodByName("CitrixCustomerId")
	if ok {
		requestValue = method.Func.Call([]reflect.Value{requestValue, reflect.ValueOf(c.ClientConfig.CustomerId)})[0]
	}

	method, ok = requestType.MethodByName("CitrixInstanceId")
	if ok {
		requestValue = method.Func.Call([]reflect.Value{requestValue, reflect.ValueOf(c.ClientConfig.SiteId)})[0]
	}

	method, ok = requestType.MethodByName("UserAgent")
	if ok {
		requestValue = method.Func.Call([]reflect.Value{requestValue, reflect.ValueOf(c.ClientConfig.UserAgent)})[0]
	}
	return requestValue.Interface().(T)
}

func RetryOperationWithExponentialBackOff(operation func() (*http.Response, error), baseDelayInSeconds int, maxRetries int) (*http.Response, error) {
	var resp *http.Response
	var err error
	baseDelay := time.Duration(baseDelayInSeconds) * time.Second
	for i := 0; i < maxRetries; i++ {
		resp, err = operation()
		if err != nil {
			return resp, err
		}

		if resp.StatusCode == 200 {
			return resp, err
		}

		if resp.StatusCode == 429 {
			delay := math.Pow(2, float64(i))
			time.Sleep(time.Duration(delay) * baseDelay)
		} else {
			return resp, err
		}

	}

	return resp, err
}
