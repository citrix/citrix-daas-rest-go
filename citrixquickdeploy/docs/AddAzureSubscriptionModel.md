# AddAzureSubscriptionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | Pointer to **string** | ID of the subscription to add | [optional] 
**CitrixManaged** | Pointer to **bool** |  | [optional] 
**CspCustomerId** | Pointer to **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 
**CspSiteId** | Pointer to **string** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 
**ConsentedBy** | Pointer to **string** | The user that consented to using Citrix managed subscription | [optional] 
**SecretExpirationDate** | Pointer to **time.Time** | The expiration date of the user provided secret, if one was given in the headers | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


