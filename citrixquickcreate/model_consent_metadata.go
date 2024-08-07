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

// checks if the ConsentMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsentMetadata{}

// ConsentMetadata struct for ConsentMetadata
type ConsentMetadata struct {
	// Client/Application for MT AppRegistration for CloudPC Registration service  This is the app that we would get consent for
	ClientId NullableString `json:"clientId,omitempty"`
	// Scope for consent request  .default would get consent for all permissions listed in AppRegistration APIPermissions section
	Scope NullableString `json:"scope,omitempty"`
}

// NewConsentMetadata instantiates a new ConsentMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsentMetadata() *ConsentMetadata {
	this := ConsentMetadata{}
	return &this
}

// NewConsentMetadataWithDefaults instantiates a new ConsentMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsentMetadataWithDefaults() *ConsentMetadata {
	this := ConsentMetadata{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsentMetadata) GetClientId() string {
	if o == nil || IsNil(o.ClientId.Get()) {
		var ret string
		return ret
	}
	return *o.ClientId.Get()
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsentMetadata) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientId.Get(), o.ClientId.IsSet()
}

// HasClientId returns a boolean if a field has been set.
func (o *ConsentMetadata) HasClientId() bool {
	if o != nil && o.ClientId.IsSet() {
		return true
	}

	return false
}

// SetClientId gets a reference to the given NullableString and assigns it to the ClientId field.
func (o *ConsentMetadata) SetClientId(v string) {
	o.ClientId.Set(&v)
}
// SetClientIdNil sets the value for ClientId to be an explicit nil
func (o *ConsentMetadata) SetClientIdNil() {
	o.ClientId.Set(nil)
}

// UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
func (o *ConsentMetadata) UnsetClientId() {
	o.ClientId.Unset()
}

// GetScope returns the Scope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsentMetadata) GetScope() string {
	if o == nil || IsNil(o.Scope.Get()) {
		var ret string
		return ret
	}
	return *o.Scope.Get()
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsentMetadata) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scope.Get(), o.Scope.IsSet()
}

// HasScope returns a boolean if a field has been set.
func (o *ConsentMetadata) HasScope() bool {
	if o != nil && o.Scope.IsSet() {
		return true
	}

	return false
}

// SetScope gets a reference to the given NullableString and assigns it to the Scope field.
func (o *ConsentMetadata) SetScope(v string) {
	o.Scope.Set(&v)
}
// SetScopeNil sets the value for Scope to be an explicit nil
func (o *ConsentMetadata) SetScopeNil() {
	o.Scope.Set(nil)
}

// UnsetScope ensures that no value is present for Scope, not even an explicit nil
func (o *ConsentMetadata) UnsetScope() {
	o.Scope.Unset()
}

func (o ConsentMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsentMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId.IsSet() {
		toSerialize["clientId"] = o.ClientId.Get()
	}
	if o.Scope.IsSet() {
		toSerialize["scope"] = o.Scope.Get()
	}
	return toSerialize, nil
}

type NullableConsentMetadata struct {
	value *ConsentMetadata
	isSet bool
}

func (v NullableConsentMetadata) Get() *ConsentMetadata {
	return v.value
}

func (v *NullableConsentMetadata) Set(val *ConsentMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableConsentMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableConsentMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsentMetadata(val *ConsentMetadata) *NullableConsentMetadata {
	return &NullableConsentMetadata{value: val, isSet: true}
}

func (v NullableConsentMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsentMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

