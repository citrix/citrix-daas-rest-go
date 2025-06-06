/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-onprem-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
)

// checks if the OfCDisplayClass10 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfCDisplayClass10{}

// OfCDisplayClass10 struct for OfCDisplayClass10
type OfCDisplayClass10 struct {
	NameToUuid map[string]string `json:"nameToUuid,omitempty"`
}

// NewOfCDisplayClass10 instantiates a new OfCDisplayClass10 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfCDisplayClass10() *OfCDisplayClass10 {
	this := OfCDisplayClass10{}
	return &this
}

// NewOfCDisplayClass10WithDefaults instantiates a new OfCDisplayClass10 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfCDisplayClass10WithDefaults() *OfCDisplayClass10 {
	this := OfCDisplayClass10{}
	return &this
}

// GetNameToUuid returns the NameToUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OfCDisplayClass10) GetNameToUuid() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.NameToUuid
}

// GetNameToUuidOk returns a tuple with the NameToUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OfCDisplayClass10) GetNameToUuidOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.NameToUuid) {
		return nil, false
	}
	return &o.NameToUuid, true
}

// HasNameToUuid returns a boolean if a field has been set.
func (o *OfCDisplayClass10) HasNameToUuid() bool {
	if o != nil && IsNil(o.NameToUuid) {
		return true
	}

	return false
}

// SetNameToUuid gets a reference to the given map[string]string and assigns it to the NameToUuid field.
func (o *OfCDisplayClass10) SetNameToUuid(v map[string]string) {
	o.NameToUuid = v
}

func (o OfCDisplayClass10) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfCDisplayClass10) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NameToUuid != nil {
		toSerialize["nameToUuid"] = o.NameToUuid
	}
	return toSerialize, nil
}

type NullableOfCDisplayClass10 struct {
	value *OfCDisplayClass10
	isSet bool
}

func (v NullableOfCDisplayClass10) Get() *OfCDisplayClass10 {
	return v.value
}

func (v *NullableOfCDisplayClass10) Set(val *OfCDisplayClass10) {
	v.value = val
	v.isSet = true
}

func (v NullableOfCDisplayClass10) IsSet() bool {
	return v.isSet
}

func (v *NullableOfCDisplayClass10) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfCDisplayClass10(val *OfCDisplayClass10) *NullableOfCDisplayClass10 {
	return &NullableOfCDisplayClass10{value: val, isSet: true}
}

func (v NullableOfCDisplayClass10) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfCDisplayClass10) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
