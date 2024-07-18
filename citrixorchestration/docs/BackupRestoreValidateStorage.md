# BackupRestoreValidateStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Validated** | Pointer to **bool** |  | [optional] 
**ValidateStorageErrors** | Pointer to **map[string]string** |  | [optional] 
**BackupRestoreStorage** | Pointer to [**BackupRestoreStorageModel2**](BackupRestoreStorageModel2.md) |  | [optional] 

## Methods

### NewBackupRestoreValidateStorage

`func NewBackupRestoreValidateStorage() *BackupRestoreValidateStorage`

NewBackupRestoreValidateStorage instantiates a new BackupRestoreValidateStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreValidateStorageWithDefaults

`func NewBackupRestoreValidateStorageWithDefaults() *BackupRestoreValidateStorage`

NewBackupRestoreValidateStorageWithDefaults instantiates a new BackupRestoreValidateStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidated

`func (o *BackupRestoreValidateStorage) GetValidated() bool`

GetValidated returns the Validated field if non-nil, zero value otherwise.

### GetValidatedOk

`func (o *BackupRestoreValidateStorage) GetValidatedOk() (*bool, bool)`

GetValidatedOk returns a tuple with the Validated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidated

`func (o *BackupRestoreValidateStorage) SetValidated(v bool)`

SetValidated sets Validated field to given value.

### HasValidated

`func (o *BackupRestoreValidateStorage) HasValidated() bool`

HasValidated returns a boolean if a field has been set.

### GetValidateStorageErrors

`func (o *BackupRestoreValidateStorage) GetValidateStorageErrors() map[string]string`

GetValidateStorageErrors returns the ValidateStorageErrors field if non-nil, zero value otherwise.

### GetValidateStorageErrorsOk

`func (o *BackupRestoreValidateStorage) GetValidateStorageErrorsOk() (*map[string]string, bool)`

GetValidateStorageErrorsOk returns a tuple with the ValidateStorageErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateStorageErrors

`func (o *BackupRestoreValidateStorage) SetValidateStorageErrors(v map[string]string)`

SetValidateStorageErrors sets ValidateStorageErrors field to given value.

### HasValidateStorageErrors

`func (o *BackupRestoreValidateStorage) HasValidateStorageErrors() bool`

HasValidateStorageErrors returns a boolean if a field has been set.

### SetValidateStorageErrorsNil

`func (o *BackupRestoreValidateStorage) SetValidateStorageErrorsNil(b bool)`

 SetValidateStorageErrorsNil sets the value for ValidateStorageErrors to be an explicit nil

### UnsetValidateStorageErrors
`func (o *BackupRestoreValidateStorage) UnsetValidateStorageErrors()`

UnsetValidateStorageErrors ensures that no value is present for ValidateStorageErrors, not even an explicit nil
### GetBackupRestoreStorage

`func (o *BackupRestoreValidateStorage) GetBackupRestoreStorage() BackupRestoreStorageModel2`

GetBackupRestoreStorage returns the BackupRestoreStorage field if non-nil, zero value otherwise.

### GetBackupRestoreStorageOk

`func (o *BackupRestoreValidateStorage) GetBackupRestoreStorageOk() (*BackupRestoreStorageModel2, bool)`

GetBackupRestoreStorageOk returns a tuple with the BackupRestoreStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRestoreStorage

`func (o *BackupRestoreValidateStorage) SetBackupRestoreStorage(v BackupRestoreStorageModel2)`

SetBackupRestoreStorage sets BackupRestoreStorage field to given value.

### HasBackupRestoreStorage

`func (o *BackupRestoreValidateStorage) HasBackupRestoreStorage() bool`

HasBackupRestoreStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


