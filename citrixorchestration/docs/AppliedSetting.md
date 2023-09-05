# AppliedSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingName** | Pointer to **string** | Setting name | [optional] 
**SettingValue** | Pointer to **string** | Setting value | [optional] 
**WinningGpo** | Pointer to **string** | Source GPO for the setting | [optional] 
**WinningPolicy** | Pointer to **string** | Source policy for the setting | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


