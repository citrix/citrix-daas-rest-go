# HypervisorResourcePoolDetailResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GpuTypes** | Pointer to [**[]HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) | GPU types available in the resource pool.  Only applicable to hypervisors that support GPU types. | [optional] 
**ConnectionType** | [**HypervisorConnectionType**](HypervisorConnectionType.md) |  | 
**UsesExplicitStorage** | Pointer to **bool** | If the hypervisor resource pool use ExplicitStorage. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | Metadata for hypervisor resource pool.  | [optional] 

## Methods

### NewHypervisorResourcePoolDetailResponseModelAllOf

`func NewHypervisorResourcePoolDetailResponseModelAllOf(connectionType HypervisorConnectionType, ) *HypervisorResourcePoolDetailResponseModelAllOf`

NewHypervisorResourcePoolDetailResponseModelAllOf instantiates a new HypervisorResourcePoolDetailResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorResourcePoolDetailResponseModelAllOfWithDefaults

`func NewHypervisorResourcePoolDetailResponseModelAllOfWithDefaults() *HypervisorResourcePoolDetailResponseModelAllOf`

NewHypervisorResourcePoolDetailResponseModelAllOfWithDefaults instantiates a new HypervisorResourcePoolDetailResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpuTypes

`func (o *HypervisorResourcePoolDetailResponseModelAllOf) GetGpuTypes() []HypervisorResourceRefResponseModel`

GetGpuTypes returns the GpuTypes field if non-nil, zero value otherwise.

### GetGpuTypesOk

`func (o *HypervisorResourcePoolDetailResponseModelAllOf) GetGpuTypesOk() (*[]HypervisorResourceRefResponseModel, bool)`

GetGpuTypesOk returns a tuple with the GpuTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuTypes

`func (o *HypervisorResourcePoolDetailResponseModelAllOf) SetGpuTypes(v []HypervisorResourceRefResponseModel)`

SetGpuTypes sets GpuTypes field to given value.

### HasGpuTypes

`func (o *HypervisorResourcePoolDetailResponseModelAllOf) HasGpuTypes() bool`

HasGpuTypes returns a boolean if a field has been set.

### GetConnectionType

`func (o *HypervisorResourcePoolDetailResponseModelAllOf) GetConnectionType() HypervisorConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *HypervisorResourcePoolDetailResponseModelAllOf) GetConnectionTypeOk() (*HypervisorConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *HypervisorResourcePoolDetailResponseModelAllOf) SetConnectionType(v HypervisorConnectionType)`

SetConnectionType sets ConnectionType field to given value.


### GetUsesExplicitStorage

`func (o *HypervisorResourcePoolDetailResponseModelAllOf) GetUsesExplicitStorage() bool`

GetUsesExplicitStorage returns the UsesExplicitStorage field if non-nil, zero value otherwise.

### GetUsesExplicitStorageOk

`func (o *HypervisorResourcePoolDetailResponseModelAllOf) GetUsesExplicitStorageOk() (*bool, bool)`

GetUsesExplicitStorageOk returns a tuple with the UsesExplicitStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesExplicitStorage

`func (o *HypervisorResourcePoolDetailResponseModelAllOf) SetUsesExplicitStorage(v bool)`

SetUsesExplicitStorage sets UsesExplicitStorage field to given value.

### HasUsesExplicitStorage

`func (o *HypervisorResourcePoolDetailResponseModelAllOf) HasUsesExplicitStorage() bool`

HasUsesExplicitStorage returns a boolean if a field has been set.

### GetMetadata

`func (o *HypervisorResourcePoolDetailResponseModelAllOf) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HypervisorResourcePoolDetailResponseModelAllOf) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HypervisorResourcePoolDetailResponseModelAllOf) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HypervisorResourcePoolDetailResponseModelAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


