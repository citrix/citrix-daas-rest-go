// Copyright © 2024. Citrix Systems, Inc.

package models

import "fmt"

type STFStoreGatewayServiceRawResponseModel struct {
	Enabled                  NullableBool   `json:"PnaEnabled,omitempty"`
	CustomerId               NullableString `json:"CustomerId,omitempty"`
	InstanceId               NullableString `json:"InstanceId,omitempty"`
	GetGatewayServiceUrl     NullableString `json:"GetGatewayServiceUrl,omitempty"`
	PrivateKey               NullableString `json:"PrivateKey,omitempty"`
	ServiceName              NullableString `json:"ServiceName,omitempty"`
	SecureTicketAuthorityUrl NullableString `json:"SecureTicketAuthorityUrl,omitempty"`
	SecureTicketLifetime     TimeModel      `json:"SecureTicketLifetime,omitempty"`
	SessionReliability       NullableBool   `json:"SessionReliability,omitempty"`
	HandleZones              []string       `json:"HandleZones,omitempty"`
	IgnoreZones              []string       `json:"IgnoreZones,omitempty"`
	GatewayDiscoveryProtocol NullableString `json:"GatewayDiscoveryProtocol,omitempty"`
}

type STFStoreGatewayServiceResponseModel struct {
	Enabled                  NullableBool   `json:"PnaEnabled,omitempty"`
	CustomerId               NullableString `json:"CustomerId,omitempty"`
	InstanceId               NullableString `json:"InstanceId,omitempty"`
	GetGatewayServiceUrl     NullableString `json:"GetGatewayServiceUrl,omitempty"`
	PrivateKey               NullableString `json:"PrivateKey,omitempty"`
	ServiceName              NullableString `json:"ServiceName,omitempty"`
	SecureTicketAuthorityUrl NullableString `json:"SecureTicketAuthorityUrl,omitempty"`
	SecureTicketLifetime     string         `json:"SecureTicketLifetime,omitempty"`
	SessionReliability       NullableBool   `json:"SessionReliability,omitempty"`
	HandleZones              []string       `json:"HandleZones,omitempty"`
	IgnoreZones              []string       `json:"IgnoreZones,omitempty"`
	GatewayDiscoveryProtocol NullableString `json:"GatewayDiscoveryProtocol,omitempty"`
}

func (rawResponse STFStoreGatewayServiceRawResponseModel) ConvertToResponseModel() STFStoreGatewayServiceResponseModel {
	response := STFStoreGatewayServiceResponseModel{
		Enabled:                  rawResponse.Enabled,
		CustomerId:               rawResponse.CustomerId,
		InstanceId:               rawResponse.InstanceId,
		GetGatewayServiceUrl:     rawResponse.GetGatewayServiceUrl,
		PrivateKey:               rawResponse.PrivateKey,
		ServiceName:              rawResponse.ServiceName,
		SecureTicketAuthorityUrl: rawResponse.SecureTicketAuthorityUrl,
		SecureTicketLifetime:     fmt.Sprintf("%02d:%02d:%02d", *rawResponse.SecureTicketLifetime.Hours.Get(), *rawResponse.SecureTicketLifetime.Minutes.Get(), *rawResponse.SecureTicketLifetime.Seconds.Get()),
		SessionReliability:       rawResponse.SessionReliability,
		HandleZones:              rawResponse.HandleZones,
		IgnoreZones:              rawResponse.IgnoreZones,
		GatewayDiscoveryProtocol: rawResponse.GatewayDiscoveryProtocol,
	}
	return response
}
