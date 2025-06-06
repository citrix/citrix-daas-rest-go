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

// checks if the ImageVersionRefResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageVersionRefResponseModel{}

// ImageVersionRefResponseModel Image version reference.
type ImageVersionRefResponseModel struct {
	// The Id of the image version.
	Id string `json:"Id"`
	// The version number associated with the image version.
	Number          int32            `json:"Number"`
	ImageDefinition RefResponseModel `json:"ImageDefinition"`
	// The image version's description
	Description NullableString `json:"Description,omitempty"`
	// The image version specifications associated with this image version.
	ImageVersionSpecs []ImageVersionSpecRefResponseModel `json:"ImageVersionSpecs,omitempty"`
}

// NewImageVersionRefResponseModel instantiates a new ImageVersionRefResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageVersionRefResponseModel(id string, number int32, imageDefinition RefResponseModel) *ImageVersionRefResponseModel {
	this := ImageVersionRefResponseModel{}
	this.Id = id
	this.Number = number
	this.ImageDefinition = imageDefinition
	return &this
}

// NewImageVersionRefResponseModelWithDefaults instantiates a new ImageVersionRefResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageVersionRefResponseModelWithDefaults() *ImageVersionRefResponseModel {
	this := ImageVersionRefResponseModel{}
	return &this
}

// GetId returns the Id field value
func (o *ImageVersionRefResponseModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ImageVersionRefResponseModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ImageVersionRefResponseModel) SetId(v string) {
	o.Id = v
}

// GetNumber returns the Number field value
func (o *ImageVersionRefResponseModel) GetNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *ImageVersionRefResponseModel) GetNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *ImageVersionRefResponseModel) SetNumber(v int32) {
	o.Number = v
}

// GetImageDefinition returns the ImageDefinition field value
func (o *ImageVersionRefResponseModel) GetImageDefinition() RefResponseModel {
	if o == nil {
		var ret RefResponseModel
		return ret
	}

	return o.ImageDefinition
}

// GetImageDefinitionOk returns a tuple with the ImageDefinition field value
// and a boolean to check if the value has been set.
func (o *ImageVersionRefResponseModel) GetImageDefinitionOk() (*RefResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageDefinition, true
}

// SetImageDefinition sets field value
func (o *ImageVersionRefResponseModel) SetImageDefinition(v RefResponseModel) {
	o.ImageDefinition = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageVersionRefResponseModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageVersionRefResponseModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ImageVersionRefResponseModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ImageVersionRefResponseModel) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ImageVersionRefResponseModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ImageVersionRefResponseModel) UnsetDescription() {
	o.Description.Unset()
}

// GetImageVersionSpecs returns the ImageVersionSpecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageVersionRefResponseModel) GetImageVersionSpecs() []ImageVersionSpecRefResponseModel {
	if o == nil {
		var ret []ImageVersionSpecRefResponseModel
		return ret
	}
	return o.ImageVersionSpecs
}

// GetImageVersionSpecsOk returns a tuple with the ImageVersionSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageVersionRefResponseModel) GetImageVersionSpecsOk() ([]ImageVersionSpecRefResponseModel, bool) {
	if o == nil || IsNil(o.ImageVersionSpecs) {
		return nil, false
	}
	return o.ImageVersionSpecs, true
}

// HasImageVersionSpecs returns a boolean if a field has been set.
func (o *ImageVersionRefResponseModel) HasImageVersionSpecs() bool {
	if o != nil && IsNil(o.ImageVersionSpecs) {
		return true
	}

	return false
}

// SetImageVersionSpecs gets a reference to the given []ImageVersionSpecRefResponseModel and assigns it to the ImageVersionSpecs field.
func (o *ImageVersionRefResponseModel) SetImageVersionSpecs(v []ImageVersionSpecRefResponseModel) {
	o.ImageVersionSpecs = v
}

func (o ImageVersionRefResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageVersionRefResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Id"] = o.Id
	toSerialize["Number"] = o.Number
	toSerialize["ImageDefinition"] = o.ImageDefinition
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if o.ImageVersionSpecs != nil {
		toSerialize["ImageVersionSpecs"] = o.ImageVersionSpecs
	}
	return toSerialize, nil
}

type NullableImageVersionRefResponseModel struct {
	value *ImageVersionRefResponseModel
	isSet bool
}

func (v NullableImageVersionRefResponseModel) Get() *ImageVersionRefResponseModel {
	return v.value
}

func (v *NullableImageVersionRefResponseModel) Set(val *ImageVersionRefResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableImageVersionRefResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableImageVersionRefResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageVersionRefResponseModel(val *ImageVersionRefResponseModel) *NullableImageVersionRefResponseModel {
	return &NullableImageVersionRefResponseModel{value: val, isSet: true}
}

func (v NullableImageVersionRefResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageVersionRefResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
