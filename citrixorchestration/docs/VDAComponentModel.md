# VDAComponentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentId** | Pointer to **NullableString** | Id of this component object. | [optional] 
**ComponentName** | Pointer to **NullableString** | Name of this component object. | [optional] 
**ComponentDescription** | Pointer to **NullableString** | Description of this component object. | [optional] 
**Parameters** | Pointer to [**[]VDAComponentParameterModel**](VDAComponentParameterModel.md) | Parameters of this component object. | [optional] 

## Methods

### NewVDAComponentModel

`func NewVDAComponentModel() *VDAComponentModel`

NewVDAComponentModel instantiates a new VDAComponentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVDAComponentModelWithDefaults

`func NewVDAComponentModelWithDefaults() *VDAComponentModel`

NewVDAComponentModelWithDefaults instantiates a new VDAComponentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentId

`func (o *VDAComponentModel) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *VDAComponentModel) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *VDAComponentModel) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *VDAComponentModel) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### SetComponentIdNil

`func (o *VDAComponentModel) SetComponentIdNil(b bool)`

 SetComponentIdNil sets the value for ComponentId to be an explicit nil

### UnsetComponentId
`func (o *VDAComponentModel) UnsetComponentId()`

UnsetComponentId ensures that no value is present for ComponentId, not even an explicit nil
### GetComponentName

`func (o *VDAComponentModel) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *VDAComponentModel) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *VDAComponentModel) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *VDAComponentModel) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.

### SetComponentNameNil

`func (o *VDAComponentModel) SetComponentNameNil(b bool)`

 SetComponentNameNil sets the value for ComponentName to be an explicit nil

### UnsetComponentName
`func (o *VDAComponentModel) UnsetComponentName()`

UnsetComponentName ensures that no value is present for ComponentName, not even an explicit nil
### GetComponentDescription

`func (o *VDAComponentModel) GetComponentDescription() string`

GetComponentDescription returns the ComponentDescription field if non-nil, zero value otherwise.

### GetComponentDescriptionOk

`func (o *VDAComponentModel) GetComponentDescriptionOk() (*string, bool)`

GetComponentDescriptionOk returns a tuple with the ComponentDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentDescription

`func (o *VDAComponentModel) SetComponentDescription(v string)`

SetComponentDescription sets ComponentDescription field to given value.

### HasComponentDescription

`func (o *VDAComponentModel) HasComponentDescription() bool`

HasComponentDescription returns a boolean if a field has been set.

### SetComponentDescriptionNil

`func (o *VDAComponentModel) SetComponentDescriptionNil(b bool)`

 SetComponentDescriptionNil sets the value for ComponentDescription to be an explicit nil

### UnsetComponentDescription
`func (o *VDAComponentModel) UnsetComponentDescription()`

UnsetComponentDescription ensures that no value is present for ComponentDescription, not even an explicit nil
### GetParameters

`func (o *VDAComponentModel) GetParameters() []VDAComponentParameterModel`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *VDAComponentModel) GetParametersOk() (*[]VDAComponentParameterModel, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *VDAComponentModel) SetParameters(v []VDAComponentParameterModel)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *VDAComponentModel) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *VDAComponentModel) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *VDAComponentModel) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


