package models

import (
	"encoding/json"
	"fmt"
)

type LogonType string

// LogonType enum values

const (
	LOGONTYPE_USED_FOR_HDX_ONLY LogonType = "UsedForHDXOnly"
	LOGONTYPE_DOMAIN            LogonType = "Domain"
	LOGONTYPE_RSA               LogonType = "RSA"
	LOGONTYPE_DOMAIN_AND_RSA    LogonType = "DomainAndRSA"
	LOGONTYPE_SMS               LogonType = "SMS"
	LOGONTYPE_GATEWAY_KNOWS     LogonType = "GatewayKnows"
	LOGONTYPE_SMART_CARD        LogonType = "SmartCard"
	LOGONTYPE_NONE              LogonType = "None"
)

// AllowedLogonTypeEnumValues is the list of allowed values for the enum

var AllowedLogonTypeEnumValues = []LogonType{
	"UsedForHDXOnly",
	"Domain",
	"RSA",
	"DomainAndRSA",
	"SMS",
	"GatewayKnows",
	"SmartCard",
	"None",
}

func (v *LogonType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LogonType(value)
	for _, existing := range AllowedLogonTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LogonType", value)
}

// NewLogonTypeFromValue returns a pointer to a valid LogonType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogonTypeFromValue(v string) (*LogonType, error) {
	ev := LogonType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogonType: valid values are %v", v, AllowedLogonTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogonType) IsValid() bool {
	for _, existing := range AllowedLogonTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogonType value
func (v LogonType) Ptr() *LogonType {
	return &v
}

type NullableLogonType struct {
	value *LogonType
	isSet bool
}

func (v NullableLogonType) Get() *LogonType {
	return v.value
}

func (v *NullableLogonType) Set(val *LogonType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogonType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogonType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogonType(val *LogonType) *NullableLogonType {
	return &NullableLogonType{value: val, isSet: true}
}

func (v NullableLogonType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogonType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
