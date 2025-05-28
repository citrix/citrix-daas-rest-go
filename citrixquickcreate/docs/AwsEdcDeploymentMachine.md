# AwsEdcDeploymentMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | Pointer to **NullableString** | Workspace Id | [optional] 
**Username** | Pointer to **NullableString** | Username associated with workspace | [optional] 
**RootVolumeSize** | Pointer to **NullableInt32** | Size of operating system volume | [optional] 
**UserVolumeSize** | Pointer to **NullableInt32** | Size of user storage volume | [optional] 
**RootVolumeEncrypted** | Pointer to **bool** | Is operating system volume encrypted | [optional] 
**UserVolumeEncrypted** | Pointer to **bool** | Is user storage volume encrypted | [optional] 
**ComputeType** | Pointer to [**NullableAwsEdcWorkspaceCompute**](AwsEdcWorkspaceCompute.md) | Compute size | [optional] 
**RunningMode** | Pointer to [**NullableAwsEdcWorkspaceRunningMode**](AwsEdcWorkspaceRunningMode.md) | Running mode | [optional] 
**WorkspaceStatus** | Pointer to [**NullableAwsEdcWorkspaceState**](AwsEdcWorkspaceState.md) | Status of workspace  Enum values PENDING, AVAILABLE, ADMIN_MAINTANCE, ERROR, STOPPED, STOPPING, SUSPENDED, TERMINATED, TERMINATING, STARTING, REBOOTING, MAINTENANCE, IMPARED, UNHEALTHY, UPDATING | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Error Message | [optional] 

## Methods

### NewAwsEdcDeploymentMachine

`func NewAwsEdcDeploymentMachine() *AwsEdcDeploymentMachine`

NewAwsEdcDeploymentMachine instantiates a new AwsEdcDeploymentMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsEdcDeploymentMachineWithDefaults

`func NewAwsEdcDeploymentMachineWithDefaults() *AwsEdcDeploymentMachine`

NewAwsEdcDeploymentMachineWithDefaults instantiates a new AwsEdcDeploymentMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *AwsEdcDeploymentMachine) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *AwsEdcDeploymentMachine) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *AwsEdcDeploymentMachine) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *AwsEdcDeploymentMachine) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### SetWorkspaceIdNil

`func (o *AwsEdcDeploymentMachine) SetWorkspaceIdNil(b bool)`

 SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil

### UnsetWorkspaceId
`func (o *AwsEdcDeploymentMachine) UnsetWorkspaceId()`

UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil
### GetUsername

`func (o *AwsEdcDeploymentMachine) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AwsEdcDeploymentMachine) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AwsEdcDeploymentMachine) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AwsEdcDeploymentMachine) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *AwsEdcDeploymentMachine) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *AwsEdcDeploymentMachine) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetRootVolumeSize

`func (o *AwsEdcDeploymentMachine) GetRootVolumeSize() int32`

GetRootVolumeSize returns the RootVolumeSize field if non-nil, zero value otherwise.

### GetRootVolumeSizeOk

`func (o *AwsEdcDeploymentMachine) GetRootVolumeSizeOk() (*int32, bool)`

GetRootVolumeSizeOk returns a tuple with the RootVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolumeSize

`func (o *AwsEdcDeploymentMachine) SetRootVolumeSize(v int32)`

SetRootVolumeSize sets RootVolumeSize field to given value.

### HasRootVolumeSize

`func (o *AwsEdcDeploymentMachine) HasRootVolumeSize() bool`

HasRootVolumeSize returns a boolean if a field has been set.

### SetRootVolumeSizeNil

`func (o *AwsEdcDeploymentMachine) SetRootVolumeSizeNil(b bool)`

 SetRootVolumeSizeNil sets the value for RootVolumeSize to be an explicit nil

### UnsetRootVolumeSize
`func (o *AwsEdcDeploymentMachine) UnsetRootVolumeSize()`

UnsetRootVolumeSize ensures that no value is present for RootVolumeSize, not even an explicit nil
### GetUserVolumeSize

`func (o *AwsEdcDeploymentMachine) GetUserVolumeSize() int32`

GetUserVolumeSize returns the UserVolumeSize field if non-nil, zero value otherwise.

### GetUserVolumeSizeOk

`func (o *AwsEdcDeploymentMachine) GetUserVolumeSizeOk() (*int32, bool)`

GetUserVolumeSizeOk returns a tuple with the UserVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVolumeSize

`func (o *AwsEdcDeploymentMachine) SetUserVolumeSize(v int32)`

SetUserVolumeSize sets UserVolumeSize field to given value.

### HasUserVolumeSize

`func (o *AwsEdcDeploymentMachine) HasUserVolumeSize() bool`

HasUserVolumeSize returns a boolean if a field has been set.

### SetUserVolumeSizeNil

`func (o *AwsEdcDeploymentMachine) SetUserVolumeSizeNil(b bool)`

 SetUserVolumeSizeNil sets the value for UserVolumeSize to be an explicit nil

### UnsetUserVolumeSize
`func (o *AwsEdcDeploymentMachine) UnsetUserVolumeSize()`

UnsetUserVolumeSize ensures that no value is present for UserVolumeSize, not even an explicit nil
### GetRootVolumeEncrypted

`func (o *AwsEdcDeploymentMachine) GetRootVolumeEncrypted() bool`

GetRootVolumeEncrypted returns the RootVolumeEncrypted field if non-nil, zero value otherwise.

### GetRootVolumeEncryptedOk

`func (o *AwsEdcDeploymentMachine) GetRootVolumeEncryptedOk() (*bool, bool)`

GetRootVolumeEncryptedOk returns a tuple with the RootVolumeEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolumeEncrypted

`func (o *AwsEdcDeploymentMachine) SetRootVolumeEncrypted(v bool)`

SetRootVolumeEncrypted sets RootVolumeEncrypted field to given value.

### HasRootVolumeEncrypted

`func (o *AwsEdcDeploymentMachine) HasRootVolumeEncrypted() bool`

HasRootVolumeEncrypted returns a boolean if a field has been set.

### GetUserVolumeEncrypted

`func (o *AwsEdcDeploymentMachine) GetUserVolumeEncrypted() bool`

GetUserVolumeEncrypted returns the UserVolumeEncrypted field if non-nil, zero value otherwise.

### GetUserVolumeEncryptedOk

`func (o *AwsEdcDeploymentMachine) GetUserVolumeEncryptedOk() (*bool, bool)`

GetUserVolumeEncryptedOk returns a tuple with the UserVolumeEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVolumeEncrypted

`func (o *AwsEdcDeploymentMachine) SetUserVolumeEncrypted(v bool)`

SetUserVolumeEncrypted sets UserVolumeEncrypted field to given value.

### HasUserVolumeEncrypted

`func (o *AwsEdcDeploymentMachine) HasUserVolumeEncrypted() bool`

HasUserVolumeEncrypted returns a boolean if a field has been set.

### GetComputeType

`func (o *AwsEdcDeploymentMachine) GetComputeType() AwsEdcWorkspaceCompute`

GetComputeType returns the ComputeType field if non-nil, zero value otherwise.

### GetComputeTypeOk

`func (o *AwsEdcDeploymentMachine) GetComputeTypeOk() (*AwsEdcWorkspaceCompute, bool)`

GetComputeTypeOk returns a tuple with the ComputeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeType

`func (o *AwsEdcDeploymentMachine) SetComputeType(v AwsEdcWorkspaceCompute)`

SetComputeType sets ComputeType field to given value.

### HasComputeType

`func (o *AwsEdcDeploymentMachine) HasComputeType() bool`

HasComputeType returns a boolean if a field has been set.

### SetComputeTypeNil

`func (o *AwsEdcDeploymentMachine) SetComputeTypeNil(b bool)`

 SetComputeTypeNil sets the value for ComputeType to be an explicit nil

### UnsetComputeType
`func (o *AwsEdcDeploymentMachine) UnsetComputeType()`

UnsetComputeType ensures that no value is present for ComputeType, not even an explicit nil
### GetRunningMode

`func (o *AwsEdcDeploymentMachine) GetRunningMode() AwsEdcWorkspaceRunningMode`

GetRunningMode returns the RunningMode field if non-nil, zero value otherwise.

### GetRunningModeOk

`func (o *AwsEdcDeploymentMachine) GetRunningModeOk() (*AwsEdcWorkspaceRunningMode, bool)`

GetRunningModeOk returns a tuple with the RunningMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningMode

`func (o *AwsEdcDeploymentMachine) SetRunningMode(v AwsEdcWorkspaceRunningMode)`

SetRunningMode sets RunningMode field to given value.

### HasRunningMode

`func (o *AwsEdcDeploymentMachine) HasRunningMode() bool`

HasRunningMode returns a boolean if a field has been set.

### SetRunningModeNil

`func (o *AwsEdcDeploymentMachine) SetRunningModeNil(b bool)`

 SetRunningModeNil sets the value for RunningMode to be an explicit nil

### UnsetRunningMode
`func (o *AwsEdcDeploymentMachine) UnsetRunningMode()`

UnsetRunningMode ensures that no value is present for RunningMode, not even an explicit nil
### GetWorkspaceStatus

`func (o *AwsEdcDeploymentMachine) GetWorkspaceStatus() AwsEdcWorkspaceState`

GetWorkspaceStatus returns the WorkspaceStatus field if non-nil, zero value otherwise.

### GetWorkspaceStatusOk

`func (o *AwsEdcDeploymentMachine) GetWorkspaceStatusOk() (*AwsEdcWorkspaceState, bool)`

GetWorkspaceStatusOk returns a tuple with the WorkspaceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceStatus

`func (o *AwsEdcDeploymentMachine) SetWorkspaceStatus(v AwsEdcWorkspaceState)`

SetWorkspaceStatus sets WorkspaceStatus field to given value.

### HasWorkspaceStatus

`func (o *AwsEdcDeploymentMachine) HasWorkspaceStatus() bool`

HasWorkspaceStatus returns a boolean if a field has been set.

### SetWorkspaceStatusNil

`func (o *AwsEdcDeploymentMachine) SetWorkspaceStatusNil(b bool)`

 SetWorkspaceStatusNil sets the value for WorkspaceStatus to be an explicit nil

### UnsetWorkspaceStatus
`func (o *AwsEdcDeploymentMachine) UnsetWorkspaceStatus()`

UnsetWorkspaceStatus ensures that no value is present for WorkspaceStatus, not even an explicit nil
### GetErrorMessage

`func (o *AwsEdcDeploymentMachine) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AwsEdcDeploymentMachine) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AwsEdcDeploymentMachine) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AwsEdcDeploymentMachine) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *AwsEdcDeploymentMachine) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *AwsEdcDeploymentMachine) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


