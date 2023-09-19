# SiteSettingsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseVerticalScalingForRdsLaunches** | Pointer to **NullableBool** |  | [optional] 
**DnsResolutionEnabled** | Pointer to **NullableBool** |  | [optional] 
**TrustRequestsSentToTheXmlServicePortEnabled** | Pointer to **NullableBool** |  | [optional] 
**WebUiPolicySetEnabled** | Pointer to **NullableBool** | The policy set in web UI is enabled or not | [optional] 
**SupportedAuthenticators** | Pointer to [**Authenticator**](Authenticator.md) |  | [optional] 

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

### SetUseVerticalScalingForRdsLaunchesNil

`func (o *SiteSettingsResponseModel) SetUseVerticalScalingForRdsLaunchesNil(b bool)`

 SetUseVerticalScalingForRdsLaunchesNil sets the value for UseVerticalScalingForRdsLaunches to be an explicit nil

### UnsetUseVerticalScalingForRdsLaunches
`func (o *SiteSettingsResponseModel) UnsetUseVerticalScalingForRdsLaunches()`

UnsetUseVerticalScalingForRdsLaunches ensures that no value is present for UseVerticalScalingForRdsLaunches, not even an explicit nil
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

### SetDnsResolutionEnabledNil

`func (o *SiteSettingsResponseModel) SetDnsResolutionEnabledNil(b bool)`

 SetDnsResolutionEnabledNil sets the value for DnsResolutionEnabled to be an explicit nil

### UnsetDnsResolutionEnabled
`func (o *SiteSettingsResponseModel) UnsetDnsResolutionEnabled()`

UnsetDnsResolutionEnabled ensures that no value is present for DnsResolutionEnabled, not even an explicit nil
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

### SetTrustRequestsSentToTheXmlServicePortEnabledNil

`func (o *SiteSettingsResponseModel) SetTrustRequestsSentToTheXmlServicePortEnabledNil(b bool)`

 SetTrustRequestsSentToTheXmlServicePortEnabledNil sets the value for TrustRequestsSentToTheXmlServicePortEnabled to be an explicit nil

### UnsetTrustRequestsSentToTheXmlServicePortEnabled
`func (o *SiteSettingsResponseModel) UnsetTrustRequestsSentToTheXmlServicePortEnabled()`

UnsetTrustRequestsSentToTheXmlServicePortEnabled ensures that no value is present for TrustRequestsSentToTheXmlServicePortEnabled, not even an explicit nil
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

### SetWebUiPolicySetEnabledNil

`func (o *SiteSettingsResponseModel) SetWebUiPolicySetEnabledNil(b bool)`

 SetWebUiPolicySetEnabledNil sets the value for WebUiPolicySetEnabled to be an explicit nil

### UnsetWebUiPolicySetEnabled
`func (o *SiteSettingsResponseModel) UnsetWebUiPolicySetEnabled()`

UnsetWebUiPolicySetEnabled ensures that no value is present for WebUiPolicySetEnabled, not even an explicit nil
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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


