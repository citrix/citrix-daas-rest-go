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

// checks if the AdministratorRightResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdministratorRightResponseModel{}

// AdministratorRightResponseModel Response object for an administrator right, which is a combination of a role and a scope.
type AdministratorRightResponseModel struct {
	Scope ScopeResponseModel `json:"Scope"`
	Role  RoleResponseModel  `json:"Role"`
}

// NewAdministratorRightResponseModel instantiates a new AdministratorRightResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdministratorRightResponseModel(scope ScopeResponseModel, role RoleResponseModel) *AdministratorRightResponseModel {
	this := AdministratorRightResponseModel{}
	this.Scope = scope
	this.Role = role
	return &this
}

// NewAdministratorRightResponseModelWithDefaults instantiates a new AdministratorRightResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdministratorRightResponseModelWithDefaults() *AdministratorRightResponseModel {
	this := AdministratorRightResponseModel{}
	return &this
}

// GetScope returns the Scope field value
func (o *AdministratorRightResponseModel) GetScope() ScopeResponseModel {
	if o == nil {
		var ret ScopeResponseModel
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *AdministratorRightResponseModel) GetScopeOk() (*ScopeResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *AdministratorRightResponseModel) SetScope(v ScopeResponseModel) {
	o.Scope = v
}

// GetRole returns the Role field value
func (o *AdministratorRightResponseModel) GetRole() RoleResponseModel {
	if o == nil {
		var ret RoleResponseModel
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *AdministratorRightResponseModel) GetRoleOk() (*RoleResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *AdministratorRightResponseModel) SetRole(v RoleResponseModel) {
	o.Role = v
}

func (o AdministratorRightResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdministratorRightResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Scope"] = o.Scope
	toSerialize["Role"] = o.Role
	return toSerialize, nil
}

type NullableAdministratorRightResponseModel struct {
	value *AdministratorRightResponseModel
	isSet bool
}

func (v NullableAdministratorRightResponseModel) Get() *AdministratorRightResponseModel {
	return v.value
}

func (v *NullableAdministratorRightResponseModel) Set(val *AdministratorRightResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAdministratorRightResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAdministratorRightResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdministratorRightResponseModel(val *AdministratorRightResponseModel) *NullableAdministratorRightResponseModel {
	return &NullableAdministratorRightResponseModel{value: val, isSet: true}
}

func (v NullableAdministratorRightResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdministratorRightResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
