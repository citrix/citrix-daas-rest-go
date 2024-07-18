# MachineIdSaveAsImageBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Id for the image builder machine | 
**ImageName** | **string** | Image Name | 
**ImageDescription** | **string** | Image Description | 
**ImageNotes** | Pointer to **NullableString** | Image Notes | [optional] 

## Methods

### NewMachineIdSaveAsImageBody

`func NewMachineIdSaveAsImageBody(accountId string, imageName string, imageDescription string, ) *MachineIdSaveAsImageBody`

NewMachineIdSaveAsImageBody instantiates a new MachineIdSaveAsImageBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineIdSaveAsImageBodyWithDefaults

`func NewMachineIdSaveAsImageBodyWithDefaults() *MachineIdSaveAsImageBody`

NewMachineIdSaveAsImageBodyWithDefaults instantiates a new MachineIdSaveAsImageBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *MachineIdSaveAsImageBody) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *MachineIdSaveAsImageBody) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *MachineIdSaveAsImageBody) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetImageName

`func (o *MachineIdSaveAsImageBody) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *MachineIdSaveAsImageBody) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *MachineIdSaveAsImageBody) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetImageDescription

`func (o *MachineIdSaveAsImageBody) GetImageDescription() string`

GetImageDescription returns the ImageDescription field if non-nil, zero value otherwise.

### GetImageDescriptionOk

`func (o *MachineIdSaveAsImageBody) GetImageDescriptionOk() (*string, bool)`

GetImageDescriptionOk returns a tuple with the ImageDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDescription

`func (o *MachineIdSaveAsImageBody) SetImageDescription(v string)`

SetImageDescription sets ImageDescription field to given value.


### GetImageNotes

`func (o *MachineIdSaveAsImageBody) GetImageNotes() string`

GetImageNotes returns the ImageNotes field if non-nil, zero value otherwise.

### GetImageNotesOk

`func (o *MachineIdSaveAsImageBody) GetImageNotesOk() (*string, bool)`

GetImageNotesOk returns a tuple with the ImageNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageNotes

`func (o *MachineIdSaveAsImageBody) SetImageNotes(v string)`

SetImageNotes sets ImageNotes field to given value.

### HasImageNotes

`func (o *MachineIdSaveAsImageBody) HasImageNotes() bool`

HasImageNotes returns a boolean if a field has been set.

### SetImageNotesNil

`func (o *MachineIdSaveAsImageBody) SetImageNotesNil(b bool)`

 SetImageNotesNil sets the value for ImageNotes to be an explicit nil

### UnsetImageNotes
`func (o *MachineIdSaveAsImageBody) UnsetImageNotes()`

UnsetImageNotes ensures that no value is present for ImageNotes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


