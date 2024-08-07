/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-onprem-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ImageVersionsAPIsDAASService ImageVersionsAPIsDAAS service
type ImageVersionsAPIsDAASService service

type ApiImageVersionsUpdateImageVersionResourcePoolsRequest struct {
	ctx context.Context
	ApiService *ImageVersionsAPIsDAASService
	citrixCustomerId *string
	citrixInstanceId *string
	id string
	updateImageVersionResourcePoolsRequestModel *UpdateImageVersionResourcePoolsRequestModel
	userAgent *string
	authorization *string
	citrixTransactionId *string
	accept *string
	citrixLocale *string
	async *bool
}

// Citrix Customer ID. Default is &#39;CitrixOnPremises&#39;
func (r ApiImageVersionsUpdateImageVersionResourcePoolsRequest) CitrixCustomerId(citrixCustomerId string) ApiImageVersionsUpdateImageVersionResourcePoolsRequest {
	r.citrixCustomerId = &citrixCustomerId
	return r
}

// Citrix Instance (Site) ID.
func (r ApiImageVersionsUpdateImageVersionResourcePoolsRequest) CitrixInstanceId(citrixInstanceId string) ApiImageVersionsUpdateImageVersionResourcePoolsRequest {
	r.citrixInstanceId = &citrixInstanceId
	return r
}

func (r ApiImageVersionsUpdateImageVersionResourcePoolsRequest) UpdateImageVersionResourcePoolsRequestModel(updateImageVersionResourcePoolsRequestModel UpdateImageVersionResourcePoolsRequestModel) ApiImageVersionsUpdateImageVersionResourcePoolsRequest {
	r.updateImageVersionResourcePoolsRequestModel = &updateImageVersionResourcePoolsRequestModel
	return r
}

// User Agent type of the request.
func (r ApiImageVersionsUpdateImageVersionResourcePoolsRequest) UserAgent(userAgent string) ApiImageVersionsUpdateImageVersionResourcePoolsRequest {
	r.userAgent = &userAgent
	return r
}

// Citrix authorization header: CWSAuth Bearer&#x3D;{token}
func (r ApiImageVersionsUpdateImageVersionResourcePoolsRequest) Authorization(authorization string) ApiImageVersionsUpdateImageVersionResourcePoolsRequest {
	r.authorization = &authorization
	return r
}

// Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned.
func (r ApiImageVersionsUpdateImageVersionResourcePoolsRequest) CitrixTransactionId(citrixTransactionId string) ApiImageVersionsUpdateImageVersionResourcePoolsRequest {
	r.citrixTransactionId = &citrixTransactionId
	return r
}

// Must accept application/json.
func (r ApiImageVersionsUpdateImageVersionResourcePoolsRequest) Accept(accept string) ApiImageVersionsUpdateImageVersionResourcePoolsRequest {
	r.accept = &accept
	return r
}

// Locale of the request.
func (r ApiImageVersionsUpdateImageVersionResourcePoolsRequest) CitrixLocale(citrixLocale string) ApiImageVersionsUpdateImageVersionResourcePoolsRequest {
	r.citrixLocale = &citrixLocale
	return r
}

func (r ApiImageVersionsUpdateImageVersionResourcePoolsRequest) Async(async bool) ApiImageVersionsUpdateImageVersionResourcePoolsRequest {
	r.async = &async
	return r
}

func (r ApiImageVersionsUpdateImageVersionResourcePoolsRequest) Execute() (*http.Response, error) {
	return r.ApiService.ImageVersionsUpdateImageVersionResourcePoolsExecute(r)
}

/*
ImageVersionsUpdateImageVersionResourcePools Method for ImageVersionsUpdateImageVersionResourcePools

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiImageVersionsUpdateImageVersionResourcePoolsRequest
*/
func (a *ImageVersionsAPIsDAASService) ImageVersionsUpdateImageVersionResourcePools(ctx context.Context, id string) ApiImageVersionsUpdateImageVersionResourcePoolsRequest {
	return ApiImageVersionsUpdateImageVersionResourcePoolsRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ImageVersionsAPIsDAASService) ImageVersionsUpdateImageVersionResourcePoolsExecute(r ApiImageVersionsUpdateImageVersionResourcePoolsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImageVersionsAPIsDAASService.ImageVersionsUpdateImageVersionResourcePools")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ImageVersions/{id}/$UpdateResourcePools"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.citrixCustomerId == nil {
		return nil, reportError("citrixCustomerId is required and must be specified")
	}
	if r.citrixInstanceId == nil {
		return nil, reportError("citrixInstanceId is required and must be specified")
	}
	if r.updateImageVersionResourcePoolsRequestModel == nil {
		return nil, reportError("updateImageVersionResourcePoolsRequestModel is required and must be specified")
	}

	if r.async != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "async", r.async, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Citrix-CustomerId", r.citrixCustomerId, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Citrix-InstanceId", r.citrixInstanceId, "")
	if r.userAgent != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "User-Agent", r.userAgent, "")
	}
	if r.authorization != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	}
	if r.citrixTransactionId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Citrix-TransactionId", r.citrixTransactionId, "")
	}
	if r.accept != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "")
	}
	if r.citrixLocale != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Citrix-Locale", r.citrixLocale, "")
	}
	// body params
	localVarPostBody = r.updateImageVersionResourcePoolsRequestModel
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 406 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 415 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}