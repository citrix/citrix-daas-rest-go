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

// DomainUserState the model 'DomainUserState'
type DomainUserState string

// List of DomainUserState
const (
	DOMAINUSERSTATE_READY           DomainUserState = "Ready"
	DOMAINUSERSTATE_DELETING        DomainUserState = "Deleting"
	DOMAINUSERSTATE_DELETION_FAILED DomainUserState = "DeletionFailed"
)

// All allowed values of DomainUserState enum
var AllowedDomainUserStateEnumValues = []DomainUserState{
	"Ready",
	"Deleting",
	"DeletionFailed",
}

func (v *DomainUserState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = DomainUserState(value)
	return nil
}

// NewDomainUserStateFromValue returns a pointer to a valid DomainUserState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDomainUserStateFromValue(v string) (*DomainUserState, error) {
	ev := DomainUserState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DomainUserState: valid values are %v", v, AllowedDomainUserStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DomainUserState) IsValid() bool {
	for _, existing := range AllowedDomainUserStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DomainUserState value
func (v DomainUserState) Ptr() *DomainUserState {
	return &v
}

type NullableDomainUserState struct {
	value *DomainUserState
	isSet bool
}

func (v NullableDomainUserState) Get() *DomainUserState {
	return v.value
}

func (v *NullableDomainUserState) Set(val *DomainUserState) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainUserState) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainUserState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainUserState(val *DomainUserState) *NullableDomainUserState {
	return &NullableDomainUserState{value: val, isSet: true}
}

func (v NullableDomainUserState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainUserState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
