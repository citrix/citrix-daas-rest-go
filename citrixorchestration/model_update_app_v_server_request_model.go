/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: devportal-onprem-ga
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixorchestration

import (
	"encoding/json"
)

// checks if the UpdateAppVServerRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAppVServerRequestModel{}

// UpdateAppVServerRequestModel Request model to update AppV server.
type UpdateAppVServerRequestModel struct {
	// The url of the App-V Management server that packages will be discovered from.
	ManagementServer NullableString `json:"ManagementServer,omitempty"`
	// The url of the App-V Publishing server that packages will be discovered from.
	PublishingServer NullableString `json:"PublishingServer,omitempty"`
}

// NewUpdateAppVServerRequestModel instantiates a new UpdateAppVServerRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAppVServerRequestModel() *UpdateAppVServerRequestModel {
	this := UpdateAppVServerRequestModel{}
	return &this
}

// NewUpdateAppVServerRequestModelWithDefaults instantiates a new UpdateAppVServerRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAppVServerRequestModelWithDefaults() *UpdateAppVServerRequestModel {
	this := UpdateAppVServerRequestModel{}
	return &this
}

// GetManagementServer returns the ManagementServer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAppVServerRequestModel) GetManagementServer() string {
	if o == nil || IsNil(o.ManagementServer.Get()) {
		var ret string
		return ret
	}
	return *o.ManagementServer.Get()
}

// GetManagementServerOk returns a tuple with the ManagementServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAppVServerRequestModel) GetManagementServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManagementServer.Get(), o.ManagementServer.IsSet()
}

// HasManagementServer returns a boolean if a field has been set.
func (o *UpdateAppVServerRequestModel) HasManagementServer() bool {
	if o != nil && o.ManagementServer.IsSet() {
		return true
	}

	return false
}

// SetManagementServer gets a reference to the given NullableString and assigns it to the ManagementServer field.
func (o *UpdateAppVServerRequestModel) SetManagementServer(v string) {
	o.ManagementServer.Set(&v)
}

// SetManagementServerNil sets the value for ManagementServer to be an explicit nil
func (o *UpdateAppVServerRequestModel) SetManagementServerNil() {
	o.ManagementServer.Set(nil)
}

// UnsetManagementServer ensures that no value is present for ManagementServer, not even an explicit nil
func (o *UpdateAppVServerRequestModel) UnsetManagementServer() {
	o.ManagementServer.Unset()
}

// GetPublishingServer returns the PublishingServer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAppVServerRequestModel) GetPublishingServer() string {
	if o == nil || IsNil(o.PublishingServer.Get()) {
		var ret string
		return ret
	}
	return *o.PublishingServer.Get()
}

// GetPublishingServerOk returns a tuple with the PublishingServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAppVServerRequestModel) GetPublishingServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublishingServer.Get(), o.PublishingServer.IsSet()
}

// HasPublishingServer returns a boolean if a field has been set.
func (o *UpdateAppVServerRequestModel) HasPublishingServer() bool {
	if o != nil && o.PublishingServer.IsSet() {
		return true
	}

	return false
}

// SetPublishingServer gets a reference to the given NullableString and assigns it to the PublishingServer field.
func (o *UpdateAppVServerRequestModel) SetPublishingServer(v string) {
	o.PublishingServer.Set(&v)
}

// SetPublishingServerNil sets the value for PublishingServer to be an explicit nil
func (o *UpdateAppVServerRequestModel) SetPublishingServerNil() {
	o.PublishingServer.Set(nil)
}

// UnsetPublishingServer ensures that no value is present for PublishingServer, not even an explicit nil
func (o *UpdateAppVServerRequestModel) UnsetPublishingServer() {
	o.PublishingServer.Unset()
}

func (o UpdateAppVServerRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAppVServerRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ManagementServer.IsSet() {
		toSerialize["ManagementServer"] = o.ManagementServer.Get()
	}
	if o.PublishingServer.IsSet() {
		toSerialize["PublishingServer"] = o.PublishingServer.Get()
	}
	return toSerialize, nil
}

type NullableUpdateAppVServerRequestModel struct {
	value *UpdateAppVServerRequestModel
	isSet bool
}

func (v NullableUpdateAppVServerRequestModel) Get() *UpdateAppVServerRequestModel {
	return v.value
}

func (v *NullableUpdateAppVServerRequestModel) Set(val *UpdateAppVServerRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAppVServerRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAppVServerRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAppVServerRequestModel(val *UpdateAppVServerRequestModel) *NullableUpdateAppVServerRequestModel {
	return &NullableUpdateAppVServerRequestModel{value: val, isSet: true}
}

func (v NullableUpdateAppVServerRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAppVServerRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
