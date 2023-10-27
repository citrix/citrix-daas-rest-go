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

// checks if the PowerShellExecutedCommandModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PowerShellExecutedCommandModel{}

// PowerShellExecutedCommandModel The executed command
type PowerShellExecutedCommandModel struct {
	// The command
	Command NullableString `json:"Command,omitempty"`
	// The errors when the command reported
	Errors []PowerShellCommandErrorModel `json:"Errors,omitempty"`
}

// NewPowerShellExecutedCommandModel instantiates a new PowerShellExecutedCommandModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerShellExecutedCommandModel() *PowerShellExecutedCommandModel {
	this := PowerShellExecutedCommandModel{}
	return &this
}

// NewPowerShellExecutedCommandModelWithDefaults instantiates a new PowerShellExecutedCommandModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerShellExecutedCommandModelWithDefaults() *PowerShellExecutedCommandModel {
	this := PowerShellExecutedCommandModel{}
	return &this
}

// GetCommand returns the Command field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerShellExecutedCommandModel) GetCommand() string {
	if o == nil || IsNil(o.Command.Get()) {
		var ret string
		return ret
	}
	return *o.Command.Get()
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerShellExecutedCommandModel) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Command.Get(), o.Command.IsSet()
}

// HasCommand returns a boolean if a field has been set.
func (o *PowerShellExecutedCommandModel) HasCommand() bool {
	if o != nil && o.Command.IsSet() {
		return true
	}

	return false
}

// SetCommand gets a reference to the given NullableString and assigns it to the Command field.
func (o *PowerShellExecutedCommandModel) SetCommand(v string) {
	o.Command.Set(&v)
}
// SetCommandNil sets the value for Command to be an explicit nil
func (o *PowerShellExecutedCommandModel) SetCommandNil() {
	o.Command.Set(nil)
}

// UnsetCommand ensures that no value is present for Command, not even an explicit nil
func (o *PowerShellExecutedCommandModel) UnsetCommand() {
	o.Command.Unset()
}

// GetErrors returns the Errors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerShellExecutedCommandModel) GetErrors() []PowerShellCommandErrorModel {
	if o == nil {
		var ret []PowerShellCommandErrorModel
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerShellExecutedCommandModel) GetErrorsOk() ([]PowerShellCommandErrorModel, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *PowerShellExecutedCommandModel) HasErrors() bool {
	if o != nil && IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []PowerShellCommandErrorModel and assigns it to the Errors field.
func (o *PowerShellExecutedCommandModel) SetErrors(v []PowerShellCommandErrorModel) {
	o.Errors = v
}

func (o PowerShellExecutedCommandModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PowerShellExecutedCommandModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Command.IsSet() {
		toSerialize["Command"] = o.Command.Get()
	}
	if o.Errors != nil {
		toSerialize["Errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullablePowerShellExecutedCommandModel struct {
	value *PowerShellExecutedCommandModel
	isSet bool
}

func (v NullablePowerShellExecutedCommandModel) Get() *PowerShellExecutedCommandModel {
	return v.value
}

func (v *NullablePowerShellExecutedCommandModel) Set(val *PowerShellExecutedCommandModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerShellExecutedCommandModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerShellExecutedCommandModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerShellExecutedCommandModel(val *PowerShellExecutedCommandModel) *NullablePowerShellExecutedCommandModel {
	return &NullablePowerShellExecutedCommandModel{value: val, isSet: true}
}

func (v NullablePowerShellExecutedCommandModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerShellExecutedCommandModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

