# HypervisorsAndResourcePoolsResponseModelCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]HypervisorsAndResourcePoolsResponseModel**](HypervisorsAndResourcePoolsResponseModel.md) | List of items. | 
**ContinuationToken** | Pointer to **NullableString** | If present, indicates to the caller that the query was not complete, and they should call the API again specifying the continuation token as a query parameter. | [optional] 
**TotalItems** | Pointer to **NullableInt32** | Indicates the total number of items in the collection, which may be more than the number of Items returned, if there is a ContinuationToken.  Only returned in the response to &#x60;$search&#x60; APIs. | [optional] 

## Methods

### NewHypervisorsAndResourcePoolsResponseModelCollection

`func NewHypervisorsAndResourcePoolsResponseModelCollection(items []HypervisorsAndResourcePoolsResponseModel, ) *HypervisorsAndResourcePoolsResponseModelCollection`

NewHypervisorsAndResourcePoolsResponseModelCollection instantiates a new HypervisorsAndResourcePoolsResponseModelCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorsAndResourcePoolsResponseModelCollectionWithDefaults

`func NewHypervisorsAndResourcePoolsResponseModelCollectionWithDefaults() *HypervisorsAndResourcePoolsResponseModelCollection`

NewHypervisorsAndResourcePoolsResponseModelCollectionWithDefaults instantiates a new HypervisorsAndResourcePoolsResponseModelCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *HypervisorsAndResourcePoolsResponseModelCollection) GetItems() []HypervisorsAndResourcePoolsResponseModel`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *HypervisorsAndResourcePoolsResponseModelCollection) GetItemsOk() (*[]HypervisorsAndResourcePoolsResponseModel, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *HypervisorsAndResourcePoolsResponseModelCollection) SetItems(v []HypervisorsAndResourcePoolsResponseModel)`

SetItems sets Items field to given value.


### GetContinuationToken

`func (o *HypervisorsAndResourcePoolsResponseModelCollection) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *HypervisorsAndResourcePoolsResponseModelCollection) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *HypervisorsAndResourcePoolsResponseModelCollection) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *HypervisorsAndResourcePoolsResponseModelCollection) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### SetContinuationTokenNil

`func (o *HypervisorsAndResourcePoolsResponseModelCollection) SetContinuationTokenNil(b bool)`

 SetContinuationTokenNil sets the value for ContinuationToken to be an explicit nil

### UnsetContinuationToken
`func (o *HypervisorsAndResourcePoolsResponseModelCollection) UnsetContinuationToken()`

UnsetContinuationToken ensures that no value is present for ContinuationToken, not even an explicit nil
### GetTotalItems

`func (o *HypervisorsAndResourcePoolsResponseModelCollection) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *HypervisorsAndResourcePoolsResponseModelCollection) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *HypervisorsAndResourcePoolsResponseModelCollection) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *HypervisorsAndResourcePoolsResponseModelCollection) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### SetTotalItemsNil

`func (o *HypervisorsAndResourcePoolsResponseModelCollection) SetTotalItemsNil(b bool)`

 SetTotalItemsNil sets the value for TotalItems to be an explicit nil

### UnsetTotalItems
`func (o *HypervisorsAndResourcePoolsResponseModelCollection) UnsetTotalItems()`

UnsetTotalItems ensures that no value is present for TotalItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


