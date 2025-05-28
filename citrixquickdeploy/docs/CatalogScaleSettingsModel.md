# CatalogScaleSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxUsers** | Pointer to **int32** | Max number of concurrent settings for the catalog | [optional] 
**MinInstances** | Pointer to **int32** | Min number of active VMs for the catalog | [optional] 
**MaxInstances** | Pointer to **int32** | Number of VMs that will be provisioned for  this catalog | [optional] 
**PendingMaxInstances** | Pointer to **int32** | Number of VMs the admin would like the catalog changed to | [optional] 
**Weekdays** | Pointer to **map[string]bool** | Days of the week that are included in peak days | [optional] 
**WeekdaysString** | Pointer to **string** | Days of the week that are included in peak days | [optional] 
**PeakStartTime** | Pointer to **int32** | Hour of day when peak usage begins | [optional] 
**PeakEndTime** | Pointer to **int32** | Hour of day when peak usage ends | [optional] 
**PeakTimeZone** | Pointer to **string** | Display of the peak usage timezone | [optional] 
**PeakTimeZoneId** | Pointer to **string** | ID of the peak usage Timezone | [optional] 
**PeakMinInstances** | Pointer to **int32** | Min number of insances that should be running durring peak hours | [optional] 
**BufferCapacity** | Pointer to **int32** | Percentage of buffer capacity during peak hours | [optional] 
**OffPeakBufferCapacity** | Pointer to **int32** | Percentage of buffer capacity during off-peak hours | [optional] 
**ServiceAccount** | Pointer to **string** | The service account to use for modifying max instances | [optional] 
**ServiceAccountPassword** | Pointer to **string** | Password of the service account used to modify max instances | [optional] 
**IsSmartScaleDisabled** | Pointer to **bool** | Indicates if the smart scale settings are enabled or disabled | [optional] 
**PeakDisconnectedSessionTimeout** | Pointer to **int32** | During Peak Hours, the time before an action is taken on disconnected session | [optional] 
**PeakDisconnectedSessionAction** | Pointer to [**SessionChangeHostingAction**](SessionChangeHostingAction.md) | During peak hours, the action to be taken on disconnected session | [optional] 
**OffPeakDisconnectedSessionTimeout** | Pointer to **int32** | During Off Peak Hours, the time before an action is taken on disconnected session | [optional] 
**OffPeakDisconnectedSessionAction** | Pointer to [**SessionChangeHostingAction**](SessionChangeHostingAction.md) | During off peak hours, the action to be taken on disconnected session | [optional] 
**PeakExtendedDisconnectTimeoutMinutes** | Pointer to **int32** | The number of minutes before the second action (Shutdown) should be performed after a user session disconnects during peak hours. Used when peak disconnection action is Suspend | [optional] 
**OffPeakExtendedDisconnectTimeoutMinutes** | Pointer to **int32** | The number of minutes before the second action (Shutdown) should be performed after a user session disconnects outside peak hours. Used when off peak disconnect action is Suspend | [optional] 
**PowerOffDelay** | Pointer to **int32** | Amount of time to delay powering off machines with no active sessions | [optional] 

## Methods

### NewCatalogScaleSettingsModel

`func NewCatalogScaleSettingsModel() *CatalogScaleSettingsModel`

NewCatalogScaleSettingsModel instantiates a new CatalogScaleSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogScaleSettingsModelWithDefaults

`func NewCatalogScaleSettingsModelWithDefaults() *CatalogScaleSettingsModel`

NewCatalogScaleSettingsModelWithDefaults instantiates a new CatalogScaleSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxUsers

`func (o *CatalogScaleSettingsModel) GetMaxUsers() int32`

GetMaxUsers returns the MaxUsers field if non-nil, zero value otherwise.

### GetMaxUsersOk

`func (o *CatalogScaleSettingsModel) GetMaxUsersOk() (*int32, bool)`

GetMaxUsersOk returns a tuple with the MaxUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUsers

`func (o *CatalogScaleSettingsModel) SetMaxUsers(v int32)`

SetMaxUsers sets MaxUsers field to given value.

### HasMaxUsers

`func (o *CatalogScaleSettingsModel) HasMaxUsers() bool`

HasMaxUsers returns a boolean if a field has been set.

### GetMinInstances

`func (o *CatalogScaleSettingsModel) GetMinInstances() int32`

GetMinInstances returns the MinInstances field if non-nil, zero value otherwise.

### GetMinInstancesOk

`func (o *CatalogScaleSettingsModel) GetMinInstancesOk() (*int32, bool)`

GetMinInstancesOk returns a tuple with the MinInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInstances

`func (o *CatalogScaleSettingsModel) SetMinInstances(v int32)`

SetMinInstances sets MinInstances field to given value.

### HasMinInstances

`func (o *CatalogScaleSettingsModel) HasMinInstances() bool`

HasMinInstances returns a boolean if a field has been set.

### GetMaxInstances

`func (o *CatalogScaleSettingsModel) GetMaxInstances() int32`

GetMaxInstances returns the MaxInstances field if non-nil, zero value otherwise.

### GetMaxInstancesOk

`func (o *CatalogScaleSettingsModel) GetMaxInstancesOk() (*int32, bool)`

GetMaxInstancesOk returns a tuple with the MaxInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstances

`func (o *CatalogScaleSettingsModel) SetMaxInstances(v int32)`

SetMaxInstances sets MaxInstances field to given value.

### HasMaxInstances

`func (o *CatalogScaleSettingsModel) HasMaxInstances() bool`

HasMaxInstances returns a boolean if a field has been set.

### GetPendingMaxInstances

`func (o *CatalogScaleSettingsModel) GetPendingMaxInstances() int32`

GetPendingMaxInstances returns the PendingMaxInstances field if non-nil, zero value otherwise.

### GetPendingMaxInstancesOk

`func (o *CatalogScaleSettingsModel) GetPendingMaxInstancesOk() (*int32, bool)`

GetPendingMaxInstancesOk returns a tuple with the PendingMaxInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingMaxInstances

`func (o *CatalogScaleSettingsModel) SetPendingMaxInstances(v int32)`

SetPendingMaxInstances sets PendingMaxInstances field to given value.

### HasPendingMaxInstances

`func (o *CatalogScaleSettingsModel) HasPendingMaxInstances() bool`

HasPendingMaxInstances returns a boolean if a field has been set.

### GetWeekdays

`func (o *CatalogScaleSettingsModel) GetWeekdays() map[string]bool`

GetWeekdays returns the Weekdays field if non-nil, zero value otherwise.

### GetWeekdaysOk

`func (o *CatalogScaleSettingsModel) GetWeekdaysOk() (*map[string]bool, bool)`

GetWeekdaysOk returns a tuple with the Weekdays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdays

`func (o *CatalogScaleSettingsModel) SetWeekdays(v map[string]bool)`

SetWeekdays sets Weekdays field to given value.

### HasWeekdays

`func (o *CatalogScaleSettingsModel) HasWeekdays() bool`

HasWeekdays returns a boolean if a field has been set.

### GetWeekdaysString

`func (o *CatalogScaleSettingsModel) GetWeekdaysString() string`

GetWeekdaysString returns the WeekdaysString field if non-nil, zero value otherwise.

### GetWeekdaysStringOk

`func (o *CatalogScaleSettingsModel) GetWeekdaysStringOk() (*string, bool)`

GetWeekdaysStringOk returns a tuple with the WeekdaysString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekdaysString

`func (o *CatalogScaleSettingsModel) SetWeekdaysString(v string)`

SetWeekdaysString sets WeekdaysString field to given value.

### HasWeekdaysString

`func (o *CatalogScaleSettingsModel) HasWeekdaysString() bool`

HasWeekdaysString returns a boolean if a field has been set.

### GetPeakStartTime

`func (o *CatalogScaleSettingsModel) GetPeakStartTime() int32`

GetPeakStartTime returns the PeakStartTime field if non-nil, zero value otherwise.

### GetPeakStartTimeOk

`func (o *CatalogScaleSettingsModel) GetPeakStartTimeOk() (*int32, bool)`

GetPeakStartTimeOk returns a tuple with the PeakStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakStartTime

`func (o *CatalogScaleSettingsModel) SetPeakStartTime(v int32)`

SetPeakStartTime sets PeakStartTime field to given value.

### HasPeakStartTime

`func (o *CatalogScaleSettingsModel) HasPeakStartTime() bool`

HasPeakStartTime returns a boolean if a field has been set.

### GetPeakEndTime

`func (o *CatalogScaleSettingsModel) GetPeakEndTime() int32`

GetPeakEndTime returns the PeakEndTime field if non-nil, zero value otherwise.

### GetPeakEndTimeOk

`func (o *CatalogScaleSettingsModel) GetPeakEndTimeOk() (*int32, bool)`

GetPeakEndTimeOk returns a tuple with the PeakEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakEndTime

`func (o *CatalogScaleSettingsModel) SetPeakEndTime(v int32)`

SetPeakEndTime sets PeakEndTime field to given value.

### HasPeakEndTime

`func (o *CatalogScaleSettingsModel) HasPeakEndTime() bool`

HasPeakEndTime returns a boolean if a field has been set.

### GetPeakTimeZone

`func (o *CatalogScaleSettingsModel) GetPeakTimeZone() string`

GetPeakTimeZone returns the PeakTimeZone field if non-nil, zero value otherwise.

### GetPeakTimeZoneOk

`func (o *CatalogScaleSettingsModel) GetPeakTimeZoneOk() (*string, bool)`

GetPeakTimeZoneOk returns a tuple with the PeakTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakTimeZone

`func (o *CatalogScaleSettingsModel) SetPeakTimeZone(v string)`

SetPeakTimeZone sets PeakTimeZone field to given value.

### HasPeakTimeZone

`func (o *CatalogScaleSettingsModel) HasPeakTimeZone() bool`

HasPeakTimeZone returns a boolean if a field has been set.

### GetPeakTimeZoneId

`func (o *CatalogScaleSettingsModel) GetPeakTimeZoneId() string`

GetPeakTimeZoneId returns the PeakTimeZoneId field if non-nil, zero value otherwise.

### GetPeakTimeZoneIdOk

`func (o *CatalogScaleSettingsModel) GetPeakTimeZoneIdOk() (*string, bool)`

GetPeakTimeZoneIdOk returns a tuple with the PeakTimeZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakTimeZoneId

`func (o *CatalogScaleSettingsModel) SetPeakTimeZoneId(v string)`

SetPeakTimeZoneId sets PeakTimeZoneId field to given value.

### HasPeakTimeZoneId

`func (o *CatalogScaleSettingsModel) HasPeakTimeZoneId() bool`

HasPeakTimeZoneId returns a boolean if a field has been set.

### GetPeakMinInstances

`func (o *CatalogScaleSettingsModel) GetPeakMinInstances() int32`

GetPeakMinInstances returns the PeakMinInstances field if non-nil, zero value otherwise.

### GetPeakMinInstancesOk

`func (o *CatalogScaleSettingsModel) GetPeakMinInstancesOk() (*int32, bool)`

GetPeakMinInstancesOk returns a tuple with the PeakMinInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakMinInstances

`func (o *CatalogScaleSettingsModel) SetPeakMinInstances(v int32)`

SetPeakMinInstances sets PeakMinInstances field to given value.

### HasPeakMinInstances

`func (o *CatalogScaleSettingsModel) HasPeakMinInstances() bool`

HasPeakMinInstances returns a boolean if a field has been set.

### GetBufferCapacity

`func (o *CatalogScaleSettingsModel) GetBufferCapacity() int32`

GetBufferCapacity returns the BufferCapacity field if non-nil, zero value otherwise.

### GetBufferCapacityOk

`func (o *CatalogScaleSettingsModel) GetBufferCapacityOk() (*int32, bool)`

GetBufferCapacityOk returns a tuple with the BufferCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferCapacity

`func (o *CatalogScaleSettingsModel) SetBufferCapacity(v int32)`

SetBufferCapacity sets BufferCapacity field to given value.

### HasBufferCapacity

`func (o *CatalogScaleSettingsModel) HasBufferCapacity() bool`

HasBufferCapacity returns a boolean if a field has been set.

### GetOffPeakBufferCapacity

`func (o *CatalogScaleSettingsModel) GetOffPeakBufferCapacity() int32`

GetOffPeakBufferCapacity returns the OffPeakBufferCapacity field if non-nil, zero value otherwise.

### GetOffPeakBufferCapacityOk

`func (o *CatalogScaleSettingsModel) GetOffPeakBufferCapacityOk() (*int32, bool)`

GetOffPeakBufferCapacityOk returns a tuple with the OffPeakBufferCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakBufferCapacity

`func (o *CatalogScaleSettingsModel) SetOffPeakBufferCapacity(v int32)`

SetOffPeakBufferCapacity sets OffPeakBufferCapacity field to given value.

### HasOffPeakBufferCapacity

`func (o *CatalogScaleSettingsModel) HasOffPeakBufferCapacity() bool`

HasOffPeakBufferCapacity returns a boolean if a field has been set.

### GetServiceAccount

`func (o *CatalogScaleSettingsModel) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *CatalogScaleSettingsModel) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *CatalogScaleSettingsModel) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *CatalogScaleSettingsModel) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetServiceAccountPassword

`func (o *CatalogScaleSettingsModel) GetServiceAccountPassword() string`

GetServiceAccountPassword returns the ServiceAccountPassword field if non-nil, zero value otherwise.

### GetServiceAccountPasswordOk

`func (o *CatalogScaleSettingsModel) GetServiceAccountPasswordOk() (*string, bool)`

GetServiceAccountPasswordOk returns a tuple with the ServiceAccountPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountPassword

`func (o *CatalogScaleSettingsModel) SetServiceAccountPassword(v string)`

SetServiceAccountPassword sets ServiceAccountPassword field to given value.

### HasServiceAccountPassword

`func (o *CatalogScaleSettingsModel) HasServiceAccountPassword() bool`

HasServiceAccountPassword returns a boolean if a field has been set.

### GetIsSmartScaleDisabled

`func (o *CatalogScaleSettingsModel) GetIsSmartScaleDisabled() bool`

GetIsSmartScaleDisabled returns the IsSmartScaleDisabled field if non-nil, zero value otherwise.

### GetIsSmartScaleDisabledOk

`func (o *CatalogScaleSettingsModel) GetIsSmartScaleDisabledOk() (*bool, bool)`

GetIsSmartScaleDisabledOk returns a tuple with the IsSmartScaleDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSmartScaleDisabled

`func (o *CatalogScaleSettingsModel) SetIsSmartScaleDisabled(v bool)`

SetIsSmartScaleDisabled sets IsSmartScaleDisabled field to given value.

### HasIsSmartScaleDisabled

`func (o *CatalogScaleSettingsModel) HasIsSmartScaleDisabled() bool`

HasIsSmartScaleDisabled returns a boolean if a field has been set.

### GetPeakDisconnectedSessionTimeout

`func (o *CatalogScaleSettingsModel) GetPeakDisconnectedSessionTimeout() int32`

GetPeakDisconnectedSessionTimeout returns the PeakDisconnectedSessionTimeout field if non-nil, zero value otherwise.

### GetPeakDisconnectedSessionTimeoutOk

`func (o *CatalogScaleSettingsModel) GetPeakDisconnectedSessionTimeoutOk() (*int32, bool)`

GetPeakDisconnectedSessionTimeoutOk returns a tuple with the PeakDisconnectedSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakDisconnectedSessionTimeout

`func (o *CatalogScaleSettingsModel) SetPeakDisconnectedSessionTimeout(v int32)`

SetPeakDisconnectedSessionTimeout sets PeakDisconnectedSessionTimeout field to given value.

### HasPeakDisconnectedSessionTimeout

`func (o *CatalogScaleSettingsModel) HasPeakDisconnectedSessionTimeout() bool`

HasPeakDisconnectedSessionTimeout returns a boolean if a field has been set.

### GetPeakDisconnectedSessionAction

`func (o *CatalogScaleSettingsModel) GetPeakDisconnectedSessionAction() SessionChangeHostingAction`

GetPeakDisconnectedSessionAction returns the PeakDisconnectedSessionAction field if non-nil, zero value otherwise.

### GetPeakDisconnectedSessionActionOk

`func (o *CatalogScaleSettingsModel) GetPeakDisconnectedSessionActionOk() (*SessionChangeHostingAction, bool)`

GetPeakDisconnectedSessionActionOk returns a tuple with the PeakDisconnectedSessionAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakDisconnectedSessionAction

`func (o *CatalogScaleSettingsModel) SetPeakDisconnectedSessionAction(v SessionChangeHostingAction)`

SetPeakDisconnectedSessionAction sets PeakDisconnectedSessionAction field to given value.

### HasPeakDisconnectedSessionAction

`func (o *CatalogScaleSettingsModel) HasPeakDisconnectedSessionAction() bool`

HasPeakDisconnectedSessionAction returns a boolean if a field has been set.

### GetOffPeakDisconnectedSessionTimeout

`func (o *CatalogScaleSettingsModel) GetOffPeakDisconnectedSessionTimeout() int32`

GetOffPeakDisconnectedSessionTimeout returns the OffPeakDisconnectedSessionTimeout field if non-nil, zero value otherwise.

### GetOffPeakDisconnectedSessionTimeoutOk

`func (o *CatalogScaleSettingsModel) GetOffPeakDisconnectedSessionTimeoutOk() (*int32, bool)`

GetOffPeakDisconnectedSessionTimeoutOk returns a tuple with the OffPeakDisconnectedSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakDisconnectedSessionTimeout

`func (o *CatalogScaleSettingsModel) SetOffPeakDisconnectedSessionTimeout(v int32)`

SetOffPeakDisconnectedSessionTimeout sets OffPeakDisconnectedSessionTimeout field to given value.

### HasOffPeakDisconnectedSessionTimeout

`func (o *CatalogScaleSettingsModel) HasOffPeakDisconnectedSessionTimeout() bool`

HasOffPeakDisconnectedSessionTimeout returns a boolean if a field has been set.

### GetOffPeakDisconnectedSessionAction

`func (o *CatalogScaleSettingsModel) GetOffPeakDisconnectedSessionAction() SessionChangeHostingAction`

GetOffPeakDisconnectedSessionAction returns the OffPeakDisconnectedSessionAction field if non-nil, zero value otherwise.

### GetOffPeakDisconnectedSessionActionOk

`func (o *CatalogScaleSettingsModel) GetOffPeakDisconnectedSessionActionOk() (*SessionChangeHostingAction, bool)`

GetOffPeakDisconnectedSessionActionOk returns a tuple with the OffPeakDisconnectedSessionAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakDisconnectedSessionAction

`func (o *CatalogScaleSettingsModel) SetOffPeakDisconnectedSessionAction(v SessionChangeHostingAction)`

SetOffPeakDisconnectedSessionAction sets OffPeakDisconnectedSessionAction field to given value.

### HasOffPeakDisconnectedSessionAction

`func (o *CatalogScaleSettingsModel) HasOffPeakDisconnectedSessionAction() bool`

HasOffPeakDisconnectedSessionAction returns a boolean if a field has been set.

### GetPeakExtendedDisconnectTimeoutMinutes

`func (o *CatalogScaleSettingsModel) GetPeakExtendedDisconnectTimeoutMinutes() int32`

GetPeakExtendedDisconnectTimeoutMinutes returns the PeakExtendedDisconnectTimeoutMinutes field if non-nil, zero value otherwise.

### GetPeakExtendedDisconnectTimeoutMinutesOk

`func (o *CatalogScaleSettingsModel) GetPeakExtendedDisconnectTimeoutMinutesOk() (*int32, bool)`

GetPeakExtendedDisconnectTimeoutMinutesOk returns a tuple with the PeakExtendedDisconnectTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakExtendedDisconnectTimeoutMinutes

`func (o *CatalogScaleSettingsModel) SetPeakExtendedDisconnectTimeoutMinutes(v int32)`

SetPeakExtendedDisconnectTimeoutMinutes sets PeakExtendedDisconnectTimeoutMinutes field to given value.

### HasPeakExtendedDisconnectTimeoutMinutes

`func (o *CatalogScaleSettingsModel) HasPeakExtendedDisconnectTimeoutMinutes() bool`

HasPeakExtendedDisconnectTimeoutMinutes returns a boolean if a field has been set.

### GetOffPeakExtendedDisconnectTimeoutMinutes

`func (o *CatalogScaleSettingsModel) GetOffPeakExtendedDisconnectTimeoutMinutes() int32`

GetOffPeakExtendedDisconnectTimeoutMinutes returns the OffPeakExtendedDisconnectTimeoutMinutes field if non-nil, zero value otherwise.

### GetOffPeakExtendedDisconnectTimeoutMinutesOk

`func (o *CatalogScaleSettingsModel) GetOffPeakExtendedDisconnectTimeoutMinutesOk() (*int32, bool)`

GetOffPeakExtendedDisconnectTimeoutMinutesOk returns a tuple with the OffPeakExtendedDisconnectTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakExtendedDisconnectTimeoutMinutes

`func (o *CatalogScaleSettingsModel) SetOffPeakExtendedDisconnectTimeoutMinutes(v int32)`

SetOffPeakExtendedDisconnectTimeoutMinutes sets OffPeakExtendedDisconnectTimeoutMinutes field to given value.

### HasOffPeakExtendedDisconnectTimeoutMinutes

`func (o *CatalogScaleSettingsModel) HasOffPeakExtendedDisconnectTimeoutMinutes() bool`

HasOffPeakExtendedDisconnectTimeoutMinutes returns a boolean if a field has been set.

### GetPowerOffDelay

`func (o *CatalogScaleSettingsModel) GetPowerOffDelay() int32`

GetPowerOffDelay returns the PowerOffDelay field if non-nil, zero value otherwise.

### GetPowerOffDelayOk

`func (o *CatalogScaleSettingsModel) GetPowerOffDelayOk() (*int32, bool)`

GetPowerOffDelayOk returns a tuple with the PowerOffDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOffDelay

`func (o *CatalogScaleSettingsModel) SetPowerOffDelay(v int32)`

SetPowerOffDelay sets PowerOffDelay field to given value.

### HasPowerOffDelay

`func (o *CatalogScaleSettingsModel) HasPowerOffDelay() bool`

HasPowerOffDelay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


