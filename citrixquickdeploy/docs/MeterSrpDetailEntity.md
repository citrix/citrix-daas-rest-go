# MeterSrpDetailEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeterId** | Pointer to **string** | The Meter Id | [optional] 
**Name** | Pointer to **string** | Meter Name | [optional] 
**SrpPrice** | Pointer to **float64** | The Meter Price | [optional] 
**Category** | Pointer to **string** | Meter Category | [optional] 
**CloudProviderName** | Pointer to **string** | Cloud Provider Name | [optional] 
**UnitOfMeasure** | Pointer to **string** | Unit of measure | [optional] 
**IncludedUnits** | Pointer to **float64** | Include units | [optional] 
**Region** | Pointer to **string** | Region | [optional] 
**EffectiveDate** | Pointer to **time.Time** | The effective date of the price. | [optional] 
**ExpiredDate** | Pointer to **time.Time** | The expiration date of the price. | [optional] 
**CreateTime** | Pointer to **time.Time** | The Time when this meter price inputted. | [optional] 

## Methods

### NewMeterSrpDetailEntity

`func NewMeterSrpDetailEntity() *MeterSrpDetailEntity`

NewMeterSrpDetailEntity instantiates a new MeterSrpDetailEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeterSrpDetailEntityWithDefaults

`func NewMeterSrpDetailEntityWithDefaults() *MeterSrpDetailEntity`

NewMeterSrpDetailEntityWithDefaults instantiates a new MeterSrpDetailEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeterId

`func (o *MeterSrpDetailEntity) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *MeterSrpDetailEntity) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *MeterSrpDetailEntity) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *MeterSrpDetailEntity) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetName

`func (o *MeterSrpDetailEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MeterSrpDetailEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MeterSrpDetailEntity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MeterSrpDetailEntity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSrpPrice

`func (o *MeterSrpDetailEntity) GetSrpPrice() float64`

GetSrpPrice returns the SrpPrice field if non-nil, zero value otherwise.

### GetSrpPriceOk

`func (o *MeterSrpDetailEntity) GetSrpPriceOk() (*float64, bool)`

GetSrpPriceOk returns a tuple with the SrpPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrpPrice

`func (o *MeterSrpDetailEntity) SetSrpPrice(v float64)`

SetSrpPrice sets SrpPrice field to given value.

### HasSrpPrice

`func (o *MeterSrpDetailEntity) HasSrpPrice() bool`

HasSrpPrice returns a boolean if a field has been set.

### GetCategory

`func (o *MeterSrpDetailEntity) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MeterSrpDetailEntity) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MeterSrpDetailEntity) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MeterSrpDetailEntity) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCloudProviderName

`func (o *MeterSrpDetailEntity) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *MeterSrpDetailEntity) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *MeterSrpDetailEntity) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.

### HasCloudProviderName

`func (o *MeterSrpDetailEntity) HasCloudProviderName() bool`

HasCloudProviderName returns a boolean if a field has been set.

### GetUnitOfMeasure

`func (o *MeterSrpDetailEntity) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *MeterSrpDetailEntity) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *MeterSrpDetailEntity) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *MeterSrpDetailEntity) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetIncludedUnits

`func (o *MeterSrpDetailEntity) GetIncludedUnits() float64`

GetIncludedUnits returns the IncludedUnits field if non-nil, zero value otherwise.

### GetIncludedUnitsOk

`func (o *MeterSrpDetailEntity) GetIncludedUnitsOk() (*float64, bool)`

GetIncludedUnitsOk returns a tuple with the IncludedUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUnits

`func (o *MeterSrpDetailEntity) SetIncludedUnits(v float64)`

SetIncludedUnits sets IncludedUnits field to given value.

### HasIncludedUnits

`func (o *MeterSrpDetailEntity) HasIncludedUnits() bool`

HasIncludedUnits returns a boolean if a field has been set.

### GetRegion

`func (o *MeterSrpDetailEntity) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MeterSrpDetailEntity) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MeterSrpDetailEntity) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MeterSrpDetailEntity) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *MeterSrpDetailEntity) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *MeterSrpDetailEntity) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *MeterSrpDetailEntity) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *MeterSrpDetailEntity) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetExpiredDate

`func (o *MeterSrpDetailEntity) GetExpiredDate() time.Time`

GetExpiredDate returns the ExpiredDate field if non-nil, zero value otherwise.

### GetExpiredDateOk

`func (o *MeterSrpDetailEntity) GetExpiredDateOk() (*time.Time, bool)`

GetExpiredDateOk returns a tuple with the ExpiredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredDate

`func (o *MeterSrpDetailEntity) SetExpiredDate(v time.Time)`

SetExpiredDate sets ExpiredDate field to given value.

### HasExpiredDate

`func (o *MeterSrpDetailEntity) HasExpiredDate() bool`

HasExpiredDate returns a boolean if a field has been set.

### GetCreateTime

`func (o *MeterSrpDetailEntity) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *MeterSrpDetailEntity) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *MeterSrpDetailEntity) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *MeterSrpDetailEntity) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


