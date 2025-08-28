# \ProvisioningSchemesAPIsDAAS

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProvisioningSchemesGetMaintenanceCycles**](ProvisioningSchemesAPIsDAAS.md#ProvisioningSchemesGetMaintenanceCycles) | **Get** /provisioningSchemes/{nameOrId}/maintenanceCycles | 
[**ProvisioningSchemesGetProvisioiningSchemeConfigurations**](ProvisioningSchemesAPIsDAAS.md#ProvisioningSchemesGetProvisioiningSchemeConfigurations) | **Get** /provisioningSchemes/{nameOrId}/configurations | Get provisioning scheme configurations.
[**ProvisioningSchemesGetVMConfigurationResultantSet**](ProvisioningSchemesAPIsDAAS.md#ProvisioningSchemesGetVMConfigurationResultantSet) | **Get** /provisioningSchemes/{nameOrId}/provisionedVirtualMachines/{vmSid}/configurationResultantSet | Provides the ability get the resultant configuration properties for virtual machine created using Machine Creation Services. This merges properties at the provisioning scheme level with those set on a machine with Set-ProvVM specifically.



## ProvisioningSchemesGetMaintenanceCycles

> MaintenanceCycleResponseModelCollection ProvisioningSchemesGetMaintenanceCycles(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()



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
	nameOrId := "nameOrId_example" // string | 
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningSchemesAPIsDAAS.ProvisioningSchemesGetMaintenanceCycles(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningSchemesAPIsDAAS.ProvisioningSchemesGetMaintenanceCycles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvisioningSchemesGetMaintenanceCycles`: MaintenanceCycleResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningSchemesAPIsDAAS.ProvisioningSchemesGetMaintenanceCycles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningSchemesGetMaintenanceCyclesRequest struct via the builder pattern


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

[**MaintenanceCycleResponseModelCollection**](MaintenanceCycleResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisioningSchemesGetProvisioiningSchemeConfigurations

> ProvisioningSchemeConfigurationResponseModelCollection ProvisioningSchemesGetProvisioiningSchemeConfigurations(ctx, nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Version(version).Async(async).Execute()

Get provisioning scheme configurations.

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
	nameOrId := "nameOrId_example" // string | Provisioning scheme name or id. 
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	version := int32(56) // int32 | The version of provisioning scheme configuration.  (optional) (default to 0)
	async := true // bool | If `true`, the get of provisioning scheme will run as a background task. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningSchemesAPIsDAAS.ProvisioningSchemesGetProvisioiningSchemeConfigurations(context.Background(), nameOrId).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Version(version).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningSchemesAPIsDAAS.ProvisioningSchemesGetProvisioiningSchemeConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvisioningSchemesGetProvisioiningSchemeConfigurations`: ProvisioningSchemeConfigurationResponseModelCollection
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningSchemesAPIsDAAS.ProvisioningSchemesGetProvisioiningSchemeConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Provisioning scheme name or id.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningSchemesGetProvisioiningSchemeConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 

 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **version** | **int32** | The version of provisioning scheme configuration.  | [default to 0]
 **async** | **bool** | If &#x60;true&#x60;, the get of provisioning scheme will run as a background task. | [default to false]

### Return type

[**ProvisioningSchemeConfigurationResponseModelCollection**](ProvisioningSchemeConfigurationResponseModelCollection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisioningSchemesGetVMConfigurationResultantSet

> ProvisionedVirtualMachineConfigurationResultantSetResponseModel ProvisioningSchemesGetVMConfigurationResultantSet(ctx, nameOrId, vmSid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()

Provides the ability get the resultant configuration properties for virtual machine created using Machine Creation Services. This merges properties at the provisioning scheme level with those set on a machine with Set-ProvVM specifically.

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
	nameOrId := "nameOrId_example" // string | Provisioning scheme name or id.
	vmSid := "vmSid_example" // string | Virtual machine SID.
	userAgent := "Mozilla/5.0" // string | User Agent type of the request. (optional)
	authorization := "authorization_example" // string | Citrix authorization header: CWSAuth Bearer={token} (optional)
	citrixTransactionId := "citrixTransactionId_example" // string | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. (optional)
	accept := "application/json" // string | Must accept application/json. (optional)
	citrixLocale := "en-US" // string | Locale of the request. (optional)
	async := true // bool | If `true`, the get of VM configuration set will run as a background task. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningSchemesAPIsDAAS.ProvisioningSchemesGetVMConfigurationResultantSet(context.Background(), nameOrId, vmSid).CitrixCustomerId(citrixCustomerId).CitrixInstanceId(citrixInstanceId).UserAgent(userAgent).Authorization(authorization).CitrixTransactionId(citrixTransactionId).Accept(accept).CitrixLocale(citrixLocale).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningSchemesAPIsDAAS.ProvisioningSchemesGetVMConfigurationResultantSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvisioningSchemesGetVMConfigurationResultantSet`: ProvisionedVirtualMachineConfigurationResultantSetResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningSchemesAPIsDAAS.ProvisioningSchemesGetVMConfigurationResultantSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** | Provisioning scheme name or id. | 
**vmSid** | **string** | Virtual machine SID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningSchemesGetVMConfigurationResultantSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **citrixCustomerId** | **string** | Citrix Customer ID. Default is &#39;CitrixOnPremises&#39; | 
 **citrixInstanceId** | **string** | Citrix Instance (Site) ID. | 


 **userAgent** | **string** | User Agent type of the request. | 
 **authorization** | **string** | Citrix authorization header: CWSAuth Bearer&#x3D;{token} | 
 **citrixTransactionId** | **string** | Transaction ID that will be used to track this request. If not provided, a new GUID will be generated and returned. | 
 **accept** | **string** | Must accept application/json. | 
 **citrixLocale** | **string** | Locale of the request. | 
 **async** | **bool** | If &#x60;true&#x60;, the get of VM configuration set will run as a background task. | [default to false]

### Return type

[**ProvisionedVirtualMachineConfigurationResultantSetResponseModel**](ProvisionedVirtualMachineConfigurationResultantSetResponseModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

