# ActionResponseModelTerminatingError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | The error message. | [optional] 
**StackTrace** | Pointer to **string** | The stack trace. | [optional] 
**ErrorData** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The error data. | [optional] 
**InnerError** | Pointer to [**ActionErrorInnerError**](ActionErrorInnerError.md) |  | [optional] 

## Methods

### NewActionResponseModelTerminatingError

`func NewActionResponseModelTerminatingError() *ActionResponseModelTerminatingError`

NewActionResponseModelTerminatingError instantiates a new ActionResponseModelTerminatingError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionResponseModelTerminatingErrorWithDefaults

`func NewActionResponseModelTerminatingErrorWithDefaults() *ActionResponseModelTerminatingError`

NewActionResponseModelTerminatingErrorWithDefaults instantiates a new ActionResponseModelTerminatingError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ActionResponseModelTerminatingError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ActionResponseModelTerminatingError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ActionResponseModelTerminatingError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ActionResponseModelTerminatingError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStackTrace

`func (o *ActionResponseModelTerminatingError) GetStackTrace() string`

GetStackTrace returns the StackTrace field if non-nil, zero value otherwise.

### GetStackTraceOk

`func (o *ActionResponseModelTerminatingError) GetStackTraceOk() (*string, bool)`

GetStackTraceOk returns a tuple with the StackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTrace

`func (o *ActionResponseModelTerminatingError) SetStackTrace(v string)`

SetStackTrace sets StackTrace field to given value.

### HasStackTrace

`func (o *ActionResponseModelTerminatingError) HasStackTrace() bool`

HasStackTrace returns a boolean if a field has been set.

### GetErrorData

`func (o *ActionResponseModelTerminatingError) GetErrorData() []NameValueStringPairModel`

GetErrorData returns the ErrorData field if non-nil, zero value otherwise.

### GetErrorDataOk

`func (o *ActionResponseModelTerminatingError) GetErrorDataOk() (*[]NameValueStringPairModel, bool)`

GetErrorDataOk returns a tuple with the ErrorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorData

`func (o *ActionResponseModelTerminatingError) SetErrorData(v []NameValueStringPairModel)`

SetErrorData sets ErrorData field to given value.

### HasErrorData

`func (o *ActionResponseModelTerminatingError) HasErrorData() bool`

HasErrorData returns a boolean if a field has been set.

### GetInnerError

`func (o *ActionResponseModelTerminatingError) GetInnerError() ActionErrorInnerError`

GetInnerError returns the InnerError field if non-nil, zero value otherwise.

### GetInnerErrorOk

`func (o *ActionResponseModelTerminatingError) GetInnerErrorOk() (*ActionErrorInnerError, bool)`

GetInnerErrorOk returns a tuple with the InnerError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerError

`func (o *ActionResponseModelTerminatingError) SetInnerError(v ActionErrorInnerError)`

SetInnerError sets InnerError field to given value.

### HasInnerError

`func (o *ActionResponseModelTerminatingError) HasInnerError() bool`

HasInnerError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


