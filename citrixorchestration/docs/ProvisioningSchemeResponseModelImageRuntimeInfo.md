# ProvisioningSchemeResponseModelImageRuntimeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | The JSON schema version of image runtime info file. | [optional] 
**OperatingSystemInfo** | Pointer to [**ImageRuntimeInfoResponseModelOperatingSystemInfo**](ImageRuntimeInfoResponseModelOperatingSystemInfo.md) |  | [optional] 
**VdaComponents** | Pointer to [**[]VdaComponentResponseModel**](VdaComponentResponseModel.md) | Installed VDA components information | [optional] 
**SystemManagementInfo** | Pointer to [**ImageRuntimeInfoResponseModelSystemManagementInfo**](ImageRuntimeInfoResponseModelSystemManagementInfo.md) |  | [optional] 

## Methods

### NewProvisioningSchemeResponseModelImageRuntimeInfo

`func NewProvisioningSchemeResponseModelImageRuntimeInfo() *ProvisioningSchemeResponseModelImageRuntimeInfo`

NewProvisioningSchemeResponseModelImageRuntimeInfo instantiates a new ProvisioningSchemeResponseModelImageRuntimeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningSchemeResponseModelImageRuntimeInfoWithDefaults

`func NewProvisioningSchemeResponseModelImageRuntimeInfoWithDefaults() *ProvisioningSchemeResponseModelImageRuntimeInfo`

NewProvisioningSchemeResponseModelImageRuntimeInfoWithDefaults instantiates a new ProvisioningSchemeResponseModelImageRuntimeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetOperatingSystemInfo

`func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) GetOperatingSystemInfo() ImageRuntimeInfoResponseModelOperatingSystemInfo`

GetOperatingSystemInfo returns the OperatingSystemInfo field if non-nil, zero value otherwise.

### GetOperatingSystemInfoOk

`func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) GetOperatingSystemInfoOk() (*ImageRuntimeInfoResponseModelOperatingSystemInfo, bool)`

GetOperatingSystemInfoOk returns a tuple with the OperatingSystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemInfo

`func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) SetOperatingSystemInfo(v ImageRuntimeInfoResponseModelOperatingSystemInfo)`

SetOperatingSystemInfo sets OperatingSystemInfo field to given value.

### HasOperatingSystemInfo

`func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) HasOperatingSystemInfo() bool`

HasOperatingSystemInfo returns a boolean if a field has been set.

### GetVdaComponents

`func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) GetVdaComponents() []VdaComponentResponseModel`

GetVdaComponents returns the VdaComponents field if non-nil, zero value otherwise.

### GetVdaComponentsOk

`func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) GetVdaComponentsOk() (*[]VdaComponentResponseModel, bool)`

GetVdaComponentsOk returns a tuple with the VdaComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaComponents

`func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) SetVdaComponents(v []VdaComponentResponseModel)`

SetVdaComponents sets VdaComponents field to given value.

### HasVdaComponents

`func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) HasVdaComponents() bool`

HasVdaComponents returns a boolean if a field has been set.

### GetSystemManagementInfo

`func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) GetSystemManagementInfo() ImageRuntimeInfoResponseModelSystemManagementInfo`

GetSystemManagementInfo returns the SystemManagementInfo field if non-nil, zero value otherwise.

### GetSystemManagementInfoOk

`func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) GetSystemManagementInfoOk() (*ImageRuntimeInfoResponseModelSystemManagementInfo, bool)`

GetSystemManagementInfoOk returns a tuple with the SystemManagementInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemManagementInfo

`func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) SetSystemManagementInfo(v ImageRuntimeInfoResponseModelSystemManagementInfo)`

SetSystemManagementInfo sets SystemManagementInfo field to given value.

### HasSystemManagementInfo

`func (o *ProvisioningSchemeResponseModelImageRuntimeInfo) HasSystemManagementInfo() bool`

HasSystemManagementInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


