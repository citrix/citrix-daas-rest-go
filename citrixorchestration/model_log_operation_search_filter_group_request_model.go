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

// checks if the LogOperationSearchFilterGroupRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogOperationSearchFilterGroupRequestModel{}

// LogOperationSearchFilterGroupRequestModel Advanced search filter group for configuration logs.
type LogOperationSearchFilterGroupRequestModel struct {
	SearchType *LogOperationSearchFilterGroupType `json:"SearchType,omitempty"`
	// Advanced search filters.
	SearchFilters []LogOperationSearchFilterRequestModel `json:"SearchFilters,omitempty"`
	// Advanced search filter groups.
	SearchFilterGroups []LogOperationSearchFilterGroupRequestModel `json:"SearchFilterGroups,omitempty"`
}

// NewLogOperationSearchFilterGroupRequestModel instantiates a new LogOperationSearchFilterGroupRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogOperationSearchFilterGroupRequestModel() *LogOperationSearchFilterGroupRequestModel {
	this := LogOperationSearchFilterGroupRequestModel{}
	return &this
}

// NewLogOperationSearchFilterGroupRequestModelWithDefaults instantiates a new LogOperationSearchFilterGroupRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogOperationSearchFilterGroupRequestModelWithDefaults() *LogOperationSearchFilterGroupRequestModel {
	this := LogOperationSearchFilterGroupRequestModel{}
	return &this
}

// GetSearchType returns the SearchType field value if set, zero value otherwise.
func (o *LogOperationSearchFilterGroupRequestModel) GetSearchType() LogOperationSearchFilterGroupType {
	if o == nil || IsNil(o.SearchType) {
		var ret LogOperationSearchFilterGroupType
		return ret
	}
	return *o.SearchType
}

// GetSearchTypeOk returns a tuple with the SearchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogOperationSearchFilterGroupRequestModel) GetSearchTypeOk() (*LogOperationSearchFilterGroupType, bool) {
	if o == nil || IsNil(o.SearchType) {
		return nil, false
	}
	return o.SearchType, true
}

// HasSearchType returns a boolean if a field has been set.
func (o *LogOperationSearchFilterGroupRequestModel) HasSearchType() bool {
	if o != nil && !IsNil(o.SearchType) {
		return true
	}

	return false
}

// SetSearchType gets a reference to the given LogOperationSearchFilterGroupType and assigns it to the SearchType field.
func (o *LogOperationSearchFilterGroupRequestModel) SetSearchType(v LogOperationSearchFilterGroupType) {
	o.SearchType = &v
}

// GetSearchFilters returns the SearchFilters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogOperationSearchFilterGroupRequestModel) GetSearchFilters() []LogOperationSearchFilterRequestModel {
	if o == nil {
		var ret []LogOperationSearchFilterRequestModel
		return ret
	}
	return o.SearchFilters
}

// GetSearchFiltersOk returns a tuple with the SearchFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogOperationSearchFilterGroupRequestModel) GetSearchFiltersOk() ([]LogOperationSearchFilterRequestModel, bool) {
	if o == nil || IsNil(o.SearchFilters) {
		return nil, false
	}
	return o.SearchFilters, true
}

// HasSearchFilters returns a boolean if a field has been set.
func (o *LogOperationSearchFilterGroupRequestModel) HasSearchFilters() bool {
	if o != nil && IsNil(o.SearchFilters) {
		return true
	}

	return false
}

// SetSearchFilters gets a reference to the given []LogOperationSearchFilterRequestModel and assigns it to the SearchFilters field.
func (o *LogOperationSearchFilterGroupRequestModel) SetSearchFilters(v []LogOperationSearchFilterRequestModel) {
	o.SearchFilters = v
}

// GetSearchFilterGroups returns the SearchFilterGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogOperationSearchFilterGroupRequestModel) GetSearchFilterGroups() []LogOperationSearchFilterGroupRequestModel {
	if o == nil {
		var ret []LogOperationSearchFilterGroupRequestModel
		return ret
	}
	return o.SearchFilterGroups
}

// GetSearchFilterGroupsOk returns a tuple with the SearchFilterGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogOperationSearchFilterGroupRequestModel) GetSearchFilterGroupsOk() ([]LogOperationSearchFilterGroupRequestModel, bool) {
	if o == nil || IsNil(o.SearchFilterGroups) {
		return nil, false
	}
	return o.SearchFilterGroups, true
}

// HasSearchFilterGroups returns a boolean if a field has been set.
func (o *LogOperationSearchFilterGroupRequestModel) HasSearchFilterGroups() bool {
	if o != nil && IsNil(o.SearchFilterGroups) {
		return true
	}

	return false
}

// SetSearchFilterGroups gets a reference to the given []LogOperationSearchFilterGroupRequestModel and assigns it to the SearchFilterGroups field.
func (o *LogOperationSearchFilterGroupRequestModel) SetSearchFilterGroups(v []LogOperationSearchFilterGroupRequestModel) {
	o.SearchFilterGroups = v
}

func (o LogOperationSearchFilterGroupRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogOperationSearchFilterGroupRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SearchType) {
		toSerialize["SearchType"] = o.SearchType
	}
	if o.SearchFilters != nil {
		toSerialize["SearchFilters"] = o.SearchFilters
	}
	if o.SearchFilterGroups != nil {
		toSerialize["SearchFilterGroups"] = o.SearchFilterGroups
	}
	return toSerialize, nil
}

type NullableLogOperationSearchFilterGroupRequestModel struct {
	value *LogOperationSearchFilterGroupRequestModel
	isSet bool
}

func (v NullableLogOperationSearchFilterGroupRequestModel) Get() *LogOperationSearchFilterGroupRequestModel {
	return v.value
}

func (v *NullableLogOperationSearchFilterGroupRequestModel) Set(val *LogOperationSearchFilterGroupRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLogOperationSearchFilterGroupRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLogOperationSearchFilterGroupRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogOperationSearchFilterGroupRequestModel(val *LogOperationSearchFilterGroupRequestModel) *NullableLogOperationSearchFilterGroupRequestModel {
	return &NullableLogOperationSearchFilterGroupRequestModel{value: val, isSet: true}
}

func (v NullableLogOperationSearchFilterGroupRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogOperationSearchFilterGroupRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

