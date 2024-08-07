# IdentityDirectoryResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Forest** | Pointer to **NullableString** | Forest of the directory. | [optional] 
**Domain** | Pointer to **NullableString** | Domain of the directory. | [optional] 
**Tenant** | Pointer to **NullableString** | Tenant of the directory. | [optional] 
**IdentityProvider** | Pointer to [**IdentityProviderType**](IdentityProviderType.md) |  | [optional] 
**IdpInstanceId** | Pointer to **NullableString** | Instance of the directory. | [optional] 
**IdpNickName** | Pointer to **NullableString** | Nickname of the directory | [optional] 

## Methods

### NewIdentityDirectoryResponseModel

`func NewIdentityDirectoryResponseModel() *IdentityDirectoryResponseModel`

NewIdentityDirectoryResponseModel instantiates a new IdentityDirectoryResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityDirectoryResponseModelWithDefaults

`func NewIdentityDirectoryResponseModelWithDefaults() *IdentityDirectoryResponseModel`

NewIdentityDirectoryResponseModelWithDefaults instantiates a new IdentityDirectoryResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForest

`func (o *IdentityDirectoryResponseModel) GetForest() string`

GetForest returns the Forest field if non-nil, zero value otherwise.

### GetForestOk

`func (o *IdentityDirectoryResponseModel) GetForestOk() (*string, bool)`

GetForestOk returns a tuple with the Forest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForest

`func (o *IdentityDirectoryResponseModel) SetForest(v string)`

SetForest sets Forest field to given value.

### HasForest

`func (o *IdentityDirectoryResponseModel) HasForest() bool`

HasForest returns a boolean if a field has been set.

### SetForestNil

`func (o *IdentityDirectoryResponseModel) SetForestNil(b bool)`

 SetForestNil sets the value for Forest to be an explicit nil

### UnsetForest
`func (o *IdentityDirectoryResponseModel) UnsetForest()`

UnsetForest ensures that no value is present for Forest, not even an explicit nil
### GetDomain

`func (o *IdentityDirectoryResponseModel) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *IdentityDirectoryResponseModel) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *IdentityDirectoryResponseModel) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *IdentityDirectoryResponseModel) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *IdentityDirectoryResponseModel) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *IdentityDirectoryResponseModel) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetTenant

`func (o *IdentityDirectoryResponseModel) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *IdentityDirectoryResponseModel) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *IdentityDirectoryResponseModel) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *IdentityDirectoryResponseModel) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *IdentityDirectoryResponseModel) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *IdentityDirectoryResponseModel) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetIdentityProvider

`func (o *IdentityDirectoryResponseModel) GetIdentityProvider() IdentityProviderType`

GetIdentityProvider returns the IdentityProvider field if non-nil, zero value otherwise.

### GetIdentityProviderOk

`func (o *IdentityDirectoryResponseModel) GetIdentityProviderOk() (*IdentityProviderType, bool)`

GetIdentityProviderOk returns a tuple with the IdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProvider

`func (o *IdentityDirectoryResponseModel) SetIdentityProvider(v IdentityProviderType)`

SetIdentityProvider sets IdentityProvider field to given value.

### HasIdentityProvider

`func (o *IdentityDirectoryResponseModel) HasIdentityProvider() bool`

HasIdentityProvider returns a boolean if a field has been set.

### GetIdpInstanceId

`func (o *IdentityDirectoryResponseModel) GetIdpInstanceId() string`

GetIdpInstanceId returns the IdpInstanceId field if non-nil, zero value otherwise.

### GetIdpInstanceIdOk

`func (o *IdentityDirectoryResponseModel) GetIdpInstanceIdOk() (*string, bool)`

GetIdpInstanceIdOk returns a tuple with the IdpInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpInstanceId

`func (o *IdentityDirectoryResponseModel) SetIdpInstanceId(v string)`

SetIdpInstanceId sets IdpInstanceId field to given value.

### HasIdpInstanceId

`func (o *IdentityDirectoryResponseModel) HasIdpInstanceId() bool`

HasIdpInstanceId returns a boolean if a field has been set.

### SetIdpInstanceIdNil

`func (o *IdentityDirectoryResponseModel) SetIdpInstanceIdNil(b bool)`

 SetIdpInstanceIdNil sets the value for IdpInstanceId to be an explicit nil

### UnsetIdpInstanceId
`func (o *IdentityDirectoryResponseModel) UnsetIdpInstanceId()`

UnsetIdpInstanceId ensures that no value is present for IdpInstanceId, not even an explicit nil
### GetIdpNickName

`func (o *IdentityDirectoryResponseModel) GetIdpNickName() string`

GetIdpNickName returns the IdpNickName field if non-nil, zero value otherwise.

### GetIdpNickNameOk

`func (o *IdentityDirectoryResponseModel) GetIdpNickNameOk() (*string, bool)`

GetIdpNickNameOk returns a tuple with the IdpNickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpNickName

`func (o *IdentityDirectoryResponseModel) SetIdpNickName(v string)`

SetIdpNickName sets IdpNickName field to given value.

### HasIdpNickName

`func (o *IdentityDirectoryResponseModel) HasIdpNickName() bool`

HasIdpNickName returns a boolean if a field has been set.

### SetIdpNickNameNil

`func (o *IdentityDirectoryResponseModel) SetIdpNickNameNil(b bool)`

 SetIdpNickNameNil sets the value for IdpNickName to be an explicit nil

### UnsetIdpNickName
`func (o *IdentityDirectoryResponseModel) UnsetIdpNickName()`

UnsetIdpNickName ensures that no value is present for IdpNickName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


