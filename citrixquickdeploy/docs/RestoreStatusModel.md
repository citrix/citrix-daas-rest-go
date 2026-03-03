# RestoreStatusModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**RestoreState**](RestoreState.md) |  | [optional] 
**SubState** | Pointer to [**NullableRestoreSubState**](RestoreSubState.md) |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 
**TransactionId** | Pointer to **NullableString** |  | [optional] 
**LastRestoredTime** | Pointer to **NullableTime** |  | [optional] 
**LastRestoredSnapshotName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRestoreStatusModel

`func NewRestoreStatusModel() *RestoreStatusModel`

NewRestoreStatusModel instantiates a new RestoreStatusModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreStatusModelWithDefaults

`func NewRestoreStatusModelWithDefaults() *RestoreStatusModel`

NewRestoreStatusModelWithDefaults instantiates a new RestoreStatusModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *RestoreStatusModel) GetState() RestoreState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RestoreStatusModel) GetStateOk() (*RestoreState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RestoreStatusModel) SetState(v RestoreState)`

SetState sets State field to given value.

### HasState

`func (o *RestoreStatusModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubState

`func (o *RestoreStatusModel) GetSubState() RestoreSubState`

GetSubState returns the SubState field if non-nil, zero value otherwise.

### GetSubStateOk

`func (o *RestoreStatusModel) GetSubStateOk() (*RestoreSubState, bool)`

GetSubStateOk returns a tuple with the SubState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubState

`func (o *RestoreStatusModel) SetSubState(v RestoreSubState)`

SetSubState sets SubState field to given value.

### HasSubState

`func (o *RestoreStatusModel) HasSubState() bool`

HasSubState returns a boolean if a field has been set.

### SetSubStateNil

`func (o *RestoreStatusModel) SetSubStateNil(b bool)`

 SetSubStateNil sets the value for SubState to be an explicit nil

### UnsetSubState
`func (o *RestoreStatusModel) UnsetSubState()`

UnsetSubState ensures that no value is present for SubState, not even an explicit nil
### GetError

`func (o *RestoreStatusModel) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RestoreStatusModel) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RestoreStatusModel) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *RestoreStatusModel) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *RestoreStatusModel) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *RestoreStatusModel) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetTransactionId

`func (o *RestoreStatusModel) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *RestoreStatusModel) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *RestoreStatusModel) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *RestoreStatusModel) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *RestoreStatusModel) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *RestoreStatusModel) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetLastRestoredTime

`func (o *RestoreStatusModel) GetLastRestoredTime() time.Time`

GetLastRestoredTime returns the LastRestoredTime field if non-nil, zero value otherwise.

### GetLastRestoredTimeOk

`func (o *RestoreStatusModel) GetLastRestoredTimeOk() (*time.Time, bool)`

GetLastRestoredTimeOk returns a tuple with the LastRestoredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRestoredTime

`func (o *RestoreStatusModel) SetLastRestoredTime(v time.Time)`

SetLastRestoredTime sets LastRestoredTime field to given value.

### HasLastRestoredTime

`func (o *RestoreStatusModel) HasLastRestoredTime() bool`

HasLastRestoredTime returns a boolean if a field has been set.

### SetLastRestoredTimeNil

`func (o *RestoreStatusModel) SetLastRestoredTimeNil(b bool)`

 SetLastRestoredTimeNil sets the value for LastRestoredTime to be an explicit nil

### UnsetLastRestoredTime
`func (o *RestoreStatusModel) UnsetLastRestoredTime()`

UnsetLastRestoredTime ensures that no value is present for LastRestoredTime, not even an explicit nil
### GetLastRestoredSnapshotName

`func (o *RestoreStatusModel) GetLastRestoredSnapshotName() string`

GetLastRestoredSnapshotName returns the LastRestoredSnapshotName field if non-nil, zero value otherwise.

### GetLastRestoredSnapshotNameOk

`func (o *RestoreStatusModel) GetLastRestoredSnapshotNameOk() (*string, bool)`

GetLastRestoredSnapshotNameOk returns a tuple with the LastRestoredSnapshotName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRestoredSnapshotName

`func (o *RestoreStatusModel) SetLastRestoredSnapshotName(v string)`

SetLastRestoredSnapshotName sets LastRestoredSnapshotName field to given value.

### HasLastRestoredSnapshotName

`func (o *RestoreStatusModel) HasLastRestoredSnapshotName() bool`

HasLastRestoredSnapshotName returns a boolean if a field has been set.

### SetLastRestoredSnapshotNameNil

`func (o *RestoreStatusModel) SetLastRestoredSnapshotNameNil(b bool)`

 SetLastRestoredSnapshotNameNil sets the value for LastRestoredSnapshotName to be an explicit nil

### UnsetLastRestoredSnapshotName
`func (o *RestoreStatusModel) UnsetLastRestoredSnapshotName()`

UnsetLastRestoredSnapshotName ensures that no value is present for LastRestoredSnapshotName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


