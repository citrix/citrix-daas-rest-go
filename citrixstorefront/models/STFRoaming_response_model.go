// Copyright © 2024. Citrix Systems, Inc.
package models

type STFRoamingDetailModel struct {
	ID                         NullableString `json:"Id"`
	Name                       NullableString `json:"Name"`
	Default                    NullableBool   `json:"Default"`
	Edition                    NullableInt    `json:"Edition"`
	Version                    NullableInt    `json:"Version"`
	Logon                      NullableInt    `json:"Logon"`
	SmartCardFallback          NullableInt    `json:"SmartCardFallback"`
	NetScalerTrustCertificates []any          `json:"NetScalerTrustCertificates"` //TODO: Placeholder for NetScalerTrustCertificates
	NetScalerGatewayMode       NullableString `json:"NetScalerGatewayMode"`
	CallbackURL                any            `json:"CallbackUrl"` //TODO: Placeholder for CallbackURL
	RWMode                     NullableString `json:"RWMode"`
	Deployment                 NullableString `json:"Deployment"`
	Location                   NullableString `json:"Location"`
	GslbLocation               any            `json:"GslbLocation"` //TODO: Placeholder for GslbLocation
	SessionReliability         NullableBool   `json:"SessionReliability"`
	RequestTicketTwoStas       NullableBool   `json:"RequestTicketTwoStas"`
	IPAddress                  NullableString `json:"IpAddress"`
	StasUseLoadBalancing       NullableBool   `json:"StasUseLoadBalancing"`
	StasBypassDuration         TimeModel      `json:"StasBypassDuration"`
	SecureTicketAuthorityUrls  []AuthorityUrl `json:"SecureTicketAuthorityUrls"`
	IsCloudGateway             NullableBool   `json:"IsCloudGateway"`
	GatewayServiceLookupURL    NullableString `json:"GatewayServiceLookupUrl"`
	AuthenticationCapable      NullableBool   `json:"AuthenticationCapable"`
	HdxRoutingCapable          NullableBool   `json:"HdxRoutingCapable"`
	NetScalerImport            NullableBool   `json:"NetScalerImport"`
}

type AuthorityUrl struct {
	AuthorityId          NullableString
	StaUrl               NullableString
	StaValidationEnabled NullableString
	StaValidationSecret  NullableString
}
