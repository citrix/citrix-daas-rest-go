// Copyright © 2026. Citrix Systems, Inc.

package models

type STFPna struct {
	FeatureData       StoreFeatureData `json:"FeatureData,omitempty"`
	PnaEnabled        NullableBool     `json:"PnaEnabled,omitempty"`
	DefaultPnaService NullableBool     `json:"DefaultPnaService,omitempty"`
}
