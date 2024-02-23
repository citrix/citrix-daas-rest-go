# ProvisionedVirtualMachineCustomVmDataResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZone** | Pointer to **NullableString** | Availability Zones | [optional] 
**DataDiskId** | Pointer to **NullableString** | WBCDiskId | [optional] 
**DataDiskStorageType** | Pointer to **NullableString** | Data Disk Storage Type | [optional] 
**EphemeralOsDiskMode** | Pointer to **NullableString** | Ephemeral OS Disk Mode | [optional] 
**EncryptionSettingDESI** | Pointer to **NullableString** | Disk Encryption Set Id | [optional] 
**FailSafeStorageType** | Pointer to **bool** | Storage Type - Fail Safe | [optional] 
**HostGroupId** | Pointer to **NullableString** | Host Group | [optional] 
**IdentityDiskId** | Pointer to **NullableString** | Identity Disk full path | [optional] 
**IdentityDiskStorageType** | Pointer to **NullableString** | Identity Disk Storage Type | [optional] 
**IsOSDiskAzureTempDisk** | Pointer to **bool** | OS Disk Use Temp Disk | [optional] 
**IsSharedImageGalleryUsed** | Pointer to **NullableBool** | Use Shared Image Gallery | [optional] 
**IsUsingEphemeralOsDisk** | Pointer to **bool** | Ephemeral OS Disk | [optional] 
**LicenseType** | Pointer to **NullableString** | License Type | [optional] 
**MachineSizeId** | Pointer to **NullableString** | Machine Size | [optional] 
**NetworkInterfaceIds** | Pointer to **[]string** | Network Maps | [optional] 
**OSDiskId** | Pointer to **NullableString** | OS Disk full path | [optional] 
**OsDiskStorageType** | Pointer to **NullableString** | OS Disk Storage Type | [optional] 
**OsType** | Pointer to **NullableString** | Operation System | [optional] 
**PersistOsDisk** | Pointer to **NullableBool** | Retaining System Disk | [optional] 
**PersistWbc** | Pointer to **NullableBool** | Retaining Write-back Cache Disk | [optional] 
**RegionId** | Pointer to **NullableString** | Region Id | [optional] 
**ResourceGroupName** | Pointer to **NullableString** | Resource Group | [optional] 
**SecurityType** | Pointer to **NullableString** | Security Type | [optional] 
**SIGMachineToReplicaRatio** | Pointer to **NullableInt32** | Azure Shared Image Gallery Machine To Replica Ratio | [optional] 
**SIGReplicaMaximum** | Pointer to **NullableInt32** | Azure Shared Image Replica Maximum | [optional] 
**StorageTypeAtShutdown** | Pointer to **NullableString** | Storage Type At Shutdown | [optional] 
**SupportsHibernation** | Pointer to **NullableBool** | VmMetadata hibernation configuration. | [optional] 

## Methods

### NewProvisionedVirtualMachineCustomVmDataResponseModel

`func NewProvisionedVirtualMachineCustomVmDataResponseModel() *ProvisionedVirtualMachineCustomVmDataResponseModel`

NewProvisionedVirtualMachineCustomVmDataResponseModel instantiates a new ProvisionedVirtualMachineCustomVmDataResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionedVirtualMachineCustomVmDataResponseModelWithDefaults

`func NewProvisionedVirtualMachineCustomVmDataResponseModelWithDefaults() *ProvisionedVirtualMachineCustomVmDataResponseModel`

NewProvisionedVirtualMachineCustomVmDataResponseModelWithDefaults instantiates a new ProvisionedVirtualMachineCustomVmDataResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZone

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetDataDiskId

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetDataDiskId() string`

GetDataDiskId returns the DataDiskId field if non-nil, zero value otherwise.

### GetDataDiskIdOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetDataDiskIdOk() (*string, bool)`

GetDataDiskIdOk returns a tuple with the DataDiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDiskId

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetDataDiskId(v string)`

SetDataDiskId sets DataDiskId field to given value.

### HasDataDiskId

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasDataDiskId() bool`

HasDataDiskId returns a boolean if a field has been set.

### SetDataDiskIdNil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetDataDiskIdNil(b bool)`

 SetDataDiskIdNil sets the value for DataDiskId to be an explicit nil

### UnsetDataDiskId
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetDataDiskId()`

UnsetDataDiskId ensures that no value is present for DataDiskId, not even an explicit nil
### GetDataDiskStorageType

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetDataDiskStorageType() string`

GetDataDiskStorageType returns the DataDiskStorageType field if non-nil, zero value otherwise.

### GetDataDiskStorageTypeOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetDataDiskStorageTypeOk() (*string, bool)`

GetDataDiskStorageTypeOk returns a tuple with the DataDiskStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDiskStorageType

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetDataDiskStorageType(v string)`

SetDataDiskStorageType sets DataDiskStorageType field to given value.

### HasDataDiskStorageType

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasDataDiskStorageType() bool`

HasDataDiskStorageType returns a boolean if a field has been set.

### SetDataDiskStorageTypeNil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetDataDiskStorageTypeNil(b bool)`

 SetDataDiskStorageTypeNil sets the value for DataDiskStorageType to be an explicit nil

### UnsetDataDiskStorageType
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetDataDiskStorageType()`

UnsetDataDiskStorageType ensures that no value is present for DataDiskStorageType, not even an explicit nil
### GetEphemeralOsDiskMode

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetEphemeralOsDiskMode() string`

GetEphemeralOsDiskMode returns the EphemeralOsDiskMode field if non-nil, zero value otherwise.

### GetEphemeralOsDiskModeOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetEphemeralOsDiskModeOk() (*string, bool)`

GetEphemeralOsDiskModeOk returns a tuple with the EphemeralOsDiskMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralOsDiskMode

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetEphemeralOsDiskMode(v string)`

SetEphemeralOsDiskMode sets EphemeralOsDiskMode field to given value.

### HasEphemeralOsDiskMode

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasEphemeralOsDiskMode() bool`

HasEphemeralOsDiskMode returns a boolean if a field has been set.

### SetEphemeralOsDiskModeNil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetEphemeralOsDiskModeNil(b bool)`

 SetEphemeralOsDiskModeNil sets the value for EphemeralOsDiskMode to be an explicit nil

### UnsetEphemeralOsDiskMode
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetEphemeralOsDiskMode()`

UnsetEphemeralOsDiskMode ensures that no value is present for EphemeralOsDiskMode, not even an explicit nil
### GetEncryptionSettingDESI

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetEncryptionSettingDESI() string`

GetEncryptionSettingDESI returns the EncryptionSettingDESI field if non-nil, zero value otherwise.

### GetEncryptionSettingDESIOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetEncryptionSettingDESIOk() (*string, bool)`

GetEncryptionSettingDESIOk returns a tuple with the EncryptionSettingDESI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSettingDESI

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetEncryptionSettingDESI(v string)`

SetEncryptionSettingDESI sets EncryptionSettingDESI field to given value.

### HasEncryptionSettingDESI

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasEncryptionSettingDESI() bool`

HasEncryptionSettingDESI returns a boolean if a field has been set.

### SetEncryptionSettingDESINil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetEncryptionSettingDESINil(b bool)`

 SetEncryptionSettingDESINil sets the value for EncryptionSettingDESI to be an explicit nil

### UnsetEncryptionSettingDESI
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetEncryptionSettingDESI()`

UnsetEncryptionSettingDESI ensures that no value is present for EncryptionSettingDESI, not even an explicit nil
### GetFailSafeStorageType

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetFailSafeStorageType() bool`

GetFailSafeStorageType returns the FailSafeStorageType field if non-nil, zero value otherwise.

### GetFailSafeStorageTypeOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetFailSafeStorageTypeOk() (*bool, bool)`

GetFailSafeStorageTypeOk returns a tuple with the FailSafeStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailSafeStorageType

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetFailSafeStorageType(v bool)`

SetFailSafeStorageType sets FailSafeStorageType field to given value.

### HasFailSafeStorageType

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasFailSafeStorageType() bool`

HasFailSafeStorageType returns a boolean if a field has been set.

### GetHostGroupId

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetHostGroupId() string`

GetHostGroupId returns the HostGroupId field if non-nil, zero value otherwise.

### GetHostGroupIdOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetHostGroupIdOk() (*string, bool)`

GetHostGroupIdOk returns a tuple with the HostGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroupId

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetHostGroupId(v string)`

SetHostGroupId sets HostGroupId field to given value.

### HasHostGroupId

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasHostGroupId() bool`

HasHostGroupId returns a boolean if a field has been set.

### SetHostGroupIdNil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetHostGroupIdNil(b bool)`

 SetHostGroupIdNil sets the value for HostGroupId to be an explicit nil

### UnsetHostGroupId
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetHostGroupId()`

UnsetHostGroupId ensures that no value is present for HostGroupId, not even an explicit nil
### GetIdentityDiskId

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetIdentityDiskId() string`

GetIdentityDiskId returns the IdentityDiskId field if non-nil, zero value otherwise.

### GetIdentityDiskIdOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetIdentityDiskIdOk() (*string, bool)`

GetIdentityDiskIdOk returns a tuple with the IdentityDiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityDiskId

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetIdentityDiskId(v string)`

SetIdentityDiskId sets IdentityDiskId field to given value.

### HasIdentityDiskId

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasIdentityDiskId() bool`

HasIdentityDiskId returns a boolean if a field has been set.

### SetIdentityDiskIdNil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetIdentityDiskIdNil(b bool)`

 SetIdentityDiskIdNil sets the value for IdentityDiskId to be an explicit nil

### UnsetIdentityDiskId
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetIdentityDiskId()`

UnsetIdentityDiskId ensures that no value is present for IdentityDiskId, not even an explicit nil
### GetIdentityDiskStorageType

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetIdentityDiskStorageType() string`

GetIdentityDiskStorageType returns the IdentityDiskStorageType field if non-nil, zero value otherwise.

### GetIdentityDiskStorageTypeOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetIdentityDiskStorageTypeOk() (*string, bool)`

GetIdentityDiskStorageTypeOk returns a tuple with the IdentityDiskStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityDiskStorageType

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetIdentityDiskStorageType(v string)`

SetIdentityDiskStorageType sets IdentityDiskStorageType field to given value.

### HasIdentityDiskStorageType

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasIdentityDiskStorageType() bool`

HasIdentityDiskStorageType returns a boolean if a field has been set.

### SetIdentityDiskStorageTypeNil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetIdentityDiskStorageTypeNil(b bool)`

 SetIdentityDiskStorageTypeNil sets the value for IdentityDiskStorageType to be an explicit nil

### UnsetIdentityDiskStorageType
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetIdentityDiskStorageType()`

UnsetIdentityDiskStorageType ensures that no value is present for IdentityDiskStorageType, not even an explicit nil
### GetIsOSDiskAzureTempDisk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetIsOSDiskAzureTempDisk() bool`

GetIsOSDiskAzureTempDisk returns the IsOSDiskAzureTempDisk field if non-nil, zero value otherwise.

### GetIsOSDiskAzureTempDiskOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetIsOSDiskAzureTempDiskOk() (*bool, bool)`

GetIsOSDiskAzureTempDiskOk returns a tuple with the IsOSDiskAzureTempDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOSDiskAzureTempDisk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetIsOSDiskAzureTempDisk(v bool)`

SetIsOSDiskAzureTempDisk sets IsOSDiskAzureTempDisk field to given value.

### HasIsOSDiskAzureTempDisk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasIsOSDiskAzureTempDisk() bool`

HasIsOSDiskAzureTempDisk returns a boolean if a field has been set.

### GetIsSharedImageGalleryUsed

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetIsSharedImageGalleryUsed() bool`

GetIsSharedImageGalleryUsed returns the IsSharedImageGalleryUsed field if non-nil, zero value otherwise.

### GetIsSharedImageGalleryUsedOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetIsSharedImageGalleryUsedOk() (*bool, bool)`

GetIsSharedImageGalleryUsedOk returns a tuple with the IsSharedImageGalleryUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSharedImageGalleryUsed

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetIsSharedImageGalleryUsed(v bool)`

SetIsSharedImageGalleryUsed sets IsSharedImageGalleryUsed field to given value.

### HasIsSharedImageGalleryUsed

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasIsSharedImageGalleryUsed() bool`

HasIsSharedImageGalleryUsed returns a boolean if a field has been set.

### SetIsSharedImageGalleryUsedNil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetIsSharedImageGalleryUsedNil(b bool)`

 SetIsSharedImageGalleryUsedNil sets the value for IsSharedImageGalleryUsed to be an explicit nil

### UnsetIsSharedImageGalleryUsed
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetIsSharedImageGalleryUsed()`

UnsetIsSharedImageGalleryUsed ensures that no value is present for IsSharedImageGalleryUsed, not even an explicit nil
### GetIsUsingEphemeralOsDisk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetIsUsingEphemeralOsDisk() bool`

GetIsUsingEphemeralOsDisk returns the IsUsingEphemeralOsDisk field if non-nil, zero value otherwise.

### GetIsUsingEphemeralOsDiskOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetIsUsingEphemeralOsDiskOk() (*bool, bool)`

GetIsUsingEphemeralOsDiskOk returns a tuple with the IsUsingEphemeralOsDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingEphemeralOsDisk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetIsUsingEphemeralOsDisk(v bool)`

SetIsUsingEphemeralOsDisk sets IsUsingEphemeralOsDisk field to given value.

### HasIsUsingEphemeralOsDisk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasIsUsingEphemeralOsDisk() bool`

HasIsUsingEphemeralOsDisk returns a boolean if a field has been set.

### GetLicenseType

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### SetLicenseTypeNil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetLicenseTypeNil(b bool)`

 SetLicenseTypeNil sets the value for LicenseType to be an explicit nil

### UnsetLicenseType
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetLicenseType()`

UnsetLicenseType ensures that no value is present for LicenseType, not even an explicit nil
### GetMachineSizeId

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetMachineSizeId() string`

GetMachineSizeId returns the MachineSizeId field if non-nil, zero value otherwise.

### GetMachineSizeIdOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetMachineSizeIdOk() (*string, bool)`

GetMachineSizeIdOk returns a tuple with the MachineSizeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineSizeId

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetMachineSizeId(v string)`

SetMachineSizeId sets MachineSizeId field to given value.

### HasMachineSizeId

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasMachineSizeId() bool`

HasMachineSizeId returns a boolean if a field has been set.

### SetMachineSizeIdNil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetMachineSizeIdNil(b bool)`

 SetMachineSizeIdNil sets the value for MachineSizeId to be an explicit nil

### UnsetMachineSizeId
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetMachineSizeId()`

UnsetMachineSizeId ensures that no value is present for MachineSizeId, not even an explicit nil
### GetNetworkInterfaceIds

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetNetworkInterfaceIds() []string`

GetNetworkInterfaceIds returns the NetworkInterfaceIds field if non-nil, zero value otherwise.

### GetNetworkInterfaceIdsOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetNetworkInterfaceIdsOk() (*[]string, bool)`

GetNetworkInterfaceIdsOk returns a tuple with the NetworkInterfaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceIds

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetNetworkInterfaceIds(v []string)`

SetNetworkInterfaceIds sets NetworkInterfaceIds field to given value.

### HasNetworkInterfaceIds

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasNetworkInterfaceIds() bool`

HasNetworkInterfaceIds returns a boolean if a field has been set.

### SetNetworkInterfaceIdsNil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetNetworkInterfaceIdsNil(b bool)`

 SetNetworkInterfaceIdsNil sets the value for NetworkInterfaceIds to be an explicit nil

### UnsetNetworkInterfaceIds
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetNetworkInterfaceIds()`

UnsetNetworkInterfaceIds ensures that no value is present for NetworkInterfaceIds, not even an explicit nil
### GetOSDiskId

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetOSDiskId() string`

GetOSDiskId returns the OSDiskId field if non-nil, zero value otherwise.

### GetOSDiskIdOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetOSDiskIdOk() (*string, bool)`

GetOSDiskIdOk returns a tuple with the OSDiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOSDiskId

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetOSDiskId(v string)`

SetOSDiskId sets OSDiskId field to given value.

### HasOSDiskId

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasOSDiskId() bool`

HasOSDiskId returns a boolean if a field has been set.

### SetOSDiskIdNil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetOSDiskIdNil(b bool)`

 SetOSDiskIdNil sets the value for OSDiskId to be an explicit nil

### UnsetOSDiskId
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetOSDiskId()`

UnsetOSDiskId ensures that no value is present for OSDiskId, not even an explicit nil
### GetOsDiskStorageType

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetOsDiskStorageType() string`

GetOsDiskStorageType returns the OsDiskStorageType field if non-nil, zero value otherwise.

### GetOsDiskStorageTypeOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetOsDiskStorageTypeOk() (*string, bool)`

GetOsDiskStorageTypeOk returns a tuple with the OsDiskStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDiskStorageType

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetOsDiskStorageType(v string)`

SetOsDiskStorageType sets OsDiskStorageType field to given value.

### HasOsDiskStorageType

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasOsDiskStorageType() bool`

HasOsDiskStorageType returns a boolean if a field has been set.

### SetOsDiskStorageTypeNil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetOsDiskStorageTypeNil(b bool)`

 SetOsDiskStorageTypeNil sets the value for OsDiskStorageType to be an explicit nil

### UnsetOsDiskStorageType
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetOsDiskStorageType()`

UnsetOsDiskStorageType ensures that no value is present for OsDiskStorageType, not even an explicit nil
### GetOsType

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### SetOsTypeNil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetOsTypeNil(b bool)`

 SetOsTypeNil sets the value for OsType to be an explicit nil

### UnsetOsType
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetOsType()`

UnsetOsType ensures that no value is present for OsType, not even an explicit nil
### GetPersistOsDisk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetPersistOsDisk() bool`

GetPersistOsDisk returns the PersistOsDisk field if non-nil, zero value otherwise.

### GetPersistOsDiskOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetPersistOsDiskOk() (*bool, bool)`

GetPersistOsDiskOk returns a tuple with the PersistOsDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistOsDisk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetPersistOsDisk(v bool)`

SetPersistOsDisk sets PersistOsDisk field to given value.

### HasPersistOsDisk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasPersistOsDisk() bool`

HasPersistOsDisk returns a boolean if a field has been set.

### SetPersistOsDiskNil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetPersistOsDiskNil(b bool)`

 SetPersistOsDiskNil sets the value for PersistOsDisk to be an explicit nil

### UnsetPersistOsDisk
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetPersistOsDisk()`

UnsetPersistOsDisk ensures that no value is present for PersistOsDisk, not even an explicit nil
### GetPersistWbc

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetPersistWbc() bool`

GetPersistWbc returns the PersistWbc field if non-nil, zero value otherwise.

### GetPersistWbcOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetPersistWbcOk() (*bool, bool)`

GetPersistWbcOk returns a tuple with the PersistWbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistWbc

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetPersistWbc(v bool)`

SetPersistWbc sets PersistWbc field to given value.

### HasPersistWbc

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasPersistWbc() bool`

HasPersistWbc returns a boolean if a field has been set.

### SetPersistWbcNil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetPersistWbcNil(b bool)`

 SetPersistWbcNil sets the value for PersistWbc to be an explicit nil

### UnsetPersistWbc
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetPersistWbc()`

UnsetPersistWbc ensures that no value is present for PersistWbc, not even an explicit nil
### GetRegionId

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetResourceGroupName

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### SetResourceGroupNameNil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetResourceGroupNameNil(b bool)`

 SetResourceGroupNameNil sets the value for ResourceGroupName to be an explicit nil

### UnsetResourceGroupName
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetResourceGroupName()`

UnsetResourceGroupName ensures that no value is present for ResourceGroupName, not even an explicit nil
### GetSecurityType

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetSecurityType() string`

GetSecurityType returns the SecurityType field if non-nil, zero value otherwise.

### GetSecurityTypeOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetSecurityTypeOk() (*string, bool)`

GetSecurityTypeOk returns a tuple with the SecurityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityType

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetSecurityType(v string)`

SetSecurityType sets SecurityType field to given value.

### HasSecurityType

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasSecurityType() bool`

HasSecurityType returns a boolean if a field has been set.

### SetSecurityTypeNil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetSecurityTypeNil(b bool)`

 SetSecurityTypeNil sets the value for SecurityType to be an explicit nil

### UnsetSecurityType
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetSecurityType()`

UnsetSecurityType ensures that no value is present for SecurityType, not even an explicit nil
### GetSIGMachineToReplicaRatio

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetSIGMachineToReplicaRatio() int32`

GetSIGMachineToReplicaRatio returns the SIGMachineToReplicaRatio field if non-nil, zero value otherwise.

### GetSIGMachineToReplicaRatioOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetSIGMachineToReplicaRatioOk() (*int32, bool)`

GetSIGMachineToReplicaRatioOk returns a tuple with the SIGMachineToReplicaRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSIGMachineToReplicaRatio

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetSIGMachineToReplicaRatio(v int32)`

SetSIGMachineToReplicaRatio sets SIGMachineToReplicaRatio field to given value.

### HasSIGMachineToReplicaRatio

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasSIGMachineToReplicaRatio() bool`

HasSIGMachineToReplicaRatio returns a boolean if a field has been set.

### SetSIGMachineToReplicaRatioNil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetSIGMachineToReplicaRatioNil(b bool)`

 SetSIGMachineToReplicaRatioNil sets the value for SIGMachineToReplicaRatio to be an explicit nil

### UnsetSIGMachineToReplicaRatio
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetSIGMachineToReplicaRatio()`

UnsetSIGMachineToReplicaRatio ensures that no value is present for SIGMachineToReplicaRatio, not even an explicit nil
### GetSIGReplicaMaximum

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetSIGReplicaMaximum() int32`

GetSIGReplicaMaximum returns the SIGReplicaMaximum field if non-nil, zero value otherwise.

### GetSIGReplicaMaximumOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetSIGReplicaMaximumOk() (*int32, bool)`

GetSIGReplicaMaximumOk returns a tuple with the SIGReplicaMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSIGReplicaMaximum

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetSIGReplicaMaximum(v int32)`

SetSIGReplicaMaximum sets SIGReplicaMaximum field to given value.

### HasSIGReplicaMaximum

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasSIGReplicaMaximum() bool`

HasSIGReplicaMaximum returns a boolean if a field has been set.

### SetSIGReplicaMaximumNil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetSIGReplicaMaximumNil(b bool)`

 SetSIGReplicaMaximumNil sets the value for SIGReplicaMaximum to be an explicit nil

### UnsetSIGReplicaMaximum
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetSIGReplicaMaximum()`

UnsetSIGReplicaMaximum ensures that no value is present for SIGReplicaMaximum, not even an explicit nil
### GetStorageTypeAtShutdown

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetStorageTypeAtShutdown() string`

GetStorageTypeAtShutdown returns the StorageTypeAtShutdown field if non-nil, zero value otherwise.

### GetStorageTypeAtShutdownOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetStorageTypeAtShutdownOk() (*string, bool)`

GetStorageTypeAtShutdownOk returns a tuple with the StorageTypeAtShutdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTypeAtShutdown

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetStorageTypeAtShutdown(v string)`

SetStorageTypeAtShutdown sets StorageTypeAtShutdown field to given value.

### HasStorageTypeAtShutdown

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasStorageTypeAtShutdown() bool`

HasStorageTypeAtShutdown returns a boolean if a field has been set.

### SetStorageTypeAtShutdownNil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetStorageTypeAtShutdownNil(b bool)`

 SetStorageTypeAtShutdownNil sets the value for StorageTypeAtShutdown to be an explicit nil

### UnsetStorageTypeAtShutdown
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetStorageTypeAtShutdown()`

UnsetStorageTypeAtShutdown ensures that no value is present for StorageTypeAtShutdown, not even an explicit nil
### GetSupportsHibernation

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetSupportsHibernation() bool`

GetSupportsHibernation returns the SupportsHibernation field if non-nil, zero value otherwise.

### GetSupportsHibernationOk

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) GetSupportsHibernationOk() (*bool, bool)`

GetSupportsHibernationOk returns a tuple with the SupportsHibernation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsHibernation

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetSupportsHibernation(v bool)`

SetSupportsHibernation sets SupportsHibernation field to given value.

### HasSupportsHibernation

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) HasSupportsHibernation() bool`

HasSupportsHibernation returns a boolean if a field has been set.

### SetSupportsHibernationNil

`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) SetSupportsHibernationNil(b bool)`

 SetSupportsHibernationNil sets the value for SupportsHibernation to be an explicit nil

### UnsetSupportsHibernation
`func (o *ProvisionedVirtualMachineCustomVmDataResponseModel) UnsetSupportsHibernation()`

UnsetSupportsHibernation ensures that no value is present for SupportsHibernation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


