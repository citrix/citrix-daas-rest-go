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

// FunctionalLevel Functional level for VDAs.
type FunctionalLevel string

// List of FunctionalLevel
const (
	FUNCTIONALLEVEL_UNKNOWN FunctionalLevel = "Unknown"
	FUNCTIONALLEVEL_L5      FunctionalLevel = "L5"
	FUNCTIONALLEVEL_LMIN    FunctionalLevel = "LMIN"
	FUNCTIONALLEVEL_L7      FunctionalLevel = "L7"
	FUNCTIONALLEVEL_L7_6    FunctionalLevel = "L7_6"
	FUNCTIONALLEVEL_L7_7    FunctionalLevel = "L7_7"
	FUNCTIONALLEVEL_L7_8    FunctionalLevel = "L7_8"
	FUNCTIONALLEVEL_L7_9    FunctionalLevel = "L7_9"
	FUNCTIONALLEVEL_L7_20   FunctionalLevel = "L7_20"
	FUNCTIONALLEVEL_L7_25   FunctionalLevel = "L7_25"
	FUNCTIONALLEVEL_L7_30   FunctionalLevel = "L7_30"
	FUNCTIONALLEVEL_L7_34   FunctionalLevel = "L7_34"
	FUNCTIONALLEVEL_LMAX    FunctionalLevel = "LMAX"
)

// All allowed values of FunctionalLevel enum
var AllowedFunctionalLevelEnumValues = []FunctionalLevel{
	"Unknown",
	"L5",
	"LMIN",
	"L7",
	"L7_6",
	"L7_7",
	"L7_8",
	"L7_9",
	"L7_20",
	"L7_25",
	"L7_30",
	"L7_34",
	"LMAX",
}

func (v *FunctionalLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = FunctionalLevel(value)
	return nil
}

// NewFunctionalLevelFromValue returns a pointer to a valid FunctionalLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFunctionalLevelFromValue(v string) (*FunctionalLevel, error) {
	ev := FunctionalLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FunctionalLevel: valid values are %v", v, AllowedFunctionalLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FunctionalLevel) IsValid() bool {
	for _, existing := range AllowedFunctionalLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FunctionalLevel value
func (v FunctionalLevel) Ptr() *FunctionalLevel {
	return &v
}

type NullableFunctionalLevel struct {
	value *FunctionalLevel
	isSet bool
}

func (v NullableFunctionalLevel) Get() *FunctionalLevel {
	return v.value
}

func (v *NullableFunctionalLevel) Set(val *FunctionalLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionalLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionalLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionalLevel(val *FunctionalLevel) *NullableFunctionalLevel {
	return &NullableFunctionalLevel{value: val, isSet: true}
}

func (v NullableFunctionalLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionalLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
