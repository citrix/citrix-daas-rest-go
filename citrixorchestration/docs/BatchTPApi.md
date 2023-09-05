# \BatchTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchTPDoBatchRequest**](BatchTPApi.md#BatchTPDoBatchRequest) | **Post** /techpreview/{customerid}/{siteid}/$batch | Perform a batch request.



## BatchTPDoBatchRequest

> BatchResponseModel BatchTPDoBatchRequest(ctx, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Perform a batch request.



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
    customerid := "customerid_example" // string | The customer Id.
    siteid := "siteid_example" // string | The site Id.
    request := *openapiclient.NewBatchRequestModel([]openapiclient.BatchRequestItemModel{*openapiclient.NewBatchRequestItemModel("Reference_example", "RelativeUrl_example")}) // BatchRequestModel | List of requests that the server should perform in parallel.
    async := true // bool | If `true`, the batch request will be processed as a background task. This is recommended, since a batch request that takes more than 90 seconds to execute may time out, and results may be lost. The task will have JobType Batch. The SubJobs will correspond to the items in the batch request. When the task is complete it will redirect to GetJobResults. The job's Parameters will contain no properties.              (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchTPApi.BatchTPDoBatchRequest(context.Background(), customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchTPApi.BatchTPDoBatchRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchTPDoBatchRequest`: BatchResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BatchTPApi.BatchTPDoBatchRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** | The customer Id. | 
**siteid** | **string** | The site Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchTPDoBatchRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**BatchRequestModel**](BatchRequestModel.md) | List of requests that the server should perform in parallel. | 
 **async** | **bool** | If &#x60;true&#x60;, the batch request will be processed as a background task. This is recommended, since a batch request that takes more than 90 seconds to execute may time out, and results may be lost. The task will have JobType Batch. The SubJobs will correspond to the items in the batch request. When the task is complete it will redirect to GetJobResults. The job&#39;s Parameters will contain no properties.              | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**BatchResponseModel**](BatchResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

