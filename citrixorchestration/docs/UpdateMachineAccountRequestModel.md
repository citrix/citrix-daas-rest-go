# UpdateMachineAccountRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResetPassword** | Pointer to **bool** | Indicates whether the machine account password should be reset. | [optional] 
**Password** | Pointer to **string** | The current password for the machine account, in the format specified by PasswordFormat. | [optional] 
**PasswordFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 

## Methods

### NewUpdateMachineAccountRequestModel

`func NewUpdateMachineAccountRequestModel() *UpdateMachineAccountRequestModel`

NewUpdateMachineAccountRequestModel instantiates a new UpdateMachineAccountRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMachineAccountRequestModelWithDefaults

`func NewUpdateMachineAccountRequestModelWithDefaults() *UpdateMachineAccountRequestModel`

NewUpdateMachineAccountRequestModelWithDefaults instantiates a new UpdateMachineAccountRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResetPassword

`func (o *UpdateMachineAccountRequestModel) GetResetPassword() bool`

GetResetPassword returns the ResetPassword field if non-nil, zero value otherwise.

### GetResetPasswordOk

`func (o *UpdateMachineAccountRequestModel) GetResetPasswordOk() (*bool, bool)`

GetResetPasswordOk returns a tuple with the ResetPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPassword

`func (o *UpdateMachineAccountRequestModel) SetResetPassword(v bool)`

SetResetPassword sets ResetPassword field to given value.

### HasResetPassword

`func (o *UpdateMachineAccountRequestModel) HasResetPassword() bool`

HasResetPassword returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateMachineAccountRequestModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateMachineAccountRequestModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateMachineAccountRequestModel) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateMachineAccountRequestModel) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordFormat

`func (o *UpdateMachineAccountRequestModel) GetPasswordFormat() IdentityPasswordFormat`

GetPasswordFormat returns the PasswordFormat field if non-nil, zero value otherwise.

### GetPasswordFormatOk

`func (o *UpdateMachineAccountRequestModel) GetPasswordFormatOk() (*IdentityPasswordFormat, bool)`

GetPasswordFormatOk returns a tuple with the PasswordFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFormat

`func (o *UpdateMachineAccountRequestModel) SetPasswordFormat(v IdentityPasswordFormat)`

SetPasswordFormat sets PasswordFormat field to given value.

### HasPasswordFormat

`func (o *UpdateMachineAccountRequestModel) HasPasswordFormat() bool`

HasPasswordFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


