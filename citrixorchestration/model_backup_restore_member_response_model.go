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

// checks if the BackupRestoreMemberResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupRestoreMemberResponseModel{}

// BackupRestoreMemberResponseModel struct for BackupRestoreMemberResponseModel
type BackupRestoreMemberResponseModel struct {
	// Member Name
	Name NullableString `json:"Name,omitempty"`
	// Member Description
	Description NullableString `json:"Description,omitempty"`
	// Member details (YAML for single emmber)
	Details NullableString `json:"Details,omitempty"`
}

// NewBackupRestoreMemberResponseModel instantiates a new BackupRestoreMemberResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupRestoreMemberResponseModel() *BackupRestoreMemberResponseModel {
	this := BackupRestoreMemberResponseModel{}
	return &this
}

// NewBackupRestoreMemberResponseModelWithDefaults instantiates a new BackupRestoreMemberResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupRestoreMemberResponseModelWithDefaults() *BackupRestoreMemberResponseModel {
	this := BackupRestoreMemberResponseModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreMemberResponseModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreMemberResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *BackupRestoreMemberResponseModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *BackupRestoreMemberResponseModel) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *BackupRestoreMemberResponseModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *BackupRestoreMemberResponseModel) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreMemberResponseModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreMemberResponseModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *BackupRestoreMemberResponseModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *BackupRestoreMemberResponseModel) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *BackupRestoreMemberResponseModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *BackupRestoreMemberResponseModel) UnsetDescription() {
	o.Description.Unset()
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreMemberResponseModel) GetDetails() string {
	if o == nil || IsNil(o.Details.Get()) {
		var ret string
		return ret
	}
	return *o.Details.Get()
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreMemberResponseModel) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Details.Get(), o.Details.IsSet()
}

// HasDetails returns a boolean if a field has been set.
func (o *BackupRestoreMemberResponseModel) HasDetails() bool {
	if o != nil && o.Details.IsSet() {
		return true
	}

	return false
}

// SetDetails gets a reference to the given NullableString and assigns it to the Details field.
func (o *BackupRestoreMemberResponseModel) SetDetails(v string) {
	o.Details.Set(&v)
}

// SetDetailsNil sets the value for Details to be an explicit nil
func (o *BackupRestoreMemberResponseModel) SetDetailsNil() {
	o.Details.Set(nil)
}

// UnsetDetails ensures that no value is present for Details, not even an explicit nil
func (o *BackupRestoreMemberResponseModel) UnsetDetails() {
	o.Details.Unset()
}

func (o BackupRestoreMemberResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupRestoreMemberResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if o.Details.IsSet() {
		toSerialize["Details"] = o.Details.Get()
	}
	return toSerialize, nil
}

type NullableBackupRestoreMemberResponseModel struct {
	value *BackupRestoreMemberResponseModel
	isSet bool
}

func (v NullableBackupRestoreMemberResponseModel) Get() *BackupRestoreMemberResponseModel {
	return v.value
}

func (v *NullableBackupRestoreMemberResponseModel) Set(val *BackupRestoreMemberResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupRestoreMemberResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupRestoreMemberResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupRestoreMemberResponseModel(val *BackupRestoreMemberResponseModel) *NullableBackupRestoreMemberResponseModel {
	return &NullableBackupRestoreMemberResponseModel{value: val, isSet: true}
}

func (v NullableBackupRestoreMemberResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupRestoreMemberResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
