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

// MachineAccountRepairOption Option for how to repair machine accounts.
type MachineAccountRepairOption string

// List of MachineAccountRepairOption
const (
	MACHINEACCOUNTREPAIROPTION_UNKNOWN MachineAccountRepairOption = "Unknown"
	MACHINEACCOUNTREPAIROPTION_USER_CERTIFICATE MachineAccountRepairOption = "UserCertificate"
	MACHINEACCOUNTREPAIROPTION_IDENTITY_INFO MachineAccountRepairOption = "IdentityInfo"
)

// All allowed values of MachineAccountRepairOption enum
var AllowedMachineAccountRepairOptionEnumValues = []MachineAccountRepairOption{
	"Unknown",
	"UserCertificate",
	"IdentityInfo",
}

func (v *MachineAccountRepairOption) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MachineAccountRepairOption(value)
	for _, existing := range AllowedMachineAccountRepairOptionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MachineAccountRepairOption", value)
}

// NewMachineAccountRepairOptionFromValue returns a pointer to a valid MachineAccountRepairOption
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMachineAccountRepairOptionFromValue(v string) (*MachineAccountRepairOption, error) {
	ev := MachineAccountRepairOption(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MachineAccountRepairOption: valid values are %v", v, AllowedMachineAccountRepairOptionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MachineAccountRepairOption) IsValid() bool {
	for _, existing := range AllowedMachineAccountRepairOptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MachineAccountRepairOption value
func (v MachineAccountRepairOption) Ptr() *MachineAccountRepairOption {
	return &v
}

type NullableMachineAccountRepairOption struct {
	value *MachineAccountRepairOption
	isSet bool
}

func (v NullableMachineAccountRepairOption) Get() *MachineAccountRepairOption {
	return v.value
}

func (v *NullableMachineAccountRepairOption) Set(val *MachineAccountRepairOption) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineAccountRepairOption) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineAccountRepairOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineAccountRepairOption(val *MachineAccountRepairOption) *NullableMachineAccountRepairOption {
	return &NullableMachineAccountRepairOption{value: val, isSet: true}
}

func (v NullableMachineAccountRepairOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineAccountRepairOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
