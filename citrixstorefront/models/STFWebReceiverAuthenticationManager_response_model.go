package models

type GetWebReceiverAuthenticationManagerResponseModel struct {
	GetUserNameUrl       NullableString `json:"GetUserNameUrl,omitempty"`
	LogoffUrl            NullableString `json:"LogoffUrl,omitempty"`
	ChangeCredentialsUrl NullableString `json:"ChangeCredentialsUrl,omitempty"`
	LoginFormTimeout     NullableInt    `json:"LoginFormTimeout,omitempty"`
}
