// Copyright © 2023. Citrix Systems, Inc.

package citrixclient

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"reflect"
	"slices"
	"strings"
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

func NewCitrixDaasClient(ctx context.Context, authUrl, hostname, customerId, clientId, clientSecret string, onPremises bool, apiGateway bool, isGov bool, disableSslVerification bool, userAgent *string, middlewareFunc MiddlewareAuthFunction) (*CitrixDaasClient, *http.Response, error) {
	daasClient := &CitrixDaasClient{}

	/* ------ Setup API Client ------ */
	localCfg := citrixorchestration.NewConfiguration()
	localCfg.Host = hostname
	localCfg.Scheme = "https"
	// When running against on-premises, set disableSslVerification to true when the DDC does not have a valid TLS/SSL certificate
	if onPremises {
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
	} else if !apiGateway {
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
	localAuthCfg.OnPremises = onPremises
	localAuthCfg.ApiGateway = apiGateway
	localAuthCfg.IsGov = isGov

	daasClient.AuthConfig = localAuthCfg

	/* ------ Setup Client Configuration ------*/
	req := daasClient.ApiClient.MeAPIsDAAS.MeGetMe(ctx)
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
	if resp == nil || len(resp.Customers) == 0 || len(resp.Customers[0].Sites) == 0 {
		return nil, httpResp, fmt.Errorf("customer does not exist or does not have a valid site")
	}
	localClientCfg.SiteId = resp.Customers[0].Sites[0].Id

	if userAgent == nil {
		defaultUserAgent := "citrix-daas-rest-go (https://github.com/citrix/citrix-daas-rest-go)"
		userAgent = &defaultUserAgent
	}
	localClientCfg.UserAgent = *userAgent

	daasClient.ClientConfig = localClientCfg

	if onPremises || !apiGateway {
		// add CustomerId and SiteId to base path for on-prem. The Me Api will no longer work.
		localCfg.Servers[0].URL += "/" + localClientCfg.CustomerId + "/" + localClientCfg.SiteId
	}

	return daasClient, httpResp, nil
}

func (c *CitrixDaasClient) WaitForJob(ctx context.Context, jobId string, maxWaitTimeInMinutes int) (*citrixorchestration.JobResponseModel, error) {
	// default polling to every 10 seconds
	pollInterval := 10
	if maxWaitTimeInMinutes >= 30 {
		// if we're expecting a long job we can poll less frequently
		pollInterval = 30
	}
	startTime := time.Now()
	getJobIdRequest := c.ApiClient.JobsAPIsDAAS.JobsGetJob(ctx, jobId)
	var jobResponseModel *citrixorchestration.JobResponseModel
	var err error

	for {
		if time.Since(startTime) > time.Minute*time.Duration(maxWaitTimeInMinutes) {
			break
		}

		jobResponseModel, _, err = AddRequestData(getJobIdRequest, c).Execute()
		if err != nil {
			return jobResponseModel, err
		}

		jobStatus := jobResponseModel.GetStatus()

		if jobStatus == citrixorchestration.JOBSTATUS_UNKNOWN || jobStatus == citrixorchestration.JOBSTATUS_NOT_STARTED || jobStatus == citrixorchestration.JOBSTATUS_IN_PROGRESS {
			if pollInterval < 30 && time.Since(startTime) > time.Second*30 {
				// increase poll interval after the first 30 seconds of processing
				pollInterval = 30
			}
			time.Sleep(time.Second * time.Duration(pollInterval))
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

func ExecuteWithRetry[ResponseBodyType any](request any, c *CitrixDaasClient) (ResponseBodyType, *http.Response, error) {
	request = AddRequestData(request, c)
	requestValue := reflect.ValueOf(request)
	requestType := reflect.TypeOf(request)
	var result ResponseBodyType
	method, ok := requestType.MethodByName("Execute")
	if ok {
		operation := func() (ResponseBodyType, *http.Response, error) {
			values := method.Func.Call([]reflect.Value{requestValue})
			responseBody, _ := values[0].Interface().(ResponseBodyType)
			httpResp, _ := values[1].Interface().(*http.Response)
			err, _ := values[2].Interface().(error)

			return responseBody, httpResp, err
		}
		return RetryOperationWithExponentialBackOffDefault(operation)
	}

	return result, nil, fmt.Errorf("an unexpected error occurred")
}

func RetryOperationWithExponentialBackOffDefault[T any](operation func() (T, *http.Response, error)) (T, *http.Response, error) {
	return RetryOperationWithExponentialBackOff(operation, 15, 3)
}

func RetryOperationWithExponentialBackOff[T any](operation func() (T, *http.Response, error), baseDelayInSeconds int, maxRetries int) (T, *http.Response, error) {

	var resp *http.Response
	var err error
	var responseBody T
	baseDelay := time.Duration(baseDelayInSeconds) * time.Second
	for i := 0; i < maxRetries; i++ {
		responseBody, resp, err = operation()

		if resp == nil {
			// error in http call. caller will handle err
			return responseBody, resp, err
		}

		if resp.StatusCode == 200 {
			return responseBody, resp, err
		}

		if resp.StatusCode == 429 || resp.StatusCode >= 500 {
			defer resp.Body.Close()
			delay := math.Pow(2, float64(i))
			time.Sleep(time.Duration(delay) * baseDelay)
		} else {
			return responseBody, resp, err
		}
	}

	return responseBody, resp, err
}

func GetJobIdFromHttpResponse(httpResponse http.Response) string {
	locationHeader := httpResponse.Header.Get("Location")
	locationHeaderParts := strings.Split(locationHeader, "/")
	jobId := locationHeaderParts[len(locationHeaderParts)-1]

	return jobId
}

func GetTransactionIdFromHttpResponse(httpResponse *http.Response) string {
	if httpResponse == nil {
		return "failed before request was sent"
	}
	return httpResponse.Header.Get("Citrix-TransactionId")
}

func (c *CitrixDaasClient) GetBatchRequestItemRelativeUrl(relativeUrl string) string {
	if !c.AuthConfig.OnPremises && c.AuthConfig.ApiGateway {
		// In case of Cloud customer going through API Gateway, use path as is
		return relativeUrl
	}

	// In case of on-premises customer or Cloud customer bypassing API Gateway, append customerId and siteId to the path
	path := fmt.Sprintf("%s/%s/%s", c.ClientConfig.CustomerId, c.ClientConfig.SiteId, relativeUrl)
	return path
}

func PerformBatchOperation(ctx context.Context, client *CitrixDaasClient, batchRequestModel citrixorchestration.BatchRequestModel) (int, string, error) {

	successCount := 0

	batchRequest := client.ApiClient.BatchAPIsDAAS.BatchDoBatchRequest(ctx).BatchRequestModel(batchRequestModel).Async(true)
	_, httpResp, err := AddRequestData(batchRequest, client).Execute()
	txnId := GetTransactionIdFromHttpResponse(httpResp)
	if err != nil {
		return successCount, txnId, err
	}
	jobId := GetJobIdFromHttpResponse(*httpResp)
	jobResponseModel, err := client.WaitForJob(ctx, jobId, 30)
	if err != nil {
		return successCount, txnId, err
	}

	jobStatus := jobResponseModel.GetStatus()

	if jobStatus != citrixorchestration.JOBSTATUS_COMPLETE {
		if jobStatus == citrixorchestration.JOBSTATUS_FAILED {
			return successCount, txnId, fmt.Errorf(jobResponseModel.GetErrorString())
		}

		return successCount, txnId, fmt.Errorf("an unexpected error occurred")
	}

	getJobResultsRequest := client.ApiClient.JobsAPIsDAAS.JobsGetJobResults(ctx, jobId)
	jobResultsResponseContents, _, _ := ExecuteWithRetry[*os.File](getJobResultsRequest, client)
	if jobResultsResponseContents == nil {
		return successCount, txnId, fmt.Errorf("an unexpected error occurred")
	}

	data, err := io.ReadAll(jobResultsResponseContents)
	if err != nil {
		return successCount, txnId, err
	}

	var batchJobResponse map[string][]citrixorchestration.BatchResponseItemModel
	json.Unmarshal(data, &batchJobResponse)
	subJobs := batchJobResponse["Items"]
	for i := 0; i < len(subJobs); i++ {
		subJob := subJobs[i]
		subJobCode := subJob.GetCode()
		if subJobCode == 202 {
			locationHeader := slices.IndexFunc(subJob.Headers, func(pair citrixorchestration.NameValueStringPairModel) bool { return pair.GetName() == "Location" })
			locationValues := strings.Split(subJob.Headers[locationHeader].GetValue(), "/")
			jobResponseModel, err := client.WaitForJob(ctx, locationValues[len(locationValues)-1], 60)

			if err != nil || jobResponseModel == nil || jobResponseModel.GetStatus() != citrixorchestration.JOBSTATUS_COMPLETE {
				continue
			}

			successCount++
		} else if subJobCode == 204 || subJobCode == 200 {
			successCount++
		}
	}

	return successCount, txnId, err
}
