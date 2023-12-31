# PoolSizeScheduleResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeRange** | Pointer to **NullableString** | Time range during which the pool size applies. | [optional] 
**PoolSize** | Pointer to **int32** | The number of machines (either as an absolute number or a percentage of the machines in the delivery group, depending on the value of PoolUsingPercentage) that are to be maintained in a running state, whether they are in use or not. | [optional] 

## Methods

### NewPoolSizeScheduleResponseModel

`func NewPoolSizeScheduleResponseModel() *PoolSizeScheduleResponseModel`

NewPoolSizeScheduleResponseModel instantiates a new PoolSizeScheduleResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolSizeScheduleResponseModelWithDefaults

`func NewPoolSizeScheduleResponseModelWithDefaults() *PoolSizeScheduleResponseModel`

NewPoolSizeScheduleResponseModelWithDefaults instantiates a new PoolSizeScheduleResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeRange

`func (o *PoolSizeScheduleResponseModel) GetTimeRange() string`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *PoolSizeScheduleResponseModel) GetTimeRangeOk() (*string, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *PoolSizeScheduleResponseModel) SetTimeRange(v string)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *PoolSizeScheduleResponseModel) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.

### SetTimeRangeNil

`func (o *PoolSizeScheduleResponseModel) SetTimeRangeNil(b bool)`

 SetTimeRangeNil sets the value for TimeRange to be an explicit nil

### UnsetTimeRange
`func (o *PoolSizeScheduleResponseModel) UnsetTimeRange()`

UnsetTimeRange ensures that no value is present for TimeRange, not even an explicit nil
### GetPoolSize

`func (o *PoolSizeScheduleResponseModel) GetPoolSize() int32`

GetPoolSize returns the PoolSize field if non-nil, zero value otherwise.

### GetPoolSizeOk

`func (o *PoolSizeScheduleResponseModel) GetPoolSizeOk() (*int32, bool)`

GetPoolSizeOk returns a tuple with the PoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSize

`func (o *PoolSizeScheduleResponseModel) SetPoolSize(v int32)`

SetPoolSize sets PoolSize field to given value.

### HasPoolSize

`func (o *PoolSizeScheduleResponseModel) HasPoolSize() bool`

HasPoolSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


