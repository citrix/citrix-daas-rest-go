# BackupRestoreStorageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileSpecificationBlobStorage** | Pointer to **NullableString** | Path to local file storage | [optional] 
**BackupRestoreCloudStorage** | Pointer to [**BackupRestoreCloudStorage**](BackupRestoreCloudStorage.md) |  | [optional] 
**PrimaryStorage** | Pointer to [**BackupRestoreBlobStorage**](BackupRestoreBlobStorage.md) |  | [optional] 
**SecondaryStorage** | Pointer to [**BackupRestoreBlobStorage**](BackupRestoreBlobStorage.md) |  | [optional] 

## Methods

### NewBackupRestoreStorageModel

`func NewBackupRestoreStorageModel() *BackupRestoreStorageModel`

NewBackupRestoreStorageModel instantiates a new BackupRestoreStorageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreStorageModelWithDefaults

`func NewBackupRestoreStorageModelWithDefaults() *BackupRestoreStorageModel`

NewBackupRestoreStorageModelWithDefaults instantiates a new BackupRestoreStorageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileSpecificationBlobStorage

`func (o *BackupRestoreStorageModel) GetFileSpecificationBlobStorage() string`

GetFileSpecificationBlobStorage returns the FileSpecificationBlobStorage field if non-nil, zero value otherwise.

### GetFileSpecificationBlobStorageOk

`func (o *BackupRestoreStorageModel) GetFileSpecificationBlobStorageOk() (*string, bool)`

GetFileSpecificationBlobStorageOk returns a tuple with the FileSpecificationBlobStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSpecificationBlobStorage

`func (o *BackupRestoreStorageModel) SetFileSpecificationBlobStorage(v string)`

SetFileSpecificationBlobStorage sets FileSpecificationBlobStorage field to given value.

### HasFileSpecificationBlobStorage

`func (o *BackupRestoreStorageModel) HasFileSpecificationBlobStorage() bool`

HasFileSpecificationBlobStorage returns a boolean if a field has been set.

### SetFileSpecificationBlobStorageNil

`func (o *BackupRestoreStorageModel) SetFileSpecificationBlobStorageNil(b bool)`

 SetFileSpecificationBlobStorageNil sets the value for FileSpecificationBlobStorage to be an explicit nil

### UnsetFileSpecificationBlobStorage
`func (o *BackupRestoreStorageModel) UnsetFileSpecificationBlobStorage()`

UnsetFileSpecificationBlobStorage ensures that no value is present for FileSpecificationBlobStorage, not even an explicit nil
### GetBackupRestoreCloudStorage

`func (o *BackupRestoreStorageModel) GetBackupRestoreCloudStorage() BackupRestoreCloudStorage`

GetBackupRestoreCloudStorage returns the BackupRestoreCloudStorage field if non-nil, zero value otherwise.

### GetBackupRestoreCloudStorageOk

`func (o *BackupRestoreStorageModel) GetBackupRestoreCloudStorageOk() (*BackupRestoreCloudStorage, bool)`

GetBackupRestoreCloudStorageOk returns a tuple with the BackupRestoreCloudStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRestoreCloudStorage

`func (o *BackupRestoreStorageModel) SetBackupRestoreCloudStorage(v BackupRestoreCloudStorage)`

SetBackupRestoreCloudStorage sets BackupRestoreCloudStorage field to given value.

### HasBackupRestoreCloudStorage

`func (o *BackupRestoreStorageModel) HasBackupRestoreCloudStorage() bool`

HasBackupRestoreCloudStorage returns a boolean if a field has been set.

### GetPrimaryStorage

`func (o *BackupRestoreStorageModel) GetPrimaryStorage() BackupRestoreBlobStorage`

GetPrimaryStorage returns the PrimaryStorage field if non-nil, zero value otherwise.

### GetPrimaryStorageOk

`func (o *BackupRestoreStorageModel) GetPrimaryStorageOk() (*BackupRestoreBlobStorage, bool)`

GetPrimaryStorageOk returns a tuple with the PrimaryStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryStorage

`func (o *BackupRestoreStorageModel) SetPrimaryStorage(v BackupRestoreBlobStorage)`

SetPrimaryStorage sets PrimaryStorage field to given value.

### HasPrimaryStorage

`func (o *BackupRestoreStorageModel) HasPrimaryStorage() bool`

HasPrimaryStorage returns a boolean if a field has been set.

### GetSecondaryStorage

`func (o *BackupRestoreStorageModel) GetSecondaryStorage() BackupRestoreBlobStorage`

GetSecondaryStorage returns the SecondaryStorage field if non-nil, zero value otherwise.

### GetSecondaryStorageOk

`func (o *BackupRestoreStorageModel) GetSecondaryStorageOk() (*BackupRestoreBlobStorage, bool)`

GetSecondaryStorageOk returns a tuple with the SecondaryStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryStorage

`func (o *BackupRestoreStorageModel) SetSecondaryStorage(v BackupRestoreBlobStorage)`

SetSecondaryStorage sets SecondaryStorage field to given value.

### HasSecondaryStorage

`func (o *BackupRestoreStorageModel) HasSecondaryStorage() bool`

HasSecondaryStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


