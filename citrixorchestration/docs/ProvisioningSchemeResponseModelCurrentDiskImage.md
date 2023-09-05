# ProvisioningSchemeResponseModelCurrentDiskImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionalLevel** | Pointer to [**FunctionalLevel**](FunctionalLevel.md) |  | [optional] 
**Image** | Pointer to [**VMImageResponseModelImage**](VMImageResponseModelImage.md) |  | [optional] 
**ImageVersion** | Pointer to [**VMImageResponseModelImageVersion**](VMImageResponseModelImageVersion.md) |  | [optional] 
**ImageStatus** | [**VMImageStatus**](VMImageStatus.md) |  | 
**Date** | **string** | The date and time when the snapshot was used in the provisioning scheme. | 
**MasterImageNote** | Pointer to **string** | The note of the provisioning scheme image. | [optional] 

## Methods

### NewProvisioningSchemeResponseModelCurrentDiskImage

`func NewProvisioningSchemeResponseModelCurrentDiskImage(imageStatus VMImageStatus, date string, ) *ProvisioningSchemeResponseModelCurrentDiskImage`

NewProvisioningSchemeResponseModelCurrentDiskImage instantiates a new ProvisioningSchemeResponseModelCurrentDiskImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningSchemeResponseModelCurrentDiskImageWithDefaults

`func NewProvisioningSchemeResponseModelCurrentDiskImageWithDefaults() *ProvisioningSchemeResponseModelCurrentDiskImage`

NewProvisioningSchemeResponseModelCurrentDiskImageWithDefaults instantiates a new ProvisioningSchemeResponseModelCurrentDiskImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionalLevel

`func (o *ProvisioningSchemeResponseModelCurrentDiskImage) GetFunctionalLevel() FunctionalLevel`

GetFunctionalLevel returns the FunctionalLevel field if non-nil, zero value otherwise.

### GetFunctionalLevelOk

`func (o *ProvisioningSchemeResponseModelCurrentDiskImage) GetFunctionalLevelOk() (*FunctionalLevel, bool)`

GetFunctionalLevelOk returns a tuple with the FunctionalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionalLevel

`func (o *ProvisioningSchemeResponseModelCurrentDiskImage) SetFunctionalLevel(v FunctionalLevel)`

SetFunctionalLevel sets FunctionalLevel field to given value.

### HasFunctionalLevel

`func (o *ProvisioningSchemeResponseModelCurrentDiskImage) HasFunctionalLevel() bool`

HasFunctionalLevel returns a boolean if a field has been set.

### GetImage

`func (o *ProvisioningSchemeResponseModelCurrentDiskImage) GetImage() VMImageResponseModelImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ProvisioningSchemeResponseModelCurrentDiskImage) GetImageOk() (*VMImageResponseModelImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ProvisioningSchemeResponseModelCurrentDiskImage) SetImage(v VMImageResponseModelImage)`

SetImage sets Image field to given value.

### HasImage

`func (o *ProvisioningSchemeResponseModelCurrentDiskImage) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetImageVersion

`func (o *ProvisioningSchemeResponseModelCurrentDiskImage) GetImageVersion() VMImageResponseModelImageVersion`

GetImageVersion returns the ImageVersion field if non-nil, zero value otherwise.

### GetImageVersionOk

`func (o *ProvisioningSchemeResponseModelCurrentDiskImage) GetImageVersionOk() (*VMImageResponseModelImageVersion, bool)`

GetImageVersionOk returns a tuple with the ImageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersion

`func (o *ProvisioningSchemeResponseModelCurrentDiskImage) SetImageVersion(v VMImageResponseModelImageVersion)`

SetImageVersion sets ImageVersion field to given value.

### HasImageVersion

`func (o *ProvisioningSchemeResponseModelCurrentDiskImage) HasImageVersion() bool`

HasImageVersion returns a boolean if a field has been set.

### GetImageStatus

`func (o *ProvisioningSchemeResponseModelCurrentDiskImage) GetImageStatus() VMImageStatus`

GetImageStatus returns the ImageStatus field if non-nil, zero value otherwise.

### GetImageStatusOk

`func (o *ProvisioningSchemeResponseModelCurrentDiskImage) GetImageStatusOk() (*VMImageStatus, bool)`

GetImageStatusOk returns a tuple with the ImageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStatus

`func (o *ProvisioningSchemeResponseModelCurrentDiskImage) SetImageStatus(v VMImageStatus)`

SetImageStatus sets ImageStatus field to given value.


### GetDate

`func (o *ProvisioningSchemeResponseModelCurrentDiskImage) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ProvisioningSchemeResponseModelCurrentDiskImage) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ProvisioningSchemeResponseModelCurrentDiskImage) SetDate(v string)`

SetDate sets Date field to given value.


### GetMasterImageNote

`func (o *ProvisioningSchemeResponseModelCurrentDiskImage) GetMasterImageNote() string`

GetMasterImageNote returns the MasterImageNote field if non-nil, zero value otherwise.

### GetMasterImageNoteOk

`func (o *ProvisioningSchemeResponseModelCurrentDiskImage) GetMasterImageNoteOk() (*string, bool)`

GetMasterImageNoteOk returns a tuple with the MasterImageNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterImageNote

`func (o *ProvisioningSchemeResponseModelCurrentDiskImage) SetMasterImageNote(v string)`

SetMasterImageNote sets MasterImageNote field to given value.

### HasMasterImageNote

`func (o *ProvisioningSchemeResponseModelCurrentDiskImage) HasMasterImageNote() bool`

HasMasterImageNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


