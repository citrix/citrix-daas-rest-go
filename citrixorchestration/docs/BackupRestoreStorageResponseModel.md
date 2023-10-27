# BackupRestoreStorageResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileSpecificationBlobStorage** | Pointer to **NullableString** | Path to local file storage | [optional] 
**BackupRestoreCloudStorage** | Pointer to [**BackupRestoreCloudStorage**](BackupRestoreCloudStorage.md) |  | [optional] 
**PrimaryStorage** | Pointer to [**BackupRestoreBlobStorage**](BackupRestoreBlobStorage.md) |  | [optional] 
**SecondaryStorage** | Pointer to [**BackupRestoreBlobStorage**](BackupRestoreBlobStorage.md) |  | [optional] 

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

### GetFileSpecificationBlobStorage

`func (o *BackupRestoreStorageResponseModel) GetFileSpecificationBlobStorage() string`

GetFileSpecificationBlobStorage returns the FileSpecificationBlobStorage field if non-nil, zero value otherwise.

### GetFileSpecificationBlobStorageOk

`func (o *BackupRestoreStorageResponseModel) GetFileSpecificationBlobStorageOk() (*string, bool)`

GetFileSpecificationBlobStorageOk returns a tuple with the FileSpecificationBlobStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSpecificationBlobStorage

`func (o *BackupRestoreStorageResponseModel) SetFileSpecificationBlobStorage(v string)`

SetFileSpecificationBlobStorage sets FileSpecificationBlobStorage field to given value.

### HasFileSpecificationBlobStorage

`func (o *BackupRestoreStorageResponseModel) HasFileSpecificationBlobStorage() bool`

HasFileSpecificationBlobStorage returns a boolean if a field has been set.

### SetFileSpecificationBlobStorageNil

`func (o *BackupRestoreStorageResponseModel) SetFileSpecificationBlobStorageNil(b bool)`

 SetFileSpecificationBlobStorageNil sets the value for FileSpecificationBlobStorage to be an explicit nil

### UnsetFileSpecificationBlobStorage
`func (o *BackupRestoreStorageResponseModel) UnsetFileSpecificationBlobStorage()`

UnsetFileSpecificationBlobStorage ensures that no value is present for FileSpecificationBlobStorage, not even an explicit nil
### GetBackupRestoreCloudStorage

`func (o *BackupRestoreStorageResponseModel) GetBackupRestoreCloudStorage() BackupRestoreCloudStorage`

GetBackupRestoreCloudStorage returns the BackupRestoreCloudStorage field if non-nil, zero value otherwise.

### GetBackupRestoreCloudStorageOk

`func (o *BackupRestoreStorageResponseModel) GetBackupRestoreCloudStorageOk() (*BackupRestoreCloudStorage, bool)`

GetBackupRestoreCloudStorageOk returns a tuple with the BackupRestoreCloudStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRestoreCloudStorage

`func (o *BackupRestoreStorageResponseModel) SetBackupRestoreCloudStorage(v BackupRestoreCloudStorage)`

SetBackupRestoreCloudStorage sets BackupRestoreCloudStorage field to given value.

### HasBackupRestoreCloudStorage

`func (o *BackupRestoreStorageResponseModel) HasBackupRestoreCloudStorage() bool`

HasBackupRestoreCloudStorage returns a boolean if a field has been set.

### GetPrimaryStorage

`func (o *BackupRestoreStorageResponseModel) GetPrimaryStorage() BackupRestoreBlobStorage`

GetPrimaryStorage returns the PrimaryStorage field if non-nil, zero value otherwise.

### GetPrimaryStorageOk

`func (o *BackupRestoreStorageResponseModel) GetPrimaryStorageOk() (*BackupRestoreBlobStorage, bool)`

GetPrimaryStorageOk returns a tuple with the PrimaryStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryStorage

`func (o *BackupRestoreStorageResponseModel) SetPrimaryStorage(v BackupRestoreBlobStorage)`

SetPrimaryStorage sets PrimaryStorage field to given value.

### HasPrimaryStorage

`func (o *BackupRestoreStorageResponseModel) HasPrimaryStorage() bool`

HasPrimaryStorage returns a boolean if a field has been set.

### GetSecondaryStorage

`func (o *BackupRestoreStorageResponseModel) GetSecondaryStorage() BackupRestoreBlobStorage`

GetSecondaryStorage returns the SecondaryStorage field if non-nil, zero value otherwise.

### GetSecondaryStorageOk

`func (o *BackupRestoreStorageResponseModel) GetSecondaryStorageOk() (*BackupRestoreBlobStorage, bool)`

GetSecondaryStorageOk returns a tuple with the SecondaryStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryStorage

`func (o *BackupRestoreStorageResponseModel) SetSecondaryStorage(v BackupRestoreBlobStorage)`

SetSecondaryStorage sets SecondaryStorage field to given value.

### HasSecondaryStorage

`func (o *BackupRestoreStorageResponseModel) HasSecondaryStorage() bool`

HasSecondaryStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


