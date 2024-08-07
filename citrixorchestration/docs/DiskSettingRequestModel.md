# DiskSettingRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageType** | Pointer to [**DiskStorageType**](DiskStorageType.md) |  | [optional] 
**SizeInGB** | Pointer to **int32** | The size of the disk in GB (gigabytes) | [optional] 
**Location** | Pointer to **NullableString** | The location of the disk | [optional] 

## Methods

### NewDiskSettingRequestModel

`func NewDiskSettingRequestModel() *DiskSettingRequestModel`

NewDiskSettingRequestModel instantiates a new DiskSettingRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskSettingRequestModelWithDefaults

`func NewDiskSettingRequestModelWithDefaults() *DiskSettingRequestModel`

NewDiskSettingRequestModelWithDefaults instantiates a new DiskSettingRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageType

`func (o *DiskSettingRequestModel) GetStorageType() DiskStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *DiskSettingRequestModel) GetStorageTypeOk() (*DiskStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *DiskSettingRequestModel) SetStorageType(v DiskStorageType)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *DiskSettingRequestModel) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetSizeInGB

`func (o *DiskSettingRequestModel) GetSizeInGB() int32`

GetSizeInGB returns the SizeInGB field if non-nil, zero value otherwise.

### GetSizeInGBOk

`func (o *DiskSettingRequestModel) GetSizeInGBOk() (*int32, bool)`

GetSizeInGBOk returns a tuple with the SizeInGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInGB

`func (o *DiskSettingRequestModel) SetSizeInGB(v int32)`

SetSizeInGB sets SizeInGB field to given value.

### HasSizeInGB

`func (o *DiskSettingRequestModel) HasSizeInGB() bool`

HasSizeInGB returns a boolean if a field has been set.

### GetLocation

`func (o *DiskSettingRequestModel) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DiskSettingRequestModel) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DiskSettingRequestModel) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DiskSettingRequestModel) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *DiskSettingRequestModel) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *DiskSettingRequestModel) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


