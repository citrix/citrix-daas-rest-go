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

// checks if the ResourceValidationReportModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceValidationReportModel{}

// ResourceValidationReportModel The validation report.
type ResourceValidationReportModel struct {
	Category *ResourceValidationCategory `json:"Category,omitempty"`
	Result   *ResourceValidationResult   `json:"Result,omitempty"`
	// A list of violations.
	Violations []ResourceValidationViolationModel `json:"Violations,omitempty"`
}

// NewResourceValidationReportModel instantiates a new ResourceValidationReportModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceValidationReportModel() *ResourceValidationReportModel {
	this := ResourceValidationReportModel{}
	return &this
}

// NewResourceValidationReportModelWithDefaults instantiates a new ResourceValidationReportModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceValidationReportModelWithDefaults() *ResourceValidationReportModel {
	this := ResourceValidationReportModel{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ResourceValidationReportModel) GetCategory() ResourceValidationCategory {
	if o == nil || IsNil(o.Category) {
		var ret ResourceValidationCategory
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceValidationReportModel) GetCategoryOk() (*ResourceValidationCategory, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ResourceValidationReportModel) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given ResourceValidationCategory and assigns it to the Category field.
func (o *ResourceValidationReportModel) SetCategory(v ResourceValidationCategory) {
	o.Category = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ResourceValidationReportModel) GetResult() ResourceValidationResult {
	if o == nil || IsNil(o.Result) {
		var ret ResourceValidationResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceValidationReportModel) GetResultOk() (*ResourceValidationResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ResourceValidationReportModel) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ResourceValidationResult and assigns it to the Result field.
func (o *ResourceValidationReportModel) SetResult(v ResourceValidationResult) {
	o.Result = &v
}

// GetViolations returns the Violations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceValidationReportModel) GetViolations() []ResourceValidationViolationModel {
	if o == nil {
		var ret []ResourceValidationViolationModel
		return ret
	}
	return o.Violations
}

// GetViolationsOk returns a tuple with the Violations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceValidationReportModel) GetViolationsOk() ([]ResourceValidationViolationModel, bool) {
	if o == nil || IsNil(o.Violations) {
		return nil, false
	}
	return o.Violations, true
}

// HasViolations returns a boolean if a field has been set.
func (o *ResourceValidationReportModel) HasViolations() bool {
	if o != nil && IsNil(o.Violations) {
		return true
	}

	return false
}

// SetViolations gets a reference to the given []ResourceValidationViolationModel and assigns it to the Violations field.
func (o *ResourceValidationReportModel) SetViolations(v []ResourceValidationViolationModel) {
	o.Violations = v
}

func (o ResourceValidationReportModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceValidationReportModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Category) {
		toSerialize["Category"] = o.Category
	}
	if !IsNil(o.Result) {
		toSerialize["Result"] = o.Result
	}
	if o.Violations != nil {
		toSerialize["Violations"] = o.Violations
	}
	return toSerialize, nil
}

type NullableResourceValidationReportModel struct {
	value *ResourceValidationReportModel
	isSet bool
}

func (v NullableResourceValidationReportModel) Get() *ResourceValidationReportModel {
	return v.value
}

func (v *NullableResourceValidationReportModel) Set(val *ResourceValidationReportModel) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceValidationReportModel) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceValidationReportModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceValidationReportModel(val *ResourceValidationReportModel) *NullableResourceValidationReportModel {
	return &NullableResourceValidationReportModel{value: val, isSet: true}
}

func (v NullableResourceValidationReportModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceValidationReportModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
