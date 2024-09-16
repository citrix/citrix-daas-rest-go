# BackupRestoreStorageRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageInfo** | Pointer to [**BackupStorageInfo**](BackupStorageInfo.md) |  | [optional] 
**MaximumBackups** | Pointer to **int32** | Maximum Backups | [optional] 
**MaximumPinnedBackups** | Pointer to **int32** | Maximum Pinned Backups  | [optional] 
**MaximumAutoPinnedBackups** | Pointer to **int32** | Maximum Auto-Pinned Backups  | [optional] 

## Methods

### NewBackupRestoreStorageRequestModel

`func NewBackupRestoreStorageRequestModel() *BackupRestoreStorageRequestModel`

NewBackupRestoreStorageRequestModel instantiates a new BackupRestoreStorageRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreStorageRequestModelWithDefaults

`func NewBackupRestoreStorageRequestModelWithDefaults() *BackupRestoreStorageRequestModel`

NewBackupRestoreStorageRequestModelWithDefaults instantiates a new BackupRestoreStorageRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageInfo

`func (o *BackupRestoreStorageRequestModel) GetStorageInfo() BackupStorageInfo`

GetStorageInfo returns the StorageInfo field if non-nil, zero value otherwise.

### GetStorageInfoOk

`func (o *BackupRestoreStorageRequestModel) GetStorageInfoOk() (*BackupStorageInfo, bool)`

GetStorageInfoOk returns a tuple with the StorageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInfo

`func (o *BackupRestoreStorageRequestModel) SetStorageInfo(v BackupStorageInfo)`

SetStorageInfo sets StorageInfo field to given value.

### HasStorageInfo

`func (o *BackupRestoreStorageRequestModel) HasStorageInfo() bool`

HasStorageInfo returns a boolean if a field has been set.

### GetMaximumBackups

`func (o *BackupRestoreStorageRequestModel) GetMaximumBackups() int32`

GetMaximumBackups returns the MaximumBackups field if non-nil, zero value otherwise.

### GetMaximumBackupsOk

`func (o *BackupRestoreStorageRequestModel) GetMaximumBackupsOk() (*int32, bool)`

GetMaximumBackupsOk returns a tuple with the MaximumBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumBackups

`func (o *BackupRestoreStorageRequestModel) SetMaximumBackups(v int32)`

SetMaximumBackups sets MaximumBackups field to given value.

### HasMaximumBackups

`func (o *BackupRestoreStorageRequestModel) HasMaximumBackups() bool`

HasMaximumBackups returns a boolean if a field has been set.

### GetMaximumPinnedBackups

`func (o *BackupRestoreStorageRequestModel) GetMaximumPinnedBackups() int32`

GetMaximumPinnedBackups returns the MaximumPinnedBackups field if non-nil, zero value otherwise.

### GetMaximumPinnedBackupsOk

`func (o *BackupRestoreStorageRequestModel) GetMaximumPinnedBackupsOk() (*int32, bool)`

GetMaximumPinnedBackupsOk returns a tuple with the MaximumPinnedBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPinnedBackups

`func (o *BackupRestoreStorageRequestModel) SetMaximumPinnedBackups(v int32)`

SetMaximumPinnedBackups sets MaximumPinnedBackups field to given value.

### HasMaximumPinnedBackups

`func (o *BackupRestoreStorageRequestModel) HasMaximumPinnedBackups() bool`

HasMaximumPinnedBackups returns a boolean if a field has been set.

### GetMaximumAutoPinnedBackups

`func (o *BackupRestoreStorageRequestModel) GetMaximumAutoPinnedBackups() int32`

GetMaximumAutoPinnedBackups returns the MaximumAutoPinnedBackups field if non-nil, zero value otherwise.

### GetMaximumAutoPinnedBackupsOk

`func (o *BackupRestoreStorageRequestModel) GetMaximumAutoPinnedBackupsOk() (*int32, bool)`

GetMaximumAutoPinnedBackupsOk returns a tuple with the MaximumAutoPinnedBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAutoPinnedBackups

`func (o *BackupRestoreStorageRequestModel) SetMaximumAutoPinnedBackups(v int32)`

SetMaximumAutoPinnedBackups sets MaximumAutoPinnedBackups field to given value.

### HasMaximumAutoPinnedBackups

`func (o *BackupRestoreStorageRequestModel) HasMaximumAutoPinnedBackups() bool`

HasMaximumAutoPinnedBackups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


