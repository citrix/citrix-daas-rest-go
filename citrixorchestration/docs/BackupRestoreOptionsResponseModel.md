# BackupRestoreOptionsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OneScheduledBackupPerDay** | Pointer to **NullableBool** |  | [optional] 
**Result** | Pointer to **bool** |  | [optional] 

## Methods

### NewBackupRestoreOptionsResponseModel

`func NewBackupRestoreOptionsResponseModel() *BackupRestoreOptionsResponseModel`

NewBackupRestoreOptionsResponseModel instantiates a new BackupRestoreOptionsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreOptionsResponseModelWithDefaults

`func NewBackupRestoreOptionsResponseModelWithDefaults() *BackupRestoreOptionsResponseModel`

NewBackupRestoreOptionsResponseModelWithDefaults instantiates a new BackupRestoreOptionsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOneScheduledBackupPerDay

`func (o *BackupRestoreOptionsResponseModel) GetOneScheduledBackupPerDay() bool`

GetOneScheduledBackupPerDay returns the OneScheduledBackupPerDay field if non-nil, zero value otherwise.

### GetOneScheduledBackupPerDayOk

`func (o *BackupRestoreOptionsResponseModel) GetOneScheduledBackupPerDayOk() (*bool, bool)`

GetOneScheduledBackupPerDayOk returns a tuple with the OneScheduledBackupPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneScheduledBackupPerDay

`func (o *BackupRestoreOptionsResponseModel) SetOneScheduledBackupPerDay(v bool)`

SetOneScheduledBackupPerDay sets OneScheduledBackupPerDay field to given value.

### HasOneScheduledBackupPerDay

`func (o *BackupRestoreOptionsResponseModel) HasOneScheduledBackupPerDay() bool`

HasOneScheduledBackupPerDay returns a boolean if a field has been set.

### SetOneScheduledBackupPerDayNil

`func (o *BackupRestoreOptionsResponseModel) SetOneScheduledBackupPerDayNil(b bool)`

 SetOneScheduledBackupPerDayNil sets the value for OneScheduledBackupPerDay to be an explicit nil

### UnsetOneScheduledBackupPerDay
`func (o *BackupRestoreOptionsResponseModel) UnsetOneScheduledBackupPerDay()`

UnsetOneScheduledBackupPerDay ensures that no value is present for OneScheduledBackupPerDay, not even an explicit nil
### GetResult

`func (o *BackupRestoreOptionsResponseModel) GetResult() bool`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *BackupRestoreOptionsResponseModel) GetResultOk() (*bool, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *BackupRestoreOptionsResponseModel) SetResult(v bool)`

SetResult sets Result field to given value.

### HasResult

`func (o *BackupRestoreOptionsResponseModel) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


