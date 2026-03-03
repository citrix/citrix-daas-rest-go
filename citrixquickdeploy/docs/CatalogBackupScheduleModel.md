# CatalogBackupScheduleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Daily** | Pointer to [**NullableCatalogBackupScheduleTime**](CatalogBackupScheduleTime.md) |  | [optional] 
**Weekly** | Pointer to [**NullableCatalogBackupScheduleTimeAndDate**](CatalogBackupScheduleTimeAndDate.md) |  | [optional] 

## Methods

### NewCatalogBackupScheduleModel

`func NewCatalogBackupScheduleModel() *CatalogBackupScheduleModel`

NewCatalogBackupScheduleModel instantiates a new CatalogBackupScheduleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogBackupScheduleModelWithDefaults

`func NewCatalogBackupScheduleModelWithDefaults() *CatalogBackupScheduleModel`

NewCatalogBackupScheduleModelWithDefaults instantiates a new CatalogBackupScheduleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaily

`func (o *CatalogBackupScheduleModel) GetDaily() CatalogBackupScheduleTime`

GetDaily returns the Daily field if non-nil, zero value otherwise.

### GetDailyOk

`func (o *CatalogBackupScheduleModel) GetDailyOk() (*CatalogBackupScheduleTime, bool)`

GetDailyOk returns a tuple with the Daily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaily

`func (o *CatalogBackupScheduleModel) SetDaily(v CatalogBackupScheduleTime)`

SetDaily sets Daily field to given value.

### HasDaily

`func (o *CatalogBackupScheduleModel) HasDaily() bool`

HasDaily returns a boolean if a field has been set.

### SetDailyNil

`func (o *CatalogBackupScheduleModel) SetDailyNil(b bool)`

 SetDailyNil sets the value for Daily to be an explicit nil

### UnsetDaily
`func (o *CatalogBackupScheduleModel) UnsetDaily()`

UnsetDaily ensures that no value is present for Daily, not even an explicit nil
### GetWeekly

`func (o *CatalogBackupScheduleModel) GetWeekly() CatalogBackupScheduleTimeAndDate`

GetWeekly returns the Weekly field if non-nil, zero value otherwise.

### GetWeeklyOk

`func (o *CatalogBackupScheduleModel) GetWeeklyOk() (*CatalogBackupScheduleTimeAndDate, bool)`

GetWeeklyOk returns a tuple with the Weekly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekly

`func (o *CatalogBackupScheduleModel) SetWeekly(v CatalogBackupScheduleTimeAndDate)`

SetWeekly sets Weekly field to given value.

### HasWeekly

`func (o *CatalogBackupScheduleModel) HasWeekly() bool`

HasWeekly returns a boolean if a field has been set.

### SetWeeklyNil

`func (o *CatalogBackupScheduleModel) SetWeeklyNil(b bool)`

 SetWeeklyNil sets the value for Weekly to be an explicit nil

### UnsetWeekly
`func (o *CatalogBackupScheduleModel) UnsetWeekly()`

UnsetWeekly ensures that no value is present for Weekly, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


