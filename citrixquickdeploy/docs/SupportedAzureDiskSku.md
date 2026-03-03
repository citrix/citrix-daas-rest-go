# SupportedAzureDiskSku

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskSku** | Pointer to **NullableString** | The Azure Disk SKU. Eg: Premium_SSD | [optional] 
**DiskSizeInGB** | Pointer to **int32** | The size of the Disk | [optional] 
**EstimatedCredits** | Pointer to **int32** | The cost of the disk SKU | [optional] 
**IsHidden** | Pointer to **bool** | Flag indicating if the SKU is hidden from the UI | [optional] 

## Methods

### NewSupportedAzureDiskSku

`func NewSupportedAzureDiskSku() *SupportedAzureDiskSku`

NewSupportedAzureDiskSku instantiates a new SupportedAzureDiskSku object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedAzureDiskSkuWithDefaults

`func NewSupportedAzureDiskSkuWithDefaults() *SupportedAzureDiskSku`

NewSupportedAzureDiskSkuWithDefaults instantiates a new SupportedAzureDiskSku object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskSku

`func (o *SupportedAzureDiskSku) GetDiskSku() string`

GetDiskSku returns the DiskSku field if non-nil, zero value otherwise.

### GetDiskSkuOk

`func (o *SupportedAzureDiskSku) GetDiskSkuOk() (*string, bool)`

GetDiskSkuOk returns a tuple with the DiskSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSku

`func (o *SupportedAzureDiskSku) SetDiskSku(v string)`

SetDiskSku sets DiskSku field to given value.

### HasDiskSku

`func (o *SupportedAzureDiskSku) HasDiskSku() bool`

HasDiskSku returns a boolean if a field has been set.

### SetDiskSkuNil

`func (o *SupportedAzureDiskSku) SetDiskSkuNil(b bool)`

 SetDiskSkuNil sets the value for DiskSku to be an explicit nil

### UnsetDiskSku
`func (o *SupportedAzureDiskSku) UnsetDiskSku()`

UnsetDiskSku ensures that no value is present for DiskSku, not even an explicit nil
### GetDiskSizeInGB

`func (o *SupportedAzureDiskSku) GetDiskSizeInGB() int32`

GetDiskSizeInGB returns the DiskSizeInGB field if non-nil, zero value otherwise.

### GetDiskSizeInGBOk

`func (o *SupportedAzureDiskSku) GetDiskSizeInGBOk() (*int32, bool)`

GetDiskSizeInGBOk returns a tuple with the DiskSizeInGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeInGB

`func (o *SupportedAzureDiskSku) SetDiskSizeInGB(v int32)`

SetDiskSizeInGB sets DiskSizeInGB field to given value.

### HasDiskSizeInGB

`func (o *SupportedAzureDiskSku) HasDiskSizeInGB() bool`

HasDiskSizeInGB returns a boolean if a field has been set.

### GetEstimatedCredits

`func (o *SupportedAzureDiskSku) GetEstimatedCredits() int32`

GetEstimatedCredits returns the EstimatedCredits field if non-nil, zero value otherwise.

### GetEstimatedCreditsOk

`func (o *SupportedAzureDiskSku) GetEstimatedCreditsOk() (*int32, bool)`

GetEstimatedCreditsOk returns a tuple with the EstimatedCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCredits

`func (o *SupportedAzureDiskSku) SetEstimatedCredits(v int32)`

SetEstimatedCredits sets EstimatedCredits field to given value.

### HasEstimatedCredits

`func (o *SupportedAzureDiskSku) HasEstimatedCredits() bool`

HasEstimatedCredits returns a boolean if a field has been set.

### GetIsHidden

`func (o *SupportedAzureDiskSku) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *SupportedAzureDiskSku) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *SupportedAzureDiskSku) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *SupportedAzureDiskSku) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


