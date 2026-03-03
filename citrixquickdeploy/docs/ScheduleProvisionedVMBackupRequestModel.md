# ScheduleProvisionedVMBackupRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupName** | Pointer to **NullableString** | The name of the backup (required, max 64 characters) | [optional] 
**BackupDescription** | Pointer to **NullableString** | The description of the backup (optional, max 512 characters) | [optional] 
**DoNotShutdownMachine** | Pointer to **NullableBool** | Whether to shutdown the virtual machine when taking backup (optional) | [optional] 
**Description** | Pointer to **NullableString** | The description of the maintenance cycle (optional, max 512 characters) | [optional] 
**ScheduledStartTime** | Pointer to **NullableString** | Maintenance cycle start time. If not set, the operation will start immediately (optional)  RFC 3339 compatible format. | [optional] 
**MaxDurationInMinutes** | Pointer to **NullableInt32** | The max duration allowed for maintenance cycle in minutes (optional) | [optional] 
**PurgeDBAfterInDays** | Pointer to **NullableInt32** | Purge history data after a certain number of days (optional) | [optional] 
**SessionWarningLogOffTitle** | Pointer to **NullableString** | The title of session log off warning (optional, max 512 characters) | [optional] 
**SessionWarningLogOffMessage** | Pointer to **NullableString** | The message of session log off warning (optional, max 512 characters) | [optional] 
**SessionWarningTimeInMinutes** | Pointer to **NullableInt32** | The countdown duration for session log off in minutes (optional) | [optional] 

## Methods

### NewScheduleProvisionedVMBackupRequestModel

`func NewScheduleProvisionedVMBackupRequestModel() *ScheduleProvisionedVMBackupRequestModel`

NewScheduleProvisionedVMBackupRequestModel instantiates a new ScheduleProvisionedVMBackupRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleProvisionedVMBackupRequestModelWithDefaults

`func NewScheduleProvisionedVMBackupRequestModelWithDefaults() *ScheduleProvisionedVMBackupRequestModel`

NewScheduleProvisionedVMBackupRequestModelWithDefaults instantiates a new ScheduleProvisionedVMBackupRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupName

`func (o *ScheduleProvisionedVMBackupRequestModel) GetBackupName() string`

GetBackupName returns the BackupName field if non-nil, zero value otherwise.

### GetBackupNameOk

`func (o *ScheduleProvisionedVMBackupRequestModel) GetBackupNameOk() (*string, bool)`

GetBackupNameOk returns a tuple with the BackupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupName

`func (o *ScheduleProvisionedVMBackupRequestModel) SetBackupName(v string)`

SetBackupName sets BackupName field to given value.

### HasBackupName

`func (o *ScheduleProvisionedVMBackupRequestModel) HasBackupName() bool`

HasBackupName returns a boolean if a field has been set.

### SetBackupNameNil

`func (o *ScheduleProvisionedVMBackupRequestModel) SetBackupNameNil(b bool)`

 SetBackupNameNil sets the value for BackupName to be an explicit nil

### UnsetBackupName
`func (o *ScheduleProvisionedVMBackupRequestModel) UnsetBackupName()`

UnsetBackupName ensures that no value is present for BackupName, not even an explicit nil
### GetBackupDescription

`func (o *ScheduleProvisionedVMBackupRequestModel) GetBackupDescription() string`

GetBackupDescription returns the BackupDescription field if non-nil, zero value otherwise.

### GetBackupDescriptionOk

`func (o *ScheduleProvisionedVMBackupRequestModel) GetBackupDescriptionOk() (*string, bool)`

GetBackupDescriptionOk returns a tuple with the BackupDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDescription

`func (o *ScheduleProvisionedVMBackupRequestModel) SetBackupDescription(v string)`

SetBackupDescription sets BackupDescription field to given value.

### HasBackupDescription

`func (o *ScheduleProvisionedVMBackupRequestModel) HasBackupDescription() bool`

HasBackupDescription returns a boolean if a field has been set.

### SetBackupDescriptionNil

`func (o *ScheduleProvisionedVMBackupRequestModel) SetBackupDescriptionNil(b bool)`

 SetBackupDescriptionNil sets the value for BackupDescription to be an explicit nil

### UnsetBackupDescription
`func (o *ScheduleProvisionedVMBackupRequestModel) UnsetBackupDescription()`

UnsetBackupDescription ensures that no value is present for BackupDescription, not even an explicit nil
### GetDoNotShutdownMachine

`func (o *ScheduleProvisionedVMBackupRequestModel) GetDoNotShutdownMachine() bool`

GetDoNotShutdownMachine returns the DoNotShutdownMachine field if non-nil, zero value otherwise.

### GetDoNotShutdownMachineOk

`func (o *ScheduleProvisionedVMBackupRequestModel) GetDoNotShutdownMachineOk() (*bool, bool)`

GetDoNotShutdownMachineOk returns a tuple with the DoNotShutdownMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotShutdownMachine

`func (o *ScheduleProvisionedVMBackupRequestModel) SetDoNotShutdownMachine(v bool)`

SetDoNotShutdownMachine sets DoNotShutdownMachine field to given value.

### HasDoNotShutdownMachine

`func (o *ScheduleProvisionedVMBackupRequestModel) HasDoNotShutdownMachine() bool`

HasDoNotShutdownMachine returns a boolean if a field has been set.

### SetDoNotShutdownMachineNil

`func (o *ScheduleProvisionedVMBackupRequestModel) SetDoNotShutdownMachineNil(b bool)`

 SetDoNotShutdownMachineNil sets the value for DoNotShutdownMachine to be an explicit nil

### UnsetDoNotShutdownMachine
`func (o *ScheduleProvisionedVMBackupRequestModel) UnsetDoNotShutdownMachine()`

UnsetDoNotShutdownMachine ensures that no value is present for DoNotShutdownMachine, not even an explicit nil
### GetDescription

`func (o *ScheduleProvisionedVMBackupRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScheduleProvisionedVMBackupRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScheduleProvisionedVMBackupRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScheduleProvisionedVMBackupRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ScheduleProvisionedVMBackupRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ScheduleProvisionedVMBackupRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetScheduledStartTime

`func (o *ScheduleProvisionedVMBackupRequestModel) GetScheduledStartTime() string`

GetScheduledStartTime returns the ScheduledStartTime field if non-nil, zero value otherwise.

### GetScheduledStartTimeOk

`func (o *ScheduleProvisionedVMBackupRequestModel) GetScheduledStartTimeOk() (*string, bool)`

GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTime

`func (o *ScheduleProvisionedVMBackupRequestModel) SetScheduledStartTime(v string)`

SetScheduledStartTime sets ScheduledStartTime field to given value.

### HasScheduledStartTime

`func (o *ScheduleProvisionedVMBackupRequestModel) HasScheduledStartTime() bool`

HasScheduledStartTime returns a boolean if a field has been set.

### SetScheduledStartTimeNil

`func (o *ScheduleProvisionedVMBackupRequestModel) SetScheduledStartTimeNil(b bool)`

 SetScheduledStartTimeNil sets the value for ScheduledStartTime to be an explicit nil

### UnsetScheduledStartTime
`func (o *ScheduleProvisionedVMBackupRequestModel) UnsetScheduledStartTime()`

UnsetScheduledStartTime ensures that no value is present for ScheduledStartTime, not even an explicit nil
### GetMaxDurationInMinutes

`func (o *ScheduleProvisionedVMBackupRequestModel) GetMaxDurationInMinutes() int32`

GetMaxDurationInMinutes returns the MaxDurationInMinutes field if non-nil, zero value otherwise.

### GetMaxDurationInMinutesOk

`func (o *ScheduleProvisionedVMBackupRequestModel) GetMaxDurationInMinutesOk() (*int32, bool)`

GetMaxDurationInMinutesOk returns a tuple with the MaxDurationInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDurationInMinutes

`func (o *ScheduleProvisionedVMBackupRequestModel) SetMaxDurationInMinutes(v int32)`

SetMaxDurationInMinutes sets MaxDurationInMinutes field to given value.

### HasMaxDurationInMinutes

`func (o *ScheduleProvisionedVMBackupRequestModel) HasMaxDurationInMinutes() bool`

HasMaxDurationInMinutes returns a boolean if a field has been set.

### SetMaxDurationInMinutesNil

`func (o *ScheduleProvisionedVMBackupRequestModel) SetMaxDurationInMinutesNil(b bool)`

 SetMaxDurationInMinutesNil sets the value for MaxDurationInMinutes to be an explicit nil

### UnsetMaxDurationInMinutes
`func (o *ScheduleProvisionedVMBackupRequestModel) UnsetMaxDurationInMinutes()`

UnsetMaxDurationInMinutes ensures that no value is present for MaxDurationInMinutes, not even an explicit nil
### GetPurgeDBAfterInDays

`func (o *ScheduleProvisionedVMBackupRequestModel) GetPurgeDBAfterInDays() int32`

GetPurgeDBAfterInDays returns the PurgeDBAfterInDays field if non-nil, zero value otherwise.

### GetPurgeDBAfterInDaysOk

`func (o *ScheduleProvisionedVMBackupRequestModel) GetPurgeDBAfterInDaysOk() (*int32, bool)`

GetPurgeDBAfterInDaysOk returns a tuple with the PurgeDBAfterInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeDBAfterInDays

`func (o *ScheduleProvisionedVMBackupRequestModel) SetPurgeDBAfterInDays(v int32)`

SetPurgeDBAfterInDays sets PurgeDBAfterInDays field to given value.

### HasPurgeDBAfterInDays

`func (o *ScheduleProvisionedVMBackupRequestModel) HasPurgeDBAfterInDays() bool`

HasPurgeDBAfterInDays returns a boolean if a field has been set.

### SetPurgeDBAfterInDaysNil

`func (o *ScheduleProvisionedVMBackupRequestModel) SetPurgeDBAfterInDaysNil(b bool)`

 SetPurgeDBAfterInDaysNil sets the value for PurgeDBAfterInDays to be an explicit nil

### UnsetPurgeDBAfterInDays
`func (o *ScheduleProvisionedVMBackupRequestModel) UnsetPurgeDBAfterInDays()`

UnsetPurgeDBAfterInDays ensures that no value is present for PurgeDBAfterInDays, not even an explicit nil
### GetSessionWarningLogOffTitle

`func (o *ScheduleProvisionedVMBackupRequestModel) GetSessionWarningLogOffTitle() string`

GetSessionWarningLogOffTitle returns the SessionWarningLogOffTitle field if non-nil, zero value otherwise.

### GetSessionWarningLogOffTitleOk

`func (o *ScheduleProvisionedVMBackupRequestModel) GetSessionWarningLogOffTitleOk() (*string, bool)`

GetSessionWarningLogOffTitleOk returns a tuple with the SessionWarningLogOffTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionWarningLogOffTitle

`func (o *ScheduleProvisionedVMBackupRequestModel) SetSessionWarningLogOffTitle(v string)`

SetSessionWarningLogOffTitle sets SessionWarningLogOffTitle field to given value.

### HasSessionWarningLogOffTitle

`func (o *ScheduleProvisionedVMBackupRequestModel) HasSessionWarningLogOffTitle() bool`

HasSessionWarningLogOffTitle returns a boolean if a field has been set.

### SetSessionWarningLogOffTitleNil

`func (o *ScheduleProvisionedVMBackupRequestModel) SetSessionWarningLogOffTitleNil(b bool)`

 SetSessionWarningLogOffTitleNil sets the value for SessionWarningLogOffTitle to be an explicit nil

### UnsetSessionWarningLogOffTitle
`func (o *ScheduleProvisionedVMBackupRequestModel) UnsetSessionWarningLogOffTitle()`

UnsetSessionWarningLogOffTitle ensures that no value is present for SessionWarningLogOffTitle, not even an explicit nil
### GetSessionWarningLogOffMessage

`func (o *ScheduleProvisionedVMBackupRequestModel) GetSessionWarningLogOffMessage() string`

GetSessionWarningLogOffMessage returns the SessionWarningLogOffMessage field if non-nil, zero value otherwise.

### GetSessionWarningLogOffMessageOk

`func (o *ScheduleProvisionedVMBackupRequestModel) GetSessionWarningLogOffMessageOk() (*string, bool)`

GetSessionWarningLogOffMessageOk returns a tuple with the SessionWarningLogOffMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionWarningLogOffMessage

`func (o *ScheduleProvisionedVMBackupRequestModel) SetSessionWarningLogOffMessage(v string)`

SetSessionWarningLogOffMessage sets SessionWarningLogOffMessage field to given value.

### HasSessionWarningLogOffMessage

`func (o *ScheduleProvisionedVMBackupRequestModel) HasSessionWarningLogOffMessage() bool`

HasSessionWarningLogOffMessage returns a boolean if a field has been set.

### SetSessionWarningLogOffMessageNil

`func (o *ScheduleProvisionedVMBackupRequestModel) SetSessionWarningLogOffMessageNil(b bool)`

 SetSessionWarningLogOffMessageNil sets the value for SessionWarningLogOffMessage to be an explicit nil

### UnsetSessionWarningLogOffMessage
`func (o *ScheduleProvisionedVMBackupRequestModel) UnsetSessionWarningLogOffMessage()`

UnsetSessionWarningLogOffMessage ensures that no value is present for SessionWarningLogOffMessage, not even an explicit nil
### GetSessionWarningTimeInMinutes

`func (o *ScheduleProvisionedVMBackupRequestModel) GetSessionWarningTimeInMinutes() int32`

GetSessionWarningTimeInMinutes returns the SessionWarningTimeInMinutes field if non-nil, zero value otherwise.

### GetSessionWarningTimeInMinutesOk

`func (o *ScheduleProvisionedVMBackupRequestModel) GetSessionWarningTimeInMinutesOk() (*int32, bool)`

GetSessionWarningTimeInMinutesOk returns a tuple with the SessionWarningTimeInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionWarningTimeInMinutes

`func (o *ScheduleProvisionedVMBackupRequestModel) SetSessionWarningTimeInMinutes(v int32)`

SetSessionWarningTimeInMinutes sets SessionWarningTimeInMinutes field to given value.

### HasSessionWarningTimeInMinutes

`func (o *ScheduleProvisionedVMBackupRequestModel) HasSessionWarningTimeInMinutes() bool`

HasSessionWarningTimeInMinutes returns a boolean if a field has been set.

### SetSessionWarningTimeInMinutesNil

`func (o *ScheduleProvisionedVMBackupRequestModel) SetSessionWarningTimeInMinutesNil(b bool)`

 SetSessionWarningTimeInMinutesNil sets the value for SessionWarningTimeInMinutes to be an explicit nil

### UnsetSessionWarningTimeInMinutes
`func (o *ScheduleProvisionedVMBackupRequestModel) UnsetSessionWarningTimeInMinutes()`

UnsetSessionWarningTimeInMinutes ensures that no value is present for SessionWarningTimeInMinutes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


