# ProvisionedVirtualMachineSearchResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveOperation** | Pointer to **NullableString** | Active Operation. | [optional] 
**VMSid** | Pointer to **NullableString** | Provisioned Virtual Machine Sid. | [optional] 
**CpuCount** | Pointer to **int32** | Cpu Count. | [optional] 
**IdentityType** | Pointer to **NullableString** | Identity Type:ActiveDirectory, AzureAD, HybridAzureAD, Workgroup. | [optional] 
**Identities** | Pointer to **NullableString** | ActiveDirectory, HybridAzureAD - Domain. AzureAD - TenantId. Workgroup - null. | [optional] 
**ImageOutOfDate** | Pointer to **bool** | Booted image isn&#39;t same as assigned image. | [optional] 
**LastBootTime** | Pointer to **NullableString** | Last boot time. | [optional] 
**MemoryMB** | Pointer to **int32** | Memory(MB). | [optional] 
**Persistency** | Pointer to **bool** | Persistency. | [optional] 
**ProvisioningSchemeName** | Pointer to **NullableString** | Provision Scheme Name. | [optional] 
**ProvisioningSchemeVersion** | Pointer to **int32** | Provision Scheme Version. | [optional] 
**ProvVMConfigurationUpdateVersion** | Pointer to **NullableInt32** | Provisioned Virtual Machine Update Configuration Version. | [optional] 
**VMName** | Pointer to **NullableString** | Provisioned virtual machine name on hypervisor. | [optional] 
**ActivationType** | Pointer to [**WindowsActivationType**](WindowsActivationType.md) |  | [optional] 
**UseWriteBackCache** | Pointer to **bool** | Whether use write back cache.  | [optional] 
**WriteBackCacheDiskSize** | Pointer to **int32** | Write back cache disk size. | [optional] 
**WriteBackCacheMemorySize** | Pointer to **int32** | Write back cache memory size. | [optional] 

## Methods

### NewProvisionedVirtualMachineSearchResponseModel

`func NewProvisionedVirtualMachineSearchResponseModel() *ProvisionedVirtualMachineSearchResponseModel`

NewProvisionedVirtualMachineSearchResponseModel instantiates a new ProvisionedVirtualMachineSearchResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionedVirtualMachineSearchResponseModelWithDefaults

`func NewProvisionedVirtualMachineSearchResponseModelWithDefaults() *ProvisionedVirtualMachineSearchResponseModel`

NewProvisionedVirtualMachineSearchResponseModelWithDefaults instantiates a new ProvisionedVirtualMachineSearchResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveOperation

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetActiveOperation() string`

GetActiveOperation returns the ActiveOperation field if non-nil, zero value otherwise.

### GetActiveOperationOk

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetActiveOperationOk() (*string, bool)`

GetActiveOperationOk returns a tuple with the ActiveOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveOperation

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetActiveOperation(v string)`

SetActiveOperation sets ActiveOperation field to given value.

### HasActiveOperation

`func (o *ProvisionedVirtualMachineSearchResponseModel) HasActiveOperation() bool`

HasActiveOperation returns a boolean if a field has been set.

### SetActiveOperationNil

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetActiveOperationNil(b bool)`

 SetActiveOperationNil sets the value for ActiveOperation to be an explicit nil

### UnsetActiveOperation
`func (o *ProvisionedVirtualMachineSearchResponseModel) UnsetActiveOperation()`

UnsetActiveOperation ensures that no value is present for ActiveOperation, not even an explicit nil
### GetVMSid

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetVMSid() string`

GetVMSid returns the VMSid field if non-nil, zero value otherwise.

### GetVMSidOk

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetVMSidOk() (*string, bool)`

GetVMSidOk returns a tuple with the VMSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMSid

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetVMSid(v string)`

SetVMSid sets VMSid field to given value.

### HasVMSid

`func (o *ProvisionedVirtualMachineSearchResponseModel) HasVMSid() bool`

HasVMSid returns a boolean if a field has been set.

### SetVMSidNil

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetVMSidNil(b bool)`

 SetVMSidNil sets the value for VMSid to be an explicit nil

### UnsetVMSid
`func (o *ProvisionedVirtualMachineSearchResponseModel) UnsetVMSid()`

UnsetVMSid ensures that no value is present for VMSid, not even an explicit nil
### GetCpuCount

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *ProvisionedVirtualMachineSearchResponseModel) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetIdentityType

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetIdentityType() string`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetIdentityTypeOk() (*string, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetIdentityType(v string)`

SetIdentityType sets IdentityType field to given value.

### HasIdentityType

`func (o *ProvisionedVirtualMachineSearchResponseModel) HasIdentityType() bool`

HasIdentityType returns a boolean if a field has been set.

### SetIdentityTypeNil

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetIdentityTypeNil(b bool)`

 SetIdentityTypeNil sets the value for IdentityType to be an explicit nil

### UnsetIdentityType
`func (o *ProvisionedVirtualMachineSearchResponseModel) UnsetIdentityType()`

UnsetIdentityType ensures that no value is present for IdentityType, not even an explicit nil
### GetIdentities

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetIdentities() string`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetIdentitiesOk() (*string, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetIdentities(v string)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *ProvisionedVirtualMachineSearchResponseModel) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.

### SetIdentitiesNil

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetIdentitiesNil(b bool)`

 SetIdentitiesNil sets the value for Identities to be an explicit nil

### UnsetIdentities
`func (o *ProvisionedVirtualMachineSearchResponseModel) UnsetIdentities()`

UnsetIdentities ensures that no value is present for Identities, not even an explicit nil
### GetImageOutOfDate

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetImageOutOfDate() bool`

GetImageOutOfDate returns the ImageOutOfDate field if non-nil, zero value otherwise.

### GetImageOutOfDateOk

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetImageOutOfDateOk() (*bool, bool)`

GetImageOutOfDateOk returns a tuple with the ImageOutOfDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOutOfDate

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetImageOutOfDate(v bool)`

SetImageOutOfDate sets ImageOutOfDate field to given value.

### HasImageOutOfDate

`func (o *ProvisionedVirtualMachineSearchResponseModel) HasImageOutOfDate() bool`

HasImageOutOfDate returns a boolean if a field has been set.

### GetLastBootTime

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetLastBootTime() string`

GetLastBootTime returns the LastBootTime field if non-nil, zero value otherwise.

### GetLastBootTimeOk

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetLastBootTimeOk() (*string, bool)`

GetLastBootTimeOk returns a tuple with the LastBootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBootTime

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetLastBootTime(v string)`

SetLastBootTime sets LastBootTime field to given value.

### HasLastBootTime

`func (o *ProvisionedVirtualMachineSearchResponseModel) HasLastBootTime() bool`

HasLastBootTime returns a boolean if a field has been set.

### SetLastBootTimeNil

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetLastBootTimeNil(b bool)`

 SetLastBootTimeNil sets the value for LastBootTime to be an explicit nil

### UnsetLastBootTime
`func (o *ProvisionedVirtualMachineSearchResponseModel) UnsetLastBootTime()`

UnsetLastBootTime ensures that no value is present for LastBootTime, not even an explicit nil
### GetMemoryMB

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetMemoryMB() int32`

GetMemoryMB returns the MemoryMB field if non-nil, zero value otherwise.

### GetMemoryMBOk

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetMemoryMBOk() (*int32, bool)`

GetMemoryMBOk returns a tuple with the MemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMB

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetMemoryMB(v int32)`

SetMemoryMB sets MemoryMB field to given value.

### HasMemoryMB

`func (o *ProvisionedVirtualMachineSearchResponseModel) HasMemoryMB() bool`

HasMemoryMB returns a boolean if a field has been set.

### GetPersistency

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetPersistency() bool`

GetPersistency returns the Persistency field if non-nil, zero value otherwise.

### GetPersistencyOk

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetPersistencyOk() (*bool, bool)`

GetPersistencyOk returns a tuple with the Persistency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistency

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetPersistency(v bool)`

SetPersistency sets Persistency field to given value.

### HasPersistency

`func (o *ProvisionedVirtualMachineSearchResponseModel) HasPersistency() bool`

HasPersistency returns a boolean if a field has been set.

### GetProvisioningSchemeName

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetProvisioningSchemeName() string`

GetProvisioningSchemeName returns the ProvisioningSchemeName field if non-nil, zero value otherwise.

### GetProvisioningSchemeNameOk

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetProvisioningSchemeNameOk() (*string, bool)`

GetProvisioningSchemeNameOk returns a tuple with the ProvisioningSchemeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeName

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetProvisioningSchemeName(v string)`

SetProvisioningSchemeName sets ProvisioningSchemeName field to given value.

### HasProvisioningSchemeName

`func (o *ProvisionedVirtualMachineSearchResponseModel) HasProvisioningSchemeName() bool`

HasProvisioningSchemeName returns a boolean if a field has been set.

### SetProvisioningSchemeNameNil

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetProvisioningSchemeNameNil(b bool)`

 SetProvisioningSchemeNameNil sets the value for ProvisioningSchemeName to be an explicit nil

### UnsetProvisioningSchemeName
`func (o *ProvisionedVirtualMachineSearchResponseModel) UnsetProvisioningSchemeName()`

UnsetProvisioningSchemeName ensures that no value is present for ProvisioningSchemeName, not even an explicit nil
### GetProvisioningSchemeVersion

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetProvisioningSchemeVersion() int32`

GetProvisioningSchemeVersion returns the ProvisioningSchemeVersion field if non-nil, zero value otherwise.

### GetProvisioningSchemeVersionOk

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetProvisioningSchemeVersionOk() (*int32, bool)`

GetProvisioningSchemeVersionOk returns a tuple with the ProvisioningSchemeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeVersion

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetProvisioningSchemeVersion(v int32)`

SetProvisioningSchemeVersion sets ProvisioningSchemeVersion field to given value.

### HasProvisioningSchemeVersion

`func (o *ProvisionedVirtualMachineSearchResponseModel) HasProvisioningSchemeVersion() bool`

HasProvisioningSchemeVersion returns a boolean if a field has been set.

### GetProvVMConfigurationUpdateVersion

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetProvVMConfigurationUpdateVersion() int32`

GetProvVMConfigurationUpdateVersion returns the ProvVMConfigurationUpdateVersion field if non-nil, zero value otherwise.

### GetProvVMConfigurationUpdateVersionOk

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetProvVMConfigurationUpdateVersionOk() (*int32, bool)`

GetProvVMConfigurationUpdateVersionOk returns a tuple with the ProvVMConfigurationUpdateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvVMConfigurationUpdateVersion

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetProvVMConfigurationUpdateVersion(v int32)`

SetProvVMConfigurationUpdateVersion sets ProvVMConfigurationUpdateVersion field to given value.

### HasProvVMConfigurationUpdateVersion

`func (o *ProvisionedVirtualMachineSearchResponseModel) HasProvVMConfigurationUpdateVersion() bool`

HasProvVMConfigurationUpdateVersion returns a boolean if a field has been set.

### SetProvVMConfigurationUpdateVersionNil

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetProvVMConfigurationUpdateVersionNil(b bool)`

 SetProvVMConfigurationUpdateVersionNil sets the value for ProvVMConfigurationUpdateVersion to be an explicit nil

### UnsetProvVMConfigurationUpdateVersion
`func (o *ProvisionedVirtualMachineSearchResponseModel) UnsetProvVMConfigurationUpdateVersion()`

UnsetProvVMConfigurationUpdateVersion ensures that no value is present for ProvVMConfigurationUpdateVersion, not even an explicit nil
### GetVMName

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetVMName() string`

GetVMName returns the VMName field if non-nil, zero value otherwise.

### GetVMNameOk

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetVMNameOk() (*string, bool)`

GetVMNameOk returns a tuple with the VMName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMName

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetVMName(v string)`

SetVMName sets VMName field to given value.

### HasVMName

`func (o *ProvisionedVirtualMachineSearchResponseModel) HasVMName() bool`

HasVMName returns a boolean if a field has been set.

### SetVMNameNil

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetVMNameNil(b bool)`

 SetVMNameNil sets the value for VMName to be an explicit nil

### UnsetVMName
`func (o *ProvisionedVirtualMachineSearchResponseModel) UnsetVMName()`

UnsetVMName ensures that no value is present for VMName, not even an explicit nil
### GetActivationType

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetActivationType() WindowsActivationType`

GetActivationType returns the ActivationType field if non-nil, zero value otherwise.

### GetActivationTypeOk

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetActivationTypeOk() (*WindowsActivationType, bool)`

GetActivationTypeOk returns a tuple with the ActivationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationType

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetActivationType(v WindowsActivationType)`

SetActivationType sets ActivationType field to given value.

### HasActivationType

`func (o *ProvisionedVirtualMachineSearchResponseModel) HasActivationType() bool`

HasActivationType returns a boolean if a field has been set.

### GetUseWriteBackCache

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetUseWriteBackCache() bool`

GetUseWriteBackCache returns the UseWriteBackCache field if non-nil, zero value otherwise.

### GetUseWriteBackCacheOk

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetUseWriteBackCacheOk() (*bool, bool)`

GetUseWriteBackCacheOk returns a tuple with the UseWriteBackCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWriteBackCache

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetUseWriteBackCache(v bool)`

SetUseWriteBackCache sets UseWriteBackCache field to given value.

### HasUseWriteBackCache

`func (o *ProvisionedVirtualMachineSearchResponseModel) HasUseWriteBackCache() bool`

HasUseWriteBackCache returns a boolean if a field has been set.

### GetWriteBackCacheDiskSize

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetWriteBackCacheDiskSize() int32`

GetWriteBackCacheDiskSize returns the WriteBackCacheDiskSize field if non-nil, zero value otherwise.

### GetWriteBackCacheDiskSizeOk

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetWriteBackCacheDiskSizeOk() (*int32, bool)`

GetWriteBackCacheDiskSizeOk returns a tuple with the WriteBackCacheDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheDiskSize

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetWriteBackCacheDiskSize(v int32)`

SetWriteBackCacheDiskSize sets WriteBackCacheDiskSize field to given value.

### HasWriteBackCacheDiskSize

`func (o *ProvisionedVirtualMachineSearchResponseModel) HasWriteBackCacheDiskSize() bool`

HasWriteBackCacheDiskSize returns a boolean if a field has been set.

### GetWriteBackCacheMemorySize

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetWriteBackCacheMemorySize() int32`

GetWriteBackCacheMemorySize returns the WriteBackCacheMemorySize field if non-nil, zero value otherwise.

### GetWriteBackCacheMemorySizeOk

`func (o *ProvisionedVirtualMachineSearchResponseModel) GetWriteBackCacheMemorySizeOk() (*int32, bool)`

GetWriteBackCacheMemorySizeOk returns a tuple with the WriteBackCacheMemorySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheMemorySize

`func (o *ProvisionedVirtualMachineSearchResponseModel) SetWriteBackCacheMemorySize(v int32)`

SetWriteBackCacheMemorySize sets WriteBackCacheMemorySize field to given value.

### HasWriteBackCacheMemorySize

`func (o *ProvisionedVirtualMachineSearchResponseModel) HasWriteBackCacheMemorySize() bool`

HasWriteBackCacheMemorySize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


