# CatalogApplicationsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]PublishedApplication**](PublishedApplication.md) |  | 
**Applications** | Pointer to [**[]PublishedApplication**](PublishedApplication.md) |  | [optional] 

## Methods

### NewCatalogApplicationsModel

`func NewCatalogApplicationsModel(items []PublishedApplication, ) *CatalogApplicationsModel`

NewCatalogApplicationsModel instantiates a new CatalogApplicationsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogApplicationsModelWithDefaults

`func NewCatalogApplicationsModelWithDefaults() *CatalogApplicationsModel`

NewCatalogApplicationsModelWithDefaults instantiates a new CatalogApplicationsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CatalogApplicationsModel) GetItems() []PublishedApplication`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CatalogApplicationsModel) GetItemsOk() (*[]PublishedApplication, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CatalogApplicationsModel) SetItems(v []PublishedApplication)`

SetItems sets Items field to given value.


### GetApplications

`func (o *CatalogApplicationsModel) GetApplications() []PublishedApplication`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *CatalogApplicationsModel) GetApplicationsOk() (*[]PublishedApplication, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *CatalogApplicationsModel) SetApplications(v []PublishedApplication)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *CatalogApplicationsModel) HasApplications() bool`

HasApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


