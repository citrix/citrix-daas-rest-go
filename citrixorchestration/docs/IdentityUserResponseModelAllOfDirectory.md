# IdentityUserResponseModelAllOfDirectory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Forest** | Pointer to **string** | Forest of the directory. | [optional] 
**Domain** | Pointer to **string** | Domain of the directory. | [optional] 
**Tenant** | Pointer to **string** | Tenant of the directory. | [optional] 
**IdentityProvider** | Pointer to [**IdentityProviderType**](IdentityProviderType.md) |  | [optional] 

## Methods

### NewIdentityUserResponseModelAllOfDirectory

`func NewIdentityUserResponseModelAllOfDirectory() *IdentityUserResponseModelAllOfDirectory`

NewIdentityUserResponseModelAllOfDirectory instantiates a new IdentityUserResponseModelAllOfDirectory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityUserResponseModelAllOfDirectoryWithDefaults

`func NewIdentityUserResponseModelAllOfDirectoryWithDefaults() *IdentityUserResponseModelAllOfDirectory`

NewIdentityUserResponseModelAllOfDirectoryWithDefaults instantiates a new IdentityUserResponseModelAllOfDirectory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForest

`func (o *IdentityUserResponseModelAllOfDirectory) GetForest() string`

GetForest returns the Forest field if non-nil, zero value otherwise.

### GetForestOk

`func (o *IdentityUserResponseModelAllOfDirectory) GetForestOk() (*string, bool)`

GetForestOk returns a tuple with the Forest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForest

`func (o *IdentityUserResponseModelAllOfDirectory) SetForest(v string)`

SetForest sets Forest field to given value.

### HasForest

`func (o *IdentityUserResponseModelAllOfDirectory) HasForest() bool`

HasForest returns a boolean if a field has been set.

### GetDomain

`func (o *IdentityUserResponseModelAllOfDirectory) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *IdentityUserResponseModelAllOfDirectory) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *IdentityUserResponseModelAllOfDirectory) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *IdentityUserResponseModelAllOfDirectory) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetTenant

`func (o *IdentityUserResponseModelAllOfDirectory) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *IdentityUserResponseModelAllOfDirectory) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *IdentityUserResponseModelAllOfDirectory) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *IdentityUserResponseModelAllOfDirectory) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetIdentityProvider

`func (o *IdentityUserResponseModelAllOfDirectory) GetIdentityProvider() IdentityProviderType`

GetIdentityProvider returns the IdentityProvider field if non-nil, zero value otherwise.

### GetIdentityProviderOk

`func (o *IdentityUserResponseModelAllOfDirectory) GetIdentityProviderOk() (*IdentityProviderType, bool)`

GetIdentityProviderOk returns a tuple with the IdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProvider

`func (o *IdentityUserResponseModelAllOfDirectory) SetIdentityProvider(v IdentityProviderType)`

SetIdentityProvider sets IdentityProvider field to given value.

### HasIdentityProvider

`func (o *IdentityUserResponseModelAllOfDirectory) HasIdentityProvider() bool`

HasIdentityProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


