# IdentityDirectoryResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Forest** | Pointer to **NullableString** | Forest of the directory. | [optional] 
**Domain** | Pointer to **NullableString** | Domain of the directory. | [optional] 
**Tenant** | Pointer to **NullableString** | Tenant of the directory. | [optional] 
**IdentityProvider** | Pointer to [**IdentityProviderType**](IdentityProviderType.md) |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


