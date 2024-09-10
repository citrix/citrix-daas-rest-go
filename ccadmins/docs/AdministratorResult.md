# AdministratorResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**UcOid** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**AdministratorType**](AdministratorType.md) |  | [optional] 
**AccessType** | Pointer to [**AdministratorAccessType**](AdministratorAccessType.md) |  | [optional] 
**ServiceProfile** | Pointer to **NullableString** |  | [optional] 
**FirstName** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**ProviderType** | Pointer to [**AdministratorProviderType**](AdministratorProviderType.md) |  | [optional] 
**ProviderId** | Pointer to **NullableString** |  | [optional] 
**ProviderProperties** | Pointer to [**NullableAdministratorProviderProperties**](AdministratorProviderProperties.md) |  | [optional] 
**ExternalOid** | Pointer to **NullableString** |  | [optional] 
**EmailPreferences** | Pointer to **NullableString** |  | [optional] 
**NotificationsEmailPreferences** | Pointer to [**AdministratorNotificationsEmailsPreferences**](AdministratorNotificationsEmailsPreferences.md) |  | [optional] 
**AuthDomain** | Pointer to **NullableString** |  | [optional] 
**Pending** | Pointer to **NullableBool** |  | [optional] 
**InvitationExpired** | Pointer to **NullableBool** |  | [optional] 
**LegacyProviders** | Pointer to **[]string** |  | [optional] 
**CreatedDate** | Pointer to **NullableTime** |  | [optional] 
**UpdatedDate** | Pointer to **NullableTime** |  | [optional] 
**ETag** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAdministratorResult

`func NewAdministratorResult() *AdministratorResult`

NewAdministratorResult instantiates a new AdministratorResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdministratorResultWithDefaults

`func NewAdministratorResultWithDefaults() *AdministratorResult`

NewAdministratorResultWithDefaults instantiates a new AdministratorResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *AdministratorResult) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AdministratorResult) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AdministratorResult) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *AdministratorResult) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *AdministratorResult) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *AdministratorResult) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetUcOid

`func (o *AdministratorResult) GetUcOid() string`

GetUcOid returns the UcOid field if non-nil, zero value otherwise.

### GetUcOidOk

`func (o *AdministratorResult) GetUcOidOk() (*string, bool)`

GetUcOidOk returns a tuple with the UcOid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcOid

`func (o *AdministratorResult) SetUcOid(v string)`

SetUcOid sets UcOid field to given value.

### HasUcOid

`func (o *AdministratorResult) HasUcOid() bool`

HasUcOid returns a boolean if a field has been set.

### SetUcOidNil

`func (o *AdministratorResult) SetUcOidNil(b bool)`

 SetUcOidNil sets the value for UcOid to be an explicit nil

### UnsetUcOid
`func (o *AdministratorResult) UnsetUcOid()`

UnsetUcOid ensures that no value is present for UcOid, not even an explicit nil
### GetUserId

`func (o *AdministratorResult) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AdministratorResult) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AdministratorResult) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AdministratorResult) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *AdministratorResult) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *AdministratorResult) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetType

`func (o *AdministratorResult) GetType() AdministratorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdministratorResult) GetTypeOk() (*AdministratorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdministratorResult) SetType(v AdministratorType)`

SetType sets Type field to given value.

### HasType

`func (o *AdministratorResult) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccessType

`func (o *AdministratorResult) GetAccessType() AdministratorAccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *AdministratorResult) GetAccessTypeOk() (*AdministratorAccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *AdministratorResult) SetAccessType(v AdministratorAccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *AdministratorResult) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetServiceProfile

`func (o *AdministratorResult) GetServiceProfile() string`

GetServiceProfile returns the ServiceProfile field if non-nil, zero value otherwise.

### GetServiceProfileOk

`func (o *AdministratorResult) GetServiceProfileOk() (*string, bool)`

GetServiceProfileOk returns a tuple with the ServiceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProfile

`func (o *AdministratorResult) SetServiceProfile(v string)`

SetServiceProfile sets ServiceProfile field to given value.

### HasServiceProfile

`func (o *AdministratorResult) HasServiceProfile() bool`

HasServiceProfile returns a boolean if a field has been set.

### SetServiceProfileNil

`func (o *AdministratorResult) SetServiceProfileNil(b bool)`

 SetServiceProfileNil sets the value for ServiceProfile to be an explicit nil

### UnsetServiceProfile
`func (o *AdministratorResult) UnsetServiceProfile()`

UnsetServiceProfile ensures that no value is present for ServiceProfile, not even an explicit nil
### GetFirstName

`func (o *AdministratorResult) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AdministratorResult) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AdministratorResult) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *AdministratorResult) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *AdministratorResult) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *AdministratorResult) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *AdministratorResult) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AdministratorResult) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AdministratorResult) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *AdministratorResult) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *AdministratorResult) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *AdministratorResult) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetDisplayName

`func (o *AdministratorResult) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AdministratorResult) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AdministratorResult) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AdministratorResult) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AdministratorResult) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AdministratorResult) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEmail

`func (o *AdministratorResult) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AdministratorResult) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AdministratorResult) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AdministratorResult) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *AdministratorResult) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *AdministratorResult) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetProviderType

`func (o *AdministratorResult) GetProviderType() AdministratorProviderType`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *AdministratorResult) GetProviderTypeOk() (*AdministratorProviderType, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *AdministratorResult) SetProviderType(v AdministratorProviderType)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *AdministratorResult) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetProviderId

`func (o *AdministratorResult) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *AdministratorResult) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *AdministratorResult) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *AdministratorResult) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *AdministratorResult) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *AdministratorResult) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetProviderProperties

`func (o *AdministratorResult) GetProviderProperties() AdministratorProviderProperties`

GetProviderProperties returns the ProviderProperties field if non-nil, zero value otherwise.

### GetProviderPropertiesOk

`func (o *AdministratorResult) GetProviderPropertiesOk() (*AdministratorProviderProperties, bool)`

GetProviderPropertiesOk returns a tuple with the ProviderProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderProperties

`func (o *AdministratorResult) SetProviderProperties(v AdministratorProviderProperties)`

SetProviderProperties sets ProviderProperties field to given value.

### HasProviderProperties

`func (o *AdministratorResult) HasProviderProperties() bool`

HasProviderProperties returns a boolean if a field has been set.

### SetProviderPropertiesNil

`func (o *AdministratorResult) SetProviderPropertiesNil(b bool)`

 SetProviderPropertiesNil sets the value for ProviderProperties to be an explicit nil

### UnsetProviderProperties
`func (o *AdministratorResult) UnsetProviderProperties()`

UnsetProviderProperties ensures that no value is present for ProviderProperties, not even an explicit nil
### GetExternalOid

`func (o *AdministratorResult) GetExternalOid() string`

GetExternalOid returns the ExternalOid field if non-nil, zero value otherwise.

### GetExternalOidOk

`func (o *AdministratorResult) GetExternalOidOk() (*string, bool)`

GetExternalOidOk returns a tuple with the ExternalOid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOid

`func (o *AdministratorResult) SetExternalOid(v string)`

SetExternalOid sets ExternalOid field to given value.

### HasExternalOid

`func (o *AdministratorResult) HasExternalOid() bool`

HasExternalOid returns a boolean if a field has been set.

### SetExternalOidNil

`func (o *AdministratorResult) SetExternalOidNil(b bool)`

 SetExternalOidNil sets the value for ExternalOid to be an explicit nil

### UnsetExternalOid
`func (o *AdministratorResult) UnsetExternalOid()`

UnsetExternalOid ensures that no value is present for ExternalOid, not even an explicit nil
### GetEmailPreferences

`func (o *AdministratorResult) GetEmailPreferences() string`

GetEmailPreferences returns the EmailPreferences field if non-nil, zero value otherwise.

### GetEmailPreferencesOk

`func (o *AdministratorResult) GetEmailPreferencesOk() (*string, bool)`

GetEmailPreferencesOk returns a tuple with the EmailPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPreferences

`func (o *AdministratorResult) SetEmailPreferences(v string)`

SetEmailPreferences sets EmailPreferences field to given value.

### HasEmailPreferences

`func (o *AdministratorResult) HasEmailPreferences() bool`

HasEmailPreferences returns a boolean if a field has been set.

### SetEmailPreferencesNil

`func (o *AdministratorResult) SetEmailPreferencesNil(b bool)`

 SetEmailPreferencesNil sets the value for EmailPreferences to be an explicit nil

### UnsetEmailPreferences
`func (o *AdministratorResult) UnsetEmailPreferences()`

UnsetEmailPreferences ensures that no value is present for EmailPreferences, not even an explicit nil
### GetNotificationsEmailPreferences

`func (o *AdministratorResult) GetNotificationsEmailPreferences() AdministratorNotificationsEmailsPreferences`

GetNotificationsEmailPreferences returns the NotificationsEmailPreferences field if non-nil, zero value otherwise.

### GetNotificationsEmailPreferencesOk

`func (o *AdministratorResult) GetNotificationsEmailPreferencesOk() (*AdministratorNotificationsEmailsPreferences, bool)`

GetNotificationsEmailPreferencesOk returns a tuple with the NotificationsEmailPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsEmailPreferences

`func (o *AdministratorResult) SetNotificationsEmailPreferences(v AdministratorNotificationsEmailsPreferences)`

SetNotificationsEmailPreferences sets NotificationsEmailPreferences field to given value.

### HasNotificationsEmailPreferences

`func (o *AdministratorResult) HasNotificationsEmailPreferences() bool`

HasNotificationsEmailPreferences returns a boolean if a field has been set.

### GetAuthDomain

`func (o *AdministratorResult) GetAuthDomain() string`

GetAuthDomain returns the AuthDomain field if non-nil, zero value otherwise.

### GetAuthDomainOk

`func (o *AdministratorResult) GetAuthDomainOk() (*string, bool)`

GetAuthDomainOk returns a tuple with the AuthDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDomain

`func (o *AdministratorResult) SetAuthDomain(v string)`

SetAuthDomain sets AuthDomain field to given value.

### HasAuthDomain

`func (o *AdministratorResult) HasAuthDomain() bool`

HasAuthDomain returns a boolean if a field has been set.

### SetAuthDomainNil

`func (o *AdministratorResult) SetAuthDomainNil(b bool)`

 SetAuthDomainNil sets the value for AuthDomain to be an explicit nil

### UnsetAuthDomain
`func (o *AdministratorResult) UnsetAuthDomain()`

UnsetAuthDomain ensures that no value is present for AuthDomain, not even an explicit nil
### GetPending

`func (o *AdministratorResult) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *AdministratorResult) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *AdministratorResult) SetPending(v bool)`

SetPending sets Pending field to given value.

### HasPending

`func (o *AdministratorResult) HasPending() bool`

HasPending returns a boolean if a field has been set.

### SetPendingNil

`func (o *AdministratorResult) SetPendingNil(b bool)`

 SetPendingNil sets the value for Pending to be an explicit nil

### UnsetPending
`func (o *AdministratorResult) UnsetPending()`

UnsetPending ensures that no value is present for Pending, not even an explicit nil
### GetInvitationExpired

`func (o *AdministratorResult) GetInvitationExpired() bool`

GetInvitationExpired returns the InvitationExpired field if non-nil, zero value otherwise.

### GetInvitationExpiredOk

`func (o *AdministratorResult) GetInvitationExpiredOk() (*bool, bool)`

GetInvitationExpiredOk returns a tuple with the InvitationExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationExpired

`func (o *AdministratorResult) SetInvitationExpired(v bool)`

SetInvitationExpired sets InvitationExpired field to given value.

### HasInvitationExpired

`func (o *AdministratorResult) HasInvitationExpired() bool`

HasInvitationExpired returns a boolean if a field has been set.

### SetInvitationExpiredNil

`func (o *AdministratorResult) SetInvitationExpiredNil(b bool)`

 SetInvitationExpiredNil sets the value for InvitationExpired to be an explicit nil

### UnsetInvitationExpired
`func (o *AdministratorResult) UnsetInvitationExpired()`

UnsetInvitationExpired ensures that no value is present for InvitationExpired, not even an explicit nil
### GetLegacyProviders

`func (o *AdministratorResult) GetLegacyProviders() []string`

GetLegacyProviders returns the LegacyProviders field if non-nil, zero value otherwise.

### GetLegacyProvidersOk

`func (o *AdministratorResult) GetLegacyProvidersOk() (*[]string, bool)`

GetLegacyProvidersOk returns a tuple with the LegacyProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyProviders

`func (o *AdministratorResult) SetLegacyProviders(v []string)`

SetLegacyProviders sets LegacyProviders field to given value.

### HasLegacyProviders

`func (o *AdministratorResult) HasLegacyProviders() bool`

HasLegacyProviders returns a boolean if a field has been set.

### SetLegacyProvidersNil

`func (o *AdministratorResult) SetLegacyProvidersNil(b bool)`

 SetLegacyProvidersNil sets the value for LegacyProviders to be an explicit nil

### UnsetLegacyProviders
`func (o *AdministratorResult) UnsetLegacyProviders()`

UnsetLegacyProviders ensures that no value is present for LegacyProviders, not even an explicit nil
### GetCreatedDate

`func (o *AdministratorResult) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *AdministratorResult) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *AdministratorResult) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *AdministratorResult) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *AdministratorResult) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *AdministratorResult) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetUpdatedDate

`func (o *AdministratorResult) GetUpdatedDate() time.Time`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *AdministratorResult) GetUpdatedDateOk() (*time.Time, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *AdministratorResult) SetUpdatedDate(v time.Time)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *AdministratorResult) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.

### SetUpdatedDateNil

`func (o *AdministratorResult) SetUpdatedDateNil(b bool)`

 SetUpdatedDateNil sets the value for UpdatedDate to be an explicit nil

### UnsetUpdatedDate
`func (o *AdministratorResult) UnsetUpdatedDate()`

UnsetUpdatedDate ensures that no value is present for UpdatedDate, not even an explicit nil
### GetETag

`func (o *AdministratorResult) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *AdministratorResult) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *AdministratorResult) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *AdministratorResult) HasETag() bool`

HasETag returns a boolean if a field has been set.

### SetETagNil

`func (o *AdministratorResult) SetETagNil(b bool)`

 SetETagNil sets the value for ETag to be an explicit nil

### UnsetETag
`func (o *AdministratorResult) UnsetETag()`

UnsetETag ensures that no value is present for ETag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


