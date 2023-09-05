# CreateHypervisorResourcePoolTraditionalRequestModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RootPath** | Pointer to **string** | Root path of the resources on the hypervisor which should be included in the resource pool.  Required. | [optional] 
**Networks** | **[]string** | Path to the network resource(s) that are available for provisioning operations in this resource pool.  At least one is required. | 
**UseLocalStorageCaching** | Pointer to **bool** | Indicates whether local storage on the hypervisor will be used for caching purposes. Not all hypervisor types support this.  Defaults to &#x60;false&#x60;. | [optional] [default to false]
**Storage** | Pointer to **[]string** | Path to the storage resource(s) that are available for provisioning operations in this resource pool.  Required for some hypervisor types. | [optional] 
**PersonalvDiskStorage** | Pointer to **[]string** | Path to the storage resource(s) that are available for provisioning operations in this resource pool.  Required for some hypervisor types. | [optional] 
**TemporaryStorage** | Pointer to **[]string** | Path to the storage resource(s) that are used for temporary operations in this resource pool.  Required for some hypervisor types. | [optional] 
**CustomProperties** | Pointer to **string** | Custom properties.  Optional.  If not specified, will not be changed.  Only used for hypervisors of type Custom. | [optional] 

## Methods

### NewCreateHypervisorResourcePoolTraditionalRequestModelAllOf

`func NewCreateHypervisorResourcePoolTraditionalRequestModelAllOf(networks []string, ) *CreateHypervisorResourcePoolTraditionalRequestModelAllOf`

NewCreateHypervisorResourcePoolTraditionalRequestModelAllOf instantiates a new CreateHypervisorResourcePoolTraditionalRequestModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHypervisorResourcePoolTraditionalRequestModelAllOfWithDefaults

`func NewCreateHypervisorResourcePoolTraditionalRequestModelAllOfWithDefaults() *CreateHypervisorResourcePoolTraditionalRequestModelAllOf`

NewCreateHypervisorResourcePoolTraditionalRequestModelAllOfWithDefaults instantiates a new CreateHypervisorResourcePoolTraditionalRequestModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRootPath

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) GetRootPath() string`

GetRootPath returns the RootPath field if non-nil, zero value otherwise.

### GetRootPathOk

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) GetRootPathOk() (*string, bool)`

GetRootPathOk returns a tuple with the RootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPath

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) SetRootPath(v string)`

SetRootPath sets RootPath field to given value.

### HasRootPath

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) HasRootPath() bool`

HasRootPath returns a boolean if a field has been set.

### GetNetworks

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.


### GetUseLocalStorageCaching

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) GetUseLocalStorageCaching() bool`

GetUseLocalStorageCaching returns the UseLocalStorageCaching field if non-nil, zero value otherwise.

### GetUseLocalStorageCachingOk

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) GetUseLocalStorageCachingOk() (*bool, bool)`

GetUseLocalStorageCachingOk returns a tuple with the UseLocalStorageCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocalStorageCaching

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) SetUseLocalStorageCaching(v bool)`

SetUseLocalStorageCaching sets UseLocalStorageCaching field to given value.

### HasUseLocalStorageCaching

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) HasUseLocalStorageCaching() bool`

HasUseLocalStorageCaching returns a boolean if a field has been set.

### GetStorage

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) GetStorage() []string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) GetStorageOk() (*[]string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) SetStorage(v []string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetPersonalvDiskStorage

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) GetPersonalvDiskStorage() []string`

GetPersonalvDiskStorage returns the PersonalvDiskStorage field if non-nil, zero value otherwise.

### GetPersonalvDiskStorageOk

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) GetPersonalvDiskStorageOk() (*[]string, bool)`

GetPersonalvDiskStorageOk returns a tuple with the PersonalvDiskStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalvDiskStorage

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) SetPersonalvDiskStorage(v []string)`

SetPersonalvDiskStorage sets PersonalvDiskStorage field to given value.

### HasPersonalvDiskStorage

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) HasPersonalvDiskStorage() bool`

HasPersonalvDiskStorage returns a boolean if a field has been set.

### GetTemporaryStorage

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) GetTemporaryStorage() []string`

GetTemporaryStorage returns the TemporaryStorage field if non-nil, zero value otherwise.

### GetTemporaryStorageOk

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) GetTemporaryStorageOk() (*[]string, bool)`

GetTemporaryStorageOk returns a tuple with the TemporaryStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryStorage

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) SetTemporaryStorage(v []string)`

SetTemporaryStorage sets TemporaryStorage field to given value.

### HasTemporaryStorage

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) HasTemporaryStorage() bool`

HasTemporaryStorage returns a boolean if a field has been set.

### GetCustomProperties

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) GetCustomProperties() string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) GetCustomPropertiesOk() (*string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) SetCustomProperties(v string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *CreateHypervisorResourcePoolTraditionalRequestModelAllOf) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


