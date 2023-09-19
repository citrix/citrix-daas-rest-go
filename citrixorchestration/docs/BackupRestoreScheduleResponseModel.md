# BackupRestoreScheduleResponseModel

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
**FrequencyFactor** | Pointer to **int32** | Frequency Factor | [optional] 
**Description** | Pointer to **NullableString** | Description | [optional] 
**Enabled** | Pointer to **bool** | Enabled              | [optional] 
**Frequency** | Pointer to [**BackupRestoreScheduleFrequency**](BackupRestoreScheduleFrequency.md) |  | [optional] 
**StartTime** | Pointer to **NullableString** | Start Time | [optional] 
**TimeZoneId** | Pointer to **NullableString** | Time Zone Id              | [optional] 

## Methods

### NewBackupRestoreScheduleResponseModel

`func NewBackupRestoreScheduleResponseModel() *BackupRestoreScheduleResponseModel`

NewBackupRestoreScheduleResponseModel instantiates a new BackupRestoreScheduleResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreScheduleResponseModelWithDefaults

`func NewBackupRestoreScheduleResponseModelWithDefaults() *BackupRestoreScheduleResponseModel`

NewBackupRestoreScheduleResponseModelWithDefaults instantiates a new BackupRestoreScheduleResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BackupRestoreScheduleResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupRestoreScheduleResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupRestoreScheduleResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BackupRestoreScheduleResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BackupRestoreScheduleResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BackupRestoreScheduleResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUid

`func (o *BackupRestoreScheduleResponseModel) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *BackupRestoreScheduleResponseModel) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *BackupRestoreScheduleResponseModel) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *BackupRestoreScheduleResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetDay

`func (o *BackupRestoreScheduleResponseModel) GetDay() BackupRestoreScheduleDays`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *BackupRestoreScheduleResponseModel) GetDayOk() (*BackupRestoreScheduleDays, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *BackupRestoreScheduleResponseModel) SetDay(v BackupRestoreScheduleDays)`

SetDay sets Day field to given value.

### HasDay

`func (o *BackupRestoreScheduleResponseModel) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetDaysInWeek

`func (o *BackupRestoreScheduleResponseModel) GetDaysInWeek() []BackupRestoreScheduleDays`

GetDaysInWeek returns the DaysInWeek field if non-nil, zero value otherwise.

### GetDaysInWeekOk

`func (o *BackupRestoreScheduleResponseModel) GetDaysInWeekOk() (*[]BackupRestoreScheduleDays, bool)`

GetDaysInWeekOk returns a tuple with the DaysInWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysInWeek

`func (o *BackupRestoreScheduleResponseModel) SetDaysInWeek(v []BackupRestoreScheduleDays)`

SetDaysInWeek sets DaysInWeek field to given value.

### HasDaysInWeek

`func (o *BackupRestoreScheduleResponseModel) HasDaysInWeek() bool`

HasDaysInWeek returns a boolean if a field has been set.

### SetDaysInWeekNil

`func (o *BackupRestoreScheduleResponseModel) SetDaysInWeekNil(b bool)`

 SetDaysInWeekNil sets the value for DaysInWeek to be an explicit nil

### UnsetDaysInWeek
`func (o *BackupRestoreScheduleResponseModel) UnsetDaysInWeek()`

UnsetDaysInWeek ensures that no value is present for DaysInWeek, not even an explicit nil
### GetDayInMonth

`func (o *BackupRestoreScheduleResponseModel) GetDayInMonth() BackupRestoreScheduleDays`

GetDayInMonth returns the DayInMonth field if non-nil, zero value otherwise.

### GetDayInMonthOk

`func (o *BackupRestoreScheduleResponseModel) GetDayInMonthOk() (*BackupRestoreScheduleDays, bool)`

GetDayInMonthOk returns a tuple with the DayInMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayInMonth

`func (o *BackupRestoreScheduleResponseModel) SetDayInMonth(v BackupRestoreScheduleDays)`

SetDayInMonth sets DayInMonth field to given value.

### HasDayInMonth

`func (o *BackupRestoreScheduleResponseModel) HasDayInMonth() bool`

HasDayInMonth returns a boolean if a field has been set.

### GetWeekInMonth

`func (o *BackupRestoreScheduleResponseModel) GetWeekInMonth() BackupRestoreScheduleWeeks`

GetWeekInMonth returns the WeekInMonth field if non-nil, zero value otherwise.

### GetWeekInMonthOk

`func (o *BackupRestoreScheduleResponseModel) GetWeekInMonthOk() (*BackupRestoreScheduleWeeks, bool)`

GetWeekInMonthOk returns a tuple with the WeekInMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekInMonth

`func (o *BackupRestoreScheduleResponseModel) SetWeekInMonth(v BackupRestoreScheduleWeeks)`

SetWeekInMonth sets WeekInMonth field to given value.

### HasWeekInMonth

`func (o *BackupRestoreScheduleResponseModel) HasWeekInMonth() bool`

HasWeekInMonth returns a boolean if a field has been set.

### GetStartDate

`func (o *BackupRestoreScheduleResponseModel) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BackupRestoreScheduleResponseModel) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BackupRestoreScheduleResponseModel) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BackupRestoreScheduleResponseModel) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *BackupRestoreScheduleResponseModel) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *BackupRestoreScheduleResponseModel) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetFrequencyFactor

`func (o *BackupRestoreScheduleResponseModel) GetFrequencyFactor() int32`

GetFrequencyFactor returns the FrequencyFactor field if non-nil, zero value otherwise.

### GetFrequencyFactorOk

`func (o *BackupRestoreScheduleResponseModel) GetFrequencyFactorOk() (*int32, bool)`

GetFrequencyFactorOk returns a tuple with the FrequencyFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyFactor

`func (o *BackupRestoreScheduleResponseModel) SetFrequencyFactor(v int32)`

SetFrequencyFactor sets FrequencyFactor field to given value.

### HasFrequencyFactor

`func (o *BackupRestoreScheduleResponseModel) HasFrequencyFactor() bool`

HasFrequencyFactor returns a boolean if a field has been set.

### GetDescription

`func (o *BackupRestoreScheduleResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BackupRestoreScheduleResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BackupRestoreScheduleResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BackupRestoreScheduleResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BackupRestoreScheduleResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BackupRestoreScheduleResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *BackupRestoreScheduleResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BackupRestoreScheduleResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BackupRestoreScheduleResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BackupRestoreScheduleResponseModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFrequency

`func (o *BackupRestoreScheduleResponseModel) GetFrequency() BackupRestoreScheduleFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *BackupRestoreScheduleResponseModel) GetFrequencyOk() (*BackupRestoreScheduleFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *BackupRestoreScheduleResponseModel) SetFrequency(v BackupRestoreScheduleFrequency)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *BackupRestoreScheduleResponseModel) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetStartTime

`func (o *BackupRestoreScheduleResponseModel) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BackupRestoreScheduleResponseModel) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BackupRestoreScheduleResponseModel) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *BackupRestoreScheduleResponseModel) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *BackupRestoreScheduleResponseModel) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *BackupRestoreScheduleResponseModel) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetTimeZoneId

`func (o *BackupRestoreScheduleResponseModel) GetTimeZoneId() string`

GetTimeZoneId returns the TimeZoneId field if non-nil, zero value otherwise.

### GetTimeZoneIdOk

`func (o *BackupRestoreScheduleResponseModel) GetTimeZoneIdOk() (*string, bool)`

GetTimeZoneIdOk returns a tuple with the TimeZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneId

`func (o *BackupRestoreScheduleResponseModel) SetTimeZoneId(v string)`

SetTimeZoneId sets TimeZoneId field to given value.

### HasTimeZoneId

`func (o *BackupRestoreScheduleResponseModel) HasTimeZoneId() bool`

HasTimeZoneId returns a boolean if a field has been set.

### SetTimeZoneIdNil

`func (o *BackupRestoreScheduleResponseModel) SetTimeZoneIdNil(b bool)`

 SetTimeZoneIdNil sets the value for TimeZoneId to be an explicit nil

### UnsetTimeZoneId
`func (o *BackupRestoreScheduleResponseModel) UnsetTimeZoneId()`

UnsetTimeZoneId ensures that no value is present for TimeZoneId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


