# HypervisorResourcePoolDetailResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualPrivateCloud** | [**HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) |  | 
**AvailabilityZone** | [**HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) |  | 
**Networks** | [**[]HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) | List of networks that may be used within the resource pool. | 
**Region** | [**HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) |  | 
**VirtualNetwork** | [**HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) |  | 
**Subnets** | [**[]HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) | List of subnets in the VirtualNetwork that may be used within the resource pool. | 
**Project** | [**HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) |  | 
**RootPath** | Pointer to [**HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) |  | [optional] 
**Storage** | [**[]HypervisorStorageResourceResponseModel**](HypervisorStorageResourceResponseModel.md) | List of hypervisor-connected storage in the resource pool that is used for OS disks of virtual machines. | 
**TemporaryStorage** | [**[]HypervisorStorageResourceResponseModel**](HypervisorStorageResourceResponseModel.md) | List of hypervisor-connected storage in the resource pool that is used for temporary data storage for virtual machines. | 
**UseLocalStorageCaching** | Pointer to **NullableBool** | Indicates whether virtual machines created within this resource pool will use local storage caching for their disk images. | [optional] 
**CustomProperties** | Pointer to **NullableString** | Custom properties.  Optional.  If not specified, will not be changed.  Only used for hypervisors of type Custom. | [optional] 
**PersonalvDiskStorage** | Pointer to [**[]HypervisorStorageResourceResponseModel**](HypervisorStorageResourceResponseModel.md) | List of hypervisor-connected storage in the resource pool that is used for personal v disk data storage for virtual machines. | [optional] 
**StorageBalanceType** | Pointer to [**StorageBalanceType**](StorageBalanceType.md) |  | [optional] 
**Id** | Pointer to **NullableString** | Id of the resource. | [optional] 
**Name** | Pointer to **NullableString** | Name of the resource. | [optional] 
**XDPath** | Pointer to **NullableString** | XenApp &amp; XenDesktop path to the resource on the hypervisor.  An example value is: &#x60;XDHyp:\\Connections\\{{hypervisor name}}\\{{vm name}}.vm\\{{snapshot name}}.snapshot&#x60; or &#x60;XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}&#x60; | [optional] 
**HypervisorConnection** | [**RefResponseModel**](RefResponseModel.md) |  | 
**ConnectionType** | [**HypervisorConnectionType**](HypervisorConnectionType.md) |  | 
**DefaultNetwork** | [**HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) |  | 
**VMTaggingEnabled** | **bool** | Indicates whether new virtual machines are tagged with metadata from the hypervisor. | 
**ResourcePoolRootPath** | Pointer to **NullableString** | Hypervisor resourcePool RootPath. | [optional] 
**ResourcePoolRootId** | Pointer to **NullableString** | Hypervisor resourcePool RootId. | [optional] 
**GpuTypes** | Pointer to [**[]HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) | GPU types available in the resource pool.  Only applicable to hypervisors that support GPU types. | [optional] 
**UsesExplicitStorage** | Pointer to **bool** | If the hypervisor resource pool use ExplicitStorage. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Metadata for hypervisor resource pool.  | [optional] 

## Methods

### NewHypervisorResourcePoolDetailResponseModel

`func NewHypervisorResourcePoolDetailResponseModel(virtualPrivateCloud HypervisorResourceRefResponseModel, availabilityZone HypervisorResourceRefResponseModel, networks []HypervisorResourceRefResponseModel, region HypervisorResourceRefResponseModel, virtualNetwork HypervisorResourceRefResponseModel, subnets []HypervisorResourceRefResponseModel, project HypervisorResourceRefResponseModel, storage []HypervisorStorageResourceResponseModel, temporaryStorage []HypervisorStorageResourceResponseModel, hypervisorConnection RefResponseModel, connectionType HypervisorConnectionType, defaultNetwork HypervisorResourceRefResponseModel, vMTaggingEnabled bool, ) *HypervisorResourcePoolDetailResponseModel`

NewHypervisorResourcePoolDetailResponseModel instantiates a new HypervisorResourcePoolDetailResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorResourcePoolDetailResponseModelWithDefaults

`func NewHypervisorResourcePoolDetailResponseModelWithDefaults() *HypervisorResourcePoolDetailResponseModel`

NewHypervisorResourcePoolDetailResponseModelWithDefaults instantiates a new HypervisorResourcePoolDetailResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualPrivateCloud

`func (o *HypervisorResourcePoolDetailResponseModel) GetVirtualPrivateCloud() HypervisorResourceRefResponseModel`

GetVirtualPrivateCloud returns the VirtualPrivateCloud field if non-nil, zero value otherwise.

### GetVirtualPrivateCloudOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetVirtualPrivateCloudOk() (*HypervisorResourceRefResponseModel, bool)`

GetVirtualPrivateCloudOk returns a tuple with the VirtualPrivateCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPrivateCloud

`func (o *HypervisorResourcePoolDetailResponseModel) SetVirtualPrivateCloud(v HypervisorResourceRefResponseModel)`

SetVirtualPrivateCloud sets VirtualPrivateCloud field to given value.


### GetAvailabilityZone

`func (o *HypervisorResourcePoolDetailResponseModel) GetAvailabilityZone() HypervisorResourceRefResponseModel`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetAvailabilityZoneOk() (*HypervisorResourceRefResponseModel, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *HypervisorResourcePoolDetailResponseModel) SetAvailabilityZone(v HypervisorResourceRefResponseModel)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetNetworks

`func (o *HypervisorResourcePoolDetailResponseModel) GetNetworks() []HypervisorResourceRefResponseModel`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetNetworksOk() (*[]HypervisorResourceRefResponseModel, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *HypervisorResourcePoolDetailResponseModel) SetNetworks(v []HypervisorResourceRefResponseModel)`

SetNetworks sets Networks field to given value.


### GetRegion

`func (o *HypervisorResourcePoolDetailResponseModel) GetRegion() HypervisorResourceRefResponseModel`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetRegionOk() (*HypervisorResourceRefResponseModel, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HypervisorResourcePoolDetailResponseModel) SetRegion(v HypervisorResourceRefResponseModel)`

SetRegion sets Region field to given value.


### GetVirtualNetwork

`func (o *HypervisorResourcePoolDetailResponseModel) GetVirtualNetwork() HypervisorResourceRefResponseModel`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetVirtualNetworkOk() (*HypervisorResourceRefResponseModel, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *HypervisorResourcePoolDetailResponseModel) SetVirtualNetwork(v HypervisorResourceRefResponseModel)`

SetVirtualNetwork sets VirtualNetwork field to given value.


### GetSubnets

`func (o *HypervisorResourcePoolDetailResponseModel) GetSubnets() []HypervisorResourceRefResponseModel`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetSubnetsOk() (*[]HypervisorResourceRefResponseModel, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *HypervisorResourcePoolDetailResponseModel) SetSubnets(v []HypervisorResourceRefResponseModel)`

SetSubnets sets Subnets field to given value.


### GetProject

`func (o *HypervisorResourcePoolDetailResponseModel) GetProject() HypervisorResourceRefResponseModel`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetProjectOk() (*HypervisorResourceRefResponseModel, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *HypervisorResourcePoolDetailResponseModel) SetProject(v HypervisorResourceRefResponseModel)`

SetProject sets Project field to given value.


### GetRootPath

`func (o *HypervisorResourcePoolDetailResponseModel) GetRootPath() HypervisorResourceRefResponseModel`

GetRootPath returns the RootPath field if non-nil, zero value otherwise.

### GetRootPathOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetRootPathOk() (*HypervisorResourceRefResponseModel, bool)`

GetRootPathOk returns a tuple with the RootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPath

`func (o *HypervisorResourcePoolDetailResponseModel) SetRootPath(v HypervisorResourceRefResponseModel)`

SetRootPath sets RootPath field to given value.

### HasRootPath

`func (o *HypervisorResourcePoolDetailResponseModel) HasRootPath() bool`

HasRootPath returns a boolean if a field has been set.

### GetStorage

`func (o *HypervisorResourcePoolDetailResponseModel) GetStorage() []HypervisorStorageResourceResponseModel`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetStorageOk() (*[]HypervisorStorageResourceResponseModel, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *HypervisorResourcePoolDetailResponseModel) SetStorage(v []HypervisorStorageResourceResponseModel)`

SetStorage sets Storage field to given value.


### GetTemporaryStorage

`func (o *HypervisorResourcePoolDetailResponseModel) GetTemporaryStorage() []HypervisorStorageResourceResponseModel`

GetTemporaryStorage returns the TemporaryStorage field if non-nil, zero value otherwise.

### GetTemporaryStorageOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetTemporaryStorageOk() (*[]HypervisorStorageResourceResponseModel, bool)`

GetTemporaryStorageOk returns a tuple with the TemporaryStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryStorage

`func (o *HypervisorResourcePoolDetailResponseModel) SetTemporaryStorage(v []HypervisorStorageResourceResponseModel)`

SetTemporaryStorage sets TemporaryStorage field to given value.


### GetUseLocalStorageCaching

`func (o *HypervisorResourcePoolDetailResponseModel) GetUseLocalStorageCaching() bool`

GetUseLocalStorageCaching returns the UseLocalStorageCaching field if non-nil, zero value otherwise.

### GetUseLocalStorageCachingOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetUseLocalStorageCachingOk() (*bool, bool)`

GetUseLocalStorageCachingOk returns a tuple with the UseLocalStorageCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocalStorageCaching

`func (o *HypervisorResourcePoolDetailResponseModel) SetUseLocalStorageCaching(v bool)`

SetUseLocalStorageCaching sets UseLocalStorageCaching field to given value.

### HasUseLocalStorageCaching

`func (o *HypervisorResourcePoolDetailResponseModel) HasUseLocalStorageCaching() bool`

HasUseLocalStorageCaching returns a boolean if a field has been set.

### SetUseLocalStorageCachingNil

`func (o *HypervisorResourcePoolDetailResponseModel) SetUseLocalStorageCachingNil(b bool)`

 SetUseLocalStorageCachingNil sets the value for UseLocalStorageCaching to be an explicit nil

### UnsetUseLocalStorageCaching
`func (o *HypervisorResourcePoolDetailResponseModel) UnsetUseLocalStorageCaching()`

UnsetUseLocalStorageCaching ensures that no value is present for UseLocalStorageCaching, not even an explicit nil
### GetCustomProperties

`func (o *HypervisorResourcePoolDetailResponseModel) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *HypervisorResourcePoolDetailResponseModel) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *HypervisorResourcePoolDetailResponseModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *HypervisorResourcePoolDetailResponseModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *HypervisorResourcePoolDetailResponseModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
### GetPersonalvDiskStorage

`func (o *HypervisorResourcePoolDetailResponseModel) GetPersonalvDiskStorage() []HypervisorStorageResourceResponseModel`

GetPersonalvDiskStorage returns the PersonalvDiskStorage field if non-nil, zero value otherwise.

### GetPersonalvDiskStorageOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetPersonalvDiskStorageOk() (*[]HypervisorStorageResourceResponseModel, bool)`

GetPersonalvDiskStorageOk returns a tuple with the PersonalvDiskStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalvDiskStorage

`func (o *HypervisorResourcePoolDetailResponseModel) SetPersonalvDiskStorage(v []HypervisorStorageResourceResponseModel)`

SetPersonalvDiskStorage sets PersonalvDiskStorage field to given value.

### HasPersonalvDiskStorage

`func (o *HypervisorResourcePoolDetailResponseModel) HasPersonalvDiskStorage() bool`

HasPersonalvDiskStorage returns a boolean if a field has been set.

### SetPersonalvDiskStorageNil

`func (o *HypervisorResourcePoolDetailResponseModel) SetPersonalvDiskStorageNil(b bool)`

 SetPersonalvDiskStorageNil sets the value for PersonalvDiskStorage to be an explicit nil

### UnsetPersonalvDiskStorage
`func (o *HypervisorResourcePoolDetailResponseModel) UnsetPersonalvDiskStorage()`

UnsetPersonalvDiskStorage ensures that no value is present for PersonalvDiskStorage, not even an explicit nil
### GetStorageBalanceType

`func (o *HypervisorResourcePoolDetailResponseModel) GetStorageBalanceType() StorageBalanceType`

GetStorageBalanceType returns the StorageBalanceType field if non-nil, zero value otherwise.

### GetStorageBalanceTypeOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetStorageBalanceTypeOk() (*StorageBalanceType, bool)`

GetStorageBalanceTypeOk returns a tuple with the StorageBalanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageBalanceType

`func (o *HypervisorResourcePoolDetailResponseModel) SetStorageBalanceType(v StorageBalanceType)`

SetStorageBalanceType sets StorageBalanceType field to given value.

### HasStorageBalanceType

`func (o *HypervisorResourcePoolDetailResponseModel) HasStorageBalanceType() bool`

HasStorageBalanceType returns a boolean if a field has been set.

### GetId

`func (o *HypervisorResourcePoolDetailResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HypervisorResourcePoolDetailResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HypervisorResourcePoolDetailResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HypervisorResourcePoolDetailResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HypervisorResourcePoolDetailResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *HypervisorResourcePoolDetailResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorResourcePoolDetailResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervisorResourcePoolDetailResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HypervisorResourcePoolDetailResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HypervisorResourcePoolDetailResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetXDPath

`func (o *HypervisorResourcePoolDetailResponseModel) GetXDPath() string`

GetXDPath returns the XDPath field if non-nil, zero value otherwise.

### GetXDPathOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetXDPathOk() (*string, bool)`

GetXDPathOk returns a tuple with the XDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXDPath

`func (o *HypervisorResourcePoolDetailResponseModel) SetXDPath(v string)`

SetXDPath sets XDPath field to given value.

### HasXDPath

`func (o *HypervisorResourcePoolDetailResponseModel) HasXDPath() bool`

HasXDPath returns a boolean if a field has been set.

### SetXDPathNil

`func (o *HypervisorResourcePoolDetailResponseModel) SetXDPathNil(b bool)`

 SetXDPathNil sets the value for XDPath to be an explicit nil

### UnsetXDPath
`func (o *HypervisorResourcePoolDetailResponseModel) UnsetXDPath()`

UnsetXDPath ensures that no value is present for XDPath, not even an explicit nil
### GetHypervisorConnection

`func (o *HypervisorResourcePoolDetailResponseModel) GetHypervisorConnection() RefResponseModel`

GetHypervisorConnection returns the HypervisorConnection field if non-nil, zero value otherwise.

### GetHypervisorConnectionOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetHypervisorConnectionOk() (*RefResponseModel, bool)`

GetHypervisorConnectionOk returns a tuple with the HypervisorConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorConnection

`func (o *HypervisorResourcePoolDetailResponseModel) SetHypervisorConnection(v RefResponseModel)`

SetHypervisorConnection sets HypervisorConnection field to given value.


### GetConnectionType

`func (o *HypervisorResourcePoolDetailResponseModel) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *HypervisorResourcePoolDetailResponseModel) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.


### GetDefaultNetwork

`func (o *HypervisorResourcePoolDetailResponseModel) GetDefaultNetwork() HypervisorResourceRefResponseModel`

GetDefaultNetwork returns the DefaultNetwork field if non-nil, zero value otherwise.

### GetDefaultNetworkOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetDefaultNetworkOk() (*HypervisorResourceRefResponseModel, bool)`

GetDefaultNetworkOk returns a tuple with the DefaultNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetwork

`func (o *HypervisorResourcePoolDetailResponseModel) SetDefaultNetwork(v HypervisorResourceRefResponseModel)`

SetDefaultNetwork sets DefaultNetwork field to given value.


### GetVMTaggingEnabled

`func (o *HypervisorResourcePoolDetailResponseModel) GetVMTaggingEnabled() bool`

GetVMTaggingEnabled returns the VMTaggingEnabled field if non-nil, zero value otherwise.

### GetVMTaggingEnabledOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetVMTaggingEnabledOk() (*bool, bool)`

GetVMTaggingEnabledOk returns a tuple with the VMTaggingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMTaggingEnabled

`func (o *HypervisorResourcePoolDetailResponseModel) SetVMTaggingEnabled(v bool)`

SetVMTaggingEnabled sets VMTaggingEnabled field to given value.


### GetResourcePoolRootPath

`func (o *HypervisorResourcePoolDetailResponseModel) GetResourcePoolRootPath() string`

GetResourcePoolRootPath returns the ResourcePoolRootPath field if non-nil, zero value otherwise.

### GetResourcePoolRootPathOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetResourcePoolRootPathOk() (*string, bool)`

GetResourcePoolRootPathOk returns a tuple with the ResourcePoolRootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolRootPath

`func (o *HypervisorResourcePoolDetailResponseModel) SetResourcePoolRootPath(v string)`

SetResourcePoolRootPath sets ResourcePoolRootPath field to given value.

### HasResourcePoolRootPath

`func (o *HypervisorResourcePoolDetailResponseModel) HasResourcePoolRootPath() bool`

HasResourcePoolRootPath returns a boolean if a field has been set.

### SetResourcePoolRootPathNil

`func (o *HypervisorResourcePoolDetailResponseModel) SetResourcePoolRootPathNil(b bool)`

 SetResourcePoolRootPathNil sets the value for ResourcePoolRootPath to be an explicit nil

### UnsetResourcePoolRootPath
`func (o *HypervisorResourcePoolDetailResponseModel) UnsetResourcePoolRootPath()`

UnsetResourcePoolRootPath ensures that no value is present for ResourcePoolRootPath, not even an explicit nil
### GetResourcePoolRootId

`func (o *HypervisorResourcePoolDetailResponseModel) GetResourcePoolRootId() string`

GetResourcePoolRootId returns the ResourcePoolRootId field if non-nil, zero value otherwise.

### GetResourcePoolRootIdOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetResourcePoolRootIdOk() (*string, bool)`

GetResourcePoolRootIdOk returns a tuple with the ResourcePoolRootId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolRootId

`func (o *HypervisorResourcePoolDetailResponseModel) SetResourcePoolRootId(v string)`

SetResourcePoolRootId sets ResourcePoolRootId field to given value.

### HasResourcePoolRootId

`func (o *HypervisorResourcePoolDetailResponseModel) HasResourcePoolRootId() bool`

HasResourcePoolRootId returns a boolean if a field has been set.

### SetResourcePoolRootIdNil

`func (o *HypervisorResourcePoolDetailResponseModel) SetResourcePoolRootIdNil(b bool)`

 SetResourcePoolRootIdNil sets the value for ResourcePoolRootId to be an explicit nil

### UnsetResourcePoolRootId
`func (o *HypervisorResourcePoolDetailResponseModel) UnsetResourcePoolRootId()`

UnsetResourcePoolRootId ensures that no value is present for ResourcePoolRootId, not even an explicit nil
### GetGpuTypes

`func (o *HypervisorResourcePoolDetailResponseModel) GetGpuTypes() []HypervisorResourceRefResponseModel`

GetGpuTypes returns the GpuTypes field if non-nil, zero value otherwise.

### GetGpuTypesOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetGpuTypesOk() (*[]HypervisorResourceRefResponseModel, bool)`

GetGpuTypesOk returns a tuple with the GpuTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuTypes

`func (o *HypervisorResourcePoolDetailResponseModel) SetGpuTypes(v []HypervisorResourceRefResponseModel)`

SetGpuTypes sets GpuTypes field to given value.

### HasGpuTypes

`func (o *HypervisorResourcePoolDetailResponseModel) HasGpuTypes() bool`

HasGpuTypes returns a boolean if a field has been set.

### SetGpuTypesNil

`func (o *HypervisorResourcePoolDetailResponseModel) SetGpuTypesNil(b bool)`

 SetGpuTypesNil sets the value for GpuTypes to be an explicit nil

### UnsetGpuTypes
`func (o *HypervisorResourcePoolDetailResponseModel) UnsetGpuTypes()`

UnsetGpuTypes ensures that no value is present for GpuTypes, not even an explicit nil
### GetUsesExplicitStorage

`func (o *HypervisorResourcePoolDetailResponseModel) GetUsesExplicitStorage() bool`

GetUsesExplicitStorage returns the UsesExplicitStorage field if non-nil, zero value otherwise.

### GetUsesExplicitStorageOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetUsesExplicitStorageOk() (*bool, bool)`

GetUsesExplicitStorageOk returns a tuple with the UsesExplicitStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesExplicitStorage

`func (o *HypervisorResourcePoolDetailResponseModel) SetUsesExplicitStorage(v bool)`

SetUsesExplicitStorage sets UsesExplicitStorage field to given value.

### HasUsesExplicitStorage

`func (o *HypervisorResourcePoolDetailResponseModel) HasUsesExplicitStorage() bool`

HasUsesExplicitStorage returns a boolean if a field has been set.

### GetMetadata

`func (o *HypervisorResourcePoolDetailResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HypervisorResourcePoolDetailResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HypervisorResourcePoolDetailResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HypervisorResourcePoolDetailResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *HypervisorResourcePoolDetailResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *HypervisorResourcePoolDetailResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


