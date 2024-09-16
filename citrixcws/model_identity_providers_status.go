/*
Citrix.CloudServices.Cws.Api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixcws

import (
	"encoding/json"
)

// checks if the IdentityProvidersStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityProvidersStatus{}

// IdentityProvidersStatus struct for IdentityProvidersStatus
type IdentityProvidersStatus struct {
	IdentityProviderEnabled map[string]bool `json:"identityProviderEnabled,omitempty"`
}

// NewIdentityProvidersStatus instantiates a new IdentityProvidersStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProvidersStatus() *IdentityProvidersStatus {
	this := IdentityProvidersStatus{}
	return &this
}

// NewIdentityProvidersStatusWithDefaults instantiates a new IdentityProvidersStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProvidersStatusWithDefaults() *IdentityProvidersStatus {
	this := IdentityProvidersStatus{}
	return &this
}

// GetIdentityProviderEnabled returns the IdentityProviderEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityProvidersStatus) GetIdentityProviderEnabled() map[string]bool {
	if o == nil {
		var ret map[string]bool
		return ret
	}
	return o.IdentityProviderEnabled
}

// GetIdentityProviderEnabledOk returns a tuple with the IdentityProviderEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityProvidersStatus) GetIdentityProviderEnabledOk() (*map[string]bool, bool) {
	if o == nil || IsNil(o.IdentityProviderEnabled) {
		return nil, false
	}
	return &o.IdentityProviderEnabled, true
}

// HasIdentityProviderEnabled returns a boolean if a field has been set.
func (o *IdentityProvidersStatus) HasIdentityProviderEnabled() bool {
	if o != nil && IsNil(o.IdentityProviderEnabled) {
		return true
	}

	return false
}

// SetIdentityProviderEnabled gets a reference to the given map[string]bool and assigns it to the IdentityProviderEnabled field.
func (o *IdentityProvidersStatus) SetIdentityProviderEnabled(v map[string]bool) {
	o.IdentityProviderEnabled = v
}

func (o IdentityProvidersStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityProvidersStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IdentityProviderEnabled != nil {
		toSerialize["identityProviderEnabled"] = o.IdentityProviderEnabled
	}
	return toSerialize, nil
}

type NullableIdentityProvidersStatus struct {
	value *IdentityProvidersStatus
	isSet bool
}

func (v NullableIdentityProvidersStatus) Get() *IdentityProvidersStatus {
	return v.value
}

func (v *NullableIdentityProvidersStatus) Set(val *IdentityProvidersStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProvidersStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProvidersStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProvidersStatus(val *IdentityProvidersStatus) *NullableIdentityProvidersStatus {
	return &NullableIdentityProvidersStatus{value: val, isSet: true}
}

func (v NullableIdentityProvidersStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProvidersStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

