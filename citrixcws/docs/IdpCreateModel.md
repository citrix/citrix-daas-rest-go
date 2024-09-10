# IdpCreateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityProviderId** | **string** |  | 
**IdentityProviderNickname** | **string** |  | 
**IdpConnectionModel** | Pointer to [**IdpConnectionModel**](IdpConnectionModel.md) |  | [optional] 

## Methods

### NewIdpCreateModel

`func NewIdpCreateModel(identityProviderId string, identityProviderNickname string, ) *IdpCreateModel`

NewIdpCreateModel instantiates a new IdpCreateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpCreateModelWithDefaults

`func NewIdpCreateModelWithDefaults() *IdpCreateModel`

NewIdpCreateModelWithDefaults instantiates a new IdpCreateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityProviderId

`func (o *IdpCreateModel) GetIdentityProviderId() string`

GetIdentityProviderId returns the IdentityProviderId field if non-nil, zero value otherwise.

### GetIdentityProviderIdOk

`func (o *IdpCreateModel) GetIdentityProviderIdOk() (*string, bool)`

GetIdentityProviderIdOk returns a tuple with the IdentityProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderId

`func (o *IdpCreateModel) SetIdentityProviderId(v string)`

SetIdentityProviderId sets IdentityProviderId field to given value.


### GetIdentityProviderNickname

`func (o *IdpCreateModel) GetIdentityProviderNickname() string`

GetIdentityProviderNickname returns the IdentityProviderNickname field if non-nil, zero value otherwise.

### GetIdentityProviderNicknameOk

`func (o *IdpCreateModel) GetIdentityProviderNicknameOk() (*string, bool)`

GetIdentityProviderNicknameOk returns a tuple with the IdentityProviderNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderNickname

`func (o *IdpCreateModel) SetIdentityProviderNickname(v string)`

SetIdentityProviderNickname sets IdentityProviderNickname field to given value.


### GetIdpConnectionModel

`func (o *IdpCreateModel) GetIdpConnectionModel() IdpConnectionModel`

GetIdpConnectionModel returns the IdpConnectionModel field if non-nil, zero value otherwise.

### GetIdpConnectionModelOk

`func (o *IdpCreateModel) GetIdpConnectionModelOk() (*IdpConnectionModel, bool)`

GetIdpConnectionModelOk returns a tuple with the IdpConnectionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpConnectionModel

`func (o *IdpCreateModel) SetIdpConnectionModel(v IdpConnectionModel)`

SetIdpConnectionModel sets IdpConnectionModel field to given value.

### HasIdpConnectionModel

`func (o *IdpCreateModel) HasIdpConnectionModel() bool`

HasIdpConnectionModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


