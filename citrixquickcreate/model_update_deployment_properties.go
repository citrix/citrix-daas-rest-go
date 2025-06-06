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

// checks if the UpdateDeploymentProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeploymentProperties{}

// UpdateDeploymentProperties Update deployment properties which includes machine performance, running mode and power management settings  depending on deployment type
type UpdateDeploymentProperties struct {
	// The type of provider associated with the account
	AccountType AccountType `json:"accountType"`
}

type _UpdateDeploymentProperties UpdateDeploymentProperties

// NewUpdateDeploymentProperties instantiates a new UpdateDeploymentProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeploymentProperties(accountType AccountType) *UpdateDeploymentProperties {
	this := UpdateDeploymentProperties{}
	this.AccountType = accountType
	return &this
}

// NewUpdateDeploymentPropertiesWithDefaults instantiates a new UpdateDeploymentProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeploymentPropertiesWithDefaults() *UpdateDeploymentProperties {
	this := UpdateDeploymentProperties{}
	return &this
}

// GetAccountType returns the AccountType field value
func (o *UpdateDeploymentProperties) GetAccountType() AccountType {
	if o == nil {
		var ret AccountType
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *UpdateDeploymentProperties) GetAccountTypeOk() (*AccountType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *UpdateDeploymentProperties) SetAccountType(v AccountType) {
	o.AccountType = v
}

func (o UpdateDeploymentProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeploymentProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountType"] = o.AccountType
	return toSerialize, nil
}

type NullableUpdateDeploymentProperties struct {
	value *UpdateDeploymentProperties
	isSet bool
}

func (v NullableUpdateDeploymentProperties) Get() *UpdateDeploymentProperties {
	return v.value
}

func (v *NullableUpdateDeploymentProperties) Set(val *UpdateDeploymentProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeploymentProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeploymentProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeploymentProperties(val *UpdateDeploymentProperties) *NullableUpdateDeploymentProperties {
	return &NullableUpdateDeploymentProperties{value: val, isSet: true}
}

func (v NullableUpdateDeploymentProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeploymentProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
