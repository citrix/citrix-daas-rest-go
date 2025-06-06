/*
Quick Create Service for DaaS Quick Deploy with AWS WorkSpaces Core - Tech Preview

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: DaaSTechPreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickcreate

import (
	"encoding/json"
)

// checks if the AccountWarning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountWarning{}

// AccountWarning struct for AccountWarning
type AccountWarning struct {
	// The warning message
	Warning NullableString `json:"warning,omitempty"`
	// Indicates if the warning should be treated as an error
	IsError *bool `json:"isError,omitempty"`
	// Indicates if the warning message can be cleared out by the user
	IsDismissible *bool `json:"isDismissible,omitempty"`
}

// NewAccountWarning instantiates a new AccountWarning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountWarning() *AccountWarning {
	this := AccountWarning{}
	return &this
}

// NewAccountWarningWithDefaults instantiates a new AccountWarning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWarningWithDefaults() *AccountWarning {
	this := AccountWarning{}
	return &this
}

// GetWarning returns the Warning field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountWarning) GetWarning() string {
	if o == nil || IsNil(o.Warning.Get()) {
		var ret string
		return ret
	}
	return *o.Warning.Get()
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountWarning) GetWarningOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Warning.Get(), o.Warning.IsSet()
}

// HasWarning returns a boolean if a field has been set.
func (o *AccountWarning) HasWarning() bool {
	if o != nil && o.Warning.IsSet() {
		return true
	}

	return false
}

// SetWarning gets a reference to the given NullableString and assigns it to the Warning field.
func (o *AccountWarning) SetWarning(v string) {
	o.Warning.Set(&v)
}

// SetWarningNil sets the value for Warning to be an explicit nil
func (o *AccountWarning) SetWarningNil() {
	o.Warning.Set(nil)
}

// UnsetWarning ensures that no value is present for Warning, not even an explicit nil
func (o *AccountWarning) UnsetWarning() {
	o.Warning.Unset()
}

// GetIsError returns the IsError field value if set, zero value otherwise.
func (o *AccountWarning) GetIsError() bool {
	if o == nil || IsNil(o.IsError) {
		var ret bool
		return ret
	}
	return *o.IsError
}

// GetIsErrorOk returns a tuple with the IsError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountWarning) GetIsErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.IsError) {
		return nil, false
	}
	return o.IsError, true
}

// HasIsError returns a boolean if a field has been set.
func (o *AccountWarning) HasIsError() bool {
	if o != nil && !IsNil(o.IsError) {
		return true
	}

	return false
}

// SetIsError gets a reference to the given bool and assigns it to the IsError field.
func (o *AccountWarning) SetIsError(v bool) {
	o.IsError = &v
}

// GetIsDismissible returns the IsDismissible field value if set, zero value otherwise.
func (o *AccountWarning) GetIsDismissible() bool {
	if o == nil || IsNil(o.IsDismissible) {
		var ret bool
		return ret
	}
	return *o.IsDismissible
}

// GetIsDismissibleOk returns a tuple with the IsDismissible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountWarning) GetIsDismissibleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDismissible) {
		return nil, false
	}
	return o.IsDismissible, true
}

// HasIsDismissible returns a boolean if a field has been set.
func (o *AccountWarning) HasIsDismissible() bool {
	if o != nil && !IsNil(o.IsDismissible) {
		return true
	}

	return false
}

// SetIsDismissible gets a reference to the given bool and assigns it to the IsDismissible field.
func (o *AccountWarning) SetIsDismissible(v bool) {
	o.IsDismissible = &v
}

func (o AccountWarning) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountWarning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Warning.IsSet() {
		toSerialize["warning"] = o.Warning.Get()
	}
	if !IsNil(o.IsError) {
		toSerialize["isError"] = o.IsError
	}
	if !IsNil(o.IsDismissible) {
		toSerialize["isDismissible"] = o.IsDismissible
	}
	return toSerialize, nil
}

type NullableAccountWarning struct {
	value *AccountWarning
	isSet bool
}

func (v NullableAccountWarning) Get() *AccountWarning {
	return v.value
}

func (v *NullableAccountWarning) Set(val *AccountWarning) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountWarning) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountWarning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountWarning(val *AccountWarning) *NullableAccountWarning {
	return &NullableAccountWarning{value: val, isSet: true}
}

func (v NullableAccountWarning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountWarning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
