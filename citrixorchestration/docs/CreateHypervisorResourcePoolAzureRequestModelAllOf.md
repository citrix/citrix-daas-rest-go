# CreateHypervisorResourcePoolAzureRequestModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | Azure region which the resource pool is connected to.  Required. | 
**VirtualNetwork** | **string** | Azure virtual network which the resource pool is connected to. Required. | 
**Subnets** | **[]string** | Path to the subnet(s) that are available for provisioning operations in this resource pool.  At least one is required. | 

## Methods

### NewCreateHypervisorResourcePoolAzureRequestModelAllOf

`func NewCreateHypervisorResourcePoolAzureRequestModelAllOf(region string, virtualNetwork string, subnets []string, ) *CreateHypervisorResourcePoolAzureRequestModelAllOf`

NewCreateHypervisorResourcePoolAzureRequestModelAllOf instantiates a new CreateHypervisorResourcePoolAzureRequestModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHypervisorResourcePoolAzureRequestModelAllOfWithDefaults

`func NewCreateHypervisorResourcePoolAzureRequestModelAllOfWithDefaults() *CreateHypervisorResourcePoolAzureRequestModelAllOf`

NewCreateHypervisorResourcePoolAzureRequestModelAllOfWithDefaults instantiates a new CreateHypervisorResourcePoolAzureRequestModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *CreateHypervisorResourcePoolAzureRequestModelAllOf) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateHypervisorResourcePoolAzureRequestModelAllOf) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateHypervisorResourcePoolAzureRequestModelAllOf) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetVirtualNetwork

`func (o *CreateHypervisorResourcePoolAzureRequestModelAllOf) GetVirtualNetwork() string`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *CreateHypervisorResourcePoolAzureRequestModelAllOf) GetVirtualNetworkOk() (*string, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *CreateHypervisorResourcePoolAzureRequestModelAllOf) SetVirtualNetwork(v string)`

SetVirtualNetwork sets VirtualNetwork field to given value.


### GetSubnets

`func (o *CreateHypervisorResourcePoolAzureRequestModelAllOf) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *CreateHypervisorResourcePoolAzureRequestModelAllOf) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *CreateHypervisorResourcePoolAzureRequestModelAllOf) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


