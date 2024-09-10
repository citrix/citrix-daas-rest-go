# \AdministratorsAPI

All URIs are relative to *https://api.cloud.com/administrators*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAdministrator**](AdministratorsAPI.md#CreateAdministrator) | **Post** /$create | Create a new CC administrator.
[**DeleteAdministrator**](AdministratorsAPI.md#DeleteAdministrator) | **Delete** /{id} | Remove a CC administrator.
[**DeleteInvitation**](AdministratorsAPI.md#DeleteInvitation) | **Delete** /invitations | Remove a CC pending user administrator invitation.
[**FetchAdministrators**](AdministratorsAPI.md#FetchAdministrators) | **Get** / | Fetch all CC administrators.
[**GetAdministratorAccess**](AdministratorsAPI.md#GetAdministratorAccess) | **Get** /{id}/access | Fetch the access of an administrator.
[**UpdateAdministratorAccess**](AdministratorsAPI.md#UpdateAdministratorAccess) | **Put** /access | Update roles and permissions of an existing CC administrator.



## CreateAdministrator

> AdministratorResult CreateAdministrator(ctx).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CreateAdministratorInputModel(createAdministratorInputModel).Execute()

Create a new CC administrator.

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
    authorization := "authorization_example" // string | Access token.
    citrixCustomerId := "citrixCustomerId_example" // string | Customer ID.
    createAdministratorInputModel := *openapiclient.NewCreateAdministratorInputModel("AccessType_example", openapiclient.AdministratorProviderType("Ad"), "Type_example") // CreateAdministratorInputModel | Administrator to be added. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdministratorsAPI.CreateAdministrator(context.Background()).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CreateAdministratorInputModel(createAdministratorInputModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsAPI.CreateAdministrator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAdministrator`: AdministratorResult
    fmt.Fprintf(os.Stdout, "Response from `AdministratorsAPI.CreateAdministrator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdministratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Access token. | 
 **citrixCustomerId** | **string** | Customer ID. | 
 **createAdministratorInputModel** | [**CreateAdministratorInputModel**](CreateAdministratorInputModel.md) | Administrator to be added. | 

### Return type

[**AdministratorResult**](AdministratorResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAdministrator

> DeleteAdministrator(ctx, id).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Execute()

Remove a CC administrator.

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
    id := "id_example" // string | CC administrator ID.
    authorization := "authorization_example" // string | Access token.
    citrixCustomerId := "citrixCustomerId_example" // string | Customer ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdministratorsAPI.DeleteAdministrator(context.Background(), id).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsAPI.DeleteAdministrator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | CC administrator ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdministratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Access token. | 
 **citrixCustomerId** | **string** | Customer ID. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInvitation

> bool DeleteInvitation(ctx).Email(email).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Execute()

Remove a CC pending user administrator invitation.

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
    email := "email_example" // string | Pending user administrator invitation email.
    authorization := "authorization_example" // string | Access token.
    citrixCustomerId := "citrixCustomerId_example" // string | Customer ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdministratorsAPI.DeleteInvitation(context.Background()).Email(email).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsAPI.DeleteInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInvitation`: bool
    fmt.Fprintf(os.Stdout, "Response from `AdministratorsAPI.DeleteInvitation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Pending user administrator invitation email. | 
 **authorization** | **string** | Access token. | 
 **citrixCustomerId** | **string** | Customer ID. | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAdministrators

> AdministratorsResult FetchAdministrators(ctx).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Type_(type_).ProviderType(providerType).MaxItemCount(maxItemCount).RequestContinuation(requestContinuation).Id(id).Execute()

Fetch all CC administrators.

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
    authorization := "authorization_example" // string | Access token.
    citrixCustomerId := "citrixCustomerId_example" // string | Customer ID.
    type_ := "type__example" // string | Optional administrator type filter. (optional)
    providerType := openapiclient.AdministratorProviderType("Ad") // AdministratorProviderType | Optional provider type. (optional)
    maxItemCount := "maxItemCount_example" // string | Optional item count. (optional)
    requestContinuation := "requestContinuation_example" // string | Optional continuation token. (optional)
    id := "id_example" // string | User ID corresponding to a user administrator. (Note: Exclusive with other parameters on this endpoint.) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdministratorsAPI.FetchAdministrators(context.Background()).Authorization(authorization).CitrixCustomerId(citrixCustomerId).Type_(type_).ProviderType(providerType).MaxItemCount(maxItemCount).RequestContinuation(requestContinuation).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsAPI.FetchAdministrators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAdministrators`: AdministratorsResult
    fmt.Fprintf(os.Stdout, "Response from `AdministratorsAPI.FetchAdministrators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchAdministratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Access token. | 
 **citrixCustomerId** | **string** | Customer ID. | 
 **type_** | **string** | Optional administrator type filter. | 
 **providerType** | [**AdministratorProviderType**](AdministratorProviderType.md) | Optional provider type. | 
 **maxItemCount** | **string** | Optional item count. | 
 **requestContinuation** | **string** | Optional continuation token. | 
 **id** | **string** | User ID corresponding to a user administrator. (Note: Exclusive with other parameters on this endpoint.) | 

### Return type

[**AdministratorsResult**](AdministratorsResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdministratorAccess

> AdministratorAccessModel GetAdministratorAccess(ctx, id).Execute()

Fetch the access of an administrator.

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
    id := "id_example" // string | UserId for AzureAd and Citrix User Administrators. UcOid for all others.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdministratorsAPI.GetAdministratorAccess(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsAPI.GetAdministratorAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdministratorAccess`: AdministratorAccessModel
    fmt.Fprintf(os.Stdout, "Response from `AdministratorsAPI.GetAdministratorAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UserId for AzureAd and Citrix User Administrators. UcOid for all others. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdministratorAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdministratorAccessModel**](AdministratorAccessModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdministratorAccess

> UpdateAdministratorAccess(ctx).Id(id).Authorization(authorization).CitrixCustomerId(citrixCustomerId).AdministratorAccessModel(administratorAccessModel).Execute()

Update roles and permissions of an existing CC administrator.

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
    id := "id_example" // string | CC administrator ID.
    authorization := "authorization_example" // string | Access token.
    citrixCustomerId := "citrixCustomerId_example" // string | Customer ID.
    administratorAccessModel := *openapiclient.NewAdministratorAccessModel(openapiclient.AdministratorAccessType("Full")) // AdministratorAccessModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdministratorsAPI.UpdateAdministratorAccess(context.Background()).Id(id).Authorization(authorization).CitrixCustomerId(citrixCustomerId).AdministratorAccessModel(administratorAccessModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsAPI.UpdateAdministratorAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdministratorAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | CC administrator ID. | 
 **authorization** | **string** | Access token. | 
 **citrixCustomerId** | **string** | Customer ID. | 
 **administratorAccessModel** | [**AdministratorAccessModel**](AdministratorAccessModel.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

