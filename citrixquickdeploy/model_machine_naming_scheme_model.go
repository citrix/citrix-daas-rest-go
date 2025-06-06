/*
Citrix Virtual App & Desktop Catalog Service 147.0.26651.57932

Catalog Service

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickdeploy

import (
	"encoding/json"
)

// checks if the MachineNamingSchemeModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MachineNamingSchemeModel{}

// MachineNamingSchemeModel struct for MachineNamingSchemeModel
type MachineNamingSchemeModel struct {
	NamingScheme        *string `json:"namingScheme,omitempty"`
	IsSchemeTypeNumeric *bool   `json:"isSchemeTypeNumeric,omitempty"`
}

// NewMachineNamingSchemeModelWithDefaults instantiates a new MachineNamingSchemeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineNamingSchemeModelWithDefaults() *MachineNamingSchemeModel {
	this := MachineNamingSchemeModel{}
	return &this
}

// GetNamingScheme returns the NamingScheme field value if set, zero value otherwise.
func (o *MachineNamingSchemeModel) GetNamingScheme() string {
	if o == nil || IsNil(o.NamingScheme) {
		var ret string
		return ret
	}
	return *o.NamingScheme
}

// GetNamingSchemeOk returns a tuple with the NamingScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineNamingSchemeModel) GetNamingSchemeOk() (*string, bool) {
	if o == nil || IsNil(o.NamingScheme) {
		return nil, false
	}
	return o.NamingScheme, true
}

// SetNamingScheme gets a reference to the given string and assigns it to the NamingScheme field.
func (o *MachineNamingSchemeModel) SetNamingScheme(v string) {
	o.NamingScheme = &v
}

// GetIsSchemeTypeNumeric returns the IsSchemeTypeNumeric field value if set, zero value otherwise.
func (o *MachineNamingSchemeModel) GetIsSchemeTypeNumeric() bool {
	if o == nil || IsNil(o.IsSchemeTypeNumeric) {
		var ret bool
		return ret
	}
	return *o.IsSchemeTypeNumeric
}

// GetIsSchemeTypeNumericOk returns a tuple with the IsSchemeTypeNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineNamingSchemeModel) GetIsSchemeTypeNumericOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSchemeTypeNumeric) {
		return nil, false
	}
	return o.IsSchemeTypeNumeric, true
}

// SetIsSchemeTypeNumeric gets a reference to the given bool and assigns it to the IsSchemeTypeNumeric field.
func (o *MachineNamingSchemeModel) SetIsSchemeTypeNumeric(v bool) {
	o.IsSchemeTypeNumeric = &v
}

func (o MachineNamingSchemeModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MachineNamingSchemeModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NamingScheme) {
		toSerialize["namingScheme"] = o.NamingScheme
	}
	if !IsNil(o.IsSchemeTypeNumeric) {
		toSerialize["isSchemeTypeNumeric"] = o.IsSchemeTypeNumeric
	}
	return toSerialize, nil
}

type NullableMachineNamingSchemeModel struct {
	value *MachineNamingSchemeModel
	isSet bool
}

func (v NullableMachineNamingSchemeModel) Get() *MachineNamingSchemeModel {
	return v.value
}

func (v *NullableMachineNamingSchemeModel) Set(val *MachineNamingSchemeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineNamingSchemeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineNamingSchemeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineNamingSchemeModel(val *MachineNamingSchemeModel) *NullableMachineNamingSchemeModel {
	return &NullableMachineNamingSchemeModel{value: val, isSet: true}
}

func (v NullableMachineNamingSchemeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineNamingSchemeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
