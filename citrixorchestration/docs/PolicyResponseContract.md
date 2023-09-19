# PolicyResponseContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyName** | Pointer to **NullableString** | Policy name as identifier, translated. | [optional] 
**Priority** | Pointer to **int32** | Policy priority | [optional] 
**IsEnabled** | Pointer to **bool** | Is policy enabled | [optional] 
**Description** | Pointer to **NullableString** | Policy description | [optional] 
**IsUnfiltered** | Pointer to **bool** | Is unfiltered is true only for the &#39;Unfiltered&#39; policy. | [optional] 
**Settings** | Pointer to [**[]SettingResponseContract**](SettingResponseContract.md) | Policy settings | [optional] 
**Filters** | Pointer to [**[]FilterResponseContract**](FilterResponseContract.md) | Policy filters | [optional] 

## Methods

### NewPolicyResponseContract

`func NewPolicyResponseContract() *PolicyResponseContract`

NewPolicyResponseContract instantiates a new PolicyResponseContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyResponseContractWithDefaults

`func NewPolicyResponseContractWithDefaults() *PolicyResponseContract`

NewPolicyResponseContractWithDefaults instantiates a new PolicyResponseContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyName

`func (o *PolicyResponseContract) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *PolicyResponseContract) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *PolicyResponseContract) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *PolicyResponseContract) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *PolicyResponseContract) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *PolicyResponseContract) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetPriority

`func (o *PolicyResponseContract) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PolicyResponseContract) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PolicyResponseContract) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PolicyResponseContract) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetIsEnabled

`func (o *PolicyResponseContract) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *PolicyResponseContract) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *PolicyResponseContract) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *PolicyResponseContract) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyResponseContract) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyResponseContract) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyResponseContract) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyResponseContract) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PolicyResponseContract) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PolicyResponseContract) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsUnfiltered

`func (o *PolicyResponseContract) GetIsUnfiltered() bool`

GetIsUnfiltered returns the IsUnfiltered field if non-nil, zero value otherwise.

### GetIsUnfilteredOk

`func (o *PolicyResponseContract) GetIsUnfilteredOk() (*bool, bool)`

GetIsUnfilteredOk returns a tuple with the IsUnfiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnfiltered

`func (o *PolicyResponseContract) SetIsUnfiltered(v bool)`

SetIsUnfiltered sets IsUnfiltered field to given value.

### HasIsUnfiltered

`func (o *PolicyResponseContract) HasIsUnfiltered() bool`

HasIsUnfiltered returns a boolean if a field has been set.

### GetSettings

`func (o *PolicyResponseContract) GetSettings() []SettingResponseContract`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *PolicyResponseContract) GetSettingsOk() (*[]SettingResponseContract, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *PolicyResponseContract) SetSettings(v []SettingResponseContract)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *PolicyResponseContract) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *PolicyResponseContract) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *PolicyResponseContract) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetFilters

`func (o *PolicyResponseContract) GetFilters() []FilterResponseContract`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PolicyResponseContract) GetFiltersOk() (*[]FilterResponseContract, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PolicyResponseContract) SetFilters(v []FilterResponseContract)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PolicyResponseContract) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *PolicyResponseContract) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *PolicyResponseContract) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


