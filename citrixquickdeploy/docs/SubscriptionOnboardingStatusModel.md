# SubscriptionOnboardingStatusModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudAwareAppId** | Pointer to **NullableString** | Application Id of the cloud aware app service principal added to the subscription | [optional] 
**AzurePoolSubscriptionDetailsList** | Pointer to [**[]AzurePoolSubscriptionDetails**](AzurePoolSubscriptionDetails.md) | Subscription details | [optional] 

## Methods

### NewSubscriptionOnboardingStatusModel

`func NewSubscriptionOnboardingStatusModel() *SubscriptionOnboardingStatusModel`

NewSubscriptionOnboardingStatusModel instantiates a new SubscriptionOnboardingStatusModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionOnboardingStatusModelWithDefaults

`func NewSubscriptionOnboardingStatusModelWithDefaults() *SubscriptionOnboardingStatusModel`

NewSubscriptionOnboardingStatusModelWithDefaults instantiates a new SubscriptionOnboardingStatusModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudAwareAppId

`func (o *SubscriptionOnboardingStatusModel) GetCloudAwareAppId() string`

GetCloudAwareAppId returns the CloudAwareAppId field if non-nil, zero value otherwise.

### GetCloudAwareAppIdOk

`func (o *SubscriptionOnboardingStatusModel) GetCloudAwareAppIdOk() (*string, bool)`

GetCloudAwareAppIdOk returns a tuple with the CloudAwareAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAwareAppId

`func (o *SubscriptionOnboardingStatusModel) SetCloudAwareAppId(v string)`

SetCloudAwareAppId sets CloudAwareAppId field to given value.

### HasCloudAwareAppId

`func (o *SubscriptionOnboardingStatusModel) HasCloudAwareAppId() bool`

HasCloudAwareAppId returns a boolean if a field has been set.

### SetCloudAwareAppIdNil

`func (o *SubscriptionOnboardingStatusModel) SetCloudAwareAppIdNil(b bool)`

 SetCloudAwareAppIdNil sets the value for CloudAwareAppId to be an explicit nil

### UnsetCloudAwareAppId
`func (o *SubscriptionOnboardingStatusModel) UnsetCloudAwareAppId()`

UnsetCloudAwareAppId ensures that no value is present for CloudAwareAppId, not even an explicit nil
### GetAzurePoolSubscriptionDetailsList

`func (o *SubscriptionOnboardingStatusModel) GetAzurePoolSubscriptionDetailsList() []AzurePoolSubscriptionDetails`

GetAzurePoolSubscriptionDetailsList returns the AzurePoolSubscriptionDetailsList field if non-nil, zero value otherwise.

### GetAzurePoolSubscriptionDetailsListOk

`func (o *SubscriptionOnboardingStatusModel) GetAzurePoolSubscriptionDetailsListOk() (*[]AzurePoolSubscriptionDetails, bool)`

GetAzurePoolSubscriptionDetailsListOk returns a tuple with the AzurePoolSubscriptionDetailsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzurePoolSubscriptionDetailsList

`func (o *SubscriptionOnboardingStatusModel) SetAzurePoolSubscriptionDetailsList(v []AzurePoolSubscriptionDetails)`

SetAzurePoolSubscriptionDetailsList sets AzurePoolSubscriptionDetailsList field to given value.

### HasAzurePoolSubscriptionDetailsList

`func (o *SubscriptionOnboardingStatusModel) HasAzurePoolSubscriptionDetailsList() bool`

HasAzurePoolSubscriptionDetailsList returns a boolean if a field has been set.

### SetAzurePoolSubscriptionDetailsListNil

`func (o *SubscriptionOnboardingStatusModel) SetAzurePoolSubscriptionDetailsListNil(b bool)`

 SetAzurePoolSubscriptionDetailsListNil sets the value for AzurePoolSubscriptionDetailsList to be an explicit nil

### UnsetAzurePoolSubscriptionDetailsList
`func (o *SubscriptionOnboardingStatusModel) UnsetAzurePoolSubscriptionDetailsList()`

UnsetAzurePoolSubscriptionDetailsList ensures that no value is present for AzurePoolSubscriptionDetailsList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


