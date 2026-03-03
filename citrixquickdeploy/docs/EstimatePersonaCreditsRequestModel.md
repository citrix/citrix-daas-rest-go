# EstimatePersonaCreditsRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmSku** | Pointer to **NullableString** |  | [optional] 
**DiskSku** | Pointer to **NullableString** |  | [optional] 
**DiskSizeInGB** | Pointer to **int32** |  | [optional] 
**SessionSupport** | Pointer to [**SessionSupport**](SessionSupport.md) | Quantity of sessions supported per-machine. | [optional] 
**SessionsPerVM** | Pointer to **int32** |  | [optional] 
**AllocationType** | Pointer to [**CatalogAllocationType**](CatalogAllocationType.md) |  | [optional] 

## Methods

### NewEstimatePersonaCreditsRequestModel

`func NewEstimatePersonaCreditsRequestModel() *EstimatePersonaCreditsRequestModel`

NewEstimatePersonaCreditsRequestModel instantiates a new EstimatePersonaCreditsRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimatePersonaCreditsRequestModelWithDefaults

`func NewEstimatePersonaCreditsRequestModelWithDefaults() *EstimatePersonaCreditsRequestModel`

NewEstimatePersonaCreditsRequestModelWithDefaults instantiates a new EstimatePersonaCreditsRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmSku

`func (o *EstimatePersonaCreditsRequestModel) GetVmSku() string`

GetVmSku returns the VmSku field if non-nil, zero value otherwise.

### GetVmSkuOk

`func (o *EstimatePersonaCreditsRequestModel) GetVmSkuOk() (*string, bool)`

GetVmSkuOk returns a tuple with the VmSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSku

`func (o *EstimatePersonaCreditsRequestModel) SetVmSku(v string)`

SetVmSku sets VmSku field to given value.

### HasVmSku

`func (o *EstimatePersonaCreditsRequestModel) HasVmSku() bool`

HasVmSku returns a boolean if a field has been set.

### SetVmSkuNil

`func (o *EstimatePersonaCreditsRequestModel) SetVmSkuNil(b bool)`

 SetVmSkuNil sets the value for VmSku to be an explicit nil

### UnsetVmSku
`func (o *EstimatePersonaCreditsRequestModel) UnsetVmSku()`

UnsetVmSku ensures that no value is present for VmSku, not even an explicit nil
### GetDiskSku

`func (o *EstimatePersonaCreditsRequestModel) GetDiskSku() string`

GetDiskSku returns the DiskSku field if non-nil, zero value otherwise.

### GetDiskSkuOk

`func (o *EstimatePersonaCreditsRequestModel) GetDiskSkuOk() (*string, bool)`

GetDiskSkuOk returns a tuple with the DiskSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSku

`func (o *EstimatePersonaCreditsRequestModel) SetDiskSku(v string)`

SetDiskSku sets DiskSku field to given value.

### HasDiskSku

`func (o *EstimatePersonaCreditsRequestModel) HasDiskSku() bool`

HasDiskSku returns a boolean if a field has been set.

### SetDiskSkuNil

`func (o *EstimatePersonaCreditsRequestModel) SetDiskSkuNil(b bool)`

 SetDiskSkuNil sets the value for DiskSku to be an explicit nil

### UnsetDiskSku
`func (o *EstimatePersonaCreditsRequestModel) UnsetDiskSku()`

UnsetDiskSku ensures that no value is present for DiskSku, not even an explicit nil
### GetDiskSizeInGB

`func (o *EstimatePersonaCreditsRequestModel) GetDiskSizeInGB() int32`

GetDiskSizeInGB returns the DiskSizeInGB field if non-nil, zero value otherwise.

### GetDiskSizeInGBOk

`func (o *EstimatePersonaCreditsRequestModel) GetDiskSizeInGBOk() (*int32, bool)`

GetDiskSizeInGBOk returns a tuple with the DiskSizeInGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeInGB

`func (o *EstimatePersonaCreditsRequestModel) SetDiskSizeInGB(v int32)`

SetDiskSizeInGB sets DiskSizeInGB field to given value.

### HasDiskSizeInGB

`func (o *EstimatePersonaCreditsRequestModel) HasDiskSizeInGB() bool`

HasDiskSizeInGB returns a boolean if a field has been set.

### GetSessionSupport

`func (o *EstimatePersonaCreditsRequestModel) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *EstimatePersonaCreditsRequestModel) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *EstimatePersonaCreditsRequestModel) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.

### HasSessionSupport

`func (o *EstimatePersonaCreditsRequestModel) HasSessionSupport() bool`

HasSessionSupport returns a boolean if a field has been set.

### GetSessionsPerVM

`func (o *EstimatePersonaCreditsRequestModel) GetSessionsPerVM() int32`

GetSessionsPerVM returns the SessionsPerVM field if non-nil, zero value otherwise.

### GetSessionsPerVMOk

`func (o *EstimatePersonaCreditsRequestModel) GetSessionsPerVMOk() (*int32, bool)`

GetSessionsPerVMOk returns a tuple with the SessionsPerVM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionsPerVM

`func (o *EstimatePersonaCreditsRequestModel) SetSessionsPerVM(v int32)`

SetSessionsPerVM sets SessionsPerVM field to given value.

### HasSessionsPerVM

`func (o *EstimatePersonaCreditsRequestModel) HasSessionsPerVM() bool`

HasSessionsPerVM returns a boolean if a field has been set.

### GetAllocationType

`func (o *EstimatePersonaCreditsRequestModel) GetAllocationType() CatalogAllocationType`

GetAllocationType returns the AllocationType field if non-nil, zero value otherwise.

### GetAllocationTypeOk

`func (o *EstimatePersonaCreditsRequestModel) GetAllocationTypeOk() (*CatalogAllocationType, bool)`

GetAllocationTypeOk returns a tuple with the AllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationType

`func (o *EstimatePersonaCreditsRequestModel) SetAllocationType(v CatalogAllocationType)`

SetAllocationType sets AllocationType field to given value.

### HasAllocationType

`func (o *EstimatePersonaCreditsRequestModel) HasAllocationType() bool`

HasAllocationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


