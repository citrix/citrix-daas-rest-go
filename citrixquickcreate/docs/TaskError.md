# TaskError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorId** | Pointer to **NullableString** | Error Id | [optional] 
**Message** | Pointer to **NullableString** | Error Message | [optional] 
**Exception** | Pointer to **interface{}** | Error Message | [optional] 

## Methods

### NewTaskError

`func NewTaskError() *TaskError`

NewTaskError instantiates a new TaskError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskErrorWithDefaults

`func NewTaskErrorWithDefaults() *TaskError`

NewTaskErrorWithDefaults instantiates a new TaskError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorId

`func (o *TaskError) GetErrorId() string`

GetErrorId returns the ErrorId field if non-nil, zero value otherwise.

### GetErrorIdOk

`func (o *TaskError) GetErrorIdOk() (*string, bool)`

GetErrorIdOk returns a tuple with the ErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorId

`func (o *TaskError) SetErrorId(v string)`

SetErrorId sets ErrorId field to given value.

### HasErrorId

`func (o *TaskError) HasErrorId() bool`

HasErrorId returns a boolean if a field has been set.

### SetErrorIdNil

`func (o *TaskError) SetErrorIdNil(b bool)`

 SetErrorIdNil sets the value for ErrorId to be an explicit nil

### UnsetErrorId
`func (o *TaskError) UnsetErrorId()`

UnsetErrorId ensures that no value is present for ErrorId, not even an explicit nil
### GetMessage

`func (o *TaskError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TaskError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TaskError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TaskError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *TaskError) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *TaskError) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetException

`func (o *TaskError) GetException() interface{}`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *TaskError) GetExceptionOk() (*interface{}, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *TaskError) SetException(v interface{})`

SetException sets Exception field to given value.

### HasException

`func (o *TaskError) HasException() bool`

HasException returns a boolean if a field has been set.

### SetExceptionNil

`func (o *TaskError) SetExceptionNil(b bool)`

 SetExceptionNil sets the value for Exception to be an explicit nil

### UnsetException
`func (o *TaskError) UnsetException()`

UnsetException ensures that no value is present for Exception, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


