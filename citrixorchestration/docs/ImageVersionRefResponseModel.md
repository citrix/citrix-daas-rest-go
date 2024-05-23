# ImageVersionRefResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Id of the image version. | 
**Number** | **int32** | The version number associated with the image version. | 
**ImageDefinition** | [**RefResponseModel**](RefResponseModel.md) |  | 
**Description** | Pointer to **NullableString** | The image version&#39;s description | [optional] 
**ImageVersionSpecs** | Pointer to [**[]ImageVersionSpecRefResponseModel**](ImageVersionSpecRefResponseModel.md) | The image version specifications associated with this image version. | [optional] 

## Methods

### NewImageVersionRefResponseModel

`func NewImageVersionRefResponseModel(id string, number int32, imageDefinition RefResponseModel, ) *ImageVersionRefResponseModel`

NewImageVersionRefResponseModel instantiates a new ImageVersionRefResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageVersionRefResponseModelWithDefaults

`func NewImageVersionRefResponseModelWithDefaults() *ImageVersionRefResponseModel`

NewImageVersionRefResponseModelWithDefaults instantiates a new ImageVersionRefResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageVersionRefResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageVersionRefResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageVersionRefResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetNumber

`func (o *ImageVersionRefResponseModel) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ImageVersionRefResponseModel) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ImageVersionRefResponseModel) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetImageDefinition

`func (o *ImageVersionRefResponseModel) GetImageDefinition() RefResponseModel`

GetImageDefinition returns the ImageDefinition field if non-nil, zero value otherwise.

### GetImageDefinitionOk

`func (o *ImageVersionRefResponseModel) GetImageDefinitionOk() (*RefResponseModel, bool)`

GetImageDefinitionOk returns a tuple with the ImageDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDefinition

`func (o *ImageVersionRefResponseModel) SetImageDefinition(v RefResponseModel)`

SetImageDefinition sets ImageDefinition field to given value.


### GetDescription

`func (o *ImageVersionRefResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImageVersionRefResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImageVersionRefResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImageVersionRefResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ImageVersionRefResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ImageVersionRefResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageVersionSpecs

`func (o *ImageVersionRefResponseModel) GetImageVersionSpecs() []ImageVersionSpecRefResponseModel`

GetImageVersionSpecs returns the ImageVersionSpecs field if non-nil, zero value otherwise.

### GetImageVersionSpecsOk

`func (o *ImageVersionRefResponseModel) GetImageVersionSpecsOk() (*[]ImageVersionSpecRefResponseModel, bool)`

GetImageVersionSpecsOk returns a tuple with the ImageVersionSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersionSpecs

`func (o *ImageVersionRefResponseModel) SetImageVersionSpecs(v []ImageVersionSpecRefResponseModel)`

SetImageVersionSpecs sets ImageVersionSpecs field to given value.

### HasImageVersionSpecs

`func (o *ImageVersionRefResponseModel) HasImageVersionSpecs() bool`

HasImageVersionSpecs returns a boolean if a field has been set.

### SetImageVersionSpecsNil

`func (o *ImageVersionRefResponseModel) SetImageVersionSpecsNil(b bool)`

 SetImageVersionSpecsNil sets the value for ImageVersionSpecs to be an explicit nil

### UnsetImageVersionSpecs
`func (o *ImageVersionRefResponseModel) UnsetImageVersionSpecs()`

UnsetImageVersionSpecs ensures that no value is present for ImageVersionSpecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


