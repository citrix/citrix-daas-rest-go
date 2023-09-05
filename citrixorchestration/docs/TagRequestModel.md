# TagRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the tag. | [optional] 
**Description** | Pointer to **string** | Description of the tag. | [optional] 
**Uuid** | Pointer to **string** | Unique Id of the tag. | [optional] 
**Scopes** | Pointer to **[]string** | Name of the Scopes. | [optional] 

## Methods

### NewTagRequestModel

`func NewTagRequestModel() *TagRequestModel`

NewTagRequestModel instantiates a new TagRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagRequestModelWithDefaults

`func NewTagRequestModelWithDefaults() *TagRequestModel`

NewTagRequestModelWithDefaults instantiates a new TagRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TagRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TagRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TagRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TagRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TagRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TagRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUuid

`func (o *TagRequestModel) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *TagRequestModel) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *TagRequestModel) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *TagRequestModel) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetScopes

`func (o *TagRequestModel) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TagRequestModel) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TagRequestModel) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TagRequestModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


