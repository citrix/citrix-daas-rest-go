# AccountTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | [**NullableAccountTaskOperationType**](AccountTaskOperationType.md) |  | 
**AccountId** | Pointer to **NullableString** | Account Id this task is working on | [optional] 

## Methods

### NewAccountTask

`func NewAccountTask(operation NullableAccountTaskOperationType, ) *AccountTask`

NewAccountTask instantiates a new AccountTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountTaskWithDefaults

`func NewAccountTaskWithDefaults() *AccountTask`

NewAccountTaskWithDefaults instantiates a new AccountTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *AccountTask) GetOperation() AccountTaskOperationType`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *AccountTask) GetOperationOk() (*AccountTaskOperationType, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *AccountTask) SetOperation(v AccountTaskOperationType)`

SetOperation sets Operation field to given value.


### SetOperationNil

`func (o *AccountTask) SetOperationNil(b bool)`

 SetOperationNil sets the value for Operation to be an explicit nil

### UnsetOperation
`func (o *AccountTask) UnsetOperation()`

UnsetOperation ensures that no value is present for Operation, not even an explicit nil
### GetAccountId

`func (o *AccountTask) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountTask) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountTask) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AccountTask) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *AccountTask) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *AccountTask) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


