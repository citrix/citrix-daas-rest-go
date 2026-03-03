# DataStoreAzureSubscriptionUpdateCredentialsJobModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | Pointer to **NullableString** | ID of the transaction associated with the update. | [optional] 
**State** | Pointer to [**UpdateAzureSubscriptionCredentialJobState**](UpdateAzureSubscriptionCredentialJobState.md) | State of the update job. | [optional] 
**FailedCatalogIds** | Pointer to **[]string** | List of catalog IDs that failed to update. | [optional] 
**PreviousKeySecretId** | Pointer to **NullableString** | ID of the transaction associated with the update. | [optional] 
**StartedAt** | Pointer to **NullableTime** | The datetime when the job started | [optional] 
**EstimatedTimeInMinute** | Pointer to **NullableInt32** | Estimated total time for the job to finish | [optional] 

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

### SetTransactionIdNil

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
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

### SetFailedCatalogIdsNil

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) SetFailedCatalogIdsNil(b bool)`

 SetFailedCatalogIdsNil sets the value for FailedCatalogIds to be an explicit nil

### UnsetFailedCatalogIds
`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) UnsetFailedCatalogIds()`

UnsetFailedCatalogIds ensures that no value is present for FailedCatalogIds, not even an explicit nil
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

### SetPreviousKeySecretIdNil

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) SetPreviousKeySecretIdNil(b bool)`

 SetPreviousKeySecretIdNil sets the value for PreviousKeySecretId to be an explicit nil

### UnsetPreviousKeySecretId
`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) UnsetPreviousKeySecretId()`

UnsetPreviousKeySecretId ensures that no value is present for PreviousKeySecretId, not even an explicit nil
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

### SetStartedAtNil

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
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

### SetEstimatedTimeInMinuteNil

`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) SetEstimatedTimeInMinuteNil(b bool)`

 SetEstimatedTimeInMinuteNil sets the value for EstimatedTimeInMinute to be an explicit nil

### UnsetEstimatedTimeInMinute
`func (o *DataStoreAzureSubscriptionUpdateCredentialsJobModel) UnsetEstimatedTimeInMinute()`

UnsetEstimatedTimeInMinute ensures that no value is present for EstimatedTimeInMinute, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


