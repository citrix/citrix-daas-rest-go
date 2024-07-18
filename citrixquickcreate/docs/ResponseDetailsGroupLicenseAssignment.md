# ResponseDetailsGroupLicenseAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to **[]string** |  | [optional] 
**ErrorCode** | Pointer to **int32** | Once of the error codes listed in the ErrorCode enum | [optional] 
**ErrorDescription** | Pointer to **NullableString** | Description about the error | [optional] 

## Methods

### NewResponseDetailsGroupLicenseAssignment

`func NewResponseDetailsGroupLicenseAssignment() *ResponseDetailsGroupLicenseAssignment`

NewResponseDetailsGroupLicenseAssignment instantiates a new ResponseDetailsGroupLicenseAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDetailsGroupLicenseAssignmentWithDefaults

`func NewResponseDetailsGroupLicenseAssignmentWithDefaults() *ResponseDetailsGroupLicenseAssignment`

NewResponseDetailsGroupLicenseAssignmentWithDefaults instantiates a new ResponseDetailsGroupLicenseAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *ResponseDetailsGroupLicenseAssignment) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ResponseDetailsGroupLicenseAssignment) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ResponseDetailsGroupLicenseAssignment) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ResponseDetailsGroupLicenseAssignment) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *ResponseDetailsGroupLicenseAssignment) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *ResponseDetailsGroupLicenseAssignment) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetErrorCode

`func (o *ResponseDetailsGroupLicenseAssignment) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ResponseDetailsGroupLicenseAssignment) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ResponseDetailsGroupLicenseAssignment) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ResponseDetailsGroupLicenseAssignment) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *ResponseDetailsGroupLicenseAssignment) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *ResponseDetailsGroupLicenseAssignment) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *ResponseDetailsGroupLicenseAssignment) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *ResponseDetailsGroupLicenseAssignment) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### SetErrorDescriptionNil

`func (o *ResponseDetailsGroupLicenseAssignment) SetErrorDescriptionNil(b bool)`

 SetErrorDescriptionNil sets the value for ErrorDescription to be an explicit nil

### UnsetErrorDescription
`func (o *ResponseDetailsGroupLicenseAssignment) UnsetErrorDescription()`

UnsetErrorDescription ensures that no value is present for ErrorDescription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


