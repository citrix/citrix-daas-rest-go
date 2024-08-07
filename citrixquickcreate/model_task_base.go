/*
Quick Create Service for DaaS Quick Deploy with AWS WorkSpaces Core - Tech Preview

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: DaaSTechPreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickcreate

import (
	"encoding/json"
	"time"
)

// checks if the TaskBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskBase{}

// TaskBase Base class for Tasks
type TaskBase struct {
	TaskType TaskType `json:"taskType"`
	// Task Id
	TaskId NullableString `json:"taskId,omitempty"`
	TaskState NullableTaskState `json:"taskState,omitempty"`
	// Status of the task
	Status NullableString `json:"status,omitempty"`
	AccountType *AccountType `json:"accountType,omitempty"`
	// DateTime when the task started
	StartedAt NullableTime `json:"startedAt,omitempty"`
	// Datetime when the task completed
	CompletedAt NullableTime `json:"completedAt,omitempty"`
	// Warnings that occurred in task processing
	Warnings []TaskWarning `json:"warnings,omitempty"`
	// Errors that occurred in task processing
	Errors []TaskError `json:"errors,omitempty"`
	// ID of the transaction the task is associated with
	TransactionId NullableString `json:"transactionId,omitempty"`
}

// NewTaskBase instantiates a new TaskBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskBase(taskType TaskType) *TaskBase {
	this := TaskBase{}
	this.TaskType = taskType
	return &this
}

// NewTaskBaseWithDefaults instantiates a new TaskBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskBaseWithDefaults() *TaskBase {
	this := TaskBase{}
	return &this
}

// GetTaskType returns the TaskType field value
func (o *TaskBase) GetTaskType() TaskType {
	if o == nil {
		var ret TaskType
		return ret
	}

	return o.TaskType
}

// GetTaskTypeOk returns a tuple with the TaskType field value
// and a boolean to check if the value has been set.
func (o *TaskBase) GetTaskTypeOk() (*TaskType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskType, true
}

// SetTaskType sets field value
func (o *TaskBase) SetTaskType(v TaskType) {
	o.TaskType = v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskBase) GetTaskId() string {
	if o == nil || IsNil(o.TaskId.Get()) {
		var ret string
		return ret
	}
	return *o.TaskId.Get()
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskBase) GetTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskId.Get(), o.TaskId.IsSet()
}

// HasTaskId returns a boolean if a field has been set.
func (o *TaskBase) HasTaskId() bool {
	if o != nil && o.TaskId.IsSet() {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given NullableString and assigns it to the TaskId field.
func (o *TaskBase) SetTaskId(v string) {
	o.TaskId.Set(&v)
}
// SetTaskIdNil sets the value for TaskId to be an explicit nil
func (o *TaskBase) SetTaskIdNil() {
	o.TaskId.Set(nil)
}

// UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
func (o *TaskBase) UnsetTaskId() {
	o.TaskId.Unset()
}

// GetTaskState returns the TaskState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskBase) GetTaskState() TaskState {
	if o == nil || IsNil(o.TaskState.Get()) {
		var ret TaskState
		return ret
	}
	return *o.TaskState.Get()
}

// GetTaskStateOk returns a tuple with the TaskState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskBase) GetTaskStateOk() (*TaskState, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskState.Get(), o.TaskState.IsSet()
}

// HasTaskState returns a boolean if a field has been set.
func (o *TaskBase) HasTaskState() bool {
	if o != nil && o.TaskState.IsSet() {
		return true
	}

	return false
}

// SetTaskState gets a reference to the given NullableTaskState and assigns it to the TaskState field.
func (o *TaskBase) SetTaskState(v TaskState) {
	o.TaskState.Set(&v)
}
// SetTaskStateNil sets the value for TaskState to be an explicit nil
func (o *TaskBase) SetTaskStateNil() {
	o.TaskState.Set(nil)
}

// UnsetTaskState ensures that no value is present for TaskState, not even an explicit nil
func (o *TaskBase) UnsetTaskState() {
	o.TaskState.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskBase) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskBase) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *TaskBase) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *TaskBase) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *TaskBase) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *TaskBase) UnsetStatus() {
	o.Status.Unset()
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *TaskBase) GetAccountType() AccountType {
	if o == nil || IsNil(o.AccountType) {
		var ret AccountType
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskBase) GetAccountTypeOk() (*AccountType, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *TaskBase) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given AccountType and assigns it to the AccountType field.
func (o *TaskBase) SetAccountType(v AccountType) {
	o.AccountType = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskBase) GetStartedAt() time.Time {
	if o == nil || IsNil(o.StartedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskBase) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// HasStartedAt returns a boolean if a field has been set.
func (o *TaskBase) HasStartedAt() bool {
	if o != nil && o.StartedAt.IsSet() {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given NullableTime and assigns it to the StartedAt field.
func (o *TaskBase) SetStartedAt(v time.Time) {
	o.StartedAt.Set(&v)
}
// SetStartedAtNil sets the value for StartedAt to be an explicit nil
func (o *TaskBase) SetStartedAtNil() {
	o.StartedAt.Set(nil)
}

// UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
func (o *TaskBase) UnsetStartedAt() {
	o.StartedAt.Unset()
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskBase) GetCompletedAt() time.Time {
	if o == nil || IsNil(o.CompletedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskBase) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *TaskBase) HasCompletedAt() bool {
	if o != nil && o.CompletedAt.IsSet() {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given NullableTime and assigns it to the CompletedAt field.
func (o *TaskBase) SetCompletedAt(v time.Time) {
	o.CompletedAt.Set(&v)
}
// SetCompletedAtNil sets the value for CompletedAt to be an explicit nil
func (o *TaskBase) SetCompletedAtNil() {
	o.CompletedAt.Set(nil)
}

// UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
func (o *TaskBase) UnsetCompletedAt() {
	o.CompletedAt.Unset()
}

// GetWarnings returns the Warnings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskBase) GetWarnings() []TaskWarning {
	if o == nil {
		var ret []TaskWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskBase) GetWarningsOk() ([]TaskWarning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *TaskBase) HasWarnings() bool {
	if o != nil && IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []TaskWarning and assigns it to the Warnings field.
func (o *TaskBase) SetWarnings(v []TaskWarning) {
	o.Warnings = v
}

// GetErrors returns the Errors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskBase) GetErrors() []TaskError {
	if o == nil {
		var ret []TaskError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskBase) GetErrorsOk() ([]TaskError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *TaskBase) HasErrors() bool {
	if o != nil && IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []TaskError and assigns it to the Errors field.
func (o *TaskBase) SetErrors(v []TaskError) {
	o.Errors = v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskBase) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId.Get()) {
		var ret string
		return ret
	}
	return *o.TransactionId.Get()
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskBase) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionId.Get(), o.TransactionId.IsSet()
}

// HasTransactionId returns a boolean if a field has been set.
func (o *TaskBase) HasTransactionId() bool {
	if o != nil && o.TransactionId.IsSet() {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given NullableString and assigns it to the TransactionId field.
func (o *TaskBase) SetTransactionId(v string) {
	o.TransactionId.Set(&v)
}
// SetTransactionIdNil sets the value for TransactionId to be an explicit nil
func (o *TaskBase) SetTransactionIdNil() {
	o.TransactionId.Set(nil)
}

// UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
func (o *TaskBase) UnsetTransactionId() {
	o.TransactionId.Unset()
}

func (o TaskBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["taskType"] = o.TaskType
	if o.TaskId.IsSet() {
		toSerialize["taskId"] = o.TaskId.Get()
	}
	if o.TaskState.IsSet() {
		toSerialize["taskState"] = o.TaskState.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if !IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	if o.StartedAt.IsSet() {
		toSerialize["startedAt"] = o.StartedAt.Get()
	}
	if o.CompletedAt.IsSet() {
		toSerialize["completedAt"] = o.CompletedAt.Get()
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.TransactionId.IsSet() {
		toSerialize["transactionId"] = o.TransactionId.Get()
	}
	return toSerialize, nil
}

type NullableTaskBase struct {
	value *TaskBase
	isSet bool
}

func (v NullableTaskBase) Get() *TaskBase {
	return v.value
}

func (v *NullableTaskBase) Set(val *TaskBase) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskBase) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskBase(val *TaskBase) *NullableTaskBase {
	return &NullableTaskBase{value: val, isSet: true}
}

func (v NullableTaskBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

