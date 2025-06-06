/*
Global App Config Admin

Describes API used by Global App Config Admin Service

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package globalappconfiguration

import (
	"encoding/json"
)

// checks if the ServiceURL type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceURL{}

// ServiceURL struct for ServiceURL
type ServiceURL struct {
	MigrationUrl []MigrationUrl `json:"migrationUrl,omitempty"`
	Url          *string        `json:"url,omitempty"`
}

// NewServiceURL instantiates a new ServiceURL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceURL() *ServiceURL {
	this := ServiceURL{}
	return &this
}

// NewServiceURLWithDefaults instantiates a new ServiceURL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceURLWithDefaults() *ServiceURL {
	this := ServiceURL{}
	return &this
}

// GetMigrationUrl returns the MigrationUrl field value if set, zero value otherwise.
func (o *ServiceURL) GetMigrationUrl() []MigrationUrl {
	if o == nil || IsNil(o.MigrationUrl) {
		var ret []MigrationUrl
		return ret
	}
	return o.MigrationUrl
}

// GetMigrationUrlOk returns a tuple with the MigrationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceURL) GetMigrationUrlOk() ([]MigrationUrl, bool) {
	if o == nil || IsNil(o.MigrationUrl) {
		return nil, false
	}
	return o.MigrationUrl, true
}

// HasMigrationUrl returns a boolean if a field has been set.
func (o *ServiceURL) HasMigrationUrl() bool {
	if o != nil && !IsNil(o.MigrationUrl) {
		return true
	}

	return false
}

// SetMigrationUrl gets a reference to the given []MigrationUrl and assigns it to the MigrationUrl field.
func (o *ServiceURL) SetMigrationUrl(v []MigrationUrl) {
	o.MigrationUrl = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ServiceURL) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceURL) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ServiceURL) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ServiceURL) SetUrl(v string) {
	o.Url = &v
}

func (o ServiceURL) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceURL) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MigrationUrl) {
		toSerialize["migrationUrl"] = o.MigrationUrl
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableServiceURL struct {
	value *ServiceURL
	isSet bool
}

func (v NullableServiceURL) Get() *ServiceURL {
	return v.value
}

func (v *NullableServiceURL) Set(val *ServiceURL) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceURL) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceURL) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceURL(val *ServiceURL) *NullableServiceURL {
	return &NullableServiceURL{value: val, isSet: true}
}

func (v NullableServiceURL) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceURL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
