# IdentityUpdateMachineRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **NullableBool** | Enable or disable the machine account. | [optional] 
**OldPassword** | Pointer to **NullableString** | Old account password, used when changing the account password. | [optional] 
**NewPassword** | Pointer to **NullableString** | New account password, used when changing the account password. | [optional] 
**AccountPasswordFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 

## Methods

### NewIdentityUpdateMachineRequestModel

`func NewIdentityUpdateMachineRequestModel() *IdentityUpdateMachineRequestModel`

NewIdentityUpdateMachineRequestModel instantiates a new IdentityUpdateMachineRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityUpdateMachineRequestModelWithDefaults

`func NewIdentityUpdateMachineRequestModelWithDefaults() *IdentityUpdateMachineRequestModel`

NewIdentityUpdateMachineRequestModelWithDefaults instantiates a new IdentityUpdateMachineRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *IdentityUpdateMachineRequestModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdentityUpdateMachineRequestModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdentityUpdateMachineRequestModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IdentityUpdateMachineRequestModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *IdentityUpdateMachineRequestModel) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *IdentityUpdateMachineRequestModel) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetOldPassword

`func (o *IdentityUpdateMachineRequestModel) GetOldPassword() string`

GetOldPassword returns the OldPassword field if non-nil, zero value otherwise.

### GetOldPasswordOk

`func (o *IdentityUpdateMachineRequestModel) GetOldPasswordOk() (*string, bool)`

GetOldPasswordOk returns a tuple with the OldPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPassword

`func (o *IdentityUpdateMachineRequestModel) SetOldPassword(v string)`

SetOldPassword sets OldPassword field to given value.

### HasOldPassword

`func (o *IdentityUpdateMachineRequestModel) HasOldPassword() bool`

HasOldPassword returns a boolean if a field has been set.

### SetOldPasswordNil

`func (o *IdentityUpdateMachineRequestModel) SetOldPasswordNil(b bool)`

 SetOldPasswordNil sets the value for OldPassword to be an explicit nil

### UnsetOldPassword
`func (o *IdentityUpdateMachineRequestModel) UnsetOldPassword()`

UnsetOldPassword ensures that no value is present for OldPassword, not even an explicit nil
### GetNewPassword

`func (o *IdentityUpdateMachineRequestModel) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *IdentityUpdateMachineRequestModel) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *IdentityUpdateMachineRequestModel) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *IdentityUpdateMachineRequestModel) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### SetNewPasswordNil

`func (o *IdentityUpdateMachineRequestModel) SetNewPasswordNil(b bool)`

 SetNewPasswordNil sets the value for NewPassword to be an explicit nil

### UnsetNewPassword
`func (o *IdentityUpdateMachineRequestModel) UnsetNewPassword()`

UnsetNewPassword ensures that no value is present for NewPassword, not even an explicit nil
### GetAccountPasswordFormat

`func (o *IdentityUpdateMachineRequestModel) GetAccountPasswordFormat() IdentityPasswordFormat`

GetAccountPasswordFormat returns the AccountPasswordFormat field if non-nil, zero value otherwise.

### GetAccountPasswordFormatOk

`func (o *IdentityUpdateMachineRequestModel) GetAccountPasswordFormatOk() (*IdentityPasswordFormat, bool)`

GetAccountPasswordFormatOk returns a tuple with the AccountPasswordFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPasswordFormat

`func (o *IdentityUpdateMachineRequestModel) SetAccountPasswordFormat(v IdentityPasswordFormat)`

SetAccountPasswordFormat sets AccountPasswordFormat field to given value.

### HasAccountPasswordFormat

`func (o *IdentityUpdateMachineRequestModel) HasAccountPasswordFormat() bool`

HasAccountPasswordFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


