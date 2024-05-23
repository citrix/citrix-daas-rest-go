# \DiscoveryControllerDAAS

All URIs are relative to *https://wsaca.cloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDiscoveryApiUsingDELETE**](DiscoveryControllerDAAS.md#DeleteDiscoveryApiUsingDELETE) | **Delete** /aca/discovery/app/{app}/domain/{domain} | Delete Discovery Record For App And Email Domain Suffix.
[**GetAllDiscoveryApiUsingGET**](DiscoveryControllerDAAS.md#GetAllDiscoveryApiUsingGET) | **Get** /aca/discovery | Retrieve All Discovery Records For The Specified Customer Id
[**GetDiscoveryApiUsingGET**](DiscoveryControllerDAAS.md#GetDiscoveryApiUsingGET) | **Get** /aca/discovery/app/{app}/domain/{domain} | Retrieve Discovery Record For App And Email Domain Suffix.
[**PostDiscoveryApiUsingPOST**](DiscoveryControllerDAAS.md#PostDiscoveryApiUsingPOST) | **Post** /aca/discovery/app/{app}/domain | Create A New Discovery Record.
[**PutDiscoveryApiUsingPUT**](DiscoveryControllerDAAS.md#PutDiscoveryApiUsingPUT) | **Put** /aca/discovery/app/{app}/domain/{domain} | Update Discovery Record For An App And Email Domain Suffix.



## DeleteDiscoveryApiUsingDELETE

> DeleteDiscoveryApiUsingDELETE(ctx, app, domain).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Delete Discovery Record For App And Email Domain Suffix.

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
	accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
	authorization := "CWSAuth bearer="<bearer_token>"" // string | Valid Citrix Cloud Authorization header.
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	app := "app_example" // string | App Name. ['Workspace']
	domain := "domain_example" // string | Domain used for user's email address.
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DiscoveryControllerDAAS.DeleteDiscoveryApiUsingDELETE(context.Background(), app, domain).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryControllerDAAS.DeleteDiscoveryApiUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App Name. [&#39;Workspace&#39;] | 
**domain** | **string** | Domain used for user&#39;s email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDiscoveryApiUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Valid Citrix Cloud Authorization header. | 
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 


 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

 (empty response body)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllDiscoveryApiUsingGET

> GetAllDiscoveryResponse GetAllDiscoveryApiUsingGET(ctx).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Retrieve All Discovery Records For The Specified Customer Id

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
	accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
	authorization := "CWSAuth bearer="<bearer_token>"" // string | Valid Citrix Cloud Authorization header.
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscoveryControllerDAAS.GetAllDiscoveryApiUsingGET(context.Background()).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryControllerDAAS.GetAllDiscoveryApiUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllDiscoveryApiUsingGET`: GetAllDiscoveryResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryControllerDAAS.GetAllDiscoveryApiUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllDiscoveryApiUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Valid Citrix Cloud Authorization header. | 
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

[**GetAllDiscoveryResponse**](GetAllDiscoveryResponse.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscoveryApiUsingGET

> GetAllDiscoveryResponse GetDiscoveryApiUsingGET(ctx, app, domain).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Retrieve Discovery Record For App And Email Domain Suffix.

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
	accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
	authorization := "CWSAuth bearer="<bearer_token>"" // string | Valid Citrix Cloud Authorization header.
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	app := "app_example" // string | App Name. ['Workspace']
	domain := "domain_example" // string | Domain used for user's email address.
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscoveryControllerDAAS.GetDiscoveryApiUsingGET(context.Background(), app, domain).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryControllerDAAS.GetDiscoveryApiUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDiscoveryApiUsingGET`: GetAllDiscoveryResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryControllerDAAS.GetDiscoveryApiUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App Name. [&#39;Workspace&#39;] | 
**domain** | **string** | Domain used for user&#39;s email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscoveryApiUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Valid Citrix Cloud Authorization header. | 
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 


 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

[**GetAllDiscoveryResponse**](GetAllDiscoveryResponse.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDiscoveryApiUsingPOST

> DiscoveryRecordModel PostDiscoveryApiUsingPOST(ctx, app).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).DiscoveryRecord(discoveryRecord).CitrixTransactionId(citrixTransactionId).Execute()

Create A New Discovery Record.

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
	accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
	authorization := "CWSAuth bearer="<bearer_token>"" // string | Valid Citrix Cloud Authorization header.
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	app := "app_example" // string | App Name. ['Workspace']
	discoveryRecord := *openapiclient.NewDiscoveryRecordModel() // DiscoveryRecordModel | Discovery record to be created.
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscoveryControllerDAAS.PostDiscoveryApiUsingPOST(context.Background(), app).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).DiscoveryRecord(discoveryRecord).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryControllerDAAS.PostDiscoveryApiUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDiscoveryApiUsingPOST`: DiscoveryRecordModel
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryControllerDAAS.PostDiscoveryApiUsingPOST`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App Name. [&#39;Workspace&#39;] | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDiscoveryApiUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Valid Citrix Cloud Authorization header. | 
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 

 **discoveryRecord** | [**DiscoveryRecordModel**](DiscoveryRecordModel.md) | Discovery record to be created. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

[**DiscoveryRecordModel**](DiscoveryRecordModel.md)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDiscoveryApiUsingPUT

> PutDiscoveryApiUsingPUT(ctx, app, domain).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).DiscoveryRecord(discoveryRecord).CitrixTransactionId(citrixTransactionId).Execute()

Update Discovery Record For An App And Email Domain Suffix.

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
	accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
	authorization := "CWSAuth bearer="<bearer_token>"" // string | Valid Citrix Cloud Authorization header.
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
	app := "app_example" // string | App Name. ['Workspace']
	domain := "domain_example" // string | Domain used for user's email address.
	discoveryRecord := *openapiclient.NewDiscoveryRecordModel() // DiscoveryRecordModel | Discovery record to be created.
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DiscoveryControllerDAAS.PutDiscoveryApiUsingPUT(context.Background(), app, domain).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).DiscoveryRecord(discoveryRecord).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryControllerDAAS.PutDiscoveryApiUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App Name. [&#39;Workspace&#39;] | 
**domain** | **string** | Domain used for user&#39;s email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDiscoveryApiUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Valid Citrix Cloud Authorization header. | 
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 


 **discoveryRecord** | [**DiscoveryRecordModel**](DiscoveryRecordModel.md) | Discovery record to be created. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | 

### Return type

 (empty response body)

### Authorization

[BearerTokenAuth](../README.md#BearerTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

