# CatalogDomainModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | **string** | Domain the VMs will join | 
**DomainOu** | Pointer to **string** | OU of the domain | [optional] 
**ServiceAccountName** | Pointer to **string** | Name of the service account that will be used to join the domain | [optional] 
**IsSecureBrowser** | Pointer to **bool** |  | [optional] 
**CspCustomer** | Pointer to **string** |  | [optional] 

## Methods

### NewCatalogDomainModel

`func NewCatalogDomainModel(domainName string, ) *CatalogDomainModel`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


