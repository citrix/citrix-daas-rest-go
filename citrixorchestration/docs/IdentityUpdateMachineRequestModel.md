# IdentityUpdateMachineRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable or disable the machine account. | [optional] 
**OldPassword** | Pointer to **string** | Old account password, used when changing the account password. | [optional] 
**NewPassword** | Pointer to **string** | New account password, used when changing the account password. | [optional] 
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


