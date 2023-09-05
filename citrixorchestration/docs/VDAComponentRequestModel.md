# VDAComponentRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentId** | **string** | Id of the component. | 
**Parameters** | Pointer to [**[]VDAComponentParameterRequestModel**](VDAComponentParameterRequestModel.md) |  | [optional] 

## Methods

### NewVDAComponentRequestModel

`func NewVDAComponentRequestModel(componentId string, ) *VDAComponentRequestModel`

NewVDAComponentRequestModel instantiates a new VDAComponentRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVDAComponentRequestModelWithDefaults

`func NewVDAComponentRequestModelWithDefaults() *VDAComponentRequestModel`

NewVDAComponentRequestModelWithDefaults instantiates a new VDAComponentRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentId

`func (o *VDAComponentRequestModel) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *VDAComponentRequestModel) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *VDAComponentRequestModel) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.


### GetParameters

`func (o *VDAComponentRequestModel) GetParameters() []VDAComponentParameterRequestModel`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *VDAComponentRequestModel) GetParametersOk() (*[]VDAComponentParameterRequestModel, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *VDAComponentRequestModel) SetParameters(v []VDAComponentParameterRequestModel)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *VDAComponentRequestModel) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


