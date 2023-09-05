# \RebootSchedulesTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RebootSchedulesTPGetRebootSchedules**](RebootSchedulesTPApi.md#RebootSchedulesTPGetRebootSchedules) | **Get** /techpreview/{customerid}/{siteid}/RebootSchedules | Get all reboot schedules in the site.



## RebootSchedulesTPGetRebootSchedules

> RebootScheduleResponseModelCollection RebootSchedulesTPGetRebootSchedules(ctx, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get all reboot schedules in the site.



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
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RebootSchedulesTPApi.RebootSchedulesTPGetRebootSchedules(context.Background(), customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RebootSchedulesTPApi.RebootSchedulesTPGetRebootSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RebootSchedulesTPGetRebootSchedules`: RebootScheduleResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `RebootSchedulesTPApi.RebootSchedulesTPGetRebootSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootSchedulesTPGetRebootSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**RebootScheduleResponseModelCollection**](RebootScheduleResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

