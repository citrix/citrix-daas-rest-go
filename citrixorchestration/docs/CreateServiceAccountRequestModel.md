# CreateServiceAccountRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityProviderType** | **string** | The identity provider type for the service account | 
**IdentityProviderIdentifier** | **string** | The identity provider id for the service account | 
**AccountId** | **string** | The identifier for the service account | 
**SecretExpiryTime** | **string** | The secret expiration time for the service account | 
**AccountSecret** | **string** | The secret for the service account. E.g. Azure application (client) secret if &#39;IdentityProviderType&#39; is AzureAD. | 
**AccountSecretFormat** | Pointer to [**IdentityPasswordFormat**](IdentityPasswordFormat.md) |  | [optional] 
**Capabilities** | Pointer to **[]string** | The capabilities for the service account | [optional] 
**Scopes** | Pointer to **[]string** | The scopes for this object | [optional] 
**TenantId** | Pointer to **NullableString** | The tenant id | [optional] 
**DisplayName** | Pointer to **NullableString** | Gets or sets the display name | [optional] 
**Description** | Pointer to **NullableString** | Gets or sets the description | [optional] 
**CustomProperties** | Pointer to **NullableString** | Gets or sets the custom properties Format is the following: \&quot;CustomProperties\&quot;:\&quot;[{\\\&quot;Name\\\&quot;:\\\&quot;ProxyHypervisorTrafficThroughConnector\\\&quot;,\\\&quot;Value\\\&quot;:\\\&quot;true\\\&quot;},{\\\&quot;Name\\\&quot;:\\\&quot;ZoneUid\\\&quot;,\\\&quot;Value\\\&quot;:\\\&quot;4e1d7040-d830-4d97-8f94-342c03846f19\\\&quot;}]\&quot;. | [optional] 

## Methods

### NewCreateServiceAccountRequestModel

`func NewCreateServiceAccountRequestModel(identityProviderType string, identityProviderIdentifier string, accountId string, secretExpiryTime string, accountSecret string, ) *CreateServiceAccountRequestModel`

NewCreateServiceAccountRequestModel instantiates a new CreateServiceAccountRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceAccountRequestModelWithDefaults

`func NewCreateServiceAccountRequestModelWithDefaults() *CreateServiceAccountRequestModel`

NewCreateServiceAccountRequestModelWithDefaults instantiates a new CreateServiceAccountRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityProviderType

`func (o *CreateServiceAccountRequestModel) GetIdentityProviderType() string`

GetIdentityProviderType returns the IdentityProviderType field if non-nil, zero value otherwise.

### GetIdentityProviderTypeOk

`func (o *CreateServiceAccountRequestModel) GetIdentityProviderTypeOk() (*string, bool)`

GetIdentityProviderTypeOk returns a tuple with the IdentityProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderType

`func (o *CreateServiceAccountRequestModel) SetIdentityProviderType(v string)`

SetIdentityProviderType sets IdentityProviderType field to given value.


### GetIdentityProviderIdentifier

`func (o *CreateServiceAccountRequestModel) GetIdentityProviderIdentifier() string`

GetIdentityProviderIdentifier returns the IdentityProviderIdentifier field if non-nil, zero value otherwise.

### GetIdentityProviderIdentifierOk

`func (o *CreateServiceAccountRequestModel) GetIdentityProviderIdentifierOk() (*string, bool)`

GetIdentityProviderIdentifierOk returns a tuple with the IdentityProviderIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderIdentifier

`func (o *CreateServiceAccountRequestModel) SetIdentityProviderIdentifier(v string)`

SetIdentityProviderIdentifier sets IdentityProviderIdentifier field to given value.


### GetAccountId

`func (o *CreateServiceAccountRequestModel) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateServiceAccountRequestModel) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateServiceAccountRequestModel) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetSecretExpiryTime

`func (o *CreateServiceAccountRequestModel) GetSecretExpiryTime() string`

GetSecretExpiryTime returns the SecretExpiryTime field if non-nil, zero value otherwise.

### GetSecretExpiryTimeOk

`func (o *CreateServiceAccountRequestModel) GetSecretExpiryTimeOk() (*string, bool)`

GetSecretExpiryTimeOk returns a tuple with the SecretExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretExpiryTime

`func (o *CreateServiceAccountRequestModel) SetSecretExpiryTime(v string)`

SetSecretExpiryTime sets SecretExpiryTime field to given value.


### GetAccountSecret

`func (o *CreateServiceAccountRequestModel) GetAccountSecret() string`

GetAccountSecret returns the AccountSecret field if non-nil, zero value otherwise.

### GetAccountSecretOk

`func (o *CreateServiceAccountRequestModel) GetAccountSecretOk() (*string, bool)`

GetAccountSecretOk returns a tuple with the AccountSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSecret

`func (o *CreateServiceAccountRequestModel) SetAccountSecret(v string)`

SetAccountSecret sets AccountSecret field to given value.


### GetAccountSecretFormat

`func (o *CreateServiceAccountRequestModel) GetAccountSecretFormat() IdentityPasswordFormat`

GetAccountSecretFormat returns the AccountSecretFormat field if non-nil, zero value otherwise.

### GetAccountSecretFormatOk

`func (o *CreateServiceAccountRequestModel) GetAccountSecretFormatOk() (*IdentityPasswordFormat, bool)`

GetAccountSecretFormatOk returns a tuple with the AccountSecretFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSecretFormat

`func (o *CreateServiceAccountRequestModel) SetAccountSecretFormat(v IdentityPasswordFormat)`

SetAccountSecretFormat sets AccountSecretFormat field to given value.

### HasAccountSecretFormat

`func (o *CreateServiceAccountRequestModel) HasAccountSecretFormat() bool`

HasAccountSecretFormat returns a boolean if a field has been set.

### GetCapabilities

`func (o *CreateServiceAccountRequestModel) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *CreateServiceAccountRequestModel) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *CreateServiceAccountRequestModel) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *CreateServiceAccountRequestModel) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *CreateServiceAccountRequestModel) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *CreateServiceAccountRequestModel) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetScopes

`func (o *CreateServiceAccountRequestModel) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CreateServiceAccountRequestModel) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CreateServiceAccountRequestModel) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *CreateServiceAccountRequestModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *CreateServiceAccountRequestModel) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *CreateServiceAccountRequestModel) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetTenantId

`func (o *CreateServiceAccountRequestModel) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateServiceAccountRequestModel) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateServiceAccountRequestModel) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CreateServiceAccountRequestModel) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CreateServiceAccountRequestModel) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CreateServiceAccountRequestModel) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetDisplayName

`func (o *CreateServiceAccountRequestModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateServiceAccountRequestModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateServiceAccountRequestModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateServiceAccountRequestModel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CreateServiceAccountRequestModel) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CreateServiceAccountRequestModel) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *CreateServiceAccountRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateServiceAccountRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateServiceAccountRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateServiceAccountRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateServiceAccountRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateServiceAccountRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCustomProperties

`func (o *CreateServiceAccountRequestModel) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *CreateServiceAccountRequestModel) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *CreateServiceAccountRequestModel) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *CreateServiceAccountRequestModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *CreateServiceAccountRequestModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *CreateServiceAccountRequestModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


