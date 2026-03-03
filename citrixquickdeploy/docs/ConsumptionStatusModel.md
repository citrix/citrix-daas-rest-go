# ConsumptionStatusModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonthlyConsumptionCommitment** | Pointer to **NullableFloat64** |  | [optional] 
**MonthlyConsumption** | Pointer to **NullableFloat64** |  | [optional] 
**TermConsumptionCommitment** | Pointer to **NullableFloat64** |  | [optional] 
**TermConsumption** | Pointer to **NullableFloat64** |  | [optional] 
**TotalConsumption** | Pointer to **float64** |  | [optional] 
**NumberOfLicensedUsers** | Pointer to **NullableInt32** |  | [optional] 
**ConsumptionSettings** | Pointer to [**NullableDataStoreConsumptionThresholdModel**](DataStoreConsumptionThresholdModel.md) |  | [optional] 
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

### SetMonthlyConsumptionCommitmentNil

`func (o *ConsumptionStatusModel) SetMonthlyConsumptionCommitmentNil(b bool)`

 SetMonthlyConsumptionCommitmentNil sets the value for MonthlyConsumptionCommitment to be an explicit nil

### UnsetMonthlyConsumptionCommitment
`func (o *ConsumptionStatusModel) UnsetMonthlyConsumptionCommitment()`

UnsetMonthlyConsumptionCommitment ensures that no value is present for MonthlyConsumptionCommitment, not even an explicit nil
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

### SetMonthlyConsumptionNil

`func (o *ConsumptionStatusModel) SetMonthlyConsumptionNil(b bool)`

 SetMonthlyConsumptionNil sets the value for MonthlyConsumption to be an explicit nil

### UnsetMonthlyConsumption
`func (o *ConsumptionStatusModel) UnsetMonthlyConsumption()`

UnsetMonthlyConsumption ensures that no value is present for MonthlyConsumption, not even an explicit nil
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

### SetTermConsumptionCommitmentNil

`func (o *ConsumptionStatusModel) SetTermConsumptionCommitmentNil(b bool)`

 SetTermConsumptionCommitmentNil sets the value for TermConsumptionCommitment to be an explicit nil

### UnsetTermConsumptionCommitment
`func (o *ConsumptionStatusModel) UnsetTermConsumptionCommitment()`

UnsetTermConsumptionCommitment ensures that no value is present for TermConsumptionCommitment, not even an explicit nil
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

### SetTermConsumptionNil

`func (o *ConsumptionStatusModel) SetTermConsumptionNil(b bool)`

 SetTermConsumptionNil sets the value for TermConsumption to be an explicit nil

### UnsetTermConsumption
`func (o *ConsumptionStatusModel) UnsetTermConsumption()`

UnsetTermConsumption ensures that no value is present for TermConsumption, not even an explicit nil
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

### SetNumberOfLicensedUsersNil

`func (o *ConsumptionStatusModel) SetNumberOfLicensedUsersNil(b bool)`

 SetNumberOfLicensedUsersNil sets the value for NumberOfLicensedUsers to be an explicit nil

### UnsetNumberOfLicensedUsers
`func (o *ConsumptionStatusModel) UnsetNumberOfLicensedUsers()`

UnsetNumberOfLicensedUsers ensures that no value is present for NumberOfLicensedUsers, not even an explicit nil
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

### SetConsumptionSettingsNil

`func (o *ConsumptionStatusModel) SetConsumptionSettingsNil(b bool)`

 SetConsumptionSettingsNil sets the value for ConsumptionSettings to be an explicit nil

### UnsetConsumptionSettings
`func (o *ConsumptionStatusModel) UnsetConsumptionSettings()`

UnsetConsumptionSettings ensures that no value is present for ConsumptionSettings, not even an explicit nil
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


