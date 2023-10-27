/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-onprem-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
	"fmt"
)

// SettingProperty Names of setting properties that can be specified in a search.
type SettingProperty string

// List of SettingProperty
const (
	SETTINGPROPERTY_SETTING_NAME SettingProperty = "SettingName"
	SETTINGPROPERTY_SETTING_VALUE SettingProperty = "SettingValue"
	SETTINGPROPERTY_USE_DEFAULT SettingProperty = "UseDefault"
)

// All allowed values of SettingProperty enum
var AllowedSettingPropertyEnumValues = []SettingProperty{
	"SettingName",
	"SettingValue",
	"UseDefault",
}

func (v *SettingProperty) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SettingProperty(value)
	for _, existing := range AllowedSettingPropertyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SettingProperty", value)
}

// NewSettingPropertyFromValue returns a pointer to a valid SettingProperty
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSettingPropertyFromValue(v string) (*SettingProperty, error) {
	ev := SettingProperty(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SettingProperty: valid values are %v", v, AllowedSettingPropertyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SettingProperty) IsValid() bool {
	for _, existing := range AllowedSettingPropertyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SettingProperty value
func (v SettingProperty) Ptr() *SettingProperty {
	return &v
}

type NullableSettingProperty struct {
	value *SettingProperty
	isSet bool
}

func (v NullableSettingProperty) Get() *SettingProperty {
	return v.value
}

func (v *NullableSettingProperty) Set(val *SettingProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingProperty(val *SettingProperty) *NullableSettingProperty {
	return &NullableSettingProperty{value: val, isSet: true}
}

func (v NullableSettingProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
