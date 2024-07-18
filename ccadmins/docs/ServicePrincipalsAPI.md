# \ServicePrincipalsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerServicePrincipalsDelete**](ServicePrincipalsAPI.md#CustomerServicePrincipalsDelete) | **Delete** /{customer}/ServicePrincipals | Delete all Service Principals in a customer [BearerToken]
[**CustomerServicePrincipalsGet**](ServicePrincipalsAPI.md#CustomerServicePrincipalsGet) | **Get** /{customer}/ServicePrincipals | Fetches all Service Principals [BearerToken]
[**CustomerServicePrincipalsIdAccessPut**](ServicePrincipalsAPI.md#CustomerServicePrincipalsIdAccessPut) | **Put** /{customer}/ServicePrincipals/{id}/access | Updates service principal access and policies. [BearerToken]
[**CustomerServicePrincipalsIdDelete**](ServicePrincipalsAPI.md#CustomerServicePrincipalsIdDelete) | **Delete** /{customer}/ServicePrincipals/{id} | Delete a Service Principal [BearerToken]
[**CustomerServicePrincipalsIdGet**](ServicePrincipalsAPI.md#CustomerServicePrincipalsIdGet) | **Get** /{customer}/ServicePrincipals/{id} | Fetches a Service Principal [BearerToken]
[**CustomerServicePrincipalsIdRotatePost**](ServicePrincipalsAPI.md#CustomerServicePrincipalsIdRotatePost) | **Post** /{customer}/ServicePrincipals/{id}/$rotate | Rotates primary secret of a service principal. [BearerToken]
[**CustomerServicePrincipalsIdSecondarySecretDelete**](ServicePrincipalsAPI.md#CustomerServicePrincipalsIdSecondarySecretDelete) | **Delete** /{customer}/ServicePrincipals/{id}/secondarySecret | Delete a Service Principal Secondary Secret [BearerToken]
[**CustomerServicePrincipalsIdUpdateLastAccessedDatePost**](ServicePrincipalsAPI.md#CustomerServicePrincipalsIdUpdateLastAccessedDatePost) | **Post** /{customer}/ServicePrincipals/{id}/$updateLastAccessedDate | Update the LastAccessedDate property for a ServicePrincipal [BearerToken]
[**CustomerServicePrincipalsMigratePost**](ServicePrincipalsAPI.md#CustomerServicePrincipalsMigratePost) | **Post** /{customer}/ServicePrincipals/$migrate | Update the LastAccessedDate property for a ServicePrincipal [BearerToken]
[**CustomerServicePrincipalsPost**](ServicePrincipalsAPI.md#CustomerServicePrincipalsPost) | **Post** /{customer}/ServicePrincipals | Create a Service Principal [BearerToken]



## CustomerServicePrincipalsDelete

> bool CustomerServicePrincipalsDelete(ctx, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).Execute()

Delete all Service Principals in a customer [BearerToken]

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
	citrixTimeoutMs := TODO // float64 | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicePrincipalsAPI.CustomerServicePrincipalsDelete(context.Background(), customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsAPI.CustomerServicePrincipalsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerServicePrincipalsDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsAPI.CustomerServicePrincipalsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerServicePrincipalsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 
 **citrixTimeoutMs** | [**float64**](float64.md) | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds | 

### Return type

**bool**

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerServicePrincipalsGet

> CitrixCloudServicesApiModelsSegmentResult CustomerServicePrincipalsGet(ctx, customer).Take(take).Continuation(continuation).Name(name).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).Execute()

Fetches all Service Principals [BearerToken]

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
	take := int32(56) // int32 |  (optional)
	continuation := "continuation_example" // string |  (optional)
	name := "name_example" // string |  (optional)
	citrixConsistencyToken := "citrixConsistencyToken_example" // string | CosmosDB consistency token. (optional)
	xCwsTransactionId := "xCwsTransactionId_example" // string | Used for Citrix Cloud telemetry correlation. (optional)
	citrixTimeoutMs := TODO // float64 | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicePrincipalsAPI.CustomerServicePrincipalsGet(context.Background(), customer).Take(take).Continuation(continuation).Name(name).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsAPI.CustomerServicePrincipalsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerServicePrincipalsGet`: CitrixCloudServicesApiModelsSegmentResult
	fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsAPI.CustomerServicePrincipalsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerServicePrincipalsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **take** | **int32** |  | 
 **continuation** | **string** |  | 
 **name** | **string** |  | 
 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 
 **citrixTimeoutMs** | [**float64**](float64.md) | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds | 

### Return type

[**CitrixCloudServicesApiModelsSegmentResult**](CitrixCloudServicesApiModelsSegmentResult.md)

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerServicePrincipalsIdAccessPut

> CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel CustomerServicePrincipalsIdAccessPut(ctx, id, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel(citrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel).Execute()

Updates service principal access and policies. [BearerToken]

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
	id := "id_example" // string | 
	customer := "capamerica" // string | 
	citrixConsistencyToken := "citrixConsistencyToken_example" // string | CosmosDB consistency token. (optional)
	xCwsTransactionId := "xCwsTransactionId_example" // string | Used for Citrix Cloud telemetry correlation. (optional)
	citrixTimeoutMs := TODO // float64 | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds (optional)
	citrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel := *openapiclient.NewCitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel(openapiclient.Citrix.CloudServices.ServiceRuntime.PolicySets.Models.AdministratorAccessType("Full")) // CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicePrincipalsAPI.CustomerServicePrincipalsIdAccessPut(context.Background(), id, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel(citrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsAPI.CustomerServicePrincipalsIdAccessPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerServicePrincipalsIdAccessPut`: CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel
	fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsAPI.CustomerServicePrincipalsIdAccessPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerServicePrincipalsIdAccessPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 
 **citrixTimeoutMs** | [**float64**](float64.md) | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds | 
 **citrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel** | [**CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel**](CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel.md) |  | 

### Return type

[**CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel**](CitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel.md)

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerServicePrincipalsIdDelete

> bool CustomerServicePrincipalsIdDelete(ctx, id, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).Execute()

Delete a Service Principal [BearerToken]

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
	id := "id_example" // string | 
	customer := "capamerica" // string | 
	citrixConsistencyToken := "citrixConsistencyToken_example" // string | CosmosDB consistency token. (optional)
	xCwsTransactionId := "xCwsTransactionId_example" // string | Used for Citrix Cloud telemetry correlation. (optional)
	citrixTimeoutMs := TODO // float64 | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicePrincipalsAPI.CustomerServicePrincipalsIdDelete(context.Background(), id, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsAPI.CustomerServicePrincipalsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerServicePrincipalsIdDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsAPI.CustomerServicePrincipalsIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerServicePrincipalsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 
 **citrixTimeoutMs** | [**float64**](float64.md) | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds | 

### Return type

**bool**

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerServicePrincipalsIdGet

> CitrixCloudServicesServicePrincipalsModelsServicePrincipal CustomerServicePrincipalsIdGet(ctx, id, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).Execute()

Fetches a Service Principal [BearerToken]

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
	id := "id_example" // string | 
	customer := "capamerica" // string | 
	citrixConsistencyToken := "citrixConsistencyToken_example" // string | CosmosDB consistency token. (optional)
	xCwsTransactionId := "xCwsTransactionId_example" // string | Used for Citrix Cloud telemetry correlation. (optional)
	citrixTimeoutMs := TODO // float64 | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicePrincipalsAPI.CustomerServicePrincipalsIdGet(context.Background(), id, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsAPI.CustomerServicePrincipalsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerServicePrincipalsIdGet`: CitrixCloudServicesServicePrincipalsModelsServicePrincipal
	fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsAPI.CustomerServicePrincipalsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerServicePrincipalsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 
 **citrixTimeoutMs** | [**float64**](float64.md) | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds | 

### Return type

[**CitrixCloudServicesServicePrincipalsModelsServicePrincipal**](CitrixCloudServicesServicePrincipalsModelsServicePrincipal.md)

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerServicePrincipalsIdRotatePost

> CitrixCloudServicesServicePrincipalsModelsSecretRotationResponse CustomerServicePrincipalsIdRotatePost(ctx, customer, id).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).CitrixCloudServicesServicePrincipalsModelsSecretRotationRequest(citrixCloudServicesServicePrincipalsModelsSecretRotationRequest).Execute()

Rotates primary secret of a service principal. [BearerToken]

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/citrix/citrix-daas-rest-go/ccadmins"
)

func main() {
	customer := "capamerica" // string | 
	id := "id_example" // string | 
	citrixConsistencyToken := "citrixConsistencyToken_example" // string | CosmosDB consistency token. (optional)
	xCwsTransactionId := "xCwsTransactionId_example" // string | Used for Citrix Cloud telemetry correlation. (optional)
	citrixTimeoutMs := TODO // float64 | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds (optional)
	citrixCloudServicesServicePrincipalsModelsSecretRotationRequest := *openapiclient.NewCitrixCloudServicesServicePrincipalsModelsSecretRotationRequest(time.Now()) // CitrixCloudServicesServicePrincipalsModelsSecretRotationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicePrincipalsAPI.CustomerServicePrincipalsIdRotatePost(context.Background(), customer, id).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).CitrixCloudServicesServicePrincipalsModelsSecretRotationRequest(citrixCloudServicesServicePrincipalsModelsSecretRotationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsAPI.CustomerServicePrincipalsIdRotatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerServicePrincipalsIdRotatePost`: CitrixCloudServicesServicePrincipalsModelsSecretRotationResponse
	fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsAPI.CustomerServicePrincipalsIdRotatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerServicePrincipalsIdRotatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 
 **citrixTimeoutMs** | [**float64**](float64.md) | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds | 
 **citrixCloudServicesServicePrincipalsModelsSecretRotationRequest** | [**CitrixCloudServicesServicePrincipalsModelsSecretRotationRequest**](CitrixCloudServicesServicePrincipalsModelsSecretRotationRequest.md) |  | 

### Return type

[**CitrixCloudServicesServicePrincipalsModelsSecretRotationResponse**](CitrixCloudServicesServicePrincipalsModelsSecretRotationResponse.md)

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerServicePrincipalsIdSecondarySecretDelete

> bool CustomerServicePrincipalsIdSecondarySecretDelete(ctx, id, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).Execute()

Delete a Service Principal Secondary Secret [BearerToken]

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
	id := "id_example" // string | 
	customer := "capamerica" // string | 
	citrixConsistencyToken := "citrixConsistencyToken_example" // string | CosmosDB consistency token. (optional)
	xCwsTransactionId := "xCwsTransactionId_example" // string | Used for Citrix Cloud telemetry correlation. (optional)
	citrixTimeoutMs := TODO // float64 | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicePrincipalsAPI.CustomerServicePrincipalsIdSecondarySecretDelete(context.Background(), id, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsAPI.CustomerServicePrincipalsIdSecondarySecretDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerServicePrincipalsIdSecondarySecretDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsAPI.CustomerServicePrincipalsIdSecondarySecretDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerServicePrincipalsIdSecondarySecretDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 
 **citrixTimeoutMs** | [**float64**](float64.md) | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds | 

### Return type

**bool**

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerServicePrincipalsIdUpdateLastAccessedDatePost

> CustomerServicePrincipalsIdUpdateLastAccessedDatePost(ctx, customer, id).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).Execute()

Update the LastAccessedDate property for a ServicePrincipal [BearerToken]

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
	id := "id_example" // string | 
	citrixConsistencyToken := "citrixConsistencyToken_example" // string | CosmosDB consistency token. (optional)
	xCwsTransactionId := "xCwsTransactionId_example" // string | Used for Citrix Cloud telemetry correlation. (optional)
	citrixTimeoutMs := TODO // float64 | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServicePrincipalsAPI.CustomerServicePrincipalsIdUpdateLastAccessedDatePost(context.Background(), customer, id).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsAPI.CustomerServicePrincipalsIdUpdateLastAccessedDatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerServicePrincipalsIdUpdateLastAccessedDatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 
 **citrixTimeoutMs** | [**float64**](float64.md) | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds | 

### Return type

 (empty response body)

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerServicePrincipalsMigratePost

> CustomerServicePrincipalsMigratePost(ctx, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).CitrixCloudServicesServicePrincipalsModelsMigrateRequest(citrixCloudServicesServicePrincipalsModelsMigrateRequest).Execute()

Update the LastAccessedDate property for a ServicePrincipal [BearerToken]

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/citrix/citrix-daas-rest-go/ccadmins"
)

func main() {
	customer := "capamerica" // string | 
	citrixConsistencyToken := "citrixConsistencyToken_example" // string | CosmosDB consistency token. (optional)
	xCwsTransactionId := "xCwsTransactionId_example" // string | Used for Citrix Cloud telemetry correlation. (optional)
	citrixTimeoutMs := TODO // float64 | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds (optional)
	citrixCloudServicesServicePrincipalsModelsMigrateRequest := *openapiclient.NewCitrixCloudServicesServicePrincipalsModelsMigrateRequest([]string{"SecureClients_example"}, time.Now()) // CitrixCloudServicesServicePrincipalsModelsMigrateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServicePrincipalsAPI.CustomerServicePrincipalsMigratePost(context.Background(), customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).CitrixCloudServicesServicePrincipalsModelsMigrateRequest(citrixCloudServicesServicePrincipalsModelsMigrateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsAPI.CustomerServicePrincipalsMigratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerServicePrincipalsMigratePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 
 **citrixTimeoutMs** | [**float64**](float64.md) | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds | 
 **citrixCloudServicesServicePrincipalsModelsMigrateRequest** | [**CitrixCloudServicesServicePrincipalsModelsMigrateRequest**](CitrixCloudServicesServicePrincipalsModelsMigrateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerServicePrincipalsPost

> CitrixCloudServicesServicePrincipalsModelsCreationResponse CustomerServicePrincipalsPost(ctx, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).CitrixCloudServicesServicePrincipalsModelsCreationRequest(citrixCloudServicesServicePrincipalsModelsCreationRequest).Execute()

Create a Service Principal [BearerToken]

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/citrix/citrix-daas-rest-go/ccadmins"
)

func main() {
	customer := "capamerica" // string | 
	citrixConsistencyToken := "citrixConsistencyToken_example" // string | CosmosDB consistency token. (optional)
	xCwsTransactionId := "xCwsTransactionId_example" // string | Used for Citrix Cloud telemetry correlation. (optional)
	citrixTimeoutMs := TODO // float64 | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds (optional)
	citrixCloudServicesServicePrincipalsModelsCreationRequest := *openapiclient.NewCitrixCloudServicesServicePrincipalsModelsCreationRequest("Name_example", *openapiclient.NewCitrixCloudServicesServiceRuntimePolicySetsModelsAdministratorAccessModel(openapiclient.Citrix.CloudServices.ServiceRuntime.PolicySets.Models.AdministratorAccessType("Full")), time.Now()) // CitrixCloudServicesServicePrincipalsModelsCreationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicePrincipalsAPI.CustomerServicePrincipalsPost(context.Background(), customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).CitrixCloudServicesServicePrincipalsModelsCreationRequest(citrixCloudServicesServicePrincipalsModelsCreationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsAPI.CustomerServicePrincipalsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerServicePrincipalsPost`: CitrixCloudServicesServicePrincipalsModelsCreationResponse
	fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsAPI.CustomerServicePrincipalsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerServicePrincipalsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 
 **citrixTimeoutMs** | [**float64**](float64.md) | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds | 
 **citrixCloudServicesServicePrincipalsModelsCreationRequest** | [**CitrixCloudServicesServicePrincipalsModelsCreationRequest**](CitrixCloudServicesServicePrincipalsModelsCreationRequest.md) |  | 

### Return type

[**CitrixCloudServicesServicePrincipalsModelsCreationResponse**](CitrixCloudServicesServicePrincipalsModelsCreationResponse.md)

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

