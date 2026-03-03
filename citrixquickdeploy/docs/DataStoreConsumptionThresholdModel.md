# DataStoreConsumptionThresholdModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonthlyConsumptionCommitmentThreshold** | Pointer to **NullableInt32** | Threshold for monthly consumption commitment (in unit) | [optional] 
**MonthlyConsumptionAlertEnabled** | Pointer to **NullableBool** | Indicator for whether monthly consumption threshold alert has been enabled | [optional] 
**TermConsumptionCommitmentThreshold** | Pointer to **NullableFloat64** | Threshold for term consumption commitment (in percentage) | [optional] 
**TermConsumptionAlertEnabled** | Pointer to **NullableBool** | Indicator for whether term consumption threshold alert has been enabled | [optional] 

## Methods

### NewDataStoreConsumptionThresholdModel

`func NewDataStoreConsumptionThresholdModel() *DataStoreConsumptionThresholdModel`

NewDataStoreConsumptionThresholdModel instantiates a new DataStoreConsumptionThresholdModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStoreConsumptionThresholdModelWithDefaults

`func NewDataStoreConsumptionThresholdModelWithDefaults() *DataStoreConsumptionThresholdModel`

NewDataStoreConsumptionThresholdModelWithDefaults instantiates a new DataStoreConsumptionThresholdModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonthlyConsumptionCommitmentThreshold

`func (o *DataStoreConsumptionThresholdModel) GetMonthlyConsumptionCommitmentThreshold() int32`

GetMonthlyConsumptionCommitmentThreshold returns the MonthlyConsumptionCommitmentThreshold field if non-nil, zero value otherwise.

### GetMonthlyConsumptionCommitmentThresholdOk

`func (o *DataStoreConsumptionThresholdModel) GetMonthlyConsumptionCommitmentThresholdOk() (*int32, bool)`

GetMonthlyConsumptionCommitmentThresholdOk returns a tuple with the MonthlyConsumptionCommitmentThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyConsumptionCommitmentThreshold

`func (o *DataStoreConsumptionThresholdModel) SetMonthlyConsumptionCommitmentThreshold(v int32)`

SetMonthlyConsumptionCommitmentThreshold sets MonthlyConsumptionCommitmentThreshold field to given value.

### HasMonthlyConsumptionCommitmentThreshold

`func (o *DataStoreConsumptionThresholdModel) HasMonthlyConsumptionCommitmentThreshold() bool`

HasMonthlyConsumptionCommitmentThreshold returns a boolean if a field has been set.

### SetMonthlyConsumptionCommitmentThresholdNil

`func (o *DataStoreConsumptionThresholdModel) SetMonthlyConsumptionCommitmentThresholdNil(b bool)`

 SetMonthlyConsumptionCommitmentThresholdNil sets the value for MonthlyConsumptionCommitmentThreshold to be an explicit nil

### UnsetMonthlyConsumptionCommitmentThreshold
`func (o *DataStoreConsumptionThresholdModel) UnsetMonthlyConsumptionCommitmentThreshold()`

UnsetMonthlyConsumptionCommitmentThreshold ensures that no value is present for MonthlyConsumptionCommitmentThreshold, not even an explicit nil
### GetMonthlyConsumptionAlertEnabled

`func (o *DataStoreConsumptionThresholdModel) GetMonthlyConsumptionAlertEnabled() bool`

GetMonthlyConsumptionAlertEnabled returns the MonthlyConsumptionAlertEnabled field if non-nil, zero value otherwise.

### GetMonthlyConsumptionAlertEnabledOk

`func (o *DataStoreConsumptionThresholdModel) GetMonthlyConsumptionAlertEnabledOk() (*bool, bool)`

GetMonthlyConsumptionAlertEnabledOk returns a tuple with the MonthlyConsumptionAlertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyConsumptionAlertEnabled

`func (o *DataStoreConsumptionThresholdModel) SetMonthlyConsumptionAlertEnabled(v bool)`

SetMonthlyConsumptionAlertEnabled sets MonthlyConsumptionAlertEnabled field to given value.

### HasMonthlyConsumptionAlertEnabled

`func (o *DataStoreConsumptionThresholdModel) HasMonthlyConsumptionAlertEnabled() bool`

HasMonthlyConsumptionAlertEnabled returns a boolean if a field has been set.

### SetMonthlyConsumptionAlertEnabledNil

`func (o *DataStoreConsumptionThresholdModel) SetMonthlyConsumptionAlertEnabledNil(b bool)`

 SetMonthlyConsumptionAlertEnabledNil sets the value for MonthlyConsumptionAlertEnabled to be an explicit nil

### UnsetMonthlyConsumptionAlertEnabled
`func (o *DataStoreConsumptionThresholdModel) UnsetMonthlyConsumptionAlertEnabled()`

UnsetMonthlyConsumptionAlertEnabled ensures that no value is present for MonthlyConsumptionAlertEnabled, not even an explicit nil
### GetTermConsumptionCommitmentThreshold

`func (o *DataStoreConsumptionThresholdModel) GetTermConsumptionCommitmentThreshold() float64`

GetTermConsumptionCommitmentThreshold returns the TermConsumptionCommitmentThreshold field if non-nil, zero value otherwise.

### GetTermConsumptionCommitmentThresholdOk

`func (o *DataStoreConsumptionThresholdModel) GetTermConsumptionCommitmentThresholdOk() (*float64, bool)`

GetTermConsumptionCommitmentThresholdOk returns a tuple with the TermConsumptionCommitmentThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermConsumptionCommitmentThreshold

`func (o *DataStoreConsumptionThresholdModel) SetTermConsumptionCommitmentThreshold(v float64)`

SetTermConsumptionCommitmentThreshold sets TermConsumptionCommitmentThreshold field to given value.

### HasTermConsumptionCommitmentThreshold

`func (o *DataStoreConsumptionThresholdModel) HasTermConsumptionCommitmentThreshold() bool`

HasTermConsumptionCommitmentThreshold returns a boolean if a field has been set.

### SetTermConsumptionCommitmentThresholdNil

`func (o *DataStoreConsumptionThresholdModel) SetTermConsumptionCommitmentThresholdNil(b bool)`

 SetTermConsumptionCommitmentThresholdNil sets the value for TermConsumptionCommitmentThreshold to be an explicit nil

### UnsetTermConsumptionCommitmentThreshold
`func (o *DataStoreConsumptionThresholdModel) UnsetTermConsumptionCommitmentThreshold()`

UnsetTermConsumptionCommitmentThreshold ensures that no value is present for TermConsumptionCommitmentThreshold, not even an explicit nil
### GetTermConsumptionAlertEnabled

`func (o *DataStoreConsumptionThresholdModel) GetTermConsumptionAlertEnabled() bool`

GetTermConsumptionAlertEnabled returns the TermConsumptionAlertEnabled field if non-nil, zero value otherwise.

### GetTermConsumptionAlertEnabledOk

`func (o *DataStoreConsumptionThresholdModel) GetTermConsumptionAlertEnabledOk() (*bool, bool)`

GetTermConsumptionAlertEnabledOk returns a tuple with the TermConsumptionAlertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermConsumptionAlertEnabled

`func (o *DataStoreConsumptionThresholdModel) SetTermConsumptionAlertEnabled(v bool)`

SetTermConsumptionAlertEnabled sets TermConsumptionAlertEnabled field to given value.

### HasTermConsumptionAlertEnabled

`func (o *DataStoreConsumptionThresholdModel) HasTermConsumptionAlertEnabled() bool`

HasTermConsumptionAlertEnabled returns a boolean if a field has been set.

### SetTermConsumptionAlertEnabledNil

`func (o *DataStoreConsumptionThresholdModel) SetTermConsumptionAlertEnabledNil(b bool)`

 SetTermConsumptionAlertEnabledNil sets the value for TermConsumptionAlertEnabled to be an explicit nil

### UnsetTermConsumptionAlertEnabled
`func (o *DataStoreConsumptionThresholdModel) UnsetTermConsumptionAlertEnabled()`

UnsetTermConsumptionAlertEnabled ensures that no value is present for TermConsumptionAlertEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


