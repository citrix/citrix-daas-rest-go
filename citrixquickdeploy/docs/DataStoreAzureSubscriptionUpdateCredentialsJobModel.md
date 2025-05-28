# DataStoreAzureSubscriptionUpdateCredentialsJobModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | Pointer to **string** | ID of the transaction associated with the update. | [optional] 
**State** | Pointer to [**UpdateAzureSubscriptionCredentialJobState**](UpdateAzureSubscriptionCredentialJobState.md) | State of the update job. | [optional] 
**FailedCatalogIds** | Pointer to **[]string** | List of catalog IDs that failed to update. | [optional] 
**PreviousKeySecretId** | Pointer to **string** | ID of the transaction associated with the update. | [optional] 
**StartedAt** | Pointer to **time.Time** | The datetime when the job started | [optional] 
**EstimatedTimeInMinute** | Pointer to **int32** | Estimated total time for the job to finish | [optional] 

## Methods

### NewDataStoreAzureSubscriptionUpdateCredentialsJobModel

`func NewDataStoreAzureSubscriptionUpdateCredentialsJobModel() *DataStoreAzureSubscriptionUpdateCredentialsJobModel`

NewDataStoreAzureSubscriptionUpdateCredentialsJobModel instantiates a new DataStoreAzureSubscriptionUpdateCredentialsJobModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStoreAzureSubscriptionUpdateCredentialsJobModelWithDefaults

`func NewDataStoreAzureSubscriptionUpdateCredentialsJobModelWithDefaults() *DataStoreAzureSubscriptionUpdateCredentialsJobModel`

NewDataStoreAzureSubscriptionUpdateCredentialsJobModelWithDefaults instantiates a new DataStoreAzureSubscriptionUpdateCredentialsJobModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetState

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) GetState() UpdateAzureSubscriptionCredentialJobState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) GetStateOk() (*UpdateAzureSubscriptionCredentialJobState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) SetState(v UpdateAzureSubscriptionCredentialJobState)`

SetState sets State field to given value.

### HasState

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetFailedCatalogIds

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) GetFailedCatalogIds() []string`

GetFailedCatalogIds returns the FailedCatalogIds field if non-nil, zero value otherwise.

### GetFailedCatalogIdsOk

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) GetFailedCatalogIdsOk() (*[]string, bool)`

GetFailedCatalogIdsOk returns a tuple with the FailedCatalogIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCatalogIds

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) SetFailedCatalogIds(v []string)`

SetFailedCatalogIds sets FailedCatalogIds field to given value.

### HasFailedCatalogIds

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) HasFailedCatalogIds() bool`

HasFailedCatalogIds returns a boolean if a field has been set.

### GetPreviousKeySecretId

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) GetPreviousKeySecretId() string`

GetPreviousKeySecretId returns the PreviousKeySecretId field if non-nil, zero value otherwise.

### GetPreviousKeySecretIdOk

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) GetPreviousKeySecretIdOk() (*string, bool)`

GetPreviousKeySecretIdOk returns a tuple with the PreviousKeySecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousKeySecretId

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) SetPreviousKeySecretId(v string)`

SetPreviousKeySecretId sets PreviousKeySecretId field to given value.

### HasPreviousKeySecretId

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) HasPreviousKeySecretId() bool`

HasPreviousKeySecretId returns a boolean if a field has been set.

### GetStartedAt

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetEstimatedTimeInMinute

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) GetEstimatedTimeInMinute() int32`

GetEstimatedTimeInMinute returns the EstimatedTimeInMinute field if non-nil, zero value otherwise.

### GetEstimatedTimeInMinuteOk

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) GetEstimatedTimeInMinuteOk() (*int32, bool)`

GetEstimatedTimeInMinuteOk returns a tuple with the EstimatedTimeInMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTimeInMinute

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) SetEstimatedTimeInMinute(v int32)`

SetEstimatedTimeInMinute sets EstimatedTimeInMinute field to given value.

### HasEstimatedTimeInMinute

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) HasEstimatedTimeInMinute() bool`

HasEstimatedTimeInMinute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


