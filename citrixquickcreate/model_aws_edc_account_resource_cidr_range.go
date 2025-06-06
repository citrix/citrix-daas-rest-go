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

// checks if the AwsEdcAccountResourceCidrRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsEdcAccountResourceCidrRange{}

// AwsEdcAccountResourceCidrRange struct for AwsEdcAccountResourceCidrRange
type AwsEdcAccountResourceCidrRange struct {
	AwsEdcAccountResource
	// Base Cidr Range
	BaseCidrRange NullableString `json:"baseCidrRange,omitempty"`
	// Derived Cidr Ranges
	DerivedCidrRanges []string `json:"derivedCidrRanges,omitempty"`
}

type _AwsEdcAccountResourceCidrRange AwsEdcAccountResourceCidrRange

// NewAwsEdcAccountResourceCidrRange instantiates a new AwsEdcAccountResourceCidrRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsEdcAccountResourceCidrRange(resourceType AwsAccountResourceType, accountType AccountType) *AwsEdcAccountResourceCidrRange {
	this := AwsEdcAccountResourceCidrRange{}
	this.ResourceType = resourceType
	this.AccountType = accountType
	return &this
}

// NewAwsEdcAccountResourceCidrRangeWithDefaults instantiates a new AwsEdcAccountResourceCidrRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsEdcAccountResourceCidrRangeWithDefaults() *AwsEdcAccountResourceCidrRange {
	this := AwsEdcAccountResourceCidrRange{}
	return &this
}

// GetBaseCidrRange returns the BaseCidrRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceCidrRange) GetBaseCidrRange() string {
	if o == nil || IsNil(o.BaseCidrRange.Get()) {
		var ret string
		return ret
	}
	return *o.BaseCidrRange.Get()
}

// GetBaseCidrRangeOk returns a tuple with the BaseCidrRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceCidrRange) GetBaseCidrRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BaseCidrRange.Get(), o.BaseCidrRange.IsSet()
}

// HasBaseCidrRange returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceCidrRange) HasBaseCidrRange() bool {
	if o != nil && o.BaseCidrRange.IsSet() {
		return true
	}

	return false
}

// SetBaseCidrRange gets a reference to the given NullableString and assigns it to the BaseCidrRange field.
func (o *AwsEdcAccountResourceCidrRange) SetBaseCidrRange(v string) {
	o.BaseCidrRange.Set(&v)
}

// SetBaseCidrRangeNil sets the value for BaseCidrRange to be an explicit nil
func (o *AwsEdcAccountResourceCidrRange) SetBaseCidrRangeNil() {
	o.BaseCidrRange.Set(nil)
}

// UnsetBaseCidrRange ensures that no value is present for BaseCidrRange, not even an explicit nil
func (o *AwsEdcAccountResourceCidrRange) UnsetBaseCidrRange() {
	o.BaseCidrRange.Unset()
}

// GetDerivedCidrRanges returns the DerivedCidrRanges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsEdcAccountResourceCidrRange) GetDerivedCidrRanges() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DerivedCidrRanges
}

// GetDerivedCidrRangesOk returns a tuple with the DerivedCidrRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsEdcAccountResourceCidrRange) GetDerivedCidrRangesOk() ([]string, bool) {
	if o == nil || IsNil(o.DerivedCidrRanges) {
		return nil, false
	}
	return o.DerivedCidrRanges, true
}

// HasDerivedCidrRanges returns a boolean if a field has been set.
func (o *AwsEdcAccountResourceCidrRange) HasDerivedCidrRanges() bool {
	if o != nil && !IsNil(o.DerivedCidrRanges) {
		return true
	}

	return false
}

// SetDerivedCidrRanges gets a reference to the given []string and assigns it to the DerivedCidrRanges field.
func (o *AwsEdcAccountResourceCidrRange) SetDerivedCidrRanges(v []string) {
	o.DerivedCidrRanges = v
}

func (o AwsEdcAccountResourceCidrRange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsEdcAccountResourceCidrRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAwsEdcAccountResource, errAwsEdcAccountResource := json.Marshal(o.AwsEdcAccountResource)
	if errAwsEdcAccountResource != nil {
		return map[string]interface{}{}, errAwsEdcAccountResource
	}
	errAwsEdcAccountResource = json.Unmarshal([]byte(serializedAwsEdcAccountResource), &toSerialize)
	if errAwsEdcAccountResource != nil {
		return map[string]interface{}{}, errAwsEdcAccountResource
	}
	if o.BaseCidrRange.IsSet() {
		toSerialize["baseCidrRange"] = o.BaseCidrRange.Get()
	}
	if o.DerivedCidrRanges != nil {
		toSerialize["derivedCidrRanges"] = o.DerivedCidrRanges
	}
	return toSerialize, nil
}

type NullableAwsEdcAccountResourceCidrRange struct {
	value *AwsEdcAccountResourceCidrRange
	isSet bool
}

func (v NullableAwsEdcAccountResourceCidrRange) Get() *AwsEdcAccountResourceCidrRange {
	return v.value
}

func (v *NullableAwsEdcAccountResourceCidrRange) Set(val *AwsEdcAccountResourceCidrRange) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsEdcAccountResourceCidrRange) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsEdcAccountResourceCidrRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsEdcAccountResourceCidrRange(val *AwsEdcAccountResourceCidrRange) *NullableAwsEdcAccountResourceCidrRange {
	return &NullableAwsEdcAccountResourceCidrRange{value: val, isSet: true}
}

func (v NullableAwsEdcAccountResourceCidrRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsEdcAccountResourceCidrRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
