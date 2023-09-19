# VDAComponentsSelectionValidationRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludedAdditionalComponents** | Pointer to [**[]VDAComponentRequestModel**](VDAComponentRequestModel.md) | New Components that are to be included/added in this catalog level VDA Upgrade schedule. | [optional] 
**ExcludedComponents** | Pointer to [**[]VDAComponentRequestModel**](VDAComponentRequestModel.md) | Installed Components that are to be excluded/omitted in this catalog level VDA Upgrade schedule. | [optional] 
**Features** | Pointer to **[]string** | Features that needs to enabled on this catalog level VDA Upgrade schedule. | [optional] 

## Methods

### NewVDAComponentsSelectionValidationRequestModel

`func NewVDAComponentsSelectionValidationRequestModel() *VDAComponentsSelectionValidationRequestModel`

NewVDAComponentsSelectionValidationRequestModel instantiates a new VDAComponentsSelectionValidationRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVDAComponentsSelectionValidationRequestModelWithDefaults

`func NewVDAComponentsSelectionValidationRequestModelWithDefaults() *VDAComponentsSelectionValidationRequestModel`

NewVDAComponentsSelectionValidationRequestModelWithDefaults instantiates a new VDAComponentsSelectionValidationRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludedAdditionalComponents

`func (o *VDAComponentsSelectionValidationRequestModel) GetIncludedAdditionalComponents() []VDAComponentRequestModel`

GetIncludedAdditionalComponents returns the IncludedAdditionalComponents field if non-nil, zero value otherwise.

### GetIncludedAdditionalComponentsOk

`func (o *VDAComponentsSelectionValidationRequestModel) GetIncludedAdditionalComponentsOk() (*[]VDAComponentRequestModel, bool)`

GetIncludedAdditionalComponentsOk returns a tuple with the IncludedAdditionalComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedAdditionalComponents

`func (o *VDAComponentsSelectionValidationRequestModel) SetIncludedAdditionalComponents(v []VDAComponentRequestModel)`

SetIncludedAdditionalComponents sets IncludedAdditionalComponents field to given value.

### HasIncludedAdditionalComponents

`func (o *VDAComponentsSelectionValidationRequestModel) HasIncludedAdditionalComponents() bool`

HasIncludedAdditionalComponents returns a boolean if a field has been set.

### SetIncludedAdditionalComponentsNil

`func (o *VDAComponentsSelectionValidationRequestModel) SetIncludedAdditionalComponentsNil(b bool)`

 SetIncludedAdditionalComponentsNil sets the value for IncludedAdditionalComponents to be an explicit nil

### UnsetIncludedAdditionalComponents
`func (o *VDAComponentsSelectionValidationRequestModel) UnsetIncludedAdditionalComponents()`

UnsetIncludedAdditionalComponents ensures that no value is present for IncludedAdditionalComponents, not even an explicit nil
### GetExcludedComponents

`func (o *VDAComponentsSelectionValidationRequestModel) GetExcludedComponents() []VDAComponentRequestModel`

GetExcludedComponents returns the ExcludedComponents field if non-nil, zero value otherwise.

### GetExcludedComponentsOk

`func (o *VDAComponentsSelectionValidationRequestModel) GetExcludedComponentsOk() (*[]VDAComponentRequestModel, bool)`

GetExcludedComponentsOk returns a tuple with the ExcludedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedComponents

`func (o *VDAComponentsSelectionValidationRequestModel) SetExcludedComponents(v []VDAComponentRequestModel)`

SetExcludedComponents sets ExcludedComponents field to given value.

### HasExcludedComponents

`func (o *VDAComponentsSelectionValidationRequestModel) HasExcludedComponents() bool`

HasExcludedComponents returns a boolean if a field has been set.

### SetExcludedComponentsNil

`func (o *VDAComponentsSelectionValidationRequestModel) SetExcludedComponentsNil(b bool)`

 SetExcludedComponentsNil sets the value for ExcludedComponents to be an explicit nil

### UnsetExcludedComponents
`func (o *VDAComponentsSelectionValidationRequestModel) UnsetExcludedComponents()`

UnsetExcludedComponents ensures that no value is present for ExcludedComponents, not even an explicit nil
### GetFeatures

`func (o *VDAComponentsSelectionValidationRequestModel) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *VDAComponentsSelectionValidationRequestModel) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *VDAComponentsSelectionValidationRequestModel) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *VDAComponentsSelectionValidationRequestModel) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### SetFeaturesNil

`func (o *VDAComponentsSelectionValidationRequestModel) SetFeaturesNil(b bool)`

 SetFeaturesNil sets the value for Features to be an explicit nil

### UnsetFeatures
`func (o *VDAComponentsSelectionValidationRequestModel) UnsetFeatures()`

UnsetFeatures ensures that no value is present for Features, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


