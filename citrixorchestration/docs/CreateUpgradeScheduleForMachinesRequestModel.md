# CreateUpgradeScheduleForMachinesRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MachineNameOrIds** | **[]string** | Machine list to create VDA upgrade schedules for. Item can be machine name or Id. | 
**StartDateTimeUtc** | Pointer to **time.Time** | UTC time to start the Vda upgrade. Must be a future time. If set to null, the upgrade should be started at once. | [optional] 
**DurationInHours** | **int32** | Timeout duration in hours. Valid range is 1 to 24. | 

## Methods

### NewCreateUpgradeScheduleForMachinesRequestModel

`func NewCreateUpgradeScheduleForMachinesRequestModel(machineNameOrIds []string, durationInHours int32, ) *CreateUpgradeScheduleForMachinesRequestModel`

NewCreateUpgradeScheduleForMachinesRequestModel instantiates a new CreateUpgradeScheduleForMachinesRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpgradeScheduleForMachinesRequestModelWithDefaults

`func NewCreateUpgradeScheduleForMachinesRequestModelWithDefaults() *CreateUpgradeScheduleForMachinesRequestModel`

NewCreateUpgradeScheduleForMachinesRequestModelWithDefaults instantiates a new CreateUpgradeScheduleForMachinesRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachineNameOrIds

`func (o *CreateUpgradeScheduleForMachinesRequestModel) GetMachineNameOrIds() []string`

GetMachineNameOrIds returns the MachineNameOrIds field if non-nil, zero value otherwise.

### GetMachineNameOrIdsOk

`func (o *CreateUpgradeScheduleForMachinesRequestModel) GetMachineNameOrIdsOk() (*[]string, bool)`

GetMachineNameOrIdsOk returns a tuple with the MachineNameOrIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineNameOrIds

`func (o *CreateUpgradeScheduleForMachinesRequestModel) SetMachineNameOrIds(v []string)`

SetMachineNameOrIds sets MachineNameOrIds field to given value.


### GetStartDateTimeUtc

`func (o *CreateUpgradeScheduleForMachinesRequestModel) GetStartDateTimeUtc() time.Time`

GetStartDateTimeUtc returns the StartDateTimeUtc field if non-nil, zero value otherwise.

### GetStartDateTimeUtcOk

`func (o *CreateUpgradeScheduleForMachinesRequestModel) GetStartDateTimeUtcOk() (*time.Time, bool)`

GetStartDateTimeUtcOk returns a tuple with the StartDateTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTimeUtc

`func (o *CreateUpgradeScheduleForMachinesRequestModel) SetStartDateTimeUtc(v time.Time)`

SetStartDateTimeUtc sets StartDateTimeUtc field to given value.

### HasStartDateTimeUtc

`func (o *CreateUpgradeScheduleForMachinesRequestModel) HasStartDateTimeUtc() bool`

HasStartDateTimeUtc returns a boolean if a field has been set.

### GetDurationInHours

`func (o *CreateUpgradeScheduleForMachinesRequestModel) GetDurationInHours() int32`

GetDurationInHours returns the DurationInHours field if non-nil, zero value otherwise.

### GetDurationInHoursOk

`func (o *CreateUpgradeScheduleForMachinesRequestModel) GetDurationInHoursOk() (*int32, bool)`

GetDurationInHoursOk returns a tuple with the DurationInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInHours

`func (o *CreateUpgradeScheduleForMachinesRequestModel) SetDurationInHours(v int32)`

SetDurationInHours sets DurationInHours field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


