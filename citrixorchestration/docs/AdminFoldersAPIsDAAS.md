# \AdminFoldersAPIsDAAS

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminFoldersCreateAdminFolder**](AdminFoldersAPIsDAAS.md#AdminFoldersCreateAdminFolder) | **Post** /AdminFolders | Create an admin folder
[**AdminFoldersDeleteAdminFolder**](AdminFoldersAPIsDAAS.md#AdminFoldersDeleteAdminFolder) | **Delete** /AdminFolders/{pathOrId} | Delete an admin folder.
[**AdminFoldersGetAdminFolder**](AdminFoldersAPIsDAAS.md#AdminFoldersGetAdminFolder) | **Get** /AdminFolders/{pathOrId} | Get details about a single admin folder.
[**AdminFoldersGetAdminFolderApplications**](AdminFoldersAPIsDAAS.md#AdminFoldersGetAdminFolderApplications) | **Get** /AdminFolders/{pathOrId}/Applications | Get the applications in an admin folder.
[**AdminFoldersGetAdminFolderMachineCatalogs**](AdminFoldersAPIsDAAS.md#AdminFoldersGetAdminFolderMachineCatalogs) | **Get** /AdminFolders/{pathOrId}/MachineCatalogs | Get the machineCatalogs in an admin folder.
[**AdminFoldersGetAdminFolders**](AdminFoldersAPIsDAAS.md#AdminFoldersGetAdminFolders) | **Get** /AdminFolders | Get admin folders.
[**AdminFoldersUpdateAdminFolder**](AdminFoldersAPIsDAAS.md#AdminFoldersUpdateAdminFolder) | **Patch** /AdminFolders/{pathOrId} | Update an admin folder.



## AdminFoldersCreateAdminFolder

> AdminFolderResponseModel AdminFoldersCreateAdminFolder(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateAdminFolderRequestModel(createAdminFolderRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	createAdminFolderRequestModel := *openapiclient.NewCreateAdminFolderRequestModel("Name of the admin folder", []openapiclient.AdminFolderObjectIdentifier{openapiclient.AdminFolderObjectIdentifier("Unknown")}) // CreateAdminFolderRequestModel | Details of the admin folder to create.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the admin folder will be created as a background task. The task will have JobType \"JobType.CreateAdminFolder\". When the task is complete it will redirect to \"GetAdminFolder(string)\" The job Parameters will contain properties:  * _Name_ - Name of the folder being created. * _Path_ - Path to the folder being created. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminFoldersAPIsDAAS.AdminFoldersCreateAdminFolder(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateAdminFolderRequestModel(createAdminFolderRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminFoldersAPIsDAAS.AdminFoldersCreateAdminFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminFoldersCreateAdminFolder`: AdminFolderResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AdminFoldersAPIsDAAS.AdminFoldersCreateAdminFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminFoldersCreateAdminFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **createAdminFolderRequestModel** | [**CreateAdminFolderRequestModel**](CreateAdminFolderRequestModel.md) | Details of the admin folder to create. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the admin folder will be created as a background task. The task will have JobType \&quot;JobType.CreateAdminFolder\&quot;. When the task is complete it will redirect to \&quot;GetAdminFolder(string)\&quot; The job Parameters will contain properties:  * _Name_ - Name of the folder being created. * _Path_ - Path to the folder being created. | [default to false]

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


## AdminFoldersDeleteAdminFolder

> AdminFoldersDeleteAdminFolder(ctx, pathOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).ObjectsToRemove(objectsToRemove).Async(async).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	pathOrId := "pathOrId_example" // string | Name or ID of the admin folder to delete.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	objectsToRemove := []openapiclient.AdminFolderObjects{openapiclient.AdminFolderObjects("Unknown")} // []AdminFolderObjects | Identify which object should be removed from the adminFolder. If set to null, then the admin folder will be deleted if it hosts no object and no folder metaData with keys AdminFolderObjectIdentifier.              If set to Applications, the applications in the folder will be deleted. The folder will be deleted if it do not host other objects(MachineCatalogs, or DeliveryGroup, or ApplicationGroups) and no folder metaData with keysAdminFolderObjectIdentifier.; If the folder hosts any other objects, then this folder will not be deleted, folder metadata with key ContainsApplications will be removed.              If set to MachineCatalogs and this folder still hosts some MachineCatalogs, then this folder is not allowed to be deleted; If this folder do not host any MachineCatalogs but host other objects(Applications, DeliveryGroups, ApplicationGroups), the folder metaData with key ContainsMachineCatalogs will be removed; If this folder do not host any objects and no folder metaData with keys AdminFolderObjectIdentifier, the folder will be deleted.              If set to DeliveryGroups or ApplicationGroups, have the same behavior as set to MachineCatalogs. (optional)
	async := true // bool | If `true`, the admin folder will be deleted as a background task. The task will have JobType DeleteAdminFolder. When the task is complete it will redirect to GetAdminFolders. The job's Parameters will contain properties:  * _Id_ - ID of the admin folder being deleted, * _Path_ - Path of the admin folder being deleted. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminFoldersAPIsDAAS.AdminFoldersDeleteAdminFolder(context.Background(), pathOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).ObjectsToRemove(objectsToRemove).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminFoldersAPIsDAAS.AdminFoldersDeleteAdminFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathOrId** | **string** | Name or ID of the admin folder to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminFoldersDeleteAdminFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **objectsToRemove** | [**[]AdminFolderObjects**](AdminFolderObjects.md) | Identify which object should be removed from the adminFolder. If set to null, then the admin folder will be deleted if it hosts no object and no folder metaData with keys AdminFolderObjectIdentifier.              If set to Applications, the applications in the folder will be deleted. The folder will be deleted if it do not host other objects(MachineCatalogs, or DeliveryGroup, or ApplicationGroups) and no folder metaData with keysAdminFolderObjectIdentifier.; If the folder hosts any other objects, then this folder will not be deleted, folder metadata with key ContainsApplications will be removed.              If set to MachineCatalogs and this folder still hosts some MachineCatalogs, then this folder is not allowed to be deleted; If this folder do not host any MachineCatalogs but host other objects(Applications, DeliveryGroups, ApplicationGroups), the folder metaData with key ContainsMachineCatalogs will be removed; If this folder do not host any objects and no folder metaData with keys AdminFolderObjectIdentifier, the folder will be deleted.              If set to DeliveryGroups or ApplicationGroups, have the same behavior as set to MachineCatalogs. | 
 **async** | **bool** | If &#x60;true&#x60;, the admin folder will be deleted as a background task. The task will have JobType DeleteAdminFolder. When the task is complete it will redirect to GetAdminFolders. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the admin folder being deleted, * _Path_ - Path of the admin folder being deleted. | [default to false]

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


## AdminFoldersGetAdminFolder

> AdminFolderResponseModel AdminFoldersGetAdminFolder(ctx, pathOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	pathOrId := "pathOrId_example" // string | Path (URL-encoded) or ID of the admin folder.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminFoldersAPIsDAAS.AdminFoldersGetAdminFolder(context.Background(), pathOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminFoldersAPIsDAAS.AdminFoldersGetAdminFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminFoldersGetAdminFolder`: AdminFolderResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AdminFoldersAPIsDAAS.AdminFoldersGetAdminFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathOrId** | **string** | Path (URL-encoded) or ID of the admin folder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminFoldersGetAdminFolderRequest struct via the builder pattern


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

[**AdminFolderResponseModel**](AdminFolderResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminFoldersGetAdminFolderApplications

> ApplicationResponseModelCollection AdminFoldersGetAdminFolderApplications(ctx, pathOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	pathOrId := "pathOrId_example" // string | Path (URL-encoded) or ID of the admin folder.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	limit := int32(56) // int32 | The max number of applications returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
	continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
	fields := "Id,Uid,ApplicationFolder" // string | Optional filter, removing unspecified properties that otherwise would have been sent by the server. (optional)
	async := true // bool | If `true`, Fetch applications under admin folder will be a background task. The task will have JobType GetAdminFolderApplications (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminFoldersAPIsDAAS.AdminFoldersGetAdminFolderApplications(context.Background(), pathOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminFoldersAPIsDAAS.AdminFoldersGetAdminFolderApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminFoldersGetAdminFolderApplications`: ApplicationResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `AdminFoldersAPIsDAAS.AdminFoldersGetAdminFolderApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathOrId** | **string** | Path (URL-encoded) or ID of the admin folder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminFoldersGetAdminFolderApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | The max number of applications returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **fields** | **string** | Optional filter, removing unspecified properties that otherwise would have been sent by the server. | 
 **async** | **bool** | If &#x60;true&#x60;, Fetch applications under admin folder will be a background task. The task will have JobType GetAdminFolderApplications | [default to false]

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


## AdminFoldersGetAdminFolderMachineCatalogs

> MachineCatalogResponseModelCollection AdminFoldersGetAdminFolderMachineCatalogs(ctx, pathOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	pathOrId := "pathOrId_example" // string | Path (URL-encoded) or ID of the admin folder.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	limit := int32(56) // int32 | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional) (default to 2147483647)
	continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)
	fields := "Name,FullName,Id" // string | Optional filter, removing unspecified properties that otherwise would have been sent by the server. (optional)
	async := true // bool | If `true`, Fetch machineCatalogs under admin folder will be a background task. The task will have JobType GetAdminFolderMachineCatalogs (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminFoldersAPIsDAAS.AdminFoldersGetAdminFolderMachineCatalogs(context.Background(), pathOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Fields(fields).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminFoldersAPIsDAAS.AdminFoldersGetAdminFolderMachineCatalogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminFoldersGetAdminFolderMachineCatalogs`: MachineCatalogResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `AdminFoldersAPIsDAAS.AdminFoldersGetAdminFolderMachineCatalogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathOrId** | **string** | Path (URL-encoded) or ID of the admin folder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminFoldersGetAdminFolderMachineCatalogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | [default to 2147483647]
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 
 **fields** | **string** | Optional filter, removing unspecified properties that otherwise would have been sent by the server. | 
 **async** | **bool** | If &#x60;true&#x60;, Fetch machineCatalogs under admin folder will be a background task. The task will have JobType GetAdminFolderMachineCatalogs | [default to false]

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


## AdminFoldersGetAdminFolders

> AdminFolderResponseModelCollection AdminFoldersGetAdminFolders(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Limit(limit).ContinuationToken(continuationToken).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the admin folders will be fetched as a background task. The task will have JobType GetAdminFolders. (optional) (default to false)
	limit := int32(56) // int32 | The max number of admin folders returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
	continuationToken := "continuationToken_example" // string | The continuationToken returned by the previous query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminFoldersAPIsDAAS.AdminFoldersGetAdminFolders(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Limit(limit).ContinuationToken(continuationToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminFoldersAPIsDAAS.AdminFoldersGetAdminFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminFoldersGetAdminFolders`: AdminFolderResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `AdminFoldersAPIsDAAS.AdminFoldersGetAdminFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminFoldersGetAdminFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the admin folders will be fetched as a background task. The task will have JobType GetAdminFolders. | [default to false]
 **limit** | **int32** | The max number of admin folders returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | The continuationToken returned by the previous query. | 

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


## AdminFoldersUpdateAdminFolder

> AdminFolderResponseModel AdminFoldersUpdateAdminFolder(ctx, pathOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).EditAdminFolderRequestModel(editAdminFolderRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	pathOrId := "pathOrId_example" // string | Path or ID of the admin folder path to update.
	editAdminFolderRequestModel := *openapiclient.NewEditAdminFolderRequestModel() // EditAdminFolderRequestModel | Details of the admin folder to update.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the admin folder will be updated as a background task. The task will have JobType UpdateAdminFolder. When the task is complete it will redirect to GetAdminFolder. The job's Parameters will contain properties:  * _Id_ - ID of the admin folder being updated. * _Path_ - Path of the admin folder being updated. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminFoldersAPIsDAAS.AdminFoldersUpdateAdminFolder(context.Background(), pathOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).EditAdminFolderRequestModel(editAdminFolderRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminFoldersAPIsDAAS.AdminFoldersUpdateAdminFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminFoldersUpdateAdminFolder`: AdminFolderResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AdminFoldersAPIsDAAS.AdminFoldersUpdateAdminFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pathOrId** | **string** | Path or ID of the admin folder path to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminFoldersUpdateAdminFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **editAdminFolderRequestModel** | [**EditAdminFolderRequestModel**](EditAdminFolderRequestModel.md) | Details of the admin folder to update. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the admin folder will be updated as a background task. The task will have JobType UpdateAdminFolder. When the task is complete it will redirect to GetAdminFolder. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the admin folder being updated. * _Path_ - Path of the admin folder being updated. | [default to false]

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

