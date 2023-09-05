# EditHypervisorResourcePoolTraditionalRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the resource pool.  Optional.  If not specified, will not be changed. | [optional] 
**ConnectionType** | [**HypervisorConnectionType**](HypervisorConnectionType.md) |  | 
**VmTagging** | Pointer to **bool** | Indicates whether VMs created by Virtual Apps &amp; Desktops provisioning operations should be tagged.  Tagged VMs are filtered out of queries by default. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Metadata for hypervisor resource pool. When set the property value equal to null/empty means to remove this property. | [optional] 
**Storage** | Pointer to [**[]HypervisorResourcePoolStorageRequestModel**](HypervisorResourcePoolStorageRequestModel.md) | Path to the storage resource(s) that are available for provisioning operations in this resource pool. Optional; if not specified, will not be changed.  If any storage is specified, all desired storage must be listed.  Any storage that was previously used for the resource pool, but not listed during an edit operation, will be removed. Note that removing storage from the resource pool will permanently disable the ability to update or rollback machines that are provisioned using that storage.  Therefore, removal of storage should _only_ be done when the storage is terminally broken or is being replaced. | [optional] 
**TemporaryStorage** | Pointer to [**[]HypervisorResourcePoolStorageRequestModel**](HypervisorResourcePoolStorageRequestModel.md) | Path to the storage resource(s) that are used for temporary operations in this resource pool. Optional; if not specified, will not be changed.  If any storage is specified, all desired storage must be listed.  Any storage that was previously used for the resource pool, but not listed during an edit operation, will be removed. Note that removing storage from the resource pool will permanently disable the ability to update or rollback machines that are provisioned using that storage.  Therefore, removal of storage should _only_ be done when the storage is terminally broken or is being replaced. | [optional] 
**PersonalvDiskStorage** | Pointer to [**[]HypervisorResourcePoolStorageRequestModel**](HypervisorResourcePoolStorageRequestModel.md) | Path to the personal virtual disk storage resource(s). Optional; if not specified, will not be changed.  If any storage is specified, all desired storage must be listed.  Any storage that was previously used for the resource pool, but not listed during an edit operation, will be removed. Note that removing storage from the resource pool will permanently disable the ability to update or rollback machines that are provisioned using that storage.  Therefore, removal of storage should _only_ be done when the storage is terminally broken or is being replaced. | [optional] 
**CustomProperties** | Pointer to **string** | Custom properties.  Optional.  If not specified, will not be changed.  Only used for hypervisors of type Custom. | [optional] 
**UseLocalStorageCaching** | Pointer to **bool** | Indicates whether local storage on the hypervisor will be used for caching purposes. Not all hypervisor types support this.  Defaults to &#x60;false&#x60;. | [optional] 
**Networks** | Pointer to **[]string** | Path to the network resource(s) that are available for provisioning operations in this resource pool.  At least one is required. | [optional] 
**Subnets** | Pointer to **[]string** | Path to the subnet(s) that are available for provisioning operations in this resource pool.  At least one is required. | [optional] 

## Methods

### NewEditHypervisorResourcePoolTraditionalRequestModel

`func NewEditHypervisorResourcePoolTraditionalRequestModel(connectionType HypervisorConnectionType, ) *EditHypervisorResourcePoolTraditionalRequestModel`

NewEditHypervisorResourcePoolTraditionalRequestModel instantiates a new EditHypervisorResourcePoolTraditionalRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditHypervisorResourcePoolTraditionalRequestModelWithDefaults

`func NewEditHypervisorResourcePoolTraditionalRequestModelWithDefaults() *EditHypervisorResourcePoolTraditionalRequestModel`

NewEditHypervisorResourcePoolTraditionalRequestModelWithDefaults instantiates a new EditHypervisorResourcePoolTraditionalRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConnectionType

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.


### GetVmTagging

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) GetVmTagging() bool`

GetVmTagging returns the VmTagging field if non-nil, zero value otherwise.

### GetVmTaggingOk

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) GetVmTaggingOk() (*bool, bool)`

GetVmTaggingOk returns a tuple with the VmTagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTagging

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) SetVmTagging(v bool)`

SetVmTagging sets VmTagging field to given value.

### HasVmTagging

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) HasVmTagging() bool`

HasVmTagging returns a boolean if a field has been set.

### GetMetadata

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStorage

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) GetStorage() []HypervisorResourcePoolStorageRequestModel`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) GetStorageOk() (*[]HypervisorResourcePoolStorageRequestModel, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) SetStorage(v []HypervisorResourcePoolStorageRequestModel)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetTemporaryStorage

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) GetTemporaryStorage() []HypervisorResourcePoolStorageRequestModel`

GetTemporaryStorage returns the TemporaryStorage field if non-nil, zero value otherwise.

### GetTemporaryStorageOk

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) GetTemporaryStorageOk() (*[]HypervisorResourcePoolStorageRequestModel, bool)`

GetTemporaryStorageOk returns a tuple with the TemporaryStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryStorage

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) SetTemporaryStorage(v []HypervisorResourcePoolStorageRequestModel)`

SetTemporaryStorage sets TemporaryStorage field to given value.

### HasTemporaryStorage

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) HasTemporaryStorage() bool`

HasTemporaryStorage returns a boolean if a field has been set.

### GetPersonalvDiskStorage

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) GetPersonalvDiskStorage() []HypervisorResourcePoolStorageRequestModel`

GetPersonalvDiskStorage returns the PersonalvDiskStorage field if non-nil, zero value otherwise.

### GetPersonalvDiskStorageOk

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) GetPersonalvDiskStorageOk() (*[]HypervisorResourcePoolStorageRequestModel, bool)`

GetPersonalvDiskStorageOk returns a tuple with the PersonalvDiskStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalvDiskStorage

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) SetPersonalvDiskStorage(v []HypervisorResourcePoolStorageRequestModel)`

SetPersonalvDiskStorage sets PersonalvDiskStorage field to given value.

### HasPersonalvDiskStorage

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) HasPersonalvDiskStorage() bool`

HasPersonalvDiskStorage returns a boolean if a field has been set.

### GetCustomProperties

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetUseLocalStorageCaching

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) GetUseLocalStorageCaching() bool`

GetUseLocalStorageCaching returns the UseLocalStorageCaching field if non-nil, zero value otherwise.

### GetUseLocalStorageCachingOk

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) GetUseLocalStorageCachingOk() (*bool, bool)`

GetUseLocalStorageCachingOk returns a tuple with the UseLocalStorageCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocalStorageCaching

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) SetUseLocalStorageCaching(v bool)`

SetUseLocalStorageCaching sets UseLocalStorageCaching field to given value.

### HasUseLocalStorageCaching

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) HasUseLocalStorageCaching() bool`

HasUseLocalStorageCaching returns a boolean if a field has been set.

### GetNetworks

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetSubnets

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *EditHypervisorResourcePoolTraditionalRequestModel) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


