# \ApplicationFoldersTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationFoldersTPCheckApplicationFolderPathExists**](ApplicationFoldersTPApi.md#ApplicationFoldersTPCheckApplicationFolderPathExists) | **Head** /techpreview/{customerid}/{siteid}/ApplicationFolders/{path} | Check for the existence of an application folder by path.
[**ApplicationFoldersTPCreateApplicationFolder**](ApplicationFoldersTPApi.md#ApplicationFoldersTPCreateApplicationFolder) | **Post** /techpreview/{customerid}/{siteid}/ApplicationFolders | Create an application folder
[**ApplicationFoldersTPDeleteApplicationFolder**](ApplicationFoldersTPApi.md#ApplicationFoldersTPDeleteApplicationFolder) | **Delete** /techpreview/{customerid}/{siteid}/ApplicationFolders/{pathOrId} | Delete an application folder.
[**ApplicationFoldersTPGetApplicationFolder**](ApplicationFoldersTPApi.md#ApplicationFoldersTPGetApplicationFolder) | **Get** /techpreview/{customerid}/{siteid}/ApplicationFolders/{pathOrId} | Get details about a single application folder.
[**ApplicationFoldersTPGetApplicationFolderApplications**](ApplicationFoldersTPApi.md#ApplicationFoldersTPGetApplicationFolderApplications) | **Get** /techpreview/{customerid}/{siteid}/ApplicationFolders/{pathOrId}/Applications | Get the applications in an application folder.
[**ApplicationFoldersTPGetApplicationFolders**](ApplicationFoldersTPApi.md#ApplicationFoldersTPGetApplicationFolders) | **Get** /techpreview/{customerid}/{siteid}/ApplicationFolders | Get application folders.
[**ApplicationFoldersTPUpdateApplicationFolder**](ApplicationFoldersTPApi.md#ApplicationFoldersTPUpdateApplicationFolder) | **Patch** /techpreview/{customerid}/{siteid}/ApplicationFolders/{pathOrId} | Update an application folder.



## ApplicationFoldersTPCheckApplicationFolderPathExists

> ApplicationFoldersTPCheckApplicationFolderPathExists(ctx, path, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Check for the existence of an application folder by path.



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
    path := "path_example" // string | Path of the application folder.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationFoldersTPApi.ApplicationFoldersTPCheckApplicationFolderPathExists(context.Background(), path, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationFoldersTPApi.ApplicationFoldersTPCheckApplicationFolderPathExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Path of the application folder. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationFoldersTPCheckApplicationFolderPathExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## ApplicationFoldersTPCreateApplicationFolder

> ApplicationFolderResponseModel ApplicationFoldersTPCreateApplicationFolder(ctx, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Create an application folder



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
    request := *openapiclient.NewCreateApplicationFolderRequestModel() // CreateApplicationFolderRequestModel | Details of the application folder to create.
    async := true // bool | If `true`, the application folder will be created as a background task. The task will have JobType \"JobType.CreateApplicationFolder\". When the task is complete it will redirect to \"GetApplicationFolder(string)\" The job Parameters will contain properties:  * _Name_ - Name of the folder being created. * _Path_ - Path to the folder being created. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationFoldersTPApi.ApplicationFoldersTPCreateApplicationFolder(context.Background(), customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationFoldersTPApi.ApplicationFoldersTPCreateApplicationFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationFoldersTPCreateApplicationFolder`: ApplicationFolderResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ApplicationFoldersTPApi.ApplicationFoldersTPCreateApplicationFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationFoldersTPCreateApplicationFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateApplicationFolderRequestModel**](CreateApplicationFolderRequestModel.md) | Details of the application folder to create. | 
 **async** | **bool** | If &#x60;true&#x60;, the application folder will be created as a background task. The task will have JobType \&quot;JobType.CreateApplicationFolder\&quot;. When the task is complete it will redirect to \&quot;GetApplicationFolder(string)\&quot; The job Parameters will contain properties:  * _Name_ - Name of the folder being created. * _Path_ - Path to the folder being created. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ApplicationFolderResponseModel**](ApplicationFolderResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationFoldersTPDeleteApplicationFolder

> ApplicationFoldersTPDeleteApplicationFolder(ctx, pathOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Delete an application folder.



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
    pathOrId := "pathOrId_example" // string | Name or ID of the application folder to delete.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the application folder (and associated objects) will be deleted as a background task. The task will have JobType DeleteApplicationFolder. When the task is complete it will redirect to GetApplicationFolders. The job's Parameters will contain properties:  * _Id_ - ID of the application folder being deleted, * _Path_ - Path of the application folder being deleted. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationFoldersTPApi.ApplicationFoldersTPDeleteApplicationFolder(context.Background(), pathOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationFoldersTPApi.ApplicationFoldersTPDeleteApplicationFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathOrId** | **string** | Name or ID of the application folder to delete. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationFoldersTPDeleteApplicationFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, the application folder (and associated objects) will be deleted as a background task. The task will have JobType DeleteApplicationFolder. When the task is complete it will redirect to GetApplicationFolders. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the application folder being deleted, * _Path_ - Path of the application folder being deleted. | [default to false]
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


## ApplicationFoldersTPGetApplicationFolder

> ApplicationFolderResponseModel ApplicationFoldersTPGetApplicationFolder(ctx, pathOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get details about a single application folder.



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
    pathOrId := "pathOrId_example" // string | Path (URL-encoded) or ID of the application folder.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationFoldersTPApi.ApplicationFoldersTPGetApplicationFolder(context.Background(), pathOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationFoldersTPApi.ApplicationFoldersTPGetApplicationFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationFoldersTPGetApplicationFolder`: ApplicationFolderResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ApplicationFoldersTPApi.ApplicationFoldersTPGetApplicationFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathOrId** | **string** | Path (URL-encoded) or ID of the application folder. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationFoldersTPGetApplicationFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ApplicationFolderResponseModel**](ApplicationFolderResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationFoldersTPGetApplicationFolderApplications

> ApplicationResponseModelCollection ApplicationFoldersTPGetApplicationFolderApplications(ctx, pathOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the applications in an application folder.



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
    pathOrId := "pathOrId_example" // string | Path (URL-encoded) or ID of the application folder.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    limit := int32(56) // int32 | The max number of applications returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "fields_example" // string | Optional filter, removing unspecified properties that otherwise would have been sent by the server. (optional)
    async := true // bool | If `true`, Fetch applications under application folder will be a background task. The task will have JobType GetApplicationFolderApplications (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationFoldersTPApi.ApplicationFoldersTPGetApplicationFolderApplications(context.Background(), pathOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationFoldersTPApi.ApplicationFoldersTPGetApplicationFolderApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationFoldersTPGetApplicationFolderApplications`: ApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ApplicationFoldersTPApi.ApplicationFoldersTPGetApplicationFolderApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathOrId** | **string** | Path (URL-encoded) or ID of the application folder. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationFoldersTPGetApplicationFolderApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of applications returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | Optional filter, removing unspecified properties that otherwise would have been sent by the server. | 
 **async** | **bool** | If &#x60;true&#x60;, Fetch applications under application folder will be a background task. The task will have JobType GetApplicationFolderApplications | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ApplicationResponseModelCollection**](ApplicationResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationFoldersTPGetApplicationFolders

> ApplicationFolderResponseModelCollection ApplicationFoldersTPGetApplicationFolders(ctx, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get application folders.



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
    async := true // bool | If `true`, the application folders will be fetched as a background task. The task will have JobType GetApplicationFolders. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationFoldersTPApi.ApplicationFoldersTPGetApplicationFolders(context.Background(), customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationFoldersTPApi.ApplicationFoldersTPGetApplicationFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationFoldersTPGetApplicationFolders`: ApplicationFolderResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ApplicationFoldersTPApi.ApplicationFoldersTPGetApplicationFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationFoldersTPGetApplicationFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **async** | **bool** | If &#x60;true&#x60;, the application folders will be fetched as a background task. The task will have JobType GetApplicationFolders. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ApplicationFolderResponseModelCollection**](ApplicationFolderResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationFoldersTPUpdateApplicationFolder

> ApplicationFolderResponseModel ApplicationFoldersTPUpdateApplicationFolder(ctx, pathOrId, customerid, siteid).Folder(folder).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update an application folder.



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
    pathOrId := "pathOrId_example" // string | Path or ID of the application path to update.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    folder := *openapiclient.NewEditApplicationFolderRequestModel() // EditApplicationFolderRequestModel | Details of the application folder to update.
    async := true // bool | If `true`, the application folder will be updated as a background task. The task will have JobType UpdateApplicationFolder. When the task is complete it will redirect to GetApplicationFolder. The job's Parameters will contain properties:  * _Id_ - ID of the application folder being updated. * _Path_ - Path of the application folder being updated. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationFoldersTPApi.ApplicationFoldersTPUpdateApplicationFolder(context.Background(), pathOrId, customerid, siteid).Folder(folder).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationFoldersTPApi.ApplicationFoldersTPUpdateApplicationFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationFoldersTPUpdateApplicationFolder`: ApplicationFolderResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ApplicationFoldersTPApi.ApplicationFoldersTPUpdateApplicationFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathOrId** | **string** | Path or ID of the application path to update. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationFoldersTPUpdateApplicationFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **folder** | [**EditApplicationFolderRequestModel**](EditApplicationFolderRequestModel.md) | Details of the application folder to update. | 
 **async** | **bool** | If &#x60;true&#x60;, the application folder will be updated as a background task. The task will have JobType UpdateApplicationFolder. When the task is complete it will redirect to GetApplicationFolder. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the application folder being updated. * _Path_ - Path of the application folder being updated. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ApplicationFolderResponseModel**](ApplicationFolderResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

