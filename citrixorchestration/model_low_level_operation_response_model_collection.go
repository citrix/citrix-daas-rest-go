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

// checks if the LowLevelOperationResponseModelCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LowLevelOperationResponseModelCollection{}

// LowLevelOperationResponseModelCollection struct for LowLevelOperationResponseModelCollection
type LowLevelOperationResponseModelCollection struct {
	// List of items.
	Items []LowLevelOperationResponseModel `json:"Items"`
	// If present, indicates to the caller that the query was not complete, and they should call the API again specifying the continuation token as a query parameter.
	ContinuationToken NullableString `json:"ContinuationToken,omitempty"`
	// Indicates the total number of items in the collection, which may be more than the number of Items returned, if there is a ContinuationToken.  Only returned in the response to `$search` APIs.
	TotalItems NullableInt32 `json:"TotalItems,omitempty"`
}

// NewLowLevelOperationResponseModelCollection instantiates a new LowLevelOperationResponseModelCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLowLevelOperationResponseModelCollection(items []LowLevelOperationResponseModel) *LowLevelOperationResponseModelCollection {
	this := LowLevelOperationResponseModelCollection{}
	this.Items = items
	return &this
}

// NewLowLevelOperationResponseModelCollectionWithDefaults instantiates a new LowLevelOperationResponseModelCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLowLevelOperationResponseModelCollectionWithDefaults() *LowLevelOperationResponseModelCollection {
	this := LowLevelOperationResponseModelCollection{}
	return &this
}

// GetItems returns the Items field value
func (o *LowLevelOperationResponseModelCollection) GetItems() []LowLevelOperationResponseModel {
	if o == nil {
		var ret []LowLevelOperationResponseModel
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *LowLevelOperationResponseModelCollection) GetItemsOk() ([]LowLevelOperationResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *LowLevelOperationResponseModelCollection) SetItems(v []LowLevelOperationResponseModel) {
	o.Items = v
}

// GetContinuationToken returns the ContinuationToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LowLevelOperationResponseModelCollection) GetContinuationToken() string {
	if o == nil || IsNil(o.ContinuationToken.Get()) {
		var ret string
		return ret
	}
	return *o.ContinuationToken.Get()
}

// GetContinuationTokenOk returns a tuple with the ContinuationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LowLevelOperationResponseModelCollection) GetContinuationTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinuationToken.Get(), o.ContinuationToken.IsSet()
}

// HasContinuationToken returns a boolean if a field has been set.
func (o *LowLevelOperationResponseModelCollection) HasContinuationToken() bool {
	if o != nil && o.ContinuationToken.IsSet() {
		return true
	}

	return false
}

// SetContinuationToken gets a reference to the given NullableString and assigns it to the ContinuationToken field.
func (o *LowLevelOperationResponseModelCollection) SetContinuationToken(v string) {
	o.ContinuationToken.Set(&v)
}

// SetContinuationTokenNil sets the value for ContinuationToken to be an explicit nil
func (o *LowLevelOperationResponseModelCollection) SetContinuationTokenNil() {
	o.ContinuationToken.Set(nil)
}

// UnsetContinuationToken ensures that no value is present for ContinuationToken, not even an explicit nil
func (o *LowLevelOperationResponseModelCollection) UnsetContinuationToken() {
	o.ContinuationToken.Unset()
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LowLevelOperationResponseModelCollection) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalItems.Get()
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LowLevelOperationResponseModelCollection) GetTotalItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalItems.Get(), o.TotalItems.IsSet()
}

// HasTotalItems returns a boolean if a field has been set.
func (o *LowLevelOperationResponseModelCollection) HasTotalItems() bool {
	if o != nil && o.TotalItems.IsSet() {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given NullableInt32 and assigns it to the TotalItems field.
func (o *LowLevelOperationResponseModelCollection) SetTotalItems(v int32) {
	o.TotalItems.Set(&v)
}

// SetTotalItemsNil sets the value for TotalItems to be an explicit nil
func (o *LowLevelOperationResponseModelCollection) SetTotalItemsNil() {
	o.TotalItems.Set(nil)
}

// UnsetTotalItems ensures that no value is present for TotalItems, not even an explicit nil
func (o *LowLevelOperationResponseModelCollection) UnsetTotalItems() {
	o.TotalItems.Unset()
}

func (o LowLevelOperationResponseModelCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LowLevelOperationResponseModelCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Items"] = o.Items
	if o.ContinuationToken.IsSet() {
		toSerialize["ContinuationToken"] = o.ContinuationToken.Get()
	}
	if o.TotalItems.IsSet() {
		toSerialize["TotalItems"] = o.TotalItems.Get()
	}
	return toSerialize, nil
}

type NullableLowLevelOperationResponseModelCollection struct {
	value *LowLevelOperationResponseModelCollection
	isSet bool
}

func (v NullableLowLevelOperationResponseModelCollection) Get() *LowLevelOperationResponseModelCollection {
	return v.value
}

func (v *NullableLowLevelOperationResponseModelCollection) Set(val *LowLevelOperationResponseModelCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableLowLevelOperationResponseModelCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableLowLevelOperationResponseModelCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLowLevelOperationResponseModelCollection(val *LowLevelOperationResponseModelCollection) *NullableLowLevelOperationResponseModelCollection {
	return &NullableLowLevelOperationResponseModelCollection{value: val, isSet: true}
}

func (v NullableLowLevelOperationResponseModelCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLowLevelOperationResponseModelCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
