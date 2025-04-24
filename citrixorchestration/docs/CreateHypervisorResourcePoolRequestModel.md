# CreateHypervisorResourcePoolRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the resource pool to create.  Required. | 
**VmTagging** | Pointer to **NullableBool** | Indicates whether VMs created by Virtual Apps &amp; Desktops provisioning operations should be tagged.  Tagged VMs are filtered out of queries by default. Optional.  Default is &#x60;true&#x60;. | [optional] [default to true]
**GpuTypes** | Pointer to **[]string** | Path to the GPU type resource(s) that are available for provisioning operations in this resource pool.  Optional.  Not supported by all hypervisor types. | [optional] 
**ConnectionType** | [**HypervisorConnectionType**](HypervisorConnectionType.md) |  | 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Metadata of the resource pool. Optional. | [optional] 
**VirtualPrivateCloud** | Pointer to **string** | AWS Virtual Private Cloud (VPC) resource which the resource pool is connected to.  Required. | [optional] 
**AvailabilityZone** | Pointer to **string** | Path to the availability zone resource to use for provisioning operations in this resource pool.  Required. | [optional] 
**Networks** | Pointer to **[]string** | Path to the network resource(s) that are available for provisioning operations in this resource pool.  At least one is required. | [optional] 
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

### NewCreateHypervisorResourcePoolRequestModel

`func NewCreateHypervisorResourcePoolRequestModel(name string, connectionType HypervisorConnectionType, ) *CreateHypervisorResourcePoolRequestModel`

NewCreateHypervisorResourcePoolRequestModel instantiates a new CreateHypervisorResourcePoolRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHypervisorResourcePoolRequestModelWithDefaults

`func NewCreateHypervisorResourcePoolRequestModelWithDefaults() *CreateHypervisorResourcePoolRequestModel`

NewCreateHypervisorResourcePoolRequestModelWithDefaults instantiates a new CreateHypervisorResourcePoolRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateHypervisorResourcePoolRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateHypervisorResourcePoolRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateHypervisorResourcePoolRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetVmTagging

`func (o *CreateHypervisorResourcePoolRequestModel) GetVmTagging() bool`

GetVmTagging returns the VmTagging field if non-nil, zero value otherwise.

### GetVmTaggingOk

`func (o *CreateHypervisorResourcePoolRequestModel) GetVmTaggingOk() (*bool, bool)`

GetVmTaggingOk returns a tuple with the VmTagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTagging

`func (o *CreateHypervisorResourcePoolRequestModel) SetVmTagging(v bool)`

SetVmTagging sets VmTagging field to given value.

### HasVmTagging

`func (o *CreateHypervisorResourcePoolRequestModel) HasVmTagging() bool`

HasVmTagging returns a boolean if a field has been set.

### SetVmTaggingNil

`func (o *CreateHypervisorResourcePoolRequestModel) SetVmTaggingNil(b bool)`

 SetVmTaggingNil sets the value for VmTagging to be an explicit nil

### UnsetVmTagging
`func (o *CreateHypervisorResourcePoolRequestModel) UnsetVmTagging()`

UnsetVmTagging ensures that no value is present for VmTagging, not even an explicit nil
### GetGpuTypes

`func (o *CreateHypervisorResourcePoolRequestModel) GetGpuTypes() []string`

GetGpuTypes returns the GpuTypes field if non-nil, zero value otherwise.

### GetGpuTypesOk

`func (o *CreateHypervisorResourcePoolRequestModel) GetGpuTypesOk() (*[]string, bool)`

GetGpuTypesOk returns a tuple with the GpuTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuTypes

`func (o *CreateHypervisorResourcePoolRequestModel) SetGpuTypes(v []string)`

SetGpuTypes sets GpuTypes field to given value.

### HasGpuTypes

`func (o *CreateHypervisorResourcePoolRequestModel) HasGpuTypes() bool`

HasGpuTypes returns a boolean if a field has been set.

### SetGpuTypesNil

`func (o *CreateHypervisorResourcePoolRequestModel) SetGpuTypesNil(b bool)`

 SetGpuTypesNil sets the value for GpuTypes to be an explicit nil

### UnsetGpuTypes
`func (o *CreateHypervisorResourcePoolRequestModel) UnsetGpuTypes()`

UnsetGpuTypes ensures that no value is present for GpuTypes, not even an explicit nil
### GetConnectionType

`func (o *CreateHypervisorResourcePoolRequestModel) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *CreateHypervisorResourcePoolRequestModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *CreateHypervisorResourcePoolRequestModel) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.


### GetMetadata

`func (o *CreateHypervisorResourcePoolRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateHypervisorResourcePoolRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateHypervisorResourcePoolRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateHypervisorResourcePoolRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CreateHypervisorResourcePoolRequestModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CreateHypervisorResourcePoolRequestModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetVirtualPrivateCloud

`func (o *CreateHypervisorResourcePoolRequestModel) GetVirtualPrivateCloud() string`

GetVirtualPrivateCloud returns the VirtualPrivateCloud field if non-nil, zero value otherwise.

### GetVirtualPrivateCloudOk

`func (o *CreateHypervisorResourcePoolRequestModel) GetVirtualPrivateCloudOk() (*string, bool)`

GetVirtualPrivateCloudOk returns a tuple with the VirtualPrivateCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPrivateCloud

`func (o *CreateHypervisorResourcePoolRequestModel) SetVirtualPrivateCloud(v string)`

SetVirtualPrivateCloud sets VirtualPrivateCloud field to given value.

### HasVirtualPrivateCloud

`func (o *CreateHypervisorResourcePoolRequestModel) HasVirtualPrivateCloud() bool`

HasVirtualPrivateCloud returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *CreateHypervisorResourcePoolRequestModel) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreateHypervisorResourcePoolRequestModel) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreateHypervisorResourcePoolRequestModel) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *CreateHypervisorResourcePoolRequestModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetNetworks

`func (o *CreateHypervisorResourcePoolRequestModel) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *CreateHypervisorResourcePoolRequestModel) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *CreateHypervisorResourcePoolRequestModel) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *CreateHypervisorResourcePoolRequestModel) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetRegion

`func (o *CreateHypervisorResourcePoolRequestModel) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateHypervisorResourcePoolRequestModel) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateHypervisorResourcePoolRequestModel) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CreateHypervisorResourcePoolRequestModel) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetVirtualNetwork

`func (o *CreateHypervisorResourcePoolRequestModel) GetVirtualNetwork() string`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *CreateHypervisorResourcePoolRequestModel) GetVirtualNetworkOk() (*string, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *CreateHypervisorResourcePoolRequestModel) SetVirtualNetwork(v string)`

SetVirtualNetwork sets VirtualNetwork field to given value.

### HasVirtualNetwork

`func (o *CreateHypervisorResourcePoolRequestModel) HasVirtualNetwork() bool`

HasVirtualNetwork returns a boolean if a field has been set.

### GetSubnets

`func (o *CreateHypervisorResourcePoolRequestModel) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *CreateHypervisorResourcePoolRequestModel) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *CreateHypervisorResourcePoolRequestModel) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *CreateHypervisorResourcePoolRequestModel) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetRootPath

`func (o *CreateHypervisorResourcePoolRequestModel) GetRootPath() string`

GetRootPath returns the RootPath field if non-nil, zero value otherwise.

### GetRootPathOk

`func (o *CreateHypervisorResourcePoolRequestModel) GetRootPathOk() (*string, bool)`

GetRootPathOk returns a tuple with the RootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPath

`func (o *CreateHypervisorResourcePoolRequestModel) SetRootPath(v string)`

SetRootPath sets RootPath field to given value.

### HasRootPath

`func (o *CreateHypervisorResourcePoolRequestModel) HasRootPath() bool`

HasRootPath returns a boolean if a field has been set.

### SetRootPathNil

`func (o *CreateHypervisorResourcePoolRequestModel) SetRootPathNil(b bool)`

 SetRootPathNil sets the value for RootPath to be an explicit nil

### UnsetRootPath
`func (o *CreateHypervisorResourcePoolRequestModel) UnsetRootPath()`

UnsetRootPath ensures that no value is present for RootPath, not even an explicit nil
### GetUseLocalStorageCaching

`func (o *CreateHypervisorResourcePoolRequestModel) GetUseLocalStorageCaching() bool`

GetUseLocalStorageCaching returns the UseLocalStorageCaching field if non-nil, zero value otherwise.

### GetUseLocalStorageCachingOk

`func (o *CreateHypervisorResourcePoolRequestModel) GetUseLocalStorageCachingOk() (*bool, bool)`

GetUseLocalStorageCachingOk returns a tuple with the UseLocalStorageCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocalStorageCaching

`func (o *CreateHypervisorResourcePoolRequestModel) SetUseLocalStorageCaching(v bool)`

SetUseLocalStorageCaching sets UseLocalStorageCaching field to given value.

### HasUseLocalStorageCaching

`func (o *CreateHypervisorResourcePoolRequestModel) HasUseLocalStorageCaching() bool`

HasUseLocalStorageCaching returns a boolean if a field has been set.

### SetUseLocalStorageCachingNil

`func (o *CreateHypervisorResourcePoolRequestModel) SetUseLocalStorageCachingNil(b bool)`

 SetUseLocalStorageCachingNil sets the value for UseLocalStorageCaching to be an explicit nil

### UnsetUseLocalStorageCaching
`func (o *CreateHypervisorResourcePoolRequestModel) UnsetUseLocalStorageCaching()`

UnsetUseLocalStorageCaching ensures that no value is present for UseLocalStorageCaching, not even an explicit nil
### GetStorage

`func (o *CreateHypervisorResourcePoolRequestModel) GetStorage() []string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CreateHypervisorResourcePoolRequestModel) GetStorageOk() (*[]string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CreateHypervisorResourcePoolRequestModel) SetStorage(v []string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *CreateHypervisorResourcePoolRequestModel) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *CreateHypervisorResourcePoolRequestModel) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *CreateHypervisorResourcePoolRequestModel) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetPersonalvDiskStorage

`func (o *CreateHypervisorResourcePoolRequestModel) GetPersonalvDiskStorage() []string`

GetPersonalvDiskStorage returns the PersonalvDiskStorage field if non-nil, zero value otherwise.

### GetPersonalvDiskStorageOk

`func (o *CreateHypervisorResourcePoolRequestModel) GetPersonalvDiskStorageOk() (*[]string, bool)`

GetPersonalvDiskStorageOk returns a tuple with the PersonalvDiskStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalvDiskStorage

`func (o *CreateHypervisorResourcePoolRequestModel) SetPersonalvDiskStorage(v []string)`

SetPersonalvDiskStorage sets PersonalvDiskStorage field to given value.

### HasPersonalvDiskStorage

`func (o *CreateHypervisorResourcePoolRequestModel) HasPersonalvDiskStorage() bool`

HasPersonalvDiskStorage returns a boolean if a field has been set.

### SetPersonalvDiskStorageNil

`func (o *CreateHypervisorResourcePoolRequestModel) SetPersonalvDiskStorageNil(b bool)`

 SetPersonalvDiskStorageNil sets the value for PersonalvDiskStorage to be an explicit nil

### UnsetPersonalvDiskStorage
`func (o *CreateHypervisorResourcePoolRequestModel) UnsetPersonalvDiskStorage()`

UnsetPersonalvDiskStorage ensures that no value is present for PersonalvDiskStorage, not even an explicit nil
### GetTemporaryStorage

`func (o *CreateHypervisorResourcePoolRequestModel) GetTemporaryStorage() []string`

GetTemporaryStorage returns the TemporaryStorage field if non-nil, zero value otherwise.

### GetTemporaryStorageOk

`func (o *CreateHypervisorResourcePoolRequestModel) GetTemporaryStorageOk() (*[]string, bool)`

GetTemporaryStorageOk returns a tuple with the TemporaryStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryStorage

`func (o *CreateHypervisorResourcePoolRequestModel) SetTemporaryStorage(v []string)`

SetTemporaryStorage sets TemporaryStorage field to given value.

### HasTemporaryStorage

`func (o *CreateHypervisorResourcePoolRequestModel) HasTemporaryStorage() bool`

HasTemporaryStorage returns a boolean if a field has been set.

### SetTemporaryStorageNil

`func (o *CreateHypervisorResourcePoolRequestModel) SetTemporaryStorageNil(b bool)`

 SetTemporaryStorageNil sets the value for TemporaryStorage to be an explicit nil

### UnsetTemporaryStorage
`func (o *CreateHypervisorResourcePoolRequestModel) UnsetTemporaryStorage()`

UnsetTemporaryStorage ensures that no value is present for TemporaryStorage, not even an explicit nil
### GetCustomProperties

`func (o *CreateHypervisorResourcePoolRequestModel) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *CreateHypervisorResourcePoolRequestModel) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *CreateHypervisorResourcePoolRequestModel) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *CreateHypervisorResourcePoolRequestModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *CreateHypervisorResourcePoolRequestModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *CreateHypervisorResourcePoolRequestModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
### GetStorageBalanceType

`func (o *CreateHypervisorResourcePoolRequestModel) GetStorageBalanceType() StorageBalanceType`

GetStorageBalanceType returns the StorageBalanceType field if non-nil, zero value otherwise.

### GetStorageBalanceTypeOk

`func (o *CreateHypervisorResourcePoolRequestModel) GetStorageBalanceTypeOk() (*StorageBalanceType, bool)`

GetStorageBalanceTypeOk returns a tuple with the StorageBalanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageBalanceType

`func (o *CreateHypervisorResourcePoolRequestModel) SetStorageBalanceType(v StorageBalanceType)`

SetStorageBalanceType sets StorageBalanceType field to given value.

### HasStorageBalanceType

`func (o *CreateHypervisorResourcePoolRequestModel) HasStorageBalanceType() bool`

HasStorageBalanceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


