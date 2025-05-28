# \SiteDAAS

All URIs are relative to *https://.wem.cloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SiteCreate**](SiteDAAS.md#SiteCreate) | **Post** /services/wem/sites | Create a Site
[**SiteDelete**](SiteDAAS.md#SiteDelete) | **Delete** /services/wem/sites/{id} | Delete a Site
[**SiteExport**](SiteDAAS.md#SiteExport) | **Get** /services/wem/sites/$export | Export a Site
[**SiteImport**](SiteDAAS.md#SiteImport) | **Post** /services/wem/sites/$import | Import a Site
[**SiteQuery**](SiteDAAS.md#SiteQuery) | **Get** /services/wem/sites | Query Sites
[**SiteQueryById**](SiteDAAS.md#SiteQueryById) | **Get** /services/wem/sites/{id} | Query Sites by Identity
[**SiteReplicate**](SiteDAAS.md#SiteReplicate) | **Post** /services/wem/sites/$replicate | Replicate a Site
[**SiteUpdate**](SiteDAAS.md#SiteUpdate) | **Put** /services/wem/sites | Update a Site



## SiteCreate

> SiteCreate(ctx).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Create a Site

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/devicemanagement"
)

func main() {
	accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
	authorization := "authorization_example" // string | Citrix Cloud authorization token.
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud customer ID.
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional) (default to "235ca558-a454-42ee-be58-54f57a0b87cc")
	body := *openapiclient.NewSiteModel() // SiteModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SiteDAAS.SiteCreate(context.Background()).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteDAAS.SiteCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSiteCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Citrix Cloud authorization token. | 
 **citrixCustomerId** | **string** | Citrix Cloud customer ID. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | [default to &quot;235ca558-a454-42ee-be58-54f57a0b87cc&quot;]
 **body** | [**SiteModel**](SiteModel.md) |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SiteDelete

> SiteDelete(ctx, id).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Delete a Site

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/devicemanagement"
)

func main() {
	id := int64(789) // int64 | The ID of the resource you want to delete.
	accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
	authorization := "authorization_example" // string | Citrix Cloud authorization token.
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud customer ID.
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional) (default to "c64a0e4a-426e-4271-ab2d-49fa019d8de3")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SiteDAAS.SiteDelete(context.Background(), id).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteDAAS.SiteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource you want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSiteDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Citrix Cloud authorization token. | 
 **citrixCustomerId** | **string** | Citrix Cloud customer ID. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | [default to &quot;c64a0e4a-426e-4271-ab2d-49fa019d8de3&quot;]

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SiteExport

> map[string]interface{} SiteExport(ctx).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Id(id).Name(name).CitrixTransactionId(citrixTransactionId).Execute()

Export a Site

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/devicemanagement"
)

func main() {
	accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
	authorization := "authorization_example" // string | Citrix Cloud authorization token.
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud customer ID.
	id := int64(789) // int64 | The ID of the resource you want to export. (optional)
	name := "name_example" // string | The name of the resource you want to export. (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional) (default to "2bd401c7-e7d0-4d18-90e2-9d92584ddde3")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteDAAS.SiteExport(context.Background()).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Id(id).Name(name).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteDAAS.SiteExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SiteExport`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SiteDAAS.SiteExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSiteExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Citrix Cloud authorization token. | 
 **citrixCustomerId** | **string** | Citrix Cloud customer ID. | 
 **id** | **int64** | The ID of the resource you want to export. | 
 **name** | **string** | The name of the resource you want to export. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | [default to &quot;2bd401c7-e7d0-4d18-90e2-9d92584ddde3&quot;]

### Return type

**map[string]interface{}**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SiteImport

> SiteImport(ctx).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Import a Site

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/devicemanagement"
)

func main() {
	accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
	authorization := "authorization_example" // string | Citrix Cloud authorization token.
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud customer ID.
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional) (default to "d0948dc7-4861-45b9-8ae2-b5b40229237a")
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SiteDAAS.SiteImport(context.Background()).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteDAAS.SiteImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSiteImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Citrix Cloud authorization token. | 
 **citrixCustomerId** | **string** | Citrix Cloud customer ID. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | [default to &quot;d0948dc7-4861-45b9-8ae2-b5b40229237a&quot;]
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SiteQuery

> SiteQuery200Response SiteQuery(ctx).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Name(name).Id(id).CitrixTransactionId(citrixTransactionId).Execute()

Query Sites

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/devicemanagement"
)

func main() {
	accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
	authorization := "authorization_example" // string | Citrix Cloud authorization token.
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud customer ID.
	name := "name_example" // string | The name is used to filter the query results. (optional)
	id := int64(789) // int64 | The ID is used to filter the query results. (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional) (default to "cecfb7dd-4c50-4b11-b4d5-6688946f69ee")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteDAAS.SiteQuery(context.Background()).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Name(name).Id(id).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteDAAS.SiteQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SiteQuery`: SiteQuery200Response
	fmt.Fprintf(os.Stdout, "Response from `SiteDAAS.SiteQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSiteQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Citrix Cloud authorization token. | 
 **citrixCustomerId** | **string** | Citrix Cloud customer ID. | 
 **name** | **string** | The name is used to filter the query results. | 
 **id** | **int64** | The ID is used to filter the query results. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | [default to &quot;cecfb7dd-4c50-4b11-b4d5-6688946f69ee&quot;]

### Return type

[**SiteQuery200Response**](SiteQuery200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8, items

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SiteQueryById

> SiteModel SiteQueryById(ctx, id).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Query Sites by Identity

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/devicemanagement"
)

func main() {
	id := int64(789) // int64 | The ID of the resource you want to query.
	accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
	authorization := "authorization_example" // string | Citrix Cloud authorization token.
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud customer ID.
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional) (default to "b6155b53-555d-4898-a4bb-3540339a1ac7")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteDAAS.SiteQueryById(context.Background(), id).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteDAAS.SiteQueryById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SiteQueryById`: SiteModel
	fmt.Fprintf(os.Stdout, "Response from `SiteDAAS.SiteQueryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource you want to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSiteQueryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Citrix Cloud authorization token. | 
 **citrixCustomerId** | **string** | Citrix Cloud customer ID. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | [default to &quot;b6155b53-555d-4898-a4bb-3540339a1ac7&quot;]

### Return type

[**SiteModel**](SiteModel.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8, id, name, description

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SiteReplicate

> SiteReplicate(ctx).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Id(id).Name(name).CitrixTransactionId(citrixTransactionId).Execute()

Replicate a Site

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/devicemanagement"
)

func main() {
	accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
	authorization := "authorization_example" // string | Citrix Cloud authorization token.
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud customer ID.
	id := int64(789) // int64 |  (optional)
	name := "name_example" // string |  (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional) (default to "ea86ad5e-3bf7-4d95-98c7-6993e25a153e")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SiteDAAS.SiteReplicate(context.Background()).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Id(id).Name(name).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteDAAS.SiteReplicate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSiteReplicateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Citrix Cloud authorization token. | 
 **citrixCustomerId** | **string** | Citrix Cloud customer ID. | 
 **id** | **int64** |  | 
 **name** | **string** |  | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | [default to &quot;ea86ad5e-3bf7-4d95-98c7-6993e25a153e&quot;]

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SiteUpdate

> SiteUpdate(ctx).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Update a Site

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/devicemanagement"
)

func main() {
	accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
	authorization := "authorization_example" // string | Citrix Cloud authorization token.
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud customer ID.
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional) (default to "773b417e-4dc8-4312-86aa-bcf7e7266467")
	body := *openapiclient.NewSiteModel() // SiteModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SiteDAAS.SiteUpdate(context.Background()).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteDAAS.SiteUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSiteUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Citrix Cloud authorization token. | 
 **citrixCustomerId** | **string** | Citrix Cloud customer ID. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | [default to &quot;773b417e-4dc8-4312-86aa-bcf7e7266467&quot;]
 **body** | [**SiteModel**](SiteModel.md) |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

