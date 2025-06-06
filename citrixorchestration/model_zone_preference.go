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

// ZonePreference Settings for zone preference and failover.
type ZonePreference string

// List of ZonePreference
const (
	ZONEPREFERENCE_UNKNOWN          ZonePreference = "Unknown"
	ZONEPREFERENCE_USER_LOCATION    ZonePreference = "UserLocation"
	ZONEPREFERENCE_USER_HOME        ZonePreference = "UserHome"
	ZONEPREFERENCE_USER_HOME_ONLY   ZonePreference = "UserHomeOnly"
	ZONEPREFERENCE_APPLICATION_HOME ZonePreference = "ApplicationHome"
)

// All allowed values of ZonePreference enum
var AllowedZonePreferenceEnumValues = []ZonePreference{
	"Unknown",
	"UserLocation",
	"UserHome",
	"UserHomeOnly",
	"ApplicationHome",
}

func (v *ZonePreference) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = ZonePreference(value)
	return nil
}

// NewZonePreferenceFromValue returns a pointer to a valid ZonePreference
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewZonePreferenceFromValue(v string) (*ZonePreference, error) {
	ev := ZonePreference(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ZonePreference: valid values are %v", v, AllowedZonePreferenceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ZonePreference) IsValid() bool {
	for _, existing := range AllowedZonePreferenceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ZonePreference value
func (v ZonePreference) Ptr() *ZonePreference {
	return &v
}

type NullableZonePreference struct {
	value *ZonePreference
	isSet bool
}

func (v NullableZonePreference) Get() *ZonePreference {
	return v.value
}

func (v *NullableZonePreference) Set(val *ZonePreference) {
	v.value = val
	v.isSet = true
}

func (v NullableZonePreference) IsSet() bool {
	return v.isSet
}

func (v *NullableZonePreference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZonePreference(val *ZonePreference) *NullableZonePreference {
	return &NullableZonePreference{value: val, isSet: true}
}

func (v NullableZonePreference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZonePreference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
