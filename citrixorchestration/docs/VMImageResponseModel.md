# VMImageResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionalLevel** | Pointer to [**FunctionalLevel**](FunctionalLevel.md) |  | [optional] 
**Image** | Pointer to [**HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) |  | [optional] 
**ImageStatus** | [**VMImageStatus**](VMImageStatus.md) |  | 
**Date** | **string** | The date and time when the snapshot was used in the provisioning scheme. | 
**MasterImageNote** | Pointer to **NullableString** | The note of the provisioning scheme image. | [optional] 

## Methods

### NewVMImageResponseModel

`func NewVMImageResponseModel(imageStatus VMImageStatus, date string, ) *VMImageResponseModel`

NewVMImageResponseModel instantiates a new VMImageResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMImageResponseModelWithDefaults

`func NewVMImageResponseModelWithDefaults() *VMImageResponseModel`

NewVMImageResponseModelWithDefaults instantiates a new VMImageResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionalLevel

`func (o *VMImageResponseModel) GetFunctionalLevel() FunctionalLevel`

GetFunctionalLevel returns the FunctionalLevel field if non-nil, zero value otherwise.

### GetFunctionalLevelOk

`func (o *VMImageResponseModel) GetFunctionalLevelOk() (*FunctionalLevel, bool)`

GetFunctionalLevelOk returns a tuple with the FunctionalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionalLevel

`func (o *VMImageResponseModel) SetFunctionalLevel(v FunctionalLevel)`

SetFunctionalLevel sets FunctionalLevel field to given value.

### HasFunctionalLevel

`func (o *VMImageResponseModel) HasFunctionalLevel() bool`

HasFunctionalLevel returns a boolean if a field has been set.

### GetImage

`func (o *VMImageResponseModel) GetImage() HypervisorResourceRefResponseModel`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *VMImageResponseModel) GetImageOk() (*HypervisorResourceRefResponseModel, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *VMImageResponseModel) SetImage(v HypervisorResourceRefResponseModel)`

SetImage sets Image field to given value.

### HasImage

`func (o *VMImageResponseModel) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetImageStatus

`func (o *VMImageResponseModel) GetImageStatus() VMImageStatus`

GetImageStatus returns the ImageStatus field if non-nil, zero value otherwise.

### GetImageStatusOk

`func (o *VMImageResponseModel) GetImageStatusOk() (*VMImageStatus, bool)`

GetImageStatusOk returns a tuple with the ImageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStatus

`func (o *VMImageResponseModel) SetImageStatus(v VMImageStatus)`

SetImageStatus sets ImageStatus field to given value.


### GetDate

`func (o *VMImageResponseModel) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *VMImageResponseModel) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *VMImageResponseModel) SetDate(v string)`

SetDate sets Date field to given value.


### GetMasterImageNote

`func (o *VMImageResponseModel) GetMasterImageNote() string`

GetMasterImageNote returns the MasterImageNote field if non-nil, zero value otherwise.

### GetMasterImageNoteOk

`func (o *VMImageResponseModel) GetMasterImageNoteOk() (*string, bool)`

GetMasterImageNoteOk returns a tuple with the MasterImageNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterImageNote

`func (o *VMImageResponseModel) SetMasterImageNote(v string)`

SetMasterImageNote sets MasterImageNote field to given value.

### HasMasterImageNote

`func (o *VMImageResponseModel) HasMasterImageNote() bool`

HasMasterImageNote returns a boolean if a field has been set.

### SetMasterImageNoteNil

`func (o *VMImageResponseModel) SetMasterImageNoteNil(b bool)`

 SetMasterImageNoteNil sets the value for MasterImageNote to be an explicit nil

### UnsetMasterImageNote
`func (o *VMImageResponseModel) UnsetMasterImageNote()`

UnsetMasterImageNote ensures that no value is present for MasterImageNote, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


