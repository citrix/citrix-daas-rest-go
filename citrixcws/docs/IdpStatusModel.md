# IdpStatusModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalStatusInfo** | Pointer to **map[string]string** |  | [optional] 
**AuthDomainName** | Pointer to **NullableString** |  | [optional] 
**AvailableActions** | Pointer to **[]string** |  | [optional] 
**AzureAdConnection** | Pointer to [**AzureAdConnection**](AzureAdConnection.md) |  | [optional] 
**ClientId** | Pointer to **NullableString** |  | [optional] 
**ConnectorsCount** | Pointer to **NullableInt32** |  | [optional] 
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**DevicesLastResetTime** | Pointer to **map[string]string** |  | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] [readonly] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ForestPerDomain** | Pointer to **map[string]string** |  | [optional] 
**IdentityProviderId** | Pointer to **NullableString** |  | [optional] 
**IdpInstanceId** | Pointer to **NullableString** |  | [optional] 
**IdpNickname** | Pointer to **NullableString** |  | [optional] 
**IssuerFqdn** | Pointer to **NullableString** |  | [optional] 
**MultipleDevices** | Pointer to **bool** |  | [optional] 
**NotificationsEnabled** | Pointer to **bool** |  | [optional] 
**NotificationsTime** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIdpStatusModel

`func NewIdpStatusModel() *IdpStatusModel`

NewIdpStatusModel instantiates a new IdpStatusModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpStatusModelWithDefaults

`func NewIdpStatusModelWithDefaults() *IdpStatusModel`

NewIdpStatusModelWithDefaults instantiates a new IdpStatusModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalStatusInfo

`func (o *IdpStatusModel) GetAdditionalStatusInfo() map[string]string`

GetAdditionalStatusInfo returns the AdditionalStatusInfo field if non-nil, zero value otherwise.

### GetAdditionalStatusInfoOk

`func (o *IdpStatusModel) GetAdditionalStatusInfoOk() (*map[string]string, bool)`

GetAdditionalStatusInfoOk returns a tuple with the AdditionalStatusInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalStatusInfo

`func (o *IdpStatusModel) SetAdditionalStatusInfo(v map[string]string)`

SetAdditionalStatusInfo sets AdditionalStatusInfo field to given value.

### HasAdditionalStatusInfo

`func (o *IdpStatusModel) HasAdditionalStatusInfo() bool`

HasAdditionalStatusInfo returns a boolean if a field has been set.

### SetAdditionalStatusInfoNil

`func (o *IdpStatusModel) SetAdditionalStatusInfoNil(b bool)`

 SetAdditionalStatusInfoNil sets the value for AdditionalStatusInfo to be an explicit nil

### UnsetAdditionalStatusInfo
`func (o *IdpStatusModel) UnsetAdditionalStatusInfo()`

UnsetAdditionalStatusInfo ensures that no value is present for AdditionalStatusInfo, not even an explicit nil
### GetAuthDomainName

`func (o *IdpStatusModel) GetAuthDomainName() string`

GetAuthDomainName returns the AuthDomainName field if non-nil, zero value otherwise.

### GetAuthDomainNameOk

`func (o *IdpStatusModel) GetAuthDomainNameOk() (*string, bool)`

GetAuthDomainNameOk returns a tuple with the AuthDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDomainName

`func (o *IdpStatusModel) SetAuthDomainName(v string)`

SetAuthDomainName sets AuthDomainName field to given value.

### HasAuthDomainName

`func (o *IdpStatusModel) HasAuthDomainName() bool`

HasAuthDomainName returns a boolean if a field has been set.

### SetAuthDomainNameNil

`func (o *IdpStatusModel) SetAuthDomainNameNil(b bool)`

 SetAuthDomainNameNil sets the value for AuthDomainName to be an explicit nil

### UnsetAuthDomainName
`func (o *IdpStatusModel) UnsetAuthDomainName()`

UnsetAuthDomainName ensures that no value is present for AuthDomainName, not even an explicit nil
### GetAvailableActions

`func (o *IdpStatusModel) GetAvailableActions() []string`

GetAvailableActions returns the AvailableActions field if non-nil, zero value otherwise.

### GetAvailableActionsOk

`func (o *IdpStatusModel) GetAvailableActionsOk() (*[]string, bool)`

GetAvailableActionsOk returns a tuple with the AvailableActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableActions

`func (o *IdpStatusModel) SetAvailableActions(v []string)`

SetAvailableActions sets AvailableActions field to given value.

### HasAvailableActions

`func (o *IdpStatusModel) HasAvailableActions() bool`

HasAvailableActions returns a boolean if a field has been set.

### SetAvailableActionsNil

`func (o *IdpStatusModel) SetAvailableActionsNil(b bool)`

 SetAvailableActionsNil sets the value for AvailableActions to be an explicit nil

### UnsetAvailableActions
`func (o *IdpStatusModel) UnsetAvailableActions()`

UnsetAvailableActions ensures that no value is present for AvailableActions, not even an explicit nil
### GetAzureAdConnection

`func (o *IdpStatusModel) GetAzureAdConnection() AzureAdConnection`

GetAzureAdConnection returns the AzureAdConnection field if non-nil, zero value otherwise.

### GetAzureAdConnectionOk

`func (o *IdpStatusModel) GetAzureAdConnectionOk() (*AzureAdConnection, bool)`

GetAzureAdConnectionOk returns a tuple with the AzureAdConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAdConnection

`func (o *IdpStatusModel) SetAzureAdConnection(v AzureAdConnection)`

SetAzureAdConnection sets AzureAdConnection field to given value.

### HasAzureAdConnection

`func (o *IdpStatusModel) HasAzureAdConnection() bool`

HasAzureAdConnection returns a boolean if a field has been set.

### GetClientId

`func (o *IdpStatusModel) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IdpStatusModel) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IdpStatusModel) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IdpStatusModel) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *IdpStatusModel) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *IdpStatusModel) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetConnectorsCount

`func (o *IdpStatusModel) GetConnectorsCount() int32`

GetConnectorsCount returns the ConnectorsCount field if non-nil, zero value otherwise.

### GetConnectorsCountOk

`func (o *IdpStatusModel) GetConnectorsCountOk() (*int32, bool)`

GetConnectorsCountOk returns a tuple with the ConnectorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorsCount

`func (o *IdpStatusModel) SetConnectorsCount(v int32)`

SetConnectorsCount sets ConnectorsCount field to given value.

### HasConnectorsCount

`func (o *IdpStatusModel) HasConnectorsCount() bool`

HasConnectorsCount returns a boolean if a field has been set.

### SetConnectorsCountNil

`func (o *IdpStatusModel) SetConnectorsCountNil(b bool)`

 SetConnectorsCountNil sets the value for ConnectorsCount to be an explicit nil

### UnsetConnectorsCount
`func (o *IdpStatusModel) UnsetConnectorsCount()`

UnsetConnectorsCount ensures that no value is present for ConnectorsCount, not even an explicit nil
### GetCustomerId

`func (o *IdpStatusModel) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *IdpStatusModel) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *IdpStatusModel) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *IdpStatusModel) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *IdpStatusModel) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *IdpStatusModel) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetDevicesLastResetTime

`func (o *IdpStatusModel) GetDevicesLastResetTime() map[string]string`

GetDevicesLastResetTime returns the DevicesLastResetTime field if non-nil, zero value otherwise.

### GetDevicesLastResetTimeOk

`func (o *IdpStatusModel) GetDevicesLastResetTimeOk() (*map[string]string, bool)`

GetDevicesLastResetTimeOk returns a tuple with the DevicesLastResetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesLastResetTime

`func (o *IdpStatusModel) SetDevicesLastResetTime(v map[string]string)`

SetDevicesLastResetTime sets DevicesLastResetTime field to given value.

### HasDevicesLastResetTime

`func (o *IdpStatusModel) HasDevicesLastResetTime() bool`

HasDevicesLastResetTime returns a boolean if a field has been set.

### SetDevicesLastResetTimeNil

`func (o *IdpStatusModel) SetDevicesLastResetTimeNil(b bool)`

 SetDevicesLastResetTimeNil sets the value for DevicesLastResetTime to be an explicit nil

### UnsetDevicesLastResetTime
`func (o *IdpStatusModel) UnsetDevicesLastResetTime()`

UnsetDevicesLastResetTime ensures that no value is present for DevicesLastResetTime, not even an explicit nil
### GetDomains

`func (o *IdpStatusModel) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *IdpStatusModel) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *IdpStatusModel) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *IdpStatusModel) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### SetDomainsNil

`func (o *IdpStatusModel) SetDomainsNil(b bool)`

 SetDomainsNil sets the value for Domains to be an explicit nil

### UnsetDomains
`func (o *IdpStatusModel) UnsetDomains()`

UnsetDomains ensures that no value is present for Domains, not even an explicit nil
### GetEnabled

`func (o *IdpStatusModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdpStatusModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdpStatusModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IdpStatusModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetForestPerDomain

`func (o *IdpStatusModel) GetForestPerDomain() map[string]string`

GetForestPerDomain returns the ForestPerDomain field if non-nil, zero value otherwise.

### GetForestPerDomainOk

`func (o *IdpStatusModel) GetForestPerDomainOk() (*map[string]string, bool)`

GetForestPerDomainOk returns a tuple with the ForestPerDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForestPerDomain

`func (o *IdpStatusModel) SetForestPerDomain(v map[string]string)`

SetForestPerDomain sets ForestPerDomain field to given value.

### HasForestPerDomain

`func (o *IdpStatusModel) HasForestPerDomain() bool`

HasForestPerDomain returns a boolean if a field has been set.

### SetForestPerDomainNil

`func (o *IdpStatusModel) SetForestPerDomainNil(b bool)`

 SetForestPerDomainNil sets the value for ForestPerDomain to be an explicit nil

### UnsetForestPerDomain
`func (o *IdpStatusModel) UnsetForestPerDomain()`

UnsetForestPerDomain ensures that no value is present for ForestPerDomain, not even an explicit nil
### GetIdentityProviderId

`func (o *IdpStatusModel) GetIdentityProviderId() string`

GetIdentityProviderId returns the IdentityProviderId field if non-nil, zero value otherwise.

### GetIdentityProviderIdOk

`func (o *IdpStatusModel) GetIdentityProviderIdOk() (*string, bool)`

GetIdentityProviderIdOk returns a tuple with the IdentityProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderId

`func (o *IdpStatusModel) SetIdentityProviderId(v string)`

SetIdentityProviderId sets IdentityProviderId field to given value.

### HasIdentityProviderId

`func (o *IdpStatusModel) HasIdentityProviderId() bool`

HasIdentityProviderId returns a boolean if a field has been set.

### SetIdentityProviderIdNil

`func (o *IdpStatusModel) SetIdentityProviderIdNil(b bool)`

 SetIdentityProviderIdNil sets the value for IdentityProviderId to be an explicit nil

### UnsetIdentityProviderId
`func (o *IdpStatusModel) UnsetIdentityProviderId()`

UnsetIdentityProviderId ensures that no value is present for IdentityProviderId, not even an explicit nil
### GetIdpInstanceId

`func (o *IdpStatusModel) GetIdpInstanceId() string`

GetIdpInstanceId returns the IdpInstanceId field if non-nil, zero value otherwise.

### GetIdpInstanceIdOk

`func (o *IdpStatusModel) GetIdpInstanceIdOk() (*string, bool)`

GetIdpInstanceIdOk returns a tuple with the IdpInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpInstanceId

`func (o *IdpStatusModel) SetIdpInstanceId(v string)`

SetIdpInstanceId sets IdpInstanceId field to given value.

### HasIdpInstanceId

`func (o *IdpStatusModel) HasIdpInstanceId() bool`

HasIdpInstanceId returns a boolean if a field has been set.

### SetIdpInstanceIdNil

`func (o *IdpStatusModel) SetIdpInstanceIdNil(b bool)`

 SetIdpInstanceIdNil sets the value for IdpInstanceId to be an explicit nil

### UnsetIdpInstanceId
`func (o *IdpStatusModel) UnsetIdpInstanceId()`

UnsetIdpInstanceId ensures that no value is present for IdpInstanceId, not even an explicit nil
### GetIdpNickname

`func (o *IdpStatusModel) GetIdpNickname() string`

GetIdpNickname returns the IdpNickname field if non-nil, zero value otherwise.

### GetIdpNicknameOk

`func (o *IdpStatusModel) GetIdpNicknameOk() (*string, bool)`

GetIdpNicknameOk returns a tuple with the IdpNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpNickname

`func (o *IdpStatusModel) SetIdpNickname(v string)`

SetIdpNickname sets IdpNickname field to given value.

### HasIdpNickname

`func (o *IdpStatusModel) HasIdpNickname() bool`

HasIdpNickname returns a boolean if a field has been set.

### SetIdpNicknameNil

`func (o *IdpStatusModel) SetIdpNicknameNil(b bool)`

 SetIdpNicknameNil sets the value for IdpNickname to be an explicit nil

### UnsetIdpNickname
`func (o *IdpStatusModel) UnsetIdpNickname()`

UnsetIdpNickname ensures that no value is present for IdpNickname, not even an explicit nil
### GetIssuerFqdn

`func (o *IdpStatusModel) GetIssuerFqdn() string`

GetIssuerFqdn returns the IssuerFqdn field if non-nil, zero value otherwise.

### GetIssuerFqdnOk

`func (o *IdpStatusModel) GetIssuerFqdnOk() (*string, bool)`

GetIssuerFqdnOk returns a tuple with the IssuerFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerFqdn

`func (o *IdpStatusModel) SetIssuerFqdn(v string)`

SetIssuerFqdn sets IssuerFqdn field to given value.

### HasIssuerFqdn

`func (o *IdpStatusModel) HasIssuerFqdn() bool`

HasIssuerFqdn returns a boolean if a field has been set.

### SetIssuerFqdnNil

`func (o *IdpStatusModel) SetIssuerFqdnNil(b bool)`

 SetIssuerFqdnNil sets the value for IssuerFqdn to be an explicit nil

### UnsetIssuerFqdn
`func (o *IdpStatusModel) UnsetIssuerFqdn()`

UnsetIssuerFqdn ensures that no value is present for IssuerFqdn, not even an explicit nil
### GetMultipleDevices

`func (o *IdpStatusModel) GetMultipleDevices() bool`

GetMultipleDevices returns the MultipleDevices field if non-nil, zero value otherwise.

### GetMultipleDevicesOk

`func (o *IdpStatusModel) GetMultipleDevicesOk() (*bool, bool)`

GetMultipleDevicesOk returns a tuple with the MultipleDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleDevices

`func (o *IdpStatusModel) SetMultipleDevices(v bool)`

SetMultipleDevices sets MultipleDevices field to given value.

### HasMultipleDevices

`func (o *IdpStatusModel) HasMultipleDevices() bool`

HasMultipleDevices returns a boolean if a field has been set.

### GetNotificationsEnabled

`func (o *IdpStatusModel) GetNotificationsEnabled() bool`

GetNotificationsEnabled returns the NotificationsEnabled field if non-nil, zero value otherwise.

### GetNotificationsEnabledOk

`func (o *IdpStatusModel) GetNotificationsEnabledOk() (*bool, bool)`

GetNotificationsEnabledOk returns a tuple with the NotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsEnabled

`func (o *IdpStatusModel) SetNotificationsEnabled(v bool)`

SetNotificationsEnabled sets NotificationsEnabled field to given value.

### HasNotificationsEnabled

`func (o *IdpStatusModel) HasNotificationsEnabled() bool`

HasNotificationsEnabled returns a boolean if a field has been set.

### GetNotificationsTime

`func (o *IdpStatusModel) GetNotificationsTime() int32`

GetNotificationsTime returns the NotificationsTime field if non-nil, zero value otherwise.

### GetNotificationsTimeOk

`func (o *IdpStatusModel) GetNotificationsTimeOk() (*int32, bool)`

GetNotificationsTimeOk returns a tuple with the NotificationsTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsTime

`func (o *IdpStatusModel) SetNotificationsTime(v int32)`

SetNotificationsTime sets NotificationsTime field to given value.

### HasNotificationsTime

`func (o *IdpStatusModel) HasNotificationsTime() bool`

HasNotificationsTime returns a boolean if a field has been set.

### GetUrl

`func (o *IdpStatusModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IdpStatusModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IdpStatusModel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IdpStatusModel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *IdpStatusModel) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *IdpStatusModel) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


