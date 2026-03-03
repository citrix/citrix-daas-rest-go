# DataStoreCatalogPowerSchemeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemeName** | Pointer to **NullableString** | Name of the custom power scheme | [optional] 
**Weekdays** | Pointer to **map[string]bool** | Days of the week that are included in peak days | [optional] 
**PeakEndTime** | Pointer to **int32** | Hour of day when peak usage ends | [optional] 
**PeakStartTime** | Pointer to **int32** | Hour of day when peak usage begins | [optional] 
**PeakTimeZoneId** | Pointer to **NullableString** | ID of the peak usage Timezone | [optional] 
**PeakDisconnectedSessionTimeout** | Pointer to **int32** | Durring Peak Hours, the time before a disconnected session is shutdown | [optional] 
**OffPeakDisconnectedSessionTimeout** | Pointer to **int32** | During Off Peak Hours, the time before a disconnected session is shutdown | [optional] 
**MultiSessionDisconnectedSessionTimeout** | Pointer to **int32** | Minutes to wait for disconnected sessions to be logged off on multi-session VMs | [optional] 
**SessionTimeout** | Pointer to **NullableInt32** | Idle timeout for session in the catalog (in mins) | [optional] 
**BufferCapacity** | Pointer to **NullableInt32** | Percentage of buffer capacity | [optional] 
**PeakDisconnectedSessionAction** | Pointer to [**NullableSessionChangeHostingAction**](SessionChangeHostingAction.md) | During peak hours, the action to be taken on disconnected session | [optional] 
**OffPeakDisconnectedSessionAction** | Pointer to [**NullableSessionChangeHostingAction**](SessionChangeHostingAction.md) | During off peak hours, the action to be taken on disconnected session | [optional] 
**PeakExtendedDisconnectTimeoutMinutes** | Pointer to **NullableInt32** | TThe number of minutes before the second action (Shutdown) should be performed after a user session disconnects during peak hours. Used when peak disconnection action is Suspend | [optional] 
**OffPeakExtendedDisconnectTimeoutMinutes** | Pointer to **NullableInt32** | The number of minutes before the second action (Shutdown) should be performed after a user session disconnects outside peak hours. Used when off peak disconnect action is Suspend | [optional] 
**OffPeakBufferCapacity** | Pointer to **NullableInt32** | Percentage of buffer capacity during off-peak hours | [optional] 
**PowerOffDelay** | Pointer to **int32** | Amount of time to delay powering off machines with no active sessions | [optional] 

## Methods

### NewDataStoreCatalogPowerSchemeModel

`func NewDataStoreCatalogPowerSchemeModel() *DataStoreCatalogPowerSchemeModel`

NewDataStoreCatalogPowerSchemeModel instantiates a new DataStoreCatalogPowerSchemeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStoreCatalogPowerSchemeModelWithDefaults

`func NewDataStoreCatalogPowerSchemeModelWithDefaults() *DataStoreCatalogPowerSchemeModel`

NewDataStoreCatalogPowerSchemeModelWithDefaults instantiates a new DataStoreCatalogPowerSchemeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemeName

`func (o *DataStoreCatalogPowerSchemeModel) GetSchemeName() string`

GetSchemeName returns the SchemeName field if non-nil, zero value otherwise.

### GetSchemeNameOk

`func (o *DataStoreCatalogPowerSchemeModel) GetSchemeNameOk() (*string, bool)`

GetSchemeNameOk returns a tuple with the SchemeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeName

`func (o *DataStoreCatalogPowerSchemeModel) SetSchemeName(v string)`

SetSchemeName sets SchemeName field to given value.

### HasSchemeName

`func (o *DataStoreCatalogPowerSchemeModel) HasSchemeName() bool`

HasSchemeName returns a boolean if a field has been set.

### SetSchemeNameNil

`func (o *DataStoreCatalogPowerSchemeModel) SetSchemeNameNil(b bool)`

 SetSchemeNameNil sets the value for SchemeName to be an explicit nil

### UnsetSchemeName
`func (o *DataStoreCatalogPowerSchemeModel) UnsetSchemeName()`

UnsetSchemeName ensures that no value is present for SchemeName, not even an explicit nil
### GetWeekdays

`func (o *DataStoreCatalogPowerSchemeModel) GetWeekdays() map[string]bool`

GetWeekdays returns the Weekdays field if non-nil, zero value otherwise.

### GetWeekdaysOk

`func (o *DataStoreCatalogPowerSchemeModel) GetWeekdaysOk() (*map[string]bool, bool)`

GetWeekdaysOk returns a tuple with the Weekdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdays

`func (o *DataStoreCatalogPowerSchemeModel) SetWeekdays(v map[string]bool)`

SetWeekdays sets Weekdays field to given value.

### HasWeekdays

`func (o *DataStoreCatalogPowerSchemeModel) HasWeekdays() bool`

HasWeekdays returns a boolean if a field has been set.

### SetWeekdaysNil

`func (o *DataStoreCatalogPowerSchemeModel) SetWeekdaysNil(b bool)`

 SetWeekdaysNil sets the value for Weekdays to be an explicit nil

### UnsetWeekdays
`func (o *DataStoreCatalogPowerSchemeModel) UnsetWeekdays()`

UnsetWeekdays ensures that no value is present for Weekdays, not even an explicit nil
### GetPeakEndTime

`func (o *DataStoreCatalogPowerSchemeModel) GetPeakEndTime() int32`

GetPeakEndTime returns the PeakEndTime field if non-nil, zero value otherwise.

### GetPeakEndTimeOk

`func (o *DataStoreCatalogPowerSchemeModel) GetPeakEndTimeOk() (*int32, bool)`

GetPeakEndTimeOk returns a tuple with the PeakEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakEndTime

`func (o *DataStoreCatalogPowerSchemeModel) SetPeakEndTime(v int32)`

SetPeakEndTime sets PeakEndTime field to given value.

### HasPeakEndTime

`func (o *DataStoreCatalogPowerSchemeModel) HasPeakEndTime() bool`

HasPeakEndTime returns a boolean if a field has been set.

### GetPeakStartTime

`func (o *DataStoreCatalogPowerSchemeModel) GetPeakStartTime() int32`

GetPeakStartTime returns the PeakStartTime field if non-nil, zero value otherwise.

### GetPeakStartTimeOk

`func (o *DataStoreCatalogPowerSchemeModel) GetPeakStartTimeOk() (*int32, bool)`

GetPeakStartTimeOk returns a tuple with the PeakStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakStartTime

`func (o *DataStoreCatalogPowerSchemeModel) SetPeakStartTime(v int32)`

SetPeakStartTime sets PeakStartTime field to given value.

### HasPeakStartTime

`func (o *DataStoreCatalogPowerSchemeModel) HasPeakStartTime() bool`

HasPeakStartTime returns a boolean if a field has been set.

### GetPeakTimeZoneId

`func (o *DataStoreCatalogPowerSchemeModel) GetPeakTimeZoneId() string`

GetPeakTimeZoneId returns the PeakTimeZoneId field if non-nil, zero value otherwise.

### GetPeakTimeZoneIdOk

`func (o *DataStoreCatalogPowerSchemeModel) GetPeakTimeZoneIdOk() (*string, bool)`

GetPeakTimeZoneIdOk returns a tuple with the PeakTimeZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakTimeZoneId

`func (o *DataStoreCatalogPowerSchemeModel) SetPeakTimeZoneId(v string)`

SetPeakTimeZoneId sets PeakTimeZoneId field to given value.

### HasPeakTimeZoneId

`func (o *DataStoreCatalogPowerSchemeModel) HasPeakTimeZoneId() bool`

HasPeakTimeZoneId returns a boolean if a field has been set.

### SetPeakTimeZoneIdNil

`func (o *DataStoreCatalogPowerSchemeModel) SetPeakTimeZoneIdNil(b bool)`

 SetPeakTimeZoneIdNil sets the value for PeakTimeZoneId to be an explicit nil

### UnsetPeakTimeZoneId
`func (o *DataStoreCatalogPowerSchemeModel) UnsetPeakTimeZoneId()`

UnsetPeakTimeZoneId ensures that no value is present for PeakTimeZoneId, not even an explicit nil
### GetPeakDisconnectedSessionTimeout

`func (o *DataStoreCatalogPowerSchemeModel) GetPeakDisconnectedSessionTimeout() int32`

GetPeakDisconnectedSessionTimeout returns the PeakDisconnectedSessionTimeout field if non-nil, zero value otherwise.

### GetPeakDisconnectedSessionTimeoutOk

`func (o *DataStoreCatalogPowerSchemeModel) GetPeakDisconnectedSessionTimeoutOk() (*int32, bool)`

GetPeakDisconnectedSessionTimeoutOk returns a tuple with the PeakDisconnectedSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakDisconnectedSessionTimeout

`func (o *DataStoreCatalogPowerSchemeModel) SetPeakDisconnectedSessionTimeout(v int32)`

SetPeakDisconnectedSessionTimeout sets PeakDisconnectedSessionTimeout field to given value.

### HasPeakDisconnectedSessionTimeout

`func (o *DataStoreCatalogPowerSchemeModel) HasPeakDisconnectedSessionTimeout() bool`

HasPeakDisconnectedSessionTimeout returns a boolean if a field has been set.

### GetOffPeakDisconnectedSessionTimeout

`func (o *DataStoreCatalogPowerSchemeModel) GetOffPeakDisconnectedSessionTimeout() int32`

GetOffPeakDisconnectedSessionTimeout returns the OffPeakDisconnectedSessionTimeout field if non-nil, zero value otherwise.

### GetOffPeakDisconnectedSessionTimeoutOk

`func (o *DataStoreCatalogPowerSchemeModel) GetOffPeakDisconnectedSessionTimeoutOk() (*int32, bool)`

GetOffPeakDisconnectedSessionTimeoutOk returns a tuple with the OffPeakDisconnectedSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakDisconnectedSessionTimeout

`func (o *DataStoreCatalogPowerSchemeModel) SetOffPeakDisconnectedSessionTimeout(v int32)`

SetOffPeakDisconnectedSessionTimeout sets OffPeakDisconnectedSessionTimeout field to given value.

### HasOffPeakDisconnectedSessionTimeout

`func (o *DataStoreCatalogPowerSchemeModel) HasOffPeakDisconnectedSessionTimeout() bool`

HasOffPeakDisconnectedSessionTimeout returns a boolean if a field has been set.

### GetMultiSessionDisconnectedSessionTimeout

`func (o *DataStoreCatalogPowerSchemeModel) GetMultiSessionDisconnectedSessionTimeout() int32`

GetMultiSessionDisconnectedSessionTimeout returns the MultiSessionDisconnectedSessionTimeout field if non-nil, zero value otherwise.

### GetMultiSessionDisconnectedSessionTimeoutOk

`func (o *DataStoreCatalogPowerSchemeModel) GetMultiSessionDisconnectedSessionTimeoutOk() (*int32, bool)`

GetMultiSessionDisconnectedSessionTimeoutOk returns a tuple with the MultiSessionDisconnectedSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSessionDisconnectedSessionTimeout

`func (o *DataStoreCatalogPowerSchemeModel) SetMultiSessionDisconnectedSessionTimeout(v int32)`

SetMultiSessionDisconnectedSessionTimeout sets MultiSessionDisconnectedSessionTimeout field to given value.

### HasMultiSessionDisconnectedSessionTimeout

`func (o *DataStoreCatalogPowerSchemeModel) HasMultiSessionDisconnectedSessionTimeout() bool`

HasMultiSessionDisconnectedSessionTimeout returns a boolean if a field has been set.

### GetSessionTimeout

`func (o *DataStoreCatalogPowerSchemeModel) GetSessionTimeout() int32`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *DataStoreCatalogPowerSchemeModel) GetSessionTimeoutOk() (*int32, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *DataStoreCatalogPowerSchemeModel) SetSessionTimeout(v int32)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *DataStoreCatalogPowerSchemeModel) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.

### SetSessionTimeoutNil

`func (o *DataStoreCatalogPowerSchemeModel) SetSessionTimeoutNil(b bool)`

 SetSessionTimeoutNil sets the value for SessionTimeout to be an explicit nil

### UnsetSessionTimeout
`func (o *DataStoreCatalogPowerSchemeModel) UnsetSessionTimeout()`

UnsetSessionTimeout ensures that no value is present for SessionTimeout, not even an explicit nil
### GetBufferCapacity

`func (o *DataStoreCatalogPowerSchemeModel) GetBufferCapacity() int32`

GetBufferCapacity returns the BufferCapacity field if non-nil, zero value otherwise.

### GetBufferCapacityOk

`func (o *DataStoreCatalogPowerSchemeModel) GetBufferCapacityOk() (*int32, bool)`

GetBufferCapacityOk returns a tuple with the BufferCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferCapacity

`func (o *DataStoreCatalogPowerSchemeModel) SetBufferCapacity(v int32)`

SetBufferCapacity sets BufferCapacity field to given value.

### HasBufferCapacity

`func (o *DataStoreCatalogPowerSchemeModel) HasBufferCapacity() bool`

HasBufferCapacity returns a boolean if a field has been set.

### SetBufferCapacityNil

`func (o *DataStoreCatalogPowerSchemeModel) SetBufferCapacityNil(b bool)`

 SetBufferCapacityNil sets the value for BufferCapacity to be an explicit nil

### UnsetBufferCapacity
`func (o *DataStoreCatalogPowerSchemeModel) UnsetBufferCapacity()`

UnsetBufferCapacity ensures that no value is present for BufferCapacity, not even an explicit nil
### GetPeakDisconnectedSessionAction

`func (o *DataStoreCatalogPowerSchemeModel) GetPeakDisconnectedSessionAction() SessionChangeHostingAction`

GetPeakDisconnectedSessionAction returns the PeakDisconnectedSessionAction field if non-nil, zero value otherwise.

### GetPeakDisconnectedSessionActionOk

`func (o *DataStoreCatalogPowerSchemeModel) GetPeakDisconnectedSessionActionOk() (*SessionChangeHostingAction, bool)`

GetPeakDisconnectedSessionActionOk returns a tuple with the PeakDisconnectedSessionAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakDisconnectedSessionAction

`func (o *DataStoreCatalogPowerSchemeModel) SetPeakDisconnectedSessionAction(v SessionChangeHostingAction)`

SetPeakDisconnectedSessionAction sets PeakDisconnectedSessionAction field to given value.

### HasPeakDisconnectedSessionAction

`func (o *DataStoreCatalogPowerSchemeModel) HasPeakDisconnectedSessionAction() bool`

HasPeakDisconnectedSessionAction returns a boolean if a field has been set.

### SetPeakDisconnectedSessionActionNil

`func (o *DataStoreCatalogPowerSchemeModel) SetPeakDisconnectedSessionActionNil(b bool)`

 SetPeakDisconnectedSessionActionNil sets the value for PeakDisconnectedSessionAction to be an explicit nil

### UnsetPeakDisconnectedSessionAction
`func (o *DataStoreCatalogPowerSchemeModel) UnsetPeakDisconnectedSessionAction()`

UnsetPeakDisconnectedSessionAction ensures that no value is present for PeakDisconnectedSessionAction, not even an explicit nil
### GetOffPeakDisconnectedSessionAction

`func (o *DataStoreCatalogPowerSchemeModel) GetOffPeakDisconnectedSessionAction() SessionChangeHostingAction`

GetOffPeakDisconnectedSessionAction returns the OffPeakDisconnectedSessionAction field if non-nil, zero value otherwise.

### GetOffPeakDisconnectedSessionActionOk

`func (o *DataStoreCatalogPowerSchemeModel) GetOffPeakDisconnectedSessionActionOk() (*SessionChangeHostingAction, bool)`

GetOffPeakDisconnectedSessionActionOk returns a tuple with the OffPeakDisconnectedSessionAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakDisconnectedSessionAction

`func (o *DataStoreCatalogPowerSchemeModel) SetOffPeakDisconnectedSessionAction(v SessionChangeHostingAction)`

SetOffPeakDisconnectedSessionAction sets OffPeakDisconnectedSessionAction field to given value.

### HasOffPeakDisconnectedSessionAction

`func (o *DataStoreCatalogPowerSchemeModel) HasOffPeakDisconnectedSessionAction() bool`

HasOffPeakDisconnectedSessionAction returns a boolean if a field has been set.

### SetOffPeakDisconnectedSessionActionNil

`func (o *DataStoreCatalogPowerSchemeModel) SetOffPeakDisconnectedSessionActionNil(b bool)`

 SetOffPeakDisconnectedSessionActionNil sets the value for OffPeakDisconnectedSessionAction to be an explicit nil

### UnsetOffPeakDisconnectedSessionAction
`func (o *DataStoreCatalogPowerSchemeModel) UnsetOffPeakDisconnectedSessionAction()`

UnsetOffPeakDisconnectedSessionAction ensures that no value is present for OffPeakDisconnectedSessionAction, not even an explicit nil
### GetPeakExtendedDisconnectTimeoutMinutes

`func (o *DataStoreCatalogPowerSchemeModel) GetPeakExtendedDisconnectTimeoutMinutes() int32`

GetPeakExtendedDisconnectTimeoutMinutes returns the PeakExtendedDisconnectTimeoutMinutes field if non-nil, zero value otherwise.

### GetPeakExtendedDisconnectTimeoutMinutesOk

`func (o *DataStoreCatalogPowerSchemeModel) GetPeakExtendedDisconnectTimeoutMinutesOk() (*int32, bool)`

GetPeakExtendedDisconnectTimeoutMinutesOk returns a tuple with the PeakExtendedDisconnectTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakExtendedDisconnectTimeoutMinutes

`func (o *DataStoreCatalogPowerSchemeModel) SetPeakExtendedDisconnectTimeoutMinutes(v int32)`

SetPeakExtendedDisconnectTimeoutMinutes sets PeakExtendedDisconnectTimeoutMinutes field to given value.

### HasPeakExtendedDisconnectTimeoutMinutes

`func (o *DataStoreCatalogPowerSchemeModel) HasPeakExtendedDisconnectTimeoutMinutes() bool`

HasPeakExtendedDisconnectTimeoutMinutes returns a boolean if a field has been set.

### SetPeakExtendedDisconnectTimeoutMinutesNil

`func (o *DataStoreCatalogPowerSchemeModel) SetPeakExtendedDisconnectTimeoutMinutesNil(b bool)`

 SetPeakExtendedDisconnectTimeoutMinutesNil sets the value for PeakExtendedDisconnectTimeoutMinutes to be an explicit nil

### UnsetPeakExtendedDisconnectTimeoutMinutes
`func (o *DataStoreCatalogPowerSchemeModel) UnsetPeakExtendedDisconnectTimeoutMinutes()`

UnsetPeakExtendedDisconnectTimeoutMinutes ensures that no value is present for PeakExtendedDisconnectTimeoutMinutes, not even an explicit nil
### GetOffPeakExtendedDisconnectTimeoutMinutes

`func (o *DataStoreCatalogPowerSchemeModel) GetOffPeakExtendedDisconnectTimeoutMinutes() int32`

GetOffPeakExtendedDisconnectTimeoutMinutes returns the OffPeakExtendedDisconnectTimeoutMinutes field if non-nil, zero value otherwise.

### GetOffPeakExtendedDisconnectTimeoutMinutesOk

`func (o *DataStoreCatalogPowerSchemeModel) GetOffPeakExtendedDisconnectTimeoutMinutesOk() (*int32, bool)`

GetOffPeakExtendedDisconnectTimeoutMinutesOk returns a tuple with the OffPeakExtendedDisconnectTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakExtendedDisconnectTimeoutMinutes

`func (o *DataStoreCatalogPowerSchemeModel) SetOffPeakExtendedDisconnectTimeoutMinutes(v int32)`

SetOffPeakExtendedDisconnectTimeoutMinutes sets OffPeakExtendedDisconnectTimeoutMinutes field to given value.

### HasOffPeakExtendedDisconnectTimeoutMinutes

`func (o *DataStoreCatalogPowerSchemeModel) HasOffPeakExtendedDisconnectTimeoutMinutes() bool`

HasOffPeakExtendedDisconnectTimeoutMinutes returns a boolean if a field has been set.

### SetOffPeakExtendedDisconnectTimeoutMinutesNil

`func (o *DataStoreCatalogPowerSchemeModel) SetOffPeakExtendedDisconnectTimeoutMinutesNil(b bool)`

 SetOffPeakExtendedDisconnectTimeoutMinutesNil sets the value for OffPeakExtendedDisconnectTimeoutMinutes to be an explicit nil

### UnsetOffPeakExtendedDisconnectTimeoutMinutes
`func (o *DataStoreCatalogPowerSchemeModel) UnsetOffPeakExtendedDisconnectTimeoutMinutes()`

UnsetOffPeakExtendedDisconnectTimeoutMinutes ensures that no value is present for OffPeakExtendedDisconnectTimeoutMinutes, not even an explicit nil
### GetOffPeakBufferCapacity

`func (o *DataStoreCatalogPowerSchemeModel) GetOffPeakBufferCapacity() int32`

GetOffPeakBufferCapacity returns the OffPeakBufferCapacity field if non-nil, zero value otherwise.

### GetOffPeakBufferCapacityOk

`func (o *DataStoreCatalogPowerSchemeModel) GetOffPeakBufferCapacityOk() (*int32, bool)`

GetOffPeakBufferCapacityOk returns a tuple with the OffPeakBufferCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakBufferCapacity

`func (o *DataStoreCatalogPowerSchemeModel) SetOffPeakBufferCapacity(v int32)`

SetOffPeakBufferCapacity sets OffPeakBufferCapacity field to given value.

### HasOffPeakBufferCapacity

`func (o *DataStoreCatalogPowerSchemeModel) HasOffPeakBufferCapacity() bool`

HasOffPeakBufferCapacity returns a boolean if a field has been set.

### SetOffPeakBufferCapacityNil

`func (o *DataStoreCatalogPowerSchemeModel) SetOffPeakBufferCapacityNil(b bool)`

 SetOffPeakBufferCapacityNil sets the value for OffPeakBufferCapacity to be an explicit nil

### UnsetOffPeakBufferCapacity
`func (o *DataStoreCatalogPowerSchemeModel) UnsetOffPeakBufferCapacity()`

UnsetOffPeakBufferCapacity ensures that no value is present for OffPeakBufferCapacity, not even an explicit nil
### GetPowerOffDelay

`func (o *DataStoreCatalogPowerSchemeModel) GetPowerOffDelay() int32`

GetPowerOffDelay returns the PowerOffDelay field if non-nil, zero value otherwise.

### GetPowerOffDelayOk

`func (o *DataStoreCatalogPowerSchemeModel) GetPowerOffDelayOk() (*int32, bool)`

GetPowerOffDelayOk returns a tuple with the PowerOffDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOffDelay

`func (o *DataStoreCatalogPowerSchemeModel) SetPowerOffDelay(v int32)`

SetPowerOffDelay sets PowerOffDelay field to given value.

### HasPowerOffDelay

`func (o *DataStoreCatalogPowerSchemeModel) HasPowerOffDelay() bool`

HasPowerOffDelay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


