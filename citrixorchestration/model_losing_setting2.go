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

// checks if the LosingSetting2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LosingSetting2{}

// LosingSetting2 Settings that are not applied.
type LosingSetting2 struct {
	// Policy GUID.
	PolicyGuid *string `json:"PolicyGuid,omitempty"`
	// The policy that has the setting.
	PolicyName NullableString `json:"PolicyName,omitempty"`
	// The setting that is not applied.
	SettingName NullableString `json:"SettingName,omitempty"`
}

// NewLosingSetting2 instantiates a new LosingSetting2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLosingSetting2() *LosingSetting2 {
	this := LosingSetting2{}
	return &this
}

// NewLosingSetting2WithDefaults instantiates a new LosingSetting2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLosingSetting2WithDefaults() *LosingSetting2 {
	this := LosingSetting2{}
	return &this
}

// GetPolicyGuid returns the PolicyGuid field value if set, zero value otherwise.
func (o *LosingSetting2) GetPolicyGuid() string {
	if o == nil || IsNil(o.PolicyGuid) {
		var ret string
		return ret
	}
	return *o.PolicyGuid
}

// GetPolicyGuidOk returns a tuple with the PolicyGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LosingSetting2) GetPolicyGuidOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyGuid) {
		return nil, false
	}
	return o.PolicyGuid, true
}

// HasPolicyGuid returns a boolean if a field has been set.
func (o *LosingSetting2) HasPolicyGuid() bool {
	if o != nil && !IsNil(o.PolicyGuid) {
		return true
	}

	return false
}

// SetPolicyGuid gets a reference to the given string and assigns it to the PolicyGuid field.
func (o *LosingSetting2) SetPolicyGuid(v string) {
	o.PolicyGuid = &v
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LosingSetting2) GetPolicyName() string {
	if o == nil || IsNil(o.PolicyName.Get()) {
		var ret string
		return ret
	}
	return *o.PolicyName.Get()
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LosingSetting2) GetPolicyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PolicyName.Get(), o.PolicyName.IsSet()
}

// HasPolicyName returns a boolean if a field has been set.
func (o *LosingSetting2) HasPolicyName() bool {
	if o != nil && o.PolicyName.IsSet() {
		return true
	}

	return false
}

// SetPolicyName gets a reference to the given NullableString and assigns it to the PolicyName field.
func (o *LosingSetting2) SetPolicyName(v string) {
	o.PolicyName.Set(&v)
}
// SetPolicyNameNil sets the value for PolicyName to be an explicit nil
func (o *LosingSetting2) SetPolicyNameNil() {
	o.PolicyName.Set(nil)
}

// UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
func (o *LosingSetting2) UnsetPolicyName() {
	o.PolicyName.Unset()
}

// GetSettingName returns the SettingName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LosingSetting2) GetSettingName() string {
	if o == nil || IsNil(o.SettingName.Get()) {
		var ret string
		return ret
	}
	return *o.SettingName.Get()
}

// GetSettingNameOk returns a tuple with the SettingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LosingSetting2) GetSettingNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SettingName.Get(), o.SettingName.IsSet()
}

// HasSettingName returns a boolean if a field has been set.
func (o *LosingSetting2) HasSettingName() bool {
	if o != nil && o.SettingName.IsSet() {
		return true
	}

	return false
}

// SetSettingName gets a reference to the given NullableString and assigns it to the SettingName field.
func (o *LosingSetting2) SetSettingName(v string) {
	o.SettingName.Set(&v)
}
// SetSettingNameNil sets the value for SettingName to be an explicit nil
func (o *LosingSetting2) SetSettingNameNil() {
	o.SettingName.Set(nil)
}

// UnsetSettingName ensures that no value is present for SettingName, not even an explicit nil
func (o *LosingSetting2) UnsetSettingName() {
	o.SettingName.Unset()
}

func (o LosingSetting2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LosingSetting2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PolicyGuid) {
		toSerialize["PolicyGuid"] = o.PolicyGuid
	}
	if o.PolicyName.IsSet() {
		toSerialize["PolicyName"] = o.PolicyName.Get()
	}
	if o.SettingName.IsSet() {
		toSerialize["SettingName"] = o.SettingName.Get()
	}
	return toSerialize, nil
}

type NullableLosingSetting2 struct {
	value *LosingSetting2
	isSet bool
}

func (v NullableLosingSetting2) Get() *LosingSetting2 {
	return v.value
}

func (v *NullableLosingSetting2) Set(val *LosingSetting2) {
	v.value = val
	v.isSet = true
}

func (v NullableLosingSetting2) IsSet() bool {
	return v.isSet
}

func (v *NullableLosingSetting2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLosingSetting2(val *LosingSetting2) *NullableLosingSetting2 {
	return &NullableLosingSetting2{value: val, isSet: true}
}

func (v NullableLosingSetting2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLosingSetting2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

