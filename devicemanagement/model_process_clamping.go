/*
WEM Public API Guide

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package devicemanagement

import (
	"encoding/json"
)

// checks if the ProcessClamping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessClamping{}

// ProcessClamping struct for ProcessClamping
type ProcessClamping struct {
	ProcessName *string `json:"processName,omitempty"`
	Percentage *int32 `json:"percentage,omitempty"`
}

// NewProcessClamping instantiates a new ProcessClamping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessClamping() *ProcessClamping {
	this := ProcessClamping{}
	return &this
}

// NewProcessClampingWithDefaults instantiates a new ProcessClamping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessClampingWithDefaults() *ProcessClamping {
	this := ProcessClamping{}
	return &this
}

// GetProcessName returns the ProcessName field value if set, zero value otherwise.
func (o *ProcessClamping) GetProcessName() string {
	if o == nil || IsNil(o.ProcessName) {
		var ret string
		return ret
	}
	return *o.ProcessName
}

// GetProcessNameOk returns a tuple with the ProcessName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessClamping) GetProcessNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessName) {
		return nil, false
	}
	return o.ProcessName, true
}

// HasProcessName returns a boolean if a field has been set.
func (o *ProcessClamping) HasProcessName() bool {
	if o != nil && !IsNil(o.ProcessName) {
		return true
	}

	return false
}

// SetProcessName gets a reference to the given string and assigns it to the ProcessName field.
func (o *ProcessClamping) SetProcessName(v string) {
	o.ProcessName = &v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *ProcessClamping) GetPercentage() int32 {
	if o == nil || IsNil(o.Percentage) {
		var ret int32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessClamping) GetPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *ProcessClamping) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given int32 and assigns it to the Percentage field.
func (o *ProcessClamping) SetPercentage(v int32) {
	o.Percentage = &v
}

func (o ProcessClamping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessClamping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProcessName) {
		toSerialize["processName"] = o.ProcessName
	}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return toSerialize, nil
}

type NullableProcessClamping struct {
	value *ProcessClamping
	isSet bool
}

func (v NullableProcessClamping) Get() *ProcessClamping {
	return v.value
}

func (v *NullableProcessClamping) Set(val *ProcessClamping) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessClamping) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessClamping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessClamping(val *ProcessClamping) *NullableProcessClamping {
	return &NullableProcessClamping{value: val, isSet: true}
}

func (v NullableProcessClamping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessClamping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

