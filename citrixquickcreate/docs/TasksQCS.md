# \TasksQCS

All URIs are relative to *https://api.cloud.com/quickcreateservice*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTaskAsync**](TasksQCS.md#GetTaskAsync) | **Get** /{customerId}/tasks/{taskId} | Get the status of a task



## GetTaskAsync

> GetTaskAsync200Response GetTaskAsync(ctx, customerId, taskId).CitrixTransactionId(citrixTransactionId).Execute()

Get the status of a task

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixquickcreate"
)

func main() {
	customerId := "customerId_example" // string | ID of the customer
	taskId := "taskId_example" // string | ID of the task
	citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksQCS.GetTaskAsync(context.Background(), customerId, taskId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksQCS.GetTaskAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskAsync`: GetTaskAsync200Response
	fmt.Fprintf(os.Stdout, "Response from `TasksQCS.GetTaskAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**taskId** | **string** | ID of the task | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**GetTaskAsync200Response**](GetTaskAsync200Response.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

