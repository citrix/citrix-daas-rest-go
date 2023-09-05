# HypervisorResourcePoolAWSDetailResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualPrivateCloud** | [**HypervisorResourcePoolAWSDetailResponseModelAllOfVirtualPrivateCloud**](HypervisorResourcePoolAWSDetailResponseModelAllOfVirtualPrivateCloud.md) |  | 
**AvailabilityZone** | [**HypervisorResourcePoolAWSDetailResponseModelAllOfAvailabilityZone**](HypervisorResourcePoolAWSDetailResponseModelAllOfAvailabilityZone.md) |  | 
**Networks** | [**[]HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) | List of networks that may be used within the resource pool. | 

## Methods

### NewHypervisorResourcePoolAWSDetailResponseModelAllOf

`func NewHypervisorResourcePoolAWSDetailResponseModelAllOf(virtualPrivateCloud HypervisorResourcePoolAWSDetailResponseModelAllOfVirtualPrivateCloud, availabilityZone HypervisorResourcePoolAWSDetailResponseModelAllOfAvailabilityZone, networks []HypervisorResourceRefResponseModel, ) *HypervisorResourcePoolAWSDetailResponseModelAllOf`

NewHypervisorResourcePoolAWSDetailResponseModelAllOf instantiates a new HypervisorResourcePoolAWSDetailResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorResourcePoolAWSDetailResponseModelAllOfWithDefaults

`func NewHypervisorResourcePoolAWSDetailResponseModelAllOfWithDefaults() *HypervisorResourcePoolAWSDetailResponseModelAllOf`

NewHypervisorResourcePoolAWSDetailResponseModelAllOfWithDefaults instantiates a new HypervisorResourcePoolAWSDetailResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualPrivateCloud

`func (o *HypervisorResourcePoolAWSDetailResponseModelAllOf) GetVirtualPrivateCloud() HypervisorResourcePoolAWSDetailResponseModelAllOfVirtualPrivateCloud`

GetVirtualPrivateCloud returns the VirtualPrivateCloud field if non-nil, zero value otherwise.

### GetVirtualPrivateCloudOk

`func (o *HypervisorResourcePoolAWSDetailResponseModelAllOf) GetVirtualPrivateCloudOk() (*HypervisorResourcePoolAWSDetailResponseModelAllOfVirtualPrivateCloud, bool)`

GetVirtualPrivateCloudOk returns a tuple with the VirtualPrivateCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPrivateCloud

`func (o *HypervisorResourcePoolAWSDetailResponseModelAllOf) SetVirtualPrivateCloud(v HypervisorResourcePoolAWSDetailResponseModelAllOfVirtualPrivateCloud)`

SetVirtualPrivateCloud sets VirtualPrivateCloud field to given value.


### GetAvailabilityZone

`func (o *HypervisorResourcePoolAWSDetailResponseModelAllOf) GetAvailabilityZone() HypervisorResourcePoolAWSDetailResponseModelAllOfAvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *HypervisorResourcePoolAWSDetailResponseModelAllOf) GetAvailabilityZoneOk() (*HypervisorResourcePoolAWSDetailResponseModelAllOfAvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *HypervisorResourcePoolAWSDetailResponseModelAllOf) SetAvailabilityZone(v HypervisorResourcePoolAWSDetailResponseModelAllOfAvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetNetworks

`func (o *HypervisorResourcePoolAWSDetailResponseModelAllOf) GetNetworks() []HypervisorResourceRefResponseModel`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *HypervisorResourcePoolAWSDetailResponseModelAllOf) GetNetworksOk() (*[]HypervisorResourceRefResponseModel, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *HypervisorResourcePoolAWSDetailResponseModelAllOf) SetNetworks(v []HypervisorResourceRefResponseModel)`

SetNetworks sets Networks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


