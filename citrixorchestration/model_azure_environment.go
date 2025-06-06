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

// AzureEnvironment Azure environments
type AzureEnvironment string

// List of AzureEnvironment
const (
	AZUREENVIRONMENT_CUSTOM              AzureEnvironment = "Custom"
	AZUREENVIRONMENT_AZURE_CLOUD         AzureEnvironment = "AzureCloud"
	AZUREENVIRONMENT_AZURE_CHINA_CLOUD   AzureEnvironment = "AzureChinaCloud"
	AZUREENVIRONMENT_AZURE_US_GOVERNMENT AzureEnvironment = "AzureUSGovernment"
	AZUREENVIRONMENT_AZURE_GERMAN_CLOUD  AzureEnvironment = "AzureGermanCloud"
)

// All allowed values of AzureEnvironment enum
var AllowedAzureEnvironmentEnumValues = []AzureEnvironment{
	"Custom",
	"AzureCloud",
	"AzureChinaCloud",
	"AzureUSGovernment",
	"AzureGermanCloud",
}

func (v *AzureEnvironment) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = AzureEnvironment(value)
	return nil
}

// NewAzureEnvironmentFromValue returns a pointer to a valid AzureEnvironment
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAzureEnvironmentFromValue(v string) (*AzureEnvironment, error) {
	ev := AzureEnvironment(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AzureEnvironment: valid values are %v", v, AllowedAzureEnvironmentEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AzureEnvironment) IsValid() bool {
	for _, existing := range AllowedAzureEnvironmentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AzureEnvironment value
func (v AzureEnvironment) Ptr() *AzureEnvironment {
	return &v
}

type NullableAzureEnvironment struct {
	value *AzureEnvironment
	isSet bool
}

func (v NullableAzureEnvironment) Get() *AzureEnvironment {
	return v.value
}

func (v *NullableAzureEnvironment) Set(val *AzureEnvironment) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureEnvironment(val *AzureEnvironment) *NullableAzureEnvironment {
	return &NullableAzureEnvironment{value: val, isSet: true}
}

func (v NullableAzureEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
