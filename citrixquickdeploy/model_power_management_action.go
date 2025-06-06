/*
Citrix Virtual App & Desktop Catalog Service 147.0.26651.57932

Catalog Service

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickdeploy

import (
	"encoding/json"
	"fmt"
)

// PowerManagementAction the model 'PowerManagementAction'
type PowerManagementAction string

// List of PowerManagementAction
const (
	POWERMANAGEMENTACTION_TURN_ON  PowerManagementAction = "TurnOn"
	POWERMANAGEMENTACTION_TURN_OFF PowerManagementAction = "TurnOff"
	POWERMANAGEMENTACTION_SHUTDOWN PowerManagementAction = "Shutdown"
	POWERMANAGEMENTACTION_RESET    PowerManagementAction = "Reset"
	POWERMANAGEMENTACTION_RESTART  PowerManagementAction = "Restart"
	POWERMANAGEMENTACTION_SUSPEND  PowerManagementAction = "Suspend"
	POWERMANAGEMENTACTION_RESUME   PowerManagementAction = "Resume"
)

// All allowed values of PowerManagementAction enum
var AllowedPowerManagementActionEnumValues = []PowerManagementAction{
	"TurnOn",
	"TurnOff",
	"Shutdown",
	"Reset",
	"Restart",
	"Suspend",
	"Resume",
}

func (v *PowerManagementAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = PowerManagementAction(value)
	return nil
}

// NewPowerManagementActionFromValue returns a pointer to a valid PowerManagementAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPowerManagementActionFromValue(v string) (*PowerManagementAction, error) {
	ev := PowerManagementAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PowerManagementAction: valid values are %v", v, AllowedPowerManagementActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PowerManagementAction) IsValid() bool {
	for _, existing := range AllowedPowerManagementActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PowerManagementAction value
func (v PowerManagementAction) Ptr() *PowerManagementAction {
	return &v
}

type NullablePowerManagementAction struct {
	value *PowerManagementAction
	isSet bool
}

func (v NullablePowerManagementAction) Get() *PowerManagementAction {
	return v.value
}

func (v *NullablePowerManagementAction) Set(val *PowerManagementAction) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerManagementAction) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerManagementAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerManagementAction(val *PowerManagementAction) *NullablePowerManagementAction {
	return &NullablePowerManagementAction{value: val, isSet: true}
}

func (v NullablePowerManagementAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerManagementAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
