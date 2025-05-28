# UpdateTemplateImageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewName** | Pointer to **string** | Updated name of the template | [optional] 
**NewNotes** | Pointer to **string** | Customer notes about template image | [optional] 
**UpdateAllowedIPs** | Pointer to **bool** | Ip Addresses allowed to RDP flag | [optional] 
**AllowedIPs** | Pointer to **[]string** | Ip Addresses allowed to RDP | [optional] 

## Methods

### NewUpdateTemplateImageModel

`func NewUpdateTemplateImageModel() *UpdateTemplateImageModel`

NewUpdateTemplateImageModel instantiates a new UpdateTemplateImageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTemplateImageModelWithDefaults

`func NewUpdateTemplateImageModelWithDefaults() *UpdateTemplateImageModel`

NewUpdateTemplateImageModelWithDefaults instantiates a new UpdateTemplateImageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewName

`func (o *UpdateTemplateImageModel) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateTemplateImageModel) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateTemplateImageModel) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateTemplateImageModel) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetNewNotes

`func (o *UpdateTemplateImageModel) GetNewNotes() string`

GetNewNotes returns the NewNotes field if non-nil, zero value otherwise.

### GetNewNotesOk

`func (o *UpdateTemplateImageModel) GetNewNotesOk() (*string, bool)`

GetNewNotesOk returns a tuple with the NewNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewNotes

`func (o *UpdateTemplateImageModel) SetNewNotes(v string)`

SetNewNotes sets NewNotes field to given value.

### HasNewNotes

`func (o *UpdateTemplateImageModel) HasNewNotes() bool`

HasNewNotes returns a boolean if a field has been set.

### GetUpdateAllowedIPs

`func (o *UpdateTemplateImageModel) GetUpdateAllowedIPs() bool`

GetUpdateAllowedIPs returns the UpdateAllowedIPs field if non-nil, zero value otherwise.

### GetUpdateAllowedIPsOk

`func (o *UpdateTemplateImageModel) GetUpdateAllowedIPsOk() (*bool, bool)`

GetUpdateAllowedIPsOk returns a tuple with the UpdateAllowedIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAllowedIPs

`func (o *UpdateTemplateImageModel) SetUpdateAllowedIPs(v bool)`

SetUpdateAllowedIPs sets UpdateAllowedIPs field to given value.

### HasUpdateAllowedIPs

`func (o *UpdateTemplateImageModel) HasUpdateAllowedIPs() bool`

HasUpdateAllowedIPs returns a boolean if a field has been set.

### GetAllowedIPs

`func (o *UpdateTemplateImageModel) GetAllowedIPs() []string`

GetAllowedIPs returns the AllowedIPs field if non-nil, zero value otherwise.

### GetAllowedIPsOk

`func (o *UpdateTemplateImageModel) GetAllowedIPsOk() (*[]string, bool)`

GetAllowedIPsOk returns a tuple with the AllowedIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIPs

`func (o *UpdateTemplateImageModel) SetAllowedIPs(v []string)`

SetAllowedIPs sets AllowedIPs field to given value.

### HasAllowedIPs

`func (o *UpdateTemplateImageModel) HasAllowedIPs() bool`

HasAllowedIPs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


