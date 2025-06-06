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

// checks if the ResourceValidationViolationModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceValidationViolationModel{}

// ResourceValidationViolationModel A resource validation violation entry.
type ResourceValidationViolationModel struct {
	Level ResourceViolationLevel `json:"Level"`
	// A violation message.
	Message string `json:"Message"`
	// The attached detail, could be null.
	Detail NullableString `json:"Detail,omitempty"`
	// The relative path of the resource which owns this violation, could be null.
	RelativePath NullableString `json:"RelativePath,omitempty"`
}

// NewResourceValidationViolationModel instantiates a new ResourceValidationViolationModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceValidationViolationModel(level ResourceViolationLevel, message string) *ResourceValidationViolationModel {
	this := ResourceValidationViolationModel{}
	this.Level = level
	this.Message = message
	return &this
}

// NewResourceValidationViolationModelWithDefaults instantiates a new ResourceValidationViolationModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceValidationViolationModelWithDefaults() *ResourceValidationViolationModel {
	this := ResourceValidationViolationModel{}
	return &this
}

// GetLevel returns the Level field value
func (o *ResourceValidationViolationModel) GetLevel() ResourceViolationLevel {
	if o == nil {
		var ret ResourceViolationLevel
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *ResourceValidationViolationModel) GetLevelOk() (*ResourceViolationLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *ResourceValidationViolationModel) SetLevel(v ResourceViolationLevel) {
	o.Level = v
}

// GetMessage returns the Message field value
func (o *ResourceValidationViolationModel) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ResourceValidationViolationModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ResourceValidationViolationModel) SetMessage(v string) {
	o.Message = v
}

// GetDetail returns the Detail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceValidationViolationModel) GetDetail() string {
	if o == nil || IsNil(o.Detail.Get()) {
		var ret string
		return ret
	}
	return *o.Detail.Get()
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceValidationViolationModel) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Detail.Get(), o.Detail.IsSet()
}

// HasDetail returns a boolean if a field has been set.
func (o *ResourceValidationViolationModel) HasDetail() bool {
	if o != nil && o.Detail.IsSet() {
		return true
	}

	return false
}

// SetDetail gets a reference to the given NullableString and assigns it to the Detail field.
func (o *ResourceValidationViolationModel) SetDetail(v string) {
	o.Detail.Set(&v)
}

// SetDetailNil sets the value for Detail to be an explicit nil
func (o *ResourceValidationViolationModel) SetDetailNil() {
	o.Detail.Set(nil)
}

// UnsetDetail ensures that no value is present for Detail, not even an explicit nil
func (o *ResourceValidationViolationModel) UnsetDetail() {
	o.Detail.Unset()
}

// GetRelativePath returns the RelativePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceValidationViolationModel) GetRelativePath() string {
	if o == nil || IsNil(o.RelativePath.Get()) {
		var ret string
		return ret
	}
	return *o.RelativePath.Get()
}

// GetRelativePathOk returns a tuple with the RelativePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceValidationViolationModel) GetRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RelativePath.Get(), o.RelativePath.IsSet()
}

// HasRelativePath returns a boolean if a field has been set.
func (o *ResourceValidationViolationModel) HasRelativePath() bool {
	if o != nil && o.RelativePath.IsSet() {
		return true
	}

	return false
}

// SetRelativePath gets a reference to the given NullableString and assigns it to the RelativePath field.
func (o *ResourceValidationViolationModel) SetRelativePath(v string) {
	o.RelativePath.Set(&v)
}

// SetRelativePathNil sets the value for RelativePath to be an explicit nil
func (o *ResourceValidationViolationModel) SetRelativePathNil() {
	o.RelativePath.Set(nil)
}

// UnsetRelativePath ensures that no value is present for RelativePath, not even an explicit nil
func (o *ResourceValidationViolationModel) UnsetRelativePath() {
	o.RelativePath.Unset()
}

func (o ResourceValidationViolationModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceValidationViolationModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Level"] = o.Level
	toSerialize["Message"] = o.Message
	if o.Detail.IsSet() {
		toSerialize["Detail"] = o.Detail.Get()
	}
	if o.RelativePath.IsSet() {
		toSerialize["RelativePath"] = o.RelativePath.Get()
	}
	return toSerialize, nil
}

type NullableResourceValidationViolationModel struct {
	value *ResourceValidationViolationModel
	isSet bool
}

func (v NullableResourceValidationViolationModel) Get() *ResourceValidationViolationModel {
	return v.value
}

func (v *NullableResourceValidationViolationModel) Set(val *ResourceValidationViolationModel) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceValidationViolationModel) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceValidationViolationModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceValidationViolationModel(val *ResourceValidationViolationModel) *NullableResourceValidationViolationModel {
	return &NullableResourceValidationViolationModel{value: val, isSet: true}
}

func (v NullableResourceValidationViolationModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceValidationViolationModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
