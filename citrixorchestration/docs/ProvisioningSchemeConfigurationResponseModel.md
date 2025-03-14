# ProvisioningSchemeConfigurationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuCount** | Pointer to **NullableInt32** | cpu count | [optional] 
**MemoryMB** | Pointer to **NullableInt32** | Memory in MB | [optional] 
**CustomProperties** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Custom properties | [optional] 
**MachineProfile** | Pointer to [**HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) |  | [optional] 
**ServiceOffering** | Pointer to **NullableString** | Service offering used by Cloud provisioned VMs. | [optional] 
**Version** | Pointer to **int32** | Provisioning scheme configuration version.  | [optional] 
**NetworkMappings** | Pointer to [**[]NetworkMapResponseModel**](NetworkMapResponseModel.md) | Network mappings.  | [optional] 
**SecurityGroups** | Pointer to [**[]SecurityGroup**](SecurityGroup.md) | Security groups. | [optional] 
**ProvisioningSchemeName** | Pointer to **NullableString** | provisioning scheme name | [optional] 
**ProvisioningSchemeId** | Pointer to **string** | provisioning scheme uid | [optional] 

## Methods

### NewProvisioningSchemeConfigurationResponseModel

`func NewProvisioningSchemeConfigurationResponseModel() *ProvisioningSchemeConfigurationResponseModel`

NewProvisioningSchemeConfigurationResponseModel instantiates a new ProvisioningSchemeConfigurationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningSchemeConfigurationResponseModelWithDefaults

`func NewProvisioningSchemeConfigurationResponseModelWithDefaults() *ProvisioningSchemeConfigurationResponseModel`

NewProvisioningSchemeConfigurationResponseModelWithDefaults instantiates a new ProvisioningSchemeConfigurationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuCount

`func (o *ProvisioningSchemeConfigurationResponseModel) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *ProvisioningSchemeConfigurationResponseModel) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *ProvisioningSchemeConfigurationResponseModel) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *ProvisioningSchemeConfigurationResponseModel) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### SetCpuCountNil

`func (o *ProvisioningSchemeConfigurationResponseModel) SetCpuCountNil(b bool)`

 SetCpuCountNil sets the value for CpuCount to be an explicit nil

### UnsetCpuCount
`func (o *ProvisioningSchemeConfigurationResponseModel) UnsetCpuCount()`

UnsetCpuCount ensures that no value is present for CpuCount, not even an explicit nil
### GetMemoryMB

`func (o *ProvisioningSchemeConfigurationResponseModel) GetMemoryMB() int32`

GetMemoryMB returns the MemoryMB field if non-nil, zero value otherwise.

### GetMemoryMBOk

`func (o *ProvisioningSchemeConfigurationResponseModel) GetMemoryMBOk() (*int32, bool)`

GetMemoryMBOk returns a tuple with the MemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMB

`func (o *ProvisioningSchemeConfigurationResponseModel) SetMemoryMB(v int32)`

SetMemoryMB sets MemoryMB field to given value.

### HasMemoryMB

`func (o *ProvisioningSchemeConfigurationResponseModel) HasMemoryMB() bool`

HasMemoryMB returns a boolean if a field has been set.

### SetMemoryMBNil

`func (o *ProvisioningSchemeConfigurationResponseModel) SetMemoryMBNil(b bool)`

 SetMemoryMBNil sets the value for MemoryMB to be an explicit nil

### UnsetMemoryMB
`func (o *ProvisioningSchemeConfigurationResponseModel) UnsetMemoryMB()`

UnsetMemoryMB ensures that no value is present for MemoryMB, not even an explicit nil
### GetCustomProperties

`func (o *ProvisioningSchemeConfigurationResponseModel) GetCustomProperties() []NameValueStringPairModel`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ProvisioningSchemeConfigurationResponseModel) GetCustomPropertiesOk() (*[]NameValueStringPairModel, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ProvisioningSchemeConfigurationResponseModel) SetCustomProperties(v []NameValueStringPairModel)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ProvisioningSchemeConfigurationResponseModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *ProvisioningSchemeConfigurationResponseModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *ProvisioningSchemeConfigurationResponseModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
### GetMachineProfile

`func (o *ProvisioningSchemeConfigurationResponseModel) GetMachineProfile() HypervisorResourceRefResponseModel`

GetMachineProfile returns the MachineProfile field if non-nil, zero value otherwise.

### GetMachineProfileOk

`func (o *ProvisioningSchemeConfigurationResponseModel) GetMachineProfileOk() (*HypervisorResourceRefResponseModel, bool)`

GetMachineProfileOk returns a tuple with the MachineProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineProfile

`func (o *ProvisioningSchemeConfigurationResponseModel) SetMachineProfile(v HypervisorResourceRefResponseModel)`

SetMachineProfile sets MachineProfile field to given value.

### HasMachineProfile

`func (o *ProvisioningSchemeConfigurationResponseModel) HasMachineProfile() bool`

HasMachineProfile returns a boolean if a field has been set.

### GetServiceOffering

`func (o *ProvisioningSchemeConfigurationResponseModel) GetServiceOffering() string`

GetServiceOffering returns the ServiceOffering field if non-nil, zero value otherwise.

### GetServiceOfferingOk

`func (o *ProvisioningSchemeConfigurationResponseModel) GetServiceOfferingOk() (*string, bool)`

GetServiceOfferingOk returns a tuple with the ServiceOffering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceOffering

`func (o *ProvisioningSchemeConfigurationResponseModel) SetServiceOffering(v string)`

SetServiceOffering sets ServiceOffering field to given value.

### HasServiceOffering

`func (o *ProvisioningSchemeConfigurationResponseModel) HasServiceOffering() bool`

HasServiceOffering returns a boolean if a field has been set.

### SetServiceOfferingNil

`func (o *ProvisioningSchemeConfigurationResponseModel) SetServiceOfferingNil(b bool)`

 SetServiceOfferingNil sets the value for ServiceOffering to be an explicit nil

### UnsetServiceOffering
`func (o *ProvisioningSchemeConfigurationResponseModel) UnsetServiceOffering()`

UnsetServiceOffering ensures that no value is present for ServiceOffering, not even an explicit nil
### GetVersion

`func (o *ProvisioningSchemeConfigurationResponseModel) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProvisioningSchemeConfigurationResponseModel) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProvisioningSchemeConfigurationResponseModel) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ProvisioningSchemeConfigurationResponseModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetNetworkMappings

`func (o *ProvisioningSchemeConfigurationResponseModel) GetNetworkMappings() []NetworkMapResponseModel`

GetNetworkMappings returns the NetworkMappings field if non-nil, zero value otherwise.

### GetNetworkMappingsOk

`func (o *ProvisioningSchemeConfigurationResponseModel) GetNetworkMappingsOk() (*[]NetworkMapResponseModel, bool)`

GetNetworkMappingsOk returns a tuple with the NetworkMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMappings

`func (o *ProvisioningSchemeConfigurationResponseModel) SetNetworkMappings(v []NetworkMapResponseModel)`

SetNetworkMappings sets NetworkMappings field to given value.

### HasNetworkMappings

`func (o *ProvisioningSchemeConfigurationResponseModel) HasNetworkMappings() bool`

HasNetworkMappings returns a boolean if a field has been set.

### SetNetworkMappingsNil

`func (o *ProvisioningSchemeConfigurationResponseModel) SetNetworkMappingsNil(b bool)`

 SetNetworkMappingsNil sets the value for NetworkMappings to be an explicit nil

### UnsetNetworkMappings
`func (o *ProvisioningSchemeConfigurationResponseModel) UnsetNetworkMappings()`

UnsetNetworkMappings ensures that no value is present for NetworkMappings, not even an explicit nil
### GetSecurityGroups

`func (o *ProvisioningSchemeConfigurationResponseModel) GetSecurityGroups() []SecurityGroup`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *ProvisioningSchemeConfigurationResponseModel) GetSecurityGroupsOk() (*[]SecurityGroup, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *ProvisioningSchemeConfigurationResponseModel) SetSecurityGroups(v []SecurityGroup)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *ProvisioningSchemeConfigurationResponseModel) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *ProvisioningSchemeConfigurationResponseModel) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *ProvisioningSchemeConfigurationResponseModel) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetProvisioningSchemeName

`func (o *ProvisioningSchemeConfigurationResponseModel) GetProvisioningSchemeName() string`

GetProvisioningSchemeName returns the ProvisioningSchemeName field if non-nil, zero value otherwise.

### GetProvisioningSchemeNameOk

`func (o *ProvisioningSchemeConfigurationResponseModel) GetProvisioningSchemeNameOk() (*string, bool)`

GetProvisioningSchemeNameOk returns a tuple with the ProvisioningSchemeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeName

`func (o *ProvisioningSchemeConfigurationResponseModel) SetProvisioningSchemeName(v string)`

SetProvisioningSchemeName sets ProvisioningSchemeName field to given value.

### HasProvisioningSchemeName

`func (o *ProvisioningSchemeConfigurationResponseModel) HasProvisioningSchemeName() bool`

HasProvisioningSchemeName returns a boolean if a field has been set.

### SetProvisioningSchemeNameNil

`func (o *ProvisioningSchemeConfigurationResponseModel) SetProvisioningSchemeNameNil(b bool)`

 SetProvisioningSchemeNameNil sets the value for ProvisioningSchemeName to be an explicit nil

### UnsetProvisioningSchemeName
`func (o *ProvisioningSchemeConfigurationResponseModel) UnsetProvisioningSchemeName()`

UnsetProvisioningSchemeName ensures that no value is present for ProvisioningSchemeName, not even an explicit nil
### GetProvisioningSchemeId

`func (o *ProvisioningSchemeConfigurationResponseModel) GetProvisioningSchemeId() string`

GetProvisioningSchemeId returns the ProvisioningSchemeId field if non-nil, zero value otherwise.

### GetProvisioningSchemeIdOk

`func (o *ProvisioningSchemeConfigurationResponseModel) GetProvisioningSchemeIdOk() (*string, bool)`

GetProvisioningSchemeIdOk returns a tuple with the ProvisioningSchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeId

`func (o *ProvisioningSchemeConfigurationResponseModel) SetProvisioningSchemeId(v string)`

SetProvisioningSchemeId sets ProvisioningSchemeId field to given value.

### HasProvisioningSchemeId

`func (o *ProvisioningSchemeConfigurationResponseModel) HasProvisioningSchemeId() bool`

HasProvisioningSchemeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


