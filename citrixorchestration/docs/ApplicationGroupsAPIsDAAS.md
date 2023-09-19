# \ApplicationGroupsAPIsDAAS

All URIs are relative to *https://api-us.cloud.com/cvad/manage*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationGroupsAddApplicationGroupTag**](ApplicationGroupsAPIsDAAS.md#ApplicationGroupsAddApplicationGroupTag) | **Post** /ApplicationGroups/{nameOrId}/Tags/{tagNameOrId} | Add a tag to an application group.
[**ApplicationGroupsAddApplications**](ApplicationGroupsAPIsDAAS.md#ApplicationGroupsAddApplications) | **Post** /ApplicationGroups/{nameOrId}/Applications | Add applications to the application group.
[**ApplicationGroupsCheckApplicationGroupExists**](ApplicationGroupsAPIsDAAS.md#ApplicationGroupsCheckApplicationGroupExists) | **Head** /ApplicationGroups/{name} | Check for the existence of an application group by name.
[**ApplicationGroupsCreateApplicationGroup**](ApplicationGroupsAPIsDAAS.md#ApplicationGroupsCreateApplicationGroup) | **Post** /ApplicationGroups | Create an application group.
[**ApplicationGroupsDeleteApplicationGroup**](ApplicationGroupsAPIsDAAS.md#ApplicationGroupsDeleteApplicationGroup) | **Delete** /ApplicationGroups/{nameOrId} | Delete an application group.
[**ApplicationGroupsGetApplicationGroup**](ApplicationGroupsAPIsDAAS.md#ApplicationGroupsGetApplicationGroup) | **Get** /ApplicationGroups/{nameOrId} | Get details of a single application group.
[**ApplicationGroupsGetApplicationGroupApplications**](ApplicationGroupsAPIsDAAS.md#ApplicationGroupsGetApplicationGroupApplications) | **Get** /ApplicationGroups/{nameOrId}/Applications | Get applications in an application group.
[**ApplicationGroupsGetApplicationGroupDeliveryGroups**](ApplicationGroupsAPIsDAAS.md#ApplicationGroupsGetApplicationGroupDeliveryGroups) | **Get** /ApplicationGroups/{nameOrId}/DeliveryGroups | Get delivery groups for an application group.
[**ApplicationGroupsGetApplicationGroupTags**](ApplicationGroupsAPIsDAAS.md#ApplicationGroupsGetApplicationGroupTags) | **Get** /ApplicationGroups/{nameOrId}/Tags | Get the tags for an application group.
[**ApplicationGroupsGetApplicationGroups**](ApplicationGroupsAPIsDAAS.md#ApplicationGroupsGetApplicationGroups) | **Get** /ApplicationGroups | Get application groups.
[**ApplicationGroupsRemoveApplication**](ApplicationGroupsAPIsDAAS.md#ApplicationGroupsRemoveApplication) | **Delete** /ApplicationGroups/{nameOrId}/Applications/{appNameOrId} | Remove an application from the application group.
[**ApplicationGroupsRemoveApplicationGroupTag**](ApplicationGroupsAPIsDAAS.md#ApplicationGroupsRemoveApplicationGroupTag) | **Delete** /ApplicationGroups/{nameOrId}/Tags/{tagNameOrId} | Remove a tag from an application group.
[**ApplicationGroupsSetApplicationGroupTags**](ApplicationGroupsAPIsDAAS.md#ApplicationGroupsSetApplicationGroupTags) | **Put** /ApplicationGroups/{nameOrId}/Tags | Set the tags for an application group.
[**ApplicationGroupsTestApplicationGroupExists**](ApplicationGroupsAPIsDAAS.md#ApplicationGroupsTestApplicationGroupExists) | **Post** /ApplicationGroups/$checkNameExists | Test the existence of an application group by name.
[**ApplicationGroupsUpdateApplicationGroup**](ApplicationGroupsAPIsDAAS.md#ApplicationGroupsUpdateApplicationGroup) | **Patch** /ApplicationGroups/{nameOrId} | Update an application group.



## ApplicationGroupsAddApplicationGroupTag

> TagResponseModelCollection ApplicationGroupsAddApplicationGroupTag(ctx, nameOrId, tagNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Add a tag to an application group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the application group. If the application group is present in an application group folder,             specify the name in this format: {application group folder path plus application group name}.             For example, FolderName1|FolderName2|ApplicationGroupName.
    tagNameOrId := "tagNameOrId_example" // string | Name or ID of the tag to add.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    async := true // bool | If `true`, the tags will be modified as a background task. The task will have JobType SetTagApplicationGroups. When the task is complete it will redirect to GetApplicationGroupTags. The job's Parameters will contain properties: * _Id_ - ID of the application group being tagged, * _Name_ - Name of the application group being tagged. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGroupsAPIsDAAS.ApplicationGroupsAddApplicationGroupTag(context.Background(), nameOrId, tagNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsAPIsDAAS.ApplicationGroupsAddApplicationGroupTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationGroupsAddApplicationGroupTag`: TagResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGroupsAPIsDAAS.ApplicationGroupsAddApplicationGroupTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application group. If the application group is present in an application group folder,             specify the name in this format: {application group folder path plus application group name}.             For example, FolderName1|FolderName2|ApplicationGroupName. | 
**tagNameOrId** | **string** | Name or ID of the tag to add. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationGroupsAddApplicationGroupTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. The task will have JobType SetTagApplicationGroups. When the task is complete it will redirect to GetApplicationGroupTags. The job&#39;s Parameters will contain properties: * _Id_ - ID of the application group being tagged, * _Name_ - Name of the application group being tagged. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## ApplicationGroupsAddApplications

> ApplicationGroupsAddApplications(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).ApplicationGroupAddApplicationsRequestModel(applicationGroupAddApplicationsRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Add applications to the application group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the application group. If the application group is present in an application group folder,             specify the name in this format: {application group folder path plus application group name}.             For example, FolderName1|FolderName2|ApplicationGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    applicationGroupAddApplicationsRequestModel := *openapiclient.NewApplicationGroupAddApplicationsRequestModel() // ApplicationGroupAddApplicationsRequestModel | Details of the applications to add.
    async := true // bool | If `true`, the applications (and associated objects) will be added as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationGroupsAPIsDAAS.ApplicationGroupsAddApplications(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).ApplicationGroupAddApplicationsRequestModel(applicationGroupAddApplicationsRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsAPIsDAAS.ApplicationGroupsAddApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application group. If the application group is present in an application group folder,             specify the name in this format: {application group folder path plus application group name}.             For example, FolderName1|FolderName2|ApplicationGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationGroupsAddApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **applicationGroupAddApplicationsRequestModel** | [**ApplicationGroupAddApplicationsRequestModel**](ApplicationGroupAddApplicationsRequestModel.md) | Details of the applications to add. | 
 **async** | **bool** | If &#x60;true&#x60;, the applications (and associated objects) will be added as a background task. | [default to false]
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


## ApplicationGroupsCheckApplicationGroupExists

> ApplicationGroupsCheckApplicationGroupExists(ctx, name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Check for the existence of an application group by name.



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
    name := "name_example" // string | Name of the application group.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationGroupsAPIsDAAS.ApplicationGroupsCheckApplicationGroupExists(context.Background(), name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsAPIsDAAS.ApplicationGroupsCheckApplicationGroupExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the application group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationGroupsCheckApplicationGroupExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
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


## ApplicationGroupsCreateApplicationGroup

> ApplicationGroupDetailResponseModel ApplicationGroupsCreateApplicationGroup(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).CreateApplicationGroupRequestModel(createApplicationGroupRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Create an application group.

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
    createApplicationGroupRequestModel := *openapiclient.NewCreateApplicationGroupRequestModel("Name of the application group") // CreateApplicationGroupRequestModel | Details of the application group to create.
    async := true // bool | If `true`, the application group (and associated objects) will be created as a background task. The task will have JobType CreateApplicationGroup. When the task is complete it will redirect to GetApplicationGroup. The job's Parameters will contain properties: * _Name_ - Name of the application group being created. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGroupsAPIsDAAS.ApplicationGroupsCreateApplicationGroup(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).CreateApplicationGroupRequestModel(createApplicationGroupRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsAPIsDAAS.ApplicationGroupsCreateApplicationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationGroupsCreateApplicationGroup`: ApplicationGroupDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGroupsAPIsDAAS.ApplicationGroupsCreateApplicationGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplicationGroupsCreateApplicationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **createApplicationGroupRequestModel** | [**CreateApplicationGroupRequestModel**](CreateApplicationGroupRequestModel.md) | Details of the application group to create. | 
 **async** | **bool** | If &#x60;true&#x60;, the application group (and associated objects) will be created as a background task. The task will have JobType CreateApplicationGroup. When the task is complete it will redirect to GetApplicationGroup. The job&#39;s Parameters will contain properties: * _Name_ - Name of the application group being created. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**ApplicationGroupDetailResponseModel**](ApplicationGroupDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationGroupsDeleteApplicationGroup

> ApplicationGroupsDeleteApplicationGroup(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Delete an application group.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application group to delete. If the application group is present in an application group folder,             specify the name in this format: {application group folder path plus application group name}.             For example, FolderName1|FolderName2|ApplicationGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    async := true // bool | If `true`, the application group (and associated objects) will be deleted as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationGroupsAPIsDAAS.ApplicationGroupsDeleteApplicationGroup(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsAPIsDAAS.ApplicationGroupsDeleteApplicationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application group to delete. If the application group is present in an application group folder,             specify the name in this format: {application group folder path plus application group name}.             For example, FolderName1|FolderName2|ApplicationGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationGroupsDeleteApplicationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the application group (and associated objects) will be deleted as a background task. | [default to false]
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


## ApplicationGroupsGetApplicationGroup

> ApplicationGroupDetailResponseModel ApplicationGroupsGetApplicationGroup(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get details of a single application group.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the application group. If the application group is present in an application group folder,              specify the name in this format: {application group folder path plus application group name}.              For example, FolderName1|FolderName2|ApplicationGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    fields := "Id,Uid,ContainerScopes" // string | Optional filter, removing unspecified properties that otherwise would have been sent by the server              (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGroupsAPIsDAAS.ApplicationGroupsGetApplicationGroup(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsAPIsDAAS.ApplicationGroupsGetApplicationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationGroupsGetApplicationGroup`: ApplicationGroupDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGroupsAPIsDAAS.ApplicationGroupsGetApplicationGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application group. If the application group is present in an application group folder,              specify the name in this format: {application group folder path plus application group name}.              For example, FolderName1|FolderName2|ApplicationGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationGroupsGetApplicationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **fields** | **string** | Optional filter, removing unspecified properties that otherwise would have been sent by the server              | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**ApplicationGroupDetailResponseModel**](ApplicationGroupDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationGroupsGetApplicationGroupApplications

> ApplicationResponseModelCollection ApplicationGroupsGetApplicationGroupApplications(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get applications in an application group.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application group. If the application group is present in an application group folder,             specify the name in this format: {application group folder path plus application group name}.             For example, FolderName1|FolderName2|ApplicationGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    limit := int32(56) // int32 | The max number of applications returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "Id,Uid,ApplicationFolder" // string | Optional filter, removing unspecified properties that otherwise would have been sent by the server. (optional)
    async := true // bool | If \"true\", the applications under the application group will be fetched as a background task. The task will have JobType GetApplicationGroupApplications (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGroupsAPIsDAAS.ApplicationGroupsGetApplicationGroupApplications(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsAPIsDAAS.ApplicationGroupsGetApplicationGroupApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationGroupsGetApplicationGroupApplications`: ApplicationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGroupsAPIsDAAS.ApplicationGroupsGetApplicationGroupApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application group. If the application group is present in an application group folder,             specify the name in this format: {application group folder path plus application group name}.             For example, FolderName1|FolderName2|ApplicationGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationGroupsGetApplicationGroupApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **limit** | **int32** | The max number of applications returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | Optional filter, removing unspecified properties that otherwise would have been sent by the server. | 
 **async** | **bool** | If \&quot;true\&quot;, the applications under the application group will be fetched as a background task. The task will have JobType GetApplicationGroupApplications | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## ApplicationGroupsGetApplicationGroupDeliveryGroups

> ApplicationGroupDeliveryGroupResponseModelCollection ApplicationGroupsGetApplicationGroupDeliveryGroups(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get delivery groups for an application group.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application group. If the application group is present in an application group folder,             specify the name in this format: {application group folder path plus application group name}.             For example, FolderName1|FolderName2|ApplicationGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    limit := int32(56) // int32 | The max number of delivery groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    async := true // bool | If `true`, the delivery groups fetch will run as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGroupsAPIsDAAS.ApplicationGroupsGetApplicationGroupDeliveryGroups(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Limit(limit).ContinuationToken(continuationToken).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsAPIsDAAS.ApplicationGroupsGetApplicationGroupDeliveryGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationGroupsGetApplicationGroupDeliveryGroups`: ApplicationGroupDeliveryGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGroupsAPIsDAAS.ApplicationGroupsGetApplicationGroupDeliveryGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application group. If the application group is present in an application group folder,             specify the name in this format: {application group folder path plus application group name}.             For example, FolderName1|FolderName2|ApplicationGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationGroupsGetApplicationGroupDeliveryGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **limit** | **int32** | The max number of delivery groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **async** | **bool** | If &#x60;true&#x60;, the delivery groups fetch will run as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

### Return type

[**ApplicationGroupDeliveryGroupResponseModelCollection**](ApplicationGroupDeliveryGroupResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationGroupsGetApplicationGroupTags

> TagResponseModelCollection ApplicationGroupsGetApplicationGroupTags(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the tags for an application group.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application group. If the application group is present in an application group folder,             specify the name in this format: {application group folder path plus application group name}.             For example, FolderName1|FolderName2|ApplicationGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    fields := "Name,Id,Description" // string | field to filter response model. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGroupsAPIsDAAS.ApplicationGroupsGetApplicationGroupTags(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Fields(fields).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsAPIsDAAS.ApplicationGroupsGetApplicationGroupTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationGroupsGetApplicationGroupTags`: TagResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGroupsAPIsDAAS.ApplicationGroupsGetApplicationGroupTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application group. If the application group is present in an application group folder,             specify the name in this format: {application group folder path plus application group name}.             For example, FolderName1|FolderName2|ApplicationGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationGroupsGetApplicationGroupTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **fields** | **string** | field to filter response model. | 
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## ApplicationGroupsGetApplicationGroups

> ApplicationGroupResponseModelCollection ApplicationGroupsGetApplicationGroups(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).AdminFolder(adminFolder).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get application groups.



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
    adminFolder := "adminFolder_example" // string | Optional folder path (URL-encoded) or ID.  If not specified, all applications will be returned from all folders. (optional)
    limit := int32(56) // int32 | The max number of application groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    fields := "Id,Uid,ContainerScopes" // string | Optional. A filter string containing object fields requested to be returned, the requested fields are separated by comma',' (optional)
    async := true // bool | If `true`, the application groups will be fetched as a background task. The task will have JobType GetApplicationGroups. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGroupsAPIsDAAS.ApplicationGroupsGetApplicationGroups(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).AdminFolder(adminFolder).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsAPIsDAAS.ApplicationGroupsGetApplicationGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationGroupsGetApplicationGroups`: ApplicationGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGroupsAPIsDAAS.ApplicationGroupsGetApplicationGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplicationGroupsGetApplicationGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **adminFolder** | **string** | Optional folder path (URL-encoded) or ID.  If not specified, all applications will be returned from all folders. | 
 **limit** | **int32** | The max number of application groups returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | Optional. A filter string containing object fields requested to be returned, the requested fields are separated by comma&#39;,&#39; | 
 **async** | **bool** | If &#x60;true&#x60;, the application groups will be fetched as a background task. The task will have JobType GetApplicationGroups. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## ApplicationGroupsRemoveApplication

> ApplicationGroupsRemoveApplication(ctx, nameOrId, appNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Remove an application from the application group.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application group. If the application group is present in an application group folder,             specify the name in this format: {application group folder path plus application group name}.             For example, FolderName1|FolderName2|ApplicationGroupName.
    appNameOrId := "appNameOrId_example" // string | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    async := true // bool | If `true`, the application will be removed from the application group as a background task. The task will have JobType RemoveApplicationGroupApplications. When the task is complete it will redirect to GetApplicationGroupApplications. The job's Parameters will contain properties: * _ApplicationGroupId_ - ID of the application group. * _ApplicationGroupName_ - Name of the application group. * _ApplicationId_ - ID of the application. * _ApplicationName_ - Name of the application. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationGroupsAPIsDAAS.ApplicationGroupsRemoveApplication(context.Background(), nameOrId, appNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsAPIsDAAS.ApplicationGroupsRemoveApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application group. If the application group is present in an application group folder,             specify the name in this format: {application group folder path plus application group name}.             For example, FolderName1|FolderName2|ApplicationGroupName. | 
**appNameOrId** | **string** | Name or ID of the application. If the application is present in an application folder,             specify the name in this format: {application folder path plus application name}.             For example, FolderName1|FolderName2|ApplicationName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationGroupsRemoveApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the application will be removed from the application group as a background task. The task will have JobType RemoveApplicationGroupApplications. When the task is complete it will redirect to GetApplicationGroupApplications. The job&#39;s Parameters will contain properties: * _ApplicationGroupId_ - ID of the application group. * _ApplicationGroupName_ - Name of the application group. * _ApplicationId_ - ID of the application. * _ApplicationName_ - Name of the application. | [default to false]
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


## ApplicationGroupsRemoveApplicationGroupTag

> ApplicationGroupsRemoveApplicationGroupTag(ctx, nameOrId, tagNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Remove a tag from an application group.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application group. If the application group is present in an application group folder,             specify the name in this format: {application group folder path plus application group name}.             For example, FolderName1|FolderName2|ApplicationGroupName.
    tagNameOrId := "tagNameOrId_example" // string | Name or ID of the tag to remove.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    async := true // bool | If `true`, the tags will be modified as a background task. The task will have JobType SetTagApplicationGroups. When the task is complete it will redirect to GetApplicationGroupTags. The job's Parameters will contain properties: * _Id_ - ID of the application group being tagged, * _Name_ - Name of the application group being tagged. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationGroupsAPIsDAAS.ApplicationGroupsRemoveApplicationGroupTag(context.Background(), nameOrId, tagNameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsAPIsDAAS.ApplicationGroupsRemoveApplicationGroupTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application group. If the application group is present in an application group folder,             specify the name in this format: {application group folder path plus application group name}.             For example, FolderName1|FolderName2|ApplicationGroupName. | 
**tagNameOrId** | **string** | Name or ID of the tag to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationGroupsRemoveApplicationGroupTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. The task will have JobType SetTagApplicationGroups. When the task is complete it will redirect to GetApplicationGroupTags. The job&#39;s Parameters will contain properties: * _Id_ - ID of the application group being tagged, * _Name_ - Name of the application group being tagged. | [default to false]
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


## ApplicationGroupsSetApplicationGroupTags

> ApplicationGroupsSetApplicationGroupTags(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).TagsRequestModel(tagsRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Set the tags for an application group.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application group. If the application group is present in an application group folder,             specify the name in this format: {application group folder path plus application group name}.             For example, FolderName1|FolderName2|ApplicationGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    tagsRequestModel := *openapiclient.NewTagsRequestModel() // TagsRequestModel | Tags to set.
    async := true // bool | If `true`, the tags will be modified as a background task. The task will have JobType SetTagApplicationGroups. When the task is complete it will redirect to GetApplicationGroupTags. The job's Parameters will contain properties: * _Id_ - ID of the application group being tagged, * _Name_ - Name of the application group being tagged. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationGroupsAPIsDAAS.ApplicationGroupsSetApplicationGroupTags(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).TagsRequestModel(tagsRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsAPIsDAAS.ApplicationGroupsSetApplicationGroupTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application group. If the application group is present in an application group folder,             specify the name in this format: {application group folder path plus application group name}.             For example, FolderName1|FolderName2|ApplicationGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationGroupsSetApplicationGroupTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **tagsRequestModel** | [**TagsRequestModel**](TagsRequestModel.md) | Tags to set. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. The task will have JobType SetTagApplicationGroups. When the task is complete it will redirect to GetApplicationGroupTags. The job&#39;s Parameters will contain properties: * _Id_ - ID of the application group being tagged, * _Name_ - Name of the application group being tagged. | [default to false]
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


## ApplicationGroupsTestApplicationGroupExists

> ApplicationGroupsTestApplicationGroupExists(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).ApplicationGroupNameCheckRequestModel(applicationGroupNameCheckRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Test the existence of an application group by name.



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
    applicationGroupNameCheckRequestModel := *openapiclient.NewApplicationGroupNameCheckRequestModel() // ApplicationGroupNameCheckRequestModel | Request model of application group name.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationGroupsAPIsDAAS.ApplicationGroupsTestApplicationGroupExists(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).ApplicationGroupNameCheckRequestModel(applicationGroupNameCheckRequestModel).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsAPIsDAAS.ApplicationGroupsTestApplicationGroupExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplicationGroupsTestApplicationGroupExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **applicationGroupNameCheckRequestModel** | [**ApplicationGroupNameCheckRequestModel**](ApplicationGroupNameCheckRequestModel.md) | Request model of application group name. | 
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


## ApplicationGroupsUpdateApplicationGroup

> ApplicationGroupsUpdateApplicationGroup(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).EditApplicationGroupRequestModel(editApplicationGroupRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

Update an application group.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the application group to update. If the application group is present in an application group folder,             specify the name in this format: {application group folder path plus application group name}.             For example, FolderName1|FolderName2|ApplicationGroupName.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    editApplicationGroupRequestModel := *openapiclient.NewEditApplicationGroupRequestModel() // EditApplicationGroupRequestModel | Details of the application group to update.
    async := true // bool | If `true`, the application group (and associated objects) will be updated as a background task. The task will have JobType UpdateApplicationGroup. When the task is complete it will redirect to GetApplicationGroup. The job's Parameters will contain properties: * _Id_ - ID of the application group being updated. * _Name_ - Name of the application group being updated. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationGroupsAPIsDAAS.ApplicationGroupsUpdateApplicationGroup(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).EditApplicationGroupRequestModel(editApplicationGroupRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsAPIsDAAS.ApplicationGroupsUpdateApplicationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the application group to update. If the application group is present in an application group folder,             specify the name in this format: {application group folder path plus application group name}.             For example, FolderName1|FolderName2|ApplicationGroupName. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationGroupsUpdateApplicationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **editApplicationGroupRequestModel** | [**EditApplicationGroupRequestModel**](EditApplicationGroupRequestModel.md) | Details of the application group to update. | 
 **async** | **bool** | If &#x60;true&#x60;, the application group (and associated objects) will be updated as a background task. The task will have JobType UpdateApplicationGroup. When the task is complete it will redirect to GetApplicationGroup. The job&#39;s Parameters will contain properties: * _Id_ - ID of the application group being updated. * _Name_ - Name of the application group being updated. | [default to false]
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

