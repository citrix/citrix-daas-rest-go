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

// checks if the UpdateImageVersionRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateImageVersionRequestModel{}

// UpdateImageVersionRequestModel Request object for updating image version.
type UpdateImageVersionRequestModel struct {
	// The description associated with the image version.
	Description NullableString `json:"Description,omitempty" validate:"regexp=^[A-Za-z]+$"`
}

// NewUpdateImageVersionRequestModel instantiates a new UpdateImageVersionRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateImageVersionRequestModel() *UpdateImageVersionRequestModel {
	this := UpdateImageVersionRequestModel{}
	return &this
}

// NewUpdateImageVersionRequestModelWithDefaults instantiates a new UpdateImageVersionRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateImageVersionRequestModelWithDefaults() *UpdateImageVersionRequestModel {
	this := UpdateImageVersionRequestModel{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateImageVersionRequestModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateImageVersionRequestModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateImageVersionRequestModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *UpdateImageVersionRequestModel) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UpdateImageVersionRequestModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UpdateImageVersionRequestModel) UnsetDescription() {
	o.Description.Unset()
}

func (o UpdateImageVersionRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateImageVersionRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	return toSerialize, nil
}

type NullableUpdateImageVersionRequestModel struct {
	value *UpdateImageVersionRequestModel
	isSet bool
}

func (v NullableUpdateImageVersionRequestModel) Get() *UpdateImageVersionRequestModel {
	return v.value
}

func (v *NullableUpdateImageVersionRequestModel) Set(val *UpdateImageVersionRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateImageVersionRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateImageVersionRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateImageVersionRequestModel(val *UpdateImageVersionRequestModel) *NullableUpdateImageVersionRequestModel {
	return &NullableUpdateImageVersionRequestModel{value: val, isSet: true}
}

func (v NullableUpdateImageVersionRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateImageVersionRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
