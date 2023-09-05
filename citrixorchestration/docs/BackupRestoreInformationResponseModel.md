# BackupRestoreInformationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageBackupDurationInSeconds** | **int32** | Average backup time in seconds used to control the WebStudio backup progress bar | 
**TotalBackupDurationSamples** | **int32** | Total backup duration samples | 
**MaxPinnedBackups** | **int32** | Max pinned backups. Pinned backups cannot be deleted; tjey must first be unpinned and then they can be manually deleted or automatically deleted if old and a new backup is done and the max backups exceeded | 
**MaxSavedBackups** | **int32** | Max saved backups. Unpinned backups will be removed when a new backup is completed and the number of saved backups exceeds this value | 
**TimeZone** | **string** | Customer&#39;s selected time zone | 
**TimeZoneOffset** | **int32** | Customer&#39;s selected time zone offset in seconds | 
**Initialized** | **bool** | Backup process initialized when true, uninitilized when false When true, tells WebStudio to prompt the admin to select the time zone to  schedule backups in and create the first backups schedule | 
**StaleStatusTimeoutInSeconds** | Pointer to **int32** | Stale status timeout in seconds; when current status is to old to be balid | [optional] 

## Methods

### NewBackupRestoreInformationResponseModel

`func NewBackupRestoreInformationResponseModel(averageBackupDurationInSeconds int32, totalBackupDurationSamples int32, maxPinnedBackups int32, maxSavedBackups int32, timeZone string, timeZoneOffset int32, initialized bool, ) *BackupRestoreInformationResponseModel`

NewBackupRestoreInformationResponseModel instantiates a new BackupRestoreInformationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreInformationResponseModelWithDefaults

`func NewBackupRestoreInformationResponseModelWithDefaults() *BackupRestoreInformationResponseModel`

NewBackupRestoreInformationResponseModelWithDefaults instantiates a new BackupRestoreInformationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageBackupDurationInSeconds

`func (o *BackupRestoreInformationResponseModel) GetAverageBackupDurationInSeconds() int32`

GetAverageBackupDurationInSeconds returns the AverageBackupDurationInSeconds field if non-nil, zero value otherwise.

### GetAverageBackupDurationInSecondsOk

`func (o *BackupRestoreInformationResponseModel) GetAverageBackupDurationInSecondsOk() (*int32, bool)`

GetAverageBackupDurationInSecondsOk returns a tuple with the AverageBackupDurationInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageBackupDurationInSeconds

`func (o *BackupRestoreInformationResponseModel) SetAverageBackupDurationInSeconds(v int32)`

SetAverageBackupDurationInSeconds sets AverageBackupDurationInSeconds field to given value.


### GetTotalBackupDurationSamples

`func (o *BackupRestoreInformationResponseModel) GetTotalBackupDurationSamples() int32`

GetTotalBackupDurationSamples returns the TotalBackupDurationSamples field if non-nil, zero value otherwise.

### GetTotalBackupDurationSamplesOk

`func (o *BackupRestoreInformationResponseModel) GetTotalBackupDurationSamplesOk() (*int32, bool)`

GetTotalBackupDurationSamplesOk returns a tuple with the TotalBackupDurationSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBackupDurationSamples

`func (o *BackupRestoreInformationResponseModel) SetTotalBackupDurationSamples(v int32)`

SetTotalBackupDurationSamples sets TotalBackupDurationSamples field to given value.


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


### GetTimeZone

`func (o *BackupRestoreInformationResponseModel) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *BackupRestoreInformationResponseModel) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *BackupRestoreInformationResponseModel) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.


### GetTimeZoneOffset

`func (o *BackupRestoreInformationResponseModel) GetTimeZoneOffset() int32`

GetTimeZoneOffset returns the TimeZoneOffset field if non-nil, zero value otherwise.

### GetTimeZoneOffsetOk

`func (o *BackupRestoreInformationResponseModel) GetTimeZoneOffsetOk() (*int32, bool)`

GetTimeZoneOffsetOk returns a tuple with the TimeZoneOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneOffset

`func (o *BackupRestoreInformationResponseModel) SetTimeZoneOffset(v int32)`

SetTimeZoneOffset sets TimeZoneOffset field to given value.


### GetInitialized

`func (o *BackupRestoreInformationResponseModel) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *BackupRestoreInformationResponseModel) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *BackupRestoreInformationResponseModel) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.


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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


