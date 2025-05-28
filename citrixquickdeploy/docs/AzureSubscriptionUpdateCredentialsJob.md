# AzureSubscriptionUpdateCredentialsJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | Pointer to **string** | ID of the transaction associated with the update. | [optional] 
**State** | Pointer to [**UpdateAzureSubscriptionCredentialJobState**](UpdateAzureSubscriptionCredentialJobState.md) | State of the update job. | [optional] 
**FailedCatalogIds** | Pointer to **[]string** | List of catalog IDs that failed to update. | [optional] 
**StartedAt** | Pointer to **time.Time** | The datetime when the job started | [optional] 

## Methods

### NewAzureSubscriptionUpdateCredentialsJob

`func NewAzureSubscriptionUpdateCredentialsJob() *AzureSubscriptionUpdateCredentialsJob`

NewAzureSubscriptionUpdateCredentialsJob instantiates a new AzureSubscriptionUpdateCredentialsJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureSubscriptionUpdateCredentialsJobWithDefaults

`func NewAzureSubscriptionUpdateCredentialsJobWithDefaults() *AzureSubscriptionUpdateCredentialsJob`

NewAzureSubscriptionUpdateCredentialsJobWithDefaults instantiates a new AzureSubscriptionUpdateCredentialsJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *AzureSubscriptionUpdateCredentialsJob) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *AzureSubscriptionUpdateCredentialsJob) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *AzureSubscriptionUpdateCredentialsJob) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *AzureSubscriptionUpdateCredentialsJob) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetState

`func (o *AzureSubscriptionUpdateCredentialsJob) GetState() UpdateAzureSubscriptionCredentialJobState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AzureSubscriptionUpdateCredentialsJob) GetStateOk() (*UpdateAzureSubscriptionCredentialJobState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AzureSubscriptionUpdateCredentialsJob) SetState(v UpdateAzureSubscriptionCredentialJobState)`

SetState sets State field to given value.

### HasState

`func (o *AzureSubscriptionUpdateCredentialsJob) HasState() bool`

HasState returns a boolean if a field has been set.

### GetFailedCatalogIds

`func (o *AzureSubscriptionUpdateCredentialsJob) GetFailedCatalogIds() []string`

GetFailedCatalogIds returns the FailedCatalogIds field if non-nil, zero value otherwise.

### GetFailedCatalogIdsOk

`func (o *AzureSubscriptionUpdateCredentialsJob) GetFailedCatalogIdsOk() (*[]string, bool)`

GetFailedCatalogIdsOk returns a tuple with the FailedCatalogIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCatalogIds

`func (o *AzureSubscriptionUpdateCredentialsJob) SetFailedCatalogIds(v []string)`

SetFailedCatalogIds sets FailedCatalogIds field to given value.

### HasFailedCatalogIds

`func (o *AzureSubscriptionUpdateCredentialsJob) HasFailedCatalogIds() bool`

HasFailedCatalogIds returns a boolean if a field has been set.

### GetStartedAt

`func (o *AzureSubscriptionUpdateCredentialsJob) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *AzureSubscriptionUpdateCredentialsJob) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *AzureSubscriptionUpdateCredentialsJob) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *AzureSubscriptionUpdateCredentialsJob) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


