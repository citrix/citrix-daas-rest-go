/*
Quick Create Service for DaaS Quick Deploy with AWS WorkSpaces Core - Tech Preview

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: DaaSTechPreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickcreate

import (
	"encoding/json"
)

// checks if the CommissionTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommissionTask{}

// CommissionTask struct for CommissionTask
type CommissionTask struct {
	TaskBase
	// Task Type
	Operation NullableCommissionTaskOperationType `json:"operation,omitempty"`
}

type _CommissionTask CommissionTask

// NewCommissionTask instantiates a new CommissionTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommissionTask(taskType TaskType) *CommissionTask {
	this := CommissionTask{}
	this.TaskType = taskType
	return &this
}

// NewCommissionTaskWithDefaults instantiates a new CommissionTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommissionTaskWithDefaults() *CommissionTask {
	this := CommissionTask{}
	return &this
}

// GetOperation returns the Operation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommissionTask) GetOperation() CommissionTaskOperationType {
	if o == nil || IsNil(o.Operation.Get()) {
		var ret CommissionTaskOperationType
		return ret
	}
	return *o.Operation.Get()
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommissionTask) GetOperationOk() (*CommissionTaskOperationType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operation.Get(), o.Operation.IsSet()
}

// HasOperation returns a boolean if a field has been set.
func (o *CommissionTask) HasOperation() bool {
	if o != nil && o.Operation.IsSet() {
		return true
	}

	return false
}

// SetOperation gets a reference to the given NullableCommissionTaskOperationType and assigns it to the Operation field.
func (o *CommissionTask) SetOperation(v CommissionTaskOperationType) {
	o.Operation.Set(&v)
}

// SetOperationNil sets the value for Operation to be an explicit nil
func (o *CommissionTask) SetOperationNil() {
	o.Operation.Set(nil)
}

// UnsetOperation ensures that no value is present for Operation, not even an explicit nil
func (o *CommissionTask) UnsetOperation() {
	o.Operation.Unset()
}

func (o CommissionTask) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommissionTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTaskBase, errTaskBase := json.Marshal(o.TaskBase)
	if errTaskBase != nil {
		return map[string]interface{}{}, errTaskBase
	}
	errTaskBase = json.Unmarshal([]byte(serializedTaskBase), &toSerialize)
	if errTaskBase != nil {
		return map[string]interface{}{}, errTaskBase
	}
	if o.Operation.IsSet() {
		toSerialize["operation"] = o.Operation.Get()
	}
	return toSerialize, nil
}

type NullableCommissionTask struct {
	value *CommissionTask
	isSet bool
}

func (v NullableCommissionTask) Get() *CommissionTask {
	return v.value
}

func (v *NullableCommissionTask) Set(val *CommissionTask) {
	v.value = val
	v.isSet = true
}

func (v NullableCommissionTask) IsSet() bool {
	return v.isSet
}

func (v *NullableCommissionTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommissionTask(val *CommissionTask) *NullableCommissionTask {
	return &NullableCommissionTask{value: val, isSet: true}
}

func (v NullableCommissionTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommissionTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
