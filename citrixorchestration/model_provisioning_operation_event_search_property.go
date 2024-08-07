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

// ProvisioningOperationEventSearchProperty Properties which can be used for provisioned virtual machines.             
type ProvisioningOperationEventSearchProperty string

// List of ProvisioningOperationEventSearchProperty
const (
	PROVISIONINGOPERATIONEVENTSEARCHPROPERTY_EVENT_ADDITIONAL_DATA ProvisioningOperationEventSearchProperty = "EventAdditionalData"
	PROVISIONINGOPERATIONEVENTSEARCHPROPERTY_EVENT_CATEGORY ProvisioningOperationEventSearchProperty = "EventCategory"
	PROVISIONINGOPERATIONEVENTSEARCHPROPERTY_EVENT_DATE_TIME ProvisioningOperationEventSearchProperty = "EventDateTime"
	PROVISIONINGOPERATIONEVENTSEARCHPROPERTY_EVENT_MESSAGE ProvisioningOperationEventSearchProperty = "EventMessage"
	PROVISIONINGOPERATIONEVENTSEARCHPROPERTY_EVENT_RECORD_ID ProvisioningOperationEventSearchProperty = "EventRecordId"
	PROVISIONINGOPERATIONEVENTSEARCHPROPERTY_EVENT_SEVERITY ProvisioningOperationEventSearchProperty = "EventSeverity"
	PROVISIONINGOPERATIONEVENTSEARCHPROPERTY_EVENT_SOURCE ProvisioningOperationEventSearchProperty = "EventSource"
	PROVISIONINGOPERATIONEVENTSEARCHPROPERTY_EVENT_STATE ProvisioningOperationEventSearchProperty = "EventState"
	PROVISIONINGOPERATIONEVENTSEARCHPROPERTY_LINKED_OBJECT_TYPE ProvisioningOperationEventSearchProperty = "LinkedObjectType"
	PROVISIONINGOPERATIONEVENTSEARCHPROPERTY_LINKED_OBJECT_UID ProvisioningOperationEventSearchProperty = "LinkedObjectUid"
	PROVISIONINGOPERATIONEVENTSEARCHPROPERTY_OPERATION_NAME ProvisioningOperationEventSearchProperty = "OperationName"
	PROVISIONINGOPERATIONEVENTSEARCHPROPERTY_OPERATION_TARGET_NAME ProvisioningOperationEventSearchProperty = "OperationTargetName"
	PROVISIONINGOPERATIONEVENTSEARCHPROPERTY_OPERATION_TARGET_TYPE ProvisioningOperationEventSearchProperty = "OperationTargetType"
	PROVISIONINGOPERATIONEVENTSEARCHPROPERTY_OPERATION_TYPE ProvisioningOperationEventSearchProperty = "OperationType"
	PROVISIONINGOPERATIONEVENTSEARCHPROPERTY_RECOMMENDATION ProvisioningOperationEventSearchProperty = "Recommendation"
)

// All allowed values of ProvisioningOperationEventSearchProperty enum
var AllowedProvisioningOperationEventSearchPropertyEnumValues = []ProvisioningOperationEventSearchProperty{
	"EventAdditionalData",
	"EventCategory",
	"EventDateTime",
	"EventMessage",
	"EventRecordId",
	"EventSeverity",
	"EventSource",
	"EventState",
	"LinkedObjectType",
	"LinkedObjectUid",
	"OperationName",
	"OperationTargetName",
	"OperationTargetType",
	"OperationType",
	"Recommendation",
}

func (v *ProvisioningOperationEventSearchProperty) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProvisioningOperationEventSearchProperty(value)
	for _, existing := range AllowedProvisioningOperationEventSearchPropertyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProvisioningOperationEventSearchProperty", value)
}

// NewProvisioningOperationEventSearchPropertyFromValue returns a pointer to a valid ProvisioningOperationEventSearchProperty
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProvisioningOperationEventSearchPropertyFromValue(v string) (*ProvisioningOperationEventSearchProperty, error) {
	ev := ProvisioningOperationEventSearchProperty(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProvisioningOperationEventSearchProperty: valid values are %v", v, AllowedProvisioningOperationEventSearchPropertyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProvisioningOperationEventSearchProperty) IsValid() bool {
	for _, existing := range AllowedProvisioningOperationEventSearchPropertyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProvisioningOperationEventSearchProperty value
func (v ProvisioningOperationEventSearchProperty) Ptr() *ProvisioningOperationEventSearchProperty {
	return &v
}

type NullableProvisioningOperationEventSearchProperty struct {
	value *ProvisioningOperationEventSearchProperty
	isSet bool
}

func (v NullableProvisioningOperationEventSearchProperty) Get() *ProvisioningOperationEventSearchProperty {
	return v.value
}

func (v *NullableProvisioningOperationEventSearchProperty) Set(val *ProvisioningOperationEventSearchProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningOperationEventSearchProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningOperationEventSearchProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningOperationEventSearchProperty(val *ProvisioningOperationEventSearchProperty) *NullableProvisioningOperationEventSearchProperty {
	return &NullableProvisioningOperationEventSearchProperty{value: val, isSet: true}
}

func (v NullableProvisioningOperationEventSearchProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningOperationEventSearchProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
