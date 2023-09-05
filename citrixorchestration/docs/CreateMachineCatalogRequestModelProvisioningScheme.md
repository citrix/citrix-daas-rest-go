# CreateMachineCatalogRequestModelProvisioningScheme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePool** | Pointer to **string** | The HostingUnit name or Id in which machine catalog locates | [optional] 
**MasterImagePath** | Pointer to **string** | The path in the resource pool to the virtual machine snapshot or VM template that will be used. This identifies the hard disk to be used and the default values for the memory and processors. This must be a path to a Snapshot or Template item in the resource pool to which the Machine Catalog is associated. | [optional] 
**ImageVersion** | Pointer to [**CreateMachineCatalogProvisioningSchemeRequestModelImageVersion**](CreateMachineCatalogProvisioningSchemeRequestModelImageVersion.md) |  | [optional] 
**MachineProfilePath** | Pointer to **string** | The path in the resource pool to the virtual machine template that will be used. This identifies the VM template to be used and the default values for the tags, virtual machine size, boot diagnostics, host cache property of OS disk, accelerated networking and availability zone. This must be a path to a Virtual machine or Template item in the resource pool to which the Machine Catalog is associated. | [optional] 
**MasterImageNote** | Pointer to **string** | The note for the master image. | [optional] 
**CpuCount** | Pointer to **int32** | The number of processors that virtual machines created from the provisioning scheme should use. | [optional] 
**CoresPerCpuCount** | Pointer to **int32** | The number of cores per processor that virtual machines created from the provisioning scheme should use. | [optional] 
**MemoryMB** | Pointer to **int32** | The maximum amount of memory that virtual machines created from the provisioning scheme should use. | [optional] 
**UseWriteBackCache** | Pointer to **bool** | Indicates whether or not write back cache is enabled for the VMs created from this provisioning scheme. Optional; default is &#x60;false&#x60;. | [optional] [default to false]
**WriteBackCacheDiskSizeGB** | Pointer to **int32** | The size in GB of any temporary storage disk used by the write back cache. Should be used in conjunction with WriteBackCacheMemorySizeMB. | [optional] 
**WriteBackCacheMemorySizeMB** | Pointer to **int32** | The size in MB of any write back cache if required. Should be used in conjunction with WriteBackCacheDiskSizeGB. | [optional] 
**PrepareImage** | Pointer to **bool** | Indicates whether image preparation should be performed on this provisioning scheme. Optional; default is &#x60;true&#x60;. | [optional] [default to true]
**NetworkMapping** | Pointer to [**[]NetworkMapRequestModel**](NetworkMapRequestModel.md) | Specifies how the attached NICs are mapped to networks.  If this parameter is omitted, provisioned VMs are created with a single NIC, which is mapped to the default network in the hypervisor resource pool.  If this parameter is supplied, machines are created with the number of NICs specified in the map, and each NIC is attached to the specified network. | [optional] 
**ServiceOfferingPath** | Pointer to **string** | The hypervisor resource path of the Cloud service offering to use when creating machines. | [optional] 
**SecurityGroups** | Pointer to **[]string** | The hypervisor resource path(s) of the Cloud security group(s) to use when creating machines. | [optional] 
**DedicatedTenancy** | Pointer to **bool** | Indicates whether to use dedicated tenancy when creating machines in Cloud Hypervisors. Optional; default is &#x60;false&#x60;. | [optional] [default to false]
**TenancyType** | Pointer to [**TenancyType**](TenancyType.md) |  | [optional] 
**AzureAdJoinType** | Pointer to [**AzureAdJoinType**](AzureAdJoinType.md) |  | [optional] 
**IdentityType** | Pointer to [**IdentityType**](IdentityType.md) |  | [optional] 
**DeviceManagementType** | Pointer to [**DeviceManagementType**](DeviceManagementType.md) |  | [optional] 
**AzureADSecurityGroupName** | Pointer to **string** | Name of Azure AD security group to be created. | [optional] 
**AzureADSecurityGroupNestId** | Pointer to **string** | The object id of the Azure AD security group to nest.  Only valid for Azure. | [optional] 
**AzureADTenantId** | Pointer to **string** | TenantId of Azure Directory. | [optional] 
**CustomProperties** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The properties of the provisioning scheme that are specific to the target hosting infrastructure. | [optional] 
**ResetAdministratorPasswords** | Pointer to **bool** | Indicates whether to reset the password for the administrator accounts on provisioned machines. Optional; default is &#x60;false&#x60;. | [optional] [default to false]
**UseFullDiskCloneProvisioning** | Pointer to **bool** | Indicates that the virtual machines created from the provisioning scheme should be created using the dedicated full disk clone feature. Optional; default is &#x60;false&#x60;. | [optional] [default to false]
**WorkGroupMachines** | Pointer to **bool** | Indicates whether to provision workgroup virtual machines e.g. non-domain joined. Optional; default is &#x60;false&#x60;. | [optional] [default to false]
**NumTotalMachines** | Pointer to **int32** | Number of machines to provision initially.  Optional; default is 1.  It is valid to request 0 machines initially. | [optional] [default to 1]
**MachineAccountCreationRules** | Pointer to [**CreateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules**](CreateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules.md) |  | [optional] 
**AvailableMachineAccounts** | Pointer to [**[]MachineAccountRequestModel**](MachineAccountRequestModel.md) | List of Active Directory machine accounts that are to be used when machines are provisioned. | [optional] 
**PVSSite** | Pointer to **string** | PVS site id. | [optional] 
**PVSVDisk** | Pointer to **string** | PVS vDisk id. | [optional] 

## Methods

### NewCreateMachineCatalogRequestModelProvisioningScheme

`func NewCreateMachineCatalogRequestModelProvisioningScheme() *CreateMachineCatalogRequestModelProvisioningScheme`

NewCreateMachineCatalogRequestModelProvisioningScheme instantiates a new CreateMachineCatalogRequestModelProvisioningScheme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMachineCatalogRequestModelProvisioningSchemeWithDefaults

`func NewCreateMachineCatalogRequestModelProvisioningSchemeWithDefaults() *CreateMachineCatalogRequestModelProvisioningScheme`

NewCreateMachineCatalogRequestModelProvisioningSchemeWithDefaults instantiates a new CreateMachineCatalogRequestModelProvisioningScheme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePool

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetMasterImagePath

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetMasterImagePath() string`

GetMasterImagePath returns the MasterImagePath field if non-nil, zero value otherwise.

### GetMasterImagePathOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetMasterImagePathOk() (*string, bool)`

GetMasterImagePathOk returns a tuple with the MasterImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterImagePath

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetMasterImagePath(v string)`

SetMasterImagePath sets MasterImagePath field to given value.

### HasMasterImagePath

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasMasterImagePath() bool`

HasMasterImagePath returns a boolean if a field has been set.

### GetImageVersion

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetImageVersion() CreateMachineCatalogProvisioningSchemeRequestModelImageVersion`

GetImageVersion returns the ImageVersion field if non-nil, zero value otherwise.

### GetImageVersionOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetImageVersionOk() (*CreateMachineCatalogProvisioningSchemeRequestModelImageVersion, bool)`

GetImageVersionOk returns a tuple with the ImageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersion

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetImageVersion(v CreateMachineCatalogProvisioningSchemeRequestModelImageVersion)`

SetImageVersion sets ImageVersion field to given value.

### HasImageVersion

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasImageVersion() bool`

HasImageVersion returns a boolean if a field has been set.

### GetMachineProfilePath

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetMachineProfilePath() string`

GetMachineProfilePath returns the MachineProfilePath field if non-nil, zero value otherwise.

### GetMachineProfilePathOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetMachineProfilePathOk() (*string, bool)`

GetMachineProfilePathOk returns a tuple with the MachineProfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineProfilePath

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetMachineProfilePath(v string)`

SetMachineProfilePath sets MachineProfilePath field to given value.

### HasMachineProfilePath

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasMachineProfilePath() bool`

HasMachineProfilePath returns a boolean if a field has been set.

### GetMasterImageNote

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetMasterImageNote() string`

GetMasterImageNote returns the MasterImageNote field if non-nil, zero value otherwise.

### GetMasterImageNoteOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetMasterImageNoteOk() (*string, bool)`

GetMasterImageNoteOk returns a tuple with the MasterImageNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterImageNote

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetMasterImageNote(v string)`

SetMasterImageNote sets MasterImageNote field to given value.

### HasMasterImageNote

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasMasterImageNote() bool`

HasMasterImageNote returns a boolean if a field has been set.

### GetCpuCount

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetCoresPerCpuCount

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetCoresPerCpuCount() int32`

GetCoresPerCpuCount returns the CoresPerCpuCount field if non-nil, zero value otherwise.

### GetCoresPerCpuCountOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetCoresPerCpuCountOk() (*int32, bool)`

GetCoresPerCpuCountOk returns a tuple with the CoresPerCpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerCpuCount

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetCoresPerCpuCount(v int32)`

SetCoresPerCpuCount sets CoresPerCpuCount field to given value.

### HasCoresPerCpuCount

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasCoresPerCpuCount() bool`

HasCoresPerCpuCount returns a boolean if a field has been set.

### GetMemoryMB

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetMemoryMB() int32`

GetMemoryMB returns the MemoryMB field if non-nil, zero value otherwise.

### GetMemoryMBOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetMemoryMBOk() (*int32, bool)`

GetMemoryMBOk returns a tuple with the MemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMB

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetMemoryMB(v int32)`

SetMemoryMB sets MemoryMB field to given value.

### HasMemoryMB

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasMemoryMB() bool`

HasMemoryMB returns a boolean if a field has been set.

### GetUseWriteBackCache

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetUseWriteBackCache() bool`

GetUseWriteBackCache returns the UseWriteBackCache field if non-nil, zero value otherwise.

### GetUseWriteBackCacheOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetUseWriteBackCacheOk() (*bool, bool)`

GetUseWriteBackCacheOk returns a tuple with the UseWriteBackCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWriteBackCache

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetUseWriteBackCache(v bool)`

SetUseWriteBackCache sets UseWriteBackCache field to given value.

### HasUseWriteBackCache

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasUseWriteBackCache() bool`

HasUseWriteBackCache returns a boolean if a field has been set.

### GetWriteBackCacheDiskSizeGB

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetWriteBackCacheDiskSizeGB() int32`

GetWriteBackCacheDiskSizeGB returns the WriteBackCacheDiskSizeGB field if non-nil, zero value otherwise.

### GetWriteBackCacheDiskSizeGBOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetWriteBackCacheDiskSizeGBOk() (*int32, bool)`

GetWriteBackCacheDiskSizeGBOk returns a tuple with the WriteBackCacheDiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheDiskSizeGB

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetWriteBackCacheDiskSizeGB(v int32)`

SetWriteBackCacheDiskSizeGB sets WriteBackCacheDiskSizeGB field to given value.

### HasWriteBackCacheDiskSizeGB

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasWriteBackCacheDiskSizeGB() bool`

HasWriteBackCacheDiskSizeGB returns a boolean if a field has been set.

### GetWriteBackCacheMemorySizeMB

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetWriteBackCacheMemorySizeMB() int32`

GetWriteBackCacheMemorySizeMB returns the WriteBackCacheMemorySizeMB field if non-nil, zero value otherwise.

### GetWriteBackCacheMemorySizeMBOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetWriteBackCacheMemorySizeMBOk() (*int32, bool)`

GetWriteBackCacheMemorySizeMBOk returns a tuple with the WriteBackCacheMemorySizeMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheMemorySizeMB

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetWriteBackCacheMemorySizeMB(v int32)`

SetWriteBackCacheMemorySizeMB sets WriteBackCacheMemorySizeMB field to given value.

### HasWriteBackCacheMemorySizeMB

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasWriteBackCacheMemorySizeMB() bool`

HasWriteBackCacheMemorySizeMB returns a boolean if a field has been set.

### GetPrepareImage

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetPrepareImage() bool`

GetPrepareImage returns the PrepareImage field if non-nil, zero value otherwise.

### GetPrepareImageOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetPrepareImageOk() (*bool, bool)`

GetPrepareImageOk returns a tuple with the PrepareImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepareImage

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetPrepareImage(v bool)`

SetPrepareImage sets PrepareImage field to given value.

### HasPrepareImage

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasPrepareImage() bool`

HasPrepareImage returns a boolean if a field has been set.

### GetNetworkMapping

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetNetworkMapping() []NetworkMapRequestModel`

GetNetworkMapping returns the NetworkMapping field if non-nil, zero value otherwise.

### GetNetworkMappingOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetNetworkMappingOk() (*[]NetworkMapRequestModel, bool)`

GetNetworkMappingOk returns a tuple with the NetworkMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMapping

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetNetworkMapping(v []NetworkMapRequestModel)`

SetNetworkMapping sets NetworkMapping field to given value.

### HasNetworkMapping

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasNetworkMapping() bool`

HasNetworkMapping returns a boolean if a field has been set.

### GetServiceOfferingPath

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetServiceOfferingPath() string`

GetServiceOfferingPath returns the ServiceOfferingPath field if non-nil, zero value otherwise.

### GetServiceOfferingPathOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetServiceOfferingPathOk() (*string, bool)`

GetServiceOfferingPathOk returns a tuple with the ServiceOfferingPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceOfferingPath

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetServiceOfferingPath(v string)`

SetServiceOfferingPath sets ServiceOfferingPath field to given value.

### HasServiceOfferingPath

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasServiceOfferingPath() bool`

HasServiceOfferingPath returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetDedicatedTenancy

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetDedicatedTenancy() bool`

GetDedicatedTenancy returns the DedicatedTenancy field if non-nil, zero value otherwise.

### GetDedicatedTenancyOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetDedicatedTenancyOk() (*bool, bool)`

GetDedicatedTenancyOk returns a tuple with the DedicatedTenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedTenancy

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetDedicatedTenancy(v bool)`

SetDedicatedTenancy sets DedicatedTenancy field to given value.

### HasDedicatedTenancy

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasDedicatedTenancy() bool`

HasDedicatedTenancy returns a boolean if a field has been set.

### GetTenancyType

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetTenancyType() TenancyType`

GetTenancyType returns the TenancyType field if non-nil, zero value otherwise.

### GetTenancyTypeOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetTenancyTypeOk() (*TenancyType, bool)`

GetTenancyTypeOk returns a tuple with the TenancyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancyType

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetTenancyType(v TenancyType)`

SetTenancyType sets TenancyType field to given value.

### HasTenancyType

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasTenancyType() bool`

HasTenancyType returns a boolean if a field has been set.

### GetAzureAdJoinType

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetAzureAdJoinType() AzureAdJoinType`

GetAzureAdJoinType returns the AzureAdJoinType field if non-nil, zero value otherwise.

### GetAzureAdJoinTypeOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetAzureAdJoinTypeOk() (*AzureAdJoinType, bool)`

GetAzureAdJoinTypeOk returns a tuple with the AzureAdJoinType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAdJoinType

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetAzureAdJoinType(v AzureAdJoinType)`

SetAzureAdJoinType sets AzureAdJoinType field to given value.

### HasAzureAdJoinType

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasAzureAdJoinType() bool`

HasAzureAdJoinType returns a boolean if a field has been set.

### GetIdentityType

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetIdentityType() IdentityType`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetIdentityTypeOk() (*IdentityType, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetIdentityType(v IdentityType)`

SetIdentityType sets IdentityType field to given value.

### HasIdentityType

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasIdentityType() bool`

HasIdentityType returns a boolean if a field has been set.

### GetDeviceManagementType

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetDeviceManagementType() DeviceManagementType`

GetDeviceManagementType returns the DeviceManagementType field if non-nil, zero value otherwise.

### GetDeviceManagementTypeOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetDeviceManagementTypeOk() (*DeviceManagementType, bool)`

GetDeviceManagementTypeOk returns a tuple with the DeviceManagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceManagementType

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetDeviceManagementType(v DeviceManagementType)`

SetDeviceManagementType sets DeviceManagementType field to given value.

### HasDeviceManagementType

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasDeviceManagementType() bool`

HasDeviceManagementType returns a boolean if a field has been set.

### GetAzureADSecurityGroupName

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetAzureADSecurityGroupName() string`

GetAzureADSecurityGroupName returns the AzureADSecurityGroupName field if non-nil, zero value otherwise.

### GetAzureADSecurityGroupNameOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetAzureADSecurityGroupNameOk() (*string, bool)`

GetAzureADSecurityGroupNameOk returns a tuple with the AzureADSecurityGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureADSecurityGroupName

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetAzureADSecurityGroupName(v string)`

SetAzureADSecurityGroupName sets AzureADSecurityGroupName field to given value.

### HasAzureADSecurityGroupName

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasAzureADSecurityGroupName() bool`

HasAzureADSecurityGroupName returns a boolean if a field has been set.

### GetAzureADSecurityGroupNestId

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetAzureADSecurityGroupNestId() string`

GetAzureADSecurityGroupNestId returns the AzureADSecurityGroupNestId field if non-nil, zero value otherwise.

### GetAzureADSecurityGroupNestIdOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetAzureADSecurityGroupNestIdOk() (*string, bool)`

GetAzureADSecurityGroupNestIdOk returns a tuple with the AzureADSecurityGroupNestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureADSecurityGroupNestId

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetAzureADSecurityGroupNestId(v string)`

SetAzureADSecurityGroupNestId sets AzureADSecurityGroupNestId field to given value.

### HasAzureADSecurityGroupNestId

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasAzureADSecurityGroupNestId() bool`

HasAzureADSecurityGroupNestId returns a boolean if a field has been set.

### GetAzureADTenantId

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetAzureADTenantId() string`

GetAzureADTenantId returns the AzureADTenantId field if non-nil, zero value otherwise.

### GetAzureADTenantIdOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetAzureADTenantIdOk() (*string, bool)`

GetAzureADTenantIdOk returns a tuple with the AzureADTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureADTenantId

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetAzureADTenantId(v string)`

SetAzureADTenantId sets AzureADTenantId field to given value.

### HasAzureADTenantId

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasAzureADTenantId() bool`

HasAzureADTenantId returns a boolean if a field has been set.

### GetCustomProperties

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetCustomProperties() []NameValueStringPairModel`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetCustomPropertiesOk() (*[]NameValueStringPairModel, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetCustomProperties(v []NameValueStringPairModel)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetResetAdministratorPasswords

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetResetAdministratorPasswords() bool`

GetResetAdministratorPasswords returns the ResetAdministratorPasswords field if non-nil, zero value otherwise.

### GetResetAdministratorPasswordsOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetResetAdministratorPasswordsOk() (*bool, bool)`

GetResetAdministratorPasswordsOk returns a tuple with the ResetAdministratorPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetAdministratorPasswords

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetResetAdministratorPasswords(v bool)`

SetResetAdministratorPasswords sets ResetAdministratorPasswords field to given value.

### HasResetAdministratorPasswords

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasResetAdministratorPasswords() bool`

HasResetAdministratorPasswords returns a boolean if a field has been set.

### GetUseFullDiskCloneProvisioning

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetUseFullDiskCloneProvisioning() bool`

GetUseFullDiskCloneProvisioning returns the UseFullDiskCloneProvisioning field if non-nil, zero value otherwise.

### GetUseFullDiskCloneProvisioningOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetUseFullDiskCloneProvisioningOk() (*bool, bool)`

GetUseFullDiskCloneProvisioningOk returns a tuple with the UseFullDiskCloneProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFullDiskCloneProvisioning

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetUseFullDiskCloneProvisioning(v bool)`

SetUseFullDiskCloneProvisioning sets UseFullDiskCloneProvisioning field to given value.

### HasUseFullDiskCloneProvisioning

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasUseFullDiskCloneProvisioning() bool`

HasUseFullDiskCloneProvisioning returns a boolean if a field has been set.

### GetWorkGroupMachines

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetWorkGroupMachines() bool`

GetWorkGroupMachines returns the WorkGroupMachines field if non-nil, zero value otherwise.

### GetWorkGroupMachinesOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetWorkGroupMachinesOk() (*bool, bool)`

GetWorkGroupMachinesOk returns a tuple with the WorkGroupMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkGroupMachines

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetWorkGroupMachines(v bool)`

SetWorkGroupMachines sets WorkGroupMachines field to given value.

### HasWorkGroupMachines

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasWorkGroupMachines() bool`

HasWorkGroupMachines returns a boolean if a field has been set.

### GetNumTotalMachines

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetNumTotalMachines() int32`

GetNumTotalMachines returns the NumTotalMachines field if non-nil, zero value otherwise.

### GetNumTotalMachinesOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetNumTotalMachinesOk() (*int32, bool)`

GetNumTotalMachinesOk returns a tuple with the NumTotalMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTotalMachines

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetNumTotalMachines(v int32)`

SetNumTotalMachines sets NumTotalMachines field to given value.

### HasNumTotalMachines

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasNumTotalMachines() bool`

HasNumTotalMachines returns a boolean if a field has been set.

### GetMachineAccountCreationRules

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetMachineAccountCreationRules() CreateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules`

GetMachineAccountCreationRules returns the MachineAccountCreationRules field if non-nil, zero value otherwise.

### GetMachineAccountCreationRulesOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetMachineAccountCreationRulesOk() (*CreateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules, bool)`

GetMachineAccountCreationRulesOk returns a tuple with the MachineAccountCreationRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineAccountCreationRules

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetMachineAccountCreationRules(v CreateMachineCatalogProvisioningSchemeRequestModelMachineAccountCreationRules)`

SetMachineAccountCreationRules sets MachineAccountCreationRules field to given value.

### HasMachineAccountCreationRules

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasMachineAccountCreationRules() bool`

HasMachineAccountCreationRules returns a boolean if a field has been set.

### GetAvailableMachineAccounts

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetAvailableMachineAccounts() []MachineAccountRequestModel`

GetAvailableMachineAccounts returns the AvailableMachineAccounts field if non-nil, zero value otherwise.

### GetAvailableMachineAccountsOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetAvailableMachineAccountsOk() (*[]MachineAccountRequestModel, bool)`

GetAvailableMachineAccountsOk returns a tuple with the AvailableMachineAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMachineAccounts

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetAvailableMachineAccounts(v []MachineAccountRequestModel)`

SetAvailableMachineAccounts sets AvailableMachineAccounts field to given value.

### HasAvailableMachineAccounts

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasAvailableMachineAccounts() bool`

HasAvailableMachineAccounts returns a boolean if a field has been set.

### GetPVSSite

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetPVSSite() string`

GetPVSSite returns the PVSSite field if non-nil, zero value otherwise.

### GetPVSSiteOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetPVSSiteOk() (*string, bool)`

GetPVSSiteOk returns a tuple with the PVSSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPVSSite

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetPVSSite(v string)`

SetPVSSite sets PVSSite field to given value.

### HasPVSSite

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasPVSSite() bool`

HasPVSSite returns a boolean if a field has been set.

### GetPVSVDisk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetPVSVDisk() string`

GetPVSVDisk returns the PVSVDisk field if non-nil, zero value otherwise.

### GetPVSVDiskOk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) GetPVSVDiskOk() (*string, bool)`

GetPVSVDiskOk returns a tuple with the PVSVDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPVSVDisk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) SetPVSVDisk(v string)`

SetPVSVDisk sets PVSVDisk field to given value.

### HasPVSVDisk

`func (o *CreateMachineCatalogRequestModelProvisioningScheme) HasPVSVDisk() bool`

HasPVSVDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


