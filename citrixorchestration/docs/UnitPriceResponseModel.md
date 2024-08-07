# UnitPriceResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **NullableFloat32** | Gets or sets the price. | [optional] 
**UnitOfMeasure** | Pointer to **NullableString** | Gets or sets the unit of measure, such as per hour, per month, etc. | [optional] 
**Currency** | Pointer to **NullableString** | Gets or sets the currency, such as USD, EUR, etc. | [optional] 

## Methods

### NewUnitPriceResponseModel

`func NewUnitPriceResponseModel() *UnitPriceResponseModel`

NewUnitPriceResponseModel instantiates a new UnitPriceResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnitPriceResponseModelWithDefaults

`func NewUnitPriceResponseModelWithDefaults() *UnitPriceResponseModel`

NewUnitPriceResponseModelWithDefaults instantiates a new UnitPriceResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *UnitPriceResponseModel) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UnitPriceResponseModel) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UnitPriceResponseModel) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *UnitPriceResponseModel) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *UnitPriceResponseModel) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *UnitPriceResponseModel) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetUnitOfMeasure

`func (o *UnitPriceResponseModel) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *UnitPriceResponseModel) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *UnitPriceResponseModel) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *UnitPriceResponseModel) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### SetUnitOfMeasureNil

`func (o *UnitPriceResponseModel) SetUnitOfMeasureNil(b bool)`

 SetUnitOfMeasureNil sets the value for UnitOfMeasure to be an explicit nil

### UnsetUnitOfMeasure
`func (o *UnitPriceResponseModel) UnsetUnitOfMeasure()`

UnsetUnitOfMeasure ensures that no value is present for UnitOfMeasure, not even an explicit nil
### GetCurrency

`func (o *UnitPriceResponseModel) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnitPriceResponseModel) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnitPriceResponseModel) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnitPriceResponseModel) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *UnitPriceResponseModel) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *UnitPriceResponseModel) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


