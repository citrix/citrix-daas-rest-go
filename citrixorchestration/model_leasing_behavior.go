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

// LeasingBehavior Leasing behavior.
type LeasingBehavior string

// List of LeasingBehavior
const (
	LEASINGBEHAVIOR_UNKNOWN    LeasingBehavior = "Unknown"
	LEASINGBEHAVIOR_ALLOWED    LeasingBehavior = "Allowed"
	LEASINGBEHAVIOR_DISALLOWED LeasingBehavior = "Disallowed"
)

// All allowed values of LeasingBehavior enum
var AllowedLeasingBehaviorEnumValues = []LeasingBehavior{
	"Unknown",
	"Allowed",
	"Disallowed",
}

func (v *LeasingBehavior) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = LeasingBehavior(value)
	return nil
}

// NewLeasingBehaviorFromValue returns a pointer to a valid LeasingBehavior
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLeasingBehaviorFromValue(v string) (*LeasingBehavior, error) {
	ev := LeasingBehavior(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LeasingBehavior: valid values are %v", v, AllowedLeasingBehaviorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LeasingBehavior) IsValid() bool {
	for _, existing := range AllowedLeasingBehaviorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LeasingBehavior value
func (v LeasingBehavior) Ptr() *LeasingBehavior {
	return &v
}

type NullableLeasingBehavior struct {
	value *LeasingBehavior
	isSet bool
}

func (v NullableLeasingBehavior) Get() *LeasingBehavior {
	return v.value
}

func (v *NullableLeasingBehavior) Set(val *LeasingBehavior) {
	v.value = val
	v.isSet = true
}

func (v NullableLeasingBehavior) IsSet() bool {
	return v.isSet
}

func (v *NullableLeasingBehavior) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLeasingBehavior(val *LeasingBehavior) *NullableLeasingBehavior {
	return &NullableLeasingBehavior{value: val, isSet: true}
}

func (v NullableLeasingBehavior) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLeasingBehavior) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
