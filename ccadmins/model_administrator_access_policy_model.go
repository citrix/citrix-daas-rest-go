/*
Administrators APIs

APIs for managing CC administrators.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ccadmins

import (
	"encoding/json"
)

// checks if the AdministratorAccessPolicyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdministratorAccessPolicyModel{}

// AdministratorAccessPolicyModel struct for AdministratorAccessPolicyModel
type AdministratorAccessPolicyModel struct {
	Name string `json:"name"`
	ServiceName string `json:"serviceName"`
	DisplayName NullableString `json:"displayName,omitempty"`
	Checkable BooleanPolicyProperty `json:"checkable"`
	ScopeChoices *AdministratorAccessScopeChoices `json:"scopeChoices,omitempty"`
}

// NewAdministratorAccessPolicyModel instantiates a new AdministratorAccessPolicyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdministratorAccessPolicyModel(name string, serviceName string, checkable BooleanPolicyProperty) *AdministratorAccessPolicyModel {
	this := AdministratorAccessPolicyModel{}
	this.Name = name
	this.ServiceName = serviceName
	this.Checkable = checkable
	return &this
}

// NewAdministratorAccessPolicyModelWithDefaults instantiates a new AdministratorAccessPolicyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdministratorAccessPolicyModelWithDefaults() *AdministratorAccessPolicyModel {
	this := AdministratorAccessPolicyModel{}
	return &this
}

// GetName returns the Name field value
func (o *AdministratorAccessPolicyModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AdministratorAccessPolicyModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AdministratorAccessPolicyModel) SetName(v string) {
	o.Name = v
}

// GetServiceName returns the ServiceName field value
func (o *AdministratorAccessPolicyModel) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *AdministratorAccessPolicyModel) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *AdministratorAccessPolicyModel) SetServiceName(v string) {
	o.ServiceName = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdministratorAccessPolicyModel) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdministratorAccessPolicyModel) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AdministratorAccessPolicyModel) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *AdministratorAccessPolicyModel) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *AdministratorAccessPolicyModel) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *AdministratorAccessPolicyModel) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetCheckable returns the Checkable field value
func (o *AdministratorAccessPolicyModel) GetCheckable() BooleanPolicyProperty {
	if o == nil {
		var ret BooleanPolicyProperty
		return ret
	}

	return o.Checkable
}

// GetCheckableOk returns a tuple with the Checkable field value
// and a boolean to check if the value has been set.
func (o *AdministratorAccessPolicyModel) GetCheckableOk() (*BooleanPolicyProperty, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Checkable, true
}

// SetCheckable sets field value
func (o *AdministratorAccessPolicyModel) SetCheckable(v BooleanPolicyProperty) {
	o.Checkable = v
}

// GetScopeChoices returns the ScopeChoices field value if set, zero value otherwise.
func (o *AdministratorAccessPolicyModel) GetScopeChoices() AdministratorAccessScopeChoices {
	if o == nil || IsNil(o.ScopeChoices) {
		var ret AdministratorAccessScopeChoices
		return ret
	}
	return *o.ScopeChoices
}

// GetScopeChoicesOk returns a tuple with the ScopeChoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministratorAccessPolicyModel) GetScopeChoicesOk() (*AdministratorAccessScopeChoices, bool) {
	if o == nil || IsNil(o.ScopeChoices) {
		return nil, false
	}
	return o.ScopeChoices, true
}

// HasScopeChoices returns a boolean if a field has been set.
func (o *AdministratorAccessPolicyModel) HasScopeChoices() bool {
	if o != nil && !IsNil(o.ScopeChoices) {
		return true
	}

	return false
}

// SetScopeChoices gets a reference to the given AdministratorAccessScopeChoices and assigns it to the ScopeChoices field.
func (o *AdministratorAccessPolicyModel) SetScopeChoices(v AdministratorAccessScopeChoices) {
	o.ScopeChoices = &v
}

func (o AdministratorAccessPolicyModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdministratorAccessPolicyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["serviceName"] = o.ServiceName
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	toSerialize["checkable"] = o.Checkable
	if !IsNil(o.ScopeChoices) {
		toSerialize["scopeChoices"] = o.ScopeChoices
	}
	return toSerialize, nil
}

type NullableAdministratorAccessPolicyModel struct {
	value *AdministratorAccessPolicyModel
	isSet bool
}

func (v NullableAdministratorAccessPolicyModel) Get() *AdministratorAccessPolicyModel {
	return v.value
}

func (v *NullableAdministratorAccessPolicyModel) Set(val *AdministratorAccessPolicyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAdministratorAccessPolicyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAdministratorAccessPolicyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdministratorAccessPolicyModel(val *AdministratorAccessPolicyModel) *NullableAdministratorAccessPolicyModel {
	return &NullableAdministratorAccessPolicyModel{value: val, isSet: true}
}

func (v NullableAdministratorAccessPolicyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdministratorAccessPolicyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

