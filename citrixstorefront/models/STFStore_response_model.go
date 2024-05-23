package models

type STFStoreDetailModel struct {
	SiteId                             NullableInt                         `json:"SiteId,omitempty"`
	VirtualPath                        NullableString                      `json:"VirtualPath,omitempty"`
	FriendlyName                       NullableString                      `json:"FriendlyName,omitempty"`
	Service                            StoreServiceModel                   `json:"Service,omitempty"`
	FarmsConfiguration                 StoreFarmsConfiguration             `json:"FarmsConfiguration,omitempty"`
	UserFarmMappings                   []NullableString                    `json:"UserFarmMappings,omitempty"`
	Authentication                     StoreAuthenticationModel            `json:"Authentication,omitempty"`
	Subscription                       StoreSubscriptionModel              `json:"Subscription,omitempty"`
	Gateways                           []NullableString                    `json:"Gateways,omitempty"`
	GatewayService                     StoreGatewayServiceModel            `json:"GatewayService,omitempty"`
	OptimalGatewaysForFarms            []NullableString                    `json:"OptimalGatewaysForFarms,omitempty"`
	DiscoveryService                   StoreDiscoveryServiceModel          `json:"DiscoveryService,omitempty"`
	PNA                                StorePNAModel                       `json:"PNA,omitempty"`
	Resources                          StoreResources                      `json:"Resources,omitempty"`
	StoreResources                     StoreResourcesModel                 `json:"StoreResources,omitempty"`
	Workflow                           StoreWorkflowModel                  `json:"Workflow,omitempty"`
	SubscriptionStoreClient            SubscriptionStoreClientModel        `json:"SubscriptionStoreClient,omitempty"`
	IcaLaunchDefaults                  IcaLaunchDefaultsModel              `json:"IcaLaunchDefaults,omitempty"`
	IcaFileSigning                     IcaFileSigningModel                 `json:"IcaFileSigning,omitempty"`
	DatabaseConnectionNullableString   DatabaseConnectionStringModel       `json:"DatabaseConnectionNullableString,omitempty"`
	AppSettings                        StoreAppSettingsModel               `json:"AppSettings,omitempty"`
	SessionManagerSettings             SessionManagerSettingsModel         `json:"SessionManagerSettings,omitempty"`
	StoreSettings                      StoreSettingsModel                  `json:"StoreSettings,omitempty"`
	ServiceId                          NullableString                      `json:"ServiceId,omitempty"`
	ConsumerService                    ServiceModel                        `json:"ConsumerService,omitempty"`
	ServiceTokenCertificate            StoreServiceTokenCertificateModel   `json:"ServiceTokenCertificate,omitempty"`
	AuthenticationServiceVirtualPath   NullableString                      `json:"AuthenticationServiceVirtualPath,omitempty"`
	HostingEnvironment                 StoreHostingEnvModel                `json:"HostingEnvironment,omitempty"`
	RunAllManagedModulesForAllRequests NullableBool                        `json:"RunAllManagedModulesForAllRequests,omitempty"`
	FeatureData                        StoreFeatureData                    `json:"FeatureData,omitempty"`
	Name                               NullableString                      `json:"Name,omitempty"`
	Hmacs                              []HmacsModel                        `json:"Hmacs,omitempty"`
	Certificates                       []CertificateModel                  `json:"Certificates,omitempty"`
	TokenManagers                      []TokenManagerModel                 `json:"TokenManagers,omitempty"`
	Routing                            RoutingModel                        `json:"Routing,omitempty"`
	CredentialWallet                   StoreCredentialWalletModel          `json:"CredentialWallet,omitempty"`
	WebApplication                     StoreWebApplicationModel            `json:"WebApplication,omitempty"`
	ServiceRef                         NullableString                      `json:"ServiceRef,omitempty"`
	PhysicalPath                       NullableString                      `json:"PhysicalPath,omitempty"`
	Diagnostics                        StoreDiagnosticsModel               `json:"Diagnostics,omitempty"`
	ApplicationInitialization          StoreApplicationInitializationModel `json:"ApplicationInitialization,omitempty"`
	FeatureInstanceId                  NullableString                      `json:"FeatureInstanceId,omitempty"`
	Configuration                      StoreConfiguration                  `json:"Configuration,omitempty"`
	ConfigurationFile                  NullableString                      `json:"ConfigurationFile,omitempty"`
	TenantId                           NullableString                      `json:"TenantId,omitempty"`
}

type StoreServiceModel struct {
	LockedDown             NullableBool `json:"LockedDown"`
	Anonymous              NullableBool `json:"Anonymous"`
	AllowSessionReconnect  NullableBool `json:"AllowSessionReconnect"`
	SubstituteDesktopImage NullableBool `json:"SubstituteDesktopImage"`
}

type StoreAuthenticationModel struct {
	KerberosDelegation NullableBool `json:"KerberosDelegation"`
}

type StoreFeatureData struct {
	EventSource      NullableString `json:"EventSource,omitempty"`
	SiteID           NullableString `json:"SiteID,omitempty"`
	VirtualPath      NullableString `json:"VirtualPath,omitempty"`
	ApplicationPool  NullableString `json:"ApplicationPool,omitempty"`
	DefaultDocuments NullableString `json:"DefaultDocuments,omitempty"`
	Name             NullableString `json:"Name,omitempty"`
	UseHttps         NullableString `json:"UseHttps,omitempty"`
	FriendlyName     NullableString `json:"FriendlyName,omitempty"`
	DeployFolder     NullableString `json:"DeployFolder,omitempty"`
	ConfigLocation   NullableString `json:"ConfigLocation,omitempty"`
	ConfigTypeName   NullableString `json:"ConfigTypeName,omitempty"`
}

type StoreDiscoveryServiceModel struct {
	ServiceRecord          []NullableString `json:"ServiceRecord,omitempty"`
	StoreEndpoint          NullableString   `json:"StoreEndpoint,omitempty"`
	AuthenticationEndpoint NullableString   `json:"AuthenticationEndpoint,omitempty"`
}
type StoreFarmsConfiguration struct {
	Name                               NullableString   `json:"Name,omitempty"`
	EnableFileTypeAssociation          NullableBool     `json:"EnableFileTypeAssociation,omitempty"`
	PooledSockets                      NullableBool     `json:"PooledSockets,omitempty"`
	ServerCommunicationAttempts        NullableInt32    `json:"ServerCommunicationAttempts,omitempty"`
	CommunicationTimeout               TimeModel        `json:"CommunicationTimeout,omitempty"`
	ConnectionTimeout                  TimeModel        `json:"ConnectionTimeout,omitempty"`
	CertRevocationPolicy               NullableString   `json:"CertRevocationPolicy,omitempty"`
	MultiFarmAuthenticationMode        NullableString   `json:"MultiFarmAuthenticationMode,omitempty"`
	LeasingStatusExpiryLeasing         TimeModel        `json:"LeasingStatusExpiryLeasing,omitempty"`
	LeasingStatusExpiryFailed          TimeModel        `json:"LeasingStatusExpiryFailed,omitempty"`
	LeasingStatusExpiryPending         TimeModel        `json:"LeasingStatusExpiryPending,omitempty"`
	BackgroundHealthCheckPollingPeriod TimeModel        `json:"BackgroundHealthCheckPollingPeriod,omitempty"`
	AdvancedHealthCheck                NullableBool     `json:"AdvancedHealthCheck,omitempty"`
	Farms                              []NullableString `json:"Farms,omitempty"`
}

type StoreSubscriptionModel struct {
	Enabled          NullableBool   `json:"Enabled"`
	ConnectionString NullableString `json:"ConnectionString"`
	FactoryName      NullableString `json:"FactoryName"`
}

type StoreGatewayServiceModel struct {
	Enabled                  NullableBool     `json:"Enabled"`
	CustomerId               NullableString   `json:"CustomerId"`
	InstanceId               NullableString   `json:"InstanceId"`
	GetGatewayServiceUrl     NullableString   `json:"GetGatewayServiceUrl"`
	PrivateKey               NullableString   `json:"PrivateKey"`
	ServiceName              NullableString   `json:"ServiceName"`
	SecureTicketAuthorityUrl NullableString   `json:"SecureTicketAuthorityUrl"`
	SecureTicketLifetime     TimeModel        `json:"SecureTicketLifetime"`
	SessionReliability       NullableBool     `json:"SessionReliability"`
	HandleZones              []NullableString `json:"HandleZones"`
	IgnoreZones              []NullableString `json:"IgnoreZones"`
	GatewayDiscoveryProtocol NullableString   `json:"GatewayDiscoveryProtocol"`
}

type StoreDiscoveryService struct {
	ServiceRecord          NullableString `json:"ServiceRecord"`
	StoreEndpoint          NullableString `json:"StoreEndpoint"`
	AuthenticationEndpoint NullableString `json:"AuthenticationEndpoint"`
}

type StorePNAModel struct {
	ChangePasswordAllowed         NullableString                `json:"ChangePasswordAllowed"`
	LogonMethod                   NullableInt                   `json:"LogonMethod"`
	ChangePasswordMethod          NullableString                `json:"ChangePasswordMethod"`
	ChangePasswordUrl             NullableString                `json:"ChangePasswordUrl"`
	IsEnabled                     NullableBool                  `json:"IsEnabled"`
	KerberosEnabled               NullableBool                  `json:"KerberosEnabled"`
	AccessConditionsTrustSettings AccessConditionsTrustSettings `json:"AccessConditionsTrustSettings"`
}

type AccessConditionsTrustSettings struct {
	HeaderValidityPeriod   NullableInt    `json:"HeaderValidityPeriod"`
	AccessConditionsTrusts NullableString `json:"AccessConditionsTrusts"`
}

type StoreResources struct {
	ApplicationAccessMethods []NullableString      `json:"ApplicationAccessMethods"`
	Launch                   StoreLaunchModel      `json:"Launch"`
	Enumeration              StoreEnumerationModel `json:"Enumeration"`
}

type StoreLaunchModel struct {
	SetNoLoadBiasFlag                      NullableBool   `json:"SetNoLoadBiasFlag"`
	AddressResolutionType                  NullableInt    `json:"AddressResolutionType"`
	RequestICAClientSecureChannel          NullableInt    `json:"RequestICAClientSecureChannel"`
	IgnoreClientProvidedClientAddress      NullableBool   `json:"IgnoreClientProvidedClientAddress"`
	OverlayAutoLoginCredentialsWithTicket  NullableBool   `json:"OverlayAutoLoginCredentialsWithTicket"`
	OverrideIcaClientName                  NullableBool   `json:"OverrideIcaClientName"`
	RequireLaunchReference                 NullableBool   `json:"RequireLaunchReference"`
	AllowFontSmoothing                     NullableBool   `json:"AllowFontSmoothing"`
	ShowDesktopViewer                      NullableBool   `json:"ShowDesktopViewer"`
	AllowSpecialFolderRedirection          NullableBool   `json:"AllowSpecialFolderRedirection"`
	ClientProxyPolicy                      NullableString `json:"ClientProxyPolicy"`
	RoutingPolicy                          NullableString `json:"RoutingPolicy"`
	VdaLogonDataProviderName               NullableString `json:"VdaLogonDataProviderName"`
	FederatedAuthenticationServiceFailover NullableBool   `json:"FederatedAuthenticationServiceFailover"`
}

type StoreEnumerationModel struct {
	EnhancedEnumeration                          NullableBool   `json:"EnhancedEnumeration"`
	MaximumConcurrentEnumerations                NullableInt    `json:"MaximumConcurrentEnumerations"`
	MinimumFarmsRequiredForConcurrentEnumeration NullableInt    `json:"MinimumFarmsRequiredForConcurrentEnumeration"`
	RequestFullIconData                          NullableInt    `json:"RequestFullIconData"`
	RequestedHighColorIcons                      NullableString `json:"RequestedHighColorIcons"`
	FilterByTypesInclude                         NullableString `json:"FilterByTypesInclude"`
	FilterByKeywordsInclude                      NullableString `json:"FilterByKeywordsInclude"`
	FilterByKeywordsExclude                      NullableString `json:"FilterByKeywordsExclude"`
	TreatDesktopsAsApps                          NullableBool   `json:"TreatDesktopsAsApps"`
	Authenticator                                NullableString `json:"Authenticator"`
}

type StoreConfiguration struct {
	AppSettings               StoreAppSettings       `json:"AppSettings"`
	ConnectionStrings         StoreConnectionStrings `json:"ConnectionStrings"`
	FilePath                  NullableString         `json:"FilePath"`
	HasFile                   NullableBool           `json:"HasFile"`
	Locations                 []NullableString       `json:"Locations"`
	EvaluationContext         StoreEvaluationContext `json:"EvaluationContext"`
	RootSectionGroup          StoreRootSectionGroup  `json:"RootSectionGroup"`
	Sections                  []NullableString       `json:"Sections"`
	SectionGroups             []NullableString       `json:"SectionGroups"`
	NamespaceDeclared         NullableBool           `json:"NamespaceDeclared"`
	TypeStringTransformer     NullableString         `json:"TypeStringTransformer"`
	AssemblyStringTransformer NullableString         `json:"AssemblyStringTransformer"`
	TargetFramework           NullableString         `json:"TargetFramework"`
}

type StoreAppSettings struct {
	Settings                NullableString `json:"Settings"`
	File                    NullableString `json:"File"`
	SectionInformation      NullableString `json:"SectionInformation"`
	LockAttributes          NullableString `json:"LockAttributes"`
	LockAllAttributesExcept NullableString `json:"LockAllAttributesExcept"`
	LockElements            NullableString `json:"LockElements"`
	LockAllElementsExcept   NullableString `json:"LockAllElementsExcept"`
	LockItem                NullableBool   `json:"LockItem"`
	ElementInformation      NullableString `json:"ElementInformation"`
	CurrentConfiguration    NullableString `json:"CurrentConfiguration"`
}

type StoreConnectionStrings struct {
	ConnectionStrings       NullableString `json:"ConnectionStrings"`
	SectionInformation      NullableString `json:"SectionInformation"`
	LockAttributes          NullableString `json:"LockAttributes"`
	LockAllAttributesExcept NullableString `json:"LockAllAttributesExcept"`
	LockElements            NullableString `json:"LockElements"`
	LockAllElementsExcept   NullableString `json:"LockAllElementsExcept"`
	LockItem                NullableBool   `json:"LockItem"`
	ElementInformation      NullableString `json:"ElementInformation"`
	CurrentConfiguration    NullableString `json:"CurrentConfiguration"`
}

type StoreRootSectionGroup struct {
	IsDeclared            NullableBool   `json:"IsDeclared"`
	IsDeclarationRequired NullableBool   `json:"IsDeclarationRequired"`
	SectionGroupName      NullableString `json:"SectionGroupName"`
	Name                  NullableString `json:"Name"`
	Type                  NullableString `json:"Type"`
	Sections              NullableString `json:"Sections"`
	SectionGroups         NullableString `json:"SectionGroups"`
}

type StoreEvaluationContext struct {
	HostingContext NullableString `json:"HostingContext"`
	IsMachineLevel NullableBool   `json:"IsMachineLevel"`
}

type StoreApplicationInitializationModel struct {
	SkipManagedModules NullableBool `json:"SkipManagedModules"`
}

type StoreDiagnosticsModel struct {
	SharedListeners StoreSharedListenersModel `json:"SharedListeners"`
}

type StoreSharedListenersModel struct {
	citrixTraceListener NullableString `json:"citrixTraceListener"`
}

type StoreCredentialWalletModel struct {
	WalletName NullableString `json:"WalletName"`
}

type StoreHostingEnvModel struct {
	ShadowCopyBinAssemblies NullableBool `json:"ShadowCopyBinAssemblies"`
}

type StoreServiceTokenCertificateModel struct {
	Id         NullableString `json:"Id"`
	Thumbprint NullableString `json:"Thumbprint"`
}

type StoreSettingsModel struct {
	documentDbSubscriptions NullableString `json:"documentDbSubscriptions"`
}

type SubscriptionStoreClientModel struct {
	Enabled        NullableBool `json:"Enabled"`
	ClientEndPoint StoreEndpointModel
}

type StoreEndpointModel struct {
	Uri                NullableString `json:"Uri,omitempty"`                // The uri of the client endpoint
	TransferMode       NullableInt    `json:"TransferMode,omitempty"`       // The transfer mode of the client endpoint
	AuthenticationMode NullableInt    `json:"AuthenticationMode,omitempty"` // The authentication mode of the client endpoint
	WsSecurityEnabled  NullableBool   `json:"WsSecurityEnabled,omitempty"`  // The ws security enabled of the client endpoint
	ClientCertificate  NullableString `json:"ClientCertificate,omitempty"`  // The client certificate of the client endpoint
	OpenTimeOut        NullableString `json:"OpenTimeOut,omitempty"`        // The open time out of the client endpoint
	CloseTimeOut       NullableString `json:"CloseTimeOut,omitempty"`       // The close time out of the client endpoint
	SendTimeOut        NullableString `json:"SendTimeOut,omitempty"`        // The send time out of the client endpoint
	PortSharingEnabled NullableBool   `json:"PortSharingEnabled,omitempty"` // The port sharing enabled of the client endpoint
}

type StoreWorkflowModel struct {
	NFuseWorkflowAdaptors []NullableString `json:"NFuseWorkflowAdaptors"`
}

type StoreResourcesModel struct {
	CommonData []NullableString `json:"CommonData"`
}

type IcaLaunchDefaultsModel struct {
	ClientType   NullableString `json:"ClientType"`
	TemplateName NullableString `json:"TemplateName"`
}

type IcaFileSigningModel struct {
	Enabled       NullableBool   `json:"Enabled,omitempty"`              // The enabled of the request signing
	CertificateId NullableString `json:"SigningCertificateId,omitempty"` // The signing certificate id of the request signing
	HashAlgorithm NullableString `json:"Algorithm,omitempty"`            // The algorithm of the request signing
}

type DatabaseConnectionStringModel struct {
	Name             NullableString `json:"Name"`
	ConnectionString NullableString `json:"ConnectionString"`
	ProviderName     NullableString `json:"ProviderName"`
}

type StoreAppSettingsModel struct {
	Settings                []NullableString             `json:"Settings"`
	File                    NullableString               `json:"File"`
	SectionInformation      StoreSectionInformationModel `json:"SectionInformation"`
	LockAttributes          []NullableString             `json:"LockAttributes"`
	LockAllAttributesExcept []NullableString             `json:"LockAllAttributesExcept"`
	LockElements            []NullableString             `json:"LockElements"`
	LockAllElementsExcept   []NullableString             `json:"LockAllElementsExcept"`
	LockItem                NullableBool                 `json:"LockItem"`
	ElementInformation      StoreElementInformationModel `json:"ElementInformation"`
	CurrentConfiguration    StoreCurrentConfigModel      `json:"CurrentConfiguration"`
}

type SessionManagerSettingsModel struct {
	CertificateValidationMode    NullableString   `json:"CertificateValidationMode"`
	SessionManagerTrustedIssuers []NullableString `json:"SessionManagerTrustedIssuers"`
}

type StoreSectionInformationModel struct {
	SectionName                NullableString `json:"SectionName"`
	Name                       NullableString `json:"Name"`
	AllowDefinition            NullableInt    `json:"AllowDefinition"`
	AllowExeDefinition         NullableInt    `json:"AllowExeDefinition"`
	OverrideModeDefault        NullableInt    `json:"OverrideModeDefault"`
	AllowLocation              NullableBool   `json:"AllowLocation"`
	AllowOverride              NullableBool   `json:"AllowOverride"`
	OverrideMode               NullableInt    `json:"OverrideMode"`
	OverrideModeEffective      NullableInt    `json:"OverrideModeEffective"`
	ConfigSource               NullableString `json:"ConfigSource"`
	InheritInChildApplications NullableBool   `json:"InheritInChildApplications"`
	IsDeclared                 NullableBool   `json:"IsDeclared"`
	IsDeclarationRequired      NullableBool   `json:"IsDeclarationRequired"`
	IsLocked                   NullableBool   `json:"IsLocked"`
	IsProtected                NullableBool   `json:"IsProtected"`
	ConfigurationBuilder       NullableString `json:"ConfigurationBuilder"`
	ProtectionProvider         NullableString `json:"ProtectionProvider"`
	RestartOnExternalChanges   NullableBool   `json:"RestartOnExternalChanges"`
	RequirePermission          NullableBool   `json:"RequirePermission"`
	Type                       NullableString `json:"Type"`
	ForceSave                  NullableBool   `json:"ForceSave"`
}

type StoreElementInformationModel struct {
	Properties   NullableString `json:"Properties"`
	IsPresent    NullableBool   `json:"IsPresent"`
	IsLocked     NullableBool   `json:"IsLocked"`
	IsCollection NullableBool   `json:"IsCollection"`
	Source       NullableString `json:"Source"`
	LineNumber   NullableInt    `json:"LineNumber"`
	Type         NullableString `json:"Type"`
	Validator    NullableString `json:"Validator"`
	Errors       NullableString `json:"Errors"`
}

type StoreCurrentConfigModel struct {
	AppSettings               NullableString `json:"AppSettings"`
	ConnectionStrings         NullableString `json:"ConnectionStrings"`
	FilePath                  NullableString `json:"FilePath"`
	HasFile                   NullableBool   `json:"HasFile"`
	Locations                 NullableString `json:"Locations"`
	EvaluationContext         NullableString `json:"EvaluationContext"`
	RootSectionGroup          NullableString `json:"RootSectionGroup"`
	Sections                  NullableString `json:"Sections"`
	SectionGroups             NullableString `json:"SectionGroups"`
	NamespaceDeclared         NullableBool   `json:"NamespaceDeclared"`
	TypeStringTransformer     NullableString `json:"TypeStringTransformer"`
	AssemblyStringTransformer NullableString `json:"AssemblyStringTransformer"`
	TargetFramework           NullableString `json:"TargetFramework"`
}

type StoreWebApplicationModel struct {
	WebSite            WebSite          `json:"WebSite"`
	VirtualPath        NullableString   `json:"VirtualPath"`
	Name               NullableString   `json:"Name"`
	AppPool            NullableString   `json:"AppPool"`
	Folder             NullableString   `json:"Folder"`
	VirtualDirectories []NullableString `json:"VirtualDirectories"`
}

type WebSite struct {
	PhysicalPath        NullableString `json:"PhysicalPath"`
	DefaultApplPoolName NullableString `json:"DefaultApplPoolName"`
	Bindings            NullableString `json:"Bindings"`
	Name                NullableString `json:"Name"`
	Id                  NullableInt    `json:"Id"`
	Applications        NullableString `json:"Applications"`
}
