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

// checks if the ResourceConnections type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceConnections{}

// ResourceConnections Enumerable of HostingUnit
type ResourceConnections struct {
	// Enumerable of HostingUnit
	Items []AwsEdcDirectoryConnection `json:"items,omitempty"`
}

// NewResourceConnections instantiates a new ResourceConnections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceConnections() *ResourceConnections {
	this := ResourceConnections{}
	return &this
}

// NewResourceConnectionsWithDefaults instantiates a new ResourceConnections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceConnectionsWithDefaults() *ResourceConnections {
	this := ResourceConnections{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceConnections) GetItems() []AwsEdcDirectoryConnection {
	if o == nil {
		var ret []AwsEdcDirectoryConnection
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceConnections) GetItemsOk() ([]AwsEdcDirectoryConnection, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ResourceConnections) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []AwsEdcDirectoryConnection and assigns it to the Items field.
func (o *ResourceConnections) SetItems(v []AwsEdcDirectoryConnection) {
	o.Items = v
}

func (o ResourceConnections) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceConnections) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableResourceConnections struct {
	value *ResourceConnections
	isSet bool
}

func (v NullableResourceConnections) Get() *ResourceConnections {
	return v.value
}

func (v *NullableResourceConnections) Set(val *ResourceConnections) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceConnections) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceConnections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceConnections(val *ResourceConnections) *NullableResourceConnections {
	return &NullableResourceConnections{value: val, isSet: true}
}

func (v NullableResourceConnections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceConnections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
