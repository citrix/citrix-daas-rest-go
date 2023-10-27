# BackupRestoreInformationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxAutoPinnedBackups** | **int32** | Max automatic pinned backups via schedules. | 
**MaxPinnedBackups** | **int32** | Max pinned backups. Pinned backups cannot be deleted; they must first be unpinned and then they can be manually deleted or automatically deleted if old and a new backup is done and the max backups exceeded | 
**MaxSavedBackups** | **int32** | Max saved backups. Unpinned backups will be removed when a new backup is completed and the number of saved backups exceeds this value | 
**StaleStatusTimeoutInSeconds** | Pointer to **int32** | Stale status timeout in seconds; when current status is to old to be balid | [optional] 
**ComponentInformation** | Pointer to [**[]BackupRestoreSingleComponentInfo**](BackupRestoreSingleComponentInfo.md) | Component information | [optional] 

## Methods

### NewBackupRestoreInformationResponseModel

`func NewBackupRestoreInformationResponseModel(maxAutoPinnedBackups int32, maxPinnedBackups int32, maxSavedBackups int32, ) *BackupRestoreInformationResponseModel`

NewBackupRestoreInformationResponseModel instantiates a new BackupRestoreInformationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreInformationResponseModelWithDefaults

`func NewBackupRestoreInformationResponseModelWithDefaults() *BackupRestoreInformationResponseModel`

NewBackupRestoreInformationResponseModelWithDefaults instantiates a new BackupRestoreInformationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxAutoPinnedBackups

`func (o *BackupRestoreInformationResponseModel) GetMaxAutoPinnedBackups() int32`

GetMaxAutoPinnedBackups returns the MaxAutoPinnedBackups field if non-nil, zero value otherwise.

### GetMaxAutoPinnedBackupsOk

`func (o *BackupRestoreInformationResponseModel) GetMaxAutoPinnedBackupsOk() (*int32, bool)`

GetMaxAutoPinnedBackupsOk returns a tuple with the MaxAutoPinnedBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAutoPinnedBackups

`func (o *BackupRestoreInformationResponseModel) SetMaxAutoPinnedBackups(v int32)`

SetMaxAutoPinnedBackups sets MaxAutoPinnedBackups field to given value.


### GetMaxPinnedBackups

`func (o *BackupRestoreInformationResponseModel) GetMaxPinnedBackups() int32`

GetMaxPinnedBackups returns the MaxPinnedBackups field if non-nil, zero value otherwise.

### GetMaxPinnedBackupsOk

`func (o *BackupRestoreInformationResponseModel) GetMaxPinnedBackupsOk() (*int32, bool)`

GetMaxPinnedBackupsOk returns a tuple with the MaxPinnedBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPinnedBackups

`func (o *BackupRestoreInformationResponseModel) SetMaxPinnedBackups(v int32)`

SetMaxPinnedBackups sets MaxPinnedBackups field to given value.


### GetMaxSavedBackups

`func (o *BackupRestoreInformationResponseModel) GetMaxSavedBackups() int32`

GetMaxSavedBackups returns the MaxSavedBackups field if non-nil, zero value otherwise.

### GetMaxSavedBackupsOk

`func (o *BackupRestoreInformationResponseModel) GetMaxSavedBackupsOk() (*int32, bool)`

GetMaxSavedBackupsOk returns a tuple with the MaxSavedBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSavedBackups

`func (o *BackupRestoreInformationResponseModel) SetMaxSavedBackups(v int32)`

SetMaxSavedBackups sets MaxSavedBackups field to given value.


### GetStaleStatusTimeoutInSeconds

`func (o *BackupRestoreInformationResponseModel) GetStaleStatusTimeoutInSeconds() int32`

GetStaleStatusTimeoutInSeconds returns the StaleStatusTimeoutInSeconds field if non-nil, zero value otherwise.

### GetStaleStatusTimeoutInSecondsOk

`func (o *BackupRestoreInformationResponseModel) GetStaleStatusTimeoutInSecondsOk() (*int32, bool)`

GetStaleStatusTimeoutInSecondsOk returns a tuple with the StaleStatusTimeoutInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleStatusTimeoutInSeconds

`func (o *BackupRestoreInformationResponseModel) SetStaleStatusTimeoutInSeconds(v int32)`

SetStaleStatusTimeoutInSeconds sets StaleStatusTimeoutInSeconds field to given value.

### HasStaleStatusTimeoutInSeconds

`func (o *BackupRestoreInformationResponseModel) HasStaleStatusTimeoutInSeconds() bool`

HasStaleStatusTimeoutInSeconds returns a boolean if a field has been set.

### GetComponentInformation

`func (o *BackupRestoreInformationResponseModel) GetComponentInformation() []BackupRestoreSingleComponentInfo`

GetComponentInformation returns the ComponentInformation field if non-nil, zero value otherwise.

### GetComponentInformationOk

`func (o *BackupRestoreInformationResponseModel) GetComponentInformationOk() (*[]BackupRestoreSingleComponentInfo, bool)`

GetComponentInformationOk returns a tuple with the ComponentInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentInformation

`func (o *BackupRestoreInformationResponseModel) SetComponentInformation(v []BackupRestoreSingleComponentInfo)`

SetComponentInformation sets ComponentInformation field to given value.

### HasComponentInformation

`func (o *BackupRestoreInformationResponseModel) HasComponentInformation() bool`

HasComponentInformation returns a boolean if a field has been set.

### SetComponentInformationNil

`func (o *BackupRestoreInformationResponseModel) SetComponentInformationNil(b bool)`

 SetComponentInformationNil sets the value for ComponentInformation to be an explicit nil

### UnsetComponentInformation
`func (o *BackupRestoreInformationResponseModel) UnsetComponentInformation()`

UnsetComponentInformation ensures that no value is present for ComponentInformation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


