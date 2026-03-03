# ScheduleProvisionedVMRestoreRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProvisionedBackupId** | Pointer to **string** | The id of the backup (required) | [optional] 
**ProvisionedBackupRestoreType** | Pointer to [**ProvisionedBackupRestoreType**](ProvisionedBackupRestoreType.md) | The type of how the backup being restored | [optional] 
**Description** | Pointer to **NullableString** | The description of the maintenance cycle (optional, max 512 characters) | [optional] 
**ScheduledStartTime** | Pointer to **NullableString** | Maintenance cycle start time. If not set, the operation will start immediately (optional)  RFC 3339 compatible format. | [optional] 
**MaxDurationInMinutes** | Pointer to **NullableInt32** | The max duration allowed for maintenance cycle in minutes (optional) | [optional] 
**PurgeDBAfterInDays** | Pointer to **NullableInt32** | Purge history data after a certain number of days (optional) | [optional] 
**SessionWarningLogOffTitle** | Pointer to **NullableString** | The title of session log off warning (optional, max 512 characters) | [optional] 
**SessionWarningLogOffMessage** | Pointer to **NullableString** | The message of session log off warning (optional, max 512 characters) | [optional] 
**SessionWarningTimeInMinutes** | Pointer to **NullableInt32** | The countdown duration for session log off in minutes (optional) | [optional] 

## Methods

### NewScheduleProvisionedVMRestoreRequestModel

`func NewScheduleProvisionedVMRestoreRequestModel() *ScheduleProvisionedVMRestoreRequestModel`

NewScheduleProvisionedVMRestoreRequestModel instantiates a new ScheduleProvisionedVMRestoreRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleProvisionedVMRestoreRequestModelWithDefaults

`func NewScheduleProvisionedVMRestoreRequestModelWithDefaults() *ScheduleProvisionedVMRestoreRequestModel`

NewScheduleProvisionedVMRestoreRequestModelWithDefaults instantiates a new ScheduleProvisionedVMRestoreRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvisionedBackupId

`func (o *ScheduleProvisionedVMRestoreRequestModel) GetProvisionedBackupId() string`

GetProvisionedBackupId returns the ProvisionedBackupId field if non-nil, zero value otherwise.

### GetProvisionedBackupIdOk

`func (o *ScheduleProvisionedVMRestoreRequestModel) GetProvisionedBackupIdOk() (*string, bool)`

GetProvisionedBackupIdOk returns a tuple with the ProvisionedBackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedBackupId

`func (o *ScheduleProvisionedVMRestoreRequestModel) SetProvisionedBackupId(v string)`

SetProvisionedBackupId sets ProvisionedBackupId field to given value.

### HasProvisionedBackupId

`func (o *ScheduleProvisionedVMRestoreRequestModel) HasProvisionedBackupId() bool`

HasProvisionedBackupId returns a boolean if a field has been set.

### GetProvisionedBackupRestoreType

`func (o *ScheduleProvisionedVMRestoreRequestModel) GetProvisionedBackupRestoreType() ProvisionedBackupRestoreType`

GetProvisionedBackupRestoreType returns the ProvisionedBackupRestoreType field if non-nil, zero value otherwise.

### GetProvisionedBackupRestoreTypeOk

`func (o *ScheduleProvisionedVMRestoreRequestModel) GetProvisionedBackupRestoreTypeOk() (*ProvisionedBackupRestoreType, bool)`

GetProvisionedBackupRestoreTypeOk returns a tuple with the ProvisionedBackupRestoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedBackupRestoreType

`func (o *ScheduleProvisionedVMRestoreRequestModel) SetProvisionedBackupRestoreType(v ProvisionedBackupRestoreType)`

SetProvisionedBackupRestoreType sets ProvisionedBackupRestoreType field to given value.

### HasProvisionedBackupRestoreType

`func (o *ScheduleProvisionedVMRestoreRequestModel) HasProvisionedBackupRestoreType() bool`

HasProvisionedBackupRestoreType returns a boolean if a field has been set.

### GetDescription

`func (o *ScheduleProvisionedVMRestoreRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScheduleProvisionedVMRestoreRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScheduleProvisionedVMRestoreRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScheduleProvisionedVMRestoreRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ScheduleProvisionedVMRestoreRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ScheduleProvisionedVMRestoreRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetScheduledStartTime

`func (o *ScheduleProvisionedVMRestoreRequestModel) GetScheduledStartTime() string`

GetScheduledStartTime returns the ScheduledStartTime field if non-nil, zero value otherwise.

### GetScheduledStartTimeOk

`func (o *ScheduleProvisionedVMRestoreRequestModel) GetScheduledStartTimeOk() (*string, bool)`

GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTime

`func (o *ScheduleProvisionedVMRestoreRequestModel) SetScheduledStartTime(v string)`

SetScheduledStartTime sets ScheduledStartTime field to given value.

### HasScheduledStartTime

`func (o *ScheduleProvisionedVMRestoreRequestModel) HasScheduledStartTime() bool`

HasScheduledStartTime returns a boolean if a field has been set.

### SetScheduledStartTimeNil

`func (o *ScheduleProvisionedVMRestoreRequestModel) SetScheduledStartTimeNil(b bool)`

 SetScheduledStartTimeNil sets the value for ScheduledStartTime to be an explicit nil

### UnsetScheduledStartTime
`func (o *ScheduleProvisionedVMRestoreRequestModel) UnsetScheduledStartTime()`

UnsetScheduledStartTime ensures that no value is present for ScheduledStartTime, not even an explicit nil
### GetMaxDurationInMinutes

`func (o *ScheduleProvisionedVMRestoreRequestModel) GetMaxDurationInMinutes() int32`

GetMaxDurationInMinutes returns the MaxDurationInMinutes field if non-nil, zero value otherwise.

### GetMaxDurationInMinutesOk

`func (o *ScheduleProvisionedVMRestoreRequestModel) GetMaxDurationInMinutesOk() (*int32, bool)`

GetMaxDurationInMinutesOk returns a tuple with the MaxDurationInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDurationInMinutes

`func (o *ScheduleProvisionedVMRestoreRequestModel) SetMaxDurationInMinutes(v int32)`

SetMaxDurationInMinutes sets MaxDurationInMinutes field to given value.

### HasMaxDurationInMinutes

`func (o *ScheduleProvisionedVMRestoreRequestModel) HasMaxDurationInMinutes() bool`

HasMaxDurationInMinutes returns a boolean if a field has been set.

### SetMaxDurationInMinutesNil

`func (o *ScheduleProvisionedVMRestoreRequestModel) SetMaxDurationInMinutesNil(b bool)`

 SetMaxDurationInMinutesNil sets the value for MaxDurationInMinutes to be an explicit nil

### UnsetMaxDurationInMinutes
`func (o *ScheduleProvisionedVMRestoreRequestModel) UnsetMaxDurationInMinutes()`

UnsetMaxDurationInMinutes ensures that no value is present for MaxDurationInMinutes, not even an explicit nil
### GetPurgeDBAfterInDays

`func (o *ScheduleProvisionedVMRestoreRequestModel) GetPurgeDBAfterInDays() int32`

GetPurgeDBAfterInDays returns the PurgeDBAfterInDays field if non-nil, zero value otherwise.

### GetPurgeDBAfterInDaysOk

`func (o *ScheduleProvisionedVMRestoreRequestModel) GetPurgeDBAfterInDaysOk() (*int32, bool)`

GetPurgeDBAfterInDaysOk returns a tuple with the PurgeDBAfterInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeDBAfterInDays

`func (o *ScheduleProvisionedVMRestoreRequestModel) SetPurgeDBAfterInDays(v int32)`

SetPurgeDBAfterInDays sets PurgeDBAfterInDays field to given value.

### HasPurgeDBAfterInDays

`func (o *ScheduleProvisionedVMRestoreRequestModel) HasPurgeDBAfterInDays() bool`

HasPurgeDBAfterInDays returns a boolean if a field has been set.

### SetPurgeDBAfterInDaysNil

`func (o *ScheduleProvisionedVMRestoreRequestModel) SetPurgeDBAfterInDaysNil(b bool)`

 SetPurgeDBAfterInDaysNil sets the value for PurgeDBAfterInDays to be an explicit nil

### UnsetPurgeDBAfterInDays
`func (o *ScheduleProvisionedVMRestoreRequestModel) UnsetPurgeDBAfterInDays()`

UnsetPurgeDBAfterInDays ensures that no value is present for PurgeDBAfterInDays, not even an explicit nil
### GetSessionWarningLogOffTitle

`func (o *ScheduleProvisionedVMRestoreRequestModel) GetSessionWarningLogOffTitle() string`

GetSessionWarningLogOffTitle returns the SessionWarningLogOffTitle field if non-nil, zero value otherwise.

### GetSessionWarningLogOffTitleOk

`func (o *ScheduleProvisionedVMRestoreRequestModel) GetSessionWarningLogOffTitleOk() (*string, bool)`

GetSessionWarningLogOffTitleOk returns a tuple with the SessionWarningLogOffTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionWarningLogOffTitle

`func (o *ScheduleProvisionedVMRestoreRequestModel) SetSessionWarningLogOffTitle(v string)`

SetSessionWarningLogOffTitle sets SessionWarningLogOffTitle field to given value.

### HasSessionWarningLogOffTitle

`func (o *ScheduleProvisionedVMRestoreRequestModel) HasSessionWarningLogOffTitle() bool`

HasSessionWarningLogOffTitle returns a boolean if a field has been set.

### SetSessionWarningLogOffTitleNil

`func (o *ScheduleProvisionedVMRestoreRequestModel) SetSessionWarningLogOffTitleNil(b bool)`

 SetSessionWarningLogOffTitleNil sets the value for SessionWarningLogOffTitle to be an explicit nil

### UnsetSessionWarningLogOffTitle
`func (o *ScheduleProvisionedVMRestoreRequestModel) UnsetSessionWarningLogOffTitle()`

UnsetSessionWarningLogOffTitle ensures that no value is present for SessionWarningLogOffTitle, not even an explicit nil
### GetSessionWarningLogOffMessage

`func (o *ScheduleProvisionedVMRestoreRequestModel) GetSessionWarningLogOffMessage() string`

GetSessionWarningLogOffMessage returns the SessionWarningLogOffMessage field if non-nil, zero value otherwise.

### GetSessionWarningLogOffMessageOk

`func (o *ScheduleProvisionedVMRestoreRequestModel) GetSessionWarningLogOffMessageOk() (*string, bool)`

GetSessionWarningLogOffMessageOk returns a tuple with the SessionWarningLogOffMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionWarningLogOffMessage

`func (o *ScheduleProvisionedVMRestoreRequestModel) SetSessionWarningLogOffMessage(v string)`

SetSessionWarningLogOffMessage sets SessionWarningLogOffMessage field to given value.

### HasSessionWarningLogOffMessage

`func (o *ScheduleProvisionedVMRestoreRequestModel) HasSessionWarningLogOffMessage() bool`

HasSessionWarningLogOffMessage returns a boolean if a field has been set.

### SetSessionWarningLogOffMessageNil

`func (o *ScheduleProvisionedVMRestoreRequestModel) SetSessionWarningLogOffMessageNil(b bool)`

 SetSessionWarningLogOffMessageNil sets the value for SessionWarningLogOffMessage to be an explicit nil

### UnsetSessionWarningLogOffMessage
`func (o *ScheduleProvisionedVMRestoreRequestModel) UnsetSessionWarningLogOffMessage()`

UnsetSessionWarningLogOffMessage ensures that no value is present for SessionWarningLogOffMessage, not even an explicit nil
### GetSessionWarningTimeInMinutes

`func (o *ScheduleProvisionedVMRestoreRequestModel) GetSessionWarningTimeInMinutes() int32`

GetSessionWarningTimeInMinutes returns the SessionWarningTimeInMinutes field if non-nil, zero value otherwise.

### GetSessionWarningTimeInMinutesOk

`func (o *ScheduleProvisionedVMRestoreRequestModel) GetSessionWarningTimeInMinutesOk() (*int32, bool)`

GetSessionWarningTimeInMinutesOk returns a tuple with the SessionWarningTimeInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionWarningTimeInMinutes

`func (o *ScheduleProvisionedVMRestoreRequestModel) SetSessionWarningTimeInMinutes(v int32)`

SetSessionWarningTimeInMinutes sets SessionWarningTimeInMinutes field to given value.

### HasSessionWarningTimeInMinutes

`func (o *ScheduleProvisionedVMRestoreRequestModel) HasSessionWarningTimeInMinutes() bool`

HasSessionWarningTimeInMinutes returns a boolean if a field has been set.

### SetSessionWarningTimeInMinutesNil

`func (o *ScheduleProvisionedVMRestoreRequestModel) SetSessionWarningTimeInMinutesNil(b bool)`

 SetSessionWarningTimeInMinutesNil sets the value for SessionWarningTimeInMinutes to be an explicit nil

### UnsetSessionWarningTimeInMinutes
`func (o *ScheduleProvisionedVMRestoreRequestModel) UnsetSessionWarningTimeInMinutes()`

UnsetSessionWarningTimeInMinutes ensures that no value is present for SessionWarningTimeInMinutes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


