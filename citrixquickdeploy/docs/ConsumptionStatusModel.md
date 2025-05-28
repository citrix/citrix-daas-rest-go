# ConsumptionStatusModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonthlyConsumptionCommitment** | Pointer to **float64** |  | [optional] 
**MonthlyConsumption** | Pointer to **float64** |  | [optional] 
**TermConsumptionCommitment** | Pointer to **float64** |  | [optional] 
**TermConsumption** | Pointer to **float64** |  | [optional] 
**TotalConsumption** | Pointer to **float64** |  | [optional] 
**NumberOfLicensedUsers** | Pointer to **int32** |  | [optional] 
**ConsumptionSettings** | Pointer to [**DataStoreConsumptionThresholdModel**](DataStoreConsumptionThresholdModel.md) |  | [optional] 
**StaleData** | Pointer to **bool** |  | [optional] 

## Methods

### NewConsumptionStatusModel

`func NewConsumptionStatusModel() *ConsumptionStatusModel`

NewConsumptionStatusModel instantiates a new ConsumptionStatusModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumptionStatusModelWithDefaults

`func NewConsumptionStatusModelWithDefaults() *ConsumptionStatusModel`

NewConsumptionStatusModelWithDefaults instantiates a new ConsumptionStatusModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonthlyConsumptionCommitment

`func (o *ConsumptionStatusModel) GetMonthlyConsumptionCommitment() float64`

GetMonthlyConsumptionCommitment returns the MonthlyConsumptionCommitment field if non-nil, zero value otherwise.

### GetMonthlyConsumptionCommitmentOk

`func (o *ConsumptionStatusModel) GetMonthlyConsumptionCommitmentOk() (*float64, bool)`

GetMonthlyConsumptionCommitmentOk returns a tuple with the MonthlyConsumptionCommitment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyConsumptionCommitment

`func (o *ConsumptionStatusModel) SetMonthlyConsumptionCommitment(v float64)`

SetMonthlyConsumptionCommitment sets MonthlyConsumptionCommitment field to given value.

### HasMonthlyConsumptionCommitment

`func (o *ConsumptionStatusModel) HasMonthlyConsumptionCommitment() bool`

HasMonthlyConsumptionCommitment returns a boolean if a field has been set.

### GetMonthlyConsumption

`func (o *ConsumptionStatusModel) GetMonthlyConsumption() float64`

GetMonthlyConsumption returns the MonthlyConsumption field if non-nil, zero value otherwise.

### GetMonthlyConsumptionOk

`func (o *ConsumptionStatusModel) GetMonthlyConsumptionOk() (*float64, bool)`

GetMonthlyConsumptionOk returns a tuple with the MonthlyConsumption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyConsumption

`func (o *ConsumptionStatusModel) SetMonthlyConsumption(v float64)`

SetMonthlyConsumption sets MonthlyConsumption field to given value.

### HasMonthlyConsumption

`func (o *ConsumptionStatusModel) HasMonthlyConsumption() bool`

HasMonthlyConsumption returns a boolean if a field has been set.

### GetTermConsumptionCommitment

`func (o *ConsumptionStatusModel) GetTermConsumptionCommitment() float64`

GetTermConsumptionCommitment returns the TermConsumptionCommitment field if non-nil, zero value otherwise.

### GetTermConsumptionCommitmentOk

`func (o *ConsumptionStatusModel) GetTermConsumptionCommitmentOk() (*float64, bool)`

GetTermConsumptionCommitmentOk returns a tuple with the TermConsumptionCommitment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermConsumptionCommitment

`func (o *ConsumptionStatusModel) SetTermConsumptionCommitment(v float64)`

SetTermConsumptionCommitment sets TermConsumptionCommitment field to given value.

### HasTermConsumptionCommitment

`func (o *ConsumptionStatusModel) HasTermConsumptionCommitment() bool`

HasTermConsumptionCommitment returns a boolean if a field has been set.

### GetTermConsumption

`func (o *ConsumptionStatusModel) GetTermConsumption() float64`

GetTermConsumption returns the TermConsumption field if non-nil, zero value otherwise.

### GetTermConsumptionOk

`func (o *ConsumptionStatusModel) GetTermConsumptionOk() (*float64, bool)`

GetTermConsumptionOk returns a tuple with the TermConsumption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermConsumption

`func (o *ConsumptionStatusModel) SetTermConsumption(v float64)`

SetTermConsumption sets TermConsumption field to given value.

### HasTermConsumption

`func (o *ConsumptionStatusModel) HasTermConsumption() bool`

HasTermConsumption returns a boolean if a field has been set.

### GetTotalConsumption

`func (o *ConsumptionStatusModel) GetTotalConsumption() float64`

GetTotalConsumption returns the TotalConsumption field if non-nil, zero value otherwise.

### GetTotalConsumptionOk

`func (o *ConsumptionStatusModel) GetTotalConsumptionOk() (*float64, bool)`

GetTotalConsumptionOk returns a tuple with the TotalConsumption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalConsumption

`func (o *ConsumptionStatusModel) SetTotalConsumption(v float64)`

SetTotalConsumption sets TotalConsumption field to given value.

### HasTotalConsumption

`func (o *ConsumptionStatusModel) HasTotalConsumption() bool`

HasTotalConsumption returns a boolean if a field has been set.

### GetNumberOfLicensedUsers

`func (o *ConsumptionStatusModel) GetNumberOfLicensedUsers() int32`

GetNumberOfLicensedUsers returns the NumberOfLicensedUsers field if non-nil, zero value otherwise.

### GetNumberOfLicensedUsersOk

`func (o *ConsumptionStatusModel) GetNumberOfLicensedUsersOk() (*int32, bool)`

GetNumberOfLicensedUsersOk returns a tuple with the NumberOfLicensedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfLicensedUsers

`func (o *ConsumptionStatusModel) SetNumberOfLicensedUsers(v int32)`

SetNumberOfLicensedUsers sets NumberOfLicensedUsers field to given value.

### HasNumberOfLicensedUsers

`func (o *ConsumptionStatusModel) HasNumberOfLicensedUsers() bool`

HasNumberOfLicensedUsers returns a boolean if a field has been set.

### GetConsumptionSettings

`func (o *ConsumptionStatusModel) GetConsumptionSettings() DataStoreConsumptionThresholdModel`

GetConsumptionSettings returns the ConsumptionSettings field if non-nil, zero value otherwise.

### GetConsumptionSettingsOk

`func (o *ConsumptionStatusModel) GetConsumptionSettingsOk() (*DataStoreConsumptionThresholdModel, bool)`

GetConsumptionSettingsOk returns a tuple with the ConsumptionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumptionSettings

`func (o *ConsumptionStatusModel) SetConsumptionSettings(v DataStoreConsumptionThresholdModel)`

SetConsumptionSettings sets ConsumptionSettings field to given value.

### HasConsumptionSettings

`func (o *ConsumptionStatusModel) HasConsumptionSettings() bool`

HasConsumptionSettings returns a boolean if a field has been set.

### GetStaleData

`func (o *ConsumptionStatusModel) GetStaleData() bool`

GetStaleData returns the StaleData field if non-nil, zero value otherwise.

### GetStaleDataOk

`func (o *ConsumptionStatusModel) GetStaleDataOk() (*bool, bool)`

GetStaleDataOk returns a tuple with the StaleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleData

`func (o *ConsumptionStatusModel) SetStaleData(v bool)`

SetStaleData sets StaleData field to given value.

### HasStaleData

`func (o *ConsumptionStatusModel) HasStaleData() bool`

HasStaleData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


