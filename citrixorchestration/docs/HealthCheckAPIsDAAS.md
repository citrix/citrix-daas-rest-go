# \HealthCheckAPIsDAAS

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HealthCheckHealthCheck**](HealthCheckAPIsDAAS.md#HealthCheckHealthCheck) | **Get** /HealthCheck | The health check endpoint.



## HealthCheckHealthCheck

> bool HealthCheckHealthCheck(ctx).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

The health check endpoint.

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
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HealthCheckAPIsDAAS.HealthCheckHealthCheck(context.Background()).CitrixCustomerId(citrixCustomerId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthCheckAPIsDAAS.HealthCheckHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HealthCheckHealthCheck`: bool
	fmt.Fprintf(os.Stdout, "Response from `HealthCheckAPIsDAAS.HealthCheckHealthCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHealthCheckHealthCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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

