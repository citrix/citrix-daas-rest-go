# VNetResourceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]VNetResourceLocations**](VNetResourceLocations.md) | List of objects with VNet Resource Locations for each Citrix Managed Subscription | [optional] 

## Methods

### NewVNetResourceDetails

`func NewVNetResourceDetails() *VNetResourceDetails`

NewVNetResourceDetails instantiates a new VNetResourceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVNetResourceDetailsWithDefaults

`func NewVNetResourceDetailsWithDefaults() *VNetResourceDetails`

NewVNetResourceDetailsWithDefaults instantiates a new VNetResourceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *VNetResourceDetails) GetItems() []VNetResourceLocations`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *VNetResourceDetails) GetItemsOk() (*[]VNetResourceLocations, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *VNetResourceDetails) SetItems(v []VNetResourceLocations)`

SetItems sets Items field to given value.

### HasItems

`func (o *VNetResourceDetails) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


