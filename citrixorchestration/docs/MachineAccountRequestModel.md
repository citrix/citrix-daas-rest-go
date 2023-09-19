# MachineAccountRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ADAccountName** | **string** | The Active Directory account name to be imported. | 
**ResetPassword** | Pointer to **NullableBool** | Indicates whether the machine account password should be reset. | [optional] [default to true]
**Password** | Pointer to **NullableString** | The current password for the machine account , in the format specified by PasswordFormat. | [optional] 
**PasswordFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 

## Methods

### NewMachineAccountRequestModel

`func NewMachineAccountRequestModel(aDAccountName string, ) *MachineAccountRequestModel`

NewMachineAccountRequestModel instantiates a new MachineAccountRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineAccountRequestModelWithDefaults

`func NewMachineAccountRequestModelWithDefaults() *MachineAccountRequestModel`

NewMachineAccountRequestModelWithDefaults instantiates a new MachineAccountRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetADAccountName

`func (o *MachineAccountRequestModel) GetADAccountName() string`

GetADAccountName returns the ADAccountName field if non-nil, zero value otherwise.

### GetADAccountNameOk

`func (o *MachineAccountRequestModel) GetADAccountNameOk() (*string, bool)`

GetADAccountNameOk returns a tuple with the ADAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADAccountName

`func (o *MachineAccountRequestModel) SetADAccountName(v string)`

SetADAccountName sets ADAccountName field to given value.


### GetResetPassword

`func (o *MachineAccountRequestModel) GetResetPassword() bool`

GetResetPassword returns the ResetPassword field if non-nil, zero value otherwise.

### GetResetPasswordOk

`func (o *MachineAccountRequestModel) GetResetPasswordOk() (*bool, bool)`

GetResetPasswordOk returns a tuple with the ResetPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPassword

`func (o *MachineAccountRequestModel) SetResetPassword(v bool)`

SetResetPassword sets ResetPassword field to given value.

### HasResetPassword

`func (o *MachineAccountRequestModel) HasResetPassword() bool`

HasResetPassword returns a boolean if a field has been set.

### SetResetPasswordNil

`func (o *MachineAccountRequestModel) SetResetPasswordNil(b bool)`

 SetResetPasswordNil sets the value for ResetPassword to be an explicit nil

### UnsetResetPassword
`func (o *MachineAccountRequestModel) UnsetResetPassword()`

UnsetResetPassword ensures that no value is present for ResetPassword, not even an explicit nil
### GetPassword

`func (o *MachineAccountRequestModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MachineAccountRequestModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MachineAccountRequestModel) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MachineAccountRequestModel) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *MachineAccountRequestModel) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *MachineAccountRequestModel) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordFormat

`func (o *MachineAccountRequestModel) GetPasswordFormat() IdentityPasswordFormat`

GetPasswordFormat returns the PasswordFormat field if non-nil, zero value otherwise.

### GetPasswordFormatOk

`func (o *MachineAccountRequestModel) GetPasswordFormatOk() (*IdentityPasswordFormat, bool)`

GetPasswordFormatOk returns a tuple with the PasswordFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFormat

`func (o *MachineAccountRequestModel) SetPasswordFormat(v IdentityPasswordFormat)`

SetPasswordFormat sets PasswordFormat field to given value.

### HasPasswordFormat

`func (o *MachineAccountRequestModel) HasPasswordFormat() bool`

HasPasswordFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


