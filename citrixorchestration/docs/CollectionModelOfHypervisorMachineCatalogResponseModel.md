# CollectionModelOfHypervisorMachineCatalogResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]HypervisorMachineCatalogResponseModel**](HypervisorMachineCatalogResponseModel.md) | List of items. | 
**ContinuationToken** | Pointer to **string** | If present, indicates to the caller that the query was not complete, and they should call the API again specifying the continuation token as a query parameter. | [optional] 
**TotalItems** | Pointer to **int32** | Indicates the total number of items in the collection, which may be more than the number of Items returned, if there is a ContinuationToken.  Only returned in the response to &#x60;$search&#x60; APIs. | [optional] 

## Methods

### NewCollectionModelOfHypervisorMachineCatalogResponseModel

`func NewCollectionModelOfHypervisorMachineCatalogResponseModel(items []HypervisorMachineCatalogResponseModel, ) *CollectionModelOfHypervisorMachineCatalogResponseModel`

NewCollectionModelOfHypervisorMachineCatalogResponseModel instantiates a new CollectionModelOfHypervisorMachineCatalogResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionModelOfHypervisorMachineCatalogResponseModelWithDefaults

`func NewCollectionModelOfHypervisorMachineCatalogResponseModelWithDefaults() *CollectionModelOfHypervisorMachineCatalogResponseModel`

NewCollectionModelOfHypervisorMachineCatalogResponseModelWithDefaults instantiates a new CollectionModelOfHypervisorMachineCatalogResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CollectionModelOfHypervisorMachineCatalogResponseModel) GetItems() []HypervisorMachineCatalogResponseModel`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CollectionModelOfHypervisorMachineCatalogResponseModel) GetItemsOk() (*[]HypervisorMachineCatalogResponseModel, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CollectionModelOfHypervisorMachineCatalogResponseModel) SetItems(v []HypervisorMachineCatalogResponseModel)`

SetItems sets Items field to given value.


### GetContinuationToken

`func (o *CollectionModelOfHypervisorMachineCatalogResponseModel) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *CollectionModelOfHypervisorMachineCatalogResponseModel) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *CollectionModelOfHypervisorMachineCatalogResponseModel) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *CollectionModelOfHypervisorMachineCatalogResponseModel) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### GetTotalItems

`func (o *CollectionModelOfHypervisorMachineCatalogResponseModel) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *CollectionModelOfHypervisorMachineCatalogResponseModel) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *CollectionModelOfHypervisorMachineCatalogResponseModel) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *CollectionModelOfHypervisorMachineCatalogResponseModel) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


