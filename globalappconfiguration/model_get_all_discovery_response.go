/*
Global App Config Admin

Describes API used by Global App Config Admin Service

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package globalappconfiguration

import (
	"encoding/json"
)

// checks if the GetAllDiscoveryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllDiscoveryResponse{}

// GetAllDiscoveryResponse struct for GetAllDiscoveryResponse
type GetAllDiscoveryResponse struct {
	Count     *int32                 `json:"count,omitempty"`
	Items     []DiscoveryRecordModel `json:"items,omitempty"`
	NextToken *string                `json:"nextToken,omitempty"`
}

// NewGetAllDiscoveryResponse instantiates a new GetAllDiscoveryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllDiscoveryResponse() *GetAllDiscoveryResponse {
	this := GetAllDiscoveryResponse{}
	return &this
}

// NewGetAllDiscoveryResponseWithDefaults instantiates a new GetAllDiscoveryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllDiscoveryResponseWithDefaults() *GetAllDiscoveryResponse {
	this := GetAllDiscoveryResponse{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GetAllDiscoveryResponse) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllDiscoveryResponse) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GetAllDiscoveryResponse) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *GetAllDiscoveryResponse) SetCount(v int32) {
	o.Count = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *GetAllDiscoveryResponse) GetItems() []DiscoveryRecordModel {
	if o == nil || IsNil(o.Items) {
		var ret []DiscoveryRecordModel
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllDiscoveryResponse) GetItemsOk() ([]DiscoveryRecordModel, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GetAllDiscoveryResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []DiscoveryRecordModel and assigns it to the Items field.
func (o *GetAllDiscoveryResponse) SetItems(v []DiscoveryRecordModel) {
	o.Items = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *GetAllDiscoveryResponse) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllDiscoveryResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *GetAllDiscoveryResponse) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *GetAllDiscoveryResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o GetAllDiscoveryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllDiscoveryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableGetAllDiscoveryResponse struct {
	value *GetAllDiscoveryResponse
	isSet bool
}

func (v NullableGetAllDiscoveryResponse) Get() *GetAllDiscoveryResponse {
	return v.value
}

func (v *NullableGetAllDiscoveryResponse) Set(val *GetAllDiscoveryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllDiscoveryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllDiscoveryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllDiscoveryResponse(val *GetAllDiscoveryResponse) *NullableGetAllDiscoveryResponse {
	return &NullableGetAllDiscoveryResponse{value: val, isSet: true}
}

func (v NullableGetAllDiscoveryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllDiscoveryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
