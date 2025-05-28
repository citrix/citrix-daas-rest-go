# \UserADObjectDAAS

All URIs are relative to *https://.wem.cloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ItemCreate**](UserADObjectDAAS.md#ItemCreate) | **Post** /services/wem/users | Create a User-level AD Object
[**ItemDelete**](UserADObjectDAAS.md#ItemDelete) | **Delete** /services/wem/users/{id} | Delete a User-level AD Object
[**ItemQuery**](UserADObjectDAAS.md#ItemQuery) | **Get** /services/wem/users | Query User-level AD Objects
[**ItemQueryById**](UserADObjectDAAS.md#ItemQueryById) | **Get** /services/wem/users/{id} | Query User-level AD Objects by Identity
[**ItemUpdate**](UserADObjectDAAS.md#ItemUpdate) | **Put** /services/wem/users | Update a User-level AD Object



## ItemCreate

> ItemCreate(ctx).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Create a User-level AD Object

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
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional) (default to "6dcad3fb-47ce-45fd-acf4-f4261bc6c590")
	body := *openapiclient.NewUserModel("Sid_example", "Name_example", "Type_example", int64(123)) // UserModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserADObjectDAAS.ItemCreate(context.Background()).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserADObjectDAAS.ItemCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiItemCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Citrix Cloud authorization token. | 
 **citrixCustomerId** | **string** | Citrix Cloud customer ID. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | [default to &quot;6dcad3fb-47ce-45fd-acf4-f4261bc6c590&quot;]
 **body** | [**UserModel**](UserModel.md) |  | 

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


## ItemDelete

> ItemDelete(ctx, id).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Delete a User-level AD Object

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
	id := "id_example" // string | The ID of the resource you want to delete.
	accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
	authorization := "authorization_example" // string | Citrix Cloud authorization token.
	citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud customer ID.
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional) (default to "5dcc1325-dab6-4c13-8388-3dc17fc482c8")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserADObjectDAAS.ItemDelete(context.Background(), id).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserADObjectDAAS.ItemDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the resource you want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiItemDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Citrix Cloud authorization token. | 
 **citrixCustomerId** | **string** | Citrix Cloud customer ID. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | [default to &quot;5dcc1325-dab6-4c13-8388-3dc17fc482c8&quot;]

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


## ItemQuery

> ItemQuery200Response ItemQuery(ctx).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Sid(sid).SiteId(siteId).Id(id).CitrixTransactionId(citrixTransactionId).Execute()

Query User-level AD Objects

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
	sid := "sid_example" // string | The SID is used to filter the query results. (optional)
	siteId := int64(789) // int64 | The name is used to filter the query results. (optional)
	id := int64(789) // int64 | The ID is used to filter the query results. (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional) (default to "1d20c409-3436-4eee-b0e6-35218b74fef9")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserADObjectDAAS.ItemQuery(context.Background()).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Sid(sid).SiteId(siteId).Id(id).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserADObjectDAAS.ItemQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ItemQuery`: ItemQuery200Response
	fmt.Fprintf(os.Stdout, "Response from `UserADObjectDAAS.ItemQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiItemQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Citrix Cloud authorization token. | 
 **citrixCustomerId** | **string** | Citrix Cloud customer ID. | 
 **sid** | **string** | The SID is used to filter the query results. | 
 **siteId** | **int64** | The name is used to filter the query results. | 
 **id** | **int64** | The ID is used to filter the query results. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | [default to &quot;1d20c409-3436-4eee-b0e6-35218b74fef9&quot;]

### Return type

[**ItemQuery200Response**](ItemQuery200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8, items

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ItemQueryById

> UserModel ItemQueryById(ctx, id).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Query User-level AD Objects by Identity

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
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional) (default to "8f8c71a8-5839-40ec-bff2-c1ed61b012f5")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserADObjectDAAS.ItemQueryById(context.Background(), id).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserADObjectDAAS.ItemQueryById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ItemQueryById`: UserModel
	fmt.Fprintf(os.Stdout, "Response from `UserADObjectDAAS.ItemQueryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource you want to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiItemQueryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Citrix Cloud authorization token. | 
 **citrixCustomerId** | **string** | Citrix Cloud customer ID. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | [default to &quot;8f8c71a8-5839-40ec-bff2-c1ed61b012f5&quot;]

### Return type

[**UserModel**](UserModel.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8, id, sid, name, type, description, siteId, enabled, priority

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ItemUpdate

> ItemUpdate(ctx).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Update a User-level AD Object

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
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional) (default to "960f4f7f-73dd-421a-9baa-03f99c5df87e")
	body := *openapiclient.NewUserModel("Sid_example", "Name_example", "Type_example", int64(123)) // UserModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserADObjectDAAS.ItemUpdate(context.Background()).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserADObjectDAAS.ItemUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiItemUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Citrix Cloud authorization token. | 
 **citrixCustomerId** | **string** | Citrix Cloud customer ID. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | [default to &quot;960f4f7f-73dd-421a-9baa-03f99c5df87e&quot;]
 **body** | [**UserModel**](UserModel.md) |  | 

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

