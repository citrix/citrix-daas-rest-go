# ImageVersionResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **NullableString** | The version number associated with the image version. | [optional] 
**Id** | Pointer to **string** | The Id of the image version. | [optional] 
**MasterImagePath** | Pointer to **NullableString** | The MasterImageVM of the image version. | [optional] 
**Description** | Pointer to **NullableString** | The Description of the image version. | [optional] 
**DiskSizeGB** | Pointer to **int32** | Size of the VM&#39;s OS disk, in GB. | [optional] 
**CreateTime** | Pointer to **NullableString** | The create time of the image version. | [optional] 
**WriteBackCacheDiskSize** | Pointer to **int32** | The WriteBackCacheDiskSize of the image version. | [optional] 
**WriteBackCacheMemorySize** | Pointer to **int32** | The WriteBackCacheMemorySize of the image version. | [optional] 
**ImageStatus** | Pointer to **NullableString** | The image status of the image version. | [optional] 
**Error** | Pointer to **NullableString** | The error info of the image version. | [optional] 
**AdditionalData** | Pointer to **map[string]string** | The additional data of the image version. | [optional] 
**ImageDefinition** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 
**ImageScheme** | Pointer to [**ImageSchemeResponseModel**](ImageSchemeResponseModel.md) |  | [optional] 

## Methods

### NewImageVersionResponseModel

`func NewImageVersionResponseModel() *ImageVersionResponseModel`

NewImageVersionResponseModel instantiates a new ImageVersionResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageVersionResponseModelWithDefaults

`func NewImageVersionResponseModelWithDefaults() *ImageVersionResponseModel`

NewImageVersionResponseModelWithDefaults instantiates a new ImageVersionResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ImageVersionResponseModel) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ImageVersionResponseModel) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ImageVersionResponseModel) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ImageVersionResponseModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ImageVersionResponseModel) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ImageVersionResponseModel) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetId

`func (o *ImageVersionResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageVersionResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageVersionResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImageVersionResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMasterImagePath

`func (o *ImageVersionResponseModel) GetMasterImagePath() string`

GetMasterImagePath returns the MasterImagePath field if non-nil, zero value otherwise.

### GetMasterImagePathOk

`func (o *ImageVersionResponseModel) GetMasterImagePathOk() (*string, bool)`

GetMasterImagePathOk returns a tuple with the MasterImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterImagePath

`func (o *ImageVersionResponseModel) SetMasterImagePath(v string)`

SetMasterImagePath sets MasterImagePath field to given value.

### HasMasterImagePath

`func (o *ImageVersionResponseModel) HasMasterImagePath() bool`

HasMasterImagePath returns a boolean if a field has been set.

### SetMasterImagePathNil

`func (o *ImageVersionResponseModel) SetMasterImagePathNil(b bool)`

 SetMasterImagePathNil sets the value for MasterImagePath to be an explicit nil

### UnsetMasterImagePath
`func (o *ImageVersionResponseModel) UnsetMasterImagePath()`

UnsetMasterImagePath ensures that no value is present for MasterImagePath, not even an explicit nil
### GetDescription

`func (o *ImageVersionResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImageVersionResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImageVersionResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImageVersionResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ImageVersionResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ImageVersionResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDiskSizeGB

`func (o *ImageVersionResponseModel) GetDiskSizeGB() int32`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *ImageVersionResponseModel) GetDiskSizeGBOk() (*int32, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *ImageVersionResponseModel) SetDiskSizeGB(v int32)`

SetDiskSizeGB sets DiskSizeGB field to given value.

### HasDiskSizeGB

`func (o *ImageVersionResponseModel) HasDiskSizeGB() bool`

HasDiskSizeGB returns a boolean if a field has been set.

### GetCreateTime

`func (o *ImageVersionResponseModel) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ImageVersionResponseModel) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ImageVersionResponseModel) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ImageVersionResponseModel) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### SetCreateTimeNil

`func (o *ImageVersionResponseModel) SetCreateTimeNil(b bool)`

 SetCreateTimeNil sets the value for CreateTime to be an explicit nil

### UnsetCreateTime
`func (o *ImageVersionResponseModel) UnsetCreateTime()`

UnsetCreateTime ensures that no value is present for CreateTime, not even an explicit nil
### GetWriteBackCacheDiskSize

`func (o *ImageVersionResponseModel) GetWriteBackCacheDiskSize() int32`

GetWriteBackCacheDiskSize returns the WriteBackCacheDiskSize field if non-nil, zero value otherwise.

### GetWriteBackCacheDiskSizeOk

`func (o *ImageVersionResponseModel) GetWriteBackCacheDiskSizeOk() (*int32, bool)`

GetWriteBackCacheDiskSizeOk returns a tuple with the WriteBackCacheDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheDiskSize

`func (o *ImageVersionResponseModel) SetWriteBackCacheDiskSize(v int32)`

SetWriteBackCacheDiskSize sets WriteBackCacheDiskSize field to given value.

### HasWriteBackCacheDiskSize

`func (o *ImageVersionResponseModel) HasWriteBackCacheDiskSize() bool`

HasWriteBackCacheDiskSize returns a boolean if a field has been set.

### GetWriteBackCacheMemorySize

`func (o *ImageVersionResponseModel) GetWriteBackCacheMemorySize() int32`

GetWriteBackCacheMemorySize returns the WriteBackCacheMemorySize field if non-nil, zero value otherwise.

### GetWriteBackCacheMemorySizeOk

`func (o *ImageVersionResponseModel) GetWriteBackCacheMemorySizeOk() (*int32, bool)`

GetWriteBackCacheMemorySizeOk returns a tuple with the WriteBackCacheMemorySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheMemorySize

`func (o *ImageVersionResponseModel) SetWriteBackCacheMemorySize(v int32)`

SetWriteBackCacheMemorySize sets WriteBackCacheMemorySize field to given value.

### HasWriteBackCacheMemorySize

`func (o *ImageVersionResponseModel) HasWriteBackCacheMemorySize() bool`

HasWriteBackCacheMemorySize returns a boolean if a field has been set.

### GetImageStatus

`func (o *ImageVersionResponseModel) GetImageStatus() string`

GetImageStatus returns the ImageStatus field if non-nil, zero value otherwise.

### GetImageStatusOk

`func (o *ImageVersionResponseModel) GetImageStatusOk() (*string, bool)`

GetImageStatusOk returns a tuple with the ImageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStatus

`func (o *ImageVersionResponseModel) SetImageStatus(v string)`

SetImageStatus sets ImageStatus field to given value.

### HasImageStatus

`func (o *ImageVersionResponseModel) HasImageStatus() bool`

HasImageStatus returns a boolean if a field has been set.

### SetImageStatusNil

`func (o *ImageVersionResponseModel) SetImageStatusNil(b bool)`

 SetImageStatusNil sets the value for ImageStatus to be an explicit nil

### UnsetImageStatus
`func (o *ImageVersionResponseModel) UnsetImageStatus()`

UnsetImageStatus ensures that no value is present for ImageStatus, not even an explicit nil
### GetError

`func (o *ImageVersionResponseModel) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ImageVersionResponseModel) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ImageVersionResponseModel) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ImageVersionResponseModel) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ImageVersionResponseModel) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ImageVersionResponseModel) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetAdditionalData

`func (o *ImageVersionResponseModel) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *ImageVersionResponseModel) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *ImageVersionResponseModel) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *ImageVersionResponseModel) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### SetAdditionalDataNil

`func (o *ImageVersionResponseModel) SetAdditionalDataNil(b bool)`

 SetAdditionalDataNil sets the value for AdditionalData to be an explicit nil

### UnsetAdditionalData
`func (o *ImageVersionResponseModel) UnsetAdditionalData()`

UnsetAdditionalData ensures that no value is present for AdditionalData, not even an explicit nil
### GetImageDefinition

`func (o *ImageVersionResponseModel) GetImageDefinition() RefResponseModel`

GetImageDefinition returns the ImageDefinition field if non-nil, zero value otherwise.

### GetImageDefinitionOk

`func (o *ImageVersionResponseModel) GetImageDefinitionOk() (*RefResponseModel, bool)`

GetImageDefinitionOk returns a tuple with the ImageDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDefinition

`func (o *ImageVersionResponseModel) SetImageDefinition(v RefResponseModel)`

SetImageDefinition sets ImageDefinition field to given value.

### HasImageDefinition

`func (o *ImageVersionResponseModel) HasImageDefinition() bool`

HasImageDefinition returns a boolean if a field has been set.

### GetImageScheme

`func (o *ImageVersionResponseModel) GetImageScheme() ImageSchemeResponseModel`

GetImageScheme returns the ImageScheme field if non-nil, zero value otherwise.

### GetImageSchemeOk

`func (o *ImageVersionResponseModel) GetImageSchemeOk() (*ImageSchemeResponseModel, bool)`

GetImageSchemeOk returns a tuple with the ImageScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageScheme

`func (o *ImageVersionResponseModel) SetImageScheme(v ImageSchemeResponseModel)`

SetImageScheme sets ImageScheme field to given value.

### HasImageScheme

`func (o *ImageVersionResponseModel) HasImageScheme() bool`

HasImageScheme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


