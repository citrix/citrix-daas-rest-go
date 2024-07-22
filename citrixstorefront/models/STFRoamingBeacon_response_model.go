// Copyright © 2024. Citrix Systems, Inc.
package models

type GetSTFRoamingInternalBeaconResponseModel struct {
	Internal string `json:"internal_ip"`
}

type GetSTFRoamingExternalBeaconResponseModel struct {
	External []string `json:"external_ips"`
}
