# \SystemOptimizationDAAS

All URIs are relative to *https://api.wem.cloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemUtilitySettingQuery**](SystemOptimizationDAAS.md#SystemUtilitySettingQuery) | **Get** /services/wem/systemOptimizations | Query System Optimization Configurations
[**SystemUtilitySettingQueryById**](SystemOptimizationDAAS.md#SystemUtilitySettingQueryById) | **Get** /services/wem/systemOptimizations/{id} | Query System Optimization Configurations by Identity
[**SystemUtilitySettingUpdate**](SystemOptimizationDAAS.md#SystemUtilitySettingUpdate) | **Put** /services/wem/systemOptimizations | Update a System Optimization Configuration



## SystemUtilitySettingQuery

> SystemUtilitySettingQuery200Response SystemUtilitySettingQuery(ctx).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Query System Optimization Configurations

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/devicemanagement"
)

func main() {
    accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
    authorization := "authorization_example" // string | Citrix Cloud authorization token.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud customer ID.
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional) (default to "ca8d8637-e47e-4f6c-9a34-d8399028eeb4")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemOptimizationDAAS.SystemUtilitySettingQuery(context.Background()).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemOptimizationDAAS.SystemUtilitySettingQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SystemUtilitySettingQuery`: SystemUtilitySettingQuery200Response
    fmt.Fprintf(os.Stdout, "Response from `SystemOptimizationDAAS.SystemUtilitySettingQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSystemUtilitySettingQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Citrix Cloud authorization token. | 
 **citrixCustomerId** | **string** | Citrix Cloud customer ID. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | [default to &quot;ca8d8637-e47e-4f6c-9a34-d8399028eeb4&quot;]

### Return type

[**SystemUtilitySettingQuery200Response**](SystemUtilitySettingQuery200Response.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8, items

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemUtilitySettingQueryById

> SystemOptimizationModel SystemUtilitySettingQueryById(ctx, id).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()

Query System Optimization Configurations by Identity

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/devicemanagement"
)

func main() {
    id := int64(789) // int64 | The ID of the resource you want to query.
    accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
    authorization := "authorization_example" // string | Citrix Cloud authorization token.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud customer ID.
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional) (default to "1456d989-9136-43e3-a8f6-43dc92dd7beb")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemOptimizationDAAS.SystemUtilitySettingQueryById(context.Background(), id).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemOptimizationDAAS.SystemUtilitySettingQueryById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SystemUtilitySettingQueryById`: SystemOptimizationModel
    fmt.Fprintf(os.Stdout, "Response from `SystemOptimizationDAAS.SystemUtilitySettingQueryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource you want to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSystemUtilitySettingQueryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Citrix Cloud authorization token. | 
 **citrixCustomerId** | **string** | Citrix Cloud customer ID. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | [default to &quot;1456d989-9136-43e3-a8f6-43dc92dd7beb&quot;]

### Return type

[**SystemOptimizationModel**](SystemOptimizationModel.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8, id, enableCpuSpikeProtection, enableCpuAutoProtection, cpuUsageLimitOfSpikeProtection, cpuUsageLimitSampleTimeOfSpikeProtection, idlePriorityTimeOfSpikeProtection, enableLimitCpuCoreUsage, cpuCoreLimitOfSpikeProtection, enableIntelligentCpuOptimization, enableIntelligentIoOptimization, excludeProcessesFromCpuSpikeProtection, processesExcludedFromCpuSpikeProtection, enableProcessCpuPriority, processesOfCpuPriority, enableProcessAffinity, processesOfCpuAffinity, enableProcessClamping, processesOfCpuClamping, enableMemoryWorkingSetOptimization, idleSampleTimeOfMemoryWorkingSetOptimization, idleStateLimitOfMemoryWorkingSetOptimization, excludeProcessesFromMemoryWorkingSetOptimization, processesExcludedFromMemoryWorkingSetOptimization, enableProcessIoPriority, processesOfIoPriority, enableFastLogoff, excludeGroupsFromFastLogoff, groupsExcludedFromFastLogoff

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemUtilitySettingUpdate

> SystemUtilitySettingUpdate(ctx).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Update a System Optimization Configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/citrix/citrix-daas-rest-go/devicemanagement"
)

func main() {
    accept := "accept_example" // string | Only \"application/json\" is supported. (default to "application/json")
    authorization := "authorization_example" // string | Citrix Cloud authorization token.
    citrixCustomerId := "citrixCustomerId_example" // string | Citrix Cloud customer ID.
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. (optional) (default to "d9fb0fe4-195a-4975-83c3-3fd6253aadd4")
    body := SystemOptimizationModel(987) // SystemOptimizationModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SystemOptimizationDAAS.SystemUtilitySettingUpdate(context.Background()).Accept(accept).Authorization(authorization).CitrixCustomerId(citrixCustomerId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemOptimizationDAAS.SystemUtilitySettingUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSystemUtilitySettingUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** | Only \&quot;application/json\&quot; is supported. | [default to &quot;application/json&quot;]
 **authorization** | **string** | Citrix Cloud authorization token. | 
 **citrixCustomerId** | **string** | Citrix Cloud customer ID. | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. | [default to &quot;d9fb0fe4-195a-4975-83c3-3fd6253aadd4&quot;]
 **body** | **SystemOptimizationModel** |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

