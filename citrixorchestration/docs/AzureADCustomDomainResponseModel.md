# AzureADCustomDomainResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **NullableString** | The tenant ID. For example, 00000000-0000-0000-0000-000000000000. | [optional] 
**TenantDisplayName** | Pointer to **NullableString** | The display name of the tenant. | [optional] 
**AuthenticationType** | Pointer to **NullableString** | The authentication type of the domain. Valid values are: Managed and Federated. | [optional] 

## Methods

### NewAzureADCustomDomainResponseModel

`func NewAzureADCustomDomainResponseModel() *AzureADCustomDomainResponseModel`

NewAzureADCustomDomainResponseModel instantiates a new AzureADCustomDomainResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureADCustomDomainResponseModelWithDefaults

`func NewAzureADCustomDomainResponseModelWithDefaults() *AzureADCustomDomainResponseModel`

NewAzureADCustomDomainResponseModelWithDefaults instantiates a new AzureADCustomDomainResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *AzureADCustomDomainResponseModel) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureADCustomDomainResponseModel) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureADCustomDomainResponseModel) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AzureADCustomDomainResponseModel) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *AzureADCustomDomainResponseModel) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *AzureADCustomDomainResponseModel) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetTenantDisplayName

`func (o *AzureADCustomDomainResponseModel) GetTenantDisplayName() string`

GetTenantDisplayName returns the TenantDisplayName field if non-nil, zero value otherwise.

### GetTenantDisplayNameOk

`func (o *AzureADCustomDomainResponseModel) GetTenantDisplayNameOk() (*string, bool)`

GetTenantDisplayNameOk returns a tuple with the TenantDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantDisplayName

`func (o *AzureADCustomDomainResponseModel) SetTenantDisplayName(v string)`

SetTenantDisplayName sets TenantDisplayName field to given value.

### HasTenantDisplayName

`func (o *AzureADCustomDomainResponseModel) HasTenantDisplayName() bool`

HasTenantDisplayName returns a boolean if a field has been set.

### SetTenantDisplayNameNil

`func (o *AzureADCustomDomainResponseModel) SetTenantDisplayNameNil(b bool)`

 SetTenantDisplayNameNil sets the value for TenantDisplayName to be an explicit nil

### UnsetTenantDisplayName
`func (o *AzureADCustomDomainResponseModel) UnsetTenantDisplayName()`

UnsetTenantDisplayName ensures that no value is present for TenantDisplayName, not even an explicit nil
### GetAuthenticationType

`func (o *AzureADCustomDomainResponseModel) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *AzureADCustomDomainResponseModel) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *AzureADCustomDomainResponseModel) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *AzureADCustomDomainResponseModel) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### SetAuthenticationTypeNil

`func (o *AzureADCustomDomainResponseModel) SetAuthenticationTypeNil(b bool)`

 SetAuthenticationTypeNil sets the value for AuthenticationType to be an explicit nil

### UnsetAuthenticationType
`func (o *AzureADCustomDomainResponseModel) UnsetAuthenticationType()`

UnsetAuthenticationType ensures that no value is present for AuthenticationType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


