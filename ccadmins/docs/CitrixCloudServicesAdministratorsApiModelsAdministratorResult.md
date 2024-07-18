# CitrixCloudServicesAdministratorsApiModelsAdministratorResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**UcOid** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**CitrixCloudServicesAdministratorsApiModelsAdministratorType**](CitrixCloudServicesAdministratorsApiModelsAdministratorType.md) |  | [optional] 
**AccessType** | Pointer to [**CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType**](CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType.md) |  | [optional] 
**ServiceProfile** | Pointer to **NullableString** |  | [optional] 
**FirstName** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**ProviderType** | Pointer to [**CitrixCloudServicesAdministratorsApiModelsAdministratorProviderType**](CitrixCloudServicesAdministratorsApiModelsAdministratorProviderType.md) |  | [optional] 
**ProviderId** | Pointer to **NullableString** |  | [optional] 
**ProviderProperties** | Pointer to **map[string]string** |  | [optional] 
**ExternalOid** | Pointer to **NullableString** |  | [optional] 
**EmailPreferences** | Pointer to **NullableString** |  | [optional] 
**NotificationsEmailPreferences** | Pointer to [**CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences**](CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences.md) |  | [optional] 
**AuthDomain** | Pointer to **NullableString** |  | [optional] 
**Pending** | Pointer to **bool** |  | [optional] 
**InvitationExpired** | Pointer to **NullableBool** |  | [optional] 
**LegacyProviders** | Pointer to **[]string** |  | [optional] 
**CreatedDate** | Pointer to **NullableTime** |  | [optional] 
**UpdatedDate** | Pointer to **NullableTime** |  | [optional] 
**ETag** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCitrixCloudServicesAdministratorsApiModelsAdministratorResult

`func NewCitrixCloudServicesAdministratorsApiModelsAdministratorResult() *CitrixCloudServicesAdministratorsApiModelsAdministratorResult`

NewCitrixCloudServicesAdministratorsApiModelsAdministratorResult instantiates a new CitrixCloudServicesAdministratorsApiModelsAdministratorResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCitrixCloudServicesAdministratorsApiModelsAdministratorResultWithDefaults

`func NewCitrixCloudServicesAdministratorsApiModelsAdministratorResultWithDefaults() *CitrixCloudServicesAdministratorsApiModelsAdministratorResult`

NewCitrixCloudServicesAdministratorsApiModelsAdministratorResultWithDefaults instantiates a new CitrixCloudServicesAdministratorsApiModelsAdministratorResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetUcOid

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetUcOid() string`

GetUcOid returns the UcOid field if non-nil, zero value otherwise.

### GetUcOidOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetUcOidOk() (*string, bool)`

GetUcOidOk returns a tuple with the UcOid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcOid

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetUcOid(v string)`

SetUcOid sets UcOid field to given value.

### HasUcOid

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) HasUcOid() bool`

HasUcOid returns a boolean if a field has been set.

### SetUcOidNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetUcOidNil(b bool)`

 SetUcOidNil sets the value for UcOid to be an explicit nil

### UnsetUcOid
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) UnsetUcOid()`

UnsetUcOid ensures that no value is present for UcOid, not even an explicit nil
### GetUserId

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetType

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetType() CitrixCloudServicesAdministratorsApiModelsAdministratorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetTypeOk() (*CitrixCloudServicesAdministratorsApiModelsAdministratorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetType(v CitrixCloudServicesAdministratorsApiModelsAdministratorType)`

SetType sets Type field to given value.

### HasType

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccessType

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetAccessType() CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetAccessTypeOk() (*CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetAccessType(v CitrixCloudServicesAdministratorsApiModelsAdministratorAccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetServiceProfile

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetServiceProfile() string`

GetServiceProfile returns the ServiceProfile field if non-nil, zero value otherwise.

### GetServiceProfileOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetServiceProfileOk() (*string, bool)`

GetServiceProfileOk returns a tuple with the ServiceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProfile

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetServiceProfile(v string)`

SetServiceProfile sets ServiceProfile field to given value.

### HasServiceProfile

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) HasServiceProfile() bool`

HasServiceProfile returns a boolean if a field has been set.

### SetServiceProfileNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetServiceProfileNil(b bool)`

 SetServiceProfileNil sets the value for ServiceProfile to be an explicit nil

### UnsetServiceProfile
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) UnsetServiceProfile()`

UnsetServiceProfile ensures that no value is present for ServiceProfile, not even an explicit nil
### GetFirstName

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetDisplayName

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEmail

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetProviderType

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetProviderType() CitrixCloudServicesAdministratorsApiModelsAdministratorProviderType`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetProviderTypeOk() (*CitrixCloudServicesAdministratorsApiModelsAdministratorProviderType, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetProviderType(v CitrixCloudServicesAdministratorsApiModelsAdministratorProviderType)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetProviderId

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetProviderProperties

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetProviderProperties() map[string]string`

GetProviderProperties returns the ProviderProperties field if non-nil, zero value otherwise.

### GetProviderPropertiesOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetProviderPropertiesOk() (*map[string]string, bool)`

GetProviderPropertiesOk returns a tuple with the ProviderProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderProperties

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetProviderProperties(v map[string]string)`

SetProviderProperties sets ProviderProperties field to given value.

### HasProviderProperties

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) HasProviderProperties() bool`

HasProviderProperties returns a boolean if a field has been set.

### SetProviderPropertiesNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetProviderPropertiesNil(b bool)`

 SetProviderPropertiesNil sets the value for ProviderProperties to be an explicit nil

### UnsetProviderProperties
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) UnsetProviderProperties()`

UnsetProviderProperties ensures that no value is present for ProviderProperties, not even an explicit nil
### GetExternalOid

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetExternalOid() string`

GetExternalOid returns the ExternalOid field if non-nil, zero value otherwise.

### GetExternalOidOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetExternalOidOk() (*string, bool)`

GetExternalOidOk returns a tuple with the ExternalOid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOid

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetExternalOid(v string)`

SetExternalOid sets ExternalOid field to given value.

### HasExternalOid

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) HasExternalOid() bool`

HasExternalOid returns a boolean if a field has been set.

### SetExternalOidNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetExternalOidNil(b bool)`

 SetExternalOidNil sets the value for ExternalOid to be an explicit nil

### UnsetExternalOid
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) UnsetExternalOid()`

UnsetExternalOid ensures that no value is present for ExternalOid, not even an explicit nil
### GetEmailPreferences

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetEmailPreferences() string`

GetEmailPreferences returns the EmailPreferences field if non-nil, zero value otherwise.

### GetEmailPreferencesOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetEmailPreferencesOk() (*string, bool)`

GetEmailPreferencesOk returns a tuple with the EmailPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPreferences

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetEmailPreferences(v string)`

SetEmailPreferences sets EmailPreferences field to given value.

### HasEmailPreferences

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) HasEmailPreferences() bool`

HasEmailPreferences returns a boolean if a field has been set.

### SetEmailPreferencesNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetEmailPreferencesNil(b bool)`

 SetEmailPreferencesNil sets the value for EmailPreferences to be an explicit nil

### UnsetEmailPreferences
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) UnsetEmailPreferences()`

UnsetEmailPreferences ensures that no value is present for EmailPreferences, not even an explicit nil
### GetNotificationsEmailPreferences

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetNotificationsEmailPreferences() CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences`

GetNotificationsEmailPreferences returns the NotificationsEmailPreferences field if non-nil, zero value otherwise.

### GetNotificationsEmailPreferencesOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetNotificationsEmailPreferencesOk() (*CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences, bool)`

GetNotificationsEmailPreferencesOk returns a tuple with the NotificationsEmailPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsEmailPreferences

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetNotificationsEmailPreferences(v CitrixCloudServicesAdministratorsApiModelsAdministratorNotificationsEmailsPreferences)`

SetNotificationsEmailPreferences sets NotificationsEmailPreferences field to given value.

### HasNotificationsEmailPreferences

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) HasNotificationsEmailPreferences() bool`

HasNotificationsEmailPreferences returns a boolean if a field has been set.

### GetAuthDomain

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetAuthDomain() string`

GetAuthDomain returns the AuthDomain field if non-nil, zero value otherwise.

### GetAuthDomainOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetAuthDomainOk() (*string, bool)`

GetAuthDomainOk returns a tuple with the AuthDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDomain

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetAuthDomain(v string)`

SetAuthDomain sets AuthDomain field to given value.

### HasAuthDomain

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) HasAuthDomain() bool`

HasAuthDomain returns a boolean if a field has been set.

### SetAuthDomainNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetAuthDomainNil(b bool)`

 SetAuthDomainNil sets the value for AuthDomain to be an explicit nil

### UnsetAuthDomain
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) UnsetAuthDomain()`

UnsetAuthDomain ensures that no value is present for AuthDomain, not even an explicit nil
### GetPending

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetPending(v bool)`

SetPending sets Pending field to given value.

### HasPending

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetInvitationExpired

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetInvitationExpired() bool`

GetInvitationExpired returns the InvitationExpired field if non-nil, zero value otherwise.

### GetInvitationExpiredOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetInvitationExpiredOk() (*bool, bool)`

GetInvitationExpiredOk returns a tuple with the InvitationExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationExpired

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetInvitationExpired(v bool)`

SetInvitationExpired sets InvitationExpired field to given value.

### HasInvitationExpired

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) HasInvitationExpired() bool`

HasInvitationExpired returns a boolean if a field has been set.

### SetInvitationExpiredNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetInvitationExpiredNil(b bool)`

 SetInvitationExpiredNil sets the value for InvitationExpired to be an explicit nil

### UnsetInvitationExpired
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) UnsetInvitationExpired()`

UnsetInvitationExpired ensures that no value is present for InvitationExpired, not even an explicit nil
### GetLegacyProviders

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetLegacyProviders() []string`

GetLegacyProviders returns the LegacyProviders field if non-nil, zero value otherwise.

### GetLegacyProvidersOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetLegacyProvidersOk() (*[]string, bool)`

GetLegacyProvidersOk returns a tuple with the LegacyProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyProviders

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetLegacyProviders(v []string)`

SetLegacyProviders sets LegacyProviders field to given value.

### HasLegacyProviders

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) HasLegacyProviders() bool`

HasLegacyProviders returns a boolean if a field has been set.

### SetLegacyProvidersNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetLegacyProvidersNil(b bool)`

 SetLegacyProvidersNil sets the value for LegacyProviders to be an explicit nil

### UnsetLegacyProviders
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) UnsetLegacyProviders()`

UnsetLegacyProviders ensures that no value is present for LegacyProviders, not even an explicit nil
### GetCreatedDate

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetUpdatedDate

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetUpdatedDate() time.Time`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetUpdatedDateOk() (*time.Time, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetUpdatedDate(v time.Time)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.

### SetUpdatedDateNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetUpdatedDateNil(b bool)`

 SetUpdatedDateNil sets the value for UpdatedDate to be an explicit nil

### UnsetUpdatedDate
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) UnsetUpdatedDate()`

UnsetUpdatedDate ensures that no value is present for UpdatedDate, not even an explicit nil
### GetETag

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) HasETag() bool`

HasETag returns a boolean if a field has been set.

### SetETagNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) SetETagNil(b bool)`

 SetETagNil sets the value for ETag to be an explicit nil

### UnsetETag
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorResult) UnsetETag()`

UnsetETag ensures that no value is present for ETag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


