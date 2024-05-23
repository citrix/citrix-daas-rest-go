package models

type STFAuthenticationServiceResponseModel struct {
	ProducerService         ServiceModel                `json:"ProducerService,omitempty"`         // The producer service of the deployment
	ServiceTokenCertificate CertificateModel            `json:"ServiceTokenCertificate,omitempty"` // The service token certificate thumbprint of the deployment
	TokenValidationService  ServiceModel                `json:"TokenValidationService,omitempty"`  // The name of token validation service of the deployment
	TokenIssuerUrl          NullableString              `json:"TokenIssuerUrl,omitempty"`          // The token issuer url of the deployment
	AuthenticationOptions   AuthenticationOptionsModel  `json:"AuthenticationOptions,omitempty"`   // The authentication options of the deployment
	Authentication          AuthenticationModel         `json:"Authentication,omitempty"`          // The authentication of the deployment
	AuthenticationSettings  AuthenticationSettingsModel `json:"AuthenticationProtocols,omitempty"` // The authentication protocols of the deployment
	PnaAuthentication       PnaAuthenticationModel      `json:"PnaAuthentication,omitempty"`       // The PNA authentication of the deployment
	SiteId                  NullableInt                 `json:"SiteId,omitempty"`                  // The IIS site id of the deployment
	VirtualPath             NullableString              `json:"VirtualPath,omitempty"`             // The virtual path of the deployment
	FriendlyName            NullableString              `json:"FriendlyName,omitempty"`            // The friendly name of the deployment
	Name                    NullableString              `json:"Name,omitempty"`                    // The name of the deployment
	Hmacs                   []HmacsModel                `json:"Hmacs,omitempty"`                   // The hmacs of the deployment
	Certificates            []CertificateModel          `json:"Certificates,omitempty"`            // The certificates of the deployment
	TokenManagers           []TokenManagerModel         `json:"TokenManagers,omitempty"`           // The token managers of the deployment
	Routing                 RoutingModel                `json:"Routing,omitempty"`                 // The routing of the deployment
	ServiceRef              NullableString              `json:"ServiceRef,omitempty"`              // The service reference of the deployment
	ConfigurationFile       NullableString              `json:"ConfigurationFile,omitempty"`       // The path to configuration file of the deployment
	TenantId                NullableString              `json:"TenantId,omitempty"`                // The tenant GUID of the deployment
}

type ServiceModel struct {
	Id               NullableString   `json:"Id,omitempty"`               // The id of the service
	DisplayName      NullableString   `json:"DisplayName,omitempty"`      // The display name of the service
	TrustedIssuers   []NullableString `json:"TrustedIssuers,omitempty"`   // The trusted issuers of the service
	RelyingParties   []NullableString `json:"RelyingParties,omitempty"`   // The relying parties of the service
	AllowedAudiences []NullableString `json:"AllowedAudiences,omitempty"` // The allowed audiences of the service
}

type TokenManagerModel struct {
	Id               NullableString `json:"Id,omitempty"`               // The id of the service
	DisplayName      NullableString `json:"DisplayName,omitempty"`      // The display name of the service
	TrustedIssuers   NullableString `json:"TrustedIssuers,omitempty"`   // The trusted issuers of the service
	RelyingParties   NullableString `json:"RelyingParties,omitempty"`   // The relying parties of the service
	AllowedAudiences NullableString `json:"AllowedAudiences,omitempty"` // The allowed audiences of the service
}

type CertificateModel struct {
	Id         NullableString `json:"Id,omitempty"`         // The id of the certificate
	Thumbprint NullableString `json:"Thumbprint,omitempty"` // The thumbprint of the certificate
}

type AuthenticationOptionsModel struct {
	LoggingLevel       NullableString `json:"LoggingLevel,omitempty"`       // The logging level of the authentication options
	HideAccountErrors  NullableBool   `json:"HideAccountErrors,omitempty"`  // The hide account errors of the authentication options
	CollectFasEvidence NullableBool   `json:"CollectFasEvidence,omitempty"` // The collect FAS evidence of the authentication options
}

type AuthenticationModel struct {
	TokenServiceId  NullableString   `json:"TokenServiceId,omitempty"`  // The token service id of the authentication
	TokenValidators []NullableString `json:"TokenValidators,omitempty"` // The token validators of the authentication
	ProtocolChoices []NullableString `json:"ProtocolChoices,omitempty"` // The protocol choices of the authentication
}

type PnaAuthenticationModel struct {
	IsEnabled             NullableBool   `json:"IsEnabled,omitempty"`             // The is enabled of the PNA authentication
	ChangePasswordEnabled NullableBool   `json:"ChangePasswordEnabled,omitempty"` // The change password enabled of the PNA authentication
	ChangePasswordUrl     NullableString `json:"ChangePasswordUrl,omitempty"`     // The change password url of the PNA authentication
}

type HmacsModel struct {
	Name NullableString `json:"Name,omitempty"` // The name of the hmac
	Key  NullableString `json:"Key,omitempty"`  // The key of the hmac
}

type RoutingModel struct {
	HostBaseUrl       NullableString   `json:"HostBaseUrl,omitempty"`       // The host base url of the routing
	ServiceName       NullableString   `json:"ServiceName,omitempty"`       // The service name of the routing
	ExternalEndpoints []NullableString `json:"ExternalEndpoints,omitempty"` // The external endpoints of the routing
}

type AuthenticationSettingsModel struct {
	IntegratedWindowsAuthentication IntegratedWindowsAuthenticationModel `json:"integratedWindowsAuthentication,omitempty"` // The integrated windows authentication of the authentication settings
	CitrixAGBasicAuthentication     CitrixAGBasicAuthenticationModel     `json:"citrixAGBasicAuthentication,omitempty"`     // The citrix AG basic authentication of the authentication settings
	CommonExplicitEndpoints         CommonExplicitEndpointsModel         `json:"commonExplicitEndpoints,omitempty"`         // The common explicit endpoints of the authentication settings
	FormsProtocol                   FormsProtocolModel                   `json:"formsProtocol,omitempty"`                   // The forms protocol of the authentication settings
	ExplicitAuthentication          ExplicitAuthenticationModel          `json:"explicitAuthentication,omitempty"`          // The explicit authentication of the authentication settings
	HttpBasicAuthentication         ClaimsFactoryModel                   `json:"httpBasicAuthentication,omitempty"`         // The http basic authentication of the authentication settings
	CertificateAuthentication       ClaimsFactoryModel                   `json:"certificateAuthentication,omitempty"`       // The certificate authentication of the authentication settings
	CitrixFederationAuthentication  ClaimsFactoryModel                   `json:"citrixFederationAuthentication,omitempty"`  // The citrix federation authentication of the authentication settings
	ExplicitBL                      ExplicitBLModel                      `json:"explicitBL,omitempty"`                      // The explicit BL of the authentication settings
	DelegatedDirectoryClaimFactory  DelegatedDirectoryClaimFactoryModel  `json:"delegatedDirectoryClaimFactory,omitempty"`  // The delegated directory claim factory of the authentication settings
	SamlForms                       SamlFormsModel                       `json:"samlForms,omitempty"`                       // The saml forms of the authentication settings
}

type IntegratedWindowsAuthenticationModel struct {
	ClaimsFactoryName NullableString `json:"ClaimsFactoryName,omitempty"` // The claims factory name of the integrated windows authentication
}

type CitrixAGBasicAuthenticationModel struct {
	CredentialValidationMode NullableInt    `json:"CredentialValidationMode,omitempty"` // The credential validation mode of the citrix AG basic authentication
	NetscalerGateways        NullableString `json:"NetscalerGateways,omitempty"`        // The netscaler gateways of the citrix AG basic authentication
	ClaimsFactoryName        NullableString `json:"ClaimsFactoryName,omitempty"`        // The claims factory name of the citrix AG basic authentication
}

type CommonExplicitEndpointsModel struct {
	ChangePassword ChangePasswordModel `json:"ChangePassword,omitempty"` // The change password endpoint of the common explicit endpoints
}

type ChangePasswordModel struct {
	ConversationFactoryName NullableString `json:"ConversationFactoryName,omitempty"` // The conversation factory name of the change password
	ClaimsFactoryName       NullableString `json:"ClaimsFactoryName,omitempty"`       // The claims factory name of the change password
}

type FormsProtocolModel struct {
	TemplatePath NullableString `json:"TemplatePath,omitempty"` // The template path of the forms protocol
	CompiledPath NullableString `json:"CompiledPath,omitempty"` // The compiled path of the forms protocol
}

type ExplicitAuthenticationModel struct {
	AccountManagerFactoryName NullableString `json:"AccountManagerFactoryName,omitempty"` // The account manager factory name of the explicit authentication
	ConversationFactoryName   NullableString `json:"ConversationFactoryName,omitempty"`   // The conversation factory name of the explicit authentication
	ClaimsFactoryName         NullableString `json:"ClaimsFactoryName,omitempty"`         // The claims factory name of the explicit authentication
}

type ClaimsFactoryModel struct {
	ClaimsFactoryName NullableString `json:"ClaimsFactoryName,omitempty"` // The claims factory name of the claims factory
}

type ExplicitBLModel struct {
	Authenticator               NullableString               `json:"Authenticator,omitempty"`               // The authenticator of the explicit BL
	HideDomainField             NullableBool                 `json:"HideDomainField,omitempty"`             // The hide domain field of the explicit BL
	AllowUserPasswordChange     NullableInt                  `json:"AllowUserPasswordChange,omitempty"`     // The allow user password change of the explicit BL
	AllowZeroLengthPassword     NullableBool                 `json:"AllowZeroLengthPassword,omitempty"`     // The allow zero length password of the explicit BL
	ShowPasswordExpiryWarning   NullableInt                  `json:"ShowPasswordExpiryWarning,omitempty"`   // The show password expiry warning of the explicit BL
	PasswordExpiryWarningPeriod NullableInt                  `json:"PasswordExpiryWarningPeriod,omitempty"` // The password expiry warning period of the explicit BL
	RequireAccountSIDs          NullableBool                 `json:"RequireAccountSIDs,omitempty"`          // The require account SIDs of the explicit BL
	DomainSelection             NullableString               `json:"DomainSelection,omitempty"`             // The domain selection of the explicit BL
	AccountManagementPolicy     AccountManagementPolicyModel `json:"AccountManagementPolicy,omitempty"`     // The account management policy of the explicit BL
	JsonServiceEnabled          NullableBool                 `json:"JsonServiceEnabled,omitempty"`          // The json service enabled of the explicit BL
}

type AccountManagementPolicyModel struct {
	AllowUnlockAccount NullableBool `json:"AllowUnlockAccount,omitempty"` // The allow unlock account of the account management policy
	AllowResetPassword NullableBool `json:"AllowResetPassword,omitempty"` // The allow reset password of the account management policy
}

type DelegatedDirectoryClaimFactoryModel struct {
	RequiredProperties NullableString                                    `json:"RequiredProperties,omitempty"` // The required properties of the delegated directory claim factory
	Enabled            NullableBool                                      `json:"Enabled,omitempty"`            // The enabled of the delegated directory claim factory
	ClientEndpoint     DelegatedDirectoryClaimFactoryClientEndpointModel `json:"ClientEndpoint,omitempty"`     // The client endpoint of the delegated directory claim factory
}

type DelegatedDirectoryClaimFactoryClientEndpointModel struct {
	Uri                NullableString                       `json:"Uri,omitempty"`                // The uri of the client endpoint
	TransferMode       NullableInt                          `json:"TransferMode,omitempty"`       // The transfer mode of the client endpoint
	AuthenticationMode NullableInt                          `json:"AuthenticationMode,omitempty"` // The authentication mode of the client endpoint
	WsSecurityEnabled  NullableBool                         `json:"WsSecurityEnabled,omitempty"`  // The ws security enabled of the client endpoint
	ClientCertificate  ClientEndpointClientCertificateModel `json:"ClientCertificate,omitempty"`  // The client certificate of the client endpoint
	OpenTimeOut        TimeModel                            `json:"OpenTimeOut,omitempty"`        // The open time out of the client endpoint
	CloseTimeOut       TimeModel                            `json:"CloseTimeOut,omitempty"`       // The close time out of the client endpoint
	SendTimeOut        TimeModel                            `json:"SendTimeOut,omitempty"`        // The send time out of the client endpoint
	PortSharingEnabled NullableBool                         `json:"PortSharingEnabled,omitempty"` // The port sharing enabled of the client endpoint
}

type ClientEndpointClientCertificateModel struct {
	Thumbprint    NullableString `json:"Thumbprint,omitempty"`    // The thumbprint of the client certificate
	StoreName     NullableString `json:"StoreName,omitempty"`     // The store name of the client certificate
	StoreLocation NullableInt    `json:"StoreLocation,omitempty"` // The store location of the client certificate
}

type TimeModel struct {
	Ticks             NullableInt     `json:"Ticks,omitempty"`             // The ticks of the time
	Days              NullableInt     `json:"Days,omitempty"`              // The days of the time
	Hours             NullableInt     `json:"Hours,omitempty"`             // The hours of the time
	Milliseconds      NullableInt     `json:"Milliseconds,omitempty"`      // The milliseconds of the time
	Minutes           NullableInt     `json:"Minutes,omitempty"`           // The minutes of the time
	Seconds           NullableInt     `json:"Seconds,omitempty"`           // The seconds of the time
	TotalDays         NullableFloat64 `json:"TotalDays,omitempty"`         // The total days of the time
	TotalHours        NullableFloat64 `json:"TotalHours,omitempty"`        // The total hours of the time
	TotalMilliseconds NullableInt     `json:"TotalMilliseconds,omitempty"` // The total milliseconds of the time
	TotalMinutes      NullableFloat64 `json:"TotalMinutes,omitempty"`      // The total minutes of the time
	TotalSeconds      NullableInt     `json:"TotalSeconds,omitempty"`      // The total seconds of the time
}

type SamlFormsModel struct {
	ConversationFactoryName NullableString       `json:"ConversationFactoryName,omitempty"` // The conversation factory name of the saml forms
	ClaimsFactoryName       NullableString       `json:"ClaimsFactoryName,omitempty"`       // The claims factory name of the saml forms
	SamlSettings            SamlSettingsModel    `json:"SamlSettings,omitempty"`            // The saml settings of the saml forms
	WebViewSettings         WebViewSettingsModel `json:"WebViewSettings,omitempty"`         // The web view settings of the saml forms
}

type SamlSettingsModel struct {
	IdentityProvider SamlIdentityProviderModel `json:"IdentityProvider,omitempty"` // The identity provider of the saml settings
	ServiceProvider  SamlServiceProviderModel  `json:"ServiceProvider,omitempty"`  // The service provider of the saml settings

}

type SamlIdentityProviderModel struct {
	SingleSignOnService       SingleSignOnServiceModel `json:"SingleSignOnService,omitempty"`       // The single sign on service of the identity provider
	SingleLogoutService       SingleSignOnServiceModel `json:"SingleLogoutService,omitempty"`       // The single logout service of the identity provider
	ArtifactResolutionAddress NullableString           `json:"ArtifactResolutionAddress,omitempty"` // The artifact resolution address of the identity provider
	SigningCertificates       CertificateModel         `json:"SigningCertificates,omitempty"`       // The signing certificates of the identity provider
}

type SingleSignOnServiceModel struct {
	PostAddress     NullableString `json:"PostAddress,omitempty"`     // The post address of the single sign on service
	RedirectAddress NullableString `json:"RedirectAddress,omitempty"` // The redirect address of the single sign on service
	Binding         NullableString `json:"Binding,omitempty"`         // The binding of the single sign on service
}

type SamlServiceProviderModel struct {
	Uri                             NullableString      `json:"Uri,omitempty"`                             // The uri of the service provider
	NameClaimType                   NullableString      `json:"NameClaimType,omitempty"`                   // The name claim type of the service provider
	SecurityTokenHandlerFactoryName NullableString      `json:"SecurityTokenHandlerFactoryName,omitempty"` // The security token handler factory name of the service provider
	RequireArtifactResolution       NullableBool        `json:"RequireArtifactResolution,omitempty"`       // The require artifact resolution of the service provider
	EncryptionCertificates          []NullableString    `json:"EncryptionCertificates,omitempty"`          // The encryption certificates of the service provider
	RequestSigning                  RequestSigningModel `json:"RequestSigning,omitempty"`                  // The request signing of the service provider
}

type RequestSigningModel struct {
	Enabled              NullableBool   `json:"Enabled,omitempty"`              // The enabled of the request signing
	SigningCertificateId NullableString `json:"SigningCertificateId,omitempty"` // The signing certificate id of the request signing
	Algorithm            NullableString `json:"Algorithm,omitempty"`            // The algorithm of the request signing
}

type WebViewSettingsModel struct {
	AllowedReturnUrls NullableString `json:"AllowedReturnUrls,omitempty"` // The allowed return urls of the web view settings
}

type STFAuthenticationServiceProtocolResponseModel struct {
	Name             NullableString `json:"Name"`
	Enabled          NullableBool   `json:"Enabled"`
	ChoiceProperties any            `json:"ChoiceProperties"` //TODO: placeholder for any type, fill in the specific type later
}
