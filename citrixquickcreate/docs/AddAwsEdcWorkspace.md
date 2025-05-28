# AddAwsEdcWorkspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **NullableString** | Username associated with workspace | [optional] 
**RootVolumeSize** | Pointer to **int32** | Size of operating system volume | [optional] 
**UserVolumeSize** | Pointer to **int32** | Size of user storage volume | [optional] 
**RunningMode** | Pointer to [**NullableAwsEdcWorkspaceRunningMode**](AwsEdcWorkspaceRunningMode.md) | Running mode of this workspace | [optional] 

## Methods

### NewAddAwsEdcWorkspace

`func NewAddAwsEdcWorkspace() *AddAwsEdcWorkspace`

NewAddAwsEdcWorkspace instantiates a new AddAwsEdcWorkspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAwsEdcWorkspaceWithDefaults

`func NewAddAwsEdcWorkspaceWithDefaults() *AddAwsEdcWorkspace`

NewAddAwsEdcWorkspaceWithDefaults instantiates a new AddAwsEdcWorkspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *AddAwsEdcWorkspace) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddAwsEdcWorkspace) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddAwsEdcWorkspace) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AddAwsEdcWorkspace) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *AddAwsEdcWorkspace) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *AddAwsEdcWorkspace) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetRootVolumeSize

`func (o *AddAwsEdcWorkspace) GetRootVolumeSize() int32`

GetRootVolumeSize returns the RootVolumeSize field if non-nil, zero value otherwise.

### GetRootVolumeSizeOk

`func (o *AddAwsEdcWorkspace) GetRootVolumeSizeOk() (*int32, bool)`

GetRootVolumeSizeOk returns a tuple with the RootVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolumeSize

`func (o *AddAwsEdcWorkspace) SetRootVolumeSize(v int32)`

SetRootVolumeSize sets RootVolumeSize field to given value.

### HasRootVolumeSize

`func (o *AddAwsEdcWorkspace) HasRootVolumeSize() bool`

HasRootVolumeSize returns a boolean if a field has been set.

### GetUserVolumeSize

`func (o *AddAwsEdcWorkspace) GetUserVolumeSize() int32`

GetUserVolumeSize returns the UserVolumeSize field if non-nil, zero value otherwise.

### GetUserVolumeSizeOk

`func (o *AddAwsEdcWorkspace) GetUserVolumeSizeOk() (*int32, bool)`

GetUserVolumeSizeOk returns a tuple with the UserVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVolumeSize

`func (o *AddAwsEdcWorkspace) SetUserVolumeSize(v int32)`

SetUserVolumeSize sets UserVolumeSize field to given value.

### HasUserVolumeSize

`func (o *AddAwsEdcWorkspace) HasUserVolumeSize() bool`

HasUserVolumeSize returns a boolean if a field has been set.

### GetRunningMode

`func (o *AddAwsEdcWorkspace) GetRunningMode() AwsEdcWorkspaceRunningMode`

GetRunningMode returns the RunningMode field if non-nil, zero value otherwise.

### GetRunningModeOk

`func (o *AddAwsEdcWorkspace) GetRunningModeOk() (*AwsEdcWorkspaceRunningMode, bool)`

GetRunningModeOk returns a tuple with the RunningMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningMode

`func (o *AddAwsEdcWorkspace) SetRunningMode(v AwsEdcWorkspaceRunningMode)`

SetRunningMode sets RunningMode field to given value.

### HasRunningMode

`func (o *AddAwsEdcWorkspace) HasRunningMode() bool`

HasRunningMode returns a boolean if a field has been set.

### SetRunningModeNil

`func (o *AddAwsEdcWorkspace) SetRunningModeNil(b bool)`

 SetRunningModeNil sets the value for RunningMode to be an explicit nil

### UnsetRunningMode
`func (o *AddAwsEdcWorkspace) UnsetRunningMode()`

UnsetRunningMode ensures that no value is present for RunningMode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


