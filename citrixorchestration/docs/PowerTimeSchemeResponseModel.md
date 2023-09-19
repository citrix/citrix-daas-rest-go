# PowerTimeSchemeResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysOfWeek** | [**[]TimeSchemeDays**](TimeSchemeDays.md) | The pattern of days of the week that the power time scheme covers. | 
**Name** | **string** | The administrative name of the power time scheme. | 
**DisplayName** | **string** | The name of the power time scheme as displayed in the Studio console. | 
**PeakHours** | Pointer to **[]bool** | &#x60;DEPRECATED. Use &lt;see cref&#x3D;&#39;PeakTimeRanges&#39;/&gt; instead.&#x60; DEPRECATED.  Use PeakTimeRanges instead.  A set of 24 boolean flag values, one for each hour of the day. The first value in the array relates to 00:00 to 01:00, the next one to 01:00 to 02:00 and so on, with the last array element relating to 23:00 to 00:00. If the flag is &#x60;true&#x60; it means that the associated hour of the day is considered a peak time; if &#x60;false&#x60; it means that it is considered off-peak. | [optional] 
**PeakTimeRanges** | Pointer to **[]string** | List of peak time ranges during the day. | [optional] 
**PoolSize** | Pointer to **[]int32** | &#x60;DEPRECATED. Use &lt;see cref&#x3D;&#39;PoolUsingPercentage&#39;/&gt; instead.&#x60; DEPRECATED.  Use PoolSizeSchedule instead.  A set of 24 integer values, one for each hour of the day. The first value in the array relates to midnight to 00:59, the next one to 1 AM to 01:59 and so on, with the last array element relating to 11 PM to 11:59. The value defines the number of machines (either as an absolute number or a percentage of the machines in the delivery group, depending on the value of ) that are to be maintained in a running state, whether they are in use or not. A value of -1 has special meaning: pool size management does not apply during such hours. | [optional] 
**PoolSizeSchedule** | Pointer to [**[]PoolSizeScheduleResponseModel**](PoolSizeScheduleResponseModel.md) | List of pool size schedules during the day.  Each is specified as a time range and an indicator of the number of machines that should be powered on during that time range. | [optional] 
**PoolUsingPercentage** | Pointer to **bool** | Indicates whether the integer values in the pool size array are to be treated as absolute values (if this value is &#x60;false&#x60;) or as percentages of the number of machines in the delivery group (if this value is &#x60;true&#x60;). | [optional] 
**Id** | **string** | Id of the power time scheme. | 

## Methods

### NewPowerTimeSchemeResponseModel

`func NewPowerTimeSchemeResponseModel(daysOfWeek []TimeSchemeDays, name string, displayName string, id string, ) *PowerTimeSchemeResponseModel`

NewPowerTimeSchemeResponseModel instantiates a new PowerTimeSchemeResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerTimeSchemeResponseModelWithDefaults

`func NewPowerTimeSchemeResponseModelWithDefaults() *PowerTimeSchemeResponseModel`

NewPowerTimeSchemeResponseModelWithDefaults instantiates a new PowerTimeSchemeResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaysOfWeek

`func (o *PowerTimeSchemeResponseModel) GetDaysOfWeek() []TimeSchemeDays`

GetDaysOfWeek returns the DaysOfWeek field if non-nil, zero value otherwise.

### GetDaysOfWeekOk

`func (o *PowerTimeSchemeResponseModel) GetDaysOfWeekOk() (*[]TimeSchemeDays, bool)`

GetDaysOfWeekOk returns a tuple with the DaysOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOfWeek

`func (o *PowerTimeSchemeResponseModel) SetDaysOfWeek(v []TimeSchemeDays)`

SetDaysOfWeek sets DaysOfWeek field to given value.


### GetName

`func (o *PowerTimeSchemeResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PowerTimeSchemeResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PowerTimeSchemeResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *PowerTimeSchemeResponseModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PowerTimeSchemeResponseModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PowerTimeSchemeResponseModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetPeakHours

`func (o *PowerTimeSchemeResponseModel) GetPeakHours() []bool`

GetPeakHours returns the PeakHours field if non-nil, zero value otherwise.

### GetPeakHoursOk

`func (o *PowerTimeSchemeResponseModel) GetPeakHoursOk() (*[]bool, bool)`

GetPeakHoursOk returns a tuple with the PeakHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakHours

`func (o *PowerTimeSchemeResponseModel) SetPeakHours(v []bool)`

SetPeakHours sets PeakHours field to given value.

### HasPeakHours

`func (o *PowerTimeSchemeResponseModel) HasPeakHours() bool`

HasPeakHours returns a boolean if a field has been set.

### SetPeakHoursNil

`func (o *PowerTimeSchemeResponseModel) SetPeakHoursNil(b bool)`

 SetPeakHoursNil sets the value for PeakHours to be an explicit nil

### UnsetPeakHours
`func (o *PowerTimeSchemeResponseModel) UnsetPeakHours()`

UnsetPeakHours ensures that no value is present for PeakHours, not even an explicit nil
### GetPeakTimeRanges

`func (o *PowerTimeSchemeResponseModel) GetPeakTimeRanges() []string`

GetPeakTimeRanges returns the PeakTimeRanges field if non-nil, zero value otherwise.

### GetPeakTimeRangesOk

`func (o *PowerTimeSchemeResponseModel) GetPeakTimeRangesOk() (*[]string, bool)`

GetPeakTimeRangesOk returns a tuple with the PeakTimeRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakTimeRanges

`func (o *PowerTimeSchemeResponseModel) SetPeakTimeRanges(v []string)`

SetPeakTimeRanges sets PeakTimeRanges field to given value.

### HasPeakTimeRanges

`func (o *PowerTimeSchemeResponseModel) HasPeakTimeRanges() bool`

HasPeakTimeRanges returns a boolean if a field has been set.

### SetPeakTimeRangesNil

`func (o *PowerTimeSchemeResponseModel) SetPeakTimeRangesNil(b bool)`

 SetPeakTimeRangesNil sets the value for PeakTimeRanges to be an explicit nil

### UnsetPeakTimeRanges
`func (o *PowerTimeSchemeResponseModel) UnsetPeakTimeRanges()`

UnsetPeakTimeRanges ensures that no value is present for PeakTimeRanges, not even an explicit nil
### GetPoolSize

`func (o *PowerTimeSchemeResponseModel) GetPoolSize() []int32`

GetPoolSize returns the PoolSize field if non-nil, zero value otherwise.

### GetPoolSizeOk

`func (o *PowerTimeSchemeResponseModel) GetPoolSizeOk() (*[]int32, bool)`

GetPoolSizeOk returns a tuple with the PoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSize

`func (o *PowerTimeSchemeResponseModel) SetPoolSize(v []int32)`

SetPoolSize sets PoolSize field to given value.

### HasPoolSize

`func (o *PowerTimeSchemeResponseModel) HasPoolSize() bool`

HasPoolSize returns a boolean if a field has been set.

### SetPoolSizeNil

`func (o *PowerTimeSchemeResponseModel) SetPoolSizeNil(b bool)`

 SetPoolSizeNil sets the value for PoolSize to be an explicit nil

### UnsetPoolSize
`func (o *PowerTimeSchemeResponseModel) UnsetPoolSize()`

UnsetPoolSize ensures that no value is present for PoolSize, not even an explicit nil
### GetPoolSizeSchedule

`func (o *PowerTimeSchemeResponseModel) GetPoolSizeSchedule() []PoolSizeScheduleResponseModel`

GetPoolSizeSchedule returns the PoolSizeSchedule field if non-nil, zero value otherwise.

### GetPoolSizeScheduleOk

`func (o *PowerTimeSchemeResponseModel) GetPoolSizeScheduleOk() (*[]PoolSizeScheduleResponseModel, bool)`

GetPoolSizeScheduleOk returns a tuple with the PoolSizeSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSizeSchedule

`func (o *PowerTimeSchemeResponseModel) SetPoolSizeSchedule(v []PoolSizeScheduleResponseModel)`

SetPoolSizeSchedule sets PoolSizeSchedule field to given value.

### HasPoolSizeSchedule

`func (o *PowerTimeSchemeResponseModel) HasPoolSizeSchedule() bool`

HasPoolSizeSchedule returns a boolean if a field has been set.

### SetPoolSizeScheduleNil

`func (o *PowerTimeSchemeResponseModel) SetPoolSizeScheduleNil(b bool)`

 SetPoolSizeScheduleNil sets the value for PoolSizeSchedule to be an explicit nil

### UnsetPoolSizeSchedule
`func (o *PowerTimeSchemeResponseModel) UnsetPoolSizeSchedule()`

UnsetPoolSizeSchedule ensures that no value is present for PoolSizeSchedule, not even an explicit nil
### GetPoolUsingPercentage

`func (o *PowerTimeSchemeResponseModel) GetPoolUsingPercentage() bool`

GetPoolUsingPercentage returns the PoolUsingPercentage field if non-nil, zero value otherwise.

### GetPoolUsingPercentageOk

`func (o *PowerTimeSchemeResponseModel) GetPoolUsingPercentageOk() (*bool, bool)`

GetPoolUsingPercentageOk returns a tuple with the PoolUsingPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolUsingPercentage

`func (o *PowerTimeSchemeResponseModel) SetPoolUsingPercentage(v bool)`

SetPoolUsingPercentage sets PoolUsingPercentage field to given value.

### HasPoolUsingPercentage

`func (o *PowerTimeSchemeResponseModel) HasPoolUsingPercentage() bool`

HasPoolUsingPercentage returns a boolean if a field has been set.

### GetId

`func (o *PowerTimeSchemeResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PowerTimeSchemeResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PowerTimeSchemeResponseModel) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


