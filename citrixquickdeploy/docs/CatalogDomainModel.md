# CatalogDomainModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **NullableString** | Domain the VMs will join | [optional] 
**DomainOu** | Pointer to **NullableString** | OU of the domain | [optional] 
**ServiceAccountName** | Pointer to **NullableString** | Name of the service account that will be used to join the domain | [optional] 
**ServiceAccountUid** | Pointer to **NullableString** | Service account to associate to the IdentityPool.  Used for Pure Entra ID joined catalogs. | [optional] 
**IsSecureBrowser** | Pointer to **bool** |  | [optional] 
**CspCustomer** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCatalogDomainModel

`func NewCatalogDomainModel() *CatalogDomainModel`

NewCatalogDomainModel instantiates a new CatalogDomainModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogDomainModelWithDefaults

`func NewCatalogDomainModelWithDefaults() *CatalogDomainModel`

NewCatalogDomainModelWithDefaults instantiates a new CatalogDomainModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *CatalogDomainModel) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *CatalogDomainModel) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *CatalogDomainModel) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *CatalogDomainModel) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *CatalogDomainModel) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *CatalogDomainModel) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetDomainOu

`func (o *CatalogDomainModel) GetDomainOu() string`

GetDomainOu returns the DomainOu field if non-nil, zero value otherwise.

### GetDomainOuOk

`func (o *CatalogDomainModel) GetDomainOuOk() (*string, bool)`

GetDomainOuOk returns a tuple with the DomainOu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainOu

`func (o *CatalogDomainModel) SetDomainOu(v string)`

SetDomainOu sets DomainOu field to given value.

### HasDomainOu

`func (o *CatalogDomainModel) HasDomainOu() bool`

HasDomainOu returns a boolean if a field has been set.

### SetDomainOuNil

`func (o *CatalogDomainModel) SetDomainOuNil(b bool)`

 SetDomainOuNil sets the value for DomainOu to be an explicit nil

### UnsetDomainOu
`func (o *CatalogDomainModel) UnsetDomainOu()`

UnsetDomainOu ensures that no value is present for DomainOu, not even an explicit nil
### GetServiceAccountName

`func (o *CatalogDomainModel) GetServiceAccountName() string`

GetServiceAccountName returns the ServiceAccountName field if non-nil, zero value otherwise.

### GetServiceAccountNameOk

`func (o *CatalogDomainModel) GetServiceAccountNameOk() (*string, bool)`

GetServiceAccountNameOk returns a tuple with the ServiceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountName

`func (o *CatalogDomainModel) SetServiceAccountName(v string)`

SetServiceAccountName sets ServiceAccountName field to given value.

### HasServiceAccountName

`func (o *CatalogDomainModel) HasServiceAccountName() bool`

HasServiceAccountName returns a boolean if a field has been set.

### SetServiceAccountNameNil

`func (o *CatalogDomainModel) SetServiceAccountNameNil(b bool)`

 SetServiceAccountNameNil sets the value for ServiceAccountName to be an explicit nil

### UnsetServiceAccountName
`func (o *CatalogDomainModel) UnsetServiceAccountName()`

UnsetServiceAccountName ensures that no value is present for ServiceAccountName, not even an explicit nil
### GetServiceAccountUid

`func (o *CatalogDomainModel) GetServiceAccountUid() string`

GetServiceAccountUid returns the ServiceAccountUid field if non-nil, zero value otherwise.

### GetServiceAccountUidOk

`func (o *CatalogDomainModel) GetServiceAccountUidOk() (*string, bool)`

GetServiceAccountUidOk returns a tuple with the ServiceAccountUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountUid

`func (o *CatalogDomainModel) SetServiceAccountUid(v string)`

SetServiceAccountUid sets ServiceAccountUid field to given value.

### HasServiceAccountUid

`func (o *CatalogDomainModel) HasServiceAccountUid() bool`

HasServiceAccountUid returns a boolean if a field has been set.

### SetServiceAccountUidNil

`func (o *CatalogDomainModel) SetServiceAccountUidNil(b bool)`

 SetServiceAccountUidNil sets the value for ServiceAccountUid to be an explicit nil

### UnsetServiceAccountUid
`func (o *CatalogDomainModel) UnsetServiceAccountUid()`

UnsetServiceAccountUid ensures that no value is present for ServiceAccountUid, not even an explicit nil
### GetIsSecureBrowser

`func (o *CatalogDomainModel) GetIsSecureBrowser() bool`

GetIsSecureBrowser returns the IsSecureBrowser field if non-nil, zero value otherwise.

### GetIsSecureBrowserOk

`func (o *CatalogDomainModel) GetIsSecureBrowserOk() (*bool, bool)`

GetIsSecureBrowserOk returns a tuple with the IsSecureBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecureBrowser

`func (o *CatalogDomainModel) SetIsSecureBrowser(v bool)`

SetIsSecureBrowser sets IsSecureBrowser field to given value.

### HasIsSecureBrowser

`func (o *CatalogDomainModel) HasIsSecureBrowser() bool`

HasIsSecureBrowser returns a boolean if a field has been set.

### GetCspCustomer

`func (o *CatalogDomainModel) GetCspCustomer() string`

GetCspCustomer returns the CspCustomer field if non-nil, zero value otherwise.

### GetCspCustomerOk

`func (o *CatalogDomainModel) GetCspCustomerOk() (*string, bool)`

GetCspCustomerOk returns a tuple with the CspCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomer

`func (o *CatalogDomainModel) SetCspCustomer(v string)`

SetCspCustomer sets CspCustomer field to given value.

### HasCspCustomer

`func (o *CatalogDomainModel) HasCspCustomer() bool`

HasCspCustomer returns a boolean if a field has been set.

### SetCspCustomerNil

`func (o *CatalogDomainModel) SetCspCustomerNil(b bool)`

 SetCspCustomerNil sets the value for CspCustomer to be an explicit nil

### UnsetCspCustomer
`func (o *CatalogDomainModel) UnsetCspCustomer()`

UnsetCspCustomer ensures that no value is present for CspCustomer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


