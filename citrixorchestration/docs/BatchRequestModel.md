# BatchRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]BatchRequestItemModel**](BatchRequestItemModel.md) | List of batch request items. | 

## Methods

### NewBatchRequestModel

`func NewBatchRequestModel(items []BatchRequestItemModel, ) *BatchRequestModel`

NewBatchRequestModel instantiates a new BatchRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchRequestModelWithDefaults

`func NewBatchRequestModelWithDefaults() *BatchRequestModel`

NewBatchRequestModelWithDefaults instantiates a new BatchRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *BatchRequestModel) GetItems() []BatchRequestItemModel`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BatchRequestModel) GetItemsOk() (*[]BatchRequestItemModel, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BatchRequestModel) SetItems(v []BatchRequestItemModel)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


