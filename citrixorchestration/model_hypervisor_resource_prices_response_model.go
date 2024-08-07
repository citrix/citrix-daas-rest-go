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

// checks if the HypervisorResourcePricesResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HypervisorResourcePricesResponseModel{}

// HypervisorResourcePricesResponseModel Response model for the pricing information of different hypervisor resource settings.
type HypervisorResourcePricesResponseModel struct {
	// The pricing information for the virtual machines settings.
	VmPrices []VmPriceResponseModel `json:"VmPrices,omitempty"`
	// The pricing information for the disk settings.
	DiskPrices []DiskPriceResponseModel `json:"DiskPrices,omitempty"`
	// The warning message.
	WarningMessage NullableString `json:"WarningMessage,omitempty"`
}

// NewHypervisorResourcePricesResponseModel instantiates a new HypervisorResourcePricesResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervisorResourcePricesResponseModel() *HypervisorResourcePricesResponseModel {
	this := HypervisorResourcePricesResponseModel{}
	return &this
}

// NewHypervisorResourcePricesResponseModelWithDefaults instantiates a new HypervisorResourcePricesResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervisorResourcePricesResponseModelWithDefaults() *HypervisorResourcePricesResponseModel {
	this := HypervisorResourcePricesResponseModel{}
	return &this
}

// GetVmPrices returns the VmPrices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorResourcePricesResponseModel) GetVmPrices() []VmPriceResponseModel {
	if o == nil {
		var ret []VmPriceResponseModel
		return ret
	}
	return o.VmPrices
}

// GetVmPricesOk returns a tuple with the VmPrices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorResourcePricesResponseModel) GetVmPricesOk() ([]VmPriceResponseModel, bool) {
	if o == nil || IsNil(o.VmPrices) {
		return nil, false
	}
	return o.VmPrices, true
}

// HasVmPrices returns a boolean if a field has been set.
func (o *HypervisorResourcePricesResponseModel) HasVmPrices() bool {
	if o != nil && IsNil(o.VmPrices) {
		return true
	}

	return false
}

// SetVmPrices gets a reference to the given []VmPriceResponseModel and assigns it to the VmPrices field.
func (o *HypervisorResourcePricesResponseModel) SetVmPrices(v []VmPriceResponseModel) {
	o.VmPrices = v
}

// GetDiskPrices returns the DiskPrices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorResourcePricesResponseModel) GetDiskPrices() []DiskPriceResponseModel {
	if o == nil {
		var ret []DiskPriceResponseModel
		return ret
	}
	return o.DiskPrices
}

// GetDiskPricesOk returns a tuple with the DiskPrices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorResourcePricesResponseModel) GetDiskPricesOk() ([]DiskPriceResponseModel, bool) {
	if o == nil || IsNil(o.DiskPrices) {
		return nil, false
	}
	return o.DiskPrices, true
}

// HasDiskPrices returns a boolean if a field has been set.
func (o *HypervisorResourcePricesResponseModel) HasDiskPrices() bool {
	if o != nil && IsNil(o.DiskPrices) {
		return true
	}

	return false
}

// SetDiskPrices gets a reference to the given []DiskPriceResponseModel and assigns it to the DiskPrices field.
func (o *HypervisorResourcePricesResponseModel) SetDiskPrices(v []DiskPriceResponseModel) {
	o.DiskPrices = v
}

// GetWarningMessage returns the WarningMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HypervisorResourcePricesResponseModel) GetWarningMessage() string {
	if o == nil || IsNil(o.WarningMessage.Get()) {
		var ret string
		return ret
	}
	return *o.WarningMessage.Get()
}

// GetWarningMessageOk returns a tuple with the WarningMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HypervisorResourcePricesResponseModel) GetWarningMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WarningMessage.Get(), o.WarningMessage.IsSet()
}

// HasWarningMessage returns a boolean if a field has been set.
func (o *HypervisorResourcePricesResponseModel) HasWarningMessage() bool {
	if o != nil && o.WarningMessage.IsSet() {
		return true
	}

	return false
}

// SetWarningMessage gets a reference to the given NullableString and assigns it to the WarningMessage field.
func (o *HypervisorResourcePricesResponseModel) SetWarningMessage(v string) {
	o.WarningMessage.Set(&v)
}
// SetWarningMessageNil sets the value for WarningMessage to be an explicit nil
func (o *HypervisorResourcePricesResponseModel) SetWarningMessageNil() {
	o.WarningMessage.Set(nil)
}

// UnsetWarningMessage ensures that no value is present for WarningMessage, not even an explicit nil
func (o *HypervisorResourcePricesResponseModel) UnsetWarningMessage() {
	o.WarningMessage.Unset()
}

func (o HypervisorResourcePricesResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HypervisorResourcePricesResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.VmPrices != nil {
		toSerialize["VmPrices"] = o.VmPrices
	}
	if o.DiskPrices != nil {
		toSerialize["DiskPrices"] = o.DiskPrices
	}
	if o.WarningMessage.IsSet() {
		toSerialize["WarningMessage"] = o.WarningMessage.Get()
	}
	return toSerialize, nil
}

type NullableHypervisorResourcePricesResponseModel struct {
	value *HypervisorResourcePricesResponseModel
	isSet bool
}

func (v NullableHypervisorResourcePricesResponseModel) Get() *HypervisorResourcePricesResponseModel {
	return v.value
}

func (v *NullableHypervisorResourcePricesResponseModel) Set(val *HypervisorResourcePricesResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorResourcePricesResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorResourcePricesResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorResourcePricesResponseModel(val *HypervisorResourcePricesResponseModel) *NullableHypervisorResourcePricesResponseModel {
	return &NullableHypervisorResourcePricesResponseModel{value: val, isSet: true}
}

func (v NullableHypervisorResourcePricesResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorResourcePricesResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

