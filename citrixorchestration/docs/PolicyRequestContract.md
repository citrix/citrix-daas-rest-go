# PolicyRequestContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyName** | Pointer to **NullableString** | Policy name as identifier, translated. | [optional] 
**IsEnabled** | Pointer to **NullableBool** | Is policy enabled | [optional] [default to false]
**Description** | Pointer to **NullableString** | Policy description | [optional] 
**Priority** | Pointer to **NullableInt32** | New policy priority | [optional] 
**Settings** | Pointer to [**[]SettingRequestContract**](SettingRequestContract.md) | Policy settings | [optional] 
**Filters** | Pointer to [**[]FilterRequestContract**](FilterRequestContract.md) | Policy filters | [optional] 

## Methods

### NewPolicyRequestContract

`func NewPolicyRequestContract() *PolicyRequestContract`

NewPolicyRequestContract instantiates a new PolicyRequestContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRequestContractWithDefaults

`func NewPolicyRequestContractWithDefaults() *PolicyRequestContract`

NewPolicyRequestContractWithDefaults instantiates a new PolicyRequestContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyName

`func (o *PolicyRequestContract) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *PolicyRequestContract) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *PolicyRequestContract) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *PolicyRequestContract) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *PolicyRequestContract) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *PolicyRequestContract) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetIsEnabled

`func (o *PolicyRequestContract) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *PolicyRequestContract) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *PolicyRequestContract) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *PolicyRequestContract) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *PolicyRequestContract) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *PolicyRequestContract) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetDescription

`func (o *PolicyRequestContract) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyRequestContract) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyRequestContract) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyRequestContract) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PolicyRequestContract) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PolicyRequestContract) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPriority

`func (o *PolicyRequestContract) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PolicyRequestContract) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PolicyRequestContract) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PolicyRequestContract) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *PolicyRequestContract) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *PolicyRequestContract) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetSettings

`func (o *PolicyRequestContract) GetSettings() []SettingRequestContract`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *PolicyRequestContract) GetSettingsOk() (*[]SettingRequestContract, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *PolicyRequestContract) SetSettings(v []SettingRequestContract)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *PolicyRequestContract) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *PolicyRequestContract) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *PolicyRequestContract) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetFilters

`func (o *PolicyRequestContract) GetFilters() []FilterRequestContract`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PolicyRequestContract) GetFiltersOk() (*[]FilterRequestContract, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PolicyRequestContract) SetFilters(v []FilterRequestContract)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PolicyRequestContract) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *PolicyRequestContract) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *PolicyRequestContract) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


