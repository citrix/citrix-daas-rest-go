# GroupUserErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserErrors** | Pointer to [**[]ResponseDetailsLicenseAssignmentEx**](ResponseDetailsLicenseAssignmentEx.md) | User errors | [optional] [readonly] 

## Methods

### NewGroupUserErrorResponse

`func NewGroupUserErrorResponse() *GroupUserErrorResponse`

NewGroupUserErrorResponse instantiates a new GroupUserErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUserErrorResponseWithDefaults

`func NewGroupUserErrorResponseWithDefaults() *GroupUserErrorResponse`

NewGroupUserErrorResponseWithDefaults instantiates a new GroupUserErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserErrors

`func (o *GroupUserErrorResponse) GetUserErrors() []ResponseDetailsLicenseAssignmentEx`

GetUserErrors returns the UserErrors field if non-nil, zero value otherwise.

### GetUserErrorsOk

`func (o *GroupUserErrorResponse) GetUserErrorsOk() (*[]ResponseDetailsLicenseAssignmentEx, bool)`

GetUserErrorsOk returns a tuple with the UserErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserErrors

`func (o *GroupUserErrorResponse) SetUserErrors(v []ResponseDetailsLicenseAssignmentEx)`

SetUserErrors sets UserErrors field to given value.

### HasUserErrors

`func (o *GroupUserErrorResponse) HasUserErrors() bool`

HasUserErrors returns a boolean if a field has been set.

### SetUserErrorsNil

`func (o *GroupUserErrorResponse) SetUserErrorsNil(b bool)`

 SetUserErrorsNil sets the value for UserErrors to be an explicit nil

### UnsetUserErrors
`func (o *GroupUserErrorResponse) UnsetUserErrors()`

UnsetUserErrors ensures that no value is present for UserErrors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


