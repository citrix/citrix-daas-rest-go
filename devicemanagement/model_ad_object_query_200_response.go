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

// checks if the AdObjectQuery200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdObjectQuery200Response{}

// AdObjectQuery200Response struct for AdObjectQuery200Response
type AdObjectQuery200Response struct {
	Items []MachineModel `json:"items"`
}

type _AdObjectQuery200Response AdObjectQuery200Response

// NewAdObjectQuery200Response instantiates a new AdObjectQuery200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdObjectQuery200Response(items []MachineModel) *AdObjectQuery200Response {
	this := AdObjectQuery200Response{}
	this.Items = items
	return &this
}

// NewAdObjectQuery200ResponseWithDefaults instantiates a new AdObjectQuery200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdObjectQuery200ResponseWithDefaults() *AdObjectQuery200Response {
	this := AdObjectQuery200Response{}
	return &this
}

// GetItems returns the Items field value
func (o *AdObjectQuery200Response) GetItems() []MachineModel {
	if o == nil {
		var ret []MachineModel
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *AdObjectQuery200Response) GetItemsOk() ([]MachineModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *AdObjectQuery200Response) SetItems(v []MachineModel) {
	o.Items = v
}

func (o AdObjectQuery200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdObjectQuery200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableAdObjectQuery200Response struct {
	value *AdObjectQuery200Response
	isSet bool
}

func (v NullableAdObjectQuery200Response) Get() *AdObjectQuery200Response {
	return v.value
}

func (v *NullableAdObjectQuery200Response) Set(val *AdObjectQuery200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAdObjectQuery200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAdObjectQuery200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdObjectQuery200Response(val *AdObjectQuery200Response) *NullableAdObjectQuery200Response {
	return &NullableAdObjectQuery200Response{value: val, isSet: true}
}

func (v NullableAdObjectQuery200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdObjectQuery200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
