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

// checks if the MachineIdSaveAsImageBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MachineIdSaveAsImageBody{}

// MachineIdSaveAsImageBody struct for MachineIdSaveAsImageBody
type MachineIdSaveAsImageBody struct {
	// Id for the image builder machine
	AccountId string `json:"accountId"`
	// Image Name
	ImageName string `json:"imageName"`
	// Image Description
	ImageDescription string `json:"imageDescription"`
	// Image Notes
	ImageNotes NullableString `json:"imageNotes,omitempty"`
}

type _MachineIdSaveAsImageBody MachineIdSaveAsImageBody

// NewMachineIdSaveAsImageBody instantiates a new MachineIdSaveAsImageBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMachineIdSaveAsImageBody(accountId string, imageName string, imageDescription string) *MachineIdSaveAsImageBody {
	this := MachineIdSaveAsImageBody{}
	this.AccountId = accountId
	this.ImageName = imageName
	this.ImageDescription = imageDescription
	return &this
}

// NewMachineIdSaveAsImageBodyWithDefaults instantiates a new MachineIdSaveAsImageBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineIdSaveAsImageBodyWithDefaults() *MachineIdSaveAsImageBody {
	this := MachineIdSaveAsImageBody{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *MachineIdSaveAsImageBody) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *MachineIdSaveAsImageBody) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *MachineIdSaveAsImageBody) SetAccountId(v string) {
	o.AccountId = v
}

// GetImageName returns the ImageName field value
func (o *MachineIdSaveAsImageBody) GetImageName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value
// and a boolean to check if the value has been set.
func (o *MachineIdSaveAsImageBody) GetImageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageName, true
}

// SetImageName sets field value
func (o *MachineIdSaveAsImageBody) SetImageName(v string) {
	o.ImageName = v
}

// GetImageDescription returns the ImageDescription field value
func (o *MachineIdSaveAsImageBody) GetImageDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageDescription
}

// GetImageDescriptionOk returns a tuple with the ImageDescription field value
// and a boolean to check if the value has been set.
func (o *MachineIdSaveAsImageBody) GetImageDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageDescription, true
}

// SetImageDescription sets field value
func (o *MachineIdSaveAsImageBody) SetImageDescription(v string) {
	o.ImageDescription = v
}

// GetImageNotes returns the ImageNotes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MachineIdSaveAsImageBody) GetImageNotes() string {
	if o == nil || IsNil(o.ImageNotes.Get()) {
		var ret string
		return ret
	}
	return *o.ImageNotes.Get()
}

// GetImageNotesOk returns a tuple with the ImageNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MachineIdSaveAsImageBody) GetImageNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageNotes.Get(), o.ImageNotes.IsSet()
}

// HasImageNotes returns a boolean if a field has been set.
func (o *MachineIdSaveAsImageBody) HasImageNotes() bool {
	if o != nil && o.ImageNotes.IsSet() {
		return true
	}

	return false
}

// SetImageNotes gets a reference to the given NullableString and assigns it to the ImageNotes field.
func (o *MachineIdSaveAsImageBody) SetImageNotes(v string) {
	o.ImageNotes.Set(&v)
}

// SetImageNotesNil sets the value for ImageNotes to be an explicit nil
func (o *MachineIdSaveAsImageBody) SetImageNotesNil() {
	o.ImageNotes.Set(nil)
}

// UnsetImageNotes ensures that no value is present for ImageNotes, not even an explicit nil
func (o *MachineIdSaveAsImageBody) UnsetImageNotes() {
	o.ImageNotes.Unset()
}

func (o MachineIdSaveAsImageBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MachineIdSaveAsImageBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountId"] = o.AccountId
	toSerialize["imageName"] = o.ImageName
	toSerialize["imageDescription"] = o.ImageDescription
	if o.ImageNotes.IsSet() {
		toSerialize["imageNotes"] = o.ImageNotes.Get()
	}
	return toSerialize, nil
}

type NullableMachineIdSaveAsImageBody struct {
	value *MachineIdSaveAsImageBody
	isSet bool
}

func (v NullableMachineIdSaveAsImageBody) Get() *MachineIdSaveAsImageBody {
	return v.value
}

func (v *NullableMachineIdSaveAsImageBody) Set(val *MachineIdSaveAsImageBody) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineIdSaveAsImageBody) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineIdSaveAsImageBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineIdSaveAsImageBody(val *MachineIdSaveAsImageBody) *NullableMachineIdSaveAsImageBody {
	return &NullableMachineIdSaveAsImageBody{value: val, isSet: true}
}

func (v NullableMachineIdSaveAsImageBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineIdSaveAsImageBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
