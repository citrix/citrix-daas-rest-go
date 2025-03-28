// Copyright © 2024. Citrix Systems, Inc.

package citrixclient

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"reflect"
	"slices"
	"strings"
	"time"

	ccadmins "github.com/citrix/citrix-daas-rest-go/ccadmins"
	resourcelocations "github.com/citrix/citrix-daas-rest-go/ccresourcelocations"
	"github.com/citrix/citrix-daas-rest-go/citrixcws"
	"github.com/citrix/citrix-daas-rest-go/citrixorchestration"
	citrixquickcreate "github.com/citrix/citrix-daas-rest-go/citrixquickcreate"
	storefrontapis "github.com/citrix/citrix-daas-rest-go/citrixstorefront/apis"
	citrixwemservice "github.com/citrix/citrix-daas-rest-go/devicemanagement"
	globalappconfiguration "github.com/citrix/citrix-daas-rest-go/globalappconfiguration"
)

type CitrixDaasClient struct {
	ApiClient           *citrixorchestration.APIClient
	AuthConfig          *AuthenticationConfiguration
	ClientConfig        *ClientConfiguration
	WemOnPremAuthConfig *WemOnPremAuthentication
	AuthToken           *AuthTokenModel
	StorefrontClient    *storefrontapis.APIClient
	QuickCreateClient   *citrixquickcreate.APIClient
	WemClient           *citrixwemservice.APIClient
	// Citrix Cloud Service Clients
	CCAdminsClient          *ccadmins.APIClient
	CwsClient               *citrixcws.APIClient
	GacClient               *globalappconfiguration.APIClient
	ResourceLocationsClient *resourcelocations.APIClient
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

func getMiddlewareWithGacClient(authClient *CitrixDaasClient, middlewareAuthFunc MiddlewareAuthFunction) globalappconfiguration.MiddlewareFunction {
	return func(r *http.Request) {
		middlewareAuthFunc(authClient, r)
	}
}

func getMiddlewareWithResourceLocationsClient(authClient *CitrixDaasClient, middlewareAuthFunc MiddlewareAuthFunction) resourcelocations.MiddlewareFunction {
	return func(r *http.Request) {
		middlewareAuthFunc(authClient, r)
	}
}

func getMiddlewareWithCCAdminsClient(authClient *CitrixDaasClient, middlewareAuthFunc MiddlewareAuthFunction) ccadmins.MiddlewareFunction {
	return func(r *http.Request) {
		middlewareAuthFunc(authClient, r)
	}
}

func getMiddlewareWithQuickcreateClient(authClient *CitrixDaasClient, middlewareAuthFunc MiddlewareAuthFunction) citrixquickcreate.MiddlewareFunction {
	return func(r *http.Request) {
		middlewareAuthFunc(authClient, r)
	}
}

func getMiddlewareWithWemClient(authClient *CitrixDaasClient, middlewareAuthFunc MiddlewareAuthFunction) citrixwemservice.MiddlewareFunction {
	return func(r *http.Request) {
		middlewareAuthFunc(authClient, r)
	}
}

func getMiddlewareWithCwsClient(authClient *CitrixDaasClient, middlewareAuthFunc MiddlewareAuthFunction) citrixcws.MiddlewareFunction {
	return func(r *http.Request) {
		middlewareAuthFunc(authClient, r)
	}
}

/* ------ Setup WEM Client ------ */
func (daasClient *CitrixDaasClient) InitializeWemClient(ctx context.Context, wemHostName string, middlewareFunc MiddlewareAuthFunction, onPremises, disableSslVerification bool) {
	wemConfig := citrixwemservice.NewConfiguration()
	wemConfig.Scheme = "https"
	if onPremises {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: disableSslVerification},
		}
		client := &http.Client{Transport: tr}
		wemConfig.HTTPClient = client
	}
	wemConfig.Servers = citrixwemservice.ServerConfigurations{
		{
			URL: wemConfig.Scheme + "://" + wemHostName,
		},
	}

	wemConfig.Middleware = getMiddlewareWithWemClient(daasClient, middlewareFunc)
	daasClient.WemClient = citrixwemservice.NewAPIClient(wemConfig)
}

func (daasClient *CitrixDaasClient) InitializeStoreFrontClient(ctx context.Context, computerName, adUserName, adUserPass string, disableSslVerification bool) {
	daasClient.StorefrontClient = storefrontapis.NewAPIClient()
	daasClient.StorefrontClient.SetComputerName(computerName)
	daasClient.StorefrontClient.SetAdUserName(adUserName)
	daasClient.StorefrontClient.SetAdPassword(adUserPass)
	daasClient.StorefrontClient.SetDisableSSL(disableSslVerification)
}

func (daasClient *CitrixDaasClient) InitializeQuickCreateClient(ctx context.Context, quickCreateHostName string, middlewareFunc MiddlewareAuthFunction) {
	/* ------ Setup QuickCreate Client ------ */
	localQuickCreateCfg := citrixquickcreate.NewConfiguration()
	localQuickCreateCfg.Scheme = "https"

	localQuickCreateCfg.Servers = citrixquickcreate.ServerConfigurations{
		{
			URL: localQuickCreateCfg.Scheme + "://" + quickCreateHostName,
		},
	}

	// Disable ssl check
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	localQuickCreateCfg.HTTPClient = &http.Client{Transport: tr}

	localQuickCreateCfg.Middleware = getMiddlewareWithQuickcreateClient(daasClient, middlewareFunc)
	daasClient.QuickCreateClient = citrixquickcreate.NewAPIClient(localQuickCreateCfg)
}

func (daasClient *CitrixDaasClient) InitializeCwsClient(ctx context.Context, cwsHostname string, middlewareFunc MiddlewareAuthFunction) {
	/* ------ Setup Citrix Cloud Cws Service Client ------ */
	localCwsCfg := citrixcws.NewConfiguration()
	localCwsCfg.Scheme = "https"

	localCwsCfg.Servers = citrixcws.ServerConfigurations{
		{
			URL: localCwsCfg.Scheme + "://" + cwsHostname,
		},
	}

	localCwsCfg.Middleware = getMiddlewareWithCwsClient(daasClient, middlewareFunc)
	daasClient.CwsClient = citrixcws.NewAPIClient(localCwsCfg)
}

/* ------ Setup API Client ------ */
func (daasClient *CitrixDaasClient) SetupApiClient(hostname string, middlewareFunc MiddlewareAuthFunction, onPremises, disableSslVerification, apiGateway bool) {
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
				URL: localCfg.Scheme + "://" + hostname + "/citrix/orchestration/api",
			},
		}
	} else if !apiGateway {
		localCfg.Servers = citrixorchestration.ServerConfigurations{
			{
				URL: localCfg.Scheme + "://" + hostname + "/citrix/orchestration/api",
			},
		}
	}
	localCfg.Middleware = getMiddlewareWithClient(daasClient, middlewareFunc)
	daasClient.ApiClient = citrixorchestration.NewAPIClient(localCfg)
}

/* ------ Setup Resource Locations Client ------ */
func (daasClient *CitrixDaasClient) SetupResourceLocationsClient(ccUrl string, middlewareFunc MiddlewareAuthFunction) {
	localResourceLocationsCfg := resourcelocations.NewConfiguration()
	localResourceLocationsCfg.Scheme = "https"
	localResourceLocationsCfg.Servers = resourcelocations.ServerConfigurations{
		{
			URL: localResourceLocationsCfg.Scheme + "://" + ccUrl + "/resourcelocations",
		},
	}
	localResourceLocationsCfg.Middleware = getMiddlewareWithResourceLocationsClient(daasClient, middlewareFunc)
	daasClient.ResourceLocationsClient = resourcelocations.NewAPIClient(localResourceLocationsCfg)
}

/* ------ Setup CC Admin Client ------ */
func (daasClient *CitrixDaasClient) SetupCCAdminClient(ccUrl string, middlewareFunc MiddlewareAuthFunction) {
	localCCAdminCfg := ccadmins.NewConfiguration()
	localCCAdminCfg.Scheme = "https"
	localCCAdminCfg.Servers = ccadmins.ServerConfigurations{
		{
			URL: localCCAdminCfg.Scheme + "://" + ccUrl + "/administrators",
		},
	}
	localCCAdminCfg.Middleware = getMiddlewareWithCCAdminsClient(daasClient, middlewareFunc)
	daasClient.CCAdminsClient = ccadmins.NewAPIClient(localCCAdminCfg)
}

/* ------ Setup Authentication Configuration ------ */
func (daasClient *CitrixDaasClient) SetupAuthConfig(authUrl string, clientId string, clientSecret string, onPremises bool, apiGateway bool, isGov bool, environment string) {
	localAuthCfg := &AuthenticationConfiguration{
		AuthUrl:      authUrl,
		ClientId:     clientId,
		ClientSecret: clientSecret,
		OnPremises:   onPremises,
		ApiGateway:   apiGateway,
		IsGov:        isGov,
		Environment:  environment,
	}
	daasClient.AuthConfig = localAuthCfg
}

/* ------ Setup Wem On-Prem Authentication Configuration ------ */
func (daasClient *CitrixDaasClient) SetupWemOnPremAuthConfig(authUrl, adminUserName, adminPassword string) {
	localAuthCfg := &WemOnPremAuthentication{
		AuthUrl:       authUrl,
		AdminUserName: adminUserName,
		AdminPassword: adminPassword,
	}
	daasClient.WemOnPremAuthConfig = localAuthCfg
}

/* ------ Setup GAC Client ------ */
func (daasClient *CitrixDaasClient) SetupGacClient(hostname string, middlewareFunc MiddlewareAuthFunction) {
	gacUrlTranslator := map[string]bool{
		"api.dev.cloud.com":    false,
		"api.cloudburrito.com": false,
	}

	localGacCfg := globalappconfiguration.NewConfiguration()
	localGacCfg.Scheme = "https"

	_, exists := gacUrlTranslator[hostname]
	if exists {
		localGacCfg.Servers = globalappconfiguration.ServerConfigurations{
			{
				URL: localGacCfg.Scheme + "://wsaca.cloudburrito.com",
			},
		}

	}
	localGacCfg.Middleware = getMiddlewareWithGacClient(daasClient, middlewareFunc)
	daasClient.GacClient = globalappconfiguration.NewAPIClient(localGacCfg)
}

func (daasClient *CitrixDaasClient) SetupWemOnPremClientContext(ctx context.Context, middlewareFunc MiddlewareAuthFunction, wemAuthUrl, wemHostName, wemAdminUserName, wemAdminPassword string, disableSslVerification bool) {
	daasClient.InitializeWemClient(ctx, wemHostName, middlewareFunc, true, disableSslVerification)
	daasClient.SetupWemOnPremAuthConfig(wemAuthUrl, wemAdminUserName, wemAdminPassword)
}

func (daasClient *CitrixDaasClient) SetupCitrixClientsContext(ctx context.Context, authUrl string, ccUrl string, hostname string, customerId string, clientId string, clientSecret string, onPremises bool, apiGateway bool, isGov bool, disableSslVerification bool, userAgent *string, environment string, middlewareFunc MiddlewareAuthFunction, middlewareFuncWithCustomerIdHeader MiddlewareAuthFunction) (string, *http.Response, error) {

	daasClient.SetupApiClient(hostname, middlewareFunc, onPremises, disableSslVerification, apiGateway)
	daasClient.SetupAuthConfig(authUrl, clientId, clientSecret, onPremises, apiGateway, isGov, environment)

	// Setup Client Configuration
	daasClient.ClientConfig = &ClientConfiguration{CustomerId: customerId}

	// Authenticate and get the token
	token, httpResp, err := daasClient.SignIn()
	return token, httpResp, err
}

func (daasClient *CitrixDaasClient) InitializeCitrixCloudClients(ctx context.Context, ccUrl, hostname string, middlewareFunc MiddlewareAuthFunction, middlewareFuncWithCustomerIdHeader MiddlewareAuthFunction) {
	daasClient.SetupResourceLocationsClient(ccUrl, middlewareFunc)
	daasClient.SetupCCAdminClient(ccUrl, middlewareFuncWithCustomerIdHeader)
	daasClient.SetupGacClient(hostname, middlewareFunc)
}

func (daasClient *CitrixDaasClient) InitializeCitrixDaasClient(ctx context.Context, customerId, token string, onPremises bool, apiGateway bool, disableSslVerification bool, userAgent *string) (*http.Response, error) {
	req := daasClient.ApiClient.MeAPIsDAAS.MeGetMe(ctx)
	req = req.Authorization(token).CitrixCustomerId(customerId)
	resp, httpResp, err := req.Execute()
	if err != nil {
		return httpResp, err
	}

	if resp == nil || len(resp.Customers) == 0 || len(resp.Customers[0].Sites) == 0 {
		return httpResp, fmt.Errorf("customer does not exist or does not have a valid site")
	}

	daasClient.ClientConfig.SiteId = resp.Customers[0].Sites[0].Id
	daasClient.ClientConfig.IsCspCustomer = resp.GetIsCspCustomer()
	daasClient.ClientConfig.Accept = "application/json"
	if userAgent == nil {
		defaultUserAgent := "citrix-daas-rest-go (https://github.com/citrix/citrix-daas-rest-go)"
		userAgent = &defaultUserAgent
	}
	daasClient.ClientConfig.UserAgent = *userAgent

	if onPremises || !apiGateway {
		// add CustomerId and SiteId to base path for on-prem. The following APIs will no longer work.
		// HealthCheck, Me, and Ping
		daasClient.ApiClient.GetConfig().Servers[0].URL += "/" + daasClient.ClientConfig.CustomerId
	}

	siteRequest := daasClient.ApiClient.SitesAPIsDAAS.SitesGetSite(ctx, daasClient.ClientConfig.SiteId)
	token, httpResp, err = daasClient.SignIn()
	if err != nil {
		return httpResp, err
	}

	siteRequest = siteRequest.Authorization(token).CitrixCustomerId(customerId)
	siteResp, httpResp, err := siteRequest.Execute()
	if err != nil {
		return httpResp, err
	}

	daasClient.ClientConfig.ProductVersion = siteResp.GetProductVersion()
	daasClient.ClientConfig.OrchestrationApiVersion = siteResp.GetOrchestationApiVersion()

	if onPremises || !apiGateway {
		// add CustomerId and SiteId to base path for on-prem. The Sites API will no longer work.
		daasClient.ApiClient.GetConfig().Servers[0].URL += "/" + daasClient.ClientConfig.SiteId
	}
	return httpResp, nil
}

func (c *CitrixDaasClient) WaitForJob(ctx context.Context, jobId string, maxWaitTimeInMinutes int32) (*citrixorchestration.JobResponseModel, error) {
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
			// Cancel the job if it timeouts
			cancelJobReq := c.ApiClient.JobsAPIsDAAS.JobsCancelJob(ctx, jobId)
			_, _, err := AddRequestData(cancelJobReq, c).Execute()
			jobResponseModel := citrixorchestration.JobResponseModel{}
			if err != nil {
				err = fmt.Errorf("job %s timed out after %d minutes, but the job cannot be cancelled", jobId, maxWaitTimeInMinutes)
				jobResponseModel.SetErrorString(err.Error())
				jobResponseModel.SetId(jobId)
				return &jobResponseModel, err
			}
			err = fmt.Errorf("job %s timed out after %d minutes", jobId, maxWaitTimeInMinutes)
			jobResponseModel.SetErrorString(err.Error())
			jobResponseModel.SetId(jobId)
			return &jobResponseModel, err
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

	method, ok = requestType.MethodByName("Accept")
	if ok {
		requestValue = method.Func.Call([]reflect.Value{requestValue, reflect.ValueOf(c.ClientConfig.Accept)})[0]
	}

	method, ok = requestType.MethodByName("Authorization")
	if ok {
		requestValue = method.Func.Call([]reflect.Value{requestValue, reflect.ValueOf(c.ClientConfig.Authorization)})[0]
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
			var responseBody ResponseBodyType
			var httpResp *http.Response
			var err error
			if len(values) > 2 {
				responseBody, _ = values[0].Interface().(ResponseBodyType)
				httpResp, _ = values[1].Interface().(*http.Response)
				err, _ = values[2].Interface().(error)
			} else {
				// This allows us to use ExecuteWithRetry when the request does not have a response body (ex: 204 No Content)
				httpResp, _ = values[0].Interface().(*http.Response)
				err, _ = values[1].Interface().(error)
			}

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
	successCount, txnId, _, err := PerformBatchOperationAndReturnSubJobResponses(ctx, client, batchRequestModel)
	return successCount, txnId, err
}

func PerformBatchOperationAndReturnSubJobResponses(ctx context.Context, client *CitrixDaasClient, batchRequestModel citrixorchestration.BatchRequestModel) (int, string, []citrixorchestration.BatchResponseItemModel, error) {

	successCount := 0
	subJobs := []citrixorchestration.BatchResponseItemModel{}

	batchRequest := client.ApiClient.BatchAPIsDAAS.BatchDoBatchRequest(ctx).BatchRequestModel(batchRequestModel).Async(true)
	_, httpResp, err := AddRequestData(batchRequest, client).Execute()
	txnId := GetTransactionIdFromHttpResponse(httpResp)
	if err != nil {
		return successCount, txnId, subJobs, err
	}
	jobId := GetJobIdFromHttpResponse(*httpResp)
	jobResponseModel, err := client.WaitForJob(ctx, jobId, 30)
	if err != nil {
		return successCount, txnId, subJobs, err
	}

	jobStatus := jobResponseModel.GetStatus()

	if jobStatus != citrixorchestration.JOBSTATUS_COMPLETE {
		if jobStatus == citrixorchestration.JOBSTATUS_FAILED {
			return successCount, txnId, subJobs, fmt.Errorf(jobResponseModel.GetErrorString())
		}

		return successCount, txnId, subJobs, fmt.Errorf("an unexpected error occurred")
	}

	// Job is completed successfully. Get results
	ss := client.ApiClient.JobsAPIsDAAS.JobsGetJobResults(ctx, jobId)
	res, _, err := AddRequestData(ss, client).Execute()
	if err != nil {
		// Job failed. Return nil and error.
		return successCount, txnId, subJobs, err
	}

	var batchJobResponse map[string][]citrixorchestration.BatchResponseItemModel
	_ = json.Unmarshal([]byte(res), &batchJobResponse)

	subJobs = batchJobResponse["Items"]
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

	return successCount, txnId, subJobs, err
}
