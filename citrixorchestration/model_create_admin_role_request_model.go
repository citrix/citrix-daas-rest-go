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

// checks if the CreateAdminRoleRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAdminRoleRequestModel{}

// CreateAdminRoleRequestModel Request object for creation of admin scope.
type CreateAdminRoleRequestModel struct {
	// The name of the role. Name is globally unique.
	Name string `json:"Name"`
	// The description of the admin role.
	Description NullableString `json:"Description,omitempty"`
	// Indicates whether the role has access to the Manage tab in Citrix Cloud.
	CanLaunchManage NullableBool `json:"CanLaunchManage,omitempty"`
	// Indicates whether the role has access to the Monitor tab in Citrix Cloud.
	CanLaunchMonitor NullableBool `json:"CanLaunchMonitor,omitempty"`
	// List of permissions granted by the role. At least one permission is required.
	Permissions []string `json:"Permissions,omitempty"`
}

// NewCreateAdminRoleRequestModel instantiates a new CreateAdminRoleRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAdminRoleRequestModel(name string) *CreateAdminRoleRequestModel {
	this := CreateAdminRoleRequestModel{}
	this.Name = name
	return &this
}

// NewCreateAdminRoleRequestModelWithDefaults instantiates a new CreateAdminRoleRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAdminRoleRequestModelWithDefaults() *CreateAdminRoleRequestModel {
	this := CreateAdminRoleRequestModel{}
	return &this
}

// GetName returns the Name field value
func (o *CreateAdminRoleRequestModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateAdminRoleRequestModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateAdminRoleRequestModel) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAdminRoleRequestModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAdminRoleRequestModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateAdminRoleRequestModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CreateAdminRoleRequestModel) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CreateAdminRoleRequestModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CreateAdminRoleRequestModel) UnsetDescription() {
	o.Description.Unset()
}

// GetCanLaunchManage returns the CanLaunchManage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAdminRoleRequestModel) GetCanLaunchManage() bool {
	if o == nil || IsNil(o.CanLaunchManage.Get()) {
		var ret bool
		return ret
	}
	return *o.CanLaunchManage.Get()
}

// GetCanLaunchManageOk returns a tuple with the CanLaunchManage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAdminRoleRequestModel) GetCanLaunchManageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanLaunchManage.Get(), o.CanLaunchManage.IsSet()
}

// HasCanLaunchManage returns a boolean if a field has been set.
func (o *CreateAdminRoleRequestModel) HasCanLaunchManage() bool {
	if o != nil && o.CanLaunchManage.IsSet() {
		return true
	}

	return false
}

// SetCanLaunchManage gets a reference to the given NullableBool and assigns it to the CanLaunchManage field.
func (o *CreateAdminRoleRequestModel) SetCanLaunchManage(v bool) {
	o.CanLaunchManage.Set(&v)
}

// SetCanLaunchManageNil sets the value for CanLaunchManage to be an explicit nil
func (o *CreateAdminRoleRequestModel) SetCanLaunchManageNil() {
	o.CanLaunchManage.Set(nil)
}

// UnsetCanLaunchManage ensures that no value is present for CanLaunchManage, not even an explicit nil
func (o *CreateAdminRoleRequestModel) UnsetCanLaunchManage() {
	o.CanLaunchManage.Unset()
}

// GetCanLaunchMonitor returns the CanLaunchMonitor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAdminRoleRequestModel) GetCanLaunchMonitor() bool {
	if o == nil || IsNil(o.CanLaunchMonitor.Get()) {
		var ret bool
		return ret
	}
	return *o.CanLaunchMonitor.Get()
}

// GetCanLaunchMonitorOk returns a tuple with the CanLaunchMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAdminRoleRequestModel) GetCanLaunchMonitorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanLaunchMonitor.Get(), o.CanLaunchMonitor.IsSet()
}

// HasCanLaunchMonitor returns a boolean if a field has been set.
func (o *CreateAdminRoleRequestModel) HasCanLaunchMonitor() bool {
	if o != nil && o.CanLaunchMonitor.IsSet() {
		return true
	}

	return false
}

// SetCanLaunchMonitor gets a reference to the given NullableBool and assigns it to the CanLaunchMonitor field.
func (o *CreateAdminRoleRequestModel) SetCanLaunchMonitor(v bool) {
	o.CanLaunchMonitor.Set(&v)
}

// SetCanLaunchMonitorNil sets the value for CanLaunchMonitor to be an explicit nil
func (o *CreateAdminRoleRequestModel) SetCanLaunchMonitorNil() {
	o.CanLaunchMonitor.Set(nil)
}

// UnsetCanLaunchMonitor ensures that no value is present for CanLaunchMonitor, not even an explicit nil
func (o *CreateAdminRoleRequestModel) UnsetCanLaunchMonitor() {
	o.CanLaunchMonitor.Unset()
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAdminRoleRequestModel) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAdminRoleRequestModel) GetPermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *CreateAdminRoleRequestModel) HasPermissions() bool {
	if o != nil && IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *CreateAdminRoleRequestModel) SetPermissions(v []string) {
	o.Permissions = v
}

func (o CreateAdminRoleRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAdminRoleRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if o.CanLaunchManage.IsSet() {
		toSerialize["CanLaunchManage"] = o.CanLaunchManage.Get()
	}
	if o.CanLaunchMonitor.IsSet() {
		toSerialize["CanLaunchMonitor"] = o.CanLaunchMonitor.Get()
	}
	if o.Permissions != nil {
		toSerialize["Permissions"] = o.Permissions
	}
	return toSerialize, nil
}

type NullableCreateAdminRoleRequestModel struct {
	value *CreateAdminRoleRequestModel
	isSet bool
}

func (v NullableCreateAdminRoleRequestModel) Get() *CreateAdminRoleRequestModel {
	return v.value
}

func (v *NullableCreateAdminRoleRequestModel) Set(val *CreateAdminRoleRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAdminRoleRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAdminRoleRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAdminRoleRequestModel(val *CreateAdminRoleRequestModel) *NullableCreateAdminRoleRequestModel {
	return &NullableCreateAdminRoleRequestModel{value: val, isSet: true}
}

func (v NullableCreateAdminRoleRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAdminRoleRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
