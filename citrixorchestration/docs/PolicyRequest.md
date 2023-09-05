# PolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Policy name. | [optional] 
**IsEnabled** | **bool** | Is policy enabled | 
**Description** | Pointer to **string** | Policy description | [optional] 
**Settings** | Pointer to [**[]SettingRequest**](SettingRequest.md) | Policy settings | [optional] 
**Filters** | Pointer to [**[]FilterRequest**](FilterRequest.md) | Policy filters | [optional] 

## Methods

### NewPolicyRequest

`func NewPolicyRequest(isEnabled bool, ) *PolicyRequest`

NewPolicyRequest instantiates a new PolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRequestWithDefaults

`func NewPolicyRequestWithDefaults() *PolicyRequest`

NewPolicyRequestWithDefaults instantiates a new PolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsEnabled

`func (o *PolicyRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *PolicyRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *PolicyRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetDescription

`func (o *PolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSettings

`func (o *PolicyRequest) GetSettings() []SettingRequest`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *PolicyRequest) GetSettingsOk() (*[]SettingRequest, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *PolicyRequest) SetSettings(v []SettingRequest)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *PolicyRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetFilters

`func (o *PolicyRequest) GetFilters() []FilterRequest`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PolicyRequest) GetFiltersOk() (*[]FilterRequest, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PolicyRequest) SetFilters(v []FilterRequest)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PolicyRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


