/*
Quick Create Service for DaaS Quick Deploy with AWS WorkSpaces Core - Tech Preview

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: DaaSTechPreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickcreate

import (
	"encoding/json"
)

// checks if the AwsEdcAccountResourceKeyPair type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsEdcAccountResourceKeyPair{}

// AwsEdcAccountResourceKeyPair struct for AwsEdcAccountResourceKeyPair
type AwsEdcAccountResourceKeyPair struct {
	AwsEdcAccountResource
	// Name of Key Pari
	KeyName NullableString `json:"keyName,omitempty"`
	// Vpc Tags
	Tags []Tag `json:"tags,omitempty"`
}

// NewAwsEdcAccountResourceKeyPair instantiates a new AwsEdcAccountResourceKeyPair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsEdcAccountResourceKeyPair(accountType AccountType) *AwsEdcAccountResourceKeyPair {
	this := AwsEdcAccountResourceKeyPair{}
	this.AccountType = accountType
	return &this
}

// NewAwsEdcAccountResourceKeyPairWithDefaults instantiates a new AwsEdcAccountResourceKeyPair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsEdcAccountResourceKeyPairWithDefaults() *AwsEdcAccountResourceKeyPair {
	this := AwsEdcAccountResourceKeyPair{}
	return &this
}

// GetKeyName returns the KeyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceKeyPair) GetKeyName() string {
	if o == nil || IsNil(o.KeyName.Get()) {
		var ret string
		return ret
	}
	return *o.KeyName.Get()
}

// GetKeyNameOk returns a tuple with the KeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceKeyPair) GetKeyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeyName.Get(), o.KeyName.IsSet()
}

// HasKeyName returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceKeyPair) HasKeyName() bool {
	if o != nil && o.KeyName.IsSet() {
		return true
	}

	return false
}

// SetKeyName gets a reference to the given NullableString and assigns it to the KeyName field.
func (o *AwsEdcAccountResourceKeyPair) SetKeyName(v string) {
	o.KeyName.Set(&v)
}
// SetKeyNameNil sets the value for KeyName to be an explicit nil
func (o *AwsEdcAccountResourceKeyPair) SetKeyNameNil() {
	o.KeyName.Set(nil)
}

// UnsetKeyName ensures that no value is present for KeyName, not even an explicit nil
func (o *AwsEdcAccountResourceKeyPair) UnsetKeyName() {
	o.KeyName.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceKeyPair) GetTags() []Tag {
	if o == nil {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceKeyPair) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceKeyPair) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *AwsEdcAccountResourceKeyPair) SetTags(v []Tag) {
	o.Tags = v
}

func (o AwsEdcAccountResourceKeyPair) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsEdcAccountResourceKeyPair) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAwsEdcAccountResource, errAwsEdcAccountResource := json.Marshal(o.AwsEdcAccountResource)
	if errAwsEdcAccountResource != nil {
		return map[string]interface{}{}, errAwsEdcAccountResource
	}
	errAwsEdcAccountResource = json.Unmarshal([]byte(serializedAwsEdcAccountResource), &toSerialize)
	if errAwsEdcAccountResource != nil {
		return map[string]interface{}{}, errAwsEdcAccountResource
	}
	if o.KeyName.IsSet() {
		toSerialize["keyName"] = o.KeyName.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableAwsEdcAccountResourceKeyPair struct {
	value *AwsEdcAccountResourceKeyPair
	isSet bool
}

func (v NullableAwsEdcAccountResourceKeyPair) Get() *AwsEdcAccountResourceKeyPair {
	return v.value
}

func (v *NullableAwsEdcAccountResourceKeyPair) Set(val *AwsEdcAccountResourceKeyPair) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsEdcAccountResourceKeyPair) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsEdcAccountResourceKeyPair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsEdcAccountResourceKeyPair(val *AwsEdcAccountResourceKeyPair) *NullableAwsEdcAccountResourceKeyPair {
	return &NullableAwsEdcAccountResourceKeyPair{value: val, isSet: true}
}

func (v NullableAwsEdcAccountResourceKeyPair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsEdcAccountResourceKeyPair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

