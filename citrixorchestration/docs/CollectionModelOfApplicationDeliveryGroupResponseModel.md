# CollectionModelOfApplicationDeliveryGroupResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ApplicationDeliveryGroupResponseModel**](ApplicationDeliveryGroupResponseModel.md) | List of items. | 
**ContinuationToken** | Pointer to **string** | If present, indicates to the caller that the query was not complete, and they should call the API again specifying the continuation token as a query parameter. | [optional] 
**TotalItems** | Pointer to **int32** | Indicates the total number of items in the collection, which may be more than the number of Items returned, if there is a ContinuationToken.  Only returned in the response to &#x60;$search&#x60; APIs. | [optional] 

## Methods

### NewCollectionModelOfApplicationDeliveryGroupResponseModel

`func NewCollectionModelOfApplicationDeliveryGroupResponseModel(items []ApplicationDeliveryGroupResponseModel, ) *CollectionModelOfApplicationDeliveryGroupResponseModel`

NewCollectionModelOfApplicationDeliveryGroupResponseModel instantiates a new CollectionModelOfApplicationDeliveryGroupResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionModelOfApplicationDeliveryGroupResponseModelWithDefaults

`func NewCollectionModelOfApplicationDeliveryGroupResponseModelWithDefaults() *CollectionModelOfApplicationDeliveryGroupResponseModel`

NewCollectionModelOfApplicationDeliveryGroupResponseModelWithDefaults instantiates a new CollectionModelOfApplicationDeliveryGroupResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CollectionModelOfApplicationDeliveryGroupResponseModel) GetItems() []ApplicationDeliveryGroupResponseModel`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CollectionModelOfApplicationDeliveryGroupResponseModel) GetItemsOk() (*[]ApplicationDeliveryGroupResponseModel, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CollectionModelOfApplicationDeliveryGroupResponseModel) SetItems(v []ApplicationDeliveryGroupResponseModel)`

SetItems sets Items field to given value.


### GetContinuationToken

`func (o *CollectionModelOfApplicationDeliveryGroupResponseModel) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *CollectionModelOfApplicationDeliveryGroupResponseModel) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *CollectionModelOfApplicationDeliveryGroupResponseModel) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *CollectionModelOfApplicationDeliveryGroupResponseModel) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### GetTotalItems

`func (o *CollectionModelOfApplicationDeliveryGroupResponseModel) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *CollectionModelOfApplicationDeliveryGroupResponseModel) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *CollectionModelOfApplicationDeliveryGroupResponseModel) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *CollectionModelOfApplicationDeliveryGroupResponseModel) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


