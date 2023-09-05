# CreateHypervisorResourcePoolGcpRequestModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | Google Cloud Platform region which the resource pool is connected to.  Required. | 
**VirtualPrivateCloud** | **string** | Google Cloud Platform virtual network which the resource pool is connected to. Required. | 
**Networks** | **[]string** | Path to the subnet(s) that are available for provisioning operations in this resource pool.  At least one is required. | 

## Methods

### NewCreateHypervisorResourcePoolGcpRequestModelAllOf

`func NewCreateHypervisorResourcePoolGcpRequestModelAllOf(region string, virtualPrivateCloud string, networks []string, ) *CreateHypervisorResourcePoolGcpRequestModelAllOf`

NewCreateHypervisorResourcePoolGcpRequestModelAllOf instantiates a new CreateHypervisorResourcePoolGcpRequestModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHypervisorResourcePoolGcpRequestModelAllOfWithDefaults

`func NewCreateHypervisorResourcePoolGcpRequestModelAllOfWithDefaults() *CreateHypervisorResourcePoolGcpRequestModelAllOf`

NewCreateHypervisorResourcePoolGcpRequestModelAllOfWithDefaults instantiates a new CreateHypervisorResourcePoolGcpRequestModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *CreateHypervisorResourcePoolGcpRequestModelAllOf) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateHypervisorResourcePoolGcpRequestModelAllOf) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateHypervisorResourcePoolGcpRequestModelAllOf) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetVirtualPrivateCloud

`func (o *CreateHypervisorResourcePoolGcpRequestModelAllOf) GetVirtualPrivateCloud() string`

GetVirtualPrivateCloud returns the VirtualPrivateCloud field if non-nil, zero value otherwise.

### GetVirtualPrivateCloudOk

`func (o *CreateHypervisorResourcePoolGcpRequestModelAllOf) GetVirtualPrivateCloudOk() (*string, bool)`

GetVirtualPrivateCloudOk returns a tuple with the VirtualPrivateCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPrivateCloud

`func (o *CreateHypervisorResourcePoolGcpRequestModelAllOf) SetVirtualPrivateCloud(v string)`

SetVirtualPrivateCloud sets VirtualPrivateCloud field to given value.


### GetNetworks

`func (o *CreateHypervisorResourcePoolGcpRequestModelAllOf) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *CreateHypervisorResourcePoolGcpRequestModelAllOf) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *CreateHypervisorResourcePoolGcpRequestModelAllOf) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


