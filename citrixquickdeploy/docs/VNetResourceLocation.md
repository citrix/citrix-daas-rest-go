# VNetResourceLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceLocationId** | Pointer to **string** | ID for the Resource Location | [optional] 
**Ports** | Pointer to [**[]NetworkSecurityGroupPort**](NetworkSecurityGroupPort.md) | List of ports | [optional] 
**NetworkSecurityGroup** | Pointer to **string** | Network Security Group Name | [optional] 
**DomainJoined** | Pointer to **string** | Domain name tied to the resource location | [optional] 
**Connectors** | Pointer to **[]string** | The list of connectors configured for the resource location | [optional] 
**VnetPeering** | Pointer to **string** | Vnet Peering associated with the resource location | [optional] 

## Methods

### NewVNetResourceLocation

`func NewVNetResourceLocation() *VNetResourceLocation`

NewVNetResourceLocation instantiates a new VNetResourceLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVNetResourceLocationWithDefaults

`func NewVNetResourceLocationWithDefaults() *VNetResourceLocation`

NewVNetResourceLocationWithDefaults instantiates a new VNetResourceLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceLocationId

`func (o *VNetResourceLocation) GetResourceLocationId() string`

GetResourceLocationId returns the ResourceLocationId field if non-nil, zero value otherwise.

### GetResourceLocationIdOk

`func (o *VNetResourceLocation) GetResourceLocationIdOk() (*string, bool)`

GetResourceLocationIdOk returns a tuple with the ResourceLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLocationId

`func (o *VNetResourceLocation) SetResourceLocationId(v string)`

SetResourceLocationId sets ResourceLocationId field to given value.

### HasResourceLocationId

`func (o *VNetResourceLocation) HasResourceLocationId() bool`

HasResourceLocationId returns a boolean if a field has been set.

### GetPorts

`func (o *VNetResourceLocation) GetPorts() []NetworkSecurityGroupPort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *VNetResourceLocation) GetPortsOk() (*[]NetworkSecurityGroupPort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *VNetResourceLocation) SetPorts(v []NetworkSecurityGroupPort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *VNetResourceLocation) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetNetworkSecurityGroup

`func (o *VNetResourceLocation) GetNetworkSecurityGroup() string`

GetNetworkSecurityGroup returns the NetworkSecurityGroup field if non-nil, zero value otherwise.

### GetNetworkSecurityGroupOk

`func (o *VNetResourceLocation) GetNetworkSecurityGroupOk() (*string, bool)`

GetNetworkSecurityGroupOk returns a tuple with the NetworkSecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSecurityGroup

`func (o *VNetResourceLocation) SetNetworkSecurityGroup(v string)`

SetNetworkSecurityGroup sets NetworkSecurityGroup field to given value.

### HasNetworkSecurityGroup

`func (o *VNetResourceLocation) HasNetworkSecurityGroup() bool`

HasNetworkSecurityGroup returns a boolean if a field has been set.

### GetDomainJoined

`func (o *VNetResourceLocation) GetDomainJoined() string`

GetDomainJoined returns the DomainJoined field if non-nil, zero value otherwise.

### GetDomainJoinedOk

`func (o *VNetResourceLocation) GetDomainJoinedOk() (*string, bool)`

GetDomainJoinedOk returns a tuple with the DomainJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainJoined

`func (o *VNetResourceLocation) SetDomainJoined(v string)`

SetDomainJoined sets DomainJoined field to given value.

### HasDomainJoined

`func (o *VNetResourceLocation) HasDomainJoined() bool`

HasDomainJoined returns a boolean if a field has been set.

### GetConnectors

`func (o *VNetResourceLocation) GetConnectors() []string`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *VNetResourceLocation) GetConnectorsOk() (*[]string, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *VNetResourceLocation) SetConnectors(v []string)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *VNetResourceLocation) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### GetVnetPeering

`func (o *VNetResourceLocation) GetVnetPeering() string`

GetVnetPeering returns the VnetPeering field if non-nil, zero value otherwise.

### GetVnetPeeringOk

`func (o *VNetResourceLocation) GetVnetPeeringOk() (*string, bool)`

GetVnetPeeringOk returns a tuple with the VnetPeering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetPeering

`func (o *VNetResourceLocation) SetVnetPeering(v string)`

SetVnetPeering sets VnetPeering field to given value.

### HasVnetPeering

`func (o *VNetResourceLocation) HasVnetPeering() bool`

HasVnetPeering returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


