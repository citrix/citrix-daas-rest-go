/*
Quick Create Service for DaaS Quick Deploy with AWS WorkSpaces Core - Tech Preview

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: DaaSTechPreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickcreate

import (
	"encoding/json"
	"fmt"
)

// TaskType Task Type
type TaskType string

// List of TaskType
const (
	TASKTYPE_ACCOUNT TaskType = "ACCOUNT"
	TASKTYPE_DEPLOYMENT TaskType = "DEPLOYMENT"
	TASKTYPE_RESOURCE_CONNECTION TaskType = "RESOURCE_CONNECTION"
	TASKTYPE_COMMISSION TaskType = "COMMISSION"
	TASKTYPE_WORKER_TEST TaskType = "WORKER_TEST"
)

// All allowed values of TaskType enum
var AllowedTaskTypeEnumValues = []TaskType{
	"ACCOUNT",
	"DEPLOYMENT",
	"RESOURCE_CONNECTION",
	"COMMISSION",
	"WORKER_TEST",
}

func (v *TaskType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TaskType(value)
	for _, existing := range AllowedTaskTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TaskType", value)
}

// NewTaskTypeFromValue returns a pointer to a valid TaskType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTaskTypeFromValue(v string) (*TaskType, error) {
	ev := TaskType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TaskType: valid values are %v", v, AllowedTaskTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TaskType) IsValid() bool {
	for _, existing := range AllowedTaskTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TaskType value
func (v TaskType) Ptr() *TaskType {
	return &v
}

type NullableTaskType struct {
	value *TaskType
	isSet bool
}

func (v NullableTaskType) Get() *TaskType {
	return v.value
}

func (v *NullableTaskType) Set(val *TaskType) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskType) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskType(val *TaskType) *NullableTaskType {
	return &NullableTaskType{value: val, isSet: true}
}

func (v NullableTaskType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
