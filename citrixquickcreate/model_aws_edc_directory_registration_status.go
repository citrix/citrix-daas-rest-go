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

// AwsEdcDirectoryRegistrationStatus the model 'AwsEdcDirectoryRegistrationStatus'
type AwsEdcDirectoryRegistrationStatus string

// List of AwsEdcDirectoryRegistrationStatus
const (
	AWSEDCDIRECTORYREGISTRATIONSTATUS_REGISTERED    AwsEdcDirectoryRegistrationStatus = "REGISTERED"
	AWSEDCDIRECTORYREGISTRATIONSTATUS_REGISTERING   AwsEdcDirectoryRegistrationStatus = "REGISTERING"
	AWSEDCDIRECTORYREGISTRATIONSTATUS_DEREGISTERED  AwsEdcDirectoryRegistrationStatus = "DEREGISTERED"
	AWSEDCDIRECTORYREGISTRATIONSTATUS_DEREGISTERING AwsEdcDirectoryRegistrationStatus = "DEREGISTERING"
	AWSEDCDIRECTORYREGISTRATIONSTATUS_ERROR         AwsEdcDirectoryRegistrationStatus = "ERROR"
)

// All allowed values of AwsEdcDirectoryRegistrationStatus enum
var AllowedAwsEdcDirectoryRegistrationStatusEnumValues = []AwsEdcDirectoryRegistrationStatus{
	"REGISTERED",
	"REGISTERING",
	"DEREGISTERED",
	"DEREGISTERING",
	"ERROR",
}

func (v *AwsEdcDirectoryRegistrationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AwsEdcDirectoryRegistrationStatus(value)
	for _, existing := range AllowedAwsEdcDirectoryRegistrationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AwsEdcDirectoryRegistrationStatus", value)
}

// NewAwsEdcDirectoryRegistrationStatusFromValue returns a pointer to a valid AwsEdcDirectoryRegistrationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAwsEdcDirectoryRegistrationStatusFromValue(v string) (*AwsEdcDirectoryRegistrationStatus, error) {
	ev := AwsEdcDirectoryRegistrationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AwsEdcDirectoryRegistrationStatus: valid values are %v", v, AllowedAwsEdcDirectoryRegistrationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AwsEdcDirectoryRegistrationStatus) IsValid() bool {
	for _, existing := range AllowedAwsEdcDirectoryRegistrationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AwsEdcDirectoryRegistrationStatus value
func (v AwsEdcDirectoryRegistrationStatus) Ptr() *AwsEdcDirectoryRegistrationStatus {
	return &v
}

type NullableAwsEdcDirectoryRegistrationStatus struct {
	value *AwsEdcDirectoryRegistrationStatus
	isSet bool
}

func (v NullableAwsEdcDirectoryRegistrationStatus) Get() *AwsEdcDirectoryRegistrationStatus {
	return v.value
}

func (v *NullableAwsEdcDirectoryRegistrationStatus) Set(val *AwsEdcDirectoryRegistrationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsEdcDirectoryRegistrationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsEdcDirectoryRegistrationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsEdcDirectoryRegistrationStatus(val *AwsEdcDirectoryRegistrationStatus) *NullableAwsEdcDirectoryRegistrationStatus {
	return &NullableAwsEdcDirectoryRegistrationStatus{value: val, isSet: true}
}

func (v NullableAwsEdcDirectoryRegistrationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsEdcDirectoryRegistrationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
