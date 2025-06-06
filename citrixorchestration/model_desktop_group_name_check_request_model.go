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

// checks if the DesktopGroupNameCheckRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DesktopGroupNameCheckRequestModel{}

// DesktopGroupNameCheckRequestModel Request Model for Checking Desktop Group Name with Admin Folder
type DesktopGroupNameCheckRequestModel struct {
	// Desktop Group Name with Admin Folder
	Name NullableString `json:"Name,omitempty"`
}

// NewDesktopGroupNameCheckRequestModel instantiates a new DesktopGroupNameCheckRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesktopGroupNameCheckRequestModel() *DesktopGroupNameCheckRequestModel {
	this := DesktopGroupNameCheckRequestModel{}
	return &this
}

// NewDesktopGroupNameCheckRequestModelWithDefaults instantiates a new DesktopGroupNameCheckRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesktopGroupNameCheckRequestModelWithDefaults() *DesktopGroupNameCheckRequestModel {
	this := DesktopGroupNameCheckRequestModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DesktopGroupNameCheckRequestModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DesktopGroupNameCheckRequestModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DesktopGroupNameCheckRequestModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DesktopGroupNameCheckRequestModel) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *DesktopGroupNameCheckRequestModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DesktopGroupNameCheckRequestModel) UnsetName() {
	o.Name.Unset()
}

func (o DesktopGroupNameCheckRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DesktopGroupNameCheckRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableDesktopGroupNameCheckRequestModel struct {
	value *DesktopGroupNameCheckRequestModel
	isSet bool
}

func (v NullableDesktopGroupNameCheckRequestModel) Get() *DesktopGroupNameCheckRequestModel {
	return v.value
}

func (v *NullableDesktopGroupNameCheckRequestModel) Set(val *DesktopGroupNameCheckRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDesktopGroupNameCheckRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDesktopGroupNameCheckRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesktopGroupNameCheckRequestModel(val *DesktopGroupNameCheckRequestModel) *NullableDesktopGroupNameCheckRequestModel {
	return &NullableDesktopGroupNameCheckRequestModel{value: val, isSet: true}
}

func (v NullableDesktopGroupNameCheckRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesktopGroupNameCheckRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
