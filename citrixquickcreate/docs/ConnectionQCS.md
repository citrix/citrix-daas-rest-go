# \ConnectionQCS

All URIs are relative to *https://api.cloud.com/quickcreateservice*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddResourceConnectionAsync**](ConnectionQCS.md#AddResourceConnectionAsync) | **Post** /{customerId}/accounts/{accountId}/connections/$initiate | Adds resource connection asyncronously
[**GetAllResourceConnectionsAsync**](ConnectionQCS.md#GetAllResourceConnectionsAsync) | **Get** /{customerId}/connections | Gets all resource connections without specifying account
[**GetResourceConnectionAsync**](ConnectionQCS.md#GetResourceConnectionAsync) | **Get** /{customerId}/accounts/{accountId}/connections/{connectionId} | Gets resource connection
[**GetResourceConnectionsAsync**](ConnectionQCS.md#GetResourceConnectionsAsync) | **Get** /{customerId}/accounts/{accountId}/connections | Gets resource connections
[**ModifyResourceConnectionAsync**](ConnectionQCS.md#ModifyResourceConnectionAsync) | **Patch** /{customerId}/accounts/{accountId}/connections/{connectionId} | Modifies connection
[**RemoveResourceConnectionAsync**](ConnectionQCS.md#RemoveResourceConnectionAsync) | **Delete** /{customerId}/accounts/{accountId}/connections/{connectionId} | Removes connection



## AddResourceConnectionAsync

> ResourceConnectionTask AddResourceConnectionAsync(ctx, customerId, accountId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Adds resource connection asyncronously

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
	accountId := "accountId_example" // string | ID of account
	citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
	body := AddAwsEdcDirectoryConnection(987) // AddAwsEdcDirectoryConnection | Connection configuration (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionQCS.AddResourceConnectionAsync(context.Background(), customerId, accountId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionQCS.AddResourceConnectionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddResourceConnectionAsync`: ResourceConnectionTask
	fmt.Fprintf(os.Stdout, "Response from `ConnectionQCS.AddResourceConnectionAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**accountId** | **string** | ID of account | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddResourceConnectionAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixTransactionId** | **string** | The Transaction Id. | 
 **body** | **AddAwsEdcDirectoryConnection** | Connection configuration | 

### Return type

[**ResourceConnectionTask**](ResourceConnectionTask.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllResourceConnectionsAsync

> ResourceConnections GetAllResourceConnectionsAsync(ctx, customerId).AccountType(accountType).CitrixTransactionId(citrixTransactionId).Execute()

Gets all resource connections without specifying account

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
	accountType := openapiclient.AccountType("AWSEDC") // AccountType | Filtering parameter for account type (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionQCS.GetAllResourceConnectionsAsync(context.Background(), customerId).AccountType(accountType).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionQCS.GetAllResourceConnectionsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllResourceConnectionsAsync`: ResourceConnections
	fmt.Fprintf(os.Stdout, "Response from `ConnectionQCS.GetAllResourceConnectionsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllResourceConnectionsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountType** | [**AccountType**](AccountType.md) | Filtering parameter for account type | 
 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**ResourceConnections**](ResourceConnections.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceConnectionAsync

> AwsEdcDirectoryConnection GetResourceConnectionAsync(ctx, customerId, accountId, connectionId).CitrixTransactionId(citrixTransactionId).Execute()

Gets resource connection

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
	accountId := "accountId_example" // string | ID of account
	connectionId := "connectionId_example" // string | ID of connection
	citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionQCS.GetResourceConnectionAsync(context.Background(), customerId, accountId, connectionId).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionQCS.GetResourceConnectionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResourceConnectionAsync`: AwsEdcDirectoryConnection
	fmt.Fprintf(os.Stdout, "Response from `ConnectionQCS.GetResourceConnectionAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**accountId** | **string** | ID of account | 
**connectionId** | **string** | ID of connection | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceConnectionAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**AwsEdcDirectoryConnection**](AwsEdcDirectoryConnection.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceConnectionsAsync

> ResourceConnections GetResourceConnectionsAsync(ctx, customerId, accountId).CitrixManaged(citrixManaged).CitrixTransactionId(citrixTransactionId).Execute()

Gets resource connections

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
	accountId := "accountId_example" // string | ID of account
	citrixManaged := true // bool | citrix managed connections (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionQCS.GetResourceConnectionsAsync(context.Background(), customerId, accountId).CitrixManaged(citrixManaged).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionQCS.GetResourceConnectionsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResourceConnectionsAsync`: ResourceConnections
	fmt.Fprintf(os.Stdout, "Response from `ConnectionQCS.GetResourceConnectionsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**accountId** | **string** | ID of account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceConnectionsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixManaged** | **bool** | citrix managed connections | 
 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**ResourceConnections**](ResourceConnections.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyResourceConnectionAsync

> AwsEdcDirectoryConnection ModifyResourceConnectionAsync(ctx, customerId, accountId, connectionId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Modifies connection

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
	accountId := "accountId_example" // string | ID of account
	connectionId := "connectionId_example" // string | ID of connection
	citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
	body := UpdateAwsEdcDirectoryConnection(987) // UpdateAwsEdcDirectoryConnection | Update configuration (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionQCS.ModifyResourceConnectionAsync(context.Background(), customerId, accountId, connectionId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionQCS.ModifyResourceConnectionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyResourceConnectionAsync`: AwsEdcDirectoryConnection
	fmt.Fprintf(os.Stdout, "Response from `ConnectionQCS.ModifyResourceConnectionAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**accountId** | **string** | ID of account | 
**connectionId** | **string** | ID of connection | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyResourceConnectionAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 
 **body** | **UpdateAwsEdcDirectoryConnection** | Update configuration | 

### Return type

[**AwsEdcDirectoryConnection**](AwsEdcDirectoryConnection.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveResourceConnectionAsync

> ResourceConnectionTask RemoveResourceConnectionAsync(ctx, customerId, accountId, connectionId).ForceDelete(forceDelete).CitrixTransactionId(citrixTransactionId).Execute()

Removes connection

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
	accountId := "accountId_example" // string | ID of account
	connectionId := "connectionId_example" // string | ID of connection
	forceDelete := true // bool | Force delete a connection (optional) (default to false)
	citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionQCS.RemoveResourceConnectionAsync(context.Background(), customerId, accountId, connectionId).ForceDelete(forceDelete).CitrixTransactionId(citrixTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionQCS.RemoveResourceConnectionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveResourceConnectionAsync`: ResourceConnectionTask
	fmt.Fprintf(os.Stdout, "Response from `ConnectionQCS.RemoveResourceConnectionAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**accountId** | **string** | ID of account | 
**connectionId** | **string** | ID of connection | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveResourceConnectionAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **forceDelete** | **bool** | Force delete a connection | [default to false]
 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**ResourceConnectionTask**](ResourceConnectionTask.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

