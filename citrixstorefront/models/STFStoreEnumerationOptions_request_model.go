// Copyright © 2025. Citrix Systems, Inc.
package models

// Set-STFStoreEnumerationOptions
var _ MappedNullable = &SetSTFStoreEnumerationOptionsRequestModel{}

type SetSTFStoreEnumerationOptionsRequestModel struct {
	EnhancedEnumeration           NullableBool   // Enable enhanced enumeration. Enumerate multiple farms in parallel to reduce operation time
	MaximumConcurrentEnumerations NullableInt64  // Maximum farms enumerated in parallel
	FilterByTypesInclude          []string       // Inclusive resource filter by type (Applications, Desktops or Documents)
	FilterByKeywordsInclude       []string       // Only include applications and desktops that match the keywords
	FilterByKeywordsExclude       []string       // Exclude applications and desktops that match the keywords
	RequestFullIconData           NullableString // Request icon data (None, Full, FullAndMulti)
}

// ToMap implements MappedNullable
func (o SetSTFStoreEnumerationOptionsRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}

	if o.EnhancedEnumeration.IsSet() {
		toSerialize["EnhancedEnumeration"] = o.EnhancedEnumeration.Get()
	}
	if o.MaximumConcurrentEnumerations.IsSet() {
		toSerialize["MaximumConcurrentEnumerations"] = o.MaximumConcurrentEnumerations.Get()
	}
	if o.FilterByTypesInclude != nil {
		toSerialize["FilterByTypesInclude"] = o.FilterByTypesInclude
	}
	if o.FilterByKeywordsInclude != nil {
		toSerialize["FilterByKeywordsInclude"] = o.FilterByKeywordsInclude
	}
	if o.FilterByKeywordsExclude != nil {
		toSerialize["FilterByKeywordsExclude"] = o.FilterByKeywordsExclude
	}
	if o.RequestFullIconData.IsSet() {
		toSerialize["RequestFullIconData"] = o.RequestFullIconData.Get()
	}
	return toSerialize, nil
}

func (o *SetSTFStoreEnumerationOptionsRequestModel) SetEnhancedEnumeration(v bool) {
	o.EnhancedEnumeration.Set(&v)
}

func (o *SetSTFStoreEnumerationOptionsRequestModel) SetMaximumConcurrentEnumerations(v int64) {
	o.MaximumConcurrentEnumerations.Set(&v)
}

func (o *SetSTFStoreEnumerationOptionsRequestModel) SetFilterByTypesInclude(v []string) {
	o.FilterByTypesInclude = v
}

func (o *SetSTFStoreEnumerationOptionsRequestModel) SetFilterByKeywordsInclude(v []string) {
	o.FilterByKeywordsInclude = v
}

func (o *SetSTFStoreEnumerationOptionsRequestModel) SetFilterByKeywordsExclude(v []string) {
	o.FilterByKeywordsExclude = v
}

func (o *SetSTFStoreEnumerationOptionsRequestModel) SetRequestFullIconData(v string) {
	o.RequestFullIconData.Set(&v)
}
