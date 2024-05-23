# CreateMachineCatalogProvisioningSchemeRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePool** | Pointer to **NullableString** | The HostingUnit name or Id in which machine catalog locates | [optional] 
**MasterImagePath** | Pointer to **NullableString** | The path in the resource pool to the virtual machine snapshot or VM template that will be used. This identifies the hard disk to be used and the default values for the memory and processors. This must be a path to a Snapshot or Template item in the resource pool to which the Machine Catalog is associated. | [optional] 
**AssignImageVersionToProvisioningScheme** | Pointer to [**AssignImageVersionToProvisioningSchemeRequestModel**](AssignImageVersionToProvisioningSchemeRequestModel.md) |  | [optional] 
**MachineProfilePath** | Pointer to **NullableString** | The path in the resource pool to the virtual machine template that will be used. This identifies the VM template to be used and the default values for the tags, virtual machine size, boot diagnostics, host cache property of OS disk, accelerated networking and availability zone. This must be a path to a Virtual machine or Template item in the resource pool to which the Machine Catalog is associated. | [optional] 
**MasterImageNote** | Pointer to **NullableString** | The note for the master image. | [optional] 
**CpuCount** | Pointer to **NullableInt32** | The number of processors that virtual machines created from the provisioning scheme should use. | [optional] 
**CoresPerCpuCount** | Pointer to **NullableInt32** | The number of cores per processor that virtual machines created from the provisioning scheme should use. | [optional] 
**MemoryMB** | Pointer to **NullableInt32** | The maximum amount of memory that virtual machines created from the provisioning scheme should use. | [optional] 
**UseWriteBackCache** | Pointer to **NullableBool** | Indicates whether or not write back cache is enabled for the VMs created from this provisioning scheme. Optional; default is &#x60;false&#x60;. | [optional] [default to false]
**WriteBackCacheDiskSizeGB** | Pointer to **NullableInt32** | The size in GB of any temporary storage disk used by the write back cache. Should be used in conjunction with WriteBackCacheMemorySizeMB. | [optional] 
**WriteBackCacheMemorySizeMB** | Pointer to **NullableInt32** | The size in MB of any write back cache if required. Should be used in conjunction with WriteBackCacheDiskSizeGB. | [optional] 
**WriteBackCacheDriveLetter** | Pointer to **string** | The drive letter assigned for write back cache disk when UseWriteBackCache is &#x60;true&#x60;. | [optional] 
**PrepareImage** | Pointer to **NullableBool** | Indicates whether image preparation should be performed on this provisioning scheme. Optional; default is &#x60;true&#x60;. | [optional] [default to true]
**NetworkMapping** | Pointer to [**[]NetworkMapRequestModel**](NetworkMapRequestModel.md) | Specifies how the attached NICs are mapped to networks.  If this parameter is omitted, provisioned VMs are created with a single NIC, which is mapped to the default network in the hypervisor resource pool.  If this parameter is supplied, machines are created with the number of NICs specified in the map, and each NIC is attached to the specified network. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Metadata for provisioning scheme | [optional] 
**ServiceOfferingPath** | Pointer to **NullableString** | The hypervisor resource path of the Cloud service offering to use when creating machines. | [optional] 
**SecurityGroups** | Pointer to **[]string** | The hypervisor resource path(s) of the Cloud security group(s) to use when creating machines. | [optional] 
**DedicatedTenancy** | Pointer to **NullableBool** | Indicates whether to use dedicated tenancy when creating machines in Cloud Hypervisors. Optional; default is &#x60;false&#x60;. | [optional] [default to false]
**TenancyType** | Pointer to [**TenancyType**](TenancyType.md) |  | [optional] 
**AzureAdJoinType** | Pointer to [**AzureAdJoinType**](AzureAdJoinType.md) |  | [optional] 
**IdentityType** | Pointer to [**IdentityType**](IdentityType.md) |  | [optional] 
**DeviceManagementType** | Pointer to [**DeviceManagementType**](DeviceManagementType.md) |  | [optional] 
**AzureADSecurityGroupName** | Pointer to **NullableString** | Name of Azure AD security group to be created. | [optional] 
**AzureADSecurityGroupNestId** | Pointer to **NullableString** | The object id of the Azure AD security group to nest.  Only valid for Azure. Mandatory when to nest created security group. | [optional] 
**AzureADSecurityGroupNestName** | Pointer to **NullableString** | The object name of the Azure AD security group to nest. Only valid for Azure. If specified, AzureADSecurityGroupNestId is required. | [optional] 
**AzureADTenantId** | Pointer to **NullableString** | TenantId of Azure Directory. | [optional] 
**CustomProperties** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The properties of the provisioning scheme that are specific to the target hosting infrastructure. | [optional] 
**CustomPropertiesInString** | Pointer to **NullableString** | The properties of the provisioning scheme in a single string, that are specific to the target hosting infrastructure. This is an optional method to set the custom properties which are not in the form of an array. | [optional] 
**ResetAdministratorPasswords** | Pointer to **NullableBool** | Indicates whether to reset the password for the administrator accounts on provisioned machines. Optional; default is &#x60;false&#x60;. | [optional] [default to false]
**UseFullDiskCloneProvisioning** | Pointer to **NullableBool** | Indicates that the virtual machines created from the provisioning scheme should be created using the dedicated full disk clone feature. Optional; default is &#x60;false&#x60;. | [optional] [default to false]
**WorkGroupMachines** | Pointer to **NullableBool** | Indicates whether to provision workgroup virtual machines e.g. non-domain joined. Optional; default is &#x60;false&#x60;. | [optional] [default to false]
**NumTotalMachines** | Pointer to **NullableInt32** | Number of machines to provision initially.  Optional; default is 1.  It is valid to request 0 machines initially. | [optional] [default to 1]
**MachineAccountCreationRules** | Pointer to [**MachineAccountCreationRulesRequestModel**](MachineAccountCreationRulesRequestModel.md) |  | [optional] 
**AvailableMachineAccounts** | Pointer to [**[]MachineAccountRequestModel**](MachineAccountRequestModel.md) | List of Active Directory machine accounts that are to be used when machines are provisioned. | [optional] 
**PVSSite** | Pointer to **NullableString** | PVS site id. | [optional] 
**PVSVDisk** | Pointer to **NullableString** | PVS vDisk id. | [optional] 

## Methods

### NewCreateMachineCatalogProvisioningSchemeRequestModel

`func NewCreateMachineCatalogProvisioningSchemeRequestModel() *CreateMachineCatalogProvisioningSchemeRequestModel`

NewCreateMachineCatalogProvisioningSchemeRequestModel instantiates a new CreateMachineCatalogProvisioningSchemeRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMachineCatalogProvisioningSchemeRequestModelWithDefaults

`func NewCreateMachineCatalogProvisioningSchemeRequestModelWithDefaults() *CreateMachineCatalogProvisioningSchemeRequestModel`

NewCreateMachineCatalogProvisioningSchemeRequestModelWithDefaults instantiates a new CreateMachineCatalogProvisioningSchemeRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePool

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### SetResourcePoolNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetResourcePoolNil(b bool)`

 SetResourcePoolNil sets the value for ResourcePool to be an explicit nil

### UnsetResourcePool
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetResourcePool()`

UnsetResourcePool ensures that no value is present for ResourcePool, not even an explicit nil
### GetMasterImagePath

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetMasterImagePath() string`

GetMasterImagePath returns the MasterImagePath field if non-nil, zero value otherwise.

### GetMasterImagePathOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetMasterImagePathOk() (*string, bool)`

GetMasterImagePathOk returns a tuple with the MasterImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterImagePath

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetMasterImagePath(v string)`

SetMasterImagePath sets MasterImagePath field to given value.

### HasMasterImagePath

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasMasterImagePath() bool`

HasMasterImagePath returns a boolean if a field has been set.

### SetMasterImagePathNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetMasterImagePathNil(b bool)`

 SetMasterImagePathNil sets the value for MasterImagePath to be an explicit nil

### UnsetMasterImagePath
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetMasterImagePath()`

UnsetMasterImagePath ensures that no value is present for MasterImagePath, not even an explicit nil
### GetAssignImageVersionToProvisioningScheme

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetAssignImageVersionToProvisioningScheme() AssignImageVersionToProvisioningSchemeRequestModel`

GetAssignImageVersionToProvisioningScheme returns the AssignImageVersionToProvisioningScheme field if non-nil, zero value otherwise.

### GetAssignImageVersionToProvisioningSchemeOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetAssignImageVersionToProvisioningSchemeOk() (*AssignImageVersionToProvisioningSchemeRequestModel, bool)`

GetAssignImageVersionToProvisioningSchemeOk returns a tuple with the AssignImageVersionToProvisioningScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignImageVersionToProvisioningScheme

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetAssignImageVersionToProvisioningScheme(v AssignImageVersionToProvisioningSchemeRequestModel)`

SetAssignImageVersionToProvisioningScheme sets AssignImageVersionToProvisioningScheme field to given value.

### HasAssignImageVersionToProvisioningScheme

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasAssignImageVersionToProvisioningScheme() bool`

HasAssignImageVersionToProvisioningScheme returns a boolean if a field has been set.

### GetMachineProfilePath

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetMachineProfilePath() string`

GetMachineProfilePath returns the MachineProfilePath field if non-nil, zero value otherwise.

### GetMachineProfilePathOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetMachineProfilePathOk() (*string, bool)`

GetMachineProfilePathOk returns a tuple with the MachineProfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineProfilePath

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetMachineProfilePath(v string)`

SetMachineProfilePath sets MachineProfilePath field to given value.

### HasMachineProfilePath

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasMachineProfilePath() bool`

HasMachineProfilePath returns a boolean if a field has been set.

### SetMachineProfilePathNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetMachineProfilePathNil(b bool)`

 SetMachineProfilePathNil sets the value for MachineProfilePath to be an explicit nil

### UnsetMachineProfilePath
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetMachineProfilePath()`

UnsetMachineProfilePath ensures that no value is present for MachineProfilePath, not even an explicit nil
### GetMasterImageNote

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetMasterImageNote() string`

GetMasterImageNote returns the MasterImageNote field if non-nil, zero value otherwise.

### GetMasterImageNoteOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetMasterImageNoteOk() (*string, bool)`

GetMasterImageNoteOk returns a tuple with the MasterImageNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterImageNote

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetMasterImageNote(v string)`

SetMasterImageNote sets MasterImageNote field to given value.

### HasMasterImageNote

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasMasterImageNote() bool`

HasMasterImageNote returns a boolean if a field has been set.

### SetMasterImageNoteNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetMasterImageNoteNil(b bool)`

 SetMasterImageNoteNil sets the value for MasterImageNote to be an explicit nil

### UnsetMasterImageNote
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetMasterImageNote()`

UnsetMasterImageNote ensures that no value is present for MasterImageNote, not even an explicit nil
### GetCpuCount

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### SetCpuCountNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetCpuCountNil(b bool)`

 SetCpuCountNil sets the value for CpuCount to be an explicit nil

### UnsetCpuCount
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetCpuCount()`

UnsetCpuCount ensures that no value is present for CpuCount, not even an explicit nil
### GetCoresPerCpuCount

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetCoresPerCpuCount() int32`

GetCoresPerCpuCount returns the CoresPerCpuCount field if non-nil, zero value otherwise.

### GetCoresPerCpuCountOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetCoresPerCpuCountOk() (*int32, bool)`

GetCoresPerCpuCountOk returns a tuple with the CoresPerCpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerCpuCount

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetCoresPerCpuCount(v int32)`

SetCoresPerCpuCount sets CoresPerCpuCount field to given value.

### HasCoresPerCpuCount

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasCoresPerCpuCount() bool`

HasCoresPerCpuCount returns a boolean if a field has been set.

### SetCoresPerCpuCountNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetCoresPerCpuCountNil(b bool)`

 SetCoresPerCpuCountNil sets the value for CoresPerCpuCount to be an explicit nil

### UnsetCoresPerCpuCount
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetCoresPerCpuCount()`

UnsetCoresPerCpuCount ensures that no value is present for CoresPerCpuCount, not even an explicit nil
### GetMemoryMB

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetMemoryMB() int32`

GetMemoryMB returns the MemoryMB field if non-nil, zero value otherwise.

### GetMemoryMBOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetMemoryMBOk() (*int32, bool)`

GetMemoryMBOk returns a tuple with the MemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMB

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetMemoryMB(v int32)`

SetMemoryMB sets MemoryMB field to given value.

### HasMemoryMB

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasMemoryMB() bool`

HasMemoryMB returns a boolean if a field has been set.

### SetMemoryMBNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetMemoryMBNil(b bool)`

 SetMemoryMBNil sets the value for MemoryMB to be an explicit nil

### UnsetMemoryMB
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetMemoryMB()`

UnsetMemoryMB ensures that no value is present for MemoryMB, not even an explicit nil
### GetUseWriteBackCache

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetUseWriteBackCache() bool`

GetUseWriteBackCache returns the UseWriteBackCache field if non-nil, zero value otherwise.

### GetUseWriteBackCacheOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetUseWriteBackCacheOk() (*bool, bool)`

GetUseWriteBackCacheOk returns a tuple with the UseWriteBackCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWriteBackCache

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetUseWriteBackCache(v bool)`

SetUseWriteBackCache sets UseWriteBackCache field to given value.

### HasUseWriteBackCache

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasUseWriteBackCache() bool`

HasUseWriteBackCache returns a boolean if a field has been set.

### SetUseWriteBackCacheNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetUseWriteBackCacheNil(b bool)`

 SetUseWriteBackCacheNil sets the value for UseWriteBackCache to be an explicit nil

### UnsetUseWriteBackCache
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetUseWriteBackCache()`

UnsetUseWriteBackCache ensures that no value is present for UseWriteBackCache, not even an explicit nil
### GetWriteBackCacheDiskSizeGB

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetWriteBackCacheDiskSizeGB() int32`

GetWriteBackCacheDiskSizeGB returns the WriteBackCacheDiskSizeGB field if non-nil, zero value otherwise.

### GetWriteBackCacheDiskSizeGBOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetWriteBackCacheDiskSizeGBOk() (*int32, bool)`

GetWriteBackCacheDiskSizeGBOk returns a tuple with the WriteBackCacheDiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheDiskSizeGB

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetWriteBackCacheDiskSizeGB(v int32)`

SetWriteBackCacheDiskSizeGB sets WriteBackCacheDiskSizeGB field to given value.

### HasWriteBackCacheDiskSizeGB

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasWriteBackCacheDiskSizeGB() bool`

HasWriteBackCacheDiskSizeGB returns a boolean if a field has been set.

### SetWriteBackCacheDiskSizeGBNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetWriteBackCacheDiskSizeGBNil(b bool)`

 SetWriteBackCacheDiskSizeGBNil sets the value for WriteBackCacheDiskSizeGB to be an explicit nil

### UnsetWriteBackCacheDiskSizeGB
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetWriteBackCacheDiskSizeGB()`

UnsetWriteBackCacheDiskSizeGB ensures that no value is present for WriteBackCacheDiskSizeGB, not even an explicit nil
### GetWriteBackCacheMemorySizeMB

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetWriteBackCacheMemorySizeMB() int32`

GetWriteBackCacheMemorySizeMB returns the WriteBackCacheMemorySizeMB field if non-nil, zero value otherwise.

### GetWriteBackCacheMemorySizeMBOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetWriteBackCacheMemorySizeMBOk() (*int32, bool)`

GetWriteBackCacheMemorySizeMBOk returns a tuple with the WriteBackCacheMemorySizeMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheMemorySizeMB

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetWriteBackCacheMemorySizeMB(v int32)`

SetWriteBackCacheMemorySizeMB sets WriteBackCacheMemorySizeMB field to given value.

### HasWriteBackCacheMemorySizeMB

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasWriteBackCacheMemorySizeMB() bool`

HasWriteBackCacheMemorySizeMB returns a boolean if a field has been set.

### SetWriteBackCacheMemorySizeMBNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetWriteBackCacheMemorySizeMBNil(b bool)`

 SetWriteBackCacheMemorySizeMBNil sets the value for WriteBackCacheMemorySizeMB to be an explicit nil

### UnsetWriteBackCacheMemorySizeMB
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetWriteBackCacheMemorySizeMB()`

UnsetWriteBackCacheMemorySizeMB ensures that no value is present for WriteBackCacheMemorySizeMB, not even an explicit nil
### GetWriteBackCacheDriveLetter

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetWriteBackCacheDriveLetter() string`

GetWriteBackCacheDriveLetter returns the WriteBackCacheDriveLetter field if non-nil, zero value otherwise.

### GetWriteBackCacheDriveLetterOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetWriteBackCacheDriveLetterOk() (*string, bool)`

GetWriteBackCacheDriveLetterOk returns a tuple with the WriteBackCacheDriveLetter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheDriveLetter

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetWriteBackCacheDriveLetter(v string)`

SetWriteBackCacheDriveLetter sets WriteBackCacheDriveLetter field to given value.

### HasWriteBackCacheDriveLetter

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasWriteBackCacheDriveLetter() bool`

HasWriteBackCacheDriveLetter returns a boolean if a field has been set.

### GetPrepareImage

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetPrepareImage() bool`

GetPrepareImage returns the PrepareImage field if non-nil, zero value otherwise.

### GetPrepareImageOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetPrepareImageOk() (*bool, bool)`

GetPrepareImageOk returns a tuple with the PrepareImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepareImage

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetPrepareImage(v bool)`

SetPrepareImage sets PrepareImage field to given value.

### HasPrepareImage

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasPrepareImage() bool`

HasPrepareImage returns a boolean if a field has been set.

### SetPrepareImageNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetPrepareImageNil(b bool)`

 SetPrepareImageNil sets the value for PrepareImage to be an explicit nil

### UnsetPrepareImage
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetPrepareImage()`

UnsetPrepareImage ensures that no value is present for PrepareImage, not even an explicit nil
### GetNetworkMapping

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetNetworkMapping() []NetworkMapRequestModel`

GetNetworkMapping returns the NetworkMapping field if non-nil, zero value otherwise.

### GetNetworkMappingOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetNetworkMappingOk() (*[]NetworkMapRequestModel, bool)`

GetNetworkMappingOk returns a tuple with the NetworkMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMapping

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetNetworkMapping(v []NetworkMapRequestModel)`

SetNetworkMapping sets NetworkMapping field to given value.

### HasNetworkMapping

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasNetworkMapping() bool`

HasNetworkMapping returns a boolean if a field has been set.

### SetNetworkMappingNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetNetworkMappingNil(b bool)`

 SetNetworkMappingNil sets the value for NetworkMapping to be an explicit nil

### UnsetNetworkMapping
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetNetworkMapping()`

UnsetNetworkMapping ensures that no value is present for NetworkMapping, not even an explicit nil
### GetMetadata

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetServiceOfferingPath

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetServiceOfferingPath() string`

GetServiceOfferingPath returns the ServiceOfferingPath field if non-nil, zero value otherwise.

### GetServiceOfferingPathOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetServiceOfferingPathOk() (*string, bool)`

GetServiceOfferingPathOk returns a tuple with the ServiceOfferingPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceOfferingPath

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetServiceOfferingPath(v string)`

SetServiceOfferingPath sets ServiceOfferingPath field to given value.

### HasServiceOfferingPath

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasServiceOfferingPath() bool`

HasServiceOfferingPath returns a boolean if a field has been set.

### SetServiceOfferingPathNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetServiceOfferingPathNil(b bool)`

 SetServiceOfferingPathNil sets the value for ServiceOfferingPath to be an explicit nil

### UnsetServiceOfferingPath
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetServiceOfferingPath()`

UnsetServiceOfferingPath ensures that no value is present for ServiceOfferingPath, not even an explicit nil
### GetSecurityGroups

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetDedicatedTenancy

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetDedicatedTenancy() bool`

GetDedicatedTenancy returns the DedicatedTenancy field if non-nil, zero value otherwise.

### GetDedicatedTenancyOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetDedicatedTenancyOk() (*bool, bool)`

GetDedicatedTenancyOk returns a tuple with the DedicatedTenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedTenancy

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetDedicatedTenancy(v bool)`

SetDedicatedTenancy sets DedicatedTenancy field to given value.

### HasDedicatedTenancy

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasDedicatedTenancy() bool`

HasDedicatedTenancy returns a boolean if a field has been set.

### SetDedicatedTenancyNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetDedicatedTenancyNil(b bool)`

 SetDedicatedTenancyNil sets the value for DedicatedTenancy to be an explicit nil

### UnsetDedicatedTenancy
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetDedicatedTenancy()`

UnsetDedicatedTenancy ensures that no value is present for DedicatedTenancy, not even an explicit nil
### GetTenancyType

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetTenancyType() TenancyType`

GetTenancyType returns the TenancyType field if non-nil, zero value otherwise.

### GetTenancyTypeOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetTenancyTypeOk() (*TenancyType, bool)`

GetTenancyTypeOk returns a tuple with the TenancyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancyType

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetTenancyType(v TenancyType)`

SetTenancyType sets TenancyType field to given value.

### HasTenancyType

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasTenancyType() bool`

HasTenancyType returns a boolean if a field has been set.

### GetAzureAdJoinType

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetAzureAdJoinType() AzureAdJoinType`

GetAzureAdJoinType returns the AzureAdJoinType field if non-nil, zero value otherwise.

### GetAzureAdJoinTypeOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetAzureAdJoinTypeOk() (*AzureAdJoinType, bool)`

GetAzureAdJoinTypeOk returns a tuple with the AzureAdJoinType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAdJoinType

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetAzureAdJoinType(v AzureAdJoinType)`

SetAzureAdJoinType sets AzureAdJoinType field to given value.

### HasAzureAdJoinType

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasAzureAdJoinType() bool`

HasAzureAdJoinType returns a boolean if a field has been set.

### GetIdentityType

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetIdentityType() IdentityType`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetIdentityTypeOk() (*IdentityType, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetIdentityType(v IdentityType)`

SetIdentityType sets IdentityType field to given value.

### HasIdentityType

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasIdentityType() bool`

HasIdentityType returns a boolean if a field has been set.

### GetDeviceManagementType

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetDeviceManagementType() DeviceManagementType`

GetDeviceManagementType returns the DeviceManagementType field if non-nil, zero value otherwise.

### GetDeviceManagementTypeOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetDeviceManagementTypeOk() (*DeviceManagementType, bool)`

GetDeviceManagementTypeOk returns a tuple with the DeviceManagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceManagementType

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetDeviceManagementType(v DeviceManagementType)`

SetDeviceManagementType sets DeviceManagementType field to given value.

### HasDeviceManagementType

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasDeviceManagementType() bool`

HasDeviceManagementType returns a boolean if a field has been set.

### GetAzureADSecurityGroupName

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetAzureADSecurityGroupName() string`

GetAzureADSecurityGroupName returns the AzureADSecurityGroupName field if non-nil, zero value otherwise.

### GetAzureADSecurityGroupNameOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetAzureADSecurityGroupNameOk() (*string, bool)`

GetAzureADSecurityGroupNameOk returns a tuple with the AzureADSecurityGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureADSecurityGroupName

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetAzureADSecurityGroupName(v string)`

SetAzureADSecurityGroupName sets AzureADSecurityGroupName field to given value.

### HasAzureADSecurityGroupName

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasAzureADSecurityGroupName() bool`

HasAzureADSecurityGroupName returns a boolean if a field has been set.

### SetAzureADSecurityGroupNameNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetAzureADSecurityGroupNameNil(b bool)`

 SetAzureADSecurityGroupNameNil sets the value for AzureADSecurityGroupName to be an explicit nil

### UnsetAzureADSecurityGroupName
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetAzureADSecurityGroupName()`

UnsetAzureADSecurityGroupName ensures that no value is present for AzureADSecurityGroupName, not even an explicit nil
### GetAzureADSecurityGroupNestId

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetAzureADSecurityGroupNestId() string`

GetAzureADSecurityGroupNestId returns the AzureADSecurityGroupNestId field if non-nil, zero value otherwise.

### GetAzureADSecurityGroupNestIdOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetAzureADSecurityGroupNestIdOk() (*string, bool)`

GetAzureADSecurityGroupNestIdOk returns a tuple with the AzureADSecurityGroupNestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureADSecurityGroupNestId

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetAzureADSecurityGroupNestId(v string)`

SetAzureADSecurityGroupNestId sets AzureADSecurityGroupNestId field to given value.

### HasAzureADSecurityGroupNestId

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasAzureADSecurityGroupNestId() bool`

HasAzureADSecurityGroupNestId returns a boolean if a field has been set.

### SetAzureADSecurityGroupNestIdNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetAzureADSecurityGroupNestIdNil(b bool)`

 SetAzureADSecurityGroupNestIdNil sets the value for AzureADSecurityGroupNestId to be an explicit nil

### UnsetAzureADSecurityGroupNestId
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetAzureADSecurityGroupNestId()`

UnsetAzureADSecurityGroupNestId ensures that no value is present for AzureADSecurityGroupNestId, not even an explicit nil
### GetAzureADSecurityGroupNestName

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetAzureADSecurityGroupNestName() string`

GetAzureADSecurityGroupNestName returns the AzureADSecurityGroupNestName field if non-nil, zero value otherwise.

### GetAzureADSecurityGroupNestNameOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetAzureADSecurityGroupNestNameOk() (*string, bool)`

GetAzureADSecurityGroupNestNameOk returns a tuple with the AzureADSecurityGroupNestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureADSecurityGroupNestName

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetAzureADSecurityGroupNestName(v string)`

SetAzureADSecurityGroupNestName sets AzureADSecurityGroupNestName field to given value.

### HasAzureADSecurityGroupNestName

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasAzureADSecurityGroupNestName() bool`

HasAzureADSecurityGroupNestName returns a boolean if a field has been set.

### SetAzureADSecurityGroupNestNameNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetAzureADSecurityGroupNestNameNil(b bool)`

 SetAzureADSecurityGroupNestNameNil sets the value for AzureADSecurityGroupNestName to be an explicit nil

### UnsetAzureADSecurityGroupNestName
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetAzureADSecurityGroupNestName()`

UnsetAzureADSecurityGroupNestName ensures that no value is present for AzureADSecurityGroupNestName, not even an explicit nil
### GetAzureADTenantId

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetAzureADTenantId() string`

GetAzureADTenantId returns the AzureADTenantId field if non-nil, zero value otherwise.

### GetAzureADTenantIdOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetAzureADTenantIdOk() (*string, bool)`

GetAzureADTenantIdOk returns a tuple with the AzureADTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureADTenantId

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetAzureADTenantId(v string)`

SetAzureADTenantId sets AzureADTenantId field to given value.

### HasAzureADTenantId

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasAzureADTenantId() bool`

HasAzureADTenantId returns a boolean if a field has been set.

### SetAzureADTenantIdNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetAzureADTenantIdNil(b bool)`

 SetAzureADTenantIdNil sets the value for AzureADTenantId to be an explicit nil

### UnsetAzureADTenantId
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetAzureADTenantId()`

UnsetAzureADTenantId ensures that no value is present for AzureADTenantId, not even an explicit nil
### GetCustomProperties

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetCustomProperties() []NameValueStringPairModel`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetCustomPropertiesOk() (*[]NameValueStringPairModel, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetCustomProperties(v []NameValueStringPairModel)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
### GetCustomPropertiesInString

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetCustomPropertiesInString() string`

GetCustomPropertiesInString returns the CustomPropertiesInString field if non-nil, zero value otherwise.

### GetCustomPropertiesInStringOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetCustomPropertiesInStringOk() (*string, bool)`

GetCustomPropertiesInStringOk returns a tuple with the CustomPropertiesInString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPropertiesInString

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetCustomPropertiesInString(v string)`

SetCustomPropertiesInString sets CustomPropertiesInString field to given value.

### HasCustomPropertiesInString

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasCustomPropertiesInString() bool`

HasCustomPropertiesInString returns a boolean if a field has been set.

### SetCustomPropertiesInStringNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetCustomPropertiesInStringNil(b bool)`

 SetCustomPropertiesInStringNil sets the value for CustomPropertiesInString to be an explicit nil

### UnsetCustomPropertiesInString
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetCustomPropertiesInString()`

UnsetCustomPropertiesInString ensures that no value is present for CustomPropertiesInString, not even an explicit nil
### GetResetAdministratorPasswords

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetResetAdministratorPasswords() bool`

GetResetAdministratorPasswords returns the ResetAdministratorPasswords field if non-nil, zero value otherwise.

### GetResetAdministratorPasswordsOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetResetAdministratorPasswordsOk() (*bool, bool)`

GetResetAdministratorPasswordsOk returns a tuple with the ResetAdministratorPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetAdministratorPasswords

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetResetAdministratorPasswords(v bool)`

SetResetAdministratorPasswords sets ResetAdministratorPasswords field to given value.

### HasResetAdministratorPasswords

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasResetAdministratorPasswords() bool`

HasResetAdministratorPasswords returns a boolean if a field has been set.

### SetResetAdministratorPasswordsNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetResetAdministratorPasswordsNil(b bool)`

 SetResetAdministratorPasswordsNil sets the value for ResetAdministratorPasswords to be an explicit nil

### UnsetResetAdministratorPasswords
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetResetAdministratorPasswords()`

UnsetResetAdministratorPasswords ensures that no value is present for ResetAdministratorPasswords, not even an explicit nil
### GetUseFullDiskCloneProvisioning

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetUseFullDiskCloneProvisioning() bool`

GetUseFullDiskCloneProvisioning returns the UseFullDiskCloneProvisioning field if non-nil, zero value otherwise.

### GetUseFullDiskCloneProvisioningOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetUseFullDiskCloneProvisioningOk() (*bool, bool)`

GetUseFullDiskCloneProvisioningOk returns a tuple with the UseFullDiskCloneProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFullDiskCloneProvisioning

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetUseFullDiskCloneProvisioning(v bool)`

SetUseFullDiskCloneProvisioning sets UseFullDiskCloneProvisioning field to given value.

### HasUseFullDiskCloneProvisioning

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasUseFullDiskCloneProvisioning() bool`

HasUseFullDiskCloneProvisioning returns a boolean if a field has been set.

### SetUseFullDiskCloneProvisioningNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetUseFullDiskCloneProvisioningNil(b bool)`

 SetUseFullDiskCloneProvisioningNil sets the value for UseFullDiskCloneProvisioning to be an explicit nil

### UnsetUseFullDiskCloneProvisioning
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetUseFullDiskCloneProvisioning()`

UnsetUseFullDiskCloneProvisioning ensures that no value is present for UseFullDiskCloneProvisioning, not even an explicit nil
### GetWorkGroupMachines

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetWorkGroupMachines() bool`

GetWorkGroupMachines returns the WorkGroupMachines field if non-nil, zero value otherwise.

### GetWorkGroupMachinesOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetWorkGroupMachinesOk() (*bool, bool)`

GetWorkGroupMachinesOk returns a tuple with the WorkGroupMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkGroupMachines

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetWorkGroupMachines(v bool)`

SetWorkGroupMachines sets WorkGroupMachines field to given value.

### HasWorkGroupMachines

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasWorkGroupMachines() bool`

HasWorkGroupMachines returns a boolean if a field has been set.

### SetWorkGroupMachinesNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetWorkGroupMachinesNil(b bool)`

 SetWorkGroupMachinesNil sets the value for WorkGroupMachines to be an explicit nil

### UnsetWorkGroupMachines
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetWorkGroupMachines()`

UnsetWorkGroupMachines ensures that no value is present for WorkGroupMachines, not even an explicit nil
### GetNumTotalMachines

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetNumTotalMachines() int32`

GetNumTotalMachines returns the NumTotalMachines field if non-nil, zero value otherwise.

### GetNumTotalMachinesOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetNumTotalMachinesOk() (*int32, bool)`

GetNumTotalMachinesOk returns a tuple with the NumTotalMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTotalMachines

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetNumTotalMachines(v int32)`

SetNumTotalMachines sets NumTotalMachines field to given value.

### HasNumTotalMachines

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasNumTotalMachines() bool`

HasNumTotalMachines returns a boolean if a field has been set.

### SetNumTotalMachinesNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetNumTotalMachinesNil(b bool)`

 SetNumTotalMachinesNil sets the value for NumTotalMachines to be an explicit nil

### UnsetNumTotalMachines
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetNumTotalMachines()`

UnsetNumTotalMachines ensures that no value is present for NumTotalMachines, not even an explicit nil
### GetMachineAccountCreationRules

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetMachineAccountCreationRules() MachineAccountCreationRulesRequestModel`

GetMachineAccountCreationRules returns the MachineAccountCreationRules field if non-nil, zero value otherwise.

### GetMachineAccountCreationRulesOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetMachineAccountCreationRulesOk() (*MachineAccountCreationRulesRequestModel, bool)`

GetMachineAccountCreationRulesOk returns a tuple with the MachineAccountCreationRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineAccountCreationRules

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetMachineAccountCreationRules(v MachineAccountCreationRulesRequestModel)`

SetMachineAccountCreationRules sets MachineAccountCreationRules field to given value.

### HasMachineAccountCreationRules

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasMachineAccountCreationRules() bool`

HasMachineAccountCreationRules returns a boolean if a field has been set.

### GetAvailableMachineAccounts

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetAvailableMachineAccounts() []MachineAccountRequestModel`

GetAvailableMachineAccounts returns the AvailableMachineAccounts field if non-nil, zero value otherwise.

### GetAvailableMachineAccountsOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetAvailableMachineAccountsOk() (*[]MachineAccountRequestModel, bool)`

GetAvailableMachineAccountsOk returns a tuple with the AvailableMachineAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMachineAccounts

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetAvailableMachineAccounts(v []MachineAccountRequestModel)`

SetAvailableMachineAccounts sets AvailableMachineAccounts field to given value.

### HasAvailableMachineAccounts

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasAvailableMachineAccounts() bool`

HasAvailableMachineAccounts returns a boolean if a field has been set.

### SetAvailableMachineAccountsNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetAvailableMachineAccountsNil(b bool)`

 SetAvailableMachineAccountsNil sets the value for AvailableMachineAccounts to be an explicit nil

### UnsetAvailableMachineAccounts
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetAvailableMachineAccounts()`

UnsetAvailableMachineAccounts ensures that no value is present for AvailableMachineAccounts, not even an explicit nil
### GetPVSSite

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetPVSSite() string`

GetPVSSite returns the PVSSite field if non-nil, zero value otherwise.

### GetPVSSiteOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetPVSSiteOk() (*string, bool)`

GetPVSSiteOk returns a tuple with the PVSSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPVSSite

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetPVSSite(v string)`

SetPVSSite sets PVSSite field to given value.

### HasPVSSite

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasPVSSite() bool`

HasPVSSite returns a boolean if a field has been set.

### SetPVSSiteNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetPVSSiteNil(b bool)`

 SetPVSSiteNil sets the value for PVSSite to be an explicit nil

### UnsetPVSSite
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetPVSSite()`

UnsetPVSSite ensures that no value is present for PVSSite, not even an explicit nil
### GetPVSVDisk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetPVSVDisk() string`

GetPVSVDisk returns the PVSVDisk field if non-nil, zero value otherwise.

### GetPVSVDiskOk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) GetPVSVDiskOk() (*string, bool)`

GetPVSVDiskOk returns a tuple with the PVSVDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPVSVDisk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetPVSVDisk(v string)`

SetPVSVDisk sets PVSVDisk field to given value.

### HasPVSVDisk

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) HasPVSVDisk() bool`

HasPVSVDisk returns a boolean if a field has been set.

### SetPVSVDiskNil

`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) SetPVSVDiskNil(b bool)`

 SetPVSVDiskNil sets the value for PVSVDisk to be an explicit nil

### UnsetPVSVDisk
`func (o *CreateMachineCatalogProvisioningSchemeRequestModel) UnsetPVSVDisk()`

UnsetPVSVDisk ensures that no value is present for PVSVDisk, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


