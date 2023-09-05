# HypervisorGpuTypeResourceResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Indicates whether the GPU type is enabled on the hypervisor. | 
**FrameBufferSizeMB** | **int32** | Frame Buffer Size in MB | 
**HasDedicatedResource** | **bool** | Indicates whether the GPU type has resources dedicated to it. | 

## Methods

### NewHypervisorGpuTypeResourceResponseModelAllOf

`func NewHypervisorGpuTypeResourceResponseModelAllOf(enabled bool, frameBufferSizeMB int32, hasDedicatedResource bool, ) *HypervisorGpuTypeResourceResponseModelAllOf`

NewHypervisorGpuTypeResourceResponseModelAllOf instantiates a new HypervisorGpuTypeResourceResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorGpuTypeResourceResponseModelAllOfWithDefaults

`func NewHypervisorGpuTypeResourceResponseModelAllOfWithDefaults() *HypervisorGpuTypeResourceResponseModelAllOf`

NewHypervisorGpuTypeResourceResponseModelAllOfWithDefaults instantiates a new HypervisorGpuTypeResourceResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *HypervisorGpuTypeResourceResponseModelAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HypervisorGpuTypeResourceResponseModelAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HypervisorGpuTypeResourceResponseModelAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFrameBufferSizeMB

`func (o *HypervisorGpuTypeResourceResponseModelAllOf) GetFrameBufferSizeMB() int32`

GetFrameBufferSizeMB returns the FrameBufferSizeMB field if non-nil, zero value otherwise.

### GetFrameBufferSizeMBOk

`func (o *HypervisorGpuTypeResourceResponseModelAllOf) GetFrameBufferSizeMBOk() (*int32, bool)`

GetFrameBufferSizeMBOk returns a tuple with the FrameBufferSizeMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameBufferSizeMB

`func (o *HypervisorGpuTypeResourceResponseModelAllOf) SetFrameBufferSizeMB(v int32)`

SetFrameBufferSizeMB sets FrameBufferSizeMB field to given value.


### GetHasDedicatedResource

`func (o *HypervisorGpuTypeResourceResponseModelAllOf) GetHasDedicatedResource() bool`

GetHasDedicatedResource returns the HasDedicatedResource field if non-nil, zero value otherwise.

### GetHasDedicatedResourceOk

`func (o *HypervisorGpuTypeResourceResponseModelAllOf) GetHasDedicatedResourceOk() (*bool, bool)`

GetHasDedicatedResourceOk returns a tuple with the HasDedicatedResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDedicatedResource

`func (o *HypervisorGpuTypeResourceResponseModelAllOf) SetHasDedicatedResource(v bool)`

SetHasDedicatedResource sets HasDedicatedResource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


