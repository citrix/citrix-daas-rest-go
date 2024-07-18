# UpdateAwsEdcDeploymentProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunningMode** | Pointer to [**NullableAwsEdcWorkspaceRunningMode**](AwsEdcWorkspaceRunningMode.md) |  | [optional] 
**ScaleSettings** | Pointer to [**NullableScaleSettings**](ScaleSettings.md) |  | [optional] 

## Methods

### NewUpdateAwsEdcDeploymentProperties

`func NewUpdateAwsEdcDeploymentProperties() *UpdateAwsEdcDeploymentProperties`

NewUpdateAwsEdcDeploymentProperties instantiates a new UpdateAwsEdcDeploymentProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAwsEdcDeploymentPropertiesWithDefaults

`func NewUpdateAwsEdcDeploymentPropertiesWithDefaults() *UpdateAwsEdcDeploymentProperties`

NewUpdateAwsEdcDeploymentPropertiesWithDefaults instantiates a new UpdateAwsEdcDeploymentProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunningMode

`func (o *UpdateAwsEdcDeploymentProperties) GetRunningMode() AwsEdcWorkspaceRunningMode`

GetRunningMode returns the RunningMode field if non-nil, zero value otherwise.

### GetRunningModeOk

`func (o *UpdateAwsEdcDeploymentProperties) GetRunningModeOk() (*AwsEdcWorkspaceRunningMode, bool)`

GetRunningModeOk returns a tuple with the RunningMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningMode

`func (o *UpdateAwsEdcDeploymentProperties) SetRunningMode(v AwsEdcWorkspaceRunningMode)`

SetRunningMode sets RunningMode field to given value.

### HasRunningMode

`func (o *UpdateAwsEdcDeploymentProperties) HasRunningMode() bool`

HasRunningMode returns a boolean if a field has been set.

### SetRunningModeNil

`func (o *UpdateAwsEdcDeploymentProperties) SetRunningModeNil(b bool)`

 SetRunningModeNil sets the value for RunningMode to be an explicit nil

### UnsetRunningMode
`func (o *UpdateAwsEdcDeploymentProperties) UnsetRunningMode()`

UnsetRunningMode ensures that no value is present for RunningMode, not even an explicit nil
### GetScaleSettings

`func (o *UpdateAwsEdcDeploymentProperties) GetScaleSettings() ScaleSettings`

GetScaleSettings returns the ScaleSettings field if non-nil, zero value otherwise.

### GetScaleSettingsOk

`func (o *UpdateAwsEdcDeploymentProperties) GetScaleSettingsOk() (*ScaleSettings, bool)`

GetScaleSettingsOk returns a tuple with the ScaleSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleSettings

`func (o *UpdateAwsEdcDeploymentProperties) SetScaleSettings(v ScaleSettings)`

SetScaleSettings sets ScaleSettings field to given value.

### HasScaleSettings

`func (o *UpdateAwsEdcDeploymentProperties) HasScaleSettings() bool`

HasScaleSettings returns a boolean if a field has been set.

### SetScaleSettingsNil

`func (o *UpdateAwsEdcDeploymentProperties) SetScaleSettingsNil(b bool)`

 SetScaleSettingsNil sets the value for ScaleSettings to be an explicit nil

### UnsetScaleSettings
`func (o *UpdateAwsEdcDeploymentProperties) UnsetScaleSettings()`

UnsetScaleSettings ensures that no value is present for ScaleSettings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


