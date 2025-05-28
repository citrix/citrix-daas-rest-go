# SitesOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]SiteModel**](SiteModel.md) | The list of virtual sites for the customer | 
**Sites** | Pointer to [**[]SiteModel**](SiteModel.md) | Alias of Items property for backward compatibility | [optional] 

## Methods

### NewSitesOverview

`func NewSitesOverview(items []SiteModel, ) *SitesOverview`

NewSitesOverview instantiates a new SitesOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSitesOverviewWithDefaults

`func NewSitesOverviewWithDefaults() *SitesOverview`

NewSitesOverviewWithDefaults instantiates a new SitesOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *SitesOverview) GetItems() []SiteModel`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SitesOverview) GetItemsOk() (*[]SiteModel, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SitesOverview) SetItems(v []SiteModel)`

SetItems sets Items field to given value.


### GetSites

`func (o *SitesOverview) GetSites() []SiteModel`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *SitesOverview) GetSitesOk() (*[]SiteModel, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *SitesOverview) SetSites(v []SiteModel)`

SetSites sets Sites field to given value.

### HasSites

`func (o *SitesOverview) HasSites() bool`

HasSites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


