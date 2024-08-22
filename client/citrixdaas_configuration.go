// Copyright © 2024. Citrix Systems, Inc.

package citrixclient

// AuthenticationConfiguration provides authentication settings for CC Athena / on-premises trust service
type AuthenticationConfiguration struct {
	AuthUrl      string `json:"auth_url"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	OnPremises   bool   `json:"on_premises"`
	ApiGateway   bool   `json:"api_gateway"`
	IsGov        bool   `json:"is_gov"`
}

// ClientConfiguration provides Citrix DaaS customer context
type ClientConfiguration struct {
	CustomerId              string
	SiteId                  string
	UserAgent               string
	Accept                  string
	Authorization           string
	OrchestrationApiVersion int32
	ProductVersion          string
	IsCspCustomer           bool
}
