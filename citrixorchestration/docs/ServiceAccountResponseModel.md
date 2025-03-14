# ServiceAccountResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAccountUid** | Pointer to **string** |  | [optional] 
**IdentityProviderType** | Pointer to **NullableString** | Gets or sets the identity provider type for the service account | [optional] 
**IdentityProviderIdentifier** | Pointer to **NullableString** | Gets or sets the identity provider id for the service account | [optional] 
**AccountId** | Pointer to **NullableString** | Gets or sets the identifier for the service account | [optional] 
**SecretExpiryTime** | Pointer to **NullableString** | Gets or sets the secret expiration time for the service account | [optional] 
**Capabilities** | Pointer to [**[]ServiceAccountCapabilityReference**](ServiceAccountCapabilityReference.md) | Gets or sets the capabilities for the service account | [optional] 
**IsHealthy** | Pointer to **bool** | Gets or sets the healthy status for the service account | [optional] 
**FailureReason** | Pointer to **NullableString** | Gets or sets the failure reason for the service account if IsHealthy is false | [optional] 
**Scopes** | Pointer to [**[]ScopeResponseModel**](ScopeResponseModel.md) | Gets or sets the array of administration scopes | [optional] 
**TenantId** | Pointer to **NullableString** | Gets or sets the tenant id | [optional] 
**DisplayName** | Pointer to **NullableString** | Gets or sets the display name | [optional] 
**Description** | Pointer to **NullableString** | Gets or sets the description | [optional] 
**CustomProperties** | Pointer to **NullableString** | Gets or sets CustomProperties Format is the following: \&quot;CustomProperties\&quot;:\&quot;[{\\\&quot;Name\\\&quot;:\\\&quot;ProxyHypervisorTrafficThroughConnector\\\&quot;,\\\&quot;Value\\\&quot;:\\\&quot;true\\\&quot;},{\\\&quot;Name\\\&quot;:\\\&quot;ZoneUid\\\&quot;,\\\&quot;Value\\\&quot;:\\\&quot;4e1d7040-d830-4d97-8f94-342c03846f19\\\&quot;}]\&quot;. | [optional] 

## Methods

### NewServiceAccountResponseModel

`func NewServiceAccountResponseModel() *ServiceAccountResponseModel`

NewServiceAccountResponseModel instantiates a new ServiceAccountResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountResponseModelWithDefaults

`func NewServiceAccountResponseModelWithDefaults() *ServiceAccountResponseModel`

NewServiceAccountResponseModelWithDefaults instantiates a new ServiceAccountResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAccountUid

`func (o *ServiceAccountResponseModel) GetServiceAccountUid() string`

GetServiceAccountUid returns the ServiceAccountUid field if non-nil, zero value otherwise.

### GetServiceAccountUidOk

`func (o *ServiceAccountResponseModel) GetServiceAccountUidOk() (*string, bool)`

GetServiceAccountUidOk returns a tuple with the ServiceAccountUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountUid

`func (o *ServiceAccountResponseModel) SetServiceAccountUid(v string)`

SetServiceAccountUid sets ServiceAccountUid field to given value.

### HasServiceAccountUid

`func (o *ServiceAccountResponseModel) HasServiceAccountUid() bool`

HasServiceAccountUid returns a boolean if a field has been set.

### GetIdentityProviderType

`func (o *ServiceAccountResponseModel) GetIdentityProviderType() string`

GetIdentityProviderType returns the IdentityProviderType field if non-nil, zero value otherwise.

### GetIdentityProviderTypeOk

`func (o *ServiceAccountResponseModel) GetIdentityProviderTypeOk() (*string, bool)`

GetIdentityProviderTypeOk returns a tuple with the IdentityProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderType

`func (o *ServiceAccountResponseModel) SetIdentityProviderType(v string)`

SetIdentityProviderType sets IdentityProviderType field to given value.

### HasIdentityProviderType

`func (o *ServiceAccountResponseModel) HasIdentityProviderType() bool`

HasIdentityProviderType returns a boolean if a field has been set.

### SetIdentityProviderTypeNil

`func (o *ServiceAccountResponseModel) SetIdentityProviderTypeNil(b bool)`

 SetIdentityProviderTypeNil sets the value for IdentityProviderType to be an explicit nil

### UnsetIdentityProviderType
`func (o *ServiceAccountResponseModel) UnsetIdentityProviderType()`

UnsetIdentityProviderType ensures that no value is present for IdentityProviderType, not even an explicit nil
### GetIdentityProviderIdentifier

`func (o *ServiceAccountResponseModel) GetIdentityProviderIdentifier() string`

GetIdentityProviderIdentifier returns the IdentityProviderIdentifier field if non-nil, zero value otherwise.

### GetIdentityProviderIdentifierOk

`func (o *ServiceAccountResponseModel) GetIdentityProviderIdentifierOk() (*string, bool)`

GetIdentityProviderIdentifierOk returns a tuple with the IdentityProviderIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderIdentifier

`func (o *ServiceAccountResponseModel) SetIdentityProviderIdentifier(v string)`

SetIdentityProviderIdentifier sets IdentityProviderIdentifier field to given value.

### HasIdentityProviderIdentifier

`func (o *ServiceAccountResponseModel) HasIdentityProviderIdentifier() bool`

HasIdentityProviderIdentifier returns a boolean if a field has been set.

### SetIdentityProviderIdentifierNil

`func (o *ServiceAccountResponseModel) SetIdentityProviderIdentifierNil(b bool)`

 SetIdentityProviderIdentifierNil sets the value for IdentityProviderIdentifier to be an explicit nil

### UnsetIdentityProviderIdentifier
`func (o *ServiceAccountResponseModel) UnsetIdentityProviderIdentifier()`

UnsetIdentityProviderIdentifier ensures that no value is present for IdentityProviderIdentifier, not even an explicit nil
### GetAccountId

`func (o *ServiceAccountResponseModel) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ServiceAccountResponseModel) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ServiceAccountResponseModel) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ServiceAccountResponseModel) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *ServiceAccountResponseModel) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *ServiceAccountResponseModel) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetSecretExpiryTime

`func (o *ServiceAccountResponseModel) GetSecretExpiryTime() string`

GetSecretExpiryTime returns the SecretExpiryTime field if non-nil, zero value otherwise.

### GetSecretExpiryTimeOk

`func (o *ServiceAccountResponseModel) GetSecretExpiryTimeOk() (*string, bool)`

GetSecretExpiryTimeOk returns a tuple with the SecretExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretExpiryTime

`func (o *ServiceAccountResponseModel) SetSecretExpiryTime(v string)`

SetSecretExpiryTime sets SecretExpiryTime field to given value.

### HasSecretExpiryTime

`func (o *ServiceAccountResponseModel) HasSecretExpiryTime() bool`

HasSecretExpiryTime returns a boolean if a field has been set.

### SetSecretExpiryTimeNil

`func (o *ServiceAccountResponseModel) SetSecretExpiryTimeNil(b bool)`

 SetSecretExpiryTimeNil sets the value for SecretExpiryTime to be an explicit nil

### UnsetSecretExpiryTime
`func (o *ServiceAccountResponseModel) UnsetSecretExpiryTime()`

UnsetSecretExpiryTime ensures that no value is present for SecretExpiryTime, not even an explicit nil
### GetCapabilities

`func (o *ServiceAccountResponseModel) GetCapabilities() []ServiceAccountCapabilityReference`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ServiceAccountResponseModel) GetCapabilitiesOk() (*[]ServiceAccountCapabilityReference, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ServiceAccountResponseModel) SetCapabilities(v []ServiceAccountCapabilityReference)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ServiceAccountResponseModel) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *ServiceAccountResponseModel) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *ServiceAccountResponseModel) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetIsHealthy

`func (o *ServiceAccountResponseModel) GetIsHealthy() bool`

GetIsHealthy returns the IsHealthy field if non-nil, zero value otherwise.

### GetIsHealthyOk

`func (o *ServiceAccountResponseModel) GetIsHealthyOk() (*bool, bool)`

GetIsHealthyOk returns a tuple with the IsHealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHealthy

`func (o *ServiceAccountResponseModel) SetIsHealthy(v bool)`

SetIsHealthy sets IsHealthy field to given value.

### HasIsHealthy

`func (o *ServiceAccountResponseModel) HasIsHealthy() bool`

HasIsHealthy returns a boolean if a field has been set.

### GetFailureReason

`func (o *ServiceAccountResponseModel) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *ServiceAccountResponseModel) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *ServiceAccountResponseModel) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *ServiceAccountResponseModel) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReasonNil

`func (o *ServiceAccountResponseModel) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *ServiceAccountResponseModel) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
### GetScopes

`func (o *ServiceAccountResponseModel) GetScopes() []ScopeResponseModel`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ServiceAccountResponseModel) GetScopesOk() (*[]ScopeResponseModel, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ServiceAccountResponseModel) SetScopes(v []ScopeResponseModel)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ServiceAccountResponseModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *ServiceAccountResponseModel) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *ServiceAccountResponseModel) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetTenantId

`func (o *ServiceAccountResponseModel) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ServiceAccountResponseModel) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ServiceAccountResponseModel) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ServiceAccountResponseModel) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ServiceAccountResponseModel) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ServiceAccountResponseModel) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetDisplayName

`func (o *ServiceAccountResponseModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ServiceAccountResponseModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ServiceAccountResponseModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ServiceAccountResponseModel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ServiceAccountResponseModel) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ServiceAccountResponseModel) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *ServiceAccountResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceAccountResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceAccountResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceAccountResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ServiceAccountResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ServiceAccountResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCustomProperties

`func (o *ServiceAccountResponseModel) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ServiceAccountResponseModel) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ServiceAccountResponseModel) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ServiceAccountResponseModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *ServiceAccountResponseModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *ServiceAccountResponseModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


