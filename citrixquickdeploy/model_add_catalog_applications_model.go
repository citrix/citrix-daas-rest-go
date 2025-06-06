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

// checks if the AddCatalogApplicationsModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCatalogApplicationsModel{}

// AddCatalogApplicationsModel struct for AddCatalogApplicationsModel
type AddCatalogApplicationsModel struct {
	// Applications to be added
	Applications []AddCatalogApplicationModel `json:"applications"`
}

// NewAddCatalogApplicationsModelWithDefaults instantiates a new AddCatalogApplicationsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCatalogApplicationsModelWithDefaults() *AddCatalogApplicationsModel {
	this := AddCatalogApplicationsModel{}
	return &this
}

// GetApplications returns the Applications field value
func (o *AddCatalogApplicationsModel) GetApplications() []AddCatalogApplicationModel {
	if o == nil {
		var ret []AddCatalogApplicationModel
		return ret
	}

	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value
// and a boolean to check if the value has been set.
func (o *AddCatalogApplicationsModel) GetApplicationsOk() ([]AddCatalogApplicationModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Applications, true
}

// SetApplications sets field value
func (o *AddCatalogApplicationsModel) SetApplications(v []AddCatalogApplicationModel) {
	o.Applications = v
}

func (o AddCatalogApplicationsModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCatalogApplicationsModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["applications"] = o.Applications
	return toSerialize, nil
}

type NullableAddCatalogApplicationsModel struct {
	value *AddCatalogApplicationsModel
	isSet bool
}

func (v NullableAddCatalogApplicationsModel) Get() *AddCatalogApplicationsModel {
	return v.value
}

func (v *NullableAddCatalogApplicationsModel) Set(val *AddCatalogApplicationsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCatalogApplicationsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCatalogApplicationsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCatalogApplicationsModel(val *AddCatalogApplicationsModel) *NullableAddCatalogApplicationsModel {
	return &NullableAddCatalogApplicationsModel{value: val, isSet: true}
}

func (v NullableAddCatalogApplicationsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCatalogApplicationsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
