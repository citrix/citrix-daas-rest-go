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

// DiskStorageType The disk storage type
type DiskStorageType string

// List of DiskStorageType
const (
	DISKSTORAGETYPE_UNKNOWN          DiskStorageType = "Unknown"
	DISKSTORAGETYPE_PREMIUM_LRS      DiskStorageType = "Premium_LRS"
	DISKSTORAGETYPE_STANDARD_LRS     DiskStorageType = "Standard_LRS"
	DISKSTORAGETYPE_STANDARD_SSD_LRS DiskStorageType = "StandardSSD_LRS"
	DISKSTORAGETYPE_PREMIUM_ZRS      DiskStorageType = "Premium_ZRS"
	DISKSTORAGETYPE_STANDARD_SSD_ZRS DiskStorageType = "StandardSSD_ZRS"
	DISKSTORAGETYPE_PREMIUM_V2_LRS   DiskStorageType = "PremiumV2_LRS"
	DISKSTORAGETYPE_ULTRA_SSD_LRS    DiskStorageType = "UltraSSD_LRS"
)

// All allowed values of DiskStorageType enum
var AllowedDiskStorageTypeEnumValues = []DiskStorageType{
	"Unknown",
	"Premium_LRS",
	"Standard_LRS",
	"StandardSSD_LRS",
	"Premium_ZRS",
	"StandardSSD_ZRS",
	"PremiumV2_LRS",
	"UltraSSD_LRS",
}

func (v *DiskStorageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = DiskStorageType(value)
	return nil
}

// NewDiskStorageTypeFromValue returns a pointer to a valid DiskStorageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDiskStorageTypeFromValue(v string) (*DiskStorageType, error) {
	ev := DiskStorageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DiskStorageType: valid values are %v", v, AllowedDiskStorageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DiskStorageType) IsValid() bool {
	for _, existing := range AllowedDiskStorageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiskStorageType value
func (v DiskStorageType) Ptr() *DiskStorageType {
	return &v
}

type NullableDiskStorageType struct {
	value *DiskStorageType
	isSet bool
}

func (v NullableDiskStorageType) Get() *DiskStorageType {
	return v.value
}

func (v *NullableDiskStorageType) Set(val *DiskStorageType) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskStorageType) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskStorageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskStorageType(val *DiskStorageType) *NullableDiskStorageType {
	return &NullableDiskStorageType{value: val, isSet: true}
}

func (v NullableDiskStorageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskStorageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
