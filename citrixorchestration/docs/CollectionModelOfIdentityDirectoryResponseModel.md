# CollectionModelOfIdentityDirectoryResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]IdentityDirectoryResponseModel**](IdentityDirectoryResponseModel.md) | List of items. | 
**ContinuationToken** | Pointer to **NullableString** | If present, indicates to the caller that the query was not complete, and they should call the API again specifying the continuation token as a query parameter. | [optional] 
**TotalItems** | Pointer to **NullableInt32** | Indicates the total number of items in the collection, which may be more than the number of Items returned, if there is a ContinuationToken.  Only returned in the response to &#x60;$search&#x60; APIs. | [optional] 

## Methods

### NewCollectionModelOfIdentityDirectoryResponseModel

`func NewCollectionModelOfIdentityDirectoryResponseModel(items []IdentityDirectoryResponseModel, ) *CollectionModelOfIdentityDirectoryResponseModel`

NewCollectionModelOfIdentityDirectoryResponseModel instantiates a new CollectionModelOfIdentityDirectoryResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionModelOfIdentityDirectoryResponseModelWithDefaults

`func NewCollectionModelOfIdentityDirectoryResponseModelWithDefaults() *CollectionModelOfIdentityDirectoryResponseModel`

NewCollectionModelOfIdentityDirectoryResponseModelWithDefaults instantiates a new CollectionModelOfIdentityDirectoryResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CollectionModelOfIdentityDirectoryResponseModel) GetItems() []IdentityDirectoryResponseModel`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CollectionModelOfIdentityDirectoryResponseModel) GetItemsOk() (*[]IdentityDirectoryResponseModel, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CollectionModelOfIdentityDirectoryResponseModel) SetItems(v []IdentityDirectoryResponseModel)`

SetItems sets Items field to given value.


### GetContinuationToken

`func (o *CollectionModelOfIdentityDirectoryResponseModel) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *CollectionModelOfIdentityDirectoryResponseModel) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *CollectionModelOfIdentityDirectoryResponseModel) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *CollectionModelOfIdentityDirectoryResponseModel) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### SetContinuationTokenNil

`func (o *CollectionModelOfIdentityDirectoryResponseModel) SetContinuationTokenNil(b bool)`

 SetContinuationTokenNil sets the value for ContinuationToken to be an explicit nil

### UnsetContinuationToken
`func (o *CollectionModelOfIdentityDirectoryResponseModel) UnsetContinuationToken()`

UnsetContinuationToken ensures that no value is present for ContinuationToken, not even an explicit nil
### GetTotalItems

`func (o *CollectionModelOfIdentityDirectoryResponseModel) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *CollectionModelOfIdentityDirectoryResponseModel) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *CollectionModelOfIdentityDirectoryResponseModel) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *CollectionModelOfIdentityDirectoryResponseModel) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### SetTotalItemsNil

`func (o *CollectionModelOfIdentityDirectoryResponseModel) SetTotalItemsNil(b bool)`

 SetTotalItemsNil sets the value for TotalItems to be an explicit nil

### UnsetTotalItems
`func (o *CollectionModelOfIdentityDirectoryResponseModel) UnsetTotalItems()`

UnsetTotalItems ensures that no value is present for TotalItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


