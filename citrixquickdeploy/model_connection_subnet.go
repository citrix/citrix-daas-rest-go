/*
Citrix Virtual App & Desktop Catalog Service 147.0.26651.57932

Catalog Service

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickdeploy

import (
	"encoding/json"
)

// checks if the ConnectionSubnet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionSubnet{}

// ConnectionSubnet struct for ConnectionSubnet
type ConnectionSubnet struct {
	// IP Address of the subnet
	SubnetAddress string `json:"subnetAddress"`
	// Subnet mask associated with the subnet
	SubnetMask int32 `json:"subnetMask"`
}

// NewConnectionSubnetWithDefaults instantiates a new ConnectionSubnet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionSubnetWithDefaults() *ConnectionSubnet {
	this := ConnectionSubnet{}
	return &this
}

// GetSubnetAddress returns the SubnetAddress field value
func (o *ConnectionSubnet) GetSubnetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubnetAddress
}

// GetSubnetAddressOk returns a tuple with the SubnetAddress field value
// and a boolean to check if the value has been set.
func (o *ConnectionSubnet) GetSubnetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubnetAddress, true
}

// SetSubnetAddress sets field value
func (o *ConnectionSubnet) SetSubnetAddress(v string) {
	o.SubnetAddress = v
}

// GetSubnetMask returns the SubnetMask field value
func (o *ConnectionSubnet) GetSubnetMask() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SubnetMask
}

// GetSubnetMaskOk returns a tuple with the SubnetMask field value
// and a boolean to check if the value has been set.
func (o *ConnectionSubnet) GetSubnetMaskOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubnetMask, true
}

// SetSubnetMask sets field value
func (o *ConnectionSubnet) SetSubnetMask(v int32) {
	o.SubnetMask = v
}

func (o ConnectionSubnet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionSubnet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subnetAddress"] = o.SubnetAddress
	toSerialize["subnetMask"] = o.SubnetMask
	return toSerialize, nil
}

type NullableConnectionSubnet struct {
	value *ConnectionSubnet
	isSet bool
}

func (v NullableConnectionSubnet) Get() *ConnectionSubnet {
	return v.value
}

func (v *NullableConnectionSubnet) Set(val *ConnectionSubnet) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionSubnet) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionSubnet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionSubnet(val *ConnectionSubnet) *NullableConnectionSubnet {
	return &NullableConnectionSubnet{value: val, isSet: true}
}

func (v NullableConnectionSubnet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionSubnet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
