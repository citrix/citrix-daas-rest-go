# ImportImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | [**AccountType**](AccountType.md) |  | 
**Name** | **string** | Image Name | 
**Description** | Pointer to **NullableString** | Image Name | [optional] 
**Notes** | Pointer to **NullableString** | Image Name | [optional] 

## Methods

### NewImportImage

`func NewImportImage(accountType AccountType, name string, ) *ImportImage`

NewImportImage instantiates a new ImportImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportImageWithDefaults

`func NewImportImageWithDefaults() *ImportImage`

NewImportImageWithDefaults instantiates a new ImportImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *ImportImage) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *ImportImage) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *ImportImage) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetName

`func (o *ImportImage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImportImage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImportImage) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ImportImage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImportImage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImportImage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImportImage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ImportImage) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ImportImage) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetNotes

`func (o *ImportImage) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ImportImage) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ImportImage) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ImportImage) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *ImportImage) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *ImportImage) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


