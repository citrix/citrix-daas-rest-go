// Copyright © 2025. Citrix Systems, Inc.
package models

var _ MappedNullable = &CreateSTFWebReceiverRequestModel{}

type STFWebReceiverDetailModel struct {
	SiteId                             NullableInt               `json:"SiteId"`
	VirtualPath                        NullableString            `json:"VirtualPath"`
	FriendlyName                       NullableString            `json:"FriendlyName"`
	WebReceiverSettings                WebReceiverSettings       `json:"WebReceiverSettings"`
	WebReceiver                        WebReceiver               `json:"WebReceiver"`
	DefaultIISSite                     NullableBool              `json:"DefaultIISSite"`
	ServiceID                          NullableString            `json:"ServiceId"`
	ProducerService                    BasicServiceModel         `json:"ProducerService"`
	ConsumerService                    BasicServiceModel         `json:"ConsumerService"`
	ServiceTokenCertificate            CertificateModel          `json:"ServiceTokenCertificate"`
	AuthenticationServiceVirtualPath   NullableString            `json:"AuthenticationServiceVirtualPath"`
	StoreServiceVirtualPath            NullableString            `json:"StoreServiceVirtualPath"`
	HostingEnvironment                 HostingEnvironmentModel   `json:"HostingEnvironment"`
	RunAllManagedModulesForAllRequests NullableBool              `json:"RunAllManagedModulesForAllRequests"`
	FeatureData                        FeatureData               `json:"FeatureData"`
	Name                               NullableString            `json:"Name"`
	Hmacs                              []HmacsModel              `json:"Hmacs"`
	Certificates                       []CertificateModel        `json:"Certificates"`
	TokenManagers                      []TokenManagers           `json:"TokenManagers"`
	Routing                            RoutingModel              `json:"Routing"`
	CredentialWallet                   CredentialWalletModel     `json:"CredentialWallet"`
	WebApplication                     WebApplication            `json:"WebApplication"`
	ServiceRef                         NullableString            `json:"ServiceRef"`
	PhysicalPath                       NullableString            `json:"PhysicalPath"`
	Diagnostics                        Diagnostics               `json:"Diagnostics"`
	ApplicationInitialization          ApplicationInitialization `json:"ApplicationInitialization"`
	FeatureInstanceID                  NullableString            `json:"FeatureInstanceId"`
	Configuration                      Configuration             `json:"Configuration"`
	ConfigurationFile                  NullableString            `json:"ConfigurationFile"`
	TenantID                           NullableString            `json:"TenantId"`
}

type WebApplication struct {
	VirtualPath        string       `json:"VirtualPath"`
	Name               string       `json:"Name"`
	AppPool            string       `json:"AppPool"`
	Folder             string       `json:"Folder"`
	VirtualDirectories []string     `json:"VirtualDirectories"`
	WebSite            WebSiteModel `json:"WebSite"`
}

type WebSiteModel struct {
	PhysicalPath        string `json:"PhysicalPath"`
	DefaultApplPoolName string `json:"DefaultApplPoolName"`
	Bindings            string `json:"Bindings"`
	Name                string `json:"Name"`
	Id                  int    `json:"Id"`
	Applications        string `json:"Applications"`
}

type WebReceiverSettings struct {
}
type ClientSettings struct {
	AuthManager     NullableString `json:"AuthManager"`
	StoreProxy      NullableString `json:"StoreProxy"`
	PluginAssistant NullableString `json:"PluginAssistant"`
	UserInterface   NullableString `json:"UserInterface"`
	AppProtection   NullableString `json:"AppProtection"`
	Plugins         NullableString `json:"Plugins"`
}
type ServerSettings struct {
	Authentication          NullableString `json:"Authentication"`
	Communication           NullableString `json:"Communication"`
	ResourcesService        NullableString `json:"ResourcesService"`
	ApplicationShortcuts    NullableString `json:"ApplicationShortcuts"`
	DiscoveryService        NullableString `json:"DiscoveryService"`
	StrictTransportSecurity NullableString `json:"StrictTransportSecurity"`
	FeaturedAppGroups       NullableString `json:"FeaturedAppGroups"`
}
type WebReceiver struct {
	ClientSettings                    ClientSettings `json:"ClientSettings"`
	ServerSettings                    ServerSettings `json:"ServerSettings"`
	WebUIExperience                   NullableInt    `json:"WebUIExperience"`
	SessionStateTimeout               NullableInt    `json:"SessionStateTimeout"`
	SessionStateRedisConnectionString any            `json:"SessionStateRedisConnectionString"`
}

type FeatureData struct {
	SiteId           NullableString `json:"SiteId"`
	VirtualPath      NullableString `json:"VirtualPath"`
	ApplicationPool  NullableString `json:"ApplicationPool"`
	DefaultDocuments NullableString `json:"DefaultDocuments"`
	Name             NullableString `json:"Name"`
	UseHTTPS         NullableString `json:"UseHttps"`
	FriendlyName     NullableString `json:"Friendly Name"`
	DeployFolder     NullableString `json:"DeployFolder"`
	ConfigLocation   NullableString `json:"ConfigLocation"`
	ConfigTypeName   NullableString `json:"ConfigTypeName"`
}

type TokenManagers struct {
	DisplayName      NullableString `json:"DisplayName"`
	ID               NullableString `json:"Id"`
	TrustedIssuers   NullableString `json:"TrustedIssuers"`
	RelyingParties   NullableString `json:"RelyingParties"`
	AllowedAudiences NullableString `json:"AllowedAudiences"`
}

type SharedListeners struct {
	CitrixTraceListener NullableString `json:"citrixTraceListener"`
}
type Diagnostics struct {
	SharedListeners SharedListeners `json:"SharedListeners"`
}
type ApplicationInitialization struct {
	SkipManagedModules NullableBool `json:"SkipManagedModules"`
}
type AppSettings struct {
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
type ConnectionStrings struct {
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
type EvaluationContext struct {
	HostingContext NullableString `json:"HostingContext"`
	IsMachineLevel NullableBool   `json:"IsMachineLevel"`
}
type RootSectionGroup struct {
	IsDeclared            NullableBool   `json:"IsDeclared"`
	IsDeclarationRequired NullableBool   `json:"IsDeclarationRequired"`
	SectionGroupName      NullableString `json:"SectionGroupName"`
	Name                  NullableString `json:"Name"`
	Type                  any            `json:"Type"`
	Sections              NullableString `json:"Sections"`
	SectionGroups         NullableString `json:"SectionGroups"`
}
type Configuration struct {
	AppSettings               AppSettings       `json:"AppSettings"`
	ConnectionStrings         ConnectionStrings `json:"ConnectionStrings"`
	FilePath                  NullableString    `json:"FilePath"`
	HasFile                   NullableBool      `json:"HasFile"`
	Locations                 []any             `json:"Locations"`
	EvaluationContext         EvaluationContext `json:"EvaluationContext"`
	RootSectionGroup          RootSectionGroup  `json:"RootSectionGroup"`
	Sections                  []NullableString  `json:"Sections"`
	SectionGroups             []NullableString  `json:"SectionGroups"`
	NamespaceDeclared         NullableBool      `json:"NamespaceDeclared"`
	TypeStringTransformer     any               `json:"TypeStringTransformer"`
	AssemblyStringTransformer any               `json:"AssemblyStringTransformer"`
	TargetFramework           any               `json:"TargetFramework"`
}
