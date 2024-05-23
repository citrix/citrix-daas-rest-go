# \LocationsDAAS

All URIs are relative to *https://api.cloud.com/resourcelocations*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocationsCreate**](LocationsDAAS.md#LocationsCreate) | **Post** / | Create a resource location for a customer.
[**LocationsDelete**](LocationsDAAS.md#LocationsDelete) | **Delete** /{Id} | Delete a resource location.
[**LocationsGet**](LocationsDAAS.md#LocationsGet) | **Get** /{Id} | Get a resource location from id.
[**LocationsGetAll**](LocationsDAAS.md#LocationsGetAll) | **Get** / | Get all resource locations for a customer.
[**LocationsUpdate**](LocationsDAAS.md#LocationsUpdate) | **Put** /{Id} | Update the customer resource location information.



## LocationsCreate

> CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel LocationsCreate(ctx).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Accept(accept).Model(model).Execute()

Create a resource location for a customer.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/ccresourcelocations"
)

func main() {
	authorization := "authorization_example" // string | The access token.
	citrixCustomerId := "citrixCustomerId_example" // string | ID of the customer.
	accept := "accept_example" // string | Only supports application/json
	model := *openapiclient.NewCitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel("Name_example") // CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel | The resource location model.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationsDAAS.LocationsCreate(context.Background()).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Accept(accept).Model(model).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationsDAAS.LocationsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LocationsCreate`: CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel
	fmt.Fprintf(os.Stdout, "Response from `LocationsDAAS.LocationsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLocationsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The access token. | 
 **citrixCustomerId** | **string** | ID of the customer. | 
 **accept** | **string** | Only supports application/json | 
 **model** | [**CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel**](CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel.md) | The resource location model. | 

### Return type

[**CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel**](CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel.md)

### Authorization

[Basic](../README.md#Basic), [CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocationsDelete

> LocationsDelete(ctx, id).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Accept(accept).Execute()

Delete a resource location.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/ccresourcelocations"
)

func main() {
	authorization := "authorization_example" // string | The access token.
	id := "id_example" // string | The resource location id.
	citrixCustomerId := "citrixCustomerId_example" // string | ID of the customer.
	accept := "accept_example" // string | Only supports application/json

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LocationsDAAS.LocationsDelete(context.Background(), id).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Accept(accept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationsDAAS.LocationsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The resource location id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLocationsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The access token. | 

 **citrixCustomerId** | **string** | ID of the customer. | 
 **accept** | **string** | Only supports application/json | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocationsGet

> CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel LocationsGet(ctx, id).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Accept(accept).Execute()

Get a resource location from id.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/ccresourcelocations"
)

func main() {
	authorization := "authorization_example" // string | The access token.
	id := "id_example" // string | The resource location id
	citrixCustomerId := "citrixCustomerId_example" // string | ID of the customer.
	accept := "accept_example" // string | Only supports application/json

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationsDAAS.LocationsGet(context.Background(), id).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Accept(accept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationsDAAS.LocationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LocationsGet`: CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel
	fmt.Fprintf(os.Stdout, "Response from `LocationsDAAS.LocationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The resource location id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLocationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The access token. | 

 **citrixCustomerId** | **string** | ID of the customer. | 
 **accept** | **string** | Only supports application/json | 

### Return type

[**CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel**](CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel.md)

### Authorization

[Basic](../README.md#Basic), [CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocationsGetAll

> CitrixCloudServicesRegistryApiModelsLocationsResourceLocationsResultsModel LocationsGetAll(ctx).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Accept(accept).Execute()

Get all resource locations for a customer.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/ccresourcelocations"
)

func main() {
	authorization := "authorization_example" // string | The access token.
	citrixCustomerId := "citrixCustomerId_example" // string | ID of the customer.
	accept := "accept_example" // string | Only supports application/json

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationsDAAS.LocationsGetAll(context.Background()).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Accept(accept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationsDAAS.LocationsGetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LocationsGetAll`: CitrixCloudServicesRegistryApiModelsLocationsResourceLocationsResultsModel
	fmt.Fprintf(os.Stdout, "Response from `LocationsDAAS.LocationsGetAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLocationsGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The access token. | 
 **citrixCustomerId** | **string** | ID of the customer. | 
 **accept** | **string** | Only supports application/json | 

### Return type

[**CitrixCloudServicesRegistryApiModelsLocationsResourceLocationsResultsModel**](CitrixCloudServicesRegistryApiModelsLocationsResourceLocationsResultsModel.md)

### Authorization

[Basic](../README.md#Basic), [CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocationsUpdate

> CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel LocationsUpdate(ctx, id).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Accept(accept).Model(model).Execute()

Update the customer resource location information.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/citrix/citrix-daas-rest-go/ccresourcelocations"
)

func main() {
	authorization := "authorization_example" // string | The access token.
	citrixCustomerId := "citrixCustomerId_example" // string | ID of the customer.
	accept := "accept_example" // string | Only supports application/json
	id := "id_example" // string | The resource location id.
	model := *openapiclient.NewCitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel() // CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel | The resource location model.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationsDAAS.LocationsUpdate(context.Background(), id).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Accept(accept).Model(model).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationsDAAS.LocationsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LocationsUpdate`: CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel
	fmt.Fprintf(os.Stdout, "Response from `LocationsDAAS.LocationsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The resource location id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLocationsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The access token. | 
 **citrixCustomerId** | **string** | ID of the customer. | 
 **accept** | **string** | Only supports application/json | 

 **model** | [**CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel**](CitrixCloudServicesRegistryApiModelsLocationsResourceLocationUpdateModel.md) | The resource location model. | 

### Return type

[**CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel**](CitrixCloudServicesRegistryApiModelsLocationsResourceLocationModel.md)

### Authorization

[Basic](../README.md#Basic), [CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

