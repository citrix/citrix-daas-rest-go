/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-onprem-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
)

// checks if the TenantResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantResponseModel{}

// TenantResponseModel Tenant object.
type TenantResponseModel struct {
	// ID of the tenant.
	Id string `json:"Id"`
	// Name of the zone.
	Name string `json:"Name"`
}

// NewTenantResponseModel instantiates a new TenantResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantResponseModel(id string, name string) *TenantResponseModel {
	this := TenantResponseModel{}
	this.Id = id
	this.Name = name
	return &this
}

// NewTenantResponseModelWithDefaults instantiates a new TenantResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantResponseModelWithDefaults() *TenantResponseModel {
	this := TenantResponseModel{}
	return &this
}

// GetId returns the Id field value
func (o *TenantResponseModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TenantResponseModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TenantResponseModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TenantResponseModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TenantResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TenantResponseModel) SetName(v string) {
	o.Name = v
}

func (o TenantResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Id"] = o.Id
	toSerialize["Name"] = o.Name
	return toSerialize, nil
}

type NullableTenantResponseModel struct {
	value *TenantResponseModel
	isSet bool
}

func (v NullableTenantResponseModel) Get() *TenantResponseModel {
	return v.value
}

func (v *NullableTenantResponseModel) Set(val *TenantResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantResponseModel(val *TenantResponseModel) *NullableTenantResponseModel {
	return &NullableTenantResponseModel{value: val, isSet: true}
}

func (v NullableTenantResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
