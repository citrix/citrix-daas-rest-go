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

// checks if the GroupUsers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupUsers{}

// GroupUsers 
type GroupUsers struct {
	// 
	Users []CoreUser `json:"users,omitempty"`
}

// NewGroupUsers instantiates a new GroupUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupUsers() *GroupUsers {
	this := GroupUsers{}
	return &this
}

// NewGroupUsersWithDefaults instantiates a new GroupUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupUsersWithDefaults() *GroupUsers {
	this := GroupUsers{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupUsers) GetUsers() []CoreUser {
	if o == nil {
		var ret []CoreUser
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupUsers) GetUsersOk() ([]CoreUser, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *GroupUsers) HasUsers() bool {
	if o != nil && IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []CoreUser and assigns it to the Users field.
func (o *GroupUsers) SetUsers(v []CoreUser) {
	o.Users = v
}

func (o GroupUsers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupUsers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableGroupUsers struct {
	value *GroupUsers
	isSet bool
}

func (v NullableGroupUsers) Get() *GroupUsers {
	return v.value
}

func (v *NullableGroupUsers) Set(val *GroupUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupUsers(val *GroupUsers) *NullableGroupUsers {
	return &NullableGroupUsers{value: val, isSet: true}
}

func (v NullableGroupUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

