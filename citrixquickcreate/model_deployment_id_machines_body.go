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

// checks if the DeploymentIdMachinesBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentIdMachinesBody{}

// DeploymentIdMachinesBody Turn on or off maintenance mode on machines by id
type DeploymentIdMachinesBody struct {
	AccountType AccountType `json:"accountType"`
	// The type of provider associated with the account
	InMaintenanceMode bool `json:"inMaintenanceMode"`
	// The list of machine ids to be updated
	MachineIds []string `json:"machineIds"`
}

// NewDeploymentIdMachinesBody instantiates a new DeploymentIdMachinesBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentIdMachinesBody(accountType AccountType, inMaintenanceMode bool, machineIds []string) *DeploymentIdMachinesBody {
	this := DeploymentIdMachinesBody{}
	this.AccountType = accountType
	this.InMaintenanceMode = inMaintenanceMode
	this.MachineIds = machineIds
	return &this
}

// NewDeploymentIdMachinesBodyWithDefaults instantiates a new DeploymentIdMachinesBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentIdMachinesBodyWithDefaults() *DeploymentIdMachinesBody {
	this := DeploymentIdMachinesBody{}
	return &this
}

// GetAccountType returns the AccountType field value
func (o *DeploymentIdMachinesBody) GetAccountType() AccountType {
	if o == nil {
		var ret AccountType
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *DeploymentIdMachinesBody) GetAccountTypeOk() (*AccountType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *DeploymentIdMachinesBody) SetAccountType(v AccountType) {
	o.AccountType = v
}

// GetInMaintenanceMode returns the InMaintenanceMode field value
func (o *DeploymentIdMachinesBody) GetInMaintenanceMode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InMaintenanceMode
}

// GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field value
// and a boolean to check if the value has been set.
func (o *DeploymentIdMachinesBody) GetInMaintenanceModeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InMaintenanceMode, true
}

// SetInMaintenanceMode sets field value
func (o *DeploymentIdMachinesBody) SetInMaintenanceMode(v bool) {
	o.InMaintenanceMode = v
}

// GetMachineIds returns the MachineIds field value
func (o *DeploymentIdMachinesBody) GetMachineIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MachineIds
}

// GetMachineIdsOk returns a tuple with the MachineIds field value
// and a boolean to check if the value has been set.
func (o *DeploymentIdMachinesBody) GetMachineIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MachineIds, true
}

// SetMachineIds sets field value
func (o *DeploymentIdMachinesBody) SetMachineIds(v []string) {
	o.MachineIds = v
}

func (o DeploymentIdMachinesBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentIdMachinesBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountType"] = o.AccountType
	toSerialize["inMaintenanceMode"] = o.InMaintenanceMode
	toSerialize["machineIds"] = o.MachineIds
	return toSerialize, nil
}

type NullableDeploymentIdMachinesBody struct {
	value *DeploymentIdMachinesBody
	isSet bool
}

func (v NullableDeploymentIdMachinesBody) Get() *DeploymentIdMachinesBody {
	return v.value
}

func (v *NullableDeploymentIdMachinesBody) Set(val *DeploymentIdMachinesBody) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentIdMachinesBody) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentIdMachinesBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentIdMachinesBody(val *DeploymentIdMachinesBody) *NullableDeploymentIdMachinesBody {
	return &NullableDeploymentIdMachinesBody{value: val, isSet: true}
}

func (v NullableDeploymentIdMachinesBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentIdMachinesBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

