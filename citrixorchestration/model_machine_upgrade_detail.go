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

// checks if the MachineUpgradeDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MachineUpgradeDetail{}

// MachineUpgradeDetail Status detail of last VDA upgrade schedule for a machine.
type MachineUpgradeDetail struct {
	// Uuid of this upgrade status object.
	Uuid           NullableString                   `json:"Uuid,omitempty"`
	ScheduleStatus *VdaUpgradeMachineScheduleStatus `json:"ScheduleStatus,omitempty"`
	// Status message resulted from the action of this upgrade status object.
	StatusMessage NullableString `json:"StatusMessage,omitempty"`
	// UTC time when this upgrade status object changed status for the last time.
	LastStateChangeTimeUtc NullableString `json:"LastStateChangeTimeUtc,omitempty"`
	// UTC time when this VDA upgrade is scheduled to start.
	ScheduledTimeUtc NullableString `json:"ScheduledTimeUtc,omitempty"`
	// Timeout duration in hours, of the current VDA upgrade schdeule.
	DurationInHours NullableInt32 `json:"DurationInHours,omitempty"`
	// Target package version of the current VDA upgrade schdeule.
	TargetPackageVersion NullableString `json:"TargetPackageVersion,omitempty"`
}

// NewMachineUpgradeDetail instantiates a new MachineUpgradeDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMachineUpgradeDetail() *MachineUpgradeDetail {
	this := MachineUpgradeDetail{}
	return &this
}

// NewMachineUpgradeDetailWithDefaults instantiates a new MachineUpgradeDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineUpgradeDetailWithDefaults() *MachineUpgradeDetail {
	this := MachineUpgradeDetail{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineUpgradeDetail) GetUuid() string {
	if o == nil || IsNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineUpgradeDetail) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *MachineUpgradeDetail) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *MachineUpgradeDetail) SetUuid(v string) {
	o.Uuid.Set(&v)
}

// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *MachineUpgradeDetail) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *MachineUpgradeDetail) UnsetUuid() {
	o.Uuid.Unset()
}

// GetScheduleStatus returns the ScheduleStatus field value if set, zero value otherwise.
func (o *MachineUpgradeDetail) GetScheduleStatus() VdaUpgradeMachineScheduleStatus {
	if o == nil || IsNil(o.ScheduleStatus) {
		var ret VdaUpgradeMachineScheduleStatus
		return ret
	}
	return *o.ScheduleStatus
}

// GetScheduleStatusOk returns a tuple with the ScheduleStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineUpgradeDetail) GetScheduleStatusOk() (*VdaUpgradeMachineScheduleStatus, bool) {
	if o == nil || IsNil(o.ScheduleStatus) {
		return nil, false
	}
	return o.ScheduleStatus, true
}

// HasScheduleStatus returns a boolean if a field has been set.
func (o *MachineUpgradeDetail) HasScheduleStatus() bool {
	if o != nil && !IsNil(o.ScheduleStatus) {
		return true
	}

	return false
}

// SetScheduleStatus gets a reference to the given VdaUpgradeMachineScheduleStatus and assigns it to the ScheduleStatus field.
func (o *MachineUpgradeDetail) SetScheduleStatus(v VdaUpgradeMachineScheduleStatus) {
	o.ScheduleStatus = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineUpgradeDetail) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage.Get()) {
		var ret string
		return ret
	}
	return *o.StatusMessage.Get()
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineUpgradeDetail) GetStatusMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatusMessage.Get(), o.StatusMessage.IsSet()
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *MachineUpgradeDetail) HasStatusMessage() bool {
	if o != nil && o.StatusMessage.IsSet() {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given NullableString and assigns it to the StatusMessage field.
func (o *MachineUpgradeDetail) SetStatusMessage(v string) {
	o.StatusMessage.Set(&v)
}

// SetStatusMessageNil sets the value for StatusMessage to be an explicit nil
func (o *MachineUpgradeDetail) SetStatusMessageNil() {
	o.StatusMessage.Set(nil)
}

// UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
func (o *MachineUpgradeDetail) UnsetStatusMessage() {
	o.StatusMessage.Unset()
}

// GetLastStateChangeTimeUtc returns the LastStateChangeTimeUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineUpgradeDetail) GetLastStateChangeTimeUtc() string {
	if o == nil || IsNil(o.LastStateChangeTimeUtc.Get()) {
		var ret string
		return ret
	}
	return *o.LastStateChangeTimeUtc.Get()
}

// GetLastStateChangeTimeUtcOk returns a tuple with the LastStateChangeTimeUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineUpgradeDetail) GetLastStateChangeTimeUtcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastStateChangeTimeUtc.Get(), o.LastStateChangeTimeUtc.IsSet()
}

// HasLastStateChangeTimeUtc returns a boolean if a field has been set.
func (o *MachineUpgradeDetail) HasLastStateChangeTimeUtc() bool {
	if o != nil && o.LastStateChangeTimeUtc.IsSet() {
		return true
	}

	return false
}

// SetLastStateChangeTimeUtc gets a reference to the given NullableString and assigns it to the LastStateChangeTimeUtc field.
func (o *MachineUpgradeDetail) SetLastStateChangeTimeUtc(v string) {
	o.LastStateChangeTimeUtc.Set(&v)
}

// SetLastStateChangeTimeUtcNil sets the value for LastStateChangeTimeUtc to be an explicit nil
func (o *MachineUpgradeDetail) SetLastStateChangeTimeUtcNil() {
	o.LastStateChangeTimeUtc.Set(nil)
}

// UnsetLastStateChangeTimeUtc ensures that no value is present for LastStateChangeTimeUtc, not even an explicit nil
func (o *MachineUpgradeDetail) UnsetLastStateChangeTimeUtc() {
	o.LastStateChangeTimeUtc.Unset()
}

// GetScheduledTimeUtc returns the ScheduledTimeUtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineUpgradeDetail) GetScheduledTimeUtc() string {
	if o == nil || IsNil(o.ScheduledTimeUtc.Get()) {
		var ret string
		return ret
	}
	return *o.ScheduledTimeUtc.Get()
}

// GetScheduledTimeUtcOk returns a tuple with the ScheduledTimeUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineUpgradeDetail) GetScheduledTimeUtcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledTimeUtc.Get(), o.ScheduledTimeUtc.IsSet()
}

// HasScheduledTimeUtc returns a boolean if a field has been set.
func (o *MachineUpgradeDetail) HasScheduledTimeUtc() bool {
	if o != nil && o.ScheduledTimeUtc.IsSet() {
		return true
	}

	return false
}

// SetScheduledTimeUtc gets a reference to the given NullableString and assigns it to the ScheduledTimeUtc field.
func (o *MachineUpgradeDetail) SetScheduledTimeUtc(v string) {
	o.ScheduledTimeUtc.Set(&v)
}

// SetScheduledTimeUtcNil sets the value for ScheduledTimeUtc to be an explicit nil
func (o *MachineUpgradeDetail) SetScheduledTimeUtcNil() {
	o.ScheduledTimeUtc.Set(nil)
}

// UnsetScheduledTimeUtc ensures that no value is present for ScheduledTimeUtc, not even an explicit nil
func (o *MachineUpgradeDetail) UnsetScheduledTimeUtc() {
	o.ScheduledTimeUtc.Unset()
}

// GetDurationInHours returns the DurationInHours field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineUpgradeDetail) GetDurationInHours() int32 {
	if o == nil || IsNil(o.DurationInHours.Get()) {
		var ret int32
		return ret
	}
	return *o.DurationInHours.Get()
}

// GetDurationInHoursOk returns a tuple with the DurationInHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineUpgradeDetail) GetDurationInHoursOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DurationInHours.Get(), o.DurationInHours.IsSet()
}

// HasDurationInHours returns a boolean if a field has been set.
func (o *MachineUpgradeDetail) HasDurationInHours() bool {
	if o != nil && o.DurationInHours.IsSet() {
		return true
	}

	return false
}

// SetDurationInHours gets a reference to the given NullableInt32 and assigns it to the DurationInHours field.
func (o *MachineUpgradeDetail) SetDurationInHours(v int32) {
	o.DurationInHours.Set(&v)
}

// SetDurationInHoursNil sets the value for DurationInHours to be an explicit nil
func (o *MachineUpgradeDetail) SetDurationInHoursNil() {
	o.DurationInHours.Set(nil)
}

// UnsetDurationInHours ensures that no value is present for DurationInHours, not even an explicit nil
func (o *MachineUpgradeDetail) UnsetDurationInHours() {
	o.DurationInHours.Unset()
}

// GetTargetPackageVersion returns the TargetPackageVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineUpgradeDetail) GetTargetPackageVersion() string {
	if o == nil || IsNil(o.TargetPackageVersion.Get()) {
		var ret string
		return ret
	}
	return *o.TargetPackageVersion.Get()
}

// GetTargetPackageVersionOk returns a tuple with the TargetPackageVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineUpgradeDetail) GetTargetPackageVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetPackageVersion.Get(), o.TargetPackageVersion.IsSet()
}

// HasTargetPackageVersion returns a boolean if a field has been set.
func (o *MachineUpgradeDetail) HasTargetPackageVersion() bool {
	if o != nil && o.TargetPackageVersion.IsSet() {
		return true
	}

	return false
}

// SetTargetPackageVersion gets a reference to the given NullableString and assigns it to the TargetPackageVersion field.
func (o *MachineUpgradeDetail) SetTargetPackageVersion(v string) {
	o.TargetPackageVersion.Set(&v)
}

// SetTargetPackageVersionNil sets the value for TargetPackageVersion to be an explicit nil
func (o *MachineUpgradeDetail) SetTargetPackageVersionNil() {
	o.TargetPackageVersion.Set(nil)
}

// UnsetTargetPackageVersion ensures that no value is present for TargetPackageVersion, not even an explicit nil
func (o *MachineUpgradeDetail) UnsetTargetPackageVersion() {
	o.TargetPackageVersion.Unset()
}

func (o MachineUpgradeDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MachineUpgradeDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid.IsSet() {
		toSerialize["Uuid"] = o.Uuid.Get()
	}
	if !IsNil(o.ScheduleStatus) {
		toSerialize["ScheduleStatus"] = o.ScheduleStatus
	}
	if o.StatusMessage.IsSet() {
		toSerialize["StatusMessage"] = o.StatusMessage.Get()
	}
	if o.LastStateChangeTimeUtc.IsSet() {
		toSerialize["LastStateChangeTimeUtc"] = o.LastStateChangeTimeUtc.Get()
	}
	if o.ScheduledTimeUtc.IsSet() {
		toSerialize["ScheduledTimeUtc"] = o.ScheduledTimeUtc.Get()
	}
	if o.DurationInHours.IsSet() {
		toSerialize["DurationInHours"] = o.DurationInHours.Get()
	}
	if o.TargetPackageVersion.IsSet() {
		toSerialize["TargetPackageVersion"] = o.TargetPackageVersion.Get()
	}
	return toSerialize, nil
}

type NullableMachineUpgradeDetail struct {
	value *MachineUpgradeDetail
	isSet bool
}

func (v NullableMachineUpgradeDetail) Get() *MachineUpgradeDetail {
	return v.value
}

func (v *NullableMachineUpgradeDetail) Set(val *MachineUpgradeDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineUpgradeDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineUpgradeDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineUpgradeDetail(val *MachineUpgradeDetail) *NullableMachineUpgradeDetail {
	return &NullableMachineUpgradeDetail{value: val, isSet: true}
}

func (v NullableMachineUpgradeDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineUpgradeDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
