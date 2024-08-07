# ReservationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Term** | Pointer to **NullableString** | Gets or sets the term. Such as &#39;1 Year&#39; and &#39;3 Years&#39;. | [optional] 
**EffectivePrice** | Pointer to **NullableFloat32** | Gets or sets the effective price. | [optional] 
**RetailPrice** | Pointer to **NullableFloat32** | Gets or sets the retail price. | [optional] 
**Currency** | Pointer to **NullableString** | Gets or sets the currency. | [optional] 

## Methods

### NewReservationResponseModel

`func NewReservationResponseModel() *ReservationResponseModel`

NewReservationResponseModel instantiates a new ReservationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationResponseModelWithDefaults

`func NewReservationResponseModelWithDefaults() *ReservationResponseModel`

NewReservationResponseModelWithDefaults instantiates a new ReservationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTerm

`func (o *ReservationResponseModel) GetTerm() string`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *ReservationResponseModel) GetTermOk() (*string, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *ReservationResponseModel) SetTerm(v string)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *ReservationResponseModel) HasTerm() bool`

HasTerm returns a boolean if a field has been set.

### SetTermNil

`func (o *ReservationResponseModel) SetTermNil(b bool)`

 SetTermNil sets the value for Term to be an explicit nil

### UnsetTerm
`func (o *ReservationResponseModel) UnsetTerm()`

UnsetTerm ensures that no value is present for Term, not even an explicit nil
### GetEffectivePrice

`func (o *ReservationResponseModel) GetEffectivePrice() float32`

GetEffectivePrice returns the EffectivePrice field if non-nil, zero value otherwise.

### GetEffectivePriceOk

`func (o *ReservationResponseModel) GetEffectivePriceOk() (*float32, bool)`

GetEffectivePriceOk returns a tuple with the EffectivePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectivePrice

`func (o *ReservationResponseModel) SetEffectivePrice(v float32)`

SetEffectivePrice sets EffectivePrice field to given value.

### HasEffectivePrice

`func (o *ReservationResponseModel) HasEffectivePrice() bool`

HasEffectivePrice returns a boolean if a field has been set.

### SetEffectivePriceNil

`func (o *ReservationResponseModel) SetEffectivePriceNil(b bool)`

 SetEffectivePriceNil sets the value for EffectivePrice to be an explicit nil

### UnsetEffectivePrice
`func (o *ReservationResponseModel) UnsetEffectivePrice()`

UnsetEffectivePrice ensures that no value is present for EffectivePrice, not even an explicit nil
### GetRetailPrice

`func (o *ReservationResponseModel) GetRetailPrice() float32`

GetRetailPrice returns the RetailPrice field if non-nil, zero value otherwise.

### GetRetailPriceOk

`func (o *ReservationResponseModel) GetRetailPriceOk() (*float32, bool)`

GetRetailPriceOk returns a tuple with the RetailPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailPrice

`func (o *ReservationResponseModel) SetRetailPrice(v float32)`

SetRetailPrice sets RetailPrice field to given value.

### HasRetailPrice

`func (o *ReservationResponseModel) HasRetailPrice() bool`

HasRetailPrice returns a boolean if a field has been set.

### SetRetailPriceNil

`func (o *ReservationResponseModel) SetRetailPriceNil(b bool)`

 SetRetailPriceNil sets the value for RetailPrice to be an explicit nil

### UnsetRetailPrice
`func (o *ReservationResponseModel) UnsetRetailPrice()`

UnsetRetailPrice ensures that no value is present for RetailPrice, not even an explicit nil
### GetCurrency

`func (o *ReservationResponseModel) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ReservationResponseModel) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ReservationResponseModel) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ReservationResponseModel) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *ReservationResponseModel) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *ReservationResponseModel) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


