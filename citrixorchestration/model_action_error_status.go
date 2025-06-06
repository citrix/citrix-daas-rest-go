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

// ActionErrorStatus Action error status.
type ActionErrorStatus string

// List of ActionErrorStatus
const (
	ACTIONERRORSTATUS_UNKNOWN      ActionErrorStatus = "Unknown"
	ACTIONERRORSTATUS_NONE         ActionErrorStatus = "None"
	ACTIONERRORSTATUS_FAILED       ActionErrorStatus = "Failed"
	ACTIONERRORSTATUS_PARTIAL_FAIL ActionErrorStatus = "PartialFail"
)

// All allowed values of ActionErrorStatus enum
var AllowedActionErrorStatusEnumValues = []ActionErrorStatus{
	"Unknown",
	"None",
	"Failed",
	"PartialFail",
}

func (v *ActionErrorStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = ActionErrorStatus(value)
	return nil
}

// NewActionErrorStatusFromValue returns a pointer to a valid ActionErrorStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewActionErrorStatusFromValue(v string) (*ActionErrorStatus, error) {
	ev := ActionErrorStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ActionErrorStatus: valid values are %v", v, AllowedActionErrorStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ActionErrorStatus) IsValid() bool {
	for _, existing := range AllowedActionErrorStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ActionErrorStatus value
func (v ActionErrorStatus) Ptr() *ActionErrorStatus {
	return &v
}

type NullableActionErrorStatus struct {
	value *ActionErrorStatus
	isSet bool
}

func (v NullableActionErrorStatus) Get() *ActionErrorStatus {
	return v.value
}

func (v *NullableActionErrorStatus) Set(val *ActionErrorStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableActionErrorStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableActionErrorStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionErrorStatus(val *ActionErrorStatus) *NullableActionErrorStatus {
	return &NullableActionErrorStatus{value: val, isSet: true}
}

func (v NullableActionErrorStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionErrorStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
