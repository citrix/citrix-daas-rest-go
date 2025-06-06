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

// checks if the ProvisioningSchemeSearchFilterRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningSchemeSearchFilterRequestModel{}

// ProvisioningSchemeSearchFilterRequestModel Advanced search filter for provisioning scheme.
type ProvisioningSchemeSearchFilterRequestModel struct {
	Property ProvisioningSchemeSearchProperty `json:"Property"`
	// Value to match.
	Value    NullableString `json:"Value,omitempty"`
	Operator SearchOperator `json:"Operator"`
}

// NewProvisioningSchemeSearchFilterRequestModel instantiates a new ProvisioningSchemeSearchFilterRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningSchemeSearchFilterRequestModel(property ProvisioningSchemeSearchProperty, operator SearchOperator) *ProvisioningSchemeSearchFilterRequestModel {
	this := ProvisioningSchemeSearchFilterRequestModel{}
	this.Property = property
	this.Operator = operator
	return &this
}

// NewProvisioningSchemeSearchFilterRequestModelWithDefaults instantiates a new ProvisioningSchemeSearchFilterRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningSchemeSearchFilterRequestModelWithDefaults() *ProvisioningSchemeSearchFilterRequestModel {
	this := ProvisioningSchemeSearchFilterRequestModel{}
	return &this
}

// GetProperty returns the Property field value
func (o *ProvisioningSchemeSearchFilterRequestModel) GetProperty() ProvisioningSchemeSearchProperty {
	if o == nil {
		var ret ProvisioningSchemeSearchProperty
		return ret
	}

	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value
// and a boolean to check if the value has been set.
func (o *ProvisioningSchemeSearchFilterRequestModel) GetPropertyOk() (*ProvisioningSchemeSearchProperty, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Property, true
}

// SetProperty sets field value
func (o *ProvisioningSchemeSearchFilterRequestModel) SetProperty(v ProvisioningSchemeSearchProperty) {
	o.Property = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningSchemeSearchFilterRequestModel) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningSchemeSearchFilterRequestModel) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ProvisioningSchemeSearchFilterRequestModel) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *ProvisioningSchemeSearchFilterRequestModel) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *ProvisioningSchemeSearchFilterRequestModel) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ProvisioningSchemeSearchFilterRequestModel) UnsetValue() {
	o.Value.Unset()
}

// GetOperator returns the Operator field value
func (o *ProvisioningSchemeSearchFilterRequestModel) GetOperator() SearchOperator {
	if o == nil {
		var ret SearchOperator
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *ProvisioningSchemeSearchFilterRequestModel) GetOperatorOk() (*SearchOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *ProvisioningSchemeSearchFilterRequestModel) SetOperator(v SearchOperator) {
	o.Operator = v
}

func (o ProvisioningSchemeSearchFilterRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningSchemeSearchFilterRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Property"] = o.Property
	if o.Value.IsSet() {
		toSerialize["Value"] = o.Value.Get()
	}
	toSerialize["Operator"] = o.Operator
	return toSerialize, nil
}

type NullableProvisioningSchemeSearchFilterRequestModel struct {
	value *ProvisioningSchemeSearchFilterRequestModel
	isSet bool
}

func (v NullableProvisioningSchemeSearchFilterRequestModel) Get() *ProvisioningSchemeSearchFilterRequestModel {
	return v.value
}

func (v *NullableProvisioningSchemeSearchFilterRequestModel) Set(val *ProvisioningSchemeSearchFilterRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningSchemeSearchFilterRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningSchemeSearchFilterRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningSchemeSearchFilterRequestModel(val *ProvisioningSchemeSearchFilterRequestModel) *NullableProvisioningSchemeSearchFilterRequestModel {
	return &NullableProvisioningSchemeSearchFilterRequestModel{value: val, isSet: true}
}

func (v NullableProvisioningSchemeSearchFilterRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningSchemeSearchFilterRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
