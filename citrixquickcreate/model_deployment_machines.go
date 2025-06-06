/*
Quick Create Service for DaaS Quick Deploy with AWS WorkSpaces Core - Tech Preview

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: DaaSTechPreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickcreate

import (
	"encoding/json"
)

// checks if the DeploymentMachines type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentMachines{}

// DeploymentMachines struct for DeploymentMachines
type DeploymentMachines struct {
	// Enumerable of HostingUnit
	Items []AwsEdcDeploymentMachine `json:"items,omitempty"`
}

// NewDeploymentMachines instantiates a new DeploymentMachines object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentMachines() *DeploymentMachines {
	this := DeploymentMachines{}
	return &this
}

// NewDeploymentMachinesWithDefaults instantiates a new DeploymentMachines object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentMachinesWithDefaults() *DeploymentMachines {
	this := DeploymentMachines{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentMachines) GetItems() []AwsEdcDeploymentMachine {
	if o == nil {
		var ret []AwsEdcDeploymentMachine
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentMachines) GetItemsOk() ([]AwsEdcDeploymentMachine, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *DeploymentMachines) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []AwsEdcDeploymentMachine and assigns it to the Items field.
func (o *DeploymentMachines) SetItems(v []AwsEdcDeploymentMachine) {
	o.Items = v
}

func (o DeploymentMachines) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentMachines) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableDeploymentMachines struct {
	value *DeploymentMachines
	isSet bool
}

func (v NullableDeploymentMachines) Get() *DeploymentMachines {
	return v.value
}

func (v *NullableDeploymentMachines) Set(val *DeploymentMachines) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentMachines) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentMachines) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentMachines(val *DeploymentMachines) *NullableDeploymentMachines {
	return &NullableDeploymentMachines{value: val, isSet: true}
}

func (v NullableDeploymentMachines) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentMachines) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
