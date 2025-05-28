# \ManagedCapacityCMD

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePortForNetworkSecurityGroup**](ManagedCapacityCMD.md#DeletePortForNetworkSecurityGroup) | **Delete** /{customerId}/{siteId}/managedCapacity/resourcelocations/{resourceLocationId}/port/{port} | Perform operation to Delete Port For CMA Vnet
[**GetAllPortsForNetworkSecurityGroup**](ManagedCapacityCMD.md#GetAllPortsForNetworkSecurityGroup) | **Get** /{customerId}/{siteId}/managedCapacity/ports | Get details of all ports in CMA VNet
[**GetDeploymentRegions**](ManagedCapacityCMD.md#GetDeploymentRegions) | **Get** /{customerId}/{siteId}/managedCapacity/regions | Get the regions that are available for deployment by the customer
[**GetOnPremConnections**](ManagedCapacityCMD.md#GetOnPremConnections) | **Get** /{customerId}/{siteId}/managedCapacity/onPremConnections | Get the On-Prem connections configured for the customer
[**GetOpenPortsForNetworkSecurityGroup**](ManagedCapacityCMD.md#GetOpenPortsForNetworkSecurityGroup) | **Get** /{customerId}/{siteId}/managedCapacity/resourcelocations/{resourceLocationId}/openPorts | Get list of open ports in CMA Vnet for provided resource location
[**OpenPortForNetworkSecurityGroup**](ManagedCapacityCMD.md#OpenPortForNetworkSecurityGroup) | **Post** /{customerId}/{siteId}/managedCapacity/resourcelocations/{resourceLocationId}/port/{port}/openPort | Perform operation to Open Port For CMA Vnet



## DeletePortForNetworkSecurityGroup

> DeletePortForNetworkSecurityGroup(ctx, customerId, siteId, resourceLocationId, port).CitrixTransactionId(citrixTransactionId).Execute()

Perform operation to Delete Port For CMA Vnet

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
    resourceLocationId := "resourceLocationId_example" // string | 
    port := "port_example" // string | 
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCapacityCMD.DeletePortForNetworkSecurityGroup(context.Background(), customerId, siteId, resourceLocationId, port).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCapacityCMD.DeletePortForNetworkSecurityGroup``: %v\n", err)
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
**resourceLocationId** | **string** |  | 
**port** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePortForNetworkSecurityGroupRequest struct via the builder pattern


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


## GetAllPortsForNetworkSecurityGroup

> VNetResourceDetails GetAllPortsForNetworkSecurityGroup(ctx, customerId, siteId).CitrixTransactionId(citrixTransactionId).Execute()

Get details of all ports in CMA VNet



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
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedCapacityCMD.GetAllPortsForNetworkSecurityGroup(context.Background(), customerId, siteId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCapacityCMD.GetAllPortsForNetworkSecurityGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllPortsForNetworkSecurityGroup`: VNetResourceDetails
    fmt.Fprintf(os.Stdout, "Response from `ManagedCapacityCMD.GetAllPortsForNetworkSecurityGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** |  | 
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPortsForNetworkSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**VNetResourceDetails**](VNetResourceDetails.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentRegions

> DeploymentRegionsModel GetDeploymentRegions(ctx, customerId, siteId).IncludeHidden(includeHidden).CitrixTransactionId(citrixTransactionId).Execute()

Get the regions that are available for deployment by the customer

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
    customerId := "customerId_example" // string | ID of the user
    siteId := "siteId_example" // string | ID of the customer's site
    includeHidden := true // bool | Should hidden regions be returned? (optional) (default to false)
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedCapacityCMD.GetDeploymentRegions(context.Background(), customerId, siteId).IncludeHidden(includeHidden).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCapacityCMD.GetDeploymentRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentRegions`: DeploymentRegionsModel
    fmt.Fprintf(os.Stdout, "Response from `ManagedCapacityCMD.GetDeploymentRegions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the user | 
**siteId** | **string** | ID of the customer&#39;s site | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeHidden** | **bool** | Should hidden regions be returned? | [default to false]
 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**DeploymentRegionsModel**](DeploymentRegionsModel.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnPremConnections

> OnPremConnectionsModel GetOnPremConnections(ctx, customerId, siteId).CspCustomerId(cspCustomerId).CspSiteId(cspSiteId).RegionId(regionId).CitrixTransactionId(citrixTransactionId).Execute()

Get the On-Prem connections configured for the customer

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
    customerId := "customerId_example" // string | ID of the user
    siteId := "siteId_example" // string | ID of the customer's site
    cspCustomerId := "cspCustomerId_example" // string | Name of tenant customer ID if partner-tenant relationship exists otherwise null (optional)
    cspSiteId := "cspSiteId_example" // string | Name of tenant site ID if partner-tenant relationship exists otherwise null (optional)
    regionId := "regionId_example" // string | Return only connections for specified region. Returns all if null. (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedCapacityCMD.GetOnPremConnections(context.Background(), customerId, siteId).CspCustomerId(cspCustomerId).CspSiteId(cspSiteId).RegionId(regionId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCapacityCMD.GetOnPremConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOnPremConnections`: OnPremConnectionsModel
    fmt.Fprintf(os.Stdout, "Response from `ManagedCapacityCMD.GetOnPremConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the user | 
**siteId** | **string** | ID of the customer&#39;s site | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnPremConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cspCustomerId** | **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | 
 **cspSiteId** | **string** | Name of tenant site ID if partner-tenant relationship exists otherwise null | 
 **regionId** | **string** | Return only connections for specified region. Returns all if null. | 
 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**OnPremConnectionsModel**](OnPremConnectionsModel.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpenPortsForNetworkSecurityGroup

> VNetResourceLocation GetOpenPortsForNetworkSecurityGroup(ctx, customerId, siteId, resourceLocationId).CitrixTransactionId(citrixTransactionId).Execute()

Get list of open ports in CMA Vnet for provided resource location

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
    resourceLocationId := "resourceLocationId_example" // string | 
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedCapacityCMD.GetOpenPortsForNetworkSecurityGroup(context.Background(), customerId, siteId, resourceLocationId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCapacityCMD.GetOpenPortsForNetworkSecurityGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOpenPortsForNetworkSecurityGroup`: VNetResourceLocation
    fmt.Fprintf(os.Stdout, "Response from `ManagedCapacityCMD.GetOpenPortsForNetworkSecurityGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** |  | 
**siteId** | **string** |  | 
**resourceLocationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenPortsForNetworkSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**VNetResourceLocation**](VNetResourceLocation.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenPortForNetworkSecurityGroup

> string OpenPortForNetworkSecurityGroup(ctx, customerId, siteId, resourceLocationId, port).CitrixTransactionId(citrixTransactionId).Execute()

Perform operation to Open Port For CMA Vnet

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
    resourceLocationId := "resourceLocationId_example" // string | 
    port := "port_example" // string | 
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedCapacityCMD.OpenPortForNetworkSecurityGroup(context.Background(), customerId, siteId, resourceLocationId, port).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCapacityCMD.OpenPortForNetworkSecurityGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpenPortForNetworkSecurityGroup`: string
    fmt.Fprintf(os.Stdout, "Response from `ManagedCapacityCMD.OpenPortForNetworkSecurityGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** |  | 
**siteId** | **string** |  | 
**resourceLocationId** | **string** |  | 
**port** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpenPortForNetworkSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

**string**

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

