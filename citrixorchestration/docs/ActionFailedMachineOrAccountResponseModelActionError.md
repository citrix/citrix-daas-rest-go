# ActionFailedMachineOrAccountResponseModelActionError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | The error message. | [optional] 
**StackTrace** | Pointer to **string** | The stack trace. | [optional] 
**ErrorData** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The error data. | [optional] 
**InnerError** | Pointer to [**ActionErrorInnerError**](ActionErrorInnerError.md) |  | [optional] 

## Methods

### NewActionFailedMachineOrAccountResponseModelActionError

`func NewActionFailedMachineOrAccountResponseModelActionError() *ActionFailedMachineOrAccountResponseModelActionError`

NewActionFailedMachineOrAccountResponseModelActionError instantiates a new ActionFailedMachineOrAccountResponseModelActionError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionFailedMachineOrAccountResponseModelActionErrorWithDefaults

`func NewActionFailedMachineOrAccountResponseModelActionErrorWithDefaults() *ActionFailedMachineOrAccountResponseModelActionError`

NewActionFailedMachineOrAccountResponseModelActionErrorWithDefaults instantiates a new ActionFailedMachineOrAccountResponseModelActionError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ActionFailedMachineOrAccountResponseModelActionError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ActionFailedMachineOrAccountResponseModelActionError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ActionFailedMachineOrAccountResponseModelActionError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ActionFailedMachineOrAccountResponseModelActionError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStackTrace

`func (o *ActionFailedMachineOrAccountResponseModelActionError) GetStackTrace() string`

GetStackTrace returns the StackTrace field if non-nil, zero value otherwise.

### GetStackTraceOk

`func (o *ActionFailedMachineOrAccountResponseModelActionError) GetStackTraceOk() (*string, bool)`

GetStackTraceOk returns a tuple with the StackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTrace

`func (o *ActionFailedMachineOrAccountResponseModelActionError) SetStackTrace(v string)`

SetStackTrace sets StackTrace field to given value.

### HasStackTrace

`func (o *ActionFailedMachineOrAccountResponseModelActionError) HasStackTrace() bool`

HasStackTrace returns a boolean if a field has been set.

### GetErrorData

`func (o *ActionFailedMachineOrAccountResponseModelActionError) GetErrorData() []NameValueStringPairModel`

GetErrorData returns the ErrorData field if non-nil, zero value otherwise.

### GetErrorDataOk

`func (o *ActionFailedMachineOrAccountResponseModelActionError) GetErrorDataOk() (*[]NameValueStringPairModel, bool)`

GetErrorDataOk returns a tuple with the ErrorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorData

`func (o *ActionFailedMachineOrAccountResponseModelActionError) SetErrorData(v []NameValueStringPairModel)`

SetErrorData sets ErrorData field to given value.

### HasErrorData

`func (o *ActionFailedMachineOrAccountResponseModelActionError) HasErrorData() bool`

HasErrorData returns a boolean if a field has been set.

### GetInnerError

`func (o *ActionFailedMachineOrAccountResponseModelActionError) GetInnerError() ActionErrorInnerError`

GetInnerError returns the InnerError field if non-nil, zero value otherwise.

### GetInnerErrorOk

`func (o *ActionFailedMachineOrAccountResponseModelActionError) GetInnerErrorOk() (*ActionErrorInnerError, bool)`

GetInnerErrorOk returns a tuple with the InnerError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerError

`func (o *ActionFailedMachineOrAccountResponseModelActionError) SetInnerError(v ActionErrorInnerError)`

SetInnerError sets InnerError field to given value.

### HasInnerError

`func (o *ActionFailedMachineOrAccountResponseModelActionError) HasInnerError() bool`

HasInnerError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


