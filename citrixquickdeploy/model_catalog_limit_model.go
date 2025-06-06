/*
Citrix Virtual App & Desktop Catalog Service 147.0.26651.57932

Catalog Service

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickdeploy

import (
	"encoding/json"
)

// checks if the CatalogLimitModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogLimitModel{}

// CatalogLimitModel struct for CatalogLimitModel
type CatalogLimitModel struct {
	// Number of existing catalogs
	ExistingCatalogsCount *int32 `json:"existingCatalogsCount,omitempty"`
	// Current service state of customer
	ServiceState *ServiceState `json:"serviceState,omitempty"`
	// Limits on the Citrix Managed catalogs
	CitrixManagedLimit *CatalogSublimitModel `json:"citrixManagedLimit,omitempty"`
	// Limits on the BYOA catalogs
	ByoaLimit *CatalogSublimitModel `json:"byoaLimit,omitempty"`
}

// NewCatalogLimitModelWithDefaults instantiates a new CatalogLimitModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogLimitModelWithDefaults() *CatalogLimitModel {
	this := CatalogLimitModel{}
	return &this
}

// GetExistingCatalogsCount returns the ExistingCatalogsCount field value if set, zero value otherwise.
func (o *CatalogLimitModel) GetExistingCatalogsCount() int32 {
	if o == nil || IsNil(o.ExistingCatalogsCount) {
		var ret int32
		return ret
	}
	return *o.ExistingCatalogsCount
}

// GetExistingCatalogsCountOk returns a tuple with the ExistingCatalogsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogLimitModel) GetExistingCatalogsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ExistingCatalogsCount) {
		return nil, false
	}
	return o.ExistingCatalogsCount, true
}

// SetExistingCatalogsCount gets a reference to the given int32 and assigns it to the ExistingCatalogsCount field.
func (o *CatalogLimitModel) SetExistingCatalogsCount(v int32) {
	o.ExistingCatalogsCount = &v
}

// GetServiceState returns the ServiceState field value if set, zero value otherwise.
func (o *CatalogLimitModel) GetServiceState() ServiceState {
	if o == nil || IsNil(o.ServiceState) {
		var ret ServiceState
		return ret
	}
	return *o.ServiceState
}

// GetServiceStateOk returns a tuple with the ServiceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogLimitModel) GetServiceStateOk() (*ServiceState, bool) {
	if o == nil || IsNil(o.ServiceState) {
		return nil, false
	}
	return o.ServiceState, true
}

// SetServiceState gets a reference to the given ServiceState and assigns it to the ServiceState field.
func (o *CatalogLimitModel) SetServiceState(v ServiceState) {
	o.ServiceState = &v
}

// GetCitrixManagedLimit returns the CitrixManagedLimit field value if set, zero value otherwise.
func (o *CatalogLimitModel) GetCitrixManagedLimit() CatalogSublimitModel {
	if o == nil || IsNil(o.CitrixManagedLimit) {
		var ret CatalogSublimitModel
		return ret
	}
	return *o.CitrixManagedLimit
}

// GetCitrixManagedLimitOk returns a tuple with the CitrixManagedLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogLimitModel) GetCitrixManagedLimitOk() (*CatalogSublimitModel, bool) {
	if o == nil || IsNil(o.CitrixManagedLimit) {
		return nil, false
	}
	return o.CitrixManagedLimit, true
}

// SetCitrixManagedLimit gets a reference to the given CatalogSublimitModel and assigns it to the CitrixManagedLimit field.
func (o *CatalogLimitModel) SetCitrixManagedLimit(v CatalogSublimitModel) {
	o.CitrixManagedLimit = &v
}

// GetByoaLimit returns the ByoaLimit field value if set, zero value otherwise.
func (o *CatalogLimitModel) GetByoaLimit() CatalogSublimitModel {
	if o == nil || IsNil(o.ByoaLimit) {
		var ret CatalogSublimitModel
		return ret
	}
	return *o.ByoaLimit
}

// GetByoaLimitOk returns a tuple with the ByoaLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogLimitModel) GetByoaLimitOk() (*CatalogSublimitModel, bool) {
	if o == nil || IsNil(o.ByoaLimit) {
		return nil, false
	}
	return o.ByoaLimit, true
}

// SetByoaLimit gets a reference to the given CatalogSublimitModel and assigns it to the ByoaLimit field.
func (o *CatalogLimitModel) SetByoaLimit(v CatalogSublimitModel) {
	o.ByoaLimit = &v
}

func (o CatalogLimitModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogLimitModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExistingCatalogsCount) {
		toSerialize["existingCatalogsCount"] = o.ExistingCatalogsCount
	}
	if !IsNil(o.ServiceState) {
		toSerialize["serviceState"] = o.ServiceState
	}
	if !IsNil(o.CitrixManagedLimit) {
		toSerialize["citrixManagedLimit"] = o.CitrixManagedLimit
	}
	if !IsNil(o.ByoaLimit) {
		toSerialize["byoaLimit"] = o.ByoaLimit
	}
	return toSerialize, nil
}

type NullableCatalogLimitModel struct {
	value *CatalogLimitModel
	isSet bool
}

func (v NullableCatalogLimitModel) Get() *CatalogLimitModel {
	return v.value
}

func (v *NullableCatalogLimitModel) Set(val *CatalogLimitModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogLimitModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogLimitModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogLimitModel(val *CatalogLimitModel) *NullableCatalogLimitModel {
	return &NullableCatalogLimitModel{value: val, isSet: true}
}

func (v NullableCatalogLimitModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogLimitModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
