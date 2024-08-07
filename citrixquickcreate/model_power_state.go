/*
Quick Create Service for DaaS Quick Deploy with AWS WorkSpaces Core - Tech Preview

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: DaaSTechPreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickcreate

import (
	"encoding/json"
	"fmt"
)

// PowerState The power state of frontline Cloud PCs.
type PowerState string

// List of PowerState
const (
	POWERSTATE_RUNNING PowerState = "Running"
	POWERSTATE_POWERED_OFF PowerState = "PoweredOff"
	POWERSTATE_UNKNOWN PowerState = "Unknown"
)

// All allowed values of PowerState enum
var AllowedPowerStateEnumValues = []PowerState{
	"Running",
	"PoweredOff",
	"Unknown",
}

func (v *PowerState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PowerState(value)
	for _, existing := range AllowedPowerStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PowerState", value)
}

// NewPowerStateFromValue returns a pointer to a valid PowerState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPowerStateFromValue(v string) (*PowerState, error) {
	ev := PowerState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PowerState: valid values are %v", v, AllowedPowerStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PowerState) IsValid() bool {
	for _, existing := range AllowedPowerStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PowerState value
func (v PowerState) Ptr() *PowerState {
	return &v
}

type NullablePowerState struct {
	value *PowerState
	isSet bool
}

func (v NullablePowerState) Get() *PowerState {
	return v.value
}

func (v *NullablePowerState) Set(val *PowerState) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerState) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerState(val *PowerState) *NullablePowerState {
	return &NullablePowerState{value: val, isSet: true}
}

func (v NullablePowerState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
