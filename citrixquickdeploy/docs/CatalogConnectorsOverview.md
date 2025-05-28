# CatalogConnectorsOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ConnectorOverview**](ConnectorOverview.md) | List of catalogs that are available to the user | 
**Connectors** | Pointer to [**[]ConnectorOverview**](ConnectorOverview.md) | Alias of Item property for backward compatibility | [optional] 

## Methods

### NewCatalogConnectorsOverview

`func NewCatalogConnectorsOverview(items []ConnectorOverview, ) *CatalogConnectorsOverview`

NewCatalogConnectorsOverview instantiates a new CatalogConnectorsOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogConnectorsOverviewWithDefaults

`func NewCatalogConnectorsOverviewWithDefaults() *CatalogConnectorsOverview`

NewCatalogConnectorsOverviewWithDefaults instantiates a new CatalogConnectorsOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CatalogConnectorsOverview) GetItems() []ConnectorOverview`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CatalogConnectorsOverview) GetItemsOk() (*[]ConnectorOverview, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CatalogConnectorsOverview) SetItems(v []ConnectorOverview)`

SetItems sets Items field to given value.


### GetConnectors

`func (o *CatalogConnectorsOverview) GetConnectors() []ConnectorOverview`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *CatalogConnectorsOverview) GetConnectorsOk() (*[]ConnectorOverview, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *CatalogConnectorsOverview) SetConnectors(v []ConnectorOverview)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *CatalogConnectorsOverview) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


