# DatabaseChangeScriptModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | Pointer to **NullableString** | The filename of the text file. | [optional] 
**Content** | Pointer to **NullableString** | The contents of the text file. | [optional] 

## Methods

### NewDatabaseChangeScriptModel

`func NewDatabaseChangeScriptModel() *DatabaseChangeScriptModel`

NewDatabaseChangeScriptModel instantiates a new DatabaseChangeScriptModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseChangeScriptModelWithDefaults

`func NewDatabaseChangeScriptModelWithDefaults() *DatabaseChangeScriptModel`

NewDatabaseChangeScriptModelWithDefaults instantiates a new DatabaseChangeScriptModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *DatabaseChangeScriptModel) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *DatabaseChangeScriptModel) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *DatabaseChangeScriptModel) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *DatabaseChangeScriptModel) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *DatabaseChangeScriptModel) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *DatabaseChangeScriptModel) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetContent

`func (o *DatabaseChangeScriptModel) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DatabaseChangeScriptModel) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DatabaseChangeScriptModel) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *DatabaseChangeScriptModel) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *DatabaseChangeScriptModel) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *DatabaseChangeScriptModel) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


