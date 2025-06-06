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

// LicenseServerStatus
type LicenseServerStatus string

// List of LicenseServerStatus
const (
	LICENSESERVERSTATUS_COMPATIBLE              LicenseServerStatus = "Compatible"
	LICENSESERVERSTATUS_INCOMPATIBLE            LicenseServerStatus = "Incompatible"
	LICENSESERVERSTATUS_INACCESSIBLE            LicenseServerStatus = "Inaccessible"
	LICENSESERVERSTATUS_INTERNAL_ERROR          LicenseServerStatus = "InternalError"
	LICENSESERVERSTATUS_LS_VERSION_INCOMPATIBLE LicenseServerStatus = "LSVersionIncompatible"
	LICENSESERVERSTATUS_CLOUD_NOT_REGISTERED    LicenseServerStatus = "CloudNotRegistered"
	LICENSESERVERSTATUS_LAS_NOT_ONBOARDED       LicenseServerStatus = "LasNotOnboarded"
	LICENSESERVERSTATUS_INTERNET_NOT_CONNECTED  LicenseServerStatus = "InternetNotConnected"
	LICENSESERVERSTATUS_LAS_COMPATIBLE          LicenseServerStatus = "LasCompatible"
	LICENSESERVERSTATUS_NO_INVENTORY            LicenseServerStatus = "NoInventory"
	LICENSESERVERSTATUS_UNKNOWN_STATUS          LicenseServerStatus = "UnknownStatus"
)

// All allowed values of LicenseServerStatus enum
var AllowedLicenseServerStatusEnumValues = []LicenseServerStatus{
	"Compatible",
	"Incompatible",
	"Inaccessible",
	"InternalError",
	"LSVersionIncompatible",
	"CloudNotRegistered",
	"LasNotOnboarded",
	"InternetNotConnected",
	"LasCompatible",
	"NoInventory",
	"UnknownStatus",
}

func (v *LicenseServerStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = LicenseServerStatus(value)
	return nil
}

// NewLicenseServerStatusFromValue returns a pointer to a valid LicenseServerStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLicenseServerStatusFromValue(v string) (*LicenseServerStatus, error) {
	ev := LicenseServerStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LicenseServerStatus: valid values are %v", v, AllowedLicenseServerStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LicenseServerStatus) IsValid() bool {
	for _, existing := range AllowedLicenseServerStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LicenseServerStatus value
func (v LicenseServerStatus) Ptr() *LicenseServerStatus {
	return &v
}

type NullableLicenseServerStatus struct {
	value *LicenseServerStatus
	isSet bool
}

func (v NullableLicenseServerStatus) Get() *LicenseServerStatus {
	return v.value
}

func (v *NullableLicenseServerStatus) Set(val *LicenseServerStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseServerStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseServerStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseServerStatus(val *LicenseServerStatus) *NullableLicenseServerStatus {
	return &NullableLicenseServerStatus{value: val, isSet: true}
}

func (v NullableLicenseServerStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseServerStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
