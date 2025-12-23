// Copyright © 2025. Citrix Systems, Inc.
package models

// Create Model
var _ MappedNullable = &CreateSTFStoreRequestModel{}

type CreateSTFStoreRequestModel struct {
	FarmName                NullableString
	VirtualPath             NullableString
	SiteId                  NullableInt64
	Anonymous               NullableBool
	FarmType                NullableString
	Servers                 []string
	ServiceUrls             NullableString
	Port                    NullableInt64
	TransportType           NullableString
	SSLRelayPort            NullableInt64
	LoadBalance             NullableBool
	AllFailedBypassDuration NullableInt64
	BypassDuration          NullableInt64
	FriendlyName            NullableString
	Zones                   NullableString
	ConnectionString        NullableString
	XMLValidationEnabled    NullableBool
	XMLValidationSecret     NullableString
}

// ToMap implements MappedNullable.
func (o CreateSTFStoreRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SiteId.IsSet() {
		toSerialize["SiteId"] = o.SiteId.Get()
	}
	if o.VirtualPath.IsSet() {
		toSerialize["VirtualPath"] = o.VirtualPath.Get()
	}
	if o.Anonymous.IsSet() {
		toSerialize["Anonymous"] = o.Anonymous.Get()
	}
	if o.FarmName.IsSet() {
		toSerialize["FarmName"] = o.FarmName.Get()
	}
	if o.FarmType.IsSet() {
		toSerialize["FarmType"] = o.FarmType.Get()
	}
	if o.Servers != nil {
		toSerialize["Servers"] = o.Servers
	}
	if o.ServiceUrls.IsSet() {
		toSerialize["ServiceUrls"] = o.ServiceUrls
	}
	if o.Port.IsSet() {
		toSerialize["Port"] = o.Port.Get()
	}
	if o.TransportType.IsSet() {
		toSerialize["TransportType"] = o.TransportType.Get()
	}
	if o.SSLRelayPort.IsSet() {
		toSerialize["SSLRelayPort"] = o.SSLRelayPort.Get()
	}
	if o.LoadBalance.IsSet() {
		toSerialize["LoadBalance"] = o.LoadBalance.Get()
	}
	if o.AllFailedBypassDuration.IsSet() {
		toSerialize["AllFailedBypassDuration"] = o.AllFailedBypassDuration.Get()
	}
	if o.BypassDuration.IsSet() {
		toSerialize["BypassDuration"] = o.BypassDuration.Get()
	}
	if o.FriendlyName.IsSet() {
		toSerialize["FriendlyName"] = o.FriendlyName.Get()
	}
	if o.Zones.IsSet() {
		toSerialize["Zones"] = o.Zones.Get()
	}
	if o.ConnectionString.IsSet() {
		toSerialize["ConnectionString"] = o.ConnectionString.Get()
	}
	if o.XMLValidationEnabled.IsSet() {
		toSerialize["XMLValidationEnabled"] = o.XMLValidationEnabled.Get()
	}
	if o.XMLValidationSecret.IsSet() {
		toSerialize["XMLValidationSecret"] = o.XMLValidationSecret.Get()
	}

	return toSerialize, nil
}

func (o *CreateSTFStoreRequestModel) SetVirtualPath(v string) {
	o.VirtualPath.Set(&v)
}

func (o *CreateSTFStoreRequestModel) SetSiteId(v int64) {
	o.SiteId.Set(&v)
}

func (o *CreateSTFStoreRequestModel) SetAnonymous(v bool) {
	o.Anonymous.Set(&v)
}

func (o *CreateSTFStoreRequestModel) SetFarmType(v string) {
	o.FarmType.Set(&v)
}

func (o *CreateSTFStoreRequestModel) SetFarmName(v string) {
	o.FarmName.Set(&v)
}

func (o *CreateSTFStoreRequestModel) SetServers(v []string) {
	o.Servers = v
}

func (o *CreateSTFStoreRequestModel) SetServiceUrls(v string) {
	o.ServiceUrls.Set(&v)
}

func (o *CreateSTFStoreRequestModel) SetPort(v int64) {
	o.Port.Set(&v)
}

func (o *CreateSTFStoreRequestModel) SetTransportType(v string) {
	o.TransportType.Set(&v)
}

func (o *CreateSTFStoreRequestModel) SetSSLRelayPort(v int64) {
	o.SSLRelayPort.Set(&v)
}

func (o *CreateSTFStoreRequestModel) SetLoadBalance(v bool) {
	o.LoadBalance.Set(&v)
}

func (o *CreateSTFStoreRequestModel) SetAllFailedBypassDuration(v int64) {
	o.AllFailedBypassDuration.Set(&v)
}

func (o *CreateSTFStoreRequestModel) SetBypassDuration(v int64) {
	o.BypassDuration.Set(&v)
}

func (o *CreateSTFStoreRequestModel) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}

func (o *CreateSTFStoreRequestModel) SetZones(v string) {
	o.Zones.Set(&v)
}

func (o *CreateSTFStoreRequestModel) SetConnectionString(v string) {
	o.ConnectionString.Set(&v)
}

func (o *CreateSTFStoreRequestModel) SetXMLValidationEnabled(v bool) {
	o.XMLValidationEnabled.Set(&v)
}

func (o *CreateSTFStoreRequestModel) SetXMLValidationSecret(v string) {
	o.XMLValidationSecret.Set(&v)
}

// Get Model
var _ MappedNullable = &GetSTFStoreRequestModel{}

type GetSTFStoreRequestModel struct {
	SiteId                NullableInt64  // The IIS site id of the deployment
	VirtualPath           NullableString //Virtual path of the Store to return
	AuthenticationService NullableString //Stores used by an Authentication servic
}

// ToMap implements MappedNullable.
func (o GetSTFStoreRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SiteId.IsSet() {
		toSerialize["SiteId"] = o.SiteId.Get()
	}
	if o.VirtualPath.IsSet() {
		toSerialize["VirtualPath"] = o.VirtualPath.Get()
	}
	if o.AuthenticationService.IsSet() {
		toSerialize["AuthenticationService"] = o.AuthenticationService.Get()
	}
	return toSerialize, nil
}

func (o *GetSTFStoreRequestModel) SetSiteId(v int64) {
	o.SiteId.Set(&v)
}

func (o *GetSTFStoreRequestModel) SetVirtualPath(v string) {
	o.VirtualPath.Set(&v)
}

func (o *GetSTFStoreRequestModel) SetGetSTFStoreRequestModel(v string) {
	o.AuthenticationService.Set(&v)
}

// Set Model
var _ MappedNullable = &GetSTFStoreRequestModel{}

type SetSTFStoreRequestModel struct {
	IcaFileSigning                      NullableBool   //Sign the ica file used for launch.
	IcaFileSigningCertificateThumbprint NullableString //Thumbprint of the certificate to use for signing the ica file.
	IcaFileSigningHashAlgorithm         NullableString //Algorithm used to sign the ica file. The setting should not normally require changing from the default.
	AuthenticationService               NullableString //The Authentication service and the Store should use. Will result in the Store using the authentication protocols of the service.
	KerberosDelegation                  NullableBool   //Turn on Kerberos delegation for the Store to support XenApp 6.5 and earlier constrained delegation.
	Force                               NullableBool   //Force settings to be applied ignoring confirmation prompts.
	AllowSessionReconnect               NullableBool   //Allow sessions to be reconnected.
	SubstituteDesktopImage              NullableBool   //Substitute the desktop image provided by XenApp\XenDesktop with the StoreFront default.
	UnifiedReceiver                     NullableString //Choose the classic or unified Receiver experience.
	LockedDown                          NullableBool   //All the resources delivered by locked-down Store are auto subscribed and do not allow for un-subscription.
	PassThru                            NullableBool   //Output the resulting Store service.
}

// ToMap implements MappedNullable.
func (o SetSTFStoreRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}

	if o.IcaFileSigning.IsSet() {
		toSerialize["IcaFileSigning"] = o.IcaFileSigning.Get()
	}
	if o.IcaFileSigningCertificateThumbprint.IsSet() {
		toSerialize["IcaFileSigningCertificateThumbprint"] = o.IcaFileSigningCertificateThumbprint.Get()
	}
	if o.IcaFileSigningHashAlgorithm.IsSet() {
		toSerialize["IcaFileSigningHashAlgorithm"] = o.IcaFileSigningHashAlgorithm.Get()
	}
	if o.AuthenticationService.IsSet() {
		toSerialize["AuthenticationService"] = o.AuthenticationService.Get()
	}
	if o.KerberosDelegation.IsSet() {
		toSerialize["KerberosDelegation"] = o.KerberosDelegation.Get()
	}
	if o.Force.IsSet() {
		toSerialize["Force"] = o.Force.Get()
	}
	if o.AllowSessionReconnect.IsSet() {
		toSerialize["AllowSessionReconnect"] = o.AllowSessionReconnect.Get()
	}
	if o.SubstituteDesktopImage.IsSet() {
		toSerialize["SubstituteDesktopImage"] = o.SubstituteDesktopImage.Get()
	}
	if o.UnifiedReceiver.IsSet() {
		toSerialize["UnifiedReceiver"] = o.UnifiedReceiver.Get()
	}
	if o.LockedDown.IsSet() {
		toSerialize["LockedDown"] = o.LockedDown.Get()
	}
	if o.PassThru.IsSet() {
		toSerialize["PassThru"] = o.PassThru.Get()
	}
	return toSerialize, nil
}

// SetIcaFileSigning enables or disables signing the ICA file used for launch.
func (o *SetSTFStoreRequestModel) SetIcaFileSigning(v bool) {
	o.IcaFileSigning.Set(&v)
}

// SetIcaFileSigningCertificateThumbprint sets the thumbprint of the certificate to use for signing the ICA file.
func (o *SetSTFStoreRequestModel) SetIcaFileSigningCertificateThumbprint(v string) {
	o.IcaFileSigningCertificateThumbprint.Set(&v)
}

// SetIcaFileSigningHashAlgorithm sets the algorithm used to sign the ICA file.
func (o *SetSTFStoreRequestModel) SetIcaFileSigningHashAlgorithm(v string) {
	o.IcaFileSigningHashAlgorithm.Set(&v)
}

// SetAuthenticationService sets the authentication service that the Store should use.
func (o *SetSTFStoreRequestModel) SetAuthenticationService(v string) {
	o.AuthenticationService.Set(&v)
}

// SetKerberosDelegation enables or disables Kerberos delegation for the Store.
func (o *SetSTFStoreRequestModel) SetKerberosDelegation(v bool) {
	o.KerberosDelegation.Set(&v)
}

// SetForce sets whether to force settings to be applied, ignoring confirmation prompts.
func (o *SetSTFStoreRequestModel) SetForce(v bool) {
	o.Force.Set(&v)
}

// SetAllowSessionReconnect sets whether to allow sessions to be reconnected.
func (o *SetSTFStoreRequestModel) SetAllowSessionReconnect(v bool) {
	o.AllowSessionReconnect.Set(&v)
}

// SetSubstituteDesktopImage sets whether to substitute the desktop image provided by XenApp\XenDesktop with the StoreFront default.
func (o *SetSTFStoreRequestModel) SetSubstituteDesktopImage(v bool) {
	o.SubstituteDesktopImage.Set(&v)
}

// SetUnifiedReceiver sets the Receiver experience to choose (classic or unified).
func (o *SetSTFStoreRequestModel) SetUnifiedReceiver(v string) {
	o.UnifiedReceiver.Set(&v)
}

// SetLockedDown sets whether all the resources delivered by the locked-down Store are auto-subscribed and do not allow un-subscription.
func (o *SetSTFStoreRequestModel) SetLockedDown(v bool) {
	o.LockedDown.Set(&v)
}

// SetPassThru sets whether to output the resulting Store service.
func (o *SetSTFStoreRequestModel) SetPassThru(v bool) {
	o.PassThru.Set(&v)
}
