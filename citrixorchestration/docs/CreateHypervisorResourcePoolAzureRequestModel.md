# CreateHypervisorResourcePoolAzureRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the resource pool to create.  Required. | 
**VmTagging** | Pointer to **bool** | Indicates whether VMs created by Virtual Apps &amp; Desktops provisioning operations should be tagged.  Tagged VMs are filtered out of queries by default. Optional.  Default is &#x60;true&#x60;. | [optional] [default to true]
**GpuTypes** | Pointer to **[]string** | Path to the GPU type resource(s) that are available for provisioning operations in this resource pool.  Optional.  Not supported by all hypervisor types. | [optional] 
**ConnectionType** | [**HypervisorConnectionType**](HypervisorConnectionType.md) |  | 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Metadata of the resource pool. Optional.  | [optional] 
**VirtualPrivateCloud** | Pointer to **string** | AWS Virtual Private Cloud (VPC) resource which the resource pool is connected to.  Required. | [optional] 
**AvailabilityZone** | Pointer to **string** | Path to the availability zone resource to use for provisioning operations in this resource pool.  Required. | [optional] 
**Networks** | Pointer to **[]string** | Path to the network resource(s) that are available for provisioning operations in this resource pool.  At least one is required. | [optional] 
**Region** | **string** | Azure region which the resource pool is connected to.  Required. | 
**VirtualNetwork** | **string** | Azure virtual network which the resource pool is connected to. Required. | 
**Subnets** | **[]string** | Path to the subnet(s) that are available for provisioning operations in this resource pool.  At least one is required. | 
**RootPath** | Pointer to **string** | Root path of the resources on the hypervisor which should be included in the resource pool.  Required. | [optional] 
**UseLocalStorageCaching** | Pointer to **bool** | Indicates whether local storage on the hypervisor will be used for caching purposes. Not all hypervisor types support this.  Defaults to &#x60;false&#x60;. | [optional] [default to false]
**Storage** | Pointer to **[]string** | Path to the storage resource(s) that are available for provisioning operations in this resource pool.  Required for some hypervisor types. | [optional] 
**PersonalvDiskStorage** | Pointer to **[]string** | Path to the storage resource(s) that are available for provisioning operations in this resource pool.  Required for some hypervisor types. | [optional] 
**TemporaryStorage** | Pointer to **[]string** | Path to the storage resource(s) that are used for temporary operations in this resource pool.  Required for some hypervisor types. | [optional] 
**CustomProperties** | Pointer to **string** | Custom properties.  Optional.  If not specified, will not be changed.  Only used for hypervisors of type Custom. | [optional] 

## Methods

### NewCreateHypervisorResourcePoolAzureRequestModel

`func NewCreateHypervisorResourcePoolAzureRequestModel(name string, connectionType HypervisorConnectionType, region string, virtualNetwork string, subnets []string, ) *CreateHypervisorResourcePoolAzureRequestModel`

NewCreateHypervisorResourcePoolAzureRequestModel instantiates a new CreateHypervisorResourcePoolAzureRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHypervisorResourcePoolAzureRequestModelWithDefaults

`func NewCreateHypervisorResourcePoolAzureRequestModelWithDefaults() *CreateHypervisorResourcePoolAzureRequestModel`

NewCreateHypervisorResourcePoolAzureRequestModelWithDefaults instantiates a new CreateHypervisorResourcePoolAzureRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateHypervisorResourcePoolAzureRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetVmTagging

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetVmTagging() bool`

GetVmTagging returns the VmTagging field if non-nil, zero value otherwise.

### GetVmTaggingOk

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetVmTaggingOk() (*bool, bool)`

GetVmTaggingOk returns a tuple with the VmTagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTagging

`func (o *CreateHypervisorResourcePoolAzureRequestModel) SetVmTagging(v bool)`

SetVmTagging sets VmTagging field to given value.

### HasVmTagging

`func (o *CreateHypervisorResourcePoolAzureRequestModel) HasVmTagging() bool`

HasVmTagging returns a boolean if a field has been set.

### GetGpuTypes

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetGpuTypes() []string`

GetGpuTypes returns the GpuTypes field if non-nil, zero value otherwise.

### GetGpuTypesOk

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetGpuTypesOk() (*[]string, bool)`

GetGpuTypesOk returns a tuple with the GpuTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuTypes

`func (o *CreateHypervisorResourcePoolAzureRequestModel) SetGpuTypes(v []string)`

SetGpuTypes sets GpuTypes field to given value.

### HasGpuTypes

`func (o *CreateHypervisorResourcePoolAzureRequestModel) HasGpuTypes() bool`

HasGpuTypes returns a boolean if a field has been set.

### GetConnectionType

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *CreateHypervisorResourcePoolAzureRequestModel) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.


### GetMetadata

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateHypervisorResourcePoolAzureRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateHypervisorResourcePoolAzureRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetVirtualPrivateCloud

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetVirtualPrivateCloud() string`

GetVirtualPrivateCloud returns the VirtualPrivateCloud field if non-nil, zero value otherwise.

### GetVirtualPrivateCloudOk

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetVirtualPrivateCloudOk() (*string, bool)`

GetVirtualPrivateCloudOk returns a tuple with the VirtualPrivateCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPrivateCloud

`func (o *CreateHypervisorResourcePoolAzureRequestModel) SetVirtualPrivateCloud(v string)`

SetVirtualPrivateCloud sets VirtualPrivateCloud field to given value.

### HasVirtualPrivateCloud

`func (o *CreateHypervisorResourcePoolAzureRequestModel) HasVirtualPrivateCloud() bool`

HasVirtualPrivateCloud returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreateHypervisorResourcePoolAzureRequestModel) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *CreateHypervisorResourcePoolAzureRequestModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetNetworks

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *CreateHypervisorResourcePoolAzureRequestModel) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *CreateHypervisorResourcePoolAzureRequestModel) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetRegion

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateHypervisorResourcePoolAzureRequestModel) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetVirtualNetwork

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetVirtualNetwork() string`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetVirtualNetworkOk() (*string, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *CreateHypervisorResourcePoolAzureRequestModel) SetVirtualNetwork(v string)`

SetVirtualNetwork sets VirtualNetwork field to given value.


### GetSubnets

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *CreateHypervisorResourcePoolAzureRequestModel) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.


### GetRootPath

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetRootPath() string`

GetRootPath returns the RootPath field if non-nil, zero value otherwise.

### GetRootPathOk

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetRootPathOk() (*string, bool)`

GetRootPathOk returns a tuple with the RootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPath

`func (o *CreateHypervisorResourcePoolAzureRequestModel) SetRootPath(v string)`

SetRootPath sets RootPath field to given value.

### HasRootPath

`func (o *CreateHypervisorResourcePoolAzureRequestModel) HasRootPath() bool`

HasRootPath returns a boolean if a field has been set.

### GetUseLocalStorageCaching

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetUseLocalStorageCaching() bool`

GetUseLocalStorageCaching returns the UseLocalStorageCaching field if non-nil, zero value otherwise.

### GetUseLocalStorageCachingOk

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetUseLocalStorageCachingOk() (*bool, bool)`

GetUseLocalStorageCachingOk returns a tuple with the UseLocalStorageCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocalStorageCaching

`func (o *CreateHypervisorResourcePoolAzureRequestModel) SetUseLocalStorageCaching(v bool)`

SetUseLocalStorageCaching sets UseLocalStorageCaching field to given value.

### HasUseLocalStorageCaching

`func (o *CreateHypervisorResourcePoolAzureRequestModel) HasUseLocalStorageCaching() bool`

HasUseLocalStorageCaching returns a boolean if a field has been set.

### GetStorage

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetStorage() []string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetStorageOk() (*[]string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CreateHypervisorResourcePoolAzureRequestModel) SetStorage(v []string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *CreateHypervisorResourcePoolAzureRequestModel) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetPersonalvDiskStorage

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetPersonalvDiskStorage() []string`

GetPersonalvDiskStorage returns the PersonalvDiskStorage field if non-nil, zero value otherwise.

### GetPersonalvDiskStorageOk

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetPersonalvDiskStorageOk() (*[]string, bool)`

GetPersonalvDiskStorageOk returns a tuple with the PersonalvDiskStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalvDiskStorage

`func (o *CreateHypervisorResourcePoolAzureRequestModel) SetPersonalvDiskStorage(v []string)`

SetPersonalvDiskStorage sets PersonalvDiskStorage field to given value.

### HasPersonalvDiskStorage

`func (o *CreateHypervisorResourcePoolAzureRequestModel) HasPersonalvDiskStorage() bool`

HasPersonalvDiskStorage returns a boolean if a field has been set.

### GetTemporaryStorage

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetTemporaryStorage() []string`

GetTemporaryStorage returns the TemporaryStorage field if non-nil, zero value otherwise.

### GetTemporaryStorageOk

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetTemporaryStorageOk() (*[]string, bool)`

GetTemporaryStorageOk returns a tuple with the TemporaryStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryStorage

`func (o *CreateHypervisorResourcePoolAzureRequestModel) SetTemporaryStorage(v []string)`

SetTemporaryStorage sets TemporaryStorage field to given value.

### HasTemporaryStorage

`func (o *CreateHypervisorResourcePoolAzureRequestModel) HasTemporaryStorage() bool`

HasTemporaryStorage returns a boolean if a field has been set.

### GetCustomProperties

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *CreateHypervisorResourcePoolAzureRequestModel) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *CreateHypervisorResourcePoolAzureRequestModel) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *CreateHypervisorResourcePoolAzureRequestModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


