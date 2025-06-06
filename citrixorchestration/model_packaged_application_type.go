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

// PackagedApplicationType Packaged Application types.
type PackagedApplicationType string

// List of PackagedApplicationType
const (
	PACKAGEDAPPLICATIONTYPE_NOT_APPLICABLE     PackagedApplicationType = "NotApplicable"
	PACKAGEDAPPLICATIONTYPE_APP_V_SINGLE_ADMIN PackagedApplicationType = "AppVSingleAdmin"
	PACKAGEDAPPLICATIONTYPE_APP_V_DUAL_ADMIN   PackagedApplicationType = "AppVDualAdmin"
	PACKAGEDAPPLICATIONTYPE_MSIX               PackagedApplicationType = "Msix"
	PACKAGEDAPPLICATIONTYPE_APP_ATTACH         PackagedApplicationType = "AppAttach"
	PACKAGEDAPPLICATIONTYPE_FLEX_APP           PackagedApplicationType = "FlexApp"
	PACKAGEDAPPLICATIONTYPE_ELASTIC_APP_LAYER  PackagedApplicationType = "ElasticAppLayer"
)

// All allowed values of PackagedApplicationType enum
var AllowedPackagedApplicationTypeEnumValues = []PackagedApplicationType{
	"NotApplicable",
	"AppVSingleAdmin",
	"AppVDualAdmin",
	"Msix",
	"AppAttach",
	"FlexApp",
	"ElasticAppLayer",
}

func (v *PackagedApplicationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = PackagedApplicationType(value)
	return nil
}

// NewPackagedApplicationTypeFromValue returns a pointer to a valid PackagedApplicationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPackagedApplicationTypeFromValue(v string) (*PackagedApplicationType, error) {
	ev := PackagedApplicationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PackagedApplicationType: valid values are %v", v, AllowedPackagedApplicationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PackagedApplicationType) IsValid() bool {
	for _, existing := range AllowedPackagedApplicationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PackagedApplicationType value
func (v PackagedApplicationType) Ptr() *PackagedApplicationType {
	return &v
}

type NullablePackagedApplicationType struct {
	value *PackagedApplicationType
	isSet bool
}

func (v NullablePackagedApplicationType) Get() *PackagedApplicationType {
	return v.value
}

func (v *NullablePackagedApplicationType) Set(val *PackagedApplicationType) {
	v.value = val
	v.isSet = true
}

func (v NullablePackagedApplicationType) IsSet() bool {
	return v.isSet
}

func (v *NullablePackagedApplicationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackagedApplicationType(val *PackagedApplicationType) *NullablePackagedApplicationType {
	return &NullablePackagedApplicationType{value: val, isSet: true}
}

func (v NullablePackagedApplicationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackagedApplicationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
