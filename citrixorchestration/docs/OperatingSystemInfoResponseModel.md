# OperatingSystemInfoResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperatingSystemType** | Pointer to [**OperatingSystemType**](OperatingSystemType.md) |  | [optional] 
**MajorVersion** | Pointer to **int32** | OS major version. | [optional] 
**MinorVersion** | Pointer to **int32** | OS minor version. | [optional] 
**BuildNumber** | Pointer to **int32** | OS build number. | [optional] 
**ProductType** | Pointer to [**ProductType**](ProductType.md) |  | [optional] 
**SuiteMask** | Pointer to [**SuiteMask**](SuiteMask.md) |  | [optional] 
**ReleaseId** | Pointer to **string** | OS release Id. It is read from registry and available for Windows 10 and higher. | [optional] 
**UpdateBuildRevision** | Pointer to **string** | OS update build revision number. It is read from registry and available for Windows 10 and higher. | [optional] 
**DisplayVersion** | Pointer to **string** | OS display version. It is read from registry and available for Windows 10 and higher. | [optional] 
**SessionSupport** | Pointer to [**SessionSupport**](SessionSupport.md) |  | [optional] 

## Methods

### NewOperatingSystemInfoResponseModel

`func NewOperatingSystemInfoResponseModel() *OperatingSystemInfoResponseModel`

NewOperatingSystemInfoResponseModel instantiates a new OperatingSystemInfoResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatingSystemInfoResponseModelWithDefaults

`func NewOperatingSystemInfoResponseModelWithDefaults() *OperatingSystemInfoResponseModel`

NewOperatingSystemInfoResponseModelWithDefaults instantiates a new OperatingSystemInfoResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperatingSystemType

`func (o *OperatingSystemInfoResponseModel) GetOperatingSystemType() OperatingSystemType`

GetOperatingSystemType returns the OperatingSystemType field if non-nil, zero value otherwise.

### GetOperatingSystemTypeOk

`func (o *OperatingSystemInfoResponseModel) GetOperatingSystemTypeOk() (*OperatingSystemType, bool)`

GetOperatingSystemTypeOk returns a tuple with the OperatingSystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemType

`func (o *OperatingSystemInfoResponseModel) SetOperatingSystemType(v OperatingSystemType)`

SetOperatingSystemType sets OperatingSystemType field to given value.

### HasOperatingSystemType

`func (o *OperatingSystemInfoResponseModel) HasOperatingSystemType() bool`

HasOperatingSystemType returns a boolean if a field has been set.

### GetMajorVersion

`func (o *OperatingSystemInfoResponseModel) GetMajorVersion() int32`

GetMajorVersion returns the MajorVersion field if non-nil, zero value otherwise.

### GetMajorVersionOk

`func (o *OperatingSystemInfoResponseModel) GetMajorVersionOk() (*int32, bool)`

GetMajorVersionOk returns a tuple with the MajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorVersion

`func (o *OperatingSystemInfoResponseModel) SetMajorVersion(v int32)`

SetMajorVersion sets MajorVersion field to given value.

### HasMajorVersion

`func (o *OperatingSystemInfoResponseModel) HasMajorVersion() bool`

HasMajorVersion returns a boolean if a field has been set.

### GetMinorVersion

`func (o *OperatingSystemInfoResponseModel) GetMinorVersion() int32`

GetMinorVersion returns the MinorVersion field if non-nil, zero value otherwise.

### GetMinorVersionOk

`func (o *OperatingSystemInfoResponseModel) GetMinorVersionOk() (*int32, bool)`

GetMinorVersionOk returns a tuple with the MinorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorVersion

`func (o *OperatingSystemInfoResponseModel) SetMinorVersion(v int32)`

SetMinorVersion sets MinorVersion field to given value.

### HasMinorVersion

`func (o *OperatingSystemInfoResponseModel) HasMinorVersion() bool`

HasMinorVersion returns a boolean if a field has been set.

### GetBuildNumber

`func (o *OperatingSystemInfoResponseModel) GetBuildNumber() int32`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *OperatingSystemInfoResponseModel) GetBuildNumberOk() (*int32, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *OperatingSystemInfoResponseModel) SetBuildNumber(v int32)`

SetBuildNumber sets BuildNumber field to given value.

### HasBuildNumber

`func (o *OperatingSystemInfoResponseModel) HasBuildNumber() bool`

HasBuildNumber returns a boolean if a field has been set.

### GetProductType

`func (o *OperatingSystemInfoResponseModel) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *OperatingSystemInfoResponseModel) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *OperatingSystemInfoResponseModel) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *OperatingSystemInfoResponseModel) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetSuiteMask

`func (o *OperatingSystemInfoResponseModel) GetSuiteMask() SuiteMask`

GetSuiteMask returns the SuiteMask field if non-nil, zero value otherwise.

### GetSuiteMaskOk

`func (o *OperatingSystemInfoResponseModel) GetSuiteMaskOk() (*SuiteMask, bool)`

GetSuiteMaskOk returns a tuple with the SuiteMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuiteMask

`func (o *OperatingSystemInfoResponseModel) SetSuiteMask(v SuiteMask)`

SetSuiteMask sets SuiteMask field to given value.

### HasSuiteMask

`func (o *OperatingSystemInfoResponseModel) HasSuiteMask() bool`

HasSuiteMask returns a boolean if a field has been set.

### GetReleaseId

`func (o *OperatingSystemInfoResponseModel) GetReleaseId() string`

GetReleaseId returns the ReleaseId field if non-nil, zero value otherwise.

### GetReleaseIdOk

`func (o *OperatingSystemInfoResponseModel) GetReleaseIdOk() (*string, bool)`

GetReleaseIdOk returns a tuple with the ReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseId

`func (o *OperatingSystemInfoResponseModel) SetReleaseId(v string)`

SetReleaseId sets ReleaseId field to given value.

### HasReleaseId

`func (o *OperatingSystemInfoResponseModel) HasReleaseId() bool`

HasReleaseId returns a boolean if a field has been set.

### GetUpdateBuildRevision

`func (o *OperatingSystemInfoResponseModel) GetUpdateBuildRevision() string`

GetUpdateBuildRevision returns the UpdateBuildRevision field if non-nil, zero value otherwise.

### GetUpdateBuildRevisionOk

`func (o *OperatingSystemInfoResponseModel) GetUpdateBuildRevisionOk() (*string, bool)`

GetUpdateBuildRevisionOk returns a tuple with the UpdateBuildRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateBuildRevision

`func (o *OperatingSystemInfoResponseModel) SetUpdateBuildRevision(v string)`

SetUpdateBuildRevision sets UpdateBuildRevision field to given value.

### HasUpdateBuildRevision

`func (o *OperatingSystemInfoResponseModel) HasUpdateBuildRevision() bool`

HasUpdateBuildRevision returns a boolean if a field has been set.

### GetDisplayVersion

`func (o *OperatingSystemInfoResponseModel) GetDisplayVersion() string`

GetDisplayVersion returns the DisplayVersion field if non-nil, zero value otherwise.

### GetDisplayVersionOk

`func (o *OperatingSystemInfoResponseModel) GetDisplayVersionOk() (*string, bool)`

GetDisplayVersionOk returns a tuple with the DisplayVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayVersion

`func (o *OperatingSystemInfoResponseModel) SetDisplayVersion(v string)`

SetDisplayVersion sets DisplayVersion field to given value.

### HasDisplayVersion

`func (o *OperatingSystemInfoResponseModel) HasDisplayVersion() bool`

HasDisplayVersion returns a boolean if a field has been set.

### GetSessionSupport

`func (o *OperatingSystemInfoResponseModel) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *OperatingSystemInfoResponseModel) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *OperatingSystemInfoResponseModel) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.

### HasSessionSupport

`func (o *OperatingSystemInfoResponseModel) HasSessionSupport() bool`

HasSessionSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


