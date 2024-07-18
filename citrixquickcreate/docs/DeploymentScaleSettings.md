# DeploymentScaleSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoScaleEnabled** | Pointer to **NullableBool** | Indicates whether the scale settings are enabled for the deployment | [optional] 
**OffPeakDisconnectTimeoutMinutes** | Pointer to **NullableInt32** | The number of minutes before the configured action should be  performed after a user session disconnects outside peak hours. | [optional] 
**OffPeakLogOffTimeoutMinutes** | Pointer to **NullableInt32** | The number of minutes before the configured action should be  performed after a user session ends outside peak hours. | [optional] 
**SessionIdleTimeoutMinutes** | Pointer to **NullableInt32** | Specifies the time in minutes after which an idle session  belonging to the delivery group is disconnected during off-peak time. | [optional] 
**OffPeakBufferSizePercent** | Pointer to **NullableInt32** | The percentage of machines in the delivery group that should be  kept available in an idle state outside peak hours. | [optional] 

## Methods

### NewDeploymentScaleSettings

`func NewDeploymentScaleSettings() *DeploymentScaleSettings`

NewDeploymentScaleSettings instantiates a new DeploymentScaleSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentScaleSettingsWithDefaults

`func NewDeploymentScaleSettingsWithDefaults() *DeploymentScaleSettings`

NewDeploymentScaleSettingsWithDefaults instantiates a new DeploymentScaleSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoScaleEnabled

`func (o *DeploymentScaleSettings) GetAutoScaleEnabled() bool`

GetAutoScaleEnabled returns the AutoScaleEnabled field if non-nil, zero value otherwise.

### GetAutoScaleEnabledOk

`func (o *DeploymentScaleSettings) GetAutoScaleEnabledOk() (*bool, bool)`

GetAutoScaleEnabledOk returns a tuple with the AutoScaleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaleEnabled

`func (o *DeploymentScaleSettings) SetAutoScaleEnabled(v bool)`

SetAutoScaleEnabled sets AutoScaleEnabled field to given value.

### HasAutoScaleEnabled

`func (o *DeploymentScaleSettings) HasAutoScaleEnabled() bool`

HasAutoScaleEnabled returns a boolean if a field has been set.

### SetAutoScaleEnabledNil

`func (o *DeploymentScaleSettings) SetAutoScaleEnabledNil(b bool)`

 SetAutoScaleEnabledNil sets the value for AutoScaleEnabled to be an explicit nil

### UnsetAutoScaleEnabled
`func (o *DeploymentScaleSettings) UnsetAutoScaleEnabled()`

UnsetAutoScaleEnabled ensures that no value is present for AutoScaleEnabled, not even an explicit nil
### GetOffPeakDisconnectTimeoutMinutes

`func (o *DeploymentScaleSettings) GetOffPeakDisconnectTimeoutMinutes() int32`

GetOffPeakDisconnectTimeoutMinutes returns the OffPeakDisconnectTimeoutMinutes field if non-nil, zero value otherwise.

### GetOffPeakDisconnectTimeoutMinutesOk

`func (o *DeploymentScaleSettings) GetOffPeakDisconnectTimeoutMinutesOk() (*int32, bool)`

GetOffPeakDisconnectTimeoutMinutesOk returns a tuple with the OffPeakDisconnectTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakDisconnectTimeoutMinutes

`func (o *DeploymentScaleSettings) SetOffPeakDisconnectTimeoutMinutes(v int32)`

SetOffPeakDisconnectTimeoutMinutes sets OffPeakDisconnectTimeoutMinutes field to given value.

### HasOffPeakDisconnectTimeoutMinutes

`func (o *DeploymentScaleSettings) HasOffPeakDisconnectTimeoutMinutes() bool`

HasOffPeakDisconnectTimeoutMinutes returns a boolean if a field has been set.

### SetOffPeakDisconnectTimeoutMinutesNil

`func (o *DeploymentScaleSettings) SetOffPeakDisconnectTimeoutMinutesNil(b bool)`

 SetOffPeakDisconnectTimeoutMinutesNil sets the value for OffPeakDisconnectTimeoutMinutes to be an explicit nil

### UnsetOffPeakDisconnectTimeoutMinutes
`func (o *DeploymentScaleSettings) UnsetOffPeakDisconnectTimeoutMinutes()`

UnsetOffPeakDisconnectTimeoutMinutes ensures that no value is present for OffPeakDisconnectTimeoutMinutes, not even an explicit nil
### GetOffPeakLogOffTimeoutMinutes

`func (o *DeploymentScaleSettings) GetOffPeakLogOffTimeoutMinutes() int32`

GetOffPeakLogOffTimeoutMinutes returns the OffPeakLogOffTimeoutMinutes field if non-nil, zero value otherwise.

### GetOffPeakLogOffTimeoutMinutesOk

`func (o *DeploymentScaleSettings) GetOffPeakLogOffTimeoutMinutesOk() (*int32, bool)`

GetOffPeakLogOffTimeoutMinutesOk returns a tuple with the OffPeakLogOffTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakLogOffTimeoutMinutes

`func (o *DeploymentScaleSettings) SetOffPeakLogOffTimeoutMinutes(v int32)`

SetOffPeakLogOffTimeoutMinutes sets OffPeakLogOffTimeoutMinutes field to given value.

### HasOffPeakLogOffTimeoutMinutes

`func (o *DeploymentScaleSettings) HasOffPeakLogOffTimeoutMinutes() bool`

HasOffPeakLogOffTimeoutMinutes returns a boolean if a field has been set.

### SetOffPeakLogOffTimeoutMinutesNil

`func (o *DeploymentScaleSettings) SetOffPeakLogOffTimeoutMinutesNil(b bool)`

 SetOffPeakLogOffTimeoutMinutesNil sets the value for OffPeakLogOffTimeoutMinutes to be an explicit nil

### UnsetOffPeakLogOffTimeoutMinutes
`func (o *DeploymentScaleSettings) UnsetOffPeakLogOffTimeoutMinutes()`

UnsetOffPeakLogOffTimeoutMinutes ensures that no value is present for OffPeakLogOffTimeoutMinutes, not even an explicit nil
### GetSessionIdleTimeoutMinutes

`func (o *DeploymentScaleSettings) GetSessionIdleTimeoutMinutes() int32`

GetSessionIdleTimeoutMinutes returns the SessionIdleTimeoutMinutes field if non-nil, zero value otherwise.

### GetSessionIdleTimeoutMinutesOk

`func (o *DeploymentScaleSettings) GetSessionIdleTimeoutMinutesOk() (*int32, bool)`

GetSessionIdleTimeoutMinutesOk returns a tuple with the SessionIdleTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIdleTimeoutMinutes

`func (o *DeploymentScaleSettings) SetSessionIdleTimeoutMinutes(v int32)`

SetSessionIdleTimeoutMinutes sets SessionIdleTimeoutMinutes field to given value.

### HasSessionIdleTimeoutMinutes

`func (o *DeploymentScaleSettings) HasSessionIdleTimeoutMinutes() bool`

HasSessionIdleTimeoutMinutes returns a boolean if a field has been set.

### SetSessionIdleTimeoutMinutesNil

`func (o *DeploymentScaleSettings) SetSessionIdleTimeoutMinutesNil(b bool)`

 SetSessionIdleTimeoutMinutesNil sets the value for SessionIdleTimeoutMinutes to be an explicit nil

### UnsetSessionIdleTimeoutMinutes
`func (o *DeploymentScaleSettings) UnsetSessionIdleTimeoutMinutes()`

UnsetSessionIdleTimeoutMinutes ensures that no value is present for SessionIdleTimeoutMinutes, not even an explicit nil
### GetOffPeakBufferSizePercent

`func (o *DeploymentScaleSettings) GetOffPeakBufferSizePercent() int32`

GetOffPeakBufferSizePercent returns the OffPeakBufferSizePercent field if non-nil, zero value otherwise.

### GetOffPeakBufferSizePercentOk

`func (o *DeploymentScaleSettings) GetOffPeakBufferSizePercentOk() (*int32, bool)`

GetOffPeakBufferSizePercentOk returns a tuple with the OffPeakBufferSizePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPeakBufferSizePercent

`func (o *DeploymentScaleSettings) SetOffPeakBufferSizePercent(v int32)`

SetOffPeakBufferSizePercent sets OffPeakBufferSizePercent field to given value.

### HasOffPeakBufferSizePercent

`func (o *DeploymentScaleSettings) HasOffPeakBufferSizePercent() bool`

HasOffPeakBufferSizePercent returns a boolean if a field has been set.

### SetOffPeakBufferSizePercentNil

`func (o *DeploymentScaleSettings) SetOffPeakBufferSizePercentNil(b bool)`

 SetOffPeakBufferSizePercentNil sets the value for OffPeakBufferSizePercent to be an explicit nil

### UnsetOffPeakBufferSizePercent
`func (o *DeploymentScaleSettings) UnsetOffPeakBufferSizePercent()`

UnsetOffPeakBufferSizePercent ensures that no value is present for OffPeakBufferSizePercent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


