# ImageRuntimeInfoResponseModelOperatingSystemInfo

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

### NewImageRuntimeInfoResponseModelOperatingSystemInfo

`func NewImageRuntimeInfoResponseModelOperatingSystemInfo() *ImageRuntimeInfoResponseModelOperatingSystemInfo`

NewImageRuntimeInfoResponseModelOperatingSystemInfo instantiates a new ImageRuntimeInfoResponseModelOperatingSystemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageRuntimeInfoResponseModelOperatingSystemInfoWithDefaults

`func NewImageRuntimeInfoResponseModelOperatingSystemInfoWithDefaults() *ImageRuntimeInfoResponseModelOperatingSystemInfo`

NewImageRuntimeInfoResponseModelOperatingSystemInfoWithDefaults instantiates a new ImageRuntimeInfoResponseModelOperatingSystemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperatingSystemType

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) GetOperatingSystemType() OperatingSystemType`

GetOperatingSystemType returns the OperatingSystemType field if non-nil, zero value otherwise.

### GetOperatingSystemTypeOk

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) GetOperatingSystemTypeOk() (*OperatingSystemType, bool)`

GetOperatingSystemTypeOk returns a tuple with the OperatingSystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemType

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) SetOperatingSystemType(v OperatingSystemType)`

SetOperatingSystemType sets OperatingSystemType field to given value.

### HasOperatingSystemType

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) HasOperatingSystemType() bool`

HasOperatingSystemType returns a boolean if a field has been set.

### GetMajorVersion

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) GetMajorVersion() int32`

GetMajorVersion returns the MajorVersion field if non-nil, zero value otherwise.

### GetMajorVersionOk

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) GetMajorVersionOk() (*int32, bool)`

GetMajorVersionOk returns a tuple with the MajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorVersion

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) SetMajorVersion(v int32)`

SetMajorVersion sets MajorVersion field to given value.

### HasMajorVersion

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) HasMajorVersion() bool`

HasMajorVersion returns a boolean if a field has been set.

### GetMinorVersion

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) GetMinorVersion() int32`

GetMinorVersion returns the MinorVersion field if non-nil, zero value otherwise.

### GetMinorVersionOk

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) GetMinorVersionOk() (*int32, bool)`

GetMinorVersionOk returns a tuple with the MinorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorVersion

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) SetMinorVersion(v int32)`

SetMinorVersion sets MinorVersion field to given value.

### HasMinorVersion

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) HasMinorVersion() bool`

HasMinorVersion returns a boolean if a field has been set.

### GetBuildNumber

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) GetBuildNumber() int32`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) GetBuildNumberOk() (*int32, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) SetBuildNumber(v int32)`

SetBuildNumber sets BuildNumber field to given value.

### HasBuildNumber

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) HasBuildNumber() bool`

HasBuildNumber returns a boolean if a field has been set.

### GetProductType

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetSuiteMask

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) GetSuiteMask() SuiteMask`

GetSuiteMask returns the SuiteMask field if non-nil, zero value otherwise.

### GetSuiteMaskOk

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) GetSuiteMaskOk() (*SuiteMask, bool)`

GetSuiteMaskOk returns a tuple with the SuiteMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuiteMask

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) SetSuiteMask(v SuiteMask)`

SetSuiteMask sets SuiteMask field to given value.

### HasSuiteMask

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) HasSuiteMask() bool`

HasSuiteMask returns a boolean if a field has been set.

### GetReleaseId

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) GetReleaseId() string`

GetReleaseId returns the ReleaseId field if non-nil, zero value otherwise.

### GetReleaseIdOk

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) GetReleaseIdOk() (*string, bool)`

GetReleaseIdOk returns a tuple with the ReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseId

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) SetReleaseId(v string)`

SetReleaseId sets ReleaseId field to given value.

### HasReleaseId

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) HasReleaseId() bool`

HasReleaseId returns a boolean if a field has been set.

### GetUpdateBuildRevision

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) GetUpdateBuildRevision() string`

GetUpdateBuildRevision returns the UpdateBuildRevision field if non-nil, zero value otherwise.

### GetUpdateBuildRevisionOk

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) GetUpdateBuildRevisionOk() (*string, bool)`

GetUpdateBuildRevisionOk returns a tuple with the UpdateBuildRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateBuildRevision

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) SetUpdateBuildRevision(v string)`

SetUpdateBuildRevision sets UpdateBuildRevision field to given value.

### HasUpdateBuildRevision

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) HasUpdateBuildRevision() bool`

HasUpdateBuildRevision returns a boolean if a field has been set.

### GetDisplayVersion

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) GetDisplayVersion() string`

GetDisplayVersion returns the DisplayVersion field if non-nil, zero value otherwise.

### GetDisplayVersionOk

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) GetDisplayVersionOk() (*string, bool)`

GetDisplayVersionOk returns a tuple with the DisplayVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayVersion

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) SetDisplayVersion(v string)`

SetDisplayVersion sets DisplayVersion field to given value.

### HasDisplayVersion

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) HasDisplayVersion() bool`

HasDisplayVersion returns a boolean if a field has been set.

### GetSessionSupport

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) GetSessionSupport() SessionSupport`

GetSessionSupport returns the SessionSupport field if non-nil, zero value otherwise.

### GetSessionSupportOk

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) GetSessionSupportOk() (*SessionSupport, bool)`

GetSessionSupportOk returns a tuple with the SessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSupport

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) SetSessionSupport(v SessionSupport)`

SetSessionSupport sets SessionSupport field to given value.

### HasSessionSupport

`func (o *ImageRuntimeInfoResponseModelOperatingSystemInfo) HasSessionSupport() bool`

HasSessionSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


