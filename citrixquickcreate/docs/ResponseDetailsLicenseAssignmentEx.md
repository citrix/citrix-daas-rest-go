# ResponseDetailsLicenseAssignmentEx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]CoreUser**](CoreUser.md) | List of users that failed with the specific errorcode | [optional] 
**ErrorCode** | Pointer to **int32** | Once of the error codes listed in the ErrorCode enum | [optional] 
**ErrorDescription** | Pointer to **NullableString** | Description about the error | [optional] 

## Methods

### NewResponseDetailsLicenseAssignmentEx

`func NewResponseDetailsLicenseAssignmentEx() *ResponseDetailsLicenseAssignmentEx`

NewResponseDetailsLicenseAssignmentEx instantiates a new ResponseDetailsLicenseAssignmentEx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDetailsLicenseAssignmentExWithDefaults

`func NewResponseDetailsLicenseAssignmentExWithDefaults() *ResponseDetailsLicenseAssignmentEx`

NewResponseDetailsLicenseAssignmentExWithDefaults instantiates a new ResponseDetailsLicenseAssignmentEx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *ResponseDetailsLicenseAssignmentEx) GetUsers() []CoreUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ResponseDetailsLicenseAssignmentEx) GetUsersOk() (*[]CoreUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ResponseDetailsLicenseAssignmentEx) SetUsers(v []CoreUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ResponseDetailsLicenseAssignmentEx) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *ResponseDetailsLicenseAssignmentEx) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *ResponseDetailsLicenseAssignmentEx) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil
### GetErrorCode

`func (o *ResponseDetailsLicenseAssignmentEx) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ResponseDetailsLicenseAssignmentEx) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ResponseDetailsLicenseAssignmentEx) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ResponseDetailsLicenseAssignmentEx) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *ResponseDetailsLicenseAssignmentEx) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *ResponseDetailsLicenseAssignmentEx) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *ResponseDetailsLicenseAssignmentEx) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *ResponseDetailsLicenseAssignmentEx) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### SetErrorDescriptionNil

`func (o *ResponseDetailsLicenseAssignmentEx) SetErrorDescriptionNil(b bool)`

 SetErrorDescriptionNil sets the value for ErrorDescription to be an explicit nil

### UnsetErrorDescription
`func (o *ResponseDetailsLicenseAssignmentEx) UnsetErrorDescription()`

UnsetErrorDescription ensures that no value is present for ErrorDescription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


