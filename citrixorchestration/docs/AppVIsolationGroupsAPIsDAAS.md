# \AppVIsolationGroupsAPIsDAAS

All URIs are relative to *https://api-us.cloud.com/cvad/manage*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppVIsolationGroupsCreateAppVIsolationGroup**](AppVIsolationGroupsAPIsDAAS.md#AppVIsolationGroupsCreateAppVIsolationGroup) | **Post** /AppVIsolationGroups | Create an App-V IsolationGroup in the site
[**AppVIsolationGroupsDeleteAppVIsolationGroup**](AppVIsolationGroupsAPIsDAAS.md#AppVIsolationGroupsDeleteAppVIsolationGroup) | **Delete** /AppVIsolationGroups/{nameOrId} | Delete an App-V IsolationGroup configured in the site.
[**AppVIsolationGroupsGetAppVIsolationGroup**](AppVIsolationGroupsAPIsDAAS.md#AppVIsolationGroupsGetAppVIsolationGroup) | **Get** /AppVIsolationGroups/{nameOrId} | Get the specified App-V IsolationGroups configured in the site
[**AppVIsolationGroupsGetAppVIsolationGroups**](AppVIsolationGroupsAPIsDAAS.md#AppVIsolationGroupsGetAppVIsolationGroups) | **Get** /AppVIsolationGroups | Get the App-V IsolationGroups configured in the site
[**AppVIsolationGroupsUpdateAppVIsolationGroup**](AppVIsolationGroupsAPIsDAAS.md#AppVIsolationGroupsUpdateAppVIsolationGroup) | **Patch** /AppVIsolationGroups/{nameOrId} | Update the App-V IsolationGroup configured in the site.



## AppVIsolationGroupsCreateAppVIsolationGroup

> AppVIsolationGroupsCreateAppVIsolationGroup(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).CreateAppVIsolationGroupRequestModel(createAppVIsolationGroupRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Create an App-V IsolationGroup in the site



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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    createAppVIsolationGroupRequestModel := *openapiclient.NewCreateAppVIsolationGroupRequestModel("Name_example") // CreateAppVIsolationGroupRequestModel | Request model of the new IsolationGroup.
    async := true // bool | If `true`, the tags will be modified as a background task. The task will have JobType CreateAppVIsolationGroup. When the task is complete it will redirect to GetAppVIsolationGroups. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppVIsolationGroupsAPIsDAAS.AppVIsolationGroupsCreateAppVIsolationGroup(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).CreateAppVIsolationGroupRequestModel(createAppVIsolationGroupRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVIsolationGroupsAPIsDAAS.AppVIsolationGroupsCreateAppVIsolationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppVIsolationGroupsCreateAppVIsolationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **createAppVIsolationGroupRequestModel** | [**CreateAppVIsolationGroupRequestModel**](CreateAppVIsolationGroupRequestModel.md) | Request model of the new IsolationGroup. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. The task will have JobType CreateAppVIsolationGroup. When the task is complete it will redirect to GetAppVIsolationGroups. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## AppVIsolationGroupsDeleteAppVIsolationGroup

> AppVIsolationGroupsDeleteAppVIsolationGroup(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Delete an App-V IsolationGroup configured in the site.



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
    nameOrId := "nameOrId_example" // string | Name or UID of an isolationGroup.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    async := true // bool | If `true`, the tags will be modified as a background task. The task will have JobType CreateAppVIsolationGroup. When the task is complete it will redirect to GetAppVIsolationGroups. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppVIsolationGroupsAPIsDAAS.AppVIsolationGroupsDeleteAppVIsolationGroup(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVIsolationGroupsAPIsDAAS.AppVIsolationGroupsDeleteAppVIsolationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or UID of an isolationGroup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVIsolationGroupsDeleteAppVIsolationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. The task will have JobType CreateAppVIsolationGroup. When the task is complete it will redirect to GetAppVIsolationGroups. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## AppVIsolationGroupsGetAppVIsolationGroup

> AppVIsolationGroupResponseModel AppVIsolationGroupsGetAppVIsolationGroup(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the specified App-V IsolationGroups configured in the site



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
    nameOrId := "nameOrId_example" // string | Name or UID of an isolationGroup.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    async := true // bool | If `true`, the tags will be modified as a background task. The task will have JobType GetAppVIsolationGroup. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVIsolationGroupsAPIsDAAS.AppVIsolationGroupsGetAppVIsolationGroup(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVIsolationGroupsAPIsDAAS.AppVIsolationGroupsGetAppVIsolationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVIsolationGroupsGetAppVIsolationGroup`: AppVIsolationGroupResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AppVIsolationGroupsAPIsDAAS.AppVIsolationGroupsGetAppVIsolationGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or UID of an isolationGroup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVIsolationGroupsGetAppVIsolationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. The task will have JobType GetAppVIsolationGroup. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**AppVIsolationGroupResponseModel**](AppVIsolationGroupResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppVIsolationGroupsGetAppVIsolationGroups

> AppVIsolationGroupResponseModelCollection AppVIsolationGroupsGetAppVIsolationGroups(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the App-V IsolationGroups configured in the site



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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    async := true // bool |  (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVIsolationGroupsAPIsDAAS.AppVIsolationGroupsGetAppVIsolationGroups(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVIsolationGroupsAPIsDAAS.AppVIsolationGroupsGetAppVIsolationGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVIsolationGroupsGetAppVIsolationGroups`: AppVIsolationGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AppVIsolationGroupsAPIsDAAS.AppVIsolationGroupsGetAppVIsolationGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppVIsolationGroupsGetAppVIsolationGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **async** | **bool** |  | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**AppVIsolationGroupResponseModelCollection**](AppVIsolationGroupResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppVIsolationGroupsUpdateAppVIsolationGroup

> AppVIsolationGroupsUpdateAppVIsolationGroup(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).UpdateAppVIsolationGroupRequestModel(updateAppVIsolationGroupRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Update the App-V IsolationGroup configured in the site.



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
    nameOrId := "nameOrId_example" // string | Name or UID of an isolationGroup.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    updateAppVIsolationGroupRequestModel := *openapiclient.NewUpdateAppVIsolationGroupRequestModel() // UpdateAppVIsolationGroupRequestModel | request model to update isolationGroup.
    async := true // bool | If `true`, the tags will be modified as a background task. The task will have JobType UpdateAppVIsolationGroup. When the task is complete it will redirect to GetAppVIsolationGroups. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppVIsolationGroupsAPIsDAAS.AppVIsolationGroupsUpdateAppVIsolationGroup(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).UpdateAppVIsolationGroupRequestModel(updateAppVIsolationGroupRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVIsolationGroupsAPIsDAAS.AppVIsolationGroupsUpdateAppVIsolationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or UID of an isolationGroup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVIsolationGroupsUpdateAppVIsolationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **updateAppVIsolationGroupRequestModel** | [**UpdateAppVIsolationGroupRequestModel**](UpdateAppVIsolationGroupRequestModel.md) | request model to update isolationGroup. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. The task will have JobType UpdateAppVIsolationGroup. When the task is complete it will redirect to GetAppVIsolationGroups. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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

