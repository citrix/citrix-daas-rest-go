# \AppLibPackageDiscoveryAPIsDAAS

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppLibPackageDiscoveryCreateAppLibPackageDiscovery**](AppLibPackageDiscoveryAPIsDAAS.md#AppLibPackageDiscoveryCreateAppLibPackageDiscovery) | **Post** /AppLibPackageDiscovery/Sessions | Create AppLib Package Discovery session
[**AppLibPackageDiscoveryCreateAppLibPackageDiscoveryProfile**](AppLibPackageDiscoveryAPIsDAAS.md#AppLibPackageDiscoveryCreateAppLibPackageDiscoveryProfile) | **Post** /AppLibPackageDiscovery/Profiles | Create an AppLib Package Discovery profile.
[**AppLibPackageDiscoveryGetAppLibPackageDiscovery**](AppLibPackageDiscoveryAPIsDAAS.md#AppLibPackageDiscoveryGetAppLibPackageDiscovery) | **Get** /AppLibPackageDiscovery/Sessions/{id} | Get details of an AppLib Package Discovery session.
[**AppLibPackageDiscoveryGetAppLibPackageDiscoveryLatestSessionByProfileId**](AppLibPackageDiscoveryAPIsDAAS.md#AppLibPackageDiscoveryGetAppLibPackageDiscoveryLatestSessionByProfileId) | **Get** /AppLibPackageDiscovery/Profiles/{uid}/LatestSession | Get the latest AppLib Package Discovery session for the specified profile id.
[**AppLibPackageDiscoveryGetAppLibPackageDiscoveryProfile**](AppLibPackageDiscoveryAPIsDAAS.md#AppLibPackageDiscoveryGetAppLibPackageDiscoveryProfile) | **Get** /AppLibPackageDiscovery/Profiles/{uid} | Get details of an AppLib Package Discovery profile.
[**AppLibPackageDiscoveryGetAppLibPackageDiscoveryProfiles**](AppLibPackageDiscoveryAPIsDAAS.md#AppLibPackageDiscoveryGetAppLibPackageDiscoveryProfiles) | **Get** /AppLibPackageDiscovery/Profiles | Get a list of AppLib Package Discovery profiles.
[**AppLibPackageDiscoveryGetAppLibPackageDiscoverySessions**](AppLibPackageDiscoveryAPIsDAAS.md#AppLibPackageDiscoveryGetAppLibPackageDiscoverySessions) | **Get** /AppLibPackageDiscovery/Sessions | Get a list of AppLib Package Discovery sessions.
[**AppLibPackageDiscoveryRemoveAppLibPackageDiscoveryProfile**](AppLibPackageDiscoveryAPIsDAAS.md#AppLibPackageDiscoveryRemoveAppLibPackageDiscoveryProfile) | **Delete** /AppLibPackageDiscovery/Profiles/{uid} | Remove an AppLib Package Discovery profile.
[**AppLibPackageDiscoveryUpdateAppLibPackageDiscoveryProfile**](AppLibPackageDiscoveryAPIsDAAS.md#AppLibPackageDiscoveryUpdateAppLibPackageDiscoveryProfile) | **Patch** /AppLibPackageDiscovery/Profiles/{uid} | Update the specified AppLib Package Discovery profile.



## AppLibPackageDiscoveryCreateAppLibPackageDiscovery

> AppLibPackageDiscoveryResponseModel AppLibPackageDiscoveryCreateAppLibPackageDiscovery(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).AppLibPackageDiscoveryRequestModel(appLibPackageDiscoveryRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Create AppLib Package Discovery session

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	appLibPackageDiscoveryRequestModel := *openapiclient.NewAppLibPackageDiscoveryRequestModel() // AppLibPackageDiscoveryRequestModel | Request model to create a new AppLib package discovery session.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the appLib package discovery session will be created as a background task. The task will have jobType CreateAppLibPackageDiscovery> (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryCreateAppLibPackageDiscovery(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).AppLibPackageDiscoveryRequestModel(appLibPackageDiscoveryRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryCreateAppLibPackageDiscovery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppLibPackageDiscoveryCreateAppLibPackageDiscovery`: AppLibPackageDiscoveryResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryCreateAppLibPackageDiscovery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppLibPackageDiscoveryCreateAppLibPackageDiscoveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **appLibPackageDiscoveryRequestModel** | [**AppLibPackageDiscoveryRequestModel**](AppLibPackageDiscoveryRequestModel.md) | Request model to create a new AppLib package discovery session. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the appLib package discovery session will be created as a background task. The task will have jobType CreateAppLibPackageDiscovery&gt; | [default to false]

### Return type

[**AppLibPackageDiscoveryResponseModel**](AppLibPackageDiscoveryResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppLibPackageDiscoveryCreateAppLibPackageDiscoveryProfile

> AppLibPackageDiscoveryProfileResponseModel AppLibPackageDiscoveryCreateAppLibPackageDiscoveryProfile(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).AppLibPackageDiscoveryProfileRequestModel(appLibPackageDiscoveryProfileRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Create an AppLib Package Discovery profile.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	appLibPackageDiscoveryProfileRequestModel := *openapiclient.NewAppLibPackageDiscoveryProfileRequestModel(openapiclient.AppLibPackageDiscoveryType("Unknown"), "Name_example", int32(123)) // AppLibPackageDiscoveryProfileRequestModel | Request model to create a new AppLib package discovery profile.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the appLib package discovery profile will be created as a background task. The task will have JobType CreateAppLibPackageDiscoveryProfile. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryCreateAppLibPackageDiscoveryProfile(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).AppLibPackageDiscoveryProfileRequestModel(appLibPackageDiscoveryProfileRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryCreateAppLibPackageDiscoveryProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppLibPackageDiscoveryCreateAppLibPackageDiscoveryProfile`: AppLibPackageDiscoveryProfileResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryCreateAppLibPackageDiscoveryProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppLibPackageDiscoveryCreateAppLibPackageDiscoveryProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **appLibPackageDiscoveryProfileRequestModel** | [**AppLibPackageDiscoveryProfileRequestModel**](AppLibPackageDiscoveryProfileRequestModel.md) | Request model to create a new AppLib package discovery profile. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the appLib package discovery profile will be created as a background task. The task will have JobType CreateAppLibPackageDiscoveryProfile. | [default to false]

### Return type

[**AppLibPackageDiscoveryProfileResponseModel**](AppLibPackageDiscoveryProfileResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppLibPackageDiscoveryGetAppLibPackageDiscovery

> AppLibPackageDiscoveryResponseModel AppLibPackageDiscoveryGetAppLibPackageDiscovery(ctx, id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Get details of an AppLib Package Discovery session.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	id := "id_example" // string | Guid of the applib package discovery session that need to be fetched.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the appLib package discovery session will be created as a background task. The task will have jobType GetAppLibPackageDiscovery> (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryGetAppLibPackageDiscovery(context.Background(), id).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryGetAppLibPackageDiscovery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppLibPackageDiscoveryGetAppLibPackageDiscovery`: AppLibPackageDiscoveryResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryGetAppLibPackageDiscovery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Guid of the applib package discovery session that need to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppLibPackageDiscoveryGetAppLibPackageDiscoveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the appLib package discovery session will be created as a background task. The task will have jobType GetAppLibPackageDiscovery&gt; | [default to false]

### Return type

[**AppLibPackageDiscoveryResponseModel**](AppLibPackageDiscoveryResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppLibPackageDiscoveryGetAppLibPackageDiscoveryLatestSessionByProfileId

> AppLibPackageDiscoveryResponseModel AppLibPackageDiscoveryGetAppLibPackageDiscoveryLatestSessionByProfileId(ctx, uid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the latest AppLib Package Discovery session for the specified profile id.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	uid := int32(56) // int32 | The profile id.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryGetAppLibPackageDiscoveryLatestSessionByProfileId(context.Background(), uid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryGetAppLibPackageDiscoveryLatestSessionByProfileId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppLibPackageDiscoveryGetAppLibPackageDiscoveryLatestSessionByProfileId`: AppLibPackageDiscoveryResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryGetAppLibPackageDiscoveryLatestSessionByProfileId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **int32** | The profile id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppLibPackageDiscoveryGetAppLibPackageDiscoveryLatestSessionByProfileIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**AppLibPackageDiscoveryResponseModel**](AppLibPackageDiscoveryResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppLibPackageDiscoveryGetAppLibPackageDiscoveryProfile

> AppLibPackageDiscoveryProfileResponseModel AppLibPackageDiscoveryGetAppLibPackageDiscoveryProfile(ctx, uid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Get details of an AppLib Package Discovery profile.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	uid := int32(56) // int32 | Uid of the appLib package discovery profile that need to be fetched.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the appLib package discovery profile will be fetched as a background task. The task will have JobType GetAppLibPackageDiscoveryProfile. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryGetAppLibPackageDiscoveryProfile(context.Background(), uid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryGetAppLibPackageDiscoveryProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppLibPackageDiscoveryGetAppLibPackageDiscoveryProfile`: AppLibPackageDiscoveryProfileResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryGetAppLibPackageDiscoveryProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **int32** | Uid of the appLib package discovery profile that need to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppLibPackageDiscoveryGetAppLibPackageDiscoveryProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the appLib package discovery profile will be fetched as a background task. The task will have JobType GetAppLibPackageDiscoveryProfile. | [default to false]

### Return type

[**AppLibPackageDiscoveryProfileResponseModel**](AppLibPackageDiscoveryProfileResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppLibPackageDiscoveryGetAppLibPackageDiscoveryProfiles

> AppLibPackageDiscoveryProfileResponseModelCollection AppLibPackageDiscoveryGetAppLibPackageDiscoveryProfiles(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Get a list of AppLib Package Discovery profiles.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the appLib package discovery profiles will be fetched as a background task. The task will have JobType GetAppLibPackageDiscoveryProfiles. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryGetAppLibPackageDiscoveryProfiles(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryGetAppLibPackageDiscoveryProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppLibPackageDiscoveryGetAppLibPackageDiscoveryProfiles`: AppLibPackageDiscoveryProfileResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryGetAppLibPackageDiscoveryProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppLibPackageDiscoveryGetAppLibPackageDiscoveryProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the appLib package discovery profiles will be fetched as a background task. The task will have JobType GetAppLibPackageDiscoveryProfiles. | [default to false]

### Return type

[**AppLibPackageDiscoveryProfileResponseModelCollection**](AppLibPackageDiscoveryProfileResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppLibPackageDiscoveryGetAppLibPackageDiscoverySessions

> AppLibPackageDiscoveryResponseModelCollection AppLibPackageDiscoveryGetAppLibPackageDiscoverySessions(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Get a list of AppLib Package Discovery sessions.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the appLib package discovery sessions will be fetched as a background task. The task will have JobType GetAppLibPackageDiscoveries. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryGetAppLibPackageDiscoverySessions(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryGetAppLibPackageDiscoverySessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppLibPackageDiscoveryGetAppLibPackageDiscoverySessions`: AppLibPackageDiscoveryResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryGetAppLibPackageDiscoverySessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppLibPackageDiscoveryGetAppLibPackageDiscoverySessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the appLib package discovery sessions will be fetched as a background task. The task will have JobType GetAppLibPackageDiscoveries. | [default to false]

### Return type

[**AppLibPackageDiscoveryResponseModelCollection**](AppLibPackageDiscoveryResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppLibPackageDiscoveryRemoveAppLibPackageDiscoveryProfile

> AppLibPackageDiscoveryRemoveAppLibPackageDiscoveryProfile(ctx, uid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Remove an AppLib Package Discovery profile.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	uid := int32(56) // int32 | Uid of the appLib package discovery profile that need to be removed.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the appLib package discovery profile will be removed as a background task. The task will have JobType RemoveAppLibPackageDiscoveryProfile. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryRemoveAppLibPackageDiscoveryProfile(context.Background(), uid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryRemoveAppLibPackageDiscoveryProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **int32** | Uid of the appLib package discovery profile that need to be removed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppLibPackageDiscoveryRemoveAppLibPackageDiscoveryProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the appLib package discovery profile will be removed as a background task. The task will have JobType RemoveAppLibPackageDiscoveryProfile. | [default to false]

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppLibPackageDiscoveryUpdateAppLibPackageDiscoveryProfile

> AppLibPackageDiscoveryUpdateAppLibPackageDiscoveryProfile(ctx, uid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).EditAppLibPackageDiscoveryProfileRequestModel(editAppLibPackageDiscoveryProfileRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Update the specified AppLib Package Discovery profile.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	uid := int32(56) // int32 | Uid of the appLib package discovery profile that need to be fetched.
	editAppLibPackageDiscoveryProfileRequestModel := *openapiclient.NewEditAppLibPackageDiscoveryProfileRequestModel() // EditAppLibPackageDiscoveryProfileRequestModel | Request model to update a new AppLib package discovery profile.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the appLib package discovery profile will be fetched as a background task. The task will have JobType UpdateAppLibPackageDiscoveryProfile. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryUpdateAppLibPackageDiscoveryProfile(context.Background(), uid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).EditAppLibPackageDiscoveryProfileRequestModel(editAppLibPackageDiscoveryProfileRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppLibPackageDiscoveryAPIsDAAS.AppLibPackageDiscoveryUpdateAppLibPackageDiscoveryProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **int32** | Uid of the appLib package discovery profile that need to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppLibPackageDiscoveryUpdateAppLibPackageDiscoveryProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **editAppLibPackageDiscoveryProfileRequestModel** | [**EditAppLibPackageDiscoveryProfileRequestModel**](EditAppLibPackageDiscoveryProfileRequestModel.md) | Request model to update a new AppLib package discovery profile. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the appLib package discovery profile will be fetched as a background task. The task will have JobType UpdateAppLibPackageDiscoveryProfile. | [default to false]

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

