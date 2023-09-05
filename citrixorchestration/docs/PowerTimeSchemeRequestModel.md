# PowerTimeSchemeRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysOfWeek** | Pointer to [**[]TimeSchemeDays**](TimeSchemeDays.md) | The pattern of days of the week that the power time scheme covers. | [optional] 
**Name** | Pointer to **string** | The administrative name of the power time scheme. | [optional] 
**DisplayName** | Pointer to **string** | The name of the power time scheme as displayed in the Studio console. | [optional] 
**PeakHours** | Pointer to **[]bool** | &#x60;DEPRECATED. Use &lt;see cref&#x3D;&#39;PeakTimeRanges&#39;/&gt; instead.&#x60; DEPRECATED.  Use PeakTimeRanges instead.  A set of 24 boolean flag values, one for each hour of the day. The first value in the array relates to midnight to 00:59, the next one to 1 AM to 01:59 and so on, with the last array element relating to 11 PM to 11:59. If the flag is &#x60;true&#x60; it means that the associated hour of the day is considered a peak time; if &#x60;false&#x60; it means that it is considered off-peak. | [optional] 
**PeakTimeRanges** | Pointer to **[]string** | List of peak time ranges during the day. | [optional] 
**PoolSize** | Pointer to **[]int32** | &#x60;DEPRECATED. Use &lt;see cref&#x3D;&#39;PoolUsingPercentage&#39;/&gt; instead.&#x60; DEPRECATED.  Use PoolSizeSchedule instead.  A set of 24 integer values, one for each hour of the day. The first value in the array relates to midnight to 00:59, the next one to 1 AM to 01:59 and so on, with the last array element relating to 11 PM to 11:59. The value defines the number of machines (either as an absolute number or a percentage of the machines in the delivery group, depending on the value of ) that are to be maintained in a running state, whether they are in use or not. A value of &#x60;-1&#x60; has special meaning: pool size management does not apply during such hours. | [optional] 
**PoolSizeSchedule** | Pointer to [**[]PoolSizeScheduleRequestModel**](PoolSizeScheduleRequestModel.md) | List of pool size schedules during the day.  Each is specified as a time range and an indicator of the number of machines that should be powered on during that time range. | [optional] 
**PoolUsingPercentage** | Pointer to **bool** | Indicates whether the integer values in the pool size array are to be treated as absolute values (if this value is &#x60;false&#x60;) or as percentages of the number of machines in the delivery group (if this value is &#x60;true&#x60;). | [optional] [default to false]

## Methods

### NewPowerTimeSchemeRequestModel

`func NewPowerTimeSchemeRequestModel() *PowerTimeSchemeRequestModel`

NewPowerTimeSchemeRequestModel instantiates a new PowerTimeSchemeRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerTimeSchemeRequestModelWithDefaults

`func NewPowerTimeSchemeRequestModelWithDefaults() *PowerTimeSchemeRequestModel`

NewPowerTimeSchemeRequestModelWithDefaults instantiates a new PowerTimeSchemeRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaysOfWeek

`func (o *PowerTimeSchemeRequestModel) GetDaysOfWeek() []TimeSchemeDays`

GetDaysOfWeek returns the DaysOfWeek field if non-nil, zero value otherwise.

### GetDaysOfWeekOk

`func (o *PowerTimeSchemeRequestModel) GetDaysOfWeekOk() (*[]TimeSchemeDays, bool)`

GetDaysOfWeekOk returns a tuple with the DaysOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOfWeek

`func (o *PowerTimeSchemeRequestModel) SetDaysOfWeek(v []TimeSchemeDays)`

SetDaysOfWeek sets DaysOfWeek field to given value.

### HasDaysOfWeek

`func (o *PowerTimeSchemeRequestModel) HasDaysOfWeek() bool`

HasDaysOfWeek returns a boolean if a field has been set.

### GetName

`func (o *PowerTimeSchemeRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PowerTimeSchemeRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PowerTimeSchemeRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PowerTimeSchemeRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *PowerTimeSchemeRequestModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PowerTimeSchemeRequestModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PowerTimeSchemeRequestModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PowerTimeSchemeRequestModel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPeakHours

`func (o *PowerTimeSchemeRequestModel) GetPeakHours() []bool`

GetPeakHours returns the PeakHours field if non-nil, zero value otherwise.

### GetPeakHoursOk

`func (o *PowerTimeSchemeRequestModel) GetPeakHoursOk() (*[]bool, bool)`

GetPeakHoursOk returns a tuple with the PeakHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakHours

`func (o *PowerTimeSchemeRequestModel) SetPeakHours(v []bool)`

SetPeakHours sets PeakHours field to given value.

### HasPeakHours

`func (o *PowerTimeSchemeRequestModel) HasPeakHours() bool`

HasPeakHours returns a boolean if a field has been set.

### GetPeakTimeRanges

`func (o *PowerTimeSchemeRequestModel) GetPeakTimeRanges() []string`

GetPeakTimeRanges returns the PeakTimeRanges field if non-nil, zero value otherwise.

### GetPeakTimeRangesOk

`func (o *PowerTimeSchemeRequestModel) GetPeakTimeRangesOk() (*[]string, bool)`

GetPeakTimeRangesOk returns a tuple with the PeakTimeRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakTimeRanges

`func (o *PowerTimeSchemeRequestModel) SetPeakTimeRanges(v []string)`

SetPeakTimeRanges sets PeakTimeRanges field to given value.

### HasPeakTimeRanges

`func (o *PowerTimeSchemeRequestModel) HasPeakTimeRanges() bool`

HasPeakTimeRanges returns a boolean if a field has been set.

### GetPoolSize

`func (o *PowerTimeSchemeRequestModel) GetPoolSize() []int32`

GetPoolSize returns the PoolSize field if non-nil, zero value otherwise.

### GetPoolSizeOk

`func (o *PowerTimeSchemeRequestModel) GetPoolSizeOk() (*[]int32, bool)`

GetPoolSizeOk returns a tuple with the PoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSize

`func (o *PowerTimeSchemeRequestModel) SetPoolSize(v []int32)`

SetPoolSize sets PoolSize field to given value.

### HasPoolSize

`func (o *PowerTimeSchemeRequestModel) HasPoolSize() bool`

HasPoolSize returns a boolean if a field has been set.

### GetPoolSizeSchedule

`func (o *PowerTimeSchemeRequestModel) GetPoolSizeSchedule() []PoolSizeScheduleRequestModel`

GetPoolSizeSchedule returns the PoolSizeSchedule field if non-nil, zero value otherwise.

### GetPoolSizeScheduleOk

`func (o *PowerTimeSchemeRequestModel) GetPoolSizeScheduleOk() (*[]PoolSizeScheduleRequestModel, bool)`

GetPoolSizeScheduleOk returns a tuple with the PoolSizeSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSizeSchedule

`func (o *PowerTimeSchemeRequestModel) SetPoolSizeSchedule(v []PoolSizeScheduleRequestModel)`

SetPoolSizeSchedule sets PoolSizeSchedule field to given value.

### HasPoolSizeSchedule

`func (o *PowerTimeSchemeRequestModel) HasPoolSizeSchedule() bool`

HasPoolSizeSchedule returns a boolean if a field has been set.

### GetPoolUsingPercentage

`func (o *PowerTimeSchemeRequestModel) GetPoolUsingPercentage() bool`

GetPoolUsingPercentage returns the PoolUsingPercentage field if non-nil, zero value otherwise.

### GetPoolUsingPercentageOk

`func (o *PowerTimeSchemeRequestModel) GetPoolUsingPercentageOk() (*bool, bool)`

GetPoolUsingPercentageOk returns a tuple with the PoolUsingPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolUsingPercentage

`func (o *PowerTimeSchemeRequestModel) SetPoolUsingPercentage(v bool)`

SetPoolUsingPercentage sets PoolUsingPercentage field to given value.

### HasPoolUsingPercentage

`func (o *PowerTimeSchemeRequestModel) HasPoolUsingPercentage() bool`

HasPoolUsingPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


