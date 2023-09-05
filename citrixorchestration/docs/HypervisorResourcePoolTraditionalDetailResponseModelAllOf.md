# HypervisorResourcePoolTraditionalDetailResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RootPath** | Pointer to [**HypervisorResourcePoolTraditionalDetailResponseModelAllOfRootPath**](HypervisorResourcePoolTraditionalDetailResponseModelAllOfRootPath.md) |  | [optional] 
**Networks** | [**[]HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) | List of networks that are allowed to be used within the resource pool. | 
**Storage** | [**[]HypervisorStorageResourceResponseModel**](HypervisorStorageResourceResponseModel.md) | List of hypervisor-connected storage in the resource pool that is used for OS disks of virtual machines. | 
**TemporaryStorage** | [**[]HypervisorStorageResourceResponseModel**](HypervisorStorageResourceResponseModel.md) | List of hypervisor-connected storage in the resource pool that is used for temporary data storage for virtual machines. | 
**UseLocalStorageCaching** | Pointer to **bool** | Indicates whether virtual machines created within this resource pool will use local storage caching for their disk images. | [optional] 
**CustomProperties** | Pointer to **string** | Custom properties.  Optional.  If not specified, will not be changed.  Only used for hypervisors of type Custom. | [optional] 
**PersonalvDiskStorage** | Pointer to [**[]HypervisorStorageResourceResponseModel**](HypervisorStorageResourceResponseModel.md) | List of hypervisor-connected storage in the resource pool that is used for personal v disk data storage for virtual machines. | [optional] 

## Methods

### NewHypervisorResourcePoolTraditionalDetailResponseModelAllOf

`func NewHypervisorResourcePoolTraditionalDetailResponseModelAllOf(networks []HypervisorResourceRefResponseModel, storage []HypervisorStorageResourceResponseModel, temporaryStorage []HypervisorStorageResourceResponseModel, ) *HypervisorResourcePoolTraditionalDetailResponseModelAllOf`

NewHypervisorResourcePoolTraditionalDetailResponseModelAllOf instantiates a new HypervisorResourcePoolTraditionalDetailResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorResourcePoolTraditionalDetailResponseModelAllOfWithDefaults

`func NewHypervisorResourcePoolTraditionalDetailResponseModelAllOfWithDefaults() *HypervisorResourcePoolTraditionalDetailResponseModelAllOf`

NewHypervisorResourcePoolTraditionalDetailResponseModelAllOfWithDefaults instantiates a new HypervisorResourcePoolTraditionalDetailResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRootPath

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) GetRootPath() HypervisorResourcePoolTraditionalDetailResponseModelAllOfRootPath`

GetRootPath returns the RootPath field if non-nil, zero value otherwise.

### GetRootPathOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) GetRootPathOk() (*HypervisorResourcePoolTraditionalDetailResponseModelAllOfRootPath, bool)`

GetRootPathOk returns a tuple with the RootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPath

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) SetRootPath(v HypervisorResourcePoolTraditionalDetailResponseModelAllOfRootPath)`

SetRootPath sets RootPath field to given value.

### HasRootPath

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) HasRootPath() bool`

HasRootPath returns a boolean if a field has been set.

### GetNetworks

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) GetNetworks() []HypervisorResourceRefResponseModel`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) GetNetworksOk() (*[]HypervisorResourceRefResponseModel, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) SetNetworks(v []HypervisorResourceRefResponseModel)`

SetNetworks sets Networks field to given value.


### GetStorage

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) GetStorage() []HypervisorStorageResourceResponseModel`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) GetStorageOk() (*[]HypervisorStorageResourceResponseModel, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) SetStorage(v []HypervisorStorageResourceResponseModel)`

SetStorage sets Storage field to given value.


### GetTemporaryStorage

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) GetTemporaryStorage() []HypervisorStorageResourceResponseModel`

GetTemporaryStorage returns the TemporaryStorage field if non-nil, zero value otherwise.

### GetTemporaryStorageOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) GetTemporaryStorageOk() (*[]HypervisorStorageResourceResponseModel, bool)`

GetTemporaryStorageOk returns a tuple with the TemporaryStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryStorage

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) SetTemporaryStorage(v []HypervisorStorageResourceResponseModel)`

SetTemporaryStorage sets TemporaryStorage field to given value.


### GetUseLocalStorageCaching

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) GetUseLocalStorageCaching() bool`

GetUseLocalStorageCaching returns the UseLocalStorageCaching field if non-nil, zero value otherwise.

### GetUseLocalStorageCachingOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) GetUseLocalStorageCachingOk() (*bool, bool)`

GetUseLocalStorageCachingOk returns a tuple with the UseLocalStorageCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocalStorageCaching

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) SetUseLocalStorageCaching(v bool)`

SetUseLocalStorageCaching sets UseLocalStorageCaching field to given value.

### HasUseLocalStorageCaching

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) HasUseLocalStorageCaching() bool`

HasUseLocalStorageCaching returns a boolean if a field has been set.

### GetCustomProperties

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetPersonalvDiskStorage

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) GetPersonalvDiskStorage() []HypervisorStorageResourceResponseModel`

GetPersonalvDiskStorage returns the PersonalvDiskStorage field if non-nil, zero value otherwise.

### GetPersonalvDiskStorageOk

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) GetPersonalvDiskStorageOk() (*[]HypervisorStorageResourceResponseModel, bool)`

GetPersonalvDiskStorageOk returns a tuple with the PersonalvDiskStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalvDiskStorage

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) SetPersonalvDiskStorage(v []HypervisorStorageResourceResponseModel)`

SetPersonalvDiskStorage sets PersonalvDiskStorage field to given value.

### HasPersonalvDiskStorage

`func (o *HypervisorResourcePoolTraditionalDetailResponseModelAllOf) HasPersonalvDiskStorage() bool`

HasPersonalvDiskStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


