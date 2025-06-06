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

// ScriptOperationType the model 'ScriptOperationType'
type ScriptOperationType string

// List of ScriptOperationType
const (
	SCRIPTOPERATIONTYPE_NORMAL   ScriptOperationType = "Normal"
	SCRIPTOPERATIONTYPE_ADVANCED ScriptOperationType = "Advanced"
)

// All allowed values of ScriptOperationType enum
var AllowedScriptOperationTypeEnumValues = []ScriptOperationType{
	"Normal",
	"Advanced",
}

func (v *ScriptOperationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = ScriptOperationType(value)
	return nil
}

// NewScriptOperationTypeFromValue returns a pointer to a valid ScriptOperationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScriptOperationTypeFromValue(v string) (*ScriptOperationType, error) {
	ev := ScriptOperationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ScriptOperationType: valid values are %v", v, AllowedScriptOperationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScriptOperationType) IsValid() bool {
	for _, existing := range AllowedScriptOperationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScriptOperationType value
func (v ScriptOperationType) Ptr() *ScriptOperationType {
	return &v
}

type NullableScriptOperationType struct {
	value *ScriptOperationType
	isSet bool
}

func (v NullableScriptOperationType) Get() *ScriptOperationType {
	return v.value
}

func (v *NullableScriptOperationType) Set(val *ScriptOperationType) {
	v.value = val
	v.isSet = true
}

func (v NullableScriptOperationType) IsSet() bool {
	return v.isSet
}

func (v *NullableScriptOperationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScriptOperationType(val *ScriptOperationType) *NullableScriptOperationType {
	return &NullableScriptOperationType{value: val, isSet: true}
}

func (v NullableScriptOperationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScriptOperationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
