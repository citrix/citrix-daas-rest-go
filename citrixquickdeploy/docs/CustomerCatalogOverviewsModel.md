# CustomerCatalogOverviewsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]CatalogConfiguration**](CatalogConfiguration.md) | List of catalogs that are available to the user | 
**Catalogs** | Pointer to [**[]CatalogConfiguration**](CatalogConfiguration.md) | Alias of Items property for backward compatibility | [optional] 
**CatalogLimit** | Pointer to [**CatalogLimitModel**](CatalogLimitModel.md) | Limit imposed on the number of catalogs for the customer | [optional] 
**StaleData** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustomerCatalogOverviewsModel

`func NewCustomerCatalogOverviewsModel(items []CatalogConfiguration, ) *CustomerCatalogOverviewsModel`

NewCustomerCatalogOverviewsModel instantiates a new CustomerCatalogOverviewsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerCatalogOverviewsModelWithDefaults

`func NewCustomerCatalogOverviewsModelWithDefaults() *CustomerCatalogOverviewsModel`

NewCustomerCatalogOverviewsModelWithDefaults instantiates a new CustomerCatalogOverviewsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CustomerCatalogOverviewsModel) GetItems() []CatalogConfiguration`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CustomerCatalogOverviewsModel) GetItemsOk() (*[]CatalogConfiguration, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CustomerCatalogOverviewsModel) SetItems(v []CatalogConfiguration)`

SetItems sets Items field to given value.


### GetCatalogs

`func (o *CustomerCatalogOverviewsModel) GetCatalogs() []CatalogConfiguration`

GetCatalogs returns the Catalogs field if non-nil, zero value otherwise.

### GetCatalogsOk

`func (o *CustomerCatalogOverviewsModel) GetCatalogsOk() (*[]CatalogConfiguration, bool)`

GetCatalogsOk returns a tuple with the Catalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogs

`func (o *CustomerCatalogOverviewsModel) SetCatalogs(v []CatalogConfiguration)`

SetCatalogs sets Catalogs field to given value.

### HasCatalogs

`func (o *CustomerCatalogOverviewsModel) HasCatalogs() bool`

HasCatalogs returns a boolean if a field has been set.

### GetCatalogLimit

`func (o *CustomerCatalogOverviewsModel) GetCatalogLimit() CatalogLimitModel`

GetCatalogLimit returns the CatalogLimit field if non-nil, zero value otherwise.

### GetCatalogLimitOk

`func (o *CustomerCatalogOverviewsModel) GetCatalogLimitOk() (*CatalogLimitModel, bool)`

GetCatalogLimitOk returns a tuple with the CatalogLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogLimit

`func (o *CustomerCatalogOverviewsModel) SetCatalogLimit(v CatalogLimitModel)`

SetCatalogLimit sets CatalogLimit field to given value.

### HasCatalogLimit

`func (o *CustomerCatalogOverviewsModel) HasCatalogLimit() bool`

HasCatalogLimit returns a boolean if a field has been set.

### GetStaleData

`func (o *CustomerCatalogOverviewsModel) GetStaleData() bool`

GetStaleData returns the StaleData field if non-nil, zero value otherwise.

### GetStaleDataOk

`func (o *CustomerCatalogOverviewsModel) GetStaleDataOk() (*bool, bool)`

GetStaleDataOk returns a tuple with the StaleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleData

`func (o *CustomerCatalogOverviewsModel) SetStaleData(v bool)`

SetStaleData sets StaleData field to given value.

### HasStaleData

`func (o *CustomerCatalogOverviewsModel) HasStaleData() bool`

HasStaleData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


