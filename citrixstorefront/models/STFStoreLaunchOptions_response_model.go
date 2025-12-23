// Copyright © 2025. Citrix Systems, Inc.
package models

type GetSTFStoreLaunchOptionsRawResponseModel struct {
	SetNoLoadBiasFlag                      NullableBool   `json:"SetNoLoadBiasFlag,omitempty"`
	AddressResolutionType                  NullableInt64  `json:"AddressResolutionType,omitempty"`
	RequestICAClientSecureChannel          NullableInt64  `json:"RequestICAClientSecureChannel,omitempty"`
	IgnoreClientProvidedClientAddress      NullableBool   `json:"IgnoreClientProvidedClientAddress,omitempty"`
	OverlayAutoLoginCredentialsWithTicket  NullableBool   `json:"OverlayAutoLoginCredentialsWithTicket,omitempty"`
	OverrideIcaClientName                  NullableBool   `json:"OverrideIcaClientName,omitempty"`
	RequireLaunchReference                 NullableBool   `json:"RequireLaunchReference,omitempty"`
	AllowFontSmoothing                     NullableBool   `json:"AllowFontSmoothing,omitempty"`
	ShowDesktopViewer                      NullableBool   `json:"ShowDesktopViewer,omitempty"`
	AllowSpecialFolderRedirection          NullableBool   `json:"AllowSpecialFolderRedirection,omitempty"`
	ClientProxyPolicy                      any            `json:"ClientProxyPolicy,omitempty"` //TODO: Placeholder for ClientProxyPolicy
	RoutingPolicy                          any            `json:"RoutingPolicy,omitempty"`     //TODO: Placeholder for RoutingPolicy
	VdaLogonDataProviderName               NullableString `json:"VdaLogonDataProviderName,omitempty"`
	FederatedAuthenticationServiceFailover NullableBool   `json:"FederatedAuthenticationServiceFailover,omitempty"`
}

type GetSTFStoreLaunchOptionsResponseModel struct {
	SetNoLoadBiasFlag                      NullableBool   `json:"SetNoLoadBiasFlag,omitempty"`
	AddressResolutionType                  string         `json:"AddressResolutionType,omitempty"`
	RequestICAClientSecureChannel          string         `json:"RequestICAClientSecureChannel,omitempty"`
	IgnoreClientProvidedClientAddress      NullableBool   `json:"IgnoreClientProvidedClientAddress,omitempty"`
	OverlayAutoLoginCredentialsWithTicket  NullableBool   `json:"OverlayAutoLoginCredentialsWithTicket,omitempty"`
	OverrideIcaClientName                  NullableBool   `json:"OverrideIcaClientName,omitempty"`
	RequireLaunchReference                 NullableBool   `json:"RequireLaunchReference,omitempty"`
	AllowFontSmoothing                     NullableBool   `json:"AllowFontSmoothing,omitempty"`
	ShowDesktopViewer                      NullableBool   `json:"ShowDesktopViewer,omitempty"`
	AllowSpecialFolderRedirection          NullableBool   `json:"AllowSpecialFolderRedirection,omitempty"`
	VdaLogonDataProviderName               NullableString `json:"VdaLogonDataProviderName,omitempty"`
	FederatedAuthenticationServiceFailover NullableBool   `json:"FederatedAuthenticationServiceFailover,omitempty"`
}

func AddressResolutionTypeToString(addressResolutionType int64) string {
	switch addressResolutionType {
	case 0:
		return "Dns"
	case 2:
		return "IPV4"
	case 3:
		return "IPV4Port"
	case 4:
		return "Dot"
	case 5:
		return "DotPort"
	case 6:
		return "Uri"
	case 7:
		return "NoChange"
	default:
		return "DnsPort"
	}
}

func RequestICAClientSecureChannelToString(requestICAClientSecureChannel int64) string {
	switch requestICAClientSecureChannel {
	case 0:
		return "SSLAnyCiphers"
	case 1:
		return "TLSGovCipers"
	default:
		return "DetectAnyCiphers"
	}
}

func (rawResponse GetSTFStoreLaunchOptionsRawResponseModel) ConvertToResponseModel() GetSTFStoreLaunchOptionsResponseModel {
	response := GetSTFStoreLaunchOptionsResponseModel{}

	response.SetNoLoadBiasFlag = rawResponse.SetNoLoadBiasFlag
	response.AddressResolutionType = AddressResolutionTypeToString(*rawResponse.AddressResolutionType.Get())
	response.RequestICAClientSecureChannel = RequestICAClientSecureChannelToString(*rawResponse.RequestICAClientSecureChannel.Get())
	response.IgnoreClientProvidedClientAddress = rawResponse.IgnoreClientProvidedClientAddress
	response.OverlayAutoLoginCredentialsWithTicket = rawResponse.OverlayAutoLoginCredentialsWithTicket
	response.OverrideIcaClientName = rawResponse.OverrideIcaClientName
	response.RequireLaunchReference = rawResponse.RequireLaunchReference
	response.AllowFontSmoothing = rawResponse.AllowFontSmoothing
	response.ShowDesktopViewer = rawResponse.ShowDesktopViewer
	response.AllowSpecialFolderRedirection = rawResponse.AllowSpecialFolderRedirection
	response.VdaLogonDataProviderName = rawResponse.VdaLogonDataProviderName
	response.FederatedAuthenticationServiceFailover = rawResponse.FederatedAuthenticationServiceFailover

	return response
}
