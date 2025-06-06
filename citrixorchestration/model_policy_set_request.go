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

// checks if the PolicySetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicySetRequest{}

// PolicySetRequest Policy set request.
type PolicySetRequest struct {
	// The policy set type.
	PolicySetType NullableString `json:"policySetType,omitempty"`
	// Name of the policy set.
	Name NullableString `json:"name,omitempty"`
	// Policy set description.
	Description NullableString `json:"description,omitempty"`
	// Delegated administration scopes for the policy set.
	Scopes []string `json:"scopes,omitempty"`
}

// NewPolicySetRequest instantiates a new PolicySetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicySetRequest() *PolicySetRequest {
	this := PolicySetRequest{}
	return &this
}

// NewPolicySetRequestWithDefaults instantiates a new PolicySetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicySetRequestWithDefaults() *PolicySetRequest {
	this := PolicySetRequest{}
	return &this
}

// GetPolicySetType returns the PolicySetType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicySetRequest) GetPolicySetType() string {
	if o == nil || IsNil(o.PolicySetType.Get()) {
		var ret string
		return ret
	}
	return *o.PolicySetType.Get()
}

// GetPolicySetTypeOk returns a tuple with the PolicySetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicySetRequest) GetPolicySetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PolicySetType.Get(), o.PolicySetType.IsSet()
}

// HasPolicySetType returns a boolean if a field has been set.
func (o *PolicySetRequest) HasPolicySetType() bool {
	if o != nil && o.PolicySetType.IsSet() {
		return true
	}

	return false
}

// SetPolicySetType gets a reference to the given NullableString and assigns it to the PolicySetType field.
func (o *PolicySetRequest) SetPolicySetType(v string) {
	o.PolicySetType.Set(&v)
}

// SetPolicySetTypeNil sets the value for PolicySetType to be an explicit nil
func (o *PolicySetRequest) SetPolicySetTypeNil() {
	o.PolicySetType.Set(nil)
}

// UnsetPolicySetType ensures that no value is present for PolicySetType, not even an explicit nil
func (o *PolicySetRequest) UnsetPolicySetType() {
	o.PolicySetType.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicySetRequest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicySetRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PolicySetRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PolicySetRequest) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *PolicySetRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PolicySetRequest) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicySetRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicySetRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PolicySetRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PolicySetRequest) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PolicySetRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PolicySetRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicySetRequest) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicySetRequest) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *PolicySetRequest) HasScopes() bool {
	if o != nil && IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *PolicySetRequest) SetScopes(v []string) {
	o.Scopes = v
}

func (o PolicySetRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicySetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PolicySetType.IsSet() {
		toSerialize["policySetType"] = o.PolicySetType.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	return toSerialize, nil
}

type NullablePolicySetRequest struct {
	value *PolicySetRequest
	isSet bool
}

func (v NullablePolicySetRequest) Get() *PolicySetRequest {
	return v.value
}

func (v *NullablePolicySetRequest) Set(val *PolicySetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicySetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicySetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicySetRequest(val *PolicySetRequest) *NullablePolicySetRequest {
	return &NullablePolicySetRequest{value: val, isSet: true}
}

func (v NullablePolicySetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicySetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
