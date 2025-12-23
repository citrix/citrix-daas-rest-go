// Copyright © 2025. Citrix Systems, Inc.
package models

import (
	"strings"
)

type GetSTFStoreEnumerationOptionsRawResponseModel struct {
	EnhancedEnumeration                          NullableBool                 `json:"EnhancedEnumeration,omitempty"`
	MaximumConcurrentEnumerations                NullableInt64                `json:"MaximumConcurrentEnumerations,omitempty"`
	MinimumFarmsRequiredForConcurrentEnumeration NullableInt64                `json:"MinimumFarmsRequiredForConcurrentEnumeration,omitempty"`
	RequestFullIconData                          NullableInt64                `json:"RequestFullIconData,omitempty"`
	RequestedHighColorIcons                      RequestedHighColorIconsModel `json:"RequestedHighColorIcons,omitempty"`
	FilterByTypesInclude                         NullableString               `json:"FilterByTypesInclude,omitempty"`
	FilterByKeywordsInclude                      NullableString               `json:"FilterByKeywordsInclude,omitempty"`
	FilterByKeywordsExclude                      NullableString               `json:"FilterByKeywordsExclude,omitempty"`
	TreatDesktopsAsApps                          NullableBool                 `json:"TreatDesktopsAsApps,omitempty"`
	Authenticator                                NullableString               `json:"Authenticator,omitempty"`
}

type GetSTFStoreEnumerationOptionsResponseModel struct {
	EnhancedEnumeration                          NullableBool                 `json:"EnhancedEnumeration,omitempty"`
	MaximumConcurrentEnumerations                NullableInt64                `json:"MaximumConcurrentEnumerations,omitempty"`
	MinimumFarmsRequiredForConcurrentEnumeration NullableInt64                `json:"MinimumFarmsRequiredForConcurrentEnumeration,omitempty"`
	RequestedHighColorIcons                      RequestedHighColorIconsModel `json:"RequestedHighColorIcons,omitempty"`
	FilterByTypesInclude                         []string                     `json:"FilterByTypesInclude,omitempty"`
	FilterByKeywordsInclude                      []string                     `json:"FilterByKeywordsInclude,omitempty"`
	FilterByKeywordsExclude                      []string                     `json:"FilterByKeywordsExclude,omitempty"`
	TreatDesktopsAsApps                          NullableBool                 `json:"TreatDesktopsAsApps,omitempty"`
	Authenticator                                NullableString               `json:"Authenticator,omitempty"`
}

type RequestedHighColorIconsModel struct {
	Small  NullableInt64 `json:"small"`
	Medium NullableInt64 `json:"medium"`
	Large  NullableInt64 `json:"large"`
}

func (rawResponse GetSTFStoreEnumerationOptionsRawResponseModel) ConvertToResponseModel() GetSTFStoreEnumerationOptionsResponseModel {
	response := GetSTFStoreEnumerationOptionsResponseModel{}

	response.EnhancedEnumeration = rawResponse.EnhancedEnumeration
	response.MaximumConcurrentEnumerations = rawResponse.MaximumConcurrentEnumerations
	response.MinimumFarmsRequiredForConcurrentEnumeration = rawResponse.MinimumFarmsRequiredForConcurrentEnumeration
	response.RequestedHighColorIcons = rawResponse.RequestedHighColorIcons
	if rawResponse.FilterByTypesInclude.IsSet() && *rawResponse.FilterByTypesInclude.value != "" {
		value := rawResponse.FilterByTypesInclude.Get()
		response.FilterByTypesInclude = strings.Split(*value, ";")
	}
	if rawResponse.FilterByKeywordsInclude.IsSet() && *rawResponse.FilterByKeywordsInclude.value != "" {
		value := rawResponse.FilterByKeywordsInclude.Get()
		response.FilterByKeywordsInclude = strings.Split(*value, ";")
	}
	if rawResponse.FilterByKeywordsExclude.IsSet() && *rawResponse.FilterByKeywordsExclude.value != "" {
		value := rawResponse.FilterByKeywordsExclude.Get()
		response.FilterByKeywordsExclude = strings.Split(*value, ";")
	}
	response.TreatDesktopsAsApps = rawResponse.TreatDesktopsAsApps
	response.Authenticator = rawResponse.Authenticator

	return response
}
