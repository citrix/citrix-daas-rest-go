# CreateHypervisorResourcePoolAWSRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the resource pool to create.  Required. | 
**VmTagging** | Pointer to **NullableBool** | Indicates whether VMs created by Virtual Apps &amp; Desktops provisioning operations should be tagged.  Tagged VMs are filtered out of queries by default. Optional.  Default is &#x60;true&#x60;. | [optional] [default to true]
**GpuTypes** | Pointer to **[]string** | Path to the GPU type resource(s) that are available for provisioning operations in this resource pool.  Optional.  Not supported by all hypervisor types. | [optional] 
**ConnectionType** | [**HypervisorConnectionType**](HypervisorConnectionType.md) |  | 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Metadata of the resource pool. Optional.  | [optional] 
**VirtualPrivateCloud** | **string** | AWS Virtual Private Cloud (VPC) resource which the resource pool is connected to.  Required. | 
**AvailabilityZone** | **string** | Path to the availability zone resource to use for provisioning operations in this resource pool.  Required. | 
**Networks** | **[]string** | Path to the network resource(s) that are available for provisioning operations in this resource pool.  At least one is required. | 
**Region** | Pointer to **string** | Azure region which the resource pool is connected to.  Required. | [optional] 
**VirtualNetwork** | Pointer to **string** | Azure virtual network which the resource pool is connected to. Required. | [optional] 
**Subnets** | Pointer to **[]string** | Path to the subnet(s) that are available for provisioning operations in this resource pool.  At least one is required. | [optional] 
**RootPath** | Pointer to **NullableString** | Root path of the resources on the hypervisor which should be included in the resource pool.  Required. | [optional] 
**UseLocalStorageCaching** | Pointer to **NullableBool** | Indicates whether local storage on the hypervisor will be used for caching purposes. Not all hypervisor types support this.  Defaults to &#x60;false&#x60;. | [optional] [default to false]
**Storage** | Pointer to **[]string** | Path to the storage resource(s) that are available for provisioning operations in this resource pool.  Required for some hypervisor types. | [optional] 
**PersonalvDiskStorage** | Pointer to **[]string** | Path to the storage resource(s) that are available for provisioning operations in this resource pool.  Required for some hypervisor types. | [optional] 
**TemporaryStorage** | Pointer to **[]string** | Path to the storage resource(s) that are used for temporary operations in this resource pool.  Required for some hypervisor types. | [optional] 
**CustomProperties** | Pointer to **NullableString** | Custom properties.  Optional.  If not specified, will not be changed.  Only used for hypervisors of type Custom. | [optional] 
**StorageBalanceType** | Pointer to [**StorageBalanceType**](StorageBalanceType.md) |  | [optional] 

## Methods

### NewCreateHypervisorResourcePoolAWSRequestModel

`func NewCreateHypervisorResourcePoolAWSRequestModel(name string, connectionType HypervisorConnectionType, virtualPrivateCloud string, availabilityZone string, networks []string, ) *CreateHypervisorResourcePoolAWSRequestModel`

NewCreateHypervisorResourcePoolAWSRequestModel instantiates a new CreateHypervisorResourcePoolAWSRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHypervisorResourcePoolAWSRequestModelWithDefaults

`func NewCreateHypervisorResourcePoolAWSRequestModelWithDefaults() *CreateHypervisorResourcePoolAWSRequestModel`

NewCreateHypervisorResourcePoolAWSRequestModelWithDefaults instantiates a new CreateHypervisorResourcePoolAWSRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetVmTagging

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetVmTagging() bool`

GetVmTagging returns the VmTagging field if non-nil, zero value otherwise.

### GetVmTaggingOk

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetVmTaggingOk() (*bool, bool)`

GetVmTaggingOk returns a tuple with the VmTagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTagging

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetVmTagging(v bool)`

SetVmTagging sets VmTagging field to given value.

### HasVmTagging

`func (o *CreateHypervisorResourcePoolAWSRequestModel) HasVmTagging() bool`

HasVmTagging returns a boolean if a field has been set.

### SetVmTaggingNil

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetVmTaggingNil(b bool)`

 SetVmTaggingNil sets the value for VmTagging to be an explicit nil

### UnsetVmTagging
`func (o *CreateHypervisorResourcePoolAWSRequestModel) UnsetVmTagging()`

UnsetVmTagging ensures that no value is present for VmTagging, not even an explicit nil
### GetGpuTypes

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetGpuTypes() []string`

GetGpuTypes returns the GpuTypes field if non-nil, zero value otherwise.

### GetGpuTypesOk

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetGpuTypesOk() (*[]string, bool)`

GetGpuTypesOk returns a tuple with the GpuTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuTypes

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetGpuTypes(v []string)`

SetGpuTypes sets GpuTypes field to given value.

### HasGpuTypes

`func (o *CreateHypervisorResourcePoolAWSRequestModel) HasGpuTypes() bool`

HasGpuTypes returns a boolean if a field has been set.

### SetGpuTypesNil

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetGpuTypesNil(b bool)`

 SetGpuTypesNil sets the value for GpuTypes to be an explicit nil

### UnsetGpuTypes
`func (o *CreateHypervisorResourcePoolAWSRequestModel) UnsetGpuTypes()`

UnsetGpuTypes ensures that no value is present for GpuTypes, not even an explicit nil
### GetConnectionType

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.


### GetMetadata

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateHypervisorResourcePoolAWSRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CreateHypervisorResourcePoolAWSRequestModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetVirtualPrivateCloud

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetVirtualPrivateCloud() string`

GetVirtualPrivateCloud returns the VirtualPrivateCloud field if non-nil, zero value otherwise.

### GetVirtualPrivateCloudOk

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetVirtualPrivateCloudOk() (*string, bool)`

GetVirtualPrivateCloudOk returns a tuple with the VirtualPrivateCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPrivateCloud

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetVirtualPrivateCloud(v string)`

SetVirtualPrivateCloud sets VirtualPrivateCloud field to given value.


### GetAvailabilityZone

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetNetworks

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.


### GetRegion

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CreateHypervisorResourcePoolAWSRequestModel) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetVirtualNetwork

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetVirtualNetwork() string`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetVirtualNetworkOk() (*string, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetVirtualNetwork(v string)`

SetVirtualNetwork sets VirtualNetwork field to given value.

### HasVirtualNetwork

`func (o *CreateHypervisorResourcePoolAWSRequestModel) HasVirtualNetwork() bool`

HasVirtualNetwork returns a boolean if a field has been set.

### GetSubnets

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *CreateHypervisorResourcePoolAWSRequestModel) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetRootPath

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetRootPath() string`

GetRootPath returns the RootPath field if non-nil, zero value otherwise.

### GetRootPathOk

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetRootPathOk() (*string, bool)`

GetRootPathOk returns a tuple with the RootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPath

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetRootPath(v string)`

SetRootPath sets RootPath field to given value.

### HasRootPath

`func (o *CreateHypervisorResourcePoolAWSRequestModel) HasRootPath() bool`

HasRootPath returns a boolean if a field has been set.

### SetRootPathNil

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetRootPathNil(b bool)`

 SetRootPathNil sets the value for RootPath to be an explicit nil

### UnsetRootPath
`func (o *CreateHypervisorResourcePoolAWSRequestModel) UnsetRootPath()`

UnsetRootPath ensures that no value is present for RootPath, not even an explicit nil
### GetUseLocalStorageCaching

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetUseLocalStorageCaching() bool`

GetUseLocalStorageCaching returns the UseLocalStorageCaching field if non-nil, zero value otherwise.

### GetUseLocalStorageCachingOk

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetUseLocalStorageCachingOk() (*bool, bool)`

GetUseLocalStorageCachingOk returns a tuple with the UseLocalStorageCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocalStorageCaching

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetUseLocalStorageCaching(v bool)`

SetUseLocalStorageCaching sets UseLocalStorageCaching field to given value.

### HasUseLocalStorageCaching

`func (o *CreateHypervisorResourcePoolAWSRequestModel) HasUseLocalStorageCaching() bool`

HasUseLocalStorageCaching returns a boolean if a field has been set.

### SetUseLocalStorageCachingNil

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetUseLocalStorageCachingNil(b bool)`

 SetUseLocalStorageCachingNil sets the value for UseLocalStorageCaching to be an explicit nil

### UnsetUseLocalStorageCaching
`func (o *CreateHypervisorResourcePoolAWSRequestModel) UnsetUseLocalStorageCaching()`

UnsetUseLocalStorageCaching ensures that no value is present for UseLocalStorageCaching, not even an explicit nil
### GetStorage

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetStorage() []string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetStorageOk() (*[]string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetStorage(v []string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *CreateHypervisorResourcePoolAWSRequestModel) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *CreateHypervisorResourcePoolAWSRequestModel) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetPersonalvDiskStorage

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetPersonalvDiskStorage() []string`

GetPersonalvDiskStorage returns the PersonalvDiskStorage field if non-nil, zero value otherwise.

### GetPersonalvDiskStorageOk

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetPersonalvDiskStorageOk() (*[]string, bool)`

GetPersonalvDiskStorageOk returns a tuple with the PersonalvDiskStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalvDiskStorage

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetPersonalvDiskStorage(v []string)`

SetPersonalvDiskStorage sets PersonalvDiskStorage field to given value.

### HasPersonalvDiskStorage

`func (o *CreateHypervisorResourcePoolAWSRequestModel) HasPersonalvDiskStorage() bool`

HasPersonalvDiskStorage returns a boolean if a field has been set.

### SetPersonalvDiskStorageNil

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetPersonalvDiskStorageNil(b bool)`

 SetPersonalvDiskStorageNil sets the value for PersonalvDiskStorage to be an explicit nil

### UnsetPersonalvDiskStorage
`func (o *CreateHypervisorResourcePoolAWSRequestModel) UnsetPersonalvDiskStorage()`

UnsetPersonalvDiskStorage ensures that no value is present for PersonalvDiskStorage, not even an explicit nil
### GetTemporaryStorage

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetTemporaryStorage() []string`

GetTemporaryStorage returns the TemporaryStorage field if non-nil, zero value otherwise.

### GetTemporaryStorageOk

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetTemporaryStorageOk() (*[]string, bool)`

GetTemporaryStorageOk returns a tuple with the TemporaryStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryStorage

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetTemporaryStorage(v []string)`

SetTemporaryStorage sets TemporaryStorage field to given value.

### HasTemporaryStorage

`func (o *CreateHypervisorResourcePoolAWSRequestModel) HasTemporaryStorage() bool`

HasTemporaryStorage returns a boolean if a field has been set.

### SetTemporaryStorageNil

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetTemporaryStorageNil(b bool)`

 SetTemporaryStorageNil sets the value for TemporaryStorage to be an explicit nil

### UnsetTemporaryStorage
`func (o *CreateHypervisorResourcePoolAWSRequestModel) UnsetTemporaryStorage()`

UnsetTemporaryStorage ensures that no value is present for TemporaryStorage, not even an explicit nil
### GetCustomProperties

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *CreateHypervisorResourcePoolAWSRequestModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *CreateHypervisorResourcePoolAWSRequestModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
### GetStorageBalanceType

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetStorageBalanceType() StorageBalanceType`

GetStorageBalanceType returns the StorageBalanceType field if non-nil, zero value otherwise.

### GetStorageBalanceTypeOk

`func (o *CreateHypervisorResourcePoolAWSRequestModel) GetStorageBalanceTypeOk() (*StorageBalanceType, bool)`

GetStorageBalanceTypeOk returns a tuple with the StorageBalanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageBalanceType

`func (o *CreateHypervisorResourcePoolAWSRequestModel) SetStorageBalanceType(v StorageBalanceType)`

SetStorageBalanceType sets StorageBalanceType field to given value.

### HasStorageBalanceType

`func (o *CreateHypervisorResourcePoolAWSRequestModel) HasStorageBalanceType() bool`

HasStorageBalanceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


