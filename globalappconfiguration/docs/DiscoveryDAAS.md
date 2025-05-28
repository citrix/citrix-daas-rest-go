# \DiscoveryDAAS

All URIs are relative to *https://wsaca.cloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDiscovery**](DiscoveryDAAS.md#CreateDiscovery) | **Post** /aca/discovery/app/{app}/domain | Create A new Discovery Record For The Specified Customer Id And A Specific Product.
[**DeleteDiscovery**](DiscoveryDAAS.md#DeleteDiscovery) | **Delete** /aca/discovery/app/{app}/domain/{domain} | Delete Discovery Record For The Specified Customer Id, Domain And Product.
[**RetrieveAllDiscovery**](DiscoveryDAAS.md#RetrieveAllDiscovery) | **Get** /aca/discovery | Retrieve Discovery Records For The Specified Customer Id
[**RetrieveDiscovery**](DiscoveryDAAS.md#RetrieveDiscovery) | **Get** /aca/discovery/app/{app}/domain/{domain} | Retrieve Discovery Record For The Specified Customer Id, Domain And A Specific Product.
[**UpdateDiscovery**](DiscoveryDAAS.md#UpdateDiscovery) | **Put** /aca/discovery/app/{app}/domain/{domain} | Update Discovery Record For The Specified Customer Id, Domain And Product.



## CreateDiscovery

> DiscoveryRecordModel CreateDiscovery(ctx, app).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).DiscoveryRecordModel(discoveryRecordModel).Execute()

Create A new Discovery Record For The Specified Customer Id And A Specific Product.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/globalappconfiguration"
)

func main() {
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	app := "app_example" // string | App Name. ['workspace']
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)
	discoveryRecordModel := *openapiclient.NewDiscoveryRecordModel() // DiscoveryRecordModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscoveryDAAS.CreateDiscovery(context.Background(), app).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).DiscoveryRecordModel(discoveryRecordModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryDAAS.CreateDiscovery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDiscovery`: DiscoveryRecordModel
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryDAAS.CreateDiscovery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App Name. [&#39;workspace&#39;] | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDiscoveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 

 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 
 **discoveryRecordModel** | [**DiscoveryRecordModel**](DiscoveryRecordModel.md) |  | 

### Return type

[**DiscoveryRecordModel**](DiscoveryRecordModel.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDiscovery

> string DeleteDiscovery(ctx, app, domain).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Delete Discovery Record For The Specified Customer Id, Domain And Product.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/globalappconfiguration"
)

func main() {
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	app := "app_example" // string | App Name. ['workspace']
	domain := "domain_example" // string | Domain used for user's email address.
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscoveryDAAS.DeleteDiscovery(context.Background(), app, domain).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryDAAS.DeleteDiscovery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDiscovery`: string
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryDAAS.DeleteDiscovery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App Name. [&#39;workspace&#39;] | 
**domain** | **string** | Domain used for user&#39;s email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDiscoveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 


 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

**string**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAllDiscovery

> GetAllDiscoveryResponse RetrieveAllDiscovery(ctx).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Retrieve Discovery Records For The Specified Customer Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/globalappconfiguration"
)

func main() {
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscoveryDAAS.RetrieveAllDiscovery(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryDAAS.RetrieveAllDiscovery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveAllDiscovery`: GetAllDiscoveryResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryDAAS.RetrieveAllDiscovery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAllDiscoveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

[**GetAllDiscoveryResponse**](GetAllDiscoveryResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDiscovery

> GetAllDiscoveryResponse RetrieveDiscovery(ctx, app, domain).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Retrieve Discovery Record For The Specified Customer Id, Domain And A Specific Product.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/globalappconfiguration"
)

func main() {
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	app := "app_example" // string | App Name. ['workspace']
	domain := "domain_example" // string | Domain used for user's email address.
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscoveryDAAS.RetrieveDiscovery(context.Background(), app, domain).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryDAAS.RetrieveDiscovery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveDiscovery`: GetAllDiscoveryResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryDAAS.RetrieveDiscovery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App Name. [&#39;workspace&#39;] | 
**domain** | **string** | Domain used for user&#39;s email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDiscoveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 


 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

[**GetAllDiscoveryResponse**](GetAllDiscoveryResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDiscovery

> string UpdateDiscovery(ctx, app, domain).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).DiscoveryRecordModel(discoveryRecordModel).Execute()

Update Discovery Record For The Specified Customer Id, Domain And Product.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/globalappconfiguration"
)

func main() {
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	app := "app_example" // string | App Name. ['workspace']
	domain := "domain_example" // string | Domain used for user's email address.
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)
	discoveryRecordModel := *openapiclient.NewDiscoveryRecordModel() // DiscoveryRecordModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscoveryDAAS.UpdateDiscovery(context.Background(), app, domain).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).DiscoveryRecordModel(discoveryRecordModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryDAAS.UpdateDiscovery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDiscovery`: string
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryDAAS.UpdateDiscovery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App Name. [&#39;workspace&#39;] | 
**domain** | **string** | Domain used for user&#39;s email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDiscoveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 


 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 
 **discoveryRecordModel** | [**DiscoveryRecordModel**](DiscoveryRecordModel.md) |  | 

### Return type

**string**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

