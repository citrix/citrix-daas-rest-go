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

// AwsEdcAmiImageType the model 'AwsEdcAmiImageType'
type AwsEdcAmiImageType string

// List of AwsEdcAmiImageType
const (
	AWSEDCAMIIMAGETYPE_KERNEL  AwsEdcAmiImageType = "KERNEL"
	AWSEDCAMIIMAGETYPE_MACHINE AwsEdcAmiImageType = "MACHINE"
	AWSEDCAMIIMAGETYPE_RAMDISK AwsEdcAmiImageType = "RAMDISK"
)

// All allowed values of AwsEdcAmiImageType enum
var AllowedAwsEdcAmiImageTypeEnumValues = []AwsEdcAmiImageType{
	"KERNEL",
	"MACHINE",
	"RAMDISK",
}

func (v *AwsEdcAmiImageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AwsEdcAmiImageType(value)
	for _, existing := range AllowedAwsEdcAmiImageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AwsEdcAmiImageType", value)
}

// NewAwsEdcAmiImageTypeFromValue returns a pointer to a valid AwsEdcAmiImageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAwsEdcAmiImageTypeFromValue(v string) (*AwsEdcAmiImageType, error) {
	ev := AwsEdcAmiImageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AwsEdcAmiImageType: valid values are %v", v, AllowedAwsEdcAmiImageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AwsEdcAmiImageType) IsValid() bool {
	for _, existing := range AllowedAwsEdcAmiImageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AwsEdcAmiImageType value
func (v AwsEdcAmiImageType) Ptr() *AwsEdcAmiImageType {
	return &v
}

type NullableAwsEdcAmiImageType struct {
	value *AwsEdcAmiImageType
	isSet bool
}

func (v NullableAwsEdcAmiImageType) Get() *AwsEdcAmiImageType {
	return v.value
}

func (v *NullableAwsEdcAmiImageType) Set(val *AwsEdcAmiImageType) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsEdcAmiImageType) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsEdcAmiImageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsEdcAmiImageType(val *AwsEdcAmiImageType) *NullableAwsEdcAmiImageType {
	return &NullableAwsEdcAmiImageType{value: val, isSet: true}
}

func (v NullableAwsEdcAmiImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsEdcAmiImageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
