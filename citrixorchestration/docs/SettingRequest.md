# SettingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingName** | Pointer to **NullableString** | Setting type. Is globally unique. | [optional] 
**UseDefault** | Pointer to **bool** | Indicate if the default setting value is used. If the setting data type is bool, this is ignored. For all other settings, if this value is true, the setting value is ignored and if this value is false, the setting value is used. | [optional] 
**SettingValue** | Pointer to **NullableString** | The setting value. | [optional] 

## Methods

### NewSettingRequest

`func NewSettingRequest() *SettingRequest`

NewSettingRequest instantiates a new SettingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingRequestWithDefaults

`func NewSettingRequestWithDefaults() *SettingRequest`

NewSettingRequestWithDefaults instantiates a new SettingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingName

`func (o *SettingRequest) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *SettingRequest) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *SettingRequest) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *SettingRequest) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### SetSettingNameNil

`func (o *SettingRequest) SetSettingNameNil(b bool)`

 SetSettingNameNil sets the value for SettingName to be an explicit nil

### UnsetSettingName
`func (o *SettingRequest) UnsetSettingName()`

UnsetSettingName ensures that no value is present for SettingName, not even an explicit nil
### GetUseDefault

`func (o *SettingRequest) GetUseDefault() bool`

GetUseDefault returns the UseDefault field if non-nil, zero value otherwise.

### GetUseDefaultOk

`func (o *SettingRequest) GetUseDefaultOk() (*bool, bool)`

GetUseDefaultOk returns a tuple with the UseDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefault

`func (o *SettingRequest) SetUseDefault(v bool)`

SetUseDefault sets UseDefault field to given value.

### HasUseDefault

`func (o *SettingRequest) HasUseDefault() bool`

HasUseDefault returns a boolean if a field has been set.

### GetSettingValue

`func (o *SettingRequest) GetSettingValue() string`

GetSettingValue returns the SettingValue field if non-nil, zero value otherwise.

### GetSettingValueOk

`func (o *SettingRequest) GetSettingValueOk() (*string, bool)`

GetSettingValueOk returns a tuple with the SettingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingValue

`func (o *SettingRequest) SetSettingValue(v string)`

SetSettingValue sets SettingValue field to given value.

### HasSettingValue

`func (o *SettingRequest) HasSettingValue() bool`

HasSettingValue returns a boolean if a field has been set.

### SetSettingValueNil

`func (o *SettingRequest) SetSettingValueNil(b bool)`

 SetSettingValueNil sets the value for SettingValue to be an explicit nil

### UnsetSettingValue
`func (o *SettingRequest) UnsetSettingValue()`

UnsetSettingValue ensures that no value is present for SettingValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


