# \DatabasesTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatabasesTPChanageDatabase**](DatabasesTPApi.md#DatabasesTPChanageDatabase) | **Patch** /techpreview/{customerid}/{siteid}/Databases/{dataStore} | Change database action.
[**DatabasesTPGenerateDatabaseChangeScript**](DatabasesTPApi.md#DatabasesTPGenerateDatabaseChangeScript) | **Post** /techpreview/{customerid}/{siteid}/Databases/{dataStore}/$generateDatabaseChangeScript | generate sql script for manually change a database.
[**DatabasesTPGetDatabase**](DatabasesTPApi.md#DatabasesTPGetDatabase) | **Get** /techpreview/{customerid}/{siteid}/Databases/{dataStore} | Get a single database.
[**DatabasesTPGetDatabases**](DatabasesTPApi.md#DatabasesTPGetDatabases) | **Get** /techpreview/{customerid}/{siteid}/Databases | Get the list of all databases.



## DatabasesTPChanageDatabase

> bool DatabasesTPChanageDatabase(ctx, dataStore, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Change database action.

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
    dataStore := "dataStore_example" // string | Datastore.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewDatabaseChangePropertiesModel("DatabaseName_example", "DatabaseServerAddress_example", false, false) // DatabaseChangePropertiesModel | Details about the database to change.
    async := true // bool | If `true`, the database will be changed in a background task. The task will have JobType ChangeDatabase. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesTPApi.DatabasesTPChanageDatabase(context.Background(), dataStore, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesTPApi.DatabasesTPChanageDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabasesTPChanageDatabase`: bool
    fmt.Fprintf(os.Stdout, "Response from `DatabasesTPApi.DatabasesTPChanageDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataStore** | **string** | Datastore. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabasesTPChanageDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**DatabaseChangePropertiesModel**](DatabaseChangePropertiesModel.md) | Details about the database to change. | 
 **async** | **bool** | If &#x60;true&#x60;, the database will be changed in a background task. The task will have JobType ChangeDatabase. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

**bool**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabasesTPGenerateDatabaseChangeScript

> DatabaseChangeScriptModelCollection DatabasesTPGenerateDatabaseChangeScript(ctx, dataStore, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

generate sql script for manually change a database.

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
    dataStore := "dataStore_example" // string | Datastore.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewDatabaseChangeScriptPropertiesModel("DatabaseName_example", "DatabaseServerAddress_example") // DatabaseChangeScriptPropertiesModel | Details about the database to generate change script.
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesTPApi.DatabasesTPGenerateDatabaseChangeScript(context.Background(), dataStore, customerid, siteid).Request(request).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesTPApi.DatabasesTPGenerateDatabaseChangeScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabasesTPGenerateDatabaseChangeScript`: DatabaseChangeScriptModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DatabasesTPApi.DatabasesTPGenerateDatabaseChangeScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataStore** | **string** | Datastore. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabasesTPGenerateDatabaseChangeScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**DatabaseChangeScriptPropertiesModel**](DatabaseChangeScriptPropertiesModel.md) | Details about the database to generate change script. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**DatabaseChangeScriptModelCollection**](DatabaseChangeScriptModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabasesTPGetDatabase

> DatabaseResponseModel DatabasesTPGetDatabase(ctx, dataStore, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get a single database.

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
    dataStore := "dataStore_example" // string | Datastore.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesTPApi.DatabasesTPGetDatabase(context.Background(), dataStore, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesTPApi.DatabasesTPGetDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabasesTPGetDatabase`: DatabaseResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DatabasesTPApi.DatabasesTPGetDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataStore** | **string** | Datastore. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabasesTPGetDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**DatabaseResponseModel**](DatabaseResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatabasesTPGetDatabases

> DatabaseResponseModelCollection DatabasesTPGetDatabases(ctx, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the list of all databases.

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
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatabasesTPApi.DatabasesTPGetDatabases(context.Background(), customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesTPApi.DatabasesTPGetDatabases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatabasesTPGetDatabases`: DatabaseResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DatabasesTPApi.DatabasesTPGetDatabases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatabasesTPGetDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**DatabaseResponseModelCollection**](DatabaseResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

