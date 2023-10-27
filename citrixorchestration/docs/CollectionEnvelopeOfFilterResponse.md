# CollectionEnvelopeOfFilterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]FilterResponse**](FilterResponse.md) | Items in the collection. | [optional] 
**ContinuationToken** | Pointer to **NullableString** | If present, indicates to the caller that the query was not complete, and it should call the API again specifying the continuation token as a query parameter. | [optional] 
**TotalItems** | Pointer to **NullableInt32** | Indicates the total number of items in the collection, which may be more than the number of items returned, if there is a ContinuationToken. This is returned only in the response to $search APIs. | [optional] 

## Methods

### NewCollectionEnvelopeOfFilterResponse

`func NewCollectionEnvelopeOfFilterResponse() *CollectionEnvelopeOfFilterResponse`

NewCollectionEnvelopeOfFilterResponse instantiates a new CollectionEnvelopeOfFilterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionEnvelopeOfFilterResponseWithDefaults

`func NewCollectionEnvelopeOfFilterResponseWithDefaults() *CollectionEnvelopeOfFilterResponse`

NewCollectionEnvelopeOfFilterResponseWithDefaults instantiates a new CollectionEnvelopeOfFilterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CollectionEnvelopeOfFilterResponse) GetItems() []FilterResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CollectionEnvelopeOfFilterResponse) GetItemsOk() (*[]FilterResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CollectionEnvelopeOfFilterResponse) SetItems(v []FilterResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *CollectionEnvelopeOfFilterResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *CollectionEnvelopeOfFilterResponse) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *CollectionEnvelopeOfFilterResponse) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetContinuationToken

`func (o *CollectionEnvelopeOfFilterResponse) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *CollectionEnvelopeOfFilterResponse) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *CollectionEnvelopeOfFilterResponse) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *CollectionEnvelopeOfFilterResponse) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### SetContinuationTokenNil

`func (o *CollectionEnvelopeOfFilterResponse) SetContinuationTokenNil(b bool)`

 SetContinuationTokenNil sets the value for ContinuationToken to be an explicit nil

### UnsetContinuationToken
`func (o *CollectionEnvelopeOfFilterResponse) UnsetContinuationToken()`

UnsetContinuationToken ensures that no value is present for ContinuationToken, not even an explicit nil
### GetTotalItems

`func (o *CollectionEnvelopeOfFilterResponse) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *CollectionEnvelopeOfFilterResponse) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *CollectionEnvelopeOfFilterResponse) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *CollectionEnvelopeOfFilterResponse) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### SetTotalItemsNil

`func (o *CollectionEnvelopeOfFilterResponse) SetTotalItemsNil(b bool)`

 SetTotalItemsNil sets the value for TotalItems to be an explicit nil

### UnsetTotalItems
`func (o *CollectionEnvelopeOfFilterResponse) UnsetTotalItems()`

UnsetTotalItems ensures that no value is present for TotalItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


