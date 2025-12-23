// Copyright © 2025. Citrix Systems, Inc.

package citrixclient

// AuthenticationConfiguration provides authentication settings for CC Athena / on-premises trust service
type AuthenticationConfiguration struct {
	AuthUrl      string `json:"auth_url"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	OnPremises   bool   `json:"on_premises"`
	ApiGateway   bool   `json:"api_gateway"`
	IsGov        bool   `json:"is_gov"`
	Environment  string `json:"environment"`
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

// WemOnPremAuthentication provides authentication settings for WEM on-premises service
type WemOnPremAuthentication struct {
	AuthUrl       string `json:"auth_url"`
	AdminUserName string `json:"admin_username"`
	AdminPassword string `json:"admin_password"`
}

// Default retry configuration values
const (
	DefaultRetryBaseDelay             = 15
	DefaultMaxRetries                 = 3
	DefaultJobPollIntervalSeconds     = 10
	DefaultJobLongPollIntervalSeconds = 30
)

// contextKey is a custom type for context keys to avoid collisions
type contextKey string

// Context keys for retry configuration
// These can be used with context.WithValue() to override retry behavior for testing
const (
	RetryBaseDelayKey             contextKey = "citrix.retry.baseDelay"
	MaxRetriesKey                 contextKey = "citrix.retry.maxRetries"
	JobPollIntervalSecondsKey     contextKey = "citrix.job.pollInterval"
	JobLongPollIntervalSecondsKey contextKey = "citrix.job.longPollInterval"
)
