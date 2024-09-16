# BackupRestoreOptionsRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OneScheduledBackupPerDay** | Pointer to **NullableBool** |  | [optional] 
**AutoPinScheduleUid** | Pointer to **NullableInt32** |  | [optional] 
**AutoPinScheduleName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBackupRestoreOptionsRequestModel

`func NewBackupRestoreOptionsRequestModel() *BackupRestoreOptionsRequestModel`

NewBackupRestoreOptionsRequestModel instantiates a new BackupRestoreOptionsRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreOptionsRequestModelWithDefaults

`func NewBackupRestoreOptionsRequestModelWithDefaults() *BackupRestoreOptionsRequestModel`

NewBackupRestoreOptionsRequestModelWithDefaults instantiates a new BackupRestoreOptionsRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOneScheduledBackupPerDay

`func (o *BackupRestoreOptionsRequestModel) GetOneScheduledBackupPerDay() bool`

GetOneScheduledBackupPerDay returns the OneScheduledBackupPerDay field if non-nil, zero value otherwise.

### GetOneScheduledBackupPerDayOk

`func (o *BackupRestoreOptionsRequestModel) GetOneScheduledBackupPerDayOk() (*bool, bool)`

GetOneScheduledBackupPerDayOk returns a tuple with the OneScheduledBackupPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneScheduledBackupPerDay

`func (o *BackupRestoreOptionsRequestModel) SetOneScheduledBackupPerDay(v bool)`

SetOneScheduledBackupPerDay sets OneScheduledBackupPerDay field to given value.

### HasOneScheduledBackupPerDay

`func (o *BackupRestoreOptionsRequestModel) HasOneScheduledBackupPerDay() bool`

HasOneScheduledBackupPerDay returns a boolean if a field has been set.

### SetOneScheduledBackupPerDayNil

`func (o *BackupRestoreOptionsRequestModel) SetOneScheduledBackupPerDayNil(b bool)`

 SetOneScheduledBackupPerDayNil sets the value for OneScheduledBackupPerDay to be an explicit nil

### UnsetOneScheduledBackupPerDay
`func (o *BackupRestoreOptionsRequestModel) UnsetOneScheduledBackupPerDay()`

UnsetOneScheduledBackupPerDay ensures that no value is present for OneScheduledBackupPerDay, not even an explicit nil
### GetAutoPinScheduleUid

`func (o *BackupRestoreOptionsRequestModel) GetAutoPinScheduleUid() int32`

GetAutoPinScheduleUid returns the AutoPinScheduleUid field if non-nil, zero value otherwise.

### GetAutoPinScheduleUidOk

`func (o *BackupRestoreOptionsRequestModel) GetAutoPinScheduleUidOk() (*int32, bool)`

GetAutoPinScheduleUidOk returns a tuple with the AutoPinScheduleUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPinScheduleUid

`func (o *BackupRestoreOptionsRequestModel) SetAutoPinScheduleUid(v int32)`

SetAutoPinScheduleUid sets AutoPinScheduleUid field to given value.

### HasAutoPinScheduleUid

`func (o *BackupRestoreOptionsRequestModel) HasAutoPinScheduleUid() bool`

HasAutoPinScheduleUid returns a boolean if a field has been set.

### SetAutoPinScheduleUidNil

`func (o *BackupRestoreOptionsRequestModel) SetAutoPinScheduleUidNil(b bool)`

 SetAutoPinScheduleUidNil sets the value for AutoPinScheduleUid to be an explicit nil

### UnsetAutoPinScheduleUid
`func (o *BackupRestoreOptionsRequestModel) UnsetAutoPinScheduleUid()`

UnsetAutoPinScheduleUid ensures that no value is present for AutoPinScheduleUid, not even an explicit nil
### GetAutoPinScheduleName

`func (o *BackupRestoreOptionsRequestModel) GetAutoPinScheduleName() string`

GetAutoPinScheduleName returns the AutoPinScheduleName field if non-nil, zero value otherwise.

### GetAutoPinScheduleNameOk

`func (o *BackupRestoreOptionsRequestModel) GetAutoPinScheduleNameOk() (*string, bool)`

GetAutoPinScheduleNameOk returns a tuple with the AutoPinScheduleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPinScheduleName

`func (o *BackupRestoreOptionsRequestModel) SetAutoPinScheduleName(v string)`

SetAutoPinScheduleName sets AutoPinScheduleName field to given value.

### HasAutoPinScheduleName

`func (o *BackupRestoreOptionsRequestModel) HasAutoPinScheduleName() bool`

HasAutoPinScheduleName returns a boolean if a field has been set.

### SetAutoPinScheduleNameNil

`func (o *BackupRestoreOptionsRequestModel) SetAutoPinScheduleNameNil(b bool)`

 SetAutoPinScheduleNameNil sets the value for AutoPinScheduleName to be an explicit nil

### UnsetAutoPinScheduleName
`func (o *BackupRestoreOptionsRequestModel) UnsetAutoPinScheduleName()`

UnsetAutoPinScheduleName ensures that no value is present for AutoPinScheduleName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


