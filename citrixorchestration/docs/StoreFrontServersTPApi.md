# \StoreFrontServersTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StoreFrontServersTPCreateStoreFrontServer**](StoreFrontServersTPApi.md#StoreFrontServersTPCreateStoreFrontServer) | **Post** /techpreview/{customerid}/{siteid}/StoreFrontServers | Create a StoreFront server.
[**StoreFrontServersTPDeleteStoreFrontServer**](StoreFrontServersTPApi.md#StoreFrontServersTPDeleteStoreFrontServer) | **Delete** /techpreview/{customerid}/{siteid}/StoreFrontServers/{nameOrId} | Delete a StoreFront server from the site.
[**StoreFrontServersTPGetStoreFrontAdministrators**](StoreFrontServersTPApi.md#StoreFrontServersTPGetStoreFrontAdministrators) | **Get** /techpreview/{customerid}/{siteid}/StoreFrontServers/{nameOrId}/Administrators | Get administrators who can administer a StoreFront server.
[**StoreFrontServersTPGetStoreFrontDeliveryGroups**](StoreFrontServersTPApi.md#StoreFrontServersTPGetStoreFrontDeliveryGroups) | **Get** /techpreview/{customerid}/{siteid}/StoreFrontServers/{nameOrId}/DeliveryGroups | GET delivery groups details for a Storefront
[**StoreFrontServersTPGetStoreFrontServer**](StoreFrontServersTPApi.md#StoreFrontServersTPGetStoreFrontServer) | **Get** /techpreview/{customerid}/{siteid}/StoreFrontServers/{nameOrId} | Get the details for a single StoreFront server.
[**StoreFrontServersTPGetStoreFrontServers**](StoreFrontServersTPApi.md#StoreFrontServersTPGetStoreFrontServers) | **Get** /techpreview/{customerid}/{siteid}/StoreFrontServers | Get all StoreFront servers.
[**StoreFrontServersTPUpdateStoreFrontServer**](StoreFrontServersTPApi.md#StoreFrontServersTPUpdateStoreFrontServer) | **Patch** /techpreview/{customerid}/{siteid}/StoreFrontServers/{nameOrId} | Update a StoreFront server.



## StoreFrontServersTPCreateStoreFrontServer

> StoreFrontServerResponseModel StoreFrontServersTPCreateStoreFrontServer(ctx, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Create a StoreFront server.



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
    request := *openapiclient.NewStoreFrontServerRequestModel() // StoreFrontServerRequestModel | Details about the StoreFront server to create.  Note: the Id property must not be specified in the request.
    async := true // bool | If `true`, the StoreFront server (and associated objects) will be created as a background task.  The task will have JobType CreateStoreFrontServer.  When the task is complete it will redirect to GetStoreFrontServer.  The job's Parameters will contain properties: * _Name_ - Name of the StoreFront server being created. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StoreFrontServersTPApi.StoreFrontServersTPCreateStoreFrontServer(context.Background(), customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StoreFrontServersTPApi.StoreFrontServersTPCreateStoreFrontServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreFrontServersTPCreateStoreFrontServer`: StoreFrontServerResponseModel
    fmt.Fprintf(os.Stdout, "Response from `StoreFrontServersTPApi.StoreFrontServersTPCreateStoreFrontServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoreFrontServersTPCreateStoreFrontServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**StoreFrontServerRequestModel**](StoreFrontServerRequestModel.md) | Details about the StoreFront server to create.  Note: the Id property must not be specified in the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the StoreFront server (and associated objects) will be created as a background task.  The task will have JobType CreateStoreFrontServer.  When the task is complete it will redirect to GetStoreFrontServer.  The job&#39;s Parameters will contain properties: * _Name_ - Name of the StoreFront server being created. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**StoreFrontServerResponseModel**](StoreFrontServerResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreFrontServersTPDeleteStoreFrontServer

> StoreFrontServersTPDeleteStoreFrontServer(ctx, nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Delete a StoreFront server from the site.



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
    nameOrId := "nameOrId_example" // string | The name or ID of StoreFront server to be deleted.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the StoreFront server (and associated objects) will be deleted as a background task.  The task will have JobType DeleteStoreFrontServer.  When the task is complete it will redirect to GetStoreFrontServers.  The job's Parameters will contain properties: * _Name_ - Name of the StoreFront server being deleted. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StoreFrontServersTPApi.StoreFrontServersTPDeleteStoreFrontServer(context.Background(), nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StoreFrontServersTPApi.StoreFrontServersTPDeleteStoreFrontServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | The name or ID of StoreFront server to be deleted. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoreFrontServersTPDeleteStoreFrontServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, the StoreFront server (and associated objects) will be deleted as a background task.  The task will have JobType DeleteStoreFrontServer.  When the task is complete it will redirect to GetStoreFrontServers.  The job&#39;s Parameters will contain properties: * _Name_ - Name of the StoreFront server being deleted. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreFrontServersTPGetStoreFrontAdministrators

> AdministratorResponseModelCollection StoreFrontServersTPGetStoreFrontAdministrators(ctx, nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get administrators who can administer a StoreFront server.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the StoreFront server.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StoreFrontServersTPApi.StoreFrontServersTPGetStoreFrontAdministrators(context.Background(), nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StoreFrontServersTPApi.StoreFrontServersTPGetStoreFrontAdministrators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreFrontServersTPGetStoreFrontAdministrators`: AdministratorResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `StoreFrontServersTPApi.StoreFrontServersTPGetStoreFrontAdministrators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the StoreFront server. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoreFrontServersTPGetStoreFrontAdministratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**AdministratorResponseModelCollection**](AdministratorResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreFrontServersTPGetStoreFrontDeliveryGroups

> StoreFrontDeliveryGroupResponseModelCollection StoreFrontServersTPGetStoreFrontDeliveryGroups(ctx, nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

GET delivery groups details for a Storefront



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
    nameOrId := "nameOrId_example" // string | The id of the Storefront
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StoreFrontServersTPApi.StoreFrontServersTPGetStoreFrontDeliveryGroups(context.Background(), nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StoreFrontServersTPApi.StoreFrontServersTPGetStoreFrontDeliveryGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreFrontServersTPGetStoreFrontDeliveryGroups`: StoreFrontDeliveryGroupResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `StoreFrontServersTPApi.StoreFrontServersTPGetStoreFrontDeliveryGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | The id of the Storefront | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoreFrontServersTPGetStoreFrontDeliveryGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**StoreFrontDeliveryGroupResponseModelCollection**](StoreFrontDeliveryGroupResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreFrontServersTPGetStoreFrontServer

> StoreFrontServerResponseModel StoreFrontServersTPGetStoreFrontServer(ctx, nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the details for a single StoreFront server.



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
    nameOrId := "nameOrId_example" // string | The name or ID of the StoreFront server.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StoreFrontServersTPApi.StoreFrontServersTPGetStoreFrontServer(context.Background(), nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StoreFrontServersTPApi.StoreFrontServersTPGetStoreFrontServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreFrontServersTPGetStoreFrontServer`: StoreFrontServerResponseModel
    fmt.Fprintf(os.Stdout, "Response from `StoreFrontServersTPApi.StoreFrontServersTPGetStoreFrontServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | The name or ID of the StoreFront server. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoreFrontServersTPGetStoreFrontServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**StoreFrontServerResponseModel**](StoreFrontServerResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreFrontServersTPGetStoreFrontServers

> StoreFrontServerResponseModelCollection StoreFrontServersTPGetStoreFrontServers(ctx, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get all StoreFront servers.



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
    resp, r, err := apiClient.StoreFrontServersTPApi.StoreFrontServersTPGetStoreFrontServers(context.Background(), customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StoreFrontServersTPApi.StoreFrontServersTPGetStoreFrontServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreFrontServersTPGetStoreFrontServers`: StoreFrontServerResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `StoreFrontServersTPApi.StoreFrontServersTPGetStoreFrontServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoreFrontServersTPGetStoreFrontServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**StoreFrontServerResponseModelCollection**](StoreFrontServerResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreFrontServersTPUpdateStoreFrontServer

> StoreFrontServerResponseModel StoreFrontServersTPUpdateStoreFrontServer(ctx, nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update a StoreFront server.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the StoreFront server to update.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewStoreFrontServerRequestModel() // StoreFrontServerRequestModel | Properties of the StoreFront server to update.  Note: the Id property must not be specified in the request.
    async := true // bool | If `true`, the StoreFront server (and associated objects) will be updated as a background task.  The task will have JobType UpdateStoreFrontServer.  When the task is complete it will redirect to GetStoreFrontServer.  The job's Parameters will contain properties: * _Name_ - Name of the StoreFront server being updated. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StoreFrontServersTPApi.StoreFrontServersTPUpdateStoreFrontServer(context.Background(), nameOrId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StoreFrontServersTPApi.StoreFrontServersTPUpdateStoreFrontServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreFrontServersTPUpdateStoreFrontServer`: StoreFrontServerResponseModel
    fmt.Fprintf(os.Stdout, "Response from `StoreFrontServersTPApi.StoreFrontServersTPUpdateStoreFrontServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the StoreFront server to update. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoreFrontServersTPUpdateStoreFrontServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**StoreFrontServerRequestModel**](StoreFrontServerRequestModel.md) | Properties of the StoreFront server to update.  Note: the Id property must not be specified in the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the StoreFront server (and associated objects) will be updated as a background task.  The task will have JobType UpdateStoreFrontServer.  When the task is complete it will redirect to GetStoreFrontServer.  The job&#39;s Parameters will contain properties: * _Name_ - Name of the StoreFront server being updated. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**StoreFrontServerResponseModel**](StoreFrontServerResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

