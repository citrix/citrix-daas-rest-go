# AzureVMSize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Instance type ID used when provisioning a VM | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Sku** | Pointer to **NullableString** | Name of the VM SKU to display to the user | [optional] 
**IsBasicVM** | Pointer to **bool** |  | [optional] [readonly] 
**SupportsPremiumStorage** | Pointer to **bool** | True if a virtual machine size supports premium storage. False otherwise. | [optional] [readonly] 
**SupportsHibernation** | Pointer to **bool** | True if a virtual machine size supports hibernation. False otherwise. | [optional] [readonly] 
**NumberOfCores** | Pointer to **NullableInt32** | the Number of cores supported by a VM size. | [optional] [readonly] 
**OsDiskSizeInMB** | Pointer to **NullableInt32** | the OS disk size allowed by a VM size. | [optional] [readonly] 
**ResourceDiskSizeInMB** | Pointer to **NullableInt32** | the Resource disk size allowed by a VM size. | [optional] [readonly] 
**MemoryInMB** | Pointer to **NullableFloat64** | the Memory size supported by a VM size. | [optional] [readonly] 
**MaxDataDiskCount** | Pointer to **NullableInt32** | the Maximum number of data disks allowed by a VM size. | [optional] [readonly] 
**HyperVGen1** | Pointer to **bool** | Is compatible with HyperVGeneration V1 | [optional] [readonly] 
**HyperVGen2** | Pointer to **bool** | Is compatible with HyperVGeneration V2 | [optional] [readonly] 
**NumberofGPUs** | Pointer to **NullableInt32** | Specifies the number of GPUs in the VM SKU | [optional] [readonly] 
**GpuMemoryInMB** | Pointer to **NullableFloat64** | The GPU memory in MB | [optional] [readonly] 
**SupportedSessionsPerVm** | Pointer to **[]int32** | A set of supported sessions per vm for this sku. Valid only if sku is a cma phoenix sku | [optional] 
**EstimatedCredits** | Pointer to **NullableFloat64** | The estimated credits per month per user for this sku. Valid only if sku is a cma phoenix sku | [optional] 
**UsageType** | Pointer to **NullableString** | The Usage Type that this VM size&#39;s Quota is associated with | [optional] [readonly] 

## Methods

### NewAzureVMSize

`func NewAzureVMSize() *AzureVMSize`

NewAzureVMSize instantiates a new AzureVMSize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureVMSizeWithDefaults

`func NewAzureVMSizeWithDefaults() *AzureVMSize`

NewAzureVMSizeWithDefaults instantiates a new AzureVMSize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AzureVMSize) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureVMSize) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureVMSize) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureVMSize) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AzureVMSize) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AzureVMSize) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDisplayName

`func (o *AzureVMSize) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AzureVMSize) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AzureVMSize) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AzureVMSize) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AzureVMSize) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AzureVMSize) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetSku

`func (o *AzureVMSize) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *AzureVMSize) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *AzureVMSize) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *AzureVMSize) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *AzureVMSize) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *AzureVMSize) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetIsBasicVM

`func (o *AzureVMSize) GetIsBasicVM() bool`

GetIsBasicVM returns the IsBasicVM field if non-nil, zero value otherwise.

### GetIsBasicVMOk

`func (o *AzureVMSize) GetIsBasicVMOk() (*bool, bool)`

GetIsBasicVMOk returns a tuple with the IsBasicVM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBasicVM

`func (o *AzureVMSize) SetIsBasicVM(v bool)`

SetIsBasicVM sets IsBasicVM field to given value.

### HasIsBasicVM

`func (o *AzureVMSize) HasIsBasicVM() bool`

HasIsBasicVM returns a boolean if a field has been set.

### GetSupportsPremiumStorage

`func (o *AzureVMSize) GetSupportsPremiumStorage() bool`

GetSupportsPremiumStorage returns the SupportsPremiumStorage field if non-nil, zero value otherwise.

### GetSupportsPremiumStorageOk

`func (o *AzureVMSize) GetSupportsPremiumStorageOk() (*bool, bool)`

GetSupportsPremiumStorageOk returns a tuple with the SupportsPremiumStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsPremiumStorage

`func (o *AzureVMSize) SetSupportsPremiumStorage(v bool)`

SetSupportsPremiumStorage sets SupportsPremiumStorage field to given value.

### HasSupportsPremiumStorage

`func (o *AzureVMSize) HasSupportsPremiumStorage() bool`

HasSupportsPremiumStorage returns a boolean if a field has been set.

### GetSupportsHibernation

`func (o *AzureVMSize) GetSupportsHibernation() bool`

GetSupportsHibernation returns the SupportsHibernation field if non-nil, zero value otherwise.

### GetSupportsHibernationOk

`func (o *AzureVMSize) GetSupportsHibernationOk() (*bool, bool)`

GetSupportsHibernationOk returns a tuple with the SupportsHibernation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsHibernation

`func (o *AzureVMSize) SetSupportsHibernation(v bool)`

SetSupportsHibernation sets SupportsHibernation field to given value.

### HasSupportsHibernation

`func (o *AzureVMSize) HasSupportsHibernation() bool`

HasSupportsHibernation returns a boolean if a field has been set.

### GetNumberOfCores

`func (o *AzureVMSize) GetNumberOfCores() int32`

GetNumberOfCores returns the NumberOfCores field if non-nil, zero value otherwise.

### GetNumberOfCoresOk

`func (o *AzureVMSize) GetNumberOfCoresOk() (*int32, bool)`

GetNumberOfCoresOk returns a tuple with the NumberOfCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCores

`func (o *AzureVMSize) SetNumberOfCores(v int32)`

SetNumberOfCores sets NumberOfCores field to given value.

### HasNumberOfCores

`func (o *AzureVMSize) HasNumberOfCores() bool`

HasNumberOfCores returns a boolean if a field has been set.

### SetNumberOfCoresNil

`func (o *AzureVMSize) SetNumberOfCoresNil(b bool)`

 SetNumberOfCoresNil sets the value for NumberOfCores to be an explicit nil

### UnsetNumberOfCores
`func (o *AzureVMSize) UnsetNumberOfCores()`

UnsetNumberOfCores ensures that no value is present for NumberOfCores, not even an explicit nil
### GetOsDiskSizeInMB

`func (o *AzureVMSize) GetOsDiskSizeInMB() int32`

GetOsDiskSizeInMB returns the OsDiskSizeInMB field if non-nil, zero value otherwise.

### GetOsDiskSizeInMBOk

`func (o *AzureVMSize) GetOsDiskSizeInMBOk() (*int32, bool)`

GetOsDiskSizeInMBOk returns a tuple with the OsDiskSizeInMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDiskSizeInMB

`func (o *AzureVMSize) SetOsDiskSizeInMB(v int32)`

SetOsDiskSizeInMB sets OsDiskSizeInMB field to given value.

### HasOsDiskSizeInMB

`func (o *AzureVMSize) HasOsDiskSizeInMB() bool`

HasOsDiskSizeInMB returns a boolean if a field has been set.

### SetOsDiskSizeInMBNil

`func (o *AzureVMSize) SetOsDiskSizeInMBNil(b bool)`

 SetOsDiskSizeInMBNil sets the value for OsDiskSizeInMB to be an explicit nil

### UnsetOsDiskSizeInMB
`func (o *AzureVMSize) UnsetOsDiskSizeInMB()`

UnsetOsDiskSizeInMB ensures that no value is present for OsDiskSizeInMB, not even an explicit nil
### GetResourceDiskSizeInMB

`func (o *AzureVMSize) GetResourceDiskSizeInMB() int32`

GetResourceDiskSizeInMB returns the ResourceDiskSizeInMB field if non-nil, zero value otherwise.

### GetResourceDiskSizeInMBOk

`func (o *AzureVMSize) GetResourceDiskSizeInMBOk() (*int32, bool)`

GetResourceDiskSizeInMBOk returns a tuple with the ResourceDiskSizeInMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDiskSizeInMB

`func (o *AzureVMSize) SetResourceDiskSizeInMB(v int32)`

SetResourceDiskSizeInMB sets ResourceDiskSizeInMB field to given value.

### HasResourceDiskSizeInMB

`func (o *AzureVMSize) HasResourceDiskSizeInMB() bool`

HasResourceDiskSizeInMB returns a boolean if a field has been set.

### SetResourceDiskSizeInMBNil

`func (o *AzureVMSize) SetResourceDiskSizeInMBNil(b bool)`

 SetResourceDiskSizeInMBNil sets the value for ResourceDiskSizeInMB to be an explicit nil

### UnsetResourceDiskSizeInMB
`func (o *AzureVMSize) UnsetResourceDiskSizeInMB()`

UnsetResourceDiskSizeInMB ensures that no value is present for ResourceDiskSizeInMB, not even an explicit nil
### GetMemoryInMB

`func (o *AzureVMSize) GetMemoryInMB() float64`

GetMemoryInMB returns the MemoryInMB field if non-nil, zero value otherwise.

### GetMemoryInMBOk

`func (o *AzureVMSize) GetMemoryInMBOk() (*float64, bool)`

GetMemoryInMBOk returns a tuple with the MemoryInMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryInMB

`func (o *AzureVMSize) SetMemoryInMB(v float64)`

SetMemoryInMB sets MemoryInMB field to given value.

### HasMemoryInMB

`func (o *AzureVMSize) HasMemoryInMB() bool`

HasMemoryInMB returns a boolean if a field has been set.

### SetMemoryInMBNil

`func (o *AzureVMSize) SetMemoryInMBNil(b bool)`

 SetMemoryInMBNil sets the value for MemoryInMB to be an explicit nil

### UnsetMemoryInMB
`func (o *AzureVMSize) UnsetMemoryInMB()`

UnsetMemoryInMB ensures that no value is present for MemoryInMB, not even an explicit nil
### GetMaxDataDiskCount

`func (o *AzureVMSize) GetMaxDataDiskCount() int32`

GetMaxDataDiskCount returns the MaxDataDiskCount field if non-nil, zero value otherwise.

### GetMaxDataDiskCountOk

`func (o *AzureVMSize) GetMaxDataDiskCountOk() (*int32, bool)`

GetMaxDataDiskCountOk returns a tuple with the MaxDataDiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataDiskCount

`func (o *AzureVMSize) SetMaxDataDiskCount(v int32)`

SetMaxDataDiskCount sets MaxDataDiskCount field to given value.

### HasMaxDataDiskCount

`func (o *AzureVMSize) HasMaxDataDiskCount() bool`

HasMaxDataDiskCount returns a boolean if a field has been set.

### SetMaxDataDiskCountNil

`func (o *AzureVMSize) SetMaxDataDiskCountNil(b bool)`

 SetMaxDataDiskCountNil sets the value for MaxDataDiskCount to be an explicit nil

### UnsetMaxDataDiskCount
`func (o *AzureVMSize) UnsetMaxDataDiskCount()`

UnsetMaxDataDiskCount ensures that no value is present for MaxDataDiskCount, not even an explicit nil
### GetHyperVGen1

`func (o *AzureVMSize) GetHyperVGen1() bool`

GetHyperVGen1 returns the HyperVGen1 field if non-nil, zero value otherwise.

### GetHyperVGen1Ok

`func (o *AzureVMSize) GetHyperVGen1Ok() (*bool, bool)`

GetHyperVGen1Ok returns a tuple with the HyperVGen1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperVGen1

`func (o *AzureVMSize) SetHyperVGen1(v bool)`

SetHyperVGen1 sets HyperVGen1 field to given value.

### HasHyperVGen1

`func (o *AzureVMSize) HasHyperVGen1() bool`

HasHyperVGen1 returns a boolean if a field has been set.

### GetHyperVGen2

`func (o *AzureVMSize) GetHyperVGen2() bool`

GetHyperVGen2 returns the HyperVGen2 field if non-nil, zero value otherwise.

### GetHyperVGen2Ok

`func (o *AzureVMSize) GetHyperVGen2Ok() (*bool, bool)`

GetHyperVGen2Ok returns a tuple with the HyperVGen2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperVGen2

`func (o *AzureVMSize) SetHyperVGen2(v bool)`

SetHyperVGen2 sets HyperVGen2 field to given value.

### HasHyperVGen2

`func (o *AzureVMSize) HasHyperVGen2() bool`

HasHyperVGen2 returns a boolean if a field has been set.

### GetNumberofGPUs

`func (o *AzureVMSize) GetNumberofGPUs() int32`

GetNumberofGPUs returns the NumberofGPUs field if non-nil, zero value otherwise.

### GetNumberofGPUsOk

`func (o *AzureVMSize) GetNumberofGPUsOk() (*int32, bool)`

GetNumberofGPUsOk returns a tuple with the NumberofGPUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberofGPUs

`func (o *AzureVMSize) SetNumberofGPUs(v int32)`

SetNumberofGPUs sets NumberofGPUs field to given value.

### HasNumberofGPUs

`func (o *AzureVMSize) HasNumberofGPUs() bool`

HasNumberofGPUs returns a boolean if a field has been set.

### SetNumberofGPUsNil

`func (o *AzureVMSize) SetNumberofGPUsNil(b bool)`

 SetNumberofGPUsNil sets the value for NumberofGPUs to be an explicit nil

### UnsetNumberofGPUs
`func (o *AzureVMSize) UnsetNumberofGPUs()`

UnsetNumberofGPUs ensures that no value is present for NumberofGPUs, not even an explicit nil
### GetGpuMemoryInMB

`func (o *AzureVMSize) GetGpuMemoryInMB() float64`

GetGpuMemoryInMB returns the GpuMemoryInMB field if non-nil, zero value otherwise.

### GetGpuMemoryInMBOk

`func (o *AzureVMSize) GetGpuMemoryInMBOk() (*float64, bool)`

GetGpuMemoryInMBOk returns a tuple with the GpuMemoryInMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuMemoryInMB

`func (o *AzureVMSize) SetGpuMemoryInMB(v float64)`

SetGpuMemoryInMB sets GpuMemoryInMB field to given value.

### HasGpuMemoryInMB

`func (o *AzureVMSize) HasGpuMemoryInMB() bool`

HasGpuMemoryInMB returns a boolean if a field has been set.

### SetGpuMemoryInMBNil

`func (o *AzureVMSize) SetGpuMemoryInMBNil(b bool)`

 SetGpuMemoryInMBNil sets the value for GpuMemoryInMB to be an explicit nil

### UnsetGpuMemoryInMB
`func (o *AzureVMSize) UnsetGpuMemoryInMB()`

UnsetGpuMemoryInMB ensures that no value is present for GpuMemoryInMB, not even an explicit nil
### GetSupportedSessionsPerVm

`func (o *AzureVMSize) GetSupportedSessionsPerVm() []int32`

GetSupportedSessionsPerVm returns the SupportedSessionsPerVm field if non-nil, zero value otherwise.

### GetSupportedSessionsPerVmOk

`func (o *AzureVMSize) GetSupportedSessionsPerVmOk() (*[]int32, bool)`

GetSupportedSessionsPerVmOk returns a tuple with the SupportedSessionsPerVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedSessionsPerVm

`func (o *AzureVMSize) SetSupportedSessionsPerVm(v []int32)`

SetSupportedSessionsPerVm sets SupportedSessionsPerVm field to given value.

### HasSupportedSessionsPerVm

`func (o *AzureVMSize) HasSupportedSessionsPerVm() bool`

HasSupportedSessionsPerVm returns a boolean if a field has been set.

### SetSupportedSessionsPerVmNil

`func (o *AzureVMSize) SetSupportedSessionsPerVmNil(b bool)`

 SetSupportedSessionsPerVmNil sets the value for SupportedSessionsPerVm to be an explicit nil

### UnsetSupportedSessionsPerVm
`func (o *AzureVMSize) UnsetSupportedSessionsPerVm()`

UnsetSupportedSessionsPerVm ensures that no value is present for SupportedSessionsPerVm, not even an explicit nil
### GetEstimatedCredits

`func (o *AzureVMSize) GetEstimatedCredits() float64`

GetEstimatedCredits returns the EstimatedCredits field if non-nil, zero value otherwise.

### GetEstimatedCreditsOk

`func (o *AzureVMSize) GetEstimatedCreditsOk() (*float64, bool)`

GetEstimatedCreditsOk returns a tuple with the EstimatedCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCredits

`func (o *AzureVMSize) SetEstimatedCredits(v float64)`

SetEstimatedCredits sets EstimatedCredits field to given value.

### HasEstimatedCredits

`func (o *AzureVMSize) HasEstimatedCredits() bool`

HasEstimatedCredits returns a boolean if a field has been set.

### SetEstimatedCreditsNil

`func (o *AzureVMSize) SetEstimatedCreditsNil(b bool)`

 SetEstimatedCreditsNil sets the value for EstimatedCredits to be an explicit nil

### UnsetEstimatedCredits
`func (o *AzureVMSize) UnsetEstimatedCredits()`

UnsetEstimatedCredits ensures that no value is present for EstimatedCredits, not even an explicit nil
### GetUsageType

`func (o *AzureVMSize) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *AzureVMSize) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *AzureVMSize) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *AzureVMSize) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### SetUsageTypeNil

`func (o *AzureVMSize) SetUsageTypeNil(b bool)`

 SetUsageTypeNil sets the value for UsageType to be an explicit nil

### UnsetUsageType
`func (o *AzureVMSize) UnsetUsageType()`

UnsetUsageType ensures that no value is present for UsageType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


