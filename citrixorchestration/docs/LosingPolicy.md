# LosingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyIdentity** | Pointer to [**PolicyIdentity**](PolicyIdentity.md) |  | [optional] 
**Reasons** | Pointer to **map[string][]string** | Reasons why the policy is not applied according to filtering results. | [optional] 
**Priority** | Pointer to **int32** | Policy priority. | [optional] 
**Settings** | Pointer to [**[]WonOverBy**](WonOverBy.md) | The settings in the policy and the reason of losing for each setting. | [optional] 

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

### GetPolicyIdentity

`func (o *LosingPolicy) GetPolicyIdentity() PolicyIdentity`

GetPolicyIdentity returns the PolicyIdentity field if non-nil, zero value otherwise.

### GetPolicyIdentityOk

`func (o *LosingPolicy) GetPolicyIdentityOk() (*PolicyIdentity, bool)`

GetPolicyIdentityOk returns a tuple with the PolicyIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyIdentity

`func (o *LosingPolicy) SetPolicyIdentity(v PolicyIdentity)`

SetPolicyIdentity sets PolicyIdentity field to given value.

### HasPolicyIdentity

`func (o *LosingPolicy) HasPolicyIdentity() bool`

HasPolicyIdentity returns a boolean if a field has been set.

### GetReasons

`func (o *LosingPolicy) GetReasons() map[string][]string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *LosingPolicy) GetReasonsOk() (*map[string][]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *LosingPolicy) SetReasons(v map[string][]string)`

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

### GetSettings

`func (o *LosingPolicy) GetSettings() []WonOverBy`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *LosingPolicy) GetSettingsOk() (*[]WonOverBy, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *LosingPolicy) SetSettings(v []WonOverBy)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *LosingPolicy) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *LosingPolicy) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *LosingPolicy) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


