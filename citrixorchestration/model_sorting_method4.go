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

// checks if the SortingMethod4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SortingMethod4{}

// SortingMethod4 Filter search result sort criteria.
type SortingMethod4 struct {
	Property FilterProperty `json:"Property"`
	SortDirection ListSortDirection `json:"SortDirection"`
}

// NewSortingMethod4 instantiates a new SortingMethod4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSortingMethod4(property FilterProperty, sortDirection ListSortDirection) *SortingMethod4 {
	this := SortingMethod4{}
	this.Property = property
	this.SortDirection = sortDirection
	return &this
}

// NewSortingMethod4WithDefaults instantiates a new SortingMethod4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSortingMethod4WithDefaults() *SortingMethod4 {
	this := SortingMethod4{}
	return &this
}

// GetProperty returns the Property field value
func (o *SortingMethod4) GetProperty() FilterProperty {
	if o == nil {
		var ret FilterProperty
		return ret
	}

	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value
// and a boolean to check if the value has been set.
func (o *SortingMethod4) GetPropertyOk() (*FilterProperty, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Property, true
}

// SetProperty sets field value
func (o *SortingMethod4) SetProperty(v FilterProperty) {
	o.Property = v
}

// GetSortDirection returns the SortDirection field value
func (o *SortingMethod4) GetSortDirection() ListSortDirection {
	if o == nil {
		var ret ListSortDirection
		return ret
	}

	return o.SortDirection
}

// GetSortDirectionOk returns a tuple with the SortDirection field value
// and a boolean to check if the value has been set.
func (o *SortingMethod4) GetSortDirectionOk() (*ListSortDirection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortDirection, true
}

// SetSortDirection sets field value
func (o *SortingMethod4) SetSortDirection(v ListSortDirection) {
	o.SortDirection = v
}

func (o SortingMethod4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SortingMethod4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Property"] = o.Property
	toSerialize["SortDirection"] = o.SortDirection
	return toSerialize, nil
}

type NullableSortingMethod4 struct {
	value *SortingMethod4
	isSet bool
}

func (v NullableSortingMethod4) Get() *SortingMethod4 {
	return v.value
}

func (v *NullableSortingMethod4) Set(val *SortingMethod4) {
	v.value = val
	v.isSet = true
}

func (v NullableSortingMethod4) IsSet() bool {
	return v.isSet
}

func (v *NullableSortingMethod4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortingMethod4(val *SortingMethod4) *NullableSortingMethod4 {
	return &NullableSortingMethod4{value: val, isSet: true}
}

func (v NullableSortingMethod4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSortingMethod4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

