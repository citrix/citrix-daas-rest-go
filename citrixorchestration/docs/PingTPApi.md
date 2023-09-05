# \PingTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PingTPGetServiceStatus**](PingTPApi.md#PingTPGetServiceStatus) | **Get** /techpreview/ping/status | Get Orchestration service status.
[**PingTPPingTP**](PingTPApi.md#PingTPPingTP) | **Get** /techpreview/ping | Test if orchestration web api is ready.



## PingTPGetServiceStatus

> StatusModel PingTPGetServiceStatus(ctx).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get Orchestration service status.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingTPApi.PingTPGetServiceStatus(context.Background()).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingTPApi.PingTPGetServiceStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PingTPGetServiceStatus`: StatusModel
    fmt.Fprintf(os.Stdout, "Response from `PingTPApi.PingTPGetServiceStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPingTPGetServiceStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**StatusModel**](StatusModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PingTPPingTP

> bool PingTPPingTP(ctx).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Test if orchestration web api is ready.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func main() {
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingTPApi.PingTPPingTP(context.Background()).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingTPApi.PingTPPingTP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PingTPPingTP`: bool
    fmt.Fprintf(os.Stdout, "Response from `PingTPApi.PingTPPingTP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPingTPPingTPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

**bool**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

