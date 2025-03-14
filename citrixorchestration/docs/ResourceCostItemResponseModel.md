# ResourceCostItemResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | Pointer to **NullableString** | Specifies the resource for which the cost is calculated. | [optional] 
**Cost** | Pointer to **float32** | The cost of the resource. | [optional] 
**Currency** | Pointer to **NullableString** | The currency of the cost. | [optional] 

## Methods

### NewResourceCostItemResponseModel

`func NewResourceCostItemResponseModel() *ResourceCostItemResponseModel`

NewResourceCostItemResponseModel instantiates a new ResourceCostItemResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceCostItemResponseModelWithDefaults

`func NewResourceCostItemResponseModelWithDefaults() *ResourceCostItemResponseModel`

NewResourceCostItemResponseModelWithDefaults instantiates a new ResourceCostItemResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *ResourceCostItemResponseModel) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ResourceCostItemResponseModel) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ResourceCostItemResponseModel) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ResourceCostItemResponseModel) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *ResourceCostItemResponseModel) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *ResourceCostItemResponseModel) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil
### GetCost

`func (o *ResourceCostItemResponseModel) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ResourceCostItemResponseModel) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ResourceCostItemResponseModel) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ResourceCostItemResponseModel) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetCurrency

`func (o *ResourceCostItemResponseModel) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ResourceCostItemResponseModel) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ResourceCostItemResponseModel) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ResourceCostItemResponseModel) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *ResourceCostItemResponseModel) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *ResourceCostItemResponseModel) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


