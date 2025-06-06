/*
Citrix.CloudServices.Cws.Api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixcws

import (
	"encoding/json"
)

// checks if the GatewayConnectionModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayConnectionModel{}

// GatewayConnectionModel struct for GatewayConnectionModel
type GatewayConnectionModel struct {
	ClientId   NullableString `json:"clientId,omitempty"`
	IssuerFqdn NullableString `json:"issuerFqdn,omitempty"`
}

// NewGatewayConnectionModel instantiates a new GatewayConnectionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayConnectionModel() *GatewayConnectionModel {
	this := GatewayConnectionModel{}
	return &this
}

// NewGatewayConnectionModelWithDefaults instantiates a new GatewayConnectionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayConnectionModelWithDefaults() *GatewayConnectionModel {
	this := GatewayConnectionModel{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GatewayConnectionModel) GetClientId() string {
	if o == nil || IsNil(o.ClientId.Get()) {
		var ret string
		return ret
	}
	return *o.ClientId.Get()
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GatewayConnectionModel) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientId.Get(), o.ClientId.IsSet()
}

// HasClientId returns a boolean if a field has been set.
func (o *GatewayConnectionModel) HasClientId() bool {
	if o != nil && o.ClientId.IsSet() {
		return true
	}

	return false
}

// SetClientId gets a reference to the given NullableString and assigns it to the ClientId field.
func (o *GatewayConnectionModel) SetClientId(v string) {
	o.ClientId.Set(&v)
}

// SetClientIdNil sets the value for ClientId to be an explicit nil
func (o *GatewayConnectionModel) SetClientIdNil() {
	o.ClientId.Set(nil)
}

// UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
func (o *GatewayConnectionModel) UnsetClientId() {
	o.ClientId.Unset()
}

// GetIssuerFqdn returns the IssuerFqdn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GatewayConnectionModel) GetIssuerFqdn() string {
	if o == nil || IsNil(o.IssuerFqdn.Get()) {
		var ret string
		return ret
	}
	return *o.IssuerFqdn.Get()
}

// GetIssuerFqdnOk returns a tuple with the IssuerFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GatewayConnectionModel) GetIssuerFqdnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssuerFqdn.Get(), o.IssuerFqdn.IsSet()
}

// HasIssuerFqdn returns a boolean if a field has been set.
func (o *GatewayConnectionModel) HasIssuerFqdn() bool {
	if o != nil && o.IssuerFqdn.IsSet() {
		return true
	}

	return false
}

// SetIssuerFqdn gets a reference to the given NullableString and assigns it to the IssuerFqdn field.
func (o *GatewayConnectionModel) SetIssuerFqdn(v string) {
	o.IssuerFqdn.Set(&v)
}

// SetIssuerFqdnNil sets the value for IssuerFqdn to be an explicit nil
func (o *GatewayConnectionModel) SetIssuerFqdnNil() {
	o.IssuerFqdn.Set(nil)
}

// UnsetIssuerFqdn ensures that no value is present for IssuerFqdn, not even an explicit nil
func (o *GatewayConnectionModel) UnsetIssuerFqdn() {
	o.IssuerFqdn.Unset()
}

func (o GatewayConnectionModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayConnectionModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId.IsSet() {
		toSerialize["clientId"] = o.ClientId.Get()
	}
	if o.IssuerFqdn.IsSet() {
		toSerialize["issuerFqdn"] = o.IssuerFqdn.Get()
	}
	return toSerialize, nil
}

type NullableGatewayConnectionModel struct {
	value *GatewayConnectionModel
	isSet bool
}

func (v NullableGatewayConnectionModel) Get() *GatewayConnectionModel {
	return v.value
}

func (v *NullableGatewayConnectionModel) Set(val *GatewayConnectionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayConnectionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayConnectionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayConnectionModel(val *GatewayConnectionModel) *NullableGatewayConnectionModel {
	return &NullableGatewayConnectionModel{value: val, isSet: true}
}

func (v NullableGatewayConnectionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayConnectionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
