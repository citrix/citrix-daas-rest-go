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

// PowerShellExecutionStatus The status of PowerShell execution
type PowerShellExecutionStatus string

// List of PowerShellExecutionStatus
const (
	POWERSHELLEXECUTIONSTATUS_UNKNOWN   PowerShellExecutionStatus = "Unknown"
	POWERSHELLEXECUTIONSTATUS_FAILED    PowerShellExecutionStatus = "Failed"
	POWERSHELLEXECUTIONSTATUS_CANCELED  PowerShellExecutionStatus = "Canceled"
	POWERSHELLEXECUTIONSTATUS_COMPLETED PowerShellExecutionStatus = "Completed"
)

// All allowed values of PowerShellExecutionStatus enum
var AllowedPowerShellExecutionStatusEnumValues = []PowerShellExecutionStatus{
	"Unknown",
	"Failed",
	"Canceled",
	"Completed",
}

func (v *PowerShellExecutionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = PowerShellExecutionStatus(value)
	return nil
}

// NewPowerShellExecutionStatusFromValue returns a pointer to a valid PowerShellExecutionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPowerShellExecutionStatusFromValue(v string) (*PowerShellExecutionStatus, error) {
	ev := PowerShellExecutionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PowerShellExecutionStatus: valid values are %v", v, AllowedPowerShellExecutionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PowerShellExecutionStatus) IsValid() bool {
	for _, existing := range AllowedPowerShellExecutionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PowerShellExecutionStatus value
func (v PowerShellExecutionStatus) Ptr() *PowerShellExecutionStatus {
	return &v
}

type NullablePowerShellExecutionStatus struct {
	value *PowerShellExecutionStatus
	isSet bool
}

func (v NullablePowerShellExecutionStatus) Get() *PowerShellExecutionStatus {
	return v.value
}

func (v *NullablePowerShellExecutionStatus) Set(val *PowerShellExecutionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerShellExecutionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerShellExecutionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerShellExecutionStatus(val *PowerShellExecutionStatus) *NullablePowerShellExecutionStatus {
	return &NullablePowerShellExecutionStatus{value: val, isSet: true}
}

func (v NullablePowerShellExecutionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerShellExecutionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
