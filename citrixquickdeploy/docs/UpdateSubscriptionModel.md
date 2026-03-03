# UpdateSubscriptionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsentedBy** | Pointer to **NullableString** | The name of the user that consented to using managed subscription | [optional] 
**ClientId** | Pointer to **NullableString** | ID of the Azure AD App to use with this subscription | [optional] 
**ClientSecret** | Pointer to **NullableString** | Client Secret of the Azure App | [optional] 
**SecretExpirationDate** | Pointer to **NullableTime** | The expiration date of the user provided secret, if one was used | [optional] 
**TenantId** | Pointer to **NullableString** | ID of the Auzre tenant the subscription and app belong to | [optional] 
**CspCustomer** | Pointer to **NullableString** | The identifier string containing the CSP customer ID and site ID. | [optional] 

## Methods

### NewUpdateSubscriptionModel

`func NewUpdateSubscriptionModel() *UpdateSubscriptionModel`

NewUpdateSubscriptionModel instantiates a new UpdateSubscriptionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSubscriptionModelWithDefaults

`func NewUpdateSubscriptionModelWithDefaults() *UpdateSubscriptionModel`

NewUpdateSubscriptionModelWithDefaults instantiates a new UpdateSubscriptionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsentedBy

`func (o *UpdateSubscriptionModel) GetConsentedBy() string`

GetConsentedBy returns the ConsentedBy field if non-nil, zero value otherwise.

### GetConsentedByOk

`func (o *UpdateSubscriptionModel) GetConsentedByOk() (*string, bool)`

GetConsentedByOk returns a tuple with the ConsentedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentedBy

`func (o *UpdateSubscriptionModel) SetConsentedBy(v string)`

SetConsentedBy sets ConsentedBy field to given value.

### HasConsentedBy

`func (o *UpdateSubscriptionModel) HasConsentedBy() bool`

HasConsentedBy returns a boolean if a field has been set.

### SetConsentedByNil

`func (o *UpdateSubscriptionModel) SetConsentedByNil(b bool)`

 SetConsentedByNil sets the value for ConsentedBy to be an explicit nil

### UnsetConsentedBy
`func (o *UpdateSubscriptionModel) UnsetConsentedBy()`

UnsetConsentedBy ensures that no value is present for ConsentedBy, not even an explicit nil
### GetClientId

`func (o *UpdateSubscriptionModel) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *UpdateSubscriptionModel) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *UpdateSubscriptionModel) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *UpdateSubscriptionModel) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *UpdateSubscriptionModel) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *UpdateSubscriptionModel) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetClientSecret

`func (o *UpdateSubscriptionModel) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *UpdateSubscriptionModel) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *UpdateSubscriptionModel) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *UpdateSubscriptionModel) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *UpdateSubscriptionModel) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *UpdateSubscriptionModel) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetSecretExpirationDate

`func (o *UpdateSubscriptionModel) GetSecretExpirationDate() time.Time`

GetSecretExpirationDate returns the SecretExpirationDate field if non-nil, zero value otherwise.

### GetSecretExpirationDateOk

`func (o *UpdateSubscriptionModel) GetSecretExpirationDateOk() (*time.Time, bool)`

GetSecretExpirationDateOk returns a tuple with the SecretExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretExpirationDate

`func (o *UpdateSubscriptionModel) SetSecretExpirationDate(v time.Time)`

SetSecretExpirationDate sets SecretExpirationDate field to given value.

### HasSecretExpirationDate

`func (o *UpdateSubscriptionModel) HasSecretExpirationDate() bool`

HasSecretExpirationDate returns a boolean if a field has been set.

### SetSecretExpirationDateNil

`func (o *UpdateSubscriptionModel) SetSecretExpirationDateNil(b bool)`

 SetSecretExpirationDateNil sets the value for SecretExpirationDate to be an explicit nil

### UnsetSecretExpirationDate
`func (o *UpdateSubscriptionModel) UnsetSecretExpirationDate()`

UnsetSecretExpirationDate ensures that no value is present for SecretExpirationDate, not even an explicit nil
### GetTenantId

`func (o *UpdateSubscriptionModel) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UpdateSubscriptionModel) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UpdateSubscriptionModel) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *UpdateSubscriptionModel) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *UpdateSubscriptionModel) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *UpdateSubscriptionModel) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetCspCustomer

`func (o *UpdateSubscriptionModel) GetCspCustomer() string`

GetCspCustomer returns the CspCustomer field if non-nil, zero value otherwise.

### GetCspCustomerOk

`func (o *UpdateSubscriptionModel) GetCspCustomerOk() (*string, bool)`

GetCspCustomerOk returns a tuple with the CspCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomer

`func (o *UpdateSubscriptionModel) SetCspCustomer(v string)`

SetCspCustomer sets CspCustomer field to given value.

### HasCspCustomer

`func (o *UpdateSubscriptionModel) HasCspCustomer() bool`

HasCspCustomer returns a boolean if a field has been set.

### SetCspCustomerNil

`func (o *UpdateSubscriptionModel) SetCspCustomerNil(b bool)`

 SetCspCustomerNil sets the value for CspCustomer to be an explicit nil

### UnsetCspCustomer
`func (o *UpdateSubscriptionModel) UnsetCspCustomer()`

UnsetCspCustomer ensures that no value is present for CspCustomer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


