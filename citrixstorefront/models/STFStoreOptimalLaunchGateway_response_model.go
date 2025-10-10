// Copyright © 2025. Citrix Systems, Inc.

package models

import "fmt"

// Models for Get-STFStoreRegisteredOptimalLaunchGateway

type STFStoreRegisteredOptimalLaunchGatewayRawResponseModel struct {
	EnabledOnDirectAccess      NullableBool     `json:"EnabledOnDirectAccess,omitempty"`
	Farms                      []string         `json:"Farms,omitempty"`
	Zones                      []string         `json:"Zones,omitempty"`
	NeverLaunchThroughGateway  NullableBool     `json:"NeverLaunchThroughGateway,omitempty"`
	Id                         NullableString   `json:"Id,omitempty"`
	Name                       NullableString   `json:"Name,omitempty"`
	Address                    NullableString   `json:"Address,omitempty"`
	SessionReliability         NullableBool     `json:"SessionReliability,omitempty"`
	RequestTicketTwoSTAs       NullableBool     `json:"RequestTicketTwoSTAs,omitempty"`
	Hostnames                  []NullableString `json:"Hostnames,omitempty"`
	SecureTicketAuthorityUrls  []STFSTAUrlModel `json:"SecureTicketAuthorityUrls,omitempty"`
	StasBypassDuration         TimeModel        `json:"StasBypassDuration,omitempty"`
	StasUseLoadBalancing       NullableBool     `json:"StasUseLoadBalancing,omitempty"`
	NetScalerTrustCertificates []string         `json:"NetScalerTrustCertificates,omitempty"`
	IsCloudGateway             NullableBool     `json:"IsCloudGateway,omitempty"`
	GatewayServiceLookupUrl    NullableString   `json:"GatewayServiceLookupUrl,omitempty"`
	CgsHdxGateway              NullableString   `json:"CgsHdxGateway,omitempty"`
	CgsHdxGatewayConnector     NullableString   `json:"CgsHdxGatewayConnector,omitempty"`
	ConnectorBaseURLs          []string         `json:"ConnectorBaseURLs,omitempty"`
	GatewayLookupUrls          []string         `json:"GatewayLookupUrls,omitempty"`
	Key                        NullableString   `json:"Key,omitempty"`
}

type STFStoreRegisteredOptimalLaunchGatewayResponseModel struct {
	EnabledOnDirectAccess      NullableBool     `json:"EnabledOnDirectAccess,omitempty"`
	Farms                      []string         `json:"Farms,omitempty"`
	Zones                      []string         `json:"Zones,omitempty"`
	NeverLaunchThroughGateway  NullableBool     `json:"NeverLaunchThroughGateway,omitempty"`
	Id                         NullableString   `json:"Id,omitempty"`
	Name                       NullableString   `json:"Name,omitempty"`
	Address                    NullableString   `json:"Address,omitempty"`
	SessionReliability         NullableBool     `json:"SessionReliability,omitempty"`
	RequestTicketTwoSTAs       NullableBool     `json:"RequestTicketTwoSTAs,omitempty"`
	Hostnames                  []NullableString `json:"Hostnames,omitempty"`
	SecureTicketAuthorityUrls  []STFSTAUrlModel `json:"SecureTicketAuthorityUrls,omitempty"`
	StasBypassDuration         NullableString   `json:"StasBypassDuration,omitempty"`
	StasUseLoadBalancing       NullableBool     `json:"StasUseLoadBalancing,omitempty"`
	NetScalerTrustCertificates []string         `json:"NetScalerTrustCertificates,omitempty"`
	IsCloudGateway             NullableBool     `json:"IsCloudGateway,omitempty"`
	GatewayServiceLookupUrl    NullableString   `json:"GatewayServiceLookupUrl,omitempty"`
	CgsHdxGateway              NullableString   `json:"CgsHdxGateway,omitempty"`
	CgsHdxGatewayConnector     NullableString   `json:"CgsHdxGatewayConnector,omitempty"`
	ConnectorBaseURLs          []string         `json:"ConnectorBaseURLs,omitempty"`
	GatewayLookupUrls          []string         `json:"GatewayLookupUrls,omitempty"`
	Key                        NullableString   `json:"Key,omitempty"`
}

func (r STFStoreRegisteredOptimalLaunchGatewayRawResponseModel) ConvertToResponseModel() STFStoreRegisteredOptimalLaunchGatewayResponseModel {
	response := STFStoreRegisteredOptimalLaunchGatewayResponseModel{}
	if r.EnabledOnDirectAccess.IsSet() && r.EnabledOnDirectAccess.Get() != nil {
		response.EnabledOnDirectAccess.Set(r.EnabledOnDirectAccess.Get())
	}
	if len(r.Farms) > 0 {
		response.Farms = r.Farms
	}
	if len(r.Zones) > 0 {
		response.Zones = r.Zones
	}
	if r.NeverLaunchThroughGateway.IsSet() && r.NeverLaunchThroughGateway.Get() != nil {
		response.NeverLaunchThroughGateway.Set(r.NeverLaunchThroughGateway.Get())
	}
	if r.Id.IsSet() && r.Id.Get() != nil {
		response.Id.Set(r.Id.Get())
	}
	if r.Name.IsSet() && r.Name.Get() != nil {
		response.Name.Set(r.Name.Get())
	}
	if r.SessionReliability.IsSet() && r.SessionReliability.Get() != nil {
		response.SessionReliability.Set(r.SessionReliability.Get())
	}
	if r.RequestTicketTwoSTAs.IsSet() && r.RequestTicketTwoSTAs.Get() != nil {
		response.RequestTicketTwoSTAs.Set(r.RequestTicketTwoSTAs.Get())
	}
	if r.Hostnames != nil {
		response.Hostnames = r.Hostnames
	}
	if r.SecureTicketAuthorityUrls != nil {
		response.SecureTicketAuthorityUrls = r.SecureTicketAuthorityUrls
	}
	// convert duration to string format if present
	if r.StasBypassDuration.Days.IsSet() {
		s := fmt.Sprintf("%d.%d:%d:%d", *r.StasBypassDuration.Days.Get(), *r.StasBypassDuration.Hours.Get(), *r.StasBypassDuration.Minutes.Get(), *r.StasBypassDuration.Seconds.Get())
		response.StasBypassDuration.Set(&s)
	}
	if r.StasUseLoadBalancing.IsSet() && r.StasUseLoadBalancing.Get() != nil {
		response.StasUseLoadBalancing.Set(r.StasUseLoadBalancing.Get())
	}

	if len(r.NetScalerTrustCertificates) > 0 {
		response.NetScalerTrustCertificates = r.NetScalerTrustCertificates
	}
	if r.Address.IsSet() && r.Address.Get() != nil {
		response.Address.Set(r.Address.Get())
	}

	if r.IsCloudGateway.IsSet() && r.IsCloudGateway.Get() != nil {
		response.IsCloudGateway.Set(r.IsCloudGateway.Get())
	}

	if r.GatewayServiceLookupUrl.IsSet() && r.GatewayServiceLookupUrl.Get() != nil {
		response.GatewayServiceLookupUrl.Set(r.GatewayServiceLookupUrl.Get())
	}
	if r.CgsHdxGateway.IsSet() && r.CgsHdxGateway.Get() != nil {
		response.CgsHdxGateway.Set(r.CgsHdxGateway.Get())
	}
	if r.CgsHdxGatewayConnector.IsSet() && r.CgsHdxGatewayConnector.Get() != nil {
		response.CgsHdxGatewayConnector.Set(r.CgsHdxGatewayConnector.Get())
	}
	if len(r.ConnectorBaseURLs) > 0 {
		response.ConnectorBaseURLs = r.ConnectorBaseURLs
	}
	if len(r.GatewayLookupUrls) > 0 {
		response.GatewayLookupUrls = r.GatewayLookupUrls
	}
	if r.Key.IsSet() && r.Key.Get() != nil {
		response.Key.Set(r.Key.Get())
	}

	return response
}
