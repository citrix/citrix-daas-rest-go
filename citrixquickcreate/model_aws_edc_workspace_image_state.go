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

// AwsEdcWorkspaceImageState the model 'AwsEdcWorkspaceImageState'
type AwsEdcWorkspaceImageState string

// List of AwsEdcWorkspaceImageState
const (
	AWSEDCWORKSPACEIMAGESTATE_AVAILABLE             AwsEdcWorkspaceImageState = "AVAILABLE"
	AWSEDCWORKSPACEIMAGESTATE_ERROR                 AwsEdcWorkspaceImageState = "ERROR"
	AWSEDCWORKSPACEIMAGESTATE_PENDING               AwsEdcWorkspaceImageState = "PENDING"
	AWSEDCWORKSPACEIMAGESTATE_ERROR_INVALID_ACCOUNT AwsEdcWorkspaceImageState = "ERROR_INVALID_ACCOUNT"
)

// All allowed values of AwsEdcWorkspaceImageState enum
var AllowedAwsEdcWorkspaceImageStateEnumValues = []AwsEdcWorkspaceImageState{
	"AVAILABLE",
	"ERROR",
	"PENDING",
	"ERROR_INVALID_ACCOUNT",
}

func (v *AwsEdcWorkspaceImageState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AwsEdcWorkspaceImageState(value)
	for _, existing := range AllowedAwsEdcWorkspaceImageStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AwsEdcWorkspaceImageState", value)
}

// NewAwsEdcWorkspaceImageStateFromValue returns a pointer to a valid AwsEdcWorkspaceImageState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAwsEdcWorkspaceImageStateFromValue(v string) (*AwsEdcWorkspaceImageState, error) {
	ev := AwsEdcWorkspaceImageState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AwsEdcWorkspaceImageState: valid values are %v", v, AllowedAwsEdcWorkspaceImageStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AwsEdcWorkspaceImageState) IsValid() bool {
	for _, existing := range AllowedAwsEdcWorkspaceImageStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AwsEdcWorkspaceImageState value
func (v AwsEdcWorkspaceImageState) Ptr() *AwsEdcWorkspaceImageState {
	return &v
}

type NullableAwsEdcWorkspaceImageState struct {
	value *AwsEdcWorkspaceImageState
	isSet bool
}

func (v NullableAwsEdcWorkspaceImageState) Get() *AwsEdcWorkspaceImageState {
	return v.value
}

func (v *NullableAwsEdcWorkspaceImageState) Set(val *AwsEdcWorkspaceImageState) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsEdcWorkspaceImageState) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsEdcWorkspaceImageState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsEdcWorkspaceImageState(val *AwsEdcWorkspaceImageState) *NullableAwsEdcWorkspaceImageState {
	return &NullableAwsEdcWorkspaceImageState{value: val, isSet: true}
}

func (v NullableAwsEdcWorkspaceImageState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsEdcWorkspaceImageState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
