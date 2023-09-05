# AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ADAccountName** | **string** | The Active Directory account name to be imported. | 
**ResetPassword** | Pointer to **bool** | Indicates whether the machine account password should be reset. | [optional] [default to true]
**Password** | Pointer to **string** | The current password for the machine account , in the format specified by PasswordFormat. | [optional] 
**PasswordFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 

## Methods

### NewAddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount

`func NewAddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount(aDAccountName string, ) *AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount`

NewAddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount instantiates a new AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccountWithDefaults

`func NewAddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccountWithDefaults() *AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount`

NewAddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccountWithDefaults instantiates a new AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetADAccountName

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount) GetADAccountName() string`

GetADAccountName returns the ADAccountName field if non-nil, zero value otherwise.

### GetADAccountNameOk

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount) GetADAccountNameOk() (*string, bool)`

GetADAccountNameOk returns a tuple with the ADAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADAccountName

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount) SetADAccountName(v string)`

SetADAccountName sets ADAccountName field to given value.


### GetResetPassword

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount) GetResetPassword() bool`

GetResetPassword returns the ResetPassword field if non-nil, zero value otherwise.

### GetResetPasswordOk

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount) GetResetPasswordOk() (*bool, bool)`

GetResetPasswordOk returns a tuple with the ResetPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPassword

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount) SetResetPassword(v bool)`

SetResetPassword sets ResetPassword field to given value.

### HasResetPassword

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount) HasResetPassword() bool`

HasResetPassword returns a boolean if a field has been set.

### GetPassword

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordFormat

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount) GetPasswordFormat() IdentityPasswordFormat`

GetPasswordFormat returns the PasswordFormat field if non-nil, zero value otherwise.

### GetPasswordFormatOk

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount) GetPasswordFormatOk() (*IdentityPasswordFormat, bool)`

GetPasswordFormatOk returns a tuple with the PasswordFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFormat

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount) SetPasswordFormat(v IdentityPasswordFormat)`

SetPasswordFormat sets PasswordFormat field to given value.

### HasPasswordFormat

`func (o *AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount) HasPasswordFormat() bool`

HasPasswordFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


