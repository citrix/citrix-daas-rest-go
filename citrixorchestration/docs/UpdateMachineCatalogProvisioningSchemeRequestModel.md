# UpdateMachineCatalogProvisioningSchemeRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterImagePath** | Pointer to **string** | The path in the resource pool to the virtual machine snapshot or VM template that will be used. This identifies the hard disk to be used and the default values for the memory and processors. This must be a path to a Snapshot or Template item in the resource pool to which the Machine Catalog is associated. | [optional] 
**MachineProfilePath** | Pointer to **string** | The path in the resource pool to the virtual machine template that will be used. This identifies the VM template to be used and the default values for the tags, virtual machine size, boot diagnostics, host cache property of OS disk, accelerated networking and availability zone. This must be a path to a Virtual machine or Template item in the resource pool to which the Machine Catalog is associated. | [optional] 
**StoreOldImage** | Pointer to **bool** | Whether the old image is stored in the catalog history. | [optional] 
**MinimumFunctionalLevel** | Pointer to [**FunctionalLevel**](FunctionalLevel.md) |  | [optional] 
**MasterImageNote** | Pointer to **string** | The note of the new image. | [optional] 
**CpuCount** | Pointer to **int32** | &#x60;DEPRECATED.&#x60; The number of processors that virtual machines created from the provisioning scheme should use. | [optional] 
**MemoryMB** | Pointer to **int32** | &#x60;DEPRECATED.&#x60; The maximum amount of memory that virtual machines created from the provisioning scheme should use. | [optional] 
**ServiceOfferingPath** | Pointer to **string** | &#x60;DEPRECATED.&#x60; The hypervisor resource path of the Cloud service offering to use when creating machines. | [optional] 
**CustomProperties** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | &#x60;DEPRECATED.&#x60; The properties of the provisioning scheme that are specific to the target hosting infrastructure. | [optional] 
**RebootOptions** | Pointer to [**UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions**](UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions.md) |  | [optional] 
**NumTotalMachines** | Pointer to **int32** | &#x60;DEPRECATED.&#x60; Total number of machines desired within the catalog. Optional; default is to leave the number of machines in the catalog unchanged. | [optional] 
**MachineAccountCreationRules** | Pointer to [**UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules**](UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules.md) |  | [optional] 
**AddAvailableMachineAccounts** | Pointer to [**[]MachineAccountRequestModel**](MachineAccountRequestModel.md) | &#x60;DEPRECATED.&#x60; List of Active Directory machine accounts to add to the pool of available accounts that are to be used when machines are provisioned. | [optional] 
**RemoveAvailableMachineAccounts** | Pointer to **[]string** | &#x60;DEPRECATED.&#x60; List of Active Directory machine accounts to remove from the pool of available accounts that are used when machines are provisioned. | [optional] 
**MachineAccountDeleteOption** | Pointer to [**MachineAccountDeleteOption**](MachineAccountDeleteOption.md) |  | [optional] 
**ImageVersion** | Pointer to [**CreateMachineCatalogProvisioningSchemeRequestModelImageVersion**](CreateMachineCatalogProvisioningSchemeRequestModelImageVersion.md) |  | [optional] 

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

### GetRebootOptions

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetRebootOptions() UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions`

GetRebootOptions returns the RebootOptions field if non-nil, zero value otherwise.

### GetRebootOptionsOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetRebootOptionsOk() (*UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions, bool)`

GetRebootOptionsOk returns a tuple with the RebootOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootOptions

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetRebootOptions(v UpdateMachineCatalogProvisioningSchemeRequestModelRebootOptions)`

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

### GetMachineAccountCreationRules

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetMachineAccountCreationRules() UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules`

GetMachineAccountCreationRules returns the MachineAccountCreationRules field if non-nil, zero value otherwise.

### GetMachineAccountCreationRulesOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetMachineAccountCreationRulesOk() (*UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules, bool)`

GetMachineAccountCreationRulesOk returns a tuple with the MachineAccountCreationRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineAccountCreationRules

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetMachineAccountCreationRules(v UpdateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules)`

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

### GetImageVersion

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetImageVersion() CreateMachineCatalogProvisioningSchemeRequestModelImageVersion`

GetImageVersion returns the ImageVersion field if non-nil, zero value otherwise.

### GetImageVersionOk

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) GetImageVersionOk() (*CreateMachineCatalogProvisioningSchemeRequestModelImageVersion, bool)`

GetImageVersionOk returns a tuple with the ImageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersion

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) SetImageVersion(v CreateMachineCatalogProvisioningSchemeRequestModelImageVersion)`

SetImageVersion sets ImageVersion field to given value.

### HasImageVersion

`func (o *UpdateMachineCatalogProvisioningSchemeRequestModel) HasImageVersion() bool`

HasImageVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


