# DiskPriceResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageType** | Pointer to [**DiskStorageType**](DiskStorageType.md) |  | [optional] 
**SizeInGB** | Pointer to **int32** | The size of the disk in GB (gigabytes) | [optional] 
**Location** | Pointer to **NullableString** | The location of the disk | [optional] 
**ResourcePrices** | Pointer to [**[]ResourcePriceResponseModel**](ResourcePriceResponseModel.md) | Gets or sets pricing information of billing components. | [optional] 

## Methods

### NewDiskPriceResponseModel

`func NewDiskPriceResponseModel() *DiskPriceResponseModel`

NewDiskPriceResponseModel instantiates a new DiskPriceResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskPriceResponseModelWithDefaults

`func NewDiskPriceResponseModelWithDefaults() *DiskPriceResponseModel`

NewDiskPriceResponseModelWithDefaults instantiates a new DiskPriceResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageType

`func (o *DiskPriceResponseModel) GetStorageType() DiskStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *DiskPriceResponseModel) GetStorageTypeOk() (*DiskStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *DiskPriceResponseModel) SetStorageType(v DiskStorageType)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *DiskPriceResponseModel) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetSizeInGB

`func (o *DiskPriceResponseModel) GetSizeInGB() int32`

GetSizeInGB returns the SizeInGB field if non-nil, zero value otherwise.

### GetSizeInGBOk

`func (o *DiskPriceResponseModel) GetSizeInGBOk() (*int32, bool)`

GetSizeInGBOk returns a tuple with the SizeInGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInGB

`func (o *DiskPriceResponseModel) SetSizeInGB(v int32)`

SetSizeInGB sets SizeInGB field to given value.

### HasSizeInGB

`func (o *DiskPriceResponseModel) HasSizeInGB() bool`

HasSizeInGB returns a boolean if a field has been set.

### GetLocation

`func (o *DiskPriceResponseModel) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DiskPriceResponseModel) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DiskPriceResponseModel) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DiskPriceResponseModel) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *DiskPriceResponseModel) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *DiskPriceResponseModel) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetResourcePrices

`func (o *DiskPriceResponseModel) GetResourcePrices() []ResourcePriceResponseModel`

GetResourcePrices returns the ResourcePrices field if non-nil, zero value otherwise.

### GetResourcePricesOk

`func (o *DiskPriceResponseModel) GetResourcePricesOk() (*[]ResourcePriceResponseModel, bool)`

GetResourcePricesOk returns a tuple with the ResourcePrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePrices

`func (o *DiskPriceResponseModel) SetResourcePrices(v []ResourcePriceResponseModel)`

SetResourcePrices sets ResourcePrices field to given value.

### HasResourcePrices

`func (o *DiskPriceResponseModel) HasResourcePrices() bool`

HasResourcePrices returns a boolean if a field has been set.

### SetResourcePricesNil

`func (o *DiskPriceResponseModel) SetResourcePricesNil(b bool)`

 SetResourcePricesNil sets the value for ResourcePrices to be an explicit nil

### UnsetResourcePrices
`func (o *DiskPriceResponseModel) UnsetResourcePrices()`

UnsetResourcePrices ensures that no value is present for ResourcePrices, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


