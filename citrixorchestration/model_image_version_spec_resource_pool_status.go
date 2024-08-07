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

// ImageVersionSpecResourcePoolStatus Status of an image version spec within a specific resource pool.
type ImageVersionSpecResourcePoolStatus string

// List of ImageVersionSpecResourcePoolStatus
const (
	IMAGEVERSIONSPECRESOURCEPOOLSTATUS_UNKNOWN ImageVersionSpecResourcePoolStatus = "Unknown"
	IMAGEVERSIONSPECRESOURCEPOOLSTATUS_IN_PROGRESS ImageVersionSpecResourcePoolStatus = "InProgress"
	IMAGEVERSIONSPECRESOURCEPOOLSTATUS_SUCCESS ImageVersionSpecResourcePoolStatus = "Success"
	IMAGEVERSIONSPECRESOURCEPOOLSTATUS_FAILED ImageVersionSpecResourcePoolStatus = "Failed"
	IMAGEVERSIONSPECRESOURCEPOOLSTATUS_DELETING ImageVersionSpecResourcePoolStatus = "Deleting"
)

// All allowed values of ImageVersionSpecResourcePoolStatus enum
var AllowedImageVersionSpecResourcePoolStatusEnumValues = []ImageVersionSpecResourcePoolStatus{
	"Unknown",
	"InProgress",
	"Success",
	"Failed",
	"Deleting",
}

func (v *ImageVersionSpecResourcePoolStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ImageVersionSpecResourcePoolStatus(value)
	for _, existing := range AllowedImageVersionSpecResourcePoolStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ImageVersionSpecResourcePoolStatus", value)
}

// NewImageVersionSpecResourcePoolStatusFromValue returns a pointer to a valid ImageVersionSpecResourcePoolStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewImageVersionSpecResourcePoolStatusFromValue(v string) (*ImageVersionSpecResourcePoolStatus, error) {
	ev := ImageVersionSpecResourcePoolStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ImageVersionSpecResourcePoolStatus: valid values are %v", v, AllowedImageVersionSpecResourcePoolStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImageVersionSpecResourcePoolStatus) IsValid() bool {
	for _, existing := range AllowedImageVersionSpecResourcePoolStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImageVersionSpecResourcePoolStatus value
func (v ImageVersionSpecResourcePoolStatus) Ptr() *ImageVersionSpecResourcePoolStatus {
	return &v
}

type NullableImageVersionSpecResourcePoolStatus struct {
	value *ImageVersionSpecResourcePoolStatus
	isSet bool
}

func (v NullableImageVersionSpecResourcePoolStatus) Get() *ImageVersionSpecResourcePoolStatus {
	return v.value
}

func (v *NullableImageVersionSpecResourcePoolStatus) Set(val *ImageVersionSpecResourcePoolStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableImageVersionSpecResourcePoolStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableImageVersionSpecResourcePoolStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageVersionSpecResourcePoolStatus(val *ImageVersionSpecResourcePoolStatus) *NullableImageVersionSpecResourcePoolStatus {
	return &NullableImageVersionSpecResourcePoolStatus{value: val, isSet: true}
}

func (v NullableImageVersionSpecResourcePoolStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageVersionSpecResourcePoolStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
