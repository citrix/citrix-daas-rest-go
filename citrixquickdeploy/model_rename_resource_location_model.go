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

// checks if the RenameResourceLocationModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenameResourceLocationModel{}

// RenameResourceLocationModel struct for RenameResourceLocationModel
type RenameResourceLocationModel struct {
	Name *string `json:"name,omitempty"`
}

// NewRenameResourceLocationModelWithDefaults instantiates a new RenameResourceLocationModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenameResourceLocationModelWithDefaults() *RenameResourceLocationModel {
	this := RenameResourceLocationModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RenameResourceLocationModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenameResourceLocationModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RenameResourceLocationModel) SetName(v string) {
	o.Name = &v
}

func (o RenameResourceLocationModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenameResourceLocationModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableRenameResourceLocationModel struct {
	value *RenameResourceLocationModel
	isSet bool
}

func (v NullableRenameResourceLocationModel) Get() *RenameResourceLocationModel {
	return v.value
}

func (v *NullableRenameResourceLocationModel) Set(val *RenameResourceLocationModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRenameResourceLocationModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRenameResourceLocationModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenameResourceLocationModel(val *RenameResourceLocationModel) *NullableRenameResourceLocationModel {
	return &NullableRenameResourceLocationModel{value: val, isSet: true}
}

func (v NullableRenameResourceLocationModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenameResourceLocationModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
