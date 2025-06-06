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

// checks if the SmartAccessTagRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmartAccessTagRequestModel{}

// SmartAccessTagRequestModel Request object specifying properties of a NetScaler SmartAccess Tag.
type SmartAccessTagRequestModel struct {
	// SmartAccess Farm name.
	Farm string `json:"Farm"`
	// SmartAccess filter name.
	Filter string `json:"Filter"`
}

// NewSmartAccessTagRequestModel instantiates a new SmartAccessTagRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartAccessTagRequestModel(farm string, filter string) *SmartAccessTagRequestModel {
	this := SmartAccessTagRequestModel{}
	this.Farm = farm
	this.Filter = filter
	return &this
}

// NewSmartAccessTagRequestModelWithDefaults instantiates a new SmartAccessTagRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartAccessTagRequestModelWithDefaults() *SmartAccessTagRequestModel {
	this := SmartAccessTagRequestModel{}
	return &this
}

// GetFarm returns the Farm field value
func (o *SmartAccessTagRequestModel) GetFarm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Farm
}

// GetFarmOk returns a tuple with the Farm field value
// and a boolean to check if the value has been set.
func (o *SmartAccessTagRequestModel) GetFarmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Farm, true
}

// SetFarm sets field value
func (o *SmartAccessTagRequestModel) SetFarm(v string) {
	o.Farm = v
}

// GetFilter returns the Filter field value
func (o *SmartAccessTagRequestModel) GetFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *SmartAccessTagRequestModel) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *SmartAccessTagRequestModel) SetFilter(v string) {
	o.Filter = v
}

func (o SmartAccessTagRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmartAccessTagRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Farm"] = o.Farm
	toSerialize["Filter"] = o.Filter
	return toSerialize, nil
}

type NullableSmartAccessTagRequestModel struct {
	value *SmartAccessTagRequestModel
	isSet bool
}

func (v NullableSmartAccessTagRequestModel) Get() *SmartAccessTagRequestModel {
	return v.value
}

func (v *NullableSmartAccessTagRequestModel) Set(val *SmartAccessTagRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartAccessTagRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartAccessTagRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartAccessTagRequestModel(val *SmartAccessTagRequestModel) *NullableSmartAccessTagRequestModel {
	return &NullableSmartAccessTagRequestModel{value: val, isSet: true}
}

func (v NullableSmartAccessTagRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartAccessTagRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
