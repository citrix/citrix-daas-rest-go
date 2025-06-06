/*
Citrix Virtual App & Desktop Catalog Service 147.0.26651.57932

Catalog Service

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickdeploy

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// CatalogApplicationsCMDService CatalogApplicationsCMD service
type CatalogApplicationsCMDService service

type ApiGetCatalogAppsRequest struct {
	ctx                 context.Context
	ApiService          *CatalogApplicationsCMDService
	customerId          string
	siteId              string
	catalogId           string
	citrixTransactionId *string
}

// The Transaction Id.
func (r ApiGetCatalogAppsRequest) CitrixTransactionId(citrixTransactionId string) ApiGetCatalogAppsRequest {
	r.citrixTransactionId = &citrixTransactionId
	return r
}

func (r ApiGetCatalogAppsRequest) Execute() (*CatalogApplicationsModel, *http.Response, error) {
	return r.ApiService.GetCatalogAppsExecute(r)
}

/*
GetCatalogApps Get the list of apps that have been published for the specified catalog

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customerId ID of the customer
	@param siteId The site ID of the customer
	@param catalogId ID of the catalog
	@return ApiGetCatalogAppsRequest
*/
func (a *CatalogApplicationsCMDService) GetCatalogApps(ctx context.Context, customerId string, siteId string, catalogId string) ApiGetCatalogAppsRequest {
	return ApiGetCatalogAppsRequest{
		ApiService: a,
		ctx:        ctx,
		customerId: customerId,
		siteId:     siteId,
		catalogId:  catalogId,
	}
}

// Execute executes the request
//
//	@return CatalogApplicationsModel
func (a *CatalogApplicationsCMDService) GetCatalogAppsExecute(r ApiGetCatalogAppsRequest) (*CatalogApplicationsModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CatalogApplicationsModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogApplicationsCMDService.GetCatalogApps")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{customerId}/{siteId}/catalogs/{catalogId}/apps"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", url.PathEscape(parameterValueToString(r.customerId, "customerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"siteId"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"catalogId"+"}", url.PathEscape(parameterValueToString(r.catalogId, "catalogId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.citrixTransactionId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Citrix-TransactionId", r.citrixTransactionId, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CWSAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPublishApplicationsRequest struct {
	ctx                 context.Context
	ApiService          *CatalogApplicationsCMDService
	customerId          string
	siteId              string
	catalogId           string
	citrixTransactionId *string
	body                *AddCatalogApplicationsModel
}

// The Transaction Id.
func (r ApiPublishApplicationsRequest) CitrixTransactionId(citrixTransactionId string) ApiPublishApplicationsRequest {
	r.citrixTransactionId = &citrixTransactionId
	return r
}

// List of applications to add
func (r ApiPublishApplicationsRequest) Body(body AddCatalogApplicationsModel) ApiPublishApplicationsRequest {
	r.body = &body
	return r
}

func (r ApiPublishApplicationsRequest) Execute() (*CatalogApplicationsModel, *http.Response, error) {
	return r.ApiService.PublishApplicationsExecute(r)
}

/*
PublishApplications Publish the specified apps to the catalog

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customerId ID of the customer
	@param siteId The site ID of the customer
	@param catalogId ID of the catalog
	@return ApiPublishApplicationsRequest
*/
func (a *CatalogApplicationsCMDService) PublishApplications(ctx context.Context, customerId string, siteId string, catalogId string) ApiPublishApplicationsRequest {
	return ApiPublishApplicationsRequest{
		ApiService: a,
		ctx:        ctx,
		customerId: customerId,
		siteId:     siteId,
		catalogId:  catalogId,
	}
}

// Execute executes the request
//
//	@return CatalogApplicationsModel
func (a *CatalogApplicationsCMDService) PublishApplicationsExecute(r ApiPublishApplicationsRequest) (*CatalogApplicationsModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CatalogApplicationsModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogApplicationsCMDService.PublishApplications")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{customerId}/{siteId}/catalogs/{catalogId}/apps"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", url.PathEscape(parameterValueToString(r.customerId, "customerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"siteId"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"catalogId"+"}", url.PathEscape(parameterValueToString(r.catalogId, "catalogId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json", "application/json", "text/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.citrixTransactionId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Citrix-TransactionId", r.citrixTransactionId, "")
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CWSAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUnpublishApplicationRequest struct {
	ctx                 context.Context
	ApiService          *CatalogApplicationsCMDService
	customerId          string
	siteId              string
	catalogId           string
	appId               string
	citrixTransactionId *string
}

// The Transaction Id.
func (r ApiUnpublishApplicationRequest) CitrixTransactionId(citrixTransactionId string) ApiUnpublishApplicationRequest {
	r.citrixTransactionId = &citrixTransactionId
	return r
}

func (r ApiUnpublishApplicationRequest) Execute() (*http.Response, error) {
	return r.ApiService.UnpublishApplicationExecute(r)
}

/*
UnpublishApplication Remove an application from the published list

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customerId ID of the customer
	@param siteId
	@param catalogId ID of the catalog
	@param appId Identifier of the application
	@return ApiUnpublishApplicationRequest
*/
func (a *CatalogApplicationsCMDService) UnpublishApplication(ctx context.Context, customerId string, siteId string, catalogId string, appId string) ApiUnpublishApplicationRequest {
	return ApiUnpublishApplicationRequest{
		ApiService: a,
		ctx:        ctx,
		customerId: customerId,
		siteId:     siteId,
		catalogId:  catalogId,
		appId:      appId,
	}
}

// Execute executes the request
func (a *CatalogApplicationsCMDService) UnpublishApplicationExecute(r ApiUnpublishApplicationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogApplicationsCMDService.UnpublishApplication")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{customerId}/{siteId}/catalogs/{catalogId}/apps/{appId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", url.PathEscape(parameterValueToString(r.customerId, "customerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"siteId"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"catalogId"+"}", url.PathEscape(parameterValueToString(r.catalogId, "catalogId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterValueToString(r.appId, "appId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.citrixTransactionId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Citrix-TransactionId", r.citrixTransactionId, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CWSAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUnpublishApplicationsRequest struct {
	ctx                 context.Context
	ApiService          *CatalogApplicationsCMDService
	customerId          string
	siteId              string
	catalogId           string
	appIds              *[]string
	citrixTransactionId *string
}

// List of application ids to remove
func (r ApiUnpublishApplicationsRequest) AppIds(appIds []string) ApiUnpublishApplicationsRequest {
	r.appIds = &appIds
	return r
}

// The Transaction Id.
func (r ApiUnpublishApplicationsRequest) CitrixTransactionId(citrixTransactionId string) ApiUnpublishApplicationsRequest {
	r.citrixTransactionId = &citrixTransactionId
	return r
}

func (r ApiUnpublishApplicationsRequest) Execute() (*http.Response, error) {
	return r.ApiService.UnpublishApplicationsExecute(r)
}

/*
UnpublishApplications Remove a list of applications from a catalog

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customerId ID of the customer
	@param siteId The site ID of the customer
	@param catalogId ID of the catalog
	@return ApiUnpublishApplicationsRequest
*/
func (a *CatalogApplicationsCMDService) UnpublishApplications(ctx context.Context, customerId string, siteId string, catalogId string) ApiUnpublishApplicationsRequest {
	return ApiUnpublishApplicationsRequest{
		ApiService: a,
		ctx:        ctx,
		customerId: customerId,
		siteId:     siteId,
		catalogId:  catalogId,
	}
}

// Execute executes the request
func (a *CatalogApplicationsCMDService) UnpublishApplicationsExecute(r ApiUnpublishApplicationsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogApplicationsCMDService.UnpublishApplications")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{customerId}/{siteId}/catalogs/{catalogId}/apps"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", url.PathEscape(parameterValueToString(r.customerId, "customerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"siteId"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"catalogId"+"}", url.PathEscape(parameterValueToString(r.catalogId, "catalogId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.appIds != nil {
		t := *r.appIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "appIds", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "appIds", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.citrixTransactionId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Citrix-TransactionId", r.citrixTransactionId, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CWSAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateApplicationRequest struct {
	ctx                 context.Context
	ApiService          *CatalogApplicationsCMDService
	customerId          string
	siteId              string
	catalogId           string
	appId               string
	citrixTransactionId *string
	body                *UpdateApplicationConfigurationModel
}

// The Transaction Id.
func (r ApiUpdateApplicationRequest) CitrixTransactionId(citrixTransactionId string) ApiUpdateApplicationRequest {
	r.citrixTransactionId = &citrixTransactionId
	return r
}

// New configuration of the application
func (r ApiUpdateApplicationRequest) Body(body UpdateApplicationConfigurationModel) ApiUpdateApplicationRequest {
	r.body = &body
	return r
}

func (r ApiUpdateApplicationRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateApplicationExecute(r)
}

/*
UpdateApplication Update the configuration of a published app

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customerId ID of the customer
	@param siteId The site ID of the customer
	@param catalogId ID of the catalog
	@param appId Identifier of the application
	@return ApiUpdateApplicationRequest
*/
func (a *CatalogApplicationsCMDService) UpdateApplication(ctx context.Context, customerId string, siteId string, catalogId string, appId string) ApiUpdateApplicationRequest {
	return ApiUpdateApplicationRequest{
		ApiService: a,
		ctx:        ctx,
		customerId: customerId,
		siteId:     siteId,
		catalogId:  catalogId,
		appId:      appId,
	}
}

// Execute executes the request
func (a *CatalogApplicationsCMDService) UpdateApplicationExecute(r ApiUpdateApplicationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogApplicationsCMDService.UpdateApplication")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{customerId}/{siteId}/catalogs/{catalogId}/apps/{appId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", url.PathEscape(parameterValueToString(r.customerId, "customerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"siteId"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"catalogId"+"}", url.PathEscape(parameterValueToString(r.catalogId, "catalogId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterValueToString(r.appId, "appId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json", "application/json", "text/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.citrixTransactionId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Citrix-TransactionId", r.citrixTransactionId, "")
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CWSAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
