/*
Citrix.CloudServices.Cws.Api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixcws

import (
	"encoding/json"
)

// checks if the GoogleConnectionModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoogleConnectionModel{}

// GoogleConnectionModel struct for GoogleConnectionModel
type GoogleConnectionModel struct {
	GoogleClientEmail string `json:"googleClientEmail"`
	GooglePrivateKey string `json:"googlePrivateKey"`
	GoogleImpersonatedUser string `json:"googleImpersonatedUser"`
}

// NewGoogleConnectionModel instantiates a new GoogleConnectionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleConnectionModel(googleClientEmail string, googlePrivateKey string, googleImpersonatedUser string) *GoogleConnectionModel {
	this := GoogleConnectionModel{}
	this.GoogleClientEmail = googleClientEmail
	this.GooglePrivateKey = googlePrivateKey
	this.GoogleImpersonatedUser = googleImpersonatedUser
	return &this
}

// NewGoogleConnectionModelWithDefaults instantiates a new GoogleConnectionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleConnectionModelWithDefaults() *GoogleConnectionModel {
	this := GoogleConnectionModel{}
	return &this
}

// GetGoogleClientEmail returns the GoogleClientEmail field value
func (o *GoogleConnectionModel) GetGoogleClientEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GoogleClientEmail
}

// GetGoogleClientEmailOk returns a tuple with the GoogleClientEmail field value
// and a boolean to check if the value has been set.
func (o *GoogleConnectionModel) GetGoogleClientEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GoogleClientEmail, true
}

// SetGoogleClientEmail sets field value
func (o *GoogleConnectionModel) SetGoogleClientEmail(v string) {
	o.GoogleClientEmail = v
}

// GetGooglePrivateKey returns the GooglePrivateKey field value
func (o *GoogleConnectionModel) GetGooglePrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GooglePrivateKey
}

// GetGooglePrivateKeyOk returns a tuple with the GooglePrivateKey field value
// and a boolean to check if the value has been set.
func (o *GoogleConnectionModel) GetGooglePrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GooglePrivateKey, true
}

// SetGooglePrivateKey sets field value
func (o *GoogleConnectionModel) SetGooglePrivateKey(v string) {
	o.GooglePrivateKey = v
}

// GetGoogleImpersonatedUser returns the GoogleImpersonatedUser field value
func (o *GoogleConnectionModel) GetGoogleImpersonatedUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GoogleImpersonatedUser
}

// GetGoogleImpersonatedUserOk returns a tuple with the GoogleImpersonatedUser field value
// and a boolean to check if the value has been set.
func (o *GoogleConnectionModel) GetGoogleImpersonatedUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GoogleImpersonatedUser, true
}

// SetGoogleImpersonatedUser sets field value
func (o *GoogleConnectionModel) SetGoogleImpersonatedUser(v string) {
	o.GoogleImpersonatedUser = v
}

func (o GoogleConnectionModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoogleConnectionModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["googleClientEmail"] = o.GoogleClientEmail
	toSerialize["googlePrivateKey"] = o.GooglePrivateKey
	toSerialize["googleImpersonatedUser"] = o.GoogleImpersonatedUser
	return toSerialize, nil
}

type NullableGoogleConnectionModel struct {
	value *GoogleConnectionModel
	isSet bool
}

func (v NullableGoogleConnectionModel) Get() *GoogleConnectionModel {
	return v.value
}

func (v *NullableGoogleConnectionModel) Set(val *GoogleConnectionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleConnectionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleConnectionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleConnectionModel(val *GoogleConnectionModel) *NullableGoogleConnectionModel {
	return &NullableGoogleConnectionModel{value: val, isSet: true}
}

func (v NullableGoogleConnectionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleConnectionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

