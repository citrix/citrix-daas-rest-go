# CreateHypervisorResourcePoolGcpRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the resource pool to create.  Required. | 
**VmTagging** | Pointer to **NullableBool** | Indicates whether VMs created by Virtual Apps &amp; Desktops provisioning operations should be tagged.  Tagged VMs are filtered out of queries by default. Optional.  Default is &#x60;true&#x60;. | [optional] [default to true]
**GpuTypes** | Pointer to **[]string** | Path to the GPU type resource(s) that are available for provisioning operations in this resource pool.  Optional.  Not supported by all hypervisor types. | [optional] 
**ConnectionType** | [**HypervisorConnectionType**](HypervisorConnectionType.md) |  | 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Metadata of the resource pool. Optional.  | [optional] 
**VirtualPrivateCloud** | **string** | Google Cloud Platform virtual network which the resource pool is connected to. Required. | 
**AvailabilityZone** | Pointer to **string** | Path to the availability zone resource to use for provisioning operations in this resource pool.  Required. | [optional] 
**Networks** | **[]string** | Path to the subnet(s) that are available for provisioning operations in this resource pool.  At least one is required. | 
**Region** | **string** | Google Cloud Platform region which the resource pool is connected to.  Required. | 
**VirtualNetwork** | Pointer to **string** | Azure virtual network which the resource pool is connected to. Required. | [optional] 
**Subnets** | Pointer to **[]string** | Path to the subnet(s) that are available for provisioning operations in this resource pool.  At least one is required. | [optional] 
**RootPath** | Pointer to **NullableString** | Root path of the resources on the hypervisor which should be included in the resource pool.  Required. | [optional] 
**UseLocalStorageCaching** | Pointer to **NullableBool** | Indicates whether local storage on the hypervisor will be used for caching purposes. Not all hypervisor types support this.  Defaults to &#x60;false&#x60;. | [optional] [default to false]
**Storage** | Pointer to **[]string** | Path to the storage resource(s) that are available for provisioning operations in this resource pool.  Required for some hypervisor types. | [optional] 
**PersonalvDiskStorage** | Pointer to **[]string** | Path to the storage resource(s) that are available for provisioning operations in this resource pool.  Required for some hypervisor types. | [optional] 
**TemporaryStorage** | Pointer to **[]string** | Path to the storage resource(s) that are used for temporary operations in this resource pool.  Required for some hypervisor types. | [optional] 
**CustomProperties** | Pointer to **NullableString** | Custom properties.  Optional.  If not specified, will not be changed.  Only used for hypervisors of type Custom. | [optional] 

## Methods

### NewCreateHypervisorResourcePoolGcpRequestModel

`func NewCreateHypervisorResourcePoolGcpRequestModel(name string, connectionType HypervisorConnectionType, virtualPrivateCloud string, networks []string, region string, ) *CreateHypervisorResourcePoolGcpRequestModel`

NewCreateHypervisorResourcePoolGcpRequestModel instantiates a new CreateHypervisorResourcePoolGcpRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHypervisorResourcePoolGcpRequestModelWithDefaults

`func NewCreateHypervisorResourcePoolGcpRequestModelWithDefaults() *CreateHypervisorResourcePoolGcpRequestModel`

NewCreateHypervisorResourcePoolGcpRequestModelWithDefaults instantiates a new CreateHypervisorResourcePoolGcpRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetVmTagging

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetVmTagging() bool`

GetVmTagging returns the VmTagging field if non-nil, zero value otherwise.

### GetVmTaggingOk

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetVmTaggingOk() (*bool, bool)`

GetVmTaggingOk returns a tuple with the VmTagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTagging

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetVmTagging(v bool)`

SetVmTagging sets VmTagging field to given value.

### HasVmTagging

`func (o *CreateHypervisorResourcePoolGcpRequestModel) HasVmTagging() bool`

HasVmTagging returns a boolean if a field has been set.

### SetVmTaggingNil

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetVmTaggingNil(b bool)`

 SetVmTaggingNil sets the value for VmTagging to be an explicit nil

### UnsetVmTagging
`func (o *CreateHypervisorResourcePoolGcpRequestModel) UnsetVmTagging()`

UnsetVmTagging ensures that no value is present for VmTagging, not even an explicit nil
### GetGpuTypes

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetGpuTypes() []string`

GetGpuTypes returns the GpuTypes field if non-nil, zero value otherwise.

### GetGpuTypesOk

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetGpuTypesOk() (*[]string, bool)`

GetGpuTypesOk returns a tuple with the GpuTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuTypes

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetGpuTypes(v []string)`

SetGpuTypes sets GpuTypes field to given value.

### HasGpuTypes

`func (o *CreateHypervisorResourcePoolGcpRequestModel) HasGpuTypes() bool`

HasGpuTypes returns a boolean if a field has been set.

### SetGpuTypesNil

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetGpuTypesNil(b bool)`

 SetGpuTypesNil sets the value for GpuTypes to be an explicit nil

### UnsetGpuTypes
`func (o *CreateHypervisorResourcePoolGcpRequestModel) UnsetGpuTypes()`

UnsetGpuTypes ensures that no value is present for GpuTypes, not even an explicit nil
### GetConnectionType

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.


### GetMetadata

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateHypervisorResourcePoolGcpRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CreateHypervisorResourcePoolGcpRequestModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetVirtualPrivateCloud

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetVirtualPrivateCloud() string`

GetVirtualPrivateCloud returns the VirtualPrivateCloud field if non-nil, zero value otherwise.

### GetVirtualPrivateCloudOk

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetVirtualPrivateCloudOk() (*string, bool)`

GetVirtualPrivateCloudOk returns a tuple with the VirtualPrivateCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPrivateCloud

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetVirtualPrivateCloud(v string)`

SetVirtualPrivateCloud sets VirtualPrivateCloud field to given value.


### GetAvailabilityZone

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *CreateHypervisorResourcePoolGcpRequestModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetNetworks

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.


### GetRegion

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetVirtualNetwork

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetVirtualNetwork() string`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetVirtualNetworkOk() (*string, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetVirtualNetwork(v string)`

SetVirtualNetwork sets VirtualNetwork field to given value.

### HasVirtualNetwork

`func (o *CreateHypervisorResourcePoolGcpRequestModel) HasVirtualNetwork() bool`

HasVirtualNetwork returns a boolean if a field has been set.

### GetSubnets

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *CreateHypervisorResourcePoolGcpRequestModel) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetRootPath

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetRootPath() string`

GetRootPath returns the RootPath field if non-nil, zero value otherwise.

### GetRootPathOk

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetRootPathOk() (*string, bool)`

GetRootPathOk returns a tuple with the RootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPath

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetRootPath(v string)`

SetRootPath sets RootPath field to given value.

### HasRootPath

`func (o *CreateHypervisorResourcePoolGcpRequestModel) HasRootPath() bool`

HasRootPath returns a boolean if a field has been set.

### SetRootPathNil

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetRootPathNil(b bool)`

 SetRootPathNil sets the value for RootPath to be an explicit nil

### UnsetRootPath
`func (o *CreateHypervisorResourcePoolGcpRequestModel) UnsetRootPath()`

UnsetRootPath ensures that no value is present for RootPath, not even an explicit nil
### GetUseLocalStorageCaching

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetUseLocalStorageCaching() bool`

GetUseLocalStorageCaching returns the UseLocalStorageCaching field if non-nil, zero value otherwise.

### GetUseLocalStorageCachingOk

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetUseLocalStorageCachingOk() (*bool, bool)`

GetUseLocalStorageCachingOk returns a tuple with the UseLocalStorageCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocalStorageCaching

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetUseLocalStorageCaching(v bool)`

SetUseLocalStorageCaching sets UseLocalStorageCaching field to given value.

### HasUseLocalStorageCaching

`func (o *CreateHypervisorResourcePoolGcpRequestModel) HasUseLocalStorageCaching() bool`

HasUseLocalStorageCaching returns a boolean if a field has been set.

### SetUseLocalStorageCachingNil

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetUseLocalStorageCachingNil(b bool)`

 SetUseLocalStorageCachingNil sets the value for UseLocalStorageCaching to be an explicit nil

### UnsetUseLocalStorageCaching
`func (o *CreateHypervisorResourcePoolGcpRequestModel) UnsetUseLocalStorageCaching()`

UnsetUseLocalStorageCaching ensures that no value is present for UseLocalStorageCaching, not even an explicit nil
### GetStorage

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetStorage() []string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetStorageOk() (*[]string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetStorage(v []string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *CreateHypervisorResourcePoolGcpRequestModel) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *CreateHypervisorResourcePoolGcpRequestModel) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetPersonalvDiskStorage

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetPersonalvDiskStorage() []string`

GetPersonalvDiskStorage returns the PersonalvDiskStorage field if non-nil, zero value otherwise.

### GetPersonalvDiskStorageOk

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetPersonalvDiskStorageOk() (*[]string, bool)`

GetPersonalvDiskStorageOk returns a tuple with the PersonalvDiskStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalvDiskStorage

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetPersonalvDiskStorage(v []string)`

SetPersonalvDiskStorage sets PersonalvDiskStorage field to given value.

### HasPersonalvDiskStorage

`func (o *CreateHypervisorResourcePoolGcpRequestModel) HasPersonalvDiskStorage() bool`

HasPersonalvDiskStorage returns a boolean if a field has been set.

### SetPersonalvDiskStorageNil

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetPersonalvDiskStorageNil(b bool)`

 SetPersonalvDiskStorageNil sets the value for PersonalvDiskStorage to be an explicit nil

### UnsetPersonalvDiskStorage
`func (o *CreateHypervisorResourcePoolGcpRequestModel) UnsetPersonalvDiskStorage()`

UnsetPersonalvDiskStorage ensures that no value is present for PersonalvDiskStorage, not even an explicit nil
### GetTemporaryStorage

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetTemporaryStorage() []string`

GetTemporaryStorage returns the TemporaryStorage field if non-nil, zero value otherwise.

### GetTemporaryStorageOk

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetTemporaryStorageOk() (*[]string, bool)`

GetTemporaryStorageOk returns a tuple with the TemporaryStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryStorage

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetTemporaryStorage(v []string)`

SetTemporaryStorage sets TemporaryStorage field to given value.

### HasTemporaryStorage

`func (o *CreateHypervisorResourcePoolGcpRequestModel) HasTemporaryStorage() bool`

HasTemporaryStorage returns a boolean if a field has been set.

### SetTemporaryStorageNil

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetTemporaryStorageNil(b bool)`

 SetTemporaryStorageNil sets the value for TemporaryStorage to be an explicit nil

### UnsetTemporaryStorage
`func (o *CreateHypervisorResourcePoolGcpRequestModel) UnsetTemporaryStorage()`

UnsetTemporaryStorage ensures that no value is present for TemporaryStorage, not even an explicit nil
### GetCustomProperties

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *CreateHypervisorResourcePoolGcpRequestModel) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *CreateHypervisorResourcePoolGcpRequestModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *CreateHypervisorResourcePoolGcpRequestModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *CreateHypervisorResourcePoolGcpRequestModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


