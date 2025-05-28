# CatalogBackupScheduleTimeAndDate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeZoneId** | Pointer to **string** | Timezone the backup should take place in | [optional] 
**DayOfWeek** | Pointer to [**DayOfWeek**](DayOfWeek.md) | Day of the week to initiate backup | [optional] 
**Hour** | Pointer to **int32** | Hour of the day to initiate backup (backup will occur sometime within the hour) | [optional] 

## Methods

### NewCatalogBackupScheduleTimeAndDate

`func NewCatalogBackupScheduleTimeAndDate() *CatalogBackupScheduleTimeAndDate`

NewCatalogBackupScheduleTimeAndDate instantiates a new CatalogBackupScheduleTimeAndDate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogBackupScheduleTimeAndDateWithDefaults

`func NewCatalogBackupScheduleTimeAndDateWithDefaults() *CatalogBackupScheduleTimeAndDate`

NewCatalogBackupScheduleTimeAndDateWithDefaults instantiates a new CatalogBackupScheduleTimeAndDate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeZoneId

`func (o *CatalogBackupScheduleTimeAndDate) GetTimeZoneId() string`

GetTimeZoneId returns the TimeZoneId field if non-nil, zero value otherwise.

### GetTimeZoneIdOk

`func (o *CatalogBackupScheduleTimeAndDate) GetTimeZoneIdOk() (*string, bool)`

GetTimeZoneIdOk returns a tuple with the TimeZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneId

`func (o *CatalogBackupScheduleTimeAndDate) SetTimeZoneId(v string)`

SetTimeZoneId sets TimeZoneId field to given value.

### HasTimeZoneId

`func (o *CatalogBackupScheduleTimeAndDate) HasTimeZoneId() bool`

HasTimeZoneId returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *CatalogBackupScheduleTimeAndDate) GetDayOfWeek() DayOfWeek`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *CatalogBackupScheduleTimeAndDate) GetDayOfWeekOk() (*DayOfWeek, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *CatalogBackupScheduleTimeAndDate) SetDayOfWeek(v DayOfWeek)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *CatalogBackupScheduleTimeAndDate) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetHour

`func (o *CatalogBackupScheduleTimeAndDate) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *CatalogBackupScheduleTimeAndDate) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *CatalogBackupScheduleTimeAndDate) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *CatalogBackupScheduleTimeAndDate) HasHour() bool`

HasHour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


