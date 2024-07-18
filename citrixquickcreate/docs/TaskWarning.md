# TaskWarning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WarningId** | Pointer to **NullableString** | Warning Id | [optional] 
**Message** | Pointer to **NullableString** | Warning Message | [optional] 

## Methods

### NewTaskWarning

`func NewTaskWarning() *TaskWarning`

NewTaskWarning instantiates a new TaskWarning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWarningWithDefaults

`func NewTaskWarningWithDefaults() *TaskWarning`

NewTaskWarningWithDefaults instantiates a new TaskWarning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarningId

`func (o *TaskWarning) GetWarningId() string`

GetWarningId returns the WarningId field if non-nil, zero value otherwise.

### GetWarningIdOk

`func (o *TaskWarning) GetWarningIdOk() (*string, bool)`

GetWarningIdOk returns a tuple with the WarningId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningId

`func (o *TaskWarning) SetWarningId(v string)`

SetWarningId sets WarningId field to given value.

### HasWarningId

`func (o *TaskWarning) HasWarningId() bool`

HasWarningId returns a boolean if a field has been set.

### SetWarningIdNil

`func (o *TaskWarning) SetWarningIdNil(b bool)`

 SetWarningIdNil sets the value for WarningId to be an explicit nil

### UnsetWarningId
`func (o *TaskWarning) UnsetWarningId()`

UnsetWarningId ensures that no value is present for WarningId, not even an explicit nil
### GetMessage

`func (o *TaskWarning) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TaskWarning) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TaskWarning) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TaskWarning) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *TaskWarning) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *TaskWarning) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


