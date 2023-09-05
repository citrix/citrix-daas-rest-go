# PagingFileSettingResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to **string** | Paging file location like: ?:\\pagefile.sys, C:\\pagefile.sys | [optional] 
**MinSize** | Pointer to **int32** | Paging file minimum size in MB. | [optional] 
**MaxSize** | Pointer to **int32** | Paging file maximum size in MB. | [optional] 

## Methods

### NewPagingFileSettingResponseModel

`func NewPagingFileSettingResponseModel() *PagingFileSettingResponseModel`

NewPagingFileSettingResponseModel instantiates a new PagingFileSettingResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagingFileSettingResponseModelWithDefaults

`func NewPagingFileSettingResponseModelWithDefaults() *PagingFileSettingResponseModel`

NewPagingFileSettingResponseModelWithDefaults instantiates a new PagingFileSettingResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *PagingFileSettingResponseModel) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PagingFileSettingResponseModel) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PagingFileSettingResponseModel) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PagingFileSettingResponseModel) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMinSize

`func (o *PagingFileSettingResponseModel) GetMinSize() int32`

GetMinSize returns the MinSize field if non-nil, zero value otherwise.

### GetMinSizeOk

`func (o *PagingFileSettingResponseModel) GetMinSizeOk() (*int32, bool)`

GetMinSizeOk returns a tuple with the MinSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSize

`func (o *PagingFileSettingResponseModel) SetMinSize(v int32)`

SetMinSize sets MinSize field to given value.

### HasMinSize

`func (o *PagingFileSettingResponseModel) HasMinSize() bool`

HasMinSize returns a boolean if a field has been set.

### GetMaxSize

`func (o *PagingFileSettingResponseModel) GetMaxSize() int32`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *PagingFileSettingResponseModel) GetMaxSizeOk() (*int32, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *PagingFileSettingResponseModel) SetMaxSize(v int32)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *PagingFileSettingResponseModel) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


