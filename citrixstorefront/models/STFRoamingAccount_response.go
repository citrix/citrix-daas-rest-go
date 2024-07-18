// Copyright © 2024. Citrix Systems, Inc.

package models

// Get-STFRoamingAccount

type GetSTFRoamingAccountResponseModel struct {
	Id                      NullableString `json:"Id,omitempty"`
	Name                    NullableString `json:"Name,omitempty"`
	Description             NullableString `json:"Description,omitempty"`
	Published               NullableBool   `json:"Published,omitempty"`
	UpdaterType             NullableInt    `json:"UpdaterType,omitempty"`
	RemoteAccessType        NullableInt    `json:"RemoteAccessType,omitempty"`
	AnnotatedServiceRecords any            `json:"AnnotatedServiceRecords,omitempty"`
	Metadata                any            `json:"Metadata,omitempty"`
}
