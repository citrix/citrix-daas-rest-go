# UpdateMachineCatalogProvisioningSchemeRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterImagePath** | Pointer to **NullableString** | The path in the resource pool to the virtual machine snapshot or VM template that will be used. This identifies the hard disk to be used and the default values for the memory and processors. This must be a path to a Snapshot or Template item in the resource pool to which the Machine Catalog is associated. | [optional] 
**MachineProfilePath** | Pointer to **NullableString** | The path in the resource pool to the virtual machine template that will be used. This identifies the VM template to be used and the default values for the tags, virtual machine size, boot diagnostics, host cache property of OS disk, accelerated networking and availability zone. This must be a path to a Virtual machine or Template item in the resource pool to which the Machine Catalog is associated. | [optional] 
**StoreOldImage** | Pointer to **NullableBool** | Whether the old image is stored in the catalog history. | [optional] 
**MinimumFunctionalLevel** | Pointer to [**FunctionalLevel**](FunctionalLevel.md) |  | [optional] 
**MasterImageNote** | Pointer to **NullableString** | The note of the new image. | [optional] 
**CpuCount** | Pointer to **NullableInt32** | &#x60;DEPRECATED.&#x60; The number of processors that virtual machines created from the provisioning scheme should use. | [optional] 
**MemoryMB** | Pointer to **NullableInt32** | &#x60;DEPRECATED.&#x60; The maximum amount of memory that virtual machines created from the provisioning scheme should use. | [optional] 
**ServiceOfferingPath** | Pointer to **NullableString** | &#x60;DEPRECATED.&#x60; The hypervisor resource path of the Cloud service offering to use when creating machines. | [optional] 
**CustomProperties** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The properties of the provisioning scheme that are specific to the target hosting infrastructure. | [optional] 
**CustomPropertiesInString** | Pointer to **NullableString** | The properties of the provisioning scheme in a single string, that are specific to the target hosting infrastructure. This is an optional method to set the custom properties which are not in the form of an array. | [optional] 
**RebootOptions** | Pointer to [**RebootMachinesRequestModel**](RebootMachinesRequestModel.md) |  | [optional] 
**NumTotalMachines** | Pointer to **NullableInt32** | &#x60;DEPRECATED.&#x60; Total number of machines desired within the catalog. Optional; default is to leave the number of machines in the catalog unchanged. | [optional] 
**MachineAccountCreationRules** | Pointer to [**UpdateMachineAccountCreationRulesRequestModel**](UpdateMachineAccountCreationRulesRequestModel.md) |  | [optional] 
**AddAvailableMachineAccounts** | Pointer to [**[]MachineAccountRequestModel**](MachineAccountRequestModel.md) | &#x60;DEPRECATED.&#x60; List of Active Directory machine accounts to add to the pool of available accounts that are to be used when machines are provisioned. | [optional] 
**RemoveAvailableMachineAccounts** | Pointer to **[]string** | &#x60;DEPRECATED.&#x60; List of Active Directory machine accounts to remove from the pool of available accounts that are used when machines are provisioned. | [optional] 
**MachineAccountDeleteOption** | Pointer to [**MachineAccountDeleteOption**](MachineAccountDeleteOption.md) |  | [optional] 
**AssignImageVersionToProvisioningScheme** | Pointer to [**AssignImageVersionToProvisioningSchemeRequestModel**](AssignImageVersionToProvisioningSchemeRequestModel.md) |  | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of provisioning scheme. Set the value of the NameValueStringPairModel is null or empty will be remove this metadata. Not existing Name and Value NameValueStringPairModel object will be added. The same Name but different value object will be updated. If Metadata is specified, note that other properties in request will be ignored. So MasterImagePath and  AssignImageVersionToProvisioningScheme only can be null or empty, to avoid unintented update. | [optional] 

## Methods

### NewUpdateMachineCatalogProvisioningSchemeRequestModel

`func NewUpdateMachineCatalogProvisioningSchemeRequestModel() *UpdateMachineCatalogProvisioningSchemeRequestModel`

NewUpdateMachineCatalogProvisioningSchemeRequestModel instantiates a new UpdateMachineCatalogProvisioningSchemeRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMachineCatalogProvisioningSchemeRequestModelWithDefaults

`func NewUpdateMachineCatalogProvisioningSchemeRequestModelWithDefaults() *UpdateMachineCatalogProvisioningSchemeRequestModel`

NewUpdateMachineCatalogProvisioningSchemeRequestModelWithDefaults instantiates a new UpdateMachineCatalogProvisioningSchemeRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMasterImagePath

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetMasterImagePath() string`

GetMasterImagePath returns the MasterImagePath field if non-nil, zero value otherwise.

### GetMasterImagePathOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetMasterImagePathOk() (*string, bool)`

GetMasterImagePathOk returns a tuple with the MasterImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterImagePath

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetMasterImagePath(v string)`

SetMasterImagePath sets MasterImagePath field to given value.

### HasMasterImagePath

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) HasMasterImagePath() bool`

HasMasterImagePath returns a boolean if a field has been set.

### SetMasterImagePathNil

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetMasterImagePathNil(b bool)`

 SetMasterImagePathNil sets the value for MasterImagePath to be an explicit nil

### UnsetMasterImagePath
`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) UnsetMasterImagePath()`

UnsetMasterImagePath ensures that no value is present for MasterImagePath, not even an explicit nil
### GetMachineProfilePath

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetMachineProfilePath() string`

GetMachineProfilePath returns the MachineProfilePath field if non-nil, zero value otherwise.

### GetMachineProfilePathOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetMachineProfilePathOk() (*string, bool)`

GetMachineProfilePathOk returns a tuple with the MachineProfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineProfilePath

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetMachineProfilePath(v string)`

SetMachineProfilePath sets MachineProfilePath field to given value.

### HasMachineProfilePath

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) HasMachineProfilePath() bool`

HasMachineProfilePath returns a boolean if a field has been set.

### SetMachineProfilePathNil

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetMachineProfilePathNil(b bool)`

 SetMachineProfilePathNil sets the value for MachineProfilePath to be an explicit nil

### UnsetMachineProfilePath
`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) UnsetMachineProfilePath()`

UnsetMachineProfilePath ensures that no value is present for MachineProfilePath, not even an explicit nil
### GetStoreOldImage

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetStoreOldImage() bool`

GetStoreOldImage returns the StoreOldImage field if non-nil, zero value otherwise.

### GetStoreOldImageOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetStoreOldImageOk() (*bool, bool)`

GetStoreOldImageOk returns a tuple with the StoreOldImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreOldImage

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetStoreOldImage(v bool)`

SetStoreOldImage sets StoreOldImage field to given value.

### HasStoreOldImage

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) HasStoreOldImage() bool`

HasStoreOldImage returns a boolean if a field has been set.

### SetStoreOldImageNil

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetStoreOldImageNil(b bool)`

 SetStoreOldImageNil sets the value for StoreOldImage to be an explicit nil

### UnsetStoreOldImage
`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) UnsetStoreOldImage()`

UnsetStoreOldImage ensures that no value is present for StoreOldImage, not even an explicit nil
### GetMinimumFunctionalLevel

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetMinimumFunctionalLevel() FunctionalLevel`

GetMinimumFunctionalLevel returns the MinimumFunctionalLevel field if non-nil, zero value otherwise.

### GetMinimumFunctionalLevelOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetMinimumFunctionalLevelOk() (*FunctionalLevel, bool)`

GetMinimumFunctionalLevelOk returns a tuple with the MinimumFunctionalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumFunctionalLevel

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetMinimumFunctionalLevel(v FunctionalLevel)`

SetMinimumFunctionalLevel sets MinimumFunctionalLevel field to given value.

### HasMinimumFunctionalLevel

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) HasMinimumFunctionalLevel() bool`

HasMinimumFunctionalLevel returns a boolean if a field has been set.

### GetMasterImageNote

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetMasterImageNote() string`

GetMasterImageNote returns the MasterImageNote field if non-nil, zero value otherwise.

### GetMasterImageNoteOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetMasterImageNoteOk() (*string, bool)`

GetMasterImageNoteOk returns a tuple with the MasterImageNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterImageNote

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetMasterImageNote(v string)`

SetMasterImageNote sets MasterImageNote field to given value.

### HasMasterImageNote

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) HasMasterImageNote() bool`

HasMasterImageNote returns a boolean if a field has been set.

### SetMasterImageNoteNil

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetMasterImageNoteNil(b bool)`

 SetMasterImageNoteNil sets the value for MasterImageNote to be an explicit nil

### UnsetMasterImageNote
`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) UnsetMasterImageNote()`

UnsetMasterImageNote ensures that no value is present for MasterImageNote, not even an explicit nil
### GetCpuCount

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### SetCpuCountNil

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetCpuCountNil(b bool)`

 SetCpuCountNil sets the value for CpuCount to be an explicit nil

### UnsetCpuCount
`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) UnsetCpuCount()`

UnsetCpuCount ensures that no value is present for CpuCount, not even an explicit nil
### GetMemoryMB

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetMemoryMB() int32`

GetMemoryMB returns the MemoryMB field if non-nil, zero value otherwise.

### GetMemoryMBOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetMemoryMBOk() (*int32, bool)`

GetMemoryMBOk returns a tuple with the MemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMB

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetMemoryMB(v int32)`

SetMemoryMB sets MemoryMB field to given value.

### HasMemoryMB

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) HasMemoryMB() bool`

HasMemoryMB returns a boolean if a field has been set.

### SetMemoryMBNil

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetMemoryMBNil(b bool)`

 SetMemoryMBNil sets the value for MemoryMB to be an explicit nil

### UnsetMemoryMB
`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) UnsetMemoryMB()`

UnsetMemoryMB ensures that no value is present for MemoryMB, not even an explicit nil
### GetServiceOfferingPath

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetServiceOfferingPath() string`

GetServiceOfferingPath returns the ServiceOfferingPath field if non-nil, zero value otherwise.

### GetServiceOfferingPathOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetServiceOfferingPathOk() (*string, bool)`

GetServiceOfferingPathOk returns a tuple with the ServiceOfferingPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceOfferingPath

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetServiceOfferingPath(v string)`

SetServiceOfferingPath sets ServiceOfferingPath field to given value.

### HasServiceOfferingPath

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) HasServiceOfferingPath() bool`

HasServiceOfferingPath returns a boolean if a field has been set.

### SetServiceOfferingPathNil

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetServiceOfferingPathNil(b bool)`

 SetServiceOfferingPathNil sets the value for ServiceOfferingPath to be an explicit nil

### UnsetServiceOfferingPath
`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) UnsetServiceOfferingPath()`

UnsetServiceOfferingPath ensures that no value is present for ServiceOfferingPath, not even an explicit nil
### GetCustomProperties

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetCustomProperties() []NameValueStringPairModel`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetCustomPropertiesOk() (*[]NameValueStringPairModel, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetCustomProperties(v []NameValueStringPairModel)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
### GetCustomPropertiesInString

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetCustomPropertiesInString() string`

GetCustomPropertiesInString returns the CustomPropertiesInString field if non-nil, zero value otherwise.

### GetCustomPropertiesInStringOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetCustomPropertiesInStringOk() (*string, bool)`

GetCustomPropertiesInStringOk returns a tuple with the CustomPropertiesInString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPropertiesInString

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetCustomPropertiesInString(v string)`

SetCustomPropertiesInString sets CustomPropertiesInString field to given value.

### HasCustomPropertiesInString

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) HasCustomPropertiesInString() bool`

HasCustomPropertiesInString returns a boolean if a field has been set.

### SetCustomPropertiesInStringNil

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetCustomPropertiesInStringNil(b bool)`

 SetCustomPropertiesInStringNil sets the value for CustomPropertiesInString to be an explicit nil

### UnsetCustomPropertiesInString
`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) UnsetCustomPropertiesInString()`

UnsetCustomPropertiesInString ensures that no value is present for CustomPropertiesInString, not even an explicit nil
### GetRebootOptions

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetRebootOptions() RebootMachinesRequestModel`

GetRebootOptions returns the RebootOptions field if non-nil, zero value otherwise.

### GetRebootOptionsOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetRebootOptionsOk() (*RebootMachinesRequestModel, bool)`

GetRebootOptionsOk returns a tuple with the RebootOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootOptions

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetRebootOptions(v RebootMachinesRequestModel)`

SetRebootOptions sets RebootOptions field to given value.

### HasRebootOptions

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) HasRebootOptions() bool`

HasRebootOptions returns a boolean if a field has been set.

### GetNumTotalMachines

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetNumTotalMachines() int32`

GetNumTotalMachines returns the NumTotalMachines field if non-nil, zero value otherwise.

### GetNumTotalMachinesOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetNumTotalMachinesOk() (*int32, bool)`

GetNumTotalMachinesOk returns a tuple with the NumTotalMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTotalMachines

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetNumTotalMachines(v int32)`

SetNumTotalMachines sets NumTotalMachines field to given value.

### HasNumTotalMachines

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) HasNumTotalMachines() bool`

HasNumTotalMachines returns a boolean if a field has been set.

### SetNumTotalMachinesNil

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetNumTotalMachinesNil(b bool)`

 SetNumTotalMachinesNil sets the value for NumTotalMachines to be an explicit nil

### UnsetNumTotalMachines
`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) UnsetNumTotalMachines()`

UnsetNumTotalMachines ensures that no value is present for NumTotalMachines, not even an explicit nil
### GetMachineAccountCreationRules

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetMachineAccountCreationRules() UpdateMachineAccountCreationRulesRequestModel`

GetMachineAccountCreationRules returns the MachineAccountCreationRules field if non-nil, zero value otherwise.

### GetMachineAccountCreationRulesOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetMachineAccountCreationRulesOk() (*UpdateMachineAccountCreationRulesRequestModel, bool)`

GetMachineAccountCreationRulesOk returns a tuple with the MachineAccountCreationRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineAccountCreationRules

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetMachineAccountCreationRules(v UpdateMachineAccountCreationRulesRequestModel)`

SetMachineAccountCreationRules sets MachineAccountCreationRules field to given value.

### HasMachineAccountCreationRules

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) HasMachineAccountCreationRules() bool`

HasMachineAccountCreationRules returns a boolean if a field has been set.

### GetAddAvailableMachineAccounts

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetAddAvailableMachineAccounts() []MachineAccountRequestModel`

GetAddAvailableMachineAccounts returns the AddAvailableMachineAccounts field if non-nil, zero value otherwise.

### GetAddAvailableMachineAccountsOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetAddAvailableMachineAccountsOk() (*[]MachineAccountRequestModel, bool)`

GetAddAvailableMachineAccountsOk returns a tuple with the AddAvailableMachineAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAvailableMachineAccounts

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetAddAvailableMachineAccounts(v []MachineAccountRequestModel)`

SetAddAvailableMachineAccounts sets AddAvailableMachineAccounts field to given value.

### HasAddAvailableMachineAccounts

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) HasAddAvailableMachineAccounts() bool`

HasAddAvailableMachineAccounts returns a boolean if a field has been set.

### SetAddAvailableMachineAccountsNil

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetAddAvailableMachineAccountsNil(b bool)`

 SetAddAvailableMachineAccountsNil sets the value for AddAvailableMachineAccounts to be an explicit nil

### UnsetAddAvailableMachineAccounts
`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) UnsetAddAvailableMachineAccounts()`

UnsetAddAvailableMachineAccounts ensures that no value is present for AddAvailableMachineAccounts, not even an explicit nil
### GetRemoveAvailableMachineAccounts

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetRemoveAvailableMachineAccounts() []string`

GetRemoveAvailableMachineAccounts returns the RemoveAvailableMachineAccounts field if non-nil, zero value otherwise.

### GetRemoveAvailableMachineAccountsOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetRemoveAvailableMachineAccountsOk() (*[]string, bool)`

GetRemoveAvailableMachineAccountsOk returns a tuple with the RemoveAvailableMachineAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAvailableMachineAccounts

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetRemoveAvailableMachineAccounts(v []string)`

SetRemoveAvailableMachineAccounts sets RemoveAvailableMachineAccounts field to given value.

### HasRemoveAvailableMachineAccounts

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) HasRemoveAvailableMachineAccounts() bool`

HasRemoveAvailableMachineAccounts returns a boolean if a field has been set.

### SetRemoveAvailableMachineAccountsNil

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetRemoveAvailableMachineAccountsNil(b bool)`

 SetRemoveAvailableMachineAccountsNil sets the value for RemoveAvailableMachineAccounts to be an explicit nil

### UnsetRemoveAvailableMachineAccounts
`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) UnsetRemoveAvailableMachineAccounts()`

UnsetRemoveAvailableMachineAccounts ensures that no value is present for RemoveAvailableMachineAccounts, not even an explicit nil
### GetMachineAccountDeleteOption

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetMachineAccountDeleteOption() MachineAccountDeleteOption`

GetMachineAccountDeleteOption returns the MachineAccountDeleteOption field if non-nil, zero value otherwise.

### GetMachineAccountDeleteOptionOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetMachineAccountDeleteOptionOk() (*MachineAccountDeleteOption, bool)`

GetMachineAccountDeleteOptionOk returns a tuple with the MachineAccountDeleteOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineAccountDeleteOption

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetMachineAccountDeleteOption(v MachineAccountDeleteOption)`

SetMachineAccountDeleteOption sets MachineAccountDeleteOption field to given value.

### HasMachineAccountDeleteOption

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) HasMachineAccountDeleteOption() bool`

HasMachineAccountDeleteOption returns a boolean if a field has been set.

### GetAssignImageVersionToProvisioningScheme

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetAssignImageVersionToProvisioningScheme() AssignImageVersionToProvisioningSchemeRequestModel`

GetAssignImageVersionToProvisioningScheme returns the AssignImageVersionToProvisioningScheme field if non-nil, zero value otherwise.

### GetAssignImageVersionToProvisioningSchemeOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetAssignImageVersionToProvisioningSchemeOk() (*AssignImageVersionToProvisioningSchemeRequestModel, bool)`

GetAssignImageVersionToProvisioningSchemeOk returns a tuple with the AssignImageVersionToProvisioningScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignImageVersionToProvisioningScheme

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetAssignImageVersionToProvisioningScheme(v AssignImageVersionToProvisioningSchemeRequestModel)`

SetAssignImageVersionToProvisioningScheme sets AssignImageVersionToProvisioningScheme field to given value.

### HasAssignImageVersionToProvisioningScheme

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) HasAssignImageVersionToProvisioningScheme() bool`

HasAssignImageVersionToProvisioningScheme returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


