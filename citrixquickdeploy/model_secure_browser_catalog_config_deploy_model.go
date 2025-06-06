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

// checks if the SecureBrowserCatalogConfigDeployModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecureBrowserCatalogConfigDeployModel{}

// SecureBrowserCatalogConfigDeployModel struct for SecureBrowserCatalogConfigDeployModel
type SecureBrowserCatalogConfigDeployModel struct {
	AddCatalog         *AddSecureBrowserCatalogModel `json:"addCatalog,omitempty"`
	AddCatalogCapacity *CatalogCapacitySettingsModel `json:"addCatalogCapacity,omitempty"`
	AddSecureBrowser   *CatalogSecureBrowserModel    `json:"addSecureBrowser,omitempty"`
}

// NewSecureBrowserCatalogConfigDeployModelWithDefaults instantiates a new SecureBrowserCatalogConfigDeployModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecureBrowserCatalogConfigDeployModelWithDefaults() *SecureBrowserCatalogConfigDeployModel {
	this := SecureBrowserCatalogConfigDeployModel{}
	return &this
}

// GetAddCatalog returns the AddCatalog field value if set, zero value otherwise.
func (o *SecureBrowserCatalogConfigDeployModel) GetAddCatalog() AddSecureBrowserCatalogModel {
	if o == nil || IsNil(o.AddCatalog) {
		var ret AddSecureBrowserCatalogModel
		return ret
	}
	return *o.AddCatalog
}

// GetAddCatalogOk returns a tuple with the AddCatalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureBrowserCatalogConfigDeployModel) GetAddCatalogOk() (*AddSecureBrowserCatalogModel, bool) {
	if o == nil || IsNil(o.AddCatalog) {
		return nil, false
	}
	return o.AddCatalog, true
}

// SetAddCatalog gets a reference to the given AddSecureBrowserCatalogModel and assigns it to the AddCatalog field.
func (o *SecureBrowserCatalogConfigDeployModel) SetAddCatalog(v AddSecureBrowserCatalogModel) {
	o.AddCatalog = &v
}

// GetAddCatalogCapacity returns the AddCatalogCapacity field value if set, zero value otherwise.
func (o *SecureBrowserCatalogConfigDeployModel) GetAddCatalogCapacity() CatalogCapacitySettingsModel {
	if o == nil || IsNil(o.AddCatalogCapacity) {
		var ret CatalogCapacitySettingsModel
		return ret
	}
	return *o.AddCatalogCapacity
}

// GetAddCatalogCapacityOk returns a tuple with the AddCatalogCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureBrowserCatalogConfigDeployModel) GetAddCatalogCapacityOk() (*CatalogCapacitySettingsModel, bool) {
	if o == nil || IsNil(o.AddCatalogCapacity) {
		return nil, false
	}
	return o.AddCatalogCapacity, true
}

// SetAddCatalogCapacity gets a reference to the given CatalogCapacitySettingsModel and assigns it to the AddCatalogCapacity field.
func (o *SecureBrowserCatalogConfigDeployModel) SetAddCatalogCapacity(v CatalogCapacitySettingsModel) {
	o.AddCatalogCapacity = &v
}

// GetAddSecureBrowser returns the AddSecureBrowser field value if set, zero value otherwise.
func (o *SecureBrowserCatalogConfigDeployModel) GetAddSecureBrowser() CatalogSecureBrowserModel {
	if o == nil || IsNil(o.AddSecureBrowser) {
		var ret CatalogSecureBrowserModel
		return ret
	}
	return *o.AddSecureBrowser
}

// GetAddSecureBrowserOk returns a tuple with the AddSecureBrowser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureBrowserCatalogConfigDeployModel) GetAddSecureBrowserOk() (*CatalogSecureBrowserModel, bool) {
	if o == nil || IsNil(o.AddSecureBrowser) {
		return nil, false
	}
	return o.AddSecureBrowser, true
}

// SetAddSecureBrowser gets a reference to the given CatalogSecureBrowserModel and assigns it to the AddSecureBrowser field.
func (o *SecureBrowserCatalogConfigDeployModel) SetAddSecureBrowser(v CatalogSecureBrowserModel) {
	o.AddSecureBrowser = &v
}

func (o SecureBrowserCatalogConfigDeployModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecureBrowserCatalogConfigDeployModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddCatalog) {
		toSerialize["addCatalog"] = o.AddCatalog
	}
	if !IsNil(o.AddCatalogCapacity) {
		toSerialize["addCatalogCapacity"] = o.AddCatalogCapacity
	}
	if !IsNil(o.AddSecureBrowser) {
		toSerialize["addSecureBrowser"] = o.AddSecureBrowser
	}
	return toSerialize, nil
}

type NullableSecureBrowserCatalogConfigDeployModel struct {
	value *SecureBrowserCatalogConfigDeployModel
	isSet bool
}

func (v NullableSecureBrowserCatalogConfigDeployModel) Get() *SecureBrowserCatalogConfigDeployModel {
	return v.value
}

func (v *NullableSecureBrowserCatalogConfigDeployModel) Set(val *SecureBrowserCatalogConfigDeployModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSecureBrowserCatalogConfigDeployModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSecureBrowserCatalogConfigDeployModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecureBrowserCatalogConfigDeployModel(val *SecureBrowserCatalogConfigDeployModel) *NullableSecureBrowserCatalogConfigDeployModel {
	return &NullableSecureBrowserCatalogConfigDeployModel{value: val, isSet: true}
}

func (v NullableSecureBrowserCatalogConfigDeployModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecureBrowserCatalogConfigDeployModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
