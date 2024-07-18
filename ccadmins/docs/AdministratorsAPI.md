# \AdministratorsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerAdministratorGet**](AdministratorsAPI.md#CustomerAdministratorGet) | **Get** /{customer}/administrator | Gets an administrator with ucOid passed as a query parameter. [ServiceKey][BearerToken]
[**CustomerAdministratorsAccessPut**](AdministratorsAPI.md#CustomerAdministratorsAccessPut) | **Put** /{customer}/Administrators/Access | Updates the access for an administrator. [ServiceKey][BearerToken]
[**CustomerAdministratorsAccessUcOidGet**](AdministratorsAPI.md#CustomerAdministratorsAccessUcOidGet) | **Get** /{customer}/administrators/access/{ucOid} | Get the access of an administrator [ServiceKey][BearerToken]
[**CustomerAdministratorsActiveAdministratorsGet**](AdministratorsAPI.md#CustomerAdministratorsActiveAdministratorsGet) | **Get** /{customer}/administrators/activeAdministrators | Gets all active administrators (users, secureClients, and groups) across all customers. [ServiceKey][RootOnly]
[**CustomerAdministratorsAdminIdDelete**](AdministratorsAPI.md#CustomerAdministratorsAdminIdDelete) | **Delete** /{customer}/administrators/{adminId} | Deletes an administrator by ucOid. [ServiceKey][BearerToken]
[**CustomerAdministratorsAdminIdPut**](AdministratorsAPI.md#CustomerAdministratorsAdminIdPut) | **Put** /{customer}/administrators/{adminId} | Updates an administrator by either ucOid or user id. [BearerToken]
[**CustomerAdministratorsAuthDomainPut**](AdministratorsAPI.md#CustomerAdministratorsAuthDomainPut) | **Put** /{customer}/administrators/authDomain | Updates auth domain for relevant administrators of given customer context.  Forwards to DAv1 if not found locally or is pending. [ServiceKey]
[**CustomerAdministratorsClaimsDecompressPost**](AdministratorsAPI.md#CustomerAdministratorsClaimsDecompressPost) | **Post** /{customer}/administrators/claims/$decompress | Decompresses the group claims. [ServiceKey][BearerToken]
[**CustomerAdministratorsCreatePost**](AdministratorsAPI.md#CustomerAdministratorsCreatePost) | **Post** /{customer}/administrators/$create | Adds an administrator. This API is a higher-level version of  &#x60;POST /{customer}/administrators&#x60; that works with any type of  administrator with stricter input validation. [BearerToken]
[**CustomerAdministratorsDeleteMultiplePost**](AdministratorsAPI.md#CustomerAdministratorsDeleteMultiplePost) | **Post** /{customer}/administrators/DeleteMultiple | Deletes an administrator. [ServiceKey][BearerToken]
[**CustomerAdministratorsDeletePost**](AdministratorsAPI.md#CustomerAdministratorsDeletePost) | **Post** /{customer}/administrators/$delete | Deletes users by UserId in Administrators or DA, and deletes groups/users by UcOids in Administrators. [ServiceKey][BearerToken]
[**CustomerAdministratorsGet**](AdministratorsAPI.md#CustomerAdministratorsGet) | **Get** /{customer}/administrators | Gets all administrators on this customer. [ServiceKey][BearerToken]
[**CustomerAdministratorsIdAccessGet**](AdministratorsAPI.md#CustomerAdministratorsIdAccessGet) | **Get** /{customer}/administrators/{id}/access | Get the access of an administrator. [ServiceKey][BearerToken]
[**CustomerAdministratorsPost**](AdministratorsAPI.md#CustomerAdministratorsPost) | **Post** /{customer}/administrators | Adds an administrator. [BearerToken]
[**CustomerAdministratorsRbacrolesSyncPost**](AdministratorsAPI.md#CustomerAdministratorsRbacrolesSyncPost) | **Post** /{customer}/administrators/rbacroles/$sync | Updates the RBAC roles for the administrators belonging to a customer [ServiceKey]
[**CustomerAdministratorsSyncGroupsPost**](AdministratorsAPI.md#CustomerAdministratorsSyncGroupsPost) | **Post** /{customer}/administrators/$syncGroups | Synchronizes administrator groups with the respective identity provider directories. [ServiceKey][BearerToken]
[**CustomerAdministratorsUcOidGet**](AdministratorsAPI.md#CustomerAdministratorsUcOidGet) | **Get** /{customer}/administrators/{ucOid} | Gets an administrator with ucOid passed a route parameter. [ServiceKey][BearerToken]



## CustomerAdministratorGet

> CitrixCloudServicesAdministratorsApiModelsAdministratorResult CustomerAdministratorGet(ctx, customer).UcOid(ucOid).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).Execute()

Gets an administrator with ucOid passed as a query parameter. [ServiceKey][BearerToken]

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
	ucOid := "ucOid_example" // string | Administrator universal claim organization identifier (OID).
	customer := "capamerica" // string | 
	citrixConsistencyToken := "citrixConsistencyToken_example" // string | CosmosDB consistency token. (optional)
	xCwsTransactionId := "xCwsTransactionId_example" // string | Used for Citrix Cloud telemetry correlation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdministratorsAPI.CustomerAdministratorGet(context.Background(), customer).UcOid(ucOid).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsAPI.CustomerAdministratorGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerAdministratorGet`: CitrixCloudServicesAdministratorsApiModelsAdministratorResult
	fmt.Fprintf(os.Stdout, "Response from `AdministratorsAPI.CustomerAdministratorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerAdministratorGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ucOid** | **string** | Administrator universal claim organization identifier (OID). | 

 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 

### Return type

[**CitrixCloudServicesAdministratorsApiModelsAdministratorResult**](CitrixCloudServicesAdministratorsApiModelsAdministratorResult.md)

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer), [CCAuthService](../README.md#CCAuthService)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerAdministratorsAccessPut

> CustomerAdministratorsAccessPut(ctx, customer).Id(id).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel(citrixCloudServicesAdministratorsApiModelsAdministratorAccessModel).Execute()

Updates the access for an administrator. [ServiceKey][BearerToken]



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
	id := "id_example" // string | The administrator's userId (optional)
	citrixConsistencyToken := "citrixConsistencyToken_example" // string | CosmosDB consistency token. (optional)
	xCwsTransactionId := "xCwsTransactionId_example" // string | Used for Citrix Cloud telemetry correlation. (optional)
	citrixTimeoutMs := TODO // float64 | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds (optional)
	citrixCloudServicesAdministratorsApiModelsAdministratorAccessModel := *openapiclient.NewCitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel(openapiclient.Citrix.CloudServices.Administrators.ApiModels.AdministratorAccessType("Full")) // CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel | The administrator's access, in terms of policies and access type. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdministratorsAPI.CustomerAdministratorsAccessPut(context.Background(), customer).Id(id).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel(citrixCloudServicesAdministratorsApiModelsAdministratorAccessModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsAPI.CustomerAdministratorsAccessPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCustomerAdministratorsAccessPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **string** | The administrator&#39;s userId | 
 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 
 **citrixTimeoutMs** | [**float64**](float64.md) | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds | 
 **citrixCloudServicesAdministratorsApiModelsAdministratorAccessModel** | [**CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel**](CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel.md) | The administrator&#39;s access, in terms of policies and access type. | 

### Return type

 (empty response body)

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer), [CCAuthService](../README.md#CCAuthService)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerAdministratorsAccessUcOidGet

> CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel CustomerAdministratorsAccessUcOidGet(ctx, ucOid, customer).AdministratorType(administratorType).GroupOnlyPermissions(groupOnlyPermissions).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).Execute()

Get the access of an administrator [ServiceKey][BearerToken]



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
	ucOid := "ucOid_example" // string | Administrator universal claim organization identifier (OID).
	customer := "capamerica" // string | 
	administratorType := "administratorType_example" // string | Administrator Type. (optional) (default to "Group")
	groupOnlyPermissions := true // bool | Whether permissions applicable to group administrators should be returned. (optional) (default to true)
	citrixConsistencyToken := "citrixConsistencyToken_example" // string | CosmosDB consistency token. (optional)
	xCwsTransactionId := "xCwsTransactionId_example" // string | Used for Citrix Cloud telemetry correlation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdministratorsAPI.CustomerAdministratorsAccessUcOidGet(context.Background(), ucOid, customer).AdministratorType(administratorType).GroupOnlyPermissions(groupOnlyPermissions).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsAPI.CustomerAdministratorsAccessUcOidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerAdministratorsAccessUcOidGet`: CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel
	fmt.Fprintf(os.Stdout, "Response from `AdministratorsAPI.CustomerAdministratorsAccessUcOidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ucOid** | **string** | Administrator universal claim organization identifier (OID). | 
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerAdministratorsAccessUcOidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **administratorType** | **string** | Administrator Type. | [default to &quot;Group&quot;]
 **groupOnlyPermissions** | **bool** | Whether permissions applicable to group administrators should be returned. | [default to true]
 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 

### Return type

[**CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel**](CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel.md)

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer), [CCAuthService](../README.md#CCAuthService)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerAdministratorsActiveAdministratorsGet

> CitrixCloudServicesAdministratorsApiModelsAdministratorsResult CustomerAdministratorsActiveAdministratorsGet(ctx, customer).MaxItemCount(maxItemCount).RequestContinuation(requestContinuation).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).Execute()

Gets all active administrators (users, secureClients, and groups) across all customers. [ServiceKey][RootOnly]

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
	customer := "root" // string | 
	maxItemCount := "maxItemCount_example" // string | Optional item count. (optional)
	requestContinuation := "requestContinuation_example" // string | Optional continuation token. (optional)
	citrixConsistencyToken := "citrixConsistencyToken_example" // string | CosmosDB consistency token. (optional)
	xCwsTransactionId := "xCwsTransactionId_example" // string | Used for Citrix Cloud telemetry correlation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdministratorsAPI.CustomerAdministratorsActiveAdministratorsGet(context.Background(), customer).MaxItemCount(maxItemCount).RequestContinuation(requestContinuation).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsAPI.CustomerAdministratorsActiveAdministratorsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerAdministratorsActiveAdministratorsGet`: CitrixCloudServicesAdministratorsApiModelsAdministratorsResult
	fmt.Fprintf(os.Stdout, "Response from `AdministratorsAPI.CustomerAdministratorsActiveAdministratorsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerAdministratorsActiveAdministratorsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxItemCount** | **string** | Optional item count. | 
 **requestContinuation** | **string** | Optional continuation token. | 
 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 

### Return type

[**CitrixCloudServicesAdministratorsApiModelsAdministratorsResult**](CitrixCloudServicesAdministratorsApiModelsAdministratorsResult.md)

### Authorization

[CCAuthService](../README.md#CCAuthService)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerAdministratorsAdminIdDelete

> CustomerAdministratorsAdminIdDelete(ctx, adminId, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).Execute()

Deletes an administrator by ucOid. [ServiceKey][BearerToken]

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
	adminId := "adminId_example" // string | Either the administrator universal claim organization identifier (OID) or user id.
	customer := "capamerica" // string | 
	citrixConsistencyToken := "citrixConsistencyToken_example" // string | CosmosDB consistency token. (optional)
	xCwsTransactionId := "xCwsTransactionId_example" // string | Used for Citrix Cloud telemetry correlation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdministratorsAPI.CustomerAdministratorsAdminIdDelete(context.Background(), adminId, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsAPI.CustomerAdministratorsAdminIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adminId** | **string** | Either the administrator universal claim organization identifier (OID) or user id. | 
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerAdministratorsAdminIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 

### Return type

 (empty response body)

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer), [CCAuthService](../README.md#CCAuthService)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerAdministratorsAdminIdPut

> CitrixCloudServicesAdministratorsApiModelsAdministratorResult CustomerAdministratorsAdminIdPut(ctx, adminId, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixCloudServicesAdministratorsApiModelsUpdateAdministratorModel(citrixCloudServicesAdministratorsApiModelsUpdateAdministratorModel).Execute()

Updates an administrator by either ucOid or user id. [BearerToken]

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
	adminId := "adminId_example" // string | Either the administrator universal claim organization identifier (OID) or user id.
	customer := "capamerica" // string | 
	citrixConsistencyToken := "citrixConsistencyToken_example" // string | CosmosDB consistency token. (optional)
	xCwsTransactionId := "xCwsTransactionId_example" // string | Used for Citrix Cloud telemetry correlation. (optional)
	citrixCloudServicesAdministratorsApiModelsUpdateAdministratorModel := *openapiclient.NewCitrixCloudServicesAdministratorsApiModelsUpdateAdministratorModel() // CitrixCloudServicesAdministratorsApiModelsUpdateAdministratorModel | Administrator to be updated. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdministratorsAPI.CustomerAdministratorsAdminIdPut(context.Background(), adminId, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixCloudServicesAdministratorsApiModelsUpdateAdministratorModel(citrixCloudServicesAdministratorsApiModelsUpdateAdministratorModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsAPI.CustomerAdministratorsAdminIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerAdministratorsAdminIdPut`: CitrixCloudServicesAdministratorsApiModelsAdministratorResult
	fmt.Fprintf(os.Stdout, "Response from `AdministratorsAPI.CustomerAdministratorsAdminIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adminId** | **string** | Either the administrator universal claim organization identifier (OID) or user id. | 
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerAdministratorsAdminIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 
 **citrixCloudServicesAdministratorsApiModelsUpdateAdministratorModel** | [**CitrixCloudServicesAdministratorsApiModelsUpdateAdministratorModel**](CitrixCloudServicesAdministratorsApiModelsUpdateAdministratorModel.md) | Administrator to be updated. | 

### Return type

[**CitrixCloudServicesAdministratorsApiModelsAdministratorResult**](CitrixCloudServicesAdministratorsApiModelsAdministratorResult.md)

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerAdministratorsAuthDomainPut

> CustomerAdministratorsAuthDomainPut(ctx, customer).OldAuthDomain(oldAuthDomain).NewAuthDomain(newAuthDomain).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).Execute()

Updates auth domain for relevant administrators of given customer context.  Forwards to DAv1 if not found locally or is pending. [ServiceKey]

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
	oldAuthDomain := "oldAuthDomain_example" // string | The old auth domain to look for.
	newAuthDomain := "newAuthDomain_example" // string | The new auth domain to update.
	customer := "capamerica" // string | 
	citrixConsistencyToken := "citrixConsistencyToken_example" // string | CosmosDB consistency token. (optional)
	xCwsTransactionId := "xCwsTransactionId_example" // string | Used for Citrix Cloud telemetry correlation. (optional)
	citrixTimeoutMs := TODO // float64 | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdministratorsAPI.CustomerAdministratorsAuthDomainPut(context.Background(), customer).OldAuthDomain(oldAuthDomain).NewAuthDomain(newAuthDomain).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixTimeoutMs(citrixTimeoutMs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsAPI.CustomerAdministratorsAuthDomainPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCustomerAdministratorsAuthDomainPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oldAuthDomain** | **string** | The old auth domain to look for. | 
 **newAuthDomain** | **string** | The new auth domain to update. | 

 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 
 **citrixTimeoutMs** | [**float64**](float64.md) | Maximum time caller is willing to wait for a response before timing out the request, in milliseconds | 

### Return type

 (empty response body)

### Authorization

[CCAuthService](../README.md#CCAuthService)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerAdministratorsClaimsDecompressPost

> CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse CustomerAdministratorsClaimsDecompressPost(ctx, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModel(citrixCloudServicesAdministratorsApiModelsDecompressClaimsModel).Execute()

Decompresses the group claims. [ServiceKey][BearerToken]

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
	citrixCloudServicesAdministratorsApiModelsDecompressClaimsModel := *openapiclient.NewCitrixCloudServicesAdministratorsApiModelsDecompressClaimsModel() // CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModel | The model. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdministratorsAPI.CustomerAdministratorsClaimsDecompressPost(context.Background(), customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModel(citrixCloudServicesAdministratorsApiModelsDecompressClaimsModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsAPI.CustomerAdministratorsClaimsDecompressPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerAdministratorsClaimsDecompressPost`: CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse
	fmt.Fprintf(os.Stdout, "Response from `AdministratorsAPI.CustomerAdministratorsClaimsDecompressPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerAdministratorsClaimsDecompressPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 
 **citrixCloudServicesAdministratorsApiModelsDecompressClaimsModel** | [**CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModel**](CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModel.md) | The model. | 

### Return type

[**CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse**](CitrixCloudServicesAdministratorsApiModelsDecompressClaimsModelResponse.md)

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer), [CCAuthService](../README.md#CCAuthService)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerAdministratorsCreatePost

> CitrixCloudServicesAdministratorsApiModelsAdministratorResult CustomerAdministratorsCreatePost(ctx, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel(citrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel).Execute()

Adds an administrator. This API is a higher-level version of  `POST /{customer}/administrators` that works with any type of  administrator with stricter input validation. [BearerToken]



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
	citrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel := *openapiclient.NewCitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel("Type_example", "AccessType_example", "ProviderType_example") // CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel | Administrator to be added. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdministratorsAPI.CustomerAdministratorsCreatePost(context.Background(), customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel(citrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsAPI.CustomerAdministratorsCreatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerAdministratorsCreatePost`: CitrixCloudServicesAdministratorsApiModelsAdministratorResult
	fmt.Fprintf(os.Stdout, "Response from `AdministratorsAPI.CustomerAdministratorsCreatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerAdministratorsCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 
 **citrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel** | [**CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel**](CitrixCloudServicesAdministratorsApiModelsCreateAdministratorInputModel.md) | Administrator to be added. | 

### Return type

[**CitrixCloudServicesAdministratorsApiModelsAdministratorResult**](CitrixCloudServicesAdministratorsApiModelsAdministratorResult.md)

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerAdministratorsDeleteMultiplePost

> CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsErrorResponseModel CustomerAdministratorsDeleteMultiplePost(ctx, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel(citrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel).Execute()

Deletes an administrator. [ServiceKey][BearerToken]

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
	citrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel := *openapiclient.NewCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel([]string{"UcoIds_example"}) // CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel | Administrator delete model. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdministratorsAPI.CustomerAdministratorsDeleteMultiplePost(context.Background(), customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel(citrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsAPI.CustomerAdministratorsDeleteMultiplePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerAdministratorsDeleteMultiplePost`: CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsErrorResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AdministratorsAPI.CustomerAdministratorsDeleteMultiplePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerAdministratorsDeleteMultiplePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 
 **citrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel** | [**CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel**](CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorModel.md) | Administrator delete model. | 

### Return type

[**CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsErrorResponseModel**](CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsErrorResponseModel.md)

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer), [CCAuthService](../README.md#CCAuthService)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerAdministratorsDeletePost

> CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsErrorResponseModel CustomerAdministratorsDeletePost(ctx, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel(citrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel).Execute()

Deletes users by UserId in Administrators or DA, and deletes groups/users by UcOids in Administrators. [ServiceKey][BearerToken]

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
	citrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel := *openapiclient.NewCitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel() // CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel | Administrator delete model. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdministratorsAPI.CustomerAdministratorsDeletePost(context.Background(), customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel(citrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsAPI.CustomerAdministratorsDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerAdministratorsDeletePost`: CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsErrorResponseModel
	fmt.Fprintf(os.Stdout, "Response from `AdministratorsAPI.CustomerAdministratorsDeletePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerAdministratorsDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 
 **citrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel** | [**CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel**](CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsModel.md) | Administrator delete model. | 

### Return type

[**CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsErrorResponseModel**](CitrixCloudServicesAdministratorsApiModelsDeleteAdministratorsErrorResponseModel.md)

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer), [CCAuthService](../README.md#CCAuthService)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerAdministratorsGet

> CitrixCloudServicesAdministratorsApiModelsAdministratorsResult CustomerAdministratorsGet(ctx, customer).Type_(type_).ProviderType(providerType).MaxItemCount(maxItemCount).RequestContinuation(requestContinuation).Id(id).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).Execute()

Gets all administrators on this customer. [ServiceKey][BearerToken]

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
	type_ := "type__example" // string | Optional administrator type filter. (optional)
	providerType := "providerType_example" // string | Optional Type of the provider. (optional)
	maxItemCount := "maxItemCount_example" // string | Optional item count. (optional)
	requestContinuation := "requestContinuation_example" // string | Optional continuation token. (optional)
	id := "id_example" // string | The User Id corresponding to a user administrator. (Note: Exclusive with other parameters on this endpoint) (optional)
	citrixConsistencyToken := "citrixConsistencyToken_example" // string | CosmosDB consistency token. (optional)
	xCwsTransactionId := "xCwsTransactionId_example" // string | Used for Citrix Cloud telemetry correlation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdministratorsAPI.CustomerAdministratorsGet(context.Background(), customer).Type_(type_).ProviderType(providerType).MaxItemCount(maxItemCount).RequestContinuation(requestContinuation).Id(id).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsAPI.CustomerAdministratorsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerAdministratorsGet`: CitrixCloudServicesAdministratorsApiModelsAdministratorsResult
	fmt.Fprintf(os.Stdout, "Response from `AdministratorsAPI.CustomerAdministratorsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerAdministratorsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Optional administrator type filter. | 
 **providerType** | **string** | Optional Type of the provider. | 
 **maxItemCount** | **string** | Optional item count. | 
 **requestContinuation** | **string** | Optional continuation token. | 
 **id** | **string** | The User Id corresponding to a user administrator. (Note: Exclusive with other parameters on this endpoint) | 
 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 

### Return type

[**CitrixCloudServicesAdministratorsApiModelsAdministratorsResult**](CitrixCloudServicesAdministratorsApiModelsAdministratorsResult.md)

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer), [CCAuthService](../README.md#CCAuthService)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerAdministratorsIdAccessGet

> CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel CustomerAdministratorsIdAccessGet(ctx, id, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).Execute()

Get the access of an administrator. [ServiceKey][BearerToken]

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
	id := "id_example" // string | Either the Administrator UserId or UcOid (universal claim organization identifier [OID]). Use UserId for Citrix and AzureAd Users, UcOid for all others.
	customer := "capamerica" // string | 
	citrixConsistencyToken := "citrixConsistencyToken_example" // string | CosmosDB consistency token. (optional)
	xCwsTransactionId := "xCwsTransactionId_example" // string | Used for Citrix Cloud telemetry correlation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdministratorsAPI.CustomerAdministratorsIdAccessGet(context.Background(), id, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsAPI.CustomerAdministratorsIdAccessGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerAdministratorsIdAccessGet`: CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel
	fmt.Fprintf(os.Stdout, "Response from `AdministratorsAPI.CustomerAdministratorsIdAccessGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Either the Administrator UserId or UcOid (universal claim organization identifier [OID]). Use UserId for Citrix and AzureAd Users, UcOid for all others. | 
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerAdministratorsIdAccessGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 

### Return type

[**CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel**](CitrixCloudServicesAdministratorsApiModelsAdministratorAccessModel.md)

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer), [CCAuthService](../README.md#CCAuthService)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerAdministratorsPost

> CitrixCloudServicesAdministratorsApiModelsAdministratorResult CustomerAdministratorsPost(ctx, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel(citrixCloudServicesAdministratorsApiModelsCreateAdministratorModel).Execute()

Adds an administrator. [BearerToken]

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
	citrixCloudServicesAdministratorsApiModelsCreateAdministratorModel := *openapiclient.NewCitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel("UcOid_example", "Type_example", "AccessType_example", "ProviderType_example", "ProviderId_example", "AuthDomain_example") // CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel | Administrator to be added. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdministratorsAPI.CustomerAdministratorsPost(context.Background(), customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel(citrixCloudServicesAdministratorsApiModelsCreateAdministratorModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsAPI.CustomerAdministratorsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerAdministratorsPost`: CitrixCloudServicesAdministratorsApiModelsAdministratorResult
	fmt.Fprintf(os.Stdout, "Response from `AdministratorsAPI.CustomerAdministratorsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerAdministratorsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 
 **citrixCloudServicesAdministratorsApiModelsCreateAdministratorModel** | [**CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel**](CitrixCloudServicesAdministratorsApiModelsCreateAdministratorModel.md) | Administrator to be added. | 

### Return type

[**CitrixCloudServicesAdministratorsApiModelsAdministratorResult**](CitrixCloudServicesAdministratorsApiModelsAdministratorResult.md)

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerAdministratorsRbacrolesSyncPost

> CustomerAdministratorsRbacrolesSyncPost(ctx, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest(citrixCloudServicesAdministratorsApiModelsRoleOperationRequest).Execute()

Updates the RBAC roles for the administrators belonging to a customer [ServiceKey]

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
	citrixCloudServicesAdministratorsApiModelsRoleOperationRequest := *openapiclient.NewCitrixCloudServicesAdministratorsApiModelsRoleOperationRequest([]string{"RbacRoles_example"}, "Operation_example") // CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdministratorsAPI.CustomerAdministratorsRbacrolesSyncPost(context.Background(), customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest(citrixCloudServicesAdministratorsApiModelsRoleOperationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsAPI.CustomerAdministratorsRbacrolesSyncPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCustomerAdministratorsRbacrolesSyncPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 
 **citrixCloudServicesAdministratorsApiModelsRoleOperationRequest** | [**CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest**](CitrixCloudServicesAdministratorsApiModelsRoleOperationRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[CCAuthService](../README.md#CCAuthService)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerAdministratorsSyncGroupsPost

> CitrixCloudServicesAdministratorsApiModelsAdministratorsResult CustomerAdministratorsSyncGroupsPost(ctx, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).Execute()

Synchronizes administrator groups with the respective identity provider directories. [ServiceKey][BearerToken]

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
	resp, r, err := apiClient.AdministratorsAPI.CustomerAdministratorsSyncGroupsPost(context.Background(), customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsAPI.CustomerAdministratorsSyncGroupsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerAdministratorsSyncGroupsPost`: CitrixCloudServicesAdministratorsApiModelsAdministratorsResult
	fmt.Fprintf(os.Stdout, "Response from `AdministratorsAPI.CustomerAdministratorsSyncGroupsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerAdministratorsSyncGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 

### Return type

[**CitrixCloudServicesAdministratorsApiModelsAdministratorsResult**](CitrixCloudServicesAdministratorsApiModelsAdministratorsResult.md)

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer), [CCAuthService](../README.md#CCAuthService)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerAdministratorsUcOidGet

> CitrixCloudServicesAdministratorsApiModelsAdministratorResult CustomerAdministratorsUcOidGet(ctx, ucOid, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).Execute()

Gets an administrator with ucOid passed a route parameter. [ServiceKey][BearerToken]

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
	ucOid := "ucOid_example" // string | Administrator universal claim organization identifier (OID).
	customer := "capamerica" // string | 
	citrixConsistencyToken := "citrixConsistencyToken_example" // string | CosmosDB consistency token. (optional)
	xCwsTransactionId := "xCwsTransactionId_example" // string | Used for Citrix Cloud telemetry correlation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdministratorsAPI.CustomerAdministratorsUcOidGet(context.Background(), ucOid, customer).CitrixConsistencyToken(citrixConsistencyToken).XCwsTransactionId(xCwsTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsAPI.CustomerAdministratorsUcOidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerAdministratorsUcOidGet`: CitrixCloudServicesAdministratorsApiModelsAdministratorResult
	fmt.Fprintf(os.Stdout, "Response from `AdministratorsAPI.CustomerAdministratorsUcOidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ucOid** | **string** | Administrator universal claim organization identifier (OID). | 
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerAdministratorsUcOidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixConsistencyToken** | **string** | CosmosDB consistency token. | 
 **xCwsTransactionId** | **string** | Used for Citrix Cloud telemetry correlation. | 

### Return type

[**CitrixCloudServicesAdministratorsApiModelsAdministratorResult**](CitrixCloudServicesAdministratorsApiModelsAdministratorResult.md)

### Authorization

[CCAuthBearer](../README.md#CCAuthBearer), [CCAuthService](../README.md#CCAuthService)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

