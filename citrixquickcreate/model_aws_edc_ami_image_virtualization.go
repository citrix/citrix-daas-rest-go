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

// AwsEdcAmiImageVirtualization the model 'AwsEdcAmiImageVirtualization'
type AwsEdcAmiImageVirtualization string

// List of AwsEdcAmiImageVirtualization
const (
	AWSEDCAMIIMAGEVIRTUALIZATION_HVM AwsEdcAmiImageVirtualization = "HVM"
	AWSEDCAMIIMAGEVIRTUALIZATION_PARAVIRTUAL AwsEdcAmiImageVirtualization = "PARAVIRTUAL"
)

// All allowed values of AwsEdcAmiImageVirtualization enum
var AllowedAwsEdcAmiImageVirtualizationEnumValues = []AwsEdcAmiImageVirtualization{
	"HVM",
	"PARAVIRTUAL",
}

func (v *AwsEdcAmiImageVirtualization) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AwsEdcAmiImageVirtualization(value)
	for _, existing := range AllowedAwsEdcAmiImageVirtualizationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AwsEdcAmiImageVirtualization", value)
}

// NewAwsEdcAmiImageVirtualizationFromValue returns a pointer to a valid AwsEdcAmiImageVirtualization
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAwsEdcAmiImageVirtualizationFromValue(v string) (*AwsEdcAmiImageVirtualization, error) {
	ev := AwsEdcAmiImageVirtualization(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AwsEdcAmiImageVirtualization: valid values are %v", v, AllowedAwsEdcAmiImageVirtualizationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AwsEdcAmiImageVirtualization) IsValid() bool {
	for _, existing := range AllowedAwsEdcAmiImageVirtualizationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AwsEdcAmiImageVirtualization value
func (v AwsEdcAmiImageVirtualization) Ptr() *AwsEdcAmiImageVirtualization {
	return &v
}

type NullableAwsEdcAmiImageVirtualization struct {
	value *AwsEdcAmiImageVirtualization
	isSet bool
}

func (v NullableAwsEdcAmiImageVirtualization) Get() *AwsEdcAmiImageVirtualization {
	return v.value
}

func (v *NullableAwsEdcAmiImageVirtualization) Set(val *AwsEdcAmiImageVirtualization) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsEdcAmiImageVirtualization) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsEdcAmiImageVirtualization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsEdcAmiImageVirtualization(val *AwsEdcAmiImageVirtualization) *NullableAwsEdcAmiImageVirtualization {
	return &NullableAwsEdcAmiImageVirtualization{value: val, isSet: true}
}

func (v NullableAwsEdcAmiImageVirtualization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsEdcAmiImageVirtualization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
