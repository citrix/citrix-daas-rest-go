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

// checks if the TroubleshootScriptsOverviewsModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TroubleshootScriptsOverviewsModel{}

// TroubleshootScriptsOverviewsModel struct for TroubleshootScriptsOverviewsModel
type TroubleshootScriptsOverviewsModel struct {
	Items []TroubleshootScriptOverview `json:"items"`
}

// NewTroubleshootScriptsOverviewsModelWithDefaults instantiates a new TroubleshootScriptsOverviewsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTroubleshootScriptsOverviewsModelWithDefaults() *TroubleshootScriptsOverviewsModel {
	this := TroubleshootScriptsOverviewsModel{}
	return &this
}

// GetItems returns the Items field value
func (o *TroubleshootScriptsOverviewsModel) GetItems() []TroubleshootScriptOverview {
	if o == nil {
		var ret []TroubleshootScriptOverview
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *TroubleshootScriptsOverviewsModel) GetItemsOk() ([]TroubleshootScriptOverview, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *TroubleshootScriptsOverviewsModel) SetItems(v []TroubleshootScriptOverview) {
	o.Items = v
}

func (o TroubleshootScriptsOverviewsModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TroubleshootScriptsOverviewsModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableTroubleshootScriptsOverviewsModel struct {
	value *TroubleshootScriptsOverviewsModel
	isSet bool
}

func (v NullableTroubleshootScriptsOverviewsModel) Get() *TroubleshootScriptsOverviewsModel {
	return v.value
}

func (v *NullableTroubleshootScriptsOverviewsModel) Set(val *TroubleshootScriptsOverviewsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTroubleshootScriptsOverviewsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTroubleshootScriptsOverviewsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTroubleshootScriptsOverviewsModel(val *TroubleshootScriptsOverviewsModel) *NullableTroubleshootScriptsOverviewsModel {
	return &NullableTroubleshootScriptsOverviewsModel{value: val, isSet: true}
}

func (v NullableTroubleshootScriptsOverviewsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTroubleshootScriptsOverviewsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
