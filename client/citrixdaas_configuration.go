// Copyright © 2023. Citrix Systems, Inc.

package citrixclient

// AuthenticationConfiguration provides authentication settings for CC Athena / on-prem trust service
type AuthenticationConfiguration struct {
	AuthUrl      string `json:"auth_url"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	OnPremises   bool   `json:"on_premises"`
}

// ClientConfiguration provides Citrix DaaS customer context
type ClientConfiguration struct {
	CustomerId string
	SiteId     string
	UserAgent  string
}
