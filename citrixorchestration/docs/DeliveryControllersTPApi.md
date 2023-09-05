# \DeliveryControllersTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeliveryControllersTPDeleteDeliveryController**](DeliveryControllersTPApi.md#DeliveryControllersTPDeleteDeliveryController) | **Delete** /techpreview/{customerid}/{siteid}/DeliveryControllers/{nameOrId} | Delete a delivery controller
[**DeliveryControllersTPGenerateDeliveryControllerDeleteScript**](DeliveryControllersTPApi.md#DeliveryControllersTPGenerateDeliveryControllerDeleteScript) | **Post** /techpreview/{customerid}/{siteid}/DeliveryControllers/{nameOrId}/$generateDeliveryControllerDeleteScript | Get the sql script when deleting a delivery controller.
[**DeliveryControllersTPGetDeliveryController**](DeliveryControllersTPApi.md#DeliveryControllersTPGetDeliveryController) | **Get** /techpreview/{customerid}/{siteid}/DeliveryControllers/{nameOrId} | Get the details about a single delivery controller.
[**DeliveryControllersTPGetDeliveryControllers**](DeliveryControllersTPApi.md#DeliveryControllersTPGetDeliveryControllers) | **Get** /techpreview/{customerid}/{siteid}/DeliveryControllers | Get the list of delivery controllers that are available to the customer and visible to the admin.



## DeliveryControllersTPDeleteDeliveryController

> bool DeliveryControllersTPDeleteDeliveryController(ctx, nameOrId, customerid, siteid).UpdateDatabase(updateDatabase).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()

Delete a delivery controller

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery controller.
    updateDatabase := true // bool | Indicates whether the database should be updated when deleting controller.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xAdminCredential := "xAdminCredential_example" // string | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard 'Authorization' request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \"Basic \" scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryControllersTPApi.DeliveryControllersTPDeleteDeliveryController(context.Background(), nameOrId, customerid, siteid).UpdateDatabase(updateDatabase).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XAdminCredential(xAdminCredential).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryControllersTPApi.DeliveryControllersTPDeleteDeliveryController``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryControllersTPDeleteDeliveryController`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeliveryControllersTPApi.DeliveryControllersTPDeleteDeliveryController`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery controller. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryControllersTPDeleteDeliveryControllerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDatabase** | **bool** | Indicates whether the database should be updated when deleting controller. | 


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xAdminCredential** | **string** | Admin credential to use when performing the operation. If not specified, the operation will be performed using the account under which the identity service is running If specified, must be in a format matching that of the standard &#39;Authorization&#39; request header; the username and password separated by a colon, UTF8-encoded, then base64-encoded, then the \&quot;Basic \&quot; scheme prepended.  Example:Basic QWxhZGRpbjpPcGVuU2VzYW1l | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

**bool**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryControllersTPGenerateDeliveryControllerDeleteScript

> DeliveryControllerDeleteScriptModelCollection DeliveryControllersTPGenerateDeliveryControllerDeleteScript(ctx, nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the sql script when deleting a delivery controller.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery controller.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the delivery controller delete sql script will be generated in a background task. The task will have JobType GenerateDeliveryControllerDeleteScript. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryControllersTPApi.DeliveryControllersTPGenerateDeliveryControllerDeleteScript(context.Background(), nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryControllersTPApi.DeliveryControllersTPGenerateDeliveryControllerDeleteScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryControllersTPGenerateDeliveryControllerDeleteScript`: DeliveryControllerDeleteScriptModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryControllersTPApi.DeliveryControllersTPGenerateDeliveryControllerDeleteScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery controller. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryControllersTPGenerateDeliveryControllerDeleteScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, the delivery controller delete sql script will be generated in a background task. The task will have JobType GenerateDeliveryControllerDeleteScript. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**DeliveryControllerDeleteScriptModelCollection**](DeliveryControllerDeleteScriptModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryControllersTPGetDeliveryController

> DeliveryControllerResponseModel DeliveryControllersTPGetDeliveryController(ctx, nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the details about a single delivery controller.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the delivery controller.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryControllersTPApi.DeliveryControllersTPGetDeliveryController(context.Background(), nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryControllersTPApi.DeliveryControllersTPGetDeliveryController``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryControllersTPGetDeliveryController`: DeliveryControllerResponseModel
    fmt.Fprintf(os.Stdout, "Response from `DeliveryControllersTPApi.DeliveryControllersTPGetDeliveryController`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the delivery controller. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryControllersTPGetDeliveryControllerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**DeliveryControllerResponseModel**](DeliveryControllerResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliveryControllersTPGetDeliveryControllers

> DeliveryControllerResponseModelCollection DeliveryControllersTPGetDeliveryControllers(ctx, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the list of delivery controllers that are available to the customer and visible to the admin.

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
    resp, r, err := apiClient.DeliveryControllersTPApi.DeliveryControllersTPGetDeliveryControllers(context.Background(), customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryControllersTPApi.DeliveryControllersTPGetDeliveryControllers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliveryControllersTPGetDeliveryControllers`: DeliveryControllerResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `DeliveryControllersTPApi.DeliveryControllersTPGetDeliveryControllers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliveryControllersTPGetDeliveryControllersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**DeliveryControllerResponseModelCollection**](DeliveryControllerResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

