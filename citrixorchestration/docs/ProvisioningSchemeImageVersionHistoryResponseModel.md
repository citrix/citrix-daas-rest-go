# ProvisioningSchemeImageVersionHistoryResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageVersion** | [**ImageVersionRefResponseModel**](ImageVersionRefResponseModel.md) |  | 
**Date** | Pointer to **NullableString** | The date and time when the snapshot was used in the provisioning scheme. | [optional] 
**ImageAssignmentNote** | Pointer to **NullableString** | The image assignment note when assigning the image version to the provisioning scheme. | [optional] 
**IsImageAvailable** | Pointer to **bool** | A flag indicating if the image version is still available for this provisioning scheme. | [optional] 

## Methods

### NewProvisioningSchemeImageVersionHistoryResponseModel

`func NewProvisioningSchemeImageVersionHistoryResponseModel(imageVersion ImageVersionRefResponseModel, ) *ProvisioningSchemeImageVersionHistoryResponseModel`

NewProvisioningSchemeImageVersionHistoryResponseModel instantiates a new ProvisioningSchemeImageVersionHistoryResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningSchemeImageVersionHistoryResponseModelWithDefaults

`func NewProvisioningSchemeImageVersionHistoryResponseModelWithDefaults() *ProvisioningSchemeImageVersionHistoryResponseModel`

NewProvisioningSchemeImageVersionHistoryResponseModelWithDefaults instantiates a new ProvisioningSchemeImageVersionHistoryResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageVersion

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) GetImageVersion() ImageVersionRefResponseModel`

GetImageVersion returns the ImageVersion field if non-nil, zero value otherwise.

### GetImageVersionOk

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) GetImageVersionOk() (*ImageVersionRefResponseModel, bool)`

GetImageVersionOk returns a tuple with the ImageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersion

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) SetImageVersion(v ImageVersionRefResponseModel)`

SetImageVersion sets ImageVersion field to given value.


### GetDate

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetImageAssignmentNote

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) GetImageAssignmentNote() string`

GetImageAssignmentNote returns the ImageAssignmentNote field if non-nil, zero value otherwise.

### GetImageAssignmentNoteOk

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) GetImageAssignmentNoteOk() (*string, bool)`

GetImageAssignmentNoteOk returns a tuple with the ImageAssignmentNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAssignmentNote

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) SetImageAssignmentNote(v string)`

SetImageAssignmentNote sets ImageAssignmentNote field to given value.

### HasImageAssignmentNote

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) HasImageAssignmentNote() bool`

HasImageAssignmentNote returns a boolean if a field has been set.

### SetImageAssignmentNoteNil

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) SetImageAssignmentNoteNil(b bool)`

 SetImageAssignmentNoteNil sets the value for ImageAssignmentNote to be an explicit nil

### UnsetImageAssignmentNote
`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) UnsetImageAssignmentNote()`

UnsetImageAssignmentNote ensures that no value is present for ImageAssignmentNote, not even an explicit nil
### GetIsImageAvailable

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) GetIsImageAvailable() bool`

GetIsImageAvailable returns the IsImageAvailable field if non-nil, zero value otherwise.

### GetIsImageAvailableOk

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) GetIsImageAvailableOk() (*bool, bool)`

GetIsImageAvailableOk returns a tuple with the IsImageAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImageAvailable

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) SetIsImageAvailable(v bool)`

SetIsImageAvailable sets IsImageAvailable field to given value.

### HasIsImageAvailable

`func (o *ProvisioningSchemeImageVersionHistoryResponseModel) HasIsImageAvailable() bool`

HasIsImageAvailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


