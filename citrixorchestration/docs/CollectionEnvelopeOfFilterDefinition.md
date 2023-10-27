# CollectionEnvelopeOfFilterDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]FilterDefinition**](FilterDefinition.md) | Items in the collection. | [optional] 
**ContinuationToken** | Pointer to **NullableString** | If present, indicates to the caller that the query was not complete, and it should call the API again specifying the continuation token as a query parameter. | [optional] 
**TotalItems** | Pointer to **NullableInt32** | Indicates the total number of items in the collection, which may be more than the number of items returned, if there is a ContinuationToken. This is returned only in the response to $search APIs. | [optional] 

## Methods

### NewCollectionEnvelopeOfFilterDefinition

`func NewCollectionEnvelopeOfFilterDefinition() *CollectionEnvelopeOfFilterDefinition`

NewCollectionEnvelopeOfFilterDefinition instantiates a new CollectionEnvelopeOfFilterDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionEnvelopeOfFilterDefinitionWithDefaults

`func NewCollectionEnvelopeOfFilterDefinitionWithDefaults() *CollectionEnvelopeOfFilterDefinition`

NewCollectionEnvelopeOfFilterDefinitionWithDefaults instantiates a new CollectionEnvelopeOfFilterDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CollectionEnvelopeOfFilterDefinition) GetItems() []FilterDefinition`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CollectionEnvelopeOfFilterDefinition) GetItemsOk() (*[]FilterDefinition, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CollectionEnvelopeOfFilterDefinition) SetItems(v []FilterDefinition)`

SetItems sets Items field to given value.

### HasItems

`func (o *CollectionEnvelopeOfFilterDefinition) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *CollectionEnvelopeOfFilterDefinition) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *CollectionEnvelopeOfFilterDefinition) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetContinuationToken

`func (o *CollectionEnvelopeOfFilterDefinition) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *CollectionEnvelopeOfFilterDefinition) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *CollectionEnvelopeOfFilterDefinition) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *CollectionEnvelopeOfFilterDefinition) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### SetContinuationTokenNil

`func (o *CollectionEnvelopeOfFilterDefinition) SetContinuationTokenNil(b bool)`

 SetContinuationTokenNil sets the value for ContinuationToken to be an explicit nil

### UnsetContinuationToken
`func (o *CollectionEnvelopeOfFilterDefinition) UnsetContinuationToken()`

UnsetContinuationToken ensures that no value is present for ContinuationToken, not even an explicit nil
### GetTotalItems

`func (o *CollectionEnvelopeOfFilterDefinition) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *CollectionEnvelopeOfFilterDefinition) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *CollectionEnvelopeOfFilterDefinition) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *CollectionEnvelopeOfFilterDefinition) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### SetTotalItemsNil

`func (o *CollectionEnvelopeOfFilterDefinition) SetTotalItemsNil(b bool)`

 SetTotalItemsNil sets the value for TotalItems to be an explicit nil

### UnsetTotalItems
`func (o *CollectionEnvelopeOfFilterDefinition) UnsetTotalItems()`

UnsetTotalItems ensures that no value is present for TotalItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


