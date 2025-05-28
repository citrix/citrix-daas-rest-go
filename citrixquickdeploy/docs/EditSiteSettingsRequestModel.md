# EditSiteSettingsRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseVerticalScalingForRdsLaunches** | Pointer to **bool** | Determines whether to use vertical scaling when considering RDS machines for launches.    Vertical scaling would saturate machines in the current pool rather than    send sessions to the least loaded machines.    This would be a trade in performance vs. cost, where vertical scaling would be less costly. | [optional] 
**DnsResolutionEnabled** | Pointer to **bool** | Changes whether ICA files returned by a broker service to    a user device contain the numeric IP address or the DNS name    of the desktop machine to which a session should be established. | [optional] 
**TrustRequestsSentToTheXmlServicePortEnabled** | Pointer to **bool** | Changes whether the XML Service (as used by Storefront)    implicitly trusts the originator of requests it receives,    or whether it fully authenticates them. | [optional] 
**WebUiPolicySetEnabled** | Pointer to **bool** | Determine whether the policy set in web UI is enabled or not | [optional] 
**ConsoleInactivityTimeoutMinutes** | Pointer to **int32** | The max number of minutes that console can be inactive. | [optional] 
**SupportedAuthenticators** | Pointer to [**Authenticator**](Authenticator.md) | Applicable only for On-Premise.    Determine the supported authentication methods. | [optional] 
**ShowSecurityBannerOnLoginPage** | Pointer to **bool** | Applicable only for On-Premise.    Indicate whether to display security banner on the login page. | [optional] 
**SecurityBannerOnLoginPage** | Pointer to **string** | Applicable only for On-Premise.    The custom security banner on the login page, max length is 5000. | [optional] 
**ShowConnectedControllerOnLoginPage** | Pointer to **bool** | Applicable only for On-Premise.    Indicate whether to display the DDC used for the current Studio session on the login page. | [optional] 
**AllowedCorsOriginsForIwa** | Pointer to **[]string** | Applicable only for On-Premise.    List of origins allowed to make cross-origin requests for Integrated Windows Authentication. | [optional] 
**MultiSites** | Pointer to [**[]MultiSiteModel**](MultiSiteModel.md) | Applicable only for On-Premise.    Multiple sites configuration. | [optional] 
**DefaultDomain** | Pointer to **string** | The default domain to be used in the login page. | [optional] 
**MultiTenantServicesAccess** | Pointer to [**MultiTenantServicesAccessModel**](MultiTenantServicesAccessModel.md) | Cloud Services access to DaaS. | [optional] 

## Methods

### NewEditSiteSettingsRequestModel

`func NewEditSiteSettingsRequestModel() *EditSiteSettingsRequestModel`

NewEditSiteSettingsRequestModel instantiates a new EditSiteSettingsRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditSiteSettingsRequestModelWithDefaults

`func NewEditSiteSettingsRequestModelWithDefaults() *EditSiteSettingsRequestModel`

NewEditSiteSettingsRequestModelWithDefaults instantiates a new EditSiteSettingsRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseVerticalScalingForRdsLaunches

`func (o *EditSiteSettingsRequestModel) GetUseVerticalScalingForRdsLaunches() bool`

GetUseVerticalScalingForRdsLaunches returns the UseVerticalScalingForRdsLaunches field if non-nil, zero value otherwise.

### GetUseVerticalScalingForRdsLaunchesOk

`func (o *EditSiteSettingsRequestModel) GetUseVerticalScalingForRdsLaunchesOk() (*bool, bool)`

GetUseVerticalScalingForRdsLaunchesOk returns a tuple with the UseVerticalScalingForRdsLaunches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVerticalScalingForRdsLaunches

`func (o *EditSiteSettingsRequestModel) SetUseVerticalScalingForRdsLaunches(v bool)`

SetUseVerticalScalingForRdsLaunches sets UseVerticalScalingForRdsLaunches field to given value.

### HasUseVerticalScalingForRdsLaunches

`func (o *EditSiteSettingsRequestModel) HasUseVerticalScalingForRdsLaunches() bool`

HasUseVerticalScalingForRdsLaunches returns a boolean if a field has been set.

### GetDnsResolutionEnabled

`func (o *EditSiteSettingsRequestModel) GetDnsResolutionEnabled() bool`

GetDnsResolutionEnabled returns the DnsResolutionEnabled field if non-nil, zero value otherwise.

### GetDnsResolutionEnabledOk

`func (o *EditSiteSettingsRequestModel) GetDnsResolutionEnabledOk() (*bool, bool)`

GetDnsResolutionEnabledOk returns a tuple with the DnsResolutionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsResolutionEnabled

`func (o *EditSiteSettingsRequestModel) SetDnsResolutionEnabled(v bool)`

SetDnsResolutionEnabled sets DnsResolutionEnabled field to given value.

### HasDnsResolutionEnabled

`func (o *EditSiteSettingsRequestModel) HasDnsResolutionEnabled() bool`

HasDnsResolutionEnabled returns a boolean if a field has been set.

### GetTrustRequestsSentToTheXmlServicePortEnabled

`func (o *EditSiteSettingsRequestModel) GetTrustRequestsSentToTheXmlServicePortEnabled() bool`

GetTrustRequestsSentToTheXmlServicePortEnabled returns the TrustRequestsSentToTheXmlServicePortEnabled field if non-nil, zero value otherwise.

### GetTrustRequestsSentToTheXmlServicePortEnabledOk

`func (o *EditSiteSettingsRequestModel) GetTrustRequestsSentToTheXmlServicePortEnabledOk() (*bool, bool)`

GetTrustRequestsSentToTheXmlServicePortEnabledOk returns a tuple with the TrustRequestsSentToTheXmlServicePortEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustRequestsSentToTheXmlServicePortEnabled

`func (o *EditSiteSettingsRequestModel) SetTrustRequestsSentToTheXmlServicePortEnabled(v bool)`

SetTrustRequestsSentToTheXmlServicePortEnabled sets TrustRequestsSentToTheXmlServicePortEnabled field to given value.

### HasTrustRequestsSentToTheXmlServicePortEnabled

`func (o *EditSiteSettingsRequestModel) HasTrustRequestsSentToTheXmlServicePortEnabled() bool`

HasTrustRequestsSentToTheXmlServicePortEnabled returns a boolean if a field has been set.

### GetWebUiPolicySetEnabled

`func (o *EditSiteSettingsRequestModel) GetWebUiPolicySetEnabled() bool`

GetWebUiPolicySetEnabled returns the WebUiPolicySetEnabled field if non-nil, zero value otherwise.

### GetWebUiPolicySetEnabledOk

`func (o *EditSiteSettingsRequestModel) GetWebUiPolicySetEnabledOk() (*bool, bool)`

GetWebUiPolicySetEnabledOk returns a tuple with the WebUiPolicySetEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUiPolicySetEnabled

`func (o *EditSiteSettingsRequestModel) SetWebUiPolicySetEnabled(v bool)`

SetWebUiPolicySetEnabled sets WebUiPolicySetEnabled field to given value.

### HasWebUiPolicySetEnabled

`func (o *EditSiteSettingsRequestModel) HasWebUiPolicySetEnabled() bool`

HasWebUiPolicySetEnabled returns a boolean if a field has been set.

### GetConsoleInactivityTimeoutMinutes

`func (o *EditSiteSettingsRequestModel) GetConsoleInactivityTimeoutMinutes() int32`

GetConsoleInactivityTimeoutMinutes returns the ConsoleInactivityTimeoutMinutes field if non-nil, zero value otherwise.

### GetConsoleInactivityTimeoutMinutesOk

`func (o *EditSiteSettingsRequestModel) GetConsoleInactivityTimeoutMinutesOk() (*int32, bool)`

GetConsoleInactivityTimeoutMinutesOk returns a tuple with the ConsoleInactivityTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleInactivityTimeoutMinutes

`func (o *EditSiteSettingsRequestModel) SetConsoleInactivityTimeoutMinutes(v int32)`

SetConsoleInactivityTimeoutMinutes sets ConsoleInactivityTimeoutMinutes field to given value.

### HasConsoleInactivityTimeoutMinutes

`func (o *EditSiteSettingsRequestModel) HasConsoleInactivityTimeoutMinutes() bool`

HasConsoleInactivityTimeoutMinutes returns a boolean if a field has been set.

### GetSupportedAuthenticators

`func (o *EditSiteSettingsRequestModel) GetSupportedAuthenticators() Authenticator`

GetSupportedAuthenticators returns the SupportedAuthenticators field if non-nil, zero value otherwise.

### GetSupportedAuthenticatorsOk

`func (o *EditSiteSettingsRequestModel) GetSupportedAuthenticatorsOk() (*Authenticator, bool)`

GetSupportedAuthenticatorsOk returns a tuple with the SupportedAuthenticators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedAuthenticators

`func (o *EditSiteSettingsRequestModel) SetSupportedAuthenticators(v Authenticator)`

SetSupportedAuthenticators sets SupportedAuthenticators field to given value.

### HasSupportedAuthenticators

`func (o *EditSiteSettingsRequestModel) HasSupportedAuthenticators() bool`

HasSupportedAuthenticators returns a boolean if a field has been set.

### GetShowSecurityBannerOnLoginPage

`func (o *EditSiteSettingsRequestModel) GetShowSecurityBannerOnLoginPage() bool`

GetShowSecurityBannerOnLoginPage returns the ShowSecurityBannerOnLoginPage field if non-nil, zero value otherwise.

### GetShowSecurityBannerOnLoginPageOk

`func (o *EditSiteSettingsRequestModel) GetShowSecurityBannerOnLoginPageOk() (*bool, bool)`

GetShowSecurityBannerOnLoginPageOk returns a tuple with the ShowSecurityBannerOnLoginPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSecurityBannerOnLoginPage

`func (o *EditSiteSettingsRequestModel) SetShowSecurityBannerOnLoginPage(v bool)`

SetShowSecurityBannerOnLoginPage sets ShowSecurityBannerOnLoginPage field to given value.

### HasShowSecurityBannerOnLoginPage

`func (o *EditSiteSettingsRequestModel) HasShowSecurityBannerOnLoginPage() bool`

HasShowSecurityBannerOnLoginPage returns a boolean if a field has been set.

### GetSecurityBannerOnLoginPage

`func (o *EditSiteSettingsRequestModel) GetSecurityBannerOnLoginPage() string`

GetSecurityBannerOnLoginPage returns the SecurityBannerOnLoginPage field if non-nil, zero value otherwise.

### GetSecurityBannerOnLoginPageOk

`func (o *EditSiteSettingsRequestModel) GetSecurityBannerOnLoginPageOk() (*string, bool)`

GetSecurityBannerOnLoginPageOk returns a tuple with the SecurityBannerOnLoginPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityBannerOnLoginPage

`func (o *EditSiteSettingsRequestModel) SetSecurityBannerOnLoginPage(v string)`

SetSecurityBannerOnLoginPage sets SecurityBannerOnLoginPage field to given value.

### HasSecurityBannerOnLoginPage

`func (o *EditSiteSettingsRequestModel) HasSecurityBannerOnLoginPage() bool`

HasSecurityBannerOnLoginPage returns a boolean if a field has been set.

### GetShowConnectedControllerOnLoginPage

`func (o *EditSiteSettingsRequestModel) GetShowConnectedControllerOnLoginPage() bool`

GetShowConnectedControllerOnLoginPage returns the ShowConnectedControllerOnLoginPage field if non-nil, zero value otherwise.

### GetShowConnectedControllerOnLoginPageOk

`func (o *EditSiteSettingsRequestModel) GetShowConnectedControllerOnLoginPageOk() (*bool, bool)`

GetShowConnectedControllerOnLoginPageOk returns a tuple with the ShowConnectedControllerOnLoginPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowConnectedControllerOnLoginPage

`func (o *EditSiteSettingsRequestModel) SetShowConnectedControllerOnLoginPage(v bool)`

SetShowConnectedControllerOnLoginPage sets ShowConnectedControllerOnLoginPage field to given value.

### HasShowConnectedControllerOnLoginPage

`func (o *EditSiteSettingsRequestModel) HasShowConnectedControllerOnLoginPage() bool`

HasShowConnectedControllerOnLoginPage returns a boolean if a field has been set.

### GetAllowedCorsOriginsForIwa

`func (o *EditSiteSettingsRequestModel) GetAllowedCorsOriginsForIwa() []string`

GetAllowedCorsOriginsForIwa returns the AllowedCorsOriginsForIwa field if non-nil, zero value otherwise.

### GetAllowedCorsOriginsForIwaOk

`func (o *EditSiteSettingsRequestModel) GetAllowedCorsOriginsForIwaOk() (*[]string, bool)`

GetAllowedCorsOriginsForIwaOk returns a tuple with the AllowedCorsOriginsForIwa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCorsOriginsForIwa

`func (o *EditSiteSettingsRequestModel) SetAllowedCorsOriginsForIwa(v []string)`

SetAllowedCorsOriginsForIwa sets AllowedCorsOriginsForIwa field to given value.

### HasAllowedCorsOriginsForIwa

`func (o *EditSiteSettingsRequestModel) HasAllowedCorsOriginsForIwa() bool`

HasAllowedCorsOriginsForIwa returns a boolean if a field has been set.

### GetMultiSites

`func (o *EditSiteSettingsRequestModel) GetMultiSites() []MultiSiteModel`

GetMultiSites returns the MultiSites field if non-nil, zero value otherwise.

### GetMultiSitesOk

`func (o *EditSiteSettingsRequestModel) GetMultiSitesOk() (*[]MultiSiteModel, bool)`

GetMultiSitesOk returns a tuple with the MultiSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSites

`func (o *EditSiteSettingsRequestModel) SetMultiSites(v []MultiSiteModel)`

SetMultiSites sets MultiSites field to given value.

### HasMultiSites

`func (o *EditSiteSettingsRequestModel) HasMultiSites() bool`

HasMultiSites returns a boolean if a field has been set.

### GetDefaultDomain

`func (o *EditSiteSettingsRequestModel) GetDefaultDomain() string`

GetDefaultDomain returns the DefaultDomain field if non-nil, zero value otherwise.

### GetDefaultDomainOk

`func (o *EditSiteSettingsRequestModel) GetDefaultDomainOk() (*string, bool)`

GetDefaultDomainOk returns a tuple with the DefaultDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDomain

`func (o *EditSiteSettingsRequestModel) SetDefaultDomain(v string)`

SetDefaultDomain sets DefaultDomain field to given value.

### HasDefaultDomain

`func (o *EditSiteSettingsRequestModel) HasDefaultDomain() bool`

HasDefaultDomain returns a boolean if a field has been set.

### GetMultiTenantServicesAccess

`func (o *EditSiteSettingsRequestModel) GetMultiTenantServicesAccess() MultiTenantServicesAccessModel`

GetMultiTenantServicesAccess returns the MultiTenantServicesAccess field if non-nil, zero value otherwise.

### GetMultiTenantServicesAccessOk

`func (o *EditSiteSettingsRequestModel) GetMultiTenantServicesAccessOk() (*MultiTenantServicesAccessModel, bool)`

GetMultiTenantServicesAccessOk returns a tuple with the MultiTenantServicesAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiTenantServicesAccess

`func (o *EditSiteSettingsRequestModel) SetMultiTenantServicesAccess(v MultiTenantServicesAccessModel)`

SetMultiTenantServicesAccess sets MultiTenantServicesAccess field to given value.

### HasMultiTenantServicesAccess

`func (o *EditSiteSettingsRequestModel) HasMultiTenantServicesAccess() bool`

HasMultiTenantServicesAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


