# SubscriptionOnboardingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** | Id of the subscription being on-boarded | 
**TenantId** | Pointer to **NullableString** | Id of the tenant being on-boarded | [optional] 
**PoolName** | Pointer to **NullableString** | Name of the Azure pool where the subscription needs to be added | [optional] 

## Methods

### NewSubscriptionOnboardingModel

`func NewSubscriptionOnboardingModel(subscriptionId string, ) *SubscriptionOnboardingModel`

NewSubscriptionOnboardingModel instantiates a new SubscriptionOnboardingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionOnboardingModelWithDefaults

`func NewSubscriptionOnboardingModelWithDefaults() *SubscriptionOnboardingModel`

NewSubscriptionOnboardingModelWithDefaults instantiates a new SubscriptionOnboardingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### SetTenantIdNil

`func (o *SubscriptionOnboardingModel) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SubscriptionOnboardingModel) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
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

### SetPoolNameNil

`func (o *SubscriptionOnboardingModel) SetPoolNameNil(b bool)`

 SetPoolNameNil sets the value for PoolName to be an explicit nil

### UnsetPoolName
`func (o *SubscriptionOnboardingModel) UnsetPoolName()`

UnsetPoolName ensures that no value is present for PoolName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


