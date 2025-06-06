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

// checks if the ResponseDetailsGroupLicenseAssignment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseDetailsGroupLicenseAssignment{}

// ResponseDetailsGroupLicenseAssignment struct for ResponseDetailsGroupLicenseAssignment
type ResponseDetailsGroupLicenseAssignment struct {
	Groups []string `json:"groups,omitempty"`
	// Once of the error codes listed in the ErrorCode enum
	ErrorCode *int32 `json:"errorCode,omitempty"`
	// Description about the error
	ErrorDescription NullableString `json:"errorDescription,omitempty"`
}

// NewResponseDetailsGroupLicenseAssignment instantiates a new ResponseDetailsGroupLicenseAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseDetailsGroupLicenseAssignment() *ResponseDetailsGroupLicenseAssignment {
	this := ResponseDetailsGroupLicenseAssignment{}
	return &this
}

// NewResponseDetailsGroupLicenseAssignmentWithDefaults instantiates a new ResponseDetailsGroupLicenseAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseDetailsGroupLicenseAssignmentWithDefaults() *ResponseDetailsGroupLicenseAssignment {
	this := ResponseDetailsGroupLicenseAssignment{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseDetailsGroupLicenseAssignment) GetGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseDetailsGroupLicenseAssignment) GetGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ResponseDetailsGroupLicenseAssignment) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *ResponseDetailsGroupLicenseAssignment) SetGroups(v []string) {
	o.Groups = v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *ResponseDetailsGroupLicenseAssignment) GetErrorCode() int32 {
	if o == nil || IsNil(o.ErrorCode) {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDetailsGroupLicenseAssignment) GetErrorCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ResponseDetailsGroupLicenseAssignment) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *ResponseDetailsGroupLicenseAssignment) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseDetailsGroupLicenseAssignment) GetErrorDescription() string {
	if o == nil || IsNil(o.ErrorDescription.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorDescription.Get()
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseDetailsGroupLicenseAssignment) GetErrorDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorDescription.Get(), o.ErrorDescription.IsSet()
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *ResponseDetailsGroupLicenseAssignment) HasErrorDescription() bool {
	if o != nil && o.ErrorDescription.IsSet() {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given NullableString and assigns it to the ErrorDescription field.
func (o *ResponseDetailsGroupLicenseAssignment) SetErrorDescription(v string) {
	o.ErrorDescription.Set(&v)
}

// SetErrorDescriptionNil sets the value for ErrorDescription to be an explicit nil
func (o *ResponseDetailsGroupLicenseAssignment) SetErrorDescriptionNil() {
	o.ErrorDescription.Set(nil)
}

// UnsetErrorDescription ensures that no value is present for ErrorDescription, not even an explicit nil
func (o *ResponseDetailsGroupLicenseAssignment) UnsetErrorDescription() {
	o.ErrorDescription.Unset()
}

func (o ResponseDetailsGroupLicenseAssignment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseDetailsGroupLicenseAssignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.ErrorDescription.IsSet() {
		toSerialize["errorDescription"] = o.ErrorDescription.Get()
	}
	return toSerialize, nil
}

type NullableResponseDetailsGroupLicenseAssignment struct {
	value *ResponseDetailsGroupLicenseAssignment
	isSet bool
}

func (v NullableResponseDetailsGroupLicenseAssignment) Get() *ResponseDetailsGroupLicenseAssignment {
	return v.value
}

func (v *NullableResponseDetailsGroupLicenseAssignment) Set(val *ResponseDetailsGroupLicenseAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseDetailsGroupLicenseAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseDetailsGroupLicenseAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseDetailsGroupLicenseAssignment(val *ResponseDetailsGroupLicenseAssignment) *NullableResponseDetailsGroupLicenseAssignment {
	return &NullableResponseDetailsGroupLicenseAssignment{value: val, isSet: true}
}

func (v NullableResponseDetailsGroupLicenseAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseDetailsGroupLicenseAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
