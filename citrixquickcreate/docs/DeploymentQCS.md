# \DeploymentQCS

All URIs are relative to *https://api.cloud.com/quickcreateservice*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMachineAsync**](DeploymentQCS.md#AddMachineAsync) | **Post** /{customerId}/deployments/{deploymentId}/$addMachines | Adds machine to deployment
[**DeleteDeploymentAsync**](DeploymentQCS.md#DeleteDeploymentAsync) | **Delete** /{customerId}/deployments/{deploymentId} | Delete deployment
[**GetDeploymentAsync**](DeploymentQCS.md#GetDeploymentAsync) | **Get** /{customerId}/deployments/{deploymentId} | Gets deployment with vdas
[**GetDeploymentsAsync**](DeploymentQCS.md#GetDeploymentsAsync) | **Get** /{customerId}/deployments | Gets deployments
[**InitiateDeleteDeploymentAsync**](DeploymentQCS.md#InitiateDeleteDeploymentAsync) | **Post** /{customerId}/deployments/{deploymentId}/$initiateDelete | Initiate delete deployment
[**InitiateDeploymentAsync**](DeploymentQCS.md#InitiateDeploymentAsync) | **Post** /{customerId}/deployments/$initiate | Initiates deployment
[**InitiateRemoveMachineAsync**](DeploymentQCS.md#InitiateRemoveMachineAsync) | **Post** /{customerId}/deployments/{deploymentId}/machines/{machineId}/$initiateDelete | Initiate removal of machine from deployment
[**PatchMachinesAsync**](DeploymentQCS.md#PatchMachinesAsync) | **Patch** /{customerId}/deployments/{deploymentId}/machines | Turn on or off the deployment mode for deployment machines
[**RemoveMachineAsync**](DeploymentQCS.md#RemoveMachineAsync) | **Delete** /{customerId}/deployments/{deploymentId}/machines/{machineId} | Removes machine to deployment
[**RemoveMachinesAsync**](DeploymentQCS.md#RemoveMachinesAsync) | **Post** /{customerId}/deployments/{deploymentId}/machines/$delete | Removes machines from deployment
[**RestartMachineAsync**](DeploymentQCS.md#RestartMachineAsync) | **Post** /{customerId}/deployments/{deploymentId}/machines/{machineId}/$restart | Restart a machine
[**SaveAsImageAsync**](DeploymentQCS.md#SaveAsImageAsync) | **Post** /{customerId}/deployments/{deploymentId}/machines/{machineId}/$saveAsImage | Save image to account
[**UpdateDeploymentImageAsync**](DeploymentQCS.md#UpdateDeploymentImageAsync) | **Post** /{customerId}/deployments/{deploymentId}/image/$update | Update image for a deployment
[**UpdateDeploymentPropertiesAsync**](DeploymentQCS.md#UpdateDeploymentPropertiesAsync) | **Post** /{customerId}/deployments/{deploymentId}/$update | Updates deployment properties
[**UpdateMachineAsync**](DeploymentQCS.md#UpdateMachineAsync) | **Post** /{customerId}/deployments/{deploymentId}/machines/{machineId}/$update | Updates machine in deployment



## AddMachineAsync

> DeploymentTask AddMachineAsync(ctx, customerId, deploymentId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Adds machine to deployment

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
    deploymentId := "deploymentId_example" // string | ID of the deployment
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    body := AddAwsEdcDeploymentMachines(987) // AddAwsEdcDeploymentMachines | Configuration of machine (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentQCS.AddMachineAsync(context.Background(), customerId, deploymentId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentQCS.AddMachineAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddMachineAsync`: DeploymentTask
    fmt.Fprintf(os.Stdout, "Response from `DeploymentQCS.AddMachineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**deploymentId** | **string** | ID of the deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMachineAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixTransactionId** | **string** | The Transaction Id. | 
 **body** | **AddAwsEdcDeploymentMachines** | Configuration of machine | 

### Return type

[**DeploymentTask**](DeploymentTask.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeploymentAsync

> DeploymentTask DeleteDeploymentAsync(ctx, customerId, deploymentId).ForceDelete(forceDelete).CitrixTransactionId(citrixTransactionId).Execute()

Delete deployment

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
    deploymentId := "deploymentId_example" // string | ID of the deployment
    forceDelete := true // bool | Force delete deployment (optional) (default to false)
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentQCS.DeleteDeploymentAsync(context.Background(), customerId, deploymentId).ForceDelete(forceDelete).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentQCS.DeleteDeploymentAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDeploymentAsync`: DeploymentTask
    fmt.Fprintf(os.Stdout, "Response from `DeploymentQCS.DeleteDeploymentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**deploymentId** | **string** | ID of the deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forceDelete** | **bool** | Force delete deployment | [default to false]
 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**DeploymentTask**](DeploymentTask.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentAsync

> AwsEdcDeployment GetDeploymentAsync(ctx, customerId, deploymentId).CitrixTransactionId(citrixTransactionId).Execute()

Gets deployment with vdas

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
    deploymentId := "deploymentId_example" // string | ID of the deployment
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentQCS.GetDeploymentAsync(context.Background(), customerId, deploymentId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentQCS.GetDeploymentAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentAsync`: AwsEdcDeployment
    fmt.Fprintf(os.Stdout, "Response from `DeploymentQCS.GetDeploymentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**deploymentId** | **string** | ID of the deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**AwsEdcDeployment**](AwsEdcDeployment.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentsAsync

> Deployments GetDeploymentsAsync(ctx, customerId).IncludeVdas(includeVdas).AccountType(accountType).CitrixTransactionId(citrixTransactionId).Execute()

Gets deployments

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
    includeVdas := true // bool | Indicates if the deployment VDAs should be included in the response (optional)
    accountType := openapiclient.AccountType("AWSEDC") // AccountType | Account Type (optional)
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentQCS.GetDeploymentsAsync(context.Background(), customerId).IncludeVdas(includeVdas).AccountType(accountType).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentQCS.GetDeploymentsAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentsAsync`: Deployments
    fmt.Fprintf(os.Stdout, "Response from `DeploymentQCS.GetDeploymentsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeVdas** | **bool** | Indicates if the deployment VDAs should be included in the response | 
 **accountType** | [**AccountType**](AccountType.md) | Account Type | 
 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**Deployments**](Deployments.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateDeleteDeploymentAsync

> DeploymentTask InitiateDeleteDeploymentAsync(ctx, customerId, deploymentId).ForceDelete(forceDelete).CitrixTransactionId(citrixTransactionId).DeploymentIdInitiateDeleteBody(deploymentIdInitiateDeleteBody).Execute()

Initiate delete deployment

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
    deploymentId := "deploymentId_example" // string | ID of the deployment
    forceDelete := true // bool | Force delete deployment (optional) (default to false)
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    deploymentIdInitiateDeleteBody := *openapiclient.NewDeploymentIdInitiateDeleteBody() // DeploymentIdInitiateDeleteBody | Deployment AD credentials (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentQCS.InitiateDeleteDeploymentAsync(context.Background(), customerId, deploymentId).ForceDelete(forceDelete).CitrixTransactionId(citrixTransactionId).DeploymentIdInitiateDeleteBody(deploymentIdInitiateDeleteBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentQCS.InitiateDeleteDeploymentAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitiateDeleteDeploymentAsync`: DeploymentTask
    fmt.Fprintf(os.Stdout, "Response from `DeploymentQCS.InitiateDeleteDeploymentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**deploymentId** | **string** | ID of the deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiInitiateDeleteDeploymentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forceDelete** | **bool** | Force delete deployment | [default to false]
 **citrixTransactionId** | **string** | The Transaction Id. | 
 **deploymentIdInitiateDeleteBody** | [**DeploymentIdInitiateDeleteBody**](DeploymentIdInitiateDeleteBody.md) | Deployment AD credentials | 

### Return type

[**DeploymentTask**](DeploymentTask.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateDeploymentAsync

> AwsEdcDeployment InitiateDeploymentAsync(ctx, customerId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Initiates deployment

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
    body := InitiateAwsEdcDeployment(987) // InitiateAwsEdcDeployment | Deployment configuration (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentQCS.InitiateDeploymentAsync(context.Background(), customerId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentQCS.InitiateDeploymentAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitiateDeploymentAsync`: AwsEdcDeployment
    fmt.Fprintf(os.Stdout, "Response from `DeploymentQCS.InitiateDeploymentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 

### Other Parameters

Other parameters are passed through a pointer to a apiInitiateDeploymentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **citrixTransactionId** | **string** | The Transaction Id. | 
 **body** | **InitiateAwsEdcDeployment** | Deployment configuration | 

### Return type

[**AwsEdcDeployment**](AwsEdcDeployment.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateRemoveMachineAsync

> DeploymentTask InitiateRemoveMachineAsync(ctx, customerId, deploymentId, machineId).ForceDelete(forceDelete).CitrixTransactionId(citrixTransactionId).MachineIdInitiateDeleteBody(machineIdInitiateDeleteBody).Execute()

Initiate removal of machine from deployment

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
    deploymentId := "deploymentId_example" // string | ID of the deployment
    machineId := "machineId_example" // string | ID of the machine
    forceDelete := true // bool | Force delete machine (optional) (default to false)
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    machineIdInitiateDeleteBody := *openapiclient.NewMachineIdInitiateDeleteBody() // MachineIdInitiateDeleteBody | Active Directory credentials for the machine (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentQCS.InitiateRemoveMachineAsync(context.Background(), customerId, deploymentId, machineId).ForceDelete(forceDelete).CitrixTransactionId(citrixTransactionId).MachineIdInitiateDeleteBody(machineIdInitiateDeleteBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentQCS.InitiateRemoveMachineAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitiateRemoveMachineAsync`: DeploymentTask
    fmt.Fprintf(os.Stdout, "Response from `DeploymentQCS.InitiateRemoveMachineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**deploymentId** | **string** | ID of the deployment | 
**machineId** | **string** | ID of the machine | 

### Other Parameters

Other parameters are passed through a pointer to a apiInitiateRemoveMachineAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **forceDelete** | **bool** | Force delete machine | [default to false]
 **citrixTransactionId** | **string** | The Transaction Id. | 
 **machineIdInitiateDeleteBody** | [**MachineIdInitiateDeleteBody**](MachineIdInitiateDeleteBody.md) | Active Directory credentials for the machine | 

### Return type

[**DeploymentTask**](DeploymentTask.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMachinesAsync

> PatchMachinesAsync(ctx, customerId, deploymentId).CitrixTransactionId(citrixTransactionId).DeploymentIdMachinesBody(deploymentIdMachinesBody).Execute()

Turn on or off the deployment mode for deployment machines

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
    deploymentId := "deploymentId_example" // string | ID of the deployment
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    deploymentIdMachinesBody := *openapiclient.NewDeploymentIdMachinesBody(openapiclient.AccountType("AWSEDC"), false, []string{"MachineIds_example"}) // DeploymentIdMachinesBody | Machines to be modified to turn on or off maintenance mode (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeploymentQCS.PatchMachinesAsync(context.Background(), customerId, deploymentId).CitrixTransactionId(citrixTransactionId).DeploymentIdMachinesBody(deploymentIdMachinesBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentQCS.PatchMachinesAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**deploymentId** | **string** | ID of the deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMachinesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixTransactionId** | **string** | The Transaction Id. | 
 **deploymentIdMachinesBody** | [**DeploymentIdMachinesBody**](DeploymentIdMachinesBody.md) | Machines to be modified to turn on or off maintenance mode | 

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


## RemoveMachineAsync

> DeploymentTask RemoveMachineAsync(ctx, customerId, deploymentId, machineId).ForceDelete(forceDelete).CitrixTransactionId(citrixTransactionId).Execute()

Removes machine to deployment

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
    deploymentId := "deploymentId_example" // string | ID of the deployment
    machineId := "machineId_example" // string | ID of the machine
    forceDelete := true // bool | Force delete machine (optional) (default to false)
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentQCS.RemoveMachineAsync(context.Background(), customerId, deploymentId, machineId).ForceDelete(forceDelete).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentQCS.RemoveMachineAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveMachineAsync`: DeploymentTask
    fmt.Fprintf(os.Stdout, "Response from `DeploymentQCS.RemoveMachineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**deploymentId** | **string** | ID of the deployment | 
**machineId** | **string** | ID of the machine | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMachineAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **forceDelete** | **bool** | Force delete machine | [default to false]
 **citrixTransactionId** | **string** | The Transaction Id. | 

### Return type

[**DeploymentTask**](DeploymentTask.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMachinesAsync

> DeploymentTask RemoveMachinesAsync(ctx, customerId, deploymentId).ForceDelete(forceDelete).CitrixTransactionId(citrixTransactionId).MachinesDeleteBody(machinesDeleteBody).Execute()

Removes machines from deployment

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
    deploymentId := "deploymentId_example" // string | ID of the deployment
    forceDelete := true // bool | Force delete machines (optional) (default to false)
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    machinesDeleteBody := *openapiclient.NewMachinesDeleteBody(openapiclient.AccountType("AWSEDC"), []string{"MachineIds_example"}) // MachinesDeleteBody | Configuration object specifying which machines to delete (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentQCS.RemoveMachinesAsync(context.Background(), customerId, deploymentId).ForceDelete(forceDelete).CitrixTransactionId(citrixTransactionId).MachinesDeleteBody(machinesDeleteBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentQCS.RemoveMachinesAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveMachinesAsync`: DeploymentTask
    fmt.Fprintf(os.Stdout, "Response from `DeploymentQCS.RemoveMachinesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**deploymentId** | **string** | ID of the deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMachinesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forceDelete** | **bool** | Force delete machines | [default to false]
 **citrixTransactionId** | **string** | The Transaction Id. | 
 **machinesDeleteBody** | [**MachinesDeleteBody**](MachinesDeleteBody.md) | Configuration object specifying which machines to delete | 

### Return type

[**DeploymentTask**](DeploymentTask.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartMachineAsync

> RestartMachineAsync(ctx, customerId, deploymentId, machineId).CitrixTransactionId(citrixTransactionId).Execute()

Restart a machine

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
    deploymentId := "deploymentId_example" // string | ID of the deployment
    machineId := "machineId_example" // string | ID of the machine
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeploymentQCS.RestartMachineAsync(context.Background(), customerId, deploymentId, machineId).CitrixTransactionId(citrixTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentQCS.RestartMachineAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**deploymentId** | **string** | ID of the deployment | 
**machineId** | **string** | ID of the machine | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartMachineAsyncRequest struct via the builder pattern


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


## SaveAsImageAsync

> DeploymentTask SaveAsImageAsync(ctx, customerId, deploymentId, machineId).CitrixTransactionId(citrixTransactionId).MachineIdSaveAsImageBody(machineIdSaveAsImageBody).Execute()

Save image to account

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
    deploymentId := "deploymentId_example" // string | ID of the deployment
    machineId := "machineId_example" // string | ID of the image builder machine
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    machineIdSaveAsImageBody := *openapiclient.NewMachineIdSaveAsImageBody("AccountId_example", "ImageName_example", "ImageDescription_example") // MachineIdSaveAsImageBody | configurations of the new image (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentQCS.SaveAsImageAsync(context.Background(), customerId, deploymentId, machineId).CitrixTransactionId(citrixTransactionId).MachineIdSaveAsImageBody(machineIdSaveAsImageBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentQCS.SaveAsImageAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveAsImageAsync`: DeploymentTask
    fmt.Fprintf(os.Stdout, "Response from `DeploymentQCS.SaveAsImageAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**deploymentId** | **string** | ID of the deployment | 
**machineId** | **string** | ID of the image builder machine | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveAsImageAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 
 **machineIdSaveAsImageBody** | [**MachineIdSaveAsImageBody**](MachineIdSaveAsImageBody.md) | configurations of the new image | 

### Return type

[**DeploymentTask**](DeploymentTask.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeploymentImageAsync

> DeploymentTask UpdateDeploymentImageAsync(ctx, customerId, deploymentId).CitrixTransactionId(citrixTransactionId).ImageUpdateBody(imageUpdateBody).Execute()

Update image for a deployment

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
    deploymentId := "deploymentId_example" // string | ID of the deployment
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    imageUpdateBody := *openapiclient.NewImageUpdateBody() // ImageUpdateBody | Details of the new image (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentQCS.UpdateDeploymentImageAsync(context.Background(), customerId, deploymentId).CitrixTransactionId(citrixTransactionId).ImageUpdateBody(imageUpdateBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentQCS.UpdateDeploymentImageAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeploymentImageAsync`: DeploymentTask
    fmt.Fprintf(os.Stdout, "Response from `DeploymentQCS.UpdateDeploymentImageAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**deploymentId** | **string** | ID of the deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeploymentImageAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixTransactionId** | **string** | The Transaction Id. | 
 **imageUpdateBody** | [**ImageUpdateBody**](ImageUpdateBody.md) | Details of the new image | 

### Return type

[**DeploymentTask**](DeploymentTask.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeploymentPropertiesAsync

> DeploymentTask UpdateDeploymentPropertiesAsync(ctx, customerId, deploymentId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Updates deployment properties

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
    deploymentId := "deploymentId_example" // string | ID of the deployment
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    body := UpdateAwsEdcDeploymentProperties(987) // UpdateAwsEdcDeploymentProperties | Deployment properties configuration (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentQCS.UpdateDeploymentPropertiesAsync(context.Background(), customerId, deploymentId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentQCS.UpdateDeploymentPropertiesAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeploymentPropertiesAsync`: DeploymentTask
    fmt.Fprintf(os.Stdout, "Response from `DeploymentQCS.UpdateDeploymentPropertiesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**deploymentId** | **string** | ID of the deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeploymentPropertiesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **citrixTransactionId** | **string** | The Transaction Id. | 
 **body** | **UpdateAwsEdcDeploymentProperties** | Deployment properties configuration | 

### Return type

[**DeploymentTask**](DeploymentTask.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMachineAsync

> DeploymentTask UpdateMachineAsync(ctx, customerId, deploymentId, machineId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()

Updates machine in deployment

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
    deploymentId := "deploymentId_example" // string | ID of the deployment
    machineId := "machineId_example" // string | ID of the machine
    citrixTransactionId := "citrixTransactionId_example" // string | The Transaction Id. (optional)
    body := UpdateAwsEdcDeploymentMachine(987) // UpdateAwsEdcDeploymentMachine | Configuration of machine (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentQCS.UpdateMachineAsync(context.Background(), customerId, deploymentId, machineId).CitrixTransactionId(citrixTransactionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentQCS.UpdateMachineAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMachineAsync`: DeploymentTask
    fmt.Fprintf(os.Stdout, "Response from `DeploymentQCS.UpdateMachineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | ID of the customer | 
**deploymentId** | **string** | ID of the deployment | 
**machineId** | **string** | ID of the machine | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMachineAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **citrixTransactionId** | **string** | The Transaction Id. | 
 **body** | **UpdateAwsEdcDeploymentMachine** | Configuration of machine | 

### Return type

[**DeploymentTask**](DeploymentTask.md)

### Authorization

[CWSAuth](../README.md#CWSAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

