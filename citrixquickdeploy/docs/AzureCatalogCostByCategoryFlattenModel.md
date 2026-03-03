# AzureCatalogCostByCategoryFlattenModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | Pointer to **NullableString** | Azure Subscription ID | [optional] 
**CatalogId** | Pointer to **NullableString** | Catalog ID | [optional] 
**CatalogName** | Pointer to **NullableString** | Catalog Name | [optional] 
**Category** | Pointer to **NullableString** | Azure Cost category | [optional] 
**Cost** | Pointer to **float64** | Azure Cost | [optional] 

## Methods

### NewAzureCatalogCostByCategoryFlattenModel

`func NewAzureCatalogCostByCategoryFlattenModel() *AzureCatalogCostByCategoryFlattenModel`

NewAzureCatalogCostByCategoryFlattenModel instantiates a new AzureCatalogCostByCategoryFlattenModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCatalogCostByCategoryFlattenModelWithDefaults

`func NewAzureCatalogCostByCategoryFlattenModelWithDefaults() *AzureCatalogCostByCategoryFlattenModel`

NewAzureCatalogCostByCategoryFlattenModelWithDefaults instantiates a new AzureCatalogCostByCategoryFlattenModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *AzureCatalogCostByCategoryFlattenModel) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureCatalogCostByCategoryFlattenModel) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureCatalogCostByCategoryFlattenModel) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AzureCatalogCostByCategoryFlattenModel) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *AzureCatalogCostByCategoryFlattenModel) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *AzureCatalogCostByCategoryFlattenModel) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetCatalogId

`func (o *AzureCatalogCostByCategoryFlattenModel) GetCatalogId() string`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *AzureCatalogCostByCategoryFlattenModel) GetCatalogIdOk() (*string, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *AzureCatalogCostByCategoryFlattenModel) SetCatalogId(v string)`

SetCatalogId sets CatalogId field to given value.

### HasCatalogId

`func (o *AzureCatalogCostByCategoryFlattenModel) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.

### SetCatalogIdNil

`func (o *AzureCatalogCostByCategoryFlattenModel) SetCatalogIdNil(b bool)`

 SetCatalogIdNil sets the value for CatalogId to be an explicit nil

### UnsetCatalogId
`func (o *AzureCatalogCostByCategoryFlattenModel) UnsetCatalogId()`

UnsetCatalogId ensures that no value is present for CatalogId, not even an explicit nil
### GetCatalogName

`func (o *AzureCatalogCostByCategoryFlattenModel) GetCatalogName() string`

GetCatalogName returns the CatalogName field if non-nil, zero value otherwise.

### GetCatalogNameOk

`func (o *AzureCatalogCostByCategoryFlattenModel) GetCatalogNameOk() (*string, bool)`

GetCatalogNameOk returns a tuple with the CatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogName

`func (o *AzureCatalogCostByCategoryFlattenModel) SetCatalogName(v string)`

SetCatalogName sets CatalogName field to given value.

### HasCatalogName

`func (o *AzureCatalogCostByCategoryFlattenModel) HasCatalogName() bool`

HasCatalogName returns a boolean if a field has been set.

### SetCatalogNameNil

`func (o *AzureCatalogCostByCategoryFlattenModel) SetCatalogNameNil(b bool)`

 SetCatalogNameNil sets the value for CatalogName to be an explicit nil

### UnsetCatalogName
`func (o *AzureCatalogCostByCategoryFlattenModel) UnsetCatalogName()`

UnsetCatalogName ensures that no value is present for CatalogName, not even an explicit nil
### GetCategory

`func (o *AzureCatalogCostByCategoryFlattenModel) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AzureCatalogCostByCategoryFlattenModel) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AzureCatalogCostByCategoryFlattenModel) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AzureCatalogCostByCategoryFlattenModel) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *AzureCatalogCostByCategoryFlattenModel) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *AzureCatalogCostByCategoryFlattenModel) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetCost

`func (o *AzureCatalogCostByCategoryFlattenModel) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *AzureCatalogCostByCategoryFlattenModel) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *AzureCatalogCostByCategoryFlattenModel) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *AzureCatalogCostByCategoryFlattenModel) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


