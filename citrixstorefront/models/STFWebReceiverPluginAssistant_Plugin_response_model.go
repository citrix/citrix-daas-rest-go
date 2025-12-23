// Copyright © 2025. Citrix Systems, Inc.
package models

type WebReceiverPluginAssistantModel struct {
	Enabled         NullableBool    `json:"Enabled"`
	UpgradeAtLogin  NullableBool    `json:"UpgradeAtLogin"`
	ShowAfterLogin  NullableBool    `json:"ShowAfterLogin"`
	Win32           Win32           `json:"Win32"`
	MacOS           MacOS           `json:"MacOS"`
	HTML5           HTML5           `json:"Html5"`
	ProtocolHandler ProtocolHandler `json:"ProtocolHandler"`
}
type Win32 struct {
	Path NullableString `json:"Path"`
}
type MacOS struct {
	Path                    NullableString `json:"Path"`
	MinimumSupportedVersion NullableString `json:"MinimumSupportedVersion"`
}
type Version struct {
	Major         NullableInt `json:"Major"`
	Minor         NullableInt `json:"Minor"`
	Build         NullableInt `json:"Build"`
	Revision      NullableInt `json:"Revision"`
	MajorRevision NullableInt `json:"MajorRevision"`
	MinorRevision NullableInt `json:"MinorRevision"`
}
type HTML5 struct {
	Enabled              NullableInt64  `json:"Enabled"`
	Platforms            NullableString `json:"Platforms"`
	LaunchURL            NullableString `json:"LaunchUrl"`
	Preferences          NullableString `json:"Preferences"`
	SingleTabLaunch      NullableBool   `json:"SingleTabLaunch"`
	ChromeAppOrigins     NullableString `json:"ChromeAppOrigins"`
	ChromeAppPreferences NullableString `json:"ChromeAppPreferences"`
	Version              Version        `json:"Version"`
	ProductCode          NullableString `json:"ProductCode"`
	IsInstalled          NullableBool   `json:"IsInstalled"`
}
type ProtocolHandler struct {
	Enabled                        NullableBool   `json:"Enabled"`
	Platforms                      NullableString `json:"Platforms"`
	SkipDoubleHopCheckWhenDisabled NullableBool   `json:"SkipDoubleHopCheckWhenDisabled"`
}
