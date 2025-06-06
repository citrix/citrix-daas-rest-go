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

// AwsEdcDirectorySize the model 'AwsEdcDirectorySize'
type AwsEdcDirectorySize string

// List of AwsEdcDirectorySize
const (
	AWSEDCDIRECTORYSIZE_SMALL AwsEdcDirectorySize = "SMALL"
	AWSEDCDIRECTORYSIZE_LARGE AwsEdcDirectorySize = "LARGE"
)

// All allowed values of AwsEdcDirectorySize enum
var AllowedAwsEdcDirectorySizeEnumValues = []AwsEdcDirectorySize{
	"SMALL",
	"LARGE",
}

func (v *AwsEdcDirectorySize) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AwsEdcDirectorySize(value)
	for _, existing := range AllowedAwsEdcDirectorySizeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AwsEdcDirectorySize", value)
}

// NewAwsEdcDirectorySizeFromValue returns a pointer to a valid AwsEdcDirectorySize
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAwsEdcDirectorySizeFromValue(v string) (*AwsEdcDirectorySize, error) {
	ev := AwsEdcDirectorySize(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AwsEdcDirectorySize: valid values are %v", v, AllowedAwsEdcDirectorySizeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AwsEdcDirectorySize) IsValid() bool {
	for _, existing := range AllowedAwsEdcDirectorySizeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AwsEdcDirectorySize value
func (v AwsEdcDirectorySize) Ptr() *AwsEdcDirectorySize {
	return &v
}

type NullableAwsEdcDirectorySize struct {
	value *AwsEdcDirectorySize
	isSet bool
}

func (v NullableAwsEdcDirectorySize) Get() *AwsEdcDirectorySize {
	return v.value
}

func (v *NullableAwsEdcDirectorySize) Set(val *AwsEdcDirectorySize) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsEdcDirectorySize) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsEdcDirectorySize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsEdcDirectorySize(val *AwsEdcDirectorySize) *NullableAwsEdcDirectorySize {
	return &NullableAwsEdcDirectorySize{value: val, isSet: true}
}

func (v NullableAwsEdcDirectorySize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsEdcDirectorySize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
