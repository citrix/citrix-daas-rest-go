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

// checks if the SecurityGroupRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityGroupRule{}

// SecurityGroupRule Data contract for SecurityGroupRuleModel.
type SecurityGroupRule struct {
	// Gets or sets the FromPort.
	FromPort int32 `json:"FromPort"`
	// Gets the SecurityGroupIds.
	SecurityGroupIds []string `json:"SecurityGroupIds"`
	// Gets the IPRanges.
	IPRanges []string `json:"IPRanges"`
	// Gets or sets the Protocol.
	Protocol string `json:"Protocol"`
	// Gets or sets the ToPort.
	ToPort int32 `json:"ToPort"`
}

// NewSecurityGroupRule instantiates a new SecurityGroupRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityGroupRule(fromPort int32, securityGroupIds []string, iPRanges []string, protocol string, toPort int32) *SecurityGroupRule {
	this := SecurityGroupRule{}
	this.FromPort = fromPort
	this.SecurityGroupIds = securityGroupIds
	this.IPRanges = iPRanges
	this.Protocol = protocol
	this.ToPort = toPort
	return &this
}

// NewSecurityGroupRuleWithDefaults instantiates a new SecurityGroupRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityGroupRuleWithDefaults() *SecurityGroupRule {
	this := SecurityGroupRule{}
	return &this
}

// GetFromPort returns the FromPort field value
func (o *SecurityGroupRule) GetFromPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FromPort
}

// GetFromPortOk returns a tuple with the FromPort field value
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetFromPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromPort, true
}

// SetFromPort sets field value
func (o *SecurityGroupRule) SetFromPort(v int32) {
	o.FromPort = v
}

// GetSecurityGroupIds returns the SecurityGroupIds field value
func (o *SecurityGroupRule) GetSecurityGroupIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SecurityGroupIds
}

// GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field value
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetSecurityGroupIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecurityGroupIds, true
}

// SetSecurityGroupIds sets field value
func (o *SecurityGroupRule) SetSecurityGroupIds(v []string) {
	o.SecurityGroupIds = v
}

// GetIPRanges returns the IPRanges field value
func (o *SecurityGroupRule) GetIPRanges() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.IPRanges
}

// GetIPRangesOk returns a tuple with the IPRanges field value
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetIPRangesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IPRanges, true
}

// SetIPRanges sets field value
func (o *SecurityGroupRule) SetIPRanges(v []string) {
	o.IPRanges = v
}

// GetProtocol returns the Protocol field value
func (o *SecurityGroupRule) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *SecurityGroupRule) SetProtocol(v string) {
	o.Protocol = v
}

// GetToPort returns the ToPort field value
func (o *SecurityGroupRule) GetToPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ToPort
}

// GetToPortOk returns a tuple with the ToPort field value
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetToPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToPort, true
}

// SetToPort sets field value
func (o *SecurityGroupRule) SetToPort(v int32) {
	o.ToPort = v
}

func (o SecurityGroupRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityGroupRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["FromPort"] = o.FromPort
	toSerialize["SecurityGroupIds"] = o.SecurityGroupIds
	toSerialize["IPRanges"] = o.IPRanges
	toSerialize["Protocol"] = o.Protocol
	toSerialize["ToPort"] = o.ToPort
	return toSerialize, nil
}

type NullableSecurityGroupRule struct {
	value *SecurityGroupRule
	isSet bool
}

func (v NullableSecurityGroupRule) Get() *SecurityGroupRule {
	return v.value
}

func (v *NullableSecurityGroupRule) Set(val *SecurityGroupRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityGroupRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityGroupRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityGroupRule(val *SecurityGroupRule) *NullableSecurityGroupRule {
	return &NullableSecurityGroupRule{value: val, isSet: true}
}

func (v NullableSecurityGroupRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityGroupRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
