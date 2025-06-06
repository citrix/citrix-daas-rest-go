/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-onprem-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
	"fmt"
)

// FlashUrlRenderAction
type FlashUrlRenderAction string

// List of FlashUrlRenderAction
const (
	FLASHURLRENDERACTION_CLIENT FlashUrlRenderAction = "Client"
	FLASHURLRENDERACTION_SERVER FlashUrlRenderAction = "Server"
	FLASHURLRENDERACTION_BLOCK  FlashUrlRenderAction = "Block"
)

// All allowed values of FlashUrlRenderAction enum
var AllowedFlashUrlRenderActionEnumValues = []FlashUrlRenderAction{
	"Client",
	"Server",
	"Block",
}

func (v *FlashUrlRenderAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = FlashUrlRenderAction(value)
	return nil
}

// NewFlashUrlRenderActionFromValue returns a pointer to a valid FlashUrlRenderAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFlashUrlRenderActionFromValue(v string) (*FlashUrlRenderAction, error) {
	ev := FlashUrlRenderAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FlashUrlRenderAction: valid values are %v", v, AllowedFlashUrlRenderActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FlashUrlRenderAction) IsValid() bool {
	for _, existing := range AllowedFlashUrlRenderActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlashUrlRenderAction value
func (v FlashUrlRenderAction) Ptr() *FlashUrlRenderAction {
	return &v
}

type NullableFlashUrlRenderAction struct {
	value *FlashUrlRenderAction
	isSet bool
}

func (v NullableFlashUrlRenderAction) Get() *FlashUrlRenderAction {
	return v.value
}

func (v *NullableFlashUrlRenderAction) Set(val *FlashUrlRenderAction) {
	v.value = val
	v.isSet = true
}

func (v NullableFlashUrlRenderAction) IsSet() bool {
	return v.isSet
}

func (v *NullableFlashUrlRenderAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlashUrlRenderAction(val *FlashUrlRenderAction) *NullableFlashUrlRenderAction {
	return &NullableFlashUrlRenderAction{value: val, isSet: true}
}

func (v NullableFlashUrlRenderAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlashUrlRenderAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
