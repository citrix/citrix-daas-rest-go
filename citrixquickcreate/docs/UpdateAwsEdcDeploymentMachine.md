# UpdateAwsEdcDeploymentMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compute** | Pointer to [**NullableCompute**](Compute.md) |  | [optional] 
**RootVolumeSize** | Pointer to **NullableInt32** | Size of operating system volume | [optional] 
**UserVolumeSize** | Pointer to **NullableInt32** | Size of user storage volume | [optional] 

## Methods

### NewUpdateAwsEdcDeploymentMachine

`func NewUpdateAwsEdcDeploymentMachine() *UpdateAwsEdcDeploymentMachine`

NewUpdateAwsEdcDeploymentMachine instantiates a new UpdateAwsEdcDeploymentMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAwsEdcDeploymentMachineWithDefaults

`func NewUpdateAwsEdcDeploymentMachineWithDefaults() *UpdateAwsEdcDeploymentMachine`

NewUpdateAwsEdcDeploymentMachineWithDefaults instantiates a new UpdateAwsEdcDeploymentMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompute

`func (o *UpdateAwsEdcDeploymentMachine) GetCompute() Compute`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *UpdateAwsEdcDeploymentMachine) GetComputeOk() (*Compute, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *UpdateAwsEdcDeploymentMachine) SetCompute(v Compute)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *UpdateAwsEdcDeploymentMachine) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### SetComputeNil

`func (o *UpdateAwsEdcDeploymentMachine) SetComputeNil(b bool)`

 SetComputeNil sets the value for Compute to be an explicit nil

### UnsetCompute
`func (o *UpdateAwsEdcDeploymentMachine) UnsetCompute()`

UnsetCompute ensures that no value is present for Compute, not even an explicit nil
### GetRootVolumeSize

`func (o *UpdateAwsEdcDeploymentMachine) GetRootVolumeSize() int32`

GetRootVolumeSize returns the RootVolumeSize field if non-nil, zero value otherwise.

### GetRootVolumeSizeOk

`func (o *UpdateAwsEdcDeploymentMachine) GetRootVolumeSizeOk() (*int32, bool)`

GetRootVolumeSizeOk returns a tuple with the RootVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolumeSize

`func (o *UpdateAwsEdcDeploymentMachine) SetRootVolumeSize(v int32)`

SetRootVolumeSize sets RootVolumeSize field to given value.

### HasRootVolumeSize

`func (o *UpdateAwsEdcDeploymentMachine) HasRootVolumeSize() bool`

HasRootVolumeSize returns a boolean if a field has been set.

### SetRootVolumeSizeNil

`func (o *UpdateAwsEdcDeploymentMachine) SetRootVolumeSizeNil(b bool)`

 SetRootVolumeSizeNil sets the value for RootVolumeSize to be an explicit nil

### UnsetRootVolumeSize
`func (o *UpdateAwsEdcDeploymentMachine) UnsetRootVolumeSize()`

UnsetRootVolumeSize ensures that no value is present for RootVolumeSize, not even an explicit nil
### GetUserVolumeSize

`func (o *UpdateAwsEdcDeploymentMachine) GetUserVolumeSize() int32`

GetUserVolumeSize returns the UserVolumeSize field if non-nil, zero value otherwise.

### GetUserVolumeSizeOk

`func (o *UpdateAwsEdcDeploymentMachine) GetUserVolumeSizeOk() (*int32, bool)`

GetUserVolumeSizeOk returns a tuple with the UserVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVolumeSize

`func (o *UpdateAwsEdcDeploymentMachine) SetUserVolumeSize(v int32)`

SetUserVolumeSize sets UserVolumeSize field to given value.

### HasUserVolumeSize

`func (o *UpdateAwsEdcDeploymentMachine) HasUserVolumeSize() bool`

HasUserVolumeSize returns a boolean if a field has been set.

### SetUserVolumeSizeNil

`func (o *UpdateAwsEdcDeploymentMachine) SetUserVolumeSizeNil(b bool)`

 SetUserVolumeSizeNil sets the value for UserVolumeSize to be an explicit nil

### UnsetUserVolumeSize
`func (o *UpdateAwsEdcDeploymentMachine) UnsetUserVolumeSize()`

UnsetUserVolumeSize ensures that no value is present for UserVolumeSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


