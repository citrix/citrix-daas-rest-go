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

// SessionSupport Quantity of sessions supported per-machine.
type SessionSupport string

// List of SessionSupport
const (
	SESSIONSUPPORT_UNKNOWN        SessionSupport = "Unknown"
	SESSIONSUPPORT_SINGLE_SESSION SessionSupport = "SingleSession"
	SESSIONSUPPORT_MULTI_SESSION  SessionSupport = "MultiSession"
)

// All allowed values of SessionSupport enum
var AllowedSessionSupportEnumValues = []SessionSupport{
	"Unknown",
	"SingleSession",
	"MultiSession",
}

func (v *SessionSupport) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = SessionSupport(value)
	return nil
}

// NewSessionSupportFromValue returns a pointer to a valid SessionSupport
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSessionSupportFromValue(v string) (*SessionSupport, error) {
	ev := SessionSupport(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SessionSupport: valid values are %v", v, AllowedSessionSupportEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SessionSupport) IsValid() bool {
	for _, existing := range AllowedSessionSupportEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SessionSupport value
func (v SessionSupport) Ptr() *SessionSupport {
	return &v
}

type NullableSessionSupport struct {
	value *SessionSupport
	isSet bool
}

func (v NullableSessionSupport) Get() *SessionSupport {
	return v.value
}

func (v *NullableSessionSupport) Set(val *SessionSupport) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionSupport) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionSupport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionSupport(val *SessionSupport) *NullableSessionSupport {
	return &NullableSessionSupport{value: val, isSet: true}
}

func (v NullableSessionSupport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionSupport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
