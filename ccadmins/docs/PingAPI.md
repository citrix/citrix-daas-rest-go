# \PingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerPingGet**](PingAPI.md#CustomerPingGet) | **Get** /{customer}/ping | Gets a health check. [Public]



## CustomerPingGet

> bool CustomerPingGet(ctx, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).Execute()

Gets a health check. [Public]

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/ccadmins"
)

func main() {
	customer := "capamerica" // string | 
	citrixConsistencyToken := "citrixConsistencyToken_example" // string | CosmosDB consistency token. (optional)
	xCwsTransactionId := "xCwsTransactionId_example" // string | Used for Citrix Cloud telemetry correlation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PingAPI.CustomerPingGet(context.Background(), customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PingAPI.CustomerPingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerPingGet`: bool
	fmt.Fprintf(os.Stdout, "Response from `PingAPI.CustomerPingGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerPingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

