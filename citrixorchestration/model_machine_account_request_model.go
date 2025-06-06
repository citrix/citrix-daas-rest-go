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

// checks if the MachineAccountRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MachineAccountRequestModel{}

// MachineAccountRequestModel Active directory machine account which was externally created.
type MachineAccountRequestModel struct {
	// The Active Directory account name to be imported.
	ADAccountName string `json:"ADAccountName" validate:"regexp=.*"`
	// Indicates whether the machine account password should be reset.
	ResetPassword NullableBool `json:"ResetPassword,omitempty"`
	// The current password for the machine account , in the format specified by PasswordFormat.
	Password       NullableString          `json:"Password,omitempty"`
	PasswordFormat *IdentityPasswordFormat `json:"PasswordFormat,omitempty"`
}

// NewMachineAccountRequestModel instantiates a new MachineAccountRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMachineAccountRequestModel(aDAccountName string) *MachineAccountRequestModel {
	this := MachineAccountRequestModel{}
	this.ADAccountName = aDAccountName
	var resetPassword bool = true
	this.ResetPassword = *NewNullableBool(&resetPassword)
	return &this
}

// NewMachineAccountRequestModelWithDefaults instantiates a new MachineAccountRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineAccountRequestModelWithDefaults() *MachineAccountRequestModel {
	this := MachineAccountRequestModel{}
	var resetPassword bool = true
	this.ResetPassword = *NewNullableBool(&resetPassword)
	return &this
}

// GetADAccountName returns the ADAccountName field value
func (o *MachineAccountRequestModel) GetADAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ADAccountName
}

// GetADAccountNameOk returns a tuple with the ADAccountName field value
// and a boolean to check if the value has been set.
func (o *MachineAccountRequestModel) GetADAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ADAccountName, true
}

// SetADAccountName sets field value
func (o *MachineAccountRequestModel) SetADAccountName(v string) {
	o.ADAccountName = v
}

// GetResetPassword returns the ResetPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineAccountRequestModel) GetResetPassword() bool {
	if o == nil || IsNil(o.ResetPassword.Get()) {
		var ret bool
		return ret
	}
	return *o.ResetPassword.Get()
}

// GetResetPasswordOk returns a tuple with the ResetPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineAccountRequestModel) GetResetPasswordOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResetPassword.Get(), o.ResetPassword.IsSet()
}

// HasResetPassword returns a boolean if a field has been set.
func (o *MachineAccountRequestModel) HasResetPassword() bool {
	if o != nil && o.ResetPassword.IsSet() {
		return true
	}

	return false
}

// SetResetPassword gets a reference to the given NullableBool and assigns it to the ResetPassword field.
func (o *MachineAccountRequestModel) SetResetPassword(v bool) {
	o.ResetPassword.Set(&v)
}

// SetResetPasswordNil sets the value for ResetPassword to be an explicit nil
func (o *MachineAccountRequestModel) SetResetPasswordNil() {
	o.ResetPassword.Set(nil)
}

// UnsetResetPassword ensures that no value is present for ResetPassword, not even an explicit nil
func (o *MachineAccountRequestModel) UnsetResetPassword() {
	o.ResetPassword.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineAccountRequestModel) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineAccountRequestModel) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *MachineAccountRequestModel) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *MachineAccountRequestModel) SetPassword(v string) {
	o.Password.Set(&v)
}

// SetPasswordNil sets the value for Password to be an explicit nil
func (o *MachineAccountRequestModel) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *MachineAccountRequestModel) UnsetPassword() {
	o.Password.Unset()
}

// GetPasswordFormat returns the PasswordFormat field value if set, zero value otherwise.
func (o *MachineAccountRequestModel) GetPasswordFormat() IdentityPasswordFormat {
	if o == nil || IsNil(o.PasswordFormat) {
		var ret IdentityPasswordFormat
		return ret
	}
	return *o.PasswordFormat
}

// GetPasswordFormatOk returns a tuple with the PasswordFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineAccountRequestModel) GetPasswordFormatOk() (*IdentityPasswordFormat, bool) {
	if o == nil || IsNil(o.PasswordFormat) {
		return nil, false
	}
	return o.PasswordFormat, true
}

// HasPasswordFormat returns a boolean if a field has been set.
func (o *MachineAccountRequestModel) HasPasswordFormat() bool {
	if o != nil && !IsNil(o.PasswordFormat) {
		return true
	}

	return false
}

// SetPasswordFormat gets a reference to the given IdentityPasswordFormat and assigns it to the PasswordFormat field.
func (o *MachineAccountRequestModel) SetPasswordFormat(v IdentityPasswordFormat) {
	o.PasswordFormat = &v
}

func (o MachineAccountRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MachineAccountRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ADAccountName"] = o.ADAccountName
	if o.ResetPassword.IsSet() {
		toSerialize["ResetPassword"] = o.ResetPassword.Get()
	}
	if o.Password.IsSet() {
		toSerialize["Password"] = o.Password.Get()
	}
	if !IsNil(o.PasswordFormat) {
		toSerialize["PasswordFormat"] = o.PasswordFormat
	}
	return toSerialize, nil
}

type NullableMachineAccountRequestModel struct {
	value *MachineAccountRequestModel
	isSet bool
}

func (v NullableMachineAccountRequestModel) Get() *MachineAccountRequestModel {
	return v.value
}

func (v *NullableMachineAccountRequestModel) Set(val *MachineAccountRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineAccountRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineAccountRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineAccountRequestModel(val *MachineAccountRequestModel) *NullableMachineAccountRequestModel {
	return &NullableMachineAccountRequestModel{value: val, isSet: true}
}

func (v NullableMachineAccountRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineAccountRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
