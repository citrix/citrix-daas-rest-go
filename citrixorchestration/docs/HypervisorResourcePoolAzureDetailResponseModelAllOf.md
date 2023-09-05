# HypervisorResourcePoolAzureDetailResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | [**HypervisorResourcePoolAzureDetailResponseModelAllOfRegion**](HypervisorResourcePoolAzureDetailResponseModelAllOfRegion.md) |  | 
**VirtualNetwork** | [**HypervisorResourcePoolAzureDetailResponseModelAllOfVirtualNetwork**](HypervisorResourcePoolAzureDetailResponseModelAllOfVirtualNetwork.md) |  | 
**Subnets** | [**[]HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) | List of subnets in the VirtualNetwork that may be used within the resource pool. | 

## Methods

### NewHypervisorResourcePoolAzureDetailResponseModelAllOf

`func NewHypervisorResourcePoolAzureDetailResponseModelAllOf(region HypervisorResourcePoolAzureDetailResponseModelAllOfRegion, virtualNetwork HypervisorResourcePoolAzureDetailResponseModelAllOfVirtualNetwork, subnets []HypervisorResourceRefResponseModel, ) *HypervisorResourcePoolAzureDetailResponseModelAllOf`

NewHypervisorResourcePoolAzureDetailResponseModelAllOf instantiates a new HypervisorResourcePoolAzureDetailResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorResourcePoolAzureDetailResponseModelAllOfWithDefaults

`func NewHypervisorResourcePoolAzureDetailResponseModelAllOfWithDefaults() *HypervisorResourcePoolAzureDetailResponseModelAllOf`

NewHypervisorResourcePoolAzureDetailResponseModelAllOfWithDefaults instantiates a new HypervisorResourcePoolAzureDetailResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *HypervisorResourcePoolAzureDetailResponseModelAllOf) GetRegion() HypervisorResourcePoolAzureDetailResponseModelAllOfRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HypervisorResourcePoolAzureDetailResponseModelAllOf) GetRegionOk() (*HypervisorResourcePoolAzureDetailResponseModelAllOfRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HypervisorResourcePoolAzureDetailResponseModelAllOf) SetRegion(v HypervisorResourcePoolAzureDetailResponseModelAllOfRegion)`

SetRegion sets Region field to given value.


### GetVirtualNetwork

`func (o *HypervisorResourcePoolAzureDetailResponseModelAllOf) GetVirtualNetwork() HypervisorResourcePoolAzureDetailResponseModelAllOfVirtualNetwork`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *HypervisorResourcePoolAzureDetailResponseModelAllOf) GetVirtualNetworkOk() (*HypervisorResourcePoolAzureDetailResponseModelAllOfVirtualNetwork, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *HypervisorResourcePoolAzureDetailResponseModelAllOf) SetVirtualNetwork(v HypervisorResourcePoolAzureDetailResponseModelAllOfVirtualNetwork)`

SetVirtualNetwork sets VirtualNetwork field to given value.


### GetSubnets

`func (o *HypervisorResourcePoolAzureDetailResponseModelAllOf) GetSubnets() []HypervisorResourceRefResponseModel`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *HypervisorResourcePoolAzureDetailResponseModelAllOf) GetSubnetsOk() (*[]HypervisorResourceRefResponseModel, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *HypervisorResourcePoolAzureDetailResponseModelAllOf) SetSubnets(v []HypervisorResourceRefResponseModel)`

SetSubnets sets Subnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


