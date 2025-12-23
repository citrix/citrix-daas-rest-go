// Copyright © 2025. Citrix Systems, Inc.
package models

type STFRoamingServiceResponseModel struct {
	Roaming                            RoamingModel                   `json:"Roaming,omitempty"`
	CasCongfiguration                  CasConfidurationModel          `json:"CasConfiguration,omitempty"`
	RoamingSettings                    any                            `json:"RoamingSettings,omitempty"` //TODO: Placeholder for RoamingSettings
	HostingEnvironment                 HostingEnvironmentModel        `json:"HostingEnvironment,omitempty"`
	RunAllManagedModulesForAllRequests NullableBool                   `json:"RunAllManagedModulesForAllRequests,omitempty"`
	SiteId                             NullableInt                    `json:"SiteId,omitempty"`
	VirtualPath                        NullableString                 `json:"VirtualPath,omitempty"`
	Name                               NullableString                 `json:"Name,omitempty"`
	FriendlyName                       NullableString                 `json:"FriendlyName,omitempty"`
	FeatureData                        FeatureDataModel               `json:"FeatureData,omitempty"`
	Hmacs                              []HmacsModel                   `json:"Hmacs,omitempty"`
	Certificates                       []CertificateModel             `json:"Certificates,omitempty"`
	TokenManagers                      []ServiceModel                 `json:"TokenManagers,omitempty"`
	Routing                            RoutingModel                   `json:"Routing,omitempty"`
	CredentialWallet                   CredentialWalletModel          `json:"CredentialWallet,omitempty"`
	WebApplication                     WebApplicationModel            `json:"WebApplication,omitempty"`
	ServiceRef                         NullableString                 `json:"ServiceRef,omitempty"`
	PhysicalPath                       NullableString                 `json:"PhysicalPath,omitempty"`
	Diagnostics                        RoamingServiceDiagnosticsModel `json:"Diagnostics,omitempty"`
	ApplicationInitialization          ApplicationInitializationModel `json:"ApplicationInitialization,omitempty"`
	FeatureInstanceId                  NullableString                 `json:"FeatureInstanceId,omitempty"`
	ConfigurationFile                  NullableString                 `json:"ConfigurationFile,omitempty"`
	TenantId                           NullableString                 `json:"TenantId,omitempty"`
}

type RoamingModel struct {
	Gateways           []GatewayModel          `json:"Gateways,omitempty"`
	CommonData         []NameValuePairModel    `json:"CommonData,omitempty"`
	Beacons            BeaconModel             `json:"Beacons,omitempty"`
	ServiceRecords     []ServiceRecordModel    `json:"ServiceRecords,omitempty"`
	Accounts           []AccountModel          `json:"Accounts,omitempty"`
	AccountRedirection AccountRedirectionModel `json:"AccountRedirection,omitempty"`
}

type GatewayModel struct {
	Id                         NullableString `json:"Id,omitempty"`
	Name                       NullableString `json:"Name,omitempty"`
	Default                    NullableBool   `json:"Default,omitempty"`
	Edition                    NullableInt    `json:"Edition,omitempty"`
	Version                    NullableInt    `json:"Version,omitempty"`
	Logon                      NullableInt    `json:"Logon,omitempty"`
	SmartCardFallback          NullableInt    `json:"SmartCardFallback,omitempty"`
	NetScalerTrustCertificates NullableString `json:"NetScalerTrustCertificates,omitempty"`
	NetScalerGatewayMode       NullableString `json:"NetScalerGatewayMode,omitempty"`
	CallbackUrl                NullableString `json:"CallbackUrl,omitempty"`
	RWMode                     NullableString `json:"RWMode,omitempty"`
	Deployment                 NullableString `json:"Deployment,omitempty"`
	Location                   NullableString `json:"Location,omitempty"`
	GslLocation                NullableString `json:"GslLocation,omitempty"`
	SessionReliability         NullableBool   `json:"SessionReliability,omitempty"`
	RequestTicketTwoStas       NullableBool   `json:"RequestTicketTwoStas,omitempty"`
	IpAddress                  NullableString `json:"IpAddress,omitempty"`
	StasUseLoadBalancing       NullableBool   `json:"StasUseLoadBalancing,omitempty"`
	StasBypassDuration         NullableString `json:"StasBypassDuration,omitempty"`
	SecureTicketAuthorityUrls  NullableString `json:"SecureTicketAuthorityUrls,omitempty"`
	IsCloudGateway             NullableBool   `json:"IsCloudGateway,omitempty"`
	GatewayServiceLookupUrl    NullableString `json:"GatewayServiceLookupUrl,omitempty"`
	ConnectorBaseUrls          NullableString `json:"ConnectorBaseUrls,omitempty"`
	GatewayLookupUrls          NullableString `json:"GatewayLookupUrls,omitempty"`
	AuthenticationCapable      NullableBool   `json:"AuthenticationCapable,omitempty"`
	HdxRoutingCapable          NullableBool   `json:"HdxRoutingCapable,omitempty"`
	NetScalerImport            NullableBool   `json:"NetScalerImport,omitempty"`
}

type BeaconModel struct {
	Internal []BeaconDetailModel `json:"Internal,omitempty"`
	External []BeaconDetailModel `json:"External,omitempty"`
}

type BeaconDetailModel struct {
	Id   NullableString `json:"Id,omitempty"`
	Path NullableString `json:"Path,omitempty"`
}

type ServiceRecordModel struct {
	ServiceRef      NullableString `json:"ServiceRef,omitempty"`
	Name            NullableString `json:"Name,omitempty"`
	Description     NullableString `json:"Description,omitempty"`
	Id              NullableString `json:"Id,omitempty"`
	Type            NullableInt    `json:"Type,omitempty"`
	Address         NullableString `json:"Address,omitempty"`
	Beacons         NullableString `json:"Beacons,omitempty"`
	Gateways        NullableString `json:"Gateways,omitempty"`
	OptimalGateways NullableString `json:"OptimalGateways,omitempty"`
}

type AccountModel struct {
	Id                      NullableString `json:"Id,omitempty"`
	Name                    NullableString `json:"Name,omitempty"`
	Description             NullableString `json:"Description,omitempty"`
	Published               NullableBool   `json:"Published,omitempty"`
	UpdaterType             NullableInt    `json:"UpdaterType,omitempty"`
	RemoteAccessType        NullableInt    `json:"RemoteAccessType,omitempty"`
	AnnotatedServiceRecords NullableString `json:"AnnotatedServiceRecords,omitempty"`
	Metadata                NullableString `json:"Metadata,omitempty"`
}

type AccountRedirectionModel struct {
	RedirectUrl         NullableString   `json:"RedirectUrl,omitempty"`
	FtuRedirects        []NullableString `json:"FtuRedirects,omitempty"`
	UserAgentsRedirects []NullableString `json:"UserAgentsRedirects,omitempty"`
}

type CasConfidurationModel struct {
	CustomerId              NullableString `json:"CustomerId,omitempty"`
	EnablementService       NullableString `json:"EnablementService,omitempty"`
	CwsServiceKey           NullableString `json:"CwsServiceKey,omitempty"`
	EnablementServiceStatus NullableString `json:"EnablementServiceStatus,omitempty"`
	InstanceId              NullableString `json:"InstanceId,omitempty"`
	Name                    NullableString `json:"Name,omitempty"`
}

type FeatureDataModel struct {
	EventSource      NullableString `json:"EventSource,omitempty"`
	SiteID           NullableString `json:"SiteID,omitempty"`
	VirtualPath      NullableString `json:"VirtualPath,omitempty"`
	ApplicationPool  NullableString `json:"ApplicationPool,omitempty"`
	DefaultDocuments NullableString `json:"DefaultDocuments,omitempty"`
	Name             NullableString `json:"Name,omitempty"`
	UseHTTPS         NullableString `json:"UseHTTPS,omitempty"`
	DeployFolder     NullableString `json:"DeployFolder,omitempty"`
	ConfigLocation   NullableString `json:"ConfigLocation,omitempty"`
	ConfigTypeName   NullableString `json:"ConfigTypeName,omitempty"`
}

type WebApplicationModel struct {
	WebSite            WebSiteDetailsModel     `json:"WebSite,omitempty"`
	VirtualPath        NullableString          `json:"VirtualPath,omitempty"`
	Name               NullableString          `json:"Name,omitempty"`
	AppPool            NullableString          `json:"AppPool,omitempty"`
	Folder             NullableString          `json:"Folder,omitempty"`
	VirtualDirectories []VirtualDirectoryModel `json:"VirtualDirectories,omitempty"`
}

type WebSiteDetailsModel struct {
	PhysicalPath        NullableString             `json:"PhysicalPath,omitempty"`
	DefaultApplPoolName NullableString             `json:"DefaultApplPoolName,omitempty"`
	Bindings            []BindingModel             `json:"Bindings,omitempty"`
	Name                NullableString             `json:"Name,omitempty"`
	Id                  NullableInt                `json:"Id,omitempty"`
	Applications        []BasicWebApplicationModel `json:"Applications,omitempty"`
}

type BindingModel struct {
	Protocol    NullableString `json:"Protocol,omitempty"`
	Host        NullableString `json:"Host,omitempty"`
	HostSource  NullableInt    `json:"HostSource,omitempty"`
	Port        NullableInt    `json:"Port,omitempty"`
	BaseUrl     NullableString `json:"BaseUrl,omitempty"`
	Information NullableString `json:"Information,omitempty"`
	Certificate any            `json:"Certificate,omitempty"`
}

type VirtualDirectoryModel struct {
	VirtualPath         NullableString `json:"VirtualPath,omitempty"`
	AbsoluteVirtualPath NullableString `json:"AbsoluteVirtualPath,omitempty"`
	PhysicalPath        NullableString `json:"PhysicalPath,omitempty"`
}

type RoamingServiceDiagnosticsModel struct {
	SharedListeners SharedListenersModel `json:"SharedListeners,omitempty"`
}

type SharedListenersModel struct {
	CitrixTraceListener CitrixTraceListenerModel `json:"citrixTraceListener,omitempty"`
}

type CitrixTraceListenerModel struct {
	Name             NullableString `json:"Name,omitempty"`
	Type             NullableString `json:"Type,omitempty"`
	InitializeData   NullableString `json:"InitializeData,omitempty"`
	MaxFileSizeKb    NullableInt    `json:"MaxFileSizeKb,omitempty"`
	HostBaseUrl      NullableString `json:"HostBaseUrl,omitempty"`
	ServiceType      NullableString `json:"ServiceType,omitempty"`
	SiteId           NullableString `json:"SiteId,omitempty"`
	EndPointUrl      NullableString `json:"EndPointUrl,omitempty"`
	CloudProvider    NullableString `json:"CloudProvider,omitempty"`
	AuthToken        NullableString `json:"AuthToken,omitempty"`
	FlushInterval    NullableInt    `json:"FlushInterval,omitempty"`
	TracesPerRequest NullableInt    `json:"TracesPerRequest,omitempty"`
	FileCount        NullableInt    `json:"FileCount,omitempty"`
}
