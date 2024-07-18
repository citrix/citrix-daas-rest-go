# CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UcOid** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**EmailVerified** | Pointer to **NullableBool** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**GroupClaims** | Pointer to **[]string** |  | [optional] 
**AdminGroups** | Pointer to **[]string** |  | [optional] 
**UserClaims** | Pointer to [**CitrixCloudServicesAdministratorsApiModelsUserClaims**](CitrixCloudServicesAdministratorsApiModelsUserClaims.md) |  | [optional] 
**LoginCustomerId** | Pointer to **NullableString** |  | [optional] 
**Customers** | Pointer to [**[]CitrixCloudServicesAdministratorsApiModelsCustomerModel**](CitrixCloudServicesAdministratorsApiModelsCustomerModel.md) |  | [optional] 
**LinkedCustomers** | Pointer to **[]string** |  | [optional] 
**IsExternal** | Pointer to **bool** |  | [optional] 
**Principal** | Pointer to **NullableString** |  | [optional] 
**Locale** | Pointer to **NullableString** |  | [optional] 
**AccessTokenScope** | Pointer to **NullableString** |  | [optional] 
**RefreshToken** | Pointer to **NullableString** |  | [optional] 
**AccessToken** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**RefreshExpiration** | Pointer to **int64** |  | [optional] 
**AuthenticationAlias** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**AuthenticationMethods** | Pointer to **[]string** |  | [optional] 
**Aud** | Pointer to **NullableString** |  | [optional] 
**Idp** | Pointer to **NullableString** |  | [optional] 
**PreviousIdps** | Pointer to **[]string** |  | [optional] 
**Discovery** | Pointer to [**CitrixCloudServicesAdministratorsApiModelsDiscoveryModel**](CitrixCloudServicesAdministratorsApiModelsDiscoveryModel.md) |  | [optional] 
**DirectoryContextClaim** | Pointer to [**CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim**](CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim.md) |  | [optional] 

## Methods

### NewCitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel

`func NewCitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel() *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel`

NewCitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel instantiates a new CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModelWithDefaults

`func NewCitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModelWithDefaults() *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel`

NewCitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModelWithDefaults instantiates a new CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUcOid

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetUcOid() string`

GetUcOid returns the UcOid field if non-nil, zero value otherwise.

### GetUcOidOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetUcOidOk() (*string, bool)`

GetUcOidOk returns a tuple with the UcOid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcOid

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetUcOid(v string)`

SetUcOid sets UcOid field to given value.

### HasUcOid

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasUcOid() bool`

HasUcOid returns a boolean if a field has been set.

### SetUcOidNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetUcOidNil(b bool)`

 SetUcOidNil sets the value for UcOid to be an explicit nil

### UnsetUcOid
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) UnsetUcOid()`

UnsetUcOid ensures that no value is present for UcOid, not even an explicit nil
### GetUserId

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetEmailVerified

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### SetEmailVerifiedNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetEmailVerifiedNil(b bool)`

 SetEmailVerifiedNil sets the value for EmailVerified to be an explicit nil

### UnsetEmailVerified
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) UnsetEmailVerified()`

UnsetEmailVerified ensures that no value is present for EmailVerified, not even an explicit nil
### GetEmail

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetGroupClaims

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetGroupClaims() []string`

GetGroupClaims returns the GroupClaims field if non-nil, zero value otherwise.

### GetGroupClaimsOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetGroupClaimsOk() (*[]string, bool)`

GetGroupClaimsOk returns a tuple with the GroupClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupClaims

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetGroupClaims(v []string)`

SetGroupClaims sets GroupClaims field to given value.

### HasGroupClaims

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasGroupClaims() bool`

HasGroupClaims returns a boolean if a field has been set.

### SetGroupClaimsNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetGroupClaimsNil(b bool)`

 SetGroupClaimsNil sets the value for GroupClaims to be an explicit nil

### UnsetGroupClaims
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) UnsetGroupClaims()`

UnsetGroupClaims ensures that no value is present for GroupClaims, not even an explicit nil
### GetAdminGroups

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetAdminGroups() []string`

GetAdminGroups returns the AdminGroups field if non-nil, zero value otherwise.

### GetAdminGroupsOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetAdminGroupsOk() (*[]string, bool)`

GetAdminGroupsOk returns a tuple with the AdminGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminGroups

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetAdminGroups(v []string)`

SetAdminGroups sets AdminGroups field to given value.

### HasAdminGroups

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasAdminGroups() bool`

HasAdminGroups returns a boolean if a field has been set.

### SetAdminGroupsNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetAdminGroupsNil(b bool)`

 SetAdminGroupsNil sets the value for AdminGroups to be an explicit nil

### UnsetAdminGroups
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) UnsetAdminGroups()`

UnsetAdminGroups ensures that no value is present for AdminGroups, not even an explicit nil
### GetUserClaims

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetUserClaims() CitrixCloudServicesAdministratorsApiModelsUserClaims`

GetUserClaims returns the UserClaims field if non-nil, zero value otherwise.

### GetUserClaimsOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetUserClaimsOk() (*CitrixCloudServicesAdministratorsApiModelsUserClaims, bool)`

GetUserClaimsOk returns a tuple with the UserClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserClaims

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetUserClaims(v CitrixCloudServicesAdministratorsApiModelsUserClaims)`

SetUserClaims sets UserClaims field to given value.

### HasUserClaims

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasUserClaims() bool`

HasUserClaims returns a boolean if a field has been set.

### GetLoginCustomerId

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetLoginCustomerId() string`

GetLoginCustomerId returns the LoginCustomerId field if non-nil, zero value otherwise.

### GetLoginCustomerIdOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetLoginCustomerIdOk() (*string, bool)`

GetLoginCustomerIdOk returns a tuple with the LoginCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCustomerId

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetLoginCustomerId(v string)`

SetLoginCustomerId sets LoginCustomerId field to given value.

### HasLoginCustomerId

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasLoginCustomerId() bool`

HasLoginCustomerId returns a boolean if a field has been set.

### SetLoginCustomerIdNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetLoginCustomerIdNil(b bool)`

 SetLoginCustomerIdNil sets the value for LoginCustomerId to be an explicit nil

### UnsetLoginCustomerId
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) UnsetLoginCustomerId()`

UnsetLoginCustomerId ensures that no value is present for LoginCustomerId, not even an explicit nil
### GetCustomers

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetCustomers() []CitrixCloudServicesAdministratorsApiModelsCustomerModel`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetCustomersOk() (*[]CitrixCloudServicesAdministratorsApiModelsCustomerModel, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetCustomers(v []CitrixCloudServicesAdministratorsApiModelsCustomerModel)`

SetCustomers sets Customers field to given value.

### HasCustomers

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasCustomers() bool`

HasCustomers returns a boolean if a field has been set.

### SetCustomersNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetCustomersNil(b bool)`

 SetCustomersNil sets the value for Customers to be an explicit nil

### UnsetCustomers
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) UnsetCustomers()`

UnsetCustomers ensures that no value is present for Customers, not even an explicit nil
### GetLinkedCustomers

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetLinkedCustomers() []string`

GetLinkedCustomers returns the LinkedCustomers field if non-nil, zero value otherwise.

### GetLinkedCustomersOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetLinkedCustomersOk() (*[]string, bool)`

GetLinkedCustomersOk returns a tuple with the LinkedCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedCustomers

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetLinkedCustomers(v []string)`

SetLinkedCustomers sets LinkedCustomers field to given value.

### HasLinkedCustomers

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasLinkedCustomers() bool`

HasLinkedCustomers returns a boolean if a field has been set.

### SetLinkedCustomersNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetLinkedCustomersNil(b bool)`

 SetLinkedCustomersNil sets the value for LinkedCustomers to be an explicit nil

### UnsetLinkedCustomers
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) UnsetLinkedCustomers()`

UnsetLinkedCustomers ensures that no value is present for LinkedCustomers, not even an explicit nil
### GetIsExternal

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetPrincipal

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### SetPrincipalNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetPrincipalNil(b bool)`

 SetPrincipalNil sets the value for Principal to be an explicit nil

### UnsetPrincipal
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) UnsetPrincipal()`

UnsetPrincipal ensures that no value is present for Principal, not even an explicit nil
### GetLocale

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetAccessTokenScope

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetAccessTokenScope() string`

GetAccessTokenScope returns the AccessTokenScope field if non-nil, zero value otherwise.

### GetAccessTokenScopeOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetAccessTokenScopeOk() (*string, bool)`

GetAccessTokenScopeOk returns a tuple with the AccessTokenScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenScope

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetAccessTokenScope(v string)`

SetAccessTokenScope sets AccessTokenScope field to given value.

### HasAccessTokenScope

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasAccessTokenScope() bool`

HasAccessTokenScope returns a boolean if a field has been set.

### SetAccessTokenScopeNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetAccessTokenScopeNil(b bool)`

 SetAccessTokenScopeNil sets the value for AccessTokenScope to be an explicit nil

### UnsetAccessTokenScope
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) UnsetAccessTokenScope()`

UnsetAccessTokenScope ensures that no value is present for AccessTokenScope, not even an explicit nil
### GetRefreshToken

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### SetRefreshTokenNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetRefreshTokenNil(b bool)`

 SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil

### UnsetRefreshToken
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) UnsetRefreshToken()`

UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
### GetAccessToken

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetDisplayName

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetRefreshExpiration

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetRefreshExpiration() int64`

GetRefreshExpiration returns the RefreshExpiration field if non-nil, zero value otherwise.

### GetRefreshExpirationOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetRefreshExpirationOk() (*int64, bool)`

GetRefreshExpirationOk returns a tuple with the RefreshExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshExpiration

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetRefreshExpiration(v int64)`

SetRefreshExpiration sets RefreshExpiration field to given value.

### HasRefreshExpiration

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasRefreshExpiration() bool`

HasRefreshExpiration returns a boolean if a field has been set.

### GetAuthenticationAlias

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetAuthenticationAlias() string`

GetAuthenticationAlias returns the AuthenticationAlias field if non-nil, zero value otherwise.

### GetAuthenticationAliasOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetAuthenticationAliasOk() (*string, bool)`

GetAuthenticationAliasOk returns a tuple with the AuthenticationAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationAlias

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetAuthenticationAlias(v string)`

SetAuthenticationAlias sets AuthenticationAlias field to given value.

### HasAuthenticationAlias

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasAuthenticationAlias() bool`

HasAuthenticationAlias returns a boolean if a field has been set.

### SetAuthenticationAliasNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetAuthenticationAliasNil(b bool)`

 SetAuthenticationAliasNil sets the value for AuthenticationAlias to be an explicit nil

### UnsetAuthenticationAlias
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) UnsetAuthenticationAlias()`

UnsetAuthenticationAlias ensures that no value is present for AuthenticationAlias, not even an explicit nil
### GetName

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSubject

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetAuthenticationMethods

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetAuthenticationMethods() []string`

GetAuthenticationMethods returns the AuthenticationMethods field if non-nil, zero value otherwise.

### GetAuthenticationMethodsOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetAuthenticationMethodsOk() (*[]string, bool)`

GetAuthenticationMethodsOk returns a tuple with the AuthenticationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethods

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetAuthenticationMethods(v []string)`

SetAuthenticationMethods sets AuthenticationMethods field to given value.

### HasAuthenticationMethods

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasAuthenticationMethods() bool`

HasAuthenticationMethods returns a boolean if a field has been set.

### SetAuthenticationMethodsNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetAuthenticationMethodsNil(b bool)`

 SetAuthenticationMethodsNil sets the value for AuthenticationMethods to be an explicit nil

### UnsetAuthenticationMethods
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) UnsetAuthenticationMethods()`

UnsetAuthenticationMethods ensures that no value is present for AuthenticationMethods, not even an explicit nil
### GetAud

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetAud() string`

GetAud returns the Aud field if non-nil, zero value otherwise.

### GetAudOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetAudOk() (*string, bool)`

GetAudOk returns a tuple with the Aud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAud

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetAud(v string)`

SetAud sets Aud field to given value.

### HasAud

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasAud() bool`

HasAud returns a boolean if a field has been set.

### SetAudNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetAudNil(b bool)`

 SetAudNil sets the value for Aud to be an explicit nil

### UnsetAud
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) UnsetAud()`

UnsetAud ensures that no value is present for Aud, not even an explicit nil
### GetIdp

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetIdp() string`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetIdpOk() (*string, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetIdp(v string)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### SetIdpNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetIdpNil(b bool)`

 SetIdpNil sets the value for Idp to be an explicit nil

### UnsetIdp
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) UnsetIdp()`

UnsetIdp ensures that no value is present for Idp, not even an explicit nil
### GetPreviousIdps

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetPreviousIdps() []string`

GetPreviousIdps returns the PreviousIdps field if non-nil, zero value otherwise.

### GetPreviousIdpsOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetPreviousIdpsOk() (*[]string, bool)`

GetPreviousIdpsOk returns a tuple with the PreviousIdps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousIdps

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetPreviousIdps(v []string)`

SetPreviousIdps sets PreviousIdps field to given value.

### HasPreviousIdps

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasPreviousIdps() bool`

HasPreviousIdps returns a boolean if a field has been set.

### SetPreviousIdpsNil

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetPreviousIdpsNil(b bool)`

 SetPreviousIdpsNil sets the value for PreviousIdps to be an explicit nil

### UnsetPreviousIdps
`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) UnsetPreviousIdps()`

UnsetPreviousIdps ensures that no value is present for PreviousIdps, not even an explicit nil
### GetDiscovery

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetDiscovery() CitrixCloudServicesAdministratorsApiModelsDiscoveryModel`

GetDiscovery returns the Discovery field if non-nil, zero value otherwise.

### GetDiscoveryOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetDiscoveryOk() (*CitrixCloudServicesAdministratorsApiModelsDiscoveryModel, bool)`

GetDiscoveryOk returns a tuple with the Discovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovery

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetDiscovery(v CitrixCloudServicesAdministratorsApiModelsDiscoveryModel)`

SetDiscovery sets Discovery field to given value.

### HasDiscovery

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasDiscovery() bool`

HasDiscovery returns a boolean if a field has been set.

### GetDirectoryContextClaim

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetDirectoryContextClaim() CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim`

GetDirectoryContextClaim returns the DirectoryContextClaim field if non-nil, zero value otherwise.

### GetDirectoryContextClaimOk

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) GetDirectoryContextClaimOk() (*CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim, bool)`

GetDirectoryContextClaimOk returns a tuple with the DirectoryContextClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryContextClaim

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) SetDirectoryContextClaim(v CitrixCloudServicesAdministratorsApiModelsDirectoryContextClaim)`

SetDirectoryContextClaim sets DirectoryContextClaim field to given value.

### HasDirectoryContextClaim

`func (o *CitrixCloudServicesAdministratorsApiModelsAdministratorsTokenClaimsModel) HasDirectoryContextClaim() bool`

HasDirectoryContextClaim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


