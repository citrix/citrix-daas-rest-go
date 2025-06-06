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

// RebootScheduleDays Days to use for a reboot schedule.
type RebootScheduleDays string

// List of RebootScheduleDays
const (
	REBOOTSCHEDULEDAYS_UNKNOWN   RebootScheduleDays = "Unknown"
	REBOOTSCHEDULEDAYS_MONDAY    RebootScheduleDays = "Monday"
	REBOOTSCHEDULEDAYS_TUESDAY   RebootScheduleDays = "Tuesday"
	REBOOTSCHEDULEDAYS_WEDNESDAY RebootScheduleDays = "Wednesday"
	REBOOTSCHEDULEDAYS_THURSDAY  RebootScheduleDays = "Thursday"
	REBOOTSCHEDULEDAYS_FRIDAY    RebootScheduleDays = "Friday"
	REBOOTSCHEDULEDAYS_SATURDAY  RebootScheduleDays = "Saturday"
	REBOOTSCHEDULEDAYS_SUNDAY    RebootScheduleDays = "Sunday"
)

// All allowed values of RebootScheduleDays enum
var AllowedRebootScheduleDaysEnumValues = []RebootScheduleDays{
	"Unknown",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
}

func (v *RebootScheduleDays) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = RebootScheduleDays(value)
	return nil
}

// NewRebootScheduleDaysFromValue returns a pointer to a valid RebootScheduleDays
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRebootScheduleDaysFromValue(v string) (*RebootScheduleDays, error) {
	ev := RebootScheduleDays(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RebootScheduleDays: valid values are %v", v, AllowedRebootScheduleDaysEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RebootScheduleDays) IsValid() bool {
	for _, existing := range AllowedRebootScheduleDaysEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RebootScheduleDays value
func (v RebootScheduleDays) Ptr() *RebootScheduleDays {
	return &v
}

type NullableRebootScheduleDays struct {
	value *RebootScheduleDays
	isSet bool
}

func (v NullableRebootScheduleDays) Get() *RebootScheduleDays {
	return v.value
}

func (v *NullableRebootScheduleDays) Set(val *RebootScheduleDays) {
	v.value = val
	v.isSet = true
}

func (v NullableRebootScheduleDays) IsSet() bool {
	return v.isSet
}

func (v *NullableRebootScheduleDays) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRebootScheduleDays(val *RebootScheduleDays) *NullableRebootScheduleDays {
	return &NullableRebootScheduleDays{value: val, isSet: true}
}

func (v NullableRebootScheduleDays) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRebootScheduleDays) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
