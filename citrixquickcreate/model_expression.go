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

// checks if the Expression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Expression{}

// Expression struct for Expression
type Expression struct {
	AwsResourceTagCitrixManaged NullableString `json:"aws:ResourceTag/CitrixManaged,omitempty"`
	IamPassedToService NullableString `json:"iam:PassedToService,omitempty"`
}

// NewExpression instantiates a new Expression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpression() *Expression {
	this := Expression{}
	return &this
}

// NewExpressionWithDefaults instantiates a new Expression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpressionWithDefaults() *Expression {
	this := Expression{}
	return &this
}

// GetAwsResourceTagCitrixManaged returns the AwsResourceTagCitrixManaged field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Expression) GetAwsResourceTagCitrixManaged() string {
	if o == nil || IsNil(o.AwsResourceTagCitrixManaged.Get()) {
		var ret string
		return ret
	}
	return *o.AwsResourceTagCitrixManaged.Get()
}

// GetAwsResourceTagCitrixManagedOk returns a tuple with the AwsResourceTagCitrixManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Expression) GetAwsResourceTagCitrixManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsResourceTagCitrixManaged.Get(), o.AwsResourceTagCitrixManaged.IsSet()
}

// HasAwsResourceTagCitrixManaged returns a boolean if a field has been set.
func (o *Expression) HasAwsResourceTagCitrixManaged() bool {
	if o != nil && o.AwsResourceTagCitrixManaged.IsSet() {
		return true
	}

	return false
}

// SetAwsResourceTagCitrixManaged gets a reference to the given NullableString and assigns it to the AwsResourceTagCitrixManaged field.
func (o *Expression) SetAwsResourceTagCitrixManaged(v string) {
	o.AwsResourceTagCitrixManaged.Set(&v)
}
// SetAwsResourceTagCitrixManagedNil sets the value for AwsResourceTagCitrixManaged to be an explicit nil
func (o *Expression) SetAwsResourceTagCitrixManagedNil() {
	o.AwsResourceTagCitrixManaged.Set(nil)
}

// UnsetAwsResourceTagCitrixManaged ensures that no value is present for AwsResourceTagCitrixManaged, not even an explicit nil
func (o *Expression) UnsetAwsResourceTagCitrixManaged() {
	o.AwsResourceTagCitrixManaged.Unset()
}

// GetIamPassedToService returns the IamPassedToService field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Expression) GetIamPassedToService() string {
	if o == nil || IsNil(o.IamPassedToService.Get()) {
		var ret string
		return ret
	}
	return *o.IamPassedToService.Get()
}

// GetIamPassedToServiceOk returns a tuple with the IamPassedToService field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Expression) GetIamPassedToServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IamPassedToService.Get(), o.IamPassedToService.IsSet()
}

// HasIamPassedToService returns a boolean if a field has been set.
func (o *Expression) HasIamPassedToService() bool {
	if o != nil && o.IamPassedToService.IsSet() {
		return true
	}

	return false
}

// SetIamPassedToService gets a reference to the given NullableString and assigns it to the IamPassedToService field.
func (o *Expression) SetIamPassedToService(v string) {
	o.IamPassedToService.Set(&v)
}
// SetIamPassedToServiceNil sets the value for IamPassedToService to be an explicit nil
func (o *Expression) SetIamPassedToServiceNil() {
	o.IamPassedToService.Set(nil)
}

// UnsetIamPassedToService ensures that no value is present for IamPassedToService, not even an explicit nil
func (o *Expression) UnsetIamPassedToService() {
	o.IamPassedToService.Unset()
}

func (o Expression) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Expression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AwsResourceTagCitrixManaged.IsSet() {
		toSerialize["aws:ResourceTag/CitrixManaged"] = o.AwsResourceTagCitrixManaged.Get()
	}
	if o.IamPassedToService.IsSet() {
		toSerialize["iam:PassedToService"] = o.IamPassedToService.Get()
	}
	return toSerialize, nil
}

type NullableExpression struct {
	value *Expression
	isSet bool
}

func (v NullableExpression) Get() *Expression {
	return v.value
}

func (v *NullableExpression) Set(val *Expression) {
	v.value = val
	v.isSet = true
}

func (v NullableExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpression(val *Expression) *NullableExpression {
	return &NullableExpression{value: val, isSet: true}
}

func (v NullableExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

