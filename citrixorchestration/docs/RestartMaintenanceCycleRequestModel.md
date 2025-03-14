# RestartMaintenanceCycleRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduledStartTimeInUTC** | **time.Time** | Maintenance Cycle start time in UTC. | 
**MaxDurationInMinutes** | **int32** | The maximum duration allowed for Maintenance Cycle  | 

## Methods

### NewRestartMaintenanceCycleRequestModel

`func NewRestartMaintenanceCycleRequestModel(scheduledStartTimeInUTC time.Time, maxDurationInMinutes int32, ) *RestartMaintenanceCycleRequestModel`

NewRestartMaintenanceCycleRequestModel instantiates a new RestartMaintenanceCycleRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestartMaintenanceCycleRequestModelWithDefaults

`func NewRestartMaintenanceCycleRequestModelWithDefaults() *RestartMaintenanceCycleRequestModel`

NewRestartMaintenanceCycleRequestModelWithDefaults instantiates a new RestartMaintenanceCycleRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduledStartTimeInUTC

`func (o *RestartMaintenanceCycleRequestModel) GetScheduledStartTimeInUTC() time.Time`

GetScheduledStartTimeInUTC returns the ScheduledStartTimeInUTC field if non-nil, zero value otherwise.

### GetScheduledStartTimeInUTCOk

`func (o *RestartMaintenanceCycleRequestModel) GetScheduledStartTimeInUTCOk() (*time.Time, bool)`

GetScheduledStartTimeInUTCOk returns a tuple with the ScheduledStartTimeInUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTimeInUTC

`func (o *RestartMaintenanceCycleRequestModel) SetScheduledStartTimeInUTC(v time.Time)`

SetScheduledStartTimeInUTC sets ScheduledStartTimeInUTC field to given value.


### GetMaxDurationInMinutes

`func (o *RestartMaintenanceCycleRequestModel) GetMaxDurationInMinutes() int32`

GetMaxDurationInMinutes returns the MaxDurationInMinutes field if non-nil, zero value otherwise.

### GetMaxDurationInMinutesOk

`func (o *RestartMaintenanceCycleRequestModel) GetMaxDurationInMinutesOk() (*int32, bool)`

GetMaxDurationInMinutesOk returns a tuple with the MaxDurationInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDurationInMinutes

`func (o *RestartMaintenanceCycleRequestModel) SetMaxDurationInMinutes(v int32)`

SetMaxDurationInMinutes sets MaxDurationInMinutes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


