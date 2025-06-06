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

// checks if the UpdateAutoscaleGroupPluginRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAutoscaleGroupPluginRequestModel{}

// UpdateAutoscaleGroupPluginRequestModel Update autoscale group plugin request model
type UpdateAutoscaleGroupPluginRequestModel struct {
	// The name of the group plugin.
	Name NullableString `json:"Name,omitempty"`
	// The priority of this group plugin with respect to other group plugins associated with the same desktop group. Priority 1 is the highest priority.
	Priority NullableInt32 `json:"Priority,omitempty"`
	// Specifies whether the group plugin is enabled.
	Enabled NullableBool `json:"Enabled,omitempty"`
	// Optional description of the group plugin.
	Description NullableString `json:"Description,omitempty"`
	// If True, indicates that any schedule override produced by the plugin is automatically accepted as the schedule for the target date. Otherwise the schedule is only accepted if the target date is on or before the date specified in ApprovedUtil. If AutoApprove is False and the ApprovedUtil date is not set, or is in the past, then any schedule override produced is discarded without being used.
	AutoApprove NullableBool `json:"AutoApprove,omitempty"`
	// Date range for the autoscale holiday plugin template.
	Dates []string `json:"Dates,omitempty"`
	// If AutoApprove is False, specifies the date until which any schedule override produced by the plugin is automatically accepted as the schedule for the target date.
	ApprovedUntil NullableString `json:"ApprovedUntil,omitempty"`
}

// NewUpdateAutoscaleGroupPluginRequestModel instantiates a new UpdateAutoscaleGroupPluginRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAutoscaleGroupPluginRequestModel() *UpdateAutoscaleGroupPluginRequestModel {
	this := UpdateAutoscaleGroupPluginRequestModel{}
	return &this
}

// NewUpdateAutoscaleGroupPluginRequestModelWithDefaults instantiates a new UpdateAutoscaleGroupPluginRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAutoscaleGroupPluginRequestModelWithDefaults() *UpdateAutoscaleGroupPluginRequestModel {
	this := UpdateAutoscaleGroupPluginRequestModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAutoscaleGroupPluginRequestModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAutoscaleGroupPluginRequestModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateAutoscaleGroupPluginRequestModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateAutoscaleGroupPluginRequestModel) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateAutoscaleGroupPluginRequestModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateAutoscaleGroupPluginRequestModel) UnsetName() {
	o.Name.Unset()
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAutoscaleGroupPluginRequestModel) GetPriority() int32 {
	if o == nil || IsNil(o.Priority.Get()) {
		var ret int32
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAutoscaleGroupPluginRequestModel) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// HasPriority returns a boolean if a field has been set.
func (o *UpdateAutoscaleGroupPluginRequestModel) HasPriority() bool {
	if o != nil && o.Priority.IsSet() {
		return true
	}

	return false
}

// SetPriority gets a reference to the given NullableInt32 and assigns it to the Priority field.
func (o *UpdateAutoscaleGroupPluginRequestModel) SetPriority(v int32) {
	o.Priority.Set(&v)
}

// SetPriorityNil sets the value for Priority to be an explicit nil
func (o *UpdateAutoscaleGroupPluginRequestModel) SetPriorityNil() {
	o.Priority.Set(nil)
}

// UnsetPriority ensures that no value is present for Priority, not even an explicit nil
func (o *UpdateAutoscaleGroupPluginRequestModel) UnsetPriority() {
	o.Priority.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAutoscaleGroupPluginRequestModel) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAutoscaleGroupPluginRequestModel) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateAutoscaleGroupPluginRequestModel) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *UpdateAutoscaleGroupPluginRequestModel) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}

// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *UpdateAutoscaleGroupPluginRequestModel) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *UpdateAutoscaleGroupPluginRequestModel) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAutoscaleGroupPluginRequestModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAutoscaleGroupPluginRequestModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateAutoscaleGroupPluginRequestModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *UpdateAutoscaleGroupPluginRequestModel) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UpdateAutoscaleGroupPluginRequestModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UpdateAutoscaleGroupPluginRequestModel) UnsetDescription() {
	o.Description.Unset()
}

// GetAutoApprove returns the AutoApprove field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAutoscaleGroupPluginRequestModel) GetAutoApprove() bool {
	if o == nil || IsNil(o.AutoApprove.Get()) {
		var ret bool
		return ret
	}
	return *o.AutoApprove.Get()
}

// GetAutoApproveOk returns a tuple with the AutoApprove field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAutoscaleGroupPluginRequestModel) GetAutoApproveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoApprove.Get(), o.AutoApprove.IsSet()
}

// HasAutoApprove returns a boolean if a field has been set.
func (o *UpdateAutoscaleGroupPluginRequestModel) HasAutoApprove() bool {
	if o != nil && o.AutoApprove.IsSet() {
		return true
	}

	return false
}

// SetAutoApprove gets a reference to the given NullableBool and assigns it to the AutoApprove field.
func (o *UpdateAutoscaleGroupPluginRequestModel) SetAutoApprove(v bool) {
	o.AutoApprove.Set(&v)
}

// SetAutoApproveNil sets the value for AutoApprove to be an explicit nil
func (o *UpdateAutoscaleGroupPluginRequestModel) SetAutoApproveNil() {
	o.AutoApprove.Set(nil)
}

// UnsetAutoApprove ensures that no value is present for AutoApprove, not even an explicit nil
func (o *UpdateAutoscaleGroupPluginRequestModel) UnsetAutoApprove() {
	o.AutoApprove.Unset()
}

// GetDates returns the Dates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAutoscaleGroupPluginRequestModel) GetDates() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Dates
}

// GetDatesOk returns a tuple with the Dates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAutoscaleGroupPluginRequestModel) GetDatesOk() ([]string, bool) {
	if o == nil || IsNil(o.Dates) {
		return nil, false
	}
	return o.Dates, true
}

// HasDates returns a boolean if a field has been set.
func (o *UpdateAutoscaleGroupPluginRequestModel) HasDates() bool {
	if o != nil && IsNil(o.Dates) {
		return true
	}

	return false
}

// SetDates gets a reference to the given []string and assigns it to the Dates field.
func (o *UpdateAutoscaleGroupPluginRequestModel) SetDates(v []string) {
	o.Dates = v
}

// GetApprovedUntil returns the ApprovedUntil field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAutoscaleGroupPluginRequestModel) GetApprovedUntil() string {
	if o == nil || IsNil(o.ApprovedUntil.Get()) {
		var ret string
		return ret
	}
	return *o.ApprovedUntil.Get()
}

// GetApprovedUntilOk returns a tuple with the ApprovedUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAutoscaleGroupPluginRequestModel) GetApprovedUntilOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApprovedUntil.Get(), o.ApprovedUntil.IsSet()
}

// HasApprovedUntil returns a boolean if a field has been set.
func (o *UpdateAutoscaleGroupPluginRequestModel) HasApprovedUntil() bool {
	if o != nil && o.ApprovedUntil.IsSet() {
		return true
	}

	return false
}

// SetApprovedUntil gets a reference to the given NullableString and assigns it to the ApprovedUntil field.
func (o *UpdateAutoscaleGroupPluginRequestModel) SetApprovedUntil(v string) {
	o.ApprovedUntil.Set(&v)
}

// SetApprovedUntilNil sets the value for ApprovedUntil to be an explicit nil
func (o *UpdateAutoscaleGroupPluginRequestModel) SetApprovedUntilNil() {
	o.ApprovedUntil.Set(nil)
}

// UnsetApprovedUntil ensures that no value is present for ApprovedUntil, not even an explicit nil
func (o *UpdateAutoscaleGroupPluginRequestModel) UnsetApprovedUntil() {
	o.ApprovedUntil.Unset()
}

func (o UpdateAutoscaleGroupPluginRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAutoscaleGroupPluginRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.Priority.IsSet() {
		toSerialize["Priority"] = o.Priority.Get()
	}
	if o.Enabled.IsSet() {
		toSerialize["Enabled"] = o.Enabled.Get()
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if o.AutoApprove.IsSet() {
		toSerialize["AutoApprove"] = o.AutoApprove.Get()
	}
	if o.Dates != nil {
		toSerialize["Dates"] = o.Dates
	}
	if o.ApprovedUntil.IsSet() {
		toSerialize["ApprovedUntil"] = o.ApprovedUntil.Get()
	}
	return toSerialize, nil
}

type NullableUpdateAutoscaleGroupPluginRequestModel struct {
	value *UpdateAutoscaleGroupPluginRequestModel
	isSet bool
}

func (v NullableUpdateAutoscaleGroupPluginRequestModel) Get() *UpdateAutoscaleGroupPluginRequestModel {
	return v.value
}

func (v *NullableUpdateAutoscaleGroupPluginRequestModel) Set(val *UpdateAutoscaleGroupPluginRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAutoscaleGroupPluginRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAutoscaleGroupPluginRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAutoscaleGroupPluginRequestModel(val *UpdateAutoscaleGroupPluginRequestModel) *NullableUpdateAutoscaleGroupPluginRequestModel {
	return &NullableUpdateAutoscaleGroupPluginRequestModel{value: val, isSet: true}
}

func (v NullableUpdateAutoscaleGroupPluginRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAutoscaleGroupPluginRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
