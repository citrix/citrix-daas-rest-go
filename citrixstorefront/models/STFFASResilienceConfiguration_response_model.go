// Copyright © 2026. Citrix Systems, Inc.
package models

type GetFASResilienceConfigurationResponseModel struct {
	Enabled                       NullableBool `json:"Enabled,omitempty"`
	NumberOfExceptionsBeforeBreak NullableInt  `json:"NumberOfExceptionsBeforeBreak,omitempty"`
	DurationOfBreakInMinutes      NullableInt  `json:"DurationOfBreakInMinutes,omitempty"`
}
