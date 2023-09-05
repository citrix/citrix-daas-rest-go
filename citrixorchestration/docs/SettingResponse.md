# SettingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyGuid** | **string** | The policy GUID. | 
**SettingGuid** | **string** | The setting GUID. | 
**SettingName** | Pointer to **string** | Setting name. Is globally unique. | [optional] 
**UseDefault** | **bool** | Indicate if the default setting value is used. If the setting data type is bool, this is ignored. For all other settings, if this value is true, the setting value is ignored and if this value is false, the setting value is used. | 
**SettingValue** | Pointer to **string** | The current setting value. For boolean types, the returned value is \&quot;True\&quot;/\&quot;False\&quot;. The caller should assume the value is case-insensitive. | [optional] 

## Methods

### NewSettingResponse

`func NewSettingResponse(policyGuid string, settingGuid string, useDefault bool, ) *SettingResponse`

NewSettingResponse instantiates a new SettingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingResponseWithDefaults

`func NewSettingResponseWithDefaults() *SettingResponse`

NewSettingResponseWithDefaults instantiates a new SettingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyGuid

`func (o *SettingResponse) GetPolicyGuid() string`

GetPolicyGuid returns the PolicyGuid field if non-nil, zero value otherwise.

### GetPolicyGuidOk

`func (o *SettingResponse) GetPolicyGuidOk() (*string, bool)`

GetPolicyGuidOk returns a tuple with the PolicyGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyGuid

`func (o *SettingResponse) SetPolicyGuid(v string)`

SetPolicyGuid sets PolicyGuid field to given value.


### GetSettingGuid

`func (o *SettingResponse) GetSettingGuid() string`

GetSettingGuid returns the SettingGuid field if non-nil, zero value otherwise.

### GetSettingGuidOk

`func (o *SettingResponse) GetSettingGuidOk() (*string, bool)`

GetSettingGuidOk returns a tuple with the SettingGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingGuid

`func (o *SettingResponse) SetSettingGuid(v string)`

SetSettingGuid sets SettingGuid field to given value.


### GetSettingName

`func (o *SettingResponse) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *SettingResponse) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *SettingResponse) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *SettingResponse) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### GetUseDefault

`func (o *SettingResponse) GetUseDefault() bool`

GetUseDefault returns the UseDefault field if non-nil, zero value otherwise.

### GetUseDefaultOk

`func (o *SettingResponse) GetUseDefaultOk() (*bool, bool)`

GetUseDefaultOk returns a tuple with the UseDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefault

`func (o *SettingResponse) SetUseDefault(v bool)`

SetUseDefault sets UseDefault field to given value.


### GetSettingValue

`func (o *SettingResponse) GetSettingValue() string`

GetSettingValue returns the SettingValue field if non-nil, zero value otherwise.

### GetSettingValueOk

`func (o *SettingResponse) GetSettingValueOk() (*string, bool)`

GetSettingValueOk returns a tuple with the SettingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingValue

`func (o *SettingResponse) SetSettingValue(v string)`

SetSettingValue sets SettingValue field to given value.

### HasSettingValue

`func (o *SettingResponse) HasSettingValue() bool`

HasSettingValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


