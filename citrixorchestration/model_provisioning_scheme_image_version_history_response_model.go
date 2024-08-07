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

// checks if the ProvisioningSchemeImageVersionHistoryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningSchemeImageVersionHistoryResponseModel{}

// ProvisioningSchemeImageVersionHistoryResponseModel Details about an image version specification used for a provisioning scheme.
type ProvisioningSchemeImageVersionHistoryResponseModel struct {
	ImageVersion ImageVersionRefResponseModel `json:"ImageVersion"`
	// The date and time when the snapshot was used in the provisioning scheme.
	Date NullableString `json:"Date,omitempty"`
}

// NewProvisioningSchemeImageVersionHistoryResponseModel instantiates a new ProvisioningSchemeImageVersionHistoryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningSchemeImageVersionHistoryResponseModel(imageVersion ImageVersionRefResponseModel) *ProvisioningSchemeImageVersionHistoryResponseModel {
	this := ProvisioningSchemeImageVersionHistoryResponseModel{}
	this.ImageVersion = imageVersion
	return &this
}

// NewProvisioningSchemeImageVersionHistoryResponseModelWithDefaults instantiates a new ProvisioningSchemeImageVersionHistoryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningSchemeImageVersionHistoryResponseModelWithDefaults() *ProvisioningSchemeImageVersionHistoryResponseModel {
	this := ProvisioningSchemeImageVersionHistoryResponseModel{}
	return &this
}

// GetImageVersion returns the ImageVersion field value
func (o *ProvisioningSchemeImageVersionHistoryResponseModel) GetImageVersion() ImageVersionRefResponseModel {
	if o == nil {
		var ret ImageVersionRefResponseModel
		return ret
	}

	return o.ImageVersion
}

// GetImageVersionOk returns a tuple with the ImageVersion field value
// and a boolean to check if the value has been set.
func (o *ProvisioningSchemeImageVersionHistoryResponseModel) GetImageVersionOk() (*ImageVersionRefResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageVersion, true
}

// SetImageVersion sets field value
func (o *ProvisioningSchemeImageVersionHistoryResponseModel) SetImageVersion(v ImageVersionRefResponseModel) {
	o.ImageVersion = v
}

// GetDate returns the Date field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisioningSchemeImageVersionHistoryResponseModel) GetDate() string {
	if o == nil || IsNil(o.Date.Get()) {
		var ret string
		return ret
	}
	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisioningSchemeImageVersionHistoryResponseModel) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// HasDate returns a boolean if a field has been set.
func (o *ProvisioningSchemeImageVersionHistoryResponseModel) HasDate() bool {
	if o != nil && o.Date.IsSet() {
		return true
	}

	return false
}

// SetDate gets a reference to the given NullableString and assigns it to the Date field.
func (o *ProvisioningSchemeImageVersionHistoryResponseModel) SetDate(v string) {
	o.Date.Set(&v)
}
// SetDateNil sets the value for Date to be an explicit nil
func (o *ProvisioningSchemeImageVersionHistoryResponseModel) SetDateNil() {
	o.Date.Set(nil)
}

// UnsetDate ensures that no value is present for Date, not even an explicit nil
func (o *ProvisioningSchemeImageVersionHistoryResponseModel) UnsetDate() {
	o.Date.Unset()
}

func (o ProvisioningSchemeImageVersionHistoryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningSchemeImageVersionHistoryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ImageVersion"] = o.ImageVersion
	if o.Date.IsSet() {
		toSerialize["Date"] = o.Date.Get()
	}
	return toSerialize, nil
}

type NullableProvisioningSchemeImageVersionHistoryResponseModel struct {
	value *ProvisioningSchemeImageVersionHistoryResponseModel
	isSet bool
}

func (v NullableProvisioningSchemeImageVersionHistoryResponseModel) Get() *ProvisioningSchemeImageVersionHistoryResponseModel {
	return v.value
}

func (v *NullableProvisioningSchemeImageVersionHistoryResponseModel) Set(val *ProvisioningSchemeImageVersionHistoryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningSchemeImageVersionHistoryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningSchemeImageVersionHistoryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningSchemeImageVersionHistoryResponseModel(val *ProvisioningSchemeImageVersionHistoryResponseModel) *NullableProvisioningSchemeImageVersionHistoryResponseModel {
	return &NullableProvisioningSchemeImageVersionHistoryResponseModel{value: val, isSet: true}
}

func (v NullableProvisioningSchemeImageVersionHistoryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningSchemeImageVersionHistoryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


