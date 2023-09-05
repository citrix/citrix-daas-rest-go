# ProvisioningSchemeResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the provisioning scheme. | 
**Name** | **string** | Name of the provisioning scheme. | 
**CleanOnBoot** | **bool** | Indicates whether the VMs that are created will be reset to a clean state on each boot. | 
**ControllerAddresses** | **[]string** | The DNS names of the controllers associated with this provisioning scheme for Quick Deploy purposes. | 
**CpuCount** | **int32** | The number of processors that VMs will be created with when using this scheme. | 
**CoresPerCpuCount** | Pointer to **int32** | The number of cores per processor that VMs will be created with when using this scheme. This property only applies when hypervisor connection is Nutanix. | [optional] 
**CurrentDiskImage** | Pointer to [**ProvisioningSchemeResponseModelCurrentDiskImage**](ProvisioningSchemeResponseModelCurrentDiskImage.md) |  | [optional] 
**HistoricalDiskImages** | Pointer to [**[]VMImageResponseModel**](VMImageResponseModel.md) | The historical disk images including current image used to provision new machines in the machine catalog. | [optional] 
**CustomProperties** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The properties of the provisioning scheme that are specific to the target hosting infrastructure. | [optional] 
**ImageRuntimeInfo** | Pointer to [**ProvisioningSchemeResponseModelImageRuntimeInfo**](ProvisioningSchemeResponseModelImageRuntimeInfo.md) |  | [optional] 
**Warning** | Pointer to **string** | The warning message if parsing CustomProperties with error. | [optional] 
**Warnings** | Pointer to [**[]ProvisioningSchemeWarningReponseModel**](ProvisioningSchemeWarningReponseModel.md) | The warning. | [optional] 
**CustomPropertiesInString** | Pointer to **string** | The properties of the provisioning scheme that are specific to the target hosting infrastructure in string format. | [optional] 
**DedicatedTenancy** | **bool** | Whether to use dedicated tenancy when creating machines in Cloud Hypervisors. | 
**TenancyType** | Pointer to **string** | Indicates whether to use tenancy type Shared, Instance or Host when creating machines in Cloud Hypervisors. | [optional] 
**AzureAdJoinType** | Pointer to **string** | Indicates whether to use azure AD join type. | [optional] 
**AzureADTenantId** | Pointer to **string** | AzureAD Tenant Id associates with this catalog | [optional] 
**AzureAdSecurityGroupName** | Pointer to **string** | AzureAd dynamic security group associates with this catalog | [optional] 
**IdentityType** | Pointer to [**IdentityType**](IdentityType.md) |  | [optional] 
**DeviceManagementType** | Pointer to [**DeviceManagementType**](DeviceManagementType.md) |  | [optional] 
**IdentityContent** | Pointer to **string** | The identity content of identity pool. | [optional] 
**DiskSizeGB** | **int32** | Size of the VM&#39;s OS disk, in GB. | 
**GpuTypeId** | Pointer to **string** | Type of GPU used for VMs, if any. | [optional] 
**ResourcePool** | [**ProvisioningSchemeResponseModelResourcePool**](ProvisioningSchemeResponseModelResourcePool.md) |  | 
**IdentityPool** | [**ProvisioningSchemeResponseModelIdentityPool**](ProvisioningSchemeResponseModelIdentityPool.md) |  | 
**MachineCount** | **int32** | Number of machines provisioned from this provisioning scheme. | 
**MasterImage** | [**ProvisioningSchemeResponseModelMasterImage**](ProvisioningSchemeResponseModelMasterImage.md) |  | 
**MachineProfile** | Pointer to [**ImageSchemeResponseModelMachineProfile**](ImageSchemeResponseModelMachineProfile.md) |  | [optional] 
**MemoryMB** | **int32** | The maximum amount of memory that VMs will be created with when using this scheme. | 
**NetworkMaps** | [**[]NetworkMapResponseModel**](NetworkMapResponseModel.md) | Networks used by provisioned VMs. | 
**NutanixContainer** | Pointer to **string** | Nutanix container that the scheme uses. | [optional] 
**ResetAdministratorPasswords** | **bool** | Indicates whether administrator passwords are reset automatically to a random password each time machines are provisioned. | 
**SecurityGroups** | Pointer to **[]string** | Network security groups used by Cloud provisioned VMs. | [optional] 
**ServiceOffering** | Pointer to **string** | Service offering used by Cloud provisioned VMs. | [optional] 
**UseFullDiskCloneProvisioning** | **bool** | Indicates whether the machines are provisioned using the dedicated full disk clone feature. | 
**UseWriteBackCache** | **bool** | True if the scheme will use the write back cache feature. | 
**VMMetadata** | Pointer to [**ImageVersionDetailResponseModelAllOfVMMetadata**](ImageVersionDetailResponseModelAllOfVMMetadata.md) |  | [optional] 
**WorkgroupMachines** | Pointer to **bool** | True if the provisioned machines are non-domain joined | [optional] 
**WriteBackCacheDiskIndex** | Pointer to **int32** | The disk on which to place the write back cache.  &#x60;0&#x60; indicates the OS disk. | [optional] 
**WriteBackCacheDiskSizeGB** | Pointer to **int32** | The size of the write back cache disk if specified in GB. | [optional] 
**WriteBackCacheMemorySizeMB** | Pointer to **int32** | The size of the in-memory write back cache if specified in MB. | [optional] 
**IsPreviousImageLegacyVda** | **bool** | True if the previous image is installed with a legacy VDA. | 
**MachineAccountCreationRules** | Pointer to [**ProvisioningSchemeResponseModelMachineAccountCreationRules**](ProvisioningSchemeResponseModelMachineAccountCreationRules.md) |  | [optional] 
**NumAvailableMachineAccounts** | **int32** | Number of machine accounts available to be used by machines that will be provisioned in the machine catalog.  This will be the identities associated with the machine catalog where State is Available. | 
**PVSSite** | Pointer to **string** | PVS Site. | [optional] 
**PVSVDisk** | Pointer to **string** | PVS vDisk. | [optional] 
**ProvisioningSchemeType** | Pointer to **string** | Provisioning scheme type. | [optional] 

## Methods

### NewProvisioningSchemeResponseModel

`func NewProvisioningSchemeResponseModel(id string, name string, cleanOnBoot bool, controllerAddresses []string, cpuCount int32, dedicatedTenancy bool, diskSizeGB int32, resourcePool ProvisioningSchemeResponseModelResourcePool, identityPool ProvisioningSchemeResponseModelIdentityPool, machineCount int32, masterImage ProvisioningSchemeResponseModelMasterImage, memoryMB int32, networkMaps []NetworkMapResponseModel, resetAdministratorPasswords bool, useFullDiskCloneProvisioning bool, useWriteBackCache bool, isPreviousImageLegacyVda bool, numAvailableMachineAccounts int32, ) *ProvisioningSchemeResponseModel`

NewProvisioningSchemeResponseModel instantiates a new ProvisioningSchemeResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningSchemeResponseModelWithDefaults

`func NewProvisioningSchemeResponseModelWithDefaults() *ProvisioningSchemeResponseModel`

NewProvisioningSchemeResponseModelWithDefaults instantiates a new ProvisioningSchemeResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProvisioningSchemeResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvisioningSchemeResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvisioningSchemeResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProvisioningSchemeResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisioningSchemeResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisioningSchemeResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetCleanOnBoot

`func (o *ProvisioningSchemeResponseModel) GetCleanOnBoot() bool`

GetCleanOnBoot returns the CleanOnBoot field if non-nil, zero value otherwise.

### GetCleanOnBootOk

`func (o *ProvisioningSchemeResponseModel) GetCleanOnBootOk() (*bool, bool)`

GetCleanOnBootOk returns a tuple with the CleanOnBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanOnBoot

`func (o *ProvisioningSchemeResponseModel) SetCleanOnBoot(v bool)`

SetCleanOnBoot sets CleanOnBoot field to given value.


### GetControllerAddresses

`func (o *ProvisioningSchemeResponseModel) GetControllerAddresses() []string`

GetControllerAddresses returns the ControllerAddresses field if non-nil, zero value otherwise.

### GetControllerAddressesOk

`func (o *ProvisioningSchemeResponseModel) GetControllerAddressesOk() (*[]string, bool)`

GetControllerAddressesOk returns a tuple with the ControllerAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerAddresses

`func (o *ProvisioningSchemeResponseModel) SetControllerAddresses(v []string)`

SetControllerAddresses sets ControllerAddresses field to given value.


### GetCpuCount

`func (o *ProvisioningSchemeResponseModel) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *ProvisioningSchemeResponseModel) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *ProvisioningSchemeResponseModel) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.


### GetCoresPerCpuCount

`func (o *ProvisioningSchemeResponseModel) GetCoresPerCpuCount() int32`

GetCoresPerCpuCount returns the CoresPerCpuCount field if non-nil, zero value otherwise.

### GetCoresPerCpuCountOk

`func (o *ProvisioningSchemeResponseModel) GetCoresPerCpuCountOk() (*int32, bool)`

GetCoresPerCpuCountOk returns a tuple with the CoresPerCpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerCpuCount

`func (o *ProvisioningSchemeResponseModel) SetCoresPerCpuCount(v int32)`

SetCoresPerCpuCount sets CoresPerCpuCount field to given value.

### HasCoresPerCpuCount

`func (o *ProvisioningSchemeResponseModel) HasCoresPerCpuCount() bool`

HasCoresPerCpuCount returns a boolean if a field has been set.

### GetCurrentDiskImage

`func (o *ProvisioningSchemeResponseModel) GetCurrentDiskImage() ProvisioningSchemeResponseModelCurrentDiskImage`

GetCurrentDiskImage returns the CurrentDiskImage field if non-nil, zero value otherwise.

### GetCurrentDiskImageOk

`func (o *ProvisioningSchemeResponseModel) GetCurrentDiskImageOk() (*ProvisioningSchemeResponseModelCurrentDiskImage, bool)`

GetCurrentDiskImageOk returns a tuple with the CurrentDiskImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDiskImage

`func (o *ProvisioningSchemeResponseModel) SetCurrentDiskImage(v ProvisioningSchemeResponseModelCurrentDiskImage)`

SetCurrentDiskImage sets CurrentDiskImage field to given value.

### HasCurrentDiskImage

`func (o *ProvisioningSchemeResponseModel) HasCurrentDiskImage() bool`

HasCurrentDiskImage returns a boolean if a field has been set.

### GetHistoricalDiskImages

`func (o *ProvisioningSchemeResponseModel) GetHistoricalDiskImages() []VMImageResponseModel`

GetHistoricalDiskImages returns the HistoricalDiskImages field if non-nil, zero value otherwise.

### GetHistoricalDiskImagesOk

`func (o *ProvisioningSchemeResponseModel) GetHistoricalDiskImagesOk() (*[]VMImageResponseModel, bool)`

GetHistoricalDiskImagesOk returns a tuple with the HistoricalDiskImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricalDiskImages

`func (o *ProvisioningSchemeResponseModel) SetHistoricalDiskImages(v []VMImageResponseModel)`

SetHistoricalDiskImages sets HistoricalDiskImages field to given value.

### HasHistoricalDiskImages

`func (o *ProvisioningSchemeResponseModel) HasHistoricalDiskImages() bool`

HasHistoricalDiskImages returns a boolean if a field has been set.

### GetCustomProperties

`func (o *ProvisioningSchemeResponseModel) GetCustomProperties() []NameValueStringPairModel`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ProvisioningSchemeResponseModel) GetCustomPropertiesOk() (*[]NameValueStringPairModel, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ProvisioningSchemeResponseModel) SetCustomProperties(v []NameValueStringPairModel)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ProvisioningSchemeResponseModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetImageRuntimeInfo

`func (o *ProvisioningSchemeResponseModel) GetImageRuntimeInfo() ProvisioningSchemeResponseModelImageRuntimeInfo`

GetImageRuntimeInfo returns the ImageRuntimeInfo field if non-nil, zero value otherwise.

### GetImageRuntimeInfoOk

`func (o *ProvisioningSchemeResponseModel) GetImageRuntimeInfoOk() (*ProvisioningSchemeResponseModelImageRuntimeInfo, bool)`

GetImageRuntimeInfoOk returns a tuple with the ImageRuntimeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRuntimeInfo

`func (o *ProvisioningSchemeResponseModel) SetImageRuntimeInfo(v ProvisioningSchemeResponseModelImageRuntimeInfo)`

SetImageRuntimeInfo sets ImageRuntimeInfo field to given value.

### HasImageRuntimeInfo

`func (o *ProvisioningSchemeResponseModel) HasImageRuntimeInfo() bool`

HasImageRuntimeInfo returns a boolean if a field has been set.

### GetWarning

`func (o *ProvisioningSchemeResponseModel) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *ProvisioningSchemeResponseModel) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *ProvisioningSchemeResponseModel) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *ProvisioningSchemeResponseModel) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### GetWarnings

`func (o *ProvisioningSchemeResponseModel) GetWarnings() []ProvisioningSchemeWarningReponseModel`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ProvisioningSchemeResponseModel) GetWarningsOk() (*[]ProvisioningSchemeWarningReponseModel, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ProvisioningSchemeResponseModel) SetWarnings(v []ProvisioningSchemeWarningReponseModel)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ProvisioningSchemeResponseModel) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetCustomPropertiesInString

`func (o *ProvisioningSchemeResponseModel) GetCustomPropertiesInString() string`

GetCustomPropertiesInString returns the CustomPropertiesInString field if non-nil, zero value otherwise.

### GetCustomPropertiesInStringOk

`func (o *ProvisioningSchemeResponseModel) GetCustomPropertiesInStringOk() (*string, bool)`

GetCustomPropertiesInStringOk returns a tuple with the CustomPropertiesInString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPropertiesInString

`func (o *ProvisioningSchemeResponseModel) SetCustomPropertiesInString(v string)`

SetCustomPropertiesInString sets CustomPropertiesInString field to given value.

### HasCustomPropertiesInString

`func (o *ProvisioningSchemeResponseModel) HasCustomPropertiesInString() bool`

HasCustomPropertiesInString returns a boolean if a field has been set.

### GetDedicatedTenancy

`func (o *ProvisioningSchemeResponseModel) GetDedicatedTenancy() bool`

GetDedicatedTenancy returns the DedicatedTenancy field if non-nil, zero value otherwise.

### GetDedicatedTenancyOk

`func (o *ProvisioningSchemeResponseModel) GetDedicatedTenancyOk() (*bool, bool)`

GetDedicatedTenancyOk returns a tuple with the DedicatedTenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedTenancy

`func (o *ProvisioningSchemeResponseModel) SetDedicatedTenancy(v bool)`

SetDedicatedTenancy sets DedicatedTenancy field to given value.


### GetTenancyType

`func (o *ProvisioningSchemeResponseModel) GetTenancyType() string`

GetTenancyType returns the TenancyType field if non-nil, zero value otherwise.

### GetTenancyTypeOk

`func (o *ProvisioningSchemeResponseModel) GetTenancyTypeOk() (*string, bool)`

GetTenancyTypeOk returns a tuple with the TenancyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancyType

`func (o *ProvisioningSchemeResponseModel) SetTenancyType(v string)`

SetTenancyType sets TenancyType field to given value.

### HasTenancyType

`func (o *ProvisioningSchemeResponseModel) HasTenancyType() bool`

HasTenancyType returns a boolean if a field has been set.

### GetAzureAdJoinType

`func (o *ProvisioningSchemeResponseModel) GetAzureAdJoinType() string`

GetAzureAdJoinType returns the AzureAdJoinType field if non-nil, zero value otherwise.

### GetAzureAdJoinTypeOk

`func (o *ProvisioningSchemeResponseModel) GetAzureAdJoinTypeOk() (*string, bool)`

GetAzureAdJoinTypeOk returns a tuple with the AzureAdJoinType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAdJoinType

`func (o *ProvisioningSchemeResponseModel) SetAzureAdJoinType(v string)`

SetAzureAdJoinType sets AzureAdJoinType field to given value.

### HasAzureAdJoinType

`func (o *ProvisioningSchemeResponseModel) HasAzureAdJoinType() bool`

HasAzureAdJoinType returns a boolean if a field has been set.

### GetAzureADTenantId

`func (o *ProvisioningSchemeResponseModel) GetAzureADTenantId() string`

GetAzureADTenantId returns the AzureADTenantId field if non-nil, zero value otherwise.

### GetAzureADTenantIdOk

`func (o *ProvisioningSchemeResponseModel) GetAzureADTenantIdOk() (*string, bool)`

GetAzureADTenantIdOk returns a tuple with the AzureADTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureADTenantId

`func (o *ProvisioningSchemeResponseModel) SetAzureADTenantId(v string)`

SetAzureADTenantId sets AzureADTenantId field to given value.

### HasAzureADTenantId

`func (o *ProvisioningSchemeResponseModel) HasAzureADTenantId() bool`

HasAzureADTenantId returns a boolean if a field has been set.

### GetAzureAdSecurityGroupName

`func (o *ProvisioningSchemeResponseModel) GetAzureAdSecurityGroupName() string`

GetAzureAdSecurityGroupName returns the AzureAdSecurityGroupName field if non-nil, zero value otherwise.

### GetAzureAdSecurityGroupNameOk

`func (o *ProvisioningSchemeResponseModel) GetAzureAdSecurityGroupNameOk() (*string, bool)`

GetAzureAdSecurityGroupNameOk returns a tuple with the AzureAdSecurityGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAdSecurityGroupName

`func (o *ProvisioningSchemeResponseModel) SetAzureAdSecurityGroupName(v string)`

SetAzureAdSecurityGroupName sets AzureAdSecurityGroupName field to given value.

### HasAzureAdSecurityGroupName

`func (o *ProvisioningSchemeResponseModel) HasAzureAdSecurityGroupName() bool`

HasAzureAdSecurityGroupName returns a boolean if a field has been set.

### GetIdentityType

`func (o *ProvisioningSchemeResponseModel) GetIdentityType() IdentityType`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *ProvisioningSchemeResponseModel) GetIdentityTypeOk() (*IdentityType, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *ProvisioningSchemeResponseModel) SetIdentityType(v IdentityType)`

SetIdentityType sets IdentityType field to given value.

### HasIdentityType

`func (o *ProvisioningSchemeResponseModel) HasIdentityType() bool`

HasIdentityType returns a boolean if a field has been set.

### GetDeviceManagementType

`func (o *ProvisioningSchemeResponseModel) GetDeviceManagementType() DeviceManagementType`

GetDeviceManagementType returns the DeviceManagementType field if non-nil, zero value otherwise.

### GetDeviceManagementTypeOk

`func (o *ProvisioningSchemeResponseModel) GetDeviceManagementTypeOk() (*DeviceManagementType, bool)`

GetDeviceManagementTypeOk returns a tuple with the DeviceManagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceManagementType

`func (o *ProvisioningSchemeResponseModel) SetDeviceManagementType(v DeviceManagementType)`

SetDeviceManagementType sets DeviceManagementType field to given value.

### HasDeviceManagementType

`func (o *ProvisioningSchemeResponseModel) HasDeviceManagementType() bool`

HasDeviceManagementType returns a boolean if a field has been set.

### GetIdentityContent

`func (o *ProvisioningSchemeResponseModel) GetIdentityContent() string`

GetIdentityContent returns the IdentityContent field if non-nil, zero value otherwise.

### GetIdentityContentOk

`func (o *ProvisioningSchemeResponseModel) GetIdentityContentOk() (*string, bool)`

GetIdentityContentOk returns a tuple with the IdentityContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityContent

`func (o *ProvisioningSchemeResponseModel) SetIdentityContent(v string)`

SetIdentityContent sets IdentityContent field to given value.

### HasIdentityContent

`func (o *ProvisioningSchemeResponseModel) HasIdentityContent() bool`

HasIdentityContent returns a boolean if a field has been set.

### GetDiskSizeGB

`func (o *ProvisioningSchemeResponseModel) GetDiskSizeGB() int32`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *ProvisioningSchemeResponseModel) GetDiskSizeGBOk() (*int32, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *ProvisioningSchemeResponseModel) SetDiskSizeGB(v int32)`

SetDiskSizeGB sets DiskSizeGB field to given value.


### GetGpuTypeId

`func (o *ProvisioningSchemeResponseModel) GetGpuTypeId() string`

GetGpuTypeId returns the GpuTypeId field if non-nil, zero value otherwise.

### GetGpuTypeIdOk

`func (o *ProvisioningSchemeResponseModel) GetGpuTypeIdOk() (*string, bool)`

GetGpuTypeIdOk returns a tuple with the GpuTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuTypeId

`func (o *ProvisioningSchemeResponseModel) SetGpuTypeId(v string)`

SetGpuTypeId sets GpuTypeId field to given value.

### HasGpuTypeId

`func (o *ProvisioningSchemeResponseModel) HasGpuTypeId() bool`

HasGpuTypeId returns a boolean if a field has been set.

### GetResourcePool

`func (o *ProvisioningSchemeResponseModel) GetResourcePool() ProvisioningSchemeResponseModelResourcePool`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *ProvisioningSchemeResponseModel) GetResourcePoolOk() (*ProvisioningSchemeResponseModelResourcePool, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *ProvisioningSchemeResponseModel) SetResourcePool(v ProvisioningSchemeResponseModelResourcePool)`

SetResourcePool sets ResourcePool field to given value.


### GetIdentityPool

`func (o *ProvisioningSchemeResponseModel) GetIdentityPool() ProvisioningSchemeResponseModelIdentityPool`

GetIdentityPool returns the IdentityPool field if non-nil, zero value otherwise.

### GetIdentityPoolOk

`func (o *ProvisioningSchemeResponseModel) GetIdentityPoolOk() (*ProvisioningSchemeResponseModelIdentityPool, bool)`

GetIdentityPoolOk returns a tuple with the IdentityPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityPool

`func (o *ProvisioningSchemeResponseModel) SetIdentityPool(v ProvisioningSchemeResponseModelIdentityPool)`

SetIdentityPool sets IdentityPool field to given value.


### GetMachineCount

`func (o *ProvisioningSchemeResponseModel) GetMachineCount() int32`

GetMachineCount returns the MachineCount field if non-nil, zero value otherwise.

### GetMachineCountOk

`func (o *ProvisioningSchemeResponseModel) GetMachineCountOk() (*int32, bool)`

GetMachineCountOk returns a tuple with the MachineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineCount

`func (o *ProvisioningSchemeResponseModel) SetMachineCount(v int32)`

SetMachineCount sets MachineCount field to given value.


### GetMasterImage

`func (o *ProvisioningSchemeResponseModel) GetMasterImage() ProvisioningSchemeResponseModelMasterImage`

GetMasterImage returns the MasterImage field if non-nil, zero value otherwise.

### GetMasterImageOk

`func (o *ProvisioningSchemeResponseModel) GetMasterImageOk() (*ProvisioningSchemeResponseModelMasterImage, bool)`

GetMasterImageOk returns a tuple with the MasterImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterImage

`func (o *ProvisioningSchemeResponseModel) SetMasterImage(v ProvisioningSchemeResponseModelMasterImage)`

SetMasterImage sets MasterImage field to given value.


### GetMachineProfile

`func (o *ProvisioningSchemeResponseModel) GetMachineProfile() ImageSchemeResponseModelMachineProfile`

GetMachineProfile returns the MachineProfile field if non-nil, zero value otherwise.

### GetMachineProfileOk

`func (o *ProvisioningSchemeResponseModel) GetMachineProfileOk() (*ImageSchemeResponseModelMachineProfile, bool)`

GetMachineProfileOk returns a tuple with the MachineProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineProfile

`func (o *ProvisioningSchemeResponseModel) SetMachineProfile(v ImageSchemeResponseModelMachineProfile)`

SetMachineProfile sets MachineProfile field to given value.

### HasMachineProfile

`func (o *ProvisioningSchemeResponseModel) HasMachineProfile() bool`

HasMachineProfile returns a boolean if a field has been set.

### GetMemoryMB

`func (o *ProvisioningSchemeResponseModel) GetMemoryMB() int32`

GetMemoryMB returns the MemoryMB field if non-nil, zero value otherwise.

### GetMemoryMBOk

`func (o *ProvisioningSchemeResponseModel) GetMemoryMBOk() (*int32, bool)`

GetMemoryMBOk returns a tuple with the MemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMB

`func (o *ProvisioningSchemeResponseModel) SetMemoryMB(v int32)`

SetMemoryMB sets MemoryMB field to given value.


### GetNetworkMaps

`func (o *ProvisioningSchemeResponseModel) GetNetworkMaps() []NetworkMapResponseModel`

GetNetworkMaps returns the NetworkMaps field if non-nil, zero value otherwise.

### GetNetworkMapsOk

`func (o *ProvisioningSchemeResponseModel) GetNetworkMapsOk() (*[]NetworkMapResponseModel, bool)`

GetNetworkMapsOk returns a tuple with the NetworkMaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMaps

`func (o *ProvisioningSchemeResponseModel) SetNetworkMaps(v []NetworkMapResponseModel)`

SetNetworkMaps sets NetworkMaps field to given value.


### GetNutanixContainer

`func (o *ProvisioningSchemeResponseModel) GetNutanixContainer() string`

GetNutanixContainer returns the NutanixContainer field if non-nil, zero value otherwise.

### GetNutanixContainerOk

`func (o *ProvisioningSchemeResponseModel) GetNutanixContainerOk() (*string, bool)`

GetNutanixContainerOk returns a tuple with the NutanixContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNutanixContainer

`func (o *ProvisioningSchemeResponseModel) SetNutanixContainer(v string)`

SetNutanixContainer sets NutanixContainer field to given value.

### HasNutanixContainer

`func (o *ProvisioningSchemeResponseModel) HasNutanixContainer() bool`

HasNutanixContainer returns a boolean if a field has been set.

### GetResetAdministratorPasswords

`func (o *ProvisioningSchemeResponseModel) GetResetAdministratorPasswords() bool`

GetResetAdministratorPasswords returns the ResetAdministratorPasswords field if non-nil, zero value otherwise.

### GetResetAdministratorPasswordsOk

`func (o *ProvisioningSchemeResponseModel) GetResetAdministratorPasswordsOk() (*bool, bool)`

GetResetAdministratorPasswordsOk returns a tuple with the ResetAdministratorPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetAdministratorPasswords

`func (o *ProvisioningSchemeResponseModel) SetResetAdministratorPasswords(v bool)`

SetResetAdministratorPasswords sets ResetAdministratorPasswords field to given value.


### GetSecurityGroups

`func (o *ProvisioningSchemeResponseModel) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *ProvisioningSchemeResponseModel) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *ProvisioningSchemeResponseModel) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *ProvisioningSchemeResponseModel) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetServiceOffering

`func (o *ProvisioningSchemeResponseModel) GetServiceOffering() string`

GetServiceOffering returns the ServiceOffering field if non-nil, zero value otherwise.

### GetServiceOfferingOk

`func (o *ProvisioningSchemeResponseModel) GetServiceOfferingOk() (*string, bool)`

GetServiceOfferingOk returns a tuple with the ServiceOffering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceOffering

`func (o *ProvisioningSchemeResponseModel) SetServiceOffering(v string)`

SetServiceOffering sets ServiceOffering field to given value.

### HasServiceOffering

`func (o *ProvisioningSchemeResponseModel) HasServiceOffering() bool`

HasServiceOffering returns a boolean if a field has been set.

### GetUseFullDiskCloneProvisioning

`func (o *ProvisioningSchemeResponseModel) GetUseFullDiskCloneProvisioning() bool`

GetUseFullDiskCloneProvisioning returns the UseFullDiskCloneProvisioning field if non-nil, zero value otherwise.

### GetUseFullDiskCloneProvisioningOk

`func (o *ProvisioningSchemeResponseModel) GetUseFullDiskCloneProvisioningOk() (*bool, bool)`

GetUseFullDiskCloneProvisioningOk returns a tuple with the UseFullDiskCloneProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFullDiskCloneProvisioning

`func (o *ProvisioningSchemeResponseModel) SetUseFullDiskCloneProvisioning(v bool)`

SetUseFullDiskCloneProvisioning sets UseFullDiskCloneProvisioning field to given value.


### GetUseWriteBackCache

`func (o *ProvisioningSchemeResponseModel) GetUseWriteBackCache() bool`

GetUseWriteBackCache returns the UseWriteBackCache field if non-nil, zero value otherwise.

### GetUseWriteBackCacheOk

`func (o *ProvisioningSchemeResponseModel) GetUseWriteBackCacheOk() (*bool, bool)`

GetUseWriteBackCacheOk returns a tuple with the UseWriteBackCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWriteBackCache

`func (o *ProvisioningSchemeResponseModel) SetUseWriteBackCache(v bool)`

SetUseWriteBackCache sets UseWriteBackCache field to given value.


### GetVMMetadata

`func (o *ProvisioningSchemeResponseModel) GetVMMetadata() ImageVersionDetailResponseModelAllOfVMMetadata`

GetVMMetadata returns the VMMetadata field if non-nil, zero value otherwise.

### GetVMMetadataOk

`func (o *ProvisioningSchemeResponseModel) GetVMMetadataOk() (*ImageVersionDetailResponseModelAllOfVMMetadata, bool)`

GetVMMetadataOk returns a tuple with the VMMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMMetadata

`func (o *ProvisioningSchemeResponseModel) SetVMMetadata(v ImageVersionDetailResponseModelAllOfVMMetadata)`

SetVMMetadata sets VMMetadata field to given value.

### HasVMMetadata

`func (o *ProvisioningSchemeResponseModel) HasVMMetadata() bool`

HasVMMetadata returns a boolean if a field has been set.

### GetWorkgroupMachines

`func (o *ProvisioningSchemeResponseModel) GetWorkgroupMachines() bool`

GetWorkgroupMachines returns the WorkgroupMachines field if non-nil, zero value otherwise.

### GetWorkgroupMachinesOk

`func (o *ProvisioningSchemeResponseModel) GetWorkgroupMachinesOk() (*bool, bool)`

GetWorkgroupMachinesOk returns a tuple with the WorkgroupMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkgroupMachines

`func (o *ProvisioningSchemeResponseModel) SetWorkgroupMachines(v bool)`

SetWorkgroupMachines sets WorkgroupMachines field to given value.

### HasWorkgroupMachines

`func (o *ProvisioningSchemeResponseModel) HasWorkgroupMachines() bool`

HasWorkgroupMachines returns a boolean if a field has been set.

### GetWriteBackCacheDiskIndex

`func (o *ProvisioningSchemeResponseModel) GetWriteBackCacheDiskIndex() int32`

GetWriteBackCacheDiskIndex returns the WriteBackCacheDiskIndex field if non-nil, zero value otherwise.

### GetWriteBackCacheDiskIndexOk

`func (o *ProvisioningSchemeResponseModel) GetWriteBackCacheDiskIndexOk() (*int32, bool)`

GetWriteBackCacheDiskIndexOk returns a tuple with the WriteBackCacheDiskIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheDiskIndex

`func (o *ProvisioningSchemeResponseModel) SetWriteBackCacheDiskIndex(v int32)`

SetWriteBackCacheDiskIndex sets WriteBackCacheDiskIndex field to given value.

### HasWriteBackCacheDiskIndex

`func (o *ProvisioningSchemeResponseModel) HasWriteBackCacheDiskIndex() bool`

HasWriteBackCacheDiskIndex returns a boolean if a field has been set.

### GetWriteBackCacheDiskSizeGB

`func (o *ProvisioningSchemeResponseModel) GetWriteBackCacheDiskSizeGB() int32`

GetWriteBackCacheDiskSizeGB returns the WriteBackCacheDiskSizeGB field if non-nil, zero value otherwise.

### GetWriteBackCacheDiskSizeGBOk

`func (o *ProvisioningSchemeResponseModel) GetWriteBackCacheDiskSizeGBOk() (*int32, bool)`

GetWriteBackCacheDiskSizeGBOk returns a tuple with the WriteBackCacheDiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheDiskSizeGB

`func (o *ProvisioningSchemeResponseModel) SetWriteBackCacheDiskSizeGB(v int32)`

SetWriteBackCacheDiskSizeGB sets WriteBackCacheDiskSizeGB field to given value.

### HasWriteBackCacheDiskSizeGB

`func (o *ProvisioningSchemeResponseModel) HasWriteBackCacheDiskSizeGB() bool`

HasWriteBackCacheDiskSizeGB returns a boolean if a field has been set.

### GetWriteBackCacheMemorySizeMB

`func (o *ProvisioningSchemeResponseModel) GetWriteBackCacheMemorySizeMB() int32`

GetWriteBackCacheMemorySizeMB returns the WriteBackCacheMemorySizeMB field if non-nil, zero value otherwise.

### GetWriteBackCacheMemorySizeMBOk

`func (o *ProvisioningSchemeResponseModel) GetWriteBackCacheMemorySizeMBOk() (*int32, bool)`

GetWriteBackCacheMemorySizeMBOk returns a tuple with the WriteBackCacheMemorySizeMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheMemorySizeMB

`func (o *ProvisioningSchemeResponseModel) SetWriteBackCacheMemorySizeMB(v int32)`

SetWriteBackCacheMemorySizeMB sets WriteBackCacheMemorySizeMB field to given value.

### HasWriteBackCacheMemorySizeMB

`func (o *ProvisioningSchemeResponseModel) HasWriteBackCacheMemorySizeMB() bool`

HasWriteBackCacheMemorySizeMB returns a boolean if a field has been set.

### GetIsPreviousImageLegacyVda

`func (o *ProvisioningSchemeResponseModel) GetIsPreviousImageLegacyVda() bool`

GetIsPreviousImageLegacyVda returns the IsPreviousImageLegacyVda field if non-nil, zero value otherwise.

### GetIsPreviousImageLegacyVdaOk

`func (o *ProvisioningSchemeResponseModel) GetIsPreviousImageLegacyVdaOk() (*bool, bool)`

GetIsPreviousImageLegacyVdaOk returns a tuple with the IsPreviousImageLegacyVda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreviousImageLegacyVda

`func (o *ProvisioningSchemeResponseModel) SetIsPreviousImageLegacyVda(v bool)`

SetIsPreviousImageLegacyVda sets IsPreviousImageLegacyVda field to given value.


### GetMachineAccountCreationRules

`func (o *ProvisioningSchemeResponseModel) GetMachineAccountCreationRules() ProvisioningSchemeResponseModelMachineAccountCreationRules`

GetMachineAccountCreationRules returns the MachineAccountCreationRules field if non-nil, zero value otherwise.

### GetMachineAccountCreationRulesOk

`func (o *ProvisioningSchemeResponseModel) GetMachineAccountCreationRulesOk() (*ProvisioningSchemeResponseModelMachineAccountCreationRules, bool)`

GetMachineAccountCreationRulesOk returns a tuple with the MachineAccountCreationRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineAccountCreationRules

`func (o *ProvisioningSchemeResponseModel) SetMachineAccountCreationRules(v ProvisioningSchemeResponseModelMachineAccountCreationRules)`

SetMachineAccountCreationRules sets MachineAccountCreationRules field to given value.

### HasMachineAccountCreationRules

`func (o *ProvisioningSchemeResponseModel) HasMachineAccountCreationRules() bool`

HasMachineAccountCreationRules returns a boolean if a field has been set.

### GetNumAvailableMachineAccounts

`func (o *ProvisioningSchemeResponseModel) GetNumAvailableMachineAccounts() int32`

GetNumAvailableMachineAccounts returns the NumAvailableMachineAccounts field if non-nil, zero value otherwise.

### GetNumAvailableMachineAccountsOk

`func (o *ProvisioningSchemeResponseModel) GetNumAvailableMachineAccountsOk() (*int32, bool)`

GetNumAvailableMachineAccountsOk returns a tuple with the NumAvailableMachineAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAvailableMachineAccounts

`func (o *ProvisioningSchemeResponseModel) SetNumAvailableMachineAccounts(v int32)`

SetNumAvailableMachineAccounts sets NumAvailableMachineAccounts field to given value.


### GetPVSSite

`func (o *ProvisioningSchemeResponseModel) GetPVSSite() string`

GetPVSSite returns the PVSSite field if non-nil, zero value otherwise.

### GetPVSSiteOk

`func (o *ProvisioningSchemeResponseModel) GetPVSSiteOk() (*string, bool)`

GetPVSSiteOk returns a tuple with the PVSSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPVSSite

`func (o *ProvisioningSchemeResponseModel) SetPVSSite(v string)`

SetPVSSite sets PVSSite field to given value.

### HasPVSSite

`func (o *ProvisioningSchemeResponseModel) HasPVSSite() bool`

HasPVSSite returns a boolean if a field has been set.

### GetPVSVDisk

`func (o *ProvisioningSchemeResponseModel) GetPVSVDisk() string`

GetPVSVDisk returns the PVSVDisk field if non-nil, zero value otherwise.

### GetPVSVDiskOk

`func (o *ProvisioningSchemeResponseModel) GetPVSVDiskOk() (*string, bool)`

GetPVSVDiskOk returns a tuple with the PVSVDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPVSVDisk

`func (o *ProvisioningSchemeResponseModel) SetPVSVDisk(v string)`

SetPVSVDisk sets PVSVDisk field to given value.

### HasPVSVDisk

`func (o *ProvisioningSchemeResponseModel) HasPVSVDisk() bool`

HasPVSVDisk returns a boolean if a field has been set.

### GetProvisioningSchemeType

`func (o *ProvisioningSchemeResponseModel) GetProvisioningSchemeType() string`

GetProvisioningSchemeType returns the ProvisioningSchemeType field if non-nil, zero value otherwise.

### GetProvisioningSchemeTypeOk

`func (o *ProvisioningSchemeResponseModel) GetProvisioningSchemeTypeOk() (*string, bool)`

GetProvisioningSchemeTypeOk returns a tuple with the ProvisioningSchemeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeType

`func (o *ProvisioningSchemeResponseModel) SetProvisioningSchemeType(v string)`

SetProvisioningSchemeType sets ProvisioningSchemeType field to given value.

### HasProvisioningSchemeType

`func (o *ProvisioningSchemeResponseModel) HasProvisioningSchemeType() bool`

HasProvisioningSchemeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


