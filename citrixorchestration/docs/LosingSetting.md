# LosingSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyIdentity** | Pointer to [**PolicyIdentity**](PolicyIdentity.md) |  | [optional] 
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

### GetPolicyIdentity

`func (o *LosingSetting) GetPolicyIdentity() PolicyIdentity`

GetPolicyIdentity returns the PolicyIdentity field if non-nil, zero value otherwise.

### GetPolicyIdentityOk

`func (o *LosingSetting) GetPolicyIdentityOk() (*PolicyIdentity, bool)`

GetPolicyIdentityOk returns a tuple with the PolicyIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyIdentity

`func (o *LosingSetting) SetPolicyIdentity(v PolicyIdentity)`

SetPolicyIdentity sets PolicyIdentity field to given value.

### HasPolicyIdentity

`func (o *LosingSetting) HasPolicyIdentity() bool`

HasPolicyIdentity returns a boolean if a field has been set.

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


