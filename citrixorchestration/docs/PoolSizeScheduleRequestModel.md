# PoolSizeScheduleRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeRange** | Pointer to **string** | Time range during which the pool size applies. | [optional] 
**PoolSize** | Pointer to **int32** | The number of machines (either as an absolute number or a percentage of the machines in the delivery group, depending on the value of PoolUsingPercentage) that are to be maintained in a running state, whether they are in use or not. | [optional] 

## Methods

### NewPoolSizeScheduleRequestModel

`func NewPoolSizeScheduleRequestModel() *PoolSizeScheduleRequestModel`

NewPoolSizeScheduleRequestModel instantiates a new PoolSizeScheduleRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolSizeScheduleRequestModelWithDefaults

`func NewPoolSizeScheduleRequestModelWithDefaults() *PoolSizeScheduleRequestModel`

NewPoolSizeScheduleRequestModelWithDefaults instantiates a new PoolSizeScheduleRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeRange

`func (o *PoolSizeScheduleRequestModel) GetTimeRange() string`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *PoolSizeScheduleRequestModel) GetTimeRangeOk() (*string, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *PoolSizeScheduleRequestModel) SetTimeRange(v string)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *PoolSizeScheduleRequestModel) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.

### GetPoolSize

`func (o *PoolSizeScheduleRequestModel) GetPoolSize() int32`

GetPoolSize returns the PoolSize field if non-nil, zero value otherwise.

### GetPoolSizeOk

`func (o *PoolSizeScheduleRequestModel) GetPoolSizeOk() (*int32, bool)`

GetPoolSizeOk returns a tuple with the PoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSize

`func (o *PoolSizeScheduleRequestModel) SetPoolSize(v int32)`

SetPoolSize sets PoolSize field to given value.

### HasPoolSize

`func (o *PoolSizeScheduleRequestModel) HasPoolSize() bool`

HasPoolSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


