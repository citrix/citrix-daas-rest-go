/*
Citrix.CloudServices.Cws.Api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixcws

import (
	"encoding/json"
	"fmt"
)

// CwsIdentityProviderType the model 'CwsIdentityProviderType'
type CwsIdentityProviderType string

// List of CwsIdentityProviderType
const (
	CWSIDENTITYPROVIDERTYPE_AD              CwsIdentityProviderType = "Ad"
	CWSIDENTITYPROVIDERTYPE_ADAPTIVE_AUTH   CwsIdentityProviderType = "AdaptiveAuth"
	CWSIDENTITYPROVIDERTYPE_AD_OPT          CwsIdentityProviderType = "AdOpt"
	CWSIDENTITYPROVIDERTYPE_AZURE_AD        CwsIdentityProviderType = "AzureAd"
	CWSIDENTITYPROVIDERTYPE_GOOGLE          CwsIdentityProviderType = "Google"
	CWSIDENTITYPROVIDERTYPE_OKTA            CwsIdentityProviderType = "Okta"
	CWSIDENTITYPROVIDERTYPE_ON_PREM_GATEWAY CwsIdentityProviderType = "OnPremGateway"
	CWSIDENTITYPROVIDERTYPE_SAML            CwsIdentityProviderType = "Saml"
)

// All allowed values of CwsIdentityProviderType enum
var AllowedCwsIdentityProviderTypeEnumValues = []CwsIdentityProviderType{
	"Ad",
	"AdaptiveAuth",
	"AdOpt",
	"AzureAd",
	"Google",
	"Okta",
	"OnPremGateway",
	"Saml",
}

func (v *CwsIdentityProviderType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CwsIdentityProviderType(value)
	for _, existing := range AllowedCwsIdentityProviderTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CwsIdentityProviderType", value)
}

// NewCwsIdentityProviderTypeFromValue returns a pointer to a valid CwsIdentityProviderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCwsIdentityProviderTypeFromValue(v string) (*CwsIdentityProviderType, error) {
	ev := CwsIdentityProviderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CwsIdentityProviderType: valid values are %v", v, AllowedCwsIdentityProviderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CwsIdentityProviderType) IsValid() bool {
	for _, existing := range AllowedCwsIdentityProviderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CwsIdentityProviderType value
func (v CwsIdentityProviderType) Ptr() *CwsIdentityProviderType {
	return &v
}

type NullableCwsIdentityProviderType struct {
	value *CwsIdentityProviderType
	isSet bool
}

func (v NullableCwsIdentityProviderType) Get() *CwsIdentityProviderType {
	return v.value
}

func (v *NullableCwsIdentityProviderType) Set(val *CwsIdentityProviderType) {
	v.value = val
	v.isSet = true
}

func (v NullableCwsIdentityProviderType) IsSet() bool {
	return v.isSet
}

func (v *NullableCwsIdentityProviderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCwsIdentityProviderType(val *CwsIdentityProviderType) *NullableCwsIdentityProviderType {
	return &NullableCwsIdentityProviderType{value: val, isSet: true}
}

func (v NullableCwsIdentityProviderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCwsIdentityProviderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
