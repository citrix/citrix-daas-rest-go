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

// checks if the RemotePCEnrollmentScopeRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemotePCEnrollmentScopeRequestModel{}

// RemotePCEnrollmentScopeRequestModel Enrollment scope for remote PCs.
type RemotePCEnrollmentScopeRequestModel struct {
	// Specifies the DN of an AD container containing machines allowed to enroll as remote PCs.
	OU string `json:"OU" validate:"regexp=.*|any"`
	// Indicates whether machines in subfolders of OU are allowed to enroll as remote PCs.
	IncludeSubfolders NullableBool `json:"IncludeSubfolders,omitempty"`
	// Indicates whether this objet is for a OU or for a machine
	IsOrganizationalUnit NullableBool `json:"IsOrganizationalUnit,omitempty"`
	// Machines which are explicitly excluded from matching the enrollment scope.
	MachinesExcluded []string `json:"MachinesExcluded,omitempty"`
	// Machines which are included in the enrollment scope.
	MachinesIncluded []string `json:"MachinesIncluded,omitempty"`
	// The user(s) to whom this machine will be assigned. Machines can be assigned to multiple users, a single client IP address, or a single client name, but only to one of these categories at a time.
	AssignedUsers []string `json:"AssignedUsers,omitempty"`
}

// NewRemotePCEnrollmentScopeRequestModel instantiates a new RemotePCEnrollmentScopeRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemotePCEnrollmentScopeRequestModel(oU string) *RemotePCEnrollmentScopeRequestModel {
	this := RemotePCEnrollmentScopeRequestModel{}
	this.OU = oU
	var includeSubfolders bool = false
	this.IncludeSubfolders = *NewNullableBool(&includeSubfolders)
	var isOrganizationalUnit bool = false
	this.IsOrganizationalUnit = *NewNullableBool(&isOrganizationalUnit)
	return &this
}

// NewRemotePCEnrollmentScopeRequestModelWithDefaults instantiates a new RemotePCEnrollmentScopeRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemotePCEnrollmentScopeRequestModelWithDefaults() *RemotePCEnrollmentScopeRequestModel {
	this := RemotePCEnrollmentScopeRequestModel{}
	var includeSubfolders bool = false
	this.IncludeSubfolders = *NewNullableBool(&includeSubfolders)
	var isOrganizationalUnit bool = false
	this.IsOrganizationalUnit = *NewNullableBool(&isOrganizationalUnit)
	return &this
}

// GetOU returns the OU field value
func (o *RemotePCEnrollmentScopeRequestModel) GetOU() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OU
}

// GetOUOk returns a tuple with the OU field value
// and a boolean to check if the value has been set.
func (o *RemotePCEnrollmentScopeRequestModel) GetOUOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OU, true
}

// SetOU sets field value
func (o *RemotePCEnrollmentScopeRequestModel) SetOU(v string) {
	o.OU = v
}

// GetIncludeSubfolders returns the IncludeSubfolders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemotePCEnrollmentScopeRequestModel) GetIncludeSubfolders() bool {
	if o == nil || IsNil(o.IncludeSubfolders.Get()) {
		var ret bool
		return ret
	}
	return *o.IncludeSubfolders.Get()
}

// GetIncludeSubfoldersOk returns a tuple with the IncludeSubfolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemotePCEnrollmentScopeRequestModel) GetIncludeSubfoldersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncludeSubfolders.Get(), o.IncludeSubfolders.IsSet()
}

// HasIncludeSubfolders returns a boolean if a field has been set.
func (o *RemotePCEnrollmentScopeRequestModel) HasIncludeSubfolders() bool {
	if o != nil && o.IncludeSubfolders.IsSet() {
		return true
	}

	return false
}

// SetIncludeSubfolders gets a reference to the given NullableBool and assigns it to the IncludeSubfolders field.
func (o *RemotePCEnrollmentScopeRequestModel) SetIncludeSubfolders(v bool) {
	o.IncludeSubfolders.Set(&v)
}

// SetIncludeSubfoldersNil sets the value for IncludeSubfolders to be an explicit nil
func (o *RemotePCEnrollmentScopeRequestModel) SetIncludeSubfoldersNil() {
	o.IncludeSubfolders.Set(nil)
}

// UnsetIncludeSubfolders ensures that no value is present for IncludeSubfolders, not even an explicit nil
func (o *RemotePCEnrollmentScopeRequestModel) UnsetIncludeSubfolders() {
	o.IncludeSubfolders.Unset()
}

// GetIsOrganizationalUnit returns the IsOrganizationalUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemotePCEnrollmentScopeRequestModel) GetIsOrganizationalUnit() bool {
	if o == nil || IsNil(o.IsOrganizationalUnit.Get()) {
		var ret bool
		return ret
	}
	return *o.IsOrganizationalUnit.Get()
}

// GetIsOrganizationalUnitOk returns a tuple with the IsOrganizationalUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemotePCEnrollmentScopeRequestModel) GetIsOrganizationalUnitOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsOrganizationalUnit.Get(), o.IsOrganizationalUnit.IsSet()
}

// HasIsOrganizationalUnit returns a boolean if a field has been set.
func (o *RemotePCEnrollmentScopeRequestModel) HasIsOrganizationalUnit() bool {
	if o != nil && o.IsOrganizationalUnit.IsSet() {
		return true
	}

	return false
}

// SetIsOrganizationalUnit gets a reference to the given NullableBool and assigns it to the IsOrganizationalUnit field.
func (o *RemotePCEnrollmentScopeRequestModel) SetIsOrganizationalUnit(v bool) {
	o.IsOrganizationalUnit.Set(&v)
}

// SetIsOrganizationalUnitNil sets the value for IsOrganizationalUnit to be an explicit nil
func (o *RemotePCEnrollmentScopeRequestModel) SetIsOrganizationalUnitNil() {
	o.IsOrganizationalUnit.Set(nil)
}

// UnsetIsOrganizationalUnit ensures that no value is present for IsOrganizationalUnit, not even an explicit nil
func (o *RemotePCEnrollmentScopeRequestModel) UnsetIsOrganizationalUnit() {
	o.IsOrganizationalUnit.Unset()
}

// GetMachinesExcluded returns the MachinesExcluded field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemotePCEnrollmentScopeRequestModel) GetMachinesExcluded() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MachinesExcluded
}

// GetMachinesExcludedOk returns a tuple with the MachinesExcluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemotePCEnrollmentScopeRequestModel) GetMachinesExcludedOk() ([]string, bool) {
	if o == nil || IsNil(o.MachinesExcluded) {
		return nil, false
	}
	return o.MachinesExcluded, true
}

// HasMachinesExcluded returns a boolean if a field has been set.
func (o *RemotePCEnrollmentScopeRequestModel) HasMachinesExcluded() bool {
	if o != nil && IsNil(o.MachinesExcluded) {
		return true
	}

	return false
}

// SetMachinesExcluded gets a reference to the given []string and assigns it to the MachinesExcluded field.
func (o *RemotePCEnrollmentScopeRequestModel) SetMachinesExcluded(v []string) {
	o.MachinesExcluded = v
}

// GetMachinesIncluded returns the MachinesIncluded field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemotePCEnrollmentScopeRequestModel) GetMachinesIncluded() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MachinesIncluded
}

// GetMachinesIncludedOk returns a tuple with the MachinesIncluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemotePCEnrollmentScopeRequestModel) GetMachinesIncludedOk() ([]string, bool) {
	if o == nil || IsNil(o.MachinesIncluded) {
		return nil, false
	}
	return o.MachinesIncluded, true
}

// HasMachinesIncluded returns a boolean if a field has been set.
func (o *RemotePCEnrollmentScopeRequestModel) HasMachinesIncluded() bool {
	if o != nil && IsNil(o.MachinesIncluded) {
		return true
	}

	return false
}

// SetMachinesIncluded gets a reference to the given []string and assigns it to the MachinesIncluded field.
func (o *RemotePCEnrollmentScopeRequestModel) SetMachinesIncluded(v []string) {
	o.MachinesIncluded = v
}

// GetAssignedUsers returns the AssignedUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemotePCEnrollmentScopeRequestModel) GetAssignedUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AssignedUsers
}

// GetAssignedUsersOk returns a tuple with the AssignedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemotePCEnrollmentScopeRequestModel) GetAssignedUsersOk() ([]string, bool) {
	if o == nil || IsNil(o.AssignedUsers) {
		return nil, false
	}
	return o.AssignedUsers, true
}

// HasAssignedUsers returns a boolean if a field has been set.
func (o *RemotePCEnrollmentScopeRequestModel) HasAssignedUsers() bool {
	if o != nil && IsNil(o.AssignedUsers) {
		return true
	}

	return false
}

// SetAssignedUsers gets a reference to the given []string and assigns it to the AssignedUsers field.
func (o *RemotePCEnrollmentScopeRequestModel) SetAssignedUsers(v []string) {
	o.AssignedUsers = v
}

func (o RemotePCEnrollmentScopeRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemotePCEnrollmentScopeRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["OU"] = o.OU
	if o.IncludeSubfolders.IsSet() {
		toSerialize["IncludeSubfolders"] = o.IncludeSubfolders.Get()
	}
	if o.IsOrganizationalUnit.IsSet() {
		toSerialize["IsOrganizationalUnit"] = o.IsOrganizationalUnit.Get()
	}
	if o.MachinesExcluded != nil {
		toSerialize["MachinesExcluded"] = o.MachinesExcluded
	}
	if o.MachinesIncluded != nil {
		toSerialize["MachinesIncluded"] = o.MachinesIncluded
	}
	if o.AssignedUsers != nil {
		toSerialize["AssignedUsers"] = o.AssignedUsers
	}
	return toSerialize, nil
}

type NullableRemotePCEnrollmentScopeRequestModel struct {
	value *RemotePCEnrollmentScopeRequestModel
	isSet bool
}

func (v NullableRemotePCEnrollmentScopeRequestModel) Get() *RemotePCEnrollmentScopeRequestModel {
	return v.value
}

func (v *NullableRemotePCEnrollmentScopeRequestModel) Set(val *RemotePCEnrollmentScopeRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRemotePCEnrollmentScopeRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRemotePCEnrollmentScopeRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemotePCEnrollmentScopeRequestModel(val *RemotePCEnrollmentScopeRequestModel) *NullableRemotePCEnrollmentScopeRequestModel {
	return &NullableRemotePCEnrollmentScopeRequestModel{value: val, isSet: true}
}

func (v NullableRemotePCEnrollmentScopeRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemotePCEnrollmentScopeRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
