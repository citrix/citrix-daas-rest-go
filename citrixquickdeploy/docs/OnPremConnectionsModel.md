# OnPremConnectionsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StaleData** | Pointer to **bool** |  | [optional] 
**Items** | Pointer to [**[]OnPremConnectionModel**](OnPremConnectionModel.md) |  | [optional] [readonly] 

## Methods

### NewOnPremConnectionsModel

`func NewOnPremConnectionsModel() *OnPremConnectionsModel`

NewOnPremConnectionsModel instantiates a new OnPremConnectionsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnPremConnectionsModelWithDefaults

`func NewOnPremConnectionsModelWithDefaults() *OnPremConnectionsModel`

NewOnPremConnectionsModelWithDefaults instantiates a new OnPremConnectionsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStaleData

`func (o *OnPremConnectionsModel) GetStaleData() bool`

GetStaleData returns the StaleData field if non-nil, zero value otherwise.

### GetStaleDataOk

`func (o *OnPremConnectionsModel) GetStaleDataOk() (*bool, bool)`

GetStaleDataOk returns a tuple with the StaleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleData

`func (o *OnPremConnectionsModel) SetStaleData(v bool)`

SetStaleData sets StaleData field to given value.

### HasStaleData

`func (o *OnPremConnectionsModel) HasStaleData() bool`

HasStaleData returns a boolean if a field has been set.

### GetItems

`func (o *OnPremConnectionsModel) GetItems() []OnPremConnectionModel`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OnPremConnectionsModel) GetItemsOk() (*[]OnPremConnectionModel, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OnPremConnectionsModel) SetItems(v []OnPremConnectionModel)`

SetItems sets Items field to given value.

### HasItems

`func (o *OnPremConnectionsModel) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


