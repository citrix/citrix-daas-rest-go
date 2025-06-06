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

// checks if the CustomerCatalogOverviewsModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerCatalogOverviewsModel{}

// CustomerCatalogOverviewsModel struct for CustomerCatalogOverviewsModel
type CustomerCatalogOverviewsModel struct {
	// List of catalogs that are available to the user
	Items []CatalogConfiguration `json:"items"`
	// Alias of Items property for backward compatibility
	Catalogs []CatalogConfiguration `json:"catalogs,omitempty"`
	// Limit imposed on the number of catalogs for the customer
	CatalogLimit *CatalogLimitModel `json:"catalogLimit,omitempty"`
	StaleData    *bool              `json:"staleData,omitempty"`
}

// NewCustomerCatalogOverviewsModelWithDefaults instantiates a new CustomerCatalogOverviewsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerCatalogOverviewsModelWithDefaults() *CustomerCatalogOverviewsModel {
	this := CustomerCatalogOverviewsModel{}
	return &this
}

// GetItems returns the Items field value
func (o *CustomerCatalogOverviewsModel) GetItems() []CatalogConfiguration {
	if o == nil {
		var ret []CatalogConfiguration
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CustomerCatalogOverviewsModel) GetItemsOk() ([]CatalogConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *CustomerCatalogOverviewsModel) SetItems(v []CatalogConfiguration) {
	o.Items = v
}

// GetCatalogs returns the Catalogs field value if set, zero value otherwise.
func (o *CustomerCatalogOverviewsModel) GetCatalogs() []CatalogConfiguration {
	if o == nil || IsNil(o.Catalogs) {
		var ret []CatalogConfiguration
		return ret
	}
	return o.Catalogs
}

// GetCatalogsOk returns a tuple with the Catalogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCatalogOverviewsModel) GetCatalogsOk() ([]CatalogConfiguration, bool) {
	if o == nil || IsNil(o.Catalogs) {
		return nil, false
	}
	return o.Catalogs, true
}

// SetCatalogs gets a reference to the given []CatalogConfiguration and assigns it to the Catalogs field.
func (o *CustomerCatalogOverviewsModel) SetCatalogs(v []CatalogConfiguration) {
	o.Catalogs = v
}

// GetCatalogLimit returns the CatalogLimit field value if set, zero value otherwise.
func (o *CustomerCatalogOverviewsModel) GetCatalogLimit() CatalogLimitModel {
	if o == nil || IsNil(o.CatalogLimit) {
		var ret CatalogLimitModel
		return ret
	}
	return *o.CatalogLimit
}

// GetCatalogLimitOk returns a tuple with the CatalogLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCatalogOverviewsModel) GetCatalogLimitOk() (*CatalogLimitModel, bool) {
	if o == nil || IsNil(o.CatalogLimit) {
		return nil, false
	}
	return o.CatalogLimit, true
}

// SetCatalogLimit gets a reference to the given CatalogLimitModel and assigns it to the CatalogLimit field.
func (o *CustomerCatalogOverviewsModel) SetCatalogLimit(v CatalogLimitModel) {
	o.CatalogLimit = &v
}

// GetStaleData returns the StaleData field value if set, zero value otherwise.
func (o *CustomerCatalogOverviewsModel) GetStaleData() bool {
	if o == nil || IsNil(o.StaleData) {
		var ret bool
		return ret
	}
	return *o.StaleData
}

// GetStaleDataOk returns a tuple with the StaleData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCatalogOverviewsModel) GetStaleDataOk() (*bool, bool) {
	if o == nil || IsNil(o.StaleData) {
		return nil, false
	}
	return o.StaleData, true
}

// SetStaleData gets a reference to the given bool and assigns it to the StaleData field.
func (o *CustomerCatalogOverviewsModel) SetStaleData(v bool) {
	o.StaleData = &v
}

func (o CustomerCatalogOverviewsModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerCatalogOverviewsModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	if !IsNil(o.Catalogs) {
		toSerialize["catalogs"] = o.Catalogs
	}
	if !IsNil(o.CatalogLimit) {
		toSerialize["catalogLimit"] = o.CatalogLimit
	}
	if !IsNil(o.StaleData) {
		toSerialize["staleData"] = o.StaleData
	}
	return toSerialize, nil
}

type NullableCustomerCatalogOverviewsModel struct {
	value *CustomerCatalogOverviewsModel
	isSet bool
}

func (v NullableCustomerCatalogOverviewsModel) Get() *CustomerCatalogOverviewsModel {
	return v.value
}

func (v *NullableCustomerCatalogOverviewsModel) Set(val *CustomerCatalogOverviewsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerCatalogOverviewsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerCatalogOverviewsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerCatalogOverviewsModel(val *CustomerCatalogOverviewsModel) *NullableCustomerCatalogOverviewsModel {
	return &NullableCustomerCatalogOverviewsModel{value: val, isSet: true}
}

func (v NullableCustomerCatalogOverviewsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerCatalogOverviewsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
