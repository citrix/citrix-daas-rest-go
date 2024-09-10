# \AuthDomainsDAAS

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerAuthDomainsCheckGet**](AuthDomainsDAAS.md#CustomerAuthDomainsCheckGet) | **Get** /{customer}/AuthDomains/Check | 
[**CustomerAuthDomainsPut**](AuthDomainsDAAS.md#CustomerAuthDomainsPut) | **Put** /{customer}/AuthDomains | 



## CustomerAuthDomainsCheckGet

> bool CustomerAuthDomainsCheckGet(ctx, customer).Name(name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixcws"
)

func main() {
    name := "name_example" // string | 
    customer := "customer_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthDomainsDAAS.CustomerAuthDomainsCheckGet(context.Background(), customer).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthDomainsDAAS.CustomerAuthDomainsCheckGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerAuthDomainsCheckGet`: bool
    fmt.Fprintf(os.Stdout, "Response from `AuthDomainsDAAS.CustomerAuthDomainsCheckGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerAuthDomainsCheckGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 


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


## CustomerAuthDomainsPut

> UpdatedCustomerModel CustomerAuthDomainsPut(ctx, customer).OldName(oldName).NewName(newName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/citrixcws"
)

func main() {
    oldName := "oldName_example" // string | 
    newName := "newName_example" // string | 
    customer := "customer_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthDomainsDAAS.CustomerAuthDomainsPut(context.Background(), customer).OldName(oldName).NewName(newName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthDomainsDAAS.CustomerAuthDomainsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerAuthDomainsPut`: UpdatedCustomerModel
    fmt.Fprintf(os.Stdout, "Response from `AuthDomainsDAAS.CustomerAuthDomainsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerAuthDomainsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oldName** | **string** |  | 
 **newName** | **string** |  | 


### Return type

[**UpdatedCustomerModel**](UpdatedCustomerModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

