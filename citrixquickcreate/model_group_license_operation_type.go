/*
Quick Create Service for DaaS Quick Deploy with AWS WorkSpaces Core - Tech Preview

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: DaaSTechPreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickcreate

import (
	"encoding/json"
	"fmt"
)

// GroupLicenseOperationType Group license operation type                DO NOT RENAME the enum values!  They are stored in DB as strings.
type GroupLicenseOperationType string

// List of GroupLicenseOperationType
const (
	GROUPLICENSEOPERATIONTYPE_ASSIGNMENT GroupLicenseOperationType = "Assignment"
	GROUPLICENSEOPERATIONTYPE_UNASSIGNMENT GroupLicenseOperationType = "Unassignment"
)

// All allowed values of GroupLicenseOperationType enum
var AllowedGroupLicenseOperationTypeEnumValues = []GroupLicenseOperationType{
	"Assignment",
	"Unassignment",
}

func (v *GroupLicenseOperationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GroupLicenseOperationType(value)
	for _, existing := range AllowedGroupLicenseOperationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GroupLicenseOperationType", value)
}

// NewGroupLicenseOperationTypeFromValue returns a pointer to a valid GroupLicenseOperationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGroupLicenseOperationTypeFromValue(v string) (*GroupLicenseOperationType, error) {
	ev := GroupLicenseOperationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GroupLicenseOperationType: valid values are %v", v, AllowedGroupLicenseOperationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GroupLicenseOperationType) IsValid() bool {
	for _, existing := range AllowedGroupLicenseOperationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GroupLicenseOperationType value
func (v GroupLicenseOperationType) Ptr() *GroupLicenseOperationType {
	return &v
}

type NullableGroupLicenseOperationType struct {
	value *GroupLicenseOperationType
	isSet bool
}

func (v NullableGroupLicenseOperationType) Get() *GroupLicenseOperationType {
	return v.value
}

func (v *NullableGroupLicenseOperationType) Set(val *GroupLicenseOperationType) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupLicenseOperationType) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupLicenseOperationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupLicenseOperationType(val *GroupLicenseOperationType) *NullableGroupLicenseOperationType {
	return &NullableGroupLicenseOperationType{value: val, isSet: true}
}

func (v NullableGroupLicenseOperationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupLicenseOperationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
