// Copyright © 2025. Citrix Systems, Inc.

package models

// Model for Enable-STFAuthenticationServiceProtocol

var _ MappedNullable = &EnableSTFAuthenticationServiceProtocolRequestModel{}

type EnableSTFAuthenticationServiceProtocolRequestModel struct {
	Name []string `json:"Name,omitempty"` // The protocol type to enable. Possible values are 'SAML', 'OAuth', 'OpenIDConnect', 'WSFederation'
}

func (o *EnableSTFAuthenticationServiceProtocolRequestModel) SetName(v []string) {
	o.Name = v
}

// ToMap implements MappedNullable.
func (o *EnableSTFAuthenticationServiceProtocolRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if len(o.Name) > 0 {
		toSerialize["Name"] = o.Name
	}
	return toSerialize, nil
}

// Model for Disable-STFAuthenticationServiceProtocol

var _ MappedNullable = &DisableSTFAuthenticationServiceProtocolRequestModel{}

type DisableSTFAuthenticationServiceProtocolRequestModel struct {
	Name []string `json:"Name,omitempty"` // The protocol type to disable. Possible values are 'SAML', 'OAuth', 'OpenIDConnect', 'WSFederation'
}

func (o *DisableSTFAuthenticationServiceProtocolRequestModel) SetName(v []string) {
	o.Name = v
}

// ToMap implements MappedNullable.
func (o *DisableSTFAuthenticationServiceProtocolRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if len(o.Name) > 0 {
		toSerialize["Name"] = o.Name
	}
	return toSerialize, nil
}
