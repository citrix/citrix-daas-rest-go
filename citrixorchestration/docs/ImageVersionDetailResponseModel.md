# ImageVersionDetailResponseModel

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
**ImageDefinition** | Pointer to [**ImageDefinitionResponseModel**](ImageDefinitionResponseModel.md) |  | [optional] 
**ImageScheme** | Pointer to [**ImageSchemeResponseModel**](ImageSchemeResponseModel.md) |  | [optional] 
**VMMetadata** | Pointer to [**ProvisioningSchemeVmMetadataResponseModel**](ProvisioningSchemeVmMetadataResponseModel.md) |  | [optional] 
**MasterImage** | [**HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) |  | 

## Methods

### NewImageVersionDetailResponseModel

`func NewImageVersionDetailResponseModel(masterImage HypervisorResourceRefResponseModel, ) *ImageVersionDetailResponseModel`

NewImageVersionDetailResponseModel instantiates a new ImageVersionDetailResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageVersionDetailResponseModelWithDefaults

`func NewImageVersionDetailResponseModelWithDefaults() *ImageVersionDetailResponseModel`

NewImageVersionDetailResponseModelWithDefaults instantiates a new ImageVersionDetailResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ImageVersionDetailResponseModel) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ImageVersionDetailResponseModel) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ImageVersionDetailResponseModel) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ImageVersionDetailResponseModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ImageVersionDetailResponseModel) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ImageVersionDetailResponseModel) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetId

`func (o *ImageVersionDetailResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageVersionDetailResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageVersionDetailResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImageVersionDetailResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMasterImagePath

`func (o *ImageVersionDetailResponseModel) GetMasterImagePath() string`

GetMasterImagePath returns the MasterImagePath field if non-nil, zero value otherwise.

### GetMasterImagePathOk

`func (o *ImageVersionDetailResponseModel) GetMasterImagePathOk() (*string, bool)`

GetMasterImagePathOk returns a tuple with the MasterImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterImagePath

`func (o *ImageVersionDetailResponseModel) SetMasterImagePath(v string)`

SetMasterImagePath sets MasterImagePath field to given value.

### HasMasterImagePath

`func (o *ImageVersionDetailResponseModel) HasMasterImagePath() bool`

HasMasterImagePath returns a boolean if a field has been set.

### SetMasterImagePathNil

`func (o *ImageVersionDetailResponseModel) SetMasterImagePathNil(b bool)`

 SetMasterImagePathNil sets the value for MasterImagePath to be an explicit nil

### UnsetMasterImagePath
`func (o *ImageVersionDetailResponseModel) UnsetMasterImagePath()`

UnsetMasterImagePath ensures that no value is present for MasterImagePath, not even an explicit nil
### GetDescription

`func (o *ImageVersionDetailResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImageVersionDetailResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImageVersionDetailResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImageVersionDetailResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ImageVersionDetailResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ImageVersionDetailResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDiskSizeGB

`func (o *ImageVersionDetailResponseModel) GetDiskSizeGB() int32`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *ImageVersionDetailResponseModel) GetDiskSizeGBOk() (*int32, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *ImageVersionDetailResponseModel) SetDiskSizeGB(v int32)`

SetDiskSizeGB sets DiskSizeGB field to given value.

### HasDiskSizeGB

`func (o *ImageVersionDetailResponseModel) HasDiskSizeGB() bool`

HasDiskSizeGB returns a boolean if a field has been set.

### GetCreateTime

`func (o *ImageVersionDetailResponseModel) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ImageVersionDetailResponseModel) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ImageVersionDetailResponseModel) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ImageVersionDetailResponseModel) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### SetCreateTimeNil

`func (o *ImageVersionDetailResponseModel) SetCreateTimeNil(b bool)`

 SetCreateTimeNil sets the value for CreateTime to be an explicit nil

### UnsetCreateTime
`func (o *ImageVersionDetailResponseModel) UnsetCreateTime()`

UnsetCreateTime ensures that no value is present for CreateTime, not even an explicit nil
### GetWriteBackCacheDiskSize

`func (o *ImageVersionDetailResponseModel) GetWriteBackCacheDiskSize() int32`

GetWriteBackCacheDiskSize returns the WriteBackCacheDiskSize field if non-nil, zero value otherwise.

### GetWriteBackCacheDiskSizeOk

`func (o *ImageVersionDetailResponseModel) GetWriteBackCacheDiskSizeOk() (*int32, bool)`

GetWriteBackCacheDiskSizeOk returns a tuple with the WriteBackCacheDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheDiskSize

`func (o *ImageVersionDetailResponseModel) SetWriteBackCacheDiskSize(v int32)`

SetWriteBackCacheDiskSize sets WriteBackCacheDiskSize field to given value.

### HasWriteBackCacheDiskSize

`func (o *ImageVersionDetailResponseModel) HasWriteBackCacheDiskSize() bool`

HasWriteBackCacheDiskSize returns a boolean if a field has been set.

### GetWriteBackCacheMemorySize

`func (o *ImageVersionDetailResponseModel) GetWriteBackCacheMemorySize() int32`

GetWriteBackCacheMemorySize returns the WriteBackCacheMemorySize field if non-nil, zero value otherwise.

### GetWriteBackCacheMemorySizeOk

`func (o *ImageVersionDetailResponseModel) GetWriteBackCacheMemorySizeOk() (*int32, bool)`

GetWriteBackCacheMemorySizeOk returns a tuple with the WriteBackCacheMemorySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheMemorySize

`func (o *ImageVersionDetailResponseModel) SetWriteBackCacheMemorySize(v int32)`

SetWriteBackCacheMemorySize sets WriteBackCacheMemorySize field to given value.

### HasWriteBackCacheMemorySize

`func (o *ImageVersionDetailResponseModel) HasWriteBackCacheMemorySize() bool`

HasWriteBackCacheMemorySize returns a boolean if a field has been set.

### GetImageStatus

`func (o *ImageVersionDetailResponseModel) GetImageStatus() string`

GetImageStatus returns the ImageStatus field if non-nil, zero value otherwise.

### GetImageStatusOk

`func (o *ImageVersionDetailResponseModel) GetImageStatusOk() (*string, bool)`

GetImageStatusOk returns a tuple with the ImageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStatus

`func (o *ImageVersionDetailResponseModel) SetImageStatus(v string)`

SetImageStatus sets ImageStatus field to given value.

### HasImageStatus

`func (o *ImageVersionDetailResponseModel) HasImageStatus() bool`

HasImageStatus returns a boolean if a field has been set.

### SetImageStatusNil

`func (o *ImageVersionDetailResponseModel) SetImageStatusNil(b bool)`

 SetImageStatusNil sets the value for ImageStatus to be an explicit nil

### UnsetImageStatus
`func (o *ImageVersionDetailResponseModel) UnsetImageStatus()`

UnsetImageStatus ensures that no value is present for ImageStatus, not even an explicit nil
### GetError

`func (o *ImageVersionDetailResponseModel) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ImageVersionDetailResponseModel) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ImageVersionDetailResponseModel) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ImageVersionDetailResponseModel) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ImageVersionDetailResponseModel) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ImageVersionDetailResponseModel) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetAdditionalData

`func (o *ImageVersionDetailResponseModel) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *ImageVersionDetailResponseModel) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *ImageVersionDetailResponseModel) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *ImageVersionDetailResponseModel) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### SetAdditionalDataNil

`func (o *ImageVersionDetailResponseModel) SetAdditionalDataNil(b bool)`

 SetAdditionalDataNil sets the value for AdditionalData to be an explicit nil

### UnsetAdditionalData
`func (o *ImageVersionDetailResponseModel) UnsetAdditionalData()`

UnsetAdditionalData ensures that no value is present for AdditionalData, not even an explicit nil
### GetImageDefinition

`func (o *ImageVersionDetailResponseModel) GetImageDefinition() ImageDefinitionResponseModel`

GetImageDefinition returns the ImageDefinition field if non-nil, zero value otherwise.

### GetImageDefinitionOk

`func (o *ImageVersionDetailResponseModel) GetImageDefinitionOk() (*ImageDefinitionResponseModel, bool)`

GetImageDefinitionOk returns a tuple with the ImageDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDefinition

`func (o *ImageVersionDetailResponseModel) SetImageDefinition(v ImageDefinitionResponseModel)`

SetImageDefinition sets ImageDefinition field to given value.

### HasImageDefinition

`func (o *ImageVersionDetailResponseModel) HasImageDefinition() bool`

HasImageDefinition returns a boolean if a field has been set.

### GetImageScheme

`func (o *ImageVersionDetailResponseModel) GetImageScheme() ImageSchemeResponseModel`

GetImageScheme returns the ImageScheme field if non-nil, zero value otherwise.

### GetImageSchemeOk

`func (o *ImageVersionDetailResponseModel) GetImageSchemeOk() (*ImageSchemeResponseModel, bool)`

GetImageSchemeOk returns a tuple with the ImageScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageScheme

`func (o *ImageVersionDetailResponseModel) SetImageScheme(v ImageSchemeResponseModel)`

SetImageScheme sets ImageScheme field to given value.

### HasImageScheme

`func (o *ImageVersionDetailResponseModel) HasImageScheme() bool`

HasImageScheme returns a boolean if a field has been set.

### GetVMMetadata

`func (o *ImageVersionDetailResponseModel) GetVMMetadata() ProvisioningSchemeVmMetadataResponseModel`

GetVMMetadata returns the VMMetadata field if non-nil, zero value otherwise.

### GetVMMetadataOk

`func (o *ImageVersionDetailResponseModel) GetVMMetadataOk() (*ProvisioningSchemeVmMetadataResponseModel, bool)`

GetVMMetadataOk returns a tuple with the VMMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMMetadata

`func (o *ImageVersionDetailResponseModel) SetVMMetadata(v ProvisioningSchemeVmMetadataResponseModel)`

SetVMMetadata sets VMMetadata field to given value.

### HasVMMetadata

`func (o *ImageVersionDetailResponseModel) HasVMMetadata() bool`

HasVMMetadata returns a boolean if a field has been set.

### GetMasterImage

`func (o *ImageVersionDetailResponseModel) GetMasterImage() HypervisorResourceRefResponseModel`

GetMasterImage returns the MasterImage field if non-nil, zero value otherwise.

### GetMasterImageOk

`func (o *ImageVersionDetailResponseModel) GetMasterImageOk() (*HypervisorResourceRefResponseModel, bool)`

GetMasterImageOk returns a tuple with the MasterImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterImage

`func (o *ImageVersionDetailResponseModel) SetMasterImage(v HypervisorResourceRefResponseModel)`

SetMasterImage sets MasterImage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


