# \BatchAPIsDAAS

All URIs are relative to *https://api-us.cloud.com/cvad/manage*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchDoBatchRequest**](BatchAPIsDAAS.md#BatchDoBatchRequest) | **Post** /$batch | Perform a batch request.



## BatchDoBatchRequest

> BatchResponseModel BatchDoBatchRequest(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).BatchRequestModel(batchRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud Customer ID.
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Cloud Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request.
    batchRequestModel := *openapiclient.NewBatchRequestModel([]openapiclient.BatchRequestItemModel{*openapiclient.NewBatchRequestItemModel("For example, 0", "For example, /techpreview/me")}) // BatchRequestModel | List of requests that the server should perform in parallel.
    async := true // bool | If `true`, the batch request will be processed as a background task. This is recommended, since a batch request that takes more than 90 seconds to execute may time out, and results may be lost. The task will have JobType Batch. The SubJobs will correspond to the items in the batch request. When the task is complete it will redirect to GetJobResults. The job's Parameters will contain no properties.              (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchAPIsDAAS.BatchDoBatchRequest(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).BatchRequestModel(batchRequestModel).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchAPIsDAAS.BatchDoBatchRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchDoBatchRequest`: BatchResponseModel
    fmt.Fprintf(os.Stdout, "Response from `BatchAPIsDAAS.BatchDoBatchRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchDoBatchRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Cloud Customer ID. | 
 **citrixInstanceId** | **string** | Citrix Cloud Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **batchRequestModel** | [**BatchRequestModel**](BatchRequestModel.md) | List of requests that the server should perform in parallel. | 
 **async** | **bool** | If &#x60;true&#x60;, the batch request will be processed as a background task. This is recommended, since a batch request that takes more than 90 seconds to execute may time out, and results may be lost. The task will have JobType Batch. The SubJobs will correspond to the items in the batch request. When the task is complete it will redirect to GetJobResults. The job&#39;s Parameters will contain no properties.              | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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

