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

// checks if the CollectionModelOfIdentityDomainResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionModelOfIdentityDomainResponseModel{}

// CollectionModelOfIdentityDomainResponseModel Response object for collections of items.
type CollectionModelOfIdentityDomainResponseModel struct {
	// List of items.
	Items []IdentityDomainResponseModel `json:"Items"`
	// If present, indicates to the caller that the query was not complete, and they should call the API again specifying the continuation token as a query parameter.
	ContinuationToken NullableString `json:"ContinuationToken,omitempty"`
	// Indicates the total number of items in the collection, which may be more than the number of Items returned, if there is a ContinuationToken.  Only returned in the response to `$search` APIs.
	TotalItems NullableInt32 `json:"TotalItems,omitempty"`
}

// NewCollectionModelOfIdentityDomainResponseModel instantiates a new CollectionModelOfIdentityDomainResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionModelOfIdentityDomainResponseModel(items []IdentityDomainResponseModel) *CollectionModelOfIdentityDomainResponseModel {
	this := CollectionModelOfIdentityDomainResponseModel{}
	this.Items = items
	return &this
}

// NewCollectionModelOfIdentityDomainResponseModelWithDefaults instantiates a new CollectionModelOfIdentityDomainResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionModelOfIdentityDomainResponseModelWithDefaults() *CollectionModelOfIdentityDomainResponseModel {
	this := CollectionModelOfIdentityDomainResponseModel{}
	return &this
}

// GetItems returns the Items field value
func (o *CollectionModelOfIdentityDomainResponseModel) GetItems() []IdentityDomainResponseModel {
	if o == nil {
		var ret []IdentityDomainResponseModel
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CollectionModelOfIdentityDomainResponseModel) GetItemsOk() ([]IdentityDomainResponseModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *CollectionModelOfIdentityDomainResponseModel) SetItems(v []IdentityDomainResponseModel) {
	o.Items = v
}

// GetContinuationToken returns the ContinuationToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionModelOfIdentityDomainResponseModel) GetContinuationToken() string {
	if o == nil || IsNil(o.ContinuationToken.Get()) {
		var ret string
		return ret
	}
	return *o.ContinuationToken.Get()
}

// GetContinuationTokenOk returns a tuple with the ContinuationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionModelOfIdentityDomainResponseModel) GetContinuationTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinuationToken.Get(), o.ContinuationToken.IsSet()
}

// HasContinuationToken returns a boolean if a field has been set.
func (o *CollectionModelOfIdentityDomainResponseModel) HasContinuationToken() bool {
	if o != nil && o.ContinuationToken.IsSet() {
		return true
	}

	return false
}

// SetContinuationToken gets a reference to the given NullableString and assigns it to the ContinuationToken field.
func (o *CollectionModelOfIdentityDomainResponseModel) SetContinuationToken(v string) {
	o.ContinuationToken.Set(&v)
}

// SetContinuationTokenNil sets the value for ContinuationToken to be an explicit nil
func (o *CollectionModelOfIdentityDomainResponseModel) SetContinuationTokenNil() {
	o.ContinuationToken.Set(nil)
}

// UnsetContinuationToken ensures that no value is present for ContinuationToken, not even an explicit nil
func (o *CollectionModelOfIdentityDomainResponseModel) UnsetContinuationToken() {
	o.ContinuationToken.Unset()
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionModelOfIdentityDomainResponseModel) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalItems.Get()
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionModelOfIdentityDomainResponseModel) GetTotalItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalItems.Get(), o.TotalItems.IsSet()
}

// HasTotalItems returns a boolean if a field has been set.
func (o *CollectionModelOfIdentityDomainResponseModel) HasTotalItems() bool {
	if o != nil && o.TotalItems.IsSet() {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given NullableInt32 and assigns it to the TotalItems field.
func (o *CollectionModelOfIdentityDomainResponseModel) SetTotalItems(v int32) {
	o.TotalItems.Set(&v)
}

// SetTotalItemsNil sets the value for TotalItems to be an explicit nil
func (o *CollectionModelOfIdentityDomainResponseModel) SetTotalItemsNil() {
	o.TotalItems.Set(nil)
}

// UnsetTotalItems ensures that no value is present for TotalItems, not even an explicit nil
func (o *CollectionModelOfIdentityDomainResponseModel) UnsetTotalItems() {
	o.TotalItems.Unset()
}

func (o CollectionModelOfIdentityDomainResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionModelOfIdentityDomainResponseModel) ToMap() (map[string]interface{}, error) {
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

type NullableCollectionModelOfIdentityDomainResponseModel struct {
	value *CollectionModelOfIdentityDomainResponseModel
	isSet bool
}

func (v NullableCollectionModelOfIdentityDomainResponseModel) Get() *CollectionModelOfIdentityDomainResponseModel {
	return v.value
}

func (v *NullableCollectionModelOfIdentityDomainResponseModel) Set(val *CollectionModelOfIdentityDomainResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionModelOfIdentityDomainResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionModelOfIdentityDomainResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionModelOfIdentityDomainResponseModel(val *CollectionModelOfIdentityDomainResponseModel) *NullableCollectionModelOfIdentityDomainResponseModel {
	return &NullableCollectionModelOfIdentityDomainResponseModel{value: val, isSet: true}
}

func (v NullableCollectionModelOfIdentityDomainResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionModelOfIdentityDomainResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
