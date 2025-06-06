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

// checks if the AdministratorResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdministratorResponseModel{}

// AdministratorResponseModel Response object representing an administrator.
type AdministratorResponseModel struct {
	User IdentityUserResponseModel `json:"User"`
	// Rights associated with the administrator. CHANGE: was public ScopeRolePair[] ScopesAndRoles { get; set; }
	ScopesAndRoles []AdministratorRightResponseModel `json:"ScopesAndRoles"`
	// Indicates whether the administrator is enabled. Disabled administrators cannot administer the site UNLESS they are a member of a different user group which is granted access by a different administrator record.
	Enabled bool `json:"Enabled"`
}

// NewAdministratorResponseModel instantiates a new AdministratorResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdministratorResponseModel(user IdentityUserResponseModel, scopesAndRoles []AdministratorRightResponseModel, enabled bool) *AdministratorResponseModel {
	this := AdministratorResponseModel{}
	this.User = user
	this.ScopesAndRoles = scopesAndRoles
	this.Enabled = enabled
	return &this
}

// NewAdministratorResponseModelWithDefaults instantiates a new AdministratorResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdministratorResponseModelWithDefaults() *AdministratorResponseModel {
	this := AdministratorResponseModel{}
	return &this
}

// GetUser returns the User field value
func (o *AdministratorResponseModel) GetUser() IdentityUserResponseModel {
	if o == nil {
		var ret IdentityUserResponseModel
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *AdministratorResponseModel) GetUserOk() (*IdentityUserResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *AdministratorResponseModel) SetUser(v IdentityUserResponseModel) {
	o.User = v
}

// GetScopesAndRoles returns the ScopesAndRoles field value
func (o *AdministratorResponseModel) GetScopesAndRoles() []AdministratorRightResponseModel {
	if o == nil {
		var ret []AdministratorRightResponseModel
		return ret
	}

	return o.ScopesAndRoles
}

// GetScopesAndRolesOk returns a tuple with the ScopesAndRoles field value
// and a boolean to check if the value has been set.
func (o *AdministratorResponseModel) GetScopesAndRolesOk() ([]AdministratorRightResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScopesAndRoles, true
}

// SetScopesAndRoles sets field value
func (o *AdministratorResponseModel) SetScopesAndRoles(v []AdministratorRightResponseModel) {
	o.ScopesAndRoles = v
}

// GetEnabled returns the Enabled field value
func (o *AdministratorResponseModel) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AdministratorResponseModel) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AdministratorResponseModel) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AdministratorResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdministratorResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["User"] = o.User
	toSerialize["ScopesAndRoles"] = o.ScopesAndRoles
	toSerialize["Enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableAdministratorResponseModel struct {
	value *AdministratorResponseModel
	isSet bool
}

func (v NullableAdministratorResponseModel) Get() *AdministratorResponseModel {
	return v.value
}

func (v *NullableAdministratorResponseModel) Set(val *AdministratorResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAdministratorResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAdministratorResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdministratorResponseModel(val *AdministratorResponseModel) *NullableAdministratorResponseModel {
	return &NullableAdministratorResponseModel{value: val, isSet: true}
}

func (v NullableAdministratorResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdministratorResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
