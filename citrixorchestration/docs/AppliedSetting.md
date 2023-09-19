# AppliedSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingName** | Pointer to **NullableString** | Setting name | [optional] 
**SettingValue** | Pointer to **NullableString** | Setting value | [optional] 
**WinningGpo** | Pointer to **NullableString** | Source GPO for the setting | [optional] 
**WinningPolicy** | Pointer to **NullableString** | Source policy for the setting | [optional] 

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
### GetWinningGpo

`func (o *AppliedSetting) GetWinningGpo() string`

GetWinningGpo returns the WinningGpo field if non-nil, zero value otherwise.

### GetWinningGpoOk

`func (o *AppliedSetting) GetWinningGpoOk() (*string, bool)`

GetWinningGpoOk returns a tuple with the WinningGpo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinningGpo

`func (o *AppliedSetting) SetWinningGpo(v string)`

SetWinningGpo sets WinningGpo field to given value.

### HasWinningGpo

`func (o *AppliedSetting) HasWinningGpo() bool`

HasWinningGpo returns a boolean if a field has been set.

### SetWinningGpoNil

`func (o *AppliedSetting) SetWinningGpoNil(b bool)`

 SetWinningGpoNil sets the value for WinningGpo to be an explicit nil

### UnsetWinningGpo
`func (o *AppliedSetting) UnsetWinningGpo()`

UnsetWinningGpo ensures that no value is present for WinningGpo, not even an explicit nil
### GetWinningPolicy

`func (o *AppliedSetting) GetWinningPolicy() string`

GetWinningPolicy returns the WinningPolicy field if non-nil, zero value otherwise.

### GetWinningPolicyOk

`func (o *AppliedSetting) GetWinningPolicyOk() (*string, bool)`

GetWinningPolicyOk returns a tuple with the WinningPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinningPolicy

`func (o *AppliedSetting) SetWinningPolicy(v string)`

SetWinningPolicy sets WinningPolicy field to given value.

### HasWinningPolicy

`func (o *AppliedSetting) HasWinningPolicy() bool`

HasWinningPolicy returns a boolean if a field has been set.

### SetWinningPolicyNil

`func (o *AppliedSetting) SetWinningPolicyNil(b bool)`

 SetWinningPolicyNil sets the value for WinningPolicy to be an explicit nil

### UnsetWinningPolicy
`func (o *AppliedSetting) UnsetWinningPolicy()`

UnsetWinningPolicy ensures that no value is present for WinningPolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


