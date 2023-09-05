# ImageVersionBaseResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | The version number associated with the image version. | [optional] 
**Id** | Pointer to **string** | The Id of the image version. | [optional] 
**MasterImagePath** | Pointer to **string** | The MasterImageVM of the image version. | [optional] 
**Description** | Pointer to **string** | The Description of the image version. | [optional] 
**DiskSizeGB** | Pointer to **int32** | Size of the VM&#39;s OS disk, in GB. | [optional] 
**CreateTime** | Pointer to **string** | The create time of the image version. | [optional] 
**WriteBackCacheDiskSize** | Pointer to **int32** | The WriteBackCacheDiskSize of the image version. | [optional] 
**WriteBackCacheMemorySize** | Pointer to **int32** | The WriteBackCacheMemorySize of the image version. | [optional] 
**ImageStatus** | Pointer to **string** | The image status of the image version. | [optional] 
**Error** | Pointer to **string** | The error info of the image version. | [optional] 
**AdditionalData** | Pointer to **map[string]string** | The additional data of the image version. | [optional] 

## Methods

### NewImageVersionBaseResponseModel

`func NewImageVersionBaseResponseModel() *ImageVersionBaseResponseModel`

NewImageVersionBaseResponseModel instantiates a new ImageVersionBaseResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageVersionBaseResponseModelWithDefaults

`func NewImageVersionBaseResponseModelWithDefaults() *ImageVersionBaseResponseModel`

NewImageVersionBaseResponseModelWithDefaults instantiates a new ImageVersionBaseResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ImageVersionBaseResponseModel) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ImageVersionBaseResponseModel) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ImageVersionBaseResponseModel) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ImageVersionBaseResponseModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetId

`func (o *ImageVersionBaseResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageVersionBaseResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageVersionBaseResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImageVersionBaseResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMasterImagePath

`func (o *ImageVersionBaseResponseModel) GetMasterImagePath() string`

GetMasterImagePath returns the MasterImagePath field if non-nil, zero value otherwise.

### GetMasterImagePathOk

`func (o *ImageVersionBaseResponseModel) GetMasterImagePathOk() (*string, bool)`

GetMasterImagePathOk returns a tuple with the MasterImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterImagePath

`func (o *ImageVersionBaseResponseModel) SetMasterImagePath(v string)`

SetMasterImagePath sets MasterImagePath field to given value.

### HasMasterImagePath

`func (o *ImageVersionBaseResponseModel) HasMasterImagePath() bool`

HasMasterImagePath returns a boolean if a field has been set.

### GetDescription

`func (o *ImageVersionBaseResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImageVersionBaseResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImageVersionBaseResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImageVersionBaseResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiskSizeGB

`func (o *ImageVersionBaseResponseModel) GetDiskSizeGB() int32`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *ImageVersionBaseResponseModel) GetDiskSizeGBOk() (*int32, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *ImageVersionBaseResponseModel) SetDiskSizeGB(v int32)`

SetDiskSizeGB sets DiskSizeGB field to given value.

### HasDiskSizeGB

`func (o *ImageVersionBaseResponseModel) HasDiskSizeGB() bool`

HasDiskSizeGB returns a boolean if a field has been set.

### GetCreateTime

`func (o *ImageVersionBaseResponseModel) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ImageVersionBaseResponseModel) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ImageVersionBaseResponseModel) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ImageVersionBaseResponseModel) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetWriteBackCacheDiskSize

`func (o *ImageVersionBaseResponseModel) GetWriteBackCacheDiskSize() int32`

GetWriteBackCacheDiskSize returns the WriteBackCacheDiskSize field if non-nil, zero value otherwise.

### GetWriteBackCacheDiskSizeOk

`func (o *ImageVersionBaseResponseModel) GetWriteBackCacheDiskSizeOk() (*int32, bool)`

GetWriteBackCacheDiskSizeOk returns a tuple with the WriteBackCacheDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheDiskSize

`func (o *ImageVersionBaseResponseModel) SetWriteBackCacheDiskSize(v int32)`

SetWriteBackCacheDiskSize sets WriteBackCacheDiskSize field to given value.

### HasWriteBackCacheDiskSize

`func (o *ImageVersionBaseResponseModel) HasWriteBackCacheDiskSize() bool`

HasWriteBackCacheDiskSize returns a boolean if a field has been set.

### GetWriteBackCacheMemorySize

`func (o *ImageVersionBaseResponseModel) GetWriteBackCacheMemorySize() int32`

GetWriteBackCacheMemorySize returns the WriteBackCacheMemorySize field if non-nil, zero value otherwise.

### GetWriteBackCacheMemorySizeOk

`func (o *ImageVersionBaseResponseModel) GetWriteBackCacheMemorySizeOk() (*int32, bool)`

GetWriteBackCacheMemorySizeOk returns a tuple with the WriteBackCacheMemorySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheMemorySize

`func (o *ImageVersionBaseResponseModel) SetWriteBackCacheMemorySize(v int32)`

SetWriteBackCacheMemorySize sets WriteBackCacheMemorySize field to given value.

### HasWriteBackCacheMemorySize

`func (o *ImageVersionBaseResponseModel) HasWriteBackCacheMemorySize() bool`

HasWriteBackCacheMemorySize returns a boolean if a field has been set.

### GetImageStatus

`func (o *ImageVersionBaseResponseModel) GetImageStatus() string`

GetImageStatus returns the ImageStatus field if non-nil, zero value otherwise.

### GetImageStatusOk

`func (o *ImageVersionBaseResponseModel) GetImageStatusOk() (*string, bool)`

GetImageStatusOk returns a tuple with the ImageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStatus

`func (o *ImageVersionBaseResponseModel) SetImageStatus(v string)`

SetImageStatus sets ImageStatus field to given value.

### HasImageStatus

`func (o *ImageVersionBaseResponseModel) HasImageStatus() bool`

HasImageStatus returns a boolean if a field has been set.

### GetError

`func (o *ImageVersionBaseResponseModel) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ImageVersionBaseResponseModel) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ImageVersionBaseResponseModel) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ImageVersionBaseResponseModel) HasError() bool`

HasError returns a boolean if a field has been set.

### GetAdditionalData

`func (o *ImageVersionBaseResponseModel) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *ImageVersionBaseResponseModel) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *ImageVersionBaseResponseModel) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *ImageVersionBaseResponseModel) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


