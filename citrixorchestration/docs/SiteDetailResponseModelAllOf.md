# SiteDetailResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrchestationApiVersion** | **int32** | Orchestration API version | 
**AppProtectionAuthorized** | Pointer to **bool** |  | [optional] 
**LicenseServerName** | Pointer to **string** | Hostname of the licensing server for the site. Will be null for XenApp &amp; XenDesktop service. | [optional] 
**LicenseServerPort** | Pointer to **int32** | License server port. Will be null for XenApp &amp; XenDesktop service. | [optional] 
**LicenseServerUri** | Pointer to **string** | License server URI. Will be null for XenApp &amp; XenDesktop service. | [optional] 
**LicensingModel** | Pointer to [**LicenseModel**](LicenseModel.md) |  | [optional] 
**SiteConfigurationComplete** | Pointer to **bool** | Indicates whether the site&#39;s initial configuration is complete. | [optional] 
**PrimaryZone** | [**SiteDetailResponseModelAllOfPrimaryZone**](SiteDetailResponseModelAllOfPrimaryZone.md) |  | 
**ProductCode** | **string** | Product code for the site license. | 
**ProductEdition** | **string** | Product edition for the site license. | 
**ProductVersion** | **string** | Product version. | 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of Site. | [optional] 
**SiteServices** | [**[]SiteServiceResponseModel**](SiteServiceResponseModel.md) | Services running within the site. | 
**EnabledFeatures** | **[]string** | Enabled features within the site. | 
**BaseOU** | Pointer to **string** | The objectGUID property identifying the base OU in Active Directory used for desktop registrations. | [optional] 
**ColorDepth** | [**ColorDepth**](ColorDepth.md) |  | 
**ConnectionLeasingEnabled** | Pointer to **bool** | Indicates whether connection leasing is enabled. | [optional] 
**DefaultMinimumFunctionalLevel** | Pointer to [**FunctionalLevel**](FunctionalLevel.md) |  | [optional] 
**DefaultDesktopIconId** | **string** | The default desktop icon used for new delivery groups. | 
**DnsResolutionEnabled** | Pointer to **bool** | Indicates whether numeric IP addresses or DNS names are present in ICA files. | [optional] 
**LicenseGraceSessionsRemaining** | Pointer to **int32** | The number of licensing grace sessions remaining. | [optional] 
**LicensedSessionsActive** | Pointer to **int32** | The number of active, licensed sessions. | [optional] 
**LicensingGraceHoursLeft** | Pointer to **int32** | When in an active licensing grace period, indicates the  number of hours remaining in the grace period. Will be null for XenApp &amp; XenDesktop service. | [optional] 
**LicensingGracePeriodActive** | Pointer to **bool** | Indicates whether a licensing grace period is active. Will be null for XenApp &amp; XenDesktop service. | [optional] 
**PeakConcurrentLicenseUsers** | Pointer to **int32** | The peak number of concurrent licensed users. Only applies if LicensingModel is Concurrent. | [optional] 
**ReuseMachinesWithoutShutdownInOutageAllowed** | Pointer to **bool** | Allow sessions to launch on shared machines without an intermediate shutdown while the system is operating in outage mode. | [optional] 
**DefaultSecureIcaRequired** | Pointer to **bool** | The default SecureICA usage for new delivery groups. | [optional] 
**TotalUniqueLicenseUsers** | Pointer to **int32** | The total count of unique license users. | [optional] 
**TrustRequestsSentToTheXmlServicePort** | Pointer to **bool** | Indicates whether the XML service trusts the caller to specify user identity. | [optional] 
**LicensingAlerts** | Pointer to [**[]LicensingAlertResponseModel**](LicensingAlertResponseModel.md) | Get active licensing alerts in the site. Will be null for XenApp &amp; XenDesktop service. | [optional] 

## Methods

### NewSiteDetailResponseModelAllOf

`func NewSiteDetailResponseModelAllOf(orchestationApiVersion int32, primaryZone SiteDetailResponseModelAllOfPrimaryZone, productCode string, productEdition string, productVersion string, siteServices []SiteServiceResponseModel, enabledFeatures []string, colorDepth ColorDepth, defaultDesktopIconId string, ) *SiteDetailResponseModelAllOf`

NewSiteDetailResponseModelAllOf instantiates a new SiteDetailResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteDetailResponseModelAllOfWithDefaults

`func NewSiteDetailResponseModelAllOfWithDefaults() *SiteDetailResponseModelAllOf`

NewSiteDetailResponseModelAllOfWithDefaults instantiates a new SiteDetailResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrchestationApiVersion

`func (o *SiteDetailResponseModelAllOf) GetOrchestationApiVersion() int32`

GetOrchestationApiVersion returns the OrchestationApiVersion field if non-nil, zero value otherwise.

### GetOrchestationApiVersionOk

`func (o *SiteDetailResponseModelAllOf) GetOrchestationApiVersionOk() (*int32, bool)`

GetOrchestationApiVersionOk returns a tuple with the OrchestationApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestationApiVersion

`func (o *SiteDetailResponseModelAllOf) SetOrchestationApiVersion(v int32)`

SetOrchestationApiVersion sets OrchestationApiVersion field to given value.


### GetAppProtectionAuthorized

`func (o *SiteDetailResponseModelAllOf) GetAppProtectionAuthorized() bool`

GetAppProtectionAuthorized returns the AppProtectionAuthorized field if non-nil, zero value otherwise.

### GetAppProtectionAuthorizedOk

`func (o *SiteDetailResponseModelAllOf) GetAppProtectionAuthorizedOk() (*bool, bool)`

GetAppProtectionAuthorizedOk returns a tuple with the AppProtectionAuthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProtectionAuthorized

`func (o *SiteDetailResponseModelAllOf) SetAppProtectionAuthorized(v bool)`

SetAppProtectionAuthorized sets AppProtectionAuthorized field to given value.

### HasAppProtectionAuthorized

`func (o *SiteDetailResponseModelAllOf) HasAppProtectionAuthorized() bool`

HasAppProtectionAuthorized returns a boolean if a field has been set.

### GetLicenseServerName

`func (o *SiteDetailResponseModelAllOf) GetLicenseServerName() string`

GetLicenseServerName returns the LicenseServerName field if non-nil, zero value otherwise.

### GetLicenseServerNameOk

`func (o *SiteDetailResponseModelAllOf) GetLicenseServerNameOk() (*string, bool)`

GetLicenseServerNameOk returns a tuple with the LicenseServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseServerName

`func (o *SiteDetailResponseModelAllOf) SetLicenseServerName(v string)`

SetLicenseServerName sets LicenseServerName field to given value.

### HasLicenseServerName

`func (o *SiteDetailResponseModelAllOf) HasLicenseServerName() bool`

HasLicenseServerName returns a boolean if a field has been set.

### GetLicenseServerPort

`func (o *SiteDetailResponseModelAllOf) GetLicenseServerPort() int32`

GetLicenseServerPort returns the LicenseServerPort field if non-nil, zero value otherwise.

### GetLicenseServerPortOk

`func (o *SiteDetailResponseModelAllOf) GetLicenseServerPortOk() (*int32, bool)`

GetLicenseServerPortOk returns a tuple with the LicenseServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseServerPort

`func (o *SiteDetailResponseModelAllOf) SetLicenseServerPort(v int32)`

SetLicenseServerPort sets LicenseServerPort field to given value.

### HasLicenseServerPort

`func (o *SiteDetailResponseModelAllOf) HasLicenseServerPort() bool`

HasLicenseServerPort returns a boolean if a field has been set.

### GetLicenseServerUri

`func (o *SiteDetailResponseModelAllOf) GetLicenseServerUri() string`

GetLicenseServerUri returns the LicenseServerUri field if non-nil, zero value otherwise.

### GetLicenseServerUriOk

`func (o *SiteDetailResponseModelAllOf) GetLicenseServerUriOk() (*string, bool)`

GetLicenseServerUriOk returns a tuple with the LicenseServerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseServerUri

`func (o *SiteDetailResponseModelAllOf) SetLicenseServerUri(v string)`

SetLicenseServerUri sets LicenseServerUri field to given value.

### HasLicenseServerUri

`func (o *SiteDetailResponseModelAllOf) HasLicenseServerUri() bool`

HasLicenseServerUri returns a boolean if a field has been set.

### GetLicensingModel

`func (o *SiteDetailResponseModelAllOf) GetLicensingModel() LicenseModel`

GetLicensingModel returns the LicensingModel field if non-nil, zero value otherwise.

### GetLicensingModelOk

`func (o *SiteDetailResponseModelAllOf) GetLicensingModelOk() (*LicenseModel, bool)`

GetLicensingModelOk returns a tuple with the LicensingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingModel

`func (o *SiteDetailResponseModelAllOf) SetLicensingModel(v LicenseModel)`

SetLicensingModel sets LicensingModel field to given value.

### HasLicensingModel

`func (o *SiteDetailResponseModelAllOf) HasLicensingModel() bool`

HasLicensingModel returns a boolean if a field has been set.

### GetSiteConfigurationComplete

`func (o *SiteDetailResponseModelAllOf) GetSiteConfigurationComplete() bool`

GetSiteConfigurationComplete returns the SiteConfigurationComplete field if non-nil, zero value otherwise.

### GetSiteConfigurationCompleteOk

`func (o *SiteDetailResponseModelAllOf) GetSiteConfigurationCompleteOk() (*bool, bool)`

GetSiteConfigurationCompleteOk returns a tuple with the SiteConfigurationComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteConfigurationComplete

`func (o *SiteDetailResponseModelAllOf) SetSiteConfigurationComplete(v bool)`

SetSiteConfigurationComplete sets SiteConfigurationComplete field to given value.

### HasSiteConfigurationComplete

`func (o *SiteDetailResponseModelAllOf) HasSiteConfigurationComplete() bool`

HasSiteConfigurationComplete returns a boolean if a field has been set.

### GetPrimaryZone

`func (o *SiteDetailResponseModelAllOf) GetPrimaryZone() SiteDetailResponseModelAllOfPrimaryZone`

GetPrimaryZone returns the PrimaryZone field if non-nil, zero value otherwise.

### GetPrimaryZoneOk

`func (o *SiteDetailResponseModelAllOf) GetPrimaryZoneOk() (*SiteDetailResponseModelAllOfPrimaryZone, bool)`

GetPrimaryZoneOk returns a tuple with the PrimaryZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryZone

`func (o *SiteDetailResponseModelAllOf) SetPrimaryZone(v SiteDetailResponseModelAllOfPrimaryZone)`

SetPrimaryZone sets PrimaryZone field to given value.


### GetProductCode

`func (o *SiteDetailResponseModelAllOf) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *SiteDetailResponseModelAllOf) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *SiteDetailResponseModelAllOf) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.


### GetProductEdition

`func (o *SiteDetailResponseModelAllOf) GetProductEdition() string`

GetProductEdition returns the ProductEdition field if non-nil, zero value otherwise.

### GetProductEditionOk

`func (o *SiteDetailResponseModelAllOf) GetProductEditionOk() (*string, bool)`

GetProductEditionOk returns a tuple with the ProductEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductEdition

`func (o *SiteDetailResponseModelAllOf) SetProductEdition(v string)`

SetProductEdition sets ProductEdition field to given value.


### GetProductVersion

`func (o *SiteDetailResponseModelAllOf) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *SiteDetailResponseModelAllOf) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *SiteDetailResponseModelAllOf) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### GetMetadata

`func (o *SiteDetailResponseModelAllOf) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SiteDetailResponseModelAllOf) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SiteDetailResponseModelAllOf) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SiteDetailResponseModelAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSiteServices

`func (o *SiteDetailResponseModelAllOf) GetSiteServices() []SiteServiceResponseModel`

GetSiteServices returns the SiteServices field if non-nil, zero value otherwise.

### GetSiteServicesOk

`func (o *SiteDetailResponseModelAllOf) GetSiteServicesOk() (*[]SiteServiceResponseModel, bool)`

GetSiteServicesOk returns a tuple with the SiteServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteServices

`func (o *SiteDetailResponseModelAllOf) SetSiteServices(v []SiteServiceResponseModel)`

SetSiteServices sets SiteServices field to given value.


### GetEnabledFeatures

`func (o *SiteDetailResponseModelAllOf) GetEnabledFeatures() []string`

GetEnabledFeatures returns the EnabledFeatures field if non-nil, zero value otherwise.

### GetEnabledFeaturesOk

`func (o *SiteDetailResponseModelAllOf) GetEnabledFeaturesOk() (*[]string, bool)`

GetEnabledFeaturesOk returns a tuple with the EnabledFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledFeatures

`func (o *SiteDetailResponseModelAllOf) SetEnabledFeatures(v []string)`

SetEnabledFeatures sets EnabledFeatures field to given value.


### GetBaseOU

`func (o *SiteDetailResponseModelAllOf) GetBaseOU() string`

GetBaseOU returns the BaseOU field if non-nil, zero value otherwise.

### GetBaseOUOk

`func (o *SiteDetailResponseModelAllOf) GetBaseOUOk() (*string, bool)`

GetBaseOUOk returns a tuple with the BaseOU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseOU

`func (o *SiteDetailResponseModelAllOf) SetBaseOU(v string)`

SetBaseOU sets BaseOU field to given value.

### HasBaseOU

`func (o *SiteDetailResponseModelAllOf) HasBaseOU() bool`

HasBaseOU returns a boolean if a field has been set.

### GetColorDepth

`func (o *SiteDetailResponseModelAllOf) GetColorDepth() ColorDepth`

GetColorDepth returns the ColorDepth field if non-nil, zero value otherwise.

### GetColorDepthOk

`func (o *SiteDetailResponseModelAllOf) GetColorDepthOk() (*ColorDepth, bool)`

GetColorDepthOk returns a tuple with the ColorDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorDepth

`func (o *SiteDetailResponseModelAllOf) SetColorDepth(v ColorDepth)`

SetColorDepth sets ColorDepth field to given value.


### GetConnectionLeasingEnabled

`func (o *SiteDetailResponseModelAllOf) GetConnectionLeasingEnabled() bool`

GetConnectionLeasingEnabled returns the ConnectionLeasingEnabled field if non-nil, zero value otherwise.

### GetConnectionLeasingEnabledOk

`func (o *SiteDetailResponseModelAllOf) GetConnectionLeasingEnabledOk() (*bool, bool)`

GetConnectionLeasingEnabledOk returns a tuple with the ConnectionLeasingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionLeasingEnabled

`func (o *SiteDetailResponseModelAllOf) SetConnectionLeasingEnabled(v bool)`

SetConnectionLeasingEnabled sets ConnectionLeasingEnabled field to given value.

### HasConnectionLeasingEnabled

`func (o *SiteDetailResponseModelAllOf) HasConnectionLeasingEnabled() bool`

HasConnectionLeasingEnabled returns a boolean if a field has been set.

### GetDefaultMinimumFunctionalLevel

`func (o *SiteDetailResponseModelAllOf) GetDefaultMinimumFunctionalLevel() FunctionalLevel`

GetDefaultMinimumFunctionalLevel returns the DefaultMinimumFunctionalLevel field if non-nil, zero value otherwise.

### GetDefaultMinimumFunctionalLevelOk

`func (o *SiteDetailResponseModelAllOf) GetDefaultMinimumFunctionalLevelOk() (*FunctionalLevel, bool)`

GetDefaultMinimumFunctionalLevelOk returns a tuple with the DefaultMinimumFunctionalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMinimumFunctionalLevel

`func (o *SiteDetailResponseModelAllOf) SetDefaultMinimumFunctionalLevel(v FunctionalLevel)`

SetDefaultMinimumFunctionalLevel sets DefaultMinimumFunctionalLevel field to given value.

### HasDefaultMinimumFunctionalLevel

`func (o *SiteDetailResponseModelAllOf) HasDefaultMinimumFunctionalLevel() bool`

HasDefaultMinimumFunctionalLevel returns a boolean if a field has been set.

### GetDefaultDesktopIconId

`func (o *SiteDetailResponseModelAllOf) GetDefaultDesktopIconId() string`

GetDefaultDesktopIconId returns the DefaultDesktopIconId field if non-nil, zero value otherwise.

### GetDefaultDesktopIconIdOk

`func (o *SiteDetailResponseModelAllOf) GetDefaultDesktopIconIdOk() (*string, bool)`

GetDefaultDesktopIconIdOk returns a tuple with the DefaultDesktopIconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDesktopIconId

`func (o *SiteDetailResponseModelAllOf) SetDefaultDesktopIconId(v string)`

SetDefaultDesktopIconId sets DefaultDesktopIconId field to given value.


### GetDnsResolutionEnabled

`func (o *SiteDetailResponseModelAllOf) GetDnsResolutionEnabled() bool`

GetDnsResolutionEnabled returns the DnsResolutionEnabled field if non-nil, zero value otherwise.

### GetDnsResolutionEnabledOk

`func (o *SiteDetailResponseModelAllOf) GetDnsResolutionEnabledOk() (*bool, bool)`

GetDnsResolutionEnabledOk returns a tuple with the DnsResolutionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsResolutionEnabled

`func (o *SiteDetailResponseModelAllOf) SetDnsResolutionEnabled(v bool)`

SetDnsResolutionEnabled sets DnsResolutionEnabled field to given value.

### HasDnsResolutionEnabled

`func (o *SiteDetailResponseModelAllOf) HasDnsResolutionEnabled() bool`

HasDnsResolutionEnabled returns a boolean if a field has been set.

### GetLicenseGraceSessionsRemaining

`func (o *SiteDetailResponseModelAllOf) GetLicenseGraceSessionsRemaining() int32`

GetLicenseGraceSessionsRemaining returns the LicenseGraceSessionsRemaining field if non-nil, zero value otherwise.

### GetLicenseGraceSessionsRemainingOk

`func (o *SiteDetailResponseModelAllOf) GetLicenseGraceSessionsRemainingOk() (*int32, bool)`

GetLicenseGraceSessionsRemainingOk returns a tuple with the LicenseGraceSessionsRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseGraceSessionsRemaining

`func (o *SiteDetailResponseModelAllOf) SetLicenseGraceSessionsRemaining(v int32)`

SetLicenseGraceSessionsRemaining sets LicenseGraceSessionsRemaining field to given value.

### HasLicenseGraceSessionsRemaining

`func (o *SiteDetailResponseModelAllOf) HasLicenseGraceSessionsRemaining() bool`

HasLicenseGraceSessionsRemaining returns a boolean if a field has been set.

### GetLicensedSessionsActive

`func (o *SiteDetailResponseModelAllOf) GetLicensedSessionsActive() int32`

GetLicensedSessionsActive returns the LicensedSessionsActive field if non-nil, zero value otherwise.

### GetLicensedSessionsActiveOk

`func (o *SiteDetailResponseModelAllOf) GetLicensedSessionsActiveOk() (*int32, bool)`

GetLicensedSessionsActiveOk returns a tuple with the LicensedSessionsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensedSessionsActive

`func (o *SiteDetailResponseModelAllOf) SetLicensedSessionsActive(v int32)`

SetLicensedSessionsActive sets LicensedSessionsActive field to given value.

### HasLicensedSessionsActive

`func (o *SiteDetailResponseModelAllOf) HasLicensedSessionsActive() bool`

HasLicensedSessionsActive returns a boolean if a field has been set.

### GetLicensingGraceHoursLeft

`func (o *SiteDetailResponseModelAllOf) GetLicensingGraceHoursLeft() int32`

GetLicensingGraceHoursLeft returns the LicensingGraceHoursLeft field if non-nil, zero value otherwise.

### GetLicensingGraceHoursLeftOk

`func (o *SiteDetailResponseModelAllOf) GetLicensingGraceHoursLeftOk() (*int32, bool)`

GetLicensingGraceHoursLeftOk returns a tuple with the LicensingGraceHoursLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingGraceHoursLeft

`func (o *SiteDetailResponseModelAllOf) SetLicensingGraceHoursLeft(v int32)`

SetLicensingGraceHoursLeft sets LicensingGraceHoursLeft field to given value.

### HasLicensingGraceHoursLeft

`func (o *SiteDetailResponseModelAllOf) HasLicensingGraceHoursLeft() bool`

HasLicensingGraceHoursLeft returns a boolean if a field has been set.

### GetLicensingGracePeriodActive

`func (o *SiteDetailResponseModelAllOf) GetLicensingGracePeriodActive() bool`

GetLicensingGracePeriodActive returns the LicensingGracePeriodActive field if non-nil, zero value otherwise.

### GetLicensingGracePeriodActiveOk

`func (o *SiteDetailResponseModelAllOf) GetLicensingGracePeriodActiveOk() (*bool, bool)`

GetLicensingGracePeriodActiveOk returns a tuple with the LicensingGracePeriodActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingGracePeriodActive

`func (o *SiteDetailResponseModelAllOf) SetLicensingGracePeriodActive(v bool)`

SetLicensingGracePeriodActive sets LicensingGracePeriodActive field to given value.

### HasLicensingGracePeriodActive

`func (o *SiteDetailResponseModelAllOf) HasLicensingGracePeriodActive() bool`

HasLicensingGracePeriodActive returns a boolean if a field has been set.

### GetPeakConcurrentLicenseUsers

`func (o *SiteDetailResponseModelAllOf) GetPeakConcurrentLicenseUsers() int32`

GetPeakConcurrentLicenseUsers returns the PeakConcurrentLicenseUsers field if non-nil, zero value otherwise.

### GetPeakConcurrentLicenseUsersOk

`func (o *SiteDetailResponseModelAllOf) GetPeakConcurrentLicenseUsersOk() (*int32, bool)`

GetPeakConcurrentLicenseUsersOk returns a tuple with the PeakConcurrentLicenseUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakConcurrentLicenseUsers

`func (o *SiteDetailResponseModelAllOf) SetPeakConcurrentLicenseUsers(v int32)`

SetPeakConcurrentLicenseUsers sets PeakConcurrentLicenseUsers field to given value.

### HasPeakConcurrentLicenseUsers

`func (o *SiteDetailResponseModelAllOf) HasPeakConcurrentLicenseUsers() bool`

HasPeakConcurrentLicenseUsers returns a boolean if a field has been set.

### GetReuseMachinesWithoutShutdownInOutageAllowed

`func (o *SiteDetailResponseModelAllOf) GetReuseMachinesWithoutShutdownInOutageAllowed() bool`

GetReuseMachinesWithoutShutdownInOutageAllowed returns the ReuseMachinesWithoutShutdownInOutageAllowed field if non-nil, zero value otherwise.

### GetReuseMachinesWithoutShutdownInOutageAllowedOk

`func (o *SiteDetailResponseModelAllOf) GetReuseMachinesWithoutShutdownInOutageAllowedOk() (*bool, bool)`

GetReuseMachinesWithoutShutdownInOutageAllowedOk returns a tuple with the ReuseMachinesWithoutShutdownInOutageAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseMachinesWithoutShutdownInOutageAllowed

`func (o *SiteDetailResponseModelAllOf) SetReuseMachinesWithoutShutdownInOutageAllowed(v bool)`

SetReuseMachinesWithoutShutdownInOutageAllowed sets ReuseMachinesWithoutShutdownInOutageAllowed field to given value.

### HasReuseMachinesWithoutShutdownInOutageAllowed

`func (o *SiteDetailResponseModelAllOf) HasReuseMachinesWithoutShutdownInOutageAllowed() bool`

HasReuseMachinesWithoutShutdownInOutageAllowed returns a boolean if a field has been set.

### GetDefaultSecureIcaRequired

`func (o *SiteDetailResponseModelAllOf) GetDefaultSecureIcaRequired() bool`

GetDefaultSecureIcaRequired returns the DefaultSecureIcaRequired field if non-nil, zero value otherwise.

### GetDefaultSecureIcaRequiredOk

`func (o *SiteDetailResponseModelAllOf) GetDefaultSecureIcaRequiredOk() (*bool, bool)`

GetDefaultSecureIcaRequiredOk returns a tuple with the DefaultSecureIcaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSecureIcaRequired

`func (o *SiteDetailResponseModelAllOf) SetDefaultSecureIcaRequired(v bool)`

SetDefaultSecureIcaRequired sets DefaultSecureIcaRequired field to given value.

### HasDefaultSecureIcaRequired

`func (o *SiteDetailResponseModelAllOf) HasDefaultSecureIcaRequired() bool`

HasDefaultSecureIcaRequired returns a boolean if a field has been set.

### GetTotalUniqueLicenseUsers

`func (o *SiteDetailResponseModelAllOf) GetTotalUniqueLicenseUsers() int32`

GetTotalUniqueLicenseUsers returns the TotalUniqueLicenseUsers field if non-nil, zero value otherwise.

### GetTotalUniqueLicenseUsersOk

`func (o *SiteDetailResponseModelAllOf) GetTotalUniqueLicenseUsersOk() (*int32, bool)`

GetTotalUniqueLicenseUsersOk returns a tuple with the TotalUniqueLicenseUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUniqueLicenseUsers

`func (o *SiteDetailResponseModelAllOf) SetTotalUniqueLicenseUsers(v int32)`

SetTotalUniqueLicenseUsers sets TotalUniqueLicenseUsers field to given value.

### HasTotalUniqueLicenseUsers

`func (o *SiteDetailResponseModelAllOf) HasTotalUniqueLicenseUsers() bool`

HasTotalUniqueLicenseUsers returns a boolean if a field has been set.

### GetTrustRequestsSentToTheXmlServicePort

`func (o *SiteDetailResponseModelAllOf) GetTrustRequestsSentToTheXmlServicePort() bool`

GetTrustRequestsSentToTheXmlServicePort returns the TrustRequestsSentToTheXmlServicePort field if non-nil, zero value otherwise.

### GetTrustRequestsSentToTheXmlServicePortOk

`func (o *SiteDetailResponseModelAllOf) GetTrustRequestsSentToTheXmlServicePortOk() (*bool, bool)`

GetTrustRequestsSentToTheXmlServicePortOk returns a tuple with the TrustRequestsSentToTheXmlServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustRequestsSentToTheXmlServicePort

`func (o *SiteDetailResponseModelAllOf) SetTrustRequestsSentToTheXmlServicePort(v bool)`

SetTrustRequestsSentToTheXmlServicePort sets TrustRequestsSentToTheXmlServicePort field to given value.

### HasTrustRequestsSentToTheXmlServicePort

`func (o *SiteDetailResponseModelAllOf) HasTrustRequestsSentToTheXmlServicePort() bool`

HasTrustRequestsSentToTheXmlServicePort returns a boolean if a field has been set.

### GetLicensingAlerts

`func (o *SiteDetailResponseModelAllOf) GetLicensingAlerts() []LicensingAlertResponseModel`

GetLicensingAlerts returns the LicensingAlerts field if non-nil, zero value otherwise.

### GetLicensingAlertsOk

`func (o *SiteDetailResponseModelAllOf) GetLicensingAlertsOk() (*[]LicensingAlertResponseModel, bool)`

GetLicensingAlertsOk returns a tuple with the LicensingAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingAlerts

`func (o *SiteDetailResponseModelAllOf) SetLicensingAlerts(v []LicensingAlertResponseModel)`

SetLicensingAlerts sets LicensingAlerts field to given value.

### HasLicensingAlerts

`func (o *SiteDetailResponseModelAllOf) HasLicensingAlerts() bool`

HasLicensingAlerts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


