# GetCustomerAccountTaskAsync200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskType** | [**TaskType**](TaskType.md) |  | 
**TaskId** | Pointer to **NullableString** | Task Id | [optional] 
**TaskState** | Pointer to [**NullableTaskState**](TaskState.md) |  | [optional] 
**Status** | Pointer to **NullableString** | Status of the task | [optional] 
**AccountType** | Pointer to [**AccountType**](AccountType.md) |  | [optional] 
**StartedAt** | Pointer to **NullableTime** | DateTime when the task started | [optional] 
**CompletedAt** | Pointer to **NullableTime** | Datetime when the task completed | [optional] 
**Warnings** | Pointer to [**[]TaskWarning**](TaskWarning.md) | Warnings that occurred in task processing | [optional] 
**Errors** | Pointer to [**[]TaskError**](TaskError.md) | Errors that occurred in task processing | [optional] 
**TransactionId** | Pointer to **NullableString** | ID of the transaction the task is associated with | [optional] 

## Methods

### NewGetCustomerAccountTaskAsync200Response

`func NewGetCustomerAccountTaskAsync200Response(taskType TaskType, ) *GetCustomerAccountTaskAsync200Response`

NewGetCustomerAccountTaskAsync200Response instantiates a new GetCustomerAccountTaskAsync200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomerAccountTaskAsync200ResponseWithDefaults

`func NewGetCustomerAccountTaskAsync200ResponseWithDefaults() *GetCustomerAccountTaskAsync200Response`

NewGetCustomerAccountTaskAsync200ResponseWithDefaults instantiates a new GetCustomerAccountTaskAsync200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskType

`func (o *GetCustomerAccountTaskAsync200Response) GetTaskType() TaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *GetCustomerAccountTaskAsync200Response) GetTaskTypeOk() (*TaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *GetCustomerAccountTaskAsync200Response) SetTaskType(v TaskType)`

SetTaskType sets TaskType field to given value.


### GetTaskId

`func (o *GetCustomerAccountTaskAsync200Response) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *GetCustomerAccountTaskAsync200Response) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *GetCustomerAccountTaskAsync200Response) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *GetCustomerAccountTaskAsync200Response) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *GetCustomerAccountTaskAsync200Response) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *GetCustomerAccountTaskAsync200Response) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetTaskState

`func (o *GetCustomerAccountTaskAsync200Response) GetTaskState() TaskState`

GetTaskState returns the TaskState field if non-nil, zero value otherwise.

### GetTaskStateOk

`func (o *GetCustomerAccountTaskAsync200Response) GetTaskStateOk() (*TaskState, bool)`

GetTaskStateOk returns a tuple with the TaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskState

`func (o *GetCustomerAccountTaskAsync200Response) SetTaskState(v TaskState)`

SetTaskState sets TaskState field to given value.

### HasTaskState

`func (o *GetCustomerAccountTaskAsync200Response) HasTaskState() bool`

HasTaskState returns a boolean if a field has been set.

### SetTaskStateNil

`func (o *GetCustomerAccountTaskAsync200Response) SetTaskStateNil(b bool)`

 SetTaskStateNil sets the value for TaskState to be an explicit nil

### UnsetTaskState
`func (o *GetCustomerAccountTaskAsync200Response) UnsetTaskState()`

UnsetTaskState ensures that no value is present for TaskState, not even an explicit nil
### GetStatus

`func (o *GetCustomerAccountTaskAsync200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCustomerAccountTaskAsync200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCustomerAccountTaskAsync200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetCustomerAccountTaskAsync200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GetCustomerAccountTaskAsync200Response) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GetCustomerAccountTaskAsync200Response) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetAccountType

`func (o *GetCustomerAccountTaskAsync200Response) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *GetCustomerAccountTaskAsync200Response) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *GetCustomerAccountTaskAsync200Response) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *GetCustomerAccountTaskAsync200Response) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetStartedAt

`func (o *GetCustomerAccountTaskAsync200Response) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GetCustomerAccountTaskAsync200Response) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GetCustomerAccountTaskAsync200Response) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GetCustomerAccountTaskAsync200Response) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *GetCustomerAccountTaskAsync200Response) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *GetCustomerAccountTaskAsync200Response) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *GetCustomerAccountTaskAsync200Response) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GetCustomerAccountTaskAsync200Response) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GetCustomerAccountTaskAsync200Response) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GetCustomerAccountTaskAsync200Response) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *GetCustomerAccountTaskAsync200Response) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *GetCustomerAccountTaskAsync200Response) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetWarnings

`func (o *GetCustomerAccountTaskAsync200Response) GetWarnings() []TaskWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *GetCustomerAccountTaskAsync200Response) GetWarningsOk() (*[]TaskWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *GetCustomerAccountTaskAsync200Response) SetWarnings(v []TaskWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *GetCustomerAccountTaskAsync200Response) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *GetCustomerAccountTaskAsync200Response) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *GetCustomerAccountTaskAsync200Response) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
### GetErrors

`func (o *GetCustomerAccountTaskAsync200Response) GetErrors() []TaskError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetCustomerAccountTaskAsync200Response) GetErrorsOk() (*[]TaskError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetCustomerAccountTaskAsync200Response) SetErrors(v []TaskError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetCustomerAccountTaskAsync200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *GetCustomerAccountTaskAsync200Response) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *GetCustomerAccountTaskAsync200Response) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetTransactionId

`func (o *GetCustomerAccountTaskAsync200Response) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *GetCustomerAccountTaskAsync200Response) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *GetCustomerAccountTaskAsync200Response) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *GetCustomerAccountTaskAsync200Response) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *GetCustomerAccountTaskAsync200Response) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *GetCustomerAccountTaskAsync200Response) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


