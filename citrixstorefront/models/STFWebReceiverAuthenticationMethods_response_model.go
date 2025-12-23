// Copyright © 2025. Citrix Systems, Inc.
package models

type STFWebReceiverAuthenticationMethodsResponse struct {
	LocationUrl   NullableString `json:"LocationUrl"`
	TokenLifeTime TokenLifeTime  `json:"TokenLifeTime"`
	Methods       []string       `json:"Methods"`
}

type TokenLifeTime struct {
	Ticks             NullableInt64   `json:"Ticks"`
	Days              NullableInt     `json:"Days"`
	Hours             NullableInt     `json:"Hours"`
	Milliseconds      NullableInt     `json:"Milliseconds"`
	Minutes           NullableInt     `json:"Minutes"`
	Seconds           NullableInt     `json:"Seconds"`
	TotalDays         NullableFloat64 `json:"TotalDays"`
	TotalHours        NullableInt     `json:"TotalHours"`
	TotalMilliseconds NullableInt     `json:"TotalMilliseconds"`
	TotalMinutes      NullableInt     `json:"TotalMinutes"`
	TotalSeconds      NullableInt     `json:"TotalSeconds"`
}
