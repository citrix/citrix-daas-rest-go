# StoreFrontAutenticationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreId** | Pointer to **NullableString** |  | [optional] 
**IdpType** | Pointer to **NullableString** |  | [optional] 
**IdpConfigId** | Pointer to **NullableString** |  | [optional] 
**StoreFrontUrl** | Pointer to **NullableString** |  | [optional] 
**CitrixManaged** | Pointer to **bool** |  | [optional] 
**WorkspaceAdministrator** | Pointer to **bool** |  | [optional] 
**LibraryAdministrator** | Pointer to **bool** |  | [optional] 
**CspCustomer** | Pointer to **NullableString** |  | [optional] 
**Tenants** | Pointer to [**[]StoreFrontAutenticationModel**](StoreFrontAutenticationModel.md) |  | [optional] 

## Methods

### NewStoreFrontAutenticationModel

`func NewStoreFrontAutenticationModel() *StoreFrontAutenticationModel`

NewStoreFrontAutenticationModel instantiates a new StoreFrontAutenticationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreFrontAutenticationModelWithDefaults

`func NewStoreFrontAutenticationModelWithDefaults() *StoreFrontAutenticationModel`

NewStoreFrontAutenticationModelWithDefaults instantiates a new StoreFrontAutenticationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreId

`func (o *StoreFrontAutenticationModel) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *StoreFrontAutenticationModel) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *StoreFrontAutenticationModel) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *StoreFrontAutenticationModel) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### SetStoreIdNil

`func (o *StoreFrontAutenticationModel) SetStoreIdNil(b bool)`

 SetStoreIdNil sets the value for StoreId to be an explicit nil

### UnsetStoreId
`func (o *StoreFrontAutenticationModel) UnsetStoreId()`

UnsetStoreId ensures that no value is present for StoreId, not even an explicit nil
### GetIdpType

`func (o *StoreFrontAutenticationModel) GetIdpType() string`

GetIdpType returns the IdpType field if non-nil, zero value otherwise.

### GetIdpTypeOk

`func (o *StoreFrontAutenticationModel) GetIdpTypeOk() (*string, bool)`

GetIdpTypeOk returns a tuple with the IdpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpType

`func (o *StoreFrontAutenticationModel) SetIdpType(v string)`

SetIdpType sets IdpType field to given value.

### HasIdpType

`func (o *StoreFrontAutenticationModel) HasIdpType() bool`

HasIdpType returns a boolean if a field has been set.

### SetIdpTypeNil

`func (o *StoreFrontAutenticationModel) SetIdpTypeNil(b bool)`

 SetIdpTypeNil sets the value for IdpType to be an explicit nil

### UnsetIdpType
`func (o *StoreFrontAutenticationModel) UnsetIdpType()`

UnsetIdpType ensures that no value is present for IdpType, not even an explicit nil
### GetIdpConfigId

`func (o *StoreFrontAutenticationModel) GetIdpConfigId() string`

GetIdpConfigId returns the IdpConfigId field if non-nil, zero value otherwise.

### GetIdpConfigIdOk

`func (o *StoreFrontAutenticationModel) GetIdpConfigIdOk() (*string, bool)`

GetIdpConfigIdOk returns a tuple with the IdpConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpConfigId

`func (o *StoreFrontAutenticationModel) SetIdpConfigId(v string)`

SetIdpConfigId sets IdpConfigId field to given value.

### HasIdpConfigId

`func (o *StoreFrontAutenticationModel) HasIdpConfigId() bool`

HasIdpConfigId returns a boolean if a field has been set.

### SetIdpConfigIdNil

`func (o *StoreFrontAutenticationModel) SetIdpConfigIdNil(b bool)`

 SetIdpConfigIdNil sets the value for IdpConfigId to be an explicit nil

### UnsetIdpConfigId
`func (o *StoreFrontAutenticationModel) UnsetIdpConfigId()`

UnsetIdpConfigId ensures that no value is present for IdpConfigId, not even an explicit nil
### GetStoreFrontUrl

`func (o *StoreFrontAutenticationModel) GetStoreFrontUrl() string`

GetStoreFrontUrl returns the StoreFrontUrl field if non-nil, zero value otherwise.

### GetStoreFrontUrlOk

`func (o *StoreFrontAutenticationModel) GetStoreFrontUrlOk() (*string, bool)`

GetStoreFrontUrlOk returns a tuple with the StoreFrontUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreFrontUrl

`func (o *StoreFrontAutenticationModel) SetStoreFrontUrl(v string)`

SetStoreFrontUrl sets StoreFrontUrl field to given value.

### HasStoreFrontUrl

`func (o *StoreFrontAutenticationModel) HasStoreFrontUrl() bool`

HasStoreFrontUrl returns a boolean if a field has been set.

### SetStoreFrontUrlNil

`func (o *StoreFrontAutenticationModel) SetStoreFrontUrlNil(b bool)`

 SetStoreFrontUrlNil sets the value for StoreFrontUrl to be an explicit nil

### UnsetStoreFrontUrl
`func (o *StoreFrontAutenticationModel) UnsetStoreFrontUrl()`

UnsetStoreFrontUrl ensures that no value is present for StoreFrontUrl, not even an explicit nil
### GetCitrixManaged

`func (o *StoreFrontAutenticationModel) GetCitrixManaged() bool`

GetCitrixManaged returns the CitrixManaged field if non-nil, zero value otherwise.

### GetCitrixManagedOk

`func (o *StoreFrontAutenticationModel) GetCitrixManagedOk() (*bool, bool)`

GetCitrixManagedOk returns a tuple with the CitrixManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitrixManaged

`func (o *StoreFrontAutenticationModel) SetCitrixManaged(v bool)`

SetCitrixManaged sets CitrixManaged field to given value.

### HasCitrixManaged

`func (o *StoreFrontAutenticationModel) HasCitrixManaged() bool`

HasCitrixManaged returns a boolean if a field has been set.

### GetWorkspaceAdministrator

`func (o *StoreFrontAutenticationModel) GetWorkspaceAdministrator() bool`

GetWorkspaceAdministrator returns the WorkspaceAdministrator field if non-nil, zero value otherwise.

### GetWorkspaceAdministratorOk

`func (o *StoreFrontAutenticationModel) GetWorkspaceAdministratorOk() (*bool, bool)`

GetWorkspaceAdministratorOk returns a tuple with the WorkspaceAdministrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceAdministrator

`func (o *StoreFrontAutenticationModel) SetWorkspaceAdministrator(v bool)`

SetWorkspaceAdministrator sets WorkspaceAdministrator field to given value.

### HasWorkspaceAdministrator

`func (o *StoreFrontAutenticationModel) HasWorkspaceAdministrator() bool`

HasWorkspaceAdministrator returns a boolean if a field has been set.

### GetLibraryAdministrator

`func (o *StoreFrontAutenticationModel) GetLibraryAdministrator() bool`

GetLibraryAdministrator returns the LibraryAdministrator field if non-nil, zero value otherwise.

### GetLibraryAdministratorOk

`func (o *StoreFrontAutenticationModel) GetLibraryAdministratorOk() (*bool, bool)`

GetLibraryAdministratorOk returns a tuple with the LibraryAdministrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryAdministrator

`func (o *StoreFrontAutenticationModel) SetLibraryAdministrator(v bool)`

SetLibraryAdministrator sets LibraryAdministrator field to given value.

### HasLibraryAdministrator

`func (o *StoreFrontAutenticationModel) HasLibraryAdministrator() bool`

HasLibraryAdministrator returns a boolean if a field has been set.

### GetCspCustomer

`func (o *StoreFrontAutenticationModel) GetCspCustomer() string`

GetCspCustomer returns the CspCustomer field if non-nil, zero value otherwise.

### GetCspCustomerOk

`func (o *StoreFrontAutenticationModel) GetCspCustomerOk() (*string, bool)`

GetCspCustomerOk returns a tuple with the CspCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomer

`func (o *StoreFrontAutenticationModel) SetCspCustomer(v string)`

SetCspCustomer sets CspCustomer field to given value.

### HasCspCustomer

`func (o *StoreFrontAutenticationModel) HasCspCustomer() bool`

HasCspCustomer returns a boolean if a field has been set.

### SetCspCustomerNil

`func (o *StoreFrontAutenticationModel) SetCspCustomerNil(b bool)`

 SetCspCustomerNil sets the value for CspCustomer to be an explicit nil

### UnsetCspCustomer
`func (o *StoreFrontAutenticationModel) UnsetCspCustomer()`

UnsetCspCustomer ensures that no value is present for CspCustomer, not even an explicit nil
### GetTenants

`func (o *StoreFrontAutenticationModel) GetTenants() []StoreFrontAutenticationModel`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *StoreFrontAutenticationModel) GetTenantsOk() (*[]StoreFrontAutenticationModel, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *StoreFrontAutenticationModel) SetTenants(v []StoreFrontAutenticationModel)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *StoreFrontAutenticationModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *StoreFrontAutenticationModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *StoreFrontAutenticationModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


