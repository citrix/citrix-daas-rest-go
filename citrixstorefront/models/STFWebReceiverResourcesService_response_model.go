// Copyright © 2024. Citrix Systems, Inc.
package models

type GetSTFWebReceiverResourcesServiceResponseModel struct {
	PersistentIconCacheEnabled NullableBool `json:"PersistentIconCacheEnabled,omitempty"`
	IcaFileCacheExpiry         NullableInt  `json:"IcaFileCacheExpiry,omitempty"`
	IconSize                   NullableInt  `json:"IconSize,omitempty"`
	ShowDesktopViewer          NullableBool `json:"ShowDesktopViewer,omitempty"`
}
