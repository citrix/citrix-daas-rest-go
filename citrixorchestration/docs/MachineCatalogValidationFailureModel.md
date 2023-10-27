# MachineCatalogValidationFailureModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorId** | Pointer to **NullableString** | Gets or sets the lookup key (ErrorId) for this string. | [optional] 
**Category** | Pointer to **NullableString** | Gets or sets the category or label for the message. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Gets or sets the human-readable message for the message. | [optional] 
**Parameters** | Pointer to **[]string** | Gets or sets the parameters impacted by the validation failure. | [optional] 

## Methods

### NewMachineCatalogValidationFailureModel

`func NewMachineCatalogValidationFailureModel() *MachineCatalogValidationFailureModel`

NewMachineCatalogValidationFailureModel instantiates a new MachineCatalogValidationFailureModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineCatalogValidationFailureModelWithDefaults

`func NewMachineCatalogValidationFailureModelWithDefaults() *MachineCatalogValidationFailureModel`

NewMachineCatalogValidationFailureModelWithDefaults instantiates a new MachineCatalogValidationFailureModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorId

`func (o *MachineCatalogValidationFailureModel) GetErrorId() string`

GetErrorId returns the ErrorId field if non-nil, zero value otherwise.

### GetErrorIdOk

`func (o *MachineCatalogValidationFailureModel) GetErrorIdOk() (*string, bool)`

GetErrorIdOk returns a tuple with the ErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorId

`func (o *MachineCatalogValidationFailureModel) SetErrorId(v string)`

SetErrorId sets ErrorId field to given value.

### HasErrorId

`func (o *MachineCatalogValidationFailureModel) HasErrorId() bool`

HasErrorId returns a boolean if a field has been set.

### SetErrorIdNil

`func (o *MachineCatalogValidationFailureModel) SetErrorIdNil(b bool)`

 SetErrorIdNil sets the value for ErrorId to be an explicit nil

### UnsetErrorId
`func (o *MachineCatalogValidationFailureModel) UnsetErrorId()`

UnsetErrorId ensures that no value is present for ErrorId, not even an explicit nil
### GetCategory

`func (o *MachineCatalogValidationFailureModel) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MachineCatalogValidationFailureModel) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MachineCatalogValidationFailureModel) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MachineCatalogValidationFailureModel) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *MachineCatalogValidationFailureModel) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *MachineCatalogValidationFailureModel) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetErrorMessage

`func (o *MachineCatalogValidationFailureModel) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *MachineCatalogValidationFailureModel) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *MachineCatalogValidationFailureModel) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *MachineCatalogValidationFailureModel) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *MachineCatalogValidationFailureModel) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *MachineCatalogValidationFailureModel) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetParameters

`func (o *MachineCatalogValidationFailureModel) GetParameters() []string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *MachineCatalogValidationFailureModel) GetParametersOk() (*[]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *MachineCatalogValidationFailureModel) SetParameters(v []string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *MachineCatalogValidationFailureModel) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *MachineCatalogValidationFailureModel) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *MachineCatalogValidationFailureModel) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


