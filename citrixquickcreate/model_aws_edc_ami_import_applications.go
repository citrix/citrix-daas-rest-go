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

// AwsEdcAmiImportApplications Applications that can be installed on imported image
type AwsEdcAmiImportApplications string

// List of AwsEdcAmiImportApplications
const (
	AWSEDCAMIIMPORTAPPLICATIONS_MICROSOFT_OFFICE_2019 AwsEdcAmiImportApplications = "Microsoft_Office_2019"
	AWSEDCAMIIMPORTAPPLICATIONS_NONE                  AwsEdcAmiImportApplications = "None"
)

// All allowed values of AwsEdcAmiImportApplications enum
var AllowedAwsEdcAmiImportApplicationsEnumValues = []AwsEdcAmiImportApplications{
	"Microsoft_Office_2019",
	"None",
}

func (v *AwsEdcAmiImportApplications) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AwsEdcAmiImportApplications(value)
	for _, existing := range AllowedAwsEdcAmiImportApplicationsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AwsEdcAmiImportApplications", value)
}

// NewAwsEdcAmiImportApplicationsFromValue returns a pointer to a valid AwsEdcAmiImportApplications
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAwsEdcAmiImportApplicationsFromValue(v string) (*AwsEdcAmiImportApplications, error) {
	ev := AwsEdcAmiImportApplications(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AwsEdcAmiImportApplications: valid values are %v", v, AllowedAwsEdcAmiImportApplicationsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AwsEdcAmiImportApplications) IsValid() bool {
	for _, existing := range AllowedAwsEdcAmiImportApplicationsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AwsEdcAmiImportApplications value
func (v AwsEdcAmiImportApplications) Ptr() *AwsEdcAmiImportApplications {
	return &v
}

type NullableAwsEdcAmiImportApplications struct {
	value *AwsEdcAmiImportApplications
	isSet bool
}

func (v NullableAwsEdcAmiImportApplications) Get() *AwsEdcAmiImportApplications {
	return v.value
}

func (v *NullableAwsEdcAmiImportApplications) Set(val *AwsEdcAmiImportApplications) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsEdcAmiImportApplications) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsEdcAmiImportApplications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsEdcAmiImportApplications(val *AwsEdcAmiImportApplications) *NullableAwsEdcAmiImportApplications {
	return &NullableAwsEdcAmiImportApplications{value: val, isSet: true}
}

func (v NullableAwsEdcAmiImportApplications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsEdcAmiImportApplications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
