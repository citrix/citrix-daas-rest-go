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

// checks if the DeploymentWarning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentWarning{}

// DeploymentWarning A warning or error that is associated with the deployment
type DeploymentWarning struct {
	// Type of deployment warning
	WarningType *DeploymentWarningType `json:"warningType,omitempty"`
	// Id of the warning
	Id NullableString `json:"id,omitempty"`
	// The warning message
	Warning NullableString `json:"warning,omitempty"`
	// ID of the task the warning was created by
	TaskId NullableString `json:"taskId,omitempty"`
	// Indicates if the warning should be treated as an error
	IsError *bool `json:"isError,omitempty"`
	// Indicates if the warning message can be cleared out by the user
	IsDismissible *bool `json:"isDismissible,omitempty"`
	// ID of the transaction the warning was created by
	TransactionId NullableString `json:"transactionId,omitempty"`
}

// NewDeploymentWarning instantiates a new DeploymentWarning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentWarning() *DeploymentWarning {
	this := DeploymentWarning{}
	return &this
}

// NewDeploymentWarningWithDefaults instantiates a new DeploymentWarning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentWarningWithDefaults() *DeploymentWarning {
	this := DeploymentWarning{}
	return &this
}

// GetWarningType returns the WarningType field value if set, zero value otherwise.
func (o *DeploymentWarning) GetWarningType() DeploymentWarningType {
	if o == nil || IsNil(o.WarningType) {
		var ret DeploymentWarningType
		return ret
	}
	return *o.WarningType
}

// GetWarningTypeOk returns a tuple with the WarningType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentWarning) GetWarningTypeOk() (*DeploymentWarningType, bool) {
	if o == nil || IsNil(o.WarningType) {
		return nil, false
	}
	return o.WarningType, true
}

// HasWarningType returns a boolean if a field has been set.
func (o *DeploymentWarning) HasWarningType() bool {
	if o != nil && !IsNil(o.WarningType) {
		return true
	}

	return false
}

// SetWarningType gets a reference to the given DeploymentWarningType and assigns it to the WarningType field.
func (o *DeploymentWarning) SetWarningType(v DeploymentWarningType) {
	o.WarningType = &v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentWarning) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentWarning) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *DeploymentWarning) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *DeploymentWarning) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *DeploymentWarning) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *DeploymentWarning) UnsetId() {
	o.Id.Unset()
}

// GetWarning returns the Warning field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentWarning) GetWarning() string {
	if o == nil || IsNil(o.Warning.Get()) {
		var ret string
		return ret
	}
	return *o.Warning.Get()
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentWarning) GetWarningOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Warning.Get(), o.Warning.IsSet()
}

// HasWarning returns a boolean if a field has been set.
func (o *DeploymentWarning) HasWarning() bool {
	if o != nil && o.Warning.IsSet() {
		return true
	}

	return false
}

// SetWarning gets a reference to the given NullableString and assigns it to the Warning field.
func (o *DeploymentWarning) SetWarning(v string) {
	o.Warning.Set(&v)
}

// SetWarningNil sets the value for Warning to be an explicit nil
func (o *DeploymentWarning) SetWarningNil() {
	o.Warning.Set(nil)
}

// UnsetWarning ensures that no value is present for Warning, not even an explicit nil
func (o *DeploymentWarning) UnsetWarning() {
	o.Warning.Unset()
}

// GetTaskId returns the TaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentWarning) GetTaskId() string {
	if o == nil || IsNil(o.TaskId.Get()) {
		var ret string
		return ret
	}
	return *o.TaskId.Get()
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentWarning) GetTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskId.Get(), o.TaskId.IsSet()
}

// HasTaskId returns a boolean if a field has been set.
func (o *DeploymentWarning) HasTaskId() bool {
	if o != nil && o.TaskId.IsSet() {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given NullableString and assigns it to the TaskId field.
func (o *DeploymentWarning) SetTaskId(v string) {
	o.TaskId.Set(&v)
}

// SetTaskIdNil sets the value for TaskId to be an explicit nil
func (o *DeploymentWarning) SetTaskIdNil() {
	o.TaskId.Set(nil)
}

// UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
func (o *DeploymentWarning) UnsetTaskId() {
	o.TaskId.Unset()
}

// GetIsError returns the IsError field value if set, zero value otherwise.
func (o *DeploymentWarning) GetIsError() bool {
	if o == nil || IsNil(o.IsError) {
		var ret bool
		return ret
	}
	return *o.IsError
}

// GetIsErrorOk returns a tuple with the IsError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentWarning) GetIsErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.IsError) {
		return nil, false
	}
	return o.IsError, true
}

// HasIsError returns a boolean if a field has been set.
func (o *DeploymentWarning) HasIsError() bool {
	if o != nil && !IsNil(o.IsError) {
		return true
	}

	return false
}

// SetIsError gets a reference to the given bool and assigns it to the IsError field.
func (o *DeploymentWarning) SetIsError(v bool) {
	o.IsError = &v
}

// GetIsDismissible returns the IsDismissible field value if set, zero value otherwise.
func (o *DeploymentWarning) GetIsDismissible() bool {
	if o == nil || IsNil(o.IsDismissible) {
		var ret bool
		return ret
	}
	return *o.IsDismissible
}

// GetIsDismissibleOk returns a tuple with the IsDismissible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentWarning) GetIsDismissibleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDismissible) {
		return nil, false
	}
	return o.IsDismissible, true
}

// HasIsDismissible returns a boolean if a field has been set.
func (o *DeploymentWarning) HasIsDismissible() bool {
	if o != nil && !IsNil(o.IsDismissible) {
		return true
	}

	return false
}

// SetIsDismissible gets a reference to the given bool and assigns it to the IsDismissible field.
func (o *DeploymentWarning) SetIsDismissible(v bool) {
	o.IsDismissible = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentWarning) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId.Get()) {
		var ret string
		return ret
	}
	return *o.TransactionId.Get()
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentWarning) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionId.Get(), o.TransactionId.IsSet()
}

// HasTransactionId returns a boolean if a field has been set.
func (o *DeploymentWarning) HasTransactionId() bool {
	if o != nil && o.TransactionId.IsSet() {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given NullableString and assigns it to the TransactionId field.
func (o *DeploymentWarning) SetTransactionId(v string) {
	o.TransactionId.Set(&v)
}

// SetTransactionIdNil sets the value for TransactionId to be an explicit nil
func (o *DeploymentWarning) SetTransactionIdNil() {
	o.TransactionId.Set(nil)
}

// UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
func (o *DeploymentWarning) UnsetTransactionId() {
	o.TransactionId.Unset()
}

func (o DeploymentWarning) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentWarning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WarningType) {
		toSerialize["warningType"] = o.WarningType
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Warning.IsSet() {
		toSerialize["warning"] = o.Warning.Get()
	}
	if o.TaskId.IsSet() {
		toSerialize["taskId"] = o.TaskId.Get()
	}
	if !IsNil(o.IsError) {
		toSerialize["isError"] = o.IsError
	}
	if !IsNil(o.IsDismissible) {
		toSerialize["isDismissible"] = o.IsDismissible
	}
	if o.TransactionId.IsSet() {
		toSerialize["transactionId"] = o.TransactionId.Get()
	}
	return toSerialize, nil
}

type NullableDeploymentWarning struct {
	value *DeploymentWarning
	isSet bool
}

func (v NullableDeploymentWarning) Get() *DeploymentWarning {
	return v.value
}

func (v *NullableDeploymentWarning) Set(val *DeploymentWarning) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentWarning) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentWarning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentWarning(val *DeploymentWarning) *NullableDeploymentWarning {
	return &NullableDeploymentWarning{value: val, isSet: true}
}

func (v NullableDeploymentWarning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentWarning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
