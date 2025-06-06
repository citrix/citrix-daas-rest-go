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

// checks if the DeliveryControllerDeleteScriptModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryControllerDeleteScriptModel{}

// DeliveryControllerDeleteScriptModel Response object for DeliveryControllerDeleteScript.
type DeliveryControllerDeleteScriptModel struct {
	// The filename of the text file.
	FileName NullableString `json:"FileName,omitempty"`
	// The contents of the text file.
	Content NullableString `json:"Content,omitempty"`
}

// NewDeliveryControllerDeleteScriptModel instantiates a new DeliveryControllerDeleteScriptModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryControllerDeleteScriptModel() *DeliveryControllerDeleteScriptModel {
	this := DeliveryControllerDeleteScriptModel{}
	return &this
}

// NewDeliveryControllerDeleteScriptModelWithDefaults instantiates a new DeliveryControllerDeleteScriptModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryControllerDeleteScriptModelWithDefaults() *DeliveryControllerDeleteScriptModel {
	this := DeliveryControllerDeleteScriptModel{}
	return &this
}

// GetFileName returns the FileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeliveryControllerDeleteScriptModel) GetFileName() string {
	if o == nil || IsNil(o.FileName.Get()) {
		var ret string
		return ret
	}
	return *o.FileName.Get()
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryControllerDeleteScriptModel) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileName.Get(), o.FileName.IsSet()
}

// HasFileName returns a boolean if a field has been set.
func (o *DeliveryControllerDeleteScriptModel) HasFileName() bool {
	if o != nil && o.FileName.IsSet() {
		return true
	}

	return false
}

// SetFileName gets a reference to the given NullableString and assigns it to the FileName field.
func (o *DeliveryControllerDeleteScriptModel) SetFileName(v string) {
	o.FileName.Set(&v)
}

// SetFileNameNil sets the value for FileName to be an explicit nil
func (o *DeliveryControllerDeleteScriptModel) SetFileNameNil() {
	o.FileName.Set(nil)
}

// UnsetFileName ensures that no value is present for FileName, not even an explicit nil
func (o *DeliveryControllerDeleteScriptModel) UnsetFileName() {
	o.FileName.Unset()
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeliveryControllerDeleteScriptModel) GetContent() string {
	if o == nil || IsNil(o.Content.Get()) {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryControllerDeleteScriptModel) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *DeliveryControllerDeleteScriptModel) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *DeliveryControllerDeleteScriptModel) SetContent(v string) {
	o.Content.Set(&v)
}

// SetContentNil sets the value for Content to be an explicit nil
func (o *DeliveryControllerDeleteScriptModel) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *DeliveryControllerDeleteScriptModel) UnsetContent() {
	o.Content.Unset()
}

func (o DeliveryControllerDeleteScriptModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryControllerDeleteScriptModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FileName.IsSet() {
		toSerialize["FileName"] = o.FileName.Get()
	}
	if o.Content.IsSet() {
		toSerialize["Content"] = o.Content.Get()
	}
	return toSerialize, nil
}

type NullableDeliveryControllerDeleteScriptModel struct {
	value *DeliveryControllerDeleteScriptModel
	isSet bool
}

func (v NullableDeliveryControllerDeleteScriptModel) Get() *DeliveryControllerDeleteScriptModel {
	return v.value
}

func (v *NullableDeliveryControllerDeleteScriptModel) Set(val *DeliveryControllerDeleteScriptModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryControllerDeleteScriptModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryControllerDeleteScriptModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryControllerDeleteScriptModel(val *DeliveryControllerDeleteScriptModel) *NullableDeliveryControllerDeleteScriptModel {
	return &NullableDeliveryControllerDeleteScriptModel{value: val, isSet: true}
}

func (v NullableDeliveryControllerDeleteScriptModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryControllerDeleteScriptModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
