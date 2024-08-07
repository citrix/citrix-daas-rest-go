# HypervisorResourceSettingsRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmSettings** | Pointer to [**[]VmSettingRequestModel**](VmSettingRequestModel.md) | The settings for the virtual machines. | [optional] 
**DiskSettings** | Pointer to [**[]DiskSettingRequestModel**](DiskSettingRequestModel.md) | The settings for the disks. | [optional] 

## Methods

### NewHypervisorResourceSettingsRequestModel

`func NewHypervisorResourceSettingsRequestModel() *HypervisorResourceSettingsRequestModel`

NewHypervisorResourceSettingsRequestModel instantiates a new HypervisorResourceSettingsRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorResourceSettingsRequestModelWithDefaults

`func NewHypervisorResourceSettingsRequestModelWithDefaults() *HypervisorResourceSettingsRequestModel`

NewHypervisorResourceSettingsRequestModelWithDefaults instantiates a new HypervisorResourceSettingsRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmSettings

`func (o *HypervisorResourceSettingsRequestModel) GetVmSettings() []VmSettingRequestModel`

GetVmSettings returns the VmSettings field if non-nil, zero value otherwise.

### GetVmSettingsOk

`func (o *HypervisorResourceSettingsRequestModel) GetVmSettingsOk() (*[]VmSettingRequestModel, bool)`

GetVmSettingsOk returns a tuple with the VmSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSettings

`func (o *HypervisorResourceSettingsRequestModel) SetVmSettings(v []VmSettingRequestModel)`

SetVmSettings sets VmSettings field to given value.

### HasVmSettings

`func (o *HypervisorResourceSettingsRequestModel) HasVmSettings() bool`

HasVmSettings returns a boolean if a field has been set.

### SetVmSettingsNil

`func (o *HypervisorResourceSettingsRequestModel) SetVmSettingsNil(b bool)`

 SetVmSettingsNil sets the value for VmSettings to be an explicit nil

### UnsetVmSettings
`func (o *HypervisorResourceSettingsRequestModel) UnsetVmSettings()`

UnsetVmSettings ensures that no value is present for VmSettings, not even an explicit nil
### GetDiskSettings

`func (o *HypervisorResourceSettingsRequestModel) GetDiskSettings() []DiskSettingRequestModel`

GetDiskSettings returns the DiskSettings field if non-nil, zero value otherwise.

### GetDiskSettingsOk

`func (o *HypervisorResourceSettingsRequestModel) GetDiskSettingsOk() (*[]DiskSettingRequestModel, bool)`

GetDiskSettingsOk returns a tuple with the DiskSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSettings

`func (o *HypervisorResourceSettingsRequestModel) SetDiskSettings(v []DiskSettingRequestModel)`

SetDiskSettings sets DiskSettings field to given value.

### HasDiskSettings

`func (o *HypervisorResourceSettingsRequestModel) HasDiskSettings() bool`

HasDiskSettings returns a boolean if a field has been set.

### SetDiskSettingsNil

`func (o *HypervisorResourceSettingsRequestModel) SetDiskSettingsNil(b bool)`

 SetDiskSettingsNil sets the value for DiskSettings to be an explicit nil

### UnsetDiskSettings
`func (o *HypervisorResourceSettingsRequestModel) UnsetDiskSettings()`

UnsetDiskSettings ensures that no value is present for DiskSettings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


