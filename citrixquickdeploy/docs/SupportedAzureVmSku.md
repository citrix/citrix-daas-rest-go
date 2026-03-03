# SupportedAzureVmSku

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmId** | Pointer to **NullableString** | The Azure VM SKU ID. Ex: \&quot;Standard_D2as_v5\&quot; | [optional] 
**VmSku** | Pointer to **NullableString** | The Azure VM SKU name. Ex: \&quot;d2asv5\&quot; | [optional] 
**DefaultBackupVmSku** | Pointer to **NullableString** | The Azure VM SKU(s) that will be used as default for this VM&#39;s backup configuration.  Supports comma-separated values for multiple fallback options. Ex: \&quot;e2asv5\&quot; or \&quot;e2asv5,e2sv5,e2sv3\&quot; | [optional] 
**IsBackupSku** | Pointer to **bool** | Flag indicating if the SKU is for backup VM configuration. | [optional] 
**IsFlexSku** | Pointer to **bool** | Flag inticating if the SKU is for Flex customer use only.  Feature toggle: \&quot;cmaphoenix\&quot;. | [optional] 
**EstimatedCredits** | Pointer to **int32** | The cost per user per month for the VM SKU | [optional] 
**SupportedSessionsPerVm** | Pointer to **[]int32** | Specifies the supported sessionsPerVm values if using this SKU for multisession catalog | [optional] 
**IsHidden** | Pointer to **bool** | Flag indicating if the SKU is hidden from the UI. | [optional] 

## Methods

### NewSupportedAzureVmSku

`func NewSupportedAzureVmSku() *SupportedAzureVmSku`

NewSupportedAzureVmSku instantiates a new SupportedAzureVmSku object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedAzureVmSkuWithDefaults

`func NewSupportedAzureVmSkuWithDefaults() *SupportedAzureVmSku`

NewSupportedAzureVmSkuWithDefaults instantiates a new SupportedAzureVmSku object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmId

`func (o *SupportedAzureVmSku) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *SupportedAzureVmSku) GetVmIdOk() (*string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *SupportedAzureVmSku) SetVmId(v string)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *SupportedAzureVmSku) HasVmId() bool`

HasVmId returns a boolean if a field has been set.

### SetVmIdNil

`func (o *SupportedAzureVmSku) SetVmIdNil(b bool)`

 SetVmIdNil sets the value for VmId to be an explicit nil

### UnsetVmId
`func (o *SupportedAzureVmSku) UnsetVmId()`

UnsetVmId ensures that no value is present for VmId, not even an explicit nil
### GetVmSku

`func (o *SupportedAzureVmSku) GetVmSku() string`

GetVmSku returns the VmSku field if non-nil, zero value otherwise.

### GetVmSkuOk

`func (o *SupportedAzureVmSku) GetVmSkuOk() (*string, bool)`

GetVmSkuOk returns a tuple with the VmSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSku

`func (o *SupportedAzureVmSku) SetVmSku(v string)`

SetVmSku sets VmSku field to given value.

### HasVmSku

`func (o *SupportedAzureVmSku) HasVmSku() bool`

HasVmSku returns a boolean if a field has been set.

### SetVmSkuNil

`func (o *SupportedAzureVmSku) SetVmSkuNil(b bool)`

 SetVmSkuNil sets the value for VmSku to be an explicit nil

### UnsetVmSku
`func (o *SupportedAzureVmSku) UnsetVmSku()`

UnsetVmSku ensures that no value is present for VmSku, not even an explicit nil
### GetDefaultBackupVmSku

`func (o *SupportedAzureVmSku) GetDefaultBackupVmSku() string`

GetDefaultBackupVmSku returns the DefaultBackupVmSku field if non-nil, zero value otherwise.

### GetDefaultBackupVmSkuOk

`func (o *SupportedAzureVmSku) GetDefaultBackupVmSkuOk() (*string, bool)`

GetDefaultBackupVmSkuOk returns a tuple with the DefaultBackupVmSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBackupVmSku

`func (o *SupportedAzureVmSku) SetDefaultBackupVmSku(v string)`

SetDefaultBackupVmSku sets DefaultBackupVmSku field to given value.

### HasDefaultBackupVmSku

`func (o *SupportedAzureVmSku) HasDefaultBackupVmSku() bool`

HasDefaultBackupVmSku returns a boolean if a field has been set.

### SetDefaultBackupVmSkuNil

`func (o *SupportedAzureVmSku) SetDefaultBackupVmSkuNil(b bool)`

 SetDefaultBackupVmSkuNil sets the value for DefaultBackupVmSku to be an explicit nil

### UnsetDefaultBackupVmSku
`func (o *SupportedAzureVmSku) UnsetDefaultBackupVmSku()`

UnsetDefaultBackupVmSku ensures that no value is present for DefaultBackupVmSku, not even an explicit nil
### GetIsBackupSku

`func (o *SupportedAzureVmSku) GetIsBackupSku() bool`

GetIsBackupSku returns the IsBackupSku field if non-nil, zero value otherwise.

### GetIsBackupSkuOk

`func (o *SupportedAzureVmSku) GetIsBackupSkuOk() (*bool, bool)`

GetIsBackupSkuOk returns a tuple with the IsBackupSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBackupSku

`func (o *SupportedAzureVmSku) SetIsBackupSku(v bool)`

SetIsBackupSku sets IsBackupSku field to given value.

### HasIsBackupSku

`func (o *SupportedAzureVmSku) HasIsBackupSku() bool`

HasIsBackupSku returns a boolean if a field has been set.

### GetIsFlexSku

`func (o *SupportedAzureVmSku) GetIsFlexSku() bool`

GetIsFlexSku returns the IsFlexSku field if non-nil, zero value otherwise.

### GetIsFlexSkuOk

`func (o *SupportedAzureVmSku) GetIsFlexSkuOk() (*bool, bool)`

GetIsFlexSkuOk returns a tuple with the IsFlexSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlexSku

`func (o *SupportedAzureVmSku) SetIsFlexSku(v bool)`

SetIsFlexSku sets IsFlexSku field to given value.

### HasIsFlexSku

`func (o *SupportedAzureVmSku) HasIsFlexSku() bool`

HasIsFlexSku returns a boolean if a field has been set.

### GetEstimatedCredits

`func (o *SupportedAzureVmSku) GetEstimatedCredits() int32`

GetEstimatedCredits returns the EstimatedCredits field if non-nil, zero value otherwise.

### GetEstimatedCreditsOk

`func (o *SupportedAzureVmSku) GetEstimatedCreditsOk() (*int32, bool)`

GetEstimatedCreditsOk returns a tuple with the EstimatedCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCredits

`func (o *SupportedAzureVmSku) SetEstimatedCredits(v int32)`

SetEstimatedCredits sets EstimatedCredits field to given value.

### HasEstimatedCredits

`func (o *SupportedAzureVmSku) HasEstimatedCredits() bool`

HasEstimatedCredits returns a boolean if a field has been set.

### GetSupportedSessionsPerVm

`func (o *SupportedAzureVmSku) GetSupportedSessionsPerVm() []int32`

GetSupportedSessionsPerVm returns the SupportedSessionsPerVm field if non-nil, zero value otherwise.

### GetSupportedSessionsPerVmOk

`func (o *SupportedAzureVmSku) GetSupportedSessionsPerVmOk() (*[]int32, bool)`

GetSupportedSessionsPerVmOk returns a tuple with the SupportedSessionsPerVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedSessionsPerVm

`func (o *SupportedAzureVmSku) SetSupportedSessionsPerVm(v []int32)`

SetSupportedSessionsPerVm sets SupportedSessionsPerVm field to given value.

### HasSupportedSessionsPerVm

`func (o *SupportedAzureVmSku) HasSupportedSessionsPerVm() bool`

HasSupportedSessionsPerVm returns a boolean if a field has been set.

### SetSupportedSessionsPerVmNil

`func (o *SupportedAzureVmSku) SetSupportedSessionsPerVmNil(b bool)`

 SetSupportedSessionsPerVmNil sets the value for SupportedSessionsPerVm to be an explicit nil

### UnsetSupportedSessionsPerVm
`func (o *SupportedAzureVmSku) UnsetSupportedSessionsPerVm()`

UnsetSupportedSessionsPerVm ensures that no value is present for SupportedSessionsPerVm, not even an explicit nil
### GetIsHidden

`func (o *SupportedAzureVmSku) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *SupportedAzureVmSku) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *SupportedAzureVmSku) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *SupportedAzureVmSku) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


