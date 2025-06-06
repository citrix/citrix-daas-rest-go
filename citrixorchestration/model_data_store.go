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

// DataStore Types of DataStore.
type DataStore string

// List of DataStore
const (
	DATASTORE_UNKNOWN DataStore = "Unknown"
	DATASTORE_SITE    DataStore = "Site"
	DATASTORE_LOGGING DataStore = "Logging"
	DATASTORE_MONITOR DataStore = "Monitor"
)

// All allowed values of DataStore enum
var AllowedDataStoreEnumValues = []DataStore{
	"Unknown",
	"Site",
	"Logging",
	"Monitor",
}

func (v *DataStore) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	*v = DataStore(value)
	return nil
}

// NewDataStoreFromValue returns a pointer to a valid DataStore
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataStoreFromValue(v string) (*DataStore, error) {
	ev := DataStore(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataStore: valid values are %v", v, AllowedDataStoreEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataStore) IsValid() bool {
	for _, existing := range AllowedDataStoreEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataStore value
func (v DataStore) Ptr() *DataStore {
	return &v
}

type NullableDataStore struct {
	value *DataStore
	isSet bool
}

func (v NullableDataStore) Get() *DataStore {
	return v.value
}

func (v *NullableDataStore) Set(val *DataStore) {
	v.value = val
	v.isSet = true
}

func (v NullableDataStore) IsSet() bool {
	return v.isSet
}

func (v *NullableDataStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataStore(val *DataStore) *NullableDataStore {
	return &NullableDataStore{value: val, isSet: true}
}

func (v NullableDataStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
