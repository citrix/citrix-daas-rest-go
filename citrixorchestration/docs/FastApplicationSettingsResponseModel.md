# FastApplicationSettingsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Indicates if the pre-launch or lingering settings are enabled. | 
**MaxAverageLoadThreshold** | **int32** | Indicates the average load threshold across the delivery group before pre-launched or lingering sessions will be terminated. When the threshold hits, sessions which have been pre-launched or which have lingered the longest will be terminated to reduce load. | 
**MaxLoadPerMachineThreshold** | **int32** | Indicates the maximum load threshold per machine in the delivery group before pre-launched or lingering sessions will be terminated. When the threshold hits, sessions which have been pre-launched or which have lingered the longest on each loaded machine will be terminated to reduce load. | 
**MaxTimeBeforeDisconnectMinutes** | **int32** | Specifies the time by which a pre-launched or lingering session will be disconnected. | 
**MaxTimeBeforeTerminateMinutes** | **int32** | Specifies the time by which a pre-launched or lingering session will be terminated. | 
**IncludedUserFilterEnabled** | **bool** | Indicates whether the IncludedUsers filter is enabled or disabled.  When the user filter is enabled, pre-launch or lingering are enabled only to users who appear in the filter (either explicitly or by virtue of group membership). | 
**IncludedUsers** | Pointer to [**[]IdentityUserResponseModel**](IdentityUserResponseModel.md) | Indicates the users and user groups to whom the pre-launch or lingering settings apply. | [optional] 

## Methods

### NewFastApplicationSettingsResponseModel

`func NewFastApplicationSettingsResponseModel(enabled bool, maxAverageLoadThreshold int32, maxLoadPerMachineThreshold int32, maxTimeBeforeDisconnectMinutes int32, maxTimeBeforeTerminateMinutes int32, includedUserFilterEnabled bool, ) *FastApplicationSettingsResponseModel`

NewFastApplicationSettingsResponseModel instantiates a new FastApplicationSettingsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFastApplicationSettingsResponseModelWithDefaults

`func NewFastApplicationSettingsResponseModelWithDefaults() *FastApplicationSettingsResponseModel`

NewFastApplicationSettingsResponseModelWithDefaults instantiates a new FastApplicationSettingsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *FastApplicationSettingsResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FastApplicationSettingsResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FastApplicationSettingsResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMaxAverageLoadThreshold

`func (o *FastApplicationSettingsResponseModel) GetMaxAverageLoadThreshold() int32`

GetMaxAverageLoadThreshold returns the MaxAverageLoadThreshold field if non-nil, zero value otherwise.

### GetMaxAverageLoadThresholdOk

`func (o *FastApplicationSettingsResponseModel) GetMaxAverageLoadThresholdOk() (*int32, bool)`

GetMaxAverageLoadThresholdOk returns a tuple with the MaxAverageLoadThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAverageLoadThreshold

`func (o *FastApplicationSettingsResponseModel) SetMaxAverageLoadThreshold(v int32)`

SetMaxAverageLoadThreshold sets MaxAverageLoadThreshold field to given value.


### GetMaxLoadPerMachineThreshold

`func (o *FastApplicationSettingsResponseModel) GetMaxLoadPerMachineThreshold() int32`

GetMaxLoadPerMachineThreshold returns the MaxLoadPerMachineThreshold field if non-nil, zero value otherwise.

### GetMaxLoadPerMachineThresholdOk

`func (o *FastApplicationSettingsResponseModel) GetMaxLoadPerMachineThresholdOk() (*int32, bool)`

GetMaxLoadPerMachineThresholdOk returns a tuple with the MaxLoadPerMachineThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLoadPerMachineThreshold

`func (o *FastApplicationSettingsResponseModel) SetMaxLoadPerMachineThreshold(v int32)`

SetMaxLoadPerMachineThreshold sets MaxLoadPerMachineThreshold field to given value.


### GetMaxTimeBeforeDisconnectMinutes

`func (o *FastApplicationSettingsResponseModel) GetMaxTimeBeforeDisconnectMinutes() int32`

GetMaxTimeBeforeDisconnectMinutes returns the MaxTimeBeforeDisconnectMinutes field if non-nil, zero value otherwise.

### GetMaxTimeBeforeDisconnectMinutesOk

`func (o *FastApplicationSettingsResponseModel) GetMaxTimeBeforeDisconnectMinutesOk() (*int32, bool)`

GetMaxTimeBeforeDisconnectMinutesOk returns a tuple with the MaxTimeBeforeDisconnectMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTimeBeforeDisconnectMinutes

`func (o *FastApplicationSettingsResponseModel) SetMaxTimeBeforeDisconnectMinutes(v int32)`

SetMaxTimeBeforeDisconnectMinutes sets MaxTimeBeforeDisconnectMinutes field to given value.


### GetMaxTimeBeforeTerminateMinutes

`func (o *FastApplicationSettingsResponseModel) GetMaxTimeBeforeTerminateMinutes() int32`

GetMaxTimeBeforeTerminateMinutes returns the MaxTimeBeforeTerminateMinutes field if non-nil, zero value otherwise.

### GetMaxTimeBeforeTerminateMinutesOk

`func (o *FastApplicationSettingsResponseModel) GetMaxTimeBeforeTerminateMinutesOk() (*int32, bool)`

GetMaxTimeBeforeTerminateMinutesOk returns a tuple with the MaxTimeBeforeTerminateMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTimeBeforeTerminateMinutes

`func (o *FastApplicationSettingsResponseModel) SetMaxTimeBeforeTerminateMinutes(v int32)`

SetMaxTimeBeforeTerminateMinutes sets MaxTimeBeforeTerminateMinutes field to given value.


### GetIncludedUserFilterEnabled

`func (o *FastApplicationSettingsResponseModel) GetIncludedUserFilterEnabled() bool`

GetIncludedUserFilterEnabled returns the IncludedUserFilterEnabled field if non-nil, zero value otherwise.

### GetIncludedUserFilterEnabledOk

`func (o *FastApplicationSettingsResponseModel) GetIncludedUserFilterEnabledOk() (*bool, bool)`

GetIncludedUserFilterEnabledOk returns a tuple with the IncludedUserFilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserFilterEnabled

`func (o *FastApplicationSettingsResponseModel) SetIncludedUserFilterEnabled(v bool)`

SetIncludedUserFilterEnabled sets IncludedUserFilterEnabled field to given value.


### GetIncludedUsers

`func (o *FastApplicationSettingsResponseModel) GetIncludedUsers() []IdentityUserResponseModel`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *FastApplicationSettingsResponseModel) GetIncludedUsersOk() (*[]IdentityUserResponseModel, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *FastApplicationSettingsResponseModel) SetIncludedUsers(v []IdentityUserResponseModel)`

SetIncludedUsers sets IncludedUsers field to given value.

### HasIncludedUsers

`func (o *FastApplicationSettingsResponseModel) HasIncludedUsers() bool`

HasIncludedUsers returns a boolean if a field has been set.

### SetIncludedUsersNil

`func (o *FastApplicationSettingsResponseModel) SetIncludedUsersNil(b bool)`

 SetIncludedUsersNil sets the value for IncludedUsers to be an explicit nil

### UnsetIncludedUsers
`func (o *FastApplicationSettingsResponseModel) UnsetIncludedUsers()`

UnsetIncludedUsers ensures that no value is present for IncludedUsers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


