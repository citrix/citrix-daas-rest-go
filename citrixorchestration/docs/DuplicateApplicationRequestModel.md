# DuplicateApplicationRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewName** | Pointer to **NullableString** | Optional name for the duplicated application.  If not specified, the system will choose a new name automatically, based on the original application name. If specified, must be unique within the site. | [optional] 
**NewFolder** | Pointer to **NullableString** | Optional folder in which to create the duplicated application.  If not specified, the new application will be created in the same folder as the original. May be specified as either Path or Id.  If specified as a path, and the target does not exist, it will be automatically created. | [optional] 
**CreateDisabled** | Pointer to **NullableBool** | Optional. If not specified, or specified as &#x60;false&#x60;, the new application will have the same visibility and enabled state as the original application.  This may be undesirable; if the original application is visible, it means the new duplicate will also be immediately visible to end users. Setting this to &#x60;true&#x60; causes the new duplicate to start out invisible and disabled, allowing it to be further modified before making it visible to end users. | [optional] [default to false]

## Methods

### NewDuplicateApplicationRequestModel

`func NewDuplicateApplicationRequestModel() *DuplicateApplicationRequestModel`

NewDuplicateApplicationRequestModel instantiates a new DuplicateApplicationRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDuplicateApplicationRequestModelWithDefaults

`func NewDuplicateApplicationRequestModelWithDefaults() *DuplicateApplicationRequestModel`

NewDuplicateApplicationRequestModelWithDefaults instantiates a new DuplicateApplicationRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewName

`func (o *DuplicateApplicationRequestModel) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *DuplicateApplicationRequestModel) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *DuplicateApplicationRequestModel) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *DuplicateApplicationRequestModel) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### SetNewNameNil

`func (o *DuplicateApplicationRequestModel) SetNewNameNil(b bool)`

 SetNewNameNil sets the value for NewName to be an explicit nil

### UnsetNewName
`func (o *DuplicateApplicationRequestModel) UnsetNewName()`

UnsetNewName ensures that no value is present for NewName, not even an explicit nil
### GetNewFolder

`func (o *DuplicateApplicationRequestModel) GetNewFolder() string`

GetNewFolder returns the NewFolder field if non-nil, zero value otherwise.

### GetNewFolderOk

`func (o *DuplicateApplicationRequestModel) GetNewFolderOk() (*string, bool)`

GetNewFolderOk returns a tuple with the NewFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewFolder

`func (o *DuplicateApplicationRequestModel) SetNewFolder(v string)`

SetNewFolder sets NewFolder field to given value.

### HasNewFolder

`func (o *DuplicateApplicationRequestModel) HasNewFolder() bool`

HasNewFolder returns a boolean if a field has been set.

### SetNewFolderNil

`func (o *DuplicateApplicationRequestModel) SetNewFolderNil(b bool)`

 SetNewFolderNil sets the value for NewFolder to be an explicit nil

### UnsetNewFolder
`func (o *DuplicateApplicationRequestModel) UnsetNewFolder()`

UnsetNewFolder ensures that no value is present for NewFolder, not even an explicit nil
### GetCreateDisabled

`func (o *DuplicateApplicationRequestModel) GetCreateDisabled() bool`

GetCreateDisabled returns the CreateDisabled field if non-nil, zero value otherwise.

### GetCreateDisabledOk

`func (o *DuplicateApplicationRequestModel) GetCreateDisabledOk() (*bool, bool)`

GetCreateDisabledOk returns a tuple with the CreateDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDisabled

`func (o *DuplicateApplicationRequestModel) SetCreateDisabled(v bool)`

SetCreateDisabled sets CreateDisabled field to given value.

### HasCreateDisabled

`func (o *DuplicateApplicationRequestModel) HasCreateDisabled() bool`

HasCreateDisabled returns a boolean if a field has been set.

### SetCreateDisabledNil

`func (o *DuplicateApplicationRequestModel) SetCreateDisabledNil(b bool)`

 SetCreateDisabledNil sets the value for CreateDisabled to be an explicit nil

### UnsetCreateDisabled
`func (o *DuplicateApplicationRequestModel) UnsetCreateDisabled()`

UnsetCreateDisabled ensures that no value is present for CreateDisabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


