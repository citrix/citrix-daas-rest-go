# EditHypervisorResourcePoolTraditionalRequestModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Storage** | Pointer to [**[]HypervisorResourcePoolStorageRequestModel**](HypervisorResourcePoolStorageRequestModel.md) | Path to the storage resource(s) that are available for provisioning operations in this resource pool. Optional; if not specified, will not be changed.  If any storage is specified, all desired storage must be listed.  Any storage that was previously used for the resource pool, but not listed during an edit operation, will be removed. Note that removing storage from the resource pool will permanently disable the ability to update or rollback machines that are provisioned using that storage.  Therefore, removal of storage should _only_ be done when the storage is terminally broken or is being replaced. | [optional] 
**TemporaryStorage** | Pointer to [**[]HypervisorResourcePoolStorageRequestModel**](HypervisorResourcePoolStorageRequestModel.md) | Path to the storage resource(s) that are used for temporary operations in this resource pool. Optional; if not specified, will not be changed.  If any storage is specified, all desired storage must be listed.  Any storage that was previously used for the resource pool, but not listed during an edit operation, will be removed. Note that removing storage from the resource pool will permanently disable the ability to update or rollback machines that are provisioned using that storage.  Therefore, removal of storage should _only_ be done when the storage is terminally broken or is being replaced. | [optional] 
**PersonalvDiskStorage** | Pointer to [**[]HypervisorResourcePoolStorageRequestModel**](HypervisorResourcePoolStorageRequestModel.md) | Path to the personal virtual disk storage resource(s). Optional; if not specified, will not be changed.  If any storage is specified, all desired storage must be listed.  Any storage that was previously used for the resource pool, but not listed during an edit operation, will be removed. Note that removing storage from the resource pool will permanently disable the ability to update or rollback machines that are provisioned using that storage.  Therefore, removal of storage should _only_ be done when the storage is terminally broken or is being replaced. | [optional] 
**CustomProperties** | Pointer to **string** | Custom properties.  Optional.  If not specified, will not be changed.  Only used for hypervisors of type Custom. | [optional] 
**UseLocalStorageCaching** | Pointer to **bool** | Indicates whether local storage on the hypervisor will be used for caching purposes. Not all hypervisor types support this.  Defaults to &#x60;false&#x60;. | [optional] 
**Networks** | Pointer to **[]string** | Path to the network resource(s) that are available for provisioning operations in this resource pool.  At least one is required. | [optional] 

## Methods

### NewEditHypervisorResourcePoolTraditionalRequestModelAllOf

`func NewEditHypervisorResourcePoolTraditionalRequestModelAllOf() *EditHypervisorResourcePoolTraditionalRequestModelAllOf`

NewEditHypervisorResourcePoolTraditionalRequestModelAllOf instantiates a new EditHypervisorResourcePoolTraditionalRequestModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditHypervisorResourcePoolTraditionalRequestModelAllOfWithDefaults

`func NewEditHypervisorResourcePoolTraditionalRequestModelAllOfWithDefaults() *EditHypervisorResourcePoolTraditionalRequestModelAllOf`

NewEditHypervisorResourcePoolTraditionalRequestModelAllOfWithDefaults instantiates a new EditHypervisorResourcePoolTraditionalRequestModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorage

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetStorage() []HypervisorResourcePoolStorageRequestModel`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetStorageOk() (*[]HypervisorResourcePoolStorageRequestModel, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) SetStorage(v []HypervisorResourcePoolStorageRequestModel)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetTemporaryStorage

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetTemporaryStorage() []HypervisorResourcePoolStorageRequestModel`

GetTemporaryStorage returns the TemporaryStorage field if non-nil, zero value otherwise.

### GetTemporaryStorageOk

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetTemporaryStorageOk() (*[]HypervisorResourcePoolStorageRequestModel, bool)`

GetTemporaryStorageOk returns a tuple with the TemporaryStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryStorage

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) SetTemporaryStorage(v []HypervisorResourcePoolStorageRequestModel)`

SetTemporaryStorage sets TemporaryStorage field to given value.

### HasTemporaryStorage

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) HasTemporaryStorage() bool`

HasTemporaryStorage returns a boolean if a field has been set.

### GetPersonalvDiskStorage

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetPersonalvDiskStorage() []HypervisorResourcePoolStorageRequestModel`

GetPersonalvDiskStorage returns the PersonalvDiskStorage field if non-nil, zero value otherwise.

### GetPersonalvDiskStorageOk

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetPersonalvDiskStorageOk() (*[]HypervisorResourcePoolStorageRequestModel, bool)`

GetPersonalvDiskStorageOk returns a tuple with the PersonalvDiskStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalvDiskStorage

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) SetPersonalvDiskStorage(v []HypervisorResourcePoolStorageRequestModel)`

SetPersonalvDiskStorage sets PersonalvDiskStorage field to given value.

### HasPersonalvDiskStorage

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) HasPersonalvDiskStorage() bool`

HasPersonalvDiskStorage returns a boolean if a field has been set.

### GetCustomProperties

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetUseLocalStorageCaching

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetUseLocalStorageCaching() bool`

GetUseLocalStorageCaching returns the UseLocalStorageCaching field if non-nil, zero value otherwise.

### GetUseLocalStorageCachingOk

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetUseLocalStorageCachingOk() (*bool, bool)`

GetUseLocalStorageCachingOk returns a tuple with the UseLocalStorageCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocalStorageCaching

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) SetUseLocalStorageCaching(v bool)`

SetUseLocalStorageCaching sets UseLocalStorageCaching field to given value.

### HasUseLocalStorageCaching

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) HasUseLocalStorageCaching() bool`

HasUseLocalStorageCaching returns a boolean if a field has been set.

### GetNetworks

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *EditHypervisorResourcePoolTraditionalRequestModelAllOf) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


