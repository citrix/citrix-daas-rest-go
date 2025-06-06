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

// checks if the SystemUtilitySettingQuery200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemUtilitySettingQuery200Response{}

// SystemUtilitySettingQuery200Response struct for SystemUtilitySettingQuery200Response
type SystemUtilitySettingQuery200Response struct {
	Items []SystemOptimizationModel `json:"items"`
}

type _SystemUtilitySettingQuery200Response SystemUtilitySettingQuery200Response

// NewSystemUtilitySettingQuery200Response instantiates a new SystemUtilitySettingQuery200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemUtilitySettingQuery200Response(items []SystemOptimizationModel) *SystemUtilitySettingQuery200Response {
	this := SystemUtilitySettingQuery200Response{}
	this.Items = items
	return &this
}

// NewSystemUtilitySettingQuery200ResponseWithDefaults instantiates a new SystemUtilitySettingQuery200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemUtilitySettingQuery200ResponseWithDefaults() *SystemUtilitySettingQuery200Response {
	this := SystemUtilitySettingQuery200Response{}
	return &this
}

// GetItems returns the Items field value
func (o *SystemUtilitySettingQuery200Response) GetItems() []SystemOptimizationModel {
	if o == nil {
		var ret []SystemOptimizationModel
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *SystemUtilitySettingQuery200Response) GetItemsOk() ([]SystemOptimizationModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *SystemUtilitySettingQuery200Response) SetItems(v []SystemOptimizationModel) {
	o.Items = v
}

func (o SystemUtilitySettingQuery200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemUtilitySettingQuery200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableSystemUtilitySettingQuery200Response struct {
	value *SystemUtilitySettingQuery200Response
	isSet bool
}

func (v NullableSystemUtilitySettingQuery200Response) Get() *SystemUtilitySettingQuery200Response {
	return v.value
}

func (v *NullableSystemUtilitySettingQuery200Response) Set(val *SystemUtilitySettingQuery200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemUtilitySettingQuery200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemUtilitySettingQuery200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemUtilitySettingQuery200Response(val *SystemUtilitySettingQuery200Response) *NullableSystemUtilitySettingQuery200Response {
	return &NullableSystemUtilitySettingQuery200Response{value: val, isSet: true}
}

func (v NullableSystemUtilitySettingQuery200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemUtilitySettingQuery200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
