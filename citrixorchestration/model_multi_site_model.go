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

// checks if the MultiSiteModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiSiteModel{}

// MultiSiteModel Site information for Multi sites
type MultiSiteModel struct {
	// Site Name
	Name NullableString `json:"Name,omitempty"`
	// The site's friendly name
	DisplayName string `json:"DisplayName"`
	// The Orchestration servers belongs to the site
	DeliveryControllers []string `json:"DeliveryControllers,omitempty"`
	// Indicate if it is the default site.
	Default *bool `json:"Default,omitempty"`
}

// NewMultiSiteModel instantiates a new MultiSiteModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiSiteModel(displayName string) *MultiSiteModel {
	this := MultiSiteModel{}
	this.DisplayName = displayName
	var default_ bool = false
	this.Default = &default_
	return &this
}

// NewMultiSiteModelWithDefaults instantiates a new MultiSiteModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiSiteModelWithDefaults() *MultiSiteModel {
	this := MultiSiteModel{}
	var default_ bool = false
	this.Default = &default_
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MultiSiteModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MultiSiteModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MultiSiteModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MultiSiteModel) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *MultiSiteModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MultiSiteModel) UnsetName() {
	o.Name.Unset()
}

// GetDisplayName returns the DisplayName field value
func (o *MultiSiteModel) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *MultiSiteModel) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *MultiSiteModel) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetDeliveryControllers returns the DeliveryControllers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MultiSiteModel) GetDeliveryControllers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DeliveryControllers
}

// GetDeliveryControllersOk returns a tuple with the DeliveryControllers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MultiSiteModel) GetDeliveryControllersOk() ([]string, bool) {
	if o == nil || IsNil(o.DeliveryControllers) {
		return nil, false
	}
	return o.DeliveryControllers, true
}

// HasDeliveryControllers returns a boolean if a field has been set.
func (o *MultiSiteModel) HasDeliveryControllers() bool {
	if o != nil && IsNil(o.DeliveryControllers) {
		return true
	}

	return false
}

// SetDeliveryControllers gets a reference to the given []string and assigns it to the DeliveryControllers field.
func (o *MultiSiteModel) SetDeliveryControllers(v []string) {
	o.DeliveryControllers = v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *MultiSiteModel) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiSiteModel) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *MultiSiteModel) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *MultiSiteModel) SetDefault(v bool) {
	o.Default = &v
}

func (o MultiSiteModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiSiteModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	toSerialize["DisplayName"] = o.DisplayName
	if o.DeliveryControllers != nil {
		toSerialize["DeliveryControllers"] = o.DeliveryControllers
	}
	if !IsNil(o.Default) {
		toSerialize["Default"] = o.Default
	}
	return toSerialize, nil
}

type NullableMultiSiteModel struct {
	value *MultiSiteModel
	isSet bool
}

func (v NullableMultiSiteModel) Get() *MultiSiteModel {
	return v.value
}

func (v *NullableMultiSiteModel) Set(val *MultiSiteModel) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiSiteModel) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiSiteModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiSiteModel(val *MultiSiteModel) *NullableMultiSiteModel {
	return &NullableMultiSiteModel{value: val, isSet: true}
}

func (v NullableMultiSiteModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiSiteModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
