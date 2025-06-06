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

// SnapshotType the model 'SnapshotType'
type SnapshotType string

// List of SnapshotType
const (
	SNAPSHOTTYPE_ANYTIME SnapshotType = "Anytime"
	SNAPSHOTTYPE_DAILY   SnapshotType = "Daily"
	SNAPSHOTTYPE_WEEKLY  SnapshotType = "Weekly"
)

// All allowed values of SnapshotType enum
var AllowedSnapshotTypeEnumValues = []SnapshotType{
	"Anytime",
	"Daily",
	"Weekly",
}

func (v *SnapshotType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = SnapshotType(value)
	return nil
}

// NewSnapshotTypeFromValue returns a pointer to a valid SnapshotType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSnapshotTypeFromValue(v string) (*SnapshotType, error) {
	ev := SnapshotType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SnapshotType: valid values are %v", v, AllowedSnapshotTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SnapshotType) IsValid() bool {
	for _, existing := range AllowedSnapshotTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SnapshotType value
func (v SnapshotType) Ptr() *SnapshotType {
	return &v
}

type NullableSnapshotType struct {
	value *SnapshotType
	isSet bool
}

func (v NullableSnapshotType) Get() *SnapshotType {
	return v.value
}

func (v *NullableSnapshotType) Set(val *SnapshotType) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotType) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotType(val *SnapshotType) *NullableSnapshotType {
	return &NullableSnapshotType{value: val, isSet: true}
}

func (v NullableSnapshotType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
