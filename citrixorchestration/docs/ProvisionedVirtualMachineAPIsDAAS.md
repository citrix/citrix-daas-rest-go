# \ProvisionedVirtualMachineAPIsDAAS

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProvisionedVirtualMachineApplyProvisionedVirtualMachineConfigurationUpdate**](ProvisionedVirtualMachineAPIsDAAS.md#ProvisionedVirtualMachineApplyProvisionedVirtualMachineConfigurationUpdate) | **Post** /provisionedVirtualMachines/{vmSid}/$applyUpdate | Performs an immediate update of the configuration for a specific provisioned virtual machine.
[**ProvisionedVirtualMachineGetMaintenanceCycleVMOperationJobs**](ProvisionedVirtualMachineAPIsDAAS.md#ProvisionedVirtualMachineGetMaintenanceCycleVMOperationJobs) | **Get** /provisionedVirtualMachines/{vmSid}/operationJobs | Get the maintenance cycle vm operation jobs for a machine.
[**ProvisionedVirtualMachineResetProvisionedVirtualMachineOSDisk**](ProvisionedVirtualMachineAPIsDAAS.md#ProvisionedVirtualMachineResetProvisionedVirtualMachineOSDisk) | **Post** /provisionedVirtualMachines/{vmSid}/$resetProvVmOSDisk | Performs an immediate reset of the OS disk for a specific provisioned virtual machine.
[**ProvisionedVirtualMachineUpdateProvisionedVirtualMachine**](ProvisionedVirtualMachineAPIsDAAS.md#ProvisionedVirtualMachineUpdateProvisionedVirtualMachine) | **Patch** /provisionedVirtualMachines/{vmSid} | Changes the configuration for a specific provisioned virtual machine.



## ProvisionedVirtualMachineApplyProvisionedVirtualMachineConfigurationUpdate

> ProvisionedVirtualMachineApplyProvisionedVirtualMachineConfigurationUpdate(ctx, vmSid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Performs an immediate update of the configuration for a specific provisioned virtual machine.



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
	vmSid := "vmSid_example" // string | Virtual Machine Sid.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the update operation will be created within a background task. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProvisionedVirtualMachineAPIsDAAS.ProvisionedVirtualMachineApplyProvisionedVirtualMachineConfigurationUpdate(context.Background(), vmSid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisionedVirtualMachineAPIsDAAS.ProvisionedVirtualMachineApplyProvisionedVirtualMachineConfigurationUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmSid** | **string** | Virtual Machine Sid. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisionedVirtualMachineApplyProvisionedVirtualMachineConfigurationUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the update operation will be created within a background task. | [default to false]

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


## ProvisionedVirtualMachineGetMaintenanceCycleVMOperationJobs

> MaintenanceCycleVMOperationJobResponseModelCollection ProvisionedVirtualMachineGetMaintenanceCycleVMOperationJobs(ctx, vmSid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Get the maintenance cycle vm operation jobs for a machine.

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
	vmSid := "vmSid_example" // string | Virtual machine SID.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisionedVirtualMachineAPIsDAAS.ProvisionedVirtualMachineGetMaintenanceCycleVMOperationJobs(context.Background(), vmSid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisionedVirtualMachineAPIsDAAS.ProvisionedVirtualMachineGetMaintenanceCycleVMOperationJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvisionedVirtualMachineGetMaintenanceCycleVMOperationJobs`: MaintenanceCycleVMOperationJobResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `ProvisionedVirtualMachineAPIsDAAS.ProvisionedVirtualMachineGetMaintenanceCycleVMOperationJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmSid** | **string** | Virtual machine SID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisionedVirtualMachineGetMaintenanceCycleVMOperationJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** |  | [default to false]

### Return type

[**MaintenanceCycleVMOperationJobResponseModelCollection**](MaintenanceCycleVMOperationJobResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisionedVirtualMachineResetProvisionedVirtualMachineOSDisk

> ProvisionedVirtualMachineResetProvisionedVirtualMachineOSDisk(ctx, vmSid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Performs an immediate reset of the OS disk for a specific provisioned virtual machine.



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
	vmSid := "vmSid_example" // string | Virtual Machine Sid.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the update operation will be created within a background task. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProvisionedVirtualMachineAPIsDAAS.ProvisionedVirtualMachineResetProvisionedVirtualMachineOSDisk(context.Background(), vmSid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisionedVirtualMachineAPIsDAAS.ProvisionedVirtualMachineResetProvisionedVirtualMachineOSDisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmSid** | **string** | Virtual Machine Sid. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisionedVirtualMachineResetProvisionedVirtualMachineOSDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the update operation will be created within a background task. | [default to false]

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


## ProvisionedVirtualMachineUpdateProvisionedVirtualMachine

> ProvisionedVirtualMachineUpdateProvisionedVirtualMachine(ctx, vmSid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateProvisionedVirtualMachineRequestModel(updateProvisionedVirtualMachineRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Changes the configuration for a specific provisioned virtual machine.



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
	vmSid := "vmSid_example" // string | Virtual Machine Sid.
	updateProvisionedVirtualMachineRequestModel := *openapiclient.NewUpdateProvisionedVirtualMachineRequestModel() // UpdateProvisionedVirtualMachineRequestModel | Properties of the provisioned virtual machine to update.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the update operation will be created within a background task. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProvisionedVirtualMachineAPIsDAAS.ProvisionedVirtualMachineUpdateProvisionedVirtualMachine(context.Background(), vmSid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UpdateProvisionedVirtualMachineRequestModel(updateProvisionedVirtualMachineRequestModel).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisionedVirtualMachineAPIsDAAS.ProvisionedVirtualMachineUpdateProvisionedVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmSid** | **string** | Virtual Machine Sid. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisionedVirtualMachineUpdateProvisionedVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **updateProvisionedVirtualMachineRequestModel** | [**UpdateProvisionedVirtualMachineRequestModel**](UpdateProvisionedVirtualMachineRequestModel.md) | Properties of the provisioned virtual machine to update. | 
 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the update operation will be created within a background task. | [default to false]

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

