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

// CatalogSearchProperty Machine catalog properties can be used for advanced catalog searches.
type CatalogSearchProperty string

// List of CatalogSearchProperty
const (
	CATALOGSEARCHPROPERTY_NAME                   CatalogSearchProperty = "Name"
	CATALOGSEARCHPROPERTY_FULL_NAME              CatalogSearchProperty = "FullName"
	CATALOGSEARCHPROPERTY_UID                    CatalogSearchProperty = "Uid"
	CATALOGSEARCHPROPERTY_ALLOCATION_TYPE        CatalogSearchProperty = "AllocationType"
	CATALOGSEARCHPROPERTY_ASSIGNED_COUNT         CatalogSearchProperty = "AssignedCount"
	CATALOGSEARCHPROPERTY_AVAILABLE_COUNT        CatalogSearchProperty = "AvailableCount"
	CATALOGSEARCHPROPERTY_DESCRIPTION            CatalogSearchProperty = "Description"
	CATALOGSEARCHPROPERTY_IS_REMOTE_PC           CatalogSearchProperty = "IsRemotePC"
	CATALOGSEARCHPROPERTY_MACHINES_ARE_PHYSICAL  CatalogSearchProperty = "MachinesArePhysical"
	CATALOGSEARCHPROPERTY_PROVISIONING_TYPE      CatalogSearchProperty = "ProvisioningType"
	CATALOGSEARCHPROPERTY_PVS_ADDRESS            CatalogSearchProperty = "PvsAddress"
	CATALOGSEARCHPROPERTY_PVS_DOMAIN             CatalogSearchProperty = "PvsDomain"
	CATALOGSEARCHPROPERTY_SESSION_SUPPORT        CatalogSearchProperty = "SessionSupport"
	CATALOGSEARCHPROPERTY_UNASSIGNED_COUNT       CatalogSearchProperty = "UnassignedCount"
	CATALOGSEARCHPROPERTY_USED_COUNT             CatalogSearchProperty = "UsedCount"
	CATALOGSEARCHPROPERTY_PERSIST_CHANGES        CatalogSearchProperty = "PersistChanges"
	CATALOGSEARCHPROPERTY_PROVISIONING_SCHEME_ID CatalogSearchProperty = "ProvisioningSchemeId"
)

// All allowed values of CatalogSearchProperty enum
var AllowedCatalogSearchPropertyEnumValues = []CatalogSearchProperty{
	"Name",
	"FullName",
	"Uid",
	"AllocationType",
	"AssignedCount",
	"AvailableCount",
	"Description",
	"IsRemotePC",
	"MachinesArePhysical",
	"ProvisioningType",
	"PvsAddress",
	"PvsDomain",
	"SessionSupport",
	"UnassignedCount",
	"UsedCount",
	"PersistChanges",
	"ProvisioningSchemeId",
}

func (v *CatalogSearchProperty) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = CatalogSearchProperty(value)
	return nil
}

// NewCatalogSearchPropertyFromValue returns a pointer to a valid CatalogSearchProperty
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCatalogSearchPropertyFromValue(v string) (*CatalogSearchProperty, error) {
	ev := CatalogSearchProperty(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CatalogSearchProperty: valid values are %v", v, AllowedCatalogSearchPropertyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CatalogSearchProperty) IsValid() bool {
	for _, existing := range AllowedCatalogSearchPropertyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CatalogSearchProperty value
func (v CatalogSearchProperty) Ptr() *CatalogSearchProperty {
	return &v
}

type NullableCatalogSearchProperty struct {
	value *CatalogSearchProperty
	isSet bool
}

func (v NullableCatalogSearchProperty) Get() *CatalogSearchProperty {
	return v.value
}

func (v *NullableCatalogSearchProperty) Set(val *CatalogSearchProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogSearchProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogSearchProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogSearchProperty(val *CatalogSearchProperty) *NullableCatalogSearchProperty {
	return &NullableCatalogSearchProperty{value: val, isSet: true}
}

func (v NullableCatalogSearchProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogSearchProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
