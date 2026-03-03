# SingleTenantSubscriptionOnboardingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | ClientId of the Citrix SPN created within the tenant being on-boarded | 
**ClientSecret** | **string** | ClientSecret of the Citrix SPN created within the tenant being on-boarded | 
**SubscriptionId** | **string** | Id of the subscription being on-boarded | 
**TenantId** | Pointer to **NullableString** | Id of the tenant being on-boarded | [optional] 
**PoolName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSingleTenantSubscriptionOnboardingModel

`func NewSingleTenantSubscriptionOnboardingModel(clientId string, clientSecret string, subscriptionId string, ) *SingleTenantSubscriptionOnboardingModel`

NewSingleTenantSubscriptionOnboardingModel instantiates a new SingleTenantSubscriptionOnboardingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleTenantSubscriptionOnboardingModelWithDefaults

`func NewSingleTenantSubscriptionOnboardingModelWithDefaults() *SingleTenantSubscriptionOnboardingModel`

NewSingleTenantSubscriptionOnboardingModelWithDefaults instantiates a new SingleTenantSubscriptionOnboardingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *SingleTenantSubscriptionOnboardingModel) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SingleTenantSubscriptionOnboardingModel) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SingleTenantSubscriptionOnboardingModel) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *SingleTenantSubscriptionOnboardingModel) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *SingleTenantSubscriptionOnboardingModel) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *SingleTenantSubscriptionOnboardingModel) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetSubscriptionId

`func (o *SingleTenantSubscriptionOnboardingModel) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SingleTenantSubscriptionOnboardingModel) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SingleTenantSubscriptionOnboardingModel) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetTenantId

`func (o *SingleTenantSubscriptionOnboardingModel) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SingleTenantSubscriptionOnboardingModel) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SingleTenantSubscriptionOnboardingModel) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SingleTenantSubscriptionOnboardingModel) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SingleTenantSubscriptionOnboardingModel) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SingleTenantSubscriptionOnboardingModel) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetPoolName

`func (o *SingleTenantSubscriptionOnboardingModel) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *SingleTenantSubscriptionOnboardingModel) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *SingleTenantSubscriptionOnboardingModel) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *SingleTenantSubscriptionOnboardingModel) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### SetPoolNameNil

`func (o *SingleTenantSubscriptionOnboardingModel) SetPoolNameNil(b bool)`

 SetPoolNameNil sets the value for PoolName to be an explicit nil

### UnsetPoolName
`func (o *SingleTenantSubscriptionOnboardingModel) UnsetPoolName()`

UnsetPoolName ensures that no value is present for PoolName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


