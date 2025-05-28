# SubscriptionOnboardingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | ClientId of the Citrix SPN created within the tenant being on-boarded | 
**ClientSecret** | **string** | ClientSecret of the Citrix SPN created within the tenant being on-boarded | 
**SubscriptionId** | **string** | Id of the subscription being on-boarded | 
**TenantId** | Pointer to **string** | Id of the tenant being on-boarded | [optional] 
**PoolName** | Pointer to **string** | Name of the Azure pool where the subscription needs to be added | [optional] 

## Methods

### NewSubscriptionOnboardingModel

`func NewSubscriptionOnboardingModel(clientId string, clientSecret string, subscriptionId string, ) *SubscriptionOnboardingModel`

NewSubscriptionOnboardingModel instantiates a new SubscriptionOnboardingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionOnboardingModelWithDefaults

`func NewSubscriptionOnboardingModelWithDefaults() *SubscriptionOnboardingModel`

NewSubscriptionOnboardingModelWithDefaults instantiates a new SubscriptionOnboardingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *SubscriptionOnboardingModel) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SubscriptionOnboardingModel) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SubscriptionOnboardingModel) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *SubscriptionOnboardingModel) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *SubscriptionOnboardingModel) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *SubscriptionOnboardingModel) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetSubscriptionId

`func (o *SubscriptionOnboardingModel) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SubscriptionOnboardingModel) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SubscriptionOnboardingModel) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetTenantId

`func (o *SubscriptionOnboardingModel) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SubscriptionOnboardingModel) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SubscriptionOnboardingModel) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SubscriptionOnboardingModel) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetPoolName

`func (o *SubscriptionOnboardingModel) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *SubscriptionOnboardingModel) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *SubscriptionOnboardingModel) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *SubscriptionOnboardingModel) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


