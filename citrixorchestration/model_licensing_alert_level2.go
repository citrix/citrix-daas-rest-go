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

// LicensingAlertLevel2
type LicensingAlertLevel2 string

// List of LicensingAlertLevel2
const (
	LICENSINGALERTLEVEL2_UNKNOWN LicensingAlertLevel2 = "Unknown"
	LICENSINGALERTLEVEL2_INFO    LicensingAlertLevel2 = "Info"
	LICENSINGALERTLEVEL2_ALERT   LicensingAlertLevel2 = "Alert"
	LICENSINGALERTLEVEL2_ALARM   LicensingAlertLevel2 = "Alarm"
)

// All allowed values of LicensingAlertLevel2 enum
var AllowedLicensingAlertLevel2EnumValues = []LicensingAlertLevel2{
	"Unknown",
	"Info",
	"Alert",
	"Alarm",
}

func (v *LicensingAlertLevel2) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = LicensingAlertLevel2(value)
	return nil
}

// NewLicensingAlertLevel2FromValue returns a pointer to a valid LicensingAlertLevel2
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLicensingAlertLevel2FromValue(v string) (*LicensingAlertLevel2, error) {
	ev := LicensingAlertLevel2(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LicensingAlertLevel2: valid values are %v", v, AllowedLicensingAlertLevel2EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LicensingAlertLevel2) IsValid() bool {
	for _, existing := range AllowedLicensingAlertLevel2EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LicensingAlertLevel2 value
func (v LicensingAlertLevel2) Ptr() *LicensingAlertLevel2 {
	return &v
}

type NullableLicensingAlertLevel2 struct {
	value *LicensingAlertLevel2
	isSet bool
}

func (v NullableLicensingAlertLevel2) Get() *LicensingAlertLevel2 {
	return v.value
}

func (v *NullableLicensingAlertLevel2) Set(val *LicensingAlertLevel2) {
	v.value = val
	v.isSet = true
}

func (v NullableLicensingAlertLevel2) IsSet() bool {
	return v.isSet
}

func (v *NullableLicensingAlertLevel2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicensingAlertLevel2(val *LicensingAlertLevel2) *NullableLicensingAlertLevel2 {
	return &NullableLicensingAlertLevel2{value: val, isSet: true}
}

func (v NullableLicensingAlertLevel2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicensingAlertLevel2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
