// Copyright © 2025. Citrix Systems, Inc.
package models

type STFWebReceiverBlockingNotificationResponse struct {
	Enabled   NullableBool   `json:"Enabled"`
	Title     NullableString `json:"Title"`
	Body      NullableString `json:"Body"`
	Button    NullableString `json:"Button"`
	Frequency int            `json:"Frequency"`
}
