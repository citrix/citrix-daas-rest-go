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

// checks if the AwsEdcAccountResourceEc2InstanceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsEdcAccountResourceEc2InstanceType{}

// AwsEdcAccountResourceEc2InstanceType struct for AwsEdcAccountResourceEc2InstanceType
type AwsEdcAccountResourceEc2InstanceType struct {
	AwsEdcAccountResource
	Name       NullableString `json:"name,omitempty"`
	VCpus      NullableInt32  `json:"vCpus,omitempty"`
	MemoryInMB NullableInt64  `json:"memoryInMB,omitempty"`
	IsDefault  *bool          `json:"isDefault,omitempty"`
}

type _AwsEdcAccountResourceEc2InstanceType AwsEdcAccountResourceEc2InstanceType

// NewAwsEdcAccountResourceEc2InstanceType instantiates a new AwsEdcAccountResourceEc2InstanceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsEdcAccountResourceEc2InstanceType(resourceType AwsAccountResourceType, accountType AccountType) *AwsEdcAccountResourceEc2InstanceType {
	this := AwsEdcAccountResourceEc2InstanceType{}
	this.ResourceType = resourceType
	this.AccountType = accountType
	return &this
}

// NewAwsEdcAccountResourceEc2InstanceTypeWithDefaults instantiates a new AwsEdcAccountResourceEc2InstanceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsEdcAccountResourceEc2InstanceTypeWithDefaults() *AwsEdcAccountResourceEc2InstanceType {
	this := AwsEdcAccountResourceEc2InstanceType{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceEc2InstanceType) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceEc2InstanceType) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceEc2InstanceType) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AwsEdcAccountResourceEc2InstanceType) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *AwsEdcAccountResourceEc2InstanceType) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AwsEdcAccountResourceEc2InstanceType) UnsetName() {
	o.Name.Unset()
}

// GetVCpus returns the VCpus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceEc2InstanceType) GetVCpus() int32 {
	if o == nil || IsNil(o.VCpus.Get()) {
		var ret int32
		return ret
	}
	return *o.VCpus.Get()
}

// GetVCpusOk returns a tuple with the VCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceEc2InstanceType) GetVCpusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VCpus.Get(), o.VCpus.IsSet()
}

// HasVCpus returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceEc2InstanceType) HasVCpus() bool {
	if o != nil && o.VCpus.IsSet() {
		return true
	}

	return false
}

// SetVCpus gets a reference to the given NullableInt32 and assigns it to the VCpus field.
func (o *AwsEdcAccountResourceEc2InstanceType) SetVCpus(v int32) {
	o.VCpus.Set(&v)
}

// SetVCpusNil sets the value for VCpus to be an explicit nil
func (o *AwsEdcAccountResourceEc2InstanceType) SetVCpusNil() {
	o.VCpus.Set(nil)
}

// UnsetVCpus ensures that no value is present for VCpus, not even an explicit nil
func (o *AwsEdcAccountResourceEc2InstanceType) UnsetVCpus() {
	o.VCpus.Unset()
}

// GetMemoryInMB returns the MemoryInMB field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceEc2InstanceType) GetMemoryInMB() int64 {
	if o == nil || IsNil(o.MemoryInMB.Get()) {
		var ret int64
		return ret
	}
	return *o.MemoryInMB.Get()
}

// GetMemoryInMBOk returns a tuple with the MemoryInMB field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceEc2InstanceType) GetMemoryInMBOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemoryInMB.Get(), o.MemoryInMB.IsSet()
}

// HasMemoryInMB returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceEc2InstanceType) HasMemoryInMB() bool {
	if o != nil && o.MemoryInMB.IsSet() {
		return true
	}

	return false
}

// SetMemoryInMB gets a reference to the given NullableInt64 and assigns it to the MemoryInMB field.
func (o *AwsEdcAccountResourceEc2InstanceType) SetMemoryInMB(v int64) {
	o.MemoryInMB.Set(&v)
}

// SetMemoryInMBNil sets the value for MemoryInMB to be an explicit nil
func (o *AwsEdcAccountResourceEc2InstanceType) SetMemoryInMBNil() {
	o.MemoryInMB.Set(nil)
}

// UnsetMemoryInMB ensures that no value is present for MemoryInMB, not even an explicit nil
func (o *AwsEdcAccountResourceEc2InstanceType) UnsetMemoryInMB() {
	o.MemoryInMB.Unset()
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *AwsEdcAccountResourceEc2InstanceType) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsEdcAccountResourceEc2InstanceType) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceEc2InstanceType) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *AwsEdcAccountResourceEc2InstanceType) SetIsDefault(v bool) {
	o.IsDefault = &v
}

func (o AwsEdcAccountResourceEc2InstanceType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsEdcAccountResourceEc2InstanceType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAwsEdcAccountResource, errAwsEdcAccountResource := json.Marshal(o.AwsEdcAccountResource)
	if errAwsEdcAccountResource != nil {
		return map[string]interface{}{}, errAwsEdcAccountResource
	}
	errAwsEdcAccountResource = json.Unmarshal([]byte(serializedAwsEdcAccountResource), &toSerialize)
	if errAwsEdcAccountResource != nil {
		return map[string]interface{}{}, errAwsEdcAccountResource
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.VCpus.IsSet() {
		toSerialize["vCpus"] = o.VCpus.Get()
	}
	if o.MemoryInMB.IsSet() {
		toSerialize["memoryInMB"] = o.MemoryInMB.Get()
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	return toSerialize, nil
}

type NullableAwsEdcAccountResourceEc2InstanceType struct {
	value *AwsEdcAccountResourceEc2InstanceType
	isSet bool
}

func (v NullableAwsEdcAccountResourceEc2InstanceType) Get() *AwsEdcAccountResourceEc2InstanceType {
	return v.value
}

func (v *NullableAwsEdcAccountResourceEc2InstanceType) Set(val *AwsEdcAccountResourceEc2InstanceType) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsEdcAccountResourceEc2InstanceType) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsEdcAccountResourceEc2InstanceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsEdcAccountResourceEc2InstanceType(val *AwsEdcAccountResourceEc2InstanceType) *NullableAwsEdcAccountResourceEc2InstanceType {
	return &NullableAwsEdcAccountResourceEc2InstanceType{value: val, isSet: true}
}

func (v NullableAwsEdcAccountResourceEc2InstanceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsEdcAccountResourceEc2InstanceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
