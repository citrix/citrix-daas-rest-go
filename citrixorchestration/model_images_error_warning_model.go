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

// checks if the ImagesErrorWarningModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImagesErrorWarningModel{}

// ImagesErrorWarningModel struct for ImagesErrorWarningModel
type ImagesErrorWarningModel struct {
	NumberOfErrors   *int32                             `json:"NumberOfErrors,omitempty"`
	NumberOfWarnings *int32                             `json:"NumberOfWarnings,omitempty"`
	ErrorWarning     []ImageDefinitionErrorWarningModel `json:"ErrorWarning,omitempty"`
}

// NewImagesErrorWarningModel instantiates a new ImagesErrorWarningModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImagesErrorWarningModel() *ImagesErrorWarningModel {
	this := ImagesErrorWarningModel{}
	return &this
}

// NewImagesErrorWarningModelWithDefaults instantiates a new ImagesErrorWarningModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImagesErrorWarningModelWithDefaults() *ImagesErrorWarningModel {
	this := ImagesErrorWarningModel{}
	return &this
}

// GetNumberOfErrors returns the NumberOfErrors field value if set, zero value otherwise.
func (o *ImagesErrorWarningModel) GetNumberOfErrors() int32 {
	if o == nil || IsNil(o.NumberOfErrors) {
		var ret int32
		return ret
	}
	return *o.NumberOfErrors
}

// GetNumberOfErrorsOk returns a tuple with the NumberOfErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImagesErrorWarningModel) GetNumberOfErrorsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfErrors) {
		return nil, false
	}
	return o.NumberOfErrors, true
}

// HasNumberOfErrors returns a boolean if a field has been set.
func (o *ImagesErrorWarningModel) HasNumberOfErrors() bool {
	if o != nil && !IsNil(o.NumberOfErrors) {
		return true
	}

	return false
}

// SetNumberOfErrors gets a reference to the given int32 and assigns it to the NumberOfErrors field.
func (o *ImagesErrorWarningModel) SetNumberOfErrors(v int32) {
	o.NumberOfErrors = &v
}

// GetNumberOfWarnings returns the NumberOfWarnings field value if set, zero value otherwise.
func (o *ImagesErrorWarningModel) GetNumberOfWarnings() int32 {
	if o == nil || IsNil(o.NumberOfWarnings) {
		var ret int32
		return ret
	}
	return *o.NumberOfWarnings
}

// GetNumberOfWarningsOk returns a tuple with the NumberOfWarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImagesErrorWarningModel) GetNumberOfWarningsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfWarnings) {
		return nil, false
	}
	return o.NumberOfWarnings, true
}

// HasNumberOfWarnings returns a boolean if a field has been set.
func (o *ImagesErrorWarningModel) HasNumberOfWarnings() bool {
	if o != nil && !IsNil(o.NumberOfWarnings) {
		return true
	}

	return false
}

// SetNumberOfWarnings gets a reference to the given int32 and assigns it to the NumberOfWarnings field.
func (o *ImagesErrorWarningModel) SetNumberOfWarnings(v int32) {
	o.NumberOfWarnings = &v
}

// GetErrorWarning returns the ErrorWarning field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImagesErrorWarningModel) GetErrorWarning() []ImageDefinitionErrorWarningModel {
	if o == nil {
		var ret []ImageDefinitionErrorWarningModel
		return ret
	}
	return o.ErrorWarning
}

// GetErrorWarningOk returns a tuple with the ErrorWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImagesErrorWarningModel) GetErrorWarningOk() ([]ImageDefinitionErrorWarningModel, bool) {
	if o == nil || IsNil(o.ErrorWarning) {
		return nil, false
	}
	return o.ErrorWarning, true
}

// HasErrorWarning returns a boolean if a field has been set.
func (o *ImagesErrorWarningModel) HasErrorWarning() bool {
	if o != nil && IsNil(o.ErrorWarning) {
		return true
	}

	return false
}

// SetErrorWarning gets a reference to the given []ImageDefinitionErrorWarningModel and assigns it to the ErrorWarning field.
func (o *ImagesErrorWarningModel) SetErrorWarning(v []ImageDefinitionErrorWarningModel) {
	o.ErrorWarning = v
}

func (o ImagesErrorWarningModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImagesErrorWarningModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumberOfErrors) {
		toSerialize["NumberOfErrors"] = o.NumberOfErrors
	}
	if !IsNil(o.NumberOfWarnings) {
		toSerialize["NumberOfWarnings"] = o.NumberOfWarnings
	}
	if o.ErrorWarning != nil {
		toSerialize["ErrorWarning"] = o.ErrorWarning
	}
	return toSerialize, nil
}

type NullableImagesErrorWarningModel struct {
	value *ImagesErrorWarningModel
	isSet bool
}

func (v NullableImagesErrorWarningModel) Get() *ImagesErrorWarningModel {
	return v.value
}

func (v *NullableImagesErrorWarningModel) Set(val *ImagesErrorWarningModel) {
	v.value = val
	v.isSet = true
}

func (v NullableImagesErrorWarningModel) IsSet() bool {
	return v.isSet
}

func (v *NullableImagesErrorWarningModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImagesErrorWarningModel(val *ImagesErrorWarningModel) *NullableImagesErrorWarningModel {
	return &NullableImagesErrorWarningModel{value: val, isSet: true}
}

func (v NullableImagesErrorWarningModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImagesErrorWarningModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
