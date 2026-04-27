# \AzureSubscriptionsCMD

All URIs are relative to *https://catalogs.apps.cloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSubscriptions**](AzureSubscriptionsCMD.md#GetSubscriptions) | **Get** /{customerId}/{siteId}/subscriptions | 
[**GetVirtualHubVnetConnections**](AzureSubscriptionsCMD.md#GetVirtualHubVnetConnections) | **Get** /{customerId}/{siteId}/subscriptions/{subscriptionId}/virtualhubs/{resourceGroup}/{virtualHubName}/vnets | Returns the VNet connections attached to the specified Virtual Hub.
[**GetVirtualHubs**](AzureSubscriptionsCMD.md#GetVirtualHubs) | **Get** /{customerId}/{siteId}/subscriptions/{subscriptionId}/virtualhubs | Returns the Virtual Hubs available in the specified Azure subscription, optionally filtered by resource group.
[**GetVirtualWans**](AzureSubscriptionsCMD.md#GetVirtualWans) | **Get** /{customerId}/{siteId}/subscriptions/{subscriptionId}/virtualwans | Returns the Virtual WANs available in the specified Azure subscription.



## GetSubscriptions

> AzureSubscriptionsModel GetSubscriptions(ctx, customerId, siteId).SkipCache(skipCache).CitrixTransactionId(citrixTransactionId).XAZUREACCESSTOKEN(xAZUREACCESSTOKEN).XAZUREGRAPHACCESSTOKEN(xAZUREGRAPHACCESSTOKEN).XAZURETENANTID(xAZURETENANTID).XAZUREAPPCLIENTID(xAZUREAPPCLIENTID).XAZUREAPPCLIENTSECRET(xAZUREAPPCLIENTSECRET).Execute()



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
    skipCache := true // bool |  (optional) (default to false)
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    xAZUREACCESSTOKEN := "xAZUREACCESSTOKEN_example" // string | Azure Access Token. (optional)
    xAZUREGRAPHACCESSTOKEN := "xAZUREGRAPHACCESSTOKEN_example" // string | Azure Graph Access Token. (optional)
    xAZURETENANTID := "xAZURETENANTID_example" // string | Azure Tenant Id. (optional)
    xAZUREAPPCLIENTID := "xAZUREAPPCLIENTID_example" // string | Azure Application Key. (optional)
    xAZUREAPPCLIENTSECRET := "xAZUREAPPCLIENTSECRET_example" // string | Azure Application Secret. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureSubscriptionsCMD.GetSubscriptions(context.Background(), customerId, siteId).SkipCache(skipCache).CitrixTransactionId(citrixTransactionId).XAZUREACCESSTOKEN(xAZUREACCESSTOKEN).XAZUREGRAPHACCESSTOKEN(xAZUREGRAPHACCESSTOKEN).XAZURETENANTID(xAZURETENANTID).XAZUREAPPCLIENTID(xAZUREAPPCLIENTID).XAZUREAPPCLIENTSECRET(xAZUREAPPCLIENTSECRET).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureSubscriptionsCMD.GetSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscriptions`: AzureSubscriptionsModel
    fmt.Fprintf(os.Stdout, "Response from `AzureSubscriptionsCMD.GetSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** |  | 
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipCache** | **bool** |  | [default to false]
 **citrixTransactionId** | **string** | The Transaction Id. | 
 **xAZUREACCESSTOKEN** | **string** | Azure Access Token. | 
 **xAZUREGRAPHACCESSTOKEN** | **string** | Azure Graph Access Token. | 
 **xAZURETENANTID** | **string** | Azure Tenant Id. | 
 **xAZUREAPPCLIENTID** | **string** | Azure Application Key. | 
 **xAZUREAPPCLIENTSECRET** | **string** | Azure Application Secret. | 

### Return type

[**AzureSubscriptionsModel**](AzureSubscriptionsModel.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualHubVnetConnections

> []AzureVirtualHubVnetConnection GetVirtualHubVnetConnections(ctx, customerId, siteId, subscriptionId, resourceGroup, virtualHubName).CitrixTransactionId(citrixTransactionId).XAZUREACCESSTOKEN(xAZUREACCESSTOKEN).XAZUREGRAPHACCESSTOKEN(xAZUREGRAPHACCESSTOKEN).XAZURETENANTID(xAZURETENANTID).XAZUREAPPCLIENTID(xAZUREAPPCLIENTID).XAZUREAPPCLIENTSECRET(xAZUREAPPCLIENTSECRET).Execute()

Returns the VNet connections attached to the specified Virtual Hub.

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
    customerId := "customerId_example" // string | Specific customer id
    siteId := "siteId_example" // string | ID of the customer's site
    subscriptionId := "subscriptionId_example" // string | ID of the Azure Subscription
    resourceGroup := "resourceGroup_example" // string | Name of the Resource Group where the Virtual Hub is located
    virtualHubName := "virtualHubName_example" // string | Name of the Virtual Hub
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    xAZUREACCESSTOKEN := "xAZUREACCESSTOKEN_example" // string | Azure Access Token. (optional)
    xAZUREGRAPHACCESSTOKEN := "xAZUREGRAPHACCESSTOKEN_example" // string | Azure Graph Access Token. (optional)
    xAZURETENANTID := "xAZURETENANTID_example" // string | Azure Tenant Id. (optional)
    xAZUREAPPCLIENTID := "xAZUREAPPCLIENTID_example" // string | Azure Application Key. (optional)
    xAZUREAPPCLIENTSECRET := "xAZUREAPPCLIENTSECRET_example" // string | Azure Application Secret. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureSubscriptionsCMD.GetVirtualHubVnetConnections(context.Background(), customerId, siteId, subscriptionId, resourceGroup, virtualHubName).CitrixTransactionId(citrixTransactionId).XAZUREACCESSTOKEN(xAZUREACCESSTOKEN).XAZUREGRAPHACCESSTOKEN(xAZUREGRAPHACCESSTOKEN).XAZURETENANTID(xAZURETENANTID).XAZUREAPPCLIENTID(xAZUREAPPCLIENTID).XAZUREAPPCLIENTSECRET(xAZUREAPPCLIENTSECRET).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureSubscriptionsCMD.GetVirtualHubVnetConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualHubVnetConnections`: []AzureVirtualHubVnetConnection
    fmt.Fprintf(os.Stdout, "Response from `AzureSubscriptionsCMD.GetVirtualHubVnetConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Specific customer id | 
**siteId** | **string** | ID of the customer&#39;s site | 
**subscriptionId** | **string** | ID of the Azure Subscription | 
**resourceGroup** | **string** | Name of the Resource Group where the Virtual Hub is located | 
**virtualHubName** | **string** | Name of the Virtual Hub | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualHubVnetConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **citrixTransactionId** | **string** | The Transaction Id. | 
 **xAZUREACCESSTOKEN** | **string** | Azure Access Token. | 
 **xAZUREGRAPHACCESSTOKEN** | **string** | Azure Graph Access Token. | 
 **xAZURETENANTID** | **string** | Azure Tenant Id. | 
 **xAZUREAPPCLIENTID** | **string** | Azure Application Key. | 
 **xAZUREAPPCLIENTSECRET** | **string** | Azure Application Secret. | 

### Return type

[**[]AzureVirtualHubVnetConnection**](AzureVirtualHubVnetConnection.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualHubs

> []AzureVirtualHub GetVirtualHubs(ctx, customerId, siteId, subscriptionId).ResourceGroup(resourceGroup).CitrixTransactionId(citrixTransactionId).XAZUREACCESSTOKEN(xAZUREACCESSTOKEN).XAZUREGRAPHACCESSTOKEN(xAZUREGRAPHACCESSTOKEN).XAZURETENANTID(xAZURETENANTID).XAZUREAPPCLIENTID(xAZUREAPPCLIENTID).XAZUREAPPCLIENTSECRET(xAZUREAPPCLIENTSECRET).Execute()

Returns the Virtual Hubs available in the specified Azure subscription, optionally filtered by resource group.

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
    customerId := "customerId_example" // string | Specific customer id
    siteId := "siteId_example" // string | ID of the customer's site
    subscriptionId := "subscriptionId_example" // string | ID of the Azure Subscription
    resourceGroup := "resourceGroup_example" // string | Optional resource group name to filter results (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    xAZUREACCESSTOKEN := "xAZUREACCESSTOKEN_example" // string | Azure Access Token. (optional)
    xAZUREGRAPHACCESSTOKEN := "xAZUREGRAPHACCESSTOKEN_example" // string | Azure Graph Access Token. (optional)
    xAZURETENANTID := "xAZURETENANTID_example" // string | Azure Tenant Id. (optional)
    xAZUREAPPCLIENTID := "xAZUREAPPCLIENTID_example" // string | Azure Application Key. (optional)
    xAZUREAPPCLIENTSECRET := "xAZUREAPPCLIENTSECRET_example" // string | Azure Application Secret. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureSubscriptionsCMD.GetVirtualHubs(context.Background(), customerId, siteId, subscriptionId).ResourceGroup(resourceGroup).CitrixTransactionId(citrixTransactionId).XAZUREACCESSTOKEN(xAZUREACCESSTOKEN).XAZUREGRAPHACCESSTOKEN(xAZUREGRAPHACCESSTOKEN).XAZURETENANTID(xAZURETENANTID).XAZUREAPPCLIENTID(xAZUREAPPCLIENTID).XAZUREAPPCLIENTSECRET(xAZUREAPPCLIENTSECRET).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureSubscriptionsCMD.GetVirtualHubs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualHubs`: []AzureVirtualHub
    fmt.Fprintf(os.Stdout, "Response from `AzureSubscriptionsCMD.GetVirtualHubs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Specific customer id | 
**siteId** | **string** | ID of the customer&#39;s site | 
**subscriptionId** | **string** | ID of the Azure Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualHubsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **resourceGroup** | **string** | Optional resource group name to filter results | 
 **citrixTransactionId** | **string** | The Transaction Id. | 
 **xAZUREACCESSTOKEN** | **string** | Azure Access Token. | 
 **xAZUREGRAPHACCESSTOKEN** | **string** | Azure Graph Access Token. | 
 **xAZURETENANTID** | **string** | Azure Tenant Id. | 
 **xAZUREAPPCLIENTID** | **string** | Azure Application Key. | 
 **xAZUREAPPCLIENTSECRET** | **string** | Azure Application Secret. | 

### Return type

[**[]AzureVirtualHub**](AzureVirtualHub.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualWans

> []AzureVirtualWanModel GetVirtualWans(ctx, customerId, siteId, subscriptionId).CitrixTransactionId(citrixTransactionId).XAZUREACCESSTOKEN(xAZUREACCESSTOKEN).XAZUREGRAPHACCESSTOKEN(xAZUREGRAPHACCESSTOKEN).XAZURETENANTID(xAZURETENANTID).XAZUREAPPCLIENTID(xAZUREAPPCLIENTID).XAZUREAPPCLIENTSECRET(xAZUREAPPCLIENTSECRET).Execute()

Returns the Virtual WANs available in the specified Azure subscription.

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
    customerId := "customerId_example" // string | Specific customer id
    siteId := "siteId_example" // string | ID of the customer's site
    subscriptionId := "subscriptionId_example" // string | ID of the Azure Subscription
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    xAZUREACCESSTOKEN := "xAZUREACCESSTOKEN_example" // string | Azure Access Token. (optional)
    xAZUREGRAPHACCESSTOKEN := "xAZUREGRAPHACCESSTOKEN_example" // string | Azure Graph Access Token. (optional)
    xAZURETENANTID := "xAZURETENANTID_example" // string | Azure Tenant Id. (optional)
    xAZUREAPPCLIENTID := "xAZUREAPPCLIENTID_example" // string | Azure Application Key. (optional)
    xAZUREAPPCLIENTSECRET := "xAZUREAPPCLIENTSECRET_example" // string | Azure Application Secret. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureSubscriptionsCMD.GetVirtualWans(context.Background(), customerId, siteId, subscriptionId).CitrixTransactionId(citrixTransactionId).XAZUREACCESSTOKEN(xAZUREACCESSTOKEN).XAZUREGRAPHACCESSTOKEN(xAZUREGRAPHACCESSTOKEN).XAZURETENANTID(xAZURETENANTID).XAZUREAPPCLIENTID(xAZUREAPPCLIENTID).XAZUREAPPCLIENTSECRET(xAZUREAPPCLIENTSECRET).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureSubscriptionsCMD.GetVirtualWans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualWans`: []AzureVirtualWanModel
    fmt.Fprintf(os.Stdout, "Response from `AzureSubscriptionsCMD.GetVirtualWans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Specific customer id | 
**siteId** | **string** | ID of the customer&#39;s site | 
**subscriptionId** | **string** | ID of the Azure Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualWansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 
 **xAZUREACCESSTOKEN** | **string** | Azure Access Token. | 
 **xAZUREGRAPHACCESSTOKEN** | **string** | Azure Graph Access Token. | 
 **xAZURETENANTID** | **string** | Azure Tenant Id. | 
 **xAZUREAPPCLIENTID** | **string** | Azure Application Key. | 
 **xAZUREAPPCLIENTSECRET** | **string** | Azure Application Secret. | 

### Return type

[**[]AzureVirtualWanModel**](AzureVirtualWanModel.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

