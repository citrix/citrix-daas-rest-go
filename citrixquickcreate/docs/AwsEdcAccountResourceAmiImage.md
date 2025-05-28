# AwsEdcAccountResourceAmiImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageId** | Pointer to **NullableString** | Image Id | [optional] 
**Name** | Pointer to **NullableString** | Image Name | [optional] 
**Description** | Pointer to **NullableString** | Image Description | [optional] 
**Status** | Pointer to [**NullableAwsEdcAmiImageStatus**](AwsEdcAmiImageStatus.md) | Image Status  Enum values AVAILABLE, DEREGISTERED, ERROR, FAILED, INVALID, PENDING, TRANSIENT | [optional] 
**Location** | Pointer to **NullableString** | Image Location | [optional] 
**ImageType** | Pointer to [**NullableAwsEdcAmiImageType**](AwsEdcAmiImageType.md) | Image Type  Enum values KERNEL, MACHINE, RAMDISK | [optional] 
**PlatformType** | Pointer to [**NullableAwsEdcAmiImagePlatform**](AwsEdcAmiImagePlatform.md) | Vpc CIDR Block  Enum values WINDOWS | [optional] 
**BootMode** | Pointer to [**NullableAwsEdcAmiImageBootMode**](AwsEdcAmiImageBootMode.md) | Image Boot Type  Enum values LEGACYBIOS, UEFI | [optional] 
**ArchitectureType** | Pointer to [**NullableAwsEdcAmiImageArchitecture**](AwsEdcAmiImageArchitecture.md) | Image Architecture  Enum values ARM64, I386, X86_64 | [optional] 
**VirtualizationType** | Pointer to [**NullableAwsEdcAmiImageVirtualization**](AwsEdcAmiImageVirtualization.md) | Image Virtualization  Enum values HVM, PARAVIRTUAL | [optional] 
**Public** | Pointer to **NullableBool** | Image IsPublic | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Image Tags | [optional] 

## Methods

### NewAwsEdcAccountResourceAmiImage

`func NewAwsEdcAccountResourceAmiImage() *AwsEdcAccountResourceAmiImage`

NewAwsEdcAccountResourceAmiImage instantiates a new AwsEdcAccountResourceAmiImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsEdcAccountResourceAmiImageWithDefaults

`func NewAwsEdcAccountResourceAmiImageWithDefaults() *AwsEdcAccountResourceAmiImage`

NewAwsEdcAccountResourceAmiImageWithDefaults instantiates a new AwsEdcAccountResourceAmiImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageId

`func (o *AwsEdcAccountResourceAmiImage) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *AwsEdcAccountResourceAmiImage) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *AwsEdcAccountResourceAmiImage) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *AwsEdcAccountResourceAmiImage) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### SetImageIdNil

`func (o *AwsEdcAccountResourceAmiImage) SetImageIdNil(b bool)`

 SetImageIdNil sets the value for ImageId to be an explicit nil

### UnsetImageId
`func (o *AwsEdcAccountResourceAmiImage) UnsetImageId()`

UnsetImageId ensures that no value is present for ImageId, not even an explicit nil
### GetName

`func (o *AwsEdcAccountResourceAmiImage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsEdcAccountResourceAmiImage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsEdcAccountResourceAmiImage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AwsEdcAccountResourceAmiImage) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AwsEdcAccountResourceAmiImage) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AwsEdcAccountResourceAmiImage) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *AwsEdcAccountResourceAmiImage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AwsEdcAccountResourceAmiImage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AwsEdcAccountResourceAmiImage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AwsEdcAccountResourceAmiImage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AwsEdcAccountResourceAmiImage) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AwsEdcAccountResourceAmiImage) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStatus

`func (o *AwsEdcAccountResourceAmiImage) GetStatus() AwsEdcAmiImageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsEdcAccountResourceAmiImage) GetStatusOk() (*AwsEdcAmiImageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsEdcAccountResourceAmiImage) SetStatus(v AwsEdcAmiImageStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsEdcAccountResourceAmiImage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AwsEdcAccountResourceAmiImage) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AwsEdcAccountResourceAmiImage) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetLocation

`func (o *AwsEdcAccountResourceAmiImage) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AwsEdcAccountResourceAmiImage) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AwsEdcAccountResourceAmiImage) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AwsEdcAccountResourceAmiImage) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *AwsEdcAccountResourceAmiImage) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *AwsEdcAccountResourceAmiImage) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetImageType

`func (o *AwsEdcAccountResourceAmiImage) GetImageType() AwsEdcAmiImageType`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *AwsEdcAccountResourceAmiImage) GetImageTypeOk() (*AwsEdcAmiImageType, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *AwsEdcAccountResourceAmiImage) SetImageType(v AwsEdcAmiImageType)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *AwsEdcAccountResourceAmiImage) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### SetImageTypeNil

`func (o *AwsEdcAccountResourceAmiImage) SetImageTypeNil(b bool)`

 SetImageTypeNil sets the value for ImageType to be an explicit nil

### UnsetImageType
`func (o *AwsEdcAccountResourceAmiImage) UnsetImageType()`

UnsetImageType ensures that no value is present for ImageType, not even an explicit nil
### GetPlatformType

`func (o *AwsEdcAccountResourceAmiImage) GetPlatformType() AwsEdcAmiImagePlatform`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *AwsEdcAccountResourceAmiImage) GetPlatformTypeOk() (*AwsEdcAmiImagePlatform, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *AwsEdcAccountResourceAmiImage) SetPlatformType(v AwsEdcAmiImagePlatform)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *AwsEdcAccountResourceAmiImage) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### SetPlatformTypeNil

`func (o *AwsEdcAccountResourceAmiImage) SetPlatformTypeNil(b bool)`

 SetPlatformTypeNil sets the value for PlatformType to be an explicit nil

### UnsetPlatformType
`func (o *AwsEdcAccountResourceAmiImage) UnsetPlatformType()`

UnsetPlatformType ensures that no value is present for PlatformType, not even an explicit nil
### GetBootMode

`func (o *AwsEdcAccountResourceAmiImage) GetBootMode() AwsEdcAmiImageBootMode`

GetBootMode returns the BootMode field if non-nil, zero value otherwise.

### GetBootModeOk

`func (o *AwsEdcAccountResourceAmiImage) GetBootModeOk() (*AwsEdcAmiImageBootMode, bool)`

GetBootModeOk returns a tuple with the BootMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootMode

`func (o *AwsEdcAccountResourceAmiImage) SetBootMode(v AwsEdcAmiImageBootMode)`

SetBootMode sets BootMode field to given value.

### HasBootMode

`func (o *AwsEdcAccountResourceAmiImage) HasBootMode() bool`

HasBootMode returns a boolean if a field has been set.

### SetBootModeNil

`func (o *AwsEdcAccountResourceAmiImage) SetBootModeNil(b bool)`

 SetBootModeNil sets the value for BootMode to be an explicit nil

### UnsetBootMode
`func (o *AwsEdcAccountResourceAmiImage) UnsetBootMode()`

UnsetBootMode ensures that no value is present for BootMode, not even an explicit nil
### GetArchitectureType

`func (o *AwsEdcAccountResourceAmiImage) GetArchitectureType() AwsEdcAmiImageArchitecture`

GetArchitectureType returns the ArchitectureType field if non-nil, zero value otherwise.

### GetArchitectureTypeOk

`func (o *AwsEdcAccountResourceAmiImage) GetArchitectureTypeOk() (*AwsEdcAmiImageArchitecture, bool)`

GetArchitectureTypeOk returns a tuple with the ArchitectureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitectureType

`func (o *AwsEdcAccountResourceAmiImage) SetArchitectureType(v AwsEdcAmiImageArchitecture)`

SetArchitectureType sets ArchitectureType field to given value.

### HasArchitectureType

`func (o *AwsEdcAccountResourceAmiImage) HasArchitectureType() bool`

HasArchitectureType returns a boolean if a field has been set.

### SetArchitectureTypeNil

`func (o *AwsEdcAccountResourceAmiImage) SetArchitectureTypeNil(b bool)`

 SetArchitectureTypeNil sets the value for ArchitectureType to be an explicit nil

### UnsetArchitectureType
`func (o *AwsEdcAccountResourceAmiImage) UnsetArchitectureType()`

UnsetArchitectureType ensures that no value is present for ArchitectureType, not even an explicit nil
### GetVirtualizationType

`func (o *AwsEdcAccountResourceAmiImage) GetVirtualizationType() AwsEdcAmiImageVirtualization`

GetVirtualizationType returns the VirtualizationType field if non-nil, zero value otherwise.

### GetVirtualizationTypeOk

`func (o *AwsEdcAccountResourceAmiImage) GetVirtualizationTypeOk() (*AwsEdcAmiImageVirtualization, bool)`

GetVirtualizationTypeOk returns a tuple with the VirtualizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationType

`func (o *AwsEdcAccountResourceAmiImage) SetVirtualizationType(v AwsEdcAmiImageVirtualization)`

SetVirtualizationType sets VirtualizationType field to given value.

### HasVirtualizationType

`func (o *AwsEdcAccountResourceAmiImage) HasVirtualizationType() bool`

HasVirtualizationType returns a boolean if a field has been set.

### SetVirtualizationTypeNil

`func (o *AwsEdcAccountResourceAmiImage) SetVirtualizationTypeNil(b bool)`

 SetVirtualizationTypeNil sets the value for VirtualizationType to be an explicit nil

### UnsetVirtualizationType
`func (o *AwsEdcAccountResourceAmiImage) UnsetVirtualizationType()`

UnsetVirtualizationType ensures that no value is present for VirtualizationType, not even an explicit nil
### GetPublic

`func (o *AwsEdcAccountResourceAmiImage) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *AwsEdcAccountResourceAmiImage) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *AwsEdcAccountResourceAmiImage) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *AwsEdcAccountResourceAmiImage) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### SetPublicNil

`func (o *AwsEdcAccountResourceAmiImage) SetPublicNil(b bool)`

 SetPublicNil sets the value for Public to be an explicit nil

### UnsetPublic
`func (o *AwsEdcAccountResourceAmiImage) UnsetPublic()`

UnsetPublic ensures that no value is present for Public, not even an explicit nil
### GetTags

`func (o *AwsEdcAccountResourceAmiImage) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AwsEdcAccountResourceAmiImage) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AwsEdcAccountResourceAmiImage) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AwsEdcAccountResourceAmiImage) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AwsEdcAccountResourceAmiImage) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AwsEdcAccountResourceAmiImage) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


