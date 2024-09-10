# IdpStatusesModel

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
**Items** | Pointer to [**[]IdpStatusModel**](IdpStatusModel.md) |  | [optional] 
**ContinuationToken** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIdpStatusesModel

`func NewIdpStatusesModel() *IdpStatusesModel`

NewIdpStatusesModel instantiates a new IdpStatusesModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpStatusesModelWithDefaults

`func NewIdpStatusesModelWithDefaults() *IdpStatusesModel`

NewIdpStatusesModelWithDefaults instantiates a new IdpStatusesModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalStatusInfo

`func (o *IdpStatusesModel) GetAdditionalStatusInfo() map[string]string`

GetAdditionalStatusInfo returns the AdditionalStatusInfo field if non-nil, zero value otherwise.

### GetAdditionalStatusInfoOk

`func (o *IdpStatusesModel) GetAdditionalStatusInfoOk() (*map[string]string, bool)`

GetAdditionalStatusInfoOk returns a tuple with the AdditionalStatusInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalStatusInfo

`func (o *IdpStatusesModel) SetAdditionalStatusInfo(v map[string]string)`

SetAdditionalStatusInfo sets AdditionalStatusInfo field to given value.

### HasAdditionalStatusInfo

`func (o *IdpStatusesModel) HasAdditionalStatusInfo() bool`

HasAdditionalStatusInfo returns a boolean if a field has been set.

### SetAdditionalStatusInfoNil

`func (o *IdpStatusesModel) SetAdditionalStatusInfoNil(b bool)`

 SetAdditionalStatusInfoNil sets the value for AdditionalStatusInfo to be an explicit nil

### UnsetAdditionalStatusInfo
`func (o *IdpStatusesModel) UnsetAdditionalStatusInfo()`

UnsetAdditionalStatusInfo ensures that no value is present for AdditionalStatusInfo, not even an explicit nil
### GetAuthDomainName

`func (o *IdpStatusesModel) GetAuthDomainName() string`

GetAuthDomainName returns the AuthDomainName field if non-nil, zero value otherwise.

### GetAuthDomainNameOk

`func (o *IdpStatusesModel) GetAuthDomainNameOk() (*string, bool)`

GetAuthDomainNameOk returns a tuple with the AuthDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDomainName

`func (o *IdpStatusesModel) SetAuthDomainName(v string)`

SetAuthDomainName sets AuthDomainName field to given value.

### HasAuthDomainName

`func (o *IdpStatusesModel) HasAuthDomainName() bool`

HasAuthDomainName returns a boolean if a field has been set.

### SetAuthDomainNameNil

`func (o *IdpStatusesModel) SetAuthDomainNameNil(b bool)`

 SetAuthDomainNameNil sets the value for AuthDomainName to be an explicit nil

### UnsetAuthDomainName
`func (o *IdpStatusesModel) UnsetAuthDomainName()`

UnsetAuthDomainName ensures that no value is present for AuthDomainName, not even an explicit nil
### GetAvailableActions

`func (o *IdpStatusesModel) GetAvailableActions() []string`

GetAvailableActions returns the AvailableActions field if non-nil, zero value otherwise.

### GetAvailableActionsOk

`func (o *IdpStatusesModel) GetAvailableActionsOk() (*[]string, bool)`

GetAvailableActionsOk returns a tuple with the AvailableActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableActions

`func (o *IdpStatusesModel) SetAvailableActions(v []string)`

SetAvailableActions sets AvailableActions field to given value.

### HasAvailableActions

`func (o *IdpStatusesModel) HasAvailableActions() bool`

HasAvailableActions returns a boolean if a field has been set.

### SetAvailableActionsNil

`func (o *IdpStatusesModel) SetAvailableActionsNil(b bool)`

 SetAvailableActionsNil sets the value for AvailableActions to be an explicit nil

### UnsetAvailableActions
`func (o *IdpStatusesModel) UnsetAvailableActions()`

UnsetAvailableActions ensures that no value is present for AvailableActions, not even an explicit nil
### GetAzureAdConnection

`func (o *IdpStatusesModel) GetAzureAdConnection() AzureAdConnection`

GetAzureAdConnection returns the AzureAdConnection field if non-nil, zero value otherwise.

### GetAzureAdConnectionOk

`func (o *IdpStatusesModel) GetAzureAdConnectionOk() (*AzureAdConnection, bool)`

GetAzureAdConnectionOk returns a tuple with the AzureAdConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAdConnection

`func (o *IdpStatusesModel) SetAzureAdConnection(v AzureAdConnection)`

SetAzureAdConnection sets AzureAdConnection field to given value.

### HasAzureAdConnection

`func (o *IdpStatusesModel) HasAzureAdConnection() bool`

HasAzureAdConnection returns a boolean if a field has been set.

### GetClientId

`func (o *IdpStatusesModel) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IdpStatusesModel) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IdpStatusesModel) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IdpStatusesModel) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *IdpStatusesModel) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *IdpStatusesModel) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetConnectorsCount

`func (o *IdpStatusesModel) GetConnectorsCount() int32`

GetConnectorsCount returns the ConnectorsCount field if non-nil, zero value otherwise.

### GetConnectorsCountOk

`func (o *IdpStatusesModel) GetConnectorsCountOk() (*int32, bool)`

GetConnectorsCountOk returns a tuple with the ConnectorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorsCount

`func (o *IdpStatusesModel) SetConnectorsCount(v int32)`

SetConnectorsCount sets ConnectorsCount field to given value.

### HasConnectorsCount

`func (o *IdpStatusesModel) HasConnectorsCount() bool`

HasConnectorsCount returns a boolean if a field has been set.

### SetConnectorsCountNil

`func (o *IdpStatusesModel) SetConnectorsCountNil(b bool)`

 SetConnectorsCountNil sets the value for ConnectorsCount to be an explicit nil

### UnsetConnectorsCount
`func (o *IdpStatusesModel) UnsetConnectorsCount()`

UnsetConnectorsCount ensures that no value is present for ConnectorsCount, not even an explicit nil
### GetCustomerId

`func (o *IdpStatusesModel) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *IdpStatusesModel) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *IdpStatusesModel) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *IdpStatusesModel) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *IdpStatusesModel) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *IdpStatusesModel) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetDevicesLastResetTime

`func (o *IdpStatusesModel) GetDevicesLastResetTime() map[string]string`

GetDevicesLastResetTime returns the DevicesLastResetTime field if non-nil, zero value otherwise.

### GetDevicesLastResetTimeOk

`func (o *IdpStatusesModel) GetDevicesLastResetTimeOk() (*map[string]string, bool)`

GetDevicesLastResetTimeOk returns a tuple with the DevicesLastResetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesLastResetTime

`func (o *IdpStatusesModel) SetDevicesLastResetTime(v map[string]string)`

SetDevicesLastResetTime sets DevicesLastResetTime field to given value.

### HasDevicesLastResetTime

`func (o *IdpStatusesModel) HasDevicesLastResetTime() bool`

HasDevicesLastResetTime returns a boolean if a field has been set.

### SetDevicesLastResetTimeNil

`func (o *IdpStatusesModel) SetDevicesLastResetTimeNil(b bool)`

 SetDevicesLastResetTimeNil sets the value for DevicesLastResetTime to be an explicit nil

### UnsetDevicesLastResetTime
`func (o *IdpStatusesModel) UnsetDevicesLastResetTime()`

UnsetDevicesLastResetTime ensures that no value is present for DevicesLastResetTime, not even an explicit nil
### GetDomains

`func (o *IdpStatusesModel) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *IdpStatusesModel) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *IdpStatusesModel) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *IdpStatusesModel) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### SetDomainsNil

`func (o *IdpStatusesModel) SetDomainsNil(b bool)`

 SetDomainsNil sets the value for Domains to be an explicit nil

### UnsetDomains
`func (o *IdpStatusesModel) UnsetDomains()`

UnsetDomains ensures that no value is present for Domains, not even an explicit nil
### GetEnabled

`func (o *IdpStatusesModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdpStatusesModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdpStatusesModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IdpStatusesModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetForestPerDomain

`func (o *IdpStatusesModel) GetForestPerDomain() map[string]string`

GetForestPerDomain returns the ForestPerDomain field if non-nil, zero value otherwise.

### GetForestPerDomainOk

`func (o *IdpStatusesModel) GetForestPerDomainOk() (*map[string]string, bool)`

GetForestPerDomainOk returns a tuple with the ForestPerDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForestPerDomain

`func (o *IdpStatusesModel) SetForestPerDomain(v map[string]string)`

SetForestPerDomain sets ForestPerDomain field to given value.

### HasForestPerDomain

`func (o *IdpStatusesModel) HasForestPerDomain() bool`

HasForestPerDomain returns a boolean if a field has been set.

### SetForestPerDomainNil

`func (o *IdpStatusesModel) SetForestPerDomainNil(b bool)`

 SetForestPerDomainNil sets the value for ForestPerDomain to be an explicit nil

### UnsetForestPerDomain
`func (o *IdpStatusesModel) UnsetForestPerDomain()`

UnsetForestPerDomain ensures that no value is present for ForestPerDomain, not even an explicit nil
### GetIdentityProviderId

`func (o *IdpStatusesModel) GetIdentityProviderId() string`

GetIdentityProviderId returns the IdentityProviderId field if non-nil, zero value otherwise.

### GetIdentityProviderIdOk

`func (o *IdpStatusesModel) GetIdentityProviderIdOk() (*string, bool)`

GetIdentityProviderIdOk returns a tuple with the IdentityProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderId

`func (o *IdpStatusesModel) SetIdentityProviderId(v string)`

SetIdentityProviderId sets IdentityProviderId field to given value.

### HasIdentityProviderId

`func (o *IdpStatusesModel) HasIdentityProviderId() bool`

HasIdentityProviderId returns a boolean if a field has been set.

### SetIdentityProviderIdNil

`func (o *IdpStatusesModel) SetIdentityProviderIdNil(b bool)`

 SetIdentityProviderIdNil sets the value for IdentityProviderId to be an explicit nil

### UnsetIdentityProviderId
`func (o *IdpStatusesModel) UnsetIdentityProviderId()`

UnsetIdentityProviderId ensures that no value is present for IdentityProviderId, not even an explicit nil
### GetIdpInstanceId

`func (o *IdpStatusesModel) GetIdpInstanceId() string`

GetIdpInstanceId returns the IdpInstanceId field if non-nil, zero value otherwise.

### GetIdpInstanceIdOk

`func (o *IdpStatusesModel) GetIdpInstanceIdOk() (*string, bool)`

GetIdpInstanceIdOk returns a tuple with the IdpInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpInstanceId

`func (o *IdpStatusesModel) SetIdpInstanceId(v string)`

SetIdpInstanceId sets IdpInstanceId field to given value.

### HasIdpInstanceId

`func (o *IdpStatusesModel) HasIdpInstanceId() bool`

HasIdpInstanceId returns a boolean if a field has been set.

### SetIdpInstanceIdNil

`func (o *IdpStatusesModel) SetIdpInstanceIdNil(b bool)`

 SetIdpInstanceIdNil sets the value for IdpInstanceId to be an explicit nil

### UnsetIdpInstanceId
`func (o *IdpStatusesModel) UnsetIdpInstanceId()`

UnsetIdpInstanceId ensures that no value is present for IdpInstanceId, not even an explicit nil
### GetIdpNickname

`func (o *IdpStatusesModel) GetIdpNickname() string`

GetIdpNickname returns the IdpNickname field if non-nil, zero value otherwise.

### GetIdpNicknameOk

`func (o *IdpStatusesModel) GetIdpNicknameOk() (*string, bool)`

GetIdpNicknameOk returns a tuple with the IdpNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpNickname

`func (o *IdpStatusesModel) SetIdpNickname(v string)`

SetIdpNickname sets IdpNickname field to given value.

### HasIdpNickname

`func (o *IdpStatusesModel) HasIdpNickname() bool`

HasIdpNickname returns a boolean if a field has been set.

### SetIdpNicknameNil

`func (o *IdpStatusesModel) SetIdpNicknameNil(b bool)`

 SetIdpNicknameNil sets the value for IdpNickname to be an explicit nil

### UnsetIdpNickname
`func (o *IdpStatusesModel) UnsetIdpNickname()`

UnsetIdpNickname ensures that no value is present for IdpNickname, not even an explicit nil
### GetIssuerFqdn

`func (o *IdpStatusesModel) GetIssuerFqdn() string`

GetIssuerFqdn returns the IssuerFqdn field if non-nil, zero value otherwise.

### GetIssuerFqdnOk

`func (o *IdpStatusesModel) GetIssuerFqdnOk() (*string, bool)`

GetIssuerFqdnOk returns a tuple with the IssuerFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerFqdn

`func (o *IdpStatusesModel) SetIssuerFqdn(v string)`

SetIssuerFqdn sets IssuerFqdn field to given value.

### HasIssuerFqdn

`func (o *IdpStatusesModel) HasIssuerFqdn() bool`

HasIssuerFqdn returns a boolean if a field has been set.

### SetIssuerFqdnNil

`func (o *IdpStatusesModel) SetIssuerFqdnNil(b bool)`

 SetIssuerFqdnNil sets the value for IssuerFqdn to be an explicit nil

### UnsetIssuerFqdn
`func (o *IdpStatusesModel) UnsetIssuerFqdn()`

UnsetIssuerFqdn ensures that no value is present for IssuerFqdn, not even an explicit nil
### GetMultipleDevices

`func (o *IdpStatusesModel) GetMultipleDevices() bool`

GetMultipleDevices returns the MultipleDevices field if non-nil, zero value otherwise.

### GetMultipleDevicesOk

`func (o *IdpStatusesModel) GetMultipleDevicesOk() (*bool, bool)`

GetMultipleDevicesOk returns a tuple with the MultipleDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleDevices

`func (o *IdpStatusesModel) SetMultipleDevices(v bool)`

SetMultipleDevices sets MultipleDevices field to given value.

### HasMultipleDevices

`func (o *IdpStatusesModel) HasMultipleDevices() bool`

HasMultipleDevices returns a boolean if a field has been set.

### GetNotificationsEnabled

`func (o *IdpStatusesModel) GetNotificationsEnabled() bool`

GetNotificationsEnabled returns the NotificationsEnabled field if non-nil, zero value otherwise.

### GetNotificationsEnabledOk

`func (o *IdpStatusesModel) GetNotificationsEnabledOk() (*bool, bool)`

GetNotificationsEnabledOk returns a tuple with the NotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsEnabled

`func (o *IdpStatusesModel) SetNotificationsEnabled(v bool)`

SetNotificationsEnabled sets NotificationsEnabled field to given value.

### HasNotificationsEnabled

`func (o *IdpStatusesModel) HasNotificationsEnabled() bool`

HasNotificationsEnabled returns a boolean if a field has been set.

### GetNotificationsTime

`func (o *IdpStatusesModel) GetNotificationsTime() int32`

GetNotificationsTime returns the NotificationsTime field if non-nil, zero value otherwise.

### GetNotificationsTimeOk

`func (o *IdpStatusesModel) GetNotificationsTimeOk() (*int32, bool)`

GetNotificationsTimeOk returns a tuple with the NotificationsTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsTime

`func (o *IdpStatusesModel) SetNotificationsTime(v int32)`

SetNotificationsTime sets NotificationsTime field to given value.

### HasNotificationsTime

`func (o *IdpStatusesModel) HasNotificationsTime() bool`

HasNotificationsTime returns a boolean if a field has been set.

### GetUrl

`func (o *IdpStatusesModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IdpStatusesModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IdpStatusesModel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IdpStatusesModel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *IdpStatusesModel) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *IdpStatusesModel) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetItems

`func (o *IdpStatusesModel) GetItems() []IdpStatusModel`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IdpStatusesModel) GetItemsOk() (*[]IdpStatusModel, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IdpStatusesModel) SetItems(v []IdpStatusModel)`

SetItems sets Items field to given value.

### HasItems

`func (o *IdpStatusesModel) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *IdpStatusesModel) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *IdpStatusesModel) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetContinuationToken

`func (o *IdpStatusesModel) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *IdpStatusesModel) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *IdpStatusesModel) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *IdpStatusesModel) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### SetContinuationTokenNil

`func (o *IdpStatusesModel) SetContinuationTokenNil(b bool)`

 SetContinuationTokenNil sets the value for ContinuationToken to be an explicit nil

### UnsetContinuationToken
`func (o *IdpStatusesModel) UnsetContinuationToken()`

UnsetContinuationToken ensures that no value is present for ContinuationToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


