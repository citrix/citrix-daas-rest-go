# LosingSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyGuid** | Pointer to **string** | Policy GUID. | [optional] 
**PolicyName** | Pointer to **NullableString** | The policy that has the setting. | [optional] 
**SettingName** | Pointer to **NullableString** | The setting that is not applied. | [optional] 

## Methods

### NewLosingSetting

`func NewLosingSetting() *LosingSetting`

NewLosingSetting instantiates a new LosingSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLosingSettingWithDefaults

`func NewLosingSettingWithDefaults() *LosingSetting`

NewLosingSettingWithDefaults instantiates a new LosingSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyGuid

`func (o *LosingSetting) GetPolicyGuid() string`

GetPolicyGuid returns the PolicyGuid field if non-nil, zero value otherwise.

### GetPolicyGuidOk

`func (o *LosingSetting) GetPolicyGuidOk() (*string, bool)`

GetPolicyGuidOk returns a tuple with the PolicyGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyGuid

`func (o *LosingSetting) SetPolicyGuid(v string)`

SetPolicyGuid sets PolicyGuid field to given value.

### HasPolicyGuid

`func (o *LosingSetting) HasPolicyGuid() bool`

HasPolicyGuid returns a boolean if a field has been set.

### GetPolicyName

`func (o *LosingSetting) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *LosingSetting) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *LosingSetting) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *LosingSetting) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *LosingSetting) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *LosingSetting) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetSettingName

`func (o *LosingSetting) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *LosingSetting) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *LosingSetting) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *LosingSetting) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### SetSettingNameNil

`func (o *LosingSetting) SetSettingNameNil(b bool)`

 SetSettingNameNil sets the value for SettingName to be an explicit nil

### UnsetSettingName
`func (o *LosingSetting) UnsetSettingName()`

UnsetSettingName ensures that no value is present for SettingName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


