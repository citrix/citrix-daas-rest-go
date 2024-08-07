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

// TenancyType Indicates whether to use tenancy type Shared, Instance or Host when creating machines in Cloud Hypervisors.
type TenancyType string

// List of TenancyType
const (
	TENANCYTYPE_UNKNOWN TenancyType = "Unknown"
	TENANCYTYPE_SHARED TenancyType = "Shared"
	TENANCYTYPE_INSTANCE TenancyType = "Instance"
	TENANCYTYPE_HOST TenancyType = "Host"
)

// All allowed values of TenancyType enum
var AllowedTenancyTypeEnumValues = []TenancyType{
	"Unknown",
	"Shared",
	"Instance",
	"Host",
}

func (v *TenancyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TenancyType(value)
	for _, existing := range AllowedTenancyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TenancyType", value)
}

// NewTenancyTypeFromValue returns a pointer to a valid TenancyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTenancyTypeFromValue(v string) (*TenancyType, error) {
	ev := TenancyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TenancyType: valid values are %v", v, AllowedTenancyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TenancyType) IsValid() bool {
	for _, existing := range AllowedTenancyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TenancyType value
func (v TenancyType) Ptr() *TenancyType {
	return &v
}

type NullableTenancyType struct {
	value *TenancyType
	isSet bool
}

func (v NullableTenancyType) Get() *TenancyType {
	return v.value
}

func (v *NullableTenancyType) Set(val *TenancyType) {
	v.value = val
	v.isSet = true
}

func (v NullableTenancyType) IsSet() bool {
	return v.isSet
}

func (v *NullableTenancyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenancyType(val *TenancyType) *NullableTenancyType {
	return &NullableTenancyType{value: val, isSet: true}
}

func (v NullableTenancyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenancyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
