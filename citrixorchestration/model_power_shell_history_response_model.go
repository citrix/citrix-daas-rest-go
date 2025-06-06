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

// checks if the PowerShellHistoryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PowerShellHistoryResponseModel{}

// PowerShellHistoryResponseModel The response model of PowerShell script history
type PowerShellHistoryResponseModel struct {
	// The title of the PowerShell script history
	Title NullableString `json:"Title,omitempty"`
	// The execution time of the PowerShell
	ExecutionTime NullableString             `json:"ExecutionTime,omitempty"`
	Status        *PowerShellExecutionStatus `json:"Status,omitempty"`
	// The total commands in the execution
	TotalCommands *int32 `json:"TotalCommands,omitempty"`
	// The executed PowerShell commands
	Commands []PowerShellExecutedCommandModel `json:"Commands,omitempty"`
}

// NewPowerShellHistoryResponseModel instantiates a new PowerShellHistoryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerShellHistoryResponseModel() *PowerShellHistoryResponseModel {
	this := PowerShellHistoryResponseModel{}
	return &this
}

// NewPowerShellHistoryResponseModelWithDefaults instantiates a new PowerShellHistoryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerShellHistoryResponseModelWithDefaults() *PowerShellHistoryResponseModel {
	this := PowerShellHistoryResponseModel{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerShellHistoryResponseModel) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerShellHistoryResponseModel) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *PowerShellHistoryResponseModel) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *PowerShellHistoryResponseModel) SetTitle(v string) {
	o.Title.Set(&v)
}

// SetTitleNil sets the value for Title to be an explicit nil
func (o *PowerShellHistoryResponseModel) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *PowerShellHistoryResponseModel) UnsetTitle() {
	o.Title.Unset()
}

// GetExecutionTime returns the ExecutionTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerShellHistoryResponseModel) GetExecutionTime() string {
	if o == nil || IsNil(o.ExecutionTime.Get()) {
		var ret string
		return ret
	}
	return *o.ExecutionTime.Get()
}

// GetExecutionTimeOk returns a tuple with the ExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerShellHistoryResponseModel) GetExecutionTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutionTime.Get(), o.ExecutionTime.IsSet()
}

// HasExecutionTime returns a boolean if a field has been set.
func (o *PowerShellHistoryResponseModel) HasExecutionTime() bool {
	if o != nil && o.ExecutionTime.IsSet() {
		return true
	}

	return false
}

// SetExecutionTime gets a reference to the given NullableString and assigns it to the ExecutionTime field.
func (o *PowerShellHistoryResponseModel) SetExecutionTime(v string) {
	o.ExecutionTime.Set(&v)
}

// SetExecutionTimeNil sets the value for ExecutionTime to be an explicit nil
func (o *PowerShellHistoryResponseModel) SetExecutionTimeNil() {
	o.ExecutionTime.Set(nil)
}

// UnsetExecutionTime ensures that no value is present for ExecutionTime, not even an explicit nil
func (o *PowerShellHistoryResponseModel) UnsetExecutionTime() {
	o.ExecutionTime.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PowerShellHistoryResponseModel) GetStatus() PowerShellExecutionStatus {
	if o == nil || IsNil(o.Status) {
		var ret PowerShellExecutionStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerShellHistoryResponseModel) GetStatusOk() (*PowerShellExecutionStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PowerShellHistoryResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PowerShellExecutionStatus and assigns it to the Status field.
func (o *PowerShellHistoryResponseModel) SetStatus(v PowerShellExecutionStatus) {
	o.Status = &v
}

// GetTotalCommands returns the TotalCommands field value if set, zero value otherwise.
func (o *PowerShellHistoryResponseModel) GetTotalCommands() int32 {
	if o == nil || IsNil(o.TotalCommands) {
		var ret int32
		return ret
	}
	return *o.TotalCommands
}

// GetTotalCommandsOk returns a tuple with the TotalCommands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerShellHistoryResponseModel) GetTotalCommandsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCommands) {
		return nil, false
	}
	return o.TotalCommands, true
}

// HasTotalCommands returns a boolean if a field has been set.
func (o *PowerShellHistoryResponseModel) HasTotalCommands() bool {
	if o != nil && !IsNil(o.TotalCommands) {
		return true
	}

	return false
}

// SetTotalCommands gets a reference to the given int32 and assigns it to the TotalCommands field.
func (o *PowerShellHistoryResponseModel) SetTotalCommands(v int32) {
	o.TotalCommands = &v
}

// GetCommands returns the Commands field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerShellHistoryResponseModel) GetCommands() []PowerShellExecutedCommandModel {
	if o == nil {
		var ret []PowerShellExecutedCommandModel
		return ret
	}
	return o.Commands
}

// GetCommandsOk returns a tuple with the Commands field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerShellHistoryResponseModel) GetCommandsOk() ([]PowerShellExecutedCommandModel, bool) {
	if o == nil || IsNil(o.Commands) {
		return nil, false
	}
	return o.Commands, true
}

// HasCommands returns a boolean if a field has been set.
func (o *PowerShellHistoryResponseModel) HasCommands() bool {
	if o != nil && IsNil(o.Commands) {
		return true
	}

	return false
}

// SetCommands gets a reference to the given []PowerShellExecutedCommandModel and assigns it to the Commands field.
func (o *PowerShellHistoryResponseModel) SetCommands(v []PowerShellExecutedCommandModel) {
	o.Commands = v
}

func (o PowerShellHistoryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PowerShellHistoryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Title.IsSet() {
		toSerialize["Title"] = o.Title.Get()
	}
	if o.ExecutionTime.IsSet() {
		toSerialize["ExecutionTime"] = o.ExecutionTime.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.TotalCommands) {
		toSerialize["TotalCommands"] = o.TotalCommands
	}
	if o.Commands != nil {
		toSerialize["Commands"] = o.Commands
	}
	return toSerialize, nil
}

type NullablePowerShellHistoryResponseModel struct {
	value *PowerShellHistoryResponseModel
	isSet bool
}

func (v NullablePowerShellHistoryResponseModel) Get() *PowerShellHistoryResponseModel {
	return v.value
}

func (v *NullablePowerShellHistoryResponseModel) Set(val *PowerShellHistoryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerShellHistoryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerShellHistoryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerShellHistoryResponseModel(val *PowerShellHistoryResponseModel) *NullablePowerShellHistoryResponseModel {
	return &NullablePowerShellHistoryResponseModel{value: val, isSet: true}
}

func (v NullablePowerShellHistoryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerShellHistoryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
