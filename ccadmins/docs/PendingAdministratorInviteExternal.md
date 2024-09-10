# PendingAdministratorInviteExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**RequestorEmail** | Pointer to **NullableString** |  | [optional] 
**FirstName** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**Access** | Pointer to [**AdministratorAccessModel**](AdministratorAccessModel.md) |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**ExternalProviderType** | [**AdministratorExternalProviderType**](AdministratorExternalProviderType.md) |  | 
**ExternalProviderId** | **string** |  | 
**ExternalUserId** | **string** |  | 
**ExternalProviderProperties** | Pointer to **map[string]interface{}** |  | [optional] 
**ExternalProviderAuthDomain** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPendingAdministratorInviteExternal

`func NewPendingAdministratorInviteExternal(email string, externalProviderType AdministratorExternalProviderType, externalProviderId string, externalUserId string, ) *PendingAdministratorInviteExternal`

NewPendingAdministratorInviteExternal instantiates a new PendingAdministratorInviteExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPendingAdministratorInviteExternalWithDefaults

`func NewPendingAdministratorInviteExternalWithDefaults() *PendingAdministratorInviteExternal`

NewPendingAdministratorInviteExternalWithDefaults instantiates a new PendingAdministratorInviteExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *PendingAdministratorInviteExternal) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PendingAdministratorInviteExternal) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PendingAdministratorInviteExternal) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRequestorEmail

`func (o *PendingAdministratorInviteExternal) GetRequestorEmail() string`

GetRequestorEmail returns the RequestorEmail field if non-nil, zero value otherwise.

### GetRequestorEmailOk

`func (o *PendingAdministratorInviteExternal) GetRequestorEmailOk() (*string, bool)`

GetRequestorEmailOk returns a tuple with the RequestorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestorEmail

`func (o *PendingAdministratorInviteExternal) SetRequestorEmail(v string)`

SetRequestorEmail sets RequestorEmail field to given value.

### HasRequestorEmail

`func (o *PendingAdministratorInviteExternal) HasRequestorEmail() bool`

HasRequestorEmail returns a boolean if a field has been set.

### SetRequestorEmailNil

`func (o *PendingAdministratorInviteExternal) SetRequestorEmailNil(b bool)`

 SetRequestorEmailNil sets the value for RequestorEmail to be an explicit nil

### UnsetRequestorEmail
`func (o *PendingAdministratorInviteExternal) UnsetRequestorEmail()`

UnsetRequestorEmail ensures that no value is present for RequestorEmail, not even an explicit nil
### GetFirstName

`func (o *PendingAdministratorInviteExternal) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PendingAdministratorInviteExternal) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PendingAdministratorInviteExternal) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PendingAdministratorInviteExternal) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *PendingAdministratorInviteExternal) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *PendingAdministratorInviteExternal) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *PendingAdministratorInviteExternal) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PendingAdministratorInviteExternal) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PendingAdministratorInviteExternal) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PendingAdministratorInviteExternal) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *PendingAdministratorInviteExternal) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *PendingAdministratorInviteExternal) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetAccess

`func (o *PendingAdministratorInviteExternal) GetAccess() AdministratorAccessModel`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *PendingAdministratorInviteExternal) GetAccessOk() (*AdministratorAccessModel, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *PendingAdministratorInviteExternal) SetAccess(v AdministratorAccessModel)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *PendingAdministratorInviteExternal) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetDisplayName

`func (o *PendingAdministratorInviteExternal) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PendingAdministratorInviteExternal) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PendingAdministratorInviteExternal) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PendingAdministratorInviteExternal) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *PendingAdministratorInviteExternal) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *PendingAdministratorInviteExternal) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetExternalProviderType

`func (o *PendingAdministratorInviteExternal) GetExternalProviderType() AdministratorExternalProviderType`

GetExternalProviderType returns the ExternalProviderType field if non-nil, zero value otherwise.

### GetExternalProviderTypeOk

`func (o *PendingAdministratorInviteExternal) GetExternalProviderTypeOk() (*AdministratorExternalProviderType, bool)`

GetExternalProviderTypeOk returns a tuple with the ExternalProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProviderType

`func (o *PendingAdministratorInviteExternal) SetExternalProviderType(v AdministratorExternalProviderType)`

SetExternalProviderType sets ExternalProviderType field to given value.


### GetExternalProviderId

`func (o *PendingAdministratorInviteExternal) GetExternalProviderId() string`

GetExternalProviderId returns the ExternalProviderId field if non-nil, zero value otherwise.

### GetExternalProviderIdOk

`func (o *PendingAdministratorInviteExternal) GetExternalProviderIdOk() (*string, bool)`

GetExternalProviderIdOk returns a tuple with the ExternalProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProviderId

`func (o *PendingAdministratorInviteExternal) SetExternalProviderId(v string)`

SetExternalProviderId sets ExternalProviderId field to given value.


### GetExternalUserId

`func (o *PendingAdministratorInviteExternal) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *PendingAdministratorInviteExternal) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *PendingAdministratorInviteExternal) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.


### GetExternalProviderProperties

`func (o *PendingAdministratorInviteExternal) GetExternalProviderProperties() map[string]interface{}`

GetExternalProviderProperties returns the ExternalProviderProperties field if non-nil, zero value otherwise.

### GetExternalProviderPropertiesOk

`func (o *PendingAdministratorInviteExternal) GetExternalProviderPropertiesOk() (*map[string]interface{}, bool)`

GetExternalProviderPropertiesOk returns a tuple with the ExternalProviderProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProviderProperties

`func (o *PendingAdministratorInviteExternal) SetExternalProviderProperties(v map[string]interface{})`

SetExternalProviderProperties sets ExternalProviderProperties field to given value.

### HasExternalProviderProperties

`func (o *PendingAdministratorInviteExternal) HasExternalProviderProperties() bool`

HasExternalProviderProperties returns a boolean if a field has been set.

### SetExternalProviderPropertiesNil

`func (o *PendingAdministratorInviteExternal) SetExternalProviderPropertiesNil(b bool)`

 SetExternalProviderPropertiesNil sets the value for ExternalProviderProperties to be an explicit nil

### UnsetExternalProviderProperties
`func (o *PendingAdministratorInviteExternal) UnsetExternalProviderProperties()`

UnsetExternalProviderProperties ensures that no value is present for ExternalProviderProperties, not even an explicit nil
### GetExternalProviderAuthDomain

`func (o *PendingAdministratorInviteExternal) GetExternalProviderAuthDomain() string`

GetExternalProviderAuthDomain returns the ExternalProviderAuthDomain field if non-nil, zero value otherwise.

### GetExternalProviderAuthDomainOk

`func (o *PendingAdministratorInviteExternal) GetExternalProviderAuthDomainOk() (*string, bool)`

GetExternalProviderAuthDomainOk returns a tuple with the ExternalProviderAuthDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProviderAuthDomain

`func (o *PendingAdministratorInviteExternal) SetExternalProviderAuthDomain(v string)`

SetExternalProviderAuthDomain sets ExternalProviderAuthDomain field to given value.

### HasExternalProviderAuthDomain

`func (o *PendingAdministratorInviteExternal) HasExternalProviderAuthDomain() bool`

HasExternalProviderAuthDomain returns a boolean if a field has been set.

### SetExternalProviderAuthDomainNil

`func (o *PendingAdministratorInviteExternal) SetExternalProviderAuthDomainNil(b bool)`

 SetExternalProviderAuthDomainNil sets the value for ExternalProviderAuthDomain to be an explicit nil

### UnsetExternalProviderAuthDomain
`func (o *PendingAdministratorInviteExternal) UnsetExternalProviderAuthDomain()`

UnsetExternalProviderAuthDomain ensures that no value is present for ExternalProviderAuthDomain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


