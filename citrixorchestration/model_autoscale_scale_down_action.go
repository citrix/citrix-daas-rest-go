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

// AutoscaleScaleDownAction
type AutoscaleScaleDownAction string

// List of AutoscaleScaleDownAction
const (
	AUTOSCALESCALEDOWNACTION_SHUTDOWN AutoscaleScaleDownAction = "Shutdown"
	AUTOSCALESCALEDOWNACTION_SUSPEND  AutoscaleScaleDownAction = "Suspend"
)

// All allowed values of AutoscaleScaleDownAction enum
var AllowedAutoscaleScaleDownActionEnumValues = []AutoscaleScaleDownAction{
	"Shutdown",
	"Suspend",
}

func (v *AutoscaleScaleDownAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = AutoscaleScaleDownAction(value)
	return nil
}

// NewAutoscaleScaleDownActionFromValue returns a pointer to a valid AutoscaleScaleDownAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAutoscaleScaleDownActionFromValue(v string) (*AutoscaleScaleDownAction, error) {
	ev := AutoscaleScaleDownAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AutoscaleScaleDownAction: valid values are %v", v, AllowedAutoscaleScaleDownActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AutoscaleScaleDownAction) IsValid() bool {
	for _, existing := range AllowedAutoscaleScaleDownActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AutoscaleScaleDownAction value
func (v AutoscaleScaleDownAction) Ptr() *AutoscaleScaleDownAction {
	return &v
}

type NullableAutoscaleScaleDownAction struct {
	value *AutoscaleScaleDownAction
	isSet bool
}

func (v NullableAutoscaleScaleDownAction) Get() *AutoscaleScaleDownAction {
	return v.value
}

func (v *NullableAutoscaleScaleDownAction) Set(val *AutoscaleScaleDownAction) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoscaleScaleDownAction) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoscaleScaleDownAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoscaleScaleDownAction(val *AutoscaleScaleDownAction) *NullableAutoscaleScaleDownAction {
	return &NullableAutoscaleScaleDownAction{value: val, isSet: true}
}

func (v NullableAutoscaleScaleDownAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoscaleScaleDownAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
