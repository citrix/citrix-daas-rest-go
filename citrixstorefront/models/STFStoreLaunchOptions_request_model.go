// Copyright © 2024. Citrix Systems, Inc.
package models

type SetSTFStoreLaunchOptionsRequestModel struct {
	AddressResolutionType                  NullableString `json:"AddressResolutionType,omitempty"`
	RequestICAClientSecureChannel          NullableString `json:"RequestIcaClientSecureChannel,omitempty"`
	AllowSpecialFolderRedirection          NullableBool   `json:"AllowSpecialFolderRedirection,omitempty"`
	AllowFontSmoothing                     NullableBool   `json:"AllowFontSmoothing,omitempty"`
	RequireLaunchReference                 NullableBool   `json:"RequireLaunchReference,omitempty"`
	OverrideIcaClientName                  NullableBool   `json:"OverrideIcaClientName,omitempty"`
	OverlayAutoLoginCredentialsWithTicket  NullableBool   `json:"OverlayAutoLoginCredentialsWithTicket,omitempty"`
	IgnoreClientProvidedClientAddress      NullableBool   `json:"IgnoreClientProvidedClientAddress,omitempty"`
	SetNoLoadBiasFlag                      NullableBool   `json:"SetNoLoadBiasFlag,omitempty"`
	RDPOnly                                NullableBool   `json:"RDPOnly,omitempty"`
	VdaLogonDataProvider                   NullableString `json:"VdaLogonDataProvider,omitempty"`
	IcaTemplateName                        NullableString `json:"IcaTemplateName,omitempty"`
	FederatedAuthenticationServiceFailover NullableBool   `json:"FederatedAuthenticationServiceFailover,omitempty"`
}

// ToMap implements MappedNullable
func (o SetSTFStoreLaunchOptionsRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}

	if o.AddressResolutionType.IsSet() {
		toSerialize["AddressResolutionType"] = o.AddressResolutionType.Get()
	}
	if o.RequestICAClientSecureChannel.IsSet() {
		toSerialize["RequestIcaClientSecureChannel"] = o.RequestICAClientSecureChannel.Get()
	}
	if o.AllowSpecialFolderRedirection.IsSet() {
		toSerialize["AllowSpecialFolderRedirection"] = o.AllowSpecialFolderRedirection.Get()
	}
	if o.AllowFontSmoothing.IsSet() {
		toSerialize["AllowFontSmoothing"] = o.AllowFontSmoothing.Get()
	}
	if o.RequireLaunchReference.IsSet() {
		toSerialize["RequireLaunchReference"] = o.RequireLaunchReference.Get()
	}
	if o.OverrideIcaClientName.IsSet() {
		toSerialize["OverrideIcaClientName"] = o.OverrideIcaClientName.Get()
	}
	if o.OverlayAutoLoginCredentialsWithTicket.IsSet() {
		toSerialize["OverlayAutoLoginCredentialsWithTicket"] = o.OverlayAutoLoginCredentialsWithTicket.Get()
	}
	if o.IgnoreClientProvidedClientAddress.IsSet() {
		toSerialize["IgnoreClientProvidedClientAddress"] = o.IgnoreClientProvidedClientAddress.Get()
	}
	if o.SetNoLoadBiasFlag.IsSet() {
		toSerialize["SetNoLoadBiasFlag"] = o.SetNoLoadBiasFlag.Get()
	}
	if o.RDPOnly.IsSet() {
		toSerialize["RDPOnly"] = o.RDPOnly.Get()
	}
	if o.VdaLogonDataProvider.IsSet() {
		toSerialize["VdaLogonDataProvider"] = o.VdaLogonDataProvider.Get()
	}
	if o.IcaTemplateName.IsSet() {
		toSerialize["IcaTemplateName"] = o.IcaTemplateName.Get()
	}
	if o.FederatedAuthenticationServiceFailover.IsSet() {
		toSerialize["FederatedAuthenticationServiceFailover"] = o.FederatedAuthenticationServiceFailover.Get()
	}
	return toSerialize, nil
}

func (o *SetSTFStoreLaunchOptionsRequestModel) SetAddressResolutionType(v string) {
	o.AddressResolutionType.Set(&v)
}

func (o *SetSTFStoreLaunchOptionsRequestModel) SetRequestICAClientSecureChannel(v string) {
	o.RequestICAClientSecureChannel.Set(&v)
}

func (o *SetSTFStoreLaunchOptionsRequestModel) SetAllowSpecialFolderRedirection(v bool) {
	o.AllowSpecialFolderRedirection.Set(&v)
}

func (o *SetSTFStoreLaunchOptionsRequestModel) SetAllowFontSmoothing(v bool) {
	o.AllowFontSmoothing.Set(&v)
}

func (o *SetSTFStoreLaunchOptionsRequestModel) SetRequireLaunchReference(v bool) {
	o.RequireLaunchReference.Set(&v)
}

func (o *SetSTFStoreLaunchOptionsRequestModel) SetOverrideIcaClientName(v bool) {
	o.OverrideIcaClientName.Set(&v)
}

func (o *SetSTFStoreLaunchOptionsRequestModel) SetOverlayAutoLoginCredentialsWithTicket(v bool) {
	o.OverlayAutoLoginCredentialsWithTicket.Set(&v)
}

func (o *SetSTFStoreLaunchOptionsRequestModel) SetIgnoreClientProvidedClientAddress(v bool) {
	o.IgnoreClientProvidedClientAddress.Set(&v)
}

func (o *SetSTFStoreLaunchOptionsRequestModel) SetSetNoLoadBiasFlag(v bool) {
	o.SetNoLoadBiasFlag.Set(&v)
}

func (o *SetSTFStoreLaunchOptionsRequestModel) SetRDPOnly(v bool) {
	o.RDPOnly.Set(&v)
}

func (o *SetSTFStoreLaunchOptionsRequestModel) SetVdaLogonDataProvider(v string) {
	o.VdaLogonDataProvider.Set(&v)
}

func (o *SetSTFStoreLaunchOptionsRequestModel) SetIcaTemplateName(v string) {
	o.IcaTemplateName.Set(&v)
}

func (o *SetSTFStoreLaunchOptionsRequestModel) SetFederatedAuthenticationServiceFailover(v bool) {
	o.FederatedAuthenticationServiceFailover.Set(&v)
}
