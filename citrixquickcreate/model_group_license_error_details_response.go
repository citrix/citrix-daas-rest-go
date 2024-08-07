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

// checks if the GroupLicenseErrorDetailsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupLicenseErrorDetailsResponse{}

// GroupLicenseErrorDetailsResponse struct for GroupLicenseErrorDetailsResponse
type GroupLicenseErrorDetailsResponse struct {
	LicenseStatus []ResponseDetailsGroupLicenseAssignment `json:"licenseStatus,omitempty"`
}

// NewGroupLicenseErrorDetailsResponse instantiates a new GroupLicenseErrorDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupLicenseErrorDetailsResponse() *GroupLicenseErrorDetailsResponse {
	this := GroupLicenseErrorDetailsResponse{}
	return &this
}

// NewGroupLicenseErrorDetailsResponseWithDefaults instantiates a new GroupLicenseErrorDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupLicenseErrorDetailsResponseWithDefaults() *GroupLicenseErrorDetailsResponse {
	this := GroupLicenseErrorDetailsResponse{}
	return &this
}

// GetLicenseStatus returns the LicenseStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupLicenseErrorDetailsResponse) GetLicenseStatus() []ResponseDetailsGroupLicenseAssignment {
	if o == nil {
		var ret []ResponseDetailsGroupLicenseAssignment
		return ret
	}
	return o.LicenseStatus
}

// GetLicenseStatusOk returns a tuple with the LicenseStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupLicenseErrorDetailsResponse) GetLicenseStatusOk() ([]ResponseDetailsGroupLicenseAssignment, bool) {
	if o == nil || IsNil(o.LicenseStatus) {
		return nil, false
	}
	return o.LicenseStatus, true
}

// HasLicenseStatus returns a boolean if a field has been set.
func (o *GroupLicenseErrorDetailsResponse) HasLicenseStatus() bool {
	if o != nil && IsNil(o.LicenseStatus) {
		return true
	}

	return false
}

// SetLicenseStatus gets a reference to the given []ResponseDetailsGroupLicenseAssignment and assigns it to the LicenseStatus field.
func (o *GroupLicenseErrorDetailsResponse) SetLicenseStatus(v []ResponseDetailsGroupLicenseAssignment) {
	o.LicenseStatus = v
}

func (o GroupLicenseErrorDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupLicenseErrorDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.LicenseStatus != nil {
		toSerialize["licenseStatus"] = o.LicenseStatus
	}
	return toSerialize, nil
}

type NullableGroupLicenseErrorDetailsResponse struct {
	value *GroupLicenseErrorDetailsResponse
	isSet bool
}

func (v NullableGroupLicenseErrorDetailsResponse) Get() *GroupLicenseErrorDetailsResponse {
	return v.value
}

func (v *NullableGroupLicenseErrorDetailsResponse) Set(val *GroupLicenseErrorDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupLicenseErrorDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupLicenseErrorDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupLicenseErrorDetailsResponse(val *GroupLicenseErrorDetailsResponse) *NullableGroupLicenseErrorDetailsResponse {
	return &NullableGroupLicenseErrorDetailsResponse{value: val, isSet: true}
}

func (v NullableGroupLicenseErrorDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupLicenseErrorDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

