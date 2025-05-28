# RestoreStatusModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**RestoreState**](RestoreState.md) |  | [optional] 
**SubState** | Pointer to [**RestoreSubState**](RestoreSubState.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**LastRestoredTime** | Pointer to **time.Time** |  | [optional] 
**LastRestoredSnapshotName** | Pointer to **string** |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


