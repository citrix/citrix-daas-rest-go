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

// AllowedConnection
type AllowedConnection string

// List of AllowedConnection
const (
	ALLOWEDCONNECTION_UNKNOWN    AllowedConnection = "Unknown"
	ALLOWEDCONNECTION_FILTERED   AllowedConnection = "Filtered"
	ALLOWEDCONNECTION_NOT_VIA_AG AllowedConnection = "NotViaAG"
	ALLOWEDCONNECTION_VIA_AG     AllowedConnection = "ViaAG"
	ALLOWEDCONNECTION_ANY_VIA_AG AllowedConnection = "AnyViaAG"
)

// All allowed values of AllowedConnection enum
var AllowedAllowedConnectionEnumValues = []AllowedConnection{
	"Unknown",
	"Filtered",
	"NotViaAG",
	"ViaAG",
	"AnyViaAG",
}

func (v *AllowedConnection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = AllowedConnection(value)
	return nil
}

// NewAllowedConnectionFromValue returns a pointer to a valid AllowedConnection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAllowedConnectionFromValue(v string) (*AllowedConnection, error) {
	ev := AllowedConnection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AllowedConnection: valid values are %v", v, AllowedAllowedConnectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AllowedConnection) IsValid() bool {
	for _, existing := range AllowedAllowedConnectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AllowedConnection value
func (v AllowedConnection) Ptr() *AllowedConnection {
	return &v
}

type NullableAllowedConnection struct {
	value *AllowedConnection
	isSet bool
}

func (v NullableAllowedConnection) Get() *AllowedConnection {
	return v.value
}

func (v *NullableAllowedConnection) Set(val *AllowedConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowedConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowedConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowedConnection(val *AllowedConnection) *NullableAllowedConnection {
	return &NullableAllowedConnection{value: val, isSet: true}
}

func (v NullableAllowedConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowedConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
