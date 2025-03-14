# \HypervisorsAPIsDAAS

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HypervisorsCheckHypervisorExists**](HypervisorsAPIsDAAS.md#HypervisorsCheckHypervisorExists) | **Head** /hypervisors/{name} | Check for the existence of a hypervisor by name.
[**HypervisorsCheckHypervisorResourcePoolExists**](HypervisorsAPIsDAAS.md#HypervisorsCheckHypervisorResourcePoolExists) | **Head** /hypervisors/{nameOrId}/resourcePools/{name} | Check for the existence of a resource pool by name.
[**HypervisorsCreateHypervisor**](HypervisorsAPIsDAAS.md#HypervisorsCreateHypervisor) | **Post** /hypervisors | Create a hypervisor, and optionally a resource pool on the hypervisor.
[**HypervisorsCreateResourcePool**](HypervisorsAPIsDAAS.md#HypervisorsCreateResourcePool) | **Post** /hypervisors/{nameOrId}/resourcePools | Create a new resource pool on an existing hypervisor.
[**HypervisorsDeleteHypervisor**](HypervisorsAPIsDAAS.md#HypervisorsDeleteHypervisor) | **Delete** /hypervisors/{nameOrId} | Delete a hypervisor and related resource pools.
[**HypervisorsDeleteHypervisorResourcePool**](HypervisorsAPIsDAAS.md#HypervisorsDeleteHypervisorResourcePool) | **Delete** /hypervisors/{nameOrId}/resourcePools/{poolId} | Delete a hypervisor resource pool.
[**HypervisorsDoHypervisorAllResourcesSearch**](HypervisorsAPIsDAAS.md#HypervisorsDoHypervisorAllResourcesSearch) | **Post** /hypervisors/{nameOrId}/allResources/$search | Search the resources within a hypervisor.
[**HypervisorsDoHypervisorAllResourcesSearchWithoutConnection**](HypervisorsAPIsDAAS.md#HypervisorsDoHypervisorAllResourcesSearchWithoutConnection) | **Post** /hypervisors/allResources/$search | Search all resources within a hypervisor, _without_ creating a persistent connection to the hypervisor.
[**HypervisorsDoHypervisorResourceSearch**](HypervisorsAPIsDAAS.md#HypervisorsDoHypervisorResourceSearch) | **Post** /hypervisors/{nameOrId}/resourcePools/{poolId}/resources/$search | Search the resources within a hypervisor resource pool.
[**HypervisorsGetHypervisor**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisor) | **Get** /hypervisors/{nameOrId} | Get the details for a single hypervisor.
[**HypervisorsGetHypervisorAdministrators**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisorAdministrators) | **Get** /hypervisors/{nameOrId}/administrators | Get administrators who can administer a hypervisor.
[**HypervisorsGetHypervisorAllResources**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisorAllResources) | **Get** /hypervisors/{nameOrId}/allResources | Get all resources within a hypervisor.
[**HypervisorsGetHypervisorAllResourcesWithoutConnection**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisorAllResourcesWithoutConnection) | **Post** /hypervisors/$GetResources | Get all resources within a hypervisor, _without_ creating a persistent connection to the hypervisor.
[**HypervisorsGetHypervisorDeletePreview**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisorDeletePreview) | **Get** /hypervisors/{nameOrId}/deletePreview | Get the hypervisor delete preview.
[**HypervisorsGetHypervisorJobs**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisorJobs) | **Get** /hypervisors/{nameOrId}/jobs | Get the currently active jobs that are using a hypervisor.
[**HypervisorsGetHypervisorMachineCatalogs**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisorMachineCatalogs) | **Get** /hypervisors/{nameOrId}/machineCatalogs | Get the machine catalogs that are using a hypervisor.
[**HypervisorsGetHypervisorOrphanedResources**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisorOrphanedResources) | **Get** /Hypervisors/{nameOrId}/OrphanedResources | Run detect on a hypervisor and retrieve orphaned resources.
[**HypervisorsGetHypervisorResourcePool**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisorResourcePool) | **Get** /hypervisors/{nameOrId}/resourcePools/{poolId} | Get details about a hypervisor resource pool.
[**HypervisorsGetHypervisorResourcePoolAdministrators**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisorResourcePoolAdministrators) | **Get** /hypervisors/{nameOrId}/resourcePools/{poolId}/administrators | Get administrators who can administer a resource pool.
[**HypervisorsGetHypervisorResourcePoolAllAvailableNetworks**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisorResourcePoolAllAvailableNetworks) | **Get** /hypervisorResourcePoolAllAvailableNetworks | Get all available networks among hypervisors and resource pools.
[**HypervisorsGetHypervisorResourcePoolDeletePreview**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisorResourcePoolDeletePreview) | **Get** /hypervisors/{nameOrId}/resourcePools/{poolId}/deletePreview | Get the hypervisor resource pool delete preview.
[**HypervisorsGetHypervisorResourcePoolJobs**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisorResourcePoolJobs) | **Get** /hypervisors/{nameOrId}/resourcePools/{poolId}/jobs | Get the currently active jobs that are using a resource pool.
[**HypervisorsGetHypervisorResourcePoolMachineCatalogs**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisorResourcePoolMachineCatalogs) | **Get** /hypervisors/{nameOrId}/resourcePools/{poolId}/machineCatalogs | Get the machine catalogs that are using a resource pool.
[**HypervisorsGetHypervisorResourcePoolPatchPreview**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisorResourcePoolPatchPreview) | **Post** /hypervisors/{nameOrId}/resourcePools/{poolId}/$getPatchPreview | Get the hypervisor resource pool patch preview.
[**HypervisorsGetHypervisorResourcePoolResources**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisorResourcePoolResources) | **Get** /hypervisors/{nameOrId}/resourcePools/{poolId}/resources | Get the resources within a hypervisor resource pool.
[**HypervisorsGetHypervisorResourcePoolTestReport**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisorResourcePoolTestReport) | **Get** /Hypervisors/{nameOrId}/ResourcePools/{poolId}/TestReport | Get the most recent test report for a resource pool.
[**HypervisorsGetHypervisorResourcePools**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisorResourcePools) | **Get** /hypervisors/{nameOrId}/resourcePools | Get the list of hypervisor resource pools.
[**HypervisorsGetHypervisorResourcePrices**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisorResourcePrices) | **Post** /hypervisors/{nameOrId}/resourcePrices | Get the hypervisor resource prices from settings.
[**HypervisorsGetHypervisorServerHAAddresses**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisorServerHAAddresses) | **Get** /hypervisors/{nameOrId}/serverHAAddresses | Get hypervisor server HA addresses. Currently, it only valid for Citrix hypervisors.
[**HypervisorsGetHypervisorTestReport**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisorTestReport) | **Get** /Hypervisors/{nameOrId}/TestReport | Get the most recent test report for a hypervisor.
[**HypervisorsGetHypervisorWithoutConnection**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisorWithoutConnection) | **Post** /$GetHypervisor | Get a hypervisor details without creating persistent connection.
[**HypervisorsGetHypervisors**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisors) | **Get** /hypervisors | Get the hypervisors.
[**HypervisorsGetHypervisorsAndResourcePools**](HypervisorsAPIsDAAS.md#HypervisorsGetHypervisorsAndResourcePools) | **Get** /hypervisorsAndResourcePools | Get hypervisors and resource pools. This API is used for the hosting main view.
[**HypervisorsGetSupportHypervisors**](HypervisorsAPIsDAAS.md#HypervisorsGetSupportHypervisors) | **Get** /hypervisors/supportedPlugins | Get current server support hypervisors.
[**HypervisorsPatchHypervisor**](HypervisorsAPIsDAAS.md#HypervisorsPatchHypervisor) | **Patch** /hypervisors/{nameOrId} | Update a hypervisor.
[**HypervisorsPatchHypervisorResourcePool**](HypervisorsAPIsDAAS.md#HypervisorsPatchHypervisorResourcePool) | **Patch** /hypervisors/{nameOrId}/resourcePools/{poolId} | Update a hypervisor resource pool.
[**HypervisorsPutBrokerHypervisorConnection**](HypervisorsAPIsDAAS.md#HypervisorsPutBrokerHypervisorConnection) | **Put** /hypervisors/{nameOrId}/connection | Set the properties on a broker hypervisor connection.
[**HypervisorsResetBrokerHypervisorConnection**](HypervisorsAPIsDAAS.md#HypervisorsResetBrokerHypervisorConnection) | **Post** /hypervisors/{nameOrId}/$resetConnection | Requests the broker hypervisor connection to be reset. The connection is dropped, details including credentials refreshed and the connection reestablished. The reset request is asynchronous and may take a moment to occur.
[**HypervisorsTestHypervisor**](HypervisorsAPIsDAAS.md#HypervisorsTestHypervisor) | **Post** /Hypervisors/{nameOrId}/$test | Run tests on a hypervisor and create a test report.
[**HypervisorsTestHypervisorResourcePool**](HypervisorsAPIsDAAS.md#HypervisorsTestHypervisorResourcePool) | **Post** /Hypervisors/{nameOrId}/ResourcePools/{poolId}/$test | Run tests on a resource pool and create a test report.
[**HypervisorsValidateHypervisorResourcePoolResource**](HypervisorsAPIsDAAS.md#HypervisorsValidateHypervisorResourcePoolResource) | **Post** /hypervisors/{nameOrId}/resourcePools/{poolId}/resources/$validate | Validate a resource for certain usages, currently default and only usage is for Machine Profile.
[**HypervisorsValidateVmPath**](HypervisorsAPIsDAAS.md#HypervisorsValidateVmPath) | **Post** /hypervisors/$ValidateVm | Validate a list of hypervisor Vm Path.



## HypervisorsCheckHypervisorExists

> HypervisorsCheckHypervisorExists(ctx, name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    name := "name_example" // string | Name of the hypervisor.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsCheckHypervisorExists(context.Background(), name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsCheckHypervisorExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the hypervisor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsCheckHypervisorExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## HypervisorsCheckHypervisorResourcePoolExists

> HypervisorsCheckHypervisorResourcePoolExists(ctx, nameOrId, name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    name := "name_example" // string | Name of the resource pool.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsCheckHypervisorResourcePoolExists(context.Background(), nameOrId, name).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsCheckHypervisorResourcePoolExists``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsCheckHypervisorResourcePoolExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## HypervisorsCreateHypervisor

> HypervisorDetailResponseModel HypervisorsCreateHypervisor(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateHypervisorRequestModel(createHypervisorRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    createHypervisorRequestModel := *openapiclient.NewCreateHypervisorRequestModel(*openapiclient.NewHypervisorConnectionDetailRequestModel("Name of the hypervisor", openapiclient.HypervisorConnectionType("Unknown"))) // CreateHypervisorRequestModel | Details of the hypervisor to create.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the hypervisor (and associated objects) will be created as a background task. The task will have JobType CreateHypervisor. When the task is complete it will redirect to GetHypervisor. The job's Parameters will contain properties: * _Name_ - Name of the hypervisor being created. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsCreateHypervisor(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateHypervisorRequestModel(createHypervisorRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsCreateHypervisor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsCreateHypervisor`: HypervisorDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsCreateHypervisor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsCreateHypervisorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **createHypervisorRequestModel** | [**CreateHypervisorRequestModel**](CreateHypervisorRequestModel.md) | Details of the hypervisor to create. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the hypervisor (and associated objects) will be created as a background task. The task will have JobType CreateHypervisor. When the task is complete it will redirect to GetHypervisor. The job&#39;s Parameters will contain properties: * _Name_ - Name of the hypervisor being created. | [default to false]

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


## HypervisorsCreateResourcePool

> HypervisorResourcePoolDetailResponseModel HypervisorsCreateResourcePool(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateHypervisorResourcePoolRequestModel(createHypervisorResourcePoolRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    createHypervisorResourcePoolRequestModel := *openapiclient.NewCreateHypervisorResourcePoolRequestModel("Name of the resource pool", openapiclient.HypervisorConnectionType("Unknown")) // CreateHypervisorResourcePoolRequestModel | Details of the resource pool to create.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the resource pool (and associated objects) will be created as a background task. The task will have JobType CreateHypervisorResourcePool. When the task is complete it will redirect to GetHypervisorResourcePool. The job's Parameters will contain properties: * _Name_ - Name of the resource pool being created. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsCreateResourcePool(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).CreateHypervisorResourcePoolRequestModel(createHypervisorResourcePoolRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsCreateResourcePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsCreateResourcePool`: HypervisorResourcePoolDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsCreateResourcePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsCreateResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **createHypervisorResourcePoolRequestModel** | [**CreateHypervisorResourcePoolRequestModel**](CreateHypervisorResourcePoolRequestModel.md) | Details of the resource pool to create. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the resource pool (and associated objects) will be created as a background task. The task will have JobType CreateHypervisorResourcePool. When the task is complete it will redirect to GetHypervisorResourcePool. The job&#39;s Parameters will contain properties: * _Name_ - Name of the resource pool being created. | [default to false]

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


## HypervisorsDeleteHypervisor

> HypervisorsDeleteHypervisor(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | The hypervisor connection name or id. 
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If the api call with async type.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsDeleteHypervisor(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsDeleteHypervisor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | The hypervisor connection name or id.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsDeleteHypervisorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If the api call with async type.  | [default to false]

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


## HypervisorsDeleteHypervisorResourcePool

> HypervisorsDeleteHypervisorResourcePool(ctx, nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | The name or id of hypervisor. 
    poolId := "poolId_example" // string | The hypervisor resource pool id. 
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If the api call with async type.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsDeleteHypervisorResourcePool(context.Background(), nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsDeleteHypervisorResourcePool``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsDeleteHypervisorResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If the api call with async type.  | [default to false]

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


## HypervisorsDoHypervisorAllResourcesSearch

> HypervisorResourceResponseModelCollection HypervisorsDoHypervisorAllResourcesSearch(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).HypervisorResourceSearchRequestModel(hypervisorResourceSearchRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Detail(detail).Async(async).Execute()

Search the resources within a hypervisor.



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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    hypervisorResourceSearchRequestModel := *openapiclient.NewHypervisorResourceSearchRequestModel() // HypervisorResourceSearchRequestModel | Specifies the advanced search parameters.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    limit := int32(56) // int32 | The max number of resources returned by this query. If not specified, the server might use a default limit of 1 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    detail := true // bool | If `true`, full details of templates will be retrieved. This can be very time consuming and will reduce the performance of the call. (optional) (default to false)
    async := true // bool | Async request to get the resources with *path.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsDoHypervisorAllResourcesSearch(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).HypervisorResourceSearchRequestModel(hypervisorResourceSearchRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Detail(detail).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsDoHypervisorAllResourcesSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsDoHypervisorAllResourcesSearch`: HypervisorResourceResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsDoHypervisorAllResourcesSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsDoHypervisorAllResourcesSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **hypervisorResourceSearchRequestModel** | [**HypervisorResourceSearchRequestModel**](HypervisorResourceSearchRequestModel.md) | Specifies the advanced search parameters. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | The max number of resources returned by this query. If not specified, the server might use a default limit of 1 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **detail** | **bool** | If &#x60;true&#x60;, full details of templates will be retrieved. This can be very time consuming and will reduce the performance of the call. | [default to false]
 **async** | **bool** | Async request to get the resources with *path.  | [default to false]

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


## HypervisorsDoHypervisorAllResourcesSearchWithoutConnection

> HypervisorResourceResponseModelCollection HypervisorsDoHypervisorAllResourcesSearchWithoutConnection(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).HypervisorAllResourceSearchRequestModel(hypervisorAllResourceSearchRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Detail(detail).Async(async).Execute()

Search all resources within a hypervisor, _without_ creating a persistent connection to the hypervisor.

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    hypervisorAllResourceSearchRequestModel := *openapiclient.NewHypervisorAllResourceSearchRequestModel(*openapiclient.NewHypervisorConnectionDetailRequestModel("Name of the hypervisor", openapiclient.HypervisorConnectionType("Unknown")), *openapiclient.NewHypervisorResourceSearchRequestModel()) // HypervisorAllResourceSearchRequestModel | Contains hypervisor details and search criteria.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    limit := int32(56) // int32 | The max number of resources returned by this query. If not specified, the server might use a default limit of 1 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    detail := true // bool | If `true`, full details of templates will be retrieved. This can be very time consuming and will reduce the performance of the call. (optional) (default to false)
    async := true // bool | Async request to get the resources with *path.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsDoHypervisorAllResourcesSearchWithoutConnection(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).HypervisorAllResourceSearchRequestModel(hypervisorAllResourceSearchRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Detail(detail).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsDoHypervisorAllResourcesSearchWithoutConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsDoHypervisorAllResourcesSearchWithoutConnection`: HypervisorResourceResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsDoHypervisorAllResourcesSearchWithoutConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsDoHypervisorAllResourcesSearchWithoutConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **hypervisorAllResourceSearchRequestModel** | [**HypervisorAllResourceSearchRequestModel**](HypervisorAllResourceSearchRequestModel.md) | Contains hypervisor details and search criteria. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | The max number of resources returned by this query. If not specified, the server might use a default limit of 1 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **detail** | **bool** | If &#x60;true&#x60;, full details of templates will be retrieved. This can be very time consuming and will reduce the performance of the call. | [default to false]
 **async** | **bool** | Async request to get the resources with *path.  | [default to false]

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


## HypervisorsDoHypervisorResourceSearch

> HypervisorResourceResponseModelCollection HypervisorsDoHypervisorResourceSearch(ctx, nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).HypervisorResourceSearchRequestModel(hypervisorResourceSearchRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Detail(detail).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    poolId := "poolId_example" // string | Name or ID of the resource pool.
    hypervisorResourceSearchRequestModel := *openapiclient.NewHypervisorResourceSearchRequestModel() // HypervisorResourceSearchRequestModel | Specifies the advanced search parameters.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    limit := int32(56) // int32 | The max number of resources returned by this query. If not specified, the server might use a default limit of 1 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)
    detail := true // bool | If `true`, full details of templates will be retrieved. This can be very time consuming and will reduce the performance of the call. (optional) (default to false)
    async := true // bool | Async request to get the resources with *path.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsDoHypervisorResourceSearch(context.Background(), nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).HypervisorResourceSearchRequestModel(hypervisorResourceSearchRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Limit(limit).ContinuationToken(continuationToken).Detail(detail).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsDoHypervisorResourceSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsDoHypervisorResourceSearch`: HypervisorResourceResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsDoHypervisorResourceSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**poolId** | **string** | Name or ID of the resource pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsDoHypervisorResourceSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **hypervisorResourceSearchRequestModel** | [**HypervisorResourceSearchRequestModel**](HypervisorResourceSearchRequestModel.md) | Specifies the advanced search parameters. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **limit** | **int32** | The max number of resources returned by this query. If not specified, the server might use a default limit of 1 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 
 **detail** | **bool** | If &#x60;true&#x60;, full details of templates will be retrieved. This can be very time consuming and will reduce the performance of the call. | [default to false]
 **async** | **bool** | Async request to get the resources with *path.  | [default to false]

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


## HypervisorsGetHypervisor

> HypervisorDetailResponseModel HypervisorsGetHypervisor(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).NoCache(noCache).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | Async request to hypervisor.  (optional) (default to false)
    noCache := true // bool | Disable the cache.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisor(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).NoCache(noCache).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisor`: HypervisorDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | Async request to hypervisor.  | [default to false]
 **noCache** | **bool** | Disable the cache.  | [default to false]

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


## HypervisorsGetHypervisorAdministrators

> AdministratorResponseModelCollection HypervisorsGetHypervisorAdministrators(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Limit(limit).ContinuationToken(continuationToken).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If async execute.  (optional) (default to false)
    limit := int32(56) // int32 | The max number of administrators returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisorAdministrators(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Limit(limit).ContinuationToken(continuationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisorAdministrators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisorAdministrators`: AdministratorResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisorAdministrators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorAdministratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If async execute.  | [default to false]
 **limit** | **int32** | The max number of administrators returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 

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


## HypervisorsGetHypervisorAllResources

> HypervisorResourceResponseModel HypervisorsGetHypervisorAllResources(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Path(path).Children(children).Type_(type_).ShowTagged(showTagged).Detail(detail).Async(async).NoCache(noCache).Force(force).Recurse(recurse).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    path := "path_example" // string | Path to the resource container within the hypervisor. The path may either be a relative path as specified by RelativePath, or may be a URL-encoded XenApp & XenDesktop resource path starting with `XDHyp:`; for example, as specified by XDPath. (optional)
    children := int32(56) // int32 | Specifies the number of levels of children to enumerate. Default is `0`, meaning that only the object referred to by `path` is returned and its Children array will be left null. A special value of `-1` indicates that the entire resource hierarchy should be enumerated.  Use with care!  It may take a very long time to enumerate a large number of resources from a hypervisor, and the call may time out before completing. (optional)
    type_ := []string{"Inner_example"} // []string | If specified, limits the results to the specified resource type(s). (optional)
    showTagged := true // bool | By default, items which are tagged by XenDesktop are not shown. Set this to `true` to override that behavior. (optional)
    detail := true // bool | If `true`, full details of VMs, snapshots, and templates will be retrieved. This can be very time consuming and will reduce the performance of the call. May only be used if `path` refers to a VM, snapshot, or template resource. <!-- Internally this calls Get-HypConfigurationObjectForItem --> (optional)
    async := true // bool | Async request to get the resources with *path. (optional) (default to false)
    noCache := true // bool | Disable the cache.  (optional) (default to false)
    force := true // bool | Get the VMs provisioned by MCS.  (optional) (default to false)
    recurse := true // bool | Fetch all resources under this folder.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisorAllResources(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Path(path).Children(children).Type_(type_).ShowTagged(showTagged).Detail(detail).Async(async).NoCache(noCache).Force(force).Recurse(recurse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisorAllResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisorAllResources`: HypervisorResourceResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisorAllResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorAllResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **path** | **string** | Path to the resource container within the hypervisor. The path may either be a relative path as specified by RelativePath, or may be a URL-encoded XenApp &amp; XenDesktop resource path starting with &#x60;XDHyp:&#x60;; for example, as specified by XDPath. | 
 **children** | **int32** | Specifies the number of levels of children to enumerate. Default is &#x60;0&#x60;, meaning that only the object referred to by &#x60;path&#x60; is returned and its Children array will be left null. A special value of &#x60;-1&#x60; indicates that the entire resource hierarchy should be enumerated.  Use with care!  It may take a very long time to enumerate a large number of resources from a hypervisor, and the call may time out before completing. | 
 **type_** | **[]string** | If specified, limits the results to the specified resource type(s). | 
 **showTagged** | **bool** | By default, items which are tagged by XenDesktop are not shown. Set this to &#x60;true&#x60; to override that behavior. | 
 **detail** | **bool** | If &#x60;true&#x60;, full details of VMs, snapshots, and templates will be retrieved. This can be very time consuming and will reduce the performance of the call. May only be used if &#x60;path&#x60; refers to a VM, snapshot, or template resource. &lt;!-- Internally this calls Get-HypConfigurationObjectForItem --&gt; | 
 **async** | **bool** | Async request to get the resources with *path. | [default to false]
 **noCache** | **bool** | Disable the cache.  | [default to false]
 **force** | **bool** | Get the VMs provisioned by MCS.  | [default to false]
 **recurse** | **bool** | Fetch all resources under this folder.  | [default to false]

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


## HypervisorsGetHypervisorAllResourcesWithoutConnection

> HypervisorResourceResponseModel HypervisorsGetHypervisorAllResourcesWithoutConnection(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).HypervisorConnectionDetailRequestModel(hypervisorConnectionDetailRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Path(path).Children(children).Type_(type_).ShowTagged(showTagged).Detail(detail).Async(async).Force(force).Recurse(recurse).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    hypervisorConnectionDetailRequestModel := *openapiclient.NewHypervisorConnectionDetailRequestModel("Name of the hypervisor", openapiclient.HypervisorConnectionType("Unknown")) // HypervisorConnectionDetailRequestModel | Details of how to connect to the hypervisor to enumerate resources from it.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    path := "path_example" // string | Path to the resource container within the hypervisor. The path may either be a relative path as specified by RelativePath, or may be a URL-encoded Virtual Apps & Desktops resource path starting with `XDHyp:`; for example, as specified by XDPath. (optional)
    children := int32(56) // int32 | Specifies the number of levels of children to enumerate. Default is `0`, meaning that only the object referred to by `path` is returned and its Children array will be left null. A special value of `-1` indicates that the entire resource hierarchy should be enumerated.  Use with care!  It may take a very long time to enumerate a large number of resources from a hypervisor, and the call may time out before completing. (optional)
    type_ := []string{"Inner_example"} // []string | If specified, limits the results to the specified resource type(s). (optional)
    showTagged := true // bool | By default, items which are tagged by XenDesktop are not shown. Set this to `true` to override that behavior. (optional)
    detail := true // bool | If `true`, full details of VMs, snapshots, and templates will be retrieved. This can be very time consuming and will reduce the performance of the call. May only be used if `path` refers to a VM, snapshot, or template resource. <!-- Internally this calls Get-HypConfigurationObjectForItem --> (optional)
    async := true // bool | Async request to get the resources with *path. (optional) (default to false)
    force := true // bool | Get the VMs provisioned by MCS.  (optional)
    recurse := true // bool | Fetch all resources under this folder.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisorAllResourcesWithoutConnection(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).HypervisorConnectionDetailRequestModel(hypervisorConnectionDetailRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Path(path).Children(children).Type_(type_).ShowTagged(showTagged).Detail(detail).Async(async).Force(force).Recurse(recurse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisorAllResourcesWithoutConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisorAllResourcesWithoutConnection`: HypervisorResourceResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisorAllResourcesWithoutConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorAllResourcesWithoutConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **hypervisorConnectionDetailRequestModel** | [**HypervisorConnectionDetailRequestModel**](HypervisorConnectionDetailRequestModel.md) | Details of how to connect to the hypervisor to enumerate resources from it. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **path** | **string** | Path to the resource container within the hypervisor. The path may either be a relative path as specified by RelativePath, or may be a URL-encoded Virtual Apps &amp; Desktops resource path starting with &#x60;XDHyp:&#x60;; for example, as specified by XDPath. | 
 **children** | **int32** | Specifies the number of levels of children to enumerate. Default is &#x60;0&#x60;, meaning that only the object referred to by &#x60;path&#x60; is returned and its Children array will be left null. A special value of &#x60;-1&#x60; indicates that the entire resource hierarchy should be enumerated.  Use with care!  It may take a very long time to enumerate a large number of resources from a hypervisor, and the call may time out before completing. | 
 **type_** | **[]string** | If specified, limits the results to the specified resource type(s). | 
 **showTagged** | **bool** | By default, items which are tagged by XenDesktop are not shown. Set this to &#x60;true&#x60; to override that behavior. | 
 **detail** | **bool** | If &#x60;true&#x60;, full details of VMs, snapshots, and templates will be retrieved. This can be very time consuming and will reduce the performance of the call. May only be used if &#x60;path&#x60; refers to a VM, snapshot, or template resource. &lt;!-- Internally this calls Get-HypConfigurationObjectForItem --&gt; | 
 **async** | **bool** | Async request to get the resources with *path. | [default to false]
 **force** | **bool** | Get the VMs provisioned by MCS.  | 
 **recurse** | **bool** | Fetch all resources under this folder.  | 

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


## HypervisorsGetHypervisorDeletePreview

> HypervisorDeletePreviewResponseModel HypervisorsGetHypervisorDeletePreview(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If execute this API asynchronous.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisorDeletePreview(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisorDeletePreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisorDeletePreview`: HypervisorDeletePreviewResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisorDeletePreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorDeletePreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If execute this API asynchronous.  | [default to false]

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


## HypervisorsGetHypervisorJobs

> JobResponseModelCollection HypervisorsGetHypervisorJobs(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisorJobs(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisorJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisorJobs`: JobResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisorJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## HypervisorsGetHypervisorMachineCatalogs

> HypervisorMachineCatalogResponseModelCollection HypervisorsGetHypervisorMachineCatalogs(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Limit(limit).ContinuationToken(continuationToken).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If async execute.  (optional) (default to false)
    limit := int32(56) // int32 | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisorMachineCatalogs(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Limit(limit).ContinuationToken(continuationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisorMachineCatalogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisorMachineCatalogs`: HypervisorMachineCatalogResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisorMachineCatalogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorMachineCatalogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If async execute.  | [default to false]
 **limit** | **int32** | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 

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


## HypervisorsGetHypervisorOrphanedResources

> DetectOrphanedResourcesResponseModel HypervisorsGetHypervisorOrphanedResources(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Run detect on a hypervisor and retrieve orphaned resources.



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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor to detect.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the tests will run as a background task. This is recommended as the tests may take some time to run, and if it exceeds 90 seconds the request may time out.  * _Id_ - ID of the hypervisor being detected, * _Name_ - Name of the hypervisor being detected. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisorOrphanedResources(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisorOrphanedResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisorOrphanedResources`: DetectOrphanedResourcesResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisorOrphanedResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor to detect. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorOrphanedResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the tests will run as a background task. This is recommended as the tests may take some time to run, and if it exceeds 90 seconds the request may time out.  * _Id_ - ID of the hypervisor being detected, * _Name_ - Name of the hypervisor being detected. | [default to false]

### Return type

[**DetectOrphanedResourcesResponseModel**](DetectOrphanedResourcesResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsGetHypervisorResourcePool

> HypervisorResourcePoolDetailResponseModel HypervisorsGetHypervisorResourcePool(ctx, nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).NoCache(noCache).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    poolId := "poolId_example" // string | Name or ID of the resource pool.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | Async request to get the resource pool.  (optional) (default to false)
    noCache := true // bool | Disable the cache.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePool(context.Background(), nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).NoCache(noCache).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisorResourcePool`: HypervisorResourcePoolDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**poolId** | **string** | Name or ID of the resource pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | Async request to get the resource pool.  | [default to false]
 **noCache** | **bool** | Disable the cache.  | [default to false]

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


## HypervisorsGetHypervisorResourcePoolAdministrators

> AdministratorResponseModelCollection HypervisorsGetHypervisorResourcePoolAdministrators(ctx, nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Limit(limit).ContinuationToken(continuationToken).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    poolId := "poolId_example" // string | Name or ID of the resource pool.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If async execute.  (optional) (default to false)
    limit := int32(56) // int32 | The max number of administrators returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolAdministrators(context.Background(), nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Limit(limit).ContinuationToken(continuationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolAdministrators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisorResourcePoolAdministrators`: AdministratorResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolAdministrators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**poolId** | **string** | Name or ID of the resource pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorResourcePoolAdministratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If async execute.  | [default to false]
 **limit** | **int32** | The max number of administrators returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 

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


## HypervisorsGetHypervisorResourcePoolAllAvailableNetworks

> HypervisorResourcePoolNetworkResponseModelCollection HypervisorsGetHypervisorResourcePoolAllAvailableNetworks(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Get all available networks among hypervisors and resource pools.

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If execute this API asynchronous.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolAllAvailableNetworks(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolAllAvailableNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisorResourcePoolAllAvailableNetworks`: HypervisorResourcePoolNetworkResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolAllAvailableNetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorResourcePoolAllAvailableNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If execute this API asynchronous.  | [default to false]

### Return type

[**HypervisorResourcePoolNetworkResponseModelCollection**](HypervisorResourcePoolNetworkResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsGetHypervisorResourcePoolDeletePreview

> HypervisorDeletePreviewResponseModel HypervisorsGetHypervisorResourcePoolDeletePreview(ctx, nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    poolId := "poolId_example" // string | Name or ID of the resource pool.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If execute this API asynchronous.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolDeletePreview(context.Background(), nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolDeletePreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisorResourcePoolDeletePreview`: HypervisorDeletePreviewResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolDeletePreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**poolId** | **string** | Name or ID of the resource pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorResourcePoolDeletePreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If execute this API asynchronous.  | [default to false]

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


## HypervisorsGetHypervisorResourcePoolJobs

> JobResponseModelCollection HypervisorsGetHypervisorResourcePoolJobs(ctx, nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    poolId := "poolId_example" // string | Name or ID of the resource pool.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolJobs(context.Background(), nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisorResourcePoolJobs`: JobResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**poolId** | **string** | Name or ID of the resource pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorResourcePoolJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## HypervisorsGetHypervisorResourcePoolMachineCatalogs

> HypervisorMachineCatalogResponseModelCollection HypervisorsGetHypervisorResourcePoolMachineCatalogs(ctx, nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Limit(limit).ContinuationToken(continuationToken).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    poolId := "poolId_example" // string | Name or ID of the resource pool.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If async execute.  (optional) (default to false)
    limit := int32(56) // int32 | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. (optional)
    continuationToken := "continuationToken_example" // string | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolMachineCatalogs(context.Background(), nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Limit(limit).ContinuationToken(continuationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolMachineCatalogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisorResourcePoolMachineCatalogs`: HypervisorMachineCatalogResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolMachineCatalogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**poolId** | **string** | Name or ID of the resource pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorResourcePoolMachineCatalogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If async execute.  | [default to false]
 **limit** | **int32** | The max number of machine catalogs returned by this query. If not specified, the server might use a default limit of 250 items. If the specified value is larger than 1000, the server might reject the call. The default and maximum values depend on server settings. | 
 **continuationToken** | **string** | If a query cannot be completed, the response will have a ContinuationToken set. To obtain more results from the query, pass the continuation token back into the query to get the next batch of results. | 

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


## HypervisorsGetHypervisorResourcePoolPatchPreview

> HypervisorResourcePoolPatchPreviewResponseModel HypervisorsGetHypervisorResourcePoolPatchPreview(ctx, nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).HypervisorResourcePoolPatchPreviewRequestModel(hypervisorResourcePoolPatchPreviewRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    poolId := "poolId_example" // string | Name or ID of the resource pool.
    hypervisorResourcePoolPatchPreviewRequestModel := *openapiclient.NewHypervisorResourcePoolPatchPreviewRequestModel() // HypervisorResourcePoolPatchPreviewRequestModel | New networks of the resource pool.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the preview job will run as a background task. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolPatchPreview(context.Background(), nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).HypervisorResourcePoolPatchPreviewRequestModel(hypervisorResourcePoolPatchPreviewRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolPatchPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisorResourcePoolPatchPreview`: HypervisorResourcePoolPatchPreviewResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolPatchPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**poolId** | **string** | Name or ID of the resource pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorResourcePoolPatchPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **hypervisorResourcePoolPatchPreviewRequestModel** | [**HypervisorResourcePoolPatchPreviewRequestModel**](HypervisorResourcePoolPatchPreviewRequestModel.md) | New networks of the resource pool. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the preview job will run as a background task. | [default to false]

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


## HypervisorsGetHypervisorResourcePoolResources

> HypervisorResourceResponseModel HypervisorsGetHypervisorResourcePoolResources(ctx, nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Path(path).Children(children).Type_(type_).ShowTagged(showTagged).Detail(detail).Async(async).NoCache(noCache).Force(force).Recurse(recurse).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    poolId := "poolId_example" // string | Name or ID of the resource pool.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    path := "path_example" // string | Path to the resource container within the resource pool. The path may either be a relative path as specified by RelativePath, or may be a URL-encoded XenApp & XenDesktop resource path starting with `XDHyp:`; for example, as specified by XDPath. (optional)
    children := int32(56) // int32 | Specifies the number of levels of children to enumerate. Default is `0`, meaning that only the object referred to by `path` is returned and its Children array will be left null. A special value of `-1` indicates that the entire resource hierarchy should be enumerated.  Use with care!  It may take a very long time to enumerate a large number of resources from a hypervisor, and the call may time out before completing. (optional)
    type_ := []string{"Inner_example"} // []string | If specified, limits the results to the specified resource type(s). (optional)
    showTagged := true // bool | By default, items which are tagged by XenDesktop are not shown. Set this to `true` to override that behavior. (optional)
    detail := true // bool | If `true`, full details of VMs, snapshots, and templates will be retrieved. This can be very time consuming and will reduce the performance of the call. May only be used if `path` refers to a VM, snapshot, or template resource. <!-- Internally this calls Get-HypConfigurationObjectForItem --> (optional)
    async := true // bool | Async request to get the resources with *path.  (optional) (default to false)
    noCache := true // bool | Disable the cache.  (optional) (default to false)
    force := true // bool | Get the VMs provisioned by MCS.  (optional) (default to false)
    recurse := true // bool | Fetch all resources under this folder.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolResources(context.Background(), nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Path(path).Children(children).Type_(type_).ShowTagged(showTagged).Detail(detail).Async(async).NoCache(noCache).Force(force).Recurse(recurse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisorResourcePoolResources`: HypervisorResourceResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**poolId** | **string** | Name or ID of the resource pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorResourcePoolResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **path** | **string** | Path to the resource container within the resource pool. The path may either be a relative path as specified by RelativePath, or may be a URL-encoded XenApp &amp; XenDesktop resource path starting with &#x60;XDHyp:&#x60;; for example, as specified by XDPath. | 
 **children** | **int32** | Specifies the number of levels of children to enumerate. Default is &#x60;0&#x60;, meaning that only the object referred to by &#x60;path&#x60; is returned and its Children array will be left null. A special value of &#x60;-1&#x60; indicates that the entire resource hierarchy should be enumerated.  Use with care!  It may take a very long time to enumerate a large number of resources from a hypervisor, and the call may time out before completing. | 
 **type_** | **[]string** | If specified, limits the results to the specified resource type(s). | 
 **showTagged** | **bool** | By default, items which are tagged by XenDesktop are not shown. Set this to &#x60;true&#x60; to override that behavior. | 
 **detail** | **bool** | If &#x60;true&#x60;, full details of VMs, snapshots, and templates will be retrieved. This can be very time consuming and will reduce the performance of the call. May only be used if &#x60;path&#x60; refers to a VM, snapshot, or template resource. &lt;!-- Internally this calls Get-HypConfigurationObjectForItem --&gt; | 
 **async** | **bool** | Async request to get the resources with *path.  | [default to false]
 **noCache** | **bool** | Disable the cache.  | [default to false]
 **force** | **bool** | Get the VMs provisioned by MCS.  | [default to false]
 **recurse** | **bool** | Fetch all resources under this folder.  | [default to false]

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


## HypervisorsGetHypervisorResourcePoolTestReport

> TestReportResponseModel HypervisorsGetHypervisorResourcePoolTestReport(ctx, nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    poolId := "poolId_example" // string | Name or ID of the resource pool.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolTestReport(context.Background(), nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolTestReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisorResourcePoolTestReport`: TestReportResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePoolTestReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 
**poolId** | **string** | Name or ID of the resource pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorResourcePoolTestReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## HypervisorsGetHypervisorResourcePools

> HypervisorResourcePoolResponseModelCollection HypervisorsGetHypervisorResourcePools(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).NoCache(noCache).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | Async request to get the resource pool.  (optional) (default to false)
    noCache := true // bool | Disable the cache.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePools(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).NoCache(noCache).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisorResourcePools`: HypervisorResourcePoolResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorResourcePoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | Async request to get the resource pool.  | [default to false]
 **noCache** | **bool** | Disable the cache.  | [default to false]

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


## HypervisorsGetHypervisorResourcePrices

> HypervisorResourcePricesResponseModel HypervisorsGetHypervisorResourcePrices(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).HypervisorResourceSettingsRequestModel(hypervisorResourceSettingsRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Get the hypervisor resource prices from settings.

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | The hypervisor connection name or id.
    hypervisorResourceSettingsRequestModel := *openapiclient.NewHypervisorResourceSettingsRequestModel() // HypervisorResourceSettingsRequestModel | The resource settings model.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If the api call with async type. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePrices(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).HypervisorResourceSettingsRequestModel(hypervisorResourceSettingsRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisorResourcePrices`: HypervisorResourcePricesResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisorResourcePrices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | The hypervisor connection name or id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorResourcePricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **hypervisorResourceSettingsRequestModel** | [**HypervisorResourceSettingsRequestModel**](HypervisorResourceSettingsRequestModel.md) | The resource settings model. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If the api call with async type. | [default to false]

### Return type

[**HypervisorResourcePricesResponseModel**](HypervisorResourcePricesResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HypervisorsGetHypervisorServerHAAddresses

> HypervisorServerHAAddressesResponseModel HypervisorsGetHypervisorServerHAAddresses(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).NoCache(noCache).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | The hypervisor connection name or id. 
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If the execution with async model.  (optional) (default to false)
    noCache := true // bool | Disable the cache.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisorServerHAAddresses(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).NoCache(noCache).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisorServerHAAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisorServerHAAddresses`: HypervisorServerHAAddressesResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisorServerHAAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | The hypervisor connection name or id.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorServerHAAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If the execution with async model.  | [default to false]
 **noCache** | **bool** | Disable the cache.  | [default to false]

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


## HypervisorsGetHypervisorTestReport

> TestReportResponseModel HypervisorsGetHypervisorTestReport(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisorTestReport(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisorTestReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisorTestReport`: TestReportResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisorTestReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorTestReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 

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


## HypervisorsGetHypervisorWithoutConnection

> HypervisorDetailResponseModel HypervisorsGetHypervisorWithoutConnection(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).HypervisorConnectionDetailRequestModel(hypervisorConnectionDetailRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    hypervisorConnectionDetailRequestModel := *openapiclient.NewHypervisorConnectionDetailRequestModel("Name of the hypervisor", openapiclient.HypervisorConnectionType("Unknown")) // HypervisorConnectionDetailRequestModel | Details of how to connect to the hypervisor to enumerate resources from it.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | Async request to get the hypervisor.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisorWithoutConnection(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).HypervisorConnectionDetailRequestModel(hypervisorConnectionDetailRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisorWithoutConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisorWithoutConnection`: HypervisorDetailResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisorWithoutConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorWithoutConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **hypervisorConnectionDetailRequestModel** | [**HypervisorConnectionDetailRequestModel**](HypervisorConnectionDetailRequestModel.md) | Details of how to connect to the hypervisor to enumerate resources from it. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | Async request to get the hypervisor.  | [default to false]

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


## HypervisorsGetHypervisors

> HypervisorResponseModelCollection HypervisorsGetHypervisors(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).NoCache(noCache).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | Async request to hypervisor.  (optional) (default to false)
    noCache := true // bool | Disable the cache.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisors(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).NoCache(noCache).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisors`: HypervisorResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | Async request to hypervisor.  | [default to false]
 **noCache** | **bool** | Disable the cache.  | [default to false]

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


## HypervisorsGetHypervisorsAndResourcePools

> HypervisorsAndResourcePoolsResponseModelCollection HypervisorsGetHypervisorsAndResourcePools(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).NoCache(noCache).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If execute this API asynchronous.  (optional) (default to false)
    noCache := true // bool | Disable the cache.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetHypervisorsAndResourcePools(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).NoCache(noCache).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetHypervisorsAndResourcePools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetHypervisorsAndResourcePools`: HypervisorsAndResourcePoolsResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetHypervisorsAndResourcePools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetHypervisorsAndResourcePoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If execute this API asynchronous.  | [default to false]
 **noCache** | **bool** | Disable the cache.  | [default to false]

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


## HypervisorsGetSupportHypervisors

> HypervisorPluginResponseModelCollection HypervisorsGetSupportHypervisors(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).ZoneNameOrId(zoneNameOrId).IncludeUnavailable(includeUnavailable).NoCache(noCache).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If execute this API asynchronous.  (optional) (default to false)
    zoneNameOrId := "zoneNameOrId_example" // string | The zone name or id.  (optional)
    includeUnavailable := true // bool | Flag to show all supported hypervisor plugins. (optional) (default to false)
    noCache := true // bool | Disable the cache.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsGetSupportHypervisors(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).ZoneNameOrId(zoneNameOrId).IncludeUnavailable(includeUnavailable).NoCache(noCache).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsGetSupportHypervisors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsGetSupportHypervisors`: HypervisorPluginResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsGetSupportHypervisors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsGetSupportHypervisorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If execute this API asynchronous.  | [default to false]
 **zoneNameOrId** | **string** | The zone name or id.  | 
 **includeUnavailable** | **bool** | Flag to show all supported hypervisor plugins. | [default to false]
 **noCache** | **bool** | Disable the cache.  | [default to false]

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


## HypervisorsPatchHypervisor

> HypervisorsPatchHypervisor(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).EditHypervisorConnectionRequestModel(editHypervisorConnectionRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    editHypervisorConnectionRequestModel := *openapiclient.NewEditHypervisorConnectionRequestModel(openapiclient.HypervisorConnectionType("Unknown")) // EditHypervisorConnectionRequestModel | Details of the hypervisor to update.  Note that each type of hypervisor requires a different update model: * AWS requires a model of type EditHypervisorAWSConnectionRequestModel. * AzureRM requires a model of type EditHypervisorAzureConnectionRequestModel. * GoogleCloudPlatform requires a model of type EditHypervisorGCPConnectionRequestModel. * OracleCloudInfrastructure requires a model of type EditHypervisorOciConnectionRequestModel. * AzureArc requires a model of type EditHypervisorAzureArcConnectionRequestModel. * OpenShift requires a model of type EditHypervisorOpenShiftConnectionRequestModel. * All other hypervisor types require a model of type EditHypervisorTraditionalConnectionRequestModel.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the hypervisor will be updated as a background task. The task will have JobType UpdateHypervisor. When the task is complete it will redirect to GetHypervisor. The job's Parameters will contain properties: * _Name_ - Name of the hypervisor being updated. * _Id_ - ID of the hypervisor being updated. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsPatchHypervisor(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).EditHypervisorConnectionRequestModel(editHypervisorConnectionRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsPatchHypervisor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsPatchHypervisorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **editHypervisorConnectionRequestModel** | [**EditHypervisorConnectionRequestModel**](EditHypervisorConnectionRequestModel.md) | Details of the hypervisor to update.  Note that each type of hypervisor requires a different update model: * AWS requires a model of type EditHypervisorAWSConnectionRequestModel. * AzureRM requires a model of type EditHypervisorAzureConnectionRequestModel. * GoogleCloudPlatform requires a model of type EditHypervisorGCPConnectionRequestModel. * OracleCloudInfrastructure requires a model of type EditHypervisorOciConnectionRequestModel. * AzureArc requires a model of type EditHypervisorAzureArcConnectionRequestModel. * OpenShift requires a model of type EditHypervisorOpenShiftConnectionRequestModel. * All other hypervisor types require a model of type EditHypervisorTraditionalConnectionRequestModel. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the hypervisor will be updated as a background task. The task will have JobType UpdateHypervisor. When the task is complete it will redirect to GetHypervisor. The job&#39;s Parameters will contain properties: * _Name_ - Name of the hypervisor being updated. * _Id_ - ID of the hypervisor being updated. | [default to false]

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


## HypervisorsPatchHypervisorResourcePool

> HypervisorsPatchHypervisorResourcePool(ctx, nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).EditHypervisorResourcePoolRequestModel(editHypervisorResourcePoolRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor.
    poolId := "poolId_example" // string | Name or ID of the resource pool.
    editHypervisorResourcePoolRequestModel := *openapiclient.NewEditHypervisorResourcePoolRequestModel(openapiclient.HypervisorConnectionType("Unknown")) // EditHypervisorResourcePoolRequestModel | Details of the resource pool to update.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the resource pool will be updated as a background task. The task will have JobType UpdateHypervisorResourcePool. When the task is complete it will redirect to GetHypervisorResourcePool. The job's Parameters will contain properties: * _Name_ - Name of the resource pool being updated. * _Id_ - ID of the resource pool being updated. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsPatchHypervisorResourcePool(context.Background(), nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).EditHypervisorResourcePoolRequestModel(editHypervisorResourcePoolRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsPatchHypervisorResourcePool``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsPatchHypervisorResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **editHypervisorResourcePoolRequestModel** | [**EditHypervisorResourcePoolRequestModel**](EditHypervisorResourcePoolRequestModel.md) | Details of the resource pool to update. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the resource pool will be updated as a background task. The task will have JobType UpdateHypervisorResourcePool. When the task is complete it will redirect to GetHypervisorResourcePool. The job&#39;s Parameters will contain properties: * _Name_ - Name of the resource pool being updated. * _Id_ - ID of the resource pool being updated. | [default to false]

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


## HypervisorsPutBrokerHypervisorConnection

> HypervisorsPutBrokerHypervisorConnection(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).EditBrokerHypervisorConnectionRequestModel(editBrokerHypervisorConnectionRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Hypervisor connection name or id. 
    editBrokerHypervisorConnectionRequestModel := *openapiclient.NewEditBrokerHypervisorConnectionRequestModel() // EditBrokerHypervisorConnectionRequestModel | EditBrokerHypervisorConnectionRequestModel or derived objects. 
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If async calling.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsPutBrokerHypervisorConnection(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).EditBrokerHypervisorConnectionRequestModel(editBrokerHypervisorConnectionRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsPutBrokerHypervisorConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Hypervisor connection name or id.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsPutBrokerHypervisorConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **editBrokerHypervisorConnectionRequestModel** | [**EditBrokerHypervisorConnectionRequestModel**](EditBrokerHypervisorConnectionRequestModel.md) | EditBrokerHypervisorConnectionRequestModel or derived objects.  | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If async calling.  | [default to false]

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


## HypervisorsResetBrokerHypervisorConnection

> HypervisorsResetBrokerHypervisorConnection(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Hypervisor connection name or id. 
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If async calling.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsResetBrokerHypervisorConnection(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsResetBrokerHypervisorConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Hypervisor connection name or id.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsResetBrokerHypervisorConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If async calling.  | [default to false]

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


## HypervisorsTestHypervisor

> HypervisorTestResponseModel HypervisorsTestHypervisor(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor to test.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the tests will run as a background task. This is recommended as the tests may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType TestHypervisor. When the task is complete it will redirect to GetJobResults. The job's Parameters will contain properties:  * _Id_ - ID of the hypervisor being tested, * _Name_ - Name of the hypervisor being tested. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsTestHypervisor(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsTestHypervisor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTestHypervisor`: HypervisorTestResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsTestHypervisor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor to test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTestHypervisorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the tests will run as a background task. This is recommended as the tests may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType TestHypervisor. When the task is complete it will redirect to GetJobResults. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the hypervisor being tested, * _Name_ - Name of the hypervisor being tested. | [default to false]

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


## HypervisorsTestHypervisorResourcePool

> HypervisorResourcePoolTestResponseModel HypervisorsTestHypervisorResourcePool(ctx, nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or ID of the hypervisor containing the resource pool.
    poolId := "poolId_example" // string | Name or ID of the resource pool to test.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the tests will run as a background task. This is recommended as the tests may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType TestHypervisorResourcePool. When the task is complete it will redirect to GetJobResults. The job's Parameters will contain properties:  * _Id_ - ID of the resource pool being tested, * _Name_ - Name of the resource pool being tested. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsTestHypervisorResourcePool(context.Background(), nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsTestHypervisorResourcePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsTestHypervisorResourcePool`: HypervisorResourcePoolTestResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsTestHypervisorResourcePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or ID of the hypervisor containing the resource pool. | 
**poolId** | **string** | Name or ID of the resource pool to test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsTestHypervisorResourcePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the tests will run as a background task. This is recommended as the tests may take some time to run, and if it exceeds 90 seconds the request may time out. The task will have JobType TestHypervisorResourcePool. When the task is complete it will redirect to GetJobResults. The job&#39;s Parameters will contain properties:  * _Id_ - ID of the resource pool being tested, * _Name_ - Name of the resource pool being tested. | [default to false]

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


## HypervisorsValidateHypervisorResourcePoolResource

> HypervisorResourceValidationResponseModel HypervisorsValidateHypervisorResourcePoolResource(ctx, nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).HypervisorResourceValidationRequestModel(hypervisorResourceValidationRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    nameOrId := "nameOrId_example" // string | Name or Id of the hypervisor.
    poolId := "poolId_example" // string | Name or Id of the resource pool.
    hypervisorResourceValidationRequestModel := *openapiclient.NewHypervisorResourceValidationRequestModel("Path_example") // HypervisorResourceValidationRequestModel | More details attached to this validation.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If async execute. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsValidateHypervisorResourcePoolResource(context.Background(), nameOrId, poolId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).HypervisorResourceValidationRequestModel(hypervisorResourceValidationRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsValidateHypervisorResourcePoolResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsValidateHypervisorResourcePoolResource`: HypervisorResourceValidationResponseModel
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsValidateHypervisorResourcePoolResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Name or Id of the hypervisor. | 
**poolId** | **string** | Name or Id of the resource pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsValidateHypervisorResourcePoolResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **hypervisorResourceValidationRequestModel** | [**HypervisorResourceValidationRequestModel**](HypervisorResourceValidationRequestModel.md) | More details attached to this validation. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If async execute. | [default to false]

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


## HypervisorsValidateVmPath

> HypervisorVmValidationResponseModelCollection HypervisorsValidateVmPath(ctx).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).HypervisorVmValidationRequestModel(hypervisorVmValidationRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

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
    citrixCustomerId := "CitrixOnPremises" // string | Citrix Customer ID. Default is 'CitrixOnPremises'
    citrixInstanceId := "citrixInstanceId_example" // string | Citrix Instance (Site) ID.
    hypervisorVmValidationRequestModel := *openapiclient.NewHypervisorVmValidationRequestModel() // HypervisorVmValidationRequestModel | A list of Vm path.
    userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
    authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
    accept := "application/json" // string | Must accept application/json. (optional)
    citrixLocale := "en-US" // string | Locale of the request. (optional)
    async := true // bool | If `true`, the resource pool will be updated as a background task. The task will have JobType HypervisorVmValidation. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HypervisorsAPIsDAAS.HypervisorsValidateVmPath(context.Background()).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).HypervisorVmValidationRequestModel(hypervisorVmValidationRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPIsDAAS.HypervisorsValidateVmPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HypervisorsValidateVmPath`: HypervisorVmValidationResponseModelCollection
    fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPIsDAAS.HypervisorsValidateVmPath`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorsValidateVmPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 
 **hypervisorVmValidationRequestModel** | [**HypervisorVmValidationRequestModel**](HypervisorVmValidationRequestModel.md) | A list of Vm path. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the resource pool will be updated as a background task. The task will have JobType HypervisorVmValidation. | [default to false]

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

