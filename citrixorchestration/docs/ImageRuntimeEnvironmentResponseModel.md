# ImageRuntimeEnvironmentResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperatingSystem** | Pointer to [**OperatingSystemInfoResponseModel**](OperatingSystemInfoResponseModel.md) |  | [optional] 
**VDASessionSupport** | Pointer to **NullableString** | The session support type of VDA in master image | [optional] 
**IdentityJoinStatus** | Pointer to **NullableString** | Identity join status. | [optional] 
**DeviceEnrollmentStatus** | Pointer to **NullableString** | The enrollment mode (MDMEnrolled/MAMEnrolled/Unknown). | [optional] 
**PagingFileSettings** | Pointer to [**[]PagingFileSettingResponseModel**](PagingFileSettingResponseModel.md) | System paging file settings. | [optional] 
**Capabilities** | Pointer to **map[string]string** | Master image&#39;s capabilities, such as &#39;AzureAdSupported&#39;, &#39;HybridAzureAdSupported&#39;, &#39;IntuneSupported&#39; | [optional] 
**VDAComponents** | Pointer to **map[string]string** | Master image&#39;s VDA&#39;s components, such as &#39;MCS IO&#39;, &#39;Machine Identity Service Agent&#39; | [optional] 

## Methods

### NewImageRuntimeEnvironmentResponseModel

`func NewImageRuntimeEnvironmentResponseModel() *ImageRuntimeEnvironmentResponseModel`

NewImageRuntimeEnvironmentResponseModel instantiates a new ImageRuntimeEnvironmentResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageRuntimeEnvironmentResponseModelWithDefaults

`func NewImageRuntimeEnvironmentResponseModelWithDefaults() *ImageRuntimeEnvironmentResponseModel`

NewImageRuntimeEnvironmentResponseModelWithDefaults instantiates a new ImageRuntimeEnvironmentResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperatingSystem

`func (o *ImageRuntimeEnvironmentResponseModel) GetOperatingSystem() OperatingSystemInfoResponseModel`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *ImageRuntimeEnvironmentResponseModel) GetOperatingSystemOk() (*OperatingSystemInfoResponseModel, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *ImageRuntimeEnvironmentResponseModel) SetOperatingSystem(v OperatingSystemInfoResponseModel)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *ImageRuntimeEnvironmentResponseModel) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetVDASessionSupport

`func (o *ImageRuntimeEnvironmentResponseModel) GetVDASessionSupport() string`

GetVDASessionSupport returns the VDASessionSupport field if non-nil, zero value otherwise.

### GetVDASessionSupportOk

`func (o *ImageRuntimeEnvironmentResponseModel) GetVDASessionSupportOk() (*string, bool)`

GetVDASessionSupportOk returns a tuple with the VDASessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVDASessionSupport

`func (o *ImageRuntimeEnvironmentResponseModel) SetVDASessionSupport(v string)`

SetVDASessionSupport sets VDASessionSupport field to given value.

### HasVDASessionSupport

`func (o *ImageRuntimeEnvironmentResponseModel) HasVDASessionSupport() bool`

HasVDASessionSupport returns a boolean if a field has been set.

### SetVDASessionSupportNil

`func (o *ImageRuntimeEnvironmentResponseModel) SetVDASessionSupportNil(b bool)`

 SetVDASessionSupportNil sets the value for VDASessionSupport to be an explicit nil

### UnsetVDASessionSupport
`func (o *ImageRuntimeEnvironmentResponseModel) UnsetVDASessionSupport()`

UnsetVDASessionSupport ensures that no value is present for VDASessionSupport, not even an explicit nil
### GetIdentityJoinStatus

`func (o *ImageRuntimeEnvironmentResponseModel) GetIdentityJoinStatus() string`

GetIdentityJoinStatus returns the IdentityJoinStatus field if non-nil, zero value otherwise.

### GetIdentityJoinStatusOk

`func (o *ImageRuntimeEnvironmentResponseModel) GetIdentityJoinStatusOk() (*string, bool)`

GetIdentityJoinStatusOk returns a tuple with the IdentityJoinStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityJoinStatus

`func (o *ImageRuntimeEnvironmentResponseModel) SetIdentityJoinStatus(v string)`

SetIdentityJoinStatus sets IdentityJoinStatus field to given value.

### HasIdentityJoinStatus

`func (o *ImageRuntimeEnvironmentResponseModel) HasIdentityJoinStatus() bool`

HasIdentityJoinStatus returns a boolean if a field has been set.

### SetIdentityJoinStatusNil

`func (o *ImageRuntimeEnvironmentResponseModel) SetIdentityJoinStatusNil(b bool)`

 SetIdentityJoinStatusNil sets the value for IdentityJoinStatus to be an explicit nil

### UnsetIdentityJoinStatus
`func (o *ImageRuntimeEnvironmentResponseModel) UnsetIdentityJoinStatus()`

UnsetIdentityJoinStatus ensures that no value is present for IdentityJoinStatus, not even an explicit nil
### GetDeviceEnrollmentStatus

`func (o *ImageRuntimeEnvironmentResponseModel) GetDeviceEnrollmentStatus() string`

GetDeviceEnrollmentStatus returns the DeviceEnrollmentStatus field if non-nil, zero value otherwise.

### GetDeviceEnrollmentStatusOk

`func (o *ImageRuntimeEnvironmentResponseModel) GetDeviceEnrollmentStatusOk() (*string, bool)`

GetDeviceEnrollmentStatusOk returns a tuple with the DeviceEnrollmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceEnrollmentStatus

`func (o *ImageRuntimeEnvironmentResponseModel) SetDeviceEnrollmentStatus(v string)`

SetDeviceEnrollmentStatus sets DeviceEnrollmentStatus field to given value.

### HasDeviceEnrollmentStatus

`func (o *ImageRuntimeEnvironmentResponseModel) HasDeviceEnrollmentStatus() bool`

HasDeviceEnrollmentStatus returns a boolean if a field has been set.

### SetDeviceEnrollmentStatusNil

`func (o *ImageRuntimeEnvironmentResponseModel) SetDeviceEnrollmentStatusNil(b bool)`

 SetDeviceEnrollmentStatusNil sets the value for DeviceEnrollmentStatus to be an explicit nil

### UnsetDeviceEnrollmentStatus
`func (o *ImageRuntimeEnvironmentResponseModel) UnsetDeviceEnrollmentStatus()`

UnsetDeviceEnrollmentStatus ensures that no value is present for DeviceEnrollmentStatus, not even an explicit nil
### GetPagingFileSettings

`func (o *ImageRuntimeEnvironmentResponseModel) GetPagingFileSettings() []PagingFileSettingResponseModel`

GetPagingFileSettings returns the PagingFileSettings field if non-nil, zero value otherwise.

### GetPagingFileSettingsOk

`func (o *ImageRuntimeEnvironmentResponseModel) GetPagingFileSettingsOk() (*[]PagingFileSettingResponseModel, bool)`

GetPagingFileSettingsOk returns a tuple with the PagingFileSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagingFileSettings

`func (o *ImageRuntimeEnvironmentResponseModel) SetPagingFileSettings(v []PagingFileSettingResponseModel)`

SetPagingFileSettings sets PagingFileSettings field to given value.

### HasPagingFileSettings

`func (o *ImageRuntimeEnvironmentResponseModel) HasPagingFileSettings() bool`

HasPagingFileSettings returns a boolean if a field has been set.

### SetPagingFileSettingsNil

`func (o *ImageRuntimeEnvironmentResponseModel) SetPagingFileSettingsNil(b bool)`

 SetPagingFileSettingsNil sets the value for PagingFileSettings to be an explicit nil

### UnsetPagingFileSettings
`func (o *ImageRuntimeEnvironmentResponseModel) UnsetPagingFileSettings()`

UnsetPagingFileSettings ensures that no value is present for PagingFileSettings, not even an explicit nil
### GetCapabilities

`func (o *ImageRuntimeEnvironmentResponseModel) GetCapabilities() map[string]string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ImageRuntimeEnvironmentResponseModel) GetCapabilitiesOk() (*map[string]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ImageRuntimeEnvironmentResponseModel) SetCapabilities(v map[string]string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ImageRuntimeEnvironmentResponseModel) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *ImageRuntimeEnvironmentResponseModel) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *ImageRuntimeEnvironmentResponseModel) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetVDAComponents

`func (o *ImageRuntimeEnvironmentResponseModel) GetVDAComponents() map[string]string`

GetVDAComponents returns the VDAComponents field if non-nil, zero value otherwise.

### GetVDAComponentsOk

`func (o *ImageRuntimeEnvironmentResponseModel) GetVDAComponentsOk() (*map[string]string, bool)`

GetVDAComponentsOk returns a tuple with the VDAComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVDAComponents

`func (o *ImageRuntimeEnvironmentResponseModel) SetVDAComponents(v map[string]string)`

SetVDAComponents sets VDAComponents field to given value.

### HasVDAComponents

`func (o *ImageRuntimeEnvironmentResponseModel) HasVDAComponents() bool`

HasVDAComponents returns a boolean if a field has been set.

### SetVDAComponentsNil

`func (o *ImageRuntimeEnvironmentResponseModel) SetVDAComponentsNil(b bool)`

 SetVDAComponentsNil sets the value for VDAComponents to be an explicit nil

### UnsetVDAComponents
`func (o *ImageRuntimeEnvironmentResponseModel) UnsetVDAComponents()`

UnsetVDAComponents ensures that no value is present for VDAComponents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


