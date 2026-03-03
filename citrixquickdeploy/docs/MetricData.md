# MetricData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeStamp** | Pointer to **time.Time** |  | [optional] 
**Average** | Pointer to **NullableFloat64** |  | [optional] 
**Minimum** | Pointer to **NullableFloat64** |  | [optional] 
**Maximum** | Pointer to **NullableFloat64** |  | [optional] 
**Total** | Pointer to **NullableFloat64** |  | [optional] 
**Count** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewMetricData

`func NewMetricData() *MetricData`

NewMetricData instantiates a new MetricData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDataWithDefaults

`func NewMetricDataWithDefaults() *MetricData`

NewMetricDataWithDefaults instantiates a new MetricData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeStamp

`func (o *MetricData) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *MetricData) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *MetricData) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *MetricData) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.

### GetAverage

`func (o *MetricData) GetAverage() float64`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *MetricData) GetAverageOk() (*float64, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *MetricData) SetAverage(v float64)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *MetricData) HasAverage() bool`

HasAverage returns a boolean if a field has been set.

### SetAverageNil

`func (o *MetricData) SetAverageNil(b bool)`

 SetAverageNil sets the value for Average to be an explicit nil

### UnsetAverage
`func (o *MetricData) UnsetAverage()`

UnsetAverage ensures that no value is present for Average, not even an explicit nil
### GetMinimum

`func (o *MetricData) GetMinimum() float64`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *MetricData) GetMinimumOk() (*float64, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *MetricData) SetMinimum(v float64)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *MetricData) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### SetMinimumNil

`func (o *MetricData) SetMinimumNil(b bool)`

 SetMinimumNil sets the value for Minimum to be an explicit nil

### UnsetMinimum
`func (o *MetricData) UnsetMinimum()`

UnsetMinimum ensures that no value is present for Minimum, not even an explicit nil
### GetMaximum

`func (o *MetricData) GetMaximum() float64`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *MetricData) GetMaximumOk() (*float64, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *MetricData) SetMaximum(v float64)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *MetricData) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### SetMaximumNil

`func (o *MetricData) SetMaximumNil(b bool)`

 SetMaximumNil sets the value for Maximum to be an explicit nil

### UnsetMaximum
`func (o *MetricData) UnsetMaximum()`

UnsetMaximum ensures that no value is present for Maximum, not even an explicit nil
### GetTotal

`func (o *MetricData) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MetricData) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MetricData) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MetricData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *MetricData) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *MetricData) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetCount

`func (o *MetricData) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MetricData) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MetricData) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *MetricData) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *MetricData) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *MetricData) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


