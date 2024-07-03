// Copyright © 2024. Citrix Systems, Inc.

package models

type STFPna struct {
	Service           StoreServiceModel `json:"StoreService,omitempty"`
	PnaEnabled        NullableBool      `json:"PnaEnabled,omitempty"`
	DefaultPnaService NullableBool      `json:"DefaultPnaService,omitempty"`
}
