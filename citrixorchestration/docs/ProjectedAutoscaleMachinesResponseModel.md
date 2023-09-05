# ProjectedAutoscaleMachinesResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DesktopGroupId** | **string** | The Id of the desired desktop group. | 
**ManagedMachineCount** | **int32** | The number of machines in the desktop group that would be managed by Autoscale. This number shall include all the power managed machines in the desktop group, except for machines in maintenance mode or, not tagged for Autoscale if tag restriction is being used. | 
**ProjectedAutoscaleMachines** | [**ProjectedAutoscaleMachinesResponseModelProjectedAutoscaleMachines**](ProjectedAutoscaleMachinesResponseModelProjectedAutoscaleMachines.md) |  | 
**TimeZone** | **string** | The time zone in which this delivery group&#39;s machines reside. | 

## Methods

### NewProjectedAutoscaleMachinesResponseModel

`func NewProjectedAutoscaleMachinesResponseModel(desktopGroupId string, managedMachineCount int32, projectedAutoscaleMachines ProjectedAutoscaleMachinesResponseModelProjectedAutoscaleMachines, timeZone string, ) *ProjectedAutoscaleMachinesResponseModel`

NewProjectedAutoscaleMachinesResponseModel instantiates a new ProjectedAutoscaleMachinesResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectedAutoscaleMachinesResponseModelWithDefaults

`func NewProjectedAutoscaleMachinesResponseModelWithDefaults() *ProjectedAutoscaleMachinesResponseModel`

NewProjectedAutoscaleMachinesResponseModelWithDefaults instantiates a new ProjectedAutoscaleMachinesResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesktopGroupId

`func (o *ProjectedAutoscaleMachinesResponseModel) GetDesktopGroupId() string`

GetDesktopGroupId returns the DesktopGroupId field if non-nil, zero value otherwise.

### GetDesktopGroupIdOk

`func (o *ProjectedAutoscaleMachinesResponseModel) GetDesktopGroupIdOk() (*string, bool)`

GetDesktopGroupIdOk returns a tuple with the DesktopGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopGroupId

`func (o *ProjectedAutoscaleMachinesResponseModel) SetDesktopGroupId(v string)`

SetDesktopGroupId sets DesktopGroupId field to given value.


### GetManagedMachineCount

`func (o *ProjectedAutoscaleMachinesResponseModel) GetManagedMachineCount() int32`

GetManagedMachineCount returns the ManagedMachineCount field if non-nil, zero value otherwise.

### GetManagedMachineCountOk

`func (o *ProjectedAutoscaleMachinesResponseModel) GetManagedMachineCountOk() (*int32, bool)`

GetManagedMachineCountOk returns a tuple with the ManagedMachineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedMachineCount

`func (o *ProjectedAutoscaleMachinesResponseModel) SetManagedMachineCount(v int32)`

SetManagedMachineCount sets ManagedMachineCount field to given value.


### GetProjectedAutoscaleMachines

`func (o *ProjectedAutoscaleMachinesResponseModel) GetProjectedAutoscaleMachines() ProjectedAutoscaleMachinesResponseModelProjectedAutoscaleMachines`

GetProjectedAutoscaleMachines returns the ProjectedAutoscaleMachines field if non-nil, zero value otherwise.

### GetProjectedAutoscaleMachinesOk

`func (o *ProjectedAutoscaleMachinesResponseModel) GetProjectedAutoscaleMachinesOk() (*ProjectedAutoscaleMachinesResponseModelProjectedAutoscaleMachines, bool)`

GetProjectedAutoscaleMachinesOk returns a tuple with the ProjectedAutoscaleMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedAutoscaleMachines

`func (o *ProjectedAutoscaleMachinesResponseModel) SetProjectedAutoscaleMachines(v ProjectedAutoscaleMachinesResponseModelProjectedAutoscaleMachines)`

SetProjectedAutoscaleMachines sets ProjectedAutoscaleMachines field to given value.


### GetTimeZone

`func (o *ProjectedAutoscaleMachinesResponseModel) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *ProjectedAutoscaleMachinesResponseModel) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *ProjectedAutoscaleMachinesResponseModel) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


