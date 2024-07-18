# UpdateServiceAccountRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **NullableString** | The identifier for the service account | [optional] 
**AccountSecret** | Pointer to **NullableString** | The secret for the service account. E.g. Azure application (client) secret if &#39;IdentityProviderType&#39; is AzureAD. The secret will be encrypted and stored in database. | [optional] 
**AccountSecretFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**SecretExpiryTime** | Pointer to **NullableString** | The secret expiration time for the service account | [optional] 
**Capabilities** | Pointer to **[]string** | Gets or sets capabilities for the service account | [optional] 
**DisplayName** | Pointer to **NullableString** | Gets or sets the display name | [optional] 
**Description** | Pointer to **NullableString** | Gets or sets the description | [optional] 
**Scopes** | Pointer to **[]string** | Gets or sets the scopes for the service account. | [optional] 
**Tenants** | Pointer to **[]string** | Tenants to associate with the service account. | [optional] 

## Methods

### NewUpdateServiceAccountRequestModel

`func NewUpdateServiceAccountRequestModel() *UpdateServiceAccountRequestModel`

NewUpdateServiceAccountRequestModel instantiates a new UpdateServiceAccountRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServiceAccountRequestModelWithDefaults

`func NewUpdateServiceAccountRequestModelWithDefaults() *UpdateServiceAccountRequestModel`

NewUpdateServiceAccountRequestModelWithDefaults instantiates a new UpdateServiceAccountRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *UpdateServiceAccountRequestModel) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateServiceAccountRequestModel) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateServiceAccountRequestModel) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UpdateServiceAccountRequestModel) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *UpdateServiceAccountRequestModel) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *UpdateServiceAccountRequestModel) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetAccountSecret

`func (o *UpdateServiceAccountRequestModel) GetAccountSecret() string`

GetAccountSecret returns the AccountSecret field if non-nil, zero value otherwise.

### GetAccountSecretOk

`func (o *UpdateServiceAccountRequestModel) GetAccountSecretOk() (*string, bool)`

GetAccountSecretOk returns a tuple with the AccountSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSecret

`func (o *UpdateServiceAccountRequestModel) SetAccountSecret(v string)`

SetAccountSecret sets AccountSecret field to given value.

### HasAccountSecret

`func (o *UpdateServiceAccountRequestModel) HasAccountSecret() bool`

HasAccountSecret returns a boolean if a field has been set.

### SetAccountSecretNil

`func (o *UpdateServiceAccountRequestModel) SetAccountSecretNil(b bool)`

 SetAccountSecretNil sets the value for AccountSecret to be an explicit nil

### UnsetAccountSecret
`func (o *UpdateServiceAccountRequestModel) UnsetAccountSecret()`

UnsetAccountSecret ensures that no value is present for AccountSecret, not even an explicit nil
### GetAccountSecretFormat

`func (o *UpdateServiceAccountRequestModel) GetAccountSecretFormat() IdentityPasswordFormat`

GetAccountSecretFormat returns the AccountSecretFormat field if non-nil, zero value otherwise.

### GetAccountSecretFormatOk

`func (o *UpdateServiceAccountRequestModel) GetAccountSecretFormatOk() (*IdentityPasswordFormat, bool)`

GetAccountSecretFormatOk returns a tuple with the AccountSecretFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSecretFormat

`func (o *UpdateServiceAccountRequestModel) SetAccountSecretFormat(v IdentityPasswordFormat)`

SetAccountSecretFormat sets AccountSecretFormat field to given value.

### HasAccountSecretFormat

`func (o *UpdateServiceAccountRequestModel) HasAccountSecretFormat() bool`

HasAccountSecretFormat returns a boolean if a field has been set.

### GetSecretExpiryTime

`func (o *UpdateServiceAccountRequestModel) GetSecretExpiryTime() string`

GetSecretExpiryTime returns the SecretExpiryTime field if non-nil, zero value otherwise.

### GetSecretExpiryTimeOk

`func (o *UpdateServiceAccountRequestModel) GetSecretExpiryTimeOk() (*string, bool)`

GetSecretExpiryTimeOk returns a tuple with the SecretExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretExpiryTime

`func (o *UpdateServiceAccountRequestModel) SetSecretExpiryTime(v string)`

SetSecretExpiryTime sets SecretExpiryTime field to given value.

### HasSecretExpiryTime

`func (o *UpdateServiceAccountRequestModel) HasSecretExpiryTime() bool`

HasSecretExpiryTime returns a boolean if a field has been set.

### SetSecretExpiryTimeNil

`func (o *UpdateServiceAccountRequestModel) SetSecretExpiryTimeNil(b bool)`

 SetSecretExpiryTimeNil sets the value for SecretExpiryTime to be an explicit nil

### UnsetSecretExpiryTime
`func (o *UpdateServiceAccountRequestModel) UnsetSecretExpiryTime()`

UnsetSecretExpiryTime ensures that no value is present for SecretExpiryTime, not even an explicit nil
### GetCapabilities

`func (o *UpdateServiceAccountRequestModel) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *UpdateServiceAccountRequestModel) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *UpdateServiceAccountRequestModel) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *UpdateServiceAccountRequestModel) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *UpdateServiceAccountRequestModel) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *UpdateServiceAccountRequestModel) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetDisplayName

`func (o *UpdateServiceAccountRequestModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateServiceAccountRequestModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateServiceAccountRequestModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateServiceAccountRequestModel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *UpdateServiceAccountRequestModel) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *UpdateServiceAccountRequestModel) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *UpdateServiceAccountRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateServiceAccountRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateServiceAccountRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateServiceAccountRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateServiceAccountRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateServiceAccountRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetScopes

`func (o *UpdateServiceAccountRequestModel) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *UpdateServiceAccountRequestModel) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *UpdateServiceAccountRequestModel) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *UpdateServiceAccountRequestModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *UpdateServiceAccountRequestModel) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *UpdateServiceAccountRequestModel) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetTenants

`func (o *UpdateServiceAccountRequestModel) GetTenants() []string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *UpdateServiceAccountRequestModel) GetTenantsOk() (*[]string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *UpdateServiceAccountRequestModel) SetTenants(v []string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *UpdateServiceAccountRequestModel) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *UpdateServiceAccountRequestModel) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *UpdateServiceAccountRequestModel) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


