# HypervisorResourcePoolGcpDetailResponseModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | [**HypervisorResourcePoolGcpDetailResponseModelAllOfProject**](HypervisorResourcePoolGcpDetailResponseModelAllOfProject.md) |  | 
**Region** | [**HypervisorResourcePoolGcpDetailResponseModelAllOfRegion**](HypervisorResourcePoolGcpDetailResponseModelAllOfRegion.md) |  | 
**VirtualPrivateCloud** | [**HypervisorResourcePoolGcpDetailResponseModelAllOfVirtualPrivateCloud**](HypervisorResourcePoolGcpDetailResponseModelAllOfVirtualPrivateCloud.md) |  | 
**Networks** | [**[]HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) | List of networks in the VirtualPrivateCloud that may be used within the resource pool. | 

## Methods

### NewHypervisorResourcePoolGcpDetailResponseModelAllOf

`func NewHypervisorResourcePoolGcpDetailResponseModelAllOf(project HypervisorResourcePoolGcpDetailResponseModelAllOfProject, region HypervisorResourcePoolGcpDetailResponseModelAllOfRegion, virtualPrivateCloud HypervisorResourcePoolGcpDetailResponseModelAllOfVirtualPrivateCloud, networks []HypervisorResourceRefResponseModel, ) *HypervisorResourcePoolGcpDetailResponseModelAllOf`

NewHypervisorResourcePoolGcpDetailResponseModelAllOf instantiates a new HypervisorResourcePoolGcpDetailResponseModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorResourcePoolGcpDetailResponseModelAllOfWithDefaults

`func NewHypervisorResourcePoolGcpDetailResponseModelAllOfWithDefaults() *HypervisorResourcePoolGcpDetailResponseModelAllOf`

NewHypervisorResourcePoolGcpDetailResponseModelAllOfWithDefaults instantiates a new HypervisorResourcePoolGcpDetailResponseModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOf) GetProject() HypervisorResourcePoolGcpDetailResponseModelAllOfProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOf) GetProjectOk() (*HypervisorResourcePoolGcpDetailResponseModelAllOfProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOf) SetProject(v HypervisorResourcePoolGcpDetailResponseModelAllOfProject)`

SetProject sets Project field to given value.


### GetRegion

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOf) GetRegion() HypervisorResourcePoolGcpDetailResponseModelAllOfRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOf) GetRegionOk() (*HypervisorResourcePoolGcpDetailResponseModelAllOfRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOf) SetRegion(v HypervisorResourcePoolGcpDetailResponseModelAllOfRegion)`

SetRegion sets Region field to given value.


### GetVirtualPrivateCloud

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOf) GetVirtualPrivateCloud() HypervisorResourcePoolGcpDetailResponseModelAllOfVirtualPrivateCloud`

GetVirtualPrivateCloud returns the VirtualPrivateCloud field if non-nil, zero value otherwise.

### GetVirtualPrivateCloudOk

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOf) GetVirtualPrivateCloudOk() (*HypervisorResourcePoolGcpDetailResponseModelAllOfVirtualPrivateCloud, bool)`

GetVirtualPrivateCloudOk returns a tuple with the VirtualPrivateCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPrivateCloud

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOf) SetVirtualPrivateCloud(v HypervisorResourcePoolGcpDetailResponseModelAllOfVirtualPrivateCloud)`

SetVirtualPrivateCloud sets VirtualPrivateCloud field to given value.


### GetNetworks

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOf) GetNetworks() []HypervisorResourceRefResponseModel`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOf) GetNetworksOk() (*[]HypervisorResourceRefResponseModel, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *HypervisorResourcePoolGcpDetailResponseModelAllOf) SetNetworks(v []HypervisorResourceRefResponseModel)`

SetNetworks sets Networks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


