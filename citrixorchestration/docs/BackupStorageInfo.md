# BackupStorageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlobStorageType** | Pointer to [**BackupRestoreBlobStorage**](BackupRestoreBlobStorage.md) |  | [optional] 
**FileStorageFolder** | Pointer to **NullableString** | File Storage folder              | [optional] 
**Info1** | Pointer to **NullableString** | Info 1 for storage | [optional] 
**Info2** | Pointer to **NullableString** | Info 2 for storage | [optional] 
**Info3** | Pointer to **NullableString** | Info 3 for storage | [optional] 
**Info4** | Pointer to **NullableString** | Info 4 for storage | [optional] 
**Initialized** | Pointer to **bool** | Storage initialized by admin | [optional] 

## Methods

### NewBackupStorageInfo

`func NewBackupStorageInfo() *BackupStorageInfo`

NewBackupStorageInfo instantiates a new BackupStorageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupStorageInfoWithDefaults

`func NewBackupStorageInfoWithDefaults() *BackupStorageInfo`

NewBackupStorageInfoWithDefaults instantiates a new BackupStorageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlobStorageType

`func (o *BackupStorageInfo) GetBlobStorageType() BackupRestoreBlobStorage`

GetBlobStorageType returns the BlobStorageType field if non-nil, zero value otherwise.

### GetBlobStorageTypeOk

`func (o *BackupStorageInfo) GetBlobStorageTypeOk() (*BackupRestoreBlobStorage, bool)`

GetBlobStorageTypeOk returns a tuple with the BlobStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobStorageType

`func (o *BackupStorageInfo) SetBlobStorageType(v BackupRestoreBlobStorage)`

SetBlobStorageType sets BlobStorageType field to given value.

### HasBlobStorageType

`func (o *BackupStorageInfo) HasBlobStorageType() bool`

HasBlobStorageType returns a boolean if a field has been set.

### GetFileStorageFolder

`func (o *BackupStorageInfo) GetFileStorageFolder() string`

GetFileStorageFolder returns the FileStorageFolder field if non-nil, zero value otherwise.

### GetFileStorageFolderOk

`func (o *BackupStorageInfo) GetFileStorageFolderOk() (*string, bool)`

GetFileStorageFolderOk returns a tuple with the FileStorageFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileStorageFolder

`func (o *BackupStorageInfo) SetFileStorageFolder(v string)`

SetFileStorageFolder sets FileStorageFolder field to given value.

### HasFileStorageFolder

`func (o *BackupStorageInfo) HasFileStorageFolder() bool`

HasFileStorageFolder returns a boolean if a field has been set.

### SetFileStorageFolderNil

`func (o *BackupStorageInfo) SetFileStorageFolderNil(b bool)`

 SetFileStorageFolderNil sets the value for FileStorageFolder to be an explicit nil

### UnsetFileStorageFolder
`func (o *BackupStorageInfo) UnsetFileStorageFolder()`

UnsetFileStorageFolder ensures that no value is present for FileStorageFolder, not even an explicit nil
### GetInfo1

`func (o *BackupStorageInfo) GetInfo1() string`

GetInfo1 returns the Info1 field if non-nil, zero value otherwise.

### GetInfo1Ok

`func (o *BackupStorageInfo) GetInfo1Ok() (*string, bool)`

GetInfo1Ok returns a tuple with the Info1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo1

`func (o *BackupStorageInfo) SetInfo1(v string)`

SetInfo1 sets Info1 field to given value.

### HasInfo1

`func (o *BackupStorageInfo) HasInfo1() bool`

HasInfo1 returns a boolean if a field has been set.

### SetInfo1Nil

`func (o *BackupStorageInfo) SetInfo1Nil(b bool)`

 SetInfo1Nil sets the value for Info1 to be an explicit nil

### UnsetInfo1
`func (o *BackupStorageInfo) UnsetInfo1()`

UnsetInfo1 ensures that no value is present for Info1, not even an explicit nil
### GetInfo2

`func (o *BackupStorageInfo) GetInfo2() string`

GetInfo2 returns the Info2 field if non-nil, zero value otherwise.

### GetInfo2Ok

`func (o *BackupStorageInfo) GetInfo2Ok() (*string, bool)`

GetInfo2Ok returns a tuple with the Info2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo2

`func (o *BackupStorageInfo) SetInfo2(v string)`

SetInfo2 sets Info2 field to given value.

### HasInfo2

`func (o *BackupStorageInfo) HasInfo2() bool`

HasInfo2 returns a boolean if a field has been set.

### SetInfo2Nil

`func (o *BackupStorageInfo) SetInfo2Nil(b bool)`

 SetInfo2Nil sets the value for Info2 to be an explicit nil

### UnsetInfo2
`func (o *BackupStorageInfo) UnsetInfo2()`

UnsetInfo2 ensures that no value is present for Info2, not even an explicit nil
### GetInfo3

`func (o *BackupStorageInfo) GetInfo3() string`

GetInfo3 returns the Info3 field if non-nil, zero value otherwise.

### GetInfo3Ok

`func (o *BackupStorageInfo) GetInfo3Ok() (*string, bool)`

GetInfo3Ok returns a tuple with the Info3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo3

`func (o *BackupStorageInfo) SetInfo3(v string)`

SetInfo3 sets Info3 field to given value.

### HasInfo3

`func (o *BackupStorageInfo) HasInfo3() bool`

HasInfo3 returns a boolean if a field has been set.

### SetInfo3Nil

`func (o *BackupStorageInfo) SetInfo3Nil(b bool)`

 SetInfo3Nil sets the value for Info3 to be an explicit nil

### UnsetInfo3
`func (o *BackupStorageInfo) UnsetInfo3()`

UnsetInfo3 ensures that no value is present for Info3, not even an explicit nil
### GetInfo4

`func (o *BackupStorageInfo) GetInfo4() string`

GetInfo4 returns the Info4 field if non-nil, zero value otherwise.

### GetInfo4Ok

`func (o *BackupStorageInfo) GetInfo4Ok() (*string, bool)`

GetInfo4Ok returns a tuple with the Info4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo4

`func (o *BackupStorageInfo) SetInfo4(v string)`

SetInfo4 sets Info4 field to given value.

### HasInfo4

`func (o *BackupStorageInfo) HasInfo4() bool`

HasInfo4 returns a boolean if a field has been set.

### SetInfo4Nil

`func (o *BackupStorageInfo) SetInfo4Nil(b bool)`

 SetInfo4Nil sets the value for Info4 to be an explicit nil

### UnsetInfo4
`func (o *BackupStorageInfo) UnsetInfo4()`

UnsetInfo4 ensures that no value is present for Info4, not even an explicit nil
### GetInitialized

`func (o *BackupStorageInfo) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *BackupStorageInfo) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *BackupStorageInfo) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *BackupStorageInfo) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


