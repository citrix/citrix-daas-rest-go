# \AppLibPackageDiscoveryTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppLibPackageDiscoveryTPCreateAppLibPackageDiscovery**](AppLibPackageDiscoveryTPApi.md#AppLibPackageDiscoveryTPCreateAppLibPackageDiscovery) | **Post** /techpreview/{customerid}/{siteid}/AppLibPackageDiscovery/Sessions | Create AppLib Package Discovery session
[**AppLibPackageDiscoveryTPCreateAppLibPackageDiscoveryProfile**](AppLibPackageDiscoveryTPApi.md#AppLibPackageDiscoveryTPCreateAppLibPackageDiscoveryProfile) | **Post** /techpreview/{customerid}/{siteid}/AppLibPackageDiscovery/Profiles | Create an AppLib Package Discovery profile.
[**AppLibPackageDiscoveryTPGetAppLibPackageDiscovery**](AppLibPackageDiscoveryTPApi.md#AppLibPackageDiscoveryTPGetAppLibPackageDiscovery) | **Get** /techpreview/{customerid}/{siteid}/AppLibPackageDiscovery/Sessions/{id} | Get details of an AppLib Package Discovery session.
[**AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryLatestSessionByProfileId**](AppLibPackageDiscoveryTPApi.md#AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryLatestSessionByProfileId) | **Get** /techpreview/{customerid}/{siteid}/AppLibPackageDiscovery/Profiles/{uid}/LatestSession | Get the latest AppLib Package Discovery session for the specified profile id.
[**AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryProfile**](AppLibPackageDiscoveryTPApi.md#AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryProfile) | **Get** /techpreview/{customerid}/{siteid}/AppLibPackageDiscovery/Profiles/{uid} | Get details of an AppLib Package Discovery profile.
[**AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryProfiles**](AppLibPackageDiscoveryTPApi.md#AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryProfiles) | **Get** /techpreview/{customerid}/{siteid}/AppLibPackageDiscovery/Profiles | Get a list of AppLib Package Discovery profiles.
[**AppLibPackageDiscoveryTPGetAppLibPackageDiscoverySessions**](AppLibPackageDiscoveryTPApi.md#AppLibPackageDiscoveryTPGetAppLibPackageDiscoverySessions) | **Get** /techpreview/{customerid}/{siteid}/AppLibPackageDiscovery/Sessions | Get a list of AppLib Package Discovery sessions.
[**AppLibPackageDiscoveryTPRemoveAppLibPackageDiscoveryProfile**](AppLibPackageDiscoveryTPApi.md#AppLibPackageDiscoveryTPRemoveAppLibPackageDiscoveryProfile) | **Delete** /techpreview/{customerid}/{siteid}/AppLibPackageDiscovery/Profiles/{uid} | Remove an AppLib Package Discovery profile.
[**AppLibPackageDiscoveryTPUpdateAppLibPackageDiscoveryProfile**](AppLibPackageDiscoveryTPApi.md#AppLibPackageDiscoveryTPUpdateAppLibPackageDiscoveryProfile) | **Patch** /techpreview/{customerid}/{siteid}/AppLibPackageDiscovery/Profiles/{uid} | Update the specified AppLib Package Discovery profile.



## AppLibPackageDiscoveryTPCreateAppLibPackageDiscovery

> AppLibPackageDiscoveryResponseModel AppLibPackageDiscoveryTPCreateAppLibPackageDiscovery(ctx, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewAppLibPackageDiscoveryRequestModel() // AppLibPackageDiscoveryRequestModel | Request model to create a new AppLib package discovery session.
    async := true // bool | If `true`, the appLib package discovery session will be created as a background task. The task will have jobType CreateAppLibPackageDiscovery> (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPCreateAppLibPackageDiscovery(context.Background(), customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPCreateAppLibPackageDiscovery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppLibPackageDiscoveryTPCreateAppLibPackageDiscovery`: AppLibPackageDiscoveryResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPCreateAppLibPackageDiscovery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppLibPackageDiscoveryTPCreateAppLibPackageDiscoveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**AppLibPackageDiscoveryRequestModel**](AppLibPackageDiscoveryRequestModel.md) | Request model to create a new AppLib package discovery session. | 
 **async** | **bool** | If &#x60;true&#x60;, the appLib package discovery session will be created as a background task. The task will have jobType CreateAppLibPackageDiscovery&gt; | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## AppLibPackageDiscoveryTPCreateAppLibPackageDiscoveryProfile

> AppLibPackageDiscoveryProfileResponseModel AppLibPackageDiscoveryTPCreateAppLibPackageDiscoveryProfile(ctx, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewAppLibPackageDiscoveryProfileRequestModel(openapiclient.AppLibPackageDiscoveryType("Unknown"), "Name_example", int32(123)) // AppLibPackageDiscoveryProfileRequestModel | Request model to create a new AppLib package discovery profile.
    async := true // bool | If `true`, the appLib package discovery profile will be created as a background task. The task will have JobType CreateAppLibPackageDiscoveryProfile. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPCreateAppLibPackageDiscoveryProfile(context.Background(), customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPCreateAppLibPackageDiscoveryProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppLibPackageDiscoveryTPCreateAppLibPackageDiscoveryProfile`: AppLibPackageDiscoveryProfileResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPCreateAppLibPackageDiscoveryProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppLibPackageDiscoveryTPCreateAppLibPackageDiscoveryProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**AppLibPackageDiscoveryProfileRequestModel**](AppLibPackageDiscoveryProfileRequestModel.md) | Request model to create a new AppLib package discovery profile. | 
 **async** | **bool** | If &#x60;true&#x60;, the appLib package discovery profile will be created as a background task. The task will have JobType CreateAppLibPackageDiscoveryProfile. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## AppLibPackageDiscoveryTPGetAppLibPackageDiscovery

> AppLibPackageDiscoveryResponseModel AppLibPackageDiscoveryTPGetAppLibPackageDiscovery(ctx, id, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    id := "id_example" // string | Guid of the applib package discovery session that need to be fetched.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the appLib package discovery session will be created as a background task. The task will have jobType GetAppLibPackageDiscovery> (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPGetAppLibPackageDiscovery(context.Background(), id, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPGetAppLibPackageDiscovery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppLibPackageDiscoveryTPGetAppLibPackageDiscovery`: AppLibPackageDiscoveryResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPGetAppLibPackageDiscovery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Guid of the applib package discovery session that need to be fetched. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppLibPackageDiscoveryTPGetAppLibPackageDiscoveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, the appLib package discovery session will be created as a background task. The task will have jobType GetAppLibPackageDiscovery&gt; | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryLatestSessionByProfileId

> AppLibPackageDiscoveryResponseModel AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryLatestSessionByProfileId(ctx, uid, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    uid := int32(56) // int32 | The profile id.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryLatestSessionByProfileId(context.Background(), uid, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryLatestSessionByProfileId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryLatestSessionByProfileId`: AppLibPackageDiscoveryResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryLatestSessionByProfileId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **int32** | The profile id. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppLibPackageDiscoveryTPGetAppLibPackageDiscoveryLatestSessionByProfileIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryProfile

> AppLibPackageDiscoveryProfileResponseModel AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryProfile(ctx, uid, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    uid := int32(56) // int32 | Uid of the appLib package discovery profile that need to be fetched.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the appLib package discovery profile will be fetched as a background task. The task will have JobType GetAppLibPackageDiscoveryProfile. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryProfile(context.Background(), uid, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryProfile`: AppLibPackageDiscoveryProfileResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **int32** | Uid of the appLib package discovery profile that need to be fetched. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppLibPackageDiscoveryTPGetAppLibPackageDiscoveryProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, the appLib package discovery profile will be fetched as a background task. The task will have JobType GetAppLibPackageDiscoveryProfile. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryProfiles

> AppLibPackageDiscoveryProfileResponseModelCollection AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryProfiles(ctx, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the appLib package discovery profiles will be fetched as a background task. The task will have JobType GetAppLibPackageDiscoveryProfiles. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryProfiles(context.Background(), customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryProfiles`: AppLibPackageDiscoveryProfileResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPGetAppLibPackageDiscoveryProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppLibPackageDiscoveryTPGetAppLibPackageDiscoveryProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **async** | **bool** | If &#x60;true&#x60;, the appLib package discovery profiles will be fetched as a background task. The task will have JobType GetAppLibPackageDiscoveryProfiles. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## AppLibPackageDiscoveryTPGetAppLibPackageDiscoverySessions

> AppLibPackageDiscoveryResponseModelCollection AppLibPackageDiscoveryTPGetAppLibPackageDiscoverySessions(ctx, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the appLib package discovery sessions will be fetched as a background task. The task will have JobType GetAppLibPackageDiscoveries. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPGetAppLibPackageDiscoverySessions(context.Background(), customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPGetAppLibPackageDiscoverySessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppLibPackageDiscoveryTPGetAppLibPackageDiscoverySessions`: AppLibPackageDiscoveryResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPGetAppLibPackageDiscoverySessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppLibPackageDiscoveryTPGetAppLibPackageDiscoverySessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **async** | **bool** | If &#x60;true&#x60;, the appLib package discovery sessions will be fetched as a background task. The task will have JobType GetAppLibPackageDiscoveries. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## AppLibPackageDiscoveryTPRemoveAppLibPackageDiscoveryProfile

> AppLibPackageDiscoveryTPRemoveAppLibPackageDiscoveryProfile(ctx, uid, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    uid := int32(56) // int32 | Uid of the appLib package discovery profile that need to be removed.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the appLib package discovery profile will be removed as a background task. The task will have JobType RemoveAppLibPackageDiscoveryProfile. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPRemoveAppLibPackageDiscoveryProfile(context.Background(), uid, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPRemoveAppLibPackageDiscoveryProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **int32** | Uid of the appLib package discovery profile that need to be removed. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppLibPackageDiscoveryTPRemoveAppLibPackageDiscoveryProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, the appLib package discovery profile will be removed as a background task. The task will have JobType RemoveAppLibPackageDiscoveryProfile. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## AppLibPackageDiscoveryTPUpdateAppLibPackageDiscoveryProfile

> AppLibPackageDiscoveryTPUpdateAppLibPackageDiscoveryProfile(ctx, uid, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    uid := int32(56) // int32 | Uid of the appLib package discovery profile that need to be fetched.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewEditAppLibPackageDiscoveryProfileRequestModel() // EditAppLibPackageDiscoveryProfileRequestModel | Request model to update a new AppLib package discovery profile.
    async := true // bool | If `true`, the appLib package discovery profile will be fetched as a background task. The task will have JobType UpdateAppLibPackageDiscoveryProfile. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPUpdateAppLibPackageDiscoveryProfile(context.Background(), uid, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppLibPackageDiscoveryTPApi.AppLibPackageDiscoveryTPUpdateAppLibPackageDiscoveryProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **int32** | Uid of the appLib package discovery profile that need to be fetched. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppLibPackageDiscoveryTPUpdateAppLibPackageDiscoveryProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**EditAppLibPackageDiscoveryProfileRequestModel**](EditAppLibPackageDiscoveryProfileRequestModel.md) | Request model to update a new AppLib package discovery profile. | 
 **async** | **bool** | If &#x60;true&#x60;, the appLib package discovery profile will be fetched as a background task. The task will have JobType UpdateAppLibPackageDiscoveryProfile. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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

