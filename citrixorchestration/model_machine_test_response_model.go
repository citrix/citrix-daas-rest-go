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

// checks if the MachineTestResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MachineTestResponseModel{}

// MachineTestResponseModel struct for MachineTestResponseModel
type MachineTestResponseModel struct {
	Machine *RefResponseModel       `json:"Machine,omitempty"`
	Status  *CloudHealthCheckStatus `json:"Status,omitempty"`
	// CommandResponse.
	CommandResponse NullableString `json:"CommandResponse,omitempty"`
	// ErrorMessage.
	ErrorMessage NullableString `json:"ErrorMessage,omitempty"`
	// CommandName.
	CommandName NullableString `json:"CommandName,omitempty"`
	// CategoryName.
	CategoryName NullableString `json:"CategoryName,omitempty"`
	// The owner of the test
	CreatedBy NullableString `json:"CreatedBy,omitempty"`
}

// NewMachineTestResponseModel instantiates a new MachineTestResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMachineTestResponseModel() *MachineTestResponseModel {
	this := MachineTestResponseModel{}
	return &this
}

// NewMachineTestResponseModelWithDefaults instantiates a new MachineTestResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineTestResponseModelWithDefaults() *MachineTestResponseModel {
	this := MachineTestResponseModel{}
	return &this
}

// GetMachine returns the Machine field value if set, zero value otherwise.
func (o *MachineTestResponseModel) GetMachine() RefResponseModel {
	if o == nil || IsNil(o.Machine) {
		var ret RefResponseModel
		return ret
	}
	return *o.Machine
}

// GetMachineOk returns a tuple with the Machine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineTestResponseModel) GetMachineOk() (*RefResponseModel, bool) {
	if o == nil || IsNil(o.Machine) {
		return nil, false
	}
	return o.Machine, true
}

// HasMachine returns a boolean if a field has been set.
func (o *MachineTestResponseModel) HasMachine() bool {
	if o != nil && !IsNil(o.Machine) {
		return true
	}

	return false
}

// SetMachine gets a reference to the given RefResponseModel and assigns it to the Machine field.
func (o *MachineTestResponseModel) SetMachine(v RefResponseModel) {
	o.Machine = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MachineTestResponseModel) GetStatus() CloudHealthCheckStatus {
	if o == nil || IsNil(o.Status) {
		var ret CloudHealthCheckStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineTestResponseModel) GetStatusOk() (*CloudHealthCheckStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MachineTestResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CloudHealthCheckStatus and assigns it to the Status field.
func (o *MachineTestResponseModel) SetStatus(v CloudHealthCheckStatus) {
	o.Status = &v
}

// GetCommandResponse returns the CommandResponse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineTestResponseModel) GetCommandResponse() string {
	if o == nil || IsNil(o.CommandResponse.Get()) {
		var ret string
		return ret
	}
	return *o.CommandResponse.Get()
}

// GetCommandResponseOk returns a tuple with the CommandResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineTestResponseModel) GetCommandResponseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommandResponse.Get(), o.CommandResponse.IsSet()
}

// HasCommandResponse returns a boolean if a field has been set.
func (o *MachineTestResponseModel) HasCommandResponse() bool {
	if o != nil && o.CommandResponse.IsSet() {
		return true
	}

	return false
}

// SetCommandResponse gets a reference to the given NullableString and assigns it to the CommandResponse field.
func (o *MachineTestResponseModel) SetCommandResponse(v string) {
	o.CommandResponse.Set(&v)
}

// SetCommandResponseNil sets the value for CommandResponse to be an explicit nil
func (o *MachineTestResponseModel) SetCommandResponseNil() {
	o.CommandResponse.Set(nil)
}

// UnsetCommandResponse ensures that no value is present for CommandResponse, not even an explicit nil
func (o *MachineTestResponseModel) UnsetCommandResponse() {
	o.CommandResponse.Unset()
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineTestResponseModel) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineTestResponseModel) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *MachineTestResponseModel) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *MachineTestResponseModel) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}

// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *MachineTestResponseModel) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *MachineTestResponseModel) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetCommandName returns the CommandName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineTestResponseModel) GetCommandName() string {
	if o == nil || IsNil(o.CommandName.Get()) {
		var ret string
		return ret
	}
	return *o.CommandName.Get()
}

// GetCommandNameOk returns a tuple with the CommandName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineTestResponseModel) GetCommandNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommandName.Get(), o.CommandName.IsSet()
}

// HasCommandName returns a boolean if a field has been set.
func (o *MachineTestResponseModel) HasCommandName() bool {
	if o != nil && o.CommandName.IsSet() {
		return true
	}

	return false
}

// SetCommandName gets a reference to the given NullableString and assigns it to the CommandName field.
func (o *MachineTestResponseModel) SetCommandName(v string) {
	o.CommandName.Set(&v)
}

// SetCommandNameNil sets the value for CommandName to be an explicit nil
func (o *MachineTestResponseModel) SetCommandNameNil() {
	o.CommandName.Set(nil)
}

// UnsetCommandName ensures that no value is present for CommandName, not even an explicit nil
func (o *MachineTestResponseModel) UnsetCommandName() {
	o.CommandName.Unset()
}

// GetCategoryName returns the CategoryName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineTestResponseModel) GetCategoryName() string {
	if o == nil || IsNil(o.CategoryName.Get()) {
		var ret string
		return ret
	}
	return *o.CategoryName.Get()
}

// GetCategoryNameOk returns a tuple with the CategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineTestResponseModel) GetCategoryNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CategoryName.Get(), o.CategoryName.IsSet()
}

// HasCategoryName returns a boolean if a field has been set.
func (o *MachineTestResponseModel) HasCategoryName() bool {
	if o != nil && o.CategoryName.IsSet() {
		return true
	}

	return false
}

// SetCategoryName gets a reference to the given NullableString and assigns it to the CategoryName field.
func (o *MachineTestResponseModel) SetCategoryName(v string) {
	o.CategoryName.Set(&v)
}

// SetCategoryNameNil sets the value for CategoryName to be an explicit nil
func (o *MachineTestResponseModel) SetCategoryNameNil() {
	o.CategoryName.Set(nil)
}

// UnsetCategoryName ensures that no value is present for CategoryName, not even an explicit nil
func (o *MachineTestResponseModel) UnsetCategoryName() {
	o.CategoryName.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineTestResponseModel) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineTestResponseModel) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *MachineTestResponseModel) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *MachineTestResponseModel) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *MachineTestResponseModel) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *MachineTestResponseModel) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

func (o MachineTestResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MachineTestResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Machine) {
		toSerialize["Machine"] = o.Machine
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if o.CommandResponse.IsSet() {
		toSerialize["CommandResponse"] = o.CommandResponse.Get()
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["ErrorMessage"] = o.ErrorMessage.Get()
	}
	if o.CommandName.IsSet() {
		toSerialize["CommandName"] = o.CommandName.Get()
	}
	if o.CategoryName.IsSet() {
		toSerialize["CategoryName"] = o.CategoryName.Get()
	}
	if o.CreatedBy.IsSet() {
		toSerialize["CreatedBy"] = o.CreatedBy.Get()
	}
	return toSerialize, nil
}

type NullableMachineTestResponseModel struct {
	value *MachineTestResponseModel
	isSet bool
}

func (v NullableMachineTestResponseModel) Get() *MachineTestResponseModel {
	return v.value
}

func (v *NullableMachineTestResponseModel) Set(val *MachineTestResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineTestResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineTestResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineTestResponseModel(val *MachineTestResponseModel) *NullableMachineTestResponseModel {
	return &NullableMachineTestResponseModel{value: val, isSet: true}
}

func (v NullableMachineTestResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineTestResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
