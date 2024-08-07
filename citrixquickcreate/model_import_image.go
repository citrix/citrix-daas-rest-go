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

// checks if the ImportImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportImage{}

// ImportImage Import image
type ImportImage struct {
	AccountType AccountType `json:"accountType"`
	// Image Name
	Name string `json:"name"`
	// Image Description
	Description NullableString `json:"description,omitempty"`
	// Image Notes
	Notes NullableString `json:"notes,omitempty"`
	SessionSupport NullableSessionSupport `json:"sessionSupport,omitempty"`
}

// NewImportImage instantiates a new ImportImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportImage(accountType AccountType, name string) *ImportImage {
	this := ImportImage{}
	this.AccountType = accountType
	this.Name = name
	return &this
}

// NewImportImageWithDefaults instantiates a new ImportImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportImageWithDefaults() *ImportImage {
	this := ImportImage{}
	return &this
}

// GetAccountType returns the AccountType field value
func (o *ImportImage) GetAccountType() AccountType {
	if o == nil {
		var ret AccountType
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *ImportImage) GetAccountTypeOk() (*AccountType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *ImportImage) SetAccountType(v AccountType) {
	o.AccountType = v
}

// GetName returns the Name field value
func (o *ImportImage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ImportImage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ImportImage) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportImage) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportImage) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ImportImage) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ImportImage) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ImportImage) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ImportImage) UnsetDescription() {
	o.Description.Unset()
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportImage) GetNotes() string {
	if o == nil || IsNil(o.Notes.Get()) {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportImage) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *ImportImage) HasNotes() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *ImportImage) SetNotes(v string) {
	o.Notes.Set(&v)
}
// SetNotesNil sets the value for Notes to be an explicit nil
func (o *ImportImage) SetNotesNil() {
	o.Notes.Set(nil)
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *ImportImage) UnsetNotes() {
	o.Notes.Unset()
}

// GetSessionSupport returns the SessionSupport field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportImage) GetSessionSupport() SessionSupport {
	if o == nil || IsNil(o.SessionSupport.Get()) {
		var ret SessionSupport
		return ret
	}
	return *o.SessionSupport.Get()
}

// GetSessionSupportOk returns a tuple with the SessionSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportImage) GetSessionSupportOk() (*SessionSupport, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionSupport.Get(), o.SessionSupport.IsSet()
}

// HasSessionSupport returns a boolean if a field has been set.
func (o *ImportImage) HasSessionSupport() bool {
	if o != nil && o.SessionSupport.IsSet() {
		return true
	}

	return false
}

// SetSessionSupport gets a reference to the given NullableSessionSupport and assigns it to the SessionSupport field.
func (o *ImportImage) SetSessionSupport(v SessionSupport) {
	o.SessionSupport.Set(&v)
}
// SetSessionSupportNil sets the value for SessionSupport to be an explicit nil
func (o *ImportImage) SetSessionSupportNil() {
	o.SessionSupport.Set(nil)
}

// UnsetSessionSupport ensures that no value is present for SessionSupport, not even an explicit nil
func (o *ImportImage) UnsetSessionSupport() {
	o.SessionSupport.Unset()
}

func (o ImportImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountType"] = o.AccountType
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Notes.IsSet() {
		toSerialize["notes"] = o.Notes.Get()
	}
	if o.SessionSupport.IsSet() {
		toSerialize["sessionSupport"] = o.SessionSupport.Get()
	}
	return toSerialize, nil
}

type NullableImportImage struct {
	value *ImportImage
	isSet bool
}

func (v NullableImportImage) Get() *ImportImage {
	return v.value
}

func (v *NullableImportImage) Set(val *ImportImage) {
	v.value = val
	v.isSet = true
}

func (v NullableImportImage) IsSet() bool {
	return v.isSet
}

func (v *NullableImportImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportImage(val *ImportImage) *NullableImportImage {
	return &NullableImportImage{value: val, isSet: true}
}

func (v NullableImportImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

