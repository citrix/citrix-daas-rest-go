# IdentityUpdateUserRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OldPassword** | Pointer to **string** | Old account password, used when changing the account password. | [optional] 
**NewPassword** | Pointer to **string** | New account password, used when changing the account password. | [optional] 
**AccountPasswordFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 

## Methods

### NewIdentityUpdateUserRequestModel

`func NewIdentityUpdateUserRequestModel() *IdentityUpdateUserRequestModel`

NewIdentityUpdateUserRequestModel instantiates a new IdentityUpdateUserRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityUpdateUserRequestModelWithDefaults

`func NewIdentityUpdateUserRequestModelWithDefaults() *IdentityUpdateUserRequestModel`

NewIdentityUpdateUserRequestModelWithDefaults instantiates a new IdentityUpdateUserRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOldPassword

`func (o *IdentityUpdateUserRequestModel) GetOldPassword() string`

GetOldPassword returns the OldPassword field if non-nil, zero value otherwise.

### GetOldPasswordOk

`func (o *IdentityUpdateUserRequestModel) GetOldPasswordOk() (*string, bool)`

GetOldPasswordOk returns a tuple with the OldPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPassword

`func (o *IdentityUpdateUserRequestModel) SetOldPassword(v string)`

SetOldPassword sets OldPassword field to given value.

### HasOldPassword

`func (o *IdentityUpdateUserRequestModel) HasOldPassword() bool`

HasOldPassword returns a boolean if a field has been set.

### GetNewPassword

`func (o *IdentityUpdateUserRequestModel) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *IdentityUpdateUserRequestModel) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *IdentityUpdateUserRequestModel) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *IdentityUpdateUserRequestModel) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### GetAccountPasswordFormat

`func (o *IdentityUpdateUserRequestModel) GetAccountPasswordFormat() IdentityPasswordFormat`

GetAccountPasswordFormat returns the AccountPasswordFormat field if non-nil, zero value otherwise.

### GetAccountPasswordFormatOk

`func (o *IdentityUpdateUserRequestModel) GetAccountPasswordFormatOk() (*IdentityPasswordFormat, bool)`

GetAccountPasswordFormatOk returns a tuple with the AccountPasswordFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPasswordFormat

`func (o *IdentityUpdateUserRequestModel) SetAccountPasswordFormat(v IdentityPasswordFormat)`

SetAccountPasswordFormat sets AccountPasswordFormat field to given value.

### HasAccountPasswordFormat

`func (o *IdentityUpdateUserRequestModel) HasAccountPasswordFormat() bool`

HasAccountPasswordFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


