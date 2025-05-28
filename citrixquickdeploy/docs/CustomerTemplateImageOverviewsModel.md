# CustomerTemplateImageOverviewsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]TemplateImageOverview**](TemplateImageOverview.md) | Template Images associated with a customer | 
**Overviews** | Pointer to [**[]TemplateImageOverview**](TemplateImageOverview.md) | Alias of Items property for backward compatibility | [optional] 
**StaleData** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustomerTemplateImageOverviewsModel

`func NewCustomerTemplateImageOverviewsModel(items []TemplateImageOverview, ) *CustomerTemplateImageOverviewsModel`

NewCustomerTemplateImageOverviewsModel instantiates a new CustomerTemplateImageOverviewsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerTemplateImageOverviewsModelWithDefaults

`func NewCustomerTemplateImageOverviewsModelWithDefaults() *CustomerTemplateImageOverviewsModel`

NewCustomerTemplateImageOverviewsModelWithDefaults instantiates a new CustomerTemplateImageOverviewsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CustomerTemplateImageOverviewsModel) GetItems() []TemplateImageOverview`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CustomerTemplateImageOverviewsModel) GetItemsOk() (*[]TemplateImageOverview, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CustomerTemplateImageOverviewsModel) SetItems(v []TemplateImageOverview)`

SetItems sets Items field to given value.


### GetOverviews

`func (o *CustomerTemplateImageOverviewsModel) GetOverviews() []TemplateImageOverview`

GetOverviews returns the Overviews field if non-nil, zero value otherwise.

### GetOverviewsOk

`func (o *CustomerTemplateImageOverviewsModel) GetOverviewsOk() (*[]TemplateImageOverview, bool)`

GetOverviewsOk returns a tuple with the Overviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverviews

`func (o *CustomerTemplateImageOverviewsModel) SetOverviews(v []TemplateImageOverview)`

SetOverviews sets Overviews field to given value.

### HasOverviews

`func (o *CustomerTemplateImageOverviewsModel) HasOverviews() bool`

HasOverviews returns a boolean if a field has been set.

### GetStaleData

`func (o *CustomerTemplateImageOverviewsModel) GetStaleData() bool`

GetStaleData returns the StaleData field if non-nil, zero value otherwise.

### GetStaleDataOk

`func (o *CustomerTemplateImageOverviewsModel) GetStaleDataOk() (*bool, bool)`

GetStaleDataOk returns a tuple with the StaleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleData

`func (o *CustomerTemplateImageOverviewsModel) SetStaleData(v bool)`

SetStaleData sets StaleData field to given value.

### HasStaleData

`func (o *CustomerTemplateImageOverviewsModel) HasStaleData() bool`

HasStaleData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


