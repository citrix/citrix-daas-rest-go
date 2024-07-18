# \AccountQCS

All URIs are relative to *https://api.cloud.com/quickcreateservice*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAccountAsync**](AccountQCS.md#AddAccountAsync) | **Post** /{customerId}/accounts | Add a account
[**DeleteCustomerAccountAsync**](AccountQCS.md#DeleteCustomerAccountAsync) | **Delete** /{customerId}/accounts/{accountId} | Deletes the account configured for the specified customer
[**GetAccountResourcesAsync**](AccountQCS.md#GetAccountResourcesAsync) | **Post** /{customerId}/accounts/{accountId}/resources/$search | Get the account resources for the specified customer
[**GetCustomerAccountAsync**](AccountQCS.md#GetCustomerAccountAsync) | **Get** /{customerId}/accounts/{accountId} | Get the account configured for the specified customer
[**GetCustomerAccountResourcesAsync**](AccountQCS.md#GetCustomerAccountResourcesAsync) | **Post** /{customerId}/accounts/resources/$search | Get the account resources for the specified customer
[**GetCustomerAccountTaskAsync**](AccountQCS.md#GetCustomerAccountTaskAsync) | **Get** /{customerId}/accounts/{accountId}/tasks/{taskId} | Gets account task
[**GetCustomerAccountsAsync**](AccountQCS.md#GetCustomerAccountsAsync) | **Get** /{customerId}/accounts | Get the accounts configured for the specified customer
[**InititateAccountTaskAsync**](AccountQCS.md#InititateAccountTaskAsync) | **Post** /{customerId}/accounts/{accountId}/tasks | Registers account BYOL account
[**UpdateCustomerAccountAsync**](AccountQCS.md#UpdateCustomerAccountAsync) | **Patch** /{customerId}/accounts/{accountId} | Updates the account access keys for the specified customer



## AddAccountAsync

> AwsEdcAccount AddAccountAsync(ctx, customerId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Add a account

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
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    body := AddAwsEdcAccount(987) // AddAwsEdcAccount | Configuration of the account to add (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountQCS.AddAccountAsync(context.Background(), customerId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountQCS.AddAccountAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAccountAsync`: AwsEdcAccount
    fmt.Fprintf(os.Stdout, "Response from `AccountQCS.AddAccountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAccountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixTransactionId** | **string** | The Transaction Id. | 
 **body** | **AddAwsEdcAccount** | Configuration of the account to add | 

### Return type

[**AwsEdcAccount**](AwsEdcAccount.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomerAccountAsync

> DeleteCustomerAccountAsync(ctx, customerId, accountId).CitrixTransactionId(citrixTransactionId).Execute()

Deletes the account configured for the specified customer

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
    accountId := "accountId_example" // string | ID of the account
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccountQCS.DeleteCustomerAccountAsync(context.Background(), customerId, accountId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountQCS.DeleteCustomerAccountAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**accountId** | **string** | ID of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomerAccountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

 (empty response body)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountResourcesAsync

> AccountResources GetAccountResourcesAsync(ctx, customerId, accountId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Get the account resources for the specified customer

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
    accountId := "accountId_example" // string | ID of the account
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    body := SearchAwsEdcAccountResourceRequest(987) // SearchAwsEdcAccountResourceRequest | Search request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountQCS.GetAccountResourcesAsync(context.Background(), customerId, accountId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountQCS.GetAccountResourcesAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountResourcesAsync`: AccountResources
    fmt.Fprintf(os.Stdout, "Response from `AccountQCS.GetAccountResourcesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**accountId** | **string** | ID of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountResourcesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixTransactionId** | **string** | The Transaction Id. | 
 **body** | **SearchAwsEdcAccountResourceRequest** | Search request | 

### Return type

[**AccountResources**](AccountResources.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerAccountAsync

> AwsEdcAccount GetCustomerAccountAsync(ctx, customerId, accountId).CitrixTransactionId(citrixTransactionId).Execute()

Get the account configured for the specified customer

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
    accountId := "accountId_example" // string | ID of the account
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountQCS.GetCustomerAccountAsync(context.Background(), customerId, accountId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountQCS.GetCustomerAccountAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerAccountAsync`: AwsEdcAccount
    fmt.Fprintf(os.Stdout, "Response from `AccountQCS.GetCustomerAccountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**accountId** | **string** | ID of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerAccountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**AwsEdcAccount**](AwsEdcAccount.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerAccountResourcesAsync

> AccountResources GetCustomerAccountResourcesAsync(ctx, customerId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Get the account resources for the specified customer

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
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    body := SearchAwsEdcAccountResourceRequest(987) // SearchAwsEdcAccountResourceRequest | Search request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountQCS.GetCustomerAccountResourcesAsync(context.Background(), customerId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountQCS.GetCustomerAccountResourcesAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerAccountResourcesAsync`: AccountResources
    fmt.Fprintf(os.Stdout, "Response from `AccountQCS.GetCustomerAccountResourcesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerAccountResourcesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixTransactionId** | **string** | The Transaction Id. | 
 **body** | **SearchAwsEdcAccountResourceRequest** | Search request | 

### Return type

[**AccountResources**](AccountResources.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerAccountTaskAsync

> GetCustomerAccountTaskAsync200Response GetCustomerAccountTaskAsync(ctx, customerId, accountId, taskId).CitrixTransactionId(citrixTransactionId).Execute()

Gets account task

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
    accountId := "accountId_example" // string | ID of the account
    taskId := "taskId_example" // string | ID of task
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountQCS.GetCustomerAccountTaskAsync(context.Background(), customerId, accountId, taskId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountQCS.GetCustomerAccountTaskAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerAccountTaskAsync`: GetCustomerAccountTaskAsync200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountQCS.GetCustomerAccountTaskAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**accountId** | **string** | ID of the account | 
**taskId** | **string** | ID of task | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerAccountTaskAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**GetCustomerAccountTaskAsync200Response**](GetCustomerAccountTaskAsync200Response.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerAccountsAsync

> Accounts GetCustomerAccountsAsync(ctx, customerId).AccountType(accountType).FetchDetails(fetchDetails).CitrixTransactionId(citrixTransactionId).Execute()

Get the accounts configured for the specified customer

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
    accountType := openapiclient.AccountType("AWSEDC") // AccountType | Account Type (optional)
    fetchDetails := true // bool | If true, checks if all accounts have the required set of permissions for assume role tasks (optional) (default to false)
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountQCS.GetCustomerAccountsAsync(context.Background(), customerId).AccountType(accountType).FetchDetails(fetchDetails).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountQCS.GetCustomerAccountsAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerAccountsAsync`: Accounts
    fmt.Fprintf(os.Stdout, "Response from `AccountQCS.GetCustomerAccountsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerAccountsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountType** | [**AccountType**](AccountType.md) | Account Type | 
 **fetchDetails** | **bool** | If true, checks if all accounts have the required set of permissions for assume role tasks | [default to false]
 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**Accounts**](Accounts.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InititateAccountTaskAsync

> AwsEdcRegisterAccount InititateAccountTaskAsync(ctx, customerId, accountId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Registers account BYOL account

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
    accountId := "accountId_example" // string | ID of the account
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    body := RegisterAwsEdcAccountTask(987) // RegisterAwsEdcAccountTask | Task request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountQCS.InititateAccountTaskAsync(context.Background(), customerId, accountId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountQCS.InititateAccountTaskAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InititateAccountTaskAsync`: AwsEdcRegisterAccount
    fmt.Fprintf(os.Stdout, "Response from `AccountQCS.InititateAccountTaskAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**accountId** | **string** | ID of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiInititateAccountTaskAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixTransactionId** | **string** | The Transaction Id. | 
 **body** | **RegisterAwsEdcAccountTask** | Task request | 

### Return type

[**AwsEdcRegisterAccount**](AwsEdcRegisterAccount.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomerAccountAsync

> UpdateCustomerAccountAsync(ctx, customerId, accountId).CitrixTransactionId(citrixTransactionId).UpdateCustomerAccountAsyncRequest(updateCustomerAccountAsyncRequest).Execute()

Updates the account access keys for the specified customer

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
    accountId := "accountId_example" // string | ID of the account
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    updateCustomerAccountAsyncRequest := openapiclient.UpdateCustomerAccountAsync_request{UpdateAccount: openapiclient.NewUpdateAccount(openapiclient.UpdateAccountOperationType("RenameAccount"))} // UpdateCustomerAccountAsyncRequest | Configuration of the account to update (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccountQCS.UpdateCustomerAccountAsync(context.Background(), customerId, accountId).CitrixTransactionId(citrixTransactionId).UpdateCustomerAccountAsyncRequest(updateCustomerAccountAsyncRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountQCS.UpdateCustomerAccountAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**accountId** | **string** | ID of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerAccountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixTransactionId** | **string** | The Transaction Id. | 
 **updateCustomerAccountAsyncRequest** | [**UpdateCustomerAccountAsyncRequest**](UpdateCustomerAccountAsyncRequest.md) | Configuration of the account to update | 

### Return type

 (empty response body)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

