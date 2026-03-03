# AzureConsumptionUsageDetailForCsv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogId** | Pointer to **NullableString** | ID of the catalog | [optional] 
**SubscriptionId** | Pointer to **NullableString** | Subscription guid. | [optional] 
**SubscriptionName** | Pointer to **NullableString** | Subscription name. | [optional] 
**Cost** | Pointer to **NullableFloat64** | The amount of cost before tax. | [optional] [readonly] 
**ResourceId** | Pointer to **NullableString** | Unique identifier of the Azure Resource Manager usage detail resource. | [optional] [readonly] 
**ResourceGroup** | Pointer to **NullableString** | Name of the resource&#39;s resource group. | [optional] [readonly] 
**ResourceLocation** | Pointer to **NullableString** | Location of the resource. | [optional] [readonly] 
**BillingPeriodStartDate** | Pointer to **NullableTime** | The billing period start date. | [optional] [readonly] 
**BillingPeriodEndDate** | Pointer to **NullableTime** | The billing period end date. | [optional] [readonly] 
**MeterId** | Pointer to **NullableString** | The meter id (GUID). Not available for marketplace. For reserved instance this represents the primary meter for which the reservation was purchased. For the actual VM Size for which the reservation is purchased see productOrderName. | [optional] [readonly] 
**Quantity** | Pointer to **NullableFloat64** | The usage quantity. | [optional] [readonly] 
**UnitPrice** | Pointer to **NullableFloat64** | The unit price of the meter. | [optional] [readonly] 
**BillingCurrency** | Pointer to **NullableString** | The currency of the meter. | [optional] [readonly] 
**MeterDetails** | Pointer to [**NullableAzureConsumptionMeterDetails**](AzureConsumptionMeterDetails.md) | The details about the meter. By default this is not populated, unless it&#39;s specified in $expand. | [optional] [readonly] 
**Tags** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAzureConsumptionUsageDetailForCsv

`func NewAzureConsumptionUsageDetailForCsv() *AzureConsumptionUsageDetailForCsv`

NewAzureConsumptionUsageDetailForCsv instantiates a new AzureConsumptionUsageDetailForCsv object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureConsumptionUsageDetailForCsvWithDefaults

`func NewAzureConsumptionUsageDetailForCsvWithDefaults() *AzureConsumptionUsageDetailForCsv`

NewAzureConsumptionUsageDetailForCsvWithDefaults instantiates a new AzureConsumptionUsageDetailForCsv object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogId

`func (o *AzureConsumptionUsageDetailForCsv) GetCatalogId() string`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *AzureConsumptionUsageDetailForCsv) GetCatalogIdOk() (*string, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *AzureConsumptionUsageDetailForCsv) SetCatalogId(v string)`

SetCatalogId sets CatalogId field to given value.

### HasCatalogId

`func (o *AzureConsumptionUsageDetailForCsv) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.

### SetCatalogIdNil

`func (o *AzureConsumptionUsageDetailForCsv) SetCatalogIdNil(b bool)`

 SetCatalogIdNil sets the value for CatalogId to be an explicit nil

### UnsetCatalogId
`func (o *AzureConsumptionUsageDetailForCsv) UnsetCatalogId()`

UnsetCatalogId ensures that no value is present for CatalogId, not even an explicit nil
### GetSubscriptionId

`func (o *AzureConsumptionUsageDetailForCsv) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureConsumptionUsageDetailForCsv) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureConsumptionUsageDetailForCsv) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AzureConsumptionUsageDetailForCsv) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *AzureConsumptionUsageDetailForCsv) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *AzureConsumptionUsageDetailForCsv) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetSubscriptionName

`func (o *AzureConsumptionUsageDetailForCsv) GetSubscriptionName() string`

GetSubscriptionName returns the SubscriptionName field if non-nil, zero value otherwise.

### GetSubscriptionNameOk

`func (o *AzureConsumptionUsageDetailForCsv) GetSubscriptionNameOk() (*string, bool)`

GetSubscriptionNameOk returns a tuple with the SubscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionName

`func (o *AzureConsumptionUsageDetailForCsv) SetSubscriptionName(v string)`

SetSubscriptionName sets SubscriptionName field to given value.

### HasSubscriptionName

`func (o *AzureConsumptionUsageDetailForCsv) HasSubscriptionName() bool`

HasSubscriptionName returns a boolean if a field has been set.

### SetSubscriptionNameNil

`func (o *AzureConsumptionUsageDetailForCsv) SetSubscriptionNameNil(b bool)`

 SetSubscriptionNameNil sets the value for SubscriptionName to be an explicit nil

### UnsetSubscriptionName
`func (o *AzureConsumptionUsageDetailForCsv) UnsetSubscriptionName()`

UnsetSubscriptionName ensures that no value is present for SubscriptionName, not even an explicit nil
### GetCost

`func (o *AzureConsumptionUsageDetailForCsv) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *AzureConsumptionUsageDetailForCsv) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *AzureConsumptionUsageDetailForCsv) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *AzureConsumptionUsageDetailForCsv) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *AzureConsumptionUsageDetailForCsv) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *AzureConsumptionUsageDetailForCsv) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetResourceId

`func (o *AzureConsumptionUsageDetailForCsv) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AzureConsumptionUsageDetailForCsv) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AzureConsumptionUsageDetailForCsv) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *AzureConsumptionUsageDetailForCsv) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *AzureConsumptionUsageDetailForCsv) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *AzureConsumptionUsageDetailForCsv) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetResourceGroup

`func (o *AzureConsumptionUsageDetailForCsv) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *AzureConsumptionUsageDetailForCsv) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *AzureConsumptionUsageDetailForCsv) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *AzureConsumptionUsageDetailForCsv) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### SetResourceGroupNil

`func (o *AzureConsumptionUsageDetailForCsv) SetResourceGroupNil(b bool)`

 SetResourceGroupNil sets the value for ResourceGroup to be an explicit nil

### UnsetResourceGroup
`func (o *AzureConsumptionUsageDetailForCsv) UnsetResourceGroup()`

UnsetResourceGroup ensures that no value is present for ResourceGroup, not even an explicit nil
### GetResourceLocation

`func (o *AzureConsumptionUsageDetailForCsv) GetResourceLocation() string`

GetResourceLocation returns the ResourceLocation field if non-nil, zero value otherwise.

### GetResourceLocationOk

`func (o *AzureConsumptionUsageDetailForCsv) GetResourceLocationOk() (*string, bool)`

GetResourceLocationOk returns a tuple with the ResourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLocation

`func (o *AzureConsumptionUsageDetailForCsv) SetResourceLocation(v string)`

SetResourceLocation sets ResourceLocation field to given value.

### HasResourceLocation

`func (o *AzureConsumptionUsageDetailForCsv) HasResourceLocation() bool`

HasResourceLocation returns a boolean if a field has been set.

### SetResourceLocationNil

`func (o *AzureConsumptionUsageDetailForCsv) SetResourceLocationNil(b bool)`

 SetResourceLocationNil sets the value for ResourceLocation to be an explicit nil

### UnsetResourceLocation
`func (o *AzureConsumptionUsageDetailForCsv) UnsetResourceLocation()`

UnsetResourceLocation ensures that no value is present for ResourceLocation, not even an explicit nil
### GetBillingPeriodStartDate

`func (o *AzureConsumptionUsageDetailForCsv) GetBillingPeriodStartDate() time.Time`

GetBillingPeriodStartDate returns the BillingPeriodStartDate field if non-nil, zero value otherwise.

### GetBillingPeriodStartDateOk

`func (o *AzureConsumptionUsageDetailForCsv) GetBillingPeriodStartDateOk() (*time.Time, bool)`

GetBillingPeriodStartDateOk returns a tuple with the BillingPeriodStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodStartDate

`func (o *AzureConsumptionUsageDetailForCsv) SetBillingPeriodStartDate(v time.Time)`

SetBillingPeriodStartDate sets BillingPeriodStartDate field to given value.

### HasBillingPeriodStartDate

`func (o *AzureConsumptionUsageDetailForCsv) HasBillingPeriodStartDate() bool`

HasBillingPeriodStartDate returns a boolean if a field has been set.

### SetBillingPeriodStartDateNil

`func (o *AzureConsumptionUsageDetailForCsv) SetBillingPeriodStartDateNil(b bool)`

 SetBillingPeriodStartDateNil sets the value for BillingPeriodStartDate to be an explicit nil

### UnsetBillingPeriodStartDate
`func (o *AzureConsumptionUsageDetailForCsv) UnsetBillingPeriodStartDate()`

UnsetBillingPeriodStartDate ensures that no value is present for BillingPeriodStartDate, not even an explicit nil
### GetBillingPeriodEndDate

`func (o *AzureConsumptionUsageDetailForCsv) GetBillingPeriodEndDate() time.Time`

GetBillingPeriodEndDate returns the BillingPeriodEndDate field if non-nil, zero value otherwise.

### GetBillingPeriodEndDateOk

`func (o *AzureConsumptionUsageDetailForCsv) GetBillingPeriodEndDateOk() (*time.Time, bool)`

GetBillingPeriodEndDateOk returns a tuple with the BillingPeriodEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodEndDate

`func (o *AzureConsumptionUsageDetailForCsv) SetBillingPeriodEndDate(v time.Time)`

SetBillingPeriodEndDate sets BillingPeriodEndDate field to given value.

### HasBillingPeriodEndDate

`func (o *AzureConsumptionUsageDetailForCsv) HasBillingPeriodEndDate() bool`

HasBillingPeriodEndDate returns a boolean if a field has been set.

### SetBillingPeriodEndDateNil

`func (o *AzureConsumptionUsageDetailForCsv) SetBillingPeriodEndDateNil(b bool)`

 SetBillingPeriodEndDateNil sets the value for BillingPeriodEndDate to be an explicit nil

### UnsetBillingPeriodEndDate
`func (o *AzureConsumptionUsageDetailForCsv) UnsetBillingPeriodEndDate()`

UnsetBillingPeriodEndDate ensures that no value is present for BillingPeriodEndDate, not even an explicit nil
### GetMeterId

`func (o *AzureConsumptionUsageDetailForCsv) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *AzureConsumptionUsageDetailForCsv) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *AzureConsumptionUsageDetailForCsv) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *AzureConsumptionUsageDetailForCsv) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### SetMeterIdNil

`func (o *AzureConsumptionUsageDetailForCsv) SetMeterIdNil(b bool)`

 SetMeterIdNil sets the value for MeterId to be an explicit nil

### UnsetMeterId
`func (o *AzureConsumptionUsageDetailForCsv) UnsetMeterId()`

UnsetMeterId ensures that no value is present for MeterId, not even an explicit nil
### GetQuantity

`func (o *AzureConsumptionUsageDetailForCsv) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *AzureConsumptionUsageDetailForCsv) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *AzureConsumptionUsageDetailForCsv) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *AzureConsumptionUsageDetailForCsv) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *AzureConsumptionUsageDetailForCsv) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *AzureConsumptionUsageDetailForCsv) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetUnitPrice

`func (o *AzureConsumptionUsageDetailForCsv) GetUnitPrice() float64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *AzureConsumptionUsageDetailForCsv) GetUnitPriceOk() (*float64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *AzureConsumptionUsageDetailForCsv) SetUnitPrice(v float64)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *AzureConsumptionUsageDetailForCsv) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### SetUnitPriceNil

`func (o *AzureConsumptionUsageDetailForCsv) SetUnitPriceNil(b bool)`

 SetUnitPriceNil sets the value for UnitPrice to be an explicit nil

### UnsetUnitPrice
`func (o *AzureConsumptionUsageDetailForCsv) UnsetUnitPrice()`

UnsetUnitPrice ensures that no value is present for UnitPrice, not even an explicit nil
### GetBillingCurrency

`func (o *AzureConsumptionUsageDetailForCsv) GetBillingCurrency() string`

GetBillingCurrency returns the BillingCurrency field if non-nil, zero value otherwise.

### GetBillingCurrencyOk

`func (o *AzureConsumptionUsageDetailForCsv) GetBillingCurrencyOk() (*string, bool)`

GetBillingCurrencyOk returns a tuple with the BillingCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCurrency

`func (o *AzureConsumptionUsageDetailForCsv) SetBillingCurrency(v string)`

SetBillingCurrency sets BillingCurrency field to given value.

### HasBillingCurrency

`func (o *AzureConsumptionUsageDetailForCsv) HasBillingCurrency() bool`

HasBillingCurrency returns a boolean if a field has been set.

### SetBillingCurrencyNil

`func (o *AzureConsumptionUsageDetailForCsv) SetBillingCurrencyNil(b bool)`

 SetBillingCurrencyNil sets the value for BillingCurrency to be an explicit nil

### UnsetBillingCurrency
`func (o *AzureConsumptionUsageDetailForCsv) UnsetBillingCurrency()`

UnsetBillingCurrency ensures that no value is present for BillingCurrency, not even an explicit nil
### GetMeterDetails

`func (o *AzureConsumptionUsageDetailForCsv) GetMeterDetails() AzureConsumptionMeterDetails`

GetMeterDetails returns the MeterDetails field if non-nil, zero value otherwise.

### GetMeterDetailsOk

`func (o *AzureConsumptionUsageDetailForCsv) GetMeterDetailsOk() (*AzureConsumptionMeterDetails, bool)`

GetMeterDetailsOk returns a tuple with the MeterDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterDetails

`func (o *AzureConsumptionUsageDetailForCsv) SetMeterDetails(v AzureConsumptionMeterDetails)`

SetMeterDetails sets MeterDetails field to given value.

### HasMeterDetails

`func (o *AzureConsumptionUsageDetailForCsv) HasMeterDetails() bool`

HasMeterDetails returns a boolean if a field has been set.

### SetMeterDetailsNil

`func (o *AzureConsumptionUsageDetailForCsv) SetMeterDetailsNil(b bool)`

 SetMeterDetailsNil sets the value for MeterDetails to be an explicit nil

### UnsetMeterDetails
`func (o *AzureConsumptionUsageDetailForCsv) UnsetMeterDetails()`

UnsetMeterDetails ensures that no value is present for MeterDetails, not even an explicit nil
### GetTags

`func (o *AzureConsumptionUsageDetailForCsv) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AzureConsumptionUsageDetailForCsv) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AzureConsumptionUsageDetailForCsv) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AzureConsumptionUsageDetailForCsv) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AzureConsumptionUsageDetailForCsv) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AzureConsumptionUsageDetailForCsv) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


