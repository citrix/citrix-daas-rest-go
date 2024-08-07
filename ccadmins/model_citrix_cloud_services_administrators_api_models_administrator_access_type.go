/*
Citrix.CloudServices.Administrators.Api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ccadmins

import (
	"encoding/json"
	"fmt"
)

// CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType the model 'CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType'
type CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType string

// List of Citrix.CloudServices.Administrators.ApiModels.AdministratorAccessType
const (
	CITRIXCLOUDSERVICESADMINISTRATORSAPIMODELSADMINISTRATORACCESSTYPE_FULL CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType = "Full"
	CITRIXCLOUDSERVICESADMINISTRATORSAPIMODELSADMINISTRATORACCESSTYPE_CUSTOM CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType = "Custom"
	CITRIXCLOUDSERVICESADMINISTRATORSAPIMODELSADMINISTRATORACCESSTYPE_UNDEFINED CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType = "Undefined"
)

// All allowed values of CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType enum
var AllowedCitrixCloudServicesAdministratorsApiModelsAdministratorAccessTypeEnumValues = []CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType{
	"Full",
	"Custom",
	"Undefined",
}

func (v *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType(value)
	for _, existing := range AllowedCitrixCloudServicesAdministratorsApiModelsAdministratorAccessTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType", value)
}

// NewCitrixCloudServicesAdministratorsApiModelsAdministratorAccessTypeFromValue returns a pointer to a valid CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCitrixCloudServicesAdministratorsApiModelsAdministratorAccessTypeFromValue(v string) (*CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType, error) {
	ev := CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType: valid values are %v", v, AllowedCitrixCloudServicesAdministratorsApiModelsAdministratorAccessTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType) IsValid() bool {
	for _, existing := range AllowedCitrixCloudServicesAdministratorsApiModelsAdministratorAccessTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Citrix.CloudServices.Administrators.ApiModels.AdministratorAccessType value
func (v CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType) Ptr() *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType {
	return &v
}

type NullableCitrixCloudServicesAdministratorsApiModelsAdministratorAccessType struct {
	value *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType
	isSet bool
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsAdministratorAccessType) Get() *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType {
	return v.value
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsAdministratorAccessType) Set(val *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType) {
	v.value = val
	v.isSet = true
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsAdministratorAccessType) IsSet() bool {
	return v.isSet
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsAdministratorAccessType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCitrixCloudServicesAdministratorsApiModelsAdministratorAccessType(val *CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType) *NullableCitrixCloudServicesAdministratorsApiModelsAdministratorAccessType {
	return &NullableCitrixCloudServicesAdministratorsApiModelsAdministratorAccessType{value: val, isSet: true}
}

func (v NullableCitrixCloudServicesAdministratorsApiModelsAdministratorAccessType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCitrixCloudServicesAdministratorsApiModelsAdministratorAccessType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
