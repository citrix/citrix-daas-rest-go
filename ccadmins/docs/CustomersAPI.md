# \CustomersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerCustomersIdDelete**](CustomersAPI.md#CustomerCustomersIdDelete) | **Delete** /{customer}/customers/{id} | Delete customer&#39;s administrators. [ServiceKey][RootOnly]
[**CustomerCustomersPost**](CustomersAPI.md#CustomerCustomersPost) | **Post** /{customer}/customers | Creates a new customer. [ServiceKey][RootOnly]



## CustomerCustomersIdDelete

> bool CustomerCustomersIdDelete(ctx, id, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).Execute()

Delete customer's administrators. [ServiceKey][RootOnly]

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/ccadmins"
)

func main() {
	id := "id_example" // string | Id of the customer.
	customer := "root" // string | 
	citrixConsistencyToken := "citrixConsistencyToken_example" // string | CosmosDB consistency token. (optional)
	xCwsTransactionId := "xCwsTransactionId_example" // string | Used for Citrix Cloud telemetry correlation. (optional)
	citrixTimeoutMs := TODO // float64 | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CustomerCustomersIdDelete(context.Background(), id, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomerCustomersIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerCustomersIdDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomerCustomersIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the customer. | 
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerCustomersIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 
 **citrixTimeoutMs** | [**float64**](float64.md) | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds | 

### Return type

**bool**

### Authorization

[CCAuthService](../README.md#CCAuthService)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerCustomersPost

> SystemNetHttpHttpResponseMessage CustomerCustomersPost(ctx, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixCloudServicesCoreApiClientModelsCustomersCustomerCreateModel(citrixCloudServicesCoreApiClientModelsCustomersCustomerCreateModel).Execute()

Creates a new customer. [ServiceKey][RootOnly]

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/ccadmins"
)

func main() {
	customer := "root" // string | 
	citrixConsistencyToken := "citrixConsistencyToken_example" // string | CosmosDB consistency token. (optional)
	xCwsTransactionId := "xCwsTransactionId_example" // string | Used for Citrix Cloud telemetry correlation. (optional)
	citrixCloudServicesCoreApiClientModelsCustomersCustomerCreateModel := *openapiclient.NewCitrixCloudServicesCoreApiClientModelsCustomersCustomerCreateModel() // CitrixCloudServicesCoreApiClientModelsCustomersCustomerCreateModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CustomerCustomersPost(context.Background(), customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixCloudServicesCoreApiClientModelsCustomersCustomerCreateModel(citrixCloudServicesCoreApiClientModelsCustomersCustomerCreateModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomerCustomersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerCustomersPost`: SystemNetHttpHttpResponseMessage
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomerCustomersPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerCustomersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 
 **citrixCloudServicesCoreApiClientModelsCustomersCustomerCreateModel** | [**CitrixCloudServicesCoreApiClientModelsCustomersCustomerCreateModel**](CitrixCloudServicesCoreApiClientModelsCustomersCustomerCreateModel.md) |  | 

### Return type

[**SystemNetHttpHttpResponseMessage**](SystemNetHttpHttpResponseMessage.md)

### Authorization

[CCAuthService](../README.md#CCAuthService)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

