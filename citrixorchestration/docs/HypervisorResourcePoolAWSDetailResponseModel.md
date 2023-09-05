# HypervisorResourcePoolAWSDetailResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | [**HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) |  | 
**VirtualNetwork** | [**HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) |  | 
**Subnets** | [**[]HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) | List of subnets in the VirtualNetwork that may be used within the resource pool. | 
**Project** | [**HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) |  | 
**RootPath** | Pointer to [**HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) |  | [optional] 
**Storage** | [**[]HypervisorStorageResourceResponseModel**](HypervisorStorageResourceResponseModel.md) | List of hypervisor-connected storage in the resource pool that is used for OS disks of virtual machines. | 
**TemporaryStorage** | [**[]HypervisorStorageResourceResponseModel**](HypervisorStorageResourceResponseModel.md) | List of hypervisor-connected storage in the resource pool that is used for temporary data storage for virtual machines. | 
**UseLocalStorageCaching** | Pointer to **bool** | Indicates whether virtual machines created within this resource pool will use local storage caching for their disk images. | [optional] 
**CustomProperties** | Pointer to **string** | Custom properties.  Optional.  If not specified, will not be changed.  Only used for hypervisors of type Custom. | [optional] 
**PersonalvDiskStorage** | Pointer to [**[]HypervisorStorageResourceResponseModel**](HypervisorStorageResourceResponseModel.md) | List of hypervisor-connected storage in the resource pool that is used for personal v disk data storage for virtual machines. | [optional] 
**GpuTypes** | Pointer to [**[]HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) | GPU types available in the resource pool.  Only applicable to hypervisors that support GPU types. | [optional] 
**ConnectionType** | [**HypervisorConnectionType**](HypervisorConnectionType.md) |  | 
**UsesExplicitStorage** | Pointer to **bool** | If the hypervisor resource pool use ExplicitStorage. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Metadata for hypervisor resource pool.  | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**XDPath** | Pointer to **string** | XenApp &amp; XenDesktop path to the resource on the hypervisor.  An example value is: &#x60;XDHyp:\\Connections\\{{hypervisor name}}\\{{vm name}}.vm\\{{snapshot name}}.snapshot&#x60; or &#x60;XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}&#x60; | [optional] 
**HypervisorConnection** | [**HypervisorResourcePoolResponseModelAllOfHypervisorConnection**](HypervisorResourcePoolResponseModelAllOfHypervisorConnection.md) |  | 
**DefaultNetwork** | [**HypervisorResourcePoolResponseModelAllOfDefaultNetwork**](HypervisorResourcePoolResponseModelAllOfDefaultNetwork.md) |  | 
**VMTaggingEnabled** | **bool** | Indicates whether new virtual machines are tagged with metadata from the hypervisor. | 
**ResourcePoolRootPath** | Pointer to **string** | Hypervisor resourcePool RootPath. | [optional] 
**ResourcePoolRootId** | Pointer to **string** | Hypervisor resourcePool RootId. | [optional] 
**VirtualPrivateCloud** | [**HypervisorResourcePoolAWSDetailResponseModelAllOfVirtualPrivateCloud**](HypervisorResourcePoolAWSDetailResponseModelAllOfVirtualPrivateCloud.md) |  | 
**AvailabilityZone** | [**HypervisorResourcePoolAWSDetailResponseModelAllOfAvailabilityZone**](HypervisorResourcePoolAWSDetailResponseModelAllOfAvailabilityZone.md) |  | 
**Networks** | [**[]HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) | List of networks that may be used within the resource pool. | 

## Methods

### NewHypervisorResourcePoolAWSDetailResponseModel

`func NewHypervisorResourcePoolAWSDetailResponseModel(region HypervisorResourceRefResponseModel, virtualNetwork HypervisorResourceRefResponseModel, subnets []HypervisorResourceRefResponseModel, project HypervisorResourceRefResponseModel, storage []HypervisorStorageResourceResponseModel, temporaryStorage []HypervisorStorageResourceResponseModel, connectionType HypervisorConnectionType, hypervisorConnection HypervisorResourcePoolResponseModelAllOfHypervisorConnection, defaultNetwork HypervisorResourcePoolResponseModelAllOfDefaultNetwork, vMTaggingEnabled bool, virtualPrivateCloud HypervisorResourcePoolAWSDetailResponseModelAllOfVirtualPrivateCloud, availabilityZone HypervisorResourcePoolAWSDetailResponseModelAllOfAvailabilityZone, networks []HypervisorResourceRefResponseModel, ) *HypervisorResourcePoolAWSDetailResponseModel`

NewHypervisorResourcePoolAWSDetailResponseModel instantiates a new HypervisorResourcePoolAWSDetailResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorResourcePoolAWSDetailResponseModelWithDefaults

`func NewHypervisorResourcePoolAWSDetailResponseModelWithDefaults() *HypervisorResourcePoolAWSDetailResponseModel`

NewHypervisorResourcePoolAWSDetailResponseModelWithDefaults instantiates a new HypervisorResourcePoolAWSDetailResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetRegion() HypervisorResourceRefResponseModel`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetRegionOk() (*HypervisorResourceRefResponseModel, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetRegion(v HypervisorResourceRefResponseModel)`

SetRegion sets Region field to given value.


### GetVirtualNetwork

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetVirtualNetwork() HypervisorResourceRefResponseModel`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetVirtualNetworkOk() (*HypervisorResourceRefResponseModel, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetVirtualNetwork(v HypervisorResourceRefResponseModel)`

SetVirtualNetwork sets VirtualNetwork field to given value.


### GetSubnets

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetSubnets() []HypervisorResourceRefResponseModel`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetSubnetsOk() (*[]HypervisorResourceRefResponseModel, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetSubnets(v []HypervisorResourceRefResponseModel)`

SetSubnets sets Subnets field to given value.


### GetProject

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetProject() HypervisorResourceRefResponseModel`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetProjectOk() (*HypervisorResourceRefResponseModel, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetProject(v HypervisorResourceRefResponseModel)`

SetProject sets Project field to given value.


### GetRootPath

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetRootPath() HypervisorResourceRefResponseModel`

GetRootPath returns the RootPath field if non-nil, zero value otherwise.

### GetRootPathOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetRootPathOk() (*HypervisorResourceRefResponseModel, bool)`

GetRootPathOk returns a tuple with the RootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPath

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetRootPath(v HypervisorResourceRefResponseModel)`

SetRootPath sets RootPath field to given value.

### HasRootPath

`func (o *HypervisorResourcePoolAWSDetailResponseModel) HasRootPath() bool`

HasRootPath returns a boolean if a field has been set.

### GetStorage

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetStorage() []HypervisorStorageResourceResponseModel`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetStorageOk() (*[]HypervisorStorageResourceResponseModel, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetStorage(v []HypervisorStorageResourceResponseModel)`

SetStorage sets Storage field to given value.


### GetTemporaryStorage

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetTemporaryStorage() []HypervisorStorageResourceResponseModel`

GetTemporaryStorage returns the TemporaryStorage field if non-nil, zero value otherwise.

### GetTemporaryStorageOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetTemporaryStorageOk() (*[]HypervisorStorageResourceResponseModel, bool)`

GetTemporaryStorageOk returns a tuple with the TemporaryStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryStorage

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetTemporaryStorage(v []HypervisorStorageResourceResponseModel)`

SetTemporaryStorage sets TemporaryStorage field to given value.


### GetUseLocalStorageCaching

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetUseLocalStorageCaching() bool`

GetUseLocalStorageCaching returns the UseLocalStorageCaching field if non-nil, zero value otherwise.

### GetUseLocalStorageCachingOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetUseLocalStorageCachingOk() (*bool, bool)`

GetUseLocalStorageCachingOk returns a tuple with the UseLocalStorageCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocalStorageCaching

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetUseLocalStorageCaching(v bool)`

SetUseLocalStorageCaching sets UseLocalStorageCaching field to given value.

### HasUseLocalStorageCaching

`func (o *HypervisorResourcePoolAWSDetailResponseModel) HasUseLocalStorageCaching() bool`

HasUseLocalStorageCaching returns a boolean if a field has been set.

### GetCustomProperties

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *HypervisorResourcePoolAWSDetailResponseModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetPersonalvDiskStorage

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetPersonalvDiskStorage() []HypervisorStorageResourceResponseModel`

GetPersonalvDiskStorage returns the PersonalvDiskStorage field if non-nil, zero value otherwise.

### GetPersonalvDiskStorageOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetPersonalvDiskStorageOk() (*[]HypervisorStorageResourceResponseModel, bool)`

GetPersonalvDiskStorageOk returns a tuple with the PersonalvDiskStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalvDiskStorage

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetPersonalvDiskStorage(v []HypervisorStorageResourceResponseModel)`

SetPersonalvDiskStorage sets PersonalvDiskStorage field to given value.

### HasPersonalvDiskStorage

`func (o *HypervisorResourcePoolAWSDetailResponseModel) HasPersonalvDiskStorage() bool`

HasPersonalvDiskStorage returns a boolean if a field has been set.

### GetGpuTypes

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetGpuTypes() []HypervisorResourceRefResponseModel`

GetGpuTypes returns the GpuTypes field if non-nil, zero value otherwise.

### GetGpuTypesOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetGpuTypesOk() (*[]HypervisorResourceRefResponseModel, bool)`

GetGpuTypesOk returns a tuple with the GpuTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuTypes

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetGpuTypes(v []HypervisorResourceRefResponseModel)`

SetGpuTypes sets GpuTypes field to given value.

### HasGpuTypes

`func (o *HypervisorResourcePoolAWSDetailResponseModel) HasGpuTypes() bool`

HasGpuTypes returns a boolean if a field has been set.

### GetConnectionType

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.


### GetUsesExplicitStorage

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetUsesExplicitStorage() bool`

GetUsesExplicitStorage returns the UsesExplicitStorage field if non-nil, zero value otherwise.

### GetUsesExplicitStorageOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetUsesExplicitStorageOk() (*bool, bool)`

GetUsesExplicitStorageOk returns a tuple with the UsesExplicitStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesExplicitStorage

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetUsesExplicitStorage(v bool)`

SetUsesExplicitStorage sets UsesExplicitStorage field to given value.

### HasUsesExplicitStorage

`func (o *HypervisorResourcePoolAWSDetailResponseModel) HasUsesExplicitStorage() bool`

HasUsesExplicitStorage returns a boolean if a field has been set.

### GetMetadata

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HypervisorResourcePoolAWSDetailResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HypervisorResourcePoolAWSDetailResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervisorResourcePoolAWSDetailResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetXDPath

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetXDPath() string`

GetXDPath returns the XDPath field if non-nil, zero value otherwise.

### GetXDPathOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetXDPathOk() (*string, bool)`

GetXDPathOk returns a tuple with the XDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXDPath

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetXDPath(v string)`

SetXDPath sets XDPath field to given value.

### HasXDPath

`func (o *HypervisorResourcePoolAWSDetailResponseModel) HasXDPath() bool`

HasXDPath returns a boolean if a field has been set.

### GetHypervisorConnection

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetHypervisorConnection() HypervisorResourcePoolResponseModelAllOfHypervisorConnection`

GetHypervisorConnection returns the HypervisorConnection field if non-nil, zero value otherwise.

### GetHypervisorConnectionOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetHypervisorConnectionOk() (*HypervisorResourcePoolResponseModelAllOfHypervisorConnection, bool)`

GetHypervisorConnectionOk returns a tuple with the HypervisorConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorConnection

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetHypervisorConnection(v HypervisorResourcePoolResponseModelAllOfHypervisorConnection)`

SetHypervisorConnection sets HypervisorConnection field to given value.


### GetDefaultNetwork

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetDefaultNetwork() HypervisorResourcePoolResponseModelAllOfDefaultNetwork`

GetDefaultNetwork returns the DefaultNetwork field if non-nil, zero value otherwise.

### GetDefaultNetworkOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetDefaultNetworkOk() (*HypervisorResourcePoolResponseModelAllOfDefaultNetwork, bool)`

GetDefaultNetworkOk returns a tuple with the DefaultNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetwork

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetDefaultNetwork(v HypervisorResourcePoolResponseModelAllOfDefaultNetwork)`

SetDefaultNetwork sets DefaultNetwork field to given value.


### GetVMTaggingEnabled

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetVMTaggingEnabled() bool`

GetVMTaggingEnabled returns the VMTaggingEnabled field if non-nil, zero value otherwise.

### GetVMTaggingEnabledOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetVMTaggingEnabledOk() (*bool, bool)`

GetVMTaggingEnabledOk returns a tuple with the VMTaggingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMTaggingEnabled

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetVMTaggingEnabled(v bool)`

SetVMTaggingEnabled sets VMTaggingEnabled field to given value.


### GetResourcePoolRootPath

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetResourcePoolRootPath() string`

GetResourcePoolRootPath returns the ResourcePoolRootPath field if non-nil, zero value otherwise.

### GetResourcePoolRootPathOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetResourcePoolRootPathOk() (*string, bool)`

GetResourcePoolRootPathOk returns a tuple with the ResourcePoolRootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolRootPath

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetResourcePoolRootPath(v string)`

SetResourcePoolRootPath sets ResourcePoolRootPath field to given value.

### HasResourcePoolRootPath

`func (o *HypervisorResourcePoolAWSDetailResponseModel) HasResourcePoolRootPath() bool`

HasResourcePoolRootPath returns a boolean if a field has been set.

### GetResourcePoolRootId

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetResourcePoolRootId() string`

GetResourcePoolRootId returns the ResourcePoolRootId field if non-nil, zero value otherwise.

### GetResourcePoolRootIdOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetResourcePoolRootIdOk() (*string, bool)`

GetResourcePoolRootIdOk returns a tuple with the ResourcePoolRootId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolRootId

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetResourcePoolRootId(v string)`

SetResourcePoolRootId sets ResourcePoolRootId field to given value.

### HasResourcePoolRootId

`func (o *HypervisorResourcePoolAWSDetailResponseModel) HasResourcePoolRootId() bool`

HasResourcePoolRootId returns a boolean if a field has been set.

### GetVirtualPrivateCloud

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetVirtualPrivateCloud() HypervisorResourcePoolAWSDetailResponseModelAllOfVirtualPrivateCloud`

GetVirtualPrivateCloud returns the VirtualPrivateCloud field if non-nil, zero value otherwise.

### GetVirtualPrivateCloudOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetVirtualPrivateCloudOk() (*HypervisorResourcePoolAWSDetailResponseModelAllOfVirtualPrivateCloud, bool)`

GetVirtualPrivateCloudOk returns a tuple with the VirtualPrivateCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPrivateCloud

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetVirtualPrivateCloud(v HypervisorResourcePoolAWSDetailResponseModelAllOfVirtualPrivateCloud)`

SetVirtualPrivateCloud sets VirtualPrivateCloud field to given value.


### GetAvailabilityZone

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetAvailabilityZone() HypervisorResourcePoolAWSDetailResponseModelAllOfAvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetAvailabilityZoneOk() (*HypervisorResourcePoolAWSDetailResponseModelAllOfAvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetAvailabilityZone(v HypervisorResourcePoolAWSDetailResponseModelAllOfAvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetNetworks

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetNetworks() []HypervisorResourceRefResponseModel`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *HypervisorResourcePoolAWSDetailResponseModel) GetNetworksOk() (*[]HypervisorResourceRefResponseModel, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *HypervisorResourcePoolAWSDetailResponseModel) SetNetworks(v []HypervisorResourceRefResponseModel)`

SetNetworks sets Networks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


