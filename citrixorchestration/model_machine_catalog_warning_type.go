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

// MachineCatalogWarningType Types of warnings that may occur on a machine catalog.
type MachineCatalogWarningType string

// List of MachineCatalogWarningType
const (
	MACHINECATALOGWARNINGTYPE_UNKNOWN                                       MachineCatalogWarningType = "Unknown"
	MACHINECATALOGWARNINGTYPE_RDS_CATALOG_LICENSE_CHECK                     MachineCatalogWarningType = "RdsCatalogLicenseCheck"
	MACHINECATALOGWARNINGTYPE_IMAGE_PREPARATION_OFFICE_REARM_COUNT_EXCEEDED MachineCatalogWarningType = "ImagePreparationOfficeRearmCountExceeded"
	MACHINECATALOGWARNINGTYPE_IMAGE_PREPARATION_OFFICE_REARM_FAILED         MachineCatalogWarningType = "ImagePreparationOfficeRearmFailed"
	MACHINECATALOGWARNINGTYPE_IMAGE_PREPARATION_OS_REARM_COUNT_EXCEEDED     MachineCatalogWarningType = "ImagePreparationOSRearmCountExceeded"
	MACHINECATALOGWARNINGTYPE_IMAGE_PREPARATION_OS_REARM_FAILED             MachineCatalogWarningType = "ImagePreparationOSRearmFailed"
	MACHINECATALOGWARNINGTYPE_CATALOG_ACTION                                MachineCatalogWarningType = "CatalogAction"
	MACHINECATALOGWARNINGTYPE_GET_CUSTOM_PROPERTIES_FAILED                  MachineCatalogWarningType = "GetCustomPropertiesFailed"
	MACHINECATALOGWARNINGTYPE_GET_VM_METADATA_FAILED                        MachineCatalogWarningType = "GetVMMetadataFailed"
)

// All allowed values of MachineCatalogWarningType enum
var AllowedMachineCatalogWarningTypeEnumValues = []MachineCatalogWarningType{
	"Unknown",
	"RdsCatalogLicenseCheck",
	"ImagePreparationOfficeRearmCountExceeded",
	"ImagePreparationOfficeRearmFailed",
	"ImagePreparationOSRearmCountExceeded",
	"ImagePreparationOSRearmFailed",
	"CatalogAction",
	"GetCustomPropertiesFailed",
	"GetVMMetadataFailed",
}

func (v *MachineCatalogWarningType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = MachineCatalogWarningType(value)
	return nil
}

// NewMachineCatalogWarningTypeFromValue returns a pointer to a valid MachineCatalogWarningType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMachineCatalogWarningTypeFromValue(v string) (*MachineCatalogWarningType, error) {
	ev := MachineCatalogWarningType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MachineCatalogWarningType: valid values are %v", v, AllowedMachineCatalogWarningTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MachineCatalogWarningType) IsValid() bool {
	for _, existing := range AllowedMachineCatalogWarningTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MachineCatalogWarningType value
func (v MachineCatalogWarningType) Ptr() *MachineCatalogWarningType {
	return &v
}

type NullableMachineCatalogWarningType struct {
	value *MachineCatalogWarningType
	isSet bool
}

func (v NullableMachineCatalogWarningType) Get() *MachineCatalogWarningType {
	return v.value
}

func (v *NullableMachineCatalogWarningType) Set(val *MachineCatalogWarningType) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineCatalogWarningType) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineCatalogWarningType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineCatalogWarningType(val *MachineCatalogWarningType) *NullableMachineCatalogWarningType {
	return &NullableMachineCatalogWarningType{value: val, isSet: true}
}

func (v NullableMachineCatalogWarningType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineCatalogWarningType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
