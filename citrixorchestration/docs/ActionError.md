# ActionError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** | The error message. | [optional] 
**StackTrace** | Pointer to **NullableString** | The stack trace. | [optional] 
**ErrorData** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The error data. | [optional] 
**InnerError** | Pointer to [**ActionError**](ActionError.md) |  | [optional] 

## Methods

### NewActionError

`func NewActionError() *ActionError`

NewActionError instantiates a new ActionError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionErrorWithDefaults

`func NewActionErrorWithDefaults() *ActionError`

NewActionErrorWithDefaults instantiates a new ActionError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ActionError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ActionError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ActionError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ActionError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ActionError) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ActionError) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetStackTrace

`func (o *ActionError) GetStackTrace() string`

GetStackTrace returns the StackTrace field if non-nil, zero value otherwise.

### GetStackTraceOk

`func (o *ActionError) GetStackTraceOk() (*string, bool)`

GetStackTraceOk returns a tuple with the StackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTrace

`func (o *ActionError) SetStackTrace(v string)`

SetStackTrace sets StackTrace field to given value.

### HasStackTrace

`func (o *ActionError) HasStackTrace() bool`

HasStackTrace returns a boolean if a field has been set.

### SetStackTraceNil

`func (o *ActionError) SetStackTraceNil(b bool)`

 SetStackTraceNil sets the value for StackTrace to be an explicit nil

### UnsetStackTrace
`func (o *ActionError) UnsetStackTrace()`

UnsetStackTrace ensures that no value is present for StackTrace, not even an explicit nil
### GetErrorData

`func (o *ActionError) GetErrorData() []NameValueStringPairModel`

GetErrorData returns the ErrorData field if non-nil, zero value otherwise.

### GetErrorDataOk

`func (o *ActionError) GetErrorDataOk() (*[]NameValueStringPairModel, bool)`

GetErrorDataOk returns a tuple with the ErrorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorData

`func (o *ActionError) SetErrorData(v []NameValueStringPairModel)`

SetErrorData sets ErrorData field to given value.

### HasErrorData

`func (o *ActionError) HasErrorData() bool`

HasErrorData returns a boolean if a field has been set.

### SetErrorDataNil

`func (o *ActionError) SetErrorDataNil(b bool)`

 SetErrorDataNil sets the value for ErrorData to be an explicit nil

### UnsetErrorData
`func (o *ActionError) UnsetErrorData()`

UnsetErrorData ensures that no value is present for ErrorData, not even an explicit nil
### GetInnerError

`func (o *ActionError) GetInnerError() ActionError`

GetInnerError returns the InnerError field if non-nil, zero value otherwise.

### GetInnerErrorOk

`func (o *ActionError) GetInnerErrorOk() (*ActionError, bool)`

GetInnerErrorOk returns a tuple with the InnerError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerError

`func (o *ActionError) SetInnerError(v ActionError)`

SetInnerError sets InnerError field to given value.

### HasInnerError

`func (o *ActionError) HasInnerError() bool`

HasInnerError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


