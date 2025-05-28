# AzureConsumptionUsageDetailForCsv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogId** | Pointer to **string** | ID of the catalog | [optional] 
**SubscriptionId** | Pointer to **string** | Subscription guid. | [optional] 
**Cost** | Pointer to **float64** | The amount of cost before tax. | [optional] [readonly] 
**ResourceId** | Pointer to **string** | Unique identifier of the Azure Resource Manager usage detail resource. | [optional] [readonly] 
**ResourceGroup** | Pointer to **string** | Name of the resource&#39;s resource group. | [optional] [readonly] 
**ResourceLocation** | Pointer to **string** | Location of the resource. | [optional] [readonly] 
**BillingPeriodStartDate** | Pointer to **time.Time** | The billing period start date. | [optional] [readonly] 
**BillingPeriodEndDate** | Pointer to **time.Time** | The billing period end date. | [optional] [readonly] 
**MeterId** | Pointer to **string** | The meter id (GUID). Not available for marketplace. For reserved instance this represents the primary meter for which the reservation was purchased. For the actual VM Size for which the reservation is purchased see productOrderName. | [optional] [readonly] 
**Quantity** | Pointer to **float64** | The usage quantity. | [optional] [readonly] 
**UnitPrice** | Pointer to **float64** | The unit price of the meter. | [optional] [readonly] 
**BillingCurrency** | Pointer to **string** | The currency of the meter. | [optional] [readonly] 
**MeterDetails** | Pointer to [**AzureConsumptionMeterDetails**](AzureConsumptionMeterDetails.md) | The details about the meter. By default this is not populated, unless it&#39;s specified in $expand. | [optional] 
**Tags** | Pointer to **string** |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


