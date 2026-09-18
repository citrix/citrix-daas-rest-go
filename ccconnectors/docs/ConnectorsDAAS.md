# \ConnectorsDAAS

All URIs are relative to *https://api.cloud.com/connectors*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectorsDelete**](ConnectorsDAAS.md#ConnectorsDelete) | **Delete** /{Id} | Delete a connector.
[**ConnectorsGetAll**](ConnectorsDAAS.md#ConnectorsGetAll) | **Get** / | Get all Edge Servers (Connectors) for a customer.



## ConnectorsDelete

> bool ConnectorsDelete(ctx, id).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Execute()

Delete a connector.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/ccconnectors"
)

func main() {
	accept := "accept_example" // string | Only supports application/json
	authorization := "authorization_example" // string | The access token.
	citrixCustomerId := "citrixCustomerId_example" // string | ID of the customer.
	id := "id_example" // string | The connector id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsDAAS.ConnectorsDelete(context.Background(), id).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsDAAS.ConnectorsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorsDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsDAAS.ConnectorsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The connector id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectorsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only supports application/json | 
 **authorization** | **string** | The access token. | 
 **citrixCustomerId** | **string** | ID of the customer. | 


### Return type

**bool**

### Authorization

[Basic](../README.md#Basic), [CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectorsGetAll

> []CitrixCloudServicesAgentHubApiEdgeServersGetResultModel ConnectorsGetAll(ctx).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Location(location).ExtendedData(extendedData).ConnectorType(connectorType).RetrieveExpectedVersion(retrieveExpectedVersion).Execute()

Get all Edge Servers (Connectors) for a customer.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/ccconnectors"
)

func main() {
	accept := "accept_example" // string | Only supports application/json
	authorization := "authorization_example" // string | The access token.
	citrixCustomerId := "citrixCustomerId_example" // string | ID of the customer.
	location := "location_example" // string |  (optional)
	extendedData := true // bool |  (optional)
	connectorType := "connectorType_example" // string |  (optional)
	retrieveExpectedVersion := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsDAAS.ConnectorsGetAll(context.Background()).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Location(location).ExtendedData(extendedData).ConnectorType(connectorType).RetrieveExpectedVersion(retrieveExpectedVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsDAAS.ConnectorsGetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectorsGetAll`: []CitrixCloudServicesAgentHubApiEdgeServersGetResultModel
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsDAAS.ConnectorsGetAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectorsGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only supports application/json | 
 **authorization** | **string** | The access token. | 
 **citrixCustomerId** | **string** | ID of the customer. | 
 **location** | **string** |  | 
 **extendedData** | **bool** |  | 
 **connectorType** | **string** |  | 
 **retrieveExpectedVersion** | **bool** |  | 

### Return type

[**[]CitrixCloudServicesAgentHubApiEdgeServersGetResultModel**](CitrixCloudServicesAgentHubApiEdgeServersGetResultModel.md)

### Authorization

[Basic](../README.md#Basic), [CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

