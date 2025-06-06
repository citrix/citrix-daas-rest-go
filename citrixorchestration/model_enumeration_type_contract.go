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

// checks if the EnumerationTypeContract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnumerationTypeContract{}

// EnumerationTypeContract Enumeration type
type EnumerationTypeContract struct {
	// Name of the type.
	TypeName NullableString `json:"TypeName,omitempty"`
	// Members of the type.
	Members []EnumerationMemberContract `json:"Members,omitempty"`
}

// NewEnumerationTypeContract instantiates a new EnumerationTypeContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnumerationTypeContract() *EnumerationTypeContract {
	this := EnumerationTypeContract{}
	return &this
}

// NewEnumerationTypeContractWithDefaults instantiates a new EnumerationTypeContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnumerationTypeContractWithDefaults() *EnumerationTypeContract {
	this := EnumerationTypeContract{}
	return &this
}

// GetTypeName returns the TypeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnumerationTypeContract) GetTypeName() string {
	if o == nil || IsNil(o.TypeName.Get()) {
		var ret string
		return ret
	}
	return *o.TypeName.Get()
}

// GetTypeNameOk returns a tuple with the TypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnumerationTypeContract) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TypeName.Get(), o.TypeName.IsSet()
}

// HasTypeName returns a boolean if a field has been set.
func (o *EnumerationTypeContract) HasTypeName() bool {
	if o != nil && o.TypeName.IsSet() {
		return true
	}

	return false
}

// SetTypeName gets a reference to the given NullableString and assigns it to the TypeName field.
func (o *EnumerationTypeContract) SetTypeName(v string) {
	o.TypeName.Set(&v)
}

// SetTypeNameNil sets the value for TypeName to be an explicit nil
func (o *EnumerationTypeContract) SetTypeNameNil() {
	o.TypeName.Set(nil)
}

// UnsetTypeName ensures that no value is present for TypeName, not even an explicit nil
func (o *EnumerationTypeContract) UnsetTypeName() {
	o.TypeName.Unset()
}

// GetMembers returns the Members field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnumerationTypeContract) GetMembers() []EnumerationMemberContract {
	if o == nil {
		var ret []EnumerationMemberContract
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnumerationTypeContract) GetMembersOk() ([]EnumerationMemberContract, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *EnumerationTypeContract) HasMembers() bool {
	if o != nil && IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []EnumerationMemberContract and assigns it to the Members field.
func (o *EnumerationTypeContract) SetMembers(v []EnumerationMemberContract) {
	o.Members = v
}

func (o EnumerationTypeContract) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnumerationTypeContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TypeName.IsSet() {
		toSerialize["TypeName"] = o.TypeName.Get()
	}
	if o.Members != nil {
		toSerialize["Members"] = o.Members
	}
	return toSerialize, nil
}

type NullableEnumerationTypeContract struct {
	value *EnumerationTypeContract
	isSet bool
}

func (v NullableEnumerationTypeContract) Get() *EnumerationTypeContract {
	return v.value
}

func (v *NullableEnumerationTypeContract) Set(val *EnumerationTypeContract) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumerationTypeContract) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumerationTypeContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumerationTypeContract(val *EnumerationTypeContract) *NullableEnumerationTypeContract {
	return &NullableEnumerationTypeContract{value: val, isSet: true}
}

func (v NullableEnumerationTypeContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumerationTypeContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
