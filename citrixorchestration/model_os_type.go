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

// OsType Operating System types.
type OsType string

// List of OsType
const (
	OSTYPE_UNKNOWN OsType = "Unknown"
	OSTYPE_WINDOWS OsType = "Windows"
	OSTYPE_LINUX OsType = "Linux"
)

// All allowed values of OsType enum
var AllowedOsTypeEnumValues = []OsType{
	"Unknown",
	"Windows",
	"Linux",
}

func (v *OsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OsType(value)
	for _, existing := range AllowedOsTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OsType", value)
}

// NewOsTypeFromValue returns a pointer to a valid OsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOsTypeFromValue(v string) (*OsType, error) {
	ev := OsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OsType: valid values are %v", v, AllowedOsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OsType) IsValid() bool {
	for _, existing := range AllowedOsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OsType value
func (v OsType) Ptr() *OsType {
	return &v
}

type NullableOsType struct {
	value *OsType
	isSet bool
}

func (v NullableOsType) Get() *OsType {
	return v.value
}

func (v *NullableOsType) Set(val *OsType) {
	v.value = val
	v.isSet = true
}

func (v NullableOsType) IsSet() bool {
	return v.isSet
}

func (v *NullableOsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsType(val *OsType) *NullableOsType {
	return &NullableOsType{value: val, isSet: true}
}

func (v NullableOsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
