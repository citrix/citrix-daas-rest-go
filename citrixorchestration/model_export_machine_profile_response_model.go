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

// checks if the ExportMachineProfileResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportMachineProfileResponseModel{}

// ExportMachineProfileResponseModel ExportMachineProfileResponseModel class.
type ExportMachineProfileResponseModel struct {
	// The content of the exported machine profile in the form of a JSON string.
	Content NullableString `json:"Content,omitempty"`
}

// NewExportMachineProfileResponseModel instantiates a new ExportMachineProfileResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportMachineProfileResponseModel() *ExportMachineProfileResponseModel {
	this := ExportMachineProfileResponseModel{}
	return &this
}

// NewExportMachineProfileResponseModelWithDefaults instantiates a new ExportMachineProfileResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportMachineProfileResponseModelWithDefaults() *ExportMachineProfileResponseModel {
	this := ExportMachineProfileResponseModel{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportMachineProfileResponseModel) GetContent() string {
	if o == nil || IsNil(o.Content.Get()) {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportMachineProfileResponseModel) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *ExportMachineProfileResponseModel) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *ExportMachineProfileResponseModel) SetContent(v string) {
	o.Content.Set(&v)
}

// SetContentNil sets the value for Content to be an explicit nil
func (o *ExportMachineProfileResponseModel) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *ExportMachineProfileResponseModel) UnsetContent() {
	o.Content.Unset()
}

func (o ExportMachineProfileResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportMachineProfileResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Content.IsSet() {
		toSerialize["Content"] = o.Content.Get()
	}
	return toSerialize, nil
}

type NullableExportMachineProfileResponseModel struct {
	value *ExportMachineProfileResponseModel
	isSet bool
}

func (v NullableExportMachineProfileResponseModel) Get() *ExportMachineProfileResponseModel {
	return v.value
}

func (v *NullableExportMachineProfileResponseModel) Set(val *ExportMachineProfileResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableExportMachineProfileResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableExportMachineProfileResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportMachineProfileResponseModel(val *ExportMachineProfileResponseModel) *NullableExportMachineProfileResponseModel {
	return &NullableExportMachineProfileResponseModel{value: val, isSet: true}
}

func (v NullableExportMachineProfileResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportMachineProfileResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
