/*
WEM Public API Guide

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package devicemanagement

import (
	"encoding/json"
)

// checks if the UserModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserModel{}

// UserModel Model of user-level AD object
type UserModel struct {
	// Identity of user-level AD object
	Id *int64 `json:"id,omitempty"`
	// SID of user-level AD object
	Sid string `json:"sid"`
	// Distingushed name of user-level AD object
	Name string `json:"name"`
	// Type of user-level AD object
	Type string `json:"type"`
	// Description of user-level AD object
	Description *string `json:"description,omitempty"`
	// Identity of site to which this user-level AD object belongs
	SiteId int64 `json:"siteId"`
	// If this user-level AD objectis enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Priority of user-level AD object
	Priority *int64 `json:"priority,omitempty"`
}

type _UserModel UserModel

// NewUserModel instantiates a new UserModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserModel(sid string, name string, type_ string, siteId int64) *UserModel {
	this := UserModel{}
	this.Sid = sid
	this.Name = name
	this.Type = type_
	this.SiteId = siteId
	return &this
}

// NewUserModelWithDefaults instantiates a new UserModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserModelWithDefaults() *UserModel {
	this := UserModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserModel) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModel) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UserModel) SetId(v int64) {
	o.Id = &v
}

// GetSid returns the Sid field value
func (o *UserModel) GetSid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sid
}

// GetSidOk returns a tuple with the Sid field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sid, true
}

// SetSid sets field value
func (o *UserModel) SetSid(v string) {
	o.Sid = v
}

// GetName returns the Name field value
func (o *UserModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserModel) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *UserModel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UserModel) SetType(v string) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UserModel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UserModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UserModel) SetDescription(v string) {
	o.Description = &v
}

// GetSiteId returns the SiteId field value
func (o *UserModel) GetSiteId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetSiteIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *UserModel) SetSiteId(v int64) {
	o.SiteId = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UserModel) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModel) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UserModel) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UserModel) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *UserModel) GetPriority() int64 {
	if o == nil || IsNil(o.Priority) {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModel) GetPriorityOk() (*int64, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *UserModel) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *UserModel) SetPriority(v int64) {
	o.Priority = &v
}

func (o UserModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["sid"] = o.Sid
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["siteId"] = o.SiteId
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	return toSerialize, nil
}

type NullableUserModel struct {
	value *UserModel
	isSet bool
}

func (v NullableUserModel) Get() *UserModel {
	return v.value
}

func (v *NullableUserModel) Set(val *UserModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUserModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUserModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserModel(val *UserModel) *NullableUserModel {
	return &NullableUserModel{value: val, isSet: true}
}

func (v NullableUserModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
