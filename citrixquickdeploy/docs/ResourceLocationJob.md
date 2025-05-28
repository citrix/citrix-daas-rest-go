# ResourceLocationJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobType** | Pointer to [**ResourceLocationJobType**](ResourceLocationJobType.md) | The type of job | [optional] 
**State** | Pointer to [**ResourceLocationJobState**](ResourceLocationJobState.md) | The state of the job | [optional] 
**Status** | Pointer to [**ResourceLocationJobStatus**](ResourceLocationJobStatus.md) | Status of the job | [optional] 
**StatusMessage** | Pointer to **string** | Error associated with the job failure | [optional] 
**EndedAt** | Pointer to **time.Time** | Time at which the job was completed at | [optional] 
**TransactionId** | Pointer to **string** | ID of the transaction performing the job | [optional] 
**Quantity** | Pointer to **int32** | Quantity of items requested in the job | [optional] 
**QuantityFailed** | Pointer to **int32** | Quantity of items that failed the job | [optional] 
**StartedAt** | Pointer to **time.Time** | The datetime when the job started | [optional] 
**EstimatedTimeInMinute** | Pointer to **int32** | Estimated total time for the job to finish | [optional] 

## Methods

### NewResourceLocationJob

`func NewResourceLocationJob() *ResourceLocationJob`

NewResourceLocationJob instantiates a new ResourceLocationJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceLocationJobWithDefaults

`func NewResourceLocationJobWithDefaults() *ResourceLocationJob`

NewResourceLocationJobWithDefaults instantiates a new ResourceLocationJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobType

`func (o *ResourceLocationJob) GetJobType() ResourceLocationJobType`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *ResourceLocationJob) GetJobTypeOk() (*ResourceLocationJobType, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *ResourceLocationJob) SetJobType(v ResourceLocationJobType)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *ResourceLocationJob) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetState

`func (o *ResourceLocationJob) GetState() ResourceLocationJobState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResourceLocationJob) GetStateOk() (*ResourceLocationJobState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResourceLocationJob) SetState(v ResourceLocationJobState)`

SetState sets State field to given value.

### HasState

`func (o *ResourceLocationJob) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *ResourceLocationJob) GetStatus() ResourceLocationJobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceLocationJob) GetStatusOk() (*ResourceLocationJobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceLocationJob) SetStatus(v ResourceLocationJobStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResourceLocationJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ResourceLocationJob) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ResourceLocationJob) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ResourceLocationJob) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ResourceLocationJob) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetEndedAt

`func (o *ResourceLocationJob) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *ResourceLocationJob) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *ResourceLocationJob) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *ResourceLocationJob) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetTransactionId

`func (o *ResourceLocationJob) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ResourceLocationJob) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ResourceLocationJob) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ResourceLocationJob) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetQuantity

`func (o *ResourceLocationJob) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ResourceLocationJob) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ResourceLocationJob) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ResourceLocationJob) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetQuantityFailed

`func (o *ResourceLocationJob) GetQuantityFailed() int32`

GetQuantityFailed returns the QuantityFailed field if non-nil, zero value otherwise.

### GetQuantityFailedOk

`func (o *ResourceLocationJob) GetQuantityFailedOk() (*int32, bool)`

GetQuantityFailedOk returns a tuple with the QuantityFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityFailed

`func (o *ResourceLocationJob) SetQuantityFailed(v int32)`

SetQuantityFailed sets QuantityFailed field to given value.

### HasQuantityFailed

`func (o *ResourceLocationJob) HasQuantityFailed() bool`

HasQuantityFailed returns a boolean if a field has been set.

### GetStartedAt

`func (o *ResourceLocationJob) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ResourceLocationJob) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ResourceLocationJob) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ResourceLocationJob) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetEstimatedTimeInMinute

`func (o *ResourceLocationJob) GetEstimatedTimeInMinute() int32`

GetEstimatedTimeInMinute returns the EstimatedTimeInMinute field if non-nil, zero value otherwise.

### GetEstimatedTimeInMinuteOk

`func (o *ResourceLocationJob) GetEstimatedTimeInMinuteOk() (*int32, bool)`

GetEstimatedTimeInMinuteOk returns a tuple with the EstimatedTimeInMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTimeInMinute

`func (o *ResourceLocationJob) SetEstimatedTimeInMinute(v int32)`

SetEstimatedTimeInMinute sets EstimatedTimeInMinute field to given value.

### HasEstimatedTimeInMinute

`func (o *ResourceLocationJob) HasEstimatedTimeInMinute() bool`

HasEstimatedTimeInMinute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


