# BackupDoBackupResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionId** | **string** | Backup Execution Id: GUID | 
**Result** | [**BackupRestoreActionStartStatus**](BackupRestoreActionStartStatus.md) |  | 

## Methods

### NewBackupDoBackupResponseModel

`func NewBackupDoBackupResponseModel(executionId string, result BackupRestoreActionStartStatus, ) *BackupDoBackupResponseModel`

NewBackupDoBackupResponseModel instantiates a new BackupDoBackupResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupDoBackupResponseModelWithDefaults

`func NewBackupDoBackupResponseModelWithDefaults() *BackupDoBackupResponseModel`

NewBackupDoBackupResponseModelWithDefaults instantiates a new BackupDoBackupResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionId

`func (o *BackupDoBackupResponseModel) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BackupDoBackupResponseModel) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BackupDoBackupResponseModel) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetResult

`func (o *BackupDoBackupResponseModel) GetResult() BackupRestoreActionStartStatus`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *BackupDoBackupResponseModel) GetResultOk() (*BackupRestoreActionStartStatus, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *BackupDoBackupResponseModel) SetResult(v BackupRestoreActionStartStatus)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


