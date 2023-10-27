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

// RequiredSleepCapability The sleep capability of this desktop group.             
type RequiredSleepCapability string

// List of RequiredSleepCapability
const (
	REQUIREDSLEEPCAPABILITY_UNKNOWN RequiredSleepCapability = "Unknown"
	REQUIREDSLEEPCAPABILITY_NONE RequiredSleepCapability = "None"
	REQUIREDSLEEPCAPABILITY_SUSPEND RequiredSleepCapability = "Suspend"
)

// All allowed values of RequiredSleepCapability enum
var AllowedRequiredSleepCapabilityEnumValues = []RequiredSleepCapability{
	"Unknown",
	"None",
	"Suspend",
}

func (v *RequiredSleepCapability) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequiredSleepCapability(value)
	for _, existing := range AllowedRequiredSleepCapabilityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequiredSleepCapability", value)
}

// NewRequiredSleepCapabilityFromValue returns a pointer to a valid RequiredSleepCapability
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequiredSleepCapabilityFromValue(v string) (*RequiredSleepCapability, error) {
	ev := RequiredSleepCapability(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequiredSleepCapability: valid values are %v", v, AllowedRequiredSleepCapabilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequiredSleepCapability) IsValid() bool {
	for _, existing := range AllowedRequiredSleepCapabilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RequiredSleepCapability value
func (v RequiredSleepCapability) Ptr() *RequiredSleepCapability {
	return &v
}

type NullableRequiredSleepCapability struct {
	value *RequiredSleepCapability
	isSet bool
}

func (v NullableRequiredSleepCapability) Get() *RequiredSleepCapability {
	return v.value
}

func (v *NullableRequiredSleepCapability) Set(val *RequiredSleepCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableRequiredSleepCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableRequiredSleepCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequiredSleepCapability(val *RequiredSleepCapability) *NullableRequiredSleepCapability {
	return &NullableRequiredSleepCapability{value: val, isSet: true}
}

func (v NullableRequiredSleepCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequiredSleepCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
