# CatalogBackupScheduleTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeZoneId** | Pointer to **string** | Timezone the backup should take place in | [optional] 
**Hour** | Pointer to **int32** | Hour of the day to initiate backup (backup will occur sometime within the hour) | [optional] 

## Methods

### NewCatalogBackupScheduleTime

`func NewCatalogBackupScheduleTime() *CatalogBackupScheduleTime`

NewCatalogBackupScheduleTime instantiates a new CatalogBackupScheduleTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogBackupScheduleTimeWithDefaults

`func NewCatalogBackupScheduleTimeWithDefaults() *CatalogBackupScheduleTime`

NewCatalogBackupScheduleTimeWithDefaults instantiates a new CatalogBackupScheduleTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeZoneId

`func (o *CatalogBackupScheduleTime) GetTimeZoneId() string`

GetTimeZoneId returns the TimeZoneId field if non-nil, zero value otherwise.

### GetTimeZoneIdOk

`func (o *CatalogBackupScheduleTime) GetTimeZoneIdOk() (*string, bool)`

GetTimeZoneIdOk returns a tuple with the TimeZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneId

`func (o *CatalogBackupScheduleTime) SetTimeZoneId(v string)`

SetTimeZoneId sets TimeZoneId field to given value.

### HasTimeZoneId

`func (o *CatalogBackupScheduleTime) HasTimeZoneId() bool`

HasTimeZoneId returns a boolean if a field has been set.

### GetHour

`func (o *CatalogBackupScheduleTime) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *CatalogBackupScheduleTime) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *CatalogBackupScheduleTime) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *CatalogBackupScheduleTime) HasHour() bool`

HasHour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


