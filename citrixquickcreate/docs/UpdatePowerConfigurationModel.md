# UpdatePowerConfigurationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunningMode** | Pointer to [**NullableAwsEdcWorkspaceRunningMode**](AwsEdcWorkspaceRunningMode.md) |  | [optional] 
**AutoScaleEnabled** | Pointer to **NullableBool** | Whether auto-scale is enabled for the delivery group. | [optional] 
**OffPeakDisconnectTimeoutMinutes** | Pointer to **NullableInt32** | The number of minutes before the configured action should be  performed after a user session disconnects outside peak hours. | [optional] 
**OffPeakBufferSizePercent** | Pointer to **NullableInt32** | The percentage of machines in the delivery group that should be  kept available in an idle state outside peak hours. | [optional] 
**DisconnectOffPeakIdleSessionAfterSeconds** | Pointer to **NullableInt32** | Specifies the time in seconds after which an idle session  belonging to the delivery group is disconnected during off-peak time. | [optional] 
**TimeZone** | Pointer to **NullableString** | The time zone in which this delivery group&#39;s machines reside. | [optional] 
**OffPeakLogOffAction** | Pointer to [**NullableSessionChangeHostingAction**](SessionChangeHostingAction.md) |  | [optional] 
**OffPeakDisconnectAction** | Pointer to [**NullableSessionChangeHostingAction**](SessionChangeHostingAction.md) |  | [optional] 
**OffPeakLogOffTimeoutMinutes** | Pointer to **NullableInt32** | The number of minutes before the configured action should be  performed after a user session ends outside peak hours. | [optional] 

## Methods

### NewUpdatePowerConfigurationModel

`func NewUpdatePowerConfigurationModel() *UpdatePowerConfigurationModel`

NewUpdatePowerConfigurationModel instantiates a new UpdatePowerConfigurationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePowerConfigurationModelWithDefaults

`func NewUpdatePowerConfigurationModelWithDefaults() *UpdatePowerConfigurationModel`

NewUpdatePowerConfigurationModelWithDefaults instantiates a new UpdatePowerConfigurationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunningMode

`func (o *UpdatePowerConfigurationModel) GetRunningMode() AwsEdcWorkspaceRunningMode`

GetRunningMode returns the RunningMode field if non-nil, zero value otherwise.

### GetRunningModeOk

`func (o *UpdatePowerConfigurationModel) GetRunningModeOk() (*AwsEdcWorkspaceRunningMode, bool)`

GetRunningModeOk returns a tuple with the RunningMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningMode

`func (o *UpdatePowerConfigurationModel) SetRunningMode(v AwsEdcWorkspaceRunningMode)`

SetRunningMode sets RunningMode field to given value.

### HasRunningMode

`func (o *UpdatePowerConfigurationModel) HasRunningMode() bool`

HasRunningMode returns a boolean if a field has been set.

### SetRunningModeNil

`func (o *UpdatePowerConfigurationModel) SetRunningModeNil(b bool)`

 SetRunningModeNil sets the value for RunningMode to be an explicit nil

### UnsetRunningMode
`func (o *UpdatePowerConfigurationModel) UnsetRunningMode()`

UnsetRunningMode ensures that no value is present for RunningMode, not even an explicit nil
### GetAutoScaleEnabled

`func (o *UpdatePowerConfigurationModel) GetAutoScaleEnabled() bool`

GetAutoScaleEnabled returns the AutoScaleEnabled field if non-nil, zero value otherwise.

### GetAutoScaleEnabledOk

`func (o *UpdatePowerConfigurationModel) GetAutoScaleEnabledOk() (*bool, bool)`

GetAutoScaleEnabledOk returns a tuple with the AutoScaleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaleEnabled

`func (o *UpdatePowerConfigurationModel) SetAutoScaleEnabled(v bool)`

SetAutoScaleEnabled sets AutoScaleEnabled field to given value.

### HasAutoScaleEnabled

`func (o *UpdatePowerConfigurationModel) HasAutoScaleEnabled() bool`

HasAutoScaleEnabled returns a boolean if a field has been set.

### SetAutoScaleEnabledNil

`func (o *UpdatePowerConfigurationModel) SetAutoScaleEnabledNil(b bool)`

 SetAutoScaleEnabledNil sets the value for AutoScaleEnabled to be an explicit nil

### UnsetAutoScaleEnabled
`func (o *UpdatePowerConfigurationModel) UnsetAutoScaleEnabled()`

UnsetAutoScaleEnabled ensures that no value is present for AutoScaleEnabled, not even an explicit nil
### GetOffPeakDisconnectTimeoutMinutes

`func (o *UpdatePowerConfigurationModel) GetOffPeakDisconnectTimeoutMinutes() int32`

GetOffPeakDisconnectTimeoutMinutes returns the OffPeakDisconnectTimeoutMinutes field if non-nil, zero value otherwise.

### GetOffPeakDisconnectTimeoutMinutesOk

`func (o *UpdatePowerConfigurationModel) GetOffPeakDisconnectTimeoutMinutesOk() (*int32, bool)`

GetOffPeakDisconnectTimeoutMinutesOk returns a tuple with the OffPeakDisconnectTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakDisconnectTimeoutMinutes

`func (o *UpdatePowerConfigurationModel) SetOffPeakDisconnectTimeoutMinutes(v int32)`

SetOffPeakDisconnectTimeoutMinutes sets OffPeakDisconnectTimeoutMinutes field to given value.

### HasOffPeakDisconnectTimeoutMinutes

`func (o *UpdatePowerConfigurationModel) HasOffPeakDisconnectTimeoutMinutes() bool`

HasOffPeakDisconnectTimeoutMinutes returns a boolean if a field has been set.

### SetOffPeakDisconnectTimeoutMinutesNil

`func (o *UpdatePowerConfigurationModel) SetOffPeakDisconnectTimeoutMinutesNil(b bool)`

 SetOffPeakDisconnectTimeoutMinutesNil sets the value for OffPeakDisconnectTimeoutMinutes to be an explicit nil

### UnsetOffPeakDisconnectTimeoutMinutes
`func (o *UpdatePowerConfigurationModel) UnsetOffPeakDisconnectTimeoutMinutes()`

UnsetOffPeakDisconnectTimeoutMinutes ensures that no value is present for OffPeakDisconnectTimeoutMinutes, not even an explicit nil
### GetOffPeakBufferSizePercent

`func (o *UpdatePowerConfigurationModel) GetOffPeakBufferSizePercent() int32`

GetOffPeakBufferSizePercent returns the OffPeakBufferSizePercent field if non-nil, zero value otherwise.

### GetOffPeakBufferSizePercentOk

`func (o *UpdatePowerConfigurationModel) GetOffPeakBufferSizePercentOk() (*int32, bool)`

GetOffPeakBufferSizePercentOk returns a tuple with the OffPeakBufferSizePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakBufferSizePercent

`func (o *UpdatePowerConfigurationModel) SetOffPeakBufferSizePercent(v int32)`

SetOffPeakBufferSizePercent sets OffPeakBufferSizePercent field to given value.

### HasOffPeakBufferSizePercent

`func (o *UpdatePowerConfigurationModel) HasOffPeakBufferSizePercent() bool`

HasOffPeakBufferSizePercent returns a boolean if a field has been set.

### SetOffPeakBufferSizePercentNil

`func (o *UpdatePowerConfigurationModel) SetOffPeakBufferSizePercentNil(b bool)`

 SetOffPeakBufferSizePercentNil sets the value for OffPeakBufferSizePercent to be an explicit nil

### UnsetOffPeakBufferSizePercent
`func (o *UpdatePowerConfigurationModel) UnsetOffPeakBufferSizePercent()`

UnsetOffPeakBufferSizePercent ensures that no value is present for OffPeakBufferSizePercent, not even an explicit nil
### GetDisconnectOffPeakIdleSessionAfterSeconds

`func (o *UpdatePowerConfigurationModel) GetDisconnectOffPeakIdleSessionAfterSeconds() int32`

GetDisconnectOffPeakIdleSessionAfterSeconds returns the DisconnectOffPeakIdleSessionAfterSeconds field if non-nil, zero value otherwise.

### GetDisconnectOffPeakIdleSessionAfterSecondsOk

`func (o *UpdatePowerConfigurationModel) GetDisconnectOffPeakIdleSessionAfterSecondsOk() (*int32, bool)`

GetDisconnectOffPeakIdleSessionAfterSecondsOk returns a tuple with the DisconnectOffPeakIdleSessionAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectOffPeakIdleSessionAfterSeconds

`func (o *UpdatePowerConfigurationModel) SetDisconnectOffPeakIdleSessionAfterSeconds(v int32)`

SetDisconnectOffPeakIdleSessionAfterSeconds sets DisconnectOffPeakIdleSessionAfterSeconds field to given value.

### HasDisconnectOffPeakIdleSessionAfterSeconds

`func (o *UpdatePowerConfigurationModel) HasDisconnectOffPeakIdleSessionAfterSeconds() bool`

HasDisconnectOffPeakIdleSessionAfterSeconds returns a boolean if a field has been set.

### SetDisconnectOffPeakIdleSessionAfterSecondsNil

`func (o *UpdatePowerConfigurationModel) SetDisconnectOffPeakIdleSessionAfterSecondsNil(b bool)`

 SetDisconnectOffPeakIdleSessionAfterSecondsNil sets the value for DisconnectOffPeakIdleSessionAfterSeconds to be an explicit nil

### UnsetDisconnectOffPeakIdleSessionAfterSeconds
`func (o *UpdatePowerConfigurationModel) UnsetDisconnectOffPeakIdleSessionAfterSeconds()`

UnsetDisconnectOffPeakIdleSessionAfterSeconds ensures that no value is present for DisconnectOffPeakIdleSessionAfterSeconds, not even an explicit nil
### GetTimeZone

`func (o *UpdatePowerConfigurationModel) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *UpdatePowerConfigurationModel) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *UpdatePowerConfigurationModel) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *UpdatePowerConfigurationModel) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *UpdatePowerConfigurationModel) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *UpdatePowerConfigurationModel) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
### GetOffPeakLogOffAction

`func (o *UpdatePowerConfigurationModel) GetOffPeakLogOffAction() SessionChangeHostingAction`

GetOffPeakLogOffAction returns the OffPeakLogOffAction field if non-nil, zero value otherwise.

### GetOffPeakLogOffActionOk

`func (o *UpdatePowerConfigurationModel) GetOffPeakLogOffActionOk() (*SessionChangeHostingAction, bool)`

GetOffPeakLogOffActionOk returns a tuple with the OffPeakLogOffAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakLogOffAction

`func (o *UpdatePowerConfigurationModel) SetOffPeakLogOffAction(v SessionChangeHostingAction)`

SetOffPeakLogOffAction sets OffPeakLogOffAction field to given value.

### HasOffPeakLogOffAction

`func (o *UpdatePowerConfigurationModel) HasOffPeakLogOffAction() bool`

HasOffPeakLogOffAction returns a boolean if a field has been set.

### SetOffPeakLogOffActionNil

`func (o *UpdatePowerConfigurationModel) SetOffPeakLogOffActionNil(b bool)`

 SetOffPeakLogOffActionNil sets the value for OffPeakLogOffAction to be an explicit nil

### UnsetOffPeakLogOffAction
`func (o *UpdatePowerConfigurationModel) UnsetOffPeakLogOffAction()`

UnsetOffPeakLogOffAction ensures that no value is present for OffPeakLogOffAction, not even an explicit nil
### GetOffPeakDisconnectAction

`func (o *UpdatePowerConfigurationModel) GetOffPeakDisconnectAction() SessionChangeHostingAction`

GetOffPeakDisconnectAction returns the OffPeakDisconnectAction field if non-nil, zero value otherwise.

### GetOffPeakDisconnectActionOk

`func (o *UpdatePowerConfigurationModel) GetOffPeakDisconnectActionOk() (*SessionChangeHostingAction, bool)`

GetOffPeakDisconnectActionOk returns a tuple with the OffPeakDisconnectAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakDisconnectAction

`func (o *UpdatePowerConfigurationModel) SetOffPeakDisconnectAction(v SessionChangeHostingAction)`

SetOffPeakDisconnectAction sets OffPeakDisconnectAction field to given value.

### HasOffPeakDisconnectAction

`func (o *UpdatePowerConfigurationModel) HasOffPeakDisconnectAction() bool`

HasOffPeakDisconnectAction returns a boolean if a field has been set.

### SetOffPeakDisconnectActionNil

`func (o *UpdatePowerConfigurationModel) SetOffPeakDisconnectActionNil(b bool)`

 SetOffPeakDisconnectActionNil sets the value for OffPeakDisconnectAction to be an explicit nil

### UnsetOffPeakDisconnectAction
`func (o *UpdatePowerConfigurationModel) UnsetOffPeakDisconnectAction()`

UnsetOffPeakDisconnectAction ensures that no value is present for OffPeakDisconnectAction, not even an explicit nil
### GetOffPeakLogOffTimeoutMinutes

`func (o *UpdatePowerConfigurationModel) GetOffPeakLogOffTimeoutMinutes() int32`

GetOffPeakLogOffTimeoutMinutes returns the OffPeakLogOffTimeoutMinutes field if non-nil, zero value otherwise.

### GetOffPeakLogOffTimeoutMinutesOk

`func (o *UpdatePowerConfigurationModel) GetOffPeakLogOffTimeoutMinutesOk() (*int32, bool)`

GetOffPeakLogOffTimeoutMinutesOk returns a tuple with the OffPeakLogOffTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakLogOffTimeoutMinutes

`func (o *UpdatePowerConfigurationModel) SetOffPeakLogOffTimeoutMinutes(v int32)`

SetOffPeakLogOffTimeoutMinutes sets OffPeakLogOffTimeoutMinutes field to given value.

### HasOffPeakLogOffTimeoutMinutes

`func (o *UpdatePowerConfigurationModel) HasOffPeakLogOffTimeoutMinutes() bool`

HasOffPeakLogOffTimeoutMinutes returns a boolean if a field has been set.

### SetOffPeakLogOffTimeoutMinutesNil

`func (o *UpdatePowerConfigurationModel) SetOffPeakLogOffTimeoutMinutesNil(b bool)`

 SetOffPeakLogOffTimeoutMinutesNil sets the value for OffPeakLogOffTimeoutMinutes to be an explicit nil

### UnsetOffPeakLogOffTimeoutMinutes
`func (o *UpdatePowerConfigurationModel) UnsetOffPeakLogOffTimeoutMinutes()`

UnsetOffPeakLogOffTimeoutMinutes ensures that no value is present for OffPeakLogOffTimeoutMinutes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


