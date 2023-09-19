# SettingRequestContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingType** | Pointer to **NullableString** | Setting type. Is globally unique. | [optional] 
**SettingValue** | Pointer to **NullableString** | The setting value. * For all setting types, excluding booleans, if null, use the default value. * For all setting types, including booleans, otherwise, use the specified value. * For boolean types, the value can be \&quot;true\&quot;/\&quot;false\&quot;, \&quot;1\&quot;/\&quot;0\&quot;, null will be accepted as using default value. | [optional] 
**TypedValue** | Pointer to **map[string]interface{}** | The value in native C# type. If this value is null, use SettingValue, otherwise use this value. | [optional] 

## Methods

### NewSettingRequestContract

`func NewSettingRequestContract() *SettingRequestContract`

NewSettingRequestContract instantiates a new SettingRequestContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingRequestContractWithDefaults

`func NewSettingRequestContractWithDefaults() *SettingRequestContract`

NewSettingRequestContractWithDefaults instantiates a new SettingRequestContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingType

`func (o *SettingRequestContract) GetSettingType() string`

GetSettingType returns the SettingType field if non-nil, zero value otherwise.

### GetSettingTypeOk

`func (o *SettingRequestContract) GetSettingTypeOk() (*string, bool)`

GetSettingTypeOk returns a tuple with the SettingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingType

`func (o *SettingRequestContract) SetSettingType(v string)`

SetSettingType sets SettingType field to given value.

### HasSettingType

`func (o *SettingRequestContract) HasSettingType() bool`

HasSettingType returns a boolean if a field has been set.

### SetSettingTypeNil

`func (o *SettingRequestContract) SetSettingTypeNil(b bool)`

 SetSettingTypeNil sets the value for SettingType to be an explicit nil

### UnsetSettingType
`func (o *SettingRequestContract) UnsetSettingType()`

UnsetSettingType ensures that no value is present for SettingType, not even an explicit nil
### GetSettingValue

`func (o *SettingRequestContract) GetSettingValue() string`

GetSettingValue returns the SettingValue field if non-nil, zero value otherwise.

### GetSettingValueOk

`func (o *SettingRequestContract) GetSettingValueOk() (*string, bool)`

GetSettingValueOk returns a tuple with the SettingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingValue

`func (o *SettingRequestContract) SetSettingValue(v string)`

SetSettingValue sets SettingValue field to given value.

### HasSettingValue

`func (o *SettingRequestContract) HasSettingValue() bool`

HasSettingValue returns a boolean if a field has been set.

### SetSettingValueNil

`func (o *SettingRequestContract) SetSettingValueNil(b bool)`

 SetSettingValueNil sets the value for SettingValue to be an explicit nil

### UnsetSettingValue
`func (o *SettingRequestContract) UnsetSettingValue()`

UnsetSettingValue ensures that no value is present for SettingValue, not even an explicit nil
### GetTypedValue

`func (o *SettingRequestContract) GetTypedValue() map[string]interface{}`

GetTypedValue returns the TypedValue field if non-nil, zero value otherwise.

### GetTypedValueOk

`func (o *SettingRequestContract) GetTypedValueOk() (*map[string]interface{}, bool)`

GetTypedValueOk returns a tuple with the TypedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypedValue

`func (o *SettingRequestContract) SetTypedValue(v map[string]interface{})`

SetTypedValue sets TypedValue field to given value.

### HasTypedValue

`func (o *SettingRequestContract) HasTypedValue() bool`

HasTypedValue returns a boolean if a field has been set.

### SetTypedValueNil

`func (o *SettingRequestContract) SetTypedValueNil(b bool)`

 SetTypedValueNil sets the value for TypedValue to be an explicit nil

### UnsetTypedValue
`func (o *SettingRequestContract) UnsetTypedValue()`

UnsetTypedValue ensures that no value is present for TypedValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


