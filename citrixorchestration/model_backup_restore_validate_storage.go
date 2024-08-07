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

// checks if the BackupRestoreValidateStorage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupRestoreValidateStorage{}

// BackupRestoreValidateStorage struct for BackupRestoreValidateStorage
type BackupRestoreValidateStorage struct {
	Validated *bool `json:"Validated,omitempty"`
	ValidateStorageErrors map[string]string `json:"ValidateStorageErrors,omitempty"`
	BackupRestoreStorage *BackupRestoreStorageModel2 `json:"BackupRestoreStorage,omitempty"`
}

// NewBackupRestoreValidateStorage instantiates a new BackupRestoreValidateStorage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupRestoreValidateStorage() *BackupRestoreValidateStorage {
	this := BackupRestoreValidateStorage{}
	return &this
}

// NewBackupRestoreValidateStorageWithDefaults instantiates a new BackupRestoreValidateStorage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupRestoreValidateStorageWithDefaults() *BackupRestoreValidateStorage {
	this := BackupRestoreValidateStorage{}
	return &this
}

// GetValidated returns the Validated field value if set, zero value otherwise.
func (o *BackupRestoreValidateStorage) GetValidated() bool {
	if o == nil || IsNil(o.Validated) {
		var ret bool
		return ret
	}
	return *o.Validated
}

// GetValidatedOk returns a tuple with the Validated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreValidateStorage) GetValidatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Validated) {
		return nil, false
	}
	return o.Validated, true
}

// HasValidated returns a boolean if a field has been set.
func (o *BackupRestoreValidateStorage) HasValidated() bool {
	if o != nil && !IsNil(o.Validated) {
		return true
	}

	return false
}

// SetValidated gets a reference to the given bool and assigns it to the Validated field.
func (o *BackupRestoreValidateStorage) SetValidated(v bool) {
	o.Validated = &v
}

// GetValidateStorageErrors returns the ValidateStorageErrors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreValidateStorage) GetValidateStorageErrors() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ValidateStorageErrors
}

// GetValidateStorageErrorsOk returns a tuple with the ValidateStorageErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreValidateStorage) GetValidateStorageErrorsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ValidateStorageErrors) {
		return nil, false
	}
	return &o.ValidateStorageErrors, true
}

// HasValidateStorageErrors returns a boolean if a field has been set.
func (o *BackupRestoreValidateStorage) HasValidateStorageErrors() bool {
	if o != nil && IsNil(o.ValidateStorageErrors) {
		return true
	}

	return false
}

// SetValidateStorageErrors gets a reference to the given map[string]string and assigns it to the ValidateStorageErrors field.
func (o *BackupRestoreValidateStorage) SetValidateStorageErrors(v map[string]string) {
	o.ValidateStorageErrors = v
}

// GetBackupRestoreStorage returns the BackupRestoreStorage field value if set, zero value otherwise.
func (o *BackupRestoreValidateStorage) GetBackupRestoreStorage() BackupRestoreStorageModel2 {
	if o == nil || IsNil(o.BackupRestoreStorage) {
		var ret BackupRestoreStorageModel2
		return ret
	}
	return *o.BackupRestoreStorage
}

// GetBackupRestoreStorageOk returns a tuple with the BackupRestoreStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreValidateStorage) GetBackupRestoreStorageOk() (*BackupRestoreStorageModel2, bool) {
	if o == nil || IsNil(o.BackupRestoreStorage) {
		return nil, false
	}
	return o.BackupRestoreStorage, true
}

// HasBackupRestoreStorage returns a boolean if a field has been set.
func (o *BackupRestoreValidateStorage) HasBackupRestoreStorage() bool {
	if o != nil && !IsNil(o.BackupRestoreStorage) {
		return true
	}

	return false
}

// SetBackupRestoreStorage gets a reference to the given BackupRestoreStorageModel2 and assigns it to the BackupRestoreStorage field.
func (o *BackupRestoreValidateStorage) SetBackupRestoreStorage(v BackupRestoreStorageModel2) {
	o.BackupRestoreStorage = &v
}

func (o BackupRestoreValidateStorage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupRestoreValidateStorage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Validated) {
		toSerialize["Validated"] = o.Validated
	}
	if o.ValidateStorageErrors != nil {
		toSerialize["ValidateStorageErrors"] = o.ValidateStorageErrors
	}
	if !IsNil(o.BackupRestoreStorage) {
		toSerialize["BackupRestoreStorage"] = o.BackupRestoreStorage
	}
	return toSerialize, nil
}

type NullableBackupRestoreValidateStorage struct {
	value *BackupRestoreValidateStorage
	isSet bool
}

func (v NullableBackupRestoreValidateStorage) Get() *BackupRestoreValidateStorage {
	return v.value
}

func (v *NullableBackupRestoreValidateStorage) Set(val *BackupRestoreValidateStorage) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupRestoreValidateStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupRestoreValidateStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupRestoreValidateStorage(val *BackupRestoreValidateStorage) *NullableBackupRestoreValidateStorage {
	return &NullableBackupRestoreValidateStorage{value: val, isSet: true}
}

func (v NullableBackupRestoreValidateStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupRestoreValidateStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

