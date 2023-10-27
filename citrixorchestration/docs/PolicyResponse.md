# PolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicySetGuid** | Pointer to **string** | The policy set that this policy belongs to. | [optional] 
**PolicyGuid** | Pointer to **string** | Policy GUID. | [optional] 
**PolicyName** | Pointer to **NullableString** | Policy name. | [optional] 
**Priority** | Pointer to **int32** | Policy priority. | [optional] 
**IsEnabled** | Pointer to **bool** | Is policy enabled. | [optional] 
**Description** | Pointer to **NullableString** | Policy description. | [optional] 
**Settings** | Pointer to [**[]SettingResponse**](SettingResponse.md) | Policy settings. | [optional] 
**Filters** | Pointer to [**[]FilterResponse**](FilterResponse.md) | Policy filters. | [optional] 

## Methods

### NewPolicyResponse

`func NewPolicyResponse() *PolicyResponse`

NewPolicyResponse instantiates a new PolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyResponseWithDefaults

`func NewPolicyResponseWithDefaults() *PolicyResponse`

NewPolicyResponseWithDefaults instantiates a new PolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicySetGuid

`func (o *PolicyResponse) GetPolicySetGuid() string`

GetPolicySetGuid returns the PolicySetGuid field if non-nil, zero value otherwise.

### GetPolicySetGuidOk

`func (o *PolicyResponse) GetPolicySetGuidOk() (*string, bool)`

GetPolicySetGuidOk returns a tuple with the PolicySetGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicySetGuid

`func (o *PolicyResponse) SetPolicySetGuid(v string)`

SetPolicySetGuid sets PolicySetGuid field to given value.

### HasPolicySetGuid

`func (o *PolicyResponse) HasPolicySetGuid() bool`

HasPolicySetGuid returns a boolean if a field has been set.

### GetPolicyGuid

`func (o *PolicyResponse) GetPolicyGuid() string`

GetPolicyGuid returns the PolicyGuid field if non-nil, zero value otherwise.

### GetPolicyGuidOk

`func (o *PolicyResponse) GetPolicyGuidOk() (*string, bool)`

GetPolicyGuidOk returns a tuple with the PolicyGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyGuid

`func (o *PolicyResponse) SetPolicyGuid(v string)`

SetPolicyGuid sets PolicyGuid field to given value.

### HasPolicyGuid

`func (o *PolicyResponse) HasPolicyGuid() bool`

HasPolicyGuid returns a boolean if a field has been set.

### GetPolicyName

`func (o *PolicyResponse) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *PolicyResponse) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *PolicyResponse) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *PolicyResponse) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *PolicyResponse) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *PolicyResponse) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetPriority

`func (o *PolicyResponse) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PolicyResponse) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PolicyResponse) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PolicyResponse) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetIsEnabled

`func (o *PolicyResponse) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *PolicyResponse) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *PolicyResponse) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *PolicyResponse) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PolicyResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PolicyResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSettings

`func (o *PolicyResponse) GetSettings() []SettingResponse`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *PolicyResponse) GetSettingsOk() (*[]SettingResponse, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *PolicyResponse) SetSettings(v []SettingResponse)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *PolicyResponse) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *PolicyResponse) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *PolicyResponse) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetFilters

`func (o *PolicyResponse) GetFilters() []FilterResponse`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PolicyResponse) GetFiltersOk() (*[]FilterResponse, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PolicyResponse) SetFilters(v []FilterResponse)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PolicyResponse) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *PolicyResponse) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *PolicyResponse) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


