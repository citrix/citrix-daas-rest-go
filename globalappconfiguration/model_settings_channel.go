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

// checks if the SettingsChannel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingsChannel{}

// SettingsChannel struct for SettingsChannel
type SettingsChannel struct {
	ChannelName *string `json:"channelName,omitempty"`
}

// NewSettingsChannel instantiates a new SettingsChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsChannel() *SettingsChannel {
	this := SettingsChannel{}
	return &this
}

// NewSettingsChannelWithDefaults instantiates a new SettingsChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsChannelWithDefaults() *SettingsChannel {
	this := SettingsChannel{}
	return &this
}

// GetChannelName returns the ChannelName field value if set, zero value otherwise.
func (o *SettingsChannel) GetChannelName() string {
	if o == nil || IsNil(o.ChannelName) {
		var ret string
		return ret
	}
	return *o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsChannel) GetChannelNameOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelName) {
		return nil, false
	}
	return o.ChannelName, true
}

// HasChannelName returns a boolean if a field has been set.
func (o *SettingsChannel) HasChannelName() bool {
	if o != nil && !IsNil(o.ChannelName) {
		return true
	}

	return false
}

// SetChannelName gets a reference to the given string and assigns it to the ChannelName field.
func (o *SettingsChannel) SetChannelName(v string) {
	o.ChannelName = &v
}

func (o SettingsChannel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingsChannel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChannelName) {
		toSerialize["channelName"] = o.ChannelName
	}
	return toSerialize, nil
}

type NullableSettingsChannel struct {
	value *SettingsChannel
	isSet bool
}

func (v NullableSettingsChannel) Get() *SettingsChannel {
	return v.value
}

func (v *NullableSettingsChannel) Set(val *SettingsChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsChannel(val *SettingsChannel) *NullableSettingsChannel {
	return &NullableSettingsChannel{value: val, isSet: true}
}

func (v NullableSettingsChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
