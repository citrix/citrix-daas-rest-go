# GroupLicenseErrorDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseStatus** | Pointer to [**[]ResponseDetailsGroupLicenseAssignment**](ResponseDetailsGroupLicenseAssignment.md) |  | [optional] [readonly] 

## Methods

### NewGroupLicenseErrorDetailsResponse

`func NewGroupLicenseErrorDetailsResponse() *GroupLicenseErrorDetailsResponse`

NewGroupLicenseErrorDetailsResponse instantiates a new GroupLicenseErrorDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupLicenseErrorDetailsResponseWithDefaults

`func NewGroupLicenseErrorDetailsResponseWithDefaults() *GroupLicenseErrorDetailsResponse`

NewGroupLicenseErrorDetailsResponseWithDefaults instantiates a new GroupLicenseErrorDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseStatus

`func (o *GroupLicenseErrorDetailsResponse) GetLicenseStatus() []ResponseDetailsGroupLicenseAssignment`

GetLicenseStatus returns the LicenseStatus field if non-nil, zero value otherwise.

### GetLicenseStatusOk

`func (o *GroupLicenseErrorDetailsResponse) GetLicenseStatusOk() (*[]ResponseDetailsGroupLicenseAssignment, bool)`

GetLicenseStatusOk returns a tuple with the LicenseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseStatus

`func (o *GroupLicenseErrorDetailsResponse) SetLicenseStatus(v []ResponseDetailsGroupLicenseAssignment)`

SetLicenseStatus sets LicenseStatus field to given value.

### HasLicenseStatus

`func (o *GroupLicenseErrorDetailsResponse) HasLicenseStatus() bool`

HasLicenseStatus returns a boolean if a field has been set.

### SetLicenseStatusNil

`func (o *GroupLicenseErrorDetailsResponse) SetLicenseStatusNil(b bool)`

 SetLicenseStatusNil sets the value for LicenseStatus to be an explicit nil

### UnsetLicenseStatus
`func (o *GroupLicenseErrorDetailsResponse) UnsetLicenseStatus()`

UnsetLicenseStatus ensures that no value is present for LicenseStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


