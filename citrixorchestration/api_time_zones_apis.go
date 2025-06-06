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
)

// TimeZonesAPIsDAASService TimeZonesAPIsDAAS service
type TimeZonesAPIsDAASService service

type ApiTimeZonesGetTimeZonesRequest struct {
	ctx                 context.Context
	ApiService          *TimeZonesAPIsDAASService
	citrixCustomerId    *string
	citrixInstanceId    *string
	userAgent           *string
	authorization       *string
	citrixTransactionId *string
	accept              *string
	citrixLocale        *string
}

// Citrix Customer ID. Default is &#39;CitrixOnPremises&#39;
func (r ApiTimeZonesGetTimeZonesRequest) CitrixCustomerId(citrixCustomerId string) ApiTimeZonesGetTimeZonesRequest {
	r.citrixCustomerId = &citrixCustomerId
	return r
}

// Citrix Instance (Site) ID.
func (r ApiTimeZonesGetTimeZonesRequest) CitrixInstanceId(citrixInstanceId string) ApiTimeZonesGetTimeZonesRequest {
	r.citrixInstanceId = &citrixInstanceId
	return r
}

// User Agent type of the request.
func (r ApiTimeZonesGetTimeZonesRequest) UserAgent(userAgent string) ApiTimeZonesGetTimeZonesRequest {
	r.userAgent = &userAgent
	return r
}

// Citrix authorization header: CWSAuth Bearer&#x3D;{token}
func (r ApiTimeZonesGetTimeZonesRequest) Authorization(authorization string) ApiTimeZonesGetTimeZonesRequest {
	r.authorization = &authorization
	return r
}

// Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned.
func (r ApiTimeZonesGetTimeZonesRequest) CitrixTransactionId(citrixTransactionId string) ApiTimeZonesGetTimeZonesRequest {
	r.citrixTransactionId = &citrixTransactionId
	return r
}

// Must accept application/json.
func (r ApiTimeZonesGetTimeZonesRequest) Accept(accept string) ApiTimeZonesGetTimeZonesRequest {
	r.accept = &accept
	return r
}

// Locale of the request.
func (r ApiTimeZonesGetTimeZonesRequest) CitrixLocale(citrixLocale string) ApiTimeZonesGetTimeZonesRequest {
	r.citrixLocale = &citrixLocale
	return r
}

func (r ApiTimeZonesGetTimeZonesRequest) Execute() (*TimeZoneResponseModelCollection, *http.Response, error) {
	return r.ApiService.TimeZonesGetTimeZonesExecute(r)
}

/*
TimeZonesGetTimeZones Get a list of time zones supported by the site.

Get a (non-exhaustive) list of time zones supported by the site,
with time zone names represented in the caller's locale.

Note that all IANA (https://www.iana.org/time-zones) time zones are
supported as inputs to time zone APIs.  However sometimes callers
need a suitable list of time zones to display as options within a
UI. This API provides that capability. The options returned
represent a reasonable palette of options for a caller to choose
from when selecting a time zone.

The localized time zone names come from the Unicode CLDR
(http://cldr.unicode.org/) project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiTimeZonesGetTimeZonesRequest
*/
func (a *TimeZonesAPIsDAASService) TimeZonesGetTimeZones(ctx context.Context) ApiTimeZonesGetTimeZonesRequest {
	return ApiTimeZonesGetTimeZonesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TimeZoneResponseModelCollection
func (a *TimeZonesAPIsDAASService) TimeZonesGetTimeZonesExecute(r ApiTimeZonesGetTimeZonesRequest) (*TimeZoneResponseModelCollection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TimeZoneResponseModelCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TimeZonesAPIsDAASService.TimeZonesGetTimeZones")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/TimeZones/All"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.citrixCustomerId == nil {
		return localVarReturnValue, nil, reportError("citrixCustomerId is required and must be specified")
	}
	if r.citrixInstanceId == nil {
		return localVarReturnValue, nil, reportError("citrixInstanceId is required and must be specified")
	}

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 406 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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
