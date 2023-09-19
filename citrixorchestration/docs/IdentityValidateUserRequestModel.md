# IdentityValidateUserRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Forest** | Pointer to **NullableString** | Forest in which the user account resides. | [optional] 
**Domain** | Pointer to **NullableString** | Domain in which the user account resides. | [optional] 
**LogOnType** | Pointer to [**IdentityLogonType**](IdentityLogonType.md) |  | [optional] 
**UserName** | **string** | Username to validate.  Required. | 
**Password** | **string** | Password to validate.  Required. Must be specified in the format indicated by PasswordFormat. | 
**PasswordFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 

## Methods

### NewIdentityValidateUserRequestModel

`func NewIdentityValidateUserRequestModel(userName string, password string, ) *IdentityValidateUserRequestModel`

NewIdentityValidateUserRequestModel instantiates a new IdentityValidateUserRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityValidateUserRequestModelWithDefaults

`func NewIdentityValidateUserRequestModelWithDefaults() *IdentityValidateUserRequestModel`

NewIdentityValidateUserRequestModelWithDefaults instantiates a new IdentityValidateUserRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForest

`func (o *IdentityValidateUserRequestModel) GetForest() string`

GetForest returns the Forest field if non-nil, zero value otherwise.

### GetForestOk

`func (o *IdentityValidateUserRequestModel) GetForestOk() (*string, bool)`

GetForestOk returns a tuple with the Forest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForest

`func (o *IdentityValidateUserRequestModel) SetForest(v string)`

SetForest sets Forest field to given value.

### HasForest

`func (o *IdentityValidateUserRequestModel) HasForest() bool`

HasForest returns a boolean if a field has been set.

### SetForestNil

`func (o *IdentityValidateUserRequestModel) SetForestNil(b bool)`

 SetForestNil sets the value for Forest to be an explicit nil

### UnsetForest
`func (o *IdentityValidateUserRequestModel) UnsetForest()`

UnsetForest ensures that no value is present for Forest, not even an explicit nil
### GetDomain

`func (o *IdentityValidateUserRequestModel) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *IdentityValidateUserRequestModel) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *IdentityValidateUserRequestModel) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *IdentityValidateUserRequestModel) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *IdentityValidateUserRequestModel) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *IdentityValidateUserRequestModel) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetLogOnType

`func (o *IdentityValidateUserRequestModel) GetLogOnType() IdentityLogonType`

GetLogOnType returns the LogOnType field if non-nil, zero value otherwise.

### GetLogOnTypeOk

`func (o *IdentityValidateUserRequestModel) GetLogOnTypeOk() (*IdentityLogonType, bool)`

GetLogOnTypeOk returns a tuple with the LogOnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogOnType

`func (o *IdentityValidateUserRequestModel) SetLogOnType(v IdentityLogonType)`

SetLogOnType sets LogOnType field to given value.

### HasLogOnType

`func (o *IdentityValidateUserRequestModel) HasLogOnType() bool`

HasLogOnType returns a boolean if a field has been set.

### GetUserName

`func (o *IdentityValidateUserRequestModel) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *IdentityValidateUserRequestModel) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *IdentityValidateUserRequestModel) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetPassword

`func (o *IdentityValidateUserRequestModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IdentityValidateUserRequestModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IdentityValidateUserRequestModel) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPasswordFormat

`func (o *IdentityValidateUserRequestModel) GetPasswordFormat() IdentityPasswordFormat`

GetPasswordFormat returns the PasswordFormat field if non-nil, zero value otherwise.

### GetPasswordFormatOk

`func (o *IdentityValidateUserRequestModel) GetPasswordFormatOk() (*IdentityPasswordFormat, bool)`

GetPasswordFormatOk returns a tuple with the PasswordFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFormat

`func (o *IdentityValidateUserRequestModel) SetPasswordFormat(v IdentityPasswordFormat)`

SetPasswordFormat sets PasswordFormat field to given value.

### HasPasswordFormat

`func (o *IdentityValidateUserRequestModel) HasPasswordFormat() bool`

HasPasswordFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


