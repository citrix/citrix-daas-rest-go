# AddAzureSubscriptionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | Pointer to **NullableString** | ID of the subscription to add | [optional] 
**CitrixManaged** | Pointer to **bool** |  | [optional] 
**CspCustomerId** | Pointer to **NullableString** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 
**CspSiteId** | Pointer to **NullableString** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 
**ConsentedBy** | Pointer to **NullableString** | The user that consented to using Citrix managed subscription | [optional] 
**SecretExpirationDate** | Pointer to **NullableTime** | The expiration date of the user provided secret, if one was given in the headers | [optional] 

## Methods

### NewAddAzureSubscriptionModel

`func NewAddAzureSubscriptionModel() *AddAzureSubscriptionModel`

NewAddAzureSubscriptionModel instantiates a new AddAzureSubscriptionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAzureSubscriptionModelWithDefaults

`func NewAddAzureSubscriptionModelWithDefaults() *AddAzureSubscriptionModel`

NewAddAzureSubscriptionModelWithDefaults instantiates a new AddAzureSubscriptionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *AddAzureSubscriptionModel) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AddAzureSubscriptionModel) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AddAzureSubscriptionModel) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AddAzureSubscriptionModel) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *AddAzureSubscriptionModel) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *AddAzureSubscriptionModel) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetCitrixManaged

`func (o *AddAzureSubscriptionModel) GetCitrixManaged() bool`

GetCitrixManaged returns the CitrixManaged field if non-nil, zero value otherwise.

### GetCitrixManagedOk

`func (o *AddAzureSubscriptionModel) GetCitrixManagedOk() (*bool, bool)`

GetCitrixManagedOk returns a tuple with the CitrixManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitrixManaged

`func (o *AddAzureSubscriptionModel) SetCitrixManaged(v bool)`

SetCitrixManaged sets CitrixManaged field to given value.

### HasCitrixManaged

`func (o *AddAzureSubscriptionModel) HasCitrixManaged() bool`

HasCitrixManaged returns a boolean if a field has been set.

### GetCspCustomerId

`func (o *AddAzureSubscriptionModel) GetCspCustomerId() string`

GetCspCustomerId returns the CspCustomerId field if non-nil, zero value otherwise.

### GetCspCustomerIdOk

`func (o *AddAzureSubscriptionModel) GetCspCustomerIdOk() (*string, bool)`

GetCspCustomerIdOk returns a tuple with the CspCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomerId

`func (o *AddAzureSubscriptionModel) SetCspCustomerId(v string)`

SetCspCustomerId sets CspCustomerId field to given value.

### HasCspCustomerId

`func (o *AddAzureSubscriptionModel) HasCspCustomerId() bool`

HasCspCustomerId returns a boolean if a field has been set.

### SetCspCustomerIdNil

`func (o *AddAzureSubscriptionModel) SetCspCustomerIdNil(b bool)`

 SetCspCustomerIdNil sets the value for CspCustomerId to be an explicit nil

### UnsetCspCustomerId
`func (o *AddAzureSubscriptionModel) UnsetCspCustomerId()`

UnsetCspCustomerId ensures that no value is present for CspCustomerId, not even an explicit nil
### GetCspSiteId

`func (o *AddAzureSubscriptionModel) GetCspSiteId() string`

GetCspSiteId returns the CspSiteId field if non-nil, zero value otherwise.

### GetCspSiteIdOk

`func (o *AddAzureSubscriptionModel) GetCspSiteIdOk() (*string, bool)`

GetCspSiteIdOk returns a tuple with the CspSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspSiteId

`func (o *AddAzureSubscriptionModel) SetCspSiteId(v string)`

SetCspSiteId sets CspSiteId field to given value.

### HasCspSiteId

`func (o *AddAzureSubscriptionModel) HasCspSiteId() bool`

HasCspSiteId returns a boolean if a field has been set.

### SetCspSiteIdNil

`func (o *AddAzureSubscriptionModel) SetCspSiteIdNil(b bool)`

 SetCspSiteIdNil sets the value for CspSiteId to be an explicit nil

### UnsetCspSiteId
`func (o *AddAzureSubscriptionModel) UnsetCspSiteId()`

UnsetCspSiteId ensures that no value is present for CspSiteId, not even an explicit nil
### GetConsentedBy

`func (o *AddAzureSubscriptionModel) GetConsentedBy() string`

GetConsentedBy returns the ConsentedBy field if non-nil, zero value otherwise.

### GetConsentedByOk

`func (o *AddAzureSubscriptionModel) GetConsentedByOk() (*string, bool)`

GetConsentedByOk returns a tuple with the ConsentedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentedBy

`func (o *AddAzureSubscriptionModel) SetConsentedBy(v string)`

SetConsentedBy sets ConsentedBy field to given value.

### HasConsentedBy

`func (o *AddAzureSubscriptionModel) HasConsentedBy() bool`

HasConsentedBy returns a boolean if a field has been set.

### SetConsentedByNil

`func (o *AddAzureSubscriptionModel) SetConsentedByNil(b bool)`

 SetConsentedByNil sets the value for ConsentedBy to be an explicit nil

### UnsetConsentedBy
`func (o *AddAzureSubscriptionModel) UnsetConsentedBy()`

UnsetConsentedBy ensures that no value is present for ConsentedBy, not even an explicit nil
### GetSecretExpirationDate

`func (o *AddAzureSubscriptionModel) GetSecretExpirationDate() time.Time`

GetSecretExpirationDate returns the SecretExpirationDate field if non-nil, zero value otherwise.

### GetSecretExpirationDateOk

`func (o *AddAzureSubscriptionModel) GetSecretExpirationDateOk() (*time.Time, bool)`

GetSecretExpirationDateOk returns a tuple with the SecretExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretExpirationDate

`func (o *AddAzureSubscriptionModel) SetSecretExpirationDate(v time.Time)`

SetSecretExpirationDate sets SecretExpirationDate field to given value.

### HasSecretExpirationDate

`func (o *AddAzureSubscriptionModel) HasSecretExpirationDate() bool`

HasSecretExpirationDate returns a boolean if a field has been set.

### SetSecretExpirationDateNil

`func (o *AddAzureSubscriptionModel) SetSecretExpirationDateNil(b bool)`

 SetSecretExpirationDateNil sets the value for SecretExpirationDate to be an explicit nil

### UnsetSecretExpirationDate
`func (o *AddAzureSubscriptionModel) UnsetSecretExpirationDate()`

UnsetSecretExpirationDate ensures that no value is present for SecretExpirationDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


