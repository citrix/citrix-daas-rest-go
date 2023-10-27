# WonOverBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingName** | Pointer to **NullableString** | Name of the setting. | [optional] 
**WinningPolicy** | Pointer to [**PolicyIdentity**](PolicyIdentity.md) |  | [optional] 

## Methods

### NewWonOverBy

`func NewWonOverBy() *WonOverBy`

NewWonOverBy instantiates a new WonOverBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWonOverByWithDefaults

`func NewWonOverByWithDefaults() *WonOverBy`

NewWonOverByWithDefaults instantiates a new WonOverBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingName

`func (o *WonOverBy) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *WonOverBy) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *WonOverBy) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *WonOverBy) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### SetSettingNameNil

`func (o *WonOverBy) SetSettingNameNil(b bool)`

 SetSettingNameNil sets the value for SettingName to be an explicit nil

### UnsetSettingName
`func (o *WonOverBy) UnsetSettingName()`

UnsetSettingName ensures that no value is present for SettingName, not even an explicit nil
### GetWinningPolicy

`func (o *WonOverBy) GetWinningPolicy() PolicyIdentity`

GetWinningPolicy returns the WinningPolicy field if non-nil, zero value otherwise.

### GetWinningPolicyOk

`func (o *WonOverBy) GetWinningPolicyOk() (*PolicyIdentity, bool)`

GetWinningPolicyOk returns a tuple with the WinningPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinningPolicy

`func (o *WonOverBy) SetWinningPolicy(v PolicyIdentity)`

SetWinningPolicy sets WinningPolicy field to given value.

### HasWinningPolicy

`func (o *WonOverBy) HasWinningPolicy() bool`

HasWinningPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


