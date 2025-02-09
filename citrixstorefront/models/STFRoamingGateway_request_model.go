// Copyright © 2024. Citrix Systems, Inc.
package models

// Add-STFRoamingGateway
var _ MappedNullable = &AddSTFRoamingGatewayRequestModel{}

type AddSTFRoamingGatewayRequestModel struct {
	Name                       NullableString `json:"Name,omitempty"`                       //The Gateway friendly name
	LogonType                  *LogonType     `json:"LogonType,omitempty"`                  // The login type required and supported by the Gateway
	SmartCardFallbackLogonType *LogonType     `json:"SmartCardFallbackLogonType,omitempty"` // The login type to use when SmartCard fails
	Version                    NullableString `json:"Version,omitempty"`                    // The Citrix NetScaler Gateway version
	GatewayUrl                 NullableString `json:"GatewayUrl,omitempty"`                 // The Gateway Url
	CallbackUrl                NullableString `json:"CallbackUrl,omitempty"`                // The Gateway authentication NetScaler call-back Url
	SubnetIPAddress            NullableString `json:"SubnetIPAddress,omitempty"`            //IP address
	StasBypassDuration         NullableString `json:"StasBypassDuration,omitempty"`         // Time before retrying a failed STA server
	GslbUrl                    NullableString `json:"GslbUrl,omitempty"`                    // An optional URL which corresponds to the GSLB domain used by multiple gateways
	IsCloudGateway             NullableBool   `json:"IsCloudGateway,omitempty"`             //Whether the Gateway is an instance of Citrix Gateway Service in the cloud.
	SessionReliability         NullableBool   `json:"SessionReliability,omitempty"`         // Enable session reliability. Session Reliability keeps sessions active and on the user’s screen when network connectivity is interrupted. Users continue to see the application they are using until network connectivity resumes.
	RequestTicketTwoSTAs       NullableBool   `json:"RequestTicketTwoSTAs,omitempty"`       //Request STA tickets from two STA servers (Requires two STA servers)
	StasUseLoadBalancing       NullableBool   `json:"StasUseLoadBalancings,omitempty"`      //Use load balancing for STAs
}

// ToMap implements MappedNullable.
func (o AddSTFRoamingGatewayRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if !IsNil(o.LogonType) {
		toSerialize["LogonType"] = o.LogonType
	}
	if !IsNil(o.SmartCardFallbackLogonType) {
		toSerialize["SmartCardFallbackLogonType"] = o.SmartCardFallbackLogonType
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
	if o.SubnetIPAddress.IsSet() {
		toSerialize["SubnetIPAddress"] = o.SubnetIPAddress.Get()
	}
	if o.StasBypassDuration.IsSet() {
		toSerialize["StasBypassDuration"] = o.StasBypassDuration.Get()
	}
	if o.GslbUrl.IsSet() {
		toSerialize["GslbUrl"] = o.GslbUrl.Get()
	}
	if o.IsCloudGateway.IsSet() {
		toSerialize["IsCloudGateway"] = o.IsCloudGateway.Get()
	}
	if o.SessionReliability.IsSet() {
		toSerialize["SessionReliability"] = o.SessionReliability.Get()
	}
	if o.RequestTicketTwoSTAs.IsSet() {
		toSerialize["RequestTicketTwoSTAs"] = o.RequestTicketTwoSTAs.Get()
	}
	if o.StasUseLoadBalancing.IsSet() {
		toSerialize["StasUseLoadBalancings"] = o.StasUseLoadBalancing.Get()
	}
	return toSerialize, nil
}

func (o *AddSTFRoamingGatewayRequestModel) SetName(v string) {
	o.Name.Set(&v)
}

func (o *AddSTFRoamingGatewayRequestModel) SetLogonType(v LogonType) {
	o.LogonType = &v
}

func (o *AddSTFRoamingGatewayRequestModel) SetSmartCardFallbackLogonType(v LogonType) {
	o.SmartCardFallbackLogonType = &v
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

func (o *AddSTFRoamingGatewayRequestModel) SetSubnetIPAddress(v string) {
	o.SubnetIPAddress.Set(&v)
}

func (o *AddSTFRoamingGatewayRequestModel) SetStasBypassDuration(v string) {
	o.StasBypassDuration.Set(&v)
}

func (o *AddSTFRoamingGatewayRequestModel) SetGslbUrl(v string) {
	o.GslbUrl.Set(&v)
}

func (o *AddSTFRoamingGatewayRequestModel) SetIsCloudGateway(v bool) {
	o.IsCloudGateway.Set(&v)
}

func (o *AddSTFRoamingGatewayRequestModel) SetSessionReliability(v bool) {
	o.SessionReliability.Set(&v)
}

func (o *AddSTFRoamingGatewayRequestModel) SetRequestTicketTwoSTAs(v bool) {
	o.RequestTicketTwoSTAs.Set(&v)
}

func (o *AddSTFRoamingGatewayRequestModel) SetStasUseLoadBalancing(v bool) {
	o.StasUseLoadBalancing.Set(&v)
}

// GetLogonType returns the LogonType field value if set, zero value otherwise.
func (o *AddSTFRoamingGatewayRequestModel) GetLogonType() LogonType {
	if o == nil || IsNil(o.LogonType) {
		var ret LogonType
		return ret
	}
	return *o.LogonType
}

// GetLogonTypeOk returns a tuple with the LogonType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSTFRoamingGatewayRequestModel) GetLogonTypeOk() (*LogonType, bool) {
	if o == nil || IsNil(o.LogonType) {
		return nil, false
	}
	return o.LogonType, true
}

// HasLogonType returns a boolean if a field has been set.
func (o *AddSTFRoamingGatewayRequestModel) HasLogonType() bool {
	if o != nil && !IsNil(o.LogonType) {
		return true
	}

	return false
}

// GetSmartCardFallbackLogonType returns the LogonType field value if set, zero value otherwise.
func (o *AddSTFRoamingGatewayRequestModel) GetSmartCardFallbackLogonType() LogonType {
	if o == nil || IsNil(o.LogonType) {
		var ret LogonType
		return ret
	}
	return *o.LogonType
}

// GetSmartCardFallbackLogonTypeOk returns a tuple with the LogonType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSTFRoamingGatewayRequestModel) GetSmartCardFallbackLogonTypeOk() (*LogonType, bool) {
	if o == nil || IsNil(o.LogonType) {
		return nil, false
	}
	return o.LogonType, true
}

// HasSmartCardFallbackLogonType returns a boolean if a field has been set.
func (o *AddSTFRoamingGatewayRequestModel) HasSmartCardFallbackLogonType() bool {
	if o != nil && !IsNil(o.LogonType) {
		return true
	}

	return false
}

// Get-STFRoamingGateway
var _ MappedNullable = &GetSTFRoamingGatewayRequestModel{}

type GetSTFRoamingGatewayRequestModel struct {
	Name NullableString `json:"Name,omitempty"` // The Gateway friendly name
}

// ToMap implements MappedNullable.
func (o GetSTFRoamingGatewayRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	return toSerialize, nil
}

func (o *GetSTFRoamingGatewayRequestModel) SetName(v string) {
	o.Name.Set(&v)
}

// Set-STFRoamingGateway
var _ MappedNullable = &SetSTFRoamingGatewayRequestModel{}

type SetSTFRoamingGatewayRequestModel struct {
	Name                       NullableString `json:"Name,omitempty"`                       //The Gateway friendly name
	LogonType                  *LogonType     `json:"LogonType,omitempty"`                  // The login type required and supported by the Gateway
	SmartCardFallbackLogonType *LogonType     `json:"SmartCardFallbackLogonType,omitempty"` // The login type to use when SmartCard fails
	Version                    NullableString `json:"Version,omitempty"`                    // The Citrix NetScaler Gateway version
	GatewayUrl                 NullableString `json:"GatewayUrl,omitempty"`                 // The Gateway Url
	CallbackUrl                NullableString `json:"CallbackUrl,omitempty"`                // The Gateway authentication NetScaler call-back Url
	SessionReliability         NullableBool   `json:"SessionReliability,omitempty"`         // Enable session reliability. Session Reliability keeps sessions active and on the user’s screen when network connectivity is interrupted. Users continue to see the application they are using until network connectivity resumes.
	RequestTicketTwoSTAs       NullableBool   `json:"RequestTicketTwoSTAs,omitempty"`       //Request STA tickets from two STA servers (Requires two STA servers)
	SubnetIPAddress            NullableString `json:"SubnetIPAddress,omitempty"`            //IP address
	GslbUrl                    NullableString `json:"GslbUrl,omitempty"`                    // An optional URL which corresponds to the GSLB domain used by multiple gateways
	IsCloudGateway             NullableBool   `json:"IsCloudGateway,omitempty"`             //Whether the Gateway is an instance of Citrix Gateway Service in the cloud.
}

// ToMap implements MappedNullable.
func (o SetSTFRoamingGatewayRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if !IsNil(o.LogonType) {
		toSerialize["LogonType"] = o.LogonType
	}
	if !IsNil(o.SmartCardFallbackLogonType) {
		toSerialize["SmartCardFallbackLogonType"] = o.SmartCardFallbackLogonType
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
	if o.GslbUrl.IsSet() {
		toSerialize["GslbUrl"] = o.GslbUrl.Get()
	}
	if o.IsCloudGateway.IsSet() {
		toSerialize["IsCloudGateway"] = o.IsCloudGateway.Get()
	}
	return toSerialize, nil
}

func (o *SetSTFRoamingGatewayRequestModel) SetName(v string) {
	o.Name.Set(&v)
}

func (o *SetSTFRoamingGatewayRequestModel) SetLogonType(v LogonType) {
	o.LogonType = &v
}

func (o *SetSTFRoamingGatewayRequestModel) SetSmartCardFallbackLogonType(v LogonType) {
	o.SmartCardFallbackLogonType = &v
}

func (o *SetSTFRoamingGatewayRequestModel) SetVersion(v string) {
	o.Version.Set(&v)
}

func (o *SetSTFRoamingGatewayRequestModel) SetGatewayUrl(v string) {
	o.GatewayUrl.Set(&v)
}

func (o *SetSTFRoamingGatewayRequestModel) SetCallbackUrl(v string) {
	o.CallbackUrl.Set(&v)
}

func (o *SetSTFRoamingGatewayRequestModel) SetSessionReliability(v bool) {
	o.SessionReliability.Set(&v)
}

func (o *SetSTFRoamingGatewayRequestModel) SetRequestTicketTwoSTAs(v bool) {
	o.RequestTicketTwoSTAs.Set(&v)
}

func (o *SetSTFRoamingGatewayRequestModel) SetSubnetIPAddress(v string) {
	o.SubnetIPAddress.Set(&v)
}

func (o *SetSTFRoamingGatewayRequestModel) SetGslbUrl(v string) {
	o.GslbUrl.Set(&v)
}

func (o *SetSTFRoamingGatewayRequestModel) SetIsCloudGateway(v bool) {
	o.IsCloudGateway.Set(&v)
}

// GetLogonType returns the LogonType field value if set, zero value otherwise.
func (o *SetSTFRoamingGatewayRequestModel) GetLogonType() LogonType {
	if o == nil || IsNil(o.LogonType) {
		var ret LogonType
		return ret
	}
	return *o.LogonType
}

// GetLogonTypeOk returns a tuple with the LogonType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetSTFRoamingGatewayRequestModel) GetLogonTypeOk() (*LogonType, bool) {
	if o == nil || IsNil(o.LogonType) {
		return nil, false
	}
	return o.LogonType, true
}

// HasLogonType returns a boolean if a field has been set.
func (o *SetSTFRoamingGatewayRequestModel) HasLogonType() bool {
	if o != nil && !IsNil(o.LogonType) {
		return true
	}

	return false
}

// GetSmartCardFallbackLogonType returns the LogonType field value if set, zero value otherwise.
func (o *SetSTFRoamingGatewayRequestModel) GetSmartCardFallbackLogonType() LogonType {
	if o == nil || IsNil(o.LogonType) {
		var ret LogonType
		return ret
	}
	return *o.LogonType
}

// GetSmartCardFallbackLogonTypeOk returns a tuple with the LogonType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetSTFRoamingGatewayRequestModel) GetSmartCardFallbackLogonTypeOk() (*LogonType, bool) {
	if o == nil || IsNil(o.LogonType) {
		return nil, false
	}
	return o.LogonType, true
}

// HasSmartCardFallbackLogonType returns a boolean if a field has been set.
func (o *SetSTFRoamingGatewayRequestModel) HasSmartCardFallbackLogonType() bool {
	if o != nil && !IsNil(o.LogonType) {
		return true
	}

	return false
}
