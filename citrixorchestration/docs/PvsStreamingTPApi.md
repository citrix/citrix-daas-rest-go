# \PvsStreamingTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PvsStreamingTPGetPvsStreamingSites**](PvsStreamingTPApi.md#PvsStreamingTPGetPvsStreamingSites) | **Get** /techpreview/{customerid}/{siteid}/PvsStreaming/Sites | Get the list of PVS sites.
[**PvsStreamingTPGetPvsStreamingStores**](PvsStreamingTPApi.md#PvsStreamingTPGetPvsStreamingStores) | **Get** /techpreview/{customerid}/{siteid}/PvsStreaming/Stores/{farmId} | Get the list of PVS stores.
[**PvsStreamingTPGetPvsStreamingVDisks**](PvsStreamingTPApi.md#PvsStreamingTPGetPvsStreamingVDisks) | **Get** /techpreview/{customerid}/{siteid}/PvsStreaming/vDisks/{farmId} | Get the list of PVS vDisks.



## PvsStreamingTPGetPvsStreamingSites

> PvsStreamingSiteResponseModelCollection PvsStreamingTPGetPvsStreamingSites(ctx, customerid, siteid).FarmId(farmId).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the list of PVS sites.

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
    farmId := "farmId_example" // string | PVS farm id. (optional)
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PvsStreamingTPApi.PvsStreamingTPGetPvsStreamingSites(context.Background(), customerid, siteid).FarmId(farmId).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PvsStreamingTPApi.PvsStreamingTPGetPvsStreamingSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PvsStreamingTPGetPvsStreamingSites`: PvsStreamingSiteResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `PvsStreamingTPApi.PvsStreamingTPGetPvsStreamingSites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPvsStreamingTPGetPvsStreamingSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **farmId** | **string** | PVS farm id. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**PvsStreamingSiteResponseModelCollection**](PvsStreamingSiteResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PvsStreamingTPGetPvsStreamingStores

> PvsStreamingStoreResponseModelCollection PvsStreamingTPGetPvsStreamingStores(ctx, farmId, customerid, siteid).PvsSiteId(pvsSiteId).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the list of PVS stores.

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
    farmId := "farmId_example" // string | PVS farm id.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    pvsSiteId := "pvsSiteId_example" // string | PVS site id. (optional)
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PvsStreamingTPApi.PvsStreamingTPGetPvsStreamingStores(context.Background(), farmId, customerid, siteid).PvsSiteId(pvsSiteId).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PvsStreamingTPApi.PvsStreamingTPGetPvsStreamingStores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PvsStreamingTPGetPvsStreamingStores`: PvsStreamingStoreResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `PvsStreamingTPApi.PvsStreamingTPGetPvsStreamingStores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**farmId** | **string** | PVS farm id. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPvsStreamingTPGetPvsStreamingStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pvsSiteId** | **string** | PVS site id. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**PvsStreamingStoreResponseModelCollection**](PvsStreamingStoreResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PvsStreamingTPGetPvsStreamingVDisks

> PvsStreamingVDiskResponseModelCollection PvsStreamingTPGetPvsStreamingVDisks(ctx, farmId, customerid, siteid).PvsSiteId(pvsSiteId).StoreId(storeId).DiskLocatorId(diskLocatorId).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the list of PVS vDisks.

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
    farmId := "farmId_example" // string | PVS farm id.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    pvsSiteId := "pvsSiteId_example" // string | PVS site id. (optional)
    storeId := "storeId_example" // string | PVS store id. (optional)
    diskLocatorId := "diskLocatorId_example" // string | PVS vDisk id. (optional)
    async := true // bool | If `true`, it will be queried as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PvsStreamingTPApi.PvsStreamingTPGetPvsStreamingVDisks(context.Background(), farmId, customerid, siteid).PvsSiteId(pvsSiteId).StoreId(storeId).DiskLocatorId(diskLocatorId).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PvsStreamingTPApi.PvsStreamingTPGetPvsStreamingVDisks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PvsStreamingTPGetPvsStreamingVDisks`: PvsStreamingVDiskResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `PvsStreamingTPApi.PvsStreamingTPGetPvsStreamingVDisks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**farmId** | **string** | PVS farm id. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPvsStreamingTPGetPvsStreamingVDisksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pvsSiteId** | **string** | PVS site id. | 
 **storeId** | **string** | PVS store id. | 
 **diskLocatorId** | **string** | PVS vDisk id. | 
 **async** | **bool** | If &#x60;true&#x60;, it will be queried as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**PvsStreamingVDiskResponseModelCollection**](PvsStreamingVDiskResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

