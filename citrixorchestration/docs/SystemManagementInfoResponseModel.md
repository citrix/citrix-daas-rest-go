# SystemManagementInfoResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureADJoinStatus** | Pointer to [**AzureADJoinStatusResponseModel**](AzureADJoinStatusResponseModel.md) |  | [optional] 
**DeviceRegistrationStatus** | Pointer to [**DeviceRegistrationStatusResponseModel**](DeviceRegistrationStatusResponseModel.md) |  | [optional] 
**PagingFileSettings** | Pointer to [**[]PagingFileSettingResponseModel**](PagingFileSettingResponseModel.md) | System paging file settings. | [optional] 

## Methods

### NewSystemManagementInfoResponseModel

`func NewSystemManagementInfoResponseModel() *SystemManagementInfoResponseModel`

NewSystemManagementInfoResponseModel instantiates a new SystemManagementInfoResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemManagementInfoResponseModelWithDefaults

`func NewSystemManagementInfoResponseModelWithDefaults() *SystemManagementInfoResponseModel`

NewSystemManagementInfoResponseModelWithDefaults instantiates a new SystemManagementInfoResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureADJoinStatus

`func (o *SystemManagementInfoResponseModel) GetAzureADJoinStatus() AzureADJoinStatusResponseModel`

GetAzureADJoinStatus returns the AzureADJoinStatus field if non-nil, zero value otherwise.

### GetAzureADJoinStatusOk

`func (o *SystemManagementInfoResponseModel) GetAzureADJoinStatusOk() (*AzureADJoinStatusResponseModel, bool)`

GetAzureADJoinStatusOk returns a tuple with the AzureADJoinStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureADJoinStatus

`func (o *SystemManagementInfoResponseModel) SetAzureADJoinStatus(v AzureADJoinStatusResponseModel)`

SetAzureADJoinStatus sets AzureADJoinStatus field to given value.

### HasAzureADJoinStatus

`func (o *SystemManagementInfoResponseModel) HasAzureADJoinStatus() bool`

HasAzureADJoinStatus returns a boolean if a field has been set.

### GetDeviceRegistrationStatus

`func (o *SystemManagementInfoResponseModel) GetDeviceRegistrationStatus() DeviceRegistrationStatusResponseModel`

GetDeviceRegistrationStatus returns the DeviceRegistrationStatus field if non-nil, zero value otherwise.

### GetDeviceRegistrationStatusOk

`func (o *SystemManagementInfoResponseModel) GetDeviceRegistrationStatusOk() (*DeviceRegistrationStatusResponseModel, bool)`

GetDeviceRegistrationStatusOk returns a tuple with the DeviceRegistrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistrationStatus

`func (o *SystemManagementInfoResponseModel) SetDeviceRegistrationStatus(v DeviceRegistrationStatusResponseModel)`

SetDeviceRegistrationStatus sets DeviceRegistrationStatus field to given value.

### HasDeviceRegistrationStatus

`func (o *SystemManagementInfoResponseModel) HasDeviceRegistrationStatus() bool`

HasDeviceRegistrationStatus returns a boolean if a field has been set.

### GetPagingFileSettings

`func (o *SystemManagementInfoResponseModel) GetPagingFileSettings() []PagingFileSettingResponseModel`

GetPagingFileSettings returns the PagingFileSettings field if non-nil, zero value otherwise.

### GetPagingFileSettingsOk

`func (o *SystemManagementInfoResponseModel) GetPagingFileSettingsOk() (*[]PagingFileSettingResponseModel, bool)`

GetPagingFileSettingsOk returns a tuple with the PagingFileSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagingFileSettings

`func (o *SystemManagementInfoResponseModel) SetPagingFileSettings(v []PagingFileSettingResponseModel)`

SetPagingFileSettings sets PagingFileSettings field to given value.

### HasPagingFileSettings

`func (o *SystemManagementInfoResponseModel) HasPagingFileSettings() bool`

HasPagingFileSettings returns a boolean if a field has been set.

### SetPagingFileSettingsNil

`func (o *SystemManagementInfoResponseModel) SetPagingFileSettingsNil(b bool)`

 SetPagingFileSettingsNil sets the value for PagingFileSettings to be an explicit nil

### UnsetPagingFileSettings
`func (o *SystemManagementInfoResponseModel) UnsetPagingFileSettings()`

UnsetPagingFileSettings ensures that no value is present for PagingFileSettings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


