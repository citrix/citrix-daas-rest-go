# UpdateProvisionedVirtualMachineRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuCount** | Pointer to **NullableInt32** | The number of processors. Not supported for cloud hypervisors | [optional] 
**MemoryMB** | Pointer to **NullableInt32** | The maximum amount of memory in megabytes. Not supported for cloud hypervisors | [optional] 
**ServiceOfferingPath** | Pointer to **NullableString** | The hypervisor resource path of the Cloud service offering to use when creating machines. | [optional] 
**MachineProfilePath** | Pointer to **NullableString** | The path in the resource pool to the virtual machine template that will be used. This identifies the VM template to be used and the default values for the tags, virtual machine size, boot diagnostics, host cache property of OS disk, accelerated networking and availability zone. This must be a path to a Virtual machine or Template item in the resource pool to which the Machine Catalog is associated. | [optional] 
**CustomProperties** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The properties of the provisioning scheme that are specific to the target hosting infrastructure. | [optional] 
**RevertToProvSchemeConfiguration** | Pointer to **NullableBool** | If supplied, all existing configuration is cleared for the given machine. It is mutually exclusive with parameters that apply configuration settings | [optional] 

## Methods

### NewUpdateProvisionedVirtualMachineRequestModel

`func NewUpdateProvisionedVirtualMachineRequestModel() *UpdateProvisionedVirtualMachineRequestModel`

NewUpdateProvisionedVirtualMachineRequestModel instantiates a new UpdateProvisionedVirtualMachineRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProvisionedVirtualMachineRequestModelWithDefaults

`func NewUpdateProvisionedVirtualMachineRequestModelWithDefaults() *UpdateProvisionedVirtualMachineRequestModel`

NewUpdateProvisionedVirtualMachineRequestModelWithDefaults instantiates a new UpdateProvisionedVirtualMachineRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuCount

`func (o *UpdateProvisionedVirtualMachineRequestModel) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *UpdateProvisionedVirtualMachineRequestModel) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *UpdateProvisionedVirtualMachineRequestModel) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *UpdateProvisionedVirtualMachineRequestModel) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### SetCpuCountNil

`func (o *UpdateProvisionedVirtualMachineRequestModel) SetCpuCountNil(b bool)`

 SetCpuCountNil sets the value for CpuCount to be an explicit nil

### UnsetCpuCount
`func (o *UpdateProvisionedVirtualMachineRequestModel) UnsetCpuCount()`

UnsetCpuCount ensures that no value is present for CpuCount, not even an explicit nil
### GetMemoryMB

`func (o *UpdateProvisionedVirtualMachineRequestModel) GetMemoryMB() int32`

GetMemoryMB returns the MemoryMB field if non-nil, zero value otherwise.

### GetMemoryMBOk

`func (o *UpdateProvisionedVirtualMachineRequestModel) GetMemoryMBOk() (*int32, bool)`

GetMemoryMBOk returns a tuple with the MemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMB

`func (o *UpdateProvisionedVirtualMachineRequestModel) SetMemoryMB(v int32)`

SetMemoryMB sets MemoryMB field to given value.

### HasMemoryMB

`func (o *UpdateProvisionedVirtualMachineRequestModel) HasMemoryMB() bool`

HasMemoryMB returns a boolean if a field has been set.

### SetMemoryMBNil

`func (o *UpdateProvisionedVirtualMachineRequestModel) SetMemoryMBNil(b bool)`

 SetMemoryMBNil sets the value for MemoryMB to be an explicit nil

### UnsetMemoryMB
`func (o *UpdateProvisionedVirtualMachineRequestModel) UnsetMemoryMB()`

UnsetMemoryMB ensures that no value is present for MemoryMB, not even an explicit nil
### GetServiceOfferingPath

`func (o *UpdateProvisionedVirtualMachineRequestModel) GetServiceOfferingPath() string`

GetServiceOfferingPath returns the ServiceOfferingPath field if non-nil, zero value otherwise.

### GetServiceOfferingPathOk

`func (o *UpdateProvisionedVirtualMachineRequestModel) GetServiceOfferingPathOk() (*string, bool)`

GetServiceOfferingPathOk returns a tuple with the ServiceOfferingPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceOfferingPath

`func (o *UpdateProvisionedVirtualMachineRequestModel) SetServiceOfferingPath(v string)`

SetServiceOfferingPath sets ServiceOfferingPath field to given value.

### HasServiceOfferingPath

`func (o *UpdateProvisionedVirtualMachineRequestModel) HasServiceOfferingPath() bool`

HasServiceOfferingPath returns a boolean if a field has been set.

### SetServiceOfferingPathNil

`func (o *UpdateProvisionedVirtualMachineRequestModel) SetServiceOfferingPathNil(b bool)`

 SetServiceOfferingPathNil sets the value for ServiceOfferingPath to be an explicit nil

### UnsetServiceOfferingPath
`func (o *UpdateProvisionedVirtualMachineRequestModel) UnsetServiceOfferingPath()`

UnsetServiceOfferingPath ensures that no value is present for ServiceOfferingPath, not even an explicit nil
### GetMachineProfilePath

`func (o *UpdateProvisionedVirtualMachineRequestModel) GetMachineProfilePath() string`

GetMachineProfilePath returns the MachineProfilePath field if non-nil, zero value otherwise.

### GetMachineProfilePathOk

`func (o *UpdateProvisionedVirtualMachineRequestModel) GetMachineProfilePathOk() (*string, bool)`

GetMachineProfilePathOk returns a tuple with the MachineProfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineProfilePath

`func (o *UpdateProvisionedVirtualMachineRequestModel) SetMachineProfilePath(v string)`

SetMachineProfilePath sets MachineProfilePath field to given value.

### HasMachineProfilePath

`func (o *UpdateProvisionedVirtualMachineRequestModel) HasMachineProfilePath() bool`

HasMachineProfilePath returns a boolean if a field has been set.

### SetMachineProfilePathNil

`func (o *UpdateProvisionedVirtualMachineRequestModel) SetMachineProfilePathNil(b bool)`

 SetMachineProfilePathNil sets the value for MachineProfilePath to be an explicit nil

### UnsetMachineProfilePath
`func (o *UpdateProvisionedVirtualMachineRequestModel) UnsetMachineProfilePath()`

UnsetMachineProfilePath ensures that no value is present for MachineProfilePath, not even an explicit nil
### GetCustomProperties

`func (o *UpdateProvisionedVirtualMachineRequestModel) GetCustomProperties() []NameValueStringPairModel`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *UpdateProvisionedVirtualMachineRequestModel) GetCustomPropertiesOk() (*[]NameValueStringPairModel, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *UpdateProvisionedVirtualMachineRequestModel) SetCustomProperties(v []NameValueStringPairModel)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *UpdateProvisionedVirtualMachineRequestModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *UpdateProvisionedVirtualMachineRequestModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *UpdateProvisionedVirtualMachineRequestModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
### GetRevertToProvSchemeConfiguration

`func (o *UpdateProvisionedVirtualMachineRequestModel) GetRevertToProvSchemeConfiguration() bool`

GetRevertToProvSchemeConfiguration returns the RevertToProvSchemeConfiguration field if non-nil, zero value otherwise.

### GetRevertToProvSchemeConfigurationOk

`func (o *UpdateProvisionedVirtualMachineRequestModel) GetRevertToProvSchemeConfigurationOk() (*bool, bool)`

GetRevertToProvSchemeConfigurationOk returns a tuple with the RevertToProvSchemeConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevertToProvSchemeConfiguration

`func (o *UpdateProvisionedVirtualMachineRequestModel) SetRevertToProvSchemeConfiguration(v bool)`

SetRevertToProvSchemeConfiguration sets RevertToProvSchemeConfiguration field to given value.

### HasRevertToProvSchemeConfiguration

`func (o *UpdateProvisionedVirtualMachineRequestModel) HasRevertToProvSchemeConfiguration() bool`

HasRevertToProvSchemeConfiguration returns a boolean if a field has been set.

### SetRevertToProvSchemeConfigurationNil

`func (o *UpdateProvisionedVirtualMachineRequestModel) SetRevertToProvSchemeConfigurationNil(b bool)`

 SetRevertToProvSchemeConfigurationNil sets the value for RevertToProvSchemeConfiguration to be an explicit nil

### UnsetRevertToProvSchemeConfiguration
`func (o *UpdateProvisionedVirtualMachineRequestModel) UnsetRevertToProvSchemeConfiguration()`

UnsetRevertToProvSchemeConfiguration ensures that no value is present for RevertToProvSchemeConfiguration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


