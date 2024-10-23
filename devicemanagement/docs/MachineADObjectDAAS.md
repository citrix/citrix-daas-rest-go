# \MachineADObjectDAAS

All URIs are relative to *https://api.wem.cloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdObjectCreate**](MachineADObjectDAAS.md#AdObjectCreate) | **Post** /services/wem/machines | Create a Machine-level AD Object
[**AdObjectDelete**](MachineADObjectDAAS.md#AdObjectDelete) | **Delete** /services/wem/machines/{id} | Delete a Machine-level AD Object
[**AdObjectQuery**](MachineADObjectDAAS.md#AdObjectQuery) | **Get** /services/wem/machines | Query Machine-level AD Objects
[**AdObjectQueryById**](MachineADObjectDAAS.md#AdObjectQueryById) | **Get** /services/wem/machines/{id} | Query Machine-level AD Objects by Identity
[**AdObjectUpdate**](MachineADObjectDAAS.md#AdObjectUpdate) | **Put** /services/wem/machines | Update a Machine-level AD Object



## AdObjectCreate

> AdObjectCreate(ctx).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Create a Machine-level AD Object

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
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional) (default to "932fb487-f91a-44f7-8b98-a4baafd2a0ed")
    body := MachineModel(987) // MachineModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineADObjectDAAS.AdObjectCreate(context.Background()).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineADObjectDAAS.AdObjectCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdObjectCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Citrix Cloud authorization token. | 
 **citrixCustomerId** | **string** | Citrix Cloud customer ID. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | [default to &quot;932fb487-f91a-44f7-8b98-a4baafd2a0ed&quot;]
 **body** | **MachineModel** |  | 

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


## AdObjectDelete

> AdObjectDelete(ctx, id).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Delete a Machine-level AD Object

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
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional) (default to "8bb6abc3-9a37-4742-8808-024f13c0e4d0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineADObjectDAAS.AdObjectDelete(context.Background(), id).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineADObjectDAAS.AdObjectDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdObjectDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Citrix Cloud authorization token. | 
 **citrixCustomerId** | **string** | Citrix Cloud customer ID. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | [default to &quot;8bb6abc3-9a37-4742-8808-024f13c0e4d0&quot;]

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


## AdObjectQuery

> AdObjectQuery200Response AdObjectQuery(ctx).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Sid(sid).SiteId(siteId).Id(id).CitrixTransactionId(citrixTransactionId).Execute()

Query Machine-level AD Objects

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
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional) (default to "ad159f5b-8961-4c6a-b3db-f2d073eb77b5")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineADObjectDAAS.AdObjectQuery(context.Background()).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Sid(sid).SiteId(siteId).Id(id).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineADObjectDAAS.AdObjectQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdObjectQuery`: AdObjectQuery200Response
    fmt.Fprintf(os.Stdout, "Response from `MachineADObjectDAAS.AdObjectQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdObjectQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Citrix Cloud authorization token. | 
 **citrixCustomerId** | **string** | Citrix Cloud customer ID. | 
 **sid** | **string** | The SID is used to filter the query results. | 
 **siteId** | **int64** | The name is used to filter the query results. | 
 **id** | **int64** | The ID is used to filter the query results. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | [default to &quot;ad159f5b-8961-4c6a-b3db-f2d073eb77b5&quot;]

### Return type

[**AdObjectQuery200Response**](AdObjectQuery200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8, items

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdObjectQueryById

> MachineModel AdObjectQueryById(ctx, id).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Query Machine-level AD Objects by Identity

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
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional) (default to "c8b677e6-b18c-4a7b-a2fa-b16e671cf430")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineADObjectDAAS.AdObjectQueryById(context.Background(), id).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineADObjectDAAS.AdObjectQueryById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdObjectQueryById`: MachineModel
    fmt.Fprintf(os.Stdout, "Response from `MachineADObjectDAAS.AdObjectQueryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource you want to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdObjectQueryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Citrix Cloud authorization token. | 
 **citrixCustomerId** | **string** | Citrix Cloud customer ID. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | [default to &quot;c8b677e6-b18c-4a7b-a2fa-b16e671cf430&quot;]

### Return type

[**MachineModel**](MachineModel.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8, id, sid, name, type, description, siteId, enabled, priority

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdObjectUpdate

> AdObjectUpdate(ctx).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Update a Machine-level AD Object

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
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional) (default to "d3116ffe-b620-4162-bc82-ac5f66895692")
    body := MachineModel(987) // MachineModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MachineADObjectDAAS.AdObjectUpdate(context.Background()).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineADObjectDAAS.AdObjectUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdObjectUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Citrix Cloud authorization token. | 
 **citrixCustomerId** | **string** | Citrix Cloud customer ID. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | [default to &quot;d3116ffe-b620-4162-bc82-ac5f66895692&quot;]
 **body** | **MachineModel** |  | 

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

