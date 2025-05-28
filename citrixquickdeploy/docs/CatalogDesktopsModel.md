# CatalogDesktopsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]PublishedDesktop**](PublishedDesktop.md) |  | 
**Desktops** | Pointer to [**[]PublishedDesktop**](PublishedDesktop.md) |  | [optional] 

## Methods

### NewCatalogDesktopsModel

`func NewCatalogDesktopsModel(items []PublishedDesktop, ) *CatalogDesktopsModel`

NewCatalogDesktopsModel instantiates a new CatalogDesktopsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogDesktopsModelWithDefaults

`func NewCatalogDesktopsModelWithDefaults() *CatalogDesktopsModel`

NewCatalogDesktopsModelWithDefaults instantiates a new CatalogDesktopsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CatalogDesktopsModel) GetItems() []PublishedDesktop`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CatalogDesktopsModel) GetItemsOk() (*[]PublishedDesktop, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CatalogDesktopsModel) SetItems(v []PublishedDesktop)`

SetItems sets Items field to given value.


### GetDesktops

`func (o *CatalogDesktopsModel) GetDesktops() []PublishedDesktop`

GetDesktops returns the Desktops field if non-nil, zero value otherwise.

### GetDesktopsOk

`func (o *CatalogDesktopsModel) GetDesktopsOk() (*[]PublishedDesktop, bool)`

GetDesktopsOk returns a tuple with the Desktops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktops

`func (o *CatalogDesktopsModel) SetDesktops(v []PublishedDesktop)`

SetDesktops sets Desktops field to given value.

### HasDesktops

`func (o *CatalogDesktopsModel) HasDesktops() bool`

HasDesktops returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


