// Copyright © 2025. Citrix Systems, Inc.
package models

type STFCitrixAGBasicOptionsResponseModel struct {
	CredentialValidationMode int                      `json:"CredentialValidationMode,omitempty"`
	NetscalerGateways        []NetscalerGatewaysModel `json:"NetscalerGateways,omitempty"` // The netscaler gateways of the citrix AG basic authentication
	ClaimsFactoryName        NullableString           `json:"ClaimsFactoryName,omitempty"` // The claims factory name of the citrix AG basic authentication
}
