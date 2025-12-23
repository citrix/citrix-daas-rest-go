// Copyright © 2025. Citrix Systems, Inc.

package models

type STFWebReceiverDiscoveryServiceResponseModel struct {
	Url                 NullableString `json:"Url"`
	RunDiscoveryAtStart NullableBool   `json:"RunDiscoveryAtStart"`
}
