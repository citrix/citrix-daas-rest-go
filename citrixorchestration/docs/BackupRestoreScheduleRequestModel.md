# BackupRestoreScheduleRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name | 
**Day** | Pointer to [**BackupRestoreScheduleDays**](BackupRestoreScheduleDays.md) |  | [optional] 
**DaysInWeek** | Pointer to [**[]BackupRestoreScheduleDays**](BackupRestoreScheduleDays.md) | Days in Week | [optional] 
**DayInMonth** | Pointer to [**BackupRestoreScheduleDays**](BackupRestoreScheduleDays.md) |  | [optional] 
**WeekInMonth** | Pointer to [**BackupRestoreScheduleWeeks**](BackupRestoreScheduleWeeks.md) |  | [optional] 
**StartDate** | **string** | Start Date | 
**FrequencyFactor** | **int32** | Frequency Factor | 
**Description** | Pointer to **string** | Description | [optional] 
**Enabled** | **bool** | Enabled              | 
**Frequency** | [**BackupRestoreScheduleFrequency**](BackupRestoreScheduleFrequency.md) |  | 
**StartTime** | **string** | Start Time | 

## Methods

### NewBackupRestoreScheduleRequestModel

`func NewBackupRestoreScheduleRequestModel(name string, startDate string, frequencyFactor int32, enabled bool, frequency BackupRestoreScheduleFrequency, startTime string, ) *BackupRestoreScheduleRequestModel`

NewBackupRestoreScheduleRequestModel instantiates a new BackupRestoreScheduleRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreScheduleRequestModelWithDefaults

`func NewBackupRestoreScheduleRequestModelWithDefaults() *BackupRestoreScheduleRequestModel`

NewBackupRestoreScheduleRequestModelWithDefaults instantiates a new BackupRestoreScheduleRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BackupRestoreScheduleRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupRestoreScheduleRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupRestoreScheduleRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetDay

`func (o *BackupRestoreScheduleRequestModel) GetDay() BackupRestoreScheduleDays`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *BackupRestoreScheduleRequestModel) GetDayOk() (*BackupRestoreScheduleDays, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *BackupRestoreScheduleRequestModel) SetDay(v BackupRestoreScheduleDays)`

SetDay sets Day field to given value.

### HasDay

`func (o *BackupRestoreScheduleRequestModel) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetDaysInWeek

`func (o *BackupRestoreScheduleRequestModel) GetDaysInWeek() []BackupRestoreScheduleDays`

GetDaysInWeek returns the DaysInWeek field if non-nil, zero value otherwise.

### GetDaysInWeekOk

`func (o *BackupRestoreScheduleRequestModel) GetDaysInWeekOk() (*[]BackupRestoreScheduleDays, bool)`

GetDaysInWeekOk returns a tuple with the DaysInWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysInWeek

`func (o *BackupRestoreScheduleRequestModel) SetDaysInWeek(v []BackupRestoreScheduleDays)`

SetDaysInWeek sets DaysInWeek field to given value.

### HasDaysInWeek

`func (o *BackupRestoreScheduleRequestModel) HasDaysInWeek() bool`

HasDaysInWeek returns a boolean if a field has been set.

### GetDayInMonth

`func (o *BackupRestoreScheduleRequestModel) GetDayInMonth() BackupRestoreScheduleDays`

GetDayInMonth returns the DayInMonth field if non-nil, zero value otherwise.

### GetDayInMonthOk

`func (o *BackupRestoreScheduleRequestModel) GetDayInMonthOk() (*BackupRestoreScheduleDays, bool)`

GetDayInMonthOk returns a tuple with the DayInMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayInMonth

`func (o *BackupRestoreScheduleRequestModel) SetDayInMonth(v BackupRestoreScheduleDays)`

SetDayInMonth sets DayInMonth field to given value.

### HasDayInMonth

`func (o *BackupRestoreScheduleRequestModel) HasDayInMonth() bool`

HasDayInMonth returns a boolean if a field has been set.

### GetWeekInMonth

`func (o *BackupRestoreScheduleRequestModel) GetWeekInMonth() BackupRestoreScheduleWeeks`

GetWeekInMonth returns the WeekInMonth field if non-nil, zero value otherwise.

### GetWeekInMonthOk

`func (o *BackupRestoreScheduleRequestModel) GetWeekInMonthOk() (*BackupRestoreScheduleWeeks, bool)`

GetWeekInMonthOk returns a tuple with the WeekInMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekInMonth

`func (o *BackupRestoreScheduleRequestModel) SetWeekInMonth(v BackupRestoreScheduleWeeks)`

SetWeekInMonth sets WeekInMonth field to given value.

### HasWeekInMonth

`func (o *BackupRestoreScheduleRequestModel) HasWeekInMonth() bool`

HasWeekInMonth returns a boolean if a field has been set.

### GetStartDate

`func (o *BackupRestoreScheduleRequestModel) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BackupRestoreScheduleRequestModel) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BackupRestoreScheduleRequestModel) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetFrequencyFactor

`func (o *BackupRestoreScheduleRequestModel) GetFrequencyFactor() int32`

GetFrequencyFactor returns the FrequencyFactor field if non-nil, zero value otherwise.

### GetFrequencyFactorOk

`func (o *BackupRestoreScheduleRequestModel) GetFrequencyFactorOk() (*int32, bool)`

GetFrequencyFactorOk returns a tuple with the FrequencyFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyFactor

`func (o *BackupRestoreScheduleRequestModel) SetFrequencyFactor(v int32)`

SetFrequencyFactor sets FrequencyFactor field to given value.


### GetDescription

`func (o *BackupRestoreScheduleRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BackupRestoreScheduleRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BackupRestoreScheduleRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BackupRestoreScheduleRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *BackupRestoreScheduleRequestModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BackupRestoreScheduleRequestModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BackupRestoreScheduleRequestModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFrequency

`func (o *BackupRestoreScheduleRequestModel) GetFrequency() BackupRestoreScheduleFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *BackupRestoreScheduleRequestModel) GetFrequencyOk() (*BackupRestoreScheduleFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *BackupRestoreScheduleRequestModel) SetFrequency(v BackupRestoreScheduleFrequency)`

SetFrequency sets Frequency field to given value.


### GetStartTime

`func (o *BackupRestoreScheduleRequestModel) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BackupRestoreScheduleRequestModel) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BackupRestoreScheduleRequestModel) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


