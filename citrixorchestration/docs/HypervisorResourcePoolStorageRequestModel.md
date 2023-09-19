# HypervisorResourcePoolStorageRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoragePath** | **string** | Path to storage within the resource pool.  Required. | 
**Superseded** | Pointer to **NullableBool** | Indicates whether the storage has been superseded. Superseded storage may be used for existing virtual machines, but is not used when provisioning new virtual machines.  Optional; default is &#x60;false&#x60;. | [optional] [default to false]

## Methods

### NewHypervisorResourcePoolStorageRequestModel

`func NewHypervisorResourcePoolStorageRequestModel(storagePath string, ) *HypervisorResourcePoolStorageRequestModel`

NewHypervisorResourcePoolStorageRequestModel instantiates a new HypervisorResourcePoolStorageRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorResourcePoolStorageRequestModelWithDefaults

`func NewHypervisorResourcePoolStorageRequestModelWithDefaults() *HypervisorResourcePoolStorageRequestModel`

NewHypervisorResourcePoolStorageRequestModelWithDefaults instantiates a new HypervisorResourcePoolStorageRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoragePath

`func (o *HypervisorResourcePoolStorageRequestModel) GetStoragePath() string`

GetStoragePath returns the StoragePath field if non-nil, zero value otherwise.

### GetStoragePathOk

`func (o *HypervisorResourcePoolStorageRequestModel) GetStoragePathOk() (*string, bool)`

GetStoragePathOk returns a tuple with the StoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePath

`func (o *HypervisorResourcePoolStorageRequestModel) SetStoragePath(v string)`

SetStoragePath sets StoragePath field to given value.


### GetSuperseded

`func (o *HypervisorResourcePoolStorageRequestModel) GetSuperseded() bool`

GetSuperseded returns the Superseded field if non-nil, zero value otherwise.

### GetSupersededOk

`func (o *HypervisorResourcePoolStorageRequestModel) GetSupersededOk() (*bool, bool)`

GetSupersededOk returns a tuple with the Superseded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperseded

`func (o *HypervisorResourcePoolStorageRequestModel) SetSuperseded(v bool)`

SetSuperseded sets Superseded field to given value.

### HasSuperseded

`func (o *HypervisorResourcePoolStorageRequestModel) HasSuperseded() bool`

HasSuperseded returns a boolean if a field has been set.

### SetSupersededNil

`func (o *HypervisorResourcePoolStorageRequestModel) SetSupersededNil(b bool)`

 SetSupersededNil sets the value for Superseded to be an explicit nil

### UnsetSuperseded
`func (o *HypervisorResourcePoolStorageRequestModel) UnsetSuperseded()`

UnsetSuperseded ensures that no value is present for Superseded, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


