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

// BackupRestoreRestoreTypes Restore Types
type BackupRestoreRestoreTypes string

// List of BackupRestoreRestoreTypes
const (
	BACKUPRESTORERESTORETYPES_UNDEFINED                          BackupRestoreRestoreTypes = "Undefined"
	BACKUPRESTORERESTORETYPES_MISSING_ONLY                       BackupRestoreRestoreTypes = "MissingOnly"
	BACKUPRESTORERESTORETYPES_MISSING_UPDATE_EXISTING            BackupRestoreRestoreTypes = "MissingUpdateExisting"
	BACKUPRESTORERESTORETYPES_MISSING_UPDATE_EXISTING_REMOVE_NEW BackupRestoreRestoreTypes = "MissingUpdateExistingRemoveNew"
)

// All allowed values of BackupRestoreRestoreTypes enum
var AllowedBackupRestoreRestoreTypesEnumValues = []BackupRestoreRestoreTypes{
	"Undefined",
	"MissingOnly",
	"MissingUpdateExisting",
	"MissingUpdateExistingRemoveNew",
}

func (v *BackupRestoreRestoreTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = BackupRestoreRestoreTypes(value)
	return nil
}

// NewBackupRestoreRestoreTypesFromValue returns a pointer to a valid BackupRestoreRestoreTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBackupRestoreRestoreTypesFromValue(v string) (*BackupRestoreRestoreTypes, error) {
	ev := BackupRestoreRestoreTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BackupRestoreRestoreTypes: valid values are %v", v, AllowedBackupRestoreRestoreTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BackupRestoreRestoreTypes) IsValid() bool {
	for _, existing := range AllowedBackupRestoreRestoreTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackupRestoreRestoreTypes value
func (v BackupRestoreRestoreTypes) Ptr() *BackupRestoreRestoreTypes {
	return &v
}

type NullableBackupRestoreRestoreTypes struct {
	value *BackupRestoreRestoreTypes
	isSet bool
}

func (v NullableBackupRestoreRestoreTypes) Get() *BackupRestoreRestoreTypes {
	return v.value
}

func (v *NullableBackupRestoreRestoreTypes) Set(val *BackupRestoreRestoreTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupRestoreRestoreTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupRestoreRestoreTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupRestoreRestoreTypes(val *BackupRestoreRestoreTypes) *NullableBackupRestoreRestoreTypes {
	return &NullableBackupRestoreRestoreTypes{value: val, isSet: true}
}

func (v NullableBackupRestoreRestoreTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupRestoreRestoreTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
