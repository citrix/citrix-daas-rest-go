// Copyright © 2025. Citrix Systems, Inc.
package models

type GetSTFRoamingInternalBeaconResponseModel struct {
	Internal string `json:"internal_address"`
}

type GetSTFRoamingExternalBeaconResponseModel struct {
	External []string `json:"external_addresses"`
}
