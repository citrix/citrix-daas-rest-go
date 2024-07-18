# ResponseDetailsLicenseAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to **[]string** | List of user IDs that failed with the specific errorcode | [optional] [readonly] 
**ErrorCode** | Pointer to **int32** | Once of the error codes listed in the ErrorCode enum | [optional] 
**ErrorDescription** | Pointer to **NullableString** | Description about the error | [optional] 

## Methods

### NewResponseDetailsLicenseAssignment

`func NewResponseDetailsLicenseAssignment() *ResponseDetailsLicenseAssignment`

NewResponseDetailsLicenseAssignment instantiates a new ResponseDetailsLicenseAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDetailsLicenseAssignmentWithDefaults

`func NewResponseDetailsLicenseAssignmentWithDefaults() *ResponseDetailsLicenseAssignment`

NewResponseDetailsLicenseAssignmentWithDefaults instantiates a new ResponseDetailsLicenseAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *ResponseDetailsLicenseAssignment) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ResponseDetailsLicenseAssignment) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ResponseDetailsLicenseAssignment) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ResponseDetailsLicenseAssignment) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *ResponseDetailsLicenseAssignment) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *ResponseDetailsLicenseAssignment) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil
### GetErrorCode

`func (o *ResponseDetailsLicenseAssignment) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ResponseDetailsLicenseAssignment) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ResponseDetailsLicenseAssignment) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ResponseDetailsLicenseAssignment) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *ResponseDetailsLicenseAssignment) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *ResponseDetailsLicenseAssignment) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *ResponseDetailsLicenseAssignment) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *ResponseDetailsLicenseAssignment) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### SetErrorDescriptionNil

`func (o *ResponseDetailsLicenseAssignment) SetErrorDescriptionNil(b bool)`

 SetErrorDescriptionNil sets the value for ErrorDescription to be an explicit nil

### UnsetErrorDescription
`func (o *ResponseDetailsLicenseAssignment) UnsetErrorDescription()`

UnsetErrorDescription ensures that no value is present for ErrorDescription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


