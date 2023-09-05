# \AdminFoldersTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminFoldersTPCreateAdminFolder**](AdminFoldersTPApi.md#AdminFoldersTPCreateAdminFolder) | **Post** /techpreview/{customerid}/{siteid}/AdminFolders | Create an admin folder
[**AdminFoldersTPDeleteAdminFolder**](AdminFoldersTPApi.md#AdminFoldersTPDeleteAdminFolder) | **Delete** /techpreview/{customerid}/{siteid}/AdminFolders/{pathOrId} | Delete an admin folder.
[**AdminFoldersTPGetAdminFolder**](AdminFoldersTPApi.md#AdminFoldersTPGetAdminFolder) | **Get** /techpreview/{customerid}/{siteid}/AdminFolders/{pathOrId} | Get details about a single admin folder.
[**AdminFoldersTPGetAdminFolderApplications**](AdminFoldersTPApi.md#AdminFoldersTPGetAdminFolderApplications) | **Get** /techpreview/{customerid}/{siteid}/AdminFolders/{pathOrId}/Applications | Get the applications in an admin folder.
[**AdminFoldersTPGetAdminFolderMachineCatalogs**](AdminFoldersTPApi.md#AdminFoldersTPGetAdminFolderMachineCatalogs) | **Get** /techpreview/{customerid}/{siteid}/AdminFolders/{pathOrId}/MachineCatalogs | Get the machineCatalogs in an admin folder.
[**AdminFoldersTPGetAdminFolders**](AdminFoldersTPApi.md#AdminFoldersTPGetAdminFolders) | **Get** /techpreview/{customerid}/{siteid}/AdminFolders | Get admin folders.
[**AdminFoldersTPUpdateAdminFolder**](AdminFoldersTPApi.md#AdminFoldersTPUpdateAdminFolder) | **Patch** /techpreview/{customerid}/{siteid}/AdminFolders/{pathOrId} | Update an admin folder.



## AdminFoldersTPCreateAdminFolder

> AdminFolderResponseModel AdminFoldersTPCreateAdminFolder(ctx, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Create an admin folder



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
    request := *openapiclient.NewCreateAdminFolderRequestModel("Name_example", []openapiclient.AdminFolderObjectIdentifier{openapiclient.AdminFolderObjectIdentifier("Unknown")}) // CreateAdminFolderRequestModel | Details of the admin folder to create.
    async := true // bool | If `true`, the admin folder will be created as a background task. The task will have JobType \"JobType.CreateAdminFolder\". When the task is complete it will redirect to \"GetAdminFolder(string)\" The job Parameters will contain properties:  * _Name_ - Name of the folder being created. * _Path_ - Path to the folder being created. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminFoldersTPApi.AdminFoldersTPCreateAdminFolder(context.Background(), customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminFoldersTPApi.AdminFoldersTPCreateAdminFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminFoldersTPCreateAdminFolder`: AdminFolderResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AdminFoldersTPApi.AdminFoldersTPCreateAdminFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminFoldersTPCreateAdminFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateAdminFolderRequestModel**](CreateAdminFolderRequestModel.md) | Details of the admin folder to create. | 
 **async** | **bool** | If &#x60;true&#x60;, the admin folder will be created as a background task. The task will have JobType \&quot;JobType.CreateAdminFolder\&quot;. When the task is complete it will redirect to \&quot;GetAdminFolder(string)\&quot; The job Parameters will contain properties:  * _Name_ - Name of the folder being created. * _Path_ - Path to the folder being created. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**AdminFolderResponseModel**](AdminFolderResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminFoldersTPDeleteAdminFolder

> AdminFoldersTPDeleteAdminFolder(ctx, pathOrId, customerid, siteid).ObjectsToRemove(objectsToRemove).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Delete an admin folder.



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
    pathOrId := "pathOrId_example" // string | Name or ID of the admin folder to delete.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    objectsToRemove := []openapiclient.AdminFolderObjects{openapiclient.AdminFolderObjects("Unknown")} // []AdminFolderObjects | Identify which object should be removed from the adminFolder. If set to null, then the admin folder will be deleted if it hosts no object and no folder metaData with keys AdminFolderObjectIdentifier.  If set to Applications, the applications in the folder will be deleted. The folder will be deleted if it do not host other objects(MachineCatalogs, or DeliveryGroup, or ApplicationGroups) and no folder metaData with keysAdminFolderObjectIdentifier.; If the folder hosts any other objects, then this folder will not be deleted, folder metadata with key ContainsApplications will be removed.  If set to MachineCatalogs and this folder still hosts some MachineCatalogs, then this folder is not allowed to be deleted; If this folder do not host any MachineCatalogs but host other objects(Applications, DeliveryGroups, ApplicationGroups), the folder metaData with key ContainsMachineCatalogs will be removed; If this folder do not host any objects and no folder metaData with keys AdminFolderObjectIdentifier, the folder will be deleted.              If set to DeliveryGroups or ApplicationGroups, have the same behavior as set to MachineCatalogs. (optional)
    async := true // bool | If `true`, the admin folder will be deleted as a background task. The task will have JobType DeleteAdminFolder. When the task is complete it will redirect to GetAdminFolders. The job's Parameters will contain properties:  * _Id_ - ID of the admin folder being deleted, * _Path_ - Path of the admin folder being deleted. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminFoldersTPApi.AdminFoldersTPDeleteAdminFolder(context.Background(), pathOrId, customerid, siteid).ObjectsToRemove(objectsToRemove).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminFoldersTPApi.AdminFoldersTPDeleteAdminFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathOrId** | **string** | Name or ID of the admin folder to delete. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminFoldersTPDeleteAdminFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **objectsToRemove** | [**[]AdminFolderObjects**](AdminFolderObjects.md) | Identify which object should be removed from the adminFolder. If set to null, then the admin folder will be deleted if it hosts no object and no folder metaData with keys AdminFolderObjectIdentifier.  If set to Applications, the applications in the folder will be deleted. The folder will be deleted if it do not host other objects(MachineCatalogs, or DeliveryGroup, or ApplicationGroups) and no folder metaData with keysAdminFolderObjectIdentifier.; If the folder hosts any other objects, then this folder will not be deleted, folder metadata with key ContainsApplications will be removed.  If set to MachineCatalogs and this folder still hosts some MachineCatalogs, then this folder is not allowed to be deleted; If this folder do not host any MachineCatalogs but host other objects(Applications, DeliveryGroups, ApplicationGroups), the folder metaData with key ContainsMachineCatalogs will be removed; If this folder do not host any objects and no folder metaData with keys AdminFolderObjectIdentifier, the folder will be deleted.              If set to DeliveryGroups or ApplicationGroups, have the same behavior as set to MachineCatalogs. | 
 **async** | **bool** | If &#x60;true&#x60;, the admin folder will be deleted as a background task. The task will have JobType DeleteAdminFolder. When the task is complete it will redirect to GetAdminFolders. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the admin folder being deleted, * _Path_ - Path of the admin folder being deleted. | [default to false]
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


## AdminFoldersTPGetAdminFolder

> AdminFolderResponseModel AdminFoldersTPGetAdminFolder(ctx, pathOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get details about a single admin folder.



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
    pathOrId := "pathOrId_example" // string | Path (URL-encoded) or ID of the admin folder.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminFoldersTPApi.AdminFoldersTPGetAdminFolder(context.Background(), pathOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminFoldersTPApi.AdminFoldersTPGetAdminFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminFoldersTPGetAdminFolder`: AdminFolderResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AdminFoldersTPApi.AdminFoldersTPGetAdminFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathOrId** | **string** | Path (URL-encoded) or ID of the admin folder. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminFoldersTPGetAdminFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**AdminFolderResponseModel**](AdminFolderResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminFoldersTPGetAdminFolderApplications

> ApplicationResponseModelCollection AdminFoldersTPGetAdminFolderApplications(ctx, pathOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the applications in an admin folder.



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
    pathOrId := "pathOrId_example" // string | Path (URL-encoded) or ID of the admin folder.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    limit := int32(56) // int32 | The max number of applications returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "fields_example" // string | Optional filter, removing unspecified properties that otherwise would have been sent by the server. (optional)
    async := true // bool | If `true`, Fetch applications under admin folder will be a background task. The task will have JobType GetAdminFolderApplications (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminFoldersTPApi.AdminFoldersTPGetAdminFolderApplications(context.Background(), pathOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminFoldersTPApi.AdminFoldersTPGetAdminFolderApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminFoldersTPGetAdminFolderApplications`: ApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AdminFoldersTPApi.AdminFoldersTPGetAdminFolderApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathOrId** | **string** | Path (URL-encoded) or ID of the admin folder. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminFoldersTPGetAdminFolderApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of applications returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | Optional filter, removing unspecified properties that otherwise would have been sent by the server. | 
 **async** | **bool** | If &#x60;true&#x60;, Fetch applications under admin folder will be a background task. The task will have JobType GetAdminFolderApplications | [default to false]
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


## AdminFoldersTPGetAdminFolderMachineCatalogs

> MachineCatalogResponseModelCollection AdminFoldersTPGetAdminFolderMachineCatalogs(ctx, pathOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the machineCatalogs in an admin folder.



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
    pathOrId := "pathOrId_example" // string | Path (URL-encoded) or ID of the admin folder.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    limit := int32(56) // int32 | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional) (default to 2147483647)
    continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)
    fields := "fields_example" // string | Optional filter, removing unspecified properties that otherwise would have been sent by the server. (optional)
    async := true // bool | If `true`, Fetch machineCatalogs under admin folder will be a background task. The task will have JobType GetAdminFolderMachineCatalogs (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminFoldersTPApi.AdminFoldersTPGetAdminFolderMachineCatalogs(context.Background(), pathOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminFoldersTPApi.AdminFoldersTPGetAdminFolderMachineCatalogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminFoldersTPGetAdminFolderMachineCatalogs`: MachineCatalogResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AdminFoldersTPApi.AdminFoldersTPGetAdminFolderMachineCatalogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathOrId** | **string** | Path (URL-encoded) or ID of the admin folder. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminFoldersTPGetAdminFolderMachineCatalogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | [default to 2147483647]
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 
 **fields** | **string** | Optional filter, removing unspecified properties that otherwise would have been sent by the server. | 
 **async** | **bool** | If &#x60;true&#x60;, Fetch machineCatalogs under admin folder will be a background task. The task will have JobType GetAdminFolderMachineCatalogs | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**MachineCatalogResponseModelCollection**](MachineCatalogResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminFoldersTPGetAdminFolders

> AdminFolderResponseModelCollection AdminFoldersTPGetAdminFolders(ctx, customerid, siteid).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get admin folders.



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
    async := true // bool | If `true`, the admin folders will be fetched as a background task. The task will have JobType GetAdminFolders. (optional) (default to false)
    limit := int32(56) // int32 | The max number of admin folders returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminFoldersTPApi.AdminFoldersTPGetAdminFolders(context.Background(), customerid, siteid).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminFoldersTPApi.AdminFoldersTPGetAdminFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminFoldersTPGetAdminFolders`: AdminFolderResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `AdminFoldersTPApi.AdminFoldersTPGetAdminFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminFoldersTPGetAdminFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **async** | **bool** | If &#x60;true&#x60;, the admin folders will be fetched as a background task. The task will have JobType GetAdminFolders. | [default to false]
 **limit** | **int32** | The max number of admin folders returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**AdminFolderResponseModelCollection**](AdminFolderResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminFoldersTPUpdateAdminFolder

> AdminFolderResponseModel AdminFoldersTPUpdateAdminFolder(ctx, pathOrId, customerid, siteid).Folder(folder).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update an admin folder.



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
    pathOrId := "pathOrId_example" // string | Path or ID of the admin folder path to update.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    folder := *openapiclient.NewEditAdminFolderRequestModel() // EditAdminFolderRequestModel | Details of the admin folder to update.
    async := true // bool | If `true`, the admin folder will be updated as a background task. The task will have JobType UpdateAdminFolder. When the task is complete it will redirect to GetAdminFolder. The job's Parameters will contain properties:  * _Id_ - ID of the admin folder being updated. * _Path_ - Path of the admin folder being updated. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminFoldersTPApi.AdminFoldersTPUpdateAdminFolder(context.Background(), pathOrId, customerid, siteid).Folder(folder).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminFoldersTPApi.AdminFoldersTPUpdateAdminFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminFoldersTPUpdateAdminFolder`: AdminFolderResponseModel
    fmt.Fprintf(os.Stdout, "Response from `AdminFoldersTPApi.AdminFoldersTPUpdateAdminFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathOrId** | **string** | Path or ID of the admin folder path to update. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminFoldersTPUpdateAdminFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **folder** | [**EditAdminFolderRequestModel**](EditAdminFolderRequestModel.md) | Details of the admin folder to update. | 
 **async** | **bool** | If &#x60;true&#x60;, the admin folder will be updated as a background task. The task will have JobType UpdateAdminFolder. When the task is complete it will redirect to GetAdminFolder. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the admin folder being updated. * _Path_ - Path of the admin folder being updated. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**AdminFolderResponseModel**](AdminFolderResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

