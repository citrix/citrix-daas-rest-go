# EditHypervisorResourcePoolAzureRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name of the resource pool.  Optional.  If not specified, will not be changed. | [optional] 
**ConnectionType** | [**HypervisorConnectionType**](HypervisorConnectionType.md) |  | 
**VmTagging** | Pointer to **NullableBool** | Indicates whether VMs created by Virtual Apps &amp; Desktops provisioning operations should be tagged.  Tagged VMs are filtered out of queries by default. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Metadata for hypervisor resource pool. When set the property value equal to null/empty means to remove this property. | [optional] 
**Storage** | Pointer to [**[]HypervisorResourcePoolStorageRequestModel**](HypervisorResourcePoolStorageRequestModel.md) | Path to the storage resource(s) that are available for provisioning operations in this resource pool. Optional; if not specified, will not be changed.  If any storage is specified, all desired storage must be listed.  Any storage that was previously used for the resource pool, but not listed during an edit operation, will be removed. Note that removing storage from the resource pool will permanently disable the ability to update or rollback machines that are provisioned using that storage.  Therefore, removal of storage should _only_ be done when the storage is terminally broken or is being replaced. | [optional] 
**TemporaryStorage** | Pointer to [**[]HypervisorResourcePoolStorageRequestModel**](HypervisorResourcePoolStorageRequestModel.md) | Path to the storage resource(s) that are used for temporary operations in this resource pool. Optional; if not specified, will not be changed.  If any storage is specified, all desired storage must be listed.  Any storage that was previously used for the resource pool, but not listed during an edit operation, will be removed. Note that removing storage from the resource pool will permanently disable the ability to update or rollback machines that are provisioned using that storage.  Therefore, removal of storage should _only_ be done when the storage is terminally broken or is being replaced. | [optional] 
**PersonalvDiskStorage** | Pointer to [**[]HypervisorResourcePoolStorageRequestModel**](HypervisorResourcePoolStorageRequestModel.md) | Path to the personal virtual disk storage resource(s). Optional; if not specified, will not be changed.  If any storage is specified, all desired storage must be listed.  Any storage that was previously used for the resource pool, but not listed during an edit operation, will be removed. Note that removing storage from the resource pool will permanently disable the ability to update or rollback machines that are provisioned using that storage.  Therefore, removal of storage should _only_ be done when the storage is terminally broken or is being replaced. | [optional] 
**CustomProperties** | Pointer to **NullableString** | Custom properties.  Optional.  If not specified, will not be changed.  Only used for hypervisors of type Custom. | [optional] 
**UseLocalStorageCaching** | Pointer to **NullableBool** | Indicates whether local storage on the hypervisor will be used for caching purposes. Not all hypervisor types support this.  Defaults to &#x60;false&#x60;. | [optional] 
**Networks** | Pointer to **[]string** | Path to the network resource(s) that are available for provisioning operations in this resource pool.  At least one is required. | [optional] 
**StorageBalanceType** | Pointer to [**StorageBalanceType**](StorageBalanceType.md) |  | [optional] 
**Subnets** | Pointer to **[]string** | Path to the subnet(s) that are available for provisioning operations in this resource pool.  At least one is required. | [optional] 

## Methods

### NewEditHypervisorResourcePoolAzureRequestModel

`func NewEditHypervisorResourcePoolAzureRequestModel(connectionType HypervisorConnectionType, ) *EditHypervisorResourcePoolAzureRequestModel`

NewEditHypervisorResourcePoolAzureRequestModel instantiates a new EditHypervisorResourcePoolAzureRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditHypervisorResourcePoolAzureRequestModelWithDefaults

`func NewEditHypervisorResourcePoolAzureRequestModelWithDefaults() *EditHypervisorResourcePoolAzureRequestModel`

NewEditHypervisorResourcePoolAzureRequestModelWithDefaults instantiates a new EditHypervisorResourcePoolAzureRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditHypervisorResourcePoolAzureRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditHypervisorResourcePoolAzureRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EditHypervisorResourcePoolAzureRequestModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EditHypervisorResourcePoolAzureRequestModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetConnectionType

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *EditHypervisorResourcePoolAzureRequestModel) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.


### GetVmTagging

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetVmTagging() bool`

GetVmTagging returns the VmTagging field if non-nil, zero value otherwise.

### GetVmTaggingOk

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetVmTaggingOk() (*bool, bool)`

GetVmTaggingOk returns a tuple with the VmTagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTagging

`func (o *EditHypervisorResourcePoolAzureRequestModel) SetVmTagging(v bool)`

SetVmTagging sets VmTagging field to given value.

### HasVmTagging

`func (o *EditHypervisorResourcePoolAzureRequestModel) HasVmTagging() bool`

HasVmTagging returns a boolean if a field has been set.

### SetVmTaggingNil

`func (o *EditHypervisorResourcePoolAzureRequestModel) SetVmTaggingNil(b bool)`

 SetVmTaggingNil sets the value for VmTagging to be an explicit nil

### UnsetVmTagging
`func (o *EditHypervisorResourcePoolAzureRequestModel) UnsetVmTagging()`

UnsetVmTagging ensures that no value is present for VmTagging, not even an explicit nil
### GetMetadata

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EditHypervisorResourcePoolAzureRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EditHypervisorResourcePoolAzureRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *EditHypervisorResourcePoolAzureRequestModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *EditHypervisorResourcePoolAzureRequestModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetStorage

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetStorage() []HypervisorResourcePoolStorageRequestModel`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetStorageOk() (*[]HypervisorResourcePoolStorageRequestModel, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *EditHypervisorResourcePoolAzureRequestModel) SetStorage(v []HypervisorResourcePoolStorageRequestModel)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *EditHypervisorResourcePoolAzureRequestModel) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *EditHypervisorResourcePoolAzureRequestModel) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *EditHypervisorResourcePoolAzureRequestModel) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetTemporaryStorage

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetTemporaryStorage() []HypervisorResourcePoolStorageRequestModel`

GetTemporaryStorage returns the TemporaryStorage field if non-nil, zero value otherwise.

### GetTemporaryStorageOk

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetTemporaryStorageOk() (*[]HypervisorResourcePoolStorageRequestModel, bool)`

GetTemporaryStorageOk returns a tuple with the TemporaryStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryStorage

`func (o *EditHypervisorResourcePoolAzureRequestModel) SetTemporaryStorage(v []HypervisorResourcePoolStorageRequestModel)`

SetTemporaryStorage sets TemporaryStorage field to given value.

### HasTemporaryStorage

`func (o *EditHypervisorResourcePoolAzureRequestModel) HasTemporaryStorage() bool`

HasTemporaryStorage returns a boolean if a field has been set.

### SetTemporaryStorageNil

`func (o *EditHypervisorResourcePoolAzureRequestModel) SetTemporaryStorageNil(b bool)`

 SetTemporaryStorageNil sets the value for TemporaryStorage to be an explicit nil

### UnsetTemporaryStorage
`func (o *EditHypervisorResourcePoolAzureRequestModel) UnsetTemporaryStorage()`

UnsetTemporaryStorage ensures that no value is present for TemporaryStorage, not even an explicit nil
### GetPersonalvDiskStorage

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetPersonalvDiskStorage() []HypervisorResourcePoolStorageRequestModel`

GetPersonalvDiskStorage returns the PersonalvDiskStorage field if non-nil, zero value otherwise.

### GetPersonalvDiskStorageOk

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetPersonalvDiskStorageOk() (*[]HypervisorResourcePoolStorageRequestModel, bool)`

GetPersonalvDiskStorageOk returns a tuple with the PersonalvDiskStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalvDiskStorage

`func (o *EditHypervisorResourcePoolAzureRequestModel) SetPersonalvDiskStorage(v []HypervisorResourcePoolStorageRequestModel)`

SetPersonalvDiskStorage sets PersonalvDiskStorage field to given value.

### HasPersonalvDiskStorage

`func (o *EditHypervisorResourcePoolAzureRequestModel) HasPersonalvDiskStorage() bool`

HasPersonalvDiskStorage returns a boolean if a field has been set.

### SetPersonalvDiskStorageNil

`func (o *EditHypervisorResourcePoolAzureRequestModel) SetPersonalvDiskStorageNil(b bool)`

 SetPersonalvDiskStorageNil sets the value for PersonalvDiskStorage to be an explicit nil

### UnsetPersonalvDiskStorage
`func (o *EditHypervisorResourcePoolAzureRequestModel) UnsetPersonalvDiskStorage()`

UnsetPersonalvDiskStorage ensures that no value is present for PersonalvDiskStorage, not even an explicit nil
### GetCustomProperties

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *EditHypervisorResourcePoolAzureRequestModel) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *EditHypervisorResourcePoolAzureRequestModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *EditHypervisorResourcePoolAzureRequestModel) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *EditHypervisorResourcePoolAzureRequestModel) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
### GetUseLocalStorageCaching

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetUseLocalStorageCaching() bool`

GetUseLocalStorageCaching returns the UseLocalStorageCaching field if non-nil, zero value otherwise.

### GetUseLocalStorageCachingOk

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetUseLocalStorageCachingOk() (*bool, bool)`

GetUseLocalStorageCachingOk returns a tuple with the UseLocalStorageCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocalStorageCaching

`func (o *EditHypervisorResourcePoolAzureRequestModel) SetUseLocalStorageCaching(v bool)`

SetUseLocalStorageCaching sets UseLocalStorageCaching field to given value.

### HasUseLocalStorageCaching

`func (o *EditHypervisorResourcePoolAzureRequestModel) HasUseLocalStorageCaching() bool`

HasUseLocalStorageCaching returns a boolean if a field has been set.

### SetUseLocalStorageCachingNil

`func (o *EditHypervisorResourcePoolAzureRequestModel) SetUseLocalStorageCachingNil(b bool)`

 SetUseLocalStorageCachingNil sets the value for UseLocalStorageCaching to be an explicit nil

### UnsetUseLocalStorageCaching
`func (o *EditHypervisorResourcePoolAzureRequestModel) UnsetUseLocalStorageCaching()`

UnsetUseLocalStorageCaching ensures that no value is present for UseLocalStorageCaching, not even an explicit nil
### GetNetworks

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *EditHypervisorResourcePoolAzureRequestModel) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *EditHypervisorResourcePoolAzureRequestModel) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### SetNetworksNil

`func (o *EditHypervisorResourcePoolAzureRequestModel) SetNetworksNil(b bool)`

 SetNetworksNil sets the value for Networks to be an explicit nil

### UnsetNetworks
`func (o *EditHypervisorResourcePoolAzureRequestModel) UnsetNetworks()`

UnsetNetworks ensures that no value is present for Networks, not even an explicit nil
### GetStorageBalanceType

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetStorageBalanceType() StorageBalanceType`

GetStorageBalanceType returns the StorageBalanceType field if non-nil, zero value otherwise.

### GetStorageBalanceTypeOk

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetStorageBalanceTypeOk() (*StorageBalanceType, bool)`

GetStorageBalanceTypeOk returns a tuple with the StorageBalanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageBalanceType

`func (o *EditHypervisorResourcePoolAzureRequestModel) SetStorageBalanceType(v StorageBalanceType)`

SetStorageBalanceType sets StorageBalanceType field to given value.

### HasStorageBalanceType

`func (o *EditHypervisorResourcePoolAzureRequestModel) HasStorageBalanceType() bool`

HasStorageBalanceType returns a boolean if a field has been set.

### GetSubnets

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *EditHypervisorResourcePoolAzureRequestModel) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *EditHypervisorResourcePoolAzureRequestModel) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *EditHypervisorResourcePoolAzureRequestModel) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### SetSubnetsNil

`func (o *EditHypervisorResourcePoolAzureRequestModel) SetSubnetsNil(b bool)`

 SetSubnetsNil sets the value for Subnets to be an explicit nil

### UnsetSubnets
`func (o *EditHypervisorResourcePoolAzureRequestModel) UnsetSubnets()`

UnsetSubnets ensures that no value is present for Subnets, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


