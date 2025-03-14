# ProvisionedVirtualMachineConfigurationResultantSetResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuCount** | Pointer to **NullableInt32** | The number of processors that VMs will be created with when using this configuration. | [optional] 
**MemoryMB** | Pointer to **NullableInt32** | The maximum amount of memory that VMs will be created with when using this configuration. | [optional] 
**MachineProfile** | Pointer to [**HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) |  | [optional] 
**ServiceOffering** | Pointer to **NullableString** | Service offering used by Cloud provisioned VMs. | [optional] 
**CustomProperties** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The properties of the provisioning scheme that are specific to the target hosting infrastructure. | [optional] 
**ProvisioningSchemeName** | Pointer to **NullableString** | provisioning scheme name | [optional] 
**ProvisioningSchemeId** | Pointer to **string** | provisioning scheme uid | [optional] 
**VMId** | Pointer to **NullableString** | VM Id | [optional] 
**VMName** | Pointer to **NullableString** | VM Name | [optional] 

## Methods

### NewProvisionedVirtualMachineConfigurationResultantSetResponseModel

`func NewProvisionedVirtualMachineConfigurationResultantSetResponseModel() *ProvisionedVirtualMachineConfigurationResultantSetResponseModel`

NewProvisionedVirtualMachineConfigurationResultantSetResponseModel instantiates a new ProvisionedVirtualMachineConfigurationResultantSetResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionedVirtualMachineConfigurationResultantSetResponseModelWithDefaults

`func NewProvisionedVirtualMachineConfigurationResultantSetResponseModelWithDefaults() *ProvisionedVirtualMachineConfigurationResultantSetResponseModel`

NewProvisionedVirtualMachineConfigurationResultantSetResponseModelWithDefaults instantiates a new ProvisionedVirtualMachineConfigurationResultantSetResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuCount

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### SetCpuCountNil

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) SetCpuCountNil(b bool)`

 SetCpuCountNil sets the value for CpuCount to be an explicit nil

### UnsetCpuCount
`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) UnsetCpuCount()`

UnsetCpuCount ensures that no value is present for CpuCount, not even an explicit nil
### GetMemoryMB

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) GetMemoryMB() int32`

GetMemoryMB returns the MemoryMB field if non-nil, zero value otherwise.

### GetMemoryMBOk

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) GetMemoryMBOk() (*int32, bool)`

GetMemoryMBOk returns a tuple with the MemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMB

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) SetMemoryMB(v int32)`

SetMemoryMB sets MemoryMB field to given value.

### HasMemoryMB

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) HasMemoryMB() bool`

HasMemoryMB returns a boolean if a field has been set.

### SetMemoryMBNil

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) SetMemoryMBNil(b bool)`

 SetMemoryMBNil sets the value for MemoryMB to be an explicit nil

### UnsetMemoryMB
`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) UnsetMemoryMB()`

UnsetMemoryMB ensures that no value is present for MemoryMB, not even an explicit nil
### GetMachineProfile

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) GetMachineProfile() HypervisorResourceRefResponseModel`

GetMachineProfile returns the MachineProfile field if non-nil, zero value otherwise.

### GetMachineProfileOk

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) GetMachineProfileOk() (*HypervisorResourceRefResponseModel, bool)`

GetMachineProfileOk returns a tuple with the MachineProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineProfile

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) SetMachineProfile(v HypervisorResourceRefResponseModel)`

SetMachineProfile sets MachineProfile field to given value.

### HasMachineProfile

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) HasMachineProfile() bool`

HasMachineProfile returns a boolean if a field has been set.

### GetServiceOffering

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) GetServiceOffering() string`

GetServiceOffering returns the ServiceOffering field if non-nil, zero value otherwise.

### GetServiceOfferingOk

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) GetServiceOfferingOk() (*string, bool)`

GetServiceOfferingOk returns a tuple with the ServiceOffering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceOffering

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) SetServiceOffering(v string)`

SetServiceOffering sets ServiceOffering field to given value.

### HasServiceOffering

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) HasServiceOffering() bool`

HasServiceOffering returns a boolean if a field has been set.

### SetServiceOfferingNil

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) SetServiceOfferingNil(b bool)`

 SetServiceOfferingNil sets the value for ServiceOffering to be an explicit nil

### UnsetServiceOffering
`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) UnsetServiceOffering()`

UnsetServiceOffering ensures that no value is present for ServiceOffering, not even an explicit nil
### GetCustomProperties

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) GetCustomProperties() []NameValueStringPairModel`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) GetCustomPropertiesOk() (*[]NameValueStringPairModel, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) SetCustomProperties(v []NameValueStringPairModel)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
### GetProvisioningSchemeName

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) GetProvisioningSchemeName() string`

GetProvisioningSchemeName returns the ProvisioningSchemeName field if non-nil, zero value otherwise.

### GetProvisioningSchemeNameOk

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) GetProvisioningSchemeNameOk() (*string, bool)`

GetProvisioningSchemeNameOk returns a tuple with the ProvisioningSchemeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeName

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) SetProvisioningSchemeName(v string)`

SetProvisioningSchemeName sets ProvisioningSchemeName field to given value.

### HasProvisioningSchemeName

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) HasProvisioningSchemeName() bool`

HasProvisioningSchemeName returns a boolean if a field has been set.

### SetProvisioningSchemeNameNil

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) SetProvisioningSchemeNameNil(b bool)`

 SetProvisioningSchemeNameNil sets the value for ProvisioningSchemeName to be an explicit nil

### UnsetProvisioningSchemeName
`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) UnsetProvisioningSchemeName()`

UnsetProvisioningSchemeName ensures that no value is present for ProvisioningSchemeName, not even an explicit nil
### GetProvisioningSchemeId

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) GetProvisioningSchemeId() string`

GetProvisioningSchemeId returns the ProvisioningSchemeId field if non-nil, zero value otherwise.

### GetProvisioningSchemeIdOk

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) GetProvisioningSchemeIdOk() (*string, bool)`

GetProvisioningSchemeIdOk returns a tuple with the ProvisioningSchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeId

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) SetProvisioningSchemeId(v string)`

SetProvisioningSchemeId sets ProvisioningSchemeId field to given value.

### HasProvisioningSchemeId

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) HasProvisioningSchemeId() bool`

HasProvisioningSchemeId returns a boolean if a field has been set.

### GetVMId

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) GetVMId() string`

GetVMId returns the VMId field if non-nil, zero value otherwise.

### GetVMIdOk

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) GetVMIdOk() (*string, bool)`

GetVMIdOk returns a tuple with the VMId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMId

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) SetVMId(v string)`

SetVMId sets VMId field to given value.

### HasVMId

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) HasVMId() bool`

HasVMId returns a boolean if a field has been set.

### SetVMIdNil

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) SetVMIdNil(b bool)`

 SetVMIdNil sets the value for VMId to be an explicit nil

### UnsetVMId
`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) UnsetVMId()`

UnsetVMId ensures that no value is present for VMId, not even an explicit nil
### GetVMName

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) GetVMName() string`

GetVMName returns the VMName field if non-nil, zero value otherwise.

### GetVMNameOk

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) GetVMNameOk() (*string, bool)`

GetVMNameOk returns a tuple with the VMName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMName

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) SetVMName(v string)`

SetVMName sets VMName field to given value.

### HasVMName

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) HasVMName() bool`

HasVMName returns a boolean if a field has been set.

### SetVMNameNil

`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) SetVMNameNil(b bool)`

 SetVMNameNil sets the value for VMName to be an explicit nil

### UnsetVMName
`func (o *ProvisionedVirtualMachineConfigurationResultantSetResponseModel) UnsetVMName()`

UnsetVMName ensures that no value is present for VMName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


