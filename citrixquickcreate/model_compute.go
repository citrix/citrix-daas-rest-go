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

// checks if the Compute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Compute{}

// Compute struct for Compute
type Compute struct {
	Value NullableString `json:"value,omitempty"`
}

// NewCompute instantiates a new Compute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompute() *Compute {
	this := Compute{}
	return &this
}

// NewComputeWithDefaults instantiates a new Compute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeWithDefaults() *Compute {
	this := Compute{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Compute) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Compute) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *Compute) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *Compute) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *Compute) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *Compute) UnsetValue() {
	o.Value.Unset()
}

func (o Compute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Compute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return toSerialize, nil
}

type NullableCompute struct {
	value *Compute
	isSet bool
}

func (v NullableCompute) Get() *Compute {
	return v.value
}

func (v *NullableCompute) Set(val *Compute) {
	v.value = val
	v.isSet = true
}

func (v NullableCompute) IsSet() bool {
	return v.isSet
}

func (v *NullableCompute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompute(val *Compute) *NullableCompute {
	return &NullableCompute{value: val, isSet: true}
}

func (v NullableCompute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
