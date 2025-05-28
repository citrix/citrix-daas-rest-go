# \AzureSubscriptionsCMD

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSubscriptions**](AzureSubscriptionsCMD.md#GetSubscriptions) | **Get** /{customerId}/{siteId}/subscriptions | Returns the subscriptions that we have a known association with



## GetSubscriptions

> AzureSubscriptionsModel GetSubscriptions(ctx, customerId, siteId).CitrixTransactionId(citrixTransactionId).XAZUREACCESSTOKEN(xAZUREACCESSTOKEN).XAZUREGRAPHACCESSTOKEN(xAZUREGRAPHACCESSTOKEN).XAZURETENANTID(xAZURETENANTID).XAZUREAPPCLIENTID(xAZUREAPPCLIENTID).XAZUREAPPCLIENTSECRET(xAZUREAPPCLIENTSECRET).Execute()

Returns the subscriptions that we have a known association with

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
    siteId := "siteId_example" // string | 
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    xAZUREACCESSTOKEN := "xAZUREACCESSTOKEN_example" // string | Azure Access Token. (optional)
    xAZUREGRAPHACCESSTOKEN := "xAZUREGRAPHACCESSTOKEN_example" // string | Azure Graph Access Token. (optional)
    xAZURETENANTID := "xAZURETENANTID_example" // string | Azure Tenant Id. (optional)
    xAZUREAPPCLIENTID := "xAZUREAPPCLIENTID_example" // string | Azure Application Key. (optional)
    xAZUREAPPCLIENTSECRET := "xAZUREAPPCLIENTSECRET_example" // string | Azure Application Secret. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureSubscriptionsCMD.GetSubscriptions(context.Background(), customerId, siteId).CitrixTransactionId(citrixTransactionId).XAZUREACCESSTOKEN(xAZUREACCESSTOKEN).XAZUREGRAPHACCESSTOKEN(xAZUREGRAPHACCESSTOKEN).XAZURETENANTID(xAZURETENANTID).XAZUREAPPCLIENTID(xAZUREAPPCLIENTID).XAZUREAPPCLIENTSECRET(xAZUREAPPCLIENTSECRET).Execute()
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
**customerId** | **string** | Specific customer id | 
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

