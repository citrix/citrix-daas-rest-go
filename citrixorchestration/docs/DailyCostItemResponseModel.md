# DailyCostItemResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsageDate** | Pointer to **int64** | The date of the usage. | [optional] 
**ResourceCostItems** | Pointer to [**[]ResourceCostItemResponseModel**](ResourceCostItemResponseModel.md) | The cost of the resources. | [optional] 

## Methods

### NewDailyCostItemResponseModel

`func NewDailyCostItemResponseModel() *DailyCostItemResponseModel`

NewDailyCostItemResponseModel instantiates a new DailyCostItemResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyCostItemResponseModelWithDefaults

`func NewDailyCostItemResponseModelWithDefaults() *DailyCostItemResponseModel`

NewDailyCostItemResponseModelWithDefaults instantiates a new DailyCostItemResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsageDate

`func (o *DailyCostItemResponseModel) GetUsageDate() int64`

GetUsageDate returns the UsageDate field if non-nil, zero value otherwise.

### GetUsageDateOk

`func (o *DailyCostItemResponseModel) GetUsageDateOk() (*int64, bool)`

GetUsageDateOk returns a tuple with the UsageDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageDate

`func (o *DailyCostItemResponseModel) SetUsageDate(v int64)`

SetUsageDate sets UsageDate field to given value.

### HasUsageDate

`func (o *DailyCostItemResponseModel) HasUsageDate() bool`

HasUsageDate returns a boolean if a field has been set.

### GetResourceCostItems

`func (o *DailyCostItemResponseModel) GetResourceCostItems() []ResourceCostItemResponseModel`

GetResourceCostItems returns the ResourceCostItems field if non-nil, zero value otherwise.

### GetResourceCostItemsOk

`func (o *DailyCostItemResponseModel) GetResourceCostItemsOk() (*[]ResourceCostItemResponseModel, bool)`

GetResourceCostItemsOk returns a tuple with the ResourceCostItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCostItems

`func (o *DailyCostItemResponseModel) SetResourceCostItems(v []ResourceCostItemResponseModel)`

SetResourceCostItems sets ResourceCostItems field to given value.

### HasResourceCostItems

`func (o *DailyCostItemResponseModel) HasResourceCostItems() bool`

HasResourceCostItems returns a boolean if a field has been set.

### SetResourceCostItemsNil

`func (o *DailyCostItemResponseModel) SetResourceCostItemsNil(b bool)`

 SetResourceCostItemsNil sets the value for ResourceCostItems to be an explicit nil

### UnsetResourceCostItems
`func (o *DailyCostItemResponseModel) UnsetResourceCostItems()`

UnsetResourceCostItems ensures that no value is present for ResourceCostItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


