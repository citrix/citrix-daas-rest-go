# AwsEdcDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputeType** | Pointer to [**NullableAwsEdcWorkspaceCompute**](AwsEdcWorkspaceCompute.md) | Performance size for bundle creation | [optional] 
**RootVolumeSize** | Pointer to **NullableInt32** | Size of operating system volume for bundle creation | [optional] 
**UserVolumeSize** | Pointer to **NullableInt32** | Size of user storage volume for bundle creation | [optional] 
**VolumesEncrypted** | Pointer to **NullableBool** | Is root and user storage volume encrypted | [optional] 
**VolumesEncryptionKey** | Pointer to **NullableString** | Volume encryption key | [optional] 
**ImageId** | Pointer to **NullableString** | Image Id | [optional] 
**ImageName** | Pointer to **NullableString** | Image Name | [optional] 
**RunningMode** | Pointer to [**NullableAwsEdcWorkspaceRunningMode**](AwsEdcWorkspaceRunningMode.md) | Running mode to apply to all workspaces | [optional] 
**Workspaces** | Pointer to [**[]AwsEdcDeploymentMachine**](AwsEdcDeploymentMachine.md) | Configuration for individual workspace | [optional] 
**UserDecoupledWorkspaces** | Pointer to **NullableBool** | Image Name | [optional] 
**DirectoryName** | Pointer to **NullableString** | The name of the directory associated with the deployment | [optional] 

## Methods

### NewAwsEdcDeployment

`func NewAwsEdcDeployment() *AwsEdcDeployment`

NewAwsEdcDeployment instantiates a new AwsEdcDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsEdcDeploymentWithDefaults

`func NewAwsEdcDeploymentWithDefaults() *AwsEdcDeployment`

NewAwsEdcDeploymentWithDefaults instantiates a new AwsEdcDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputeType

`func (o *AwsEdcDeployment) GetComputeType() AwsEdcWorkspaceCompute`

GetComputeType returns the ComputeType field if non-nil, zero value otherwise.

### GetComputeTypeOk

`func (o *AwsEdcDeployment) GetComputeTypeOk() (*AwsEdcWorkspaceCompute, bool)`

GetComputeTypeOk returns a tuple with the ComputeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeType

`func (o *AwsEdcDeployment) SetComputeType(v AwsEdcWorkspaceCompute)`

SetComputeType sets ComputeType field to given value.

### HasComputeType

`func (o *AwsEdcDeployment) HasComputeType() bool`

HasComputeType returns a boolean if a field has been set.

### SetComputeTypeNil

`func (o *AwsEdcDeployment) SetComputeTypeNil(b bool)`

 SetComputeTypeNil sets the value for ComputeType to be an explicit nil

### UnsetComputeType
`func (o *AwsEdcDeployment) UnsetComputeType()`

UnsetComputeType ensures that no value is present for ComputeType, not even an explicit nil
### GetRootVolumeSize

`func (o *AwsEdcDeployment) GetRootVolumeSize() int32`

GetRootVolumeSize returns the RootVolumeSize field if non-nil, zero value otherwise.

### GetRootVolumeSizeOk

`func (o *AwsEdcDeployment) GetRootVolumeSizeOk() (*int32, bool)`

GetRootVolumeSizeOk returns a tuple with the RootVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolumeSize

`func (o *AwsEdcDeployment) SetRootVolumeSize(v int32)`

SetRootVolumeSize sets RootVolumeSize field to given value.

### HasRootVolumeSize

`func (o *AwsEdcDeployment) HasRootVolumeSize() bool`

HasRootVolumeSize returns a boolean if a field has been set.

### SetRootVolumeSizeNil

`func (o *AwsEdcDeployment) SetRootVolumeSizeNil(b bool)`

 SetRootVolumeSizeNil sets the value for RootVolumeSize to be an explicit nil

### UnsetRootVolumeSize
`func (o *AwsEdcDeployment) UnsetRootVolumeSize()`

UnsetRootVolumeSize ensures that no value is present for RootVolumeSize, not even an explicit nil
### GetUserVolumeSize

`func (o *AwsEdcDeployment) GetUserVolumeSize() int32`

GetUserVolumeSize returns the UserVolumeSize field if non-nil, zero value otherwise.

### GetUserVolumeSizeOk

`func (o *AwsEdcDeployment) GetUserVolumeSizeOk() (*int32, bool)`

GetUserVolumeSizeOk returns a tuple with the UserVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVolumeSize

`func (o *AwsEdcDeployment) SetUserVolumeSize(v int32)`

SetUserVolumeSize sets UserVolumeSize field to given value.

### HasUserVolumeSize

`func (o *AwsEdcDeployment) HasUserVolumeSize() bool`

HasUserVolumeSize returns a boolean if a field has been set.

### SetUserVolumeSizeNil

`func (o *AwsEdcDeployment) SetUserVolumeSizeNil(b bool)`

 SetUserVolumeSizeNil sets the value for UserVolumeSize to be an explicit nil

### UnsetUserVolumeSize
`func (o *AwsEdcDeployment) UnsetUserVolumeSize()`

UnsetUserVolumeSize ensures that no value is present for UserVolumeSize, not even an explicit nil
### GetVolumesEncrypted

`func (o *AwsEdcDeployment) GetVolumesEncrypted() bool`

GetVolumesEncrypted returns the VolumesEncrypted field if non-nil, zero value otherwise.

### GetVolumesEncryptedOk

`func (o *AwsEdcDeployment) GetVolumesEncryptedOk() (*bool, bool)`

GetVolumesEncryptedOk returns a tuple with the VolumesEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesEncrypted

`func (o *AwsEdcDeployment) SetVolumesEncrypted(v bool)`

SetVolumesEncrypted sets VolumesEncrypted field to given value.

### HasVolumesEncrypted

`func (o *AwsEdcDeployment) HasVolumesEncrypted() bool`

HasVolumesEncrypted returns a boolean if a field has been set.

### SetVolumesEncryptedNil

`func (o *AwsEdcDeployment) SetVolumesEncryptedNil(b bool)`

 SetVolumesEncryptedNil sets the value for VolumesEncrypted to be an explicit nil

### UnsetVolumesEncrypted
`func (o *AwsEdcDeployment) UnsetVolumesEncrypted()`

UnsetVolumesEncrypted ensures that no value is present for VolumesEncrypted, not even an explicit nil
### GetVolumesEncryptionKey

`func (o *AwsEdcDeployment) GetVolumesEncryptionKey() string`

GetVolumesEncryptionKey returns the VolumesEncryptionKey field if non-nil, zero value otherwise.

### GetVolumesEncryptionKeyOk

`func (o *AwsEdcDeployment) GetVolumesEncryptionKeyOk() (*string, bool)`

GetVolumesEncryptionKeyOk returns a tuple with the VolumesEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesEncryptionKey

`func (o *AwsEdcDeployment) SetVolumesEncryptionKey(v string)`

SetVolumesEncryptionKey sets VolumesEncryptionKey field to given value.

### HasVolumesEncryptionKey

`func (o *AwsEdcDeployment) HasVolumesEncryptionKey() bool`

HasVolumesEncryptionKey returns a boolean if a field has been set.

### SetVolumesEncryptionKeyNil

`func (o *AwsEdcDeployment) SetVolumesEncryptionKeyNil(b bool)`

 SetVolumesEncryptionKeyNil sets the value for VolumesEncryptionKey to be an explicit nil

### UnsetVolumesEncryptionKey
`func (o *AwsEdcDeployment) UnsetVolumesEncryptionKey()`

UnsetVolumesEncryptionKey ensures that no value is present for VolumesEncryptionKey, not even an explicit nil
### GetImageId

`func (o *AwsEdcDeployment) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *AwsEdcDeployment) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *AwsEdcDeployment) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *AwsEdcDeployment) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### SetImageIdNil

`func (o *AwsEdcDeployment) SetImageIdNil(b bool)`

 SetImageIdNil sets the value for ImageId to be an explicit nil

### UnsetImageId
`func (o *AwsEdcDeployment) UnsetImageId()`

UnsetImageId ensures that no value is present for ImageId, not even an explicit nil
### GetImageName

`func (o *AwsEdcDeployment) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *AwsEdcDeployment) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *AwsEdcDeployment) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *AwsEdcDeployment) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### SetImageNameNil

`func (o *AwsEdcDeployment) SetImageNameNil(b bool)`

 SetImageNameNil sets the value for ImageName to be an explicit nil

### UnsetImageName
`func (o *AwsEdcDeployment) UnsetImageName()`

UnsetImageName ensures that no value is present for ImageName, not even an explicit nil
### GetRunningMode

`func (o *AwsEdcDeployment) GetRunningMode() AwsEdcWorkspaceRunningMode`

GetRunningMode returns the RunningMode field if non-nil, zero value otherwise.

### GetRunningModeOk

`func (o *AwsEdcDeployment) GetRunningModeOk() (*AwsEdcWorkspaceRunningMode, bool)`

GetRunningModeOk returns a tuple with the RunningMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningMode

`func (o *AwsEdcDeployment) SetRunningMode(v AwsEdcWorkspaceRunningMode)`

SetRunningMode sets RunningMode field to given value.

### HasRunningMode

`func (o *AwsEdcDeployment) HasRunningMode() bool`

HasRunningMode returns a boolean if a field has been set.

### SetRunningModeNil

`func (o *AwsEdcDeployment) SetRunningModeNil(b bool)`

 SetRunningModeNil sets the value for RunningMode to be an explicit nil

### UnsetRunningMode
`func (o *AwsEdcDeployment) UnsetRunningMode()`

UnsetRunningMode ensures that no value is present for RunningMode, not even an explicit nil
### GetWorkspaces

`func (o *AwsEdcDeployment) GetWorkspaces() []AwsEdcDeploymentMachine`

GetWorkspaces returns the Workspaces field if non-nil, zero value otherwise.

### GetWorkspacesOk

`func (o *AwsEdcDeployment) GetWorkspacesOk() (*[]AwsEdcDeploymentMachine, bool)`

GetWorkspacesOk returns a tuple with the Workspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaces

`func (o *AwsEdcDeployment) SetWorkspaces(v []AwsEdcDeploymentMachine)`

SetWorkspaces sets Workspaces field to given value.

### HasWorkspaces

`func (o *AwsEdcDeployment) HasWorkspaces() bool`

HasWorkspaces returns a boolean if a field has been set.

### SetWorkspacesNil

`func (o *AwsEdcDeployment) SetWorkspacesNil(b bool)`

 SetWorkspacesNil sets the value for Workspaces to be an explicit nil

### UnsetWorkspaces
`func (o *AwsEdcDeployment) UnsetWorkspaces()`

UnsetWorkspaces ensures that no value is present for Workspaces, not even an explicit nil
### GetUserDecoupledWorkspaces

`func (o *AwsEdcDeployment) GetUserDecoupledWorkspaces() bool`

GetUserDecoupledWorkspaces returns the UserDecoupledWorkspaces field if non-nil, zero value otherwise.

### GetUserDecoupledWorkspacesOk

`func (o *AwsEdcDeployment) GetUserDecoupledWorkspacesOk() (*bool, bool)`

GetUserDecoupledWorkspacesOk returns a tuple with the UserDecoupledWorkspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDecoupledWorkspaces

`func (o *AwsEdcDeployment) SetUserDecoupledWorkspaces(v bool)`

SetUserDecoupledWorkspaces sets UserDecoupledWorkspaces field to given value.

### HasUserDecoupledWorkspaces

`func (o *AwsEdcDeployment) HasUserDecoupledWorkspaces() bool`

HasUserDecoupledWorkspaces returns a boolean if a field has been set.

### SetUserDecoupledWorkspacesNil

`func (o *AwsEdcDeployment) SetUserDecoupledWorkspacesNil(b bool)`

 SetUserDecoupledWorkspacesNil sets the value for UserDecoupledWorkspaces to be an explicit nil

### UnsetUserDecoupledWorkspaces
`func (o *AwsEdcDeployment) UnsetUserDecoupledWorkspaces()`

UnsetUserDecoupledWorkspaces ensures that no value is present for UserDecoupledWorkspaces, not even an explicit nil
### GetDirectoryName

`func (o *AwsEdcDeployment) GetDirectoryName() string`

GetDirectoryName returns the DirectoryName field if non-nil, zero value otherwise.

### GetDirectoryNameOk

`func (o *AwsEdcDeployment) GetDirectoryNameOk() (*string, bool)`

GetDirectoryNameOk returns a tuple with the DirectoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryName

`func (o *AwsEdcDeployment) SetDirectoryName(v string)`

SetDirectoryName sets DirectoryName field to given value.

### HasDirectoryName

`func (o *AwsEdcDeployment) HasDirectoryName() bool`

HasDirectoryName returns a boolean if a field has been set.

### SetDirectoryNameNil

`func (o *AwsEdcDeployment) SetDirectoryNameNil(b bool)`

 SetDirectoryNameNil sets the value for DirectoryName to be an explicit nil

### UnsetDirectoryName
`func (o *AwsEdcDeployment) UnsetDirectoryName()`

UnsetDirectoryName ensures that no value is present for DirectoryName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


