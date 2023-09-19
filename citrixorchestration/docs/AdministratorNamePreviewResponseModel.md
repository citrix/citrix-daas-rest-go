# AdministratorNamePreviewResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The requested administrator name for previewing. | [optional] 
**Error** | Pointer to **NullableString** | Localized error message when the request administrator request name is illegal. If the name is fine, the property should be null. | [optional] 

## Methods

### NewAdministratorNamePreviewResponseModel

`func NewAdministratorNamePreviewResponseModel() *AdministratorNamePreviewResponseModel`

NewAdministratorNamePreviewResponseModel instantiates a new AdministratorNamePreviewResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdministratorNamePreviewResponseModelWithDefaults

`func NewAdministratorNamePreviewResponseModelWithDefaults() *AdministratorNamePreviewResponseModel`

NewAdministratorNamePreviewResponseModelWithDefaults instantiates a new AdministratorNamePreviewResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AdministratorNamePreviewResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdministratorNamePreviewResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdministratorNamePreviewResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdministratorNamePreviewResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AdministratorNamePreviewResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AdministratorNamePreviewResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetError

`func (o *AdministratorNamePreviewResponseModel) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AdministratorNamePreviewResponseModel) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AdministratorNamePreviewResponseModel) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AdministratorNamePreviewResponseModel) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *AdministratorNamePreviewResponseModel) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *AdministratorNamePreviewResponseModel) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


