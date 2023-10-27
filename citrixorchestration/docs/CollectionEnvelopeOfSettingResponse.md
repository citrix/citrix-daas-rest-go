# CollectionEnvelopeOfSettingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]SettingResponse**](SettingResponse.md) | Items in the collection. | [optional] 
**ContinuationToken** | Pointer to **NullableString** | If present, indicates to the caller that the query was not complete, and it should call the API again specifying the continuation token as a query parameter. | [optional] 
**TotalItems** | Pointer to **NullableInt32** | Indicates the total number of items in the collection, which may be more than the number of items returned, if there is a ContinuationToken. This is returned only in the response to $search APIs. | [optional] 

## Methods

### NewCollectionEnvelopeOfSettingResponse

`func NewCollectionEnvelopeOfSettingResponse() *CollectionEnvelopeOfSettingResponse`

NewCollectionEnvelopeOfSettingResponse instantiates a new CollectionEnvelopeOfSettingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionEnvelopeOfSettingResponseWithDefaults

`func NewCollectionEnvelopeOfSettingResponseWithDefaults() *CollectionEnvelopeOfSettingResponse`

NewCollectionEnvelopeOfSettingResponseWithDefaults instantiates a new CollectionEnvelopeOfSettingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CollectionEnvelopeOfSettingResponse) GetItems() []SettingResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CollectionEnvelopeOfSettingResponse) GetItemsOk() (*[]SettingResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CollectionEnvelopeOfSettingResponse) SetItems(v []SettingResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *CollectionEnvelopeOfSettingResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *CollectionEnvelopeOfSettingResponse) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *CollectionEnvelopeOfSettingResponse) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetContinuationToken

`func (o *CollectionEnvelopeOfSettingResponse) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *CollectionEnvelopeOfSettingResponse) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *CollectionEnvelopeOfSettingResponse) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *CollectionEnvelopeOfSettingResponse) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### SetContinuationTokenNil

`func (o *CollectionEnvelopeOfSettingResponse) SetContinuationTokenNil(b bool)`

 SetContinuationTokenNil sets the value for ContinuationToken to be an explicit nil

### UnsetContinuationToken
`func (o *CollectionEnvelopeOfSettingResponse) UnsetContinuationToken()`

UnsetContinuationToken ensures that no value is present for ContinuationToken, not even an explicit nil
### GetTotalItems

`func (o *CollectionEnvelopeOfSettingResponse) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *CollectionEnvelopeOfSettingResponse) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *CollectionEnvelopeOfSettingResponse) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *CollectionEnvelopeOfSettingResponse) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### SetTotalItemsNil

`func (o *CollectionEnvelopeOfSettingResponse) SetTotalItemsNil(b bool)`

 SetTotalItemsNil sets the value for TotalItems to be an explicit nil

### UnsetTotalItems
`func (o *CollectionEnvelopeOfSettingResponse) UnsetTotalItems()`

UnsetTotalItems ensures that no value is present for TotalItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


