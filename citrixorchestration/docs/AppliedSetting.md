# AppliedSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingName** | Pointer to **NullableString** | Setting name. | [optional] 
**SettingValue** | Pointer to **NullableString** | Setting value. | [optional] 
**PolicyName** | Pointer to **NullableString** | Name of the policy. | [optional] 
**GpoName** | Pointer to **NullableString** | Name of the GPO that contains the policy that uses this setting. | [optional] 

## Methods

### NewAppliedSetting

`func NewAppliedSetting() *AppliedSetting`

NewAppliedSetting instantiates a new AppliedSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppliedSettingWithDefaults

`func NewAppliedSettingWithDefaults() *AppliedSetting`

NewAppliedSettingWithDefaults instantiates a new AppliedSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingName

`func (o *AppliedSetting) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *AppliedSetting) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *AppliedSetting) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *AppliedSetting) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### SetSettingNameNil

`func (o *AppliedSetting) SetSettingNameNil(b bool)`

 SetSettingNameNil sets the value for SettingName to be an explicit nil

### UnsetSettingName
`func (o *AppliedSetting) UnsetSettingName()`

UnsetSettingName ensures that no value is present for SettingName, not even an explicit nil
### GetSettingValue

`func (o *AppliedSetting) GetSettingValue() string`

GetSettingValue returns the SettingValue field if non-nil, zero value otherwise.

### GetSettingValueOk

`func (o *AppliedSetting) GetSettingValueOk() (*string, bool)`

GetSettingValueOk returns a tuple with the SettingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingValue

`func (o *AppliedSetting) SetSettingValue(v string)`

SetSettingValue sets SettingValue field to given value.

### HasSettingValue

`func (o *AppliedSetting) HasSettingValue() bool`

HasSettingValue returns a boolean if a field has been set.

### SetSettingValueNil

`func (o *AppliedSetting) SetSettingValueNil(b bool)`

 SetSettingValueNil sets the value for SettingValue to be an explicit nil

### UnsetSettingValue
`func (o *AppliedSetting) UnsetSettingValue()`

UnsetSettingValue ensures that no value is present for SettingValue, not even an explicit nil
### GetPolicyName

`func (o *AppliedSetting) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *AppliedSetting) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *AppliedSetting) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *AppliedSetting) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *AppliedSetting) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *AppliedSetting) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetGpoName

`func (o *AppliedSetting) GetGpoName() string`

GetGpoName returns the GpoName field if non-nil, zero value otherwise.

### GetGpoNameOk

`func (o *AppliedSetting) GetGpoNameOk() (*string, bool)`

GetGpoNameOk returns a tuple with the GpoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpoName

`func (o *AppliedSetting) SetGpoName(v string)`

SetGpoName sets GpoName field to given value.

### HasGpoName

`func (o *AppliedSetting) HasGpoName() bool`

HasGpoName returns a boolean if a field has been set.

### SetGpoNameNil

`func (o *AppliedSetting) SetGpoNameNil(b bool)`

 SetGpoNameNil sets the value for GpoName to be an explicit nil

### UnsetGpoName
`func (o *AppliedSetting) UnsetGpoName()`

UnsetGpoName ensures that no value is present for GpoName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


