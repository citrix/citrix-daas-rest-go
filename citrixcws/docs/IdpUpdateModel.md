# IdpUpdateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdOtpDetails** | Pointer to [**AdOtpConnectionModel**](AdOtpConnectionModel.md) |  | [optional] 
**AuthDomainName** | Pointer to **NullableString** |  | [optional] 
**Nickname** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIdpUpdateModel

`func NewIdpUpdateModel() *IdpUpdateModel`

NewIdpUpdateModel instantiates a new IdpUpdateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpUpdateModelWithDefaults

`func NewIdpUpdateModelWithDefaults() *IdpUpdateModel`

NewIdpUpdateModelWithDefaults instantiates a new IdpUpdateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdOtpDetails

`func (o *IdpUpdateModel) GetAdOtpDetails() AdOtpConnectionModel`

GetAdOtpDetails returns the AdOtpDetails field if non-nil, zero value otherwise.

### GetAdOtpDetailsOk

`func (o *IdpUpdateModel) GetAdOtpDetailsOk() (*AdOtpConnectionModel, bool)`

GetAdOtpDetailsOk returns a tuple with the AdOtpDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdOtpDetails

`func (o *IdpUpdateModel) SetAdOtpDetails(v AdOtpConnectionModel)`

SetAdOtpDetails sets AdOtpDetails field to given value.

### HasAdOtpDetails

`func (o *IdpUpdateModel) HasAdOtpDetails() bool`

HasAdOtpDetails returns a boolean if a field has been set.

### GetAuthDomainName

`func (o *IdpUpdateModel) GetAuthDomainName() string`

GetAuthDomainName returns the AuthDomainName field if non-nil, zero value otherwise.

### GetAuthDomainNameOk

`func (o *IdpUpdateModel) GetAuthDomainNameOk() (*string, bool)`

GetAuthDomainNameOk returns a tuple with the AuthDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDomainName

`func (o *IdpUpdateModel) SetAuthDomainName(v string)`

SetAuthDomainName sets AuthDomainName field to given value.

### HasAuthDomainName

`func (o *IdpUpdateModel) HasAuthDomainName() bool`

HasAuthDomainName returns a boolean if a field has been set.

### SetAuthDomainNameNil

`func (o *IdpUpdateModel) SetAuthDomainNameNil(b bool)`

 SetAuthDomainNameNil sets the value for AuthDomainName to be an explicit nil

### UnsetAuthDomainName
`func (o *IdpUpdateModel) UnsetAuthDomainName()`

UnsetAuthDomainName ensures that no value is present for AuthDomainName, not even an explicit nil
### GetNickname

`func (o *IdpUpdateModel) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *IdpUpdateModel) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *IdpUpdateModel) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *IdpUpdateModel) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### SetNicknameNil

`func (o *IdpUpdateModel) SetNicknameNil(b bool)`

 SetNicknameNil sets the value for Nickname to be an explicit nil

### UnsetNickname
`func (o *IdpUpdateModel) UnsetNickname()`

UnsetNickname ensures that no value is present for Nickname, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


