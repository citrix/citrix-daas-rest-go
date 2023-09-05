# CreateImageVersionRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterImageVM** | Pointer to **string** | The path in the resource pool to the virtual machine snapshot or VM template that will be used. This identifies the hard disk to be used and the default values for the memory and processors. This must be a path to a Snapshot or Template item in the resource pool to which the Image Version is associated. | [optional] 
**ServiceOffering** | Pointer to **string** | The hypervisor resource path of the Cloud service offering to use when creating machines. | [optional] 
**MachineProfile** | Pointer to **string** | The path in the resource pool to the virtual machine template that will be used. This identifies the VM template to be used and the default values for the tags, virtual machine size, boot diagnostics, host cache property of OS disk, accelerated networking and availability zone. This must be a path to a Virtual machine or Template item in the resource pool to which the Image Version is associated. | [optional] 
**VMCpuCount** | Pointer to **int32** | The number of processors that virtual machines created for the image preparing should use. | [optional] 
**VMMemoryMB** | Pointer to **int32** | The maximum amount of memory that virtual machines created for the image preparing should use. | [optional] 
**WriteBackCacheDiskSizeGB** | Pointer to **int32** | The size in GB of any temporary storage disk used by the write back cache. Should be used in conjunction with WriteBackCacheMemorySizeMB. | [optional] 
**WriteBackCacheMemorySizeMB** | Pointer to **int32** | The size in MB of any write back cache if required. Should be used in conjunction with WriteBackCacheDiskSizeGB. | [optional] 
**Scopes** | Pointer to **[]string** | Administrative scopes which the newly created image version will be a part of. | [optional] 
**NetworkMapping** | Pointer to [**[]NetworkMapRequestModel**](NetworkMapRequestModel.md) | Specifies how the attached NICs are mapped to networks.  If this parameter is omitted, provisioned VMs are created with a single NIC, which is mapped to the default network in the hypervisor resource pool.  If this parameter is supplied, machines are created with the number of NICs specified in the map, and each NIC is attached to the specified network. | [optional] 
**CustomProperties** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The properties of the image version that are specific to the target hosting infrastructure. | [optional] 
**ImageVersionDescription** | Pointer to **string** | The description associated with the image version. | [optional] 

## Methods

### NewCreateImageVersionRequestModel

`func NewCreateImageVersionRequestModel() *CreateImageVersionRequestModel`

NewCreateImageVersionRequestModel instantiates a new CreateImageVersionRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateImageVersionRequestModelWithDefaults

`func NewCreateImageVersionRequestModelWithDefaults() *CreateImageVersionRequestModel`

NewCreateImageVersionRequestModelWithDefaults instantiates a new CreateImageVersionRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMasterImageVM

`func (o *CreateImageVersionRequestModel) GetMasterImageVM() string`

GetMasterImageVM returns the MasterImageVM field if non-nil, zero value otherwise.

### GetMasterImageVMOk

`func (o *CreateImageVersionRequestModel) GetMasterImageVMOk() (*string, bool)`

GetMasterImageVMOk returns a tuple with the MasterImageVM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterImageVM

`func (o *CreateImageVersionRequestModel) SetMasterImageVM(v string)`

SetMasterImageVM sets MasterImageVM field to given value.

### HasMasterImageVM

`func (o *CreateImageVersionRequestModel) HasMasterImageVM() bool`

HasMasterImageVM returns a boolean if a field has been set.

### GetServiceOffering

`func (o *CreateImageVersionRequestModel) GetServiceOffering() string`

GetServiceOffering returns the ServiceOffering field if non-nil, zero value otherwise.

### GetServiceOfferingOk

`func (o *CreateImageVersionRequestModel) GetServiceOfferingOk() (*string, bool)`

GetServiceOfferingOk returns a tuple with the ServiceOffering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceOffering

`func (o *CreateImageVersionRequestModel) SetServiceOffering(v string)`

SetServiceOffering sets ServiceOffering field to given value.

### HasServiceOffering

`func (o *CreateImageVersionRequestModel) HasServiceOffering() bool`

HasServiceOffering returns a boolean if a field has been set.

### GetMachineProfile

`func (o *CreateImageVersionRequestModel) GetMachineProfile() string`

GetMachineProfile returns the MachineProfile field if non-nil, zero value otherwise.

### GetMachineProfileOk

`func (o *CreateImageVersionRequestModel) GetMachineProfileOk() (*string, bool)`

GetMachineProfileOk returns a tuple with the MachineProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineProfile

`func (o *CreateImageVersionRequestModel) SetMachineProfile(v string)`

SetMachineProfile sets MachineProfile field to given value.

### HasMachineProfile

`func (o *CreateImageVersionRequestModel) HasMachineProfile() bool`

HasMachineProfile returns a boolean if a field has been set.

### GetVMCpuCount

`func (o *CreateImageVersionRequestModel) GetVMCpuCount() int32`

GetVMCpuCount returns the VMCpuCount field if non-nil, zero value otherwise.

### GetVMCpuCountOk

`func (o *CreateImageVersionRequestModel) GetVMCpuCountOk() (*int32, bool)`

GetVMCpuCountOk returns a tuple with the VMCpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMCpuCount

`func (o *CreateImageVersionRequestModel) SetVMCpuCount(v int32)`

SetVMCpuCount sets VMCpuCount field to given value.

### HasVMCpuCount

`func (o *CreateImageVersionRequestModel) HasVMCpuCount() bool`

HasVMCpuCount returns a boolean if a field has been set.

### GetVMMemoryMB

`func (o *CreateImageVersionRequestModel) GetVMMemoryMB() int32`

GetVMMemoryMB returns the VMMemoryMB field if non-nil, zero value otherwise.

### GetVMMemoryMBOk

`func (o *CreateImageVersionRequestModel) GetVMMemoryMBOk() (*int32, bool)`

GetVMMemoryMBOk returns a tuple with the VMMemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMMemoryMB

`func (o *CreateImageVersionRequestModel) SetVMMemoryMB(v int32)`

SetVMMemoryMB sets VMMemoryMB field to given value.

### HasVMMemoryMB

`func (o *CreateImageVersionRequestModel) HasVMMemoryMB() bool`

HasVMMemoryMB returns a boolean if a field has been set.

### GetWriteBackCacheDiskSizeGB

`func (o *CreateImageVersionRequestModel) GetWriteBackCacheDiskSizeGB() int32`

GetWriteBackCacheDiskSizeGB returns the WriteBackCacheDiskSizeGB field if non-nil, zero value otherwise.

### GetWriteBackCacheDiskSizeGBOk

`func (o *CreateImageVersionRequestModel) GetWriteBackCacheDiskSizeGBOk() (*int32, bool)`

GetWriteBackCacheDiskSizeGBOk returns a tuple with the WriteBackCacheDiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheDiskSizeGB

`func (o *CreateImageVersionRequestModel) SetWriteBackCacheDiskSizeGB(v int32)`

SetWriteBackCacheDiskSizeGB sets WriteBackCacheDiskSizeGB field to given value.

### HasWriteBackCacheDiskSizeGB

`func (o *CreateImageVersionRequestModel) HasWriteBackCacheDiskSizeGB() bool`

HasWriteBackCacheDiskSizeGB returns a boolean if a field has been set.

### GetWriteBackCacheMemorySizeMB

`func (o *CreateImageVersionRequestModel) GetWriteBackCacheMemorySizeMB() int32`

GetWriteBackCacheMemorySizeMB returns the WriteBackCacheMemorySizeMB field if non-nil, zero value otherwise.

### GetWriteBackCacheMemorySizeMBOk

`func (o *CreateImageVersionRequestModel) GetWriteBackCacheMemorySizeMBOk() (*int32, bool)`

GetWriteBackCacheMemorySizeMBOk returns a tuple with the WriteBackCacheMemorySizeMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackCacheMemorySizeMB

`func (o *CreateImageVersionRequestModel) SetWriteBackCacheMemorySizeMB(v int32)`

SetWriteBackCacheMemorySizeMB sets WriteBackCacheMemorySizeMB field to given value.

### HasWriteBackCacheMemorySizeMB

`func (o *CreateImageVersionRequestModel) HasWriteBackCacheMemorySizeMB() bool`

HasWriteBackCacheMemorySizeMB returns a boolean if a field has been set.

### GetScopes

`func (o *CreateImageVersionRequestModel) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CreateImageVersionRequestModel) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CreateImageVersionRequestModel) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *CreateImageVersionRequestModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetNetworkMapping

`func (o *CreateImageVersionRequestModel) GetNetworkMapping() []NetworkMapRequestModel`

GetNetworkMapping returns the NetworkMapping field if non-nil, zero value otherwise.

### GetNetworkMappingOk

`func (o *CreateImageVersionRequestModel) GetNetworkMappingOk() (*[]NetworkMapRequestModel, bool)`

GetNetworkMappingOk returns a tuple with the NetworkMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMapping

`func (o *CreateImageVersionRequestModel) SetNetworkMapping(v []NetworkMapRequestModel)`

SetNetworkMapping sets NetworkMapping field to given value.

### HasNetworkMapping

`func (o *CreateImageVersionRequestModel) HasNetworkMapping() bool`

HasNetworkMapping returns a boolean if a field has been set.

### GetCustomProperties

`func (o *CreateImageVersionRequestModel) GetCustomProperties() []NameValueStringPairModel`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *CreateImageVersionRequestModel) GetCustomPropertiesOk() (*[]NameValueStringPairModel, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *CreateImageVersionRequestModel) SetCustomProperties(v []NameValueStringPairModel)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *CreateImageVersionRequestModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetImageVersionDescription

`func (o *CreateImageVersionRequestModel) GetImageVersionDescription() string`

GetImageVersionDescription returns the ImageVersionDescription field if non-nil, zero value otherwise.

### GetImageVersionDescriptionOk

`func (o *CreateImageVersionRequestModel) GetImageVersionDescriptionOk() (*string, bool)`

GetImageVersionDescriptionOk returns a tuple with the ImageVersionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersionDescription

`func (o *CreateImageVersionRequestModel) SetImageVersionDescription(v string)`

SetImageVersionDescription sets ImageVersionDescription field to given value.

### HasImageVersionDescription

`func (o *CreateImageVersionRequestModel) HasImageVersionDescription() bool`

HasImageVersionDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


