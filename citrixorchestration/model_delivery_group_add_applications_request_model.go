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

// checks if the DeliveryGroupAddApplicationsRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryGroupAddApplicationsRequestModel{}

// DeliveryGroupAddApplicationsRequestModel Request object for adding applications to a delivery group.
type DeliveryGroupAddApplicationsRequestModel struct {
	// List of existing applications to be associated with the delivery group.
	ExistingApplications []PriorityRefRequestModel `json:"ExistingApplications,omitempty"`
	// List of applications which should be created and associated with the delivery group.
	NewApplications []CreateApplicationRequestModel `json:"NewApplications,omitempty"`
	// List of existing application groups to be associated with the delivery group.
	ExistingApplicationGroups []PriorityRefRequestModel `json:"ExistingApplicationGroups,omitempty"`
}

// NewDeliveryGroupAddApplicationsRequestModel instantiates a new DeliveryGroupAddApplicationsRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryGroupAddApplicationsRequestModel() *DeliveryGroupAddApplicationsRequestModel {
	this := DeliveryGroupAddApplicationsRequestModel{}
	return &this
}

// NewDeliveryGroupAddApplicationsRequestModelWithDefaults instantiates a new DeliveryGroupAddApplicationsRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryGroupAddApplicationsRequestModelWithDefaults() *DeliveryGroupAddApplicationsRequestModel {
	this := DeliveryGroupAddApplicationsRequestModel{}
	return &this
}

// GetExistingApplications returns the ExistingApplications field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeliveryGroupAddApplicationsRequestModel) GetExistingApplications() []PriorityRefRequestModel {
	if o == nil {
		var ret []PriorityRefRequestModel
		return ret
	}
	return o.ExistingApplications
}

// GetExistingApplicationsOk returns a tuple with the ExistingApplications field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryGroupAddApplicationsRequestModel) GetExistingApplicationsOk() ([]PriorityRefRequestModel, bool) {
	if o == nil || IsNil(o.ExistingApplications) {
		return nil, false
	}
	return o.ExistingApplications, true
}

// HasExistingApplications returns a boolean if a field has been set.
func (o *DeliveryGroupAddApplicationsRequestModel) HasExistingApplications() bool {
	if o != nil && IsNil(o.ExistingApplications) {
		return true
	}

	return false
}

// SetExistingApplications gets a reference to the given []PriorityRefRequestModel and assigns it to the ExistingApplications field.
func (o *DeliveryGroupAddApplicationsRequestModel) SetExistingApplications(v []PriorityRefRequestModel) {
	o.ExistingApplications = v
}

// GetNewApplications returns the NewApplications field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeliveryGroupAddApplicationsRequestModel) GetNewApplications() []CreateApplicationRequestModel {
	if o == nil {
		var ret []CreateApplicationRequestModel
		return ret
	}
	return o.NewApplications
}

// GetNewApplicationsOk returns a tuple with the NewApplications field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryGroupAddApplicationsRequestModel) GetNewApplicationsOk() ([]CreateApplicationRequestModel, bool) {
	if o == nil || IsNil(o.NewApplications) {
		return nil, false
	}
	return o.NewApplications, true
}

// HasNewApplications returns a boolean if a field has been set.
func (o *DeliveryGroupAddApplicationsRequestModel) HasNewApplications() bool {
	if o != nil && IsNil(o.NewApplications) {
		return true
	}

	return false
}

// SetNewApplications gets a reference to the given []CreateApplicationRequestModel and assigns it to the NewApplications field.
func (o *DeliveryGroupAddApplicationsRequestModel) SetNewApplications(v []CreateApplicationRequestModel) {
	o.NewApplications = v
}

// GetExistingApplicationGroups returns the ExistingApplicationGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeliveryGroupAddApplicationsRequestModel) GetExistingApplicationGroups() []PriorityRefRequestModel {
	if o == nil {
		var ret []PriorityRefRequestModel
		return ret
	}
	return o.ExistingApplicationGroups
}

// GetExistingApplicationGroupsOk returns a tuple with the ExistingApplicationGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryGroupAddApplicationsRequestModel) GetExistingApplicationGroupsOk() ([]PriorityRefRequestModel, bool) {
	if o == nil || IsNil(o.ExistingApplicationGroups) {
		return nil, false
	}
	return o.ExistingApplicationGroups, true
}

// HasExistingApplicationGroups returns a boolean if a field has been set.
func (o *DeliveryGroupAddApplicationsRequestModel) HasExistingApplicationGroups() bool {
	if o != nil && IsNil(o.ExistingApplicationGroups) {
		return true
	}

	return false
}

// SetExistingApplicationGroups gets a reference to the given []PriorityRefRequestModel and assigns it to the ExistingApplicationGroups field.
func (o *DeliveryGroupAddApplicationsRequestModel) SetExistingApplicationGroups(v []PriorityRefRequestModel) {
	o.ExistingApplicationGroups = v
}

func (o DeliveryGroupAddApplicationsRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryGroupAddApplicationsRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExistingApplications != nil {
		toSerialize["ExistingApplications"] = o.ExistingApplications
	}
	if o.NewApplications != nil {
		toSerialize["NewApplications"] = o.NewApplications
	}
	if o.ExistingApplicationGroups != nil {
		toSerialize["ExistingApplicationGroups"] = o.ExistingApplicationGroups
	}
	return toSerialize, nil
}

type NullableDeliveryGroupAddApplicationsRequestModel struct {
	value *DeliveryGroupAddApplicationsRequestModel
	isSet bool
}

func (v NullableDeliveryGroupAddApplicationsRequestModel) Get() *DeliveryGroupAddApplicationsRequestModel {
	return v.value
}

func (v *NullableDeliveryGroupAddApplicationsRequestModel) Set(val *DeliveryGroupAddApplicationsRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryGroupAddApplicationsRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryGroupAddApplicationsRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryGroupAddApplicationsRequestModel(val *DeliveryGroupAddApplicationsRequestModel) *NullableDeliveryGroupAddApplicationsRequestModel {
	return &NullableDeliveryGroupAddApplicationsRequestModel{value: val, isSet: true}
}

func (v NullableDeliveryGroupAddApplicationsRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryGroupAddApplicationsRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
