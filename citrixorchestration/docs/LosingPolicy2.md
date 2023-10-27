# LosingPolicy2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyGuid** | Pointer to **string** | Policy GUID. | [optional] 
**PolicyName** | Pointer to **NullableString** | The policy name. | [optional] 
**Priority** | Pointer to **int32** | Policy priority. | [optional] 
**Settings** | Pointer to [**[]WonOverBy2**](WonOverBy2.md) | The settings in the policy and the reason of losing for each setting. | [optional] 

## Methods

### NewLosingPolicy2

`func NewLosingPolicy2() *LosingPolicy2`

NewLosingPolicy2 instantiates a new LosingPolicy2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLosingPolicy2WithDefaults

`func NewLosingPolicy2WithDefaults() *LosingPolicy2`

NewLosingPolicy2WithDefaults instantiates a new LosingPolicy2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyGuid

`func (o *LosingPolicy2) GetPolicyGuid() string`

GetPolicyGuid returns the PolicyGuid field if non-nil, zero value otherwise.

### GetPolicyGuidOk

`func (o *LosingPolicy2) GetPolicyGuidOk() (*string, bool)`

GetPolicyGuidOk returns a tuple with the PolicyGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyGuid

`func (o *LosingPolicy2) SetPolicyGuid(v string)`

SetPolicyGuid sets PolicyGuid field to given value.

### HasPolicyGuid

`func (o *LosingPolicy2) HasPolicyGuid() bool`

HasPolicyGuid returns a boolean if a field has been set.

### GetPolicyName

`func (o *LosingPolicy2) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *LosingPolicy2) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *LosingPolicy2) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *LosingPolicy2) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *LosingPolicy2) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *LosingPolicy2) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetPriority

`func (o *LosingPolicy2) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *LosingPolicy2) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *LosingPolicy2) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *LosingPolicy2) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSettings

`func (o *LosingPolicy2) GetSettings() []WonOverBy2`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *LosingPolicy2) GetSettingsOk() (*[]WonOverBy2, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *LosingPolicy2) SetSettings(v []WonOverBy2)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *LosingPolicy2) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *LosingPolicy2) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *LosingPolicy2) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


