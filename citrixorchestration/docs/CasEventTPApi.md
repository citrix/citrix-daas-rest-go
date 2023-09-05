# \CasEventTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CasEventTPSendCasEvent**](CasEventTPApi.md#CasEventTPSendCasEvent) | **Post** /techpreview/{customerid}/{siteid}/CasEvents | Send Event to Cas.



## CasEventTPSendCasEvent

> CasEventTPSendCasEvent(ctx, customerid, siteid).CasEventRequest(casEventRequest).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Send Event to Cas.



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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    casEventRequest := *openapiclient.NewCasEventRequestModel("EventType_example") // CasEventRequestModel | Details of the application folder to create.
    async := true // bool | If `true`, the even will be sent to cas as a background task. The task will have JobType \"JobType.SendCasEvent\". (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CasEventTPApi.CasEventTPSendCasEvent(context.Background(), customerid, siteid).CasEventRequest(casEventRequest).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CasEventTPApi.CasEventTPSendCasEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCasEventTPSendCasEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **casEventRequest** | [**CasEventRequestModel**](CasEventRequestModel.md) | Details of the application folder to create. | 
 **async** | **bool** | If &#x60;true&#x60;, the even will be sent to cas as a background task. The task will have JobType \&quot;JobType.SendCasEvent\&quot;. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

