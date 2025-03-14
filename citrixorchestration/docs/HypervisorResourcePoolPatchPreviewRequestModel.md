# HypervisorResourcePoolPatchPreviewRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Networks** | Pointer to **[]string** | Path to the networks/subnets that are available for provisioning operations in this resource pool. At least one is required. | [optional] 
**NetworkIds** | Pointer to **[]string** | Ids of the networks/subnets that are available for provisioning operations, at least one is required. | [optional] 

## Methods

### NewHypervisorResourcePoolPatchPreviewRequestModel

`func NewHypervisorResourcePoolPatchPreviewRequestModel() *HypervisorResourcePoolPatchPreviewRequestModel`

NewHypervisorResourcePoolPatchPreviewRequestModel instantiates a new HypervisorResourcePoolPatchPreviewRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorResourcePoolPatchPreviewRequestModelWithDefaults

`func NewHypervisorResourcePoolPatchPreviewRequestModelWithDefaults() *HypervisorResourcePoolPatchPreviewRequestModel`

NewHypervisorResourcePoolPatchPreviewRequestModelWithDefaults instantiates a new HypervisorResourcePoolPatchPreviewRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworks

`func (o *HypervisorResourcePoolPatchPreviewRequestModel) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *HypervisorResourcePoolPatchPreviewRequestModel) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *HypervisorResourcePoolPatchPreviewRequestModel) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *HypervisorResourcePoolPatchPreviewRequestModel) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### SetNetworksNil

`func (o *HypervisorResourcePoolPatchPreviewRequestModel) SetNetworksNil(b bool)`

 SetNetworksNil sets the value for Networks to be an explicit nil

### UnsetNetworks
`func (o *HypervisorResourcePoolPatchPreviewRequestModel) UnsetNetworks()`

UnsetNetworks ensures that no value is present for Networks, not even an explicit nil
### GetNetworkIds

`func (o *HypervisorResourcePoolPatchPreviewRequestModel) GetNetworkIds() []string`

GetNetworkIds returns the NetworkIds field if non-nil, zero value otherwise.

### GetNetworkIdsOk

`func (o *HypervisorResourcePoolPatchPreviewRequestModel) GetNetworkIdsOk() (*[]string, bool)`

GetNetworkIdsOk returns a tuple with the NetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIds

`func (o *HypervisorResourcePoolPatchPreviewRequestModel) SetNetworkIds(v []string)`

SetNetworkIds sets NetworkIds field to given value.

### HasNetworkIds

`func (o *HypervisorResourcePoolPatchPreviewRequestModel) HasNetworkIds() bool`

HasNetworkIds returns a boolean if a field has been set.

### SetNetworkIdsNil

`func (o *HypervisorResourcePoolPatchPreviewRequestModel) SetNetworkIdsNil(b bool)`

 SetNetworkIdsNil sets the value for NetworkIds to be an explicit nil

### UnsetNetworkIds
`func (o *HypervisorResourcePoolPatchPreviewRequestModel) UnsetNetworkIds()`

UnsetNetworkIds ensures that no value is present for NetworkIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


