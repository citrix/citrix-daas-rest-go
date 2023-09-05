# RebootScheduleResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the reboot schedule. needs to be globally unique | 
**Name** | **string** | A friendly name for the reboot schedule. | 
**DeliveryGroup** | [**RebootScheduleResponseModelDeliveryGroup**](RebootScheduleResponseModelDeliveryGroup.md) |  | 
**RebootDurationMinutes** | **int32** | Approximate maximum number of minutes over which the scheduled reboot cycle runs. | 
**Day** | Pointer to [**RebootScheduleDays**](RebootScheduleDays.md) |  | [optional] 
**DaysInWeek** | Pointer to [**[]RebootScheduleDays**](RebootScheduleDays.md) | For weekly cycles, days of the week on which the scheduled reboot cycle starts. | [optional] 
**WeekInMonth** | Pointer to [**RebootScheduleWeeks**](RebootScheduleWeeks.md) |  | [optional] 
**DayInMonth** | Pointer to [**RebootScheduleDays**](RebootScheduleDays.md) |  | [optional] 
**Description** | Pointer to **string** | An optional description for the reboot schedule. | [optional] 
**Enabled** | **bool** | Indicates whether the reboot schedule is enabled. | 
**IgnoreMaintenanceMode** | **bool** | Indicates whether the reboot schedule is active for maintained machines. | 
**UseNaturalReboot** | Pointer to **bool** | Indicates whether the reboot will be a natural reboot, where the machines will be rebooted when they have no sessions. | [optional] 
**Frequency** | [**RebootScheduleFrequency**](RebootScheduleFrequency.md) |  | 
**FrequencyFactor** | Pointer to **int32** | The frequency factor for the schedule. | [optional] 
**RestrictToTag** | Pointer to [**RebootScheduleResponseModelRestrictToTag**](RebootScheduleResponseModelRestrictToTag.md) |  | [optional] 
**StartDate** | Pointer to **string** | The date on which the schedule is expected to firstly run. RFC 3339 compatible format. | [optional] 
**StartTime** | **string** | Time of day at which the scheduled reboot cycle starts. RFC 3339 compatible format. | 
**WarningDurationMinutes** | **int32** | Time prior to the start of a machine reboot at which a warning message is displayed to all the users on the machine. | 
**WarningMessage** | Pointer to **string** | Warning message displayed in user sessions on a machine scheduled for reboot. | [optional] 
**WarningRepeatIntervalMinutes** | **int32** | Time to wait after the previous reboot warning before displaying the warning message in all user sessions on that machine again. | 
**WarningTitle** | Pointer to **string** | The window title used when showing the warning message in user sessions on a machine scheduled for reboot. | [optional] 

## Methods

### NewRebootScheduleResponseModel

`func NewRebootScheduleResponseModel(id string, name string, deliveryGroup RebootScheduleResponseModelDeliveryGroup, rebootDurationMinutes int32, enabled bool, ignoreMaintenanceMode bool, frequency RebootScheduleFrequency, startTime string, warningDurationMinutes int32, warningRepeatIntervalMinutes int32, ) *RebootScheduleResponseModel`

NewRebootScheduleResponseModel instantiates a new RebootScheduleResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebootScheduleResponseModelWithDefaults

`func NewRebootScheduleResponseModelWithDefaults() *RebootScheduleResponseModel`

NewRebootScheduleResponseModelWithDefaults instantiates a new RebootScheduleResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RebootScheduleResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RebootScheduleResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RebootScheduleResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RebootScheduleResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RebootScheduleResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RebootScheduleResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetDeliveryGroup

`func (o *RebootScheduleResponseModel) GetDeliveryGroup() RebootScheduleResponseModelDeliveryGroup`

GetDeliveryGroup returns the DeliveryGroup field if non-nil, zero value otherwise.

### GetDeliveryGroupOk

`func (o *RebootScheduleResponseModel) GetDeliveryGroupOk() (*RebootScheduleResponseModelDeliveryGroup, bool)`

GetDeliveryGroupOk returns a tuple with the DeliveryGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGroup

`func (o *RebootScheduleResponseModel) SetDeliveryGroup(v RebootScheduleResponseModelDeliveryGroup)`

SetDeliveryGroup sets DeliveryGroup field to given value.


### GetRebootDurationMinutes

`func (o *RebootScheduleResponseModel) GetRebootDurationMinutes() int32`

GetRebootDurationMinutes returns the RebootDurationMinutes field if non-nil, zero value otherwise.

### GetRebootDurationMinutesOk

`func (o *RebootScheduleResponseModel) GetRebootDurationMinutesOk() (*int32, bool)`

GetRebootDurationMinutesOk returns a tuple with the RebootDurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootDurationMinutes

`func (o *RebootScheduleResponseModel) SetRebootDurationMinutes(v int32)`

SetRebootDurationMinutes sets RebootDurationMinutes field to given value.


### GetDay

`func (o *RebootScheduleResponseModel) GetDay() RebootScheduleDays`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *RebootScheduleResponseModel) GetDayOk() (*RebootScheduleDays, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *RebootScheduleResponseModel) SetDay(v RebootScheduleDays)`

SetDay sets Day field to given value.

### HasDay

`func (o *RebootScheduleResponseModel) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetDaysInWeek

`func (o *RebootScheduleResponseModel) GetDaysInWeek() []RebootScheduleDays`

GetDaysInWeek returns the DaysInWeek field if non-nil, zero value otherwise.

### GetDaysInWeekOk

`func (o *RebootScheduleResponseModel) GetDaysInWeekOk() (*[]RebootScheduleDays, bool)`

GetDaysInWeekOk returns a tuple with the DaysInWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysInWeek

`func (o *RebootScheduleResponseModel) SetDaysInWeek(v []RebootScheduleDays)`

SetDaysInWeek sets DaysInWeek field to given value.

### HasDaysInWeek

`func (o *RebootScheduleResponseModel) HasDaysInWeek() bool`

HasDaysInWeek returns a boolean if a field has been set.

### GetWeekInMonth

`func (o *RebootScheduleResponseModel) GetWeekInMonth() RebootScheduleWeeks`

GetWeekInMonth returns the WeekInMonth field if non-nil, zero value otherwise.

### GetWeekInMonthOk

`func (o *RebootScheduleResponseModel) GetWeekInMonthOk() (*RebootScheduleWeeks, bool)`

GetWeekInMonthOk returns a tuple with the WeekInMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekInMonth

`func (o *RebootScheduleResponseModel) SetWeekInMonth(v RebootScheduleWeeks)`

SetWeekInMonth sets WeekInMonth field to given value.

### HasWeekInMonth

`func (o *RebootScheduleResponseModel) HasWeekInMonth() bool`

HasWeekInMonth returns a boolean if a field has been set.

### GetDayInMonth

`func (o *RebootScheduleResponseModel) GetDayInMonth() RebootScheduleDays`

GetDayInMonth returns the DayInMonth field if non-nil, zero value otherwise.

### GetDayInMonthOk

`func (o *RebootScheduleResponseModel) GetDayInMonthOk() (*RebootScheduleDays, bool)`

GetDayInMonthOk returns a tuple with the DayInMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayInMonth

`func (o *RebootScheduleResponseModel) SetDayInMonth(v RebootScheduleDays)`

SetDayInMonth sets DayInMonth field to given value.

### HasDayInMonth

`func (o *RebootScheduleResponseModel) HasDayInMonth() bool`

HasDayInMonth returns a boolean if a field has been set.

### GetDescription

`func (o *RebootScheduleResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RebootScheduleResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RebootScheduleResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RebootScheduleResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *RebootScheduleResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RebootScheduleResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RebootScheduleResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIgnoreMaintenanceMode

`func (o *RebootScheduleResponseModel) GetIgnoreMaintenanceMode() bool`

GetIgnoreMaintenanceMode returns the IgnoreMaintenanceMode field if non-nil, zero value otherwise.

### GetIgnoreMaintenanceModeOk

`func (o *RebootScheduleResponseModel) GetIgnoreMaintenanceModeOk() (*bool, bool)`

GetIgnoreMaintenanceModeOk returns a tuple with the IgnoreMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreMaintenanceMode

`func (o *RebootScheduleResponseModel) SetIgnoreMaintenanceMode(v bool)`

SetIgnoreMaintenanceMode sets IgnoreMaintenanceMode field to given value.


### GetUseNaturalReboot

`func (o *RebootScheduleResponseModel) GetUseNaturalReboot() bool`

GetUseNaturalReboot returns the UseNaturalReboot field if non-nil, zero value otherwise.

### GetUseNaturalRebootOk

`func (o *RebootScheduleResponseModel) GetUseNaturalRebootOk() (*bool, bool)`

GetUseNaturalRebootOk returns a tuple with the UseNaturalReboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNaturalReboot

`func (o *RebootScheduleResponseModel) SetUseNaturalReboot(v bool)`

SetUseNaturalReboot sets UseNaturalReboot field to given value.

### HasUseNaturalReboot

`func (o *RebootScheduleResponseModel) HasUseNaturalReboot() bool`

HasUseNaturalReboot returns a boolean if a field has been set.

### GetFrequency

`func (o *RebootScheduleResponseModel) GetFrequency() RebootScheduleFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *RebootScheduleResponseModel) GetFrequencyOk() (*RebootScheduleFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *RebootScheduleResponseModel) SetFrequency(v RebootScheduleFrequency)`

SetFrequency sets Frequency field to given value.


### GetFrequencyFactor

`func (o *RebootScheduleResponseModel) GetFrequencyFactor() int32`

GetFrequencyFactor returns the FrequencyFactor field if non-nil, zero value otherwise.

### GetFrequencyFactorOk

`func (o *RebootScheduleResponseModel) GetFrequencyFactorOk() (*int32, bool)`

GetFrequencyFactorOk returns a tuple with the FrequencyFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyFactor

`func (o *RebootScheduleResponseModel) SetFrequencyFactor(v int32)`

SetFrequencyFactor sets FrequencyFactor field to given value.

### HasFrequencyFactor

`func (o *RebootScheduleResponseModel) HasFrequencyFactor() bool`

HasFrequencyFactor returns a boolean if a field has been set.

### GetRestrictToTag

`func (o *RebootScheduleResponseModel) GetRestrictToTag() RebootScheduleResponseModelRestrictToTag`

GetRestrictToTag returns the RestrictToTag field if non-nil, zero value otherwise.

### GetRestrictToTagOk

`func (o *RebootScheduleResponseModel) GetRestrictToTagOk() (*RebootScheduleResponseModelRestrictToTag, bool)`

GetRestrictToTagOk returns a tuple with the RestrictToTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictToTag

`func (o *RebootScheduleResponseModel) SetRestrictToTag(v RebootScheduleResponseModelRestrictToTag)`

SetRestrictToTag sets RestrictToTag field to given value.

### HasRestrictToTag

`func (o *RebootScheduleResponseModel) HasRestrictToTag() bool`

HasRestrictToTag returns a boolean if a field has been set.

### GetStartDate

`func (o *RebootScheduleResponseModel) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *RebootScheduleResponseModel) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *RebootScheduleResponseModel) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *RebootScheduleResponseModel) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStartTime

`func (o *RebootScheduleResponseModel) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RebootScheduleResponseModel) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RebootScheduleResponseModel) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetWarningDurationMinutes

`func (o *RebootScheduleResponseModel) GetWarningDurationMinutes() int32`

GetWarningDurationMinutes returns the WarningDurationMinutes field if non-nil, zero value otherwise.

### GetWarningDurationMinutesOk

`func (o *RebootScheduleResponseModel) GetWarningDurationMinutesOk() (*int32, bool)`

GetWarningDurationMinutesOk returns a tuple with the WarningDurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningDurationMinutes

`func (o *RebootScheduleResponseModel) SetWarningDurationMinutes(v int32)`

SetWarningDurationMinutes sets WarningDurationMinutes field to given value.


### GetWarningMessage

`func (o *RebootScheduleResponseModel) GetWarningMessage() string`

GetWarningMessage returns the WarningMessage field if non-nil, zero value otherwise.

### GetWarningMessageOk

`func (o *RebootScheduleResponseModel) GetWarningMessageOk() (*string, bool)`

GetWarningMessageOk returns a tuple with the WarningMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningMessage

`func (o *RebootScheduleResponseModel) SetWarningMessage(v string)`

SetWarningMessage sets WarningMessage field to given value.

### HasWarningMessage

`func (o *RebootScheduleResponseModel) HasWarningMessage() bool`

HasWarningMessage returns a boolean if a field has been set.

### GetWarningRepeatIntervalMinutes

`func (o *RebootScheduleResponseModel) GetWarningRepeatIntervalMinutes() int32`

GetWarningRepeatIntervalMinutes returns the WarningRepeatIntervalMinutes field if non-nil, zero value otherwise.

### GetWarningRepeatIntervalMinutesOk

`func (o *RebootScheduleResponseModel) GetWarningRepeatIntervalMinutesOk() (*int32, bool)`

GetWarningRepeatIntervalMinutesOk returns a tuple with the WarningRepeatIntervalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningRepeatIntervalMinutes

`func (o *RebootScheduleResponseModel) SetWarningRepeatIntervalMinutes(v int32)`

SetWarningRepeatIntervalMinutes sets WarningRepeatIntervalMinutes field to given value.


### GetWarningTitle

`func (o *RebootScheduleResponseModel) GetWarningTitle() string`

GetWarningTitle returns the WarningTitle field if non-nil, zero value otherwise.

### GetWarningTitleOk

`func (o *RebootScheduleResponseModel) GetWarningTitleOk() (*string, bool)`

GetWarningTitleOk returns a tuple with the WarningTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningTitle

`func (o *RebootScheduleResponseModel) SetWarningTitle(v string)`

SetWarningTitle sets WarningTitle field to given value.

### HasWarningTitle

`func (o *RebootScheduleResponseModel) HasWarningTitle() bool`

HasWarningTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


