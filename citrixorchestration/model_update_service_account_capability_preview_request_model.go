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

// checks if the UpdateServiceAccountCapabilityPreviewRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateServiceAccountCapabilityPreviewRequestModel{}

// UpdateServiceAccountCapabilityPreviewRequestModel Request model for updating service account capability preview
type UpdateServiceAccountCapabilityPreviewRequestModel struct {
	// Gets or sets capabilities for the service account
	Capabilities []string `json:"Capabilities,omitempty"`
}

// NewUpdateServiceAccountCapabilityPreviewRequestModel instantiates a new UpdateServiceAccountCapabilityPreviewRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateServiceAccountCapabilityPreviewRequestModel() *UpdateServiceAccountCapabilityPreviewRequestModel {
	this := UpdateServiceAccountCapabilityPreviewRequestModel{}
	return &this
}

// NewUpdateServiceAccountCapabilityPreviewRequestModelWithDefaults instantiates a new UpdateServiceAccountCapabilityPreviewRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateServiceAccountCapabilityPreviewRequestModelWithDefaults() *UpdateServiceAccountCapabilityPreviewRequestModel {
	this := UpdateServiceAccountCapabilityPreviewRequestModel{}
	return &this
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateServiceAccountCapabilityPreviewRequestModel) GetCapabilities() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateServiceAccountCapabilityPreviewRequestModel) GetCapabilitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *UpdateServiceAccountCapabilityPreviewRequestModel) HasCapabilities() bool {
	if o != nil && IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []string and assigns it to the Capabilities field.
func (o *UpdateServiceAccountCapabilityPreviewRequestModel) SetCapabilities(v []string) {
	o.Capabilities = v
}

func (o UpdateServiceAccountCapabilityPreviewRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateServiceAccountCapabilityPreviewRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Capabilities != nil {
		toSerialize["Capabilities"] = o.Capabilities
	}
	return toSerialize, nil
}

type NullableUpdateServiceAccountCapabilityPreviewRequestModel struct {
	value *UpdateServiceAccountCapabilityPreviewRequestModel
	isSet bool
}

func (v NullableUpdateServiceAccountCapabilityPreviewRequestModel) Get() *UpdateServiceAccountCapabilityPreviewRequestModel {
	return v.value
}

func (v *NullableUpdateServiceAccountCapabilityPreviewRequestModel) Set(val *UpdateServiceAccountCapabilityPreviewRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateServiceAccountCapabilityPreviewRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateServiceAccountCapabilityPreviewRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateServiceAccountCapabilityPreviewRequestModel(val *UpdateServiceAccountCapabilityPreviewRequestModel) *NullableUpdateServiceAccountCapabilityPreviewRequestModel {
	return &NullableUpdateServiceAccountCapabilityPreviewRequestModel{value: val, isSet: true}
}

func (v NullableUpdateServiceAccountCapabilityPreviewRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateServiceAccountCapabilityPreviewRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

