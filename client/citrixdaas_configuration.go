package citrixclient

// AuthenticationConfiguration provides authentication settings for CC Athena / on-prem trust service
type AuthenticationConfiguration struct {
	AuthUrl      string `json:"auth_url"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	OnPremise    bool   `json:"on_premise"`
}

// ClientConfiguration provides Citrix DaaS customer context
type ClientConfiguration struct {
	CustomerId string
	SiteId     string
}
