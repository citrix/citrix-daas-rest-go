# BackupRestoreScheduleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name | [optional] 
**Uid** | Pointer to **int32** | Id | [optional] 
**Day** | Pointer to [**BackupRestoreScheduleDays**](BackupRestoreScheduleDays.md) |  | [optional] 
**DaysInWeek** | Pointer to [**[]BackupRestoreScheduleDays**](BackupRestoreScheduleDays.md) | Days in Week | [optional] 
**DayInMonth** | Pointer to [**BackupRestoreScheduleDays**](BackupRestoreScheduleDays.md) |  | [optional] 
**WeekInMonth** | Pointer to [**BackupRestoreScheduleWeeks**](BackupRestoreScheduleWeeks.md) |  | [optional] 
**StartDate** | Pointer to **NullableString** | Start Date | [optional] 
**FrequencyFactor** | Pointer to **NullableInt32** | Frequency Factor | [optional] 
**Description** | Pointer to **NullableString** | Description | [optional] 
**Enabled** | Pointer to **bool** | Enabled              | [optional] 
**Frequency** | Pointer to [**BackupRestoreScheduleFrequency**](BackupRestoreScheduleFrequency.md) |  | [optional] 
**StartTime** | Pointer to **NullableString** | Start Time | [optional] 
**TimeZoneId** | Pointer to **NullableString** | Time Zone Id              | [optional] 

## Methods

### NewBackupRestoreScheduleModel

`func NewBackupRestoreScheduleModel() *BackupRestoreScheduleModel`

NewBackupRestoreScheduleModel instantiates a new BackupRestoreScheduleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreScheduleModelWithDefaults

`func NewBackupRestoreScheduleModelWithDefaults() *BackupRestoreScheduleModel`

NewBackupRestoreScheduleModelWithDefaults instantiates a new BackupRestoreScheduleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BackupRestoreScheduleModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupRestoreScheduleModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupRestoreScheduleModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BackupRestoreScheduleModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BackupRestoreScheduleModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BackupRestoreScheduleModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUid

`func (o *BackupRestoreScheduleModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *BackupRestoreScheduleModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *BackupRestoreScheduleModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *BackupRestoreScheduleModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetDay

`func (o *BackupRestoreScheduleModel) GetDay() BackupRestoreScheduleDays`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *BackupRestoreScheduleModel) GetDayOk() (*BackupRestoreScheduleDays, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *BackupRestoreScheduleModel) SetDay(v BackupRestoreScheduleDays)`

SetDay sets Day field to given value.

### HasDay

`func (o *BackupRestoreScheduleModel) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetDaysInWeek

`func (o *BackupRestoreScheduleModel) GetDaysInWeek() []BackupRestoreScheduleDays`

GetDaysInWeek returns the DaysInWeek field if non-nil, zero value otherwise.

### GetDaysInWeekOk

`func (o *BackupRestoreScheduleModel) GetDaysInWeekOk() (*[]BackupRestoreScheduleDays, bool)`

GetDaysInWeekOk returns a tuple with the DaysInWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysInWeek

`func (o *BackupRestoreScheduleModel) SetDaysInWeek(v []BackupRestoreScheduleDays)`

SetDaysInWeek sets DaysInWeek field to given value.

### HasDaysInWeek

`func (o *BackupRestoreScheduleModel) HasDaysInWeek() bool`

HasDaysInWeek returns a boolean if a field has been set.

### SetDaysInWeekNil

`func (o *BackupRestoreScheduleModel) SetDaysInWeekNil(b bool)`

 SetDaysInWeekNil sets the value for DaysInWeek to be an explicit nil

### UnsetDaysInWeek
`func (o *BackupRestoreScheduleModel) UnsetDaysInWeek()`

UnsetDaysInWeek ensures that no value is present for DaysInWeek, not even an explicit nil
### GetDayInMonth

`func (o *BackupRestoreScheduleModel) GetDayInMonth() BackupRestoreScheduleDays`

GetDayInMonth returns the DayInMonth field if non-nil, zero value otherwise.

### GetDayInMonthOk

`func (o *BackupRestoreScheduleModel) GetDayInMonthOk() (*BackupRestoreScheduleDays, bool)`

GetDayInMonthOk returns a tuple with the DayInMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayInMonth

`func (o *BackupRestoreScheduleModel) SetDayInMonth(v BackupRestoreScheduleDays)`

SetDayInMonth sets DayInMonth field to given value.

### HasDayInMonth

`func (o *BackupRestoreScheduleModel) HasDayInMonth() bool`

HasDayInMonth returns a boolean if a field has been set.

### GetWeekInMonth

`func (o *BackupRestoreScheduleModel) GetWeekInMonth() BackupRestoreScheduleWeeks`

GetWeekInMonth returns the WeekInMonth field if non-nil, zero value otherwise.

### GetWeekInMonthOk

`func (o *BackupRestoreScheduleModel) GetWeekInMonthOk() (*BackupRestoreScheduleWeeks, bool)`

GetWeekInMonthOk returns a tuple with the WeekInMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekInMonth

`func (o *BackupRestoreScheduleModel) SetWeekInMonth(v BackupRestoreScheduleWeeks)`

SetWeekInMonth sets WeekInMonth field to given value.

### HasWeekInMonth

`func (o *BackupRestoreScheduleModel) HasWeekInMonth() bool`

HasWeekInMonth returns a boolean if a field has been set.

### GetStartDate

`func (o *BackupRestoreScheduleModel) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BackupRestoreScheduleModel) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BackupRestoreScheduleModel) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BackupRestoreScheduleModel) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *BackupRestoreScheduleModel) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *BackupRestoreScheduleModel) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetFrequencyFactor

`func (o *BackupRestoreScheduleModel) GetFrequencyFactor() int32`

GetFrequencyFactor returns the FrequencyFactor field if non-nil, zero value otherwise.

### GetFrequencyFactorOk

`func (o *BackupRestoreScheduleModel) GetFrequencyFactorOk() (*int32, bool)`

GetFrequencyFactorOk returns a tuple with the FrequencyFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyFactor

`func (o *BackupRestoreScheduleModel) SetFrequencyFactor(v int32)`

SetFrequencyFactor sets FrequencyFactor field to given value.

### HasFrequencyFactor

`func (o *BackupRestoreScheduleModel) HasFrequencyFactor() bool`

HasFrequencyFactor returns a boolean if a field has been set.

### SetFrequencyFactorNil

`func (o *BackupRestoreScheduleModel) SetFrequencyFactorNil(b bool)`

 SetFrequencyFactorNil sets the value for FrequencyFactor to be an explicit nil

### UnsetFrequencyFactor
`func (o *BackupRestoreScheduleModel) UnsetFrequencyFactor()`

UnsetFrequencyFactor ensures that no value is present for FrequencyFactor, not even an explicit nil
### GetDescription

`func (o *BackupRestoreScheduleModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BackupRestoreScheduleModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BackupRestoreScheduleModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BackupRestoreScheduleModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BackupRestoreScheduleModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BackupRestoreScheduleModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *BackupRestoreScheduleModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BackupRestoreScheduleModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BackupRestoreScheduleModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BackupRestoreScheduleModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFrequency

`func (o *BackupRestoreScheduleModel) GetFrequency() BackupRestoreScheduleFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *BackupRestoreScheduleModel) GetFrequencyOk() (*BackupRestoreScheduleFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *BackupRestoreScheduleModel) SetFrequency(v BackupRestoreScheduleFrequency)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *BackupRestoreScheduleModel) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetStartTime

`func (o *BackupRestoreScheduleModel) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BackupRestoreScheduleModel) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BackupRestoreScheduleModel) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *BackupRestoreScheduleModel) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *BackupRestoreScheduleModel) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *BackupRestoreScheduleModel) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetTimeZoneId

`func (o *BackupRestoreScheduleModel) GetTimeZoneId() string`

GetTimeZoneId returns the TimeZoneId field if non-nil, zero value otherwise.

### GetTimeZoneIdOk

`func (o *BackupRestoreScheduleModel) GetTimeZoneIdOk() (*string, bool)`

GetTimeZoneIdOk returns a tuple with the TimeZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneId

`func (o *BackupRestoreScheduleModel) SetTimeZoneId(v string)`

SetTimeZoneId sets TimeZoneId field to given value.

### HasTimeZoneId

`func (o *BackupRestoreScheduleModel) HasTimeZoneId() bool`

HasTimeZoneId returns a boolean if a field has been set.

### SetTimeZoneIdNil

`func (o *BackupRestoreScheduleModel) SetTimeZoneIdNil(b bool)`

 SetTimeZoneIdNil sets the value for TimeZoneId to be an explicit nil

### UnsetTimeZoneId
`func (o *BackupRestoreScheduleModel) UnsetTimeZoneId()`

UnsetTimeZoneId ensures that no value is present for TimeZoneId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


