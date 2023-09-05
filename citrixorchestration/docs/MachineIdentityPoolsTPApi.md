# \MachineIdentityPoolsTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MachineIdentityPoolsTPGetMachineIdentityPool**](MachineIdentityPoolsTPApi.md#MachineIdentityPoolsTPGetMachineIdentityPool) | **Get** /techpreview/{customerid}/{siteid}/MachineIdentityPools/{id} | Get existing machine identity pool with id.
[**MachineIdentityPoolsTPGetMachineIdentityPools**](MachineIdentityPoolsTPApi.md#MachineIdentityPoolsTPGetMachineIdentityPools) | **Get** /techpreview/{customerid}/{siteid}/MachineIdentityPools | Get all existing machine identity pools.
[**MachineIdentityPoolsTPGetProvSchemeReferences**](MachineIdentityPoolsTPApi.md#MachineIdentityPoolsTPGetProvSchemeReferences) | **Get** /techpreview/{customerid}/{siteid}/MachineIdentityPools/{identityPoolId}/Provschemes | List all provschemes that reference the given identitypool.



## MachineIdentityPoolsTPGetMachineIdentityPool

> MachineIdentityPoolResponseModel MachineIdentityPoolsTPGetMachineIdentityPool(ctx, id, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get existing machine identity pool with id.

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
    id := "id_example" // string | Machine identity pool id
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineIdentityPoolsTPApi.MachineIdentityPoolsTPGetMachineIdentityPool(context.Background(), id, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineIdentityPoolsTPApi.MachineIdentityPoolsTPGetMachineIdentityPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineIdentityPoolsTPGetMachineIdentityPool`: MachineIdentityPoolResponseModel
    fmt.Fprintf(os.Stdout, "Response from `MachineIdentityPoolsTPApi.MachineIdentityPoolsTPGetMachineIdentityPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Machine identity pool id | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineIdentityPoolsTPGetMachineIdentityPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**MachineIdentityPoolResponseModel**](MachineIdentityPoolResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineIdentityPoolsTPGetMachineIdentityPools

> MachineIdentityPoolResponseModelCollection MachineIdentityPoolsTPGetMachineIdentityPools(ctx, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get all existing machine identity pools.



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
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineIdentityPoolsTPApi.MachineIdentityPoolsTPGetMachineIdentityPools(context.Background(), customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineIdentityPoolsTPApi.MachineIdentityPoolsTPGetMachineIdentityPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineIdentityPoolsTPGetMachineIdentityPools`: MachineIdentityPoolResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineIdentityPoolsTPApi.MachineIdentityPoolsTPGetMachineIdentityPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineIdentityPoolsTPGetMachineIdentityPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**MachineIdentityPoolResponseModelCollection**](MachineIdentityPoolResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MachineIdentityPoolsTPGetProvSchemeReferences

> ProvisioningSchemeResponseModelCollection MachineIdentityPoolsTPGetProvSchemeReferences(ctx, identityPoolId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

List all provschemes that reference the given identitypool.

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
    identityPoolId := "identityPoolId_example" // string | Identity pool id.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineIdentityPoolsTPApi.MachineIdentityPoolsTPGetProvSchemeReferences(context.Background(), identityPoolId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MachineIdentityPoolsTPApi.MachineIdentityPoolsTPGetProvSchemeReferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MachineIdentityPoolsTPGetProvSchemeReferences`: ProvisioningSchemeResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `MachineIdentityPoolsTPApi.MachineIdentityPoolsTPGetProvSchemeReferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityPoolId** | **string** | Identity pool id. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMachineIdentityPoolsTPGetProvSchemeReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**ProvisioningSchemeResponseModelCollection**](ProvisioningSchemeResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

