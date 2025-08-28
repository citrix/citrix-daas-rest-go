# \StoreFrontServersAPIsDAAS

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StoreFrontServersCreateStoreFrontServer**](StoreFrontServersAPIsDAAS.md#StoreFrontServersCreateStoreFrontServer) | **Post** /StoreFrontServers | Create a StoreFront server.
[**StoreFrontServersDeleteStoreFrontServer**](StoreFrontServersAPIsDAAS.md#StoreFrontServersDeleteStoreFrontServer) | **Delete** /StoreFrontServers/{nameOrId} | Delete a StoreFront server from the site.
[**StoreFrontServersGetStoreFrontDeliveryGroups**](StoreFrontServersAPIsDAAS.md#StoreFrontServersGetStoreFrontDeliveryGroups) | **Get** /StoreFrontServers/{nameOrId}/DeliveryGroups | GET delivery groups details for a Storefront
[**StoreFrontServersGetStoreFrontServer**](StoreFrontServersAPIsDAAS.md#StoreFrontServersGetStoreFrontServer) | **Get** /StoreFrontServers/{nameOrId} | Get the details for a single StoreFront server.
[**StoreFrontServersGetStoreFrontServers**](StoreFrontServersAPIsDAAS.md#StoreFrontServersGetStoreFrontServers) | **Get** /StoreFrontServers | Get all StoreFront servers.
[**StoreFrontServersUpdateStoreFrontServer**](StoreFrontServersAPIsDAAS.md#StoreFrontServersUpdateStoreFrontServer) | **Patch** /StoreFrontServers/{nameOrId} | Update a StoreFront server.



## StoreFrontServersCreateStoreFrontServer

> StoreFrontServerResponseModel StoreFrontServersCreateStoreFrontServer(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).StoreFrontServerRequestModel(storeFrontServerRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Create a StoreFront server.



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
	storeFrontServerRequestModel := *openapiclient.NewStoreFrontServerRequestModel() // StoreFrontServerRequestModel | Details about the StoreFront server to create.  Note: the Id property must not be specified in the request.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the StoreFront server (and associated objects) will be created as a background task.  The task will have JobType CreateStoreFrontServer.  When the task is complete it will redirect to GetStoreFrontServer.  The job's Parameters will contain properties: * _Name_ - Name of the StoreFront server being created. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreFrontServersAPIsDAAS.StoreFrontServersCreateStoreFrontServer(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).StoreFrontServerRequestModel(storeFrontServerRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreFrontServersAPIsDAAS.StoreFrontServersCreateStoreFrontServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreFrontServersCreateStoreFrontServer`: StoreFrontServerResponseModel
	fmt.Fprintf(os.Stdout, "Response from `StoreFrontServersAPIsDAAS.StoreFrontServersCreateStoreFrontServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreFrontServersCreateStoreFrontServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **storeFrontServerRequestModel** | [**StoreFrontServerRequestModel**](StoreFrontServerRequestModel.md) | Details about the StoreFront server to create.  Note: the Id property must not be specified in the request. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the StoreFront server (and associated objects) will be created as a background task.  The task will have JobType CreateStoreFrontServer.  When the task is complete it will redirect to GetStoreFrontServer.  The job&#39;s Parameters will contain properties: * _Name_ - Name of the StoreFront server being created. | [default to false]

### Return type

[**StoreFrontServerResponseModel**](StoreFrontServerResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreFrontServersDeleteStoreFrontServer

> StoreFrontServersDeleteStoreFrontServer(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Delete a StoreFront server from the site.



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
	nameOrId := "nameOrId_example" // string | The name or ID of StoreFront server to be deleted.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the StoreFront server (and associated objects) will be deleted as a background task.  The task will have JobType DeleteStoreFrontServer.  When the task is complete it will redirect to GetStoreFrontServers.  The job's Parameters will contain properties: * _Name_ - Name of the StoreFront server being deleted. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreFrontServersAPIsDAAS.StoreFrontServersDeleteStoreFrontServer(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreFrontServersAPIsDAAS.StoreFrontServersDeleteStoreFrontServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | The name or ID of StoreFront server to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoreFrontServersDeleteStoreFrontServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the StoreFront server (and associated objects) will be deleted as a background task.  The task will have JobType DeleteStoreFrontServer.  When the task is complete it will redirect to GetStoreFrontServers.  The job&#39;s Parameters will contain properties: * _Name_ - Name of the StoreFront server being deleted. | [default to false]

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


## StoreFrontServersGetStoreFrontDeliveryGroups

> StoreFrontDeliveryGroupResponseModelCollection StoreFrontServersGetStoreFrontDeliveryGroups(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

GET delivery groups details for a Storefront



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
	nameOrId := "nameOrId_example" // string | The id of the Storefront
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreFrontServersAPIsDAAS.StoreFrontServersGetStoreFrontDeliveryGroups(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreFrontServersAPIsDAAS.StoreFrontServersGetStoreFrontDeliveryGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreFrontServersGetStoreFrontDeliveryGroups`: StoreFrontDeliveryGroupResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `StoreFrontServersAPIsDAAS.StoreFrontServersGetStoreFrontDeliveryGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | The id of the Storefront | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoreFrontServersGetStoreFrontDeliveryGroupsRequest struct via the builder pattern


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

[**StoreFrontDeliveryGroupResponseModelCollection**](StoreFrontDeliveryGroupResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreFrontServersGetStoreFrontServer

> StoreFrontServerResponseModel StoreFrontServersGetStoreFrontServer(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get the details for a single StoreFront server.



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
	nameOrId := "nameOrId_example" // string | The name or ID of the StoreFront server.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreFrontServersAPIsDAAS.StoreFrontServersGetStoreFrontServer(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreFrontServersAPIsDAAS.StoreFrontServersGetStoreFrontServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreFrontServersGetStoreFrontServer`: StoreFrontServerResponseModel
	fmt.Fprintf(os.Stdout, "Response from `StoreFrontServersAPIsDAAS.StoreFrontServersGetStoreFrontServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | The name or ID of the StoreFront server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoreFrontServersGetStoreFrontServerRequest struct via the builder pattern


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

[**StoreFrontServerResponseModel**](StoreFrontServerResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreFrontServersGetStoreFrontServers

> StoreFrontServerResponseModelCollection StoreFrontServersGetStoreFrontServers(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

Get all StoreFront servers.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreFrontServersAPIsDAAS.StoreFrontServersGetStoreFrontServers(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreFrontServersAPIsDAAS.StoreFrontServersGetStoreFrontServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreFrontServersGetStoreFrontServers`: StoreFrontServerResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `StoreFrontServersAPIsDAAS.StoreFrontServersGetStoreFrontServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreFrontServersGetStoreFrontServersRequest struct via the builder pattern


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

[**StoreFrontServerResponseModelCollection**](StoreFrontServerResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreFrontServersUpdateStoreFrontServer

> StoreFrontServerResponseModel StoreFrontServersUpdateStoreFrontServer(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).StoreFrontServerRequestModel(storeFrontServerRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Update a StoreFront server.



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
	nameOrId := "nameOrId_example" // string | Name or ID of the StoreFront server to update.
	storeFrontServerRequestModel := *openapiclient.NewStoreFrontServerRequestModel() // StoreFrontServerRequestModel | Properties of the StoreFront server to update.  Note: the Id property must not be specified in the request.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the StoreFront server (and associated objects) will be updated as a background task.  The task will have JobType UpdateStoreFrontServer.  When the task is complete it will redirect to GetStoreFrontServer.  The job's Parameters will contain properties: * _Name_ - Name of the StoreFront server being updated. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreFrontServersAPIsDAAS.StoreFrontServersUpdateStoreFrontServer(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).StoreFrontServerRequestModel(storeFrontServerRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreFrontServersAPIsDAAS.StoreFrontServersUpdateStoreFrontServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreFrontServersUpdateStoreFrontServer`: StoreFrontServerResponseModel
	fmt.Fprintf(os.Stdout, "Response from `StoreFrontServersAPIsDAAS.StoreFrontServersUpdateStoreFrontServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the StoreFront server to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoreFrontServersUpdateStoreFrontServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **storeFrontServerRequestModel** | [**StoreFrontServerRequestModel**](StoreFrontServerRequestModel.md) | Properties of the StoreFront server to update.  Note: the Id property must not be specified in the request. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the StoreFront server (and associated objects) will be updated as a background task.  The task will have JobType UpdateStoreFrontServer.  When the task is complete it will redirect to GetStoreFrontServer.  The job&#39;s Parameters will contain properties: * _Name_ - Name of the StoreFront server being updated. | [default to false]

### Return type

[**StoreFrontServerResponseModel**](StoreFrontServerResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

