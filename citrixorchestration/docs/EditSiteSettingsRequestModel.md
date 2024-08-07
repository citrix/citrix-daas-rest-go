# EditSiteSettingsRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseVerticalScalingForRdsLaunches** | Pointer to **NullableBool** | Determines whether to use vertical scaling when considering RDS machines for launches. Vertical scaling would saturate machines in the current pool rather than send sessions to the least loaded machines. This would be a trade in performance vs. cost, where vertical scaling would be less costly. | [optional] 
**DnsResolutionEnabled** | Pointer to **NullableBool** | Changes whether ICA files returned by a broker service to a user device contain the numeric IP address or the DNS name of the desktop machine to which a session should be established. | [optional] 
**TrustRequestsSentToTheXmlServicePortEnabled** | Pointer to **NullableBool** | Changes whether the XML Service (as used by Storefront) implicitly trusts the originator of requests it receives, or whether it fully authenticates them. | [optional] 
**WebUiPolicySetEnabled** | Pointer to **NullableBool** | Determine whether the policy set in web UI is enabled or not | [optional] 
**ConsoleInactivityTimeoutMinutes** | Pointer to **NullableInt32** | The max number of minutes that console can be inactive. | [optional] 
**SupportedAuthenticators** | Pointer to [**Authenticator**](Authenticator.md) |  | [optional] 
**AllowedCorsOriginsForIwa** | Pointer to **[]string** | Applicable only for On-Premise. List of origins allowed to make cross-origin requests for Integrated Windows Authentication. | [optional] 
**MultiSites** | Pointer to [**[]MultiSiteModel**](MultiSiteModel.md) | Applicable only for On-Premise. Multiple sites configuration. | [optional] 
**DefaultDomain** | Pointer to **NullableString** | The default domain to be used in the login page. | [optional] 

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

### SetUseVerticalScalingForRdsLaunchesNil

`func (o *EditSiteSettingsRequestModel) SetUseVerticalScalingForRdsLaunchesNil(b bool)`

 SetUseVerticalScalingForRdsLaunchesNil sets the value for UseVerticalScalingForRdsLaunches to be an explicit nil

### UnsetUseVerticalScalingForRdsLaunches
`func (o *EditSiteSettingsRequestModel) UnsetUseVerticalScalingForRdsLaunches()`

UnsetUseVerticalScalingForRdsLaunches ensures that no value is present for UseVerticalScalingForRdsLaunches, not even an explicit nil
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

### SetDnsResolutionEnabledNil

`func (o *EditSiteSettingsRequestModel) SetDnsResolutionEnabledNil(b bool)`

 SetDnsResolutionEnabledNil sets the value for DnsResolutionEnabled to be an explicit nil

### UnsetDnsResolutionEnabled
`func (o *EditSiteSettingsRequestModel) UnsetDnsResolutionEnabled()`

UnsetDnsResolutionEnabled ensures that no value is present for DnsResolutionEnabled, not even an explicit nil
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

### SetTrustRequestsSentToTheXmlServicePortEnabledNil

`func (o *EditSiteSettingsRequestModel) SetTrustRequestsSentToTheXmlServicePortEnabledNil(b bool)`

 SetTrustRequestsSentToTheXmlServicePortEnabledNil sets the value for TrustRequestsSentToTheXmlServicePortEnabled to be an explicit nil

### UnsetTrustRequestsSentToTheXmlServicePortEnabled
`func (o *EditSiteSettingsRequestModel) UnsetTrustRequestsSentToTheXmlServicePortEnabled()`

UnsetTrustRequestsSentToTheXmlServicePortEnabled ensures that no value is present for TrustRequestsSentToTheXmlServicePortEnabled, not even an explicit nil
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

### SetWebUiPolicySetEnabledNil

`func (o *EditSiteSettingsRequestModel) SetWebUiPolicySetEnabledNil(b bool)`

 SetWebUiPolicySetEnabledNil sets the value for WebUiPolicySetEnabled to be an explicit nil

### UnsetWebUiPolicySetEnabled
`func (o *EditSiteSettingsRequestModel) UnsetWebUiPolicySetEnabled()`

UnsetWebUiPolicySetEnabled ensures that no value is present for WebUiPolicySetEnabled, not even an explicit nil
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

### SetConsoleInactivityTimeoutMinutesNil

`func (o *EditSiteSettingsRequestModel) SetConsoleInactivityTimeoutMinutesNil(b bool)`

 SetConsoleInactivityTimeoutMinutesNil sets the value for ConsoleInactivityTimeoutMinutes to be an explicit nil

### UnsetConsoleInactivityTimeoutMinutes
`func (o *EditSiteSettingsRequestModel) UnsetConsoleInactivityTimeoutMinutes()`

UnsetConsoleInactivityTimeoutMinutes ensures that no value is present for ConsoleInactivityTimeoutMinutes, not even an explicit nil
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

### SetAllowedCorsOriginsForIwaNil

`func (o *EditSiteSettingsRequestModel) SetAllowedCorsOriginsForIwaNil(b bool)`

 SetAllowedCorsOriginsForIwaNil sets the value for AllowedCorsOriginsForIwa to be an explicit nil

### UnsetAllowedCorsOriginsForIwa
`func (o *EditSiteSettingsRequestModel) UnsetAllowedCorsOriginsForIwa()`

UnsetAllowedCorsOriginsForIwa ensures that no value is present for AllowedCorsOriginsForIwa, not even an explicit nil
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

### SetMultiSitesNil

`func (o *EditSiteSettingsRequestModel) SetMultiSitesNil(b bool)`

 SetMultiSitesNil sets the value for MultiSites to be an explicit nil

### UnsetMultiSites
`func (o *EditSiteSettingsRequestModel) UnsetMultiSites()`

UnsetMultiSites ensures that no value is present for MultiSites, not even an explicit nil
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

### SetDefaultDomainNil

`func (o *EditSiteSettingsRequestModel) SetDefaultDomainNil(b bool)`

 SetDefaultDomainNil sets the value for DefaultDomain to be an explicit nil

### UnsetDefaultDomain
`func (o *EditSiteSettingsRequestModel) UnsetDefaultDomain()`

UnsetDefaultDomain ensures that no value is present for DefaultDomain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


