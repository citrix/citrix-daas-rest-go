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

// checks if the AwsEdcAccountByolRegistrationTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsEdcAccountByolRegistrationTask{}

// AwsEdcAccountByolRegistrationTask Task status for a Aws BYOL Registration task
type AwsEdcAccountByolRegistrationTask struct {
	AccountTask
	// 
	DedicatedTenancyManagementCidrRange NullableString `json:"dedicatedTenancyManagementCidrRange,omitempty"`
	// 
	DedicatedTenancySupportEnabled NullableBool `json:"dedicatedTenancySupportEnabled,omitempty"`
	ModificationState NullableAwsEdcDedicatedTenancyState `json:"modificationState,omitempty"`
	// 
	ErrorCode NullableString `json:"errorCode,omitempty"`
	// 
	ErrorMessage NullableString `json:"errorMessage,omitempty"`
}

// NewAwsEdcAccountByolRegistrationTask instantiates a new AwsEdcAccountByolRegistrationTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsEdcAccountByolRegistrationTask(taskType TaskType) *AwsEdcAccountByolRegistrationTask {
	this := AwsEdcAccountByolRegistrationTask{}
	this.TaskType = taskType
	return &this
}

// NewAwsEdcAccountByolRegistrationTaskWithDefaults instantiates a new AwsEdcAccountByolRegistrationTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsEdcAccountByolRegistrationTaskWithDefaults() *AwsEdcAccountByolRegistrationTask {
	this := AwsEdcAccountByolRegistrationTask{}
	return &this
}

// GetDedicatedTenancyManagementCidrRange returns the DedicatedTenancyManagementCidrRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountByolRegistrationTask) GetDedicatedTenancyManagementCidrRange() string {
	if o == nil || IsNil(o.DedicatedTenancyManagementCidrRange.Get()) {
		var ret string
		return ret
	}
	return *o.DedicatedTenancyManagementCidrRange.Get()
}

// GetDedicatedTenancyManagementCidrRangeOk returns a tuple with the DedicatedTenancyManagementCidrRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountByolRegistrationTask) GetDedicatedTenancyManagementCidrRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DedicatedTenancyManagementCidrRange.Get(), o.DedicatedTenancyManagementCidrRange.IsSet()
}

// HasDedicatedTenancyManagementCidrRange returns a boolean if a field has been set.
func (o *AwsEdcAccountByolRegistrationTask) HasDedicatedTenancyManagementCidrRange() bool {
	if o != nil && o.DedicatedTenancyManagementCidrRange.IsSet() {
		return true
	}

	return false
}

// SetDedicatedTenancyManagementCidrRange gets a reference to the given NullableString and assigns it to the DedicatedTenancyManagementCidrRange field.
func (o *AwsEdcAccountByolRegistrationTask) SetDedicatedTenancyManagementCidrRange(v string) {
	o.DedicatedTenancyManagementCidrRange.Set(&v)
}
// SetDedicatedTenancyManagementCidrRangeNil sets the value for DedicatedTenancyManagementCidrRange to be an explicit nil
func (o *AwsEdcAccountByolRegistrationTask) SetDedicatedTenancyManagementCidrRangeNil() {
	o.DedicatedTenancyManagementCidrRange.Set(nil)
}

// UnsetDedicatedTenancyManagementCidrRange ensures that no value is present for DedicatedTenancyManagementCidrRange, not even an explicit nil
func (o *AwsEdcAccountByolRegistrationTask) UnsetDedicatedTenancyManagementCidrRange() {
	o.DedicatedTenancyManagementCidrRange.Unset()
}

// GetDedicatedTenancySupportEnabled returns the DedicatedTenancySupportEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountByolRegistrationTask) GetDedicatedTenancySupportEnabled() bool {
	if o == nil || IsNil(o.DedicatedTenancySupportEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.DedicatedTenancySupportEnabled.Get()
}

// GetDedicatedTenancySupportEnabledOk returns a tuple with the DedicatedTenancySupportEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountByolRegistrationTask) GetDedicatedTenancySupportEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DedicatedTenancySupportEnabled.Get(), o.DedicatedTenancySupportEnabled.IsSet()
}

// HasDedicatedTenancySupportEnabled returns a boolean if a field has been set.
func (o *AwsEdcAccountByolRegistrationTask) HasDedicatedTenancySupportEnabled() bool {
	if o != nil && o.DedicatedTenancySupportEnabled.IsSet() {
		return true
	}

	return false
}

// SetDedicatedTenancySupportEnabled gets a reference to the given NullableBool and assigns it to the DedicatedTenancySupportEnabled field.
func (o *AwsEdcAccountByolRegistrationTask) SetDedicatedTenancySupportEnabled(v bool) {
	o.DedicatedTenancySupportEnabled.Set(&v)
}
// SetDedicatedTenancySupportEnabledNil sets the value for DedicatedTenancySupportEnabled to be an explicit nil
func (o *AwsEdcAccountByolRegistrationTask) SetDedicatedTenancySupportEnabledNil() {
	o.DedicatedTenancySupportEnabled.Set(nil)
}

// UnsetDedicatedTenancySupportEnabled ensures that no value is present for DedicatedTenancySupportEnabled, not even an explicit nil
func (o *AwsEdcAccountByolRegistrationTask) UnsetDedicatedTenancySupportEnabled() {
	o.DedicatedTenancySupportEnabled.Unset()
}

// GetModificationState returns the ModificationState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountByolRegistrationTask) GetModificationState() AwsEdcDedicatedTenancyState {
	if o == nil || IsNil(o.ModificationState.Get()) {
		var ret AwsEdcDedicatedTenancyState
		return ret
	}
	return *o.ModificationState.Get()
}

// GetModificationStateOk returns a tuple with the ModificationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountByolRegistrationTask) GetModificationStateOk() (*AwsEdcDedicatedTenancyState, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModificationState.Get(), o.ModificationState.IsSet()
}

// HasModificationState returns a boolean if a field has been set.
func (o *AwsEdcAccountByolRegistrationTask) HasModificationState() bool {
	if o != nil && o.ModificationState.IsSet() {
		return true
	}

	return false
}

// SetModificationState gets a reference to the given NullableAwsEdcDedicatedTenancyState and assigns it to the ModificationState field.
func (o *AwsEdcAccountByolRegistrationTask) SetModificationState(v AwsEdcDedicatedTenancyState) {
	o.ModificationState.Set(&v)
}
// SetModificationStateNil sets the value for ModificationState to be an explicit nil
func (o *AwsEdcAccountByolRegistrationTask) SetModificationStateNil() {
	o.ModificationState.Set(nil)
}

// UnsetModificationState ensures that no value is present for ModificationState, not even an explicit nil
func (o *AwsEdcAccountByolRegistrationTask) UnsetModificationState() {
	o.ModificationState.Unset()
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountByolRegistrationTask) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorCode.Get()
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountByolRegistrationTask) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorCode.Get(), o.ErrorCode.IsSet()
}

// HasErrorCode returns a boolean if a field has been set.
func (o *AwsEdcAccountByolRegistrationTask) HasErrorCode() bool {
	if o != nil && o.ErrorCode.IsSet() {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableString and assigns it to the ErrorCode field.
func (o *AwsEdcAccountByolRegistrationTask) SetErrorCode(v string) {
	o.ErrorCode.Set(&v)
}
// SetErrorCodeNil sets the value for ErrorCode to be an explicit nil
func (o *AwsEdcAccountByolRegistrationTask) SetErrorCodeNil() {
	o.ErrorCode.Set(nil)
}

// UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
func (o *AwsEdcAccountByolRegistrationTask) UnsetErrorCode() {
	o.ErrorCode.Unset()
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountByolRegistrationTask) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountByolRegistrationTask) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *AwsEdcAccountByolRegistrationTask) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *AwsEdcAccountByolRegistrationTask) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *AwsEdcAccountByolRegistrationTask) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *AwsEdcAccountByolRegistrationTask) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

func (o AwsEdcAccountByolRegistrationTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsEdcAccountByolRegistrationTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAccountTask, errAccountTask := json.Marshal(o.AccountTask)
	if errAccountTask != nil {
		return map[string]interface{}{}, errAccountTask
	}
	errAccountTask = json.Unmarshal([]byte(serializedAccountTask), &toSerialize)
	if errAccountTask != nil {
		return map[string]interface{}{}, errAccountTask
	}
	if o.DedicatedTenancyManagementCidrRange.IsSet() {
		toSerialize["dedicatedTenancyManagementCidrRange"] = o.DedicatedTenancyManagementCidrRange.Get()
	}
	if o.DedicatedTenancySupportEnabled.IsSet() {
		toSerialize["dedicatedTenancySupportEnabled"] = o.DedicatedTenancySupportEnabled.Get()
	}
	if o.ModificationState.IsSet() {
		toSerialize["modificationState"] = o.ModificationState.Get()
	}
	if o.ErrorCode.IsSet() {
		toSerialize["errorCode"] = o.ErrorCode.Get()
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["errorMessage"] = o.ErrorMessage.Get()
	}
	return toSerialize, nil
}

type NullableAwsEdcAccountByolRegistrationTask struct {
	value *AwsEdcAccountByolRegistrationTask
	isSet bool
}

func (v NullableAwsEdcAccountByolRegistrationTask) Get() *AwsEdcAccountByolRegistrationTask {
	return v.value
}

func (v *NullableAwsEdcAccountByolRegistrationTask) Set(val *AwsEdcAccountByolRegistrationTask) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsEdcAccountByolRegistrationTask) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsEdcAccountByolRegistrationTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsEdcAccountByolRegistrationTask(val *AwsEdcAccountByolRegistrationTask) *NullableAwsEdcAccountByolRegistrationTask {
	return &NullableAwsEdcAccountByolRegistrationTask{value: val, isSet: true}
}

func (v NullableAwsEdcAccountByolRegistrationTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsEdcAccountByolRegistrationTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

