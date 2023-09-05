# DesktopUsageResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **string** | Time when data point was collected. | 
**FormattedTime** | **string** | Formatted time when data point was collected. RFC 3339 compatible format. | 
**Usage** | **int32** | Specifies how many desktop are in use at the time the timestamp corresponds to. | 

## Methods

### NewDesktopUsageResponseModel

`func NewDesktopUsageResponseModel(time string, formattedTime string, usage int32, ) *DesktopUsageResponseModel`

NewDesktopUsageResponseModel instantiates a new DesktopUsageResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopUsageResponseModelWithDefaults

`func NewDesktopUsageResponseModelWithDefaults() *DesktopUsageResponseModel`

NewDesktopUsageResponseModelWithDefaults instantiates a new DesktopUsageResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *DesktopUsageResponseModel) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *DesktopUsageResponseModel) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *DesktopUsageResponseModel) SetTime(v string)`

SetTime sets Time field to given value.


### GetFormattedTime

`func (o *DesktopUsageResponseModel) GetFormattedTime() string`

GetFormattedTime returns the FormattedTime field if non-nil, zero value otherwise.

### GetFormattedTimeOk

`func (o *DesktopUsageResponseModel) GetFormattedTimeOk() (*string, bool)`

GetFormattedTimeOk returns a tuple with the FormattedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTime

`func (o *DesktopUsageResponseModel) SetFormattedTime(v string)`

SetFormattedTime sets FormattedTime field to given value.


### GetUsage

`func (o *DesktopUsageResponseModel) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *DesktopUsageResponseModel) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *DesktopUsageResponseModel) SetUsage(v int32)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


