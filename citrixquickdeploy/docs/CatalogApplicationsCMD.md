# \CatalogApplicationsCMD

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCatalogApps**](CatalogApplicationsCMD.md#GetCatalogApps) | **Get** /{customerId}/{siteId}/catalogs/{catalogId}/apps | Get the list of apps that have been published for the specified catalog
[**PublishApplications**](CatalogApplicationsCMD.md#PublishApplications) | **Post** /{customerId}/{siteId}/catalogs/{catalogId}/apps | Publish the specified apps to the catalog
[**UnpublishApplication**](CatalogApplicationsCMD.md#UnpublishApplication) | **Delete** /{customerId}/{siteId}/catalogs/{catalogId}/apps/{appId} | Remove an application from the published list
[**UnpublishApplications**](CatalogApplicationsCMD.md#UnpublishApplications) | **Delete** /{customerId}/{siteId}/catalogs/{catalogId}/apps | Remove a list of applications from a catalog
[**UpdateApplication**](CatalogApplicationsCMD.md#UpdateApplication) | **Put** /{customerId}/{siteId}/catalogs/{catalogId}/apps/{appId} | Update the configuration of a published app



## GetCatalogApps

> CatalogApplicationsModel GetCatalogApps(ctx, customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).Execute()

Get the list of apps that have been published for the specified catalog

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
    resp, r, err := apiClient.CatalogApplicationsCMD.GetCatalogApps(context.Background(), customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApplicationsCMD.GetCatalogApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCatalogApps`: CatalogApplicationsModel
    fmt.Fprintf(os.Stdout, "Response from `CatalogApplicationsCMD.GetCatalogApps`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetCatalogAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**CatalogApplicationsModel**](CatalogApplicationsModel.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishApplications

> CatalogApplicationsModel PublishApplications(ctx, customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Publish the specified apps to the catalog

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
    body := *openapiclient.NewAddCatalogApplicationsModel([]openapiclient.AddCatalogApplicationModel{*openapiclient.NewAddCatalogApplicationModel("Name_example", "Path_example")}) // AddCatalogApplicationsModel | List of applications to add (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogApplicationsCMD.PublishApplications(context.Background(), customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApplicationsCMD.PublishApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublishApplications`: CatalogApplicationsModel
    fmt.Fprintf(os.Stdout, "Response from `CatalogApplicationsCMD.PublishApplications`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPublishApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 
 **body** | [**AddCatalogApplicationsModel**](AddCatalogApplicationsModel.md) | List of applications to add | 

### Return type

[**CatalogApplicationsModel**](CatalogApplicationsModel.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnpublishApplication

> UnpublishApplication(ctx, customerId, siteId, catalogId, appId).CitrixTransactionId(citrixTransactionId).Execute()

Remove an application from the published list

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
    appId := "appId_example" // string | Identifier of the application
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CatalogApplicationsCMD.UnpublishApplication(context.Background(), customerId, siteId, catalogId, appId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApplicationsCMD.UnpublishApplication``: %v\n", err)
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
**appId** | **string** | Identifier of the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnpublishApplicationRequest struct via the builder pattern


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


## UnpublishApplications

> UnpublishApplications(ctx, customerId, siteId, catalogId).AppIds(appIds).CitrixTransactionId(citrixTransactionId).Execute()

Remove a list of applications from a catalog

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
    appIds := []string{"Inner_example"} // []string | List of application ids to remove (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CatalogApplicationsCMD.UnpublishApplications(context.Background(), customerId, siteId, catalogId).AppIds(appIds).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApplicationsCMD.UnpublishApplications``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUnpublishApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **appIds** | **[]string** | List of application ids to remove | 
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


## UpdateApplication

> UpdateApplication(ctx, customerId, siteId, catalogId, appId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Update the configuration of a published app

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
    appId := "appId_example" // string | Identifier of the application
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    body := *openapiclient.NewUpdateApplicationConfigurationModel("Name_example", "ApplicationPath_example") // UpdateApplicationConfigurationModel | New configuration of the application (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CatalogApplicationsCMD.UpdateApplication(context.Background(), customerId, siteId, catalogId, appId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApplicationsCMD.UpdateApplication``: %v\n", err)
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
**appId** | **string** | Identifier of the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **citrixTransactionId** | **string** | The Transaction Id. | 
 **body** | [**UpdateApplicationConfigurationModel**](UpdateApplicationConfigurationModel.md) | New configuration of the application | 

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

