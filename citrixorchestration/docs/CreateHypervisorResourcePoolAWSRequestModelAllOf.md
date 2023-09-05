# CreateHypervisorResourcePoolAWSRequestModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualPrivateCloud** | **string** | AWS Virtual Private Cloud (VPC) resource which the resource pool is connected to.  Required. | 
**AvailabilityZone** | **string** | Path to the availability zone resource to use for provisioning operations in this resource pool.  Required. | 
**Networks** | **[]string** | Path to the network resource(s) that are available for provisioning operations in this resource pool.  At least one is required. | 

## Methods

### NewCreateHypervisorResourcePoolAWSRequestModelAllOf

`func NewCreateHypervisorResourcePoolAWSRequestModelAllOf(virtualPrivateCloud string, availabilityZone string, networks []string, ) *CreateHypervisorResourcePoolAWSRequestModelAllOf`

NewCreateHypervisorResourcePoolAWSRequestModelAllOf instantiates a new CreateHypervisorResourcePoolAWSRequestModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHypervisorResourcePoolAWSRequestModelAllOfWithDefaults

`func NewCreateHypervisorResourcePoolAWSRequestModelAllOfWithDefaults() *CreateHypervisorResourcePoolAWSRequestModelAllOf`

NewCreateHypervisorResourcePoolAWSRequestModelAllOfWithDefaults instantiates a new CreateHypervisorResourcePoolAWSRequestModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualPrivateCloud

`func (o *CreateHypervisorResourcePoolAWSRequestModelAllOf) GetVirtualPrivateCloud() string`

GetVirtualPrivateCloud returns the VirtualPrivateCloud field if non-nil, zero value otherwise.

### GetVirtualPrivateCloudOk

`func (o *CreateHypervisorResourcePoolAWSRequestModelAllOf) GetVirtualPrivateCloudOk() (*string, bool)`

GetVirtualPrivateCloudOk returns a tuple with the VirtualPrivateCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPrivateCloud

`func (o *CreateHypervisorResourcePoolAWSRequestModelAllOf) SetVirtualPrivateCloud(v string)`

SetVirtualPrivateCloud sets VirtualPrivateCloud field to given value.


### GetAvailabilityZone

`func (o *CreateHypervisorResourcePoolAWSRequestModelAllOf) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreateHypervisorResourcePoolAWSRequestModelAllOf) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreateHypervisorResourcePoolAWSRequestModelAllOf) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetNetworks

`func (o *CreateHypervisorResourcePoolAWSRequestModelAllOf) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *CreateHypervisorResourcePoolAWSRequestModelAllOf) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *CreateHypervisorResourcePoolAWSRequestModelAllOf) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


