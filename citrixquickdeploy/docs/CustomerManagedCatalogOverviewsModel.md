# CustomerManagedCatalogOverviewsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxCatalogVdas** | Pointer to **int32** | Deprecated, check CatalogLimit.(CitrixManagedLimit | ByoaLimit).MaxVdaCount instead | [optional] [readonly] 
**Items** | [**[]CatalogConfiguration**](CatalogConfiguration.md) | List of catalogs that are available to the user | 
**Catalogs** | Pointer to [**[]CatalogConfiguration**](CatalogConfiguration.md) | Alias of Items property for backward compatibility | [optional] 
**CatalogLimit** | Pointer to [**CatalogLimitModel**](CatalogLimitModel.md) | Limit imposed on the number of catalogs for the customer | [optional] 
**StaleData** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustomerManagedCatalogOverviewsModel

`func NewCustomerManagedCatalogOverviewsModel(items []CatalogConfiguration, ) *CustomerManagedCatalogOverviewsModel`

NewCustomerManagedCatalogOverviewsModel instantiates a new CustomerManagedCatalogOverviewsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerManagedCatalogOverviewsModelWithDefaults

`func NewCustomerManagedCatalogOverviewsModelWithDefaults() *CustomerManagedCatalogOverviewsModel`

NewCustomerManagedCatalogOverviewsModelWithDefaults instantiates a new CustomerManagedCatalogOverviewsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxCatalogVdas

`func (o *CustomerManagedCatalogOverviewsModel) GetMaxCatalogVdas() int32`

GetMaxCatalogVdas returns the MaxCatalogVdas field if non-nil, zero value otherwise.

### GetMaxCatalogVdasOk

`func (o *CustomerManagedCatalogOverviewsModel) GetMaxCatalogVdasOk() (*int32, bool)`

GetMaxCatalogVdasOk returns a tuple with the MaxCatalogVdas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCatalogVdas

`func (o *CustomerManagedCatalogOverviewsModel) SetMaxCatalogVdas(v int32)`

SetMaxCatalogVdas sets MaxCatalogVdas field to given value.

### HasMaxCatalogVdas

`func (o *CustomerManagedCatalogOverviewsModel) HasMaxCatalogVdas() bool`

HasMaxCatalogVdas returns a boolean if a field has been set.

### GetItems

`func (o *CustomerManagedCatalogOverviewsModel) GetItems() []CatalogConfiguration`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CustomerManagedCatalogOverviewsModel) GetItemsOk() (*[]CatalogConfiguration, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CustomerManagedCatalogOverviewsModel) SetItems(v []CatalogConfiguration)`

SetItems sets Items field to given value.


### GetCatalogs

`func (o *CustomerManagedCatalogOverviewsModel) GetCatalogs() []CatalogConfiguration`

GetCatalogs returns the Catalogs field if non-nil, zero value otherwise.

### GetCatalogsOk

`func (o *CustomerManagedCatalogOverviewsModel) GetCatalogsOk() (*[]CatalogConfiguration, bool)`

GetCatalogsOk returns a tuple with the Catalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogs

`func (o *CustomerManagedCatalogOverviewsModel) SetCatalogs(v []CatalogConfiguration)`

SetCatalogs sets Catalogs field to given value.

### HasCatalogs

`func (o *CustomerManagedCatalogOverviewsModel) HasCatalogs() bool`

HasCatalogs returns a boolean if a field has been set.

### GetCatalogLimit

`func (o *CustomerManagedCatalogOverviewsModel) GetCatalogLimit() CatalogLimitModel`

GetCatalogLimit returns the CatalogLimit field if non-nil, zero value otherwise.

### GetCatalogLimitOk

`func (o *CustomerManagedCatalogOverviewsModel) GetCatalogLimitOk() (*CatalogLimitModel, bool)`

GetCatalogLimitOk returns a tuple with the CatalogLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogLimit

`func (o *CustomerManagedCatalogOverviewsModel) SetCatalogLimit(v CatalogLimitModel)`

SetCatalogLimit sets CatalogLimit field to given value.

### HasCatalogLimit

`func (o *CustomerManagedCatalogOverviewsModel) HasCatalogLimit() bool`

HasCatalogLimit returns a boolean if a field has been set.

### GetStaleData

`func (o *CustomerManagedCatalogOverviewsModel) GetStaleData() bool`

GetStaleData returns the StaleData field if non-nil, zero value otherwise.

### GetStaleDataOk

`func (o *CustomerManagedCatalogOverviewsModel) GetStaleDataOk() (*bool, bool)`

GetStaleDataOk returns a tuple with the StaleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleData

`func (o *CustomerManagedCatalogOverviewsModel) SetStaleData(v bool)`

SetStaleData sets StaleData field to given value.

### HasStaleData

`func (o *CustomerManagedCatalogOverviewsModel) HasStaleData() bool`

HasStaleData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


