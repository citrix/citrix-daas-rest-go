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

// checks if the TaskWarning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskWarning{}

// TaskWarning Deployment Task Warning
type TaskWarning struct {
	// Warning Id
	WarningId NullableString `json:"warningId,omitempty"`
	// Warning Message
	Message NullableString `json:"message,omitempty"`
}

// NewTaskWarning instantiates a new TaskWarning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskWarning() *TaskWarning {
	this := TaskWarning{}
	return &this
}

// NewTaskWarningWithDefaults instantiates a new TaskWarning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskWarningWithDefaults() *TaskWarning {
	this := TaskWarning{}
	return &this
}

// GetWarningId returns the WarningId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskWarning) GetWarningId() string {
	if o == nil || IsNil(o.WarningId.Get()) {
		var ret string
		return ret
	}
	return *o.WarningId.Get()
}

// GetWarningIdOk returns a tuple with the WarningId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskWarning) GetWarningIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WarningId.Get(), o.WarningId.IsSet()
}

// HasWarningId returns a boolean if a field has been set.
func (o *TaskWarning) HasWarningId() bool {
	if o != nil && o.WarningId.IsSet() {
		return true
	}

	return false
}

// SetWarningId gets a reference to the given NullableString and assigns it to the WarningId field.
func (o *TaskWarning) SetWarningId(v string) {
	o.WarningId.Set(&v)
}
// SetWarningIdNil sets the value for WarningId to be an explicit nil
func (o *TaskWarning) SetWarningIdNil() {
	o.WarningId.Set(nil)
}

// UnsetWarningId ensures that no value is present for WarningId, not even an explicit nil
func (o *TaskWarning) UnsetWarningId() {
	o.WarningId.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskWarning) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskWarning) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *TaskWarning) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *TaskWarning) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *TaskWarning) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *TaskWarning) UnsetMessage() {
	o.Message.Unset()
}

func (o TaskWarning) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskWarning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.WarningId.IsSet() {
		toSerialize["warningId"] = o.WarningId.Get()
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableTaskWarning struct {
	value *TaskWarning
	isSet bool
}

func (v NullableTaskWarning) Get() *TaskWarning {
	return v.value
}

func (v *NullableTaskWarning) Set(val *TaskWarning) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskWarning) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskWarning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskWarning(val *TaskWarning) *NullableTaskWarning {
	return &NullableTaskWarning{value: val, isSet: true}
}

func (v NullableTaskWarning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskWarning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

