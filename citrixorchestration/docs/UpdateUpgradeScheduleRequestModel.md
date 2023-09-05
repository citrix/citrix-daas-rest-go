# UpdateUpgradeScheduleRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDateTimeUtc** | **time.Time** | UTC time to start the Vda upgrade. Must be a future time. | 
**DurationInHours** | **int32** | Timeout duration in hours. Valid range is 1 to 24. | 

## Methods

### NewUpdateUpgradeScheduleRequestModel

`func NewUpdateUpgradeScheduleRequestModel(startDateTimeUtc time.Time, durationInHours int32, ) *UpdateUpgradeScheduleRequestModel`

NewUpdateUpgradeScheduleRequestModel instantiates a new UpdateUpgradeScheduleRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUpgradeScheduleRequestModelWithDefaults

`func NewUpdateUpgradeScheduleRequestModelWithDefaults() *UpdateUpgradeScheduleRequestModel`

NewUpdateUpgradeScheduleRequestModelWithDefaults instantiates a new UpdateUpgradeScheduleRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDateTimeUtc

`func (o *UpdateUpgradeScheduleRequestModel) GetStartDateTimeUtc() time.Time`

GetStartDateTimeUtc returns the StartDateTimeUtc field if non-nil, zero value otherwise.

### GetStartDateTimeUtcOk

`func (o *UpdateUpgradeScheduleRequestModel) GetStartDateTimeUtcOk() (*time.Time, bool)`

GetStartDateTimeUtcOk returns a tuple with the StartDateTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTimeUtc

`func (o *UpdateUpgradeScheduleRequestModel) SetStartDateTimeUtc(v time.Time)`

SetStartDateTimeUtc sets StartDateTimeUtc field to given value.


### GetDurationInHours

`func (o *UpdateUpgradeScheduleRequestModel) GetDurationInHours() int32`

GetDurationInHours returns the DurationInHours field if non-nil, zero value otherwise.

### GetDurationInHoursOk

`func (o *UpdateUpgradeScheduleRequestModel) GetDurationInHoursOk() (*int32, bool)`

GetDurationInHoursOk returns a tuple with the DurationInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInHours

`func (o *UpdateUpgradeScheduleRequestModel) SetDurationInHours(v int32)`

SetDurationInHours sets DurationInHours field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


