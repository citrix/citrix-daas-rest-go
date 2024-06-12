// Copyright © 2024. Citrix Systems, Inc.
package models

// Add-STFRoamingGateway
var _ MappedNullable = &AddSTFRoamingGatewayRequestModel{}

type AddSTFRoamingGatewayRequestModel struct {
	RoamingGateway             NullableString `json:"RoamingGateway,omitempty"`             // The Gateway to add
	Name                       NullableString `json:"Name,omitempty"`                       //The Gateway friendly name
	LogonType                  NullableString `json:"LogonType,omitempty"`                  // The login type required and supported by the Gateway
	SmartCardFallbackLogonType NullableString `json:"SmartCardFallbackLogonType,omitempty"` // The login type to use when SmartCard fails
	Version                    NullableString `json:"Version,omitempty"`                    // The Citrix NetScaler Gateway version
	GatewayUrl                 NullableString `json:"GatewayUrl,omitempty"`                 // The Gateway Url
	CallbackUrl                NullableString `json:"CallbackUrl,omitempty"`                // The Gateway authentication NetScaler call-back Url
	SessionReliability         NullableBool   `json:"SessionReliability,omitempty"`         // Enable session reliability. Session Reliability keeps sessions active and on the user’s screen when network connectivity is interrupted. Users continue to see the application they are using until network connectivity resumes.
	RequestTicketTwoSTAs       NullableBool   `json:"RequestTicketTwoSTAs,omitempty"`       //Request STA tickets from two STA servers (Requires two STA servers)
	SubnetIPAddress            NullableString `json:"SubnetIPAddress,omitempty"`            //IP address
	SecureTicketAuthorityUrls  []string       `json:"SecureTicketAuthorityUrls,omitempty"`  // Secure Ticket Authority server Urls. The Secure Ticket Authority (STA) is responsible for issuing session tickets in response to connection requests for published resources on XenApp. These session tickets form the basis of authentication and authorization for access to published resources.
	StasUseLoadBalancing       NullableBool   `json:"StasUseLoadBalancing,omitempty"`       //Load balance between the configured STA servers (requires two or more STA servers)
	StasBypassDuration         NullableString `json:"StasBypassDuration,omitempty"`         // Time before retrying a failed STA server
	PassThru                   NullableBool   `json:"PassThru,omitempty"`                   //Output the added Gateway
	GslbUrl                    NullableString `json:"GslbUrl,omitempty"`                    // An optional URL which corresponds to the GSLB domain used by multiple gateways
	SecureTicketAuthorityObjs  []string       `json:"SecureTicketAuthorityObjs,omitempty"`  // Secure Ticket Authority objects. The Secure Ticket Authority (STA) is responsible for issuing session tickets in response to connection requests for published resources on XenApp. These session tickets form the basis of authentication and authorization for access to published resources. This new parameter is added to enable the secure sta endpoint feature, and is encouraged to be used, and for backward compatibility we still support the SecureTicketAuthorityUrls parameter.
	IsCloudGateway             NullableBool   `json:"IsCloudGateway,omitempty"`             //Whether the Gateway is an instance of Citrix Gateway Service in the cloud.
	RoamingService             NullableString `json:"RoamingService,omitempty"`             // The Roaming service
}

// ToMap implements MappedNullable.
func (o AddSTFRoamingGatewayRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RoamingGateway.IsSet() {
		toSerialize["RoamingGateway"] = o.RoamingGateway.Get()
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.LogonType.IsSet() {
		toSerialize["LogonType"] = o.LogonType.Get()
	}
	if o.SmartCardFallbackLogonType.IsSet() {
		toSerialize["SmartCardFallbackLogonType"] = o.SmartCardFallbackLogonType.Get()
	}
	if o.Version.IsSet() {
		toSerialize["Version"] = o.Version.Get()
	}
	if o.GatewayUrl.IsSet() {
		toSerialize["GatewayUrl"] = o.GatewayUrl.Get()
	}
	if o.CallbackUrl.IsSet() {
		toSerialize["CallbackUrl"] = o.CallbackUrl.Get()
	}
	if o.SessionReliability.IsSet() {
		toSerialize["SessionReliability"] = o.SessionReliability.Get()
	}
	if o.RequestTicketTwoSTAs.IsSet() {
		toSerialize["RequestTicketTwoSTAs"] = o.RequestTicketTwoSTAs.Get()
	}
	if o.SubnetIPAddress.IsSet() {
		toSerialize["SubnetIPAddress"] = o.SubnetIPAddress.Get()
	}
	if o.SecureTicketAuthorityUrls != nil {
		toSerialize["SecureTicketAuthorityUrls"] = o.SecureTicketAuthorityUrls
	}
	if o.StasUseLoadBalancing.IsSet() {
		toSerialize["StasUseLoadBalancing"] = o.StasUseLoadBalancing.Get()
	}
	if o.StasBypassDuration.IsSet() {
		toSerialize["StasBypassDuration"] = o.StasBypassDuration.Get()
	}
	if o.PassThru.IsSet() {
		toSerialize["PassThru"] = o.PassThru.Get()
	}
	if o.GslbUrl.IsSet() {
		toSerialize["GslbUrl"] = o.GslbUrl.Get()
	}
	if o.SecureTicketAuthorityObjs != nil {
		toSerialize["SecureTicketAuthorityObjs"] = o.SecureTicketAuthorityObjs
	}
	if o.IsCloudGateway.IsSet() {
		toSerialize["IsCloudGateway"] = o.IsCloudGateway.Get()
	}
	if o.RoamingService.IsSet() {
		toSerialize["RoamingService"] = o.RoamingService.Get()
	}
	return toSerialize, nil
}

func (o *AddSTFRoamingGatewayRequestModel) SetRoamingGateway(v string) {
	o.RoamingGateway.Set(&v)
}
func (o *AddSTFRoamingGatewayRequestModel) SetName(v string) {
	o.Name.Set(&v)
}

func (o *AddSTFRoamingGatewayRequestModel) SetLogonType(v string) {
	o.LogonType.Set(&v)
}

func (o *AddSTFRoamingGatewayRequestModel) SetSmartCardFallbackLogonType(v string) {
	o.SmartCardFallbackLogonType.Set(&v)
}

func (o *AddSTFRoamingGatewayRequestModel) SetVersion(v string) {
	o.Version.Set(&v)
}

func (o *AddSTFRoamingGatewayRequestModel) SetGatewayUrl(v string) {
	o.GatewayUrl.Set(&v)
}

func (o *AddSTFRoamingGatewayRequestModel) SetCallbackUrl(v string) {
	o.CallbackUrl.Set(&v)
}

func (o *AddSTFRoamingGatewayRequestModel) SetSessionReliability(v bool) {
	o.SessionReliability.Set(&v)
}

func (o *AddSTFRoamingGatewayRequestModel) SetRequestTicketTwoSTAs(v bool) {
	o.RequestTicketTwoSTAs.Set(&v)
}

func (o *AddSTFRoamingGatewayRequestModel) SetSubnetIPAddress(v string) {
	o.SubnetIPAddress.Set(&v)
}

func (o *AddSTFRoamingGatewayRequestModel) SetStasUseLoadBalancing(v bool) {
	o.StasUseLoadBalancing.Set(&v)
}

func (o *AddSTFRoamingGatewayRequestModel) SetStasBypassDuration(v string) {
	o.StasBypassDuration.Set(&v)
}

func (o *AddSTFRoamingGatewayRequestModel) SetPassThru(v bool) {
	o.PassThru.Set(&v)
}

func (o *AddSTFRoamingGatewayRequestModel) SetGslbUrl(v string) {
	o.GslbUrl.Set(&v)
}

func (o *AddSTFRoamingGatewayRequestModel) SetSecureTicketAuthorityObjs(v []string) {
	o.SecureTicketAuthorityObjs = v
}

func (o *AddSTFRoamingGatewayRequestModel) SetIsCloudGateway(v bool) {
	o.IsCloudGateway.Set(&v)
}

func (o *AddSTFRoamingGatewayRequestModel) SetRoamingService(v string) {
	o.RoamingService.Set(&v)
}

// New-STFRoamingGateway

var _ MappedNullable = &NewSTFRoamingGatewayRequestModel{}

type NewSTFRoamingGatewayRequestModel struct {
	Name                       NullableString `json:"Name,omitempty"`                       // Gateway friendly name
	LogonType                  NullableString `json:"LogonType,omitempty"`                  // The login type required and supported by the Gateway
	SmartCardFallbackLogonType NullableString `json:"SmartCardFallbackLogonType,omitempty"` // The login type to use when SmartCard fails
	Version                    NullableString `json:"Version,omitempty"`                    // The Citrix NetScaler Gateway version
	GatewayUrl                 NullableString `json:"GatewayUrl,omitempty"`                 // The Gateway Url
	CallbackUrl                NullableString `json:"CallbackUrl,omitempty"`                // The Gateway authentication call-back Url
	SessionReliability         NullableBool   `json:"SessionReliability,omitempty"`         // Enable session reliability. Session Reliability keeps sessions active and on the user’s screen when network connectivity is interrupted. Users continue to see the application they are using until network connectivity resumes
	RequestTicketTwoSTAs       NullableBool   `json:"RequestTicketTwoSTAs,omitempty"`       // Request STA tickets from two STA servers (Requires two STA servers)
	SubnetIPAddress            NullableString `json:"SubnetIPAddress,omitempty"`            // IP address
	SecureTicketAuthorityUrls  []string       `json:"SecureTicketAuthorityUrls,omitempty"`  // Secure Ticket Authority server Urls. The Secure Ticket Authority (STA) is responsible for issuing session tickets in response to connection requests for published resources on XenApp. These session tickets form the basis of authentication and authorization for access to published resources
	StasUseLoadBalancing       NullableBool   `json:"StasUseLoadBalancing,omitempty"`       // Load balance between the configured STA servers (requires two or more STA servers)
	StasBypassDuration         NullableString `json:"StasBypassDuration,omitempty"`         // Time before retrying a failed STA server
	GslbUrl                    NullableString `json:"GslbUrl,omitempty"`                    // An optional URL which corresponds to the GSLB domain used by multiple gateways
	SecureTicketAuthorityObjs  []string       `json:"SecureTicketAuthorityObjs,omitempty"`  // 	Secure Ticket Authority objects. The Secure Ticket Authority (STA) is responsible for issuing session tickets in response to connection requests for published resources on XenApp. These session tickets form the basis of authentication and authorization for access to published resources. This new parameter is added to enable the secure sta endpoint feature, and is encouraged to be used, and for backward compatibility we still support the SecureTicketAuthorityUrls parameter.
	IsCloudGateway             NullableBool   `json:"IsCloudGateway,omitempty"`             // Whether the Gateway is an instance of Citrix Gateway Service in the cloud.m
}

// ToMap implements MappedNullable.
func (o NewSTFRoamingGatewayRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.LogonType.IsSet() {
		toSerialize["LogonType"] = o.LogonType.Get()
	}
	if o.SmartCardFallbackLogonType.IsSet() {
		toSerialize["SmartCardFallbackLogonType"] = o.SmartCardFallbackLogonType.Get()
	}
	if o.Version.IsSet() {
		toSerialize["Version"] = o.Version.Get()
	}
	if o.GatewayUrl.IsSet() {
		toSerialize["GatewayUrl"] = o.GatewayUrl.Get()
	}
	if o.CallbackUrl.IsSet() {
		toSerialize["CallbackUrl"] = o.CallbackUrl.Get()
	}
	if o.SessionReliability.IsSet() {
		toSerialize["SessionReliability"] = o.SessionReliability.Get()
	}
	if o.RequestTicketTwoSTAs.IsSet() {
		toSerialize["RequestTicketTwoSTAs"] = o.RequestTicketTwoSTAs.Get()
	}
	if o.SubnetIPAddress.IsSet() {
		toSerialize["SubnetIPAddress"] = o.SubnetIPAddress.Get()
	}
	if len(o.SecureTicketAuthorityUrls) > 0 {
		toSerialize["SecureTicketAuthorityUrls"] = o.SecureTicketAuthorityUrls
	}
	if o.StasUseLoadBalancing.IsSet() {
		toSerialize["StasUseLoadBalancing"] = o.StasUseLoadBalancing.Get()
	}
	if o.StasBypassDuration.IsSet() {
		toSerialize["StasBypassDuration"] = o.StasBypassDuration.Get()
	}
	if o.GslbUrl.IsSet() {
		toSerialize["GslbUrl"] = o.GslbUrl.Get()
	}
	if len(o.SecureTicketAuthorityObjs) > 0 {
		toSerialize["SecureTicketAuthorityObjs"] = o.SecureTicketAuthorityObjs
	}
	if o.IsCloudGateway.IsSet() {
		toSerialize["IsCloudGateway"] = o.IsCloudGateway.Get()
	}
	return toSerialize, nil

}

func (o *NewSTFRoamingGatewayRequestModel) SetName(v string) {
	o.Name.Set(&v)
}
func (o *NewSTFRoamingGatewayRequestModel) SetLogonType(v string) {
	o.LogonType.Set(&v)
}
func (o *NewSTFRoamingGatewayRequestModel) SetSmartCardFallbackLogonType(v string) {
	o.SmartCardFallbackLogonType.Set(&v)
}
func (o *NewSTFRoamingGatewayRequestModel) SetVersion(v string) {
	o.Version.Set(&v)
}
func (o *NewSTFRoamingGatewayRequestModel) SetGatewayUrl(v string) {
	o.GatewayUrl.Set(&v)
}
func (o *NewSTFRoamingGatewayRequestModel) SetCallbackUrl(v string) {
	o.CallbackUrl.Set(&v)
}
func (o *NewSTFRoamingGatewayRequestModel) SetSessionReliability(v bool) {
	o.SessionReliability.Set(&v)
}
func (o *NewSTFRoamingGatewayRequestModel) SetRequestTicketTwoSTAs(v bool) {
	o.RequestTicketTwoSTAs.Set(&v)
}
func (o *NewSTFRoamingGatewayRequestModel) SetSubnetIPAddress(v string) {
	o.SubnetIPAddress.Set(&v)
}
func (o *NewSTFRoamingGatewayRequestModel) SetSecureTicketAuthorityUrls(v []string) {
	o.SecureTicketAuthorityUrls = v
}
func (o *NewSTFRoamingGatewayRequestModel) SetStasUseLoadBalancing(v bool) {
	o.StasUseLoadBalancing.Set(&v)
}
func (o *NewSTFRoamingGatewayRequestModel) SetStasBypassDuration(v string) {
	o.StasBypassDuration.Set(&v)
}
func (o *NewSTFRoamingGatewayRequestModel) SetGslbUrl(v string) {
	o.GslbUrl.Set(&v)
}
func (o *NewSTFRoamingGatewayRequestModel) SetSecureTicketAuthorityObjs(v []string) {
	o.SecureTicketAuthorityObjs = v
}
func (o *NewSTFRoamingGatewayRequestModel) SetIsCloudGateway(v bool) {
	o.IsCloudGateway.Set(&v)
}

// Get-STFRoamingGateway

var _ MappedNullable = &GetSTFRoamingGatewayRequestModel{}

type GetSTFRoamingGatewayRequestModel struct {
	RoamingService NullableString `json:"RoamingGateway,omitempty"` // The Gateway to get
	Name           NullableString `json:"Name,omitempty"`           // The Gateway friendly name
}

// ToMap implements MappedNullable.
func (o GetSTFRoamingGatewayRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RoamingService.IsSet() {
		toSerialize["RoamingGateway"] = o.RoamingService.Get()
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	return toSerialize, nil
}

func (o *GetSTFRoamingGatewayRequestModel) SetRoamingService(v string) {
	o.RoamingService.Set(&v)
}

func (o *GetSTFRoamingGatewayRequestModel) SetName(v string) {
	o.Name.Set(&v)
}

// Remove-STFRoamingGateway

var _ MappedNullable = &RemoveSTFRoamingGatewayRequestModel{}

type RemoveSTFRoamingGatewayRequestModel struct {
	RoamingService NullableString `json:"RoamingGateway,omitempty"` // The Roaming service
	Name           NullableString `json:"Name,omitempty"`           // Remove the Gateway named from the global list
	Gateway        NullableString `json:"Gateway,omitempty"`        // Remove the Gateway supplied from the global list
}

// ToMap implements MappedNullable.
func (o RemoveSTFRoamingGatewayRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RoamingService.IsSet() {
		toSerialize["RoamingGateway"] = o.RoamingService.Get()
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.Gateway.IsSet() {
		toSerialize["Gateway"] = o.Gateway.Get()
	}
	return toSerialize, nil
}

func (o *RemoveSTFRoamingGatewayRequestModel) SetRoamingService(v string) {
	o.RoamingService.Set(&v)
}

func (o *RemoveSTFRoamingGatewayRequestModel) SetName(v string) {
	o.Name.Set(&v)
}

func (o *RemoveSTFRoamingGatewayRequestModel) SetGateway(v string) {
	o.Gateway.Set(&v)
}
