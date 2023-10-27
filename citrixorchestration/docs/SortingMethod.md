# SortingMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | [**SearchProperty**](SearchProperty.md) |  | 
**SortDirection** | [**ListSortDirection**](ListSortDirection.md) |  | 

## Methods

### NewSortingMethod

`func NewSortingMethod(property SearchProperty, sortDirection ListSortDirection, ) *SortingMethod`

NewSortingMethod instantiates a new SortingMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSortingMethodWithDefaults

`func NewSortingMethodWithDefaults() *SortingMethod`

NewSortingMethodWithDefaults instantiates a new SortingMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *SortingMethod) GetProperty() SearchProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *SortingMethod) GetPropertyOk() (*SearchProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *SortingMethod) SetProperty(v SearchProperty)`

SetProperty sets Property field to given value.


### GetSortDirection

`func (o *SortingMethod) GetSortDirection() ListSortDirection`

GetSortDirection returns the SortDirection field if non-nil, zero value otherwise.

### GetSortDirectionOk

`func (o *SortingMethod) GetSortDirectionOk() (*ListSortDirection, bool)`

GetSortDirectionOk returns a tuple with the SortDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortDirection

`func (o *SortingMethod) SetSortDirection(v ListSortDirection)`

SetSortDirection sets SortDirection field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


