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

// checks if the PolicyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyRequest{}

// PolicyRequest Data for policy request.
type PolicyRequest struct {
	// Policy name.
	Name NullableString `json:"name,omitempty"`
	// Is policy enabled.
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Policy description.
	Description NullableString `json:"description,omitempty"`
	// Policy settings.
	Settings []SettingRequest `json:"settings,omitempty"`
	// Policy filters.
	Filters []FilterRequest `json:"filters,omitempty"`
}

// NewPolicyRequest instantiates a new PolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyRequest() *PolicyRequest {
	this := PolicyRequest{}
	return &this
}

// NewPolicyRequestWithDefaults instantiates a new PolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyRequestWithDefaults() *PolicyRequest {
	this := PolicyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyRequest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PolicyRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PolicyRequest) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *PolicyRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PolicyRequest) UnsetName() {
	o.Name.Unset()
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *PolicyRequest) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRequest) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *PolicyRequest) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *PolicyRequest) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PolicyRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PolicyRequest) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PolicyRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PolicyRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetSettings returns the Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyRequest) GetSettings() []SettingRequest {
	if o == nil {
		var ret []SettingRequest
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyRequest) GetSettingsOk() ([]SettingRequest, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *PolicyRequest) HasSettings() bool {
	if o != nil && IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given []SettingRequest and assigns it to the Settings field.
func (o *PolicyRequest) SetSettings(v []SettingRequest) {
	o.Settings = v
}

// GetFilters returns the Filters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyRequest) GetFilters() []FilterRequest {
	if o == nil {
		var ret []FilterRequest
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyRequest) GetFiltersOk() ([]FilterRequest, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *PolicyRequest) HasFilters() bool {
	if o != nil && IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []FilterRequest and assigns it to the Filters field.
func (o *PolicyRequest) SetFilters(v []FilterRequest) {
	o.Filters = v
}

func (o PolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	return toSerialize, nil
}

type NullablePolicyRequest struct {
	value *PolicyRequest
	isSet bool
}

func (v NullablePolicyRequest) Get() *PolicyRequest {
	return v.value
}

func (v *NullablePolicyRequest) Set(val *PolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyRequest(val *PolicyRequest) *NullablePolicyRequest {
	return &NullablePolicyRequest{value: val, isSet: true}
}

func (v NullablePolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
