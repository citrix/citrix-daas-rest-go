# RegisterAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | [**AccountType**](AccountType.md) |  | 
**TaskType** | Pointer to [**TaskType**](TaskType.md) |  | [optional] 
**TaskId** | Pointer to **NullableString** | Task Id | [optional] 

## Methods

### NewRegisterAccount

`func NewRegisterAccount(accountType AccountType, ) *RegisterAccount`

NewRegisterAccount instantiates a new RegisterAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterAccountWithDefaults

`func NewRegisterAccountWithDefaults() *RegisterAccount`

NewRegisterAccountWithDefaults instantiates a new RegisterAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *RegisterAccount) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *RegisterAccount) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *RegisterAccount) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetTaskType

`func (o *RegisterAccount) GetTaskType() TaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *RegisterAccount) GetTaskTypeOk() (*TaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *RegisterAccount) SetTaskType(v TaskType)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *RegisterAccount) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetTaskId

`func (o *RegisterAccount) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *RegisterAccount) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *RegisterAccount) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *RegisterAccount) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *RegisterAccount) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *RegisterAccount) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


