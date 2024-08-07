# InitiateAwsEdcDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageId** | **string** | Image Id | 
**ComputeType** | Pointer to [**AwsEdcWorkspaceCompute**](AwsEdcWorkspaceCompute.md) |  | [optional] 
**RootVolumeSize** | Pointer to **int32** | Size of operating system volume for bundle creation | [optional] 
**UserVolumeSize** | Pointer to **int32** | Size of user storage volume for bundle creation | [optional] 
**VolumesEncrypted** | Pointer to **NullableBool** | Is root and user storage volume encrypted | [optional] 
**VolumesEncryptionKey** | Pointer to **NullableString** | Volume encryption key | [optional] 
**RunningMode** | Pointer to [**NullableAwsEdcWorkspaceRunningMode**](AwsEdcWorkspaceRunningMode.md) |  | [optional] 
**ScaleSettings** | Pointer to [**NullableScaleSettings**](ScaleSettings.md) |  | [optional] 
**UserDecoupledWorkspaces** | Pointer to **NullableBool** | Should workspaces be decoupled from the user during creation | [optional] 
**Workspaces** | Pointer to [**[]AddAwsEdcWorkspace**](AddAwsEdcWorkspace.md) | Configuration for individual workspace | [optional] 

## Methods

### NewInitiateAwsEdcDeployment

`func NewInitiateAwsEdcDeployment(imageId string, ) *InitiateAwsEdcDeployment`

NewInitiateAwsEdcDeployment instantiates a new InitiateAwsEdcDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiateAwsEdcDeploymentWithDefaults

`func NewInitiateAwsEdcDeploymentWithDefaults() *InitiateAwsEdcDeployment`

NewInitiateAwsEdcDeploymentWithDefaults instantiates a new InitiateAwsEdcDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageId

`func (o *InitiateAwsEdcDeployment) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *InitiateAwsEdcDeployment) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *InitiateAwsEdcDeployment) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetComputeType

`func (o *InitiateAwsEdcDeployment) GetComputeType() AwsEdcWorkspaceCompute`

GetComputeType returns the ComputeType field if non-nil, zero value otherwise.

### GetComputeTypeOk

`func (o *InitiateAwsEdcDeployment) GetComputeTypeOk() (*AwsEdcWorkspaceCompute, bool)`

GetComputeTypeOk returns a tuple with the ComputeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeType

`func (o *InitiateAwsEdcDeployment) SetComputeType(v AwsEdcWorkspaceCompute)`

SetComputeType sets ComputeType field to given value.

### HasComputeType

`func (o *InitiateAwsEdcDeployment) HasComputeType() bool`

HasComputeType returns a boolean if a field has been set.

### GetRootVolumeSize

`func (o *InitiateAwsEdcDeployment) GetRootVolumeSize() int32`

GetRootVolumeSize returns the RootVolumeSize field if non-nil, zero value otherwise.

### GetRootVolumeSizeOk

`func (o *InitiateAwsEdcDeployment) GetRootVolumeSizeOk() (*int32, bool)`

GetRootVolumeSizeOk returns a tuple with the RootVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolumeSize

`func (o *InitiateAwsEdcDeployment) SetRootVolumeSize(v int32)`

SetRootVolumeSize sets RootVolumeSize field to given value.

### HasRootVolumeSize

`func (o *InitiateAwsEdcDeployment) HasRootVolumeSize() bool`

HasRootVolumeSize returns a boolean if a field has been set.

### GetUserVolumeSize

`func (o *InitiateAwsEdcDeployment) GetUserVolumeSize() int32`

GetUserVolumeSize returns the UserVolumeSize field if non-nil, zero value otherwise.

### GetUserVolumeSizeOk

`func (o *InitiateAwsEdcDeployment) GetUserVolumeSizeOk() (*int32, bool)`

GetUserVolumeSizeOk returns a tuple with the UserVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVolumeSize

`func (o *InitiateAwsEdcDeployment) SetUserVolumeSize(v int32)`

SetUserVolumeSize sets UserVolumeSize field to given value.

### HasUserVolumeSize

`func (o *InitiateAwsEdcDeployment) HasUserVolumeSize() bool`

HasUserVolumeSize returns a boolean if a field has been set.

### GetVolumesEncrypted

`func (o *InitiateAwsEdcDeployment) GetVolumesEncrypted() bool`

GetVolumesEncrypted returns the VolumesEncrypted field if non-nil, zero value otherwise.

### GetVolumesEncryptedOk

`func (o *InitiateAwsEdcDeployment) GetVolumesEncryptedOk() (*bool, bool)`

GetVolumesEncryptedOk returns a tuple with the VolumesEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesEncrypted

`func (o *InitiateAwsEdcDeployment) SetVolumesEncrypted(v bool)`

SetVolumesEncrypted sets VolumesEncrypted field to given value.

### HasVolumesEncrypted

`func (o *InitiateAwsEdcDeployment) HasVolumesEncrypted() bool`

HasVolumesEncrypted returns a boolean if a field has been set.

### SetVolumesEncryptedNil

`func (o *InitiateAwsEdcDeployment) SetVolumesEncryptedNil(b bool)`

 SetVolumesEncryptedNil sets the value for VolumesEncrypted to be an explicit nil

### UnsetVolumesEncrypted
`func (o *InitiateAwsEdcDeployment) UnsetVolumesEncrypted()`

UnsetVolumesEncrypted ensures that no value is present for VolumesEncrypted, not even an explicit nil
### GetVolumesEncryptionKey

`func (o *InitiateAwsEdcDeployment) GetVolumesEncryptionKey() string`

GetVolumesEncryptionKey returns the VolumesEncryptionKey field if non-nil, zero value otherwise.

### GetVolumesEncryptionKeyOk

`func (o *InitiateAwsEdcDeployment) GetVolumesEncryptionKeyOk() (*string, bool)`

GetVolumesEncryptionKeyOk returns a tuple with the VolumesEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesEncryptionKey

`func (o *InitiateAwsEdcDeployment) SetVolumesEncryptionKey(v string)`

SetVolumesEncryptionKey sets VolumesEncryptionKey field to given value.

### HasVolumesEncryptionKey

`func (o *InitiateAwsEdcDeployment) HasVolumesEncryptionKey() bool`

HasVolumesEncryptionKey returns a boolean if a field has been set.

### SetVolumesEncryptionKeyNil

`func (o *InitiateAwsEdcDeployment) SetVolumesEncryptionKeyNil(b bool)`

 SetVolumesEncryptionKeyNil sets the value for VolumesEncryptionKey to be an explicit nil

### UnsetVolumesEncryptionKey
`func (o *InitiateAwsEdcDeployment) UnsetVolumesEncryptionKey()`

UnsetVolumesEncryptionKey ensures that no value is present for VolumesEncryptionKey, not even an explicit nil
### GetRunningMode

`func (o *InitiateAwsEdcDeployment) GetRunningMode() AwsEdcWorkspaceRunningMode`

GetRunningMode returns the RunningMode field if non-nil, zero value otherwise.

### GetRunningModeOk

`func (o *InitiateAwsEdcDeployment) GetRunningModeOk() (*AwsEdcWorkspaceRunningMode, bool)`

GetRunningModeOk returns a tuple with the RunningMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningMode

`func (o *InitiateAwsEdcDeployment) SetRunningMode(v AwsEdcWorkspaceRunningMode)`

SetRunningMode sets RunningMode field to given value.

### HasRunningMode

`func (o *InitiateAwsEdcDeployment) HasRunningMode() bool`

HasRunningMode returns a boolean if a field has been set.

### SetRunningModeNil

`func (o *InitiateAwsEdcDeployment) SetRunningModeNil(b bool)`

 SetRunningModeNil sets the value for RunningMode to be an explicit nil

### UnsetRunningMode
`func (o *InitiateAwsEdcDeployment) UnsetRunningMode()`

UnsetRunningMode ensures that no value is present for RunningMode, not even an explicit nil
### GetScaleSettings

`func (o *InitiateAwsEdcDeployment) GetScaleSettings() ScaleSettings`

GetScaleSettings returns the ScaleSettings field if non-nil, zero value otherwise.

### GetScaleSettingsOk

`func (o *InitiateAwsEdcDeployment) GetScaleSettingsOk() (*ScaleSettings, bool)`

GetScaleSettingsOk returns a tuple with the ScaleSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleSettings

`func (o *InitiateAwsEdcDeployment) SetScaleSettings(v ScaleSettings)`

SetScaleSettings sets ScaleSettings field to given value.

### HasScaleSettings

`func (o *InitiateAwsEdcDeployment) HasScaleSettings() bool`

HasScaleSettings returns a boolean if a field has been set.

### SetScaleSettingsNil

`func (o *InitiateAwsEdcDeployment) SetScaleSettingsNil(b bool)`

 SetScaleSettingsNil sets the value for ScaleSettings to be an explicit nil

### UnsetScaleSettings
`func (o *InitiateAwsEdcDeployment) UnsetScaleSettings()`

UnsetScaleSettings ensures that no value is present for ScaleSettings, not even an explicit nil
### GetUserDecoupledWorkspaces

`func (o *InitiateAwsEdcDeployment) GetUserDecoupledWorkspaces() bool`

GetUserDecoupledWorkspaces returns the UserDecoupledWorkspaces field if non-nil, zero value otherwise.

### GetUserDecoupledWorkspacesOk

`func (o *InitiateAwsEdcDeployment) GetUserDecoupledWorkspacesOk() (*bool, bool)`

GetUserDecoupledWorkspacesOk returns a tuple with the UserDecoupledWorkspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDecoupledWorkspaces

`func (o *InitiateAwsEdcDeployment) SetUserDecoupledWorkspaces(v bool)`

SetUserDecoupledWorkspaces sets UserDecoupledWorkspaces field to given value.

### HasUserDecoupledWorkspaces

`func (o *InitiateAwsEdcDeployment) HasUserDecoupledWorkspaces() bool`

HasUserDecoupledWorkspaces returns a boolean if a field has been set.

### SetUserDecoupledWorkspacesNil

`func (o *InitiateAwsEdcDeployment) SetUserDecoupledWorkspacesNil(b bool)`

 SetUserDecoupledWorkspacesNil sets the value for UserDecoupledWorkspaces to be an explicit nil

### UnsetUserDecoupledWorkspaces
`func (o *InitiateAwsEdcDeployment) UnsetUserDecoupledWorkspaces()`

UnsetUserDecoupledWorkspaces ensures that no value is present for UserDecoupledWorkspaces, not even an explicit nil
### GetWorkspaces

`func (o *InitiateAwsEdcDeployment) GetWorkspaces() []AddAwsEdcWorkspace`

GetWorkspaces returns the Workspaces field if non-nil, zero value otherwise.

### GetWorkspacesOk

`func (o *InitiateAwsEdcDeployment) GetWorkspacesOk() (*[]AddAwsEdcWorkspace, bool)`

GetWorkspacesOk returns a tuple with the Workspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaces

`func (o *InitiateAwsEdcDeployment) SetWorkspaces(v []AddAwsEdcWorkspace)`

SetWorkspaces sets Workspaces field to given value.

### HasWorkspaces

`func (o *InitiateAwsEdcDeployment) HasWorkspaces() bool`

HasWorkspaces returns a boolean if a field has been set.

### SetWorkspacesNil

`func (o *InitiateAwsEdcDeployment) SetWorkspacesNil(b bool)`

 SetWorkspacesNil sets the value for Workspaces to be an explicit nil

### UnsetWorkspaces
`func (o *InitiateAwsEdcDeployment) UnsetWorkspaces()`

UnsetWorkspaces ensures that no value is present for Workspaces, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


