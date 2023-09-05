# ImageRuntimeInfoResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | The JSON schema version of image runtime info file. | [optional] 
**OperatingSystemInfo** | Pointer to [**ImageRuntimeInfoResponseModelOperatingSystemInfo**](ImageRuntimeInfoResponseModelOperatingSystemInfo.md) |  | [optional] 
**VdaComponents** | Pointer to [**[]VdaComponentResponseModel**](VdaComponentResponseModel.md) | Installed VDA components information | [optional] 
**SystemManagementInfo** | Pointer to [**ImageRuntimeInfoResponseModelSystemManagementInfo**](ImageRuntimeInfoResponseModelSystemManagementInfo.md) |  | [optional] 

## Methods

### NewImageRuntimeInfoResponseModel

`func NewImageRuntimeInfoResponseModel() *ImageRuntimeInfoResponseModel`

NewImageRuntimeInfoResponseModel instantiates a new ImageRuntimeInfoResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageRuntimeInfoResponseModelWithDefaults

`func NewImageRuntimeInfoResponseModelWithDefaults() *ImageRuntimeInfoResponseModel`

NewImageRuntimeInfoResponseModelWithDefaults instantiates a new ImageRuntimeInfoResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ImageRuntimeInfoResponseModel) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ImageRuntimeInfoResponseModel) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ImageRuntimeInfoResponseModel) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ImageRuntimeInfoResponseModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetOperatingSystemInfo

`func (o *ImageRuntimeInfoResponseModel) GetOperatingSystemInfo() ImageRuntimeInfoResponseModelOperatingSystemInfo`

GetOperatingSystemInfo returns the OperatingSystemInfo field if non-nil, zero value otherwise.

### GetOperatingSystemInfoOk

`func (o *ImageRuntimeInfoResponseModel) GetOperatingSystemInfoOk() (*ImageRuntimeInfoResponseModelOperatingSystemInfo, bool)`

GetOperatingSystemInfoOk returns a tuple with the OperatingSystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemInfo

`func (o *ImageRuntimeInfoResponseModel) SetOperatingSystemInfo(v ImageRuntimeInfoResponseModelOperatingSystemInfo)`

SetOperatingSystemInfo sets OperatingSystemInfo field to given value.

### HasOperatingSystemInfo

`func (o *ImageRuntimeInfoResponseModel) HasOperatingSystemInfo() bool`

HasOperatingSystemInfo returns a boolean if a field has been set.

### GetVdaComponents

`func (o *ImageRuntimeInfoResponseModel) GetVdaComponents() []VdaComponentResponseModel`

GetVdaComponents returns the VdaComponents field if non-nil, zero value otherwise.

### GetVdaComponentsOk

`func (o *ImageRuntimeInfoResponseModel) GetVdaComponentsOk() (*[]VdaComponentResponseModel, bool)`

GetVdaComponentsOk returns a tuple with the VdaComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaComponents

`func (o *ImageRuntimeInfoResponseModel) SetVdaComponents(v []VdaComponentResponseModel)`

SetVdaComponents sets VdaComponents field to given value.

### HasVdaComponents

`func (o *ImageRuntimeInfoResponseModel) HasVdaComponents() bool`

HasVdaComponents returns a boolean if a field has been set.

### GetSystemManagementInfo

`func (o *ImageRuntimeInfoResponseModel) GetSystemManagementInfo() ImageRuntimeInfoResponseModelSystemManagementInfo`

GetSystemManagementInfo returns the SystemManagementInfo field if non-nil, zero value otherwise.

### GetSystemManagementInfoOk

`func (o *ImageRuntimeInfoResponseModel) GetSystemManagementInfoOk() (*ImageRuntimeInfoResponseModelSystemManagementInfo, bool)`

GetSystemManagementInfoOk returns a tuple with the SystemManagementInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemManagementInfo

`func (o *ImageRuntimeInfoResponseModel) SetSystemManagementInfo(v ImageRuntimeInfoResponseModelSystemManagementInfo)`

SetSystemManagementInfo sets SystemManagementInfo field to given value.

### HasSystemManagementInfo

`func (o *ImageRuntimeInfoResponseModel) HasSystemManagementInfo() bool`

HasSystemManagementInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


