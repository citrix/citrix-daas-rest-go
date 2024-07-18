// Copyright © 2024. Citrix Systems, Inc.
package models

type STFWebReceiverSiteStyleResponseModel struct {
	HeaderLogoPath        NullableString `json:"HeaderLogoPath,omitempty"`
	LogonLogoPath         NullableString `json:"LogonLogoPath,omitempty"`
	HeaderBackgroundColor NullableString `json:"HeaderBackgroundColor,omitempty"`
	HeaderForegroundColor NullableString `json:"HeaderForegroundColor,omitempty"`
	LinkColor             NullableString `json:"LinkColor,omitempty"`
}
