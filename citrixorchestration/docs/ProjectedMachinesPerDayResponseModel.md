# ProjectedMachinesPerDayResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Day** | **string** | Day | 
**TimeSlots** | [**[]ProjectedMachinesTimeSlotResponseModel**](ProjectedMachinesTimeSlotResponseModel.md) | TimeSlots | 

## Methods

### NewProjectedMachinesPerDayResponseModel

`func NewProjectedMachinesPerDayResponseModel(day string, timeSlots []ProjectedMachinesTimeSlotResponseModel, ) *ProjectedMachinesPerDayResponseModel`

NewProjectedMachinesPerDayResponseModel instantiates a new ProjectedMachinesPerDayResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectedMachinesPerDayResponseModelWithDefaults

`func NewProjectedMachinesPerDayResponseModelWithDefaults() *ProjectedMachinesPerDayResponseModel`

NewProjectedMachinesPerDayResponseModelWithDefaults instantiates a new ProjectedMachinesPerDayResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDay

`func (o *ProjectedMachinesPerDayResponseModel) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *ProjectedMachinesPerDayResponseModel) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *ProjectedMachinesPerDayResponseModel) SetDay(v string)`

SetDay sets Day field to given value.


### GetTimeSlots

`func (o *ProjectedMachinesPerDayResponseModel) GetTimeSlots() []ProjectedMachinesTimeSlotResponseModel`

GetTimeSlots returns the TimeSlots field if non-nil, zero value otherwise.

### GetTimeSlotsOk

`func (o *ProjectedMachinesPerDayResponseModel) GetTimeSlotsOk() (*[]ProjectedMachinesTimeSlotResponseModel, bool)`

GetTimeSlotsOk returns a tuple with the TimeSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlots

`func (o *ProjectedMachinesPerDayResponseModel) SetTimeSlots(v []ProjectedMachinesTimeSlotResponseModel)`

SetTimeSlots sets TimeSlots field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


