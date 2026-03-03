# \CatalogCMD

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRemotePcCatalogMachineAssignments**](CatalogCMD.md#AddRemotePcCatalogMachineAssignments) | **Post** /{customerId}/{siteId}/catalogs/{catalogId}/machinesassignments | Add machine assignments to a remote pc catalog.  New machines will be added to the catalog with the specified users  Existing machines will be updated with the included assignments
[**ConfigureAndDeployCitrixManagedCatalogApi**](CatalogCMD.md#ConfigureAndDeployCitrixManagedCatalogApi) | **Post** /{customerId}/{siteId}/catalogs/$manageddeploy | Configure all the Citrix managed catalog deployment steps and initiate the catalog deployment
[**CreateRemotePcCatalog**](CatalogCMD.md#CreateRemotePcCatalog) | **Post** /{customerId}/{siteId}/catalogs/remotePc | Create a Remote PC catalog.
[**DeleteCustomerCatalog**](CatalogCMD.md#DeleteCustomerCatalog) | **Delete** /{customerId}/{siteId}/catalogs/{catalogId} | Delete a catalog along with all the pubhishd apps for the catalog
[**GetAllTroubleshootScripts**](CatalogCMD.md#GetAllTroubleshootScripts) | **Get** /{customerId}/catalogs/{catalogId}/jobs/troubleshoot/{machineName} | Get the statuses of all the troubleshoot scripts assocaited with a vda
[**GetAllTroubleshootScripts_0**](CatalogCMD.md#GetAllTroubleshootScripts_0) | **Get** /{customerId}/{siteId}/catalogs/{catalogId}/jobs/troubleshoot/{machineName} | Get the statuses of all the troubleshoot scripts associated with a vda
[**GetCatalogCapacityConfiguration**](CatalogCMD.md#GetCatalogCapacityConfiguration) | **Get** /{customerId}/{siteId}/catalogs/{catalogId}/capacity | Get the performance information configured for this catalog
[**GetCustomerCatalog**](CatalogCMD.md#GetCustomerCatalog) | **Get** /{customerId}/{siteId}/catalogs/{catalogId} | Returns a specific catalog for a specific customer
[**GetCustomerCatalogs**](CatalogCMD.md#GetCustomerCatalogs) | **Get** /{customerId}/{siteId}/catalogs | Returns all the catalogs that the specified customer has created
[**GetCustomerManagedCatalogs**](CatalogCMD.md#GetCustomerManagedCatalogs) | **Get** /{customerId}/{siteId}/managedcatalogs | Returns all the catalogs that the specified Citrix managed customer has created
[**GetCustomerManagedCatalogsById**](CatalogCMD.md#GetCustomerManagedCatalogsById) | **Get** /{customerId}/{siteId}/managedcatalogs/{catalogId} | Returns all the catalogs that the specified Citrix managed customer has created
[**GetScriptsAndParameters**](CatalogCMD.md#GetScriptsAndParameters) | **Get** /{customerId}/catalogs/{catalogId}/scripts | Gets available scripts and corresponding default parameters for available operations
[**GetScriptsAndParameters_0**](CatalogCMD.md#GetScriptsAndParameters_0) | **Get** /{customerId}/{siteId}/catalogs/{catalogId}/scripts | Gets available scripts and corresponding default parameters for available operations
[**GetTroubleshootScriptStatus**](CatalogCMD.md#GetTroubleshootScriptStatus) | **Get** /{customerId}/catalogs/{catalogId}/jobs/troubleshoot/{machineName}/{troubleshootId} | Get the status of the troubleshoot script
[**GetTroubleshootScriptStatus_0**](CatalogCMD.md#GetTroubleshootScriptStatus_0) | **Get** /{customerId}/{siteId}/catalogs/{catalogId}/jobs/troubleshoot/{machineName}/{troubleshootId} | Get the status of the troubleshoot script
[**GetTroubleshootScriptsOnlyForFailedConnectorInstalls**](CatalogCMD.md#GetTroubleshootScriptsOnlyForFailedConnectorInstalls) | **Get** /{customerId}/catalogs/{catalogId}/jobs/troubleshoot/failedConnectorInstalls | Get the statuses of all the troubleshoot scripts assocaited with catalog connectors
[**GetTroubleshootScriptsOnlyForFailedConnectorInstalls_0**](CatalogCMD.md#GetTroubleshootScriptsOnlyForFailedConnectorInstalls_0) | **Get** /{customerId}/{siteId}/catalogs/{catalogId}/jobs/troubleshoot/failedConnectorInstalls | Get the statuses of all the troubleshoot scripts associated with catalog connectors
[**RunScriptOnVm**](CatalogCMD.md#RunScriptOnVm) | **Post** /{customerId}/catalogs/{catalogId}/vmresources/{machineName}/$runscriptWithModel | Run a troubleshoot script on a VDA
[**RunScriptOnVm_0**](CatalogCMD.md#RunScriptOnVm_0) | **Post** /{customerId}/{siteId}/catalogs/{catalogId}/vmresources/{machineName}/$runscriptWithModel | Run a troubleshoot script on a VDA
[**UpdateCatalogImage**](CatalogCMD.md#UpdateCatalogImage) | **Post** /{customerId}/{siteId}/catalogs/{catalogId}/updateImage | Update the catalog&#39;s master image.
[**UpdateCatalogImageApi**](CatalogCMD.md#UpdateCatalogImageApi) | **Post** /{customerId}/{siteId}/catalogs/{catalogId}/$updateImage | Update the catalog&#39;s master image.
[**UpdateCatalogScaleConfiguration**](CatalogCMD.md#UpdateCatalogScaleConfiguration) | **Patch** /{customerId}/{siteId}/catalogs/{catalogId}/capacity | Update the performance information configured for this catalog
[**UpdateRemotePcCatalogScopes**](CatalogCMD.md#UpdateRemotePcCatalogScopes) | **Put** /{customerId}/{siteId}/catalogs/{catalogId}/scopes | Configure all the Citrix managed catalog deployment steps and initiate the catalog deployment



## AddRemotePcCatalogMachineAssignments

> string AddRemotePcCatalogMachineAssignments(ctx, customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).AddCatalogMachineAssignments(addCatalogMachineAssignments).Execute()

Add machine assignments to a remote pc catalog.  New machines will be added to the catalog with the specified users  Existing machines will be updated with the included assignments

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
    catalogId := "catalogId_example" // string | Id of the catalog to add machines to.
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    addCatalogMachineAssignments := *openapiclient.NewAddCatalogMachineAssignments([]openapiclient.AddCatalogMachineAssignment{*openapiclient.NewAddCatalogMachineAssignment()}) // AddCatalogMachineAssignments | The configuration of the machines to add. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogCMD.AddRemotePcCatalogMachineAssignments(context.Background(), customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).AddCatalogMachineAssignments(addCatalogMachineAssignments).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogCMD.AddRemotePcCatalogMachineAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRemotePcCatalogMachineAssignments`: string
    fmt.Fprintf(os.Stdout, "Response from `CatalogCMD.AddRemotePcCatalogMachineAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**siteId** | **string** | The site ID of the customer | 
**catalogId** | **string** | Id of the catalog to add machines to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRemotePcCatalogMachineAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 
 **addCatalogMachineAssignments** | [**AddCatalogMachineAssignments**](AddCatalogMachineAssignments.md) | The configuration of the machines to add. | 

### Return type

**string**

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigureAndDeployCitrixManagedCatalogApi

> string ConfigureAndDeployCitrixManagedCatalogApi(ctx, customerId, siteId).CitrixTransactionId(citrixTransactionId).CitrixManagedCatalogConfigDeployModel(citrixManagedCatalogConfigDeployModel).Execute()

Configure all the Citrix managed catalog deployment steps and initiate the catalog deployment

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
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    citrixManagedCatalogConfigDeployModel := *openapiclient.NewCitrixManagedCatalogConfigDeployModel() // CitrixManagedCatalogConfigDeployModel | The configuration of the catalog (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogCMD.ConfigureAndDeployCitrixManagedCatalogApi(context.Background(), customerId, siteId).CitrixTransactionId(citrixTransactionId).CitrixManagedCatalogConfigDeployModel(citrixManagedCatalogConfigDeployModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogCMD.ConfigureAndDeployCitrixManagedCatalogApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfigureAndDeployCitrixManagedCatalogApi`: string
    fmt.Fprintf(os.Stdout, "Response from `CatalogCMD.ConfigureAndDeployCitrixManagedCatalogApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**siteId** | **string** | The site ID of the customer | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigureAndDeployCitrixManagedCatalogApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixTransactionId** | **string** | The Transaction Id. | 
 **citrixManagedCatalogConfigDeployModel** | [**CitrixManagedCatalogConfigDeployModel**](CitrixManagedCatalogConfigDeployModel.md) | The configuration of the catalog | 

### Return type

**string**

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRemotePcCatalog

> string CreateRemotePcCatalog(ctx, customerId, siteId).CitrixTransactionId(citrixTransactionId).AddRemotePcCatalog(addRemotePcCatalog).Execute()

Create a Remote PC catalog.

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
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    addRemotePcCatalog := *openapiclient.NewAddRemotePcCatalog("Name_example", "ResourceLocationId_example") // AddRemotePcCatalog | The configuration of the catalog (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogCMD.CreateRemotePcCatalog(context.Background(), customerId, siteId).CitrixTransactionId(citrixTransactionId).AddRemotePcCatalog(addRemotePcCatalog).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogCMD.CreateRemotePcCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRemotePcCatalog`: string
    fmt.Fprintf(os.Stdout, "Response from `CatalogCMD.CreateRemotePcCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**siteId** | **string** | The site ID of the customer | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRemotePcCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixTransactionId** | **string** | The Transaction Id. | 
 **addRemotePcCatalog** | [**AddRemotePcCatalog**](AddRemotePcCatalog.md) | The configuration of the catalog | 

### Return type

**string**

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomerCatalog

> DeleteCustomerCatalog(ctx, customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).DeleteCatalogModel(deleteCatalogModel).Execute()

Delete a catalog along with all the pubhishd apps for the catalog

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
    catalogId := "catalogId_example" // string | ID of the catalog to be deleted
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    deleteCatalogModel := *openapiclient.NewDeleteCatalogModel() // DeleteCatalogModel | Configuration for deleting the catalog (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CatalogCMD.DeleteCustomerCatalog(context.Background(), customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).DeleteCatalogModel(deleteCatalogModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogCMD.DeleteCustomerCatalog``: %v\n", err)
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
**catalogId** | **string** | ID of the catalog to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomerCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 
 **deleteCatalogModel** | [**DeleteCatalogModel**](DeleteCatalogModel.md) | Configuration for deleting the catalog | 

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


## GetAllTroubleshootScripts

> TroubleshootScriptsOverviewsModel GetAllTroubleshootScripts(ctx, customerId, catalogId, machineName).CitrixTransactionId(citrixTransactionId).Execute()

Get the statuses of all the troubleshoot scripts assocaited with a vda

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
    catalogId := "catalogId_example" // string | ID of the catalog
    machineName := "machineName_example" // string | Name of the machine
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogCMD.GetAllTroubleshootScripts(context.Background(), customerId, catalogId, machineName).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogCMD.GetAllTroubleshootScripts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllTroubleshootScripts`: TroubleshootScriptsOverviewsModel
    fmt.Fprintf(os.Stdout, "Response from `CatalogCMD.GetAllTroubleshootScripts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**catalogId** | **string** | ID of the catalog | 
**machineName** | **string** | Name of the machine | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTroubleshootScriptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**TroubleshootScriptsOverviewsModel**](TroubleshootScriptsOverviewsModel.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllTroubleshootScripts_0

> TroubleshootScriptsOverviewsModel GetAllTroubleshootScripts_0(ctx, customerId, siteId, catalogId, machineName).CitrixTransactionId(citrixTransactionId).Execute()

Get the statuses of all the troubleshoot scripts associated with a vda

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
    siteId := "siteId_example" // string | The site Id pf the customer
    catalogId := "catalogId_example" // string | ID of the catalog
    machineName := "machineName_example" // string | Name of the machine
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogCMD.GetAllTroubleshootScripts_0(context.Background(), customerId, siteId, catalogId, machineName).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogCMD.GetAllTroubleshootScripts_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllTroubleshootScripts_0`: TroubleshootScriptsOverviewsModel
    fmt.Fprintf(os.Stdout, "Response from `CatalogCMD.GetAllTroubleshootScripts_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**siteId** | **string** | The site Id pf the customer | 
**catalogId** | **string** | ID of the catalog | 
**machineName** | **string** | Name of the machine | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTroubleshootScripts_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**TroubleshootScriptsOverviewsModel**](TroubleshootScriptsOverviewsModel.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogCapacityConfiguration

> CatalogCapacitySettingsModel GetCatalogCapacityConfiguration(ctx, customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).Execute()

Get the performance information configured for this catalog

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
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogCMD.GetCatalogCapacityConfiguration(context.Background(), customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogCMD.GetCatalogCapacityConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCatalogCapacityConfiguration`: CatalogCapacitySettingsModel
    fmt.Fprintf(os.Stdout, "Response from `CatalogCMD.GetCatalogCapacityConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**siteId** | **string** |  | 
**catalogId** | **string** | ID of the catalog | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogCapacityConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**CatalogCapacitySettingsModel**](CatalogCapacitySettingsModel.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerCatalog

> CatalogOverview GetCustomerCatalog(ctx, customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).Execute()

Returns a specific catalog for a specific customer

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
    customerId := "customerId_example" // string | Specific customerId to retrieve catalog for
    siteId := "siteId_example" // string | 
    catalogId := "catalogId_example" // string | ID of the catalog to retrieve
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogCMD.GetCustomerCatalog(context.Background(), customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogCMD.GetCustomerCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerCatalog`: CatalogOverview
    fmt.Fprintf(os.Stdout, "Response from `CatalogCMD.GetCustomerCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Specific customerId to retrieve catalog for | 
**siteId** | **string** |  | 
**catalogId** | **string** | ID of the catalog to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**CatalogOverview**](CatalogOverview.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerCatalogs

> CustomerCatalogOverviewsModel GetCustomerCatalogs(ctx, customerId, siteId).CatalogIdsOnly(catalogIdsOnly).CitrixTransactionId(citrixTransactionId).Execute()

Returns all the catalogs that the specified customer has created

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
    customerId := "customerId_example" // string | Specific customerId to retrieve catalogs for
    siteId := "siteId_example" // string | The site ID of the customer
    catalogIdsOnly := true // bool | Indicates if we should only return a list of the catalog ids (optional) (default to false)
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogCMD.GetCustomerCatalogs(context.Background(), customerId, siteId).CatalogIdsOnly(catalogIdsOnly).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogCMD.GetCustomerCatalogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerCatalogs`: CustomerCatalogOverviewsModel
    fmt.Fprintf(os.Stdout, "Response from `CatalogCMD.GetCustomerCatalogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Specific customerId to retrieve catalogs for | 
**siteId** | **string** | The site ID of the customer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerCatalogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **catalogIdsOnly** | **bool** | Indicates if we should only return a list of the catalog ids | [default to false]
 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**CustomerCatalogOverviewsModel**](CustomerCatalogOverviewsModel.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerManagedCatalogs

> CustomerManagedCatalogOverviewsModel GetCustomerManagedCatalogs(ctx, customerId, siteId).CspCustomerId(cspCustomerId).CspSiteId(cspSiteId).CitrixTransactionId(citrixTransactionId).Execute()

Returns all the catalogs that the specified Citrix managed customer has created

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
    customerId := "customerId_example" // string | Specific customerId to retrieve catalogs for
    siteId := "siteId_example" // string | The site ID of the customer
    cspCustomerId := "cspCustomerId_example" // string | Name of tenant customer ID if partner-tenant relationship exists otherwise null (optional)
    cspSiteId := "cspSiteId_example" // string | Name of tenant site ID if partner-tenant relationship exists otherwise null (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogCMD.GetCustomerManagedCatalogs(context.Background(), customerId, siteId).CspCustomerId(cspCustomerId).CspSiteId(cspSiteId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogCMD.GetCustomerManagedCatalogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerManagedCatalogs`: CustomerManagedCatalogOverviewsModel
    fmt.Fprintf(os.Stdout, "Response from `CatalogCMD.GetCustomerManagedCatalogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Specific customerId to retrieve catalogs for | 
**siteId** | **string** | The site ID of the customer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerManagedCatalogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cspCustomerId** | **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | 
 **cspSiteId** | **string** | Name of tenant site ID if partner-tenant relationship exists otherwise null | 
 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**CustomerManagedCatalogOverviewsModel**](CustomerManagedCatalogOverviewsModel.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerManagedCatalogsById

> CatalogOverview GetCustomerManagedCatalogsById(ctx, customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).Execute()

Returns all the catalogs that the specified Citrix managed customer has created

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
    customerId := "customerId_example" // string | Specific customerId to retrieve catalogs for
    siteId := "siteId_example" // string | The site ID of the customer
    catalogId := "catalogId_example" // string | The catalog ids for which to get the status
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogCMD.GetCustomerManagedCatalogsById(context.Background(), customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogCMD.GetCustomerManagedCatalogsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerManagedCatalogsById`: CatalogOverview
    fmt.Fprintf(os.Stdout, "Response from `CatalogCMD.GetCustomerManagedCatalogsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Specific customerId to retrieve catalogs for | 
**siteId** | **string** | The site ID of the customer | 
**catalogId** | **string** | The catalog ids for which to get the status | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerManagedCatalogsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**CatalogOverview**](CatalogOverview.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScriptsAndParameters

> GetScriptsAndParameters(ctx, customerId, catalogId).VmType(vmType).OperationType(operationType).CitrixTransactionId(citrixTransactionId).Execute()

Gets available scripts and corresponding default parameters for available operations

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
    customerId := "customerId_example" // string | 
    catalogId := "catalogId_example" // string | 
    vmType := openapiclient.CatalogResourceType("Unknown") // CatalogResourceType |  (optional)
    operationType := openapiclient.ScriptOperationType("Normal") // ScriptOperationType |  (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CatalogCMD.GetScriptsAndParameters(context.Background(), customerId, catalogId).VmType(vmType).OperationType(operationType).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogCMD.GetScriptsAndParameters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** |  | 
**catalogId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScriptsAndParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vmType** | [**CatalogResourceType**](CatalogResourceType.md) |  | 
 **operationType** | [**ScriptOperationType**](ScriptOperationType.md) |  | 
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


## GetScriptsAndParameters_0

> GetScriptsAndParameters_0(ctx, customerId, siteId, catalogId).VmType(vmType).OperationType(operationType).CitrixTransactionId(citrixTransactionId).Execute()

Gets available scripts and corresponding default parameters for available operations

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
    customerId := "customerId_example" // string | 
    siteId := "siteId_example" // string | 
    catalogId := "catalogId_example" // string | 
    vmType := openapiclient.CatalogResourceType("Unknown") // CatalogResourceType |  (optional)
    operationType := openapiclient.ScriptOperationType("Normal") // ScriptOperationType |  (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CatalogCMD.GetScriptsAndParameters_0(context.Background(), customerId, siteId, catalogId).VmType(vmType).OperationType(operationType).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogCMD.GetScriptsAndParameters_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** |  | 
**siteId** | **string** |  | 
**catalogId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScriptsAndParameters_2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **vmType** | [**CatalogResourceType**](CatalogResourceType.md) |  | 
 **operationType** | [**ScriptOperationType**](ScriptOperationType.md) |  | 
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


## GetTroubleshootScriptStatus

> TroubleshootScriptOverview GetTroubleshootScriptStatus(ctx, customerId, catalogId, machineName, troubleshootId).CitrixTransactionId(citrixTransactionId).Execute()

Get the status of the troubleshoot script

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
    catalogId := "catalogId_example" // string | ID of the catalog
    machineName := "machineName_example" // string | Name of the machine
    troubleshootId := "troubleshootId_example" // string | Id of the troubleshoot script
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogCMD.GetTroubleshootScriptStatus(context.Background(), customerId, catalogId, machineName, troubleshootId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogCMD.GetTroubleshootScriptStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTroubleshootScriptStatus`: TroubleshootScriptOverview
    fmt.Fprintf(os.Stdout, "Response from `CatalogCMD.GetTroubleshootScriptStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**catalogId** | **string** | ID of the catalog | 
**machineName** | **string** | Name of the machine | 
**troubleshootId** | **string** | Id of the troubleshoot script | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTroubleshootScriptStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**TroubleshootScriptOverview**](TroubleshootScriptOverview.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTroubleshootScriptStatus_0

> TroubleshootScriptOverview GetTroubleshootScriptStatus_0(ctx, customerId, siteId, catalogId, machineName, troubleshootId).CitrixTransactionId(citrixTransactionId).Execute()

Get the status of the troubleshoot script

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
    siteId := "siteId_example" // string | The site Id pf the customer
    catalogId := "catalogId_example" // string | ID of the catalog
    machineName := "machineName_example" // string | Name of the machine
    troubleshootId := "troubleshootId_example" // string | Id of the troubleshoot script
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogCMD.GetTroubleshootScriptStatus_0(context.Background(), customerId, siteId, catalogId, machineName, troubleshootId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogCMD.GetTroubleshootScriptStatus_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTroubleshootScriptStatus_0`: TroubleshootScriptOverview
    fmt.Fprintf(os.Stdout, "Response from `CatalogCMD.GetTroubleshootScriptStatus_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**siteId** | **string** | The site Id pf the customer | 
**catalogId** | **string** | ID of the catalog | 
**machineName** | **string** | Name of the machine | 
**troubleshootId** | **string** | Id of the troubleshoot script | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTroubleshootScriptStatus_3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**TroubleshootScriptOverview**](TroubleshootScriptOverview.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTroubleshootScriptsOnlyForFailedConnectorInstalls

> TroubleshootScriptsOverviewsModel GetTroubleshootScriptsOnlyForFailedConnectorInstalls(ctx, customerId, catalogId).CitrixTransactionId(citrixTransactionId).Execute()

Get the statuses of all the troubleshoot scripts assocaited with catalog connectors

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
    catalogId := "catalogId_example" // string | ID of the catalog
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogCMD.GetTroubleshootScriptsOnlyForFailedConnectorInstalls(context.Background(), customerId, catalogId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogCMD.GetTroubleshootScriptsOnlyForFailedConnectorInstalls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTroubleshootScriptsOnlyForFailedConnectorInstalls`: TroubleshootScriptsOverviewsModel
    fmt.Fprintf(os.Stdout, "Response from `CatalogCMD.GetTroubleshootScriptsOnlyForFailedConnectorInstalls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**catalogId** | **string** | ID of the catalog | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTroubleshootScriptsOnlyForFailedConnectorInstallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**TroubleshootScriptsOverviewsModel**](TroubleshootScriptsOverviewsModel.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTroubleshootScriptsOnlyForFailedConnectorInstalls_0

> TroubleshootScriptsOverviewsModel GetTroubleshootScriptsOnlyForFailedConnectorInstalls_0(ctx, customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).Execute()

Get the statuses of all the troubleshoot scripts associated with catalog connectors

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
    siteId := "siteId_example" // string | The site Id pf the customer
    catalogId := "catalogId_example" // string | ID of the catalog
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogCMD.GetTroubleshootScriptsOnlyForFailedConnectorInstalls_0(context.Background(), customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogCMD.GetTroubleshootScriptsOnlyForFailedConnectorInstalls_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTroubleshootScriptsOnlyForFailedConnectorInstalls_0`: TroubleshootScriptsOverviewsModel
    fmt.Fprintf(os.Stdout, "Response from `CatalogCMD.GetTroubleshootScriptsOnlyForFailedConnectorInstalls_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**siteId** | **string** | The site Id pf the customer | 
**catalogId** | **string** | ID of the catalog | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTroubleshootScriptsOnlyForFailedConnectorInstalls_4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**TroubleshootScriptsOverviewsModel**](TroubleshootScriptsOverviewsModel.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunScriptOnVm

> bool RunScriptOnVm(ctx, customerId, catalogId, machineName).CitrixTransactionId(citrixTransactionId).RunScriptOnVmModel(runScriptOnVmModel).Execute()

Run a troubleshoot script on a VDA

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
    catalogId := "catalogId_example" // string | ID of the catalog
    machineName := "machineName_example" // string | Name of the VDA
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    runScriptOnVmModel := *openapiclient.NewRunScriptOnVmModel(openapiclient.CatalogResourceType("Unknown"), "ScriptName_example") // RunScriptOnVmModel | The model containing the details required to run script on the specific VM resource (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogCMD.RunScriptOnVm(context.Background(), customerId, catalogId, machineName).CitrixTransactionId(citrixTransactionId).RunScriptOnVmModel(runScriptOnVmModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogCMD.RunScriptOnVm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunScriptOnVm`: bool
    fmt.Fprintf(os.Stdout, "Response from `CatalogCMD.RunScriptOnVm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**catalogId** | **string** | ID of the catalog | 
**machineName** | **string** | Name of the VDA | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunScriptOnVmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 
 **runScriptOnVmModel** | [**RunScriptOnVmModel**](RunScriptOnVmModel.md) | The model containing the details required to run script on the specific VM resource | 

### Return type

**bool**

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunScriptOnVm_0

> bool RunScriptOnVm_0(ctx, customerId, siteId, catalogId, machineName).CitrixTransactionId(citrixTransactionId).RunScriptOnVmModel(runScriptOnVmModel).Execute()

Run a troubleshoot script on a VDA

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
    machineName := "machineName_example" // string | Name of the VDA
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    runScriptOnVmModel := *openapiclient.NewRunScriptOnVmModel(openapiclient.CatalogResourceType("Unknown"), "ScriptName_example") // RunScriptOnVmModel | The model containing the details required to run script on the specific VM resource (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogCMD.RunScriptOnVm_0(context.Background(), customerId, siteId, catalogId, machineName).CitrixTransactionId(citrixTransactionId).RunScriptOnVmModel(runScriptOnVmModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogCMD.RunScriptOnVm_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunScriptOnVm_0`: bool
    fmt.Fprintf(os.Stdout, "Response from `CatalogCMD.RunScriptOnVm_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**siteId** | **string** | The site ID of the customer | 
**catalogId** | **string** | ID of the catalog | 
**machineName** | **string** | Name of the VDA | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunScriptOnVm_5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **citrixTransactionId** | **string** | The Transaction Id. | 
 **runScriptOnVmModel** | [**RunScriptOnVmModel**](RunScriptOnVmModel.md) | The model containing the details required to run script on the specific VM resource | 

### Return type

**bool**

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCatalogImage

> string UpdateCatalogImage(ctx, customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).UpdateCatalogTemplateImageModel(updateCatalogTemplateImageModel).Execute()

Update the catalog's master image.

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
    customerId := "customerId_example" // string | Specific customerId to update image for
    siteId := "siteId_example" // string | The site ID of the customer
    catalogId := "catalogId_example" // string | ID of the catalog to update
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    updateCatalogTemplateImageModel := *openapiclient.NewUpdateCatalogTemplateImageModel("TemplateId_example") // UpdateCatalogTemplateImageModel | Details of the image to update (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogCMD.UpdateCatalogImage(context.Background(), customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).UpdateCatalogTemplateImageModel(updateCatalogTemplateImageModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogCMD.UpdateCatalogImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCatalogImage`: string
    fmt.Fprintf(os.Stdout, "Response from `CatalogCMD.UpdateCatalogImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Specific customerId to update image for | 
**siteId** | **string** | The site ID of the customer | 
**catalogId** | **string** | ID of the catalog to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCatalogImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 
 **updateCatalogTemplateImageModel** | [**UpdateCatalogTemplateImageModel**](UpdateCatalogTemplateImageModel.md) | Details of the image to update | 

### Return type

**string**

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCatalogImageApi

> string UpdateCatalogImageApi(ctx, customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).UpdateCatalogTemplateImageModel(updateCatalogTemplateImageModel).Execute()

Update the catalog's master image.

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
    catalogId := "catalogId_example" // string | ID of the catalog to deploy
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    updateCatalogTemplateImageModel := *openapiclient.NewUpdateCatalogTemplateImageModel("TemplateId_example") // UpdateCatalogTemplateImageModel | Details of the image to update to (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogCMD.UpdateCatalogImageApi(context.Background(), customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).UpdateCatalogTemplateImageModel(updateCatalogTemplateImageModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogCMD.UpdateCatalogImageApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCatalogImageApi`: string
    fmt.Fprintf(os.Stdout, "Response from `CatalogCMD.UpdateCatalogImageApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**siteId** | **string** | The site ID of the customer | 
**catalogId** | **string** | ID of the catalog to deploy | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCatalogImageApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 
 **updateCatalogTemplateImageModel** | [**UpdateCatalogTemplateImageModel**](UpdateCatalogTemplateImageModel.md) | Details of the image to update to | 

### Return type

**string**

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCatalogScaleConfiguration

> UpdateCatalogScaleConfiguration(ctx, customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).CatalogCapacitySettingsModel(catalogCapacitySettingsModel).Execute()

Update the performance information configured for this catalog

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
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    catalogCapacitySettingsModel := *openapiclient.NewCatalogCapacitySettingsModel() // CatalogCapacitySettingsModel | Capacity informaton for this catalog (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CatalogCMD.UpdateCatalogScaleConfiguration(context.Background(), customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).CatalogCapacitySettingsModel(catalogCapacitySettingsModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogCMD.UpdateCatalogScaleConfiguration``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCatalogScaleConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 
 **catalogCapacitySettingsModel** | [**CatalogCapacitySettingsModel**](CatalogCapacitySettingsModel.md) | Capacity informaton for this catalog | 

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


## UpdateRemotePcCatalogScopes

> string UpdateRemotePcCatalogScopes(ctx, customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).UpdateRemotePcCatalogScopesModel(updateRemotePcCatalogScopesModel).Execute()

Configure all the Citrix managed catalog deployment steps and initiate the catalog deployment

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
    catalogId := "catalogId_example" // string | 
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    updateRemotePcCatalogScopesModel := *openapiclient.NewUpdateRemotePcCatalogScopesModel() // UpdateRemotePcCatalogScopesModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogCMD.UpdateRemotePcCatalogScopes(context.Background(), customerId, siteId, catalogId).CitrixTransactionId(citrixTransactionId).UpdateRemotePcCatalogScopesModel(updateRemotePcCatalogScopesModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogCMD.UpdateRemotePcCatalogScopes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRemotePcCatalogScopes`: string
    fmt.Fprintf(os.Stdout, "Response from `CatalogCMD.UpdateRemotePcCatalogScopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**siteId** | **string** | The site ID of the customer | 
**catalogId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRemotePcCatalogScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 
 **updateRemotePcCatalogScopesModel** | [**UpdateRemotePcCatalogScopesModel**](UpdateRemotePcCatalogScopesModel.md) |  | 

### Return type

**string**

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

