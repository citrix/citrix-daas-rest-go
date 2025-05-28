# CustomerResourceLocations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ResourceLocationDetails**](ResourceLocationDetails.md) | The list of Resource Locations | [optional] 
**StaleData** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustomerResourceLocations

`func NewCustomerResourceLocations() *CustomerResourceLocations`

NewCustomerResourceLocations instantiates a new CustomerResourceLocations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerResourceLocationsWithDefaults

`func NewCustomerResourceLocationsWithDefaults() *CustomerResourceLocations`

NewCustomerResourceLocationsWithDefaults instantiates a new CustomerResourceLocations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CustomerResourceLocations) GetItems() []ResourceLocationDetails`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CustomerResourceLocations) GetItemsOk() (*[]ResourceLocationDetails, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CustomerResourceLocations) SetItems(v []ResourceLocationDetails)`

SetItems sets Items field to given value.

### HasItems

`func (o *CustomerResourceLocations) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetStaleData

`func (o *CustomerResourceLocations) GetStaleData() bool`

GetStaleData returns the StaleData field if non-nil, zero value otherwise.

### GetStaleDataOk

`func (o *CustomerResourceLocations) GetStaleDataOk() (*bool, bool)`

GetStaleDataOk returns a tuple with the StaleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleData

`func (o *CustomerResourceLocations) SetStaleData(v bool)`

SetStaleData sets StaleData field to given value.

### HasStaleData

`func (o *CustomerResourceLocations) HasStaleData() bool`

HasStaleData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


