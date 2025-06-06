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

// BackupRestoreActionStartStatus Action start status
type BackupRestoreActionStartStatus string

// List of BackupRestoreActionStartStatus
const (
	BACKUPRESTOREACTIONSTARTSTATUS_FAILED                       BackupRestoreActionStartStatus = "Failed"
	BACKUPRESTOREACTIONSTARTSTATUS_SUCCESS                      BackupRestoreActionStartStatus = "Success"
	BACKUPRESTOREACTIONSTARTSTATUS_EXISTING_ACTION_IN_PROGRESS  BackupRestoreActionStartStatus = "ExistingActionInProgress"
	BACKUPRESTOREACTIONSTARTSTATUS_EXISTING_BACKUP_IN_PROGRESS  BackupRestoreActionStartStatus = "ExistingBackupInProgress"
	BACKUPRESTOREACTIONSTARTSTATUS_EXISTING_RESTORE_IN_PROGRESS BackupRestoreActionStartStatus = "ExistingRestoreInProgress"
	BACKUPRESTOREACTIONSTARTSTATUS_STORAGE_NOT_VALID            BackupRestoreActionStartStatus = "StorageNotValid"
)

// All allowed values of BackupRestoreActionStartStatus enum
var AllowedBackupRestoreActionStartStatusEnumValues = []BackupRestoreActionStartStatus{
	"Failed",
	"Success",
	"ExistingActionInProgress",
	"ExistingBackupInProgress",
	"ExistingRestoreInProgress",
	"StorageNotValid",
}

func (v *BackupRestoreActionStartStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = BackupRestoreActionStartStatus(value)
	return nil
}

// NewBackupRestoreActionStartStatusFromValue returns a pointer to a valid BackupRestoreActionStartStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBackupRestoreActionStartStatusFromValue(v string) (*BackupRestoreActionStartStatus, error) {
	ev := BackupRestoreActionStartStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BackupRestoreActionStartStatus: valid values are %v", v, AllowedBackupRestoreActionStartStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BackupRestoreActionStartStatus) IsValid() bool {
	for _, existing := range AllowedBackupRestoreActionStartStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackupRestoreActionStartStatus value
func (v BackupRestoreActionStartStatus) Ptr() *BackupRestoreActionStartStatus {
	return &v
}

type NullableBackupRestoreActionStartStatus struct {
	value *BackupRestoreActionStartStatus
	isSet bool
}

func (v NullableBackupRestoreActionStartStatus) Get() *BackupRestoreActionStartStatus {
	return v.value
}

func (v *NullableBackupRestoreActionStartStatus) Set(val *BackupRestoreActionStartStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupRestoreActionStartStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupRestoreActionStartStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupRestoreActionStartStatus(val *BackupRestoreActionStartStatus) *NullableBackupRestoreActionStartStatus {
	return &NullableBackupRestoreActionStartStatus{value: val, isSet: true}
}

func (v NullableBackupRestoreActionStartStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupRestoreActionStartStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
