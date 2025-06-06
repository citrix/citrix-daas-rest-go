/*
Citrix Virtual App & Desktop Catalog Service 147.0.26651.57932

Catalog Service

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickdeploy

import (
	"encoding/json"
	"fmt"
)

// ResourceLocationJobType Type of job being performed on the Resource Location
type ResourceLocationJobType string

// List of ResourceLocationJobType
const (
	RESOURCELOCATIONJOBTYPE_ADD_CONNECTORS_JOB           ResourceLocationJobType = "AddConnectorsJob"
	RESOURCELOCATIONJOBTYPE_RESTART_CONNECTORS_JOB       ResourceLocationJobType = "RestartConnectorsJob"
	RESOURCELOCATIONJOBTYPE_DELETE_RESOURCE_LOCATION_JOB ResourceLocationJobType = "DeleteResourceLocationJob"
	RESOURCELOCATIONJOBTYPE_ADD_NAT_GATEWAY_JOB          ResourceLocationJobType = "AddNatGatewayJob"
	RESOURCELOCATIONJOBTYPE_DELETE_NAT_GATEWAY_JOB       ResourceLocationJobType = "DeleteNatGatewayJob"
)

// All allowed values of ResourceLocationJobType enum
var AllowedResourceLocationJobTypeEnumValues = []ResourceLocationJobType{
	"AddConnectorsJob",
	"RestartConnectorsJob",
	"DeleteResourceLocationJob",
	"AddNatGatewayJob",
	"DeleteNatGatewayJob",
}

func (v *ResourceLocationJobType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = ResourceLocationJobType(value)
	return nil
}

// NewResourceLocationJobTypeFromValue returns a pointer to a valid ResourceLocationJobType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResourceLocationJobTypeFromValue(v string) (*ResourceLocationJobType, error) {
	ev := ResourceLocationJobType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResourceLocationJobType: valid values are %v", v, AllowedResourceLocationJobTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResourceLocationJobType) IsValid() bool {
	for _, existing := range AllowedResourceLocationJobTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResourceLocationJobType value
func (v ResourceLocationJobType) Ptr() *ResourceLocationJobType {
	return &v
}

type NullableResourceLocationJobType struct {
	value *ResourceLocationJobType
	isSet bool
}

func (v NullableResourceLocationJobType) Get() *ResourceLocationJobType {
	return v.value
}

func (v *NullableResourceLocationJobType) Set(val *ResourceLocationJobType) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceLocationJobType) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceLocationJobType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceLocationJobType(val *ResourceLocationJobType) *NullableResourceLocationJobType {
	return &NullableResourceLocationJobType{value: val, isSet: true}
}

func (v NullableResourceLocationJobType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceLocationJobType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
