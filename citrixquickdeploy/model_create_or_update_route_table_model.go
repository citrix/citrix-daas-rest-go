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

// checks if the CreateOrUpdateRouteTableModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdateRouteTableModel{}

// CreateOrUpdateRouteTableModel struct for CreateOrUpdateRouteTableModel
type CreateOrUpdateRouteTableModel struct {
	// The list of routes.
	AzureRoutes []AzureRoute `json:"azureRoutes"`
	// Indicates whether disable the border gateway protocol route propagation.
	DisableRoutePropagation *bool `json:"disableRoutePropagation,omitempty"`
}

// NewCreateOrUpdateRouteTableModelWithDefaults instantiates a new CreateOrUpdateRouteTableModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateRouteTableModelWithDefaults() *CreateOrUpdateRouteTableModel {
	this := CreateOrUpdateRouteTableModel{}
	return &this
}

// GetAzureRoutes returns the AzureRoutes field value
func (o *CreateOrUpdateRouteTableModel) GetAzureRoutes() []AzureRoute {
	if o == nil {
		var ret []AzureRoute
		return ret
	}

	return o.AzureRoutes
}

// GetAzureRoutesOk returns a tuple with the AzureRoutes field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateRouteTableModel) GetAzureRoutesOk() ([]AzureRoute, bool) {
	if o == nil {
		return nil, false
	}
	return o.AzureRoutes, true
}

// SetAzureRoutes sets field value
func (o *CreateOrUpdateRouteTableModel) SetAzureRoutes(v []AzureRoute) {
	o.AzureRoutes = v
}

// GetDisableRoutePropagation returns the DisableRoutePropagation field value if set, zero value otherwise.
func (o *CreateOrUpdateRouteTableModel) GetDisableRoutePropagation() bool {
	if o == nil || IsNil(o.DisableRoutePropagation) {
		var ret bool
		return ret
	}
	return *o.DisableRoutePropagation
}

// GetDisableRoutePropagationOk returns a tuple with the DisableRoutePropagation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateRouteTableModel) GetDisableRoutePropagationOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableRoutePropagation) {
		return nil, false
	}
	return o.DisableRoutePropagation, true
}

// SetDisableRoutePropagation gets a reference to the given bool and assigns it to the DisableRoutePropagation field.
func (o *CreateOrUpdateRouteTableModel) SetDisableRoutePropagation(v bool) {
	o.DisableRoutePropagation = &v
}

func (o CreateOrUpdateRouteTableModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrUpdateRouteTableModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["azureRoutes"] = o.AzureRoutes
	if !IsNil(o.DisableRoutePropagation) {
		toSerialize["disableRoutePropagation"] = o.DisableRoutePropagation
	}
	return toSerialize, nil
}

type NullableCreateOrUpdateRouteTableModel struct {
	value *CreateOrUpdateRouteTableModel
	isSet bool
}

func (v NullableCreateOrUpdateRouteTableModel) Get() *CreateOrUpdateRouteTableModel {
	return v.value
}

func (v *NullableCreateOrUpdateRouteTableModel) Set(val *CreateOrUpdateRouteTableModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateRouteTableModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateRouteTableModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateRouteTableModel(val *CreateOrUpdateRouteTableModel) *NullableCreateOrUpdateRouteTableModel {
	return &NullableCreateOrUpdateRouteTableModel{value: val, isSet: true}
}

func (v NullableCreateOrUpdateRouteTableModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateRouteTableModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
