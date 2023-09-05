# BatchResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]BatchResponseItemModel**](BatchResponseItemModel.md) | List of batch item responses. | 

## Methods

### NewBatchResponseModel

`func NewBatchResponseModel(items []BatchResponseItemModel, ) *BatchResponseModel`

NewBatchResponseModel instantiates a new BatchResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResponseModelWithDefaults

`func NewBatchResponseModelWithDefaults() *BatchResponseModel`

NewBatchResponseModelWithDefaults instantiates a new BatchResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *BatchResponseModel) GetItems() []BatchResponseItemModel`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BatchResponseModel) GetItemsOk() (*[]BatchResponseItemModel, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BatchResponseModel) SetItems(v []BatchResponseItemModel)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


