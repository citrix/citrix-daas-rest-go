# \ApplicationsTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationsTPAddApplicationTags**](ApplicationsTPApi.md#ApplicationsTPAddApplicationTags) | **Post** /techpreview/{customerid}/{siteid}/Applications/{nameOrId}/Tags/{tagNameOrId} | Add a tag to an application.
[**ApplicationsTPAddApplications**](ApplicationsTPApi.md#ApplicationsTPAddApplications) | **Post** /techpreview/{customerid}/{siteid}/Applications | Add applications to the site.
[**ApplicationsTPCheckApplicationExists**](ApplicationsTPApi.md#ApplicationsTPCheckApplicationExists) | **Head** /techpreview/{customerid}/{siteid}/Applications/{name} | Check for the existence of an application by name
[**ApplicationsTPDeleteApplication**](ApplicationsTPApi.md#ApplicationsTPDeleteApplication) | **Delete** /techpreview/{customerid}/{siteid}/Applications/{nameOrId} | Delete an application!
[**ApplicationsTPDisableApplicationFta**](ApplicationsTPApi.md#ApplicationsTPDisableApplicationFta) | **Delete** /techpreview/{customerid}/{siteid}/Applications/{nameOrId}/FileTypes/{extensionName} | Disable a file-type association for an application.
[**ApplicationsTPDiscoverApplications**](ApplicationsTPApi.md#ApplicationsTPDiscoverApplications) | **Post** /techpreview/{customerid}/{siteid}/Applications/$discover | Search path for applications to add manually.
[**ApplicationsTPDoApplicationSearch**](ApplicationsTPApi.md#ApplicationsTPDoApplicationSearch) | **Post** /techpreview/{customerid}/{siteid}/Applications/$search | Perform an advanced search for applications.
[**ApplicationsTPDuplicateApplication**](ApplicationsTPApi.md#ApplicationsTPDuplicateApplication) | **Post** /techpreview/{customerid}/{siteid}/Applications/{nameOrId}/$duplicate | Duplicate an application.
[**ApplicationsTPEnableApplicationFta**](ApplicationsTPApi.md#ApplicationsTPEnableApplicationFta) | **Post** /techpreview/{customerid}/{siteid}/Applications/{nameOrId}/FileTypes | Enable a file-type association for an application.
[**ApplicationsTPGetApplication**](ApplicationsTPApi.md#ApplicationsTPGetApplication) | **Get** /techpreview/{customerid}/{siteid}/Applications/{nameOrId} | Get details of a single application.
[**ApplicationsTPGetApplicationAdministrators**](ApplicationsTPApi.md#ApplicationsTPGetApplicationAdministrators) | **Get** /techpreview/{customerid}/{siteid}/Applications/{nameOrId}/Administrators | Get administrators authorized to administer an application.
[**ApplicationsTPGetApplicationApplicationGroups**](ApplicationsTPApi.md#ApplicationsTPGetApplicationApplicationGroups) | **Get** /techpreview/{customerid}/{siteid}/Applications/{nameOrId}/ApplicationGroups | Get application groups associated with an application.
[**ApplicationsTPGetApplicationDeliveryGroups**](ApplicationsTPApi.md#ApplicationsTPGetApplicationDeliveryGroups) | **Get** /techpreview/{customerid}/{siteid}/Applications/{nameOrId}/DeliveryGroups | Get delivery groups associated with an application.
[**ApplicationsTPGetApplicationFta**](ApplicationsTPApi.md#ApplicationsTPGetApplicationFta) | **Get** /techpreview/{customerid}/{siteid}/Applications/{nameOrId}/FileTypes/{extensionName} | Get a single file-type for an application.
[**ApplicationsTPGetApplicationFtas**](ApplicationsTPApi.md#ApplicationsTPGetApplicationFtas) | **Get** /techpreview/{customerid}/{siteid}/Applications/{nameOrId}/FileTypes | Get all file-types for an application.
[**ApplicationsTPGetApplicationRecommendName**](ApplicationsTPApi.md#ApplicationsTPGetApplicationRecommendName) | **Get** /techpreview/{customerid}/{siteid}/Applications/{nameOrId}/RecommendName | Get an application&#39;s recommend name.
[**ApplicationsTPGetApplicationSessions**](ApplicationsTPApi.md#ApplicationsTPGetApplicationSessions) | **Get** /techpreview/{customerid}/{siteid}/Applications/{nameOrId}/Sessions | Get sessions in which the application is currently running.
[**ApplicationsTPGetApplicationTags**](ApplicationsTPApi.md#ApplicationsTPGetApplicationTags) | **Get** /techpreview/{customerid}/{siteid}/Applications/{nameOrId}/Tags | Get tags associated with an application.
[**ApplicationsTPGetApplications**](ApplicationsTPApi.md#ApplicationsTPGetApplications) | **Get** /techpreview/{customerid}/{siteid}/Applications | Get all applications.
[**ApplicationsTPImportFileTypes**](ApplicationsTPApi.md#ApplicationsTPImportFileTypes) | **Post** /techpreview/{customerid}/{siteid}/Applications/{nameOrId}/FileTypes/$import | Import file types for the application.
[**ApplicationsTPPatchApplication**](ApplicationsTPApi.md#ApplicationsTPPatchApplication) | **Patch** /techpreview/{customerid}/{siteid}/Applications/{nameOrId} | Modify an application.
[**ApplicationsTPRefreshAppVApplications**](ApplicationsTPApi.md#ApplicationsTPRefreshAppVApplications) | **Get** /techpreview/{customerid}/{siteid}/Applications/RefreshAppVApplications | Refresh the AppV Applications.
[**ApplicationsTPRemoveApplicationTags**](ApplicationsTPApi.md#ApplicationsTPRemoveApplicationTags) | **Delete** /techpreview/{customerid}/{siteid}/Applications/{nameOrId}/Tags/{tagNameOrId} | Remove a tag from an application.
[**ApplicationsTPSearchStartMenuApplications**](ApplicationsTPApi.md#ApplicationsTPSearchStartMenuApplications) | **Post** /techpreview/{customerid}/{siteid}/Applications/StartMenuApplications | Search Start Menu Applications.
[**ApplicationsTPSetApplicationTags**](ApplicationsTPApi.md#ApplicationsTPSetApplicationTags) | **Put** /techpreview/{customerid}/{siteid}/Applications/{nameOrId}/Tags | Set tags associated with an application.



## ApplicationsTPAddApplicationTags

> ApplicationResponseModelCollection ApplicationsTPAddApplicationTags(ctx, nameOrId, tagNameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Add a tag to an application.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName.
    tagNameOrId := "tagNameOrId_example" // string | Name or ID of the tag to add.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the tags will be modified as a background task. The task will have JobType SetTagApplications. When the task is complete it will redirect to GetApplicationTags. The job's Parameters will contain properties: * _Id_ - ID of the application being tagged, * _Name_ - Name of the application being tagged. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsTPApi.ApplicationsTPAddApplicationTags(context.Background(), nameOrId, tagNameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPAddApplicationTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsTPAddApplicationTags`: ApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsTPApi.ApplicationsTPAddApplicationTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName. | 
**tagNameOrId** | **string** | Name or ID of the tag to add. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPAddApplicationTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. The task will have JobType SetTagApplications. When the task is complete it will redirect to GetApplicationTags. The job&#39;s Parameters will contain properties: * _Id_ - ID of the application being tagged, * _Name_ - Name of the application being tagged. | [default to false]
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


## ApplicationsTPAddApplications

> ApplicationsTPAddApplications(ctx, customerid, siteid).Apps(apps).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Add applications to the site.



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
    apps := *openapiclient.NewAddApplicationsRequestModel() // AddApplicationsRequestModel | Details of the applications to add.
    async := true // bool | If `true`, the applications (and associated objects) will be added as a background task. The task will have JobType AddApplications. When the task is complete it will redirect to GetApplications. The job's Parameters will contain properties: * _NumApplications_ - Number of applications being added. * _Name0_ - Name of the first application being added. * _Name1_ - Name of the second application being added. * _Name{n}_ - ...And so on. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationsTPApi.ApplicationsTPAddApplications(context.Background(), customerid, siteid).Apps(apps).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPAddApplications``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApplicationsTPAddApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apps** | [**AddApplicationsRequestModel**](AddApplicationsRequestModel.md) | Details of the applications to add. | 
 **async** | **bool** | If &#x60;true&#x60;, the applications (and associated objects) will be added as a background task. The task will have JobType AddApplications. When the task is complete it will redirect to GetApplications. The job&#39;s Parameters will contain properties: * _NumApplications_ - Number of applications being added. * _Name0_ - Name of the first application being added. * _Name1_ - Name of the second application being added. * _Name{n}_ - ...And so on. | [default to false]
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


## ApplicationsTPCheckApplicationExists

> ApplicationsTPCheckApplicationExists(ctx, name, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Check for the existence of an application by name



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
    name := "name_example" // string | Name or ID of the application to check for. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationsTPApi.ApplicationsTPCheckApplicationExists(context.Background(), name, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPCheckApplicationExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name or ID of the application to check for. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPCheckApplicationExistsRequest struct via the builder pattern


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


## ApplicationsTPDeleteApplication

> ApplicationsTPDeleteApplication(ctx, nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Delete an application!



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application to delete. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the application (and associated objects) will be deleted as a background task. The task will have JobType DeleteApplication. When the task is complete it will redirect to GetApplications. The job's Parameters will contain properties:  * _Id_ - ID of the application being deleted, * _Name_ - Name of the application being deleted. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationsTPApi.ApplicationsTPDeleteApplication(context.Background(), nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPDeleteApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application to delete. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPDeleteApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, the application (and associated objects) will be deleted as a background task. The task will have JobType DeleteApplication. When the task is complete it will redirect to GetApplications. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the application being deleted, * _Name_ - Name of the application being deleted. | [default to false]
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


## ApplicationsTPDisableApplicationFta

> ApplicationsTPDisableApplicationFta(ctx, nameOrId, extensionName, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Disable a file-type association for an application.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName.
    extensionName := "extensionName_example" // string | File-type extension name to disable.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the file type will be disabled as a background task. The task will have JobType ModifyFTA. When the task is complete it will redirect to GetApplicationFta. The job's Parameters will contain properties: * _Id_ - ID of the application whose FTA is being modified, * _Name_ - Name of the application whose FTA is being modified. * _ExtensionName_ - File type extension name being modified. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationsTPApi.ApplicationsTPDisableApplicationFta(context.Background(), nameOrId, extensionName, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPDisableApplicationFta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName. | 
**extensionName** | **string** | File-type extension name to disable. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPDisableApplicationFtaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **async** | **bool** | If &#x60;true&#x60;, the file type will be disabled as a background task. The task will have JobType ModifyFTA. When the task is complete it will redirect to GetApplicationFta. The job&#39;s Parameters will contain properties: * _Id_ - ID of the application whose FTA is being modified, * _Name_ - Name of the application whose FTA is being modified. * _ExtensionName_ - File type extension name being modified. | [default to false]
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


## ApplicationsTPDiscoverApplications

> ApplicationDiscoveryResponseModel ApplicationsTPDiscoverApplications(ctx, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Search path for applications to add manually.

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
    request := *openapiclient.NewApplicationDiscoveryRequestModel() // ApplicationDiscoveryRequestModel | Search request for applications discovery. ApplicationDiscoveryRequestModel
    async := true // bool | If `true`, the search operation will run as a background task. The task will have JobType ApplicationsDiscovery (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsTPApi.ApplicationsTPDiscoverApplications(context.Background(), customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPDiscoverApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsTPDiscoverApplications`: ApplicationDiscoveryResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsTPApi.ApplicationsTPDiscoverApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPDiscoverApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**ApplicationDiscoveryRequestModel**](ApplicationDiscoveryRequestModel.md) | Search request for applications discovery. ApplicationDiscoveryRequestModel | 
 **async** | **bool** | If &#x60;true&#x60;, the search operation will run as a background task. The task will have JobType ApplicationsDiscovery | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ApplicationDiscoveryResponseModel**](ApplicationDiscoveryResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsTPDoApplicationSearch

> ApplicationResponseModelCollection ApplicationsTPDoApplicationSearch(ctx, customerid, siteid).SearchRequest(searchRequest).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Perform an advanced search for applications.



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
    searchRequest := *openapiclient.NewApplicationSearchRequestModel() // ApplicationSearchRequestModel | Specifies the advanced search parameters.
    limit := int32(56) // int32 | The max number of applications returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "fields_example" // string | Optional filter, removing unspecified properties that otherwise would have been sent by the server. (optional)
    async := true // bool | If `true`, the applications (and associated objects) will be added as a background task. The task will have JobType SearchApplications. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsTPApi.ApplicationsTPDoApplicationSearch(context.Background(), customerid, siteid).SearchRequest(searchRequest).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPDoApplicationSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsTPDoApplicationSearch`: ApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsTPApi.ApplicationsTPDoApplicationSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPDoApplicationSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **searchRequest** | [**ApplicationSearchRequestModel**](ApplicationSearchRequestModel.md) | Specifies the advanced search parameters. | 
 **limit** | **int32** | The max number of applications returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | Optional filter, removing unspecified properties that otherwise would have been sent by the server. | 
 **async** | **bool** | If &#x60;true&#x60;, the applications (and associated objects) will be added as a background task. The task will have JobType SearchApplications. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ApplicationResponseModelCollection**](ApplicationResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsTPDuplicateApplication

> ApplicationDetailResponseModel ApplicationsTPDuplicateApplication(ctx, nameOrId, customerid, siteid).Options(options).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Duplicate an application.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application to duplicate. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    options := *openapiclient.NewDuplicateApplicationRequestModel() // DuplicateApplicationRequestModel | Optional properties controlling how the new application is created.
    async := true // bool | If `true`, the application (and associated objects) will be duplicated as a background task. The task will have JobType AddApplications. When the task is complete it will redirect to GetApplications. The job's Parameters will contain properties: * _NumApplications_ - Number of applications being added.  For a duplicate operation, this will always be 1. * _Name0_ - Name of the first application being added.  For a duplicate operation, this will be the name of the new duplicate of the original application. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsTPApi.ApplicationsTPDuplicateApplication(context.Background(), nameOrId, customerid, siteid).Options(options).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPDuplicateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsTPDuplicateApplication`: ApplicationDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsTPApi.ApplicationsTPDuplicateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application to duplicate. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPDuplicateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **options** | [**DuplicateApplicationRequestModel**](DuplicateApplicationRequestModel.md) | Optional properties controlling how the new application is created. | 
 **async** | **bool** | If &#x60;true&#x60;, the application (and associated objects) will be duplicated as a background task. The task will have JobType AddApplications. When the task is complete it will redirect to GetApplications. The job&#39;s Parameters will contain properties: * _NumApplications_ - Number of applications being added.  For a duplicate operation, this will always be 1. * _Name0_ - Name of the first application being added.  For a duplicate operation, this will be the name of the new duplicate of the original application. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ApplicationDetailResponseModel**](ApplicationDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsTPEnableApplicationFta

> ApplicationsTPEnableApplicationFta(ctx, nameOrId, customerid, siteid).Fta(fta).SkipVdaChecking(skipVdaChecking).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Enable a file-type association for an application.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    fta := *openapiclient.NewFtaRequestModel("ExtensionName_example") // FtaRequestModel | Details of the file-type association to enable.
    skipVdaChecking := true // bool | If true, don't check the status of VDAs before             enabling the file-type association for the application. (optional) (default to false)
    async := true // bool | If `true`, the file type will be enabled as a background task. The task will have JobType ModifyFTA. When the task is complete it will redirect to GetApplicationFta. The job's Parameters will contain properties: * _Id_ - ID of the application whose FTA is being modified, * _Name_ - Name of the application whose FTA is being modified. * _ExtensionName_ - File type extension name being modified. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationsTPApi.ApplicationsTPEnableApplicationFta(context.Background(), nameOrId, customerid, siteid).Fta(fta).SkipVdaChecking(skipVdaChecking).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPEnableApplicationFta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPEnableApplicationFtaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fta** | [**FtaRequestModel**](FtaRequestModel.md) | Details of the file-type association to enable. | 
 **skipVdaChecking** | **bool** | If true, don&#39;t check the status of VDAs before             enabling the file-type association for the application. | [default to false]
 **async** | **bool** | If &#x60;true&#x60;, the file type will be enabled as a background task. The task will have JobType ModifyFTA. When the task is complete it will redirect to GetApplicationFta. The job&#39;s Parameters will contain properties: * _Id_ - ID of the application whose FTA is being modified, * _Name_ - Name of the application whose FTA is being modified. * _ExtensionName_ - File type extension name being modified. | [default to false]
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


## ApplicationsTPGetApplication

> ApplicationDetailResponseModel ApplicationsTPGetApplication(ctx, nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get details of a single application.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsTPApi.ApplicationsTPGetApplication(context.Background(), nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPGetApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsTPGetApplication`: ApplicationDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsTPApi.ApplicationsTPGetApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ApplicationDetailResponseModel**](ApplicationDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsTPGetApplicationAdministrators

> AdministratorResponseModelCollection ApplicationsTPGetApplicationAdministrators(ctx, nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get administrators authorized to administer an application.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    limit := int32(56) // int32 | The max number of administrators returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsTPApi.ApplicationsTPGetApplicationAdministrators(context.Background(), nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPGetApplicationAdministrators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsTPGetApplicationAdministrators`: AdministratorResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsTPApi.ApplicationsTPGetApplicationAdministrators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPGetApplicationAdministratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of administrators returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**AdministratorResponseModelCollection**](AdministratorResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsTPGetApplicationApplicationGroups

> ApplicationGroupResponseModelCollection ApplicationsTPGetApplicationApplicationGroups(ctx, nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get application groups associated with an application.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    limit := int32(56) // int32 | The max number of application groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    async := true // bool | If `true`, the application groups associated with the application will be fetched as a background task. The task will have JobType GetApplicationApplicationGroups. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsTPApi.ApplicationsTPGetApplicationApplicationGroups(context.Background(), nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPGetApplicationApplicationGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsTPGetApplicationApplicationGroups`: ApplicationGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsTPApi.ApplicationsTPGetApplicationApplicationGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPGetApplicationApplicationGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of application groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **async** | **bool** | If &#x60;true&#x60;, the application groups associated with the application will be fetched as a background task. The task will have JobType GetApplicationApplicationGroups. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ApplicationGroupResponseModelCollection**](ApplicationGroupResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsTPGetApplicationDeliveryGroups

> ApplicationDeliveryGroupResponseModelCollection ApplicationsTPGetApplicationDeliveryGroups(ctx, nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get delivery groups associated with an application.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    limit := int32(56) // int32 | The max number of delivery groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "fields_example" // string | Optional filter, removing unspecified properties that otherwise would have been sent by the server (optional)
    async := true // bool | If `true`, the delivery groups associated with the application will be fetched as a background task. The task will have JobType GetApplicationDeliveryGroups. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsTPApi.ApplicationsTPGetApplicationDeliveryGroups(context.Background(), nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPGetApplicationDeliveryGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsTPGetApplicationDeliveryGroups`: ApplicationDeliveryGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsTPApi.ApplicationsTPGetApplicationDeliveryGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPGetApplicationDeliveryGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of delivery groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | Optional filter, removing unspecified properties that otherwise would have been sent by the server | 
 **async** | **bool** | If &#x60;true&#x60;, the delivery groups associated with the application will be fetched as a background task. The task will have JobType GetApplicationDeliveryGroups. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ApplicationDeliveryGroupResponseModelCollection**](ApplicationDeliveryGroupResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsTPGetApplicationFta

> FtaResponseModel ApplicationsTPGetApplicationFta(ctx, nameOrId, extensionName, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get a single file-type for an application.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName.
    extensionName := "extensionName_example" // string | Extension name of the file-type.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsTPApi.ApplicationsTPGetApplicationFta(context.Background(), nameOrId, extensionName, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPGetApplicationFta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsTPGetApplicationFta`: FtaResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsTPApi.ApplicationsTPGetApplicationFta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName. | 
**extensionName** | **string** | Extension name of the file-type. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPGetApplicationFtaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**FtaResponseModel**](FtaResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsTPGetApplicationFtas

> FtaResponseModelCollection ApplicationsTPGetApplicationFtas(ctx, nameOrId, customerid, siteid).SkipVdaChecking(skipVdaChecking).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get all file-types for an application.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    skipVdaChecking := true // bool | If true, don't check the status of VDAs before             getting the file-type association for the application. (optional) (default to false)
    async := true // bool | If 'true', the file types will be gotten as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsTPApi.ApplicationsTPGetApplicationFtas(context.Background(), nameOrId, customerid, siteid).SkipVdaChecking(skipVdaChecking).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPGetApplicationFtas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsTPGetApplicationFtas`: FtaResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsTPApi.ApplicationsTPGetApplicationFtas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPGetApplicationFtasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **skipVdaChecking** | **bool** | If true, don&#39;t check the status of VDAs before             getting the file-type association for the application. | [default to false]
 **async** | **bool** | If &#39;true&#39;, the file types will be gotten as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**FtaResponseModelCollection**](FtaResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsTPGetApplicationRecommendName

> string ApplicationsTPGetApplicationRecommendName(ctx, nameOrId, customerid, siteid).ApplicationFolder(applicationFolder).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get an application's recommend name.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    applicationFolder := "applicationFolder_example" // string | Name or ID of the application Folder (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsTPApi.ApplicationsTPGetApplicationRecommendName(context.Background(), nameOrId, customerid, siteid).ApplicationFolder(applicationFolder).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPGetApplicationRecommendName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsTPGetApplicationRecommendName`: string
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsTPApi.ApplicationsTPGetApplicationRecommendName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPGetApplicationRecommendNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applicationFolder** | **string** | Name or ID of the application Folder | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsTPGetApplicationSessions

> SessionResponseModelCollection ApplicationsTPGetApplicationSessions(ctx, nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get sessions in which the application is currently running.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    limit := int32(56) // int32 | The max number of sessions returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    async := true // bool | If `true`, the application sessions will be fetched as a background task. The task will have JobType GetApplicationSessions. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsTPApi.ApplicationsTPGetApplicationSessions(context.Background(), nameOrId, customerid, siteid).Limit(limit).ContinuationToken(continuationToken).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPGetApplicationSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsTPGetApplicationSessions`: SessionResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsTPApi.ApplicationsTPGetApplicationSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPGetApplicationSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of sessions returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **async** | **bool** | If &#x60;true&#x60;, the application sessions will be fetched as a background task. The task will have JobType GetApplicationSessions. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**SessionResponseModelCollection**](SessionResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsTPGetApplicationTags

> TagResponseModelCollection ApplicationsTPGetApplicationTags(ctx, nameOrId, customerid, siteid).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get tags associated with an application.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    fields := "fields_example" // string | field to filter response model. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsTPApi.ApplicationsTPGetApplicationTags(context.Background(), nameOrId, customerid, siteid).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPGetApplicationTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsTPGetApplicationTags`: TagResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsTPApi.ApplicationsTPGetApplicationTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPGetApplicationTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | **string** | field to filter response model. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**TagResponseModelCollection**](TagResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsTPGetApplications

> ApplicationResponseModelCollection ApplicationsTPGetApplications(ctx, customerid, siteid).ApplicationFolder(applicationFolder).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get all applications.



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
    applicationFolder := "applicationFolder_example" // string | Optional folder path (URL-encoded) or ID.  If not specified, all applications will be returned from all folders. (optional)
    limit := int32(56) // int32 | The max number of applications returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "fields_example" // string | Optional filter, removing unspecified properties that otherwise would have been sent by the server. (optional)
    async := true // bool | If `true`, the applications will be fetched as a background task. The task will have JobType GetApplications. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsTPApi.ApplicationsTPGetApplications(context.Background(), customerid, siteid).ApplicationFolder(applicationFolder).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPGetApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsTPGetApplications`: ApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsTPApi.ApplicationsTPGetApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPGetApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationFolder** | **string** | Optional folder path (URL-encoded) or ID.  If not specified, all applications will be returned from all folders. | 
 **limit** | **int32** | The max number of applications returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | Optional filter, removing unspecified properties that otherwise would have been sent by the server. | 
 **async** | **bool** | If &#x60;true&#x60;, the applications will be fetched as a background task. The task will have JobType GetApplications. | [default to false]
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


## ApplicationsTPImportFileTypes

> ApplicationsTPImportFileTypes(ctx, nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Import file types for the application.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the file types will be imported as a background task. This is recommended as this operation may cause a power action, turning on a machine in order to gather the data. This may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType ImportFTA. When the task is complete it will redirect to GetApplicationFtas. The job's Parameters will contain properties: * _Id_ - ID of the application whose FTA is being imported, * _Name_ - Name of the application whose FTA is being imported. * _MachineId_ - ID of the machine selected, from which the file types are being imported; will be present in Parameters only after a machine is selected. * _MachineName_ - Name of the machine selected, from which the file types are being imported; will be present in Parameters only after a machine is selected. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationsTPApi.ApplicationsTPImportFileTypes(context.Background(), nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPImportFileTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPImportFileTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, the file types will be imported as a background task. This is recommended as this operation may cause a power action, turning on a machine in order to gather the data. This may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType ImportFTA. When the task is complete it will redirect to GetApplicationFtas. The job&#39;s Parameters will contain properties: * _Id_ - ID of the application whose FTA is being imported, * _Name_ - Name of the application whose FTA is being imported. * _MachineId_ - ID of the machine selected, from which the file types are being imported; will be present in Parameters only after a machine is selected. * _MachineName_ - Name of the machine selected, from which the file types are being imported; will be present in Parameters only after a machine is selected. | [default to false]
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


## ApplicationsTPPatchApplication

> ApplicationsTPPatchApplication(ctx, nameOrId, customerid, siteid).App(app).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Modify an application.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    app := *openapiclient.NewEditApplicationRequestModel() // EditApplicationRequestModel | Properties of the application to modify.
    async := true // bool | If `true`, the application will be updated as a background task. The task will have JobType UpdateApplication. When the task is complete it will redirect to GetApplication. The job's Parameters will contain properties: * _Id_ - ID of the application being updated, * _Name_ - Name of the application being updated. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationsTPApi.ApplicationsTPPatchApplication(context.Background(), nameOrId, customerid, siteid).App(app).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPPatchApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPPatchApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **app** | [**EditApplicationRequestModel**](EditApplicationRequestModel.md) | Properties of the application to modify. | 
 **async** | **bool** | If &#x60;true&#x60;, the application will be updated as a background task. The task will have JobType UpdateApplication. When the task is complete it will redirect to GetApplication. The job&#39;s Parameters will contain properties: * _Id_ - ID of the application being updated, * _Name_ - Name of the application being updated. | [default to false]
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


## ApplicationsTPRefreshAppVApplications

> ApplicationResponseModelCollection ApplicationsTPRefreshAppVApplications(ctx, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Refresh the AppV Applications.

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
    async := true // bool | If `true`, the refresh operation will run as a background task. The task will have JobType RefreshAppVApplications. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsTPApi.ApplicationsTPRefreshAppVApplications(context.Background(), customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPRefreshAppVApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsTPRefreshAppVApplications`: ApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsTPApi.ApplicationsTPRefreshAppVApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPRefreshAppVApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **async** | **bool** | If &#x60;true&#x60;, the refresh operation will run as a background task. The task will have JobType RefreshAppVApplications. | [default to false]
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


## ApplicationsTPRemoveApplicationTags

> ApplicationsTPRemoveApplicationTags(ctx, nameOrId, tagNameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Remove a tag from an application.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName.
    tagNameOrId := "tagNameOrId_example" // string | Name or ID of the tag to remove.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the tags will be modified as a background task. The task will have JobType SetTagApplications. When the task is complete it will redirect to GetApplicationTags. The job's Parameters will contain properties: * _Id_ - ID of the application being tagged, * _Name_ - Name of the application being tagged. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationsTPApi.ApplicationsTPRemoveApplicationTags(context.Background(), nameOrId, tagNameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPRemoveApplicationTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName. | 
**tagNameOrId** | **string** | Name or ID of the tag to remove. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPRemoveApplicationTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. The task will have JobType SetTagApplications. When the task is complete it will redirect to GetApplicationTags. The job&#39;s Parameters will contain properties: * _Id_ - ID of the application being tagged, * _Name_ - Name of the application being tagged. | [default to false]
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


## ApplicationsTPSearchStartMenuApplications

> StartMenuApplicationResponseModelCollection ApplicationsTPSearchStartMenuApplications(ctx, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Search Start Menu Applications.

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
    request := *openapiclient.NewSearchStartMenuApplicationsRequestModel() // SearchStartMenuApplicationsRequestModel | Search request for querying start menu applications. SearchStartMenuApplicationsRequestModel
    async := true // bool | If `true`, the search operation will run as a background task. The task will have JobType SearchStartMenuApplications (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsTPApi.ApplicationsTPSearchStartMenuApplications(context.Background(), customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPSearchStartMenuApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsTPSearchStartMenuApplications`: StartMenuApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsTPApi.ApplicationsTPSearchStartMenuApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPSearchStartMenuApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**SearchStartMenuApplicationsRequestModel**](SearchStartMenuApplicationsRequestModel.md) | Search request for querying start menu applications. SearchStartMenuApplicationsRequestModel | 
 **async** | **bool** | If &#x60;true&#x60;, the search operation will run as a background task. The task will have JobType SearchStartMenuApplications | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**StartMenuApplicationResponseModelCollection**](StartMenuApplicationResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsTPSetApplicationTags

> ApplicationsTPSetApplicationTags(ctx, nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Set tags associated with an application.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewTagsRequestModel() // TagsRequestModel | Details of which tags to set.
    async := true // bool | If `true`, the tags will be modified as a background task. The task will have JobType SetTagApplications. When the task is complete it will redirect to GetApplicationTags. The job's Parameters will contain properties: * _Id_ - ID of the application being tagged, * _Name_ - Name of the application being tagged. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationsTPApi.ApplicationsTPSetApplicationTags(context.Background(), nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTPApi.ApplicationsTPSetApplicationTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsTPSetApplicationTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**TagsRequestModel**](TagsRequestModel.md) | Details of which tags to set. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. The task will have JobType SetTagApplications. When the task is complete it will redirect to GetApplicationTags. The job&#39;s Parameters will contain properties: * _Id_ - ID of the application being tagged, * _Name_ - Name of the application being tagged. | [default to false]
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

