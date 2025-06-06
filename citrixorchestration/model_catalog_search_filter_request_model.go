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

// checks if the CatalogSearchFilterRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogSearchFilterRequestModel{}

// CatalogSearchFilterRequestModel Advanced search filter for machine catalogs.
type CatalogSearchFilterRequestModel struct {
	Property CatalogSearchProperty `json:"Property"`
	// Value to match.
	Value    NullableString `json:"Value,omitempty"`
	Operator SearchOperator `json:"Operator"`
}

// NewCatalogSearchFilterRequestModel instantiates a new CatalogSearchFilterRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogSearchFilterRequestModel(property CatalogSearchProperty, operator SearchOperator) *CatalogSearchFilterRequestModel {
	this := CatalogSearchFilterRequestModel{}
	this.Property = property
	this.Operator = operator
	return &this
}

// NewCatalogSearchFilterRequestModelWithDefaults instantiates a new CatalogSearchFilterRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogSearchFilterRequestModelWithDefaults() *CatalogSearchFilterRequestModel {
	this := CatalogSearchFilterRequestModel{}
	return &this
}

// GetProperty returns the Property field value
func (o *CatalogSearchFilterRequestModel) GetProperty() CatalogSearchProperty {
	if o == nil {
		var ret CatalogSearchProperty
		return ret
	}

	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value
// and a boolean to check if the value has been set.
func (o *CatalogSearchFilterRequestModel) GetPropertyOk() (*CatalogSearchProperty, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Property, true
}

// SetProperty sets field value
func (o *CatalogSearchFilterRequestModel) SetProperty(v CatalogSearchProperty) {
	o.Property = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogSearchFilterRequestModel) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogSearchFilterRequestModel) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *CatalogSearchFilterRequestModel) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *CatalogSearchFilterRequestModel) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *CatalogSearchFilterRequestModel) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *CatalogSearchFilterRequestModel) UnsetValue() {
	o.Value.Unset()
}

// GetOperator returns the Operator field value
func (o *CatalogSearchFilterRequestModel) GetOperator() SearchOperator {
	if o == nil {
		var ret SearchOperator
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *CatalogSearchFilterRequestModel) GetOperatorOk() (*SearchOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *CatalogSearchFilterRequestModel) SetOperator(v SearchOperator) {
	o.Operator = v
}

func (o CatalogSearchFilterRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogSearchFilterRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Property"] = o.Property
	if o.Value.IsSet() {
		toSerialize["Value"] = o.Value.Get()
	}
	toSerialize["Operator"] = o.Operator
	return toSerialize, nil
}

type NullableCatalogSearchFilterRequestModel struct {
	value *CatalogSearchFilterRequestModel
	isSet bool
}

func (v NullableCatalogSearchFilterRequestModel) Get() *CatalogSearchFilterRequestModel {
	return v.value
}

func (v *NullableCatalogSearchFilterRequestModel) Set(val *CatalogSearchFilterRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogSearchFilterRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogSearchFilterRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogSearchFilterRequestModel(val *CatalogSearchFilterRequestModel) *NullableCatalogSearchFilterRequestModel {
	return &NullableCatalogSearchFilterRequestModel{value: val, isSet: true}
}

func (v NullableCatalogSearchFilterRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogSearchFilterRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
