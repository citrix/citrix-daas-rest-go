# CollectionModelOfHypervisorPluginResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]HypervisorPluginResponseModel**](HypervisorPluginResponseModel.md) | List of items. | 
**ContinuationToken** | Pointer to **string** | If present, indicates to the caller that the query was not complete, and they should call the API again specifying the continuation token as a query parameter. | [optional] 
**TotalItems** | Pointer to **int32** | Indicates the total number of items in the collection, which may be more than the number of Items returned, if there is a ContinuationToken.  Only returned in the response to &#x60;$search&#x60; APIs. | [optional] 

## Methods

### NewCollectionModelOfHypervisorPluginResponseModel

`func NewCollectionModelOfHypervisorPluginResponseModel(items []HypervisorPluginResponseModel, ) *CollectionModelOfHypervisorPluginResponseModel`

NewCollectionModelOfHypervisorPluginResponseModel instantiates a new CollectionModelOfHypervisorPluginResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionModelOfHypervisorPluginResponseModelWithDefaults

`func NewCollectionModelOfHypervisorPluginResponseModelWithDefaults() *CollectionModelOfHypervisorPluginResponseModel`

NewCollectionModelOfHypervisorPluginResponseModelWithDefaults instantiates a new CollectionModelOfHypervisorPluginResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CollectionModelOfHypervisorPluginResponseModel) GetItems() []HypervisorPluginResponseModel`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CollectionModelOfHypervisorPluginResponseModel) GetItemsOk() (*[]HypervisorPluginResponseModel, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CollectionModelOfHypervisorPluginResponseModel) SetItems(v []HypervisorPluginResponseModel)`

SetItems sets Items field to given value.


### GetContinuationToken

`func (o *CollectionModelOfHypervisorPluginResponseModel) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *CollectionModelOfHypervisorPluginResponseModel) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *CollectionModelOfHypervisorPluginResponseModel) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *CollectionModelOfHypervisorPluginResponseModel) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### GetTotalItems

`func (o *CollectionModelOfHypervisorPluginResponseModel) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *CollectionModelOfHypervisorPluginResponseModel) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *CollectionModelOfHypervisorPluginResponseModel) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *CollectionModelOfHypervisorPluginResponseModel) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


