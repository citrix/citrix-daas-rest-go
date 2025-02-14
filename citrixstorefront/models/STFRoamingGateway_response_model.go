// Copyright © 2024. Citrix Systems, Inc.
package models

import (
	"fmt"
	"strconv"
)

type STFRoamingGatewayRawResponseModel struct {
	SiteId                     int64              `json:"SiteId,omitempty"`
	Name                       NullableString     `json:"Name,omitempty"`
	Default                    NullableBool       `json:"Default,omitempty"`
	Edition                    NullableInt        `json:"Edition,omitempty"`
	Version                    NullableInt        `json:"Version,omitempty"`
	Logon                      NullableInt        `json:"Logon,omitempty"`
	SmartCardFallback          NullableInt        `json:"SmartCardFallback,omitempty"`
	NetScalerTrustCertificates []CertificateModel `json:"NetScalerTrustCertificates,omitempty"`
	NetScalerGatewayMode       NullableString     `json:"NetScalerGatewayMode,omitempty"`
	CallbackUrl                NullableString     `json:"CallbackUrl,omitempty"`
	RWMode                     NullableString     `json:"RWMode,omitempty"`
	Deployment                 NullableString     `json:"Deployment,omitempty"`
	Location                   NullableString     `json:"Location,omitempty"`
	GslbLocation               NullableString     `json:"GslbLocation,omitempty"`
	SessionReliability         NullableBool       `json:"SessionReliability,omitempty"`
	RequestTicketTwoStas       NullableBool       `json:"RequestTicketTwoStas,omitempty"`
	IPAddress                  NullableString     `json:"IpAddress,omitempty"`
	StasUseLoadBalancing       NullableBool       `json:"StasUseLoadBalancing,omitempty"`
	StasBypassDuration         TimeModel          `json:"StasBypassDuration,omitempty"`
	SecureTicketAuthorityUrls  []STFSTAUrlModel   `json:"SecureTicketAuthorityUrls,omitempty"`
	IsCloudGateway             NullableBool       `json:"IsCloudGateway,omitempty"`
	GatewayServiceLookupURL    NullableString     `json:"GatewayServiceLookupUrl,omitempty"`
	AuthenticationCapable      NullableBool       `json:"AuthenticationCapable,omitempty"`
	HdxRoutingCapable          NullableBool       `json:"HdxRoutingCapable,omitempty"`
	NetScalerImport            NullableBool       `json:"NetScalerImport,omitempty"`
}

type STFSTAUrlModel struct {
	StaUrl               NullableString `json:"StaUrl,omitempty"`
	StaValidationEnabled NullableBool   `json:"StaValidationEnabled,omitempty"`
	StaValidationSecret  NullableString `json:"StaValidationSecret,omitempty"`
}

func (r *STFSTAUrlModel) SetStaUrl(staUrl string) {
	r.StaUrl.Set(&staUrl)
}

func (r *STFSTAUrlModel) SetStaValidationEnabled(staValidationEnabled bool) {
	r.StaValidationEnabled.Set(&staValidationEnabled)
}

func (r *STFSTAUrlModel) SetStaValidationSecret(staValidationSecret string) {
	r.StaValidationSecret.Set(&staValidationSecret)
}

func (r *STFRoamingGatewayRawResponseModel) ConvertToResponseModel() STFRoamingGatewayResponseModel {
	response := STFRoamingGatewayResponseModel{}
	response.SiteId = strconv.Itoa(int(r.SiteId))
	if r.Name.IsSet() && r.Name.Get() != nil {
		response.SetName(*r.Name.Get())
	}

	if r.Default.IsSet() && r.Default.Get() != nil {
		response.SetDefault(*r.Default.Get())
	}

	if r.Edition.IsSet() && r.Edition.Get() != nil {
		response.SetEdition(*r.Edition.Get())
	}

	if r.Version.IsSet() && r.Version.Get() != nil {
		response.SetVersion(*r.Version.Get())
	}

	if r.Logon.IsSet() && r.Logon.Get() != nil {
		response.SetLogonType(*r.Logon.Get())
	}

	if r.SmartCardFallback.IsSet() && r.SmartCardFallback.Get() != nil {
		response.SetSmartCardFallbackLogonType(*r.SmartCardFallback.Get())
	}

	if r.NetScalerTrustCertificates != nil {
		response.SetNetScalerTrustCertificates(r.NetScalerTrustCertificates)
	}

	if r.NetScalerGatewayMode.IsSet() && r.NetScalerGatewayMode.Get() != nil {
		response.SetNetScalerGatewayMode(*r.NetScalerGatewayMode.Get())
	}

	if r.CallbackUrl.IsSet() && r.CallbackUrl.Get() != nil {
		response.SetCallbackUrl(*r.CallbackUrl.Get())
	}

	if r.RWMode.IsSet() && r.RWMode.Get() != nil {
		response.SetRWMode(*r.RWMode.Get())
	}

	if r.Deployment.IsSet() && r.Deployment.Get() != nil {
		response.SetDeployment(*r.Deployment.Get())
	}

	if r.Location.IsSet() && r.Location.Get() != nil {
		response.SetGatewayUrl(*r.Location.Get())
	}

	if r.GslbLocation.IsSet() && r.GslbLocation.Get() != nil {
		response.SetGslbUrl(*r.GslbLocation.Get())
	}

	if r.SessionReliability.IsSet() && r.SessionReliability.Get() != nil {
		response.SetSessionReliability(*r.SessionReliability.Get())
	}

	if r.RequestTicketTwoStas.IsSet() && r.RequestTicketTwoStas.Get() != nil {
		response.SetRequestTicketTwoSTAs(*r.RequestTicketTwoStas.Get())
	}

	if r.IPAddress.IsSet() && r.IPAddress.Get() != nil {
		response.SetSubnetIPAddress(*r.IPAddress.Get())
	}

	if r.StasUseLoadBalancing.IsSet() && r.StasUseLoadBalancing.Get() != nil {
		response.SetStasUseLoadBalancing(*r.StasUseLoadBalancing.Get())
	}

	response.SetStasBypassDuration(fmt.Sprintf("%d.%d:%d:%d", *r.StasBypassDuration.Days.Get(), *r.StasBypassDuration.Hours.Get(), *r.StasBypassDuration.Minutes.Get(), *r.StasBypassDuration.Seconds.Get()))

	if r.SecureTicketAuthorityUrls != nil {
		response.SetSecureTicketAuthorityUrls(r.SecureTicketAuthorityUrls)
	}

	if r.IsCloudGateway.IsSet() && r.IsCloudGateway.Get() != nil {
		response.SetIsCloudGateway(*r.IsCloudGateway.Get())
	}

	if r.GatewayServiceLookupURL.IsSet() && r.GatewayServiceLookupURL.Get() != nil {
		response.SetGatewayServiceLookupURL(*r.GatewayServiceLookupURL.Get())
	}

	if r.AuthenticationCapable.IsSet() && r.AuthenticationCapable.Get() != nil {
		response.SetAuthenticationCapable(*r.AuthenticationCapable.Get())
	}

	if r.HdxRoutingCapable.IsSet() && r.HdxRoutingCapable.Get() != nil {
		response.SetHdxRoutingCapable(*r.HdxRoutingCapable.Get())
	}

	if r.NetScalerImport.IsSet() && r.NetScalerImport.Get() != nil {
		response.SetNetScalerImport(*r.NetScalerImport.Get())
	}

	return response
}

type STFRoamingGatewayResponseModel struct {
	SiteId                     string             `json:"SiteId,omitempty"`
	Name                       NullableString     `json:"Name,omitempty"`
	Default                    NullableBool       `json:"Default,omitempty"`
	Edition                    NullableString     `json:"Edition,omitempty"`
	Version                    NullableString     `json:"Version,omitempty"`
	LogonType                  NullableString     `json:"LogonType,omitempty"`
	SmartCardFallbackLogonType NullableString     `json:"SmartCardFallbackLogonType,omitempty"`
	NetScalerTrustCertificates []CertificateModel `json:"NetScalerTrustCertificates,omitempty"`
	NetScalerGatewayMode       NullableString     `json:"NetScalerGatewayMode,omitempty"`
	CallbackUrl                NullableString     `json:"CallbackUrl,omitempty"`
	RWMode                     NullableString     `json:"RWMode,omitempty"`
	Deployment                 NullableString     `json:"Deployment,omitempty"`
	GatewayUrl                 NullableString     `json:"GatewayUrl,omitempty"`
	GslbUrl                    NullableString     `json:"GslbUrl,omitempty"`
	SessionReliability         NullableBool       `json:"SessionReliability,omitempty"`
	RequestTicketTwoSTAs       NullableBool       `json:"RequestTicketTwoSTAs,omitempty"`
	SubnetIPAddress            NullableString     `json:"SubnetIPAddress,omitempty"`
	StasUseLoadBalancing       NullableBool       `json:"StasUseLoadBalancing,omitempty"`
	StasBypassDuration         NullableString     `json:"StasBypassDuration,omitempty"`
	SecureTicketAuthorityUrls  []STFSTAUrlModel   `json:"SecureTicketAuthorityUrls,omitempty"`
	IsCloudGateway             NullableBool       `json:"IsCloudGateway,omitempty"`
	GatewayServiceLookupURL    NullableString     `json:"GatewayServiceLookupUrl,omitempty"`
	AuthenticationCapable      NullableBool       `json:"AuthenticationCapable,omitempty"`
	HdxRoutingCapable          NullableBool       `json:"HdxRoutingCapable,omitempty"`
	NetScalerImport            NullableBool       `json:"NetScalerImport,omitempty"`
}

func (r *STFRoamingGatewayResponseModel) SetSiteId(siteId int64) {
	r.SiteId = strconv.Itoa(int(siteId))
}

func (r *STFRoamingGatewayResponseModel) SetName(name string) {
	r.Name.Set(&name)
}

func (r *STFRoamingGatewayResponseModel) SetDefault(defaultValue bool) {
	r.Default.Set(&defaultValue)
}

func (r *STFRoamingGatewayResponseModel) SetEdition(edition int) {
	editionString := "Enterprise"
	switch edition {
	case 0:
		editionString = "Advanced"
	}
	r.Edition.Set(&editionString)

}

func (r *STFRoamingGatewayResponseModel) SetVersion(version int) {
	versionString := "Version10_0_69_4"
	switch version {
	case 1:
		versionString = "Version9x"
	}
	r.Version.Set(&versionString)
}

func (r *STFRoamingGatewayResponseModel) SetLogonType(logonType int) {
	logonTypeString := "None"
	switch logonType {
	case 0:
		logonTypeString = "UsedForHDXOnly"
	case 1:
		logonTypeString = "Domain"
	case 2:
		logonTypeString = "RSA"
	case 3:
		logonTypeString = "DomainAndRSA"
	case 4:
		logonTypeString = "SMS"
	case 5:
		logonTypeString = "GatewayKnows"
	case 6:
		logonTypeString = "SmartCard"
	}

	r.LogonType.Set(&logonTypeString)
}

func (r *STFRoamingGatewayResponseModel) SetSmartCardFallbackLogonType(smartCardFallback int) {
	smartCardFallbackLogonTypeString := "None"
	switch smartCardFallback {
	case 0:
		smartCardFallbackLogonTypeString = "UsedForHDXOnly"
	case 1:
		smartCardFallbackLogonTypeString = "Domain"
	case 2:
		smartCardFallbackLogonTypeString = "RSA"
	case 3:
		smartCardFallbackLogonTypeString = "DomainAndRSA"
	case 4:
		smartCardFallbackLogonTypeString = "SMS"
	case 5:
		smartCardFallbackLogonTypeString = "GatewayKnows"
	case 6:
		smartCardFallbackLogonTypeString = "SmartCard"
	}
	r.SmartCardFallbackLogonType.Set(&smartCardFallbackLogonTypeString)
}

func (r *STFRoamingGatewayResponseModel) SetNetScalerTrustCertificates(certificates []CertificateModel) {
	r.NetScalerTrustCertificates = certificates
}

func (r *STFRoamingGatewayResponseModel) SetNetScalerGatewayMode(netScalerGatewayMode string) {
	r.NetScalerGatewayMode.Set(&netScalerGatewayMode)
}

func (r *STFRoamingGatewayResponseModel) SetCallbackUrl(callbackUrl string) {
	r.CallbackUrl.Set(&callbackUrl)
}

func (r *STFRoamingGatewayResponseModel) SetRWMode(rwMode string) {
	r.RWMode.Set(&rwMode)
}

func (r *STFRoamingGatewayResponseModel) SetDeployment(deployment string) {
	r.Deployment.Set(&deployment)
}

func (r *STFRoamingGatewayResponseModel) SetGatewayUrl(gatewayUrl string) {
	r.GatewayUrl.Set(&gatewayUrl)
}

func (r *STFRoamingGatewayResponseModel) SetGslbUrl(gslbUrl string) {
	r.GslbUrl.Set(&gslbUrl)
}

func (r *STFRoamingGatewayResponseModel) SetSessionReliability(sessionReliability bool) {
	r.SessionReliability.Set(&sessionReliability)
}

func (r *STFRoamingGatewayResponseModel) SetRequestTicketTwoSTAs(requestTicketTwoSTAs bool) {
	r.RequestTicketTwoSTAs.Set(&requestTicketTwoSTAs)
}

func (r *STFRoamingGatewayResponseModel) SetSubnetIPAddress(subnetIPAddress string) {
	r.SubnetIPAddress.Set(&subnetIPAddress)
}

func (r *STFRoamingGatewayResponseModel) SetStasUseLoadBalancing(stasUseLoadBalancing bool) {
	r.StasUseLoadBalancing.Set(&stasUseLoadBalancing)
}

func (r *STFRoamingGatewayResponseModel) SetStasBypassDuration(stasBypassDuration string) {
	r.StasBypassDuration.Set(&stasBypassDuration)
}

func (r *STFRoamingGatewayResponseModel) SetSecureTicketAuthorityUrls(secureTicketAuthorityUrls []STFSTAUrlModel) {
	r.SecureTicketAuthorityUrls = secureTicketAuthorityUrls
}

func (r *STFRoamingGatewayResponseModel) SetIsCloudGateway(isCloudGateway bool) {
	r.IsCloudGateway.Set(&isCloudGateway)
}

func (r *STFRoamingGatewayResponseModel) SetGatewayServiceLookupURL(gatewayServiceLookupURL string) {
	r.GatewayServiceLookupURL.Set(&gatewayServiceLookupURL)
}

func (r *STFRoamingGatewayResponseModel) SetAuthenticationCapable(authenticationCapable bool) {
	r.AuthenticationCapable.Set(&authenticationCapable)
}

func (r *STFRoamingGatewayResponseModel) SetHdxRoutingCapable(hdxRoutingCapable bool) {
	r.HdxRoutingCapable.Set(&hdxRoutingCapable)
}

func (r *STFRoamingGatewayResponseModel) SetNetScalerImport(netScalerImport bool) {
	r.NetScalerImport.Set(&netScalerImport)
}
