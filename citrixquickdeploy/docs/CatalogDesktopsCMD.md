# \CatalogDesktopsCMD

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCatalogDesktops**](CatalogDesktopsCMD.md#GetCatalogDesktops) | **Get** /{customerId}/{siteId}/catalogs/{catalogId}/desktops | Get the list of desktops that have been published for the specified catalog
[**PublishDesktops**](CatalogDesktopsCMD.md#PublishDesktops) | **Post** /{customerId}/{siteId}/catalogs/{catalogId}/desktops | Publish the specified desktops to the catalog
[**UnpublishDesktop**](CatalogDesktopsCMD.md#UnpublishDesktop) | **Delete** /{customerId}/{siteId}/catalogs/{catalogId}/desktops/{desktopId} | Remove a desktop from the published list
[**UnpublishDesktops**](CatalogDesktopsCMD.md#UnpublishDesktops) | **Delete** /{customerId}/{siteId}/catalogs/{catalogId}/desktops | Remove a list of desktops from a catalog
[**UpdateDesktop**](CatalogDesktopsCMD.md#UpdateDesktop) | **Put** /{customerId}/{siteId}/catalogs/{catalogId}/desktops/{desktopId} | Update the configuration of a published desktop



## GetCatalogDesktops

> CatalogDesktopsModel GetCatalogDesktops(ctx, customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).Execute()

Get the list of desktops that have been published for the specified catalog

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixquickdeploy"
)

func main() {
    customerId := "customerId_example" // string | ID of the customer
    siteId := "siteId_example" // string | The site ID of the customer
    catalogId := "catalogId_example" // string | ID of the catalog
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogDesktopsCMD.GetCatalogDesktops(context.Background(), customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogDesktopsCMD.GetCatalogDesktops``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCatalogDesktops`: CatalogDesktopsModel
    fmt.Fprintf(os.Stdout, "Response from `CatalogDesktopsCMD.GetCatalogDesktops`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**siteId** | **string** | The site ID of the customer | 
**catalogId** | **string** | ID of the catalog | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogDesktopsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**CatalogDesktopsModel**](CatalogDesktopsModel.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishDesktops

> CatalogDesktopsModel PublishDesktops(ctx, customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Publish the specified desktops to the catalog

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixquickdeploy"
)

func main() {
    customerId := "customerId_example" // string | ID of the customer
    siteId := "siteId_example" // string | The site ID of the customer
    catalogId := "catalogId_example" // string | ID of the catalog
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    body := *openapiclient.NewAddCatalogDesktopsModel([]openapiclient.AddCatalogDesktopModel{*openapiclient.NewAddCatalogDesktopModel("Name_example")}) // AddCatalogDesktopsModel | List of desktops to add (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogDesktopsCMD.PublishDesktops(context.Background(), customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogDesktopsCMD.PublishDesktops``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublishDesktops`: CatalogDesktopsModel
    fmt.Fprintf(os.Stdout, "Response from `CatalogDesktopsCMD.PublishDesktops`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**siteId** | **string** | The site ID of the customer | 
**catalogId** | **string** | ID of the catalog | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishDesktopsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 
 **body** | [**AddCatalogDesktopsModel**](AddCatalogDesktopsModel.md) | List of desktops to add | 

### Return type

[**CatalogDesktopsModel**](CatalogDesktopsModel.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnpublishDesktop

> UnpublishDesktop(ctx, customerId, siteId, catalogId, desktopId).CitrixTransactionId(citrixTransactionId).Execute()

Remove a desktop from the published list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixquickdeploy"
)

func main() {
    customerId := "customerId_example" // string | ID of the customer
    siteId := "siteId_example" // string | 
    catalogId := "catalogId_example" // string | ID of the catalog
    desktopId := "desktopId_example" // string | Identifier of the desktop
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CatalogDesktopsCMD.UnpublishDesktop(context.Background(), customerId, siteId, catalogId, desktopId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogDesktopsCMD.UnpublishDesktop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**siteId** | **string** |  | 
**catalogId** | **string** | ID of the catalog | 
**desktopId** | **string** | Identifier of the desktop | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnpublishDesktopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

 (empty response body)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnpublishDesktops

> UnpublishDesktops(ctx, customerId, siteId, catalogId).DesktopIds(desktopIds).CitrixTransactionId(citrixTransactionId).Execute()

Remove a list of desktops from a catalog

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixquickdeploy"
)

func main() {
    customerId := "customerId_example" // string | ID of the customer
    siteId := "siteId_example" // string | The site ID of the customer
    catalogId := "catalogId_example" // string | ID of the catalog
    desktopIds := []string{"Inner_example"} // []string | List of desktop ids to remove (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CatalogDesktopsCMD.UnpublishDesktops(context.Background(), customerId, siteId, catalogId).DesktopIds(desktopIds).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogDesktopsCMD.UnpublishDesktops``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**siteId** | **string** | The site ID of the customer | 
**catalogId** | **string** | ID of the catalog | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnpublishDesktopsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **desktopIds** | **[]string** | List of desktop ids to remove | 
 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

 (empty response body)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDesktop

> UpdateDesktop(ctx, customerId, siteId, catalogId, desktopId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Update the configuration of a published desktop

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixquickdeploy"
)

func main() {
    customerId := "customerId_example" // string | ID of the customer
    siteId := "siteId_example" // string | The site ID of the customer
    catalogId := "catalogId_example" // string | ID of the catalog
    desktopId := "desktopId_example" // string | Identifier of the desktop
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    body := *openapiclient.NewUpdateDesktopConfigurationModel("Name_example") // UpdateDesktopConfigurationModel | New configuration of the application (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CatalogDesktopsCMD.UpdateDesktop(context.Background(), customerId, siteId, catalogId, desktopId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogDesktopsCMD.UpdateDesktop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**siteId** | **string** | The site ID of the customer | 
**catalogId** | **string** | ID of the catalog | 
**desktopId** | **string** | Identifier of the desktop | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDesktopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **citrixTransactionId** | **string** | The Transaction Id. | 
 **body** | [**UpdateDesktopConfigurationModel**](UpdateDesktopConfigurationModel.md) | New configuration of the application | 

### Return type

 (empty response body)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

