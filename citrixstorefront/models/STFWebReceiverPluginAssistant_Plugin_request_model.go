// Copyright © 2024. Citrix Systems, Inc.
package models

//Get Model
var _ MappedNullable = &GetSTFWebReceiverPluginAssistantRequestModel{}

type GetSTFWebReceiverPluginAssistantRequestModel struct {
	WebReceiverService NullableString `json:"WebReceiverService,omitempty"`
}

// ToMap implements MappedNullable.
func (o GetSTFWebReceiverPluginAssistantRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.WebReceiverService.IsSet() {
		toSerialize["WebReceiverService"] = o.WebReceiverService.Get()
	}
	return toSerialize, nil
}

func (o *GetSTFWebReceiverPluginAssistantRequestModel) SetWebReceiverService(v string) {
	o.WebReceiverService.Set(&v)
}

// Update Model
var _ MappedNullable = &UpdateSTFWebReceiverPluginAssistantRequestModel{}

type UpdateSTFWebReceiverPluginAssistantRequestModel struct {
	WebReceiverService                            NullableString `json:"WebReceiverService,omitempty"`                            // The WebReceiver service.
	Enabled                                       NullableBool   `json:"Enabled,omitempty"`                                       // Enable Receiver client detection.
	UpgradeAtLogin                                NullableBool   `json:"UpgradeAtLogin,omitempty"`                                // Prompt to upgrade older clients.
	ShowAfterLogin                                NullableBool   `json:"ShowAfterLogin,omitempty"`                                // Show Receiver client detection after the user logs in.
	Win32Path                                     NullableString `json:"Win32Path,omitempty"`                                     // Path to the Windows Receiver.
	MacOSPath                                     NullableString `json:"MacOSPath,omitempty"`                                     // Path to the MacOS Receiver.
	MacOSMinimumSupportedVersion                  NullableString `json:"MacOSMinimumSupportedVersion,omitempty"`                  // Minimum version of the MacOS supported.
	Html5Enabled                                  NullableString `json:"Html5Enabled,omitempty"`                                  // Method of deploying and using the Html5 Receiver.
	Html5Platforms                                NullableString `json:"Html5Platforms,omitempty"`                                // The supported Html5 platforms.
	Html5Preferences                              NullableString `json:"Html5Preferences,omitempty"`                              // Html5 Receiver preferences.
	Html5SingleTabLaunch                          NullableBool   `json:"Html5SingleTabLaunch,omitempty"`                          // Launch Html5 Receiver in the same browser tab.
	Html5ChromeAppOrigins                         NullableString `json:"Html5ChromeAppOrigins,omitempty"`                         // The Html5 Chrome Application Origins settings.
	Html5ChromeAppPreferences                     NullableString `json:"Html5ChromeAppPreferences,omitempty"`                     // The Html5 Chrome Application preferences.
	ProtocolHandlerEnabled                        NullableBool   `json:"ProtocolHandlerEnabled,omitempty"`                        // Enable the Receiver Protocol Handler.
	ProtocolHandlerPlatforms                      NullableString `json:"ProtocolHandlerPlatforms,omitempty"`                      // The supported Protocol Handler platforms.
	ProtocolHandlerSkipDoubleHopCheckWhenDisabled NullableBool   `json:"ProtocolHandlerSkipDoubleHopCheckWhenDisabled,omitempty"` // Skip the Protocol Handle double hop check.
	AllowSkipNativeAppDetection                   NullableBool   `json:"AllowSkipNativeAppDetection,omitempty"`                   // Allow the user to skip native app detection.
}

// ToMap implements MappedNullable.
func (o UpdateSTFWebReceiverPluginAssistantRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.WebReceiverService.IsSet() {
		toSerialize["WebReceiverService"] = o.WebReceiverService.Get()
	}
	if o.Enabled.IsSet() {
		toSerialize["Enabled"] = o.Enabled.Get()
	}
	if o.UpgradeAtLogin.IsSet() {
		toSerialize["UpgradeAtLogin"] = o.UpgradeAtLogin.Get()
	}
	if o.ShowAfterLogin.IsSet() {
		toSerialize["ShowAfterLogin"] = o.ShowAfterLogin.Get()
	}
	if o.Win32Path.IsSet() {
		toSerialize["Win32Path"] = o.Win32Path.Get()
	}
	if o.MacOSPath.IsSet() {
		toSerialize["MacOSPath"] = o.MacOSPath.Get()
	}
	if o.MacOSMinimumSupportedVersion.IsSet() {
		toSerialize["MacOSMinimumSupportedVersion"] = o.MacOSMinimumSupportedVersion.Get()
	}
	if o.Html5Enabled.IsSet() {
		toSerialize["Html5Enabled"] = o.Html5Enabled.Get()
	}
	if o.Html5Platforms.IsSet() {
		toSerialize["Html5Platforms"] = o.Html5Platforms.Get()
	}
	if o.Html5Preferences.IsSet() {
		toSerialize["Html5Preferences"] = o.Html5Preferences.Get()
	}
	if o.Html5SingleTabLaunch.IsSet() {
		toSerialize["Html5SingleTabLaunch"] = o.Html5SingleTabLaunch.Get()
	}
	if o.Html5ChromeAppOrigins.IsSet() {
		toSerialize["Html5ChromeAppOrigins"] = o.Html5ChromeAppOrigins.Get()
	}
	if o.Html5ChromeAppPreferences.IsSet() {
		toSerialize["Html5ChromeAppPreferences"] = o.Html5ChromeAppPreferences.Get()
	}
	if o.ProtocolHandlerEnabled.IsSet() {
		toSerialize["ProtocolHandlerEnabled"] = o.ProtocolHandlerEnabled.Get()
	}
	if o.ProtocolHandlerPlatforms.IsSet() {
		toSerialize["ProtocolHandlerPlatforms"] = o.ProtocolHandlerPlatforms.Get()
	}
	if o.ProtocolHandlerSkipDoubleHopCheckWhenDisabled.IsSet() {
		toSerialize["ProtocolHandlerSkipDoubleHopCheckWhenDisabled"] = o.ProtocolHandlerSkipDoubleHopCheckWhenDisabled.Get()
	}
	if o.AllowSkipNativeAppDetection.IsSet() {
		toSerialize["AllowSkipNativeAppDetection"] = o.AllowSkipNativeAppDetection.Get()
	}

	return toSerialize, nil
}

func (o *UpdateSTFWebReceiverPluginAssistantRequestModel) SetWebReceiverService(v string) {
	o.WebReceiverService.Set(&v)
}
func (o *UpdateSTFWebReceiverPluginAssistantRequestModel) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
func (o *UpdateSTFWebReceiverPluginAssistantRequestModel) SetUpgradeAtLogin(v bool) {
	o.UpgradeAtLogin.Set(&v)
}
func (o *UpdateSTFWebReceiverPluginAssistantRequestModel) SetShowAfterLogin(v bool) {
	o.ShowAfterLogin.Set(&v)
}
func (o *UpdateSTFWebReceiverPluginAssistantRequestModel) SetWin32Path(v string) {
	o.Win32Path.Set(&v)
}
func (o *UpdateSTFWebReceiverPluginAssistantRequestModel) SetMacOSPath(v string) {
	o.MacOSPath.Set(&v)
}
func (o *UpdateSTFWebReceiverPluginAssistantRequestModel) SetMacOSMinimumSupportedVersion(v string) {
	o.MacOSMinimumSupportedVersion.Set(&v)
}
func (o *UpdateSTFWebReceiverPluginAssistantRequestModel) SetHtml5Enabled(v string) {
	o.Html5Enabled.Set(&v)
}
func (o *UpdateSTFWebReceiverPluginAssistantRequestModel) SetHtml5Platforms(v string) {
	o.Html5Platforms.Set(&v)
}
func (o *UpdateSTFWebReceiverPluginAssistantRequestModel) SetHtml5Preferences(v string) {
	o.Html5Preferences.Set(&v)
}
func (o *UpdateSTFWebReceiverPluginAssistantRequestModel) SetHtml5SingleTabLaunch(v bool) {
	o.Html5SingleTabLaunch.Set(&v)
}
func (o *UpdateSTFWebReceiverPluginAssistantRequestModel) SetHtml5ChromeAppOrigins(v string) {
	o.Html5ChromeAppOrigins.Set(&v)
}
func (o *UpdateSTFWebReceiverPluginAssistantRequestModel) SetHtml5ChromeAppPreferences(v string) {
	o.Html5ChromeAppPreferences.Set(&v)
}
func (o *UpdateSTFWebReceiverPluginAssistantRequestModel) SetProtocolHandlerEnabled(v bool) {
	o.ProtocolHandlerEnabled.Set(&v)
}
func (o *UpdateSTFWebReceiverPluginAssistantRequestModel) SetProtocolHandlerPlatforms(v string) {
	o.ProtocolHandlerPlatforms.Set(&v)
}
func (o *UpdateSTFWebReceiverPluginAssistantRequestModel) SetProtocolHandlerSkipDoubleHopCheckWhenDisabled(v bool) {
	o.ProtocolHandlerSkipDoubleHopCheckWhenDisabled.Set(&v)
}

func (o *UpdateSTFWebReceiverPluginAssistantRequestModel) SetAllowSkipNativeAppDetection(v bool) {
	o.AllowSkipNativeAppDetection.Set(&v)
}
