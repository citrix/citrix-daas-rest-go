# BackupRestoreStorageResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageInfo** | Pointer to [**BackupStorageInfo**](BackupStorageInfo.md) |  | [optional] 
**MaximumBackups** | Pointer to **int32** | Maximum Backups | [optional] 
**MaximumPinnedBackups** | Pointer to **int32** | Maximum Pinned Backups  | [optional] 

## Methods

### NewBackupRestoreStorageResponseModel

`func NewBackupRestoreStorageResponseModel() *BackupRestoreStorageResponseModel`

NewBackupRestoreStorageResponseModel instantiates a new BackupRestoreStorageResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreStorageResponseModelWithDefaults

`func NewBackupRestoreStorageResponseModelWithDefaults() *BackupRestoreStorageResponseModel`

NewBackupRestoreStorageResponseModelWithDefaults instantiates a new BackupRestoreStorageResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageInfo

`func (o *BackupRestoreStorageResponseModel) GetStorageInfo() BackupStorageInfo`

GetStorageInfo returns the StorageInfo field if non-nil, zero value otherwise.

### GetStorageInfoOk

`func (o *BackupRestoreStorageResponseModel) GetStorageInfoOk() (*BackupStorageInfo, bool)`

GetStorageInfoOk returns a tuple with the StorageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInfo

`func (o *BackupRestoreStorageResponseModel) SetStorageInfo(v BackupStorageInfo)`

SetStorageInfo sets StorageInfo field to given value.

### HasStorageInfo

`func (o *BackupRestoreStorageResponseModel) HasStorageInfo() bool`

HasStorageInfo returns a boolean if a field has been set.

### GetMaximumBackups

`func (o *BackupRestoreStorageResponseModel) GetMaximumBackups() int32`

GetMaximumBackups returns the MaximumBackups field if non-nil, zero value otherwise.

### GetMaximumBackupsOk

`func (o *BackupRestoreStorageResponseModel) GetMaximumBackupsOk() (*int32, bool)`

GetMaximumBackupsOk returns a tuple with the MaximumBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumBackups

`func (o *BackupRestoreStorageResponseModel) SetMaximumBackups(v int32)`

SetMaximumBackups sets MaximumBackups field to given value.

### HasMaximumBackups

`func (o *BackupRestoreStorageResponseModel) HasMaximumBackups() bool`

HasMaximumBackups returns a boolean if a field has been set.

### GetMaximumPinnedBackups

`func (o *BackupRestoreStorageResponseModel) GetMaximumPinnedBackups() int32`

GetMaximumPinnedBackups returns the MaximumPinnedBackups field if non-nil, zero value otherwise.

### GetMaximumPinnedBackupsOk

`func (o *BackupRestoreStorageResponseModel) GetMaximumPinnedBackupsOk() (*int32, bool)`

GetMaximumPinnedBackupsOk returns a tuple with the MaximumPinnedBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPinnedBackups

`func (o *BackupRestoreStorageResponseModel) SetMaximumPinnedBackups(v int32)`

SetMaximumPinnedBackups sets MaximumPinnedBackups field to given value.

### HasMaximumPinnedBackups

`func (o *BackupRestoreStorageResponseModel) HasMaximumPinnedBackups() bool`

HasMaximumPinnedBackups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


