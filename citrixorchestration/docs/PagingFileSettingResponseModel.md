# PagingFileSettingResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **NullableString** | Paging file location like: ?:\\pagefile.sys, C:\\pagefile.sys | [optional] 
**InitialSizeInMb** | Pointer to **int32** | Paging file minimum size in MB. | [optional] 
**MaxSizeInMb** | Pointer to **int32** | Paging file maximum size in MB. | [optional] 
**IsSystemManaged** | Pointer to **bool** | If system managed, return true, OtherWise return false. | [optional] 

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

### GetPath

`func (o *PagingFileSettingResponseModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PagingFileSettingResponseModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PagingFileSettingResponseModel) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PagingFileSettingResponseModel) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *PagingFileSettingResponseModel) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *PagingFileSettingResponseModel) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetInitialSizeInMb

`func (o *PagingFileSettingResponseModel) GetInitialSizeInMb() int32`

GetInitialSizeInMb returns the InitialSizeInMb field if non-nil, zero value otherwise.

### GetInitialSizeInMbOk

`func (o *PagingFileSettingResponseModel) GetInitialSizeInMbOk() (*int32, bool)`

GetInitialSizeInMbOk returns a tuple with the InitialSizeInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialSizeInMb

`func (o *PagingFileSettingResponseModel) SetInitialSizeInMb(v int32)`

SetInitialSizeInMb sets InitialSizeInMb field to given value.

### HasInitialSizeInMb

`func (o *PagingFileSettingResponseModel) HasInitialSizeInMb() bool`

HasInitialSizeInMb returns a boolean if a field has been set.

### GetMaxSizeInMb

`func (o *PagingFileSettingResponseModel) GetMaxSizeInMb() int32`

GetMaxSizeInMb returns the MaxSizeInMb field if non-nil, zero value otherwise.

### GetMaxSizeInMbOk

`func (o *PagingFileSettingResponseModel) GetMaxSizeInMbOk() (*int32, bool)`

GetMaxSizeInMbOk returns a tuple with the MaxSizeInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSizeInMb

`func (o *PagingFileSettingResponseModel) SetMaxSizeInMb(v int32)`

SetMaxSizeInMb sets MaxSizeInMb field to given value.

### HasMaxSizeInMb

`func (o *PagingFileSettingResponseModel) HasMaxSizeInMb() bool`

HasMaxSizeInMb returns a boolean if a field has been set.

### GetIsSystemManaged

`func (o *PagingFileSettingResponseModel) GetIsSystemManaged() bool`

GetIsSystemManaged returns the IsSystemManaged field if non-nil, zero value otherwise.

### GetIsSystemManagedOk

`func (o *PagingFileSettingResponseModel) GetIsSystemManagedOk() (*bool, bool)`

GetIsSystemManagedOk returns a tuple with the IsSystemManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemManaged

`func (o *PagingFileSettingResponseModel) SetIsSystemManaged(v bool)`

SetIsSystemManaged sets IsSystemManaged field to given value.

### HasIsSystemManaged

`func (o *PagingFileSettingResponseModel) HasIsSystemManaged() bool`

HasIsSystemManaged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


