# \HypervisorsTPApi

All URIs are relative to */citrix/orchestration/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HypervisorsTPCheckHypervisorExists**](HypervisorsTPApi.md#HypervisorsTPCheckHypervisorExists) | **Head** /techpreview/{customerid}/{siteid}/hypervisors/{name} | Check for the existence of a hypervisor by name.
[**HypervisorsTPCheckHypervisorResourcePoolExists**](HypervisorsTPApi.md#HypervisorsTPCheckHypervisorResourcePoolExists) | **Head** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId}/resourcePools/{name} | Check for the existence of a resource pool by name.
[**HypervisorsTPCreateHypervisor**](HypervisorsTPApi.md#HypervisorsTPCreateHypervisor) | **Post** /techpreview/{customerid}/{siteid}/hypervisors | Create a hypervisor, and optionally a resource pool on the hypervisor.
[**HypervisorsTPCreateResourcePool**](HypervisorsTPApi.md#HypervisorsTPCreateResourcePool) | **Post** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId}/resourcePools | Create a new resource pool on an existing hypervisor.
[**HypervisorsTPDeleteHypervisor**](HypervisorsTPApi.md#HypervisorsTPDeleteHypervisor) | **Delete** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId} | Delete a hypervisor and related resource pools.
[**HypervisorsTPDeleteHypervisorResourcePool**](HypervisorsTPApi.md#HypervisorsTPDeleteHypervisorResourcePool) | **Delete** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId}/resourcePools/{poolId} | Delete a hypervisor resource pool.
[**HypervisorsTPDoHypervisorResourceSearch**](HypervisorsTPApi.md#HypervisorsTPDoHypervisorResourceSearch) | **Post** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId}/resourcePools/{poolId}/resources/$search | Search the resources within a hypervisor resource pool.
[**HypervisorsTPGetHypervisor**](HypervisorsTPApi.md#HypervisorsTPGetHypervisor) | **Get** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId} | Get the details for a single hypervisor.
[**HypervisorsTPGetHypervisorAdministrators**](HypervisorsTPApi.md#HypervisorsTPGetHypervisorAdministrators) | **Get** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId}/administrators | Get administrators who can administer a hypervisor.
[**HypervisorsTPGetHypervisorAllResources**](HypervisorsTPApi.md#HypervisorsTPGetHypervisorAllResources) | **Get** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId}/allResources | Get all resources within a hypervisor.
[**HypervisorsTPGetHypervisorAllResourcesWithoutConnection**](HypervisorsTPApi.md#HypervisorsTPGetHypervisorAllResourcesWithoutConnection) | **Post** /techpreview/{customerid}/{siteid}/hypervisors/$GetResources | Get all resources within a hypervisor, _without_ creating a persistent connection to the hypervisor.
[**HypervisorsTPGetHypervisorDeletePreview**](HypervisorsTPApi.md#HypervisorsTPGetHypervisorDeletePreview) | **Get** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId}/deletePreview | Get the hypervisor delete preview.
[**HypervisorsTPGetHypervisorJobs**](HypervisorsTPApi.md#HypervisorsTPGetHypervisorJobs) | **Get** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId}/jobs | Get the currently active jobs that are using a hypervisor.
[**HypervisorsTPGetHypervisorMachineCatalogs**](HypervisorsTPApi.md#HypervisorsTPGetHypervisorMachineCatalogs) | **Get** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId}/machineCatalogs | Get the machine catalogs that are using a hypervisor.
[**HypervisorsTPGetHypervisorResourcePool**](HypervisorsTPApi.md#HypervisorsTPGetHypervisorResourcePool) | **Get** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId}/resourcePools/{poolId} | Get details about a hypervisor resource pool.
[**HypervisorsTPGetHypervisorResourcePoolAdministrators**](HypervisorsTPApi.md#HypervisorsTPGetHypervisorResourcePoolAdministrators) | **Get** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId}/resourcePools/{poolId}/administrators | Get administrators who can administer a resource pool.
[**HypervisorsTPGetHypervisorResourcePoolDeletePreview**](HypervisorsTPApi.md#HypervisorsTPGetHypervisorResourcePoolDeletePreview) | **Get** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId}/resourcePools/{poolId}/deletePreview | Get the hypervisor resource pool delete preview.
[**HypervisorsTPGetHypervisorResourcePoolJobs**](HypervisorsTPApi.md#HypervisorsTPGetHypervisorResourcePoolJobs) | **Get** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId}/resourcePools/{poolId}/jobs | Get the currently active jobs that are using a resource pool.
[**HypervisorsTPGetHypervisorResourcePoolMachineCatalogs**](HypervisorsTPApi.md#HypervisorsTPGetHypervisorResourcePoolMachineCatalogs) | **Get** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId}/resourcePools/{poolId}/machineCatalogs | Get the machine catalogs that are using a resource pool.
[**HypervisorsTPGetHypervisorResourcePoolPatchPreview**](HypervisorsTPApi.md#HypervisorsTPGetHypervisorResourcePoolPatchPreview) | **Post** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId}/resourcePools/{poolId}/$getPatchPreview | Get the hypervisor resource pool patch preview.
[**HypervisorsTPGetHypervisorResourcePoolResources**](HypervisorsTPApi.md#HypervisorsTPGetHypervisorResourcePoolResources) | **Get** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId}/resourcePools/{poolId}/resources | Get the resources within a hypervisor resource pool.
[**HypervisorsTPGetHypervisorResourcePoolTestReport**](HypervisorsTPApi.md#HypervisorsTPGetHypervisorResourcePoolTestReport) | **Get** /techpreview/{customerid}/{siteid}/Hypervisors/{nameOrId}/ResourcePools/{poolId}/TestReport | Get the most recent test report for a resource pool.
[**HypervisorsTPGetHypervisorResourcePools**](HypervisorsTPApi.md#HypervisorsTPGetHypervisorResourcePools) | **Get** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId}/resourcePools | Get the list of hypervisor resource pools.
[**HypervisorsTPGetHypervisorServerHAAddresses**](HypervisorsTPApi.md#HypervisorsTPGetHypervisorServerHAAddresses) | **Get** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId}/serverHAAddresses | Get hypervisor server HA addresses. Currently, it only valid for Citrix hypervisors.
[**HypervisorsTPGetHypervisorTestReport**](HypervisorsTPApi.md#HypervisorsTPGetHypervisorTestReport) | **Get** /techpreview/{customerid}/{siteid}/Hypervisors/{nameOrId}/TestReport | Get the most recent test report for a hypervisor.
[**HypervisorsTPGetHypervisorWithoutConnection**](HypervisorsTPApi.md#HypervisorsTPGetHypervisorWithoutConnection) | **Post** /techpreview/{customerid}/{siteid}/$GetHypervisor | Get a hypervisor details without creating persistent connection.
[**HypervisorsTPGetHypervisors**](HypervisorsTPApi.md#HypervisorsTPGetHypervisors) | **Get** /techpreview/{customerid}/{siteid}/hypervisors | Get the hypervisors.
[**HypervisorsTPGetHypervisorsAndResourcePools**](HypervisorsTPApi.md#HypervisorsTPGetHypervisorsAndResourcePools) | **Get** /techpreview/{customerid}/{siteid}/hypervisorsAndResourcePools | Get hypervisors and resource pools. This API is used for the hosting main view.
[**HypervisorsTPGetSupportHypervisors**](HypervisorsTPApi.md#HypervisorsTPGetSupportHypervisors) | **Get** /techpreview/{customerid}/{siteid}/hypervisors/supportedPlugins | Get current server support hypervisors.
[**HypervisorsTPPatchHypervisor**](HypervisorsTPApi.md#HypervisorsTPPatchHypervisor) | **Patch** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId} | Update a hypervisor.
[**HypervisorsTPPatchHypervisorResourcePool**](HypervisorsTPApi.md#HypervisorsTPPatchHypervisorResourcePool) | **Patch** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId}/resourcePools/{poolId} | Update a hypervisor resource pool.
[**HypervisorsTPPutBrokerHypervisorConnection**](HypervisorsTPApi.md#HypervisorsTPPutBrokerHypervisorConnection) | **Put** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId}/connection | Set the properties on a broker hypervisor connection.
[**HypervisorsTPResetBrokerHypervisorConnection**](HypervisorsTPApi.md#HypervisorsTPResetBrokerHypervisorConnection) | **Post** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId}/$resetConnection | Requests the broker hypervisor connection to be reset. The connection is dropped, details including credentials refreshed and the connection reestablished. The reset request is asynchronous and may take a moment to occur.
[**HypervisorsTPTestHypervisor**](HypervisorsTPApi.md#HypervisorsTPTestHypervisor) | **Post** /techpreview/{customerid}/{siteid}/Hypervisors/{nameOrId}/$test | Run tests on a hypervisor and create a test report.
[**HypervisorsTPTestHypervisorResourcePool**](HypervisorsTPApi.md#HypervisorsTPTestHypervisorResourcePool) | **Post** /techpreview/{customerid}/{siteid}/Hypervisors/{nameOrId}/ResourcePools/{poolId}/$test | Run tests on a resource pool and create a test report.
[**HypervisorsTPValidateHypervisorResourcePoolResource**](HypervisorsTPApi.md#HypervisorsTPValidateHypervisorResourcePoolResource) | **Post** /techpreview/{customerid}/{siteid}/hypervisors/{nameOrId}/resourcePools/{poolId}/resources/$validate | Validate a resource for certain usages, currently default and only usage is for Machine Profile.
[**HypervisorsTPValidateVmPath**](HypervisorsTPApi.md#HypervisorsTPValidateVmPath) | **Post** /techpreview/{customerid}/{siteid}/hypervisors/$ValidateVm | Validate a list of hypervisor Vm Path.



## HypervisorsTPCheckHypervisorExists

> HypervisorsTPCheckHypervisorExists(ctx, name, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Check for the existence of a hypervisor by name.



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
    name := "name_example" // string | Name of the hypervisor.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HypervisorsTPApi.HypervisorsTPCheckHypervisorExists(context.Background(), name, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPCheckHypervisorExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the hypervisor. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPCheckHypervisorExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## HypervisorsTPCheckHypervisorResourcePoolExists

> HypervisorsTPCheckHypervisorResourcePoolExists(ctx, nameOrId, name, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Check for the existence of a resource pool by name.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    name := "name_example" // string | Name of the resource pool.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HypervisorsTPApi.HypervisorsTPCheckHypervisorResourcePoolExists(context.Background(), nameOrId, name, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPCheckHypervisorResourcePoolExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**name** | **string** | Name of the resource pool. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPCheckHypervisorResourcePoolExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## HypervisorsTPCreateHypervisor

> HypervisorDetailResponseModel HypervisorsTPCreateHypervisor(ctx, customerid, siteid).Model(model).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Create a hypervisor, and optionally a resource pool on the hypervisor.



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
    model := *openapiclient.NewCreateHypervisorRequestModel(*openapiclient.NewCreateHypervisorRequestModelConnectionDetails("Name_example", openapiclient.HypervisorConnectionType("Unknown"))) // CreateHypervisorRequestModel | Details of the hypervisor to create.
    async := true // bool | If `true`, the hypervisor (and associated objects) will be created as a background task. The task will have JobType CreateHypervisor. When the task is complete it will redirect to GetHypervisor. The job's Parameters will contain properties: * _Name_ - Name of the hypervisor being created. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPCreateHypervisor(context.Background(), customerid, siteid).Model(model).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPCreateHypervisor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPCreateHypervisor`: HypervisorDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPCreateHypervisor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPCreateHypervisorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**CreateHypervisorRequestModel**](CreateHypervisorRequestModel.md) | Details of the hypervisor to create. | 
 **async** | **bool** | If &#x60;true&#x60;, the hypervisor (and associated objects) will be created as a background task. The task will have JobType CreateHypervisor. When the task is complete it will redirect to GetHypervisor. The job&#39;s Parameters will contain properties: * _Name_ - Name of the hypervisor being created. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**HypervisorDetailResponseModel**](HypervisorDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPCreateResourcePool

> HypervisorResourcePoolDetailResponseModel HypervisorsTPCreateResourcePool(ctx, nameOrId, customerid, siteid).Model(model).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Create a new resource pool on an existing hypervisor.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    model := *openapiclient.NewCreateHypervisorResourcePoolRequestModel("Name_example", openapiclient.HypervisorConnectionType("Unknown")) // CreateHypervisorResourcePoolRequestModel | Details of the resource pool to create.
    async := true // bool | If `true`, the resource pool (and associated objects) will be created as a background task. The task will have JobType CreateHypervisorResourcePool. When the task is complete it will redirect to GetHypervisorResourcePool. The job's Parameters will contain properties: * _Name_ - Name of the resource pool being created. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPCreateResourcePool(context.Background(), nameOrId, customerid, siteid).Model(model).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPCreateResourcePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPCreateResourcePool`: HypervisorResourcePoolDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPCreateResourcePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPCreateResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **model** | [**CreateHypervisorResourcePoolRequestModel**](CreateHypervisorResourcePoolRequestModel.md) | Details of the resource pool to create. | 
 **async** | **bool** | If &#x60;true&#x60;, the resource pool (and associated objects) will be created as a background task. The task will have JobType CreateHypervisorResourcePool. When the task is complete it will redirect to GetHypervisorResourcePool. The job&#39;s Parameters will contain properties: * _Name_ - Name of the resource pool being created. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**HypervisorResourcePoolDetailResponseModel**](HypervisorResourcePoolDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPDeleteHypervisor

> HypervisorsTPDeleteHypervisor(ctx, nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Delete a hypervisor and related resource pools.

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
    nameOrId := "nameOrId_example" // string | The hypervisor connection name or id. 
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If the api call with async type.  (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HypervisorsTPApi.HypervisorsTPDeleteHypervisor(context.Background(), nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPDeleteHypervisor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | The hypervisor connection name or id.  | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPDeleteHypervisorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If the api call with async type.  | [default to false]
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


## HypervisorsTPDeleteHypervisorResourcePool

> HypervisorsTPDeleteHypervisorResourcePool(ctx, nameOrId, poolId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Delete a hypervisor resource pool.

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
    nameOrId := "nameOrId_example" // string | The name or id of hypervisor. 
    poolId := "poolId_example" // string | The hypervisor resource pool id. 
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If the api call with async type.  (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HypervisorsTPApi.HypervisorsTPDeleteHypervisorResourcePool(context.Background(), nameOrId, poolId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPDeleteHypervisorResourcePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | The name or id of hypervisor.  | 
**poolId** | **string** | The hypervisor resource pool id.  | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPDeleteHypervisorResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **async** | **bool** | If the api call with async type.  | [default to false]
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


## HypervisorsTPDoHypervisorResourceSearch

> HypervisorResourceResponseModelCollection HypervisorsTPDoHypervisorResourceSearch(ctx, nameOrId, poolId, customerid, siteid).SearchRequest(searchRequest).Limit(limit).ContinuationToken(continuationToken).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Search the resources within a hypervisor resource pool.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    poolId := "poolId_example" // string | Name or ID of the resource pool.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    searchRequest := *openapiclient.NewHypervisorResourceSearchRequestModel() // HypervisorResourceSearchRequestModel | Specifies the advanced search parameters.
    limit := int32(56) // int32 | The max number of resources returned by this query. If not specified, the server might use a default limit of 1 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    async := true // bool | Async request to get the resources with *path.  (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPDoHypervisorResourceSearch(context.Background(), nameOrId, poolId, customerid, siteid).SearchRequest(searchRequest).Limit(limit).ContinuationToken(continuationToken).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPDoHypervisorResourceSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPDoHypervisorResourceSearch`: HypervisorResourceResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPDoHypervisorResourceSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**poolId** | **string** | Name or ID of the resource pool. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPDoHypervisorResourceSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **searchRequest** | [**HypervisorResourceSearchRequestModel**](HypervisorResourceSearchRequestModel.md) | Specifies the advanced search parameters. | 
 **limit** | **int32** | The max number of resources returned by this query. If not specified, the server might use a default limit of 1 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **async** | **bool** | Async request to get the resources with *path.  | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**HypervisorResourceResponseModelCollection**](HypervisorResourceResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPGetHypervisor

> HypervisorDetailResponseModel HypervisorsTPGetHypervisor(ctx, nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the details for a single hypervisor.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | Async request to hypervisor.  (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPGetHypervisor(context.Background(), nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPGetHypervisor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPGetHypervisor`: HypervisorDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPGetHypervisor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPGetHypervisorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | Async request to hypervisor.  | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**HypervisorDetailResponseModel**](HypervisorDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPGetHypervisorAdministrators

> AdministratorResponseModelCollection HypervisorsTPGetHypervisorAdministrators(ctx, nameOrId, customerid, siteid).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get administrators who can administer a hypervisor.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If async execute.  (optional) (default to false)
    limit := int32(56) // int32 | The max number of administrators returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPGetHypervisorAdministrators(context.Background(), nameOrId, customerid, siteid).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPGetHypervisorAdministrators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPGetHypervisorAdministrators`: AdministratorResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPGetHypervisorAdministrators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPGetHypervisorAdministratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If async execute.  | [default to false]
 **limit** | **int32** | The max number of administrators returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
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


## HypervisorsTPGetHypervisorAllResources

> HypervisorResourceResponseModel HypervisorsTPGetHypervisorAllResources(ctx, nameOrId, customerid, siteid).Path(path).Children(children).Type_(type_).ShowTagged(showTagged).Detail(detail).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get all resources within a hypervisor.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    path := "path_example" // string | Path to the resource container within the hypervisor. The path may either be a relative path as specified by RelativePath, or may be a URL-encoded XenApp & XenDesktop resource path starting with `XDHyp:`; for example, as specified by XDPath. (optional)
    children := int32(56) // int32 | Specifies the number of levels of children to enumerate. Default is `0`, meaning that only the object referred to by `path` is returned and its Children array will be left null. A special value of `-1` indicates that the entire resource hierarchy should be enumerated.  Use with care!  It may take a very long time to enumerate a large number of resources from a hypervisor, and the call may time out before completing. (optional)
    type_ := []string{"Inner_example"} // []string | If specified, limits the results to the specified resource type(s). (optional)
    showTagged := true // bool | By default, items which are tagged by XenDesktop are not shown. Set this to `true` to override that behavior. (optional)
    detail := true // bool | If `true`, full details of VMs, snapshots, and templates will be retrieved. This can be very time consuming and will reduce the performance of the call. May only be used if `path` refers to a VM, snapshot, or template resource. <!-- Internally this calls Get-HypConfigurationObjectForItem --> (optional)
    async := true // bool | Async request to get the resources with *path. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPGetHypervisorAllResources(context.Background(), nameOrId, customerid, siteid).Path(path).Children(children).Type_(type_).ShowTagged(showTagged).Detail(detail).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPGetHypervisorAllResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPGetHypervisorAllResources`: HypervisorResourceResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPGetHypervisorAllResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPGetHypervisorAllResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **path** | **string** | Path to the resource container within the hypervisor. The path may either be a relative path as specified by RelativePath, or may be a URL-encoded XenApp &amp; XenDesktop resource path starting with &#x60;XDHyp:&#x60;; for example, as specified by XDPath. | 
 **children** | **int32** | Specifies the number of levels of children to enumerate. Default is &#x60;0&#x60;, meaning that only the object referred to by &#x60;path&#x60; is returned and its Children array will be left null. A special value of &#x60;-1&#x60; indicates that the entire resource hierarchy should be enumerated.  Use with care!  It may take a very long time to enumerate a large number of resources from a hypervisor, and the call may time out before completing. | 
 **type_** | **[]string** | If specified, limits the results to the specified resource type(s). | 
 **showTagged** | **bool** | By default, items which are tagged by XenDesktop are not shown. Set this to &#x60;true&#x60; to override that behavior. | 
 **detail** | **bool** | If &#x60;true&#x60;, full details of VMs, snapshots, and templates will be retrieved. This can be very time consuming and will reduce the performance of the call. May only be used if &#x60;path&#x60; refers to a VM, snapshot, or template resource. &lt;!-- Internally this calls Get-HypConfigurationObjectForItem --&gt; | 
 **async** | **bool** | Async request to get the resources with *path. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**HypervisorResourceResponseModel**](HypervisorResourceResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPGetHypervisorAllResourcesWithoutConnection

> HypervisorResourceResponseModel HypervisorsTPGetHypervisorAllResourcesWithoutConnection(ctx, customerid, siteid).ConnectionDetail(connectionDetail).Path(path).Children(children).Type_(type_).ShowTagged(showTagged).Detail(detail).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get all resources within a hypervisor, _without_ creating a persistent connection to the hypervisor.



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
    connectionDetail := *openapiclient.NewHypervisorConnectionDetailRequestModel("Name_example", openapiclient.HypervisorConnectionType("Unknown")) // HypervisorConnectionDetailRequestModel | Details of how to connect to the hypervisor to enumerate resources from it.
    path := "path_example" // string | Path to the resource container within the hypervisor. The path may either be a relative path as specified by RelativePath, or may be a URL-encoded Virtual Apps & Desktops resource path starting with `XDHyp:`; for example, as specified by XDPath. (optional)
    children := int32(56) // int32 | Specifies the number of levels of children to enumerate. Default is `0`, meaning that only the object referred to by `path` is returned and its Children array will be left null. A special value of `-1` indicates that the entire resource hierarchy should be enumerated.  Use with care!  It may take a very long time to enumerate a large number of resources from a hypervisor, and the call may time out before completing. (optional)
    type_ := []string{"Inner_example"} // []string | If specified, limits the results to the specified resource type(s). (optional)
    showTagged := true // bool | By default, items which are tagged by XenDesktop are not shown. Set this to `true` to override that behavior. (optional)
    detail := true // bool | If `true`, full details of VMs, snapshots, and templates will be retrieved. This can be very time consuming and will reduce the performance of the call. May only be used if `path` refers to a VM, snapshot, or template resource. <!-- Internally this calls Get-HypConfigurationObjectForItem --> (optional)
    async := true // bool | Async request to get the resources with *path. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPGetHypervisorAllResourcesWithoutConnection(context.Background(), customerid, siteid).ConnectionDetail(connectionDetail).Path(path).Children(children).Type_(type_).ShowTagged(showTagged).Detail(detail).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPGetHypervisorAllResourcesWithoutConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPGetHypervisorAllResourcesWithoutConnection`: HypervisorResourceResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPGetHypervisorAllResourcesWithoutConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPGetHypervisorAllResourcesWithoutConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **connectionDetail** | [**HypervisorConnectionDetailRequestModel**](HypervisorConnectionDetailRequestModel.md) | Details of how to connect to the hypervisor to enumerate resources from it. | 
 **path** | **string** | Path to the resource container within the hypervisor. The path may either be a relative path as specified by RelativePath, or may be a URL-encoded Virtual Apps &amp; Desktops resource path starting with &#x60;XDHyp:&#x60;; for example, as specified by XDPath. | 
 **children** | **int32** | Specifies the number of levels of children to enumerate. Default is &#x60;0&#x60;, meaning that only the object referred to by &#x60;path&#x60; is returned and its Children array will be left null. A special value of &#x60;-1&#x60; indicates that the entire resource hierarchy should be enumerated.  Use with care!  It may take a very long time to enumerate a large number of resources from a hypervisor, and the call may time out before completing. | 
 **type_** | **[]string** | If specified, limits the results to the specified resource type(s). | 
 **showTagged** | **bool** | By default, items which are tagged by XenDesktop are not shown. Set this to &#x60;true&#x60; to override that behavior. | 
 **detail** | **bool** | If &#x60;true&#x60;, full details of VMs, snapshots, and templates will be retrieved. This can be very time consuming and will reduce the performance of the call. May only be used if &#x60;path&#x60; refers to a VM, snapshot, or template resource. &lt;!-- Internally this calls Get-HypConfigurationObjectForItem --&gt; | 
 **async** | **bool** | Async request to get the resources with *path. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**HypervisorResourceResponseModel**](HypervisorResourceResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPGetHypervisorDeletePreview

> HypervisorDeletePreviewResponseModel HypervisorsTPGetHypervisorDeletePreview(ctx, nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the hypervisor delete preview.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If execute this API asynchronous.  (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPGetHypervisorDeletePreview(context.Background(), nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPGetHypervisorDeletePreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPGetHypervisorDeletePreview`: HypervisorDeletePreviewResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPGetHypervisorDeletePreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPGetHypervisorDeletePreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If execute this API asynchronous.  | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**HypervisorDeletePreviewResponseModel**](HypervisorDeletePreviewResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPGetHypervisorJobs

> JobResponseModelCollection HypervisorsTPGetHypervisorJobs(ctx, nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the currently active jobs that are using a hypervisor.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPGetHypervisorJobs(context.Background(), nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPGetHypervisorJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPGetHypervisorJobs`: JobResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPGetHypervisorJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPGetHypervisorJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**JobResponseModelCollection**](JobResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPGetHypervisorMachineCatalogs

> HypervisorMachineCatalogResponseModelCollection HypervisorsTPGetHypervisorMachineCatalogs(ctx, nameOrId, customerid, siteid).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the machine catalogs that are using a hypervisor.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If async execute.  (optional) (default to false)
    limit := int32(56) // int32 | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPGetHypervisorMachineCatalogs(context.Background(), nameOrId, customerid, siteid).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPGetHypervisorMachineCatalogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPGetHypervisorMachineCatalogs`: HypervisorMachineCatalogResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPGetHypervisorMachineCatalogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPGetHypervisorMachineCatalogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If async execute.  | [default to false]
 **limit** | **int32** | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**HypervisorMachineCatalogResponseModelCollection**](HypervisorMachineCatalogResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPGetHypervisorResourcePool

> HypervisorResourcePoolDetailResponseModel HypervisorsTPGetHypervisorResourcePool(ctx, nameOrId, poolId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get details about a hypervisor resource pool.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    poolId := "poolId_example" // string | Name or ID of the resource pool.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | Async request to get the resource pool.  (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePool(context.Background(), nameOrId, poolId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPGetHypervisorResourcePool`: HypervisorResourcePoolDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**poolId** | **string** | Name or ID of the resource pool. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPGetHypervisorResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **async** | **bool** | Async request to get the resource pool.  | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**HypervisorResourcePoolDetailResponseModel**](HypervisorResourcePoolDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPGetHypervisorResourcePoolAdministrators

> AdministratorResponseModelCollection HypervisorsTPGetHypervisorResourcePoolAdministrators(ctx, nameOrId, poolId, customerid, siteid).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get administrators who can administer a resource pool.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    poolId := "poolId_example" // string | Name or ID of the resource pool.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If async execute.  (optional) (default to false)
    limit := int32(56) // int32 | The max number of administrators returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePoolAdministrators(context.Background(), nameOrId, poolId, customerid, siteid).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePoolAdministrators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPGetHypervisorResourcePoolAdministrators`: AdministratorResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePoolAdministrators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**poolId** | **string** | Name or ID of the resource pool. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPGetHypervisorResourcePoolAdministratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **async** | **bool** | If async execute.  | [default to false]
 **limit** | **int32** | The max number of administrators returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
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


## HypervisorsTPGetHypervisorResourcePoolDeletePreview

> HypervisorDeletePreviewResponseModel HypervisorsTPGetHypervisorResourcePoolDeletePreview(ctx, nameOrId, poolId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the hypervisor resource pool delete preview.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    poolId := "poolId_example" // string | Name or ID of the resource pool.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If execute this API asynchronous.  (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePoolDeletePreview(context.Background(), nameOrId, poolId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePoolDeletePreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPGetHypervisorResourcePoolDeletePreview`: HypervisorDeletePreviewResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePoolDeletePreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**poolId** | **string** | Name or ID of the resource pool. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPGetHypervisorResourcePoolDeletePreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **async** | **bool** | If execute this API asynchronous.  | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**HypervisorDeletePreviewResponseModel**](HypervisorDeletePreviewResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPGetHypervisorResourcePoolJobs

> JobResponseModelCollection HypervisorsTPGetHypervisorResourcePoolJobs(ctx, nameOrId, poolId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the currently active jobs that are using a resource pool.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    poolId := "poolId_example" // string | Name or ID of the resource pool.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePoolJobs(context.Background(), nameOrId, poolId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePoolJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPGetHypervisorResourcePoolJobs`: JobResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePoolJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**poolId** | **string** | Name or ID of the resource pool. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPGetHypervisorResourcePoolJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**JobResponseModelCollection**](JobResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPGetHypervisorResourcePoolMachineCatalogs

> HypervisorMachineCatalogResponseModelCollection HypervisorsTPGetHypervisorResourcePoolMachineCatalogs(ctx, nameOrId, poolId, customerid, siteid).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the machine catalogs that are using a resource pool.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    poolId := "poolId_example" // string | Name or ID of the resource pool.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If async execute.  (optional) (default to false)
    limit := int32(56) // int32 | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePoolMachineCatalogs(context.Background(), nameOrId, poolId, customerid, siteid).Async(async).Limit(limit).ContinuationToken(continuationToken).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePoolMachineCatalogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPGetHypervisorResourcePoolMachineCatalogs`: HypervisorMachineCatalogResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePoolMachineCatalogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**poolId** | **string** | Name or ID of the resource pool. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPGetHypervisorResourcePoolMachineCatalogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **async** | **bool** | If async execute.  | [default to false]
 **limit** | **int32** | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**HypervisorMachineCatalogResponseModelCollection**](HypervisorMachineCatalogResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPGetHypervisorResourcePoolPatchPreview

> HypervisorResourcePoolPatchPreviewResponseModel HypervisorsTPGetHypervisorResourcePoolPatchPreview(ctx, nameOrId, poolId, customerid, siteid).Model(model).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the hypervisor resource pool patch preview.

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
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    poolId := "poolId_example" // string | Name or ID of the resource pool.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    model := *openapiclient.NewHypervisorResourcePoolPatchPreviewRequestModel() // HypervisorResourcePoolPatchPreviewRequestModel | New networks of the resource pool.
    async := true // bool | If `true`, the preview job will run as a background task. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePoolPatchPreview(context.Background(), nameOrId, poolId, customerid, siteid).Model(model).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePoolPatchPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPGetHypervisorResourcePoolPatchPreview`: HypervisorResourcePoolPatchPreviewResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePoolPatchPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**poolId** | **string** | Name or ID of the resource pool. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPGetHypervisorResourcePoolPatchPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **model** | [**HypervisorResourcePoolPatchPreviewRequestModel**](HypervisorResourcePoolPatchPreviewRequestModel.md) | New networks of the resource pool. | 
 **async** | **bool** | If &#x60;true&#x60;, the preview job will run as a background task. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**HypervisorResourcePoolPatchPreviewResponseModel**](HypervisorResourcePoolPatchPreviewResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPGetHypervisorResourcePoolResources

> HypervisorResourceResponseModel HypervisorsTPGetHypervisorResourcePoolResources(ctx, nameOrId, poolId, customerid, siteid).Path(path).Children(children).Type_(type_).ShowTagged(showTagged).Detail(detail).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the resources within a hypervisor resource pool.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    poolId := "poolId_example" // string | Name or ID of the resource pool.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    path := "path_example" // string | Path to the resource container within the resource pool. The path may either be a relative path as specified by RelativePath, or may be a URL-encoded XenApp & XenDesktop resource path starting with `XDHyp:`; for example, as specified by XDPath. (optional)
    children := int32(56) // int32 | Specifies the number of levels of children to enumerate. Default is `0`, meaning that only the object referred to by `path` is returned and its Children array will be left null. A special value of `-1` indicates that the entire resource hierarchy should be enumerated.  Use with care!  It may take a very long time to enumerate a large number of resources from a hypervisor, and the call may time out before completing. (optional)
    type_ := []string{"Inner_example"} // []string | If specified, limits the results to the specified resource type(s). (optional)
    showTagged := true // bool | By default, items which are tagged by XenDesktop are not shown. Set this to `true` to override that behavior. (optional)
    detail := true // bool | If `true`, full details of VMs, snapshots, and templates will be retrieved. This can be very time consuming and will reduce the performance of the call. May only be used if `path` refers to a VM, snapshot, or template resource. <!-- Internally this calls Get-HypConfigurationObjectForItem --> (optional)
    async := true // bool | Async request to get the resources with *path.  (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePoolResources(context.Background(), nameOrId, poolId, customerid, siteid).Path(path).Children(children).Type_(type_).ShowTagged(showTagged).Detail(detail).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePoolResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPGetHypervisorResourcePoolResources`: HypervisorResourceResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePoolResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**poolId** | **string** | Name or ID of the resource pool. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPGetHypervisorResourcePoolResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **path** | **string** | Path to the resource container within the resource pool. The path may either be a relative path as specified by RelativePath, or may be a URL-encoded XenApp &amp; XenDesktop resource path starting with &#x60;XDHyp:&#x60;; for example, as specified by XDPath. | 
 **children** | **int32** | Specifies the number of levels of children to enumerate. Default is &#x60;0&#x60;, meaning that only the object referred to by &#x60;path&#x60; is returned and its Children array will be left null. A special value of &#x60;-1&#x60; indicates that the entire resource hierarchy should be enumerated.  Use with care!  It may take a very long time to enumerate a large number of resources from a hypervisor, and the call may time out before completing. | 
 **type_** | **[]string** | If specified, limits the results to the specified resource type(s). | 
 **showTagged** | **bool** | By default, items which are tagged by XenDesktop are not shown. Set this to &#x60;true&#x60; to override that behavior. | 
 **detail** | **bool** | If &#x60;true&#x60;, full details of VMs, snapshots, and templates will be retrieved. This can be very time consuming and will reduce the performance of the call. May only be used if &#x60;path&#x60; refers to a VM, snapshot, or template resource. &lt;!-- Internally this calls Get-HypConfigurationObjectForItem --&gt; | 
 **async** | **bool** | Async request to get the resources with *path.  | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**HypervisorResourceResponseModel**](HypervisorResourceResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPGetHypervisorResourcePoolTestReport

> TestReportResponseModel HypervisorsTPGetHypervisorResourcePoolTestReport(ctx, nameOrId, poolId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the most recent test report for a resource pool.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    poolId := "poolId_example" // string | Name or ID of the resource pool.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePoolTestReport(context.Background(), nameOrId, poolId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePoolTestReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPGetHypervisorResourcePoolTestReport`: TestReportResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePoolTestReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**poolId** | **string** | Name or ID of the resource pool. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPGetHypervisorResourcePoolTestReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**TestReportResponseModel**](TestReportResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPGetHypervisorResourcePools

> HypervisorResourcePoolResponseModelCollection HypervisorsTPGetHypervisorResourcePools(ctx, nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the list of hypervisor resource pools.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | Async request to get the resource pool.  (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePools(context.Background(), nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPGetHypervisorResourcePools`: HypervisorResourcePoolResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPGetHypervisorResourcePools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPGetHypervisorResourcePoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | Async request to get the resource pool.  | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**HypervisorResourcePoolResponseModelCollection**](HypervisorResourcePoolResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPGetHypervisorServerHAAddresses

> HypervisorServerHAAddressesResponseModel HypervisorsTPGetHypervisorServerHAAddresses(ctx, nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get hypervisor server HA addresses. Currently, it only valid for Citrix hypervisors.

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
    nameOrId := "nameOrId_example" // string | The hypervisor connection name or id. 
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If the execution with async model.  (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPGetHypervisorServerHAAddresses(context.Background(), nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPGetHypervisorServerHAAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPGetHypervisorServerHAAddresses`: HypervisorServerHAAddressesResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPGetHypervisorServerHAAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | The hypervisor connection name or id.  | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPGetHypervisorServerHAAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If the execution with async model.  | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**HypervisorServerHAAddressesResponseModel**](HypervisorServerHAAddressesResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPGetHypervisorTestReport

> TestReportResponseModel HypervisorsTPGetHypervisorTestReport(ctx, nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the most recent test report for a hypervisor.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPGetHypervisorTestReport(context.Background(), nameOrId, customerid, siteid).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPGetHypervisorTestReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPGetHypervisorTestReport`: TestReportResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPGetHypervisorTestReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPGetHypervisorTestReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**TestReportResponseModel**](TestReportResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPGetHypervisorWithoutConnection

> HypervisorDetailResponseModel HypervisorsTPGetHypervisorWithoutConnection(ctx, customerid, siteid).ConnectionDetail(connectionDetail).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get a hypervisor details without creating persistent connection.

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
    connectionDetail := *openapiclient.NewHypervisorConnectionDetailRequestModel("Name_example", openapiclient.HypervisorConnectionType("Unknown")) // HypervisorConnectionDetailRequestModel | Details of how to connect to the hypervisor to enumerate resources from it.
    async := true // bool | Async request to get the hypervisor.  (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPGetHypervisorWithoutConnection(context.Background(), customerid, siteid).ConnectionDetail(connectionDetail).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPGetHypervisorWithoutConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPGetHypervisorWithoutConnection`: HypervisorDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPGetHypervisorWithoutConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPGetHypervisorWithoutConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **connectionDetail** | [**HypervisorConnectionDetailRequestModel**](HypervisorConnectionDetailRequestModel.md) | Details of how to connect to the hypervisor to enumerate resources from it. | 
 **async** | **bool** | Async request to get the hypervisor.  | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**HypervisorDetailResponseModel**](HypervisorDetailResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPGetHypervisors

> HypervisorResponseModelCollection HypervisorsTPGetHypervisors(ctx, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get the hypervisors.



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
    async := true // bool | Async request to hypervisor.  (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPGetHypervisors(context.Background(), customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPGetHypervisors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPGetHypervisors`: HypervisorResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPGetHypervisors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPGetHypervisorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **async** | **bool** | Async request to hypervisor.  | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**HypervisorResponseModelCollection**](HypervisorResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPGetHypervisorsAndResourcePools

> HypervisorsAndResourcePoolsResponseModelCollection HypervisorsTPGetHypervisorsAndResourcePools(ctx, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get hypervisors and resource pools. This API is used for the hosting main view.

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
    async := true // bool | If execute this API asynchronous.  (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPGetHypervisorsAndResourcePools(context.Background(), customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPGetHypervisorsAndResourcePools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPGetHypervisorsAndResourcePools`: HypervisorsAndResourcePoolsResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPGetHypervisorsAndResourcePools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPGetHypervisorsAndResourcePoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **async** | **bool** | If execute this API asynchronous.  | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**HypervisorsAndResourcePoolsResponseModelCollection**](HypervisorsAndResourcePoolsResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPGetSupportHypervisors

> HypervisorPluginResponseModelCollection HypervisorsTPGetSupportHypervisors(ctx, customerid, siteid).Async(async).ZoneNameOrId(zoneNameOrId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Get current server support hypervisors.

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
    async := true // bool | If execute this API asynchronous.  (optional) (default to false)
    zoneNameOrId := "zoneNameOrId_example" // string | The zone name or id.  (optional)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPGetSupportHypervisors(context.Background(), customerid, siteid).Async(async).ZoneNameOrId(zoneNameOrId).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPGetSupportHypervisors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPGetSupportHypervisors`: HypervisorPluginResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPGetSupportHypervisors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPGetSupportHypervisorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **async** | **bool** | If execute this API asynchronous.  | [default to false]
 **zoneNameOrId** | **string** | The zone name or id.  | 
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**HypervisorPluginResponseModelCollection**](HypervisorPluginResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPPatchHypervisor

> HypervisorsTPPatchHypervisor(ctx, nameOrId, customerid, siteid).Model(model).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update a hypervisor.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    model := *openapiclient.NewEditHypervisorConnectionRequestModel(openapiclient.HypervisorConnectionType("Unknown")) // EditHypervisorConnectionRequestModel | Details of the hypervisor to update.  Note that each type of hypervisor requires a different update model: * AWS requires a model of type EditHypervisorAWSConnectionRequestModel. * AzureRM requires a model of type EditHypervisorAzureConnectionRequestModel. * GoogleCloudPlatform requires a model of type EditHypervisorGCPConnectionRequestModel. * OracleCloudInfrastructure requires a model of type EditHypervisorOciConnectionRequestModel. * All other hypervisor types require a model of type EditHypervisorTraditionalConnectionRequestModel.
    async := true // bool | If `true`, the hypervisor will be updated as a background task. The task will have JobType UpdateHypervisor. When the task is complete it will redirect to GetHypervisor. The job's Parameters will contain properties: * _Name_ - Name of the hypervisor being updated. * _Id_ - ID of the hypervisor being updated. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HypervisorsTPApi.HypervisorsTPPatchHypervisor(context.Background(), nameOrId, customerid, siteid).Model(model).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPPatchHypervisor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPPatchHypervisorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **model** | [**EditHypervisorConnectionRequestModel**](EditHypervisorConnectionRequestModel.md) | Details of the hypervisor to update.  Note that each type of hypervisor requires a different update model: * AWS requires a model of type EditHypervisorAWSConnectionRequestModel. * AzureRM requires a model of type EditHypervisorAzureConnectionRequestModel. * GoogleCloudPlatform requires a model of type EditHypervisorGCPConnectionRequestModel. * OracleCloudInfrastructure requires a model of type EditHypervisorOciConnectionRequestModel. * All other hypervisor types require a model of type EditHypervisorTraditionalConnectionRequestModel. | 
 **async** | **bool** | If &#x60;true&#x60;, the hypervisor will be updated as a background task. The task will have JobType UpdateHypervisor. When the task is complete it will redirect to GetHypervisor. The job&#39;s Parameters will contain properties: * _Name_ - Name of the hypervisor being updated. * _Id_ - ID of the hypervisor being updated. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPPatchHypervisorResourcePool

> HypervisorsTPPatchHypervisorResourcePool(ctx, nameOrId, poolId, customerid, siteid).Model(model).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Update a hypervisor resource pool.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    poolId := "poolId_example" // string | Name or ID of the resource pool.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    model := *openapiclient.NewEditHypervisorResourcePoolRequestModel(openapiclient.HypervisorConnectionType("Unknown")) // EditHypervisorResourcePoolRequestModel | Details of the resource pool to update.
    async := true // bool | If `true`, the resource pool will be updated as a background task. The task will have JobType UpdateHypervisorResourcePool. When the task is complete it will redirect to GetHypervisorResourcePool. The job's Parameters will contain properties: * _Name_ - Name of the resource pool being updated. * _Id_ - ID of the resource pool being updated. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HypervisorsTPApi.HypervisorsTPPatchHypervisorResourcePool(context.Background(), nameOrId, poolId, customerid, siteid).Model(model).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPPatchHypervisorResourcePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**poolId** | **string** | Name or ID of the resource pool. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPPatchHypervisorResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **model** | [**EditHypervisorResourcePoolRequestModel**](EditHypervisorResourcePoolRequestModel.md) | Details of the resource pool to update. | 
 **async** | **bool** | If &#x60;true&#x60;, the resource pool will be updated as a background task. The task will have JobType UpdateHypervisorResourcePool. When the task is complete it will redirect to GetHypervisorResourcePool. The job&#39;s Parameters will contain properties: * _Name_ - Name of the resource pool being updated. * _Id_ - ID of the resource pool being updated. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPPutBrokerHypervisorConnection

> HypervisorsTPPutBrokerHypervisorConnection(ctx, nameOrId, customerid, siteid).Model(model).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Set the properties on a broker hypervisor connection.

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
    nameOrId := "nameOrId_example" // string | Hypervisor connection name or id. 
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    model := *openapiclient.NewEditBrokerHypervisorConnectionRequestModel() // EditBrokerHypervisorConnectionRequestModel | EditBrokerHypervisorConnectionRequestModel or derived objects. 
    async := true // bool | If async calling.  (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HypervisorsTPApi.HypervisorsTPPutBrokerHypervisorConnection(context.Background(), nameOrId, customerid, siteid).Model(model).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPPutBrokerHypervisorConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Hypervisor connection name or id.  | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPPutBrokerHypervisorConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **model** | [**EditBrokerHypervisorConnectionRequestModel**](EditBrokerHypervisorConnectionRequestModel.md) | EditBrokerHypervisorConnectionRequestModel or derived objects.  | 
 **async** | **bool** | If async calling.  | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPResetBrokerHypervisorConnection

> HypervisorsTPResetBrokerHypervisorConnection(ctx, nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Requests the broker hypervisor connection to be reset. The connection is dropped, details including credentials refreshed and the connection reestablished. The reset request is asynchronous and may take a moment to occur.

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
    nameOrId := "nameOrId_example" // string | Hypervisor connection name or id. 
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If async calling.  (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HypervisorsTPApi.HypervisorsTPResetBrokerHypervisorConnection(context.Background(), nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPResetBrokerHypervisorConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Hypervisor connection name or id.  | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPResetBrokerHypervisorConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If async calling.  | [default to false]
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


## HypervisorsTPTestHypervisor

> HypervisorTestResponseModel HypervisorsTPTestHypervisor(ctx, nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Run tests on a hypervisor and create a test report.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor to test.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the tests will run as a background task. This is recommended as the tests may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType TestHypervisor. When the task is complete it will redirect to GetJobResults. The job's Parameters will contain properties:  * _Id_ - ID of the hypervisor being tested, * _Name_ - Name of the hypervisor being tested. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPTestHypervisor(context.Background(), nameOrId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPTestHypervisor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPTestHypervisor`: HypervisorTestResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPTestHypervisor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor to test. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPTestHypervisorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **async** | **bool** | If &#x60;true&#x60;, the tests will run as a background task. This is recommended as the tests may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType TestHypervisor. When the task is complete it will redirect to GetJobResults. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the hypervisor being tested, * _Name_ - Name of the hypervisor being tested. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**HypervisorTestResponseModel**](HypervisorTestResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPTestHypervisorResourcePool

> HypervisorResourcePoolTestResponseModel HypervisorsTPTestHypervisorResourcePool(ctx, nameOrId, poolId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Run tests on a resource pool and create a test report.



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
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor containing the resource pool.
    poolId := "poolId_example" // string | Name or ID of the resource pool to test.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    async := true // bool | If `true`, the tests will run as a background task. This is recommended as the tests may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType TestHypervisorResourcePool. When the task is complete it will redirect to GetJobResults. The job's Parameters will contain properties:  * _Id_ - ID of the resource pool being tested, * _Name_ - Name of the resource pool being tested. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPTestHypervisorResourcePool(context.Background(), nameOrId, poolId, customerid, siteid).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPTestHypervisorResourcePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPTestHypervisorResourcePool`: HypervisorResourcePoolTestResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPTestHypervisorResourcePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor containing the resource pool. | 
**poolId** | **string** | Name or ID of the resource pool to test. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPTestHypervisorResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **async** | **bool** | If &#x60;true&#x60;, the tests will run as a background task. This is recommended as the tests may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType TestHypervisorResourcePool. When the task is complete it will redirect to GetJobResults. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the resource pool being tested, * _Name_ - Name of the resource pool being tested. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**HypervisorResourcePoolTestResponseModel**](HypervisorResourcePoolTestResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPValidateHypervisorResourcePoolResource

> HypervisorResourceValidationResponseModel HypervisorsTPValidateHypervisorResourcePoolResource(ctx, nameOrId, poolId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Validate a resource for certain usages, currently default and only usage is for Machine Profile.

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
    nameOrId := "nameOrId_example" // string | Name or Id of the hypervisor.
    poolId := "poolId_example" // string | Name or Id of the resource pool.
    customerid := "customerid_example" // string | 
    siteid := "siteid_example" // string | 
    request := *openapiclient.NewHypervisorResourceValidationRequestModel("Path_example") // HypervisorResourceValidationRequestModel | More details attached to this validation.
    async := true // bool | If async execute. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPValidateHypervisorResourcePoolResource(context.Background(), nameOrId, poolId, customerid, siteid).Request(request).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPValidateHypervisorResourcePoolResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPValidateHypervisorResourcePoolResource`: HypervisorResourceValidationResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPValidateHypervisorResourcePoolResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or Id of the hypervisor. | 
**poolId** | **string** | Name or Id of the resource pool. | 
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPValidateHypervisorResourcePoolResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **request** | [**HypervisorResourceValidationRequestModel**](HypervisorResourceValidationRequestModel.md) | More details attached to this validation. | 
 **async** | **bool** | If async execute. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**HypervisorResourceValidationResponseModel**](HypervisorResourceValidationResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsTPValidateVmPath

> HypervisorVmValidationResponseModelCollection HypervisorsTPValidateVmPath(ctx, customerid, siteid).VmPaths(vmPaths).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()

Validate a list of hypervisor Vm Path.



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
    vmPaths := *openapiclient.NewHypervisorVmValidationRequestModel() // HypervisorVmValidationRequestModel | A list of Vm path.
    async := true // bool | If `true`, the resource pool will be updated as a background task. The task will have JobType HypervisorVmValidation. (optional) (default to false)
    authorization := "authorization_example" // string | Citrix Cloud authorization header: Bearer {token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    xActionName := "xActionName_example" // string | Orchestration Action Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsTPApi.HypervisorsTPValidateVmPath(context.Background(), customerid, siteid).VmPaths(vmPaths).Async(async).Authorization(authorization).CitrixTransactionId(citrixTransactionId).XActionName(xActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsTPApi.HypervisorsTPValidateVmPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTPValidateVmPath`: HypervisorVmValidationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsTPApi.HypervisorsTPValidateVmPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerid** | **string** |  | 
**siteid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTPValidateVmPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vmPaths** | [**HypervisorVmValidationRequestModel**](HypervisorVmValidationRequestModel.md) | A list of Vm path. | 
 **async** | **bool** | If &#x60;true&#x60;, the resource pool will be updated as a background task. The task will have JobType HypervisorVmValidation. | [default to false]
 **authorization** | **string** | Citrix Cloud authorization header: Bearer {token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **xActionName** | **string** | Orchestration Action Name | 

### Return type

[**HypervisorVmValidationResponseModelCollection**](HypervisorVmValidationResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

