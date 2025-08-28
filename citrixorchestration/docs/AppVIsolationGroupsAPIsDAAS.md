# \AppVIsolationGroupsAPIsDAAS

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppVIsolationGroupsCreateAppVIsolationGroup**](AppVIsolationGroupsAPIsDAAS.md#AppVIsolationGroupsCreateAppVIsolationGroup) | **Post** /AppVIsolationGroups | Create an App-V IsolationGroup in the site
[**AppVIsolationGroupsDeleteAppVIsolationGroup**](AppVIsolationGroupsAPIsDAAS.md#AppVIsolationGroupsDeleteAppVIsolationGroup) | **Delete** /AppVIsolationGroups/{nameOrId} | Delete an App-V IsolationGroup configured in the site.
[**AppVIsolationGroupsGetAppVIsolationGroup**](AppVIsolationGroupsAPIsDAAS.md#AppVIsolationGroupsGetAppVIsolationGroup) | **Get** /AppVIsolationGroups/{nameOrId} | Get the specified App-V IsolationGroups configured in the site
[**AppVIsolationGroupsGetAppVIsolationGroups**](AppVIsolationGroupsAPIsDAAS.md#AppVIsolationGroupsGetAppVIsolationGroups) | **Get** /AppVIsolationGroups | Get the App-V IsolationGroups configured in the site
[**AppVIsolationGroupsUpdateAppVIsolationGroup**](AppVIsolationGroupsAPIsDAAS.md#AppVIsolationGroupsUpdateAppVIsolationGroup) | **Patch** /AppVIsolationGroups/{nameOrId} | Update the App-V IsolationGroup configured in the site.



## AppVIsolationGroupsCreateAppVIsolationGroup

> AppVIsolationGroupsCreateAppVIsolationGroup(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateAppVIsolationGroupRequestModel(createAppVIsolationGroupRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	createAppVIsolationGroupRequestModel := *openapiclient.NewCreateAppVIsolationGroupRequestModel(int32(123), "Name_example") // CreateAppVIsolationGroupRequestModel | Request model of the new IsolationGroup.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the tags will be modified as a background task. The task will have JobType CreateAppVIsolationGroup. When the task is complete it will redirect to GetAppVIsolationGroups. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppVIsolationGroupsAPIsDAAS.AppVIsolationGroupsCreateAppVIsolationGroup(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateAppVIsolationGroupRequestModel(createAppVIsolationGroupRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **createAppVIsolationGroupRequestModel** | [**CreateAppVIsolationGroupRequestModel**](CreateAppVIsolationGroupRequestModel.md) | Request model of the new IsolationGroup. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. The task will have JobType CreateAppVIsolationGroup. When the task is complete it will redirect to GetAppVIsolationGroups. | [default to false]

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

> AppVIsolationGroupsDeleteAppVIsolationGroup(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	nameOrId := "nameOrId_example" // string | Name or UID of an isolationGroup.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the tags will be modified as a background task. The task will have JobType CreateAppVIsolationGroup. When the task is complete it will redirect to GetAppVIsolationGroups. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppVIsolationGroupsAPIsDAAS.AppVIsolationGroupsDeleteAppVIsolationGroup(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. The task will have JobType CreateAppVIsolationGroup. When the task is complete it will redirect to GetAppVIsolationGroups. | [default to false]

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

> AppVIsolationGroupResponseModel AppVIsolationGroupsGetAppVIsolationGroup(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	nameOrId := "nameOrId_example" // string | Name or UID of an isolationGroup.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the tags will be modified as a background task. The task will have JobType GetAppVIsolationGroup. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppVIsolationGroupsAPIsDAAS.AppVIsolationGroupsGetAppVIsolationGroup(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. The task will have JobType GetAppVIsolationGroup. | [default to false]

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

> AppVIsolationGroupResponseModelCollection AppVIsolationGroupsGetAppVIsolationGroups(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppVIsolationGroupsAPIsDAAS.AppVIsolationGroupsGetAppVIsolationGroups(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** |  | [default to false]

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

> AppVIsolationGroupsUpdateAppVIsolationGroup(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateAppVIsolationGroupRequestModel(updateAppVIsolationGroupRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
	citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
	citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
	nameOrId := "nameOrId_example" // string | Name or UID of an isolationGroup.
	updateAppVIsolationGroupRequestModel := *openapiclient.NewUpdateAppVIsolationGroupRequestModel() // UpdateAppVIsolationGroupRequestModel | request model to update isolationGroup.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the tags will be modified as a background task. The task will have JobType UpdateAppVIsolationGroup. When the task is complete it will redirect to GetAppVIsolationGroups. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppVIsolationGroupsAPIsDAAS.AppVIsolationGroupsUpdateAppVIsolationGroup(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateAppVIsolationGroupRequestModel(updateAppVIsolationGroupRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **updateAppVIsolationGroupRequestModel** | [**UpdateAppVIsolationGroupRequestModel**](UpdateAppVIsolationGroupRequestModel.md) | request model to update isolationGroup. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the tags will be modified as a background task. The task will have JobType UpdateAppVIsolationGroup. When the task is complete it will redirect to GetAppVIsolationGroups. | [default to false]

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

