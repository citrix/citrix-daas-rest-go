# HypervisorResourcePoolTraditionalDetailResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**XDPath** | Pointer to **string** | XenApp &amp; XenDesktop path to the resource on the hypervisor.  An example value is: &#x60;XDHyp:\\Connections\\{{hypervisor name}}\\{{vm name}}.vm\\{{snapshot name}}.snapshot&#x60; or &#x60;XDHyp:\\HostingUnits\\{{resource pool name}}\\{{resource name}}.{{resource type}}&#x60; | [optional] 
**HypervisorConnection** | [**HypervisorResourcePoolResponseModelAllOfHypervisorConnection**](HypervisorResourcePoolResponseModelAllOfHypervisorConnection.md) |  | 
**ConnectionType** | [**HypervisorConnectionType**](HypervisorConnectionType.md) |  | 
**DefaultNetwork** | [**HypervisorResourcePoolResponseModelAllOfDefaultNetwork**](HypervisorResourcePoolResponseModelAllOfDefaultNetwork.md) |  | 
**VMTaggingEnabled** | **bool** | Indicates whether new virtual machines are tagged with metadata from the hypervisor. | 
**ResourcePoolRootPath** | Pointer to **string** | Hypervisor resourcePool RootPath. | [optional] 
**ResourcePoolRootId** | Pointer to **string** | Hypervisor resourcePool RootId. | [optional] 
**GpuTypes** | Pointer to [**[]HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) | GPU types available in the resource pool.  Only applicable to hypervisors that support GPU types. | [optional] 
**UsesExplicitStorage** | Pointer to **bool** | If the hypervisor resource pool use ExplicitStorage. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Metadata for hypervisor resource pool.  | [optional] 
**RootPath** | Pointer to [**HypervisorResourcePoolTraditionalDetailResponseModelAllOfRootPath**](HypervisorResourcePoolTraditionalDetailResponseModelAllOfRootPath.md) |  | [optional] 
**Networks** | [**[]HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) | List of networks that are allowed to be used within the resource pool. | 
**Storage** | [**[]HypervisorStorageResourceResponseModel**](HypervisorStorageResourceResponseModel.md) | List of hypervisor-connected storage in the resource pool that is used for OS disks of virtual machines. | 
**TemporaryStorage** | [**[]HypervisorStorageResourceResponseModel**](HypervisorStorageResourceResponseModel.md) | List of hypervisor-connected storage in the resource pool that is used for temporary data storage for virtual machines. | 
**UseLocalStorageCaching** | Pointer to **bool** | Indicates whether virtual machines created within this resource pool will use local storage caching for their disk images. | [optional] 
**CustomProperties** | Pointer to **string** | Custom properties.  Optional.  If not specified, will not be changed.  Only used for hypervisors of type Custom. | [optional] 
**PersonalvDiskStorage** | Pointer to [**[]HypervisorStorageResourceResponseModel**](HypervisorStorageResourceResponseModel.md) | List of hypervisor-connected storage in the resource pool that is used for personal v disk data storage for virtual machines. | [optional] 

## Methods

### NewHypervisorResourcePoolTraditionalDetailResponseModel

`func NewHypervisorResourcePoolTraditionalDetailResponseModel(hypervisorConnection HypervisorResourcePoolResponseModelAllOfHypervisorConnection, connectionType HypervisorConnectionType, defaultNetwork HypervisorResourcePoolResponseModelAllOfDefaultNetwork, vMTaggingEnabled bool, networks []HypervisorResourceRefResponseModel, storage []HypervisorStorageResourceResponseModel, temporaryStorage []HypervisorStorageResourceResponseModel, ) *HypervisorResourcePoolTraditionalDetailResponseModel`

NewHypervisorResourcePoolTraditionalDetailResponseModel instantiates a new HypervisorResourcePoolTraditionalDetailResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorResourcePoolTraditionalDetailResponseModelWithDefaults

`func NewHypervisorResourcePoolTraditionalDetailResponseModelWithDefaults() *HypervisorResourcePoolTraditionalDetailResponseModel`

NewHypervisorResourcePoolTraditionalDetailResponseModelWithDefaults instantiates a new HypervisorResourcePoolTraditionalDetailResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetXDPath

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetXDPath() string`

GetXDPath returns the XDPath field if non-nil, zero value otherwise.

### GetXDPathOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetXDPathOk() (*string, bool)`

GetXDPathOk returns a tuple with the XDPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXDPath

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) SetXDPath(v string)`

SetXDPath sets XDPath field to given value.

### HasXDPath

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) HasXDPath() bool`

HasXDPath returns a boolean if a field has been set.

### GetHypervisorConnection

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetHypervisorConnection() HypervisorResourcePoolResponseModelAllOfHypervisorConnection`

GetHypervisorConnection returns the HypervisorConnection field if non-nil, zero value otherwise.

### GetHypervisorConnectionOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetHypervisorConnectionOk() (*HypervisorResourcePoolResponseModelAllOfHypervisorConnection, bool)`

GetHypervisorConnectionOk returns a tuple with the HypervisorConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorConnection

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) SetHypervisorConnection(v HypervisorResourcePoolResponseModelAllOfHypervisorConnection)`

SetHypervisorConnection sets HypervisorConnection field to given value.


### GetConnectionType

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.


### GetDefaultNetwork

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetDefaultNetwork() HypervisorResourcePoolResponseModelAllOfDefaultNetwork`

GetDefaultNetwork returns the DefaultNetwork field if non-nil, zero value otherwise.

### GetDefaultNetworkOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetDefaultNetworkOk() (*HypervisorResourcePoolResponseModelAllOfDefaultNetwork, bool)`

GetDefaultNetworkOk returns a tuple with the DefaultNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetwork

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) SetDefaultNetwork(v HypervisorResourcePoolResponseModelAllOfDefaultNetwork)`

SetDefaultNetwork sets DefaultNetwork field to given value.


### GetVMTaggingEnabled

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetVMTaggingEnabled() bool`

GetVMTaggingEnabled returns the VMTaggingEnabled field if non-nil, zero value otherwise.

### GetVMTaggingEnabledOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetVMTaggingEnabledOk() (*bool, bool)`

GetVMTaggingEnabledOk returns a tuple with the VMTaggingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMTaggingEnabled

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) SetVMTaggingEnabled(v bool)`

SetVMTaggingEnabled sets VMTaggingEnabled field to given value.


### GetResourcePoolRootPath

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetResourcePoolRootPath() string`

GetResourcePoolRootPath returns the ResourcePoolRootPath field if non-nil, zero value otherwise.

### GetResourcePoolRootPathOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetResourcePoolRootPathOk() (*string, bool)`

GetResourcePoolRootPathOk returns a tuple with the ResourcePoolRootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolRootPath

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) SetResourcePoolRootPath(v string)`

SetResourcePoolRootPath sets ResourcePoolRootPath field to given value.

### HasResourcePoolRootPath

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) HasResourcePoolRootPath() bool`

HasResourcePoolRootPath returns a boolean if a field has been set.

### GetResourcePoolRootId

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetResourcePoolRootId() string`

GetResourcePoolRootId returns the ResourcePoolRootId field if non-nil, zero value otherwise.

### GetResourcePoolRootIdOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetResourcePoolRootIdOk() (*string, bool)`

GetResourcePoolRootIdOk returns a tuple with the ResourcePoolRootId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolRootId

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) SetResourcePoolRootId(v string)`

SetResourcePoolRootId sets ResourcePoolRootId field to given value.

### HasResourcePoolRootId

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) HasResourcePoolRootId() bool`

HasResourcePoolRootId returns a boolean if a field has been set.

### GetGpuTypes

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetGpuTypes() []HypervisorResourceRefResponseModel`

GetGpuTypes returns the GpuTypes field if non-nil, zero value otherwise.

### GetGpuTypesOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetGpuTypesOk() (*[]HypervisorResourceRefResponseModel, bool)`

GetGpuTypesOk returns a tuple with the GpuTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuTypes

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) SetGpuTypes(v []HypervisorResourceRefResponseModel)`

SetGpuTypes sets GpuTypes field to given value.

### HasGpuTypes

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) HasGpuTypes() bool`

HasGpuTypes returns a boolean if a field has been set.

### GetUsesExplicitStorage

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetUsesExplicitStorage() bool`

GetUsesExplicitStorage returns the UsesExplicitStorage field if non-nil, zero value otherwise.

### GetUsesExplicitStorageOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetUsesExplicitStorageOk() (*bool, bool)`

GetUsesExplicitStorageOk returns a tuple with the UsesExplicitStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesExplicitStorage

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) SetUsesExplicitStorage(v bool)`

SetUsesExplicitStorage sets UsesExplicitStorage field to given value.

### HasUsesExplicitStorage

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) HasUsesExplicitStorage() bool`

HasUsesExplicitStorage returns a boolean if a field has been set.

### GetMetadata

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRootPath

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetRootPath() HypervisorResourcePoolTraditionalDetailResponseModelAllOfRootPath`

GetRootPath returns the RootPath field if non-nil, zero value otherwise.

### GetRootPathOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetRootPathOk() (*HypervisorResourcePoolTraditionalDetailResponseModelAllOfRootPath, bool)`

GetRootPathOk returns a tuple with the RootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPath

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) SetRootPath(v HypervisorResourcePoolTraditionalDetailResponseModelAllOfRootPath)`

SetRootPath sets RootPath field to given value.

### HasRootPath

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) HasRootPath() bool`

HasRootPath returns a boolean if a field has been set.

### GetNetworks

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetNetworks() []HypervisorResourceRefResponseModel`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetNetworksOk() (*[]HypervisorResourceRefResponseModel, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) SetNetworks(v []HypervisorResourceRefResponseModel)`

SetNetworks sets Networks field to given value.


### GetStorage

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetStorage() []HypervisorStorageResourceResponseModel`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetStorageOk() (*[]HypervisorStorageResourceResponseModel, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) SetStorage(v []HypervisorStorageResourceResponseModel)`

SetStorage sets Storage field to given value.


### GetTemporaryStorage

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetTemporaryStorage() []HypervisorStorageResourceResponseModel`

GetTemporaryStorage returns the TemporaryStorage field if non-nil, zero value otherwise.

### GetTemporaryStorageOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetTemporaryStorageOk() (*[]HypervisorStorageResourceResponseModel, bool)`

GetTemporaryStorageOk returns a tuple with the TemporaryStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryStorage

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) SetTemporaryStorage(v []HypervisorStorageResourceResponseModel)`

SetTemporaryStorage sets TemporaryStorage field to given value.


### GetUseLocalStorageCaching

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetUseLocalStorageCaching() bool`

GetUseLocalStorageCaching returns the UseLocalStorageCaching field if non-nil, zero value otherwise.

### GetUseLocalStorageCachingOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetUseLocalStorageCachingOk() (*bool, bool)`

GetUseLocalStorageCachingOk returns a tuple with the UseLocalStorageCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocalStorageCaching

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) SetUseLocalStorageCaching(v bool)`

SetUseLocalStorageCaching sets UseLocalStorageCaching field to given value.

### HasUseLocalStorageCaching

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) HasUseLocalStorageCaching() bool`

HasUseLocalStorageCaching returns a boolean if a field has been set.

### GetCustomProperties

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetPersonalvDiskStorage

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetPersonalvDiskStorage() []HypervisorStorageResourceResponseModel`

GetPersonalvDiskStorage returns the PersonalvDiskStorage field if non-nil, zero value otherwise.

### GetPersonalvDiskStorageOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) GetPersonalvDiskStorageOk() (*[]HypervisorStorageResourceResponseModel, bool)`

GetPersonalvDiskStorageOk returns a tuple with the PersonalvDiskStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalvDiskStorage

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) SetPersonalvDiskStorage(v []HypervisorStorageResourceResponseModel)`

SetPersonalvDiskStorage sets PersonalvDiskStorage field to given value.

### HasPersonalvDiskStorage

`func (o *HypervisorResourcePoolTraditionalDetailResponseModel) HasPersonalvDiskStorage() bool`

HasPersonalvDiskStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


