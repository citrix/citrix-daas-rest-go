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

// CpuPriorityLevel Starting process priority level for an application.
type CpuPriorityLevel string

// List of CpuPriorityLevel
const (
	CPUPRIORITYLEVEL_UNKNOWN      CpuPriorityLevel = "Unknown"
	CPUPRIORITYLEVEL_LOW          CpuPriorityLevel = "Low"
	CPUPRIORITYLEVEL_BELOW_NORMAL CpuPriorityLevel = "BelowNormal"
	CPUPRIORITYLEVEL_NORMAL       CpuPriorityLevel = "Normal"
	CPUPRIORITYLEVEL_ABOVE_NORMAL CpuPriorityLevel = "AboveNormal"
	CPUPRIORITYLEVEL_HIGH         CpuPriorityLevel = "High"
)

// All allowed values of CpuPriorityLevel enum
var AllowedCpuPriorityLevelEnumValues = []CpuPriorityLevel{
	"Unknown",
	"Low",
	"BelowNormal",
	"Normal",
	"AboveNormal",
	"High",
}

func (v *CpuPriorityLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = CpuPriorityLevel(value)
	return nil
}

// NewCpuPriorityLevelFromValue returns a pointer to a valid CpuPriorityLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCpuPriorityLevelFromValue(v string) (*CpuPriorityLevel, error) {
	ev := CpuPriorityLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CpuPriorityLevel: valid values are %v", v, AllowedCpuPriorityLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CpuPriorityLevel) IsValid() bool {
	for _, existing := range AllowedCpuPriorityLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CpuPriorityLevel value
func (v CpuPriorityLevel) Ptr() *CpuPriorityLevel {
	return &v
}

type NullableCpuPriorityLevel struct {
	value *CpuPriorityLevel
	isSet bool
}

func (v NullableCpuPriorityLevel) Get() *CpuPriorityLevel {
	return v.value
}

func (v *NullableCpuPriorityLevel) Set(val *CpuPriorityLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableCpuPriorityLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableCpuPriorityLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpuPriorityLevel(val *CpuPriorityLevel) *NullableCpuPriorityLevel {
	return &NullableCpuPriorityLevel{value: val, isSet: true}
}

func (v NullableCpuPriorityLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpuPriorityLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
