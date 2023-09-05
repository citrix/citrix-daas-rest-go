# ProjectedMachinesTimeSlotResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **string** | The start time (in 24-hour clock) for the 30-minute period represented that is based on the time zone of the delivery group. | 
**PoolSizeCount** | **int32** | The projected machines as per the pool size configured for the current 30 minute period. | 
**BufferSizeCount** | **int32** | The projected machines as per the buffer percentage configured for the current 30 minute period. | 

## Methods

### NewProjectedMachinesTimeSlotResponseModel

`func NewProjectedMachinesTimeSlotResponseModel(time string, poolSizeCount int32, bufferSizeCount int32, ) *ProjectedMachinesTimeSlotResponseModel`

NewProjectedMachinesTimeSlotResponseModel instantiates a new ProjectedMachinesTimeSlotResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectedMachinesTimeSlotResponseModelWithDefaults

`func NewProjectedMachinesTimeSlotResponseModelWithDefaults() *ProjectedMachinesTimeSlotResponseModel`

NewProjectedMachinesTimeSlotResponseModelWithDefaults instantiates a new ProjectedMachinesTimeSlotResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *ProjectedMachinesTimeSlotResponseModel) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ProjectedMachinesTimeSlotResponseModel) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ProjectedMachinesTimeSlotResponseModel) SetTime(v string)`

SetTime sets Time field to given value.


### GetPoolSizeCount

`func (o *ProjectedMachinesTimeSlotResponseModel) GetPoolSizeCount() int32`

GetPoolSizeCount returns the PoolSizeCount field if non-nil, zero value otherwise.

### GetPoolSizeCountOk

`func (o *ProjectedMachinesTimeSlotResponseModel) GetPoolSizeCountOk() (*int32, bool)`

GetPoolSizeCountOk returns a tuple with the PoolSizeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSizeCount

`func (o *ProjectedMachinesTimeSlotResponseModel) SetPoolSizeCount(v int32)`

SetPoolSizeCount sets PoolSizeCount field to given value.


### GetBufferSizeCount

`func (o *ProjectedMachinesTimeSlotResponseModel) GetBufferSizeCount() int32`

GetBufferSizeCount returns the BufferSizeCount field if non-nil, zero value otherwise.

### GetBufferSizeCountOk

`func (o *ProjectedMachinesTimeSlotResponseModel) GetBufferSizeCountOk() (*int32, bool)`

GetBufferSizeCountOk returns a tuple with the BufferSizeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferSizeCount

`func (o *ProjectedMachinesTimeSlotResponseModel) SetBufferSizeCount(v int32)`

SetBufferSizeCount sets BufferSizeCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


