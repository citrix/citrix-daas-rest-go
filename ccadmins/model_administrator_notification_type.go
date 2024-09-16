/*
Administrators APIs

APIs for managing CC administrators.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ccadmins

import (
	"encoding/json"
	"fmt"
)

// AdministratorNotificationType the model 'AdministratorNotificationType'
type AdministratorNotificationType string

// List of AdministratorNotificationType
const (
	ADMINISTRATORNOTIFICATIONTYPE_ERROR AdministratorNotificationType = "Error"
	ADMINISTRATORNOTIFICATIONTYPE_WARNING AdministratorNotificationType = "Warning"
	ADMINISTRATORNOTIFICATIONTYPE_INFORMATION AdministratorNotificationType = "Information"
	ADMINISTRATORNOTIFICATIONTYPE_CRITICAL AdministratorNotificationType = "Critical"
)

// All allowed values of AdministratorNotificationType enum
var AllowedAdministratorNotificationTypeEnumValues = []AdministratorNotificationType{
	"Error",
	"Warning",
	"Information",
	"Critical",
}

func (v *AdministratorNotificationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AdministratorNotificationType(value)
	for _, existing := range AllowedAdministratorNotificationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AdministratorNotificationType", value)
}

// NewAdministratorNotificationTypeFromValue returns a pointer to a valid AdministratorNotificationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdministratorNotificationTypeFromValue(v string) (*AdministratorNotificationType, error) {
	ev := AdministratorNotificationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdministratorNotificationType: valid values are %v", v, AllowedAdministratorNotificationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdministratorNotificationType) IsValid() bool {
	for _, existing := range AllowedAdministratorNotificationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AdministratorNotificationType value
func (v AdministratorNotificationType) Ptr() *AdministratorNotificationType {
	return &v
}

type NullableAdministratorNotificationType struct {
	value *AdministratorNotificationType
	isSet bool
}

func (v NullableAdministratorNotificationType) Get() *AdministratorNotificationType {
	return v.value
}

func (v *NullableAdministratorNotificationType) Set(val *AdministratorNotificationType) {
	v.value = val
	v.isSet = true
}

func (v NullableAdministratorNotificationType) IsSet() bool {
	return v.isSet
}

func (v *NullableAdministratorNotificationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdministratorNotificationType(val *AdministratorNotificationType) *NullableAdministratorNotificationType {
	return &NullableAdministratorNotificationType{value: val, isSet: true}
}

func (v NullableAdministratorNotificationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdministratorNotificationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
