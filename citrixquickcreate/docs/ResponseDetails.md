# ResponseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** | Once of the error codes listed in the ErrorCode enum | [optional] 
**ErrorDescription** | Pointer to **NullableString** | Description about the error | [optional] 

## Methods

### NewResponseDetails

`func NewResponseDetails() *ResponseDetails`

NewResponseDetails instantiates a new ResponseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDetailsWithDefaults

`func NewResponseDetailsWithDefaults() *ResponseDetails`

NewResponseDetailsWithDefaults instantiates a new ResponseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *ResponseDetails) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ResponseDetails) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ResponseDetails) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ResponseDetails) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *ResponseDetails) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *ResponseDetails) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *ResponseDetails) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *ResponseDetails) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### SetErrorDescriptionNil

`func (o *ResponseDetails) SetErrorDescriptionNil(b bool)`

 SetErrorDescriptionNil sets the value for ErrorDescription to be an explicit nil

### UnsetErrorDescription
`func (o *ResponseDetails) UnsetErrorDescription()`

UnsetErrorDescription ensures that no value is present for ErrorDescription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


