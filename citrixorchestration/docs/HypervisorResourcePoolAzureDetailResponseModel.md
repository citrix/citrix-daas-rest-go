# HypervisorResourcePoolAzureDetailResponseModel

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
**ContainerScopes** | Pointer to [**[]ContainerScopeResponseModel**](ContainerScopeResponseModel.md) | Delegated admin scopes in which the containers of the resource pool reside. | [optional] 

## Methods

### NewHypervisorResourcePoolAzureDetailResponseModel

`func NewHypervisorResourcePoolAzureDetailResponseModel(virtualPrivateCloud HypervisorResourceRefResponseModel, availabilityZone HypervisorResourceRefResponseModel, networks []HypervisorResourceRefResponseModel, region HypervisorResourceRefResponseModel, virtualNetwork HypervisorResourceRefResponseModel, subnets []HypervisorResourceRefResponseModel, project HypervisorResourceRefResponseModel, storage []HypervisorStorageResourceResponseModel, temporaryStorage []HypervisorStorageResourceResponseModel, hypervisorConnection RefResponseModel, connectionType HypervisorConnectionType, defaultNetwork HypervisorResourceRefResponseModel, vMTaggingEnabled bool, ) *HypervisorResourcePoolAzureDetailResponseModel`

NewHypervisorResourcePoolAzureDetailResponseModel instantiates a new HypervisorResourcePoolAzureDetailResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorResourcePoolAzureDetailResponseModelWithDefaults

`func NewHypervisorResourcePoolAzureDetailResponseModelWithDefaults() *HypervisorResourcePoolAzureDetailResponseModel`

NewHypervisorResourcePoolAzureDetailResponseModelWithDefaults instantiates a new HypervisorResourcePoolAzureDetailResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualPrivateCloud

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetVirtualPrivateCloud() HypervisorResourceRefResponseModel`

GetVirtualPrivateCloud returns the VirtualPrivateCloud field if non-nil, zero value otherwise.

### GetVirtualPrivateCloudOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetVirtualPrivateCloudOk() (*HypervisorResourceRefResponseModel, bool)`

GetVirtualPrivateCloudOk returns a tuple with the VirtualPrivateCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPrivateCloud

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetVirtualPrivateCloud(v HypervisorResourceRefResponseModel)`

SetVirtualPrivateCloud sets VirtualPrivateCloud field to given value.


### GetAvailabilityZone

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetAvailabilityZone() HypervisorResourceRefResponseModel`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetAvailabilityZoneOk() (*HypervisorResourceRefResponseModel, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetAvailabilityZone(v HypervisorResourceRefResponseModel)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetNetworks

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetNetworks() []HypervisorResourceRefResponseModel`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetNetworksOk() (*[]HypervisorResourceRefResponseModel, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetNetworks(v []HypervisorResourceRefResponseModel)`

SetNetworks sets Networks field to given value.


### GetRegion

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetRegion() HypervisorResourceRefResponseModel`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetRegionOk() (*HypervisorResourceRefResponseModel, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetRegion(v HypervisorResourceRefResponseModel)`

SetRegion sets Region field to given value.


### GetVirtualNetwork

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetVirtualNetwork() HypervisorResourceRefResponseModel`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetVirtualNetworkOk() (*HypervisorResourceRefResponseModel, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetVirtualNetwork(v HypervisorResourceRefResponseModel)`

SetVirtualNetwork sets VirtualNetwork field to given value.


### GetSubnets

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetSubnets() []HypervisorResourceRefResponseModel`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetSubnetsOk() (*[]HypervisorResourceRefResponseModel, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetSubnets(v []HypervisorResourceRefResponseModel)`

SetSubnets sets Subnets field to given value.


### GetProject

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetProject() HypervisorResourceRefResponseModel`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetProjectOk() (*HypervisorResourceRefResponseModel, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetProject(v HypervisorResourceRefResponseModel)`

SetProject sets Project field to given value.


### GetRootPath

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetRootPath() HypervisorResourceRefResponseModel`

GetRootPath returns the RootPath field if non-nil, zero value otherwise.

### GetRootPathOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetRootPathOk() (*HypervisorResourceRefResponseModel, bool)`

GetRootPathOk returns a tuple with the RootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPath

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetRootPath(v HypervisorResourceRefResponseModel)`

SetRootPath sets RootPath field to given value.

### HasRootPath

`func (o *HypervisorResourcePoolAzureDetailResponseModel) HasRootPath() bool`

HasRootPath returns a boolean if a field has been set.

### GetStorage

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetStorage() []HypervisorStorageResourceResponseModel`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetStorageOk() (*[]HypervisorStorageResourceResponseModel, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetStorage(v []HypervisorStorageResourceResponseModel)`

SetStorage sets Storage field to given value.


### GetTemporaryStorage

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetTemporaryStorage() []HypervisorStorageResourceResponseModel`

GetTemporaryStorage returns the TemporaryStorage field if non-nil, zero value otherwise.

### GetTemporaryStorageOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetTemporaryStorageOk() (*[]HypervisorStorageResourceResponseModel, bool)`

GetTemporaryStorageOk returns a tuple with the TemporaryStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryStorage

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetTemporaryStorage(v []HypervisorStorageResourceResponseModel)`

SetTemporaryStorage sets TemporaryStorage field to given value.


### GetUseLocalStorageCaching

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetUseLocalStorageCaching() bool`

GetUseLocalStorageCaching returns the UseLocalStorageCaching field if non-nil, zero value otherwise.

### GetUseLocalStorageCachingOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetUseLocalStorageCachingOk() (*bool, bool)`

GetUseLocalStorageCachingOk returns a tuple with the UseLocalStorageCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocalStorageCaching

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetUseLocalStorageCaching(v bool)`

SetUseLocalStorageCaching sets UseLocalStorageCaching field to given value.

### HasUseLocalStorageCaching

`func (o *HypervisorResourcePoolAzureDetailResponseModel) HasUseLocalStorageCaching() bool`

HasUseLocalStorageCaching returns a boolean if a field has been set.

### SetUseLocalStorageCachingNil

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetUseLocalStorageCachingNil(b bool)`

 SetUseLocalStorageCachingNil sets the value for UseLocalStorageCaching to be an explicit nil

### UnsetUseLocalStorageCaching
`func (o *HypervisorResourcePoolAzureDetailResponseModel) UnsetUseLocalStorageCaching()`

UnsetUseLocalStorageCaching ensures that no value is present for UseLocalStorageCaching, not even an explicit nil
### GetCustomProperties

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *HypervisorResourcePoolAzureDetailResponseModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *HypervisorResourcePoolAzureDetailResponseModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
### GetPersonalvDiskStorage

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetPersonalvDiskStorage() []HypervisorStorageResourceResponseModel`

GetPersonalvDiskStorage returns the PersonalvDiskStorage field if non-nil, zero value otherwise.

### GetPersonalvDiskStorageOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetPersonalvDiskStorageOk() (*[]HypervisorStorageResourceResponseModel, bool)`

GetPersonalvDiskStorageOk returns a tuple with the PersonalvDiskStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalvDiskStorage

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetPersonalvDiskStorage(v []HypervisorStorageResourceResponseModel)`

SetPersonalvDiskStorage sets PersonalvDiskStorage field to given value.

### HasPersonalvDiskStorage

`func (o *HypervisorResourcePoolAzureDetailResponseModel) HasPersonalvDiskStorage() bool`

HasPersonalvDiskStorage returns a boolean if a field has been set.

### SetPersonalvDiskStorageNil

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetPersonalvDiskStorageNil(b bool)`

 SetPersonalvDiskStorageNil sets the value for PersonalvDiskStorage to be an explicit nil

### UnsetPersonalvDiskStorage
`func (o *HypervisorResourcePoolAzureDetailResponseModel) UnsetPersonalvDiskStorage()`

UnsetPersonalvDiskStorage ensures that no value is present for PersonalvDiskStorage, not even an explicit nil
### GetStorageBalanceType

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetStorageBalanceType() StorageBalanceType`

GetStorageBalanceType returns the StorageBalanceType field if non-nil, zero value otherwise.

### GetStorageBalanceTypeOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetStorageBalanceTypeOk() (*StorageBalanceType, bool)`

GetStorageBalanceTypeOk returns a tuple with the StorageBalanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageBalanceType

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetStorageBalanceType(v StorageBalanceType)`

SetStorageBalanceType sets StorageBalanceType field to given value.

### HasStorageBalanceType

`func (o *HypervisorResourcePoolAzureDetailResponseModel) HasStorageBalanceType() bool`

HasStorageBalanceType returns a boolean if a field has been set.

### GetId

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HypervisorResourcePoolAzureDetailResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HypervisorResourcePoolAzureDetailResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervisorResourcePoolAzureDetailResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HypervisorResourcePoolAzureDetailResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetXDPath

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetXDPath() string`

GetXDPath returns the XDPath field if non-nil, zero value otherwise.

### GetXDPathOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetXDPathOk() (*string, bool)`

GetXDPathOk returns a tuple with the XDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXDPath

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetXDPath(v string)`

SetXDPath sets XDPath field to given value.

### HasXDPath

`func (o *HypervisorResourcePoolAzureDetailResponseModel) HasXDPath() bool`

HasXDPath returns a boolean if a field has been set.

### SetXDPathNil

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetXDPathNil(b bool)`

 SetXDPathNil sets the value for XDPath to be an explicit nil

### UnsetXDPath
`func (o *HypervisorResourcePoolAzureDetailResponseModel) UnsetXDPath()`

UnsetXDPath ensures that no value is present for XDPath, not even an explicit nil
### GetHypervisorConnection

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetHypervisorConnection() RefResponseModel`

GetHypervisorConnection returns the HypervisorConnection field if non-nil, zero value otherwise.

### GetHypervisorConnectionOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetHypervisorConnectionOk() (*RefResponseModel, bool)`

GetHypervisorConnectionOk returns a tuple with the HypervisorConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorConnection

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetHypervisorConnection(v RefResponseModel)`

SetHypervisorConnection sets HypervisorConnection field to given value.


### GetConnectionType

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.


### GetDefaultNetwork

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetDefaultNetwork() HypervisorResourceRefResponseModel`

GetDefaultNetwork returns the DefaultNetwork field if non-nil, zero value otherwise.

### GetDefaultNetworkOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetDefaultNetworkOk() (*HypervisorResourceRefResponseModel, bool)`

GetDefaultNetworkOk returns a tuple with the DefaultNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetwork

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetDefaultNetwork(v HypervisorResourceRefResponseModel)`

SetDefaultNetwork sets DefaultNetwork field to given value.


### GetVMTaggingEnabled

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetVMTaggingEnabled() bool`

GetVMTaggingEnabled returns the VMTaggingEnabled field if non-nil, zero value otherwise.

### GetVMTaggingEnabledOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetVMTaggingEnabledOk() (*bool, bool)`

GetVMTaggingEnabledOk returns a tuple with the VMTaggingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMTaggingEnabled

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetVMTaggingEnabled(v bool)`

SetVMTaggingEnabled sets VMTaggingEnabled field to given value.


### GetResourcePoolRootPath

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetResourcePoolRootPath() string`

GetResourcePoolRootPath returns the ResourcePoolRootPath field if non-nil, zero value otherwise.

### GetResourcePoolRootPathOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetResourcePoolRootPathOk() (*string, bool)`

GetResourcePoolRootPathOk returns a tuple with the ResourcePoolRootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolRootPath

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetResourcePoolRootPath(v string)`

SetResourcePoolRootPath sets ResourcePoolRootPath field to given value.

### HasResourcePoolRootPath

`func (o *HypervisorResourcePoolAzureDetailResponseModel) HasResourcePoolRootPath() bool`

HasResourcePoolRootPath returns a boolean if a field has been set.

### SetResourcePoolRootPathNil

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetResourcePoolRootPathNil(b bool)`

 SetResourcePoolRootPathNil sets the value for ResourcePoolRootPath to be an explicit nil

### UnsetResourcePoolRootPath
`func (o *HypervisorResourcePoolAzureDetailResponseModel) UnsetResourcePoolRootPath()`

UnsetResourcePoolRootPath ensures that no value is present for ResourcePoolRootPath, not even an explicit nil
### GetResourcePoolRootId

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetResourcePoolRootId() string`

GetResourcePoolRootId returns the ResourcePoolRootId field if non-nil, zero value otherwise.

### GetResourcePoolRootIdOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetResourcePoolRootIdOk() (*string, bool)`

GetResourcePoolRootIdOk returns a tuple with the ResourcePoolRootId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolRootId

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetResourcePoolRootId(v string)`

SetResourcePoolRootId sets ResourcePoolRootId field to given value.

### HasResourcePoolRootId

`func (o *HypervisorResourcePoolAzureDetailResponseModel) HasResourcePoolRootId() bool`

HasResourcePoolRootId returns a boolean if a field has been set.

### SetResourcePoolRootIdNil

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetResourcePoolRootIdNil(b bool)`

 SetResourcePoolRootIdNil sets the value for ResourcePoolRootId to be an explicit nil

### UnsetResourcePoolRootId
`func (o *HypervisorResourcePoolAzureDetailResponseModel) UnsetResourcePoolRootId()`

UnsetResourcePoolRootId ensures that no value is present for ResourcePoolRootId, not even an explicit nil
### GetGpuTypes

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetGpuTypes() []HypervisorResourceRefResponseModel`

GetGpuTypes returns the GpuTypes field if non-nil, zero value otherwise.

### GetGpuTypesOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetGpuTypesOk() (*[]HypervisorResourceRefResponseModel, bool)`

GetGpuTypesOk returns a tuple with the GpuTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuTypes

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetGpuTypes(v []HypervisorResourceRefResponseModel)`

SetGpuTypes sets GpuTypes field to given value.

### HasGpuTypes

`func (o *HypervisorResourcePoolAzureDetailResponseModel) HasGpuTypes() bool`

HasGpuTypes returns a boolean if a field has been set.

### SetGpuTypesNil

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetGpuTypesNil(b bool)`

 SetGpuTypesNil sets the value for GpuTypes to be an explicit nil

### UnsetGpuTypes
`func (o *HypervisorResourcePoolAzureDetailResponseModel) UnsetGpuTypes()`

UnsetGpuTypes ensures that no value is present for GpuTypes, not even an explicit nil
### GetUsesExplicitStorage

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetUsesExplicitStorage() bool`

GetUsesExplicitStorage returns the UsesExplicitStorage field if non-nil, zero value otherwise.

### GetUsesExplicitStorageOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetUsesExplicitStorageOk() (*bool, bool)`

GetUsesExplicitStorageOk returns a tuple with the UsesExplicitStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesExplicitStorage

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetUsesExplicitStorage(v bool)`

SetUsesExplicitStorage sets UsesExplicitStorage field to given value.

### HasUsesExplicitStorage

`func (o *HypervisorResourcePoolAzureDetailResponseModel) HasUsesExplicitStorage() bool`

HasUsesExplicitStorage returns a boolean if a field has been set.

### GetMetadata

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HypervisorResourcePoolAzureDetailResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *HypervisorResourcePoolAzureDetailResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetContainerScopes

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetContainerScopes() []ContainerScopeResponseModel`

GetContainerScopes returns the ContainerScopes field if non-nil, zero value otherwise.

### GetContainerScopesOk

`func (o *HypervisorResourcePoolAzureDetailResponseModel) GetContainerScopesOk() (*[]ContainerScopeResponseModel, bool)`

GetContainerScopesOk returns a tuple with the ContainerScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScopes

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetContainerScopes(v []ContainerScopeResponseModel)`

SetContainerScopes sets ContainerScopes field to given value.

### HasContainerScopes

`func (o *HypervisorResourcePoolAzureDetailResponseModel) HasContainerScopes() bool`

HasContainerScopes returns a boolean if a field has been set.

### SetContainerScopesNil

`func (o *HypervisorResourcePoolAzureDetailResponseModel) SetContainerScopesNil(b bool)`

 SetContainerScopesNil sets the value for ContainerScopes to be an explicit nil

### UnsetContainerScopes
`func (o *HypervisorResourcePoolAzureDetailResponseModel) UnsetContainerScopes()`

UnsetContainerScopes ensures that no value is present for ContainerScopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


