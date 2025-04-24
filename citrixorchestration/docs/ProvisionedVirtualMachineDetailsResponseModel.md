# ProvisionedVirtualMachineDetailsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Provisioned Machine Name. | [optional] 
**ProvisioningSchemeName** | Pointer to **NullableString** | Provisioning Scheme Name. | [optional] 
**ProvisioningSchemeId** | Pointer to **NullableString** | Provisioning Scheme Id. | [optional] 
**CpuCount** | Pointer to **int32** | Cpu Count. | [optional] 
**MemoryMB** | Pointer to **int32** | Memory in MB. | [optional] 
**Persistency** | Pointer to **bool** | Persistency. | [optional] 
**IdentityType** | Pointer to **NullableString** | Identity Type. | [optional] 
**ProvisioningSchemeVersion** | Pointer to **int32** | Provisioning Scheme Version. | [optional] 
**ImageOutOfDate** | Pointer to **bool** | Image Out Of Date. | [optional] 
**LastBootTime** | Pointer to **time.Time** | Last Boot Time. | [optional] 
**UseWriteBackCache** | Pointer to **bool** | Use Write Back Cache. | [optional] 
**WriteBackCacheDiskSize** | Pointer to **int32** | Write Back Cache Disk Size. | [optional] 
**WriteBackCacheDriveLetter** | Pointer to **string** | Write Back Cache Drive Letter. | [optional] 
**WriteBackCacheMemorySize** | Pointer to **int32** | Write Back Cache Memory Size. | [optional] 
**Identities** | Pointer to **NullableString** | Identities. | [optional] 
**ProvVMConfigurationVersion** | Pointer to **NullableInt32** | Provisioned VM Configuration Version. | [optional] 
**WindowsActivationType** | Pointer to **NullableString** | Windows Activation Type. | [optional] 
**DiskInfo** | Pointer to **NullableString** | Disk Information. | [optional] 
**ImageAndTemplateInfo** | Pointer to **NullableString** | Image and Template Information. | [optional] 
**MCSIOInfo** | Pointer to **NullableString** | MCS IO information. | [optional] 
**NetworkInfo** | Pointer to **NullableString** | Network Information. | [optional] 
**SecurityInfo** | Pointer to **NullableString** | Security Information. | [optional] 
**VMInfo** | Pointer to **NullableString** | Provisioned VM Information. | [optional] 
**ResourcePool** | Pointer to [**HypervisorResourcePoolRefResponseModel**](HypervisorResourcePoolRefResponseModel.md) |  | [optional] 
**ConfigurationResultantSet** | Pointer to [**ProvisionedVirtualMachineConfigurationResultantSetResponseModel**](ProvisionedVirtualMachineConfigurationResultantSetResponseModel.md) |  | [optional] 
**FailureReason** | Pointer to **NullableString** | Failure reason of power action. | [optional] 
**FailedActionCompletionTime** | Pointer to **NullableString** | Time when the power action failed. | [optional] 

## Methods

### NewProvisionedVirtualMachineDetailsResponseModel

`func NewProvisionedVirtualMachineDetailsResponseModel() *ProvisionedVirtualMachineDetailsResponseModel`

NewProvisionedVirtualMachineDetailsResponseModel instantiates a new ProvisionedVirtualMachineDetailsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionedVirtualMachineDetailsResponseModelWithDefaults

`func NewProvisionedVirtualMachineDetailsResponseModelWithDefaults() *ProvisionedVirtualMachineDetailsResponseModel`

NewProvisionedVirtualMachineDetailsResponseModelWithDefaults instantiates a new ProvisionedVirtualMachineDetailsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProvisioningSchemeName

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetProvisioningSchemeName() string`

GetProvisioningSchemeName returns the ProvisioningSchemeName field if non-nil, zero value otherwise.

### GetProvisioningSchemeNameOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetProvisioningSchemeNameOk() (*string, bool)`

GetProvisioningSchemeNameOk returns a tuple with the ProvisioningSchemeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeName

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetProvisioningSchemeName(v string)`

SetProvisioningSchemeName sets ProvisioningSchemeName field to given value.

### HasProvisioningSchemeName

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasProvisioningSchemeName() bool`

HasProvisioningSchemeName returns a boolean if a field has been set.

### SetProvisioningSchemeNameNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetProvisioningSchemeNameNil(b bool)`

 SetProvisioningSchemeNameNil sets the value for ProvisioningSchemeName to be an explicit nil

### UnsetProvisioningSchemeName
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetProvisioningSchemeName()`

UnsetProvisioningSchemeName ensures that no value is present for ProvisioningSchemeName, not even an explicit nil
### GetProvisioningSchemeId

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetProvisioningSchemeId() string`

GetProvisioningSchemeId returns the ProvisioningSchemeId field if non-nil, zero value otherwise.

### GetProvisioningSchemeIdOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetProvisioningSchemeIdOk() (*string, bool)`

GetProvisioningSchemeIdOk returns a tuple with the ProvisioningSchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeId

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetProvisioningSchemeId(v string)`

SetProvisioningSchemeId sets ProvisioningSchemeId field to given value.

### HasProvisioningSchemeId

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasProvisioningSchemeId() bool`

HasProvisioningSchemeId returns a boolean if a field has been set.

### SetProvisioningSchemeIdNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetProvisioningSchemeIdNil(b bool)`

 SetProvisioningSchemeIdNil sets the value for ProvisioningSchemeId to be an explicit nil

### UnsetProvisioningSchemeId
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetProvisioningSchemeId()`

UnsetProvisioningSchemeId ensures that no value is present for ProvisioningSchemeId, not even an explicit nil
### GetCpuCount

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetMemoryMB

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetMemoryMB() int32`

GetMemoryMB returns the MemoryMB field if non-nil, zero value otherwise.

### GetMemoryMBOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetMemoryMBOk() (*int32, bool)`

GetMemoryMBOk returns a tuple with the MemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMB

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetMemoryMB(v int32)`

SetMemoryMB sets MemoryMB field to given value.

### HasMemoryMB

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasMemoryMB() bool`

HasMemoryMB returns a boolean if a field has been set.

### GetPersistency

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetPersistency() bool`

GetPersistency returns the Persistency field if non-nil, zero value otherwise.

### GetPersistencyOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetPersistencyOk() (*bool, bool)`

GetPersistencyOk returns a tuple with the Persistency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistency

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetPersistency(v bool)`

SetPersistency sets Persistency field to given value.

### HasPersistency

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasPersistency() bool`

HasPersistency returns a boolean if a field has been set.

### GetIdentityType

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetIdentityType() string`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetIdentityTypeOk() (*string, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetIdentityType(v string)`

SetIdentityType sets IdentityType field to given value.

### HasIdentityType

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasIdentityType() bool`

HasIdentityType returns a boolean if a field has been set.

### SetIdentityTypeNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetIdentityTypeNil(b bool)`

 SetIdentityTypeNil sets the value for IdentityType to be an explicit nil

### UnsetIdentityType
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetIdentityType()`

UnsetIdentityType ensures that no value is present for IdentityType, not even an explicit nil
### GetProvisioningSchemeVersion

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetProvisioningSchemeVersion() int32`

GetProvisioningSchemeVersion returns the ProvisioningSchemeVersion field if non-nil, zero value otherwise.

### GetProvisioningSchemeVersionOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetProvisioningSchemeVersionOk() (*int32, bool)`

GetProvisioningSchemeVersionOk returns a tuple with the ProvisioningSchemeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeVersion

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetProvisioningSchemeVersion(v int32)`

SetProvisioningSchemeVersion sets ProvisioningSchemeVersion field to given value.

### HasProvisioningSchemeVersion

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasProvisioningSchemeVersion() bool`

HasProvisioningSchemeVersion returns a boolean if a field has been set.

### GetImageOutOfDate

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetImageOutOfDate() bool`

GetImageOutOfDate returns the ImageOutOfDate field if non-nil, zero value otherwise.

### GetImageOutOfDateOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetImageOutOfDateOk() (*bool, bool)`

GetImageOutOfDateOk returns a tuple with the ImageOutOfDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOutOfDate

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetImageOutOfDate(v bool)`

SetImageOutOfDate sets ImageOutOfDate field to given value.

### HasImageOutOfDate

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasImageOutOfDate() bool`

HasImageOutOfDate returns a boolean if a field has been set.

### GetLastBootTime

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetLastBootTime() time.Time`

GetLastBootTime returns the LastBootTime field if non-nil, zero value otherwise.

### GetLastBootTimeOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetLastBootTimeOk() (*time.Time, bool)`

GetLastBootTimeOk returns a tuple with the LastBootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBootTime

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetLastBootTime(v time.Time)`

SetLastBootTime sets LastBootTime field to given value.

### HasLastBootTime

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasLastBootTime() bool`

HasLastBootTime returns a boolean if a field has been set.

### GetUseWriteBackCache

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetUseWriteBackCache() bool`

GetUseWriteBackCache returns the UseWriteBackCache field if non-nil, zero value otherwise.

### GetUseWriteBackCacheOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetUseWriteBackCacheOk() (*bool, bool)`

GetUseWriteBackCacheOk returns a tuple with the UseWriteBackCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWriteBackCache

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetUseWriteBackCache(v bool)`

SetUseWriteBackCache sets UseWriteBackCache field to given value.

### HasUseWriteBackCache

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasUseWriteBackCache() bool`

HasUseWriteBackCache returns a boolean if a field has been set.

### GetWriteBackCacheDiskSize

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWriteBackCacheDiskSize() int32`

GetWriteBackCacheDiskSize returns the WriteBackCacheDiskSize field if non-nil, zero value otherwise.

### GetWriteBackCacheDiskSizeOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWriteBackCacheDiskSizeOk() (*int32, bool)`

GetWriteBackCacheDiskSizeOk returns a tuple with the WriteBackCacheDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheDiskSize

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetWriteBackCacheDiskSize(v int32)`

SetWriteBackCacheDiskSize sets WriteBackCacheDiskSize field to given value.

### HasWriteBackCacheDiskSize

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasWriteBackCacheDiskSize() bool`

HasWriteBackCacheDiskSize returns a boolean if a field has been set.

### GetWriteBackCacheDriveLetter

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWriteBackCacheDriveLetter() string`

GetWriteBackCacheDriveLetter returns the WriteBackCacheDriveLetter field if non-nil, zero value otherwise.

### GetWriteBackCacheDriveLetterOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWriteBackCacheDriveLetterOk() (*string, bool)`

GetWriteBackCacheDriveLetterOk returns a tuple with the WriteBackCacheDriveLetter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheDriveLetter

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetWriteBackCacheDriveLetter(v string)`

SetWriteBackCacheDriveLetter sets WriteBackCacheDriveLetter field to given value.

### HasWriteBackCacheDriveLetter

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasWriteBackCacheDriveLetter() bool`

HasWriteBackCacheDriveLetter returns a boolean if a field has been set.

### GetWriteBackCacheMemorySize

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWriteBackCacheMemorySize() int32`

GetWriteBackCacheMemorySize returns the WriteBackCacheMemorySize field if non-nil, zero value otherwise.

### GetWriteBackCacheMemorySizeOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWriteBackCacheMemorySizeOk() (*int32, bool)`

GetWriteBackCacheMemorySizeOk returns a tuple with the WriteBackCacheMemorySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheMemorySize

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetWriteBackCacheMemorySize(v int32)`

SetWriteBackCacheMemorySize sets WriteBackCacheMemorySize field to given value.

### HasWriteBackCacheMemorySize

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasWriteBackCacheMemorySize() bool`

HasWriteBackCacheMemorySize returns a boolean if a field has been set.

### GetIdentities

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetIdentities() string`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetIdentitiesOk() (*string, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetIdentities(v string)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.

### SetIdentitiesNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetIdentitiesNil(b bool)`

 SetIdentitiesNil sets the value for Identities to be an explicit nil

### UnsetIdentities
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetIdentities()`

UnsetIdentities ensures that no value is present for Identities, not even an explicit nil
### GetProvVMConfigurationVersion

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetProvVMConfigurationVersion() int32`

GetProvVMConfigurationVersion returns the ProvVMConfigurationVersion field if non-nil, zero value otherwise.

### GetProvVMConfigurationVersionOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetProvVMConfigurationVersionOk() (*int32, bool)`

GetProvVMConfigurationVersionOk returns a tuple with the ProvVMConfigurationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvVMConfigurationVersion

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetProvVMConfigurationVersion(v int32)`

SetProvVMConfigurationVersion sets ProvVMConfigurationVersion field to given value.

### HasProvVMConfigurationVersion

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasProvVMConfigurationVersion() bool`

HasProvVMConfigurationVersion returns a boolean if a field has been set.

### SetProvVMConfigurationVersionNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetProvVMConfigurationVersionNil(b bool)`

 SetProvVMConfigurationVersionNil sets the value for ProvVMConfigurationVersion to be an explicit nil

### UnsetProvVMConfigurationVersion
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetProvVMConfigurationVersion()`

UnsetProvVMConfigurationVersion ensures that no value is present for ProvVMConfigurationVersion, not even an explicit nil
### GetWindowsActivationType

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWindowsActivationType() string`

GetWindowsActivationType returns the WindowsActivationType field if non-nil, zero value otherwise.

### GetWindowsActivationTypeOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetWindowsActivationTypeOk() (*string, bool)`

GetWindowsActivationTypeOk returns a tuple with the WindowsActivationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsActivationType

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetWindowsActivationType(v string)`

SetWindowsActivationType sets WindowsActivationType field to given value.

### HasWindowsActivationType

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasWindowsActivationType() bool`

HasWindowsActivationType returns a boolean if a field has been set.

### SetWindowsActivationTypeNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetWindowsActivationTypeNil(b bool)`

 SetWindowsActivationTypeNil sets the value for WindowsActivationType to be an explicit nil

### UnsetWindowsActivationType
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetWindowsActivationType()`

UnsetWindowsActivationType ensures that no value is present for WindowsActivationType, not even an explicit nil
### GetDiskInfo

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetDiskInfo() string`

GetDiskInfo returns the DiskInfo field if non-nil, zero value otherwise.

### GetDiskInfoOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetDiskInfoOk() (*string, bool)`

GetDiskInfoOk returns a tuple with the DiskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskInfo

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetDiskInfo(v string)`

SetDiskInfo sets DiskInfo field to given value.

### HasDiskInfo

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasDiskInfo() bool`

HasDiskInfo returns a boolean if a field has been set.

### SetDiskInfoNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetDiskInfoNil(b bool)`

 SetDiskInfoNil sets the value for DiskInfo to be an explicit nil

### UnsetDiskInfo
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetDiskInfo()`

UnsetDiskInfo ensures that no value is present for DiskInfo, not even an explicit nil
### GetImageAndTemplateInfo

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetImageAndTemplateInfo() string`

GetImageAndTemplateInfo returns the ImageAndTemplateInfo field if non-nil, zero value otherwise.

### GetImageAndTemplateInfoOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetImageAndTemplateInfoOk() (*string, bool)`

GetImageAndTemplateInfoOk returns a tuple with the ImageAndTemplateInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAndTemplateInfo

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetImageAndTemplateInfo(v string)`

SetImageAndTemplateInfo sets ImageAndTemplateInfo field to given value.

### HasImageAndTemplateInfo

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasImageAndTemplateInfo() bool`

HasImageAndTemplateInfo returns a boolean if a field has been set.

### SetImageAndTemplateInfoNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetImageAndTemplateInfoNil(b bool)`

 SetImageAndTemplateInfoNil sets the value for ImageAndTemplateInfo to be an explicit nil

### UnsetImageAndTemplateInfo
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetImageAndTemplateInfo()`

UnsetImageAndTemplateInfo ensures that no value is present for ImageAndTemplateInfo, not even an explicit nil
### GetMCSIOInfo

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetMCSIOInfo() string`

GetMCSIOInfo returns the MCSIOInfo field if non-nil, zero value otherwise.

### GetMCSIOInfoOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetMCSIOInfoOk() (*string, bool)`

GetMCSIOInfoOk returns a tuple with the MCSIOInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMCSIOInfo

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetMCSIOInfo(v string)`

SetMCSIOInfo sets MCSIOInfo field to given value.

### HasMCSIOInfo

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasMCSIOInfo() bool`

HasMCSIOInfo returns a boolean if a field has been set.

### SetMCSIOInfoNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetMCSIOInfoNil(b bool)`

 SetMCSIOInfoNil sets the value for MCSIOInfo to be an explicit nil

### UnsetMCSIOInfo
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetMCSIOInfo()`

UnsetMCSIOInfo ensures that no value is present for MCSIOInfo, not even an explicit nil
### GetNetworkInfo

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetNetworkInfo() string`

GetNetworkInfo returns the NetworkInfo field if non-nil, zero value otherwise.

### GetNetworkInfoOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetNetworkInfoOk() (*string, bool)`

GetNetworkInfoOk returns a tuple with the NetworkInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInfo

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetNetworkInfo(v string)`

SetNetworkInfo sets NetworkInfo field to given value.

### HasNetworkInfo

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasNetworkInfo() bool`

HasNetworkInfo returns a boolean if a field has been set.

### SetNetworkInfoNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetNetworkInfoNil(b bool)`

 SetNetworkInfoNil sets the value for NetworkInfo to be an explicit nil

### UnsetNetworkInfo
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetNetworkInfo()`

UnsetNetworkInfo ensures that no value is present for NetworkInfo, not even an explicit nil
### GetSecurityInfo

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetSecurityInfo() string`

GetSecurityInfo returns the SecurityInfo field if non-nil, zero value otherwise.

### GetSecurityInfoOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetSecurityInfoOk() (*string, bool)`

GetSecurityInfoOk returns a tuple with the SecurityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityInfo

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetSecurityInfo(v string)`

SetSecurityInfo sets SecurityInfo field to given value.

### HasSecurityInfo

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasSecurityInfo() bool`

HasSecurityInfo returns a boolean if a field has been set.

### SetSecurityInfoNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetSecurityInfoNil(b bool)`

 SetSecurityInfoNil sets the value for SecurityInfo to be an explicit nil

### UnsetSecurityInfo
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetSecurityInfo()`

UnsetSecurityInfo ensures that no value is present for SecurityInfo, not even an explicit nil
### GetVMInfo

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetVMInfo() string`

GetVMInfo returns the VMInfo field if non-nil, zero value otherwise.

### GetVMInfoOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetVMInfoOk() (*string, bool)`

GetVMInfoOk returns a tuple with the VMInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMInfo

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetVMInfo(v string)`

SetVMInfo sets VMInfo field to given value.

### HasVMInfo

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasVMInfo() bool`

HasVMInfo returns a boolean if a field has been set.

### SetVMInfoNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetVMInfoNil(b bool)`

 SetVMInfoNil sets the value for VMInfo to be an explicit nil

### UnsetVMInfo
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetVMInfo()`

UnsetVMInfo ensures that no value is present for VMInfo, not even an explicit nil
### GetResourcePool

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetResourcePool() HypervisorResourcePoolRefResponseModel`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetResourcePoolOk() (*HypervisorResourcePoolRefResponseModel, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetResourcePool(v HypervisorResourcePoolRefResponseModel)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetConfigurationResultantSet

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetConfigurationResultantSet() ProvisionedVirtualMachineConfigurationResultantSetResponseModel`

GetConfigurationResultantSet returns the ConfigurationResultantSet field if non-nil, zero value otherwise.

### GetConfigurationResultantSetOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetConfigurationResultantSetOk() (*ProvisionedVirtualMachineConfigurationResultantSetResponseModel, bool)`

GetConfigurationResultantSetOk returns a tuple with the ConfigurationResultantSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationResultantSet

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetConfigurationResultantSet(v ProvisionedVirtualMachineConfigurationResultantSetResponseModel)`

SetConfigurationResultantSet sets ConfigurationResultantSet field to given value.

### HasConfigurationResultantSet

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasConfigurationResultantSet() bool`

HasConfigurationResultantSet returns a boolean if a field has been set.

### GetFailureReason

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReasonNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
### GetFailedActionCompletionTime

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetFailedActionCompletionTime() string`

GetFailedActionCompletionTime returns the FailedActionCompletionTime field if non-nil, zero value otherwise.

### GetFailedActionCompletionTimeOk

`func (o *ProvisionedVirtualMachineDetailsResponseModel) GetFailedActionCompletionTimeOk() (*string, bool)`

GetFailedActionCompletionTimeOk returns a tuple with the FailedActionCompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedActionCompletionTime

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetFailedActionCompletionTime(v string)`

SetFailedActionCompletionTime sets FailedActionCompletionTime field to given value.

### HasFailedActionCompletionTime

`func (o *ProvisionedVirtualMachineDetailsResponseModel) HasFailedActionCompletionTime() bool`

HasFailedActionCompletionTime returns a boolean if a field has been set.

### SetFailedActionCompletionTimeNil

`func (o *ProvisionedVirtualMachineDetailsResponseModel) SetFailedActionCompletionTimeNil(b bool)`

 SetFailedActionCompletionTimeNil sets the value for FailedActionCompletionTime to be an explicit nil

### UnsetFailedActionCompletionTime
`func (o *ProvisionedVirtualMachineDetailsResponseModel) UnsetFailedActionCompletionTime()`

UnsetFailedActionCompletionTime ensures that no value is present for FailedActionCompletionTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


