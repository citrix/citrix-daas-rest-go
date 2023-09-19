# SettingResponseContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingName** | Pointer to **NullableString** | Setting name. Is globally unique. | [optional] 
**CurrentValue** | Pointer to **NullableString** | The current setting value. * For boolean types, the returned value is \&quot;True\&quot;/\&quot;False\&quot;. The caller should assume the value is   case-insensitive. * For all other types, if the value is null, use default value, otherwise use the specified value. | [optional] 

## Methods

### NewSettingResponseContract

`func NewSettingResponseContract() *SettingResponseContract`

NewSettingResponseContract instantiates a new SettingResponseContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingResponseContractWithDefaults

`func NewSettingResponseContractWithDefaults() *SettingResponseContract`

NewSettingResponseContractWithDefaults instantiates a new SettingResponseContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingName

`func (o *SettingResponseContract) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *SettingResponseContract) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *SettingResponseContract) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *SettingResponseContract) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### SetSettingNameNil

`func (o *SettingResponseContract) SetSettingNameNil(b bool)`

 SetSettingNameNil sets the value for SettingName to be an explicit nil

### UnsetSettingName
`func (o *SettingResponseContract) UnsetSettingName()`

UnsetSettingName ensures that no value is present for SettingName, not even an explicit nil
### GetCurrentValue

`func (o *SettingResponseContract) GetCurrentValue() string`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *SettingResponseContract) GetCurrentValueOk() (*string, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *SettingResponseContract) SetCurrentValue(v string)`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *SettingResponseContract) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### SetCurrentValueNil

`func (o *SettingResponseContract) SetCurrentValueNil(b bool)`

 SetCurrentValueNil sets the value for CurrentValue to be an explicit nil

### UnsetCurrentValue
`func (o *SettingResponseContract) UnsetCurrentValue()`

UnsetCurrentValue ensures that no value is present for CurrentValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


