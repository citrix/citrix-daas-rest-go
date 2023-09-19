# CreateImageDefinitionRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageDefinition** | Pointer to [**CreateImageDefinitionDetailRequestModel**](CreateImageDefinitionDetailRequestModel.md) |  | [optional] 
**ImageVersion** | Pointer to [**CreateImageVersionRequestModel**](CreateImageVersionRequestModel.md) |  | [optional] 

## Methods

### NewCreateImageDefinitionRequestModel

`func NewCreateImageDefinitionRequestModel() *CreateImageDefinitionRequestModel`

NewCreateImageDefinitionRequestModel instantiates a new CreateImageDefinitionRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateImageDefinitionRequestModelWithDefaults

`func NewCreateImageDefinitionRequestModelWithDefaults() *CreateImageDefinitionRequestModel`

NewCreateImageDefinitionRequestModelWithDefaults instantiates a new CreateImageDefinitionRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageDefinition

`func (o *CreateImageDefinitionRequestModel) GetImageDefinition() CreateImageDefinitionDetailRequestModel`

GetImageDefinition returns the ImageDefinition field if non-nil, zero value otherwise.

### GetImageDefinitionOk

`func (o *CreateImageDefinitionRequestModel) GetImageDefinitionOk() (*CreateImageDefinitionDetailRequestModel, bool)`

GetImageDefinitionOk returns a tuple with the ImageDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDefinition

`func (o *CreateImageDefinitionRequestModel) SetImageDefinition(v CreateImageDefinitionDetailRequestModel)`

SetImageDefinition sets ImageDefinition field to given value.

### HasImageDefinition

`func (o *CreateImageDefinitionRequestModel) HasImageDefinition() bool`

HasImageDefinition returns a boolean if a field has been set.

### GetImageVersion

`func (o *CreateImageDefinitionRequestModel) GetImageVersion() CreateImageVersionRequestModel`

GetImageVersion returns the ImageVersion field if non-nil, zero value otherwise.

### GetImageVersionOk

`func (o *CreateImageDefinitionRequestModel) GetImageVersionOk() (*CreateImageVersionRequestModel, bool)`

GetImageVersionOk returns a tuple with the ImageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersion

`func (o *CreateImageDefinitionRequestModel) SetImageVersion(v CreateImageVersionRequestModel)`

SetImageVersion sets ImageVersion field to given value.

### HasImageVersion

`func (o *CreateImageDefinitionRequestModel) HasImageVersion() bool`

HasImageVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


