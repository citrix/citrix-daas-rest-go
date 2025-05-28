# GetTaskAsync200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | Pointer to [**NullableResourceConnectionTaskOperationType**](ResourceConnectionTaskOperationType.md) | Task Type | [optional] 
**DeploymentId** | Pointer to **NullableString** | Deployment Id this task is working on | [optional] 
**TaskType** | [**TaskType**](TaskType.md) | The type of task | 
**TaskId** | Pointer to **NullableString** | Task Id | [optional] 
**TaskState** | Pointer to [**NullableTaskState**](TaskState.md) | The state of the task | [optional] 
**Status** | Pointer to **NullableString** | Status of the task | [optional] 
**AccountType** | Pointer to [**AccountType**](AccountType.md) | The type of account the task is associated with | [optional] 
**StartedAt** | Pointer to **NullableTime** | DateTime when the task started | [optional] 
**CompletedAt** | Pointer to **NullableTime** | Datetime when the task completed | [optional] 
**Warnings** | Pointer to [**[]TaskWarning**](TaskWarning.md) | Warnings that occurred in task processing | [optional] 
**Errors** | Pointer to [**[]TaskError**](TaskError.md) | Errors that occurred in task processing | [optional] 
**TransactionId** | Pointer to **NullableString** | ID of the transaction the task is associated with | [optional] 
**ResourceConnectionId** | Pointer to **NullableString** | The ResourceConnectionId this task is working on | [optional] 

## Methods

### NewGetTaskAsync200Response

`func NewGetTaskAsync200Response(taskType TaskType, ) *GetTaskAsync200Response`

NewGetTaskAsync200Response instantiates a new GetTaskAsync200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTaskAsync200ResponseWithDefaults

`func NewGetTaskAsync200ResponseWithDefaults() *GetTaskAsync200Response`

NewGetTaskAsync200ResponseWithDefaults instantiates a new GetTaskAsync200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *GetTaskAsync200Response) GetOperation() ResourceConnectionTaskOperationType`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *GetTaskAsync200Response) GetOperationOk() (*ResourceConnectionTaskOperationType, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *GetTaskAsync200Response) SetOperation(v ResourceConnectionTaskOperationType)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *GetTaskAsync200Response) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### SetOperationNil

`func (o *GetTaskAsync200Response) SetOperationNil(b bool)`

 SetOperationNil sets the value for Operation to be an explicit nil

### UnsetOperation
`func (o *GetTaskAsync200Response) UnsetOperation()`

UnsetOperation ensures that no value is present for Operation, not even an explicit nil
### GetDeploymentId

`func (o *GetTaskAsync200Response) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *GetTaskAsync200Response) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *GetTaskAsync200Response) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *GetTaskAsync200Response) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### SetDeploymentIdNil

`func (o *GetTaskAsync200Response) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *GetTaskAsync200Response) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
### GetTaskType

`func (o *GetTaskAsync200Response) GetTaskType() TaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *GetTaskAsync200Response) GetTaskTypeOk() (*TaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *GetTaskAsync200Response) SetTaskType(v TaskType)`

SetTaskType sets TaskType field to given value.


### GetTaskId

`func (o *GetTaskAsync200Response) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *GetTaskAsync200Response) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *GetTaskAsync200Response) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *GetTaskAsync200Response) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *GetTaskAsync200Response) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *GetTaskAsync200Response) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetTaskState

`func (o *GetTaskAsync200Response) GetTaskState() TaskState`

GetTaskState returns the TaskState field if non-nil, zero value otherwise.

### GetTaskStateOk

`func (o *GetTaskAsync200Response) GetTaskStateOk() (*TaskState, bool)`

GetTaskStateOk returns a tuple with the TaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskState

`func (o *GetTaskAsync200Response) SetTaskState(v TaskState)`

SetTaskState sets TaskState field to given value.

### HasTaskState

`func (o *GetTaskAsync200Response) HasTaskState() bool`

HasTaskState returns a boolean if a field has been set.

### SetTaskStateNil

`func (o *GetTaskAsync200Response) SetTaskStateNil(b bool)`

 SetTaskStateNil sets the value for TaskState to be an explicit nil

### UnsetTaskState
`func (o *GetTaskAsync200Response) UnsetTaskState()`

UnsetTaskState ensures that no value is present for TaskState, not even an explicit nil
### GetStatus

`func (o *GetTaskAsync200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTaskAsync200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTaskAsync200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTaskAsync200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GetTaskAsync200Response) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GetTaskAsync200Response) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetAccountType

`func (o *GetTaskAsync200Response) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *GetTaskAsync200Response) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *GetTaskAsync200Response) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *GetTaskAsync200Response) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetStartedAt

`func (o *GetTaskAsync200Response) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GetTaskAsync200Response) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GetTaskAsync200Response) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GetTaskAsync200Response) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *GetTaskAsync200Response) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *GetTaskAsync200Response) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *GetTaskAsync200Response) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GetTaskAsync200Response) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GetTaskAsync200Response) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GetTaskAsync200Response) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *GetTaskAsync200Response) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *GetTaskAsync200Response) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetWarnings

`func (o *GetTaskAsync200Response) GetWarnings() []TaskWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *GetTaskAsync200Response) GetWarningsOk() (*[]TaskWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *GetTaskAsync200Response) SetWarnings(v []TaskWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *GetTaskAsync200Response) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *GetTaskAsync200Response) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *GetTaskAsync200Response) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
### GetErrors

`func (o *GetTaskAsync200Response) GetErrors() []TaskError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetTaskAsync200Response) GetErrorsOk() (*[]TaskError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetTaskAsync200Response) SetErrors(v []TaskError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetTaskAsync200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *GetTaskAsync200Response) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *GetTaskAsync200Response) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetTransactionId

`func (o *GetTaskAsync200Response) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *GetTaskAsync200Response) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *GetTaskAsync200Response) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *GetTaskAsync200Response) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *GetTaskAsync200Response) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *GetTaskAsync200Response) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetResourceConnectionId

`func (o *GetTaskAsync200Response) GetResourceConnectionId() string`

GetResourceConnectionId returns the ResourceConnectionId field if non-nil, zero value otherwise.

### GetResourceConnectionIdOk

`func (o *GetTaskAsync200Response) GetResourceConnectionIdOk() (*string, bool)`

GetResourceConnectionIdOk returns a tuple with the ResourceConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceConnectionId

`func (o *GetTaskAsync200Response) SetResourceConnectionId(v string)`

SetResourceConnectionId sets ResourceConnectionId field to given value.

### HasResourceConnectionId

`func (o *GetTaskAsync200Response) HasResourceConnectionId() bool`

HasResourceConnectionId returns a boolean if a field has been set.

### SetResourceConnectionIdNil

`func (o *GetTaskAsync200Response) SetResourceConnectionIdNil(b bool)`

 SetResourceConnectionIdNil sets the value for ResourceConnectionId to be an explicit nil

### UnsetResourceConnectionId
`func (o *GetTaskAsync200Response) UnsetResourceConnectionId()`

UnsetResourceConnectionId ensures that no value is present for ResourceConnectionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


