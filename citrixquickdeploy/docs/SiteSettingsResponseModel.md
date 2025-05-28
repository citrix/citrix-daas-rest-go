# SiteSettingsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseVerticalScalingForRdsLaunches** | Pointer to **bool** |  | [optional] 
**DnsResolutionEnabled** | Pointer to **bool** |  | [optional] 
**TrustRequestsSentToTheXmlServicePortEnabled** | Pointer to **bool** |  | [optional] 
**WebUiPolicySetEnabled** | Pointer to **bool** | The policy set in web UI is enabled or not | [optional] 
**ConsoleInactivityTimeoutMinutes** | Pointer to **int32** | The max number of minutes that console can be inactive. | [optional] 
**SupportedAuthenticators** | Pointer to [**Authenticator**](Authenticator.md) | Applicable only for On-Premise.    Determine the supported authentication methods. | [optional] 
**ShowSecurityBannerOnLoginPage** | Pointer to **bool** | Applicable only for On-Premise.    Indicate whether to display security banner on the login page. | [optional] 
**SecurityBannerOnLoginPage** | Pointer to **string** | Applicable only for On-Premise.    The custom security banner on the login page, max length is 5000. | [optional] 
**ShowConnectedControllerOnLoginPage** | Pointer to **bool** | Applicable only for On-Premise.    Indicate whether to display the DDC used for the current Studio session on the login page. | [optional] 
**AllowedCorsOriginsForIwa** | Pointer to **[]string** | Application only for On-Premise.    List of origins allowed to make cross-origin requests for Integrated Windows Authentication. | [optional] 
**MultiSites** | Pointer to [**[]MultiSiteModel**](MultiSiteModel.md) | Applicable only for On-Premise    Multiple sites configuration | [optional] 
**DefaultDomain** | Pointer to **string** | The default domain to be used in the login page. | [optional] 
**XmlServicesSslConfigurations** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | XML service configurations of enabling listening for https/http requests. Only applicable for On-Premise deployments. | [optional] 
**MultiTenantServicesAccess** | Pointer to [**MultiTenantServicesAccessModel**](MultiTenantServicesAccessModel.md) | Indicate whether Citrix Cloud Services are allowed to access the DaaS. | [optional] 

## Methods

### NewSiteSettingsResponseModel

`func NewSiteSettingsResponseModel() *SiteSettingsResponseModel`

NewSiteSettingsResponseModel instantiates a new SiteSettingsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteSettingsResponseModelWithDefaults

`func NewSiteSettingsResponseModelWithDefaults() *SiteSettingsResponseModel`

NewSiteSettingsResponseModelWithDefaults instantiates a new SiteSettingsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseVerticalScalingForRdsLaunches

`func (o *SiteSettingsResponseModel) GetUseVerticalScalingForRdsLaunches() bool`

GetUseVerticalScalingForRdsLaunches returns the UseVerticalScalingForRdsLaunches field if non-nil, zero value otherwise.

### GetUseVerticalScalingForRdsLaunchesOk

`func (o *SiteSettingsResponseModel) GetUseVerticalScalingForRdsLaunchesOk() (*bool, bool)`

GetUseVerticalScalingForRdsLaunchesOk returns a tuple with the UseVerticalScalingForRdsLaunches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVerticalScalingForRdsLaunches

`func (o *SiteSettingsResponseModel) SetUseVerticalScalingForRdsLaunches(v bool)`

SetUseVerticalScalingForRdsLaunches sets UseVerticalScalingForRdsLaunches field to given value.

### HasUseVerticalScalingForRdsLaunches

`func (o *SiteSettingsResponseModel) HasUseVerticalScalingForRdsLaunches() bool`

HasUseVerticalScalingForRdsLaunches returns a boolean if a field has been set.

### GetDnsResolutionEnabled

`func (o *SiteSettingsResponseModel) GetDnsResolutionEnabled() bool`

GetDnsResolutionEnabled returns the DnsResolutionEnabled field if non-nil, zero value otherwise.

### GetDnsResolutionEnabledOk

`func (o *SiteSettingsResponseModel) GetDnsResolutionEnabledOk() (*bool, bool)`

GetDnsResolutionEnabledOk returns a tuple with the DnsResolutionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsResolutionEnabled

`func (o *SiteSettingsResponseModel) SetDnsResolutionEnabled(v bool)`

SetDnsResolutionEnabled sets DnsResolutionEnabled field to given value.

### HasDnsResolutionEnabled

`func (o *SiteSettingsResponseModel) HasDnsResolutionEnabled() bool`

HasDnsResolutionEnabled returns a boolean if a field has been set.

### GetTrustRequestsSentToTheXmlServicePortEnabled

`func (o *SiteSettingsResponseModel) GetTrustRequestsSentToTheXmlServicePortEnabled() bool`

GetTrustRequestsSentToTheXmlServicePortEnabled returns the TrustRequestsSentToTheXmlServicePortEnabled field if non-nil, zero value otherwise.

### GetTrustRequestsSentToTheXmlServicePortEnabledOk

`func (o *SiteSettingsResponseModel) GetTrustRequestsSentToTheXmlServicePortEnabledOk() (*bool, bool)`

GetTrustRequestsSentToTheXmlServicePortEnabledOk returns a tuple with the TrustRequestsSentToTheXmlServicePortEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustRequestsSentToTheXmlServicePortEnabled

`func (o *SiteSettingsResponseModel) SetTrustRequestsSentToTheXmlServicePortEnabled(v bool)`

SetTrustRequestsSentToTheXmlServicePortEnabled sets TrustRequestsSentToTheXmlServicePortEnabled field to given value.

### HasTrustRequestsSentToTheXmlServicePortEnabled

`func (o *SiteSettingsResponseModel) HasTrustRequestsSentToTheXmlServicePortEnabled() bool`

HasTrustRequestsSentToTheXmlServicePortEnabled returns a boolean if a field has been set.

### GetWebUiPolicySetEnabled

`func (o *SiteSettingsResponseModel) GetWebUiPolicySetEnabled() bool`

GetWebUiPolicySetEnabled returns the WebUiPolicySetEnabled field if non-nil, zero value otherwise.

### GetWebUiPolicySetEnabledOk

`func (o *SiteSettingsResponseModel) GetWebUiPolicySetEnabledOk() (*bool, bool)`

GetWebUiPolicySetEnabledOk returns a tuple with the WebUiPolicySetEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUiPolicySetEnabled

`func (o *SiteSettingsResponseModel) SetWebUiPolicySetEnabled(v bool)`

SetWebUiPolicySetEnabled sets WebUiPolicySetEnabled field to given value.

### HasWebUiPolicySetEnabled

`func (o *SiteSettingsResponseModel) HasWebUiPolicySetEnabled() bool`

HasWebUiPolicySetEnabled returns a boolean if a field has been set.

### GetConsoleInactivityTimeoutMinutes

`func (o *SiteSettingsResponseModel) GetConsoleInactivityTimeoutMinutes() int32`

GetConsoleInactivityTimeoutMinutes returns the ConsoleInactivityTimeoutMinutes field if non-nil, zero value otherwise.

### GetConsoleInactivityTimeoutMinutesOk

`func (o *SiteSettingsResponseModel) GetConsoleInactivityTimeoutMinutesOk() (*int32, bool)`

GetConsoleInactivityTimeoutMinutesOk returns a tuple with the ConsoleInactivityTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleInactivityTimeoutMinutes

`func (o *SiteSettingsResponseModel) SetConsoleInactivityTimeoutMinutes(v int32)`

SetConsoleInactivityTimeoutMinutes sets ConsoleInactivityTimeoutMinutes field to given value.

### HasConsoleInactivityTimeoutMinutes

`func (o *SiteSettingsResponseModel) HasConsoleInactivityTimeoutMinutes() bool`

HasConsoleInactivityTimeoutMinutes returns a boolean if a field has been set.

### GetSupportedAuthenticators

`func (o *SiteSettingsResponseModel) GetSupportedAuthenticators() Authenticator`

GetSupportedAuthenticators returns the SupportedAuthenticators field if non-nil, zero value otherwise.

### GetSupportedAuthenticatorsOk

`func (o *SiteSettingsResponseModel) GetSupportedAuthenticatorsOk() (*Authenticator, bool)`

GetSupportedAuthenticatorsOk returns a tuple with the SupportedAuthenticators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedAuthenticators

`func (o *SiteSettingsResponseModel) SetSupportedAuthenticators(v Authenticator)`

SetSupportedAuthenticators sets SupportedAuthenticators field to given value.

### HasSupportedAuthenticators

`func (o *SiteSettingsResponseModel) HasSupportedAuthenticators() bool`

HasSupportedAuthenticators returns a boolean if a field has been set.

### GetShowSecurityBannerOnLoginPage

`func (o *SiteSettingsResponseModel) GetShowSecurityBannerOnLoginPage() bool`

GetShowSecurityBannerOnLoginPage returns the ShowSecurityBannerOnLoginPage field if non-nil, zero value otherwise.

### GetShowSecurityBannerOnLoginPageOk

`func (o *SiteSettingsResponseModel) GetShowSecurityBannerOnLoginPageOk() (*bool, bool)`

GetShowSecurityBannerOnLoginPageOk returns a tuple with the ShowSecurityBannerOnLoginPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSecurityBannerOnLoginPage

`func (o *SiteSettingsResponseModel) SetShowSecurityBannerOnLoginPage(v bool)`

SetShowSecurityBannerOnLoginPage sets ShowSecurityBannerOnLoginPage field to given value.

### HasShowSecurityBannerOnLoginPage

`func (o *SiteSettingsResponseModel) HasShowSecurityBannerOnLoginPage() bool`

HasShowSecurityBannerOnLoginPage returns a boolean if a field has been set.

### GetSecurityBannerOnLoginPage

`func (o *SiteSettingsResponseModel) GetSecurityBannerOnLoginPage() string`

GetSecurityBannerOnLoginPage returns the SecurityBannerOnLoginPage field if non-nil, zero value otherwise.

### GetSecurityBannerOnLoginPageOk

`func (o *SiteSettingsResponseModel) GetSecurityBannerOnLoginPageOk() (*string, bool)`

GetSecurityBannerOnLoginPageOk returns a tuple with the SecurityBannerOnLoginPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityBannerOnLoginPage

`func (o *SiteSettingsResponseModel) SetSecurityBannerOnLoginPage(v string)`

SetSecurityBannerOnLoginPage sets SecurityBannerOnLoginPage field to given value.

### HasSecurityBannerOnLoginPage

`func (o *SiteSettingsResponseModel) HasSecurityBannerOnLoginPage() bool`

HasSecurityBannerOnLoginPage returns a boolean if a field has been set.

### GetShowConnectedControllerOnLoginPage

`func (o *SiteSettingsResponseModel) GetShowConnectedControllerOnLoginPage() bool`

GetShowConnectedControllerOnLoginPage returns the ShowConnectedControllerOnLoginPage field if non-nil, zero value otherwise.

### GetShowConnectedControllerOnLoginPageOk

`func (o *SiteSettingsResponseModel) GetShowConnectedControllerOnLoginPageOk() (*bool, bool)`

GetShowConnectedControllerOnLoginPageOk returns a tuple with the ShowConnectedControllerOnLoginPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowConnectedControllerOnLoginPage

`func (o *SiteSettingsResponseModel) SetShowConnectedControllerOnLoginPage(v bool)`

SetShowConnectedControllerOnLoginPage sets ShowConnectedControllerOnLoginPage field to given value.

### HasShowConnectedControllerOnLoginPage

`func (o *SiteSettingsResponseModel) HasShowConnectedControllerOnLoginPage() bool`

HasShowConnectedControllerOnLoginPage returns a boolean if a field has been set.

### GetAllowedCorsOriginsForIwa

`func (o *SiteSettingsResponseModel) GetAllowedCorsOriginsForIwa() []string`

GetAllowedCorsOriginsForIwa returns the AllowedCorsOriginsForIwa field if non-nil, zero value otherwise.

### GetAllowedCorsOriginsForIwaOk

`func (o *SiteSettingsResponseModel) GetAllowedCorsOriginsForIwaOk() (*[]string, bool)`

GetAllowedCorsOriginsForIwaOk returns a tuple with the AllowedCorsOriginsForIwa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCorsOriginsForIwa

`func (o *SiteSettingsResponseModel) SetAllowedCorsOriginsForIwa(v []string)`

SetAllowedCorsOriginsForIwa sets AllowedCorsOriginsForIwa field to given value.

### HasAllowedCorsOriginsForIwa

`func (o *SiteSettingsResponseModel) HasAllowedCorsOriginsForIwa() bool`

HasAllowedCorsOriginsForIwa returns a boolean if a field has been set.

### GetMultiSites

`func (o *SiteSettingsResponseModel) GetMultiSites() []MultiSiteModel`

GetMultiSites returns the MultiSites field if non-nil, zero value otherwise.

### GetMultiSitesOk

`func (o *SiteSettingsResponseModel) GetMultiSitesOk() (*[]MultiSiteModel, bool)`

GetMultiSitesOk returns a tuple with the MultiSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSites

`func (o *SiteSettingsResponseModel) SetMultiSites(v []MultiSiteModel)`

SetMultiSites sets MultiSites field to given value.

### HasMultiSites

`func (o *SiteSettingsResponseModel) HasMultiSites() bool`

HasMultiSites returns a boolean if a field has been set.

### GetDefaultDomain

`func (o *SiteSettingsResponseModel) GetDefaultDomain() string`

GetDefaultDomain returns the DefaultDomain field if non-nil, zero value otherwise.

### GetDefaultDomainOk

`func (o *SiteSettingsResponseModel) GetDefaultDomainOk() (*string, bool)`

GetDefaultDomainOk returns a tuple with the DefaultDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDomain

`func (o *SiteSettingsResponseModel) SetDefaultDomain(v string)`

SetDefaultDomain sets DefaultDomain field to given value.

### HasDefaultDomain

`func (o *SiteSettingsResponseModel) HasDefaultDomain() bool`

HasDefaultDomain returns a boolean if a field has been set.

### GetXmlServicesSslConfigurations

`func (o *SiteSettingsResponseModel) GetXmlServicesSslConfigurations() []NameValueStringPairModel`

GetXmlServicesSslConfigurations returns the XmlServicesSslConfigurations field if non-nil, zero value otherwise.

### GetXmlServicesSslConfigurationsOk

`func (o *SiteSettingsResponseModel) GetXmlServicesSslConfigurationsOk() (*[]NameValueStringPairModel, bool)`

GetXmlServicesSslConfigurationsOk returns a tuple with the XmlServicesSslConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlServicesSslConfigurations

`func (o *SiteSettingsResponseModel) SetXmlServicesSslConfigurations(v []NameValueStringPairModel)`

SetXmlServicesSslConfigurations sets XmlServicesSslConfigurations field to given value.

### HasXmlServicesSslConfigurations

`func (o *SiteSettingsResponseModel) HasXmlServicesSslConfigurations() bool`

HasXmlServicesSslConfigurations returns a boolean if a field has been set.

### GetMultiTenantServicesAccess

`func (o *SiteSettingsResponseModel) GetMultiTenantServicesAccess() MultiTenantServicesAccessModel`

GetMultiTenantServicesAccess returns the MultiTenantServicesAccess field if non-nil, zero value otherwise.

### GetMultiTenantServicesAccessOk

`func (o *SiteSettingsResponseModel) GetMultiTenantServicesAccessOk() (*MultiTenantServicesAccessModel, bool)`

GetMultiTenantServicesAccessOk returns a tuple with the MultiTenantServicesAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiTenantServicesAccess

`func (o *SiteSettingsResponseModel) SetMultiTenantServicesAccess(v MultiTenantServicesAccessModel)`

SetMultiTenantServicesAccess sets MultiTenantServicesAccess field to given value.

### HasMultiTenantServicesAccess

`func (o *SiteSettingsResponseModel) HasMultiTenantServicesAccess() bool`

HasMultiTenantServicesAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


