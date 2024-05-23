# CreateImageSchemeRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceOfferingPath** | Pointer to **NullableString** | The hypervisor resource path of the Cloud service offering to use when creating machines. | [optional] 
**MachineProfile** | Pointer to **NullableString** | The path in the resource pool to the virtual machine template that will be used. This identifies the VM template to be used and the default values for the tags, virtual machine size, boot diagnostics, host cache property of OS disk, accelerated networking and availability zone. This must be a path to a Virtual machine or Template item in the resource pool to which the Image Version Specification is associated. | [optional] 
**CpuCount** | Pointer to **NullableInt32** | The number of processors that virtual machines created from the image preparing should use. | [optional] 
**MemoryMB** | Pointer to **NullableInt32** | The maximum amount of memory that virtual machines created from the image preparing should use. | [optional] 
**NetworkMapping** | Pointer to [**[]NetworkMapRequestModel**](NetworkMapRequestModel.md) | Specifies how the attached NICs are mapped to networks.  If this parameter is omitted, provisioned VMs are created with a single NIC, which is mapped to the default network in the hypervisor resource pool.  If this parameter is supplied, machines are created with the number of NICs specified in the map, and each NIC is attached to the specified network. | [optional] 
**CustomProperties** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The properties of the image version specification that are specific to the target hosting infrastructure. | [optional] 

## Methods

### NewCreateImageSchemeRequestModel

`func NewCreateImageSchemeRequestModel() *CreateImageSchemeRequestModel`

NewCreateImageSchemeRequestModel instantiates a new CreateImageSchemeRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateImageSchemeRequestModelWithDefaults

`func NewCreateImageSchemeRequestModelWithDefaults() *CreateImageSchemeRequestModel`

NewCreateImageSchemeRequestModelWithDefaults instantiates a new CreateImageSchemeRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceOfferingPath

`func (o *CreateImageSchemeRequestModel) GetServiceOfferingPath() string`

GetServiceOfferingPath returns the ServiceOfferingPath field if non-nil, zero value otherwise.

### GetServiceOfferingPathOk

`func (o *CreateImageSchemeRequestModel) GetServiceOfferingPathOk() (*string, bool)`

GetServiceOfferingPathOk returns a tuple with the ServiceOfferingPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceOfferingPath

`func (o *CreateImageSchemeRequestModel) SetServiceOfferingPath(v string)`

SetServiceOfferingPath sets ServiceOfferingPath field to given value.

### HasServiceOfferingPath

`func (o *CreateImageSchemeRequestModel) HasServiceOfferingPath() bool`

HasServiceOfferingPath returns a boolean if a field has been set.

### SetServiceOfferingPathNil

`func (o *CreateImageSchemeRequestModel) SetServiceOfferingPathNil(b bool)`

 SetServiceOfferingPathNil sets the value for ServiceOfferingPath to be an explicit nil

### UnsetServiceOfferingPath
`func (o *CreateImageSchemeRequestModel) UnsetServiceOfferingPath()`

UnsetServiceOfferingPath ensures that no value is present for ServiceOfferingPath, not even an explicit nil
### GetMachineProfile

`func (o *CreateImageSchemeRequestModel) GetMachineProfile() string`

GetMachineProfile returns the MachineProfile field if non-nil, zero value otherwise.

### GetMachineProfileOk

`func (o *CreateImageSchemeRequestModel) GetMachineProfileOk() (*string, bool)`

GetMachineProfileOk returns a tuple with the MachineProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineProfile

`func (o *CreateImageSchemeRequestModel) SetMachineProfile(v string)`

SetMachineProfile sets MachineProfile field to given value.

### HasMachineProfile

`func (o *CreateImageSchemeRequestModel) HasMachineProfile() bool`

HasMachineProfile returns a boolean if a field has been set.

### SetMachineProfileNil

`func (o *CreateImageSchemeRequestModel) SetMachineProfileNil(b bool)`

 SetMachineProfileNil sets the value for MachineProfile to be an explicit nil

### UnsetMachineProfile
`func (o *CreateImageSchemeRequestModel) UnsetMachineProfile()`

UnsetMachineProfile ensures that no value is present for MachineProfile, not even an explicit nil
### GetCpuCount

`func (o *CreateImageSchemeRequestModel) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *CreateImageSchemeRequestModel) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *CreateImageSchemeRequestModel) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *CreateImageSchemeRequestModel) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### SetCpuCountNil

`func (o *CreateImageSchemeRequestModel) SetCpuCountNil(b bool)`

 SetCpuCountNil sets the value for CpuCount to be an explicit nil

### UnsetCpuCount
`func (o *CreateImageSchemeRequestModel) UnsetCpuCount()`

UnsetCpuCount ensures that no value is present for CpuCount, not even an explicit nil
### GetMemoryMB

`func (o *CreateImageSchemeRequestModel) GetMemoryMB() int32`

GetMemoryMB returns the MemoryMB field if non-nil, zero value otherwise.

### GetMemoryMBOk

`func (o *CreateImageSchemeRequestModel) GetMemoryMBOk() (*int32, bool)`

GetMemoryMBOk returns a tuple with the MemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMB

`func (o *CreateImageSchemeRequestModel) SetMemoryMB(v int32)`

SetMemoryMB sets MemoryMB field to given value.

### HasMemoryMB

`func (o *CreateImageSchemeRequestModel) HasMemoryMB() bool`

HasMemoryMB returns a boolean if a field has been set.

### SetMemoryMBNil

`func (o *CreateImageSchemeRequestModel) SetMemoryMBNil(b bool)`

 SetMemoryMBNil sets the value for MemoryMB to be an explicit nil

### UnsetMemoryMB
`func (o *CreateImageSchemeRequestModel) UnsetMemoryMB()`

UnsetMemoryMB ensures that no value is present for MemoryMB, not even an explicit nil
### GetNetworkMapping

`func (o *CreateImageSchemeRequestModel) GetNetworkMapping() []NetworkMapRequestModel`

GetNetworkMapping returns the NetworkMapping field if non-nil, zero value otherwise.

### GetNetworkMappingOk

`func (o *CreateImageSchemeRequestModel) GetNetworkMappingOk() (*[]NetworkMapRequestModel, bool)`

GetNetworkMappingOk returns a tuple with the NetworkMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMapping

`func (o *CreateImageSchemeRequestModel) SetNetworkMapping(v []NetworkMapRequestModel)`

SetNetworkMapping sets NetworkMapping field to given value.

### HasNetworkMapping

`func (o *CreateImageSchemeRequestModel) HasNetworkMapping() bool`

HasNetworkMapping returns a boolean if a field has been set.

### SetNetworkMappingNil

`func (o *CreateImageSchemeRequestModel) SetNetworkMappingNil(b bool)`

 SetNetworkMappingNil sets the value for NetworkMapping to be an explicit nil

### UnsetNetworkMapping
`func (o *CreateImageSchemeRequestModel) UnsetNetworkMapping()`

UnsetNetworkMapping ensures that no value is present for NetworkMapping, not even an explicit nil
### GetCustomProperties

`func (o *CreateImageSchemeRequestModel) GetCustomProperties() []NameValueStringPairModel`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *CreateImageSchemeRequestModel) GetCustomPropertiesOk() (*[]NameValueStringPairModel, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *CreateImageSchemeRequestModel) SetCustomProperties(v []NameValueStringPairModel)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *CreateImageSchemeRequestModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *CreateImageSchemeRequestModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *CreateImageSchemeRequestModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


