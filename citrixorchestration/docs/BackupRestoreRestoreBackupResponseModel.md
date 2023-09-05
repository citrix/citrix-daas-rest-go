# BackupRestoreRestoreBackupResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionId** | **string** | Restore Execution Id: GUID | 
**Result** | [**BackupRestoreActionStartStatus**](BackupRestoreActionStartStatus.md) |  | 

## Methods

### NewBackupRestoreRestoreBackupResponseModel

`func NewBackupRestoreRestoreBackupResponseModel(executionId string, result BackupRestoreActionStartStatus, ) *BackupRestoreRestoreBackupResponseModel`

NewBackupRestoreRestoreBackupResponseModel instantiates a new BackupRestoreRestoreBackupResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreRestoreBackupResponseModelWithDefaults

`func NewBackupRestoreRestoreBackupResponseModelWithDefaults() *BackupRestoreRestoreBackupResponseModel`

NewBackupRestoreRestoreBackupResponseModelWithDefaults instantiates a new BackupRestoreRestoreBackupResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionId

`func (o *BackupRestoreRestoreBackupResponseModel) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BackupRestoreRestoreBackupResponseModel) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BackupRestoreRestoreBackupResponseModel) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetResult

`func (o *BackupRestoreRestoreBackupResponseModel) GetResult() BackupRestoreActionStartStatus`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *BackupRestoreRestoreBackupResponseModel) GetResultOk() (*BackupRestoreActionStartStatus, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *BackupRestoreRestoreBackupResponseModel) SetResult(v BackupRestoreActionStartStatus)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


