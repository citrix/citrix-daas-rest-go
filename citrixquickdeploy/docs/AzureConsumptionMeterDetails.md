# AzureConsumptionMeterDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeterName** | Pointer to **NullableString** | The name of the meter, within the given meter category. | [optional] 
**MeterCategory** | Pointer to **NullableString** | The category of the meter, for example, &#39;Cloud services&#39;, &#39;Networking&#39;, etc.. | [optional] 
**MeterSubCategory** | Pointer to **NullableString** | The subcategory of the meter, for example, &#39;A6 Cloud services&#39;, &#39;ExpressRoute (IXP)&#39;, etc.. | [optional] 
**Unit** | Pointer to **NullableString** | The unit in which the meter consumption is charged, for example, &#39;Hours&#39;, &#39;GB&#39;, etc. | [optional] 
**UnitOfMeasure** | Pointer to **NullableString** | The unit in which the meter consumption is charged, for example, &#39;Hours&#39;, &#39;GB&#39;, etc. | [optional] 
**MeterLocation** | Pointer to **NullableString** | The location in which the Azure service is available. | [optional] 

## Methods

### NewAzureConsumptionMeterDetails

`func NewAzureConsumptionMeterDetails() *AzureConsumptionMeterDetails`

NewAzureConsumptionMeterDetails instantiates a new AzureConsumptionMeterDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureConsumptionMeterDetailsWithDefaults

`func NewAzureConsumptionMeterDetailsWithDefaults() *AzureConsumptionMeterDetails`

NewAzureConsumptionMeterDetailsWithDefaults instantiates a new AzureConsumptionMeterDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeterName

`func (o *AzureConsumptionMeterDetails) GetMeterName() string`

GetMeterName returns the MeterName field if non-nil, zero value otherwise.

### GetMeterNameOk

`func (o *AzureConsumptionMeterDetails) GetMeterNameOk() (*string, bool)`

GetMeterNameOk returns a tuple with the MeterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterName

`func (o *AzureConsumptionMeterDetails) SetMeterName(v string)`

SetMeterName sets MeterName field to given value.

### HasMeterName

`func (o *AzureConsumptionMeterDetails) HasMeterName() bool`

HasMeterName returns a boolean if a field has been set.

### SetMeterNameNil

`func (o *AzureConsumptionMeterDetails) SetMeterNameNil(b bool)`

 SetMeterNameNil sets the value for MeterName to be an explicit nil

### UnsetMeterName
`func (o *AzureConsumptionMeterDetails) UnsetMeterName()`

UnsetMeterName ensures that no value is present for MeterName, not even an explicit nil
### GetMeterCategory

`func (o *AzureConsumptionMeterDetails) GetMeterCategory() string`

GetMeterCategory returns the MeterCategory field if non-nil, zero value otherwise.

### GetMeterCategoryOk

`func (o *AzureConsumptionMeterDetails) GetMeterCategoryOk() (*string, bool)`

GetMeterCategoryOk returns a tuple with the MeterCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterCategory

`func (o *AzureConsumptionMeterDetails) SetMeterCategory(v string)`

SetMeterCategory sets MeterCategory field to given value.

### HasMeterCategory

`func (o *AzureConsumptionMeterDetails) HasMeterCategory() bool`

HasMeterCategory returns a boolean if a field has been set.

### SetMeterCategoryNil

`func (o *AzureConsumptionMeterDetails) SetMeterCategoryNil(b bool)`

 SetMeterCategoryNil sets the value for MeterCategory to be an explicit nil

### UnsetMeterCategory
`func (o *AzureConsumptionMeterDetails) UnsetMeterCategory()`

UnsetMeterCategory ensures that no value is present for MeterCategory, not even an explicit nil
### GetMeterSubCategory

`func (o *AzureConsumptionMeterDetails) GetMeterSubCategory() string`

GetMeterSubCategory returns the MeterSubCategory field if non-nil, zero value otherwise.

### GetMeterSubCategoryOk

`func (o *AzureConsumptionMeterDetails) GetMeterSubCategoryOk() (*string, bool)`

GetMeterSubCategoryOk returns a tuple with the MeterSubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterSubCategory

`func (o *AzureConsumptionMeterDetails) SetMeterSubCategory(v string)`

SetMeterSubCategory sets MeterSubCategory field to given value.

### HasMeterSubCategory

`func (o *AzureConsumptionMeterDetails) HasMeterSubCategory() bool`

HasMeterSubCategory returns a boolean if a field has been set.

### SetMeterSubCategoryNil

`func (o *AzureConsumptionMeterDetails) SetMeterSubCategoryNil(b bool)`

 SetMeterSubCategoryNil sets the value for MeterSubCategory to be an explicit nil

### UnsetMeterSubCategory
`func (o *AzureConsumptionMeterDetails) UnsetMeterSubCategory()`

UnsetMeterSubCategory ensures that no value is present for MeterSubCategory, not even an explicit nil
### GetUnit

`func (o *AzureConsumptionMeterDetails) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *AzureConsumptionMeterDetails) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *AzureConsumptionMeterDetails) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *AzureConsumptionMeterDetails) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *AzureConsumptionMeterDetails) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *AzureConsumptionMeterDetails) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetUnitOfMeasure

`func (o *AzureConsumptionMeterDetails) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *AzureConsumptionMeterDetails) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *AzureConsumptionMeterDetails) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *AzureConsumptionMeterDetails) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### SetUnitOfMeasureNil

`func (o *AzureConsumptionMeterDetails) SetUnitOfMeasureNil(b bool)`

 SetUnitOfMeasureNil sets the value for UnitOfMeasure to be an explicit nil

### UnsetUnitOfMeasure
`func (o *AzureConsumptionMeterDetails) UnsetUnitOfMeasure()`

UnsetUnitOfMeasure ensures that no value is present for UnitOfMeasure, not even an explicit nil
### GetMeterLocation

`func (o *AzureConsumptionMeterDetails) GetMeterLocation() string`

GetMeterLocation returns the MeterLocation field if non-nil, zero value otherwise.

### GetMeterLocationOk

`func (o *AzureConsumptionMeterDetails) GetMeterLocationOk() (*string, bool)`

GetMeterLocationOk returns a tuple with the MeterLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterLocation

`func (o *AzureConsumptionMeterDetails) SetMeterLocation(v string)`

SetMeterLocation sets MeterLocation field to given value.

### HasMeterLocation

`func (o *AzureConsumptionMeterDetails) HasMeterLocation() bool`

HasMeterLocation returns a boolean if a field has been set.

### SetMeterLocationNil

`func (o *AzureConsumptionMeterDetails) SetMeterLocationNil(b bool)`

 SetMeterLocationNil sets the value for MeterLocation to be an explicit nil

### UnsetMeterLocation
`func (o *AzureConsumptionMeterDetails) UnsetMeterLocation()`

UnsetMeterLocation ensures that no value is present for MeterLocation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


