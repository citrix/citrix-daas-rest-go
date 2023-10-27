# \BatchAPIsDAAS

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchDoBatchRequest**](BatchAPIsDAAS.md#BatchDoBatchRequest) | **Post** /$batch | Perform a batch request.



## BatchDoBatchRequest

> BatchResponseModel BatchDoBatchRequest(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).BatchRequestModel(batchRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    batchRequestModel := *openapiclient.NewBatchRequestModel([]openapiclient.BatchRequestItemModel{*openapiclient.NewBatchRequestItemModel("For example, 0", "For example, /techpreview/me")}) // BatchRequestModel | List of requests that the server should perform in parallel.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the batch request will be processed as a background task. This is recommended, since a batch request that takes more than 90 seconds to execute may time out, and results may be lost. The task will have JobType Batch. The SubJobs will correspond to the items in the batch request. When the task is complete it will redirect to GetJobResults. The job's Parameters will contain no properties.              (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchAPIsDAAS.BatchDoBatchRequest(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).BatchRequestModel(batchRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
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
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **batchRequestModel** | [**BatchRequestModel**](BatchRequestModel.md) | List of requests that the server should perform in parallel. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the batch request will be processed as a background task. This is recommended, since a batch request that takes more than 90 seconds to execute may time out, and results may be lost. The task will have JobType Batch. The SubJobs will correspond to the items in the batch request. When the task is complete it will redirect to GetJobResults. The job&#39;s Parameters will contain no properties.              | [default to false]

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

