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

// checks if the SupportedTimeZones type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupportedTimeZones{}

// SupportedTimeZones struct for SupportedTimeZones
type SupportedTimeZones struct {
	// List of timezones supported by .Net.
	Items []map[string]interface{} `json:"items,omitempty"`
}

// NewSupportedTimeZonesWithDefaults instantiates a new SupportedTimeZones object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportedTimeZonesWithDefaults() *SupportedTimeZones {
	this := SupportedTimeZones{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SupportedTimeZones) GetItems() []map[string]interface{} {
	if o == nil || IsNil(o.Items) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedTimeZones) GetItemsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// SetItems gets a reference to the given []map[string]interface{} and assigns it to the Items field.
func (o *SupportedTimeZones) SetItems(v []map[string]interface{}) {
	o.Items = v
}

func (o SupportedTimeZones) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupportedTimeZones) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableSupportedTimeZones struct {
	value *SupportedTimeZones
	isSet bool
}

func (v NullableSupportedTimeZones) Get() *SupportedTimeZones {
	return v.value
}

func (v *NullableSupportedTimeZones) Set(val *SupportedTimeZones) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedTimeZones) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedTimeZones) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedTimeZones(val *SupportedTimeZones) *NullableSupportedTimeZones {
	return &NullableSupportedTimeZones{value: val, isSet: true}
}

func (v NullableSupportedTimeZones) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedTimeZones) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
