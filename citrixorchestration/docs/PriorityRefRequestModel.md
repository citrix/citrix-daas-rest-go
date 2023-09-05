# PriorityRefRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | Pointer to **string** | Reference to the item. | [optional] 
**Priority** | Pointer to **int32** | Priority.  See the containing object context for details about how the Priority affects the object relationship. | [optional] [default to 0]

## Methods

### NewPriorityRefRequestModel

`func NewPriorityRefRequestModel() *PriorityRefRequestModel`

NewPriorityRefRequestModel instantiates a new PriorityRefRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriorityRefRequestModelWithDefaults

`func NewPriorityRefRequestModelWithDefaults() *PriorityRefRequestModel`

NewPriorityRefRequestModelWithDefaults instantiates a new PriorityRefRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *PriorityRefRequestModel) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *PriorityRefRequestModel) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *PriorityRefRequestModel) SetItem(v string)`

SetItem sets Item field to given value.

### HasItem

`func (o *PriorityRefRequestModel) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetPriority

`func (o *PriorityRefRequestModel) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PriorityRefRequestModel) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PriorityRefRequestModel) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PriorityRefRequestModel) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


