# ProvisionedBackupResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Id of the backup. | [optional] 
**Name** | Pointer to **NullableString** | Name of the backup. | [optional] 
**Description** | Pointer to **NullableString** | Description of the backup. | [optional] 
**Version** | Pointer to **NullableInt32** | Version of the backup. | [optional] 
**CreationTime** | Pointer to **NullableString** | Time the backup was created.    RFC 3339 compatible format. | [optional] 
**VmSid** | Pointer to **NullableString** | Provisioned Virtual Machine Sid. | [optional] 
**VmName** | Pointer to **NullableString** | Provisioned virtual machine name on hypervisor. | [optional] 
**VmId** | Pointer to **NullableString** | Id of the VM, as defined by the hypervisor. | [optional] 
**ProvisioningSchemeName** | Pointer to **NullableString** | Name of the provisioning scheme which the vm is a member of. | [optional] 
**ProvisioningSchemeId** | Pointer to **NullableString** | Id of the provisioning scheme which the vm is a member of. | [optional] 
**ProvisionedVmConfigurationVersion** | Pointer to **NullableInt32** | The VM configuration version associated to this backup. | [optional] 
**ProvisioningSchemeConfigurationVersion** | Pointer to **NullableInt32** | The provisioning scheme configuration version associated to this backup. | [optional] 

## Methods

### NewProvisionedBackupResponseModel

`func NewProvisionedBackupResponseModel() *ProvisionedBackupResponseModel`

NewProvisionedBackupResponseModel instantiates a new ProvisionedBackupResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionedBackupResponseModelWithDefaults

`func NewProvisionedBackupResponseModelWithDefaults() *ProvisionedBackupResponseModel`

NewProvisionedBackupResponseModelWithDefaults instantiates a new ProvisionedBackupResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProvisionedBackupResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvisionedBackupResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvisionedBackupResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProvisionedBackupResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProvisionedBackupResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProvisionedBackupResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ProvisionedBackupResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisionedBackupResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisionedBackupResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProvisionedBackupResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProvisionedBackupResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProvisionedBackupResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ProvisionedBackupResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProvisionedBackupResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProvisionedBackupResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProvisionedBackupResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProvisionedBackupResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProvisionedBackupResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVersion

`func (o *ProvisionedBackupResponseModel) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProvisionedBackupResponseModel) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProvisionedBackupResponseModel) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ProvisionedBackupResponseModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ProvisionedBackupResponseModel) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ProvisionedBackupResponseModel) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetCreationTime

`func (o *ProvisionedBackupResponseModel) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ProvisionedBackupResponseModel) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ProvisionedBackupResponseModel) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *ProvisionedBackupResponseModel) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### SetCreationTimeNil

`func (o *ProvisionedBackupResponseModel) SetCreationTimeNil(b bool)`

 SetCreationTimeNil sets the value for CreationTime to be an explicit nil

### UnsetCreationTime
`func (o *ProvisionedBackupResponseModel) UnsetCreationTime()`

UnsetCreationTime ensures that no value is present for CreationTime, not even an explicit nil
### GetVmSid

`func (o *ProvisionedBackupResponseModel) GetVmSid() string`

GetVmSid returns the VmSid field if non-nil, zero value otherwise.

### GetVmSidOk

`func (o *ProvisionedBackupResponseModel) GetVmSidOk() (*string, bool)`

GetVmSidOk returns a tuple with the VmSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSid

`func (o *ProvisionedBackupResponseModel) SetVmSid(v string)`

SetVmSid sets VmSid field to given value.

### HasVmSid

`func (o *ProvisionedBackupResponseModel) HasVmSid() bool`

HasVmSid returns a boolean if a field has been set.

### SetVmSidNil

`func (o *ProvisionedBackupResponseModel) SetVmSidNil(b bool)`

 SetVmSidNil sets the value for VmSid to be an explicit nil

### UnsetVmSid
`func (o *ProvisionedBackupResponseModel) UnsetVmSid()`

UnsetVmSid ensures that no value is present for VmSid, not even an explicit nil
### GetVmName

`func (o *ProvisionedBackupResponseModel) GetVmName() string`

GetVmName returns the VmName field if non-nil, zero value otherwise.

### GetVmNameOk

`func (o *ProvisionedBackupResponseModel) GetVmNameOk() (*string, bool)`

GetVmNameOk returns a tuple with the VmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmName

`func (o *ProvisionedBackupResponseModel) SetVmName(v string)`

SetVmName sets VmName field to given value.

### HasVmName

`func (o *ProvisionedBackupResponseModel) HasVmName() bool`

HasVmName returns a boolean if a field has been set.

### SetVmNameNil

`func (o *ProvisionedBackupResponseModel) SetVmNameNil(b bool)`

 SetVmNameNil sets the value for VmName to be an explicit nil

### UnsetVmName
`func (o *ProvisionedBackupResponseModel) UnsetVmName()`

UnsetVmName ensures that no value is present for VmName, not even an explicit nil
### GetVmId

`func (o *ProvisionedBackupResponseModel) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *ProvisionedBackupResponseModel) GetVmIdOk() (*string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *ProvisionedBackupResponseModel) SetVmId(v string)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *ProvisionedBackupResponseModel) HasVmId() bool`

HasVmId returns a boolean if a field has been set.

### SetVmIdNil

`func (o *ProvisionedBackupResponseModel) SetVmIdNil(b bool)`

 SetVmIdNil sets the value for VmId to be an explicit nil

### UnsetVmId
`func (o *ProvisionedBackupResponseModel) UnsetVmId()`

UnsetVmId ensures that no value is present for VmId, not even an explicit nil
### GetProvisioningSchemeName

`func (o *ProvisionedBackupResponseModel) GetProvisioningSchemeName() string`

GetProvisioningSchemeName returns the ProvisioningSchemeName field if non-nil, zero value otherwise.

### GetProvisioningSchemeNameOk

`func (o *ProvisionedBackupResponseModel) GetProvisioningSchemeNameOk() (*string, bool)`

GetProvisioningSchemeNameOk returns a tuple with the ProvisioningSchemeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeName

`func (o *ProvisionedBackupResponseModel) SetProvisioningSchemeName(v string)`

SetProvisioningSchemeName sets ProvisioningSchemeName field to given value.

### HasProvisioningSchemeName

`func (o *ProvisionedBackupResponseModel) HasProvisioningSchemeName() bool`

HasProvisioningSchemeName returns a boolean if a field has been set.

### SetProvisioningSchemeNameNil

`func (o *ProvisionedBackupResponseModel) SetProvisioningSchemeNameNil(b bool)`

 SetProvisioningSchemeNameNil sets the value for ProvisioningSchemeName to be an explicit nil

### UnsetProvisioningSchemeName
`func (o *ProvisionedBackupResponseModel) UnsetProvisioningSchemeName()`

UnsetProvisioningSchemeName ensures that no value is present for ProvisioningSchemeName, not even an explicit nil
### GetProvisioningSchemeId

`func (o *ProvisionedBackupResponseModel) GetProvisioningSchemeId() string`

GetProvisioningSchemeId returns the ProvisioningSchemeId field if non-nil, zero value otherwise.

### GetProvisioningSchemeIdOk

`func (o *ProvisionedBackupResponseModel) GetProvisioningSchemeIdOk() (*string, bool)`

GetProvisioningSchemeIdOk returns a tuple with the ProvisioningSchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeId

`func (o *ProvisionedBackupResponseModel) SetProvisioningSchemeId(v string)`

SetProvisioningSchemeId sets ProvisioningSchemeId field to given value.

### HasProvisioningSchemeId

`func (o *ProvisionedBackupResponseModel) HasProvisioningSchemeId() bool`

HasProvisioningSchemeId returns a boolean if a field has been set.

### SetProvisioningSchemeIdNil

`func (o *ProvisionedBackupResponseModel) SetProvisioningSchemeIdNil(b bool)`

 SetProvisioningSchemeIdNil sets the value for ProvisioningSchemeId to be an explicit nil

### UnsetProvisioningSchemeId
`func (o *ProvisionedBackupResponseModel) UnsetProvisioningSchemeId()`

UnsetProvisioningSchemeId ensures that no value is present for ProvisioningSchemeId, not even an explicit nil
### GetProvisionedVmConfigurationVersion

`func (o *ProvisionedBackupResponseModel) GetProvisionedVmConfigurationVersion() int32`

GetProvisionedVmConfigurationVersion returns the ProvisionedVmConfigurationVersion field if non-nil, zero value otherwise.

### GetProvisionedVmConfigurationVersionOk

`func (o *ProvisionedBackupResponseModel) GetProvisionedVmConfigurationVersionOk() (*int32, bool)`

GetProvisionedVmConfigurationVersionOk returns a tuple with the ProvisionedVmConfigurationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedVmConfigurationVersion

`func (o *ProvisionedBackupResponseModel) SetProvisionedVmConfigurationVersion(v int32)`

SetProvisionedVmConfigurationVersion sets ProvisionedVmConfigurationVersion field to given value.

### HasProvisionedVmConfigurationVersion

`func (o *ProvisionedBackupResponseModel) HasProvisionedVmConfigurationVersion() bool`

HasProvisionedVmConfigurationVersion returns a boolean if a field has been set.

### SetProvisionedVmConfigurationVersionNil

`func (o *ProvisionedBackupResponseModel) SetProvisionedVmConfigurationVersionNil(b bool)`

 SetProvisionedVmConfigurationVersionNil sets the value for ProvisionedVmConfigurationVersion to be an explicit nil

### UnsetProvisionedVmConfigurationVersion
`func (o *ProvisionedBackupResponseModel) UnsetProvisionedVmConfigurationVersion()`

UnsetProvisionedVmConfigurationVersion ensures that no value is present for ProvisionedVmConfigurationVersion, not even an explicit nil
### GetProvisioningSchemeConfigurationVersion

`func (o *ProvisionedBackupResponseModel) GetProvisioningSchemeConfigurationVersion() int32`

GetProvisioningSchemeConfigurationVersion returns the ProvisioningSchemeConfigurationVersion field if non-nil, zero value otherwise.

### GetProvisioningSchemeConfigurationVersionOk

`func (o *ProvisionedBackupResponseModel) GetProvisioningSchemeConfigurationVersionOk() (*int32, bool)`

GetProvisioningSchemeConfigurationVersionOk returns a tuple with the ProvisioningSchemeConfigurationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeConfigurationVersion

`func (o *ProvisionedBackupResponseModel) SetProvisioningSchemeConfigurationVersion(v int32)`

SetProvisioningSchemeConfigurationVersion sets ProvisioningSchemeConfigurationVersion field to given value.

### HasProvisioningSchemeConfigurationVersion

`func (o *ProvisionedBackupResponseModel) HasProvisioningSchemeConfigurationVersion() bool`

HasProvisioningSchemeConfigurationVersion returns a boolean if a field has been set.

### SetProvisioningSchemeConfigurationVersionNil

`func (o *ProvisionedBackupResponseModel) SetProvisioningSchemeConfigurationVersionNil(b bool)`

 SetProvisioningSchemeConfigurationVersionNil sets the value for ProvisioningSchemeConfigurationVersion to be an explicit nil

### UnsetProvisioningSchemeConfigurationVersion
`func (o *ProvisionedBackupResponseModel) UnsetProvisioningSchemeConfigurationVersion()`

UnsetProvisioningSchemeConfigurationVersion ensures that no value is present for ProvisioningSchemeConfigurationVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


