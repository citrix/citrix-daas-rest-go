# PowerStartResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to [**PowerResponseCode**](PowerResponseCode.md) |  | [optional] 
**ErrorDescription** | Pointer to **NullableString** | Description about the error | [optional] [readonly] 

## Methods

### NewPowerStartResponse

`func NewPowerStartResponse() *PowerStartResponse`

NewPowerStartResponse instantiates a new PowerStartResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerStartResponseWithDefaults

`func NewPowerStartResponseWithDefaults() *PowerStartResponse`

NewPowerStartResponseWithDefaults instantiates a new PowerStartResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PowerStartResponse) GetCode() PowerResponseCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PowerStartResponse) GetCodeOk() (*PowerResponseCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PowerStartResponse) SetCode(v PowerResponseCode)`

SetCode sets Code field to given value.

### HasCode

`func (o *PowerStartResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *PowerStartResponse) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *PowerStartResponse) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *PowerStartResponse) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *PowerStartResponse) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### SetErrorDescriptionNil

`func (o *PowerStartResponse) SetErrorDescriptionNil(b bool)`

 SetErrorDescriptionNil sets the value for ErrorDescription to be an explicit nil

### UnsetErrorDescription
`func (o *PowerStartResponse) UnsetErrorDescription()`

UnsetErrorDescription ensures that no value is present for ErrorDescription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


