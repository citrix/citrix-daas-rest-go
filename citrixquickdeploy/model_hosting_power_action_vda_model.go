/*
Citrix Virtual App & Desktop Catalog Service 147.0.26651.57932

Catalog Service

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickdeploy

import (
	"encoding/json"
)

// checks if the HostingPowerActionVdaModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostingPowerActionVdaModel{}

// HostingPowerActionVdaModel XD hosting power action - applied to vda
type HostingPowerActionVdaModel struct {
	ActionType *PowerManagementAction `json:"actionType,omitempty"`
}

// NewHostingPowerActionVdaModelWithDefaults instantiates a new HostingPowerActionVdaModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostingPowerActionVdaModelWithDefaults() *HostingPowerActionVdaModel {
	this := HostingPowerActionVdaModel{}
	return &this
}

// GetActionType returns the ActionType field value if set, zero value otherwise.
func (o *HostingPowerActionVdaModel) GetActionType() PowerManagementAction {
	if o == nil || IsNil(o.ActionType) {
		var ret PowerManagementAction
		return ret
	}
	return *o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostingPowerActionVdaModel) GetActionTypeOk() (*PowerManagementAction, bool) {
	if o == nil || IsNil(o.ActionType) {
		return nil, false
	}
	return o.ActionType, true
}

// SetActionType gets a reference to the given PowerManagementAction and assigns it to the ActionType field.
func (o *HostingPowerActionVdaModel) SetActionType(v PowerManagementAction) {
	o.ActionType = &v
}

func (o HostingPowerActionVdaModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostingPowerActionVdaModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionType) {
		toSerialize["actionType"] = o.ActionType
	}
	return toSerialize, nil
}

type NullableHostingPowerActionVdaModel struct {
	value *HostingPowerActionVdaModel
	isSet bool
}

func (v NullableHostingPowerActionVdaModel) Get() *HostingPowerActionVdaModel {
	return v.value
}

func (v *NullableHostingPowerActionVdaModel) Set(val *HostingPowerActionVdaModel) {
	v.value = val
	v.isSet = true
}

func (v NullableHostingPowerActionVdaModel) IsSet() bool {
	return v.isSet
}

func (v *NullableHostingPowerActionVdaModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostingPowerActionVdaModel(val *HostingPowerActionVdaModel) *NullableHostingPowerActionVdaModel {
	return &NullableHostingPowerActionVdaModel{value: val, isSet: true}
}

func (v NullableHostingPowerActionVdaModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostingPowerActionVdaModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
