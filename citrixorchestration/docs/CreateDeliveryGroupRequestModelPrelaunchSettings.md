# CreateDeliveryGroupRequestModelPrelaunchSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates if the pre-launch or lingering settings are enabled. | [optional] [default to true]
**MaxAverageLoadThreshold** | Pointer to **int32** | Specifies the average load threshold across the delivery group before pre-launched or lingering sessions will be terminated. When the threshold hits, sessions which have been pre-launched or which have lingered the longest will be terminated to reduce load. | [optional] [default to 0]
**MaxLoadPerMachineThreshold** | Pointer to **int32** | Specifies the maximum load threshold per machine in the delivery group before pre-launched or lingering sessions will be terminated. When the threshold hits, sessions which have been pre-launched or which have lingered the longest on each loaded machine will be terminated to reduce load. Was: public int? MaxPerDesktopLoadThreshold { get; set; } | [optional] [default to 0]
**MaxTimeBeforeDisconnectMinutes** | Pointer to **int32** | Specifies the time by which a pre-launched or lingering session will be disconnected. | [optional] [default to 0]
**MaxTimeBeforeTerminateMinutes** | Pointer to **int32** | Specifies the time by which a pre-launched or lingering session will be terminated. | [optional] [default to 0]
**IncludedUserFilterEnabled** | Pointer to **bool** | Specifies whether the IncludedUsers filter is enabled or disabled.  When the user filter is enabled, pre-launch or lingering are enabled only to users who appear in the filter (either explicitly or by virtue of group membership). Was: UserFilterEnabled | [optional] [default to false]
**IncludedUsers** | Pointer to **[]string** | Specifies the users and user groups to whom the pre-launch or lingering settings apply. Was: users | [optional] 

## Methods

### NewCreateDeliveryGroupRequestModelPrelaunchSettings

`func NewCreateDeliveryGroupRequestModelPrelaunchSettings() *CreateDeliveryGroupRequestModelPrelaunchSettings`

NewCreateDeliveryGroupRequestModelPrelaunchSettings instantiates a new CreateDeliveryGroupRequestModelPrelaunchSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeliveryGroupRequestModelPrelaunchSettingsWithDefaults

`func NewCreateDeliveryGroupRequestModelPrelaunchSettingsWithDefaults() *CreateDeliveryGroupRequestModelPrelaunchSettings`

NewCreateDeliveryGroupRequestModelPrelaunchSettingsWithDefaults instantiates a new CreateDeliveryGroupRequestModelPrelaunchSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaxAverageLoadThreshold

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetMaxAverageLoadThreshold() int32`

GetMaxAverageLoadThreshold returns the MaxAverageLoadThreshold field if non-nil, zero value otherwise.

### GetMaxAverageLoadThresholdOk

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetMaxAverageLoadThresholdOk() (*int32, bool)`

GetMaxAverageLoadThresholdOk returns a tuple with the MaxAverageLoadThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAverageLoadThreshold

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) SetMaxAverageLoadThreshold(v int32)`

SetMaxAverageLoadThreshold sets MaxAverageLoadThreshold field to given value.

### HasMaxAverageLoadThreshold

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) HasMaxAverageLoadThreshold() bool`

HasMaxAverageLoadThreshold returns a boolean if a field has been set.

### GetMaxLoadPerMachineThreshold

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetMaxLoadPerMachineThreshold() int32`

GetMaxLoadPerMachineThreshold returns the MaxLoadPerMachineThreshold field if non-nil, zero value otherwise.

### GetMaxLoadPerMachineThresholdOk

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetMaxLoadPerMachineThresholdOk() (*int32, bool)`

GetMaxLoadPerMachineThresholdOk returns a tuple with the MaxLoadPerMachineThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLoadPerMachineThreshold

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) SetMaxLoadPerMachineThreshold(v int32)`

SetMaxLoadPerMachineThreshold sets MaxLoadPerMachineThreshold field to given value.

### HasMaxLoadPerMachineThreshold

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) HasMaxLoadPerMachineThreshold() bool`

HasMaxLoadPerMachineThreshold returns a boolean if a field has been set.

### GetMaxTimeBeforeDisconnectMinutes

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetMaxTimeBeforeDisconnectMinutes() int32`

GetMaxTimeBeforeDisconnectMinutes returns the MaxTimeBeforeDisconnectMinutes field if non-nil, zero value otherwise.

### GetMaxTimeBeforeDisconnectMinutesOk

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetMaxTimeBeforeDisconnectMinutesOk() (*int32, bool)`

GetMaxTimeBeforeDisconnectMinutesOk returns a tuple with the MaxTimeBeforeDisconnectMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTimeBeforeDisconnectMinutes

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) SetMaxTimeBeforeDisconnectMinutes(v int32)`

SetMaxTimeBeforeDisconnectMinutes sets MaxTimeBeforeDisconnectMinutes field to given value.

### HasMaxTimeBeforeDisconnectMinutes

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) HasMaxTimeBeforeDisconnectMinutes() bool`

HasMaxTimeBeforeDisconnectMinutes returns a boolean if a field has been set.

### GetMaxTimeBeforeTerminateMinutes

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetMaxTimeBeforeTerminateMinutes() int32`

GetMaxTimeBeforeTerminateMinutes returns the MaxTimeBeforeTerminateMinutes field if non-nil, zero value otherwise.

### GetMaxTimeBeforeTerminateMinutesOk

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetMaxTimeBeforeTerminateMinutesOk() (*int32, bool)`

GetMaxTimeBeforeTerminateMinutesOk returns a tuple with the MaxTimeBeforeTerminateMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTimeBeforeTerminateMinutes

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) SetMaxTimeBeforeTerminateMinutes(v int32)`

SetMaxTimeBeforeTerminateMinutes sets MaxTimeBeforeTerminateMinutes field to given value.

### HasMaxTimeBeforeTerminateMinutes

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) HasMaxTimeBeforeTerminateMinutes() bool`

HasMaxTimeBeforeTerminateMinutes returns a boolean if a field has been set.

### GetIncludedUserFilterEnabled

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetIncludedUserFilterEnabled() bool`

GetIncludedUserFilterEnabled returns the IncludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedUserFilterEnabledOk

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetIncludedUserFilterEnabledOk() (*bool, bool)`

GetIncludedUserFilterEnabledOk returns a tuple with the IncludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserFilterEnabled

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) SetIncludedUserFilterEnabled(v bool)`

SetIncludedUserFilterEnabled sets IncludedUserFilterEnabled field to given value.

### HasIncludedUserFilterEnabled

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) HasIncludedUserFilterEnabled() bool`

HasIncludedUserFilterEnabled returns a boolean if a field has been set.

### GetIncludedUsers

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetIncludedUsers() []string`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) GetIncludedUsersOk() (*[]string, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) SetIncludedUsers(v []string)`

SetIncludedUsers sets IncludedUsers field to given value.

### HasIncludedUsers

`func (o *CreateDeliveryGroupRequestModelPrelaunchSettings) HasIncludedUsers() bool`

HasIncludedUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


