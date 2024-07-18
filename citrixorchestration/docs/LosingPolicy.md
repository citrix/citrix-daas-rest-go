# LosingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyName** | Pointer to **NullableString** | Name of the policy. | [optional] 
**GpoName** | Pointer to **NullableString** | Name of the GPO that contains the policy that uses this setting. | [optional] 
**Reasons** | Pointer to [**map[string][]ReasonDetail**](array.md) | Reasons why the policy is not applied according to filtering results. | [optional] 
**Priority** | Pointer to **int32** | Policy priority. | [optional] 
**WinningSettings** | Pointer to [**[]WonOverBy**](WonOverBy.md) | The settings in the policy and the reason of losing for each setting. | [optional] 

## Methods

### NewLosingPolicy

`func NewLosingPolicy() *LosingPolicy`

NewLosingPolicy instantiates a new LosingPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLosingPolicyWithDefaults

`func NewLosingPolicyWithDefaults() *LosingPolicy`

NewLosingPolicyWithDefaults instantiates a new LosingPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyName

`func (o *LosingPolicy) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *LosingPolicy) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *LosingPolicy) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *LosingPolicy) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *LosingPolicy) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *LosingPolicy) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetGpoName

`func (o *LosingPolicy) GetGpoName() string`

GetGpoName returns the GpoName field if non-nil, zero value otherwise.

### GetGpoNameOk

`func (o *LosingPolicy) GetGpoNameOk() (*string, bool)`

GetGpoNameOk returns a tuple with the GpoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpoName

`func (o *LosingPolicy) SetGpoName(v string)`

SetGpoName sets GpoName field to given value.

### HasGpoName

`func (o *LosingPolicy) HasGpoName() bool`

HasGpoName returns a boolean if a field has been set.

### SetGpoNameNil

`func (o *LosingPolicy) SetGpoNameNil(b bool)`

 SetGpoNameNil sets the value for GpoName to be an explicit nil

### UnsetGpoName
`func (o *LosingPolicy) UnsetGpoName()`

UnsetGpoName ensures that no value is present for GpoName, not even an explicit nil
### GetReasons

`func (o *LosingPolicy) GetReasons() map[string][]ReasonDetail`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *LosingPolicy) GetReasonsOk() (*map[string][]ReasonDetail, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *LosingPolicy) SetReasons(v map[string][]ReasonDetail)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *LosingPolicy) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### SetReasonsNil

`func (o *LosingPolicy) SetReasonsNil(b bool)`

 SetReasonsNil sets the value for Reasons to be an explicit nil

### UnsetReasons
`func (o *LosingPolicy) UnsetReasons()`

UnsetReasons ensures that no value is present for Reasons, not even an explicit nil
### GetPriority

`func (o *LosingPolicy) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *LosingPolicy) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *LosingPolicy) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *LosingPolicy) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetWinningSettings

`func (o *LosingPolicy) GetWinningSettings() []WonOverBy`

GetWinningSettings returns the WinningSettings field if non-nil, zero value otherwise.

### GetWinningSettingsOk

`func (o *LosingPolicy) GetWinningSettingsOk() (*[]WonOverBy, bool)`

GetWinningSettingsOk returns a tuple with the WinningSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinningSettings

`func (o *LosingPolicy) SetWinningSettings(v []WonOverBy)`

SetWinningSettings sets WinningSettings field to given value.

### HasWinningSettings

`func (o *LosingPolicy) HasWinningSettings() bool`

HasWinningSettings returns a boolean if a field has been set.

### SetWinningSettingsNil

`func (o *LosingPolicy) SetWinningSettingsNil(b bool)`

 SetWinningSettingsNil sets the value for WinningSettings to be an explicit nil

### UnsetWinningSettings
`func (o *LosingPolicy) UnsetWinningSettings()`

UnsetWinningSettings ensures that no value is present for WinningSettings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


