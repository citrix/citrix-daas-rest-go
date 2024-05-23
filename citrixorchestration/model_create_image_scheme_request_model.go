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

// checks if the CreateImageSchemeRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateImageSchemeRequestModel{}

// CreateImageSchemeRequestModel Request object for creation of image scheme.
type CreateImageSchemeRequestModel struct {
	// The hypervisor resource path of the Cloud service offering to use when creating machines.
	ServiceOfferingPath NullableString `json:"ServiceOfferingPath,omitempty"`
	// The path in the resource pool to the virtual machine template that will be used. This identifies the VM template to be used and the default values for the tags, virtual machine size, boot diagnostics, host cache property of OS disk, accelerated networking and availability zone. This must be a path to a Virtual machine or Template item in the resource pool to which the Image Version Specification is associated.
	MachineProfile NullableString `json:"MachineProfile,omitempty"`
	// The number of processors that virtual machines created from the image preparing should use.
	CpuCount NullableInt32 `json:"CpuCount,omitempty"`
	// The maximum amount of memory that virtual machines created from the image preparing should use.
	MemoryMB NullableInt32 `json:"MemoryMB,omitempty"`
	// Specifies how the attached NICs are mapped to networks.  If this parameter is omitted, provisioned VMs are created with a single NIC, which is mapped to the default network in the hypervisor resource pool.  If this parameter is supplied, machines are created with the number of NICs specified in the map, and each NIC is attached to the specified network.
	NetworkMapping []NetworkMapRequestModel `json:"NetworkMapping,omitempty"`
	// The properties of the image version specification that are specific to the target hosting infrastructure.
	CustomProperties []NameValueStringPairModel `json:"CustomProperties,omitempty"`
}

// NewCreateImageSchemeRequestModel instantiates a new CreateImageSchemeRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateImageSchemeRequestModel() *CreateImageSchemeRequestModel {
	this := CreateImageSchemeRequestModel{}
	return &this
}

// NewCreateImageSchemeRequestModelWithDefaults instantiates a new CreateImageSchemeRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateImageSchemeRequestModelWithDefaults() *CreateImageSchemeRequestModel {
	this := CreateImageSchemeRequestModel{}
	return &this
}

// GetServiceOfferingPath returns the ServiceOfferingPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateImageSchemeRequestModel) GetServiceOfferingPath() string {
	if o == nil || IsNil(o.ServiceOfferingPath.Get()) {
		var ret string
		return ret
	}
	return *o.ServiceOfferingPath.Get()
}

// GetServiceOfferingPathOk returns a tuple with the ServiceOfferingPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateImageSchemeRequestModel) GetServiceOfferingPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceOfferingPath.Get(), o.ServiceOfferingPath.IsSet()
}

// HasServiceOfferingPath returns a boolean if a field has been set.
func (o *CreateImageSchemeRequestModel) HasServiceOfferingPath() bool {
	if o != nil && o.ServiceOfferingPath.IsSet() {
		return true
	}

	return false
}

// SetServiceOfferingPath gets a reference to the given NullableString and assigns it to the ServiceOfferingPath field.
func (o *CreateImageSchemeRequestModel) SetServiceOfferingPath(v string) {
	o.ServiceOfferingPath.Set(&v)
}
// SetServiceOfferingPathNil sets the value for ServiceOfferingPath to be an explicit nil
func (o *CreateImageSchemeRequestModel) SetServiceOfferingPathNil() {
	o.ServiceOfferingPath.Set(nil)
}

// UnsetServiceOfferingPath ensures that no value is present for ServiceOfferingPath, not even an explicit nil
func (o *CreateImageSchemeRequestModel) UnsetServiceOfferingPath() {
	o.ServiceOfferingPath.Unset()
}

// GetMachineProfile returns the MachineProfile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateImageSchemeRequestModel) GetMachineProfile() string {
	if o == nil || IsNil(o.MachineProfile.Get()) {
		var ret string
		return ret
	}
	return *o.MachineProfile.Get()
}

// GetMachineProfileOk returns a tuple with the MachineProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateImageSchemeRequestModel) GetMachineProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MachineProfile.Get(), o.MachineProfile.IsSet()
}

// HasMachineProfile returns a boolean if a field has been set.
func (o *CreateImageSchemeRequestModel) HasMachineProfile() bool {
	if o != nil && o.MachineProfile.IsSet() {
		return true
	}

	return false
}

// SetMachineProfile gets a reference to the given NullableString and assigns it to the MachineProfile field.
func (o *CreateImageSchemeRequestModel) SetMachineProfile(v string) {
	o.MachineProfile.Set(&v)
}
// SetMachineProfileNil sets the value for MachineProfile to be an explicit nil
func (o *CreateImageSchemeRequestModel) SetMachineProfileNil() {
	o.MachineProfile.Set(nil)
}

// UnsetMachineProfile ensures that no value is present for MachineProfile, not even an explicit nil
func (o *CreateImageSchemeRequestModel) UnsetMachineProfile() {
	o.MachineProfile.Unset()
}

// GetCpuCount returns the CpuCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateImageSchemeRequestModel) GetCpuCount() int32 {
	if o == nil || IsNil(o.CpuCount.Get()) {
		var ret int32
		return ret
	}
	return *o.CpuCount.Get()
}

// GetCpuCountOk returns a tuple with the CpuCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateImageSchemeRequestModel) GetCpuCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpuCount.Get(), o.CpuCount.IsSet()
}

// HasCpuCount returns a boolean if a field has been set.
func (o *CreateImageSchemeRequestModel) HasCpuCount() bool {
	if o != nil && o.CpuCount.IsSet() {
		return true
	}

	return false
}

// SetCpuCount gets a reference to the given NullableInt32 and assigns it to the CpuCount field.
func (o *CreateImageSchemeRequestModel) SetCpuCount(v int32) {
	o.CpuCount.Set(&v)
}
// SetCpuCountNil sets the value for CpuCount to be an explicit nil
func (o *CreateImageSchemeRequestModel) SetCpuCountNil() {
	o.CpuCount.Set(nil)
}

// UnsetCpuCount ensures that no value is present for CpuCount, not even an explicit nil
func (o *CreateImageSchemeRequestModel) UnsetCpuCount() {
	o.CpuCount.Unset()
}

// GetMemoryMB returns the MemoryMB field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateImageSchemeRequestModel) GetMemoryMB() int32 {
	if o == nil || IsNil(o.MemoryMB.Get()) {
		var ret int32
		return ret
	}
	return *o.MemoryMB.Get()
}

// GetMemoryMBOk returns a tuple with the MemoryMB field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateImageSchemeRequestModel) GetMemoryMBOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemoryMB.Get(), o.MemoryMB.IsSet()
}

// HasMemoryMB returns a boolean if a field has been set.
func (o *CreateImageSchemeRequestModel) HasMemoryMB() bool {
	if o != nil && o.MemoryMB.IsSet() {
		return true
	}

	return false
}

// SetMemoryMB gets a reference to the given NullableInt32 and assigns it to the MemoryMB field.
func (o *CreateImageSchemeRequestModel) SetMemoryMB(v int32) {
	o.MemoryMB.Set(&v)
}
// SetMemoryMBNil sets the value for MemoryMB to be an explicit nil
func (o *CreateImageSchemeRequestModel) SetMemoryMBNil() {
	o.MemoryMB.Set(nil)
}

// UnsetMemoryMB ensures that no value is present for MemoryMB, not even an explicit nil
func (o *CreateImageSchemeRequestModel) UnsetMemoryMB() {
	o.MemoryMB.Unset()
}

// GetNetworkMapping returns the NetworkMapping field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateImageSchemeRequestModel) GetNetworkMapping() []NetworkMapRequestModel {
	if o == nil {
		var ret []NetworkMapRequestModel
		return ret
	}
	return o.NetworkMapping
}

// GetNetworkMappingOk returns a tuple with the NetworkMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateImageSchemeRequestModel) GetNetworkMappingOk() ([]NetworkMapRequestModel, bool) {
	if o == nil || IsNil(o.NetworkMapping) {
		return nil, false
	}
	return o.NetworkMapping, true
}

// HasNetworkMapping returns a boolean if a field has been set.
func (o *CreateImageSchemeRequestModel) HasNetworkMapping() bool {
	if o != nil && IsNil(o.NetworkMapping) {
		return true
	}

	return false
}

// SetNetworkMapping gets a reference to the given []NetworkMapRequestModel and assigns it to the NetworkMapping field.
func (o *CreateImageSchemeRequestModel) SetNetworkMapping(v []NetworkMapRequestModel) {
	o.NetworkMapping = v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateImageSchemeRequestModel) GetCustomProperties() []NameValueStringPairModel {
	if o == nil {
		var ret []NameValueStringPairModel
		return ret
	}
	return o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateImageSchemeRequestModel) GetCustomPropertiesOk() ([]NameValueStringPairModel, bool) {
	if o == nil || IsNil(o.CustomProperties) {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *CreateImageSchemeRequestModel) HasCustomProperties() bool {
	if o != nil && IsNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given []NameValueStringPairModel and assigns it to the CustomProperties field.
func (o *CreateImageSchemeRequestModel) SetCustomProperties(v []NameValueStringPairModel) {
	o.CustomProperties = v
}

func (o CreateImageSchemeRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateImageSchemeRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceOfferingPath.IsSet() {
		toSerialize["ServiceOfferingPath"] = o.ServiceOfferingPath.Get()
	}
	if o.MachineProfile.IsSet() {
		toSerialize["MachineProfile"] = o.MachineProfile.Get()
	}
	if o.CpuCount.IsSet() {
		toSerialize["CpuCount"] = o.CpuCount.Get()
	}
	if o.MemoryMB.IsSet() {
		toSerialize["MemoryMB"] = o.MemoryMB.Get()
	}
	if o.NetworkMapping != nil {
		toSerialize["NetworkMapping"] = o.NetworkMapping
	}
	if o.CustomProperties != nil {
		toSerialize["CustomProperties"] = o.CustomProperties
	}
	return toSerialize, nil
}

type NullableCreateImageSchemeRequestModel struct {
	value *CreateImageSchemeRequestModel
	isSet bool
}

func (v NullableCreateImageSchemeRequestModel) Get() *CreateImageSchemeRequestModel {
	return v.value
}

func (v *NullableCreateImageSchemeRequestModel) Set(val *CreateImageSchemeRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateImageSchemeRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateImageSchemeRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateImageSchemeRequestModel(val *CreateImageSchemeRequestModel) *NullableCreateImageSchemeRequestModel {
	return &NullableCreateImageSchemeRequestModel{value: val, isSet: true}
}

func (v NullableCreateImageSchemeRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateImageSchemeRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

