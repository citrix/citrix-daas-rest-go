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

// checks if the AddCatalogMachineAssignment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCatalogMachineAssignment{}

// AddCatalogMachineAssignment struct for AddCatalogMachineAssignment
type AddCatalogMachineAssignment struct {
	// Name of the machine (in UPN format) to be assigned to a catalog.
	MachineName *string `json:"machineName,omitempty"`
	// List of users to assign to the machine.
	AssignedUsers []string `json:"assignedUsers,omitempty"`
}

// NewAddCatalogMachineAssignmentWithDefaults instantiates a new AddCatalogMachineAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCatalogMachineAssignmentWithDefaults() *AddCatalogMachineAssignment {
	this := AddCatalogMachineAssignment{}
	return &this
}

// GetMachineName returns the MachineName field value if set, zero value otherwise.
func (o *AddCatalogMachineAssignment) GetMachineName() string {
	if o == nil || IsNil(o.MachineName) {
		var ret string
		return ret
	}
	return *o.MachineName
}

// GetMachineNameOk returns a tuple with the MachineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCatalogMachineAssignment) GetMachineNameOk() (*string, bool) {
	if o == nil || IsNil(o.MachineName) {
		return nil, false
	}
	return o.MachineName, true
}

// SetMachineName gets a reference to the given string and assigns it to the MachineName field.
func (o *AddCatalogMachineAssignment) SetMachineName(v string) {
	o.MachineName = &v
}

// GetAssignedUsers returns the AssignedUsers field value if set, zero value otherwise.
func (o *AddCatalogMachineAssignment) GetAssignedUsers() []string {
	if o == nil || IsNil(o.AssignedUsers) {
		var ret []string
		return ret
	}
	return o.AssignedUsers
}

// GetAssignedUsersOk returns a tuple with the AssignedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCatalogMachineAssignment) GetAssignedUsersOk() ([]string, bool) {
	if o == nil || IsNil(o.AssignedUsers) {
		return nil, false
	}
	return o.AssignedUsers, true
}

// SetAssignedUsers gets a reference to the given []string and assigns it to the AssignedUsers field.
func (o *AddCatalogMachineAssignment) SetAssignedUsers(v []string) {
	o.AssignedUsers = v
}

func (o AddCatalogMachineAssignment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCatalogMachineAssignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MachineName) {
		toSerialize["machineName"] = o.MachineName
	}
	if !IsNil(o.AssignedUsers) {
		toSerialize["assignedUsers"] = o.AssignedUsers
	}
	return toSerialize, nil
}

type NullableAddCatalogMachineAssignment struct {
	value *AddCatalogMachineAssignment
	isSet bool
}

func (v NullableAddCatalogMachineAssignment) Get() *AddCatalogMachineAssignment {
	return v.value
}

func (v *NullableAddCatalogMachineAssignment) Set(val *AddCatalogMachineAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCatalogMachineAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCatalogMachineAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCatalogMachineAssignment(val *AddCatalogMachineAssignment) *NullableAddCatalogMachineAssignment {
	return &NullableAddCatalogMachineAssignment{value: val, isSet: true}
}

func (v NullableAddCatalogMachineAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCatalogMachineAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
