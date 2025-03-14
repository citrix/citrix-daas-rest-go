# ProvisioningSchemeSearchResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationType** | Pointer to [**WindowsActivationType**](WindowsActivationType.md) |  | [optional] 
**CpuCount** | Pointer to **int32** | Cpu Count. | [optional] 
**DiskSize** | Pointer to **int32** | Disk Size. | [optional] 
**Id** | Pointer to **NullableString** | Provisioning Scheme Id. | [optional] 
**IdentityPoolUid** | Pointer to **string** | Identity Pool Uid. | [optional] 
**IdentityType** | Pointer to **NullableString** | Identity Type:ActiveDirectory, AzureAD, HybridAzureAD, Workgroup. | [optional] 
**MachineCount** | Pointer to **int32** | Machine Count. | [optional] 
**MemoryMB** | Pointer to **int32** | Memory(MB). | [optional] 
**Name** | Pointer to **NullableString** | Provisioning Scheme Name. | [optional] 
**Persistency** | Pointer to **bool** | Persistency. | [optional] 
**ProvisioningSchemeType** | Pointer to **NullableString** | Provisioning Scheme Type. | [optional] 
**State** | Pointer to **NullableString** | Provisioning Scheme State. | [optional] 
**UseWriteBackCache** | Pointer to **bool** | Whether Use Write Back Cache. | [optional] 
**WriteBackCacheDiskSize** | Pointer to **int32** | Write Back Cache Disk Size. | [optional] 
**WriteBackCacheMemorySize** | Pointer to **int32** | Use Write Back Cache Memory Size. | [optional] 
**Version** | Pointer to **int32** | Provisioning Scheme Version. | [optional] 

## Methods

### NewProvisioningSchemeSearchResponseModel

`func NewProvisioningSchemeSearchResponseModel() *ProvisioningSchemeSearchResponseModel`

NewProvisioningSchemeSearchResponseModel instantiates a new ProvisioningSchemeSearchResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningSchemeSearchResponseModelWithDefaults

`func NewProvisioningSchemeSearchResponseModelWithDefaults() *ProvisioningSchemeSearchResponseModel`

NewProvisioningSchemeSearchResponseModelWithDefaults instantiates a new ProvisioningSchemeSearchResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationType

`func (o *ProvisioningSchemeSearchResponseModel) GetActivationType() WindowsActivationType`

GetActivationType returns the ActivationType field if non-nil, zero value otherwise.

### GetActivationTypeOk

`func (o *ProvisioningSchemeSearchResponseModel) GetActivationTypeOk() (*WindowsActivationType, bool)`

GetActivationTypeOk returns a tuple with the ActivationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationType

`func (o *ProvisioningSchemeSearchResponseModel) SetActivationType(v WindowsActivationType)`

SetActivationType sets ActivationType field to given value.

### HasActivationType

`func (o *ProvisioningSchemeSearchResponseModel) HasActivationType() bool`

HasActivationType returns a boolean if a field has been set.

### GetCpuCount

`func (o *ProvisioningSchemeSearchResponseModel) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *ProvisioningSchemeSearchResponseModel) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *ProvisioningSchemeSearchResponseModel) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *ProvisioningSchemeSearchResponseModel) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetDiskSize

`func (o *ProvisioningSchemeSearchResponseModel) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *ProvisioningSchemeSearchResponseModel) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *ProvisioningSchemeSearchResponseModel) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *ProvisioningSchemeSearchResponseModel) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetId

`func (o *ProvisioningSchemeSearchResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvisioningSchemeSearchResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvisioningSchemeSearchResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProvisioningSchemeSearchResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProvisioningSchemeSearchResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProvisioningSchemeSearchResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIdentityPoolUid

`func (o *ProvisioningSchemeSearchResponseModel) GetIdentityPoolUid() string`

GetIdentityPoolUid returns the IdentityPoolUid field if non-nil, zero value otherwise.

### GetIdentityPoolUidOk

`func (o *ProvisioningSchemeSearchResponseModel) GetIdentityPoolUidOk() (*string, bool)`

GetIdentityPoolUidOk returns a tuple with the IdentityPoolUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityPoolUid

`func (o *ProvisioningSchemeSearchResponseModel) SetIdentityPoolUid(v string)`

SetIdentityPoolUid sets IdentityPoolUid field to given value.

### HasIdentityPoolUid

`func (o *ProvisioningSchemeSearchResponseModel) HasIdentityPoolUid() bool`

HasIdentityPoolUid returns a boolean if a field has been set.

### GetIdentityType

`func (o *ProvisioningSchemeSearchResponseModel) GetIdentityType() string`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *ProvisioningSchemeSearchResponseModel) GetIdentityTypeOk() (*string, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *ProvisioningSchemeSearchResponseModel) SetIdentityType(v string)`

SetIdentityType sets IdentityType field to given value.

### HasIdentityType

`func (o *ProvisioningSchemeSearchResponseModel) HasIdentityType() bool`

HasIdentityType returns a boolean if a field has been set.

### SetIdentityTypeNil

`func (o *ProvisioningSchemeSearchResponseModel) SetIdentityTypeNil(b bool)`

 SetIdentityTypeNil sets the value for IdentityType to be an explicit nil

### UnsetIdentityType
`func (o *ProvisioningSchemeSearchResponseModel) UnsetIdentityType()`

UnsetIdentityType ensures that no value is present for IdentityType, not even an explicit nil
### GetMachineCount

`func (o *ProvisioningSchemeSearchResponseModel) GetMachineCount() int32`

GetMachineCount returns the MachineCount field if non-nil, zero value otherwise.

### GetMachineCountOk

`func (o *ProvisioningSchemeSearchResponseModel) GetMachineCountOk() (*int32, bool)`

GetMachineCountOk returns a tuple with the MachineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineCount

`func (o *ProvisioningSchemeSearchResponseModel) SetMachineCount(v int32)`

SetMachineCount sets MachineCount field to given value.

### HasMachineCount

`func (o *ProvisioningSchemeSearchResponseModel) HasMachineCount() bool`

HasMachineCount returns a boolean if a field has been set.

### GetMemoryMB

`func (o *ProvisioningSchemeSearchResponseModel) GetMemoryMB() int32`

GetMemoryMB returns the MemoryMB field if non-nil, zero value otherwise.

### GetMemoryMBOk

`func (o *ProvisioningSchemeSearchResponseModel) GetMemoryMBOk() (*int32, bool)`

GetMemoryMBOk returns a tuple with the MemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMB

`func (o *ProvisioningSchemeSearchResponseModel) SetMemoryMB(v int32)`

SetMemoryMB sets MemoryMB field to given value.

### HasMemoryMB

`func (o *ProvisioningSchemeSearchResponseModel) HasMemoryMB() bool`

HasMemoryMB returns a boolean if a field has been set.

### GetName

`func (o *ProvisioningSchemeSearchResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisioningSchemeSearchResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisioningSchemeSearchResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProvisioningSchemeSearchResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProvisioningSchemeSearchResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProvisioningSchemeSearchResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPersistency

`func (o *ProvisioningSchemeSearchResponseModel) GetPersistency() bool`

GetPersistency returns the Persistency field if non-nil, zero value otherwise.

### GetPersistencyOk

`func (o *ProvisioningSchemeSearchResponseModel) GetPersistencyOk() (*bool, bool)`

GetPersistencyOk returns a tuple with the Persistency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistency

`func (o *ProvisioningSchemeSearchResponseModel) SetPersistency(v bool)`

SetPersistency sets Persistency field to given value.

### HasPersistency

`func (o *ProvisioningSchemeSearchResponseModel) HasPersistency() bool`

HasPersistency returns a boolean if a field has been set.

### GetProvisioningSchemeType

`func (o *ProvisioningSchemeSearchResponseModel) GetProvisioningSchemeType() string`

GetProvisioningSchemeType returns the ProvisioningSchemeType field if non-nil, zero value otherwise.

### GetProvisioningSchemeTypeOk

`func (o *ProvisioningSchemeSearchResponseModel) GetProvisioningSchemeTypeOk() (*string, bool)`

GetProvisioningSchemeTypeOk returns a tuple with the ProvisioningSchemeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSchemeType

`func (o *ProvisioningSchemeSearchResponseModel) SetProvisioningSchemeType(v string)`

SetProvisioningSchemeType sets ProvisioningSchemeType field to given value.

### HasProvisioningSchemeType

`func (o *ProvisioningSchemeSearchResponseModel) HasProvisioningSchemeType() bool`

HasProvisioningSchemeType returns a boolean if a field has been set.

### SetProvisioningSchemeTypeNil

`func (o *ProvisioningSchemeSearchResponseModel) SetProvisioningSchemeTypeNil(b bool)`

 SetProvisioningSchemeTypeNil sets the value for ProvisioningSchemeType to be an explicit nil

### UnsetProvisioningSchemeType
`func (o *ProvisioningSchemeSearchResponseModel) UnsetProvisioningSchemeType()`

UnsetProvisioningSchemeType ensures that no value is present for ProvisioningSchemeType, not even an explicit nil
### GetState

`func (o *ProvisioningSchemeSearchResponseModel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ProvisioningSchemeSearchResponseModel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ProvisioningSchemeSearchResponseModel) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ProvisioningSchemeSearchResponseModel) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ProvisioningSchemeSearchResponseModel) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ProvisioningSchemeSearchResponseModel) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetUseWriteBackCache

`func (o *ProvisioningSchemeSearchResponseModel) GetUseWriteBackCache() bool`

GetUseWriteBackCache returns the UseWriteBackCache field if non-nil, zero value otherwise.

### GetUseWriteBackCacheOk

`func (o *ProvisioningSchemeSearchResponseModel) GetUseWriteBackCacheOk() (*bool, bool)`

GetUseWriteBackCacheOk returns a tuple with the UseWriteBackCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWriteBackCache

`func (o *ProvisioningSchemeSearchResponseModel) SetUseWriteBackCache(v bool)`

SetUseWriteBackCache sets UseWriteBackCache field to given value.

### HasUseWriteBackCache

`func (o *ProvisioningSchemeSearchResponseModel) HasUseWriteBackCache() bool`

HasUseWriteBackCache returns a boolean if a field has been set.

### GetWriteBackCacheDiskSize

`func (o *ProvisioningSchemeSearchResponseModel) GetWriteBackCacheDiskSize() int32`

GetWriteBackCacheDiskSize returns the WriteBackCacheDiskSize field if non-nil, zero value otherwise.

### GetWriteBackCacheDiskSizeOk

`func (o *ProvisioningSchemeSearchResponseModel) GetWriteBackCacheDiskSizeOk() (*int32, bool)`

GetWriteBackCacheDiskSizeOk returns a tuple with the WriteBackCacheDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheDiskSize

`func (o *ProvisioningSchemeSearchResponseModel) SetWriteBackCacheDiskSize(v int32)`

SetWriteBackCacheDiskSize sets WriteBackCacheDiskSize field to given value.

### HasWriteBackCacheDiskSize

`func (o *ProvisioningSchemeSearchResponseModel) HasWriteBackCacheDiskSize() bool`

HasWriteBackCacheDiskSize returns a boolean if a field has been set.

### GetWriteBackCacheMemorySize

`func (o *ProvisioningSchemeSearchResponseModel) GetWriteBackCacheMemorySize() int32`

GetWriteBackCacheMemorySize returns the WriteBackCacheMemorySize field if non-nil, zero value otherwise.

### GetWriteBackCacheMemorySizeOk

`func (o *ProvisioningSchemeSearchResponseModel) GetWriteBackCacheMemorySizeOk() (*int32, bool)`

GetWriteBackCacheMemorySizeOk returns a tuple with the WriteBackCacheMemorySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheMemorySize

`func (o *ProvisioningSchemeSearchResponseModel) SetWriteBackCacheMemorySize(v int32)`

SetWriteBackCacheMemorySize sets WriteBackCacheMemorySize field to given value.

### HasWriteBackCacheMemorySize

`func (o *ProvisioningSchemeSearchResponseModel) HasWriteBackCacheMemorySize() bool`

HasWriteBackCacheMemorySize returns a boolean if a field has been set.

### GetVersion

`func (o *ProvisioningSchemeSearchResponseModel) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProvisioningSchemeSearchResponseModel) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProvisioningSchemeSearchResponseModel) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ProvisioningSchemeSearchResponseModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


