// Copyright © 2025. Citrix Systems, Inc.
package models

type STFVersionModel struct {
	Major         NullableInt `json:"Major,omitempty"`
	Minor         NullableInt `json:"Minor,omitempty"`
	Build         NullableInt `json:"Build,omitempty"`
	Revision      NullableInt `json:"Revision,omitempty"`
	MajorRevision NullableInt `json:"MajorRevision,omitempty"`
	MinorRevision NullableInt `json:"MinorRevision,omitempty"`
}
