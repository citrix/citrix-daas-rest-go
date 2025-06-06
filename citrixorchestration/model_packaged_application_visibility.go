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

// PackagedApplicationVisibility Packaged Application visibility.
type PackagedApplicationVisibility string

// List of PackagedApplicationVisibility
const (
	PACKAGEDAPPLICATIONVISIBILITY_NONE                  PackagedApplicationVisibility = "None"
	PACKAGEDAPPLICATIONVISIBILITY_WORKSPACE             PackagedApplicationVisibility = "Workspace"
	PACKAGEDAPPLICATIONVISIBILITY_DESKTOP               PackagedApplicationVisibility = "Desktop"
	PACKAGEDAPPLICATIONVISIBILITY_WORKSPACE_AND_DESKTOP PackagedApplicationVisibility = "WorkspaceAndDesktop"
)

// All allowed values of PackagedApplicationVisibility enum
var AllowedPackagedApplicationVisibilityEnumValues = []PackagedApplicationVisibility{
	"None",
	"Workspace",
	"Desktop",
	"WorkspaceAndDesktop",
}

func (v *PackagedApplicationVisibility) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = PackagedApplicationVisibility(value)
	return nil
}

// NewPackagedApplicationVisibilityFromValue returns a pointer to a valid PackagedApplicationVisibility
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPackagedApplicationVisibilityFromValue(v string) (*PackagedApplicationVisibility, error) {
	ev := PackagedApplicationVisibility(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PackagedApplicationVisibility: valid values are %v", v, AllowedPackagedApplicationVisibilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PackagedApplicationVisibility) IsValid() bool {
	for _, existing := range AllowedPackagedApplicationVisibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PackagedApplicationVisibility value
func (v PackagedApplicationVisibility) Ptr() *PackagedApplicationVisibility {
	return &v
}

type NullablePackagedApplicationVisibility struct {
	value *PackagedApplicationVisibility
	isSet bool
}

func (v NullablePackagedApplicationVisibility) Get() *PackagedApplicationVisibility {
	return v.value
}

func (v *NullablePackagedApplicationVisibility) Set(val *PackagedApplicationVisibility) {
	v.value = val
	v.isSet = true
}

func (v NullablePackagedApplicationVisibility) IsSet() bool {
	return v.isSet
}

func (v *NullablePackagedApplicationVisibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackagedApplicationVisibility(val *PackagedApplicationVisibility) *NullablePackagedApplicationVisibility {
	return &NullablePackagedApplicationVisibility{value: val, isSet: true}
}

func (v NullablePackagedApplicationVisibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackagedApplicationVisibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
