# HypervisorVmSnapshotResourceResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuCount** | Pointer to **int32** | Number of CPUs, if known. | [optional] 
**MemoryMB** | Pointer to **int32** | Memory in megabytes, if known. | [optional] 
**HardDiskSizeGB** | Pointer to **int32** | Hard disk size in gigabytes, if known. | [optional] 
**MinMemoryMB** | Pointer to **int32** | Minimum memory required to run this VM or snapshot, in megabytes, if known. | [optional] 
**NetworkMappings** | Pointer to [**[]NetworkMapResponseModel**](NetworkMapResponseModel.md) | Network mappings associated with the VM, if known. | [optional] 
**AttachedDisks** | Pointer to [**[]AttachedDiskResponseModel**](AttachedDiskResponseModel.md) | Disks attached to the VM, if known. | [optional] 

## Methods

### NewHypervisorVmSnapshotResourceResponseModelAllOf

`func NewHypervisorVmSnapshotResourceResponseModelAllOf() *HypervisorVmSnapshotResourceResponseModelAllOf`

NewHypervisorVmSnapshotResourceResponseModelAllOf instantiates a new HypervisorVmSnapshotResourceResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorVmSnapshotResourceResponseModelAllOfWithDefaults

`func NewHypervisorVmSnapshotResourceResponseModelAllOfWithDefaults() *HypervisorVmSnapshotResourceResponseModelAllOf`

NewHypervisorVmSnapshotResourceResponseModelAllOfWithDefaults instantiates a new HypervisorVmSnapshotResourceResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuCount

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetMemoryMB

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) GetMemoryMB() int32`

GetMemoryMB returns the MemoryMB field if non-nil, zero value otherwise.

### GetMemoryMBOk

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) GetMemoryMBOk() (*int32, bool)`

GetMemoryMBOk returns a tuple with the MemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMB

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) SetMemoryMB(v int32)`

SetMemoryMB sets MemoryMB field to given value.

### HasMemoryMB

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) HasMemoryMB() bool`

HasMemoryMB returns a boolean if a field has been set.

### GetHardDiskSizeGB

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) GetHardDiskSizeGB() int32`

GetHardDiskSizeGB returns the HardDiskSizeGB field if non-nil, zero value otherwise.

### GetHardDiskSizeGBOk

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) GetHardDiskSizeGBOk() (*int32, bool)`

GetHardDiskSizeGBOk returns a tuple with the HardDiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardDiskSizeGB

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) SetHardDiskSizeGB(v int32)`

SetHardDiskSizeGB sets HardDiskSizeGB field to given value.

### HasHardDiskSizeGB

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) HasHardDiskSizeGB() bool`

HasHardDiskSizeGB returns a boolean if a field has been set.

### GetMinMemoryMB

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) GetMinMemoryMB() int32`

GetMinMemoryMB returns the MinMemoryMB field if non-nil, zero value otherwise.

### GetMinMemoryMBOk

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) GetMinMemoryMBOk() (*int32, bool)`

GetMinMemoryMBOk returns a tuple with the MinMemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemoryMB

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) SetMinMemoryMB(v int32)`

SetMinMemoryMB sets MinMemoryMB field to given value.

### HasMinMemoryMB

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) HasMinMemoryMB() bool`

HasMinMemoryMB returns a boolean if a field has been set.

### GetNetworkMappings

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) GetNetworkMappings() []NetworkMapResponseModel`

GetNetworkMappings returns the NetworkMappings field if non-nil, zero value otherwise.

### GetNetworkMappingsOk

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) GetNetworkMappingsOk() (*[]NetworkMapResponseModel, bool)`

GetNetworkMappingsOk returns a tuple with the NetworkMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMappings

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) SetNetworkMappings(v []NetworkMapResponseModel)`

SetNetworkMappings sets NetworkMappings field to given value.

### HasNetworkMappings

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) HasNetworkMappings() bool`

HasNetworkMappings returns a boolean if a field has been set.

### GetAttachedDisks

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) GetAttachedDisks() []AttachedDiskResponseModel`

GetAttachedDisks returns the AttachedDisks field if non-nil, zero value otherwise.

### GetAttachedDisksOk

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) GetAttachedDisksOk() (*[]AttachedDiskResponseModel, bool)`

GetAttachedDisksOk returns a tuple with the AttachedDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedDisks

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) SetAttachedDisks(v []AttachedDiskResponseModel)`

SetAttachedDisks sets AttachedDisks field to given value.

### HasAttachedDisks

`func (o *HypervisorVmSnapshotResourceResponseModelAllOf) HasAttachedDisks() bool`

HasAttachedDisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


