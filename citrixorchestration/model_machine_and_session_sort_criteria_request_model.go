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

// checks if the MachineAndSessionSortCriteriaRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MachineAndSessionSortCriteriaRequestModel{}

// MachineAndSessionSortCriteriaRequestModel Machine sort criteria.
type MachineAndSessionSortCriteriaRequestModel struct {
	Property      MachineAndSessionSearchProperty `json:"Property"`
	SortDirection ListSortDirection               `json:"SortDirection"`
}

// NewMachineAndSessionSortCriteriaRequestModel instantiates a new MachineAndSessionSortCriteriaRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMachineAndSessionSortCriteriaRequestModel(property MachineAndSessionSearchProperty, sortDirection ListSortDirection) *MachineAndSessionSortCriteriaRequestModel {
	this := MachineAndSessionSortCriteriaRequestModel{}
	this.Property = property
	this.SortDirection = sortDirection
	return &this
}

// NewMachineAndSessionSortCriteriaRequestModelWithDefaults instantiates a new MachineAndSessionSortCriteriaRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineAndSessionSortCriteriaRequestModelWithDefaults() *MachineAndSessionSortCriteriaRequestModel {
	this := MachineAndSessionSortCriteriaRequestModel{}
	return &this
}

// GetProperty returns the Property field value
func (o *MachineAndSessionSortCriteriaRequestModel) GetProperty() MachineAndSessionSearchProperty {
	if o == nil {
		var ret MachineAndSessionSearchProperty
		return ret
	}

	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value
// and a boolean to check if the value has been set.
func (o *MachineAndSessionSortCriteriaRequestModel) GetPropertyOk() (*MachineAndSessionSearchProperty, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Property, true
}

// SetProperty sets field value
func (o *MachineAndSessionSortCriteriaRequestModel) SetProperty(v MachineAndSessionSearchProperty) {
	o.Property = v
}

// GetSortDirection returns the SortDirection field value
func (o *MachineAndSessionSortCriteriaRequestModel) GetSortDirection() ListSortDirection {
	if o == nil {
		var ret ListSortDirection
		return ret
	}

	return o.SortDirection
}

// GetSortDirectionOk returns a tuple with the SortDirection field value
// and a boolean to check if the value has been set.
func (o *MachineAndSessionSortCriteriaRequestModel) GetSortDirectionOk() (*ListSortDirection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortDirection, true
}

// SetSortDirection sets field value
func (o *MachineAndSessionSortCriteriaRequestModel) SetSortDirection(v ListSortDirection) {
	o.SortDirection = v
}

func (o MachineAndSessionSortCriteriaRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MachineAndSessionSortCriteriaRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Property"] = o.Property
	toSerialize["SortDirection"] = o.SortDirection
	return toSerialize, nil
}

type NullableMachineAndSessionSortCriteriaRequestModel struct {
	value *MachineAndSessionSortCriteriaRequestModel
	isSet bool
}

func (v NullableMachineAndSessionSortCriteriaRequestModel) Get() *MachineAndSessionSortCriteriaRequestModel {
	return v.value
}

func (v *NullableMachineAndSessionSortCriteriaRequestModel) Set(val *MachineAndSessionSortCriteriaRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineAndSessionSortCriteriaRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineAndSessionSortCriteriaRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineAndSessionSortCriteriaRequestModel(val *MachineAndSessionSortCriteriaRequestModel) *NullableMachineAndSessionSortCriteriaRequestModel {
	return &NullableMachineAndSessionSortCriteriaRequestModel{value: val, isSet: true}
}

func (v NullableMachineAndSessionSortCriteriaRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineAndSessionSortCriteriaRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
