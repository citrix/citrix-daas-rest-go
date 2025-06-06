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

// RegistrationState Machine registration states.
type RegistrationState string

// List of RegistrationState
const (
	REGISTRATIONSTATE_UNKNOWN      RegistrationState = "Unknown"
	REGISTRATIONSTATE_UNREGISTERED RegistrationState = "Unregistered"
	REGISTRATIONSTATE_INITIALIZING RegistrationState = "Initializing"
	REGISTRATIONSTATE_REGISTERED   RegistrationState = "Registered"
	REGISTRATIONSTATE_AGENT_ERROR  RegistrationState = "AgentError"
)

// All allowed values of RegistrationState enum
var AllowedRegistrationStateEnumValues = []RegistrationState{
	"Unknown",
	"Unregistered",
	"Initializing",
	"Registered",
	"AgentError",
}

func (v *RegistrationState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RegistrationState(value)
	for _, existing := range AllowedRegistrationStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RegistrationState", value)
}

// NewRegistrationStateFromValue returns a pointer to a valid RegistrationState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRegistrationStateFromValue(v string) (*RegistrationState, error) {
	ev := RegistrationState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RegistrationState: valid values are %v", v, AllowedRegistrationStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RegistrationState) IsValid() bool {
	for _, existing := range AllowedRegistrationStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RegistrationState value
func (v RegistrationState) Ptr() *RegistrationState {
	return &v
}

type NullableRegistrationState struct {
	value *RegistrationState
	isSet bool
}

func (v NullableRegistrationState) Get() *RegistrationState {
	return v.value
}

func (v *NullableRegistrationState) Set(val *RegistrationState) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationState) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationState(val *RegistrationState) *NullableRegistrationState {
	return &NullableRegistrationState{value: val, isSet: true}
}

func (v NullableRegistrationState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
