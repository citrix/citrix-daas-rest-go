# RebootScheduleRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Id of the reboot schedule. needs to be globally unique | [optional] 
**Name** | Pointer to **NullableString** | A friendly name for the reboot schedule. | [optional] 
**RebootDurationMinutes** | Pointer to **NullableInt32** | Approximate maximum number of minutes over which the scheduled reboot cycle runs. | [optional] 
**Day** | Pointer to [**RebootScheduleDays**](RebootScheduleDays.md) |  | [optional] 
**DaysInWeek** | Pointer to [**[]RebootScheduleDays**](RebootScheduleDays.md) | For weekly cycles, days of the week on which the scheduled reboot cycle starts. | [optional] 
**WeekInMonth** | Pointer to [**RebootScheduleWeeks**](RebootScheduleWeeks.md) |  | [optional] 
**DayInMonth** | Pointer to [**RebootScheduleDays**](RebootScheduleDays.md) |  | [optional] 
**Description** | Pointer to **NullableString** | An optional description for the reboot schedule. | [optional] 
**Enabled** | Pointer to **NullableBool** | Indicates whether the reboot schedule is enabled. | [optional] 
**IgnoreMaintenanceMode** | Pointer to **NullableBool** | Indicates whether the reboot schedule is active for maintained machines. | [optional] 
**UseNaturalReboot** | Pointer to **NullableBool** | Indicates whether the reboot will be a natural reboot, where the machines will be rebooted when they have no sessions. Once UseNaturalReboot is set to true, RebootDurationMinutes won&#39;t have any effect. | [optional] 
**Frequency** | Pointer to [**RebootScheduleFrequency**](RebootScheduleFrequency.md) |  | [optional] 
**FrequencyFactor** | Pointer to **NullableInt32** | The frequency factor for the schedule. | [optional] 
**RestrictToTag** | Pointer to **NullableString** | If set, the reboot schedule only applies to machines in the delivery group with the specified tag. | [optional] 
**StartDate** | Pointer to **NullableString** | The date on which the schedule is expected to firstly run. RFC 3339 compatible format. | [optional] 
**StartTime** | Pointer to **NullableString** | Time of day at which the scheduled reboot cycle starts. RFC 3339 compatible format. | [optional] 
**WarningDurationMinutes** | Pointer to **NullableInt32** | Time prior to the start of a machine reboot at which a warning message is displayed to all the users on the machine. | [optional] 
**WarningMessage** | Pointer to **NullableString** | Warning message displayed in user sessions on a machine scheduled for reboot. | [optional] 
**WarningRepeatIntervalMinutes** | Pointer to **NullableInt32** | Time to wait after the previous reboot warning before displaying the warning message in all user sessions on that machine again. | [optional] 
**WarningTitle** | Pointer to **NullableString** | The window title used when showing the warning message in user sessions on a machine scheduled for reboot. | [optional] 

## Methods

### NewRebootScheduleRequestModel

`func NewRebootScheduleRequestModel() *RebootScheduleRequestModel`

NewRebootScheduleRequestModel instantiates a new RebootScheduleRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebootScheduleRequestModelWithDefaults

`func NewRebootScheduleRequestModelWithDefaults() *RebootScheduleRequestModel`

NewRebootScheduleRequestModelWithDefaults instantiates a new RebootScheduleRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RebootScheduleRequestModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RebootScheduleRequestModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RebootScheduleRequestModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RebootScheduleRequestModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *RebootScheduleRequestModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RebootScheduleRequestModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *RebootScheduleRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RebootScheduleRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RebootScheduleRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RebootScheduleRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RebootScheduleRequestModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RebootScheduleRequestModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRebootDurationMinutes

`func (o *RebootScheduleRequestModel) GetRebootDurationMinutes() int32`

GetRebootDurationMinutes returns the RebootDurationMinutes field if non-nil, zero value otherwise.

### GetRebootDurationMinutesOk

`func (o *RebootScheduleRequestModel) GetRebootDurationMinutesOk() (*int32, bool)`

GetRebootDurationMinutesOk returns a tuple with the RebootDurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootDurationMinutes

`func (o *RebootScheduleRequestModel) SetRebootDurationMinutes(v int32)`

SetRebootDurationMinutes sets RebootDurationMinutes field to given value.

### HasRebootDurationMinutes

`func (o *RebootScheduleRequestModel) HasRebootDurationMinutes() bool`

HasRebootDurationMinutes returns a boolean if a field has been set.

### SetRebootDurationMinutesNil

`func (o *RebootScheduleRequestModel) SetRebootDurationMinutesNil(b bool)`

 SetRebootDurationMinutesNil sets the value for RebootDurationMinutes to be an explicit nil

### UnsetRebootDurationMinutes
`func (o *RebootScheduleRequestModel) UnsetRebootDurationMinutes()`

UnsetRebootDurationMinutes ensures that no value is present for RebootDurationMinutes, not even an explicit nil
### GetDay

`func (o *RebootScheduleRequestModel) GetDay() RebootScheduleDays`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *RebootScheduleRequestModel) GetDayOk() (*RebootScheduleDays, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *RebootScheduleRequestModel) SetDay(v RebootScheduleDays)`

SetDay sets Day field to given value.

### HasDay

`func (o *RebootScheduleRequestModel) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetDaysInWeek

`func (o *RebootScheduleRequestModel) GetDaysInWeek() []RebootScheduleDays`

GetDaysInWeek returns the DaysInWeek field if non-nil, zero value otherwise.

### GetDaysInWeekOk

`func (o *RebootScheduleRequestModel) GetDaysInWeekOk() (*[]RebootScheduleDays, bool)`

GetDaysInWeekOk returns a tuple with the DaysInWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysInWeek

`func (o *RebootScheduleRequestModel) SetDaysInWeek(v []RebootScheduleDays)`

SetDaysInWeek sets DaysInWeek field to given value.

### HasDaysInWeek

`func (o *RebootScheduleRequestModel) HasDaysInWeek() bool`

HasDaysInWeek returns a boolean if a field has been set.

### SetDaysInWeekNil

`func (o *RebootScheduleRequestModel) SetDaysInWeekNil(b bool)`

 SetDaysInWeekNil sets the value for DaysInWeek to be an explicit nil

### UnsetDaysInWeek
`func (o *RebootScheduleRequestModel) UnsetDaysInWeek()`

UnsetDaysInWeek ensures that no value is present for DaysInWeek, not even an explicit nil
### GetWeekInMonth

`func (o *RebootScheduleRequestModel) GetWeekInMonth() RebootScheduleWeeks`

GetWeekInMonth returns the WeekInMonth field if non-nil, zero value otherwise.

### GetWeekInMonthOk

`func (o *RebootScheduleRequestModel) GetWeekInMonthOk() (*RebootScheduleWeeks, bool)`

GetWeekInMonthOk returns a tuple with the WeekInMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekInMonth

`func (o *RebootScheduleRequestModel) SetWeekInMonth(v RebootScheduleWeeks)`

SetWeekInMonth sets WeekInMonth field to given value.

### HasWeekInMonth

`func (o *RebootScheduleRequestModel) HasWeekInMonth() bool`

HasWeekInMonth returns a boolean if a field has been set.

### GetDayInMonth

`func (o *RebootScheduleRequestModel) GetDayInMonth() RebootScheduleDays`

GetDayInMonth returns the DayInMonth field if non-nil, zero value otherwise.

### GetDayInMonthOk

`func (o *RebootScheduleRequestModel) GetDayInMonthOk() (*RebootScheduleDays, bool)`

GetDayInMonthOk returns a tuple with the DayInMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayInMonth

`func (o *RebootScheduleRequestModel) SetDayInMonth(v RebootScheduleDays)`

SetDayInMonth sets DayInMonth field to given value.

### HasDayInMonth

`func (o *RebootScheduleRequestModel) HasDayInMonth() bool`

HasDayInMonth returns a boolean if a field has been set.

### GetDescription

`func (o *RebootScheduleRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RebootScheduleRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RebootScheduleRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RebootScheduleRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RebootScheduleRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RebootScheduleRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *RebootScheduleRequestModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RebootScheduleRequestModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RebootScheduleRequestModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RebootScheduleRequestModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *RebootScheduleRequestModel) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *RebootScheduleRequestModel) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetIgnoreMaintenanceMode

`func (o *RebootScheduleRequestModel) GetIgnoreMaintenanceMode() bool`

GetIgnoreMaintenanceMode returns the IgnoreMaintenanceMode field if non-nil, zero value otherwise.

### GetIgnoreMaintenanceModeOk

`func (o *RebootScheduleRequestModel) GetIgnoreMaintenanceModeOk() (*bool, bool)`

GetIgnoreMaintenanceModeOk returns a tuple with the IgnoreMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreMaintenanceMode

`func (o *RebootScheduleRequestModel) SetIgnoreMaintenanceMode(v bool)`

SetIgnoreMaintenanceMode sets IgnoreMaintenanceMode field to given value.

### HasIgnoreMaintenanceMode

`func (o *RebootScheduleRequestModel) HasIgnoreMaintenanceMode() bool`

HasIgnoreMaintenanceMode returns a boolean if a field has been set.

### SetIgnoreMaintenanceModeNil

`func (o *RebootScheduleRequestModel) SetIgnoreMaintenanceModeNil(b bool)`

 SetIgnoreMaintenanceModeNil sets the value for IgnoreMaintenanceMode to be an explicit nil

### UnsetIgnoreMaintenanceMode
`func (o *RebootScheduleRequestModel) UnsetIgnoreMaintenanceMode()`

UnsetIgnoreMaintenanceMode ensures that no value is present for IgnoreMaintenanceMode, not even an explicit nil
### GetUseNaturalReboot

`func (o *RebootScheduleRequestModel) GetUseNaturalReboot() bool`

GetUseNaturalReboot returns the UseNaturalReboot field if non-nil, zero value otherwise.

### GetUseNaturalRebootOk

`func (o *RebootScheduleRequestModel) GetUseNaturalRebootOk() (*bool, bool)`

GetUseNaturalRebootOk returns a tuple with the UseNaturalReboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNaturalReboot

`func (o *RebootScheduleRequestModel) SetUseNaturalReboot(v bool)`

SetUseNaturalReboot sets UseNaturalReboot field to given value.

### HasUseNaturalReboot

`func (o *RebootScheduleRequestModel) HasUseNaturalReboot() bool`

HasUseNaturalReboot returns a boolean if a field has been set.

### SetUseNaturalRebootNil

`func (o *RebootScheduleRequestModel) SetUseNaturalRebootNil(b bool)`

 SetUseNaturalRebootNil sets the value for UseNaturalReboot to be an explicit nil

### UnsetUseNaturalReboot
`func (o *RebootScheduleRequestModel) UnsetUseNaturalReboot()`

UnsetUseNaturalReboot ensures that no value is present for UseNaturalReboot, not even an explicit nil
### GetFrequency

`func (o *RebootScheduleRequestModel) GetFrequency() RebootScheduleFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *RebootScheduleRequestModel) GetFrequencyOk() (*RebootScheduleFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *RebootScheduleRequestModel) SetFrequency(v RebootScheduleFrequency)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *RebootScheduleRequestModel) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetFrequencyFactor

`func (o *RebootScheduleRequestModel) GetFrequencyFactor() int32`

GetFrequencyFactor returns the FrequencyFactor field if non-nil, zero value otherwise.

### GetFrequencyFactorOk

`func (o *RebootScheduleRequestModel) GetFrequencyFactorOk() (*int32, bool)`

GetFrequencyFactorOk returns a tuple with the FrequencyFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyFactor

`func (o *RebootScheduleRequestModel) SetFrequencyFactor(v int32)`

SetFrequencyFactor sets FrequencyFactor field to given value.

### HasFrequencyFactor

`func (o *RebootScheduleRequestModel) HasFrequencyFactor() bool`

HasFrequencyFactor returns a boolean if a field has been set.

### SetFrequencyFactorNil

`func (o *RebootScheduleRequestModel) SetFrequencyFactorNil(b bool)`

 SetFrequencyFactorNil sets the value for FrequencyFactor to be an explicit nil

### UnsetFrequencyFactor
`func (o *RebootScheduleRequestModel) UnsetFrequencyFactor()`

UnsetFrequencyFactor ensures that no value is present for FrequencyFactor, not even an explicit nil
### GetRestrictToTag

`func (o *RebootScheduleRequestModel) GetRestrictToTag() string`

GetRestrictToTag returns the RestrictToTag field if non-nil, zero value otherwise.

### GetRestrictToTagOk

`func (o *RebootScheduleRequestModel) GetRestrictToTagOk() (*string, bool)`

GetRestrictToTagOk returns a tuple with the RestrictToTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictToTag

`func (o *RebootScheduleRequestModel) SetRestrictToTag(v string)`

SetRestrictToTag sets RestrictToTag field to given value.

### HasRestrictToTag

`func (o *RebootScheduleRequestModel) HasRestrictToTag() bool`

HasRestrictToTag returns a boolean if a field has been set.

### SetRestrictToTagNil

`func (o *RebootScheduleRequestModel) SetRestrictToTagNil(b bool)`

 SetRestrictToTagNil sets the value for RestrictToTag to be an explicit nil

### UnsetRestrictToTag
`func (o *RebootScheduleRequestModel) UnsetRestrictToTag()`

UnsetRestrictToTag ensures that no value is present for RestrictToTag, not even an explicit nil
### GetStartDate

`func (o *RebootScheduleRequestModel) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *RebootScheduleRequestModel) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *RebootScheduleRequestModel) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *RebootScheduleRequestModel) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *RebootScheduleRequestModel) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *RebootScheduleRequestModel) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetStartTime

`func (o *RebootScheduleRequestModel) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RebootScheduleRequestModel) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RebootScheduleRequestModel) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *RebootScheduleRequestModel) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *RebootScheduleRequestModel) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *RebootScheduleRequestModel) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetWarningDurationMinutes

`func (o *RebootScheduleRequestModel) GetWarningDurationMinutes() int32`

GetWarningDurationMinutes returns the WarningDurationMinutes field if non-nil, zero value otherwise.

### GetWarningDurationMinutesOk

`func (o *RebootScheduleRequestModel) GetWarningDurationMinutesOk() (*int32, bool)`

GetWarningDurationMinutesOk returns a tuple with the WarningDurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningDurationMinutes

`func (o *RebootScheduleRequestModel) SetWarningDurationMinutes(v int32)`

SetWarningDurationMinutes sets WarningDurationMinutes field to given value.

### HasWarningDurationMinutes

`func (o *RebootScheduleRequestModel) HasWarningDurationMinutes() bool`

HasWarningDurationMinutes returns a boolean if a field has been set.

### SetWarningDurationMinutesNil

`func (o *RebootScheduleRequestModel) SetWarningDurationMinutesNil(b bool)`

 SetWarningDurationMinutesNil sets the value for WarningDurationMinutes to be an explicit nil

### UnsetWarningDurationMinutes
`func (o *RebootScheduleRequestModel) UnsetWarningDurationMinutes()`

UnsetWarningDurationMinutes ensures that no value is present for WarningDurationMinutes, not even an explicit nil
### GetWarningMessage

`func (o *RebootScheduleRequestModel) GetWarningMessage() string`

GetWarningMessage returns the WarningMessage field if non-nil, zero value otherwise.

### GetWarningMessageOk

`func (o *RebootScheduleRequestModel) GetWarningMessageOk() (*string, bool)`

GetWarningMessageOk returns a tuple with the WarningMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningMessage

`func (o *RebootScheduleRequestModel) SetWarningMessage(v string)`

SetWarningMessage sets WarningMessage field to given value.

### HasWarningMessage

`func (o *RebootScheduleRequestModel) HasWarningMessage() bool`

HasWarningMessage returns a boolean if a field has been set.

### SetWarningMessageNil

`func (o *RebootScheduleRequestModel) SetWarningMessageNil(b bool)`

 SetWarningMessageNil sets the value for WarningMessage to be an explicit nil

### UnsetWarningMessage
`func (o *RebootScheduleRequestModel) UnsetWarningMessage()`

UnsetWarningMessage ensures that no value is present for WarningMessage, not even an explicit nil
### GetWarningRepeatIntervalMinutes

`func (o *RebootScheduleRequestModel) GetWarningRepeatIntervalMinutes() int32`

GetWarningRepeatIntervalMinutes returns the WarningRepeatIntervalMinutes field if non-nil, zero value otherwise.

### GetWarningRepeatIntervalMinutesOk

`func (o *RebootScheduleRequestModel) GetWarningRepeatIntervalMinutesOk() (*int32, bool)`

GetWarningRepeatIntervalMinutesOk returns a tuple with the WarningRepeatIntervalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningRepeatIntervalMinutes

`func (o *RebootScheduleRequestModel) SetWarningRepeatIntervalMinutes(v int32)`

SetWarningRepeatIntervalMinutes sets WarningRepeatIntervalMinutes field to given value.

### HasWarningRepeatIntervalMinutes

`func (o *RebootScheduleRequestModel) HasWarningRepeatIntervalMinutes() bool`

HasWarningRepeatIntervalMinutes returns a boolean if a field has been set.

### SetWarningRepeatIntervalMinutesNil

`func (o *RebootScheduleRequestModel) SetWarningRepeatIntervalMinutesNil(b bool)`

 SetWarningRepeatIntervalMinutesNil sets the value for WarningRepeatIntervalMinutes to be an explicit nil

### UnsetWarningRepeatIntervalMinutes
`func (o *RebootScheduleRequestModel) UnsetWarningRepeatIntervalMinutes()`

UnsetWarningRepeatIntervalMinutes ensures that no value is present for WarningRepeatIntervalMinutes, not even an explicit nil
### GetWarningTitle

`func (o *RebootScheduleRequestModel) GetWarningTitle() string`

GetWarningTitle returns the WarningTitle field if non-nil, zero value otherwise.

### GetWarningTitleOk

`func (o *RebootScheduleRequestModel) GetWarningTitleOk() (*string, bool)`

GetWarningTitleOk returns a tuple with the WarningTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningTitle

`func (o *RebootScheduleRequestModel) SetWarningTitle(v string)`

SetWarningTitle sets WarningTitle field to given value.

### HasWarningTitle

`func (o *RebootScheduleRequestModel) HasWarningTitle() bool`

HasWarningTitle returns a boolean if a field has been set.

### SetWarningTitleNil

`func (o *RebootScheduleRequestModel) SetWarningTitleNil(b bool)`

 SetWarningTitleNil sets the value for WarningTitle to be an explicit nil

### UnsetWarningTitle
`func (o *RebootScheduleRequestModel) UnsetWarningTitle()`

UnsetWarningTitle ensures that no value is present for WarningTitle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


