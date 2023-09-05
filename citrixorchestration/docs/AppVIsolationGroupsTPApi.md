# \AppVIsolationGroupsTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppVIsolationGroupsTPCreateAppVIsolationGroup**](AppVIsolationGroupsTPApi.md#AppVIsolationGroupsTPCreateAppVIsolationGroup) | **Post** /techpreview/{customerid}/{siteid}/AppVIsolationGroups | Create an App-V IsolationGroup in the site
[**AppVIsolationGroupsTPDeleteAppVIsolationGroup**](AppVIsolationGroupsTPApi.md#AppVIsolationGroupsTPDeleteAppVIsolationGroup) | **Delete** /techpreview/{customerid}/{siteid}/AppVIsolationGroups/{nameOrId} | Delete an App-V IsolationGroup configured in the site.
[**AppVIsolationGroupsTPGetAppVIsolationGroup**](AppVIsolationGroupsTPApi.md#AppVIsolationGroupsTPGetAppVIsolationGroup) | **Get** /techpreview/{customerid}/{siteid}/AppVIsolationGroups/{nameOrId} | Get the specified App-V IsolationGroups configured in the site
[**AppVIsolationGroupsTPGetAppVIsolationGroups**](AppVIsolationGroupsTPApi.md#AppVIsolationGroupsTPGetAppVIsolationGroups) | **Get** /techpreview/{customerid}/{siteid}/AppVIsolationGroups | Get the App-V IsolationGroups configured in the site
[**AppVIsolationGroupsTPUpdateAppVIsolationGroup**](AppVIsolationGroupsTPApi.md#AppVIsolationGroupsTPUpdateAppVIsolationGroup) | **Patch** /techpreview/{customerid}/{siteid}/AppVIsolationGroups/{nameOrId} | Update the App-V IsolationGroup configured in the site.



## AppVIsolationGroupsTPCreateAppVIsolationGroup

> AppVIsolationGroupsTPCreateAppVIsolationGroup(ctx, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewCreateAppVIsolationGroupRequestModel("Name_example") // CreateAppVIsolationGroupRequestModel | Request model of the new IsolationGroup.
    async := true // bool | If `true`, the tags will be modified as a background task. The task will have JobType CreateAppVIsolationGroup. When the task is complete it will redirect to GetAppVIsolationGroups. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppVIsolationGroupsTPApi.AppVIsolationGroupsTPCreateAppVIsolationGroup(context.Background(), customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVIsolationGroupsTPApi.AppVIsolationGroupsTPCreateAppVIsolationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVIsolationGroupsTPCreateAppVIsolationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateAppVIsolationGroupRequestModel**](CreateAppVIsolationGroupRequestModel.md) | Request model of the new IsolationGroup. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. The task will have JobType CreateAppVIsolationGroup. When the task is complete it will redirect to GetAppVIsolationGroups. | [default to false]
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


## AppVIsolationGroupsTPDeleteAppVIsolationGroup

> AppVIsolationGroupsTPDeleteAppVIsolationGroup(ctx, nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the tags will be modified as a background task. The task will have JobType CreateAppVIsolationGroup. When the task is complete it will redirect to GetAppVIsolationGroups. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppVIsolationGroupsTPApi.AppVIsolationGroupsTPDeleteAppVIsolationGroup(context.Background(), nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVIsolationGroupsTPApi.AppVIsolationGroupsTPDeleteAppVIsolationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or UID of an isolationGroup. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVIsolationGroupsTPDeleteAppVIsolationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. The task will have JobType CreateAppVIsolationGroup. When the task is complete it will redirect to GetAppVIsolationGroups. | [default to false]
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


## AppVIsolationGroupsTPGetAppVIsolationGroup

> AppVIsolationGroupResponseModel AppVIsolationGroupsTPGetAppVIsolationGroup(ctx, nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the tags will be modified as a background task. The task will have JobType GetAppVIsolationGroup. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVIsolationGroupsTPApi.AppVIsolationGroupsTPGetAppVIsolationGroup(context.Background(), nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVIsolationGroupsTPApi.AppVIsolationGroupsTPGetAppVIsolationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVIsolationGroupsTPGetAppVIsolationGroup`: AppVIsolationGroupResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AppVIsolationGroupsTPApi.AppVIsolationGroupsTPGetAppVIsolationGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or UID of an isolationGroup. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVIsolationGroupsTPGetAppVIsolationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. The task will have JobType GetAppVIsolationGroup. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## AppVIsolationGroupsTPGetAppVIsolationGroups

> AppVIsolationGroupResponseModelCollection AppVIsolationGroupsTPGetAppVIsolationGroups(ctx, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool |  (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppVIsolationGroupsTPApi.AppVIsolationGroupsTPGetAppVIsolationGroups(context.Background(), customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVIsolationGroupsTPApi.AppVIsolationGroupsTPGetAppVIsolationGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVIsolationGroupsTPGetAppVIsolationGroups`: AppVIsolationGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AppVIsolationGroupsTPApi.AppVIsolationGroupsTPGetAppVIsolationGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVIsolationGroupsTPGetAppVIsolationGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **async** | **bool** |  | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

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


## AppVIsolationGroupsTPUpdateAppVIsolationGroup

> AppVIsolationGroupsTPUpdateAppVIsolationGroup(ctx, nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewUpdateAppVIsolationGroupRequestModel() // UpdateAppVIsolationGroupRequestModel | request model to update isolationGroup.
    async := true // bool | If `true`, the tags will be modified as a background task. The task will have JobType UpdateAppVIsolationGroup. When the task is complete it will redirect to GetAppVIsolationGroups. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppVIsolationGroupsTPApi.AppVIsolationGroupsTPUpdateAppVIsolationGroup(context.Background(), nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppVIsolationGroupsTPApi.AppVIsolationGroupsTPUpdateAppVIsolationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or UID of an isolationGroup. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVIsolationGroupsTPUpdateAppVIsolationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**UpdateAppVIsolationGroupRequestModel**](UpdateAppVIsolationGroupRequestModel.md) | request model to update isolationGroup. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. The task will have JobType UpdateAppVIsolationGroup. When the task is complete it will redirect to GetAppVIsolationGroups. | [default to false]
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

