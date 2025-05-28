# \IconCMD

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExtractIcon**](IconCMD.md#ExtractIcon) | **Post** /{customerId}/{siteId}/icons | Extract an icon from the specified file data



## ExtractIcon

> string ExtractIcon(ctx, customerId, siteId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Extract an icon from the specified file data

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
    body := *openapiclient.NewExtractIconModel() // ExtractIconModel | Name and bytes of the file to extract icon from (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IconCMD.ExtractIcon(context.Background(), customerId, siteId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IconCMD.ExtractIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtractIcon`: string
    fmt.Fprintf(os.Stdout, "Response from `IconCMD.ExtractIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**siteId** | **string** | The site ID of the customer | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtractIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixTransactionId** | **string** | The Transaction Id. | 
 **body** | [**ExtractIconModel**](ExtractIconModel.md) | Name and bytes of the file to extract icon from | 

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

