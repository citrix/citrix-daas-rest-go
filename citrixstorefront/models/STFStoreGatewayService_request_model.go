// Copyright © 2024. Citrix Systems, Inc.
package models

// Set STFStoreGatewayServiceSetRequestModel Model
var _ MappedNullable = &STFStoreGatewayServiceSetRequestModel{}

type STFStoreGatewayServiceSetRequestModel struct {
	Enabled                  NullableBool   `json:"Enabled,omitempty"`                  // Enable use of the gateway service
	CustomerId               NullableString `json:"CustomerId,omitempty"`               // The CWC customer id
	GetGatewayServiceUrl     NullableString `json:"GetGatewayServiceUrl,omitempty"`     // The URL of the service used to retrieve gateway address (FQDN)
	PrivateKey               NullableString `json:"PrivateKey,omitempty"`               // The private key for CWC trust
	ServiceName              NullableString `json:"ServiceName,omitempty"`              // The service name for CWC trust
	InstanceId               NullableString `json:"InstanceId,omitempty"`               // The instance id for CWC trust
	SecureTicketAuthorityUrl NullableString `json:"SecureTicketAuthorityUrl,omitempty"` // The URL of the CWC STA service
	SecureTicketLifetime     NullableString `json:"SecureTicketLifetime,omitempty"`     // The lifetime requested for CWC STA service tickets
	SessionReliability       NullableBool   `json:"SessionReliability,omitempty"`       // A value indicating whether session reliability should be enabled
	GatewayDiscoveryProtocol NullableString `json:"GatewayDiscoveryProtocol,omitempty"` // A value indicating the protocol the gateway service should use. Valid values are GWaaS or ConfigService.
	IgnoreZones              []string       `json:"IgnoreZones,omitempty"`              // A value indicating that the gateway service should not be used for the specified zones
	HandleZones              []string       `json:"HandleZones,omitempty"`              // A value indicating that the gateway service should be used for the specified zones
}

func (o STFStoreGatewayServiceSetRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled.IsSet() {
		toSerialize["Enabled"] = o.Enabled.Get()
	}
	if o.CustomerId.IsSet() {
		toSerialize["CustomerId"] = o.CustomerId.Get()
	}
	if o.GetGatewayServiceUrl.IsSet() {
		toSerialize["GetGatewayServiceUrl"] = o.GetGatewayServiceUrl.Get()
	}
	if o.PrivateKey.IsSet() {
		toSerialize["PrivateKey"] = o.PrivateKey.Get()
	}
	if o.ServiceName.IsSet() {
		toSerialize["ServiceName"] = o.ServiceName.Get()
	}
	if o.InstanceId.IsSet() {
		toSerialize["InstanceId"] = o.InstanceId.Get()
	}
	if o.SecureTicketAuthorityUrl.IsSet() {
		toSerialize["SecureTicketAuthorityUrl"] = o.SecureTicketAuthorityUrl.Get()
	}
	if o.SecureTicketLifetime.IsSet() {
		toSerialize["SecureTicketLifetime"] = o.SecureTicketLifetime.Get()
	}
	if o.SessionReliability.IsSet() {
		toSerialize["SessionReliability"] = o.SessionReliability.Get()
	}
	if o.GatewayDiscoveryProtocol.IsSet() {
		toSerialize["GatewayDiscoveryProtocol"] = o.GatewayDiscoveryProtocol.Get()
	}
	if len(o.IgnoreZones) > 0 {
		toSerialize["IgnoreZones"] = o.IgnoreZones
	}
	if len(o.HandleZones) > 0 {
		toSerialize["HandleZones"] = o.HandleZones
	}
	return toSerialize, nil
}

func (o *STFStoreGatewayServiceSetRequestModel) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}

func (o *STFStoreGatewayServiceSetRequestModel) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}

func (o *STFStoreGatewayServiceSetRequestModel) SetGetGatewayServiceUrl(v string) {
	o.GetGatewayServiceUrl.Set(&v)
}

func (o *STFStoreGatewayServiceSetRequestModel) SetPrivateKey(v string) {
	o.PrivateKey.Set(&v)
}

func (o *STFStoreGatewayServiceSetRequestModel) SetServiceName(v string) {
	o.ServiceName.Set(&v)
}

func (o *STFStoreGatewayServiceSetRequestModel) SetInstanceId(v string) {
	o.InstanceId.Set(&v)
}

func (o *STFStoreGatewayServiceSetRequestModel) SetSecureTicketAuthorityUrl(v string) {
	o.SecureTicketAuthorityUrl.Set(&v)
}

func (o *STFStoreGatewayServiceSetRequestModel) SetSecureTicketLifetime(v string) {
	o.SecureTicketLifetime.Set(&v)
}

func (o *STFStoreGatewayServiceSetRequestModel) SetSessionReliability(v bool) {
	o.SessionReliability.Set(&v)
}

func (o *STFStoreGatewayServiceSetRequestModel) SetGatewayDiscoveryProtocol(v string) {
	o.GatewayDiscoveryProtocol.Set(&v)
}

func (o *STFStoreGatewayServiceSetRequestModel) SetIgnoreZones(v []string) {
	o.IgnoreZones = v
}

func (o *STFStoreGatewayServiceSetRequestModel) SetHandleZones(v []string) {
	o.HandleZones = v
}
