# CreateUpgradeScheduleRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDateTimeUtc** | Pointer to **time.Time** | UTC time to start the Vda upgrade. Must be a future time. If set to null, the upgrade should be started at once. | [optional] 
**DurationInHours** | **int32** | Timeout duration in hours. Valid range is 1 to 24. | 
**VDAComponentsAndFeaturesRequestModel** | Pointer to [**CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel**](CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel.md) |  | [optional] 

## Methods

### NewCreateUpgradeScheduleRequestModel

`func NewCreateUpgradeScheduleRequestModel(durationInHours int32, ) *CreateUpgradeScheduleRequestModel`

NewCreateUpgradeScheduleRequestModel instantiates a new CreateUpgradeScheduleRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpgradeScheduleRequestModelWithDefaults

`func NewCreateUpgradeScheduleRequestModelWithDefaults() *CreateUpgradeScheduleRequestModel`

NewCreateUpgradeScheduleRequestModelWithDefaults instantiates a new CreateUpgradeScheduleRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDateTimeUtc

`func (o *CreateUpgradeScheduleRequestModel) GetStartDateTimeUtc() time.Time`

GetStartDateTimeUtc returns the StartDateTimeUtc field if non-nil, zero value otherwise.

### GetStartDateTimeUtcOk

`func (o *CreateUpgradeScheduleRequestModel) GetStartDateTimeUtcOk() (*time.Time, bool)`

GetStartDateTimeUtcOk returns a tuple with the StartDateTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTimeUtc

`func (o *CreateUpgradeScheduleRequestModel) SetStartDateTimeUtc(v time.Time)`

SetStartDateTimeUtc sets StartDateTimeUtc field to given value.

### HasStartDateTimeUtc

`func (o *CreateUpgradeScheduleRequestModel) HasStartDateTimeUtc() bool`

HasStartDateTimeUtc returns a boolean if a field has been set.

### GetDurationInHours

`func (o *CreateUpgradeScheduleRequestModel) GetDurationInHours() int32`

GetDurationInHours returns the DurationInHours field if non-nil, zero value otherwise.

### GetDurationInHoursOk

`func (o *CreateUpgradeScheduleRequestModel) GetDurationInHoursOk() (*int32, bool)`

GetDurationInHoursOk returns a tuple with the DurationInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInHours

`func (o *CreateUpgradeScheduleRequestModel) SetDurationInHours(v int32)`

SetDurationInHours sets DurationInHours field to given value.


### GetVDAComponentsAndFeaturesRequestModel

`func (o *CreateUpgradeScheduleRequestModel) GetVDAComponentsAndFeaturesRequestModel() CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel`

GetVDAComponentsAndFeaturesRequestModel returns the VDAComponentsAndFeaturesRequestModel field if non-nil, zero value otherwise.

### GetVDAComponentsAndFeaturesRequestModelOk

`func (o *CreateUpgradeScheduleRequestModel) GetVDAComponentsAndFeaturesRequestModelOk() (*CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel, bool)`

GetVDAComponentsAndFeaturesRequestModelOk returns a tuple with the VDAComponentsAndFeaturesRequestModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVDAComponentsAndFeaturesRequestModel

`func (o *CreateUpgradeScheduleRequestModel) SetVDAComponentsAndFeaturesRequestModel(v CreateUpgradeScheduleRequestModelVDAComponentsAndFeaturesRequestModel)`

SetVDAComponentsAndFeaturesRequestModel sets VDAComponentsAndFeaturesRequestModel field to given value.

### HasVDAComponentsAndFeaturesRequestModel

`func (o *CreateUpgradeScheduleRequestModel) HasVDAComponentsAndFeaturesRequestModel() bool`

HasVDAComponentsAndFeaturesRequestModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


