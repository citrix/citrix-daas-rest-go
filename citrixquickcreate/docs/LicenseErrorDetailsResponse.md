# LicenseErrorDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseStatus** | Pointer to [**[]ResponseDetailsLicenseAssignment**](ResponseDetailsLicenseAssignment.md) |  | [optional] [readonly] 

## Methods

### NewLicenseErrorDetailsResponse

`func NewLicenseErrorDetailsResponse() *LicenseErrorDetailsResponse`

NewLicenseErrorDetailsResponse instantiates a new LicenseErrorDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseErrorDetailsResponseWithDefaults

`func NewLicenseErrorDetailsResponseWithDefaults() *LicenseErrorDetailsResponse`

NewLicenseErrorDetailsResponseWithDefaults instantiates a new LicenseErrorDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseStatus

`func (o *LicenseErrorDetailsResponse) GetLicenseStatus() []ResponseDetailsLicenseAssignment`

GetLicenseStatus returns the LicenseStatus field if non-nil, zero value otherwise.

### GetLicenseStatusOk

`func (o *LicenseErrorDetailsResponse) GetLicenseStatusOk() (*[]ResponseDetailsLicenseAssignment, bool)`

GetLicenseStatusOk returns a tuple with the LicenseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseStatus

`func (o *LicenseErrorDetailsResponse) SetLicenseStatus(v []ResponseDetailsLicenseAssignment)`

SetLicenseStatus sets LicenseStatus field to given value.

### HasLicenseStatus

`func (o *LicenseErrorDetailsResponse) HasLicenseStatus() bool`

HasLicenseStatus returns a boolean if a field has been set.

### SetLicenseStatusNil

`func (o *LicenseErrorDetailsResponse) SetLicenseStatusNil(b bool)`

 SetLicenseStatusNil sets the value for LicenseStatus to be an explicit nil

### UnsetLicenseStatus
`func (o *LicenseErrorDetailsResponse) UnsetLicenseStatus()`

UnsetLicenseStatus ensures that no value is present for LicenseStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


