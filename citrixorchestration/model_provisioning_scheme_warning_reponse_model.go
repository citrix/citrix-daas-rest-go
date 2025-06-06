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

// checks if the ProvisioningSchemeWarningReponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningSchemeWarningReponseModel{}

// ProvisioningSchemeWarningReponseModel struct for ProvisioningSchemeWarningReponseModel
type ProvisioningSchemeWarningReponseModel struct {
	Type *ProvSchemeWarningType `json:"Type,omitempty"`
	// Message associated with warning
	Message NullableString `json:"Message,omitempty"`
}

// NewProvisioningSchemeWarningReponseModel instantiates a new ProvisioningSchemeWarningReponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningSchemeWarningReponseModel() *ProvisioningSchemeWarningReponseModel {
	this := ProvisioningSchemeWarningReponseModel{}
	return &this
}

// NewProvisioningSchemeWarningReponseModelWithDefaults instantiates a new ProvisioningSchemeWarningReponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningSchemeWarningReponseModelWithDefaults() *ProvisioningSchemeWarningReponseModel {
	this := ProvisioningSchemeWarningReponseModel{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProvisioningSchemeWarningReponseModel) GetType() ProvSchemeWarningType {
	if o == nil || IsNil(o.Type) {
		var ret ProvSchemeWarningType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningSchemeWarningReponseModel) GetTypeOk() (*ProvSchemeWarningType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProvisioningSchemeWarningReponseModel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ProvSchemeWarningType and assigns it to the Type field.
func (o *ProvisioningSchemeWarningReponseModel) SetType(v ProvSchemeWarningType) {
	o.Type = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningSchemeWarningReponseModel) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningSchemeWarningReponseModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *ProvisioningSchemeWarningReponseModel) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *ProvisioningSchemeWarningReponseModel) SetMessage(v string) {
	o.Message.Set(&v)
}

// SetMessageNil sets the value for Message to be an explicit nil
func (o *ProvisioningSchemeWarningReponseModel) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *ProvisioningSchemeWarningReponseModel) UnsetMessage() {
	o.Message.Unset()
}

func (o ProvisioningSchemeWarningReponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningSchemeWarningReponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if o.Message.IsSet() {
		toSerialize["Message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableProvisioningSchemeWarningReponseModel struct {
	value *ProvisioningSchemeWarningReponseModel
	isSet bool
}

func (v NullableProvisioningSchemeWarningReponseModel) Get() *ProvisioningSchemeWarningReponseModel {
	return v.value
}

func (v *NullableProvisioningSchemeWarningReponseModel) Set(val *ProvisioningSchemeWarningReponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningSchemeWarningReponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningSchemeWarningReponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningSchemeWarningReponseModel(val *ProvisioningSchemeWarningReponseModel) *NullableProvisioningSchemeWarningReponseModel {
	return &NullableProvisioningSchemeWarningReponseModel{value: val, isSet: true}
}

func (v NullableProvisioningSchemeWarningReponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningSchemeWarningReponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
