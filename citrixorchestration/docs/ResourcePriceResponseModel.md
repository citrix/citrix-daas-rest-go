# ResourcePriceResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to **NullableString** | Gets or sets the location. | [optional] 
**EffectivePrice** | Pointer to [**UnitPriceResponseModel**](UnitPriceResponseModel.md) |  | [optional] 
**RetailPrice** | Pointer to [**UnitPriceResponseModel**](UnitPriceResponseModel.md) |  | [optional] 
**SavingPlans** | Pointer to [**[]SavingPlanResponseModel**](SavingPlanResponseModel.md) | Gets or sets the savings plans. | [optional] 
**Reservations** | Pointer to [**[]ReservationResponseModel**](ReservationResponseModel.md) | Gets or sets the reservations. | [optional] 
**AdditionalData** | Pointer to **map[string]string** | Gets or sets the additional data. | [optional] 

## Methods

### NewResourcePriceResponseModel

`func NewResourcePriceResponseModel() *ResourcePriceResponseModel`

NewResourcePriceResponseModel instantiates a new ResourcePriceResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePriceResponseModelWithDefaults

`func NewResourcePriceResponseModelWithDefaults() *ResourcePriceResponseModel`

NewResourcePriceResponseModelWithDefaults instantiates a new ResourcePriceResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *ResourcePriceResponseModel) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ResourcePriceResponseModel) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ResourcePriceResponseModel) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ResourcePriceResponseModel) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *ResourcePriceResponseModel) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *ResourcePriceResponseModel) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetEffectivePrice

`func (o *ResourcePriceResponseModel) GetEffectivePrice() UnitPriceResponseModel`

GetEffectivePrice returns the EffectivePrice field if non-nil, zero value otherwise.

### GetEffectivePriceOk

`func (o *ResourcePriceResponseModel) GetEffectivePriceOk() (*UnitPriceResponseModel, bool)`

GetEffectivePriceOk returns a tuple with the EffectivePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectivePrice

`func (o *ResourcePriceResponseModel) SetEffectivePrice(v UnitPriceResponseModel)`

SetEffectivePrice sets EffectivePrice field to given value.

### HasEffectivePrice

`func (o *ResourcePriceResponseModel) HasEffectivePrice() bool`

HasEffectivePrice returns a boolean if a field has been set.

### GetRetailPrice

`func (o *ResourcePriceResponseModel) GetRetailPrice() UnitPriceResponseModel`

GetRetailPrice returns the RetailPrice field if non-nil, zero value otherwise.

### GetRetailPriceOk

`func (o *ResourcePriceResponseModel) GetRetailPriceOk() (*UnitPriceResponseModel, bool)`

GetRetailPriceOk returns a tuple with the RetailPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailPrice

`func (o *ResourcePriceResponseModel) SetRetailPrice(v UnitPriceResponseModel)`

SetRetailPrice sets RetailPrice field to given value.

### HasRetailPrice

`func (o *ResourcePriceResponseModel) HasRetailPrice() bool`

HasRetailPrice returns a boolean if a field has been set.

### GetSavingPlans

`func (o *ResourcePriceResponseModel) GetSavingPlans() []SavingPlanResponseModel`

GetSavingPlans returns the SavingPlans field if non-nil, zero value otherwise.

### GetSavingPlansOk

`func (o *ResourcePriceResponseModel) GetSavingPlansOk() (*[]SavingPlanResponseModel, bool)`

GetSavingPlansOk returns a tuple with the SavingPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingPlans

`func (o *ResourcePriceResponseModel) SetSavingPlans(v []SavingPlanResponseModel)`

SetSavingPlans sets SavingPlans field to given value.

### HasSavingPlans

`func (o *ResourcePriceResponseModel) HasSavingPlans() bool`

HasSavingPlans returns a boolean if a field has been set.

### SetSavingPlansNil

`func (o *ResourcePriceResponseModel) SetSavingPlansNil(b bool)`

 SetSavingPlansNil sets the value for SavingPlans to be an explicit nil

### UnsetSavingPlans
`func (o *ResourcePriceResponseModel) UnsetSavingPlans()`

UnsetSavingPlans ensures that no value is present for SavingPlans, not even an explicit nil
### GetReservations

`func (o *ResourcePriceResponseModel) GetReservations() []ReservationResponseModel`

GetReservations returns the Reservations field if non-nil, zero value otherwise.

### GetReservationsOk

`func (o *ResourcePriceResponseModel) GetReservationsOk() (*[]ReservationResponseModel, bool)`

GetReservationsOk returns a tuple with the Reservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservations

`func (o *ResourcePriceResponseModel) SetReservations(v []ReservationResponseModel)`

SetReservations sets Reservations field to given value.

### HasReservations

`func (o *ResourcePriceResponseModel) HasReservations() bool`

HasReservations returns a boolean if a field has been set.

### SetReservationsNil

`func (o *ResourcePriceResponseModel) SetReservationsNil(b bool)`

 SetReservationsNil sets the value for Reservations to be an explicit nil

### UnsetReservations
`func (o *ResourcePriceResponseModel) UnsetReservations()`

UnsetReservations ensures that no value is present for Reservations, not even an explicit nil
### GetAdditionalData

`func (o *ResourcePriceResponseModel) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *ResourcePriceResponseModel) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *ResourcePriceResponseModel) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *ResourcePriceResponseModel) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### SetAdditionalDataNil

`func (o *ResourcePriceResponseModel) SetAdditionalDataNil(b bool)`

 SetAdditionalDataNil sets the value for AdditionalData to be an explicit nil

### UnsetAdditionalData
`func (o *ResourcePriceResponseModel) UnsetAdditionalData()`

UnsetAdditionalData ensures that no value is present for AdditionalData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


