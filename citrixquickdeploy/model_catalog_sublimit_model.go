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

// checks if the CatalogSublimitModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogSublimitModel{}

// CatalogSublimitModel struct for CatalogSublimitModel
type CatalogSublimitModel struct {
	// Limit on number of catalog creation.
	Limit *int32 `json:"limit,omitempty"`
	// Number of existing catalogs.
	ExistingCatalogsCount *int32 `json:"existingCatalogsCount,omitempty"`
	// Indicates whether a customer's number of catalog has reached the limit.
	HasLimitReached *bool `json:"hasLimitReached,omitempty"`
	// The maximum number of VDAs per catalog.
	MaxVdaCount *int32 `json:"maxVdaCount,omitempty"`
	// The max number of VDAs supported per Azure Subscription.
	MaxVdaPerSubscription *int32 `json:"maxVdaPerSubscription,omitempty"`
}

// NewCatalogSublimitModelWithDefaults instantiates a new CatalogSublimitModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogSublimitModelWithDefaults() *CatalogSublimitModel {
	this := CatalogSublimitModel{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *CatalogSublimitModel) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogSublimitModel) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *CatalogSublimitModel) SetLimit(v int32) {
	o.Limit = &v
}

// GetExistingCatalogsCount returns the ExistingCatalogsCount field value if set, zero value otherwise.
func (o *CatalogSublimitModel) GetExistingCatalogsCount() int32 {
	if o == nil || IsNil(o.ExistingCatalogsCount) {
		var ret int32
		return ret
	}
	return *o.ExistingCatalogsCount
}

// GetExistingCatalogsCountOk returns a tuple with the ExistingCatalogsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogSublimitModel) GetExistingCatalogsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ExistingCatalogsCount) {
		return nil, false
	}
	return o.ExistingCatalogsCount, true
}

// SetExistingCatalogsCount gets a reference to the given int32 and assigns it to the ExistingCatalogsCount field.
func (o *CatalogSublimitModel) SetExistingCatalogsCount(v int32) {
	o.ExistingCatalogsCount = &v
}

// GetHasLimitReached returns the HasLimitReached field value if set, zero value otherwise.
func (o *CatalogSublimitModel) GetHasLimitReached() bool {
	if o == nil || IsNil(o.HasLimitReached) {
		var ret bool
		return ret
	}
	return *o.HasLimitReached
}

// GetHasLimitReachedOk returns a tuple with the HasLimitReached field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogSublimitModel) GetHasLimitReachedOk() (*bool, bool) {
	if o == nil || IsNil(o.HasLimitReached) {
		return nil, false
	}
	return o.HasLimitReached, true
}

// SetHasLimitReached gets a reference to the given bool and assigns it to the HasLimitReached field.
func (o *CatalogSublimitModel) SetHasLimitReached(v bool) {
	o.HasLimitReached = &v
}

// GetMaxVdaCount returns the MaxVdaCount field value if set, zero value otherwise.
func (o *CatalogSublimitModel) GetMaxVdaCount() int32 {
	if o == nil || IsNil(o.MaxVdaCount) {
		var ret int32
		return ret
	}
	return *o.MaxVdaCount
}

// GetMaxVdaCountOk returns a tuple with the MaxVdaCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogSublimitModel) GetMaxVdaCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxVdaCount) {
		return nil, false
	}
	return o.MaxVdaCount, true
}

// SetMaxVdaCount gets a reference to the given int32 and assigns it to the MaxVdaCount field.
func (o *CatalogSublimitModel) SetMaxVdaCount(v int32) {
	o.MaxVdaCount = &v
}

// GetMaxVdaPerSubscription returns the MaxVdaPerSubscription field value if set, zero value otherwise.
func (o *CatalogSublimitModel) GetMaxVdaPerSubscription() int32 {
	if o == nil || IsNil(o.MaxVdaPerSubscription) {
		var ret int32
		return ret
	}
	return *o.MaxVdaPerSubscription
}

// GetMaxVdaPerSubscriptionOk returns a tuple with the MaxVdaPerSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogSublimitModel) GetMaxVdaPerSubscriptionOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxVdaPerSubscription) {
		return nil, false
	}
	return o.MaxVdaPerSubscription, true
}

// SetMaxVdaPerSubscription gets a reference to the given int32 and assigns it to the MaxVdaPerSubscription field.
func (o *CatalogSublimitModel) SetMaxVdaPerSubscription(v int32) {
	o.MaxVdaPerSubscription = &v
}

func (o CatalogSublimitModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogSublimitModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.ExistingCatalogsCount) {
		toSerialize["existingCatalogsCount"] = o.ExistingCatalogsCount
	}
	if !IsNil(o.HasLimitReached) {
		toSerialize["hasLimitReached"] = o.HasLimitReached
	}
	if !IsNil(o.MaxVdaCount) {
		toSerialize["maxVdaCount"] = o.MaxVdaCount
	}
	if !IsNil(o.MaxVdaPerSubscription) {
		toSerialize["maxVdaPerSubscription"] = o.MaxVdaPerSubscription
	}
	return toSerialize, nil
}

type NullableCatalogSublimitModel struct {
	value *CatalogSublimitModel
	isSet bool
}

func (v NullableCatalogSublimitModel) Get() *CatalogSublimitModel {
	return v.value
}

func (v *NullableCatalogSublimitModel) Set(val *CatalogSublimitModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogSublimitModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogSublimitModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogSublimitModel(val *CatalogSublimitModel) *NullableCatalogSublimitModel {
	return &NullableCatalogSublimitModel{value: val, isSet: true}
}

func (v NullableCatalogSublimitModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogSublimitModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
