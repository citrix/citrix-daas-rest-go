# CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludedAdditionalComponents** | Pointer to [**[]VDAComponentRequestModel**](VDAComponentRequestModel.md) | New Components that are to be included/added in this catalog level VDA Upgrade schedule. | [optional] 
**ExcludedComponents** | Pointer to [**[]VDAComponentRequestModel**](VDAComponentRequestModel.md) | Installed Components that are to be excluded/omitted in this catalog level VDA Upgrade schedule. | [optional] 
**Features** | Pointer to **[]string** | Features that needs to enabled on this catalog level VDA Upgrade schedule. | [optional] 

## Methods

### NewCreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel

`func NewCreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel() *CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel`

NewCreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel instantiates a new CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModelWithDefaults

`func NewCreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModelWithDefaults() *CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel`

NewCreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModelWithDefaults instantiates a new CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludedAdditionalComponents

`func (o *CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel) GetIncludedAdditionalComponents() []VDAComponentRequestModel`

GetIncludedAdditionalComponents returns the IncludedAdditionalComponents field if non-nil, zero value otherwise.

### GetIncludedAdditionalComponentsOk

`func (o *CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel) GetIncludedAdditionalComponentsOk() (*[]VDAComponentRequestModel, bool)`

GetIncludedAdditionalComponentsOk returns a tuple with the IncludedAdditionalComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedAdditionalComponents

`func (o *CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel) SetIncludedAdditionalComponents(v []VDAComponentRequestModel)`

SetIncludedAdditionalComponents sets IncludedAdditionalComponents field to given value.

### HasIncludedAdditionalComponents

`func (o *CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel) HasIncludedAdditionalComponents() bool`

HasIncludedAdditionalComponents returns a boolean if a field has been set.

### GetExcludedComponents

`func (o *CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel) GetExcludedComponents() []VDAComponentRequestModel`

GetExcludedComponents returns the ExcludedComponents field if non-nil, zero value otherwise.

### GetExcludedComponentsOk

`func (o *CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel) GetExcludedComponentsOk() (*[]VDAComponentRequestModel, bool)`

GetExcludedComponentsOk returns a tuple with the ExcludedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedComponents

`func (o *CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel) SetExcludedComponents(v []VDAComponentRequestModel)`

SetExcludedComponents sets ExcludedComponents field to given value.

### HasExcludedComponents

`func (o *CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel) HasExcludedComponents() bool`

HasExcludedComponents returns a boolean if a field has been set.

### GetFeatures

`func (o *CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


