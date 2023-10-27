# BackupRestoreStorageRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileSpecificationBlobStorage** | Pointer to **NullableString** | Path to local file storage | [optional] 
**BackupRestoreCloudStorage** | Pointer to [**BackupRestoreCloudStorage**](BackupRestoreCloudStorage.md) |  | [optional] 
**PrimaryStorage** | Pointer to [**BackupRestoreBlobStorage**](BackupRestoreBlobStorage.md) |  | [optional] 
**SecondaryStorage** | Pointer to [**BackupRestoreBlobStorage**](BackupRestoreBlobStorage.md) |  | [optional] 

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

### GetFileSpecificationBlobStorage

`func (o *BackupRestoreStorageRequestModel) GetFileSpecificationBlobStorage() string`

GetFileSpecificationBlobStorage returns the FileSpecificationBlobStorage field if non-nil, zero value otherwise.

### GetFileSpecificationBlobStorageOk

`func (o *BackupRestoreStorageRequestModel) GetFileSpecificationBlobStorageOk() (*string, bool)`

GetFileSpecificationBlobStorageOk returns a tuple with the FileSpecificationBlobStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSpecificationBlobStorage

`func (o *BackupRestoreStorageRequestModel) SetFileSpecificationBlobStorage(v string)`

SetFileSpecificationBlobStorage sets FileSpecificationBlobStorage field to given value.

### HasFileSpecificationBlobStorage

`func (o *BackupRestoreStorageRequestModel) HasFileSpecificationBlobStorage() bool`

HasFileSpecificationBlobStorage returns a boolean if a field has been set.

### SetFileSpecificationBlobStorageNil

`func (o *BackupRestoreStorageRequestModel) SetFileSpecificationBlobStorageNil(b bool)`

 SetFileSpecificationBlobStorageNil sets the value for FileSpecificationBlobStorage to be an explicit nil

### UnsetFileSpecificationBlobStorage
`func (o *BackupRestoreStorageRequestModel) UnsetFileSpecificationBlobStorage()`

UnsetFileSpecificationBlobStorage ensures that no value is present for FileSpecificationBlobStorage, not even an explicit nil
### GetBackupRestoreCloudStorage

`func (o *BackupRestoreStorageRequestModel) GetBackupRestoreCloudStorage() BackupRestoreCloudStorage`

GetBackupRestoreCloudStorage returns the BackupRestoreCloudStorage field if non-nil, zero value otherwise.

### GetBackupRestoreCloudStorageOk

`func (o *BackupRestoreStorageRequestModel) GetBackupRestoreCloudStorageOk() (*BackupRestoreCloudStorage, bool)`

GetBackupRestoreCloudStorageOk returns a tuple with the BackupRestoreCloudStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRestoreCloudStorage

`func (o *BackupRestoreStorageRequestModel) SetBackupRestoreCloudStorage(v BackupRestoreCloudStorage)`

SetBackupRestoreCloudStorage sets BackupRestoreCloudStorage field to given value.

### HasBackupRestoreCloudStorage

`func (o *BackupRestoreStorageRequestModel) HasBackupRestoreCloudStorage() bool`

HasBackupRestoreCloudStorage returns a boolean if a field has been set.

### GetPrimaryStorage

`func (o *BackupRestoreStorageRequestModel) GetPrimaryStorage() BackupRestoreBlobStorage`

GetPrimaryStorage returns the PrimaryStorage field if non-nil, zero value otherwise.

### GetPrimaryStorageOk

`func (o *BackupRestoreStorageRequestModel) GetPrimaryStorageOk() (*BackupRestoreBlobStorage, bool)`

GetPrimaryStorageOk returns a tuple with the PrimaryStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryStorage

`func (o *BackupRestoreStorageRequestModel) SetPrimaryStorage(v BackupRestoreBlobStorage)`

SetPrimaryStorage sets PrimaryStorage field to given value.

### HasPrimaryStorage

`func (o *BackupRestoreStorageRequestModel) HasPrimaryStorage() bool`

HasPrimaryStorage returns a boolean if a field has been set.

### GetSecondaryStorage

`func (o *BackupRestoreStorageRequestModel) GetSecondaryStorage() BackupRestoreBlobStorage`

GetSecondaryStorage returns the SecondaryStorage field if non-nil, zero value otherwise.

### GetSecondaryStorageOk

`func (o *BackupRestoreStorageRequestModel) GetSecondaryStorageOk() (*BackupRestoreBlobStorage, bool)`

GetSecondaryStorageOk returns a tuple with the SecondaryStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryStorage

`func (o *BackupRestoreStorageRequestModel) SetSecondaryStorage(v BackupRestoreBlobStorage)`

SetSecondaryStorage sets SecondaryStorage field to given value.

### HasSecondaryStorage

`func (o *BackupRestoreStorageRequestModel) HasSecondaryStorage() bool`

HasSecondaryStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


