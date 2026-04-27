# AddVnetPeeringModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name to assign the peering | 
**SubscriptionId** | **string** | ID of the Azure subscription that contains the user&#39;s VNet | 
**ResourceGroup** | **string** | Name of the resource group that contain&#39;s the customer&#39;s VNET | 
**Vnet** | **string** | Name of the vnet to peer | 
**Subnet** | **string** | Subnet to assign to the Vnet that is created | 
**SubnetMask** | **int32** | Mask for the created subnet | 
**UseGateway** | Pointer to **bool** | Indicates if the peered vnet will be using the customers gateway | [optional] 
**ManagedVnetRegion** | Pointer to **NullableString** | The region the Citrix managed vnet will be located in.  If unspecified, assume that the customer wants to peer to the region where the vnet is located in | [optional] 
**CspCustomerId** | Pointer to **NullableString** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 
**CspSiteId** | Pointer to **NullableString** | Name of tenant customer ID if partner-tenant relationship exists otherwise null | [optional] 
**Routes** | Pointer to [**[]AzureRoute**](AzureRoute.md) | A list of user defined routes if any | [optional] 
**ManagedSubscriptionId** | Pointer to **NullableString** | ID of the managed subscription the peering should be added to. | [optional] 
**DnsServers** | Pointer to **[]string** | An array containing the dns servers | [optional] 
**DisableRoutePropagation** | Pointer to **bool** | Indicates if route propagation in the route table should be disabled (option is valid only if customer is using gateway). | [optional] 
**NatGatewayConfig** | Pointer to [**NullableNatGatewayModel**](NatGatewayModel.md) | The NAT config. | [optional] 
**VirtualHubId** | Pointer to **NullableString** | Azure Resource ID of the Virtual Hub (for VWAN connections).  If populated, creates a VWAN hub connection instead of traditional VNet peering.  If null/empty, creates traditional bidirectional VNet peering. | [optional] 
**VirtualHubResourceGroup** | Pointer to **NullableString** | Resource group containing the Virtual Hub (for VWAN connections).  Only used when VirtualHubId is populated. | [optional] 
**VirtualHubName** | Pointer to **NullableString** | Name of the Virtual Hub (for VWAN connections).  Only used when VirtualHubId is populated. | [optional] 
**EnableInternetSecurity** | Pointer to **bool** | Enable internet security for VWAN connection (forces traffic through Azure Firewall).  Only used when VirtualHubId is populated. | [optional] 
**RoutingConfiguration** | Pointer to **NullableString** | Routing configuration for VWAN connection (JSON string).  Only used when VirtualHubId is populated. | [optional] 

## Methods

### NewAddVnetPeeringModel

`func NewAddVnetPeeringModel(name string, subscriptionId string, resourceGroup string, vnet string, subnet string, subnetMask int32, ) *AddVnetPeeringModel`

NewAddVnetPeeringModel instantiates a new AddVnetPeeringModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVnetPeeringModelWithDefaults

`func NewAddVnetPeeringModelWithDefaults() *AddVnetPeeringModel`

NewAddVnetPeeringModelWithDefaults instantiates a new AddVnetPeeringModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddVnetPeeringModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddVnetPeeringModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddVnetPeeringModel) SetName(v string)`

SetName sets Name field to given value.


### GetSubscriptionId

`func (o *AddVnetPeeringModel) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AddVnetPeeringModel) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AddVnetPeeringModel) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetResourceGroup

`func (o *AddVnetPeeringModel) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *AddVnetPeeringModel) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *AddVnetPeeringModel) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.


### GetVnet

`func (o *AddVnetPeeringModel) GetVnet() string`

GetVnet returns the Vnet field if non-nil, zero value otherwise.

### GetVnetOk

`func (o *AddVnetPeeringModel) GetVnetOk() (*string, bool)`

GetVnetOk returns a tuple with the Vnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnet

`func (o *AddVnetPeeringModel) SetVnet(v string)`

SetVnet sets Vnet field to given value.


### GetSubnet

`func (o *AddVnetPeeringModel) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *AddVnetPeeringModel) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *AddVnetPeeringModel) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetSubnetMask

`func (o *AddVnetPeeringModel) GetSubnetMask() int32`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *AddVnetPeeringModel) GetSubnetMaskOk() (*int32, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *AddVnetPeeringModel) SetSubnetMask(v int32)`

SetSubnetMask sets SubnetMask field to given value.


### GetUseGateway

`func (o *AddVnetPeeringModel) GetUseGateway() bool`

GetUseGateway returns the UseGateway field if non-nil, zero value otherwise.

### GetUseGatewayOk

`func (o *AddVnetPeeringModel) GetUseGatewayOk() (*bool, bool)`

GetUseGatewayOk returns a tuple with the UseGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGateway

`func (o *AddVnetPeeringModel) SetUseGateway(v bool)`

SetUseGateway sets UseGateway field to given value.

### HasUseGateway

`func (o *AddVnetPeeringModel) HasUseGateway() bool`

HasUseGateway returns a boolean if a field has been set.

### GetManagedVnetRegion

`func (o *AddVnetPeeringModel) GetManagedVnetRegion() string`

GetManagedVnetRegion returns the ManagedVnetRegion field if non-nil, zero value otherwise.

### GetManagedVnetRegionOk

`func (o *AddVnetPeeringModel) GetManagedVnetRegionOk() (*string, bool)`

GetManagedVnetRegionOk returns a tuple with the ManagedVnetRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedVnetRegion

`func (o *AddVnetPeeringModel) SetManagedVnetRegion(v string)`

SetManagedVnetRegion sets ManagedVnetRegion field to given value.

### HasManagedVnetRegion

`func (o *AddVnetPeeringModel) HasManagedVnetRegion() bool`

HasManagedVnetRegion returns a boolean if a field has been set.

### SetManagedVnetRegionNil

`func (o *AddVnetPeeringModel) SetManagedVnetRegionNil(b bool)`

 SetManagedVnetRegionNil sets the value for ManagedVnetRegion to be an explicit nil

### UnsetManagedVnetRegion
`func (o *AddVnetPeeringModel) UnsetManagedVnetRegion()`

UnsetManagedVnetRegion ensures that no value is present for ManagedVnetRegion, not even an explicit nil
### GetCspCustomerId

`func (o *AddVnetPeeringModel) GetCspCustomerId() string`

GetCspCustomerId returns the CspCustomerId field if non-nil, zero value otherwise.

### GetCspCustomerIdOk

`func (o *AddVnetPeeringModel) GetCspCustomerIdOk() (*string, bool)`

GetCspCustomerIdOk returns a tuple with the CspCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomerId

`func (o *AddVnetPeeringModel) SetCspCustomerId(v string)`

SetCspCustomerId sets CspCustomerId field to given value.

### HasCspCustomerId

`func (o *AddVnetPeeringModel) HasCspCustomerId() bool`

HasCspCustomerId returns a boolean if a field has been set.

### SetCspCustomerIdNil

`func (o *AddVnetPeeringModel) SetCspCustomerIdNil(b bool)`

 SetCspCustomerIdNil sets the value for CspCustomerId to be an explicit nil

### UnsetCspCustomerId
`func (o *AddVnetPeeringModel) UnsetCspCustomerId()`

UnsetCspCustomerId ensures that no value is present for CspCustomerId, not even an explicit nil
### GetCspSiteId

`func (o *AddVnetPeeringModel) GetCspSiteId() string`

GetCspSiteId returns the CspSiteId field if non-nil, zero value otherwise.

### GetCspSiteIdOk

`func (o *AddVnetPeeringModel) GetCspSiteIdOk() (*string, bool)`

GetCspSiteIdOk returns a tuple with the CspSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspSiteId

`func (o *AddVnetPeeringModel) SetCspSiteId(v string)`

SetCspSiteId sets CspSiteId field to given value.

### HasCspSiteId

`func (o *AddVnetPeeringModel) HasCspSiteId() bool`

HasCspSiteId returns a boolean if a field has been set.

### SetCspSiteIdNil

`func (o *AddVnetPeeringModel) SetCspSiteIdNil(b bool)`

 SetCspSiteIdNil sets the value for CspSiteId to be an explicit nil

### UnsetCspSiteId
`func (o *AddVnetPeeringModel) UnsetCspSiteId()`

UnsetCspSiteId ensures that no value is present for CspSiteId, not even an explicit nil
### GetRoutes

`func (o *AddVnetPeeringModel) GetRoutes() []AzureRoute`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *AddVnetPeeringModel) GetRoutesOk() (*[]AzureRoute, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *AddVnetPeeringModel) SetRoutes(v []AzureRoute)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *AddVnetPeeringModel) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### SetRoutesNil

`func (o *AddVnetPeeringModel) SetRoutesNil(b bool)`

 SetRoutesNil sets the value for Routes to be an explicit nil

### UnsetRoutes
`func (o *AddVnetPeeringModel) UnsetRoutes()`

UnsetRoutes ensures that no value is present for Routes, not even an explicit nil
### GetManagedSubscriptionId

`func (o *AddVnetPeeringModel) GetManagedSubscriptionId() string`

GetManagedSubscriptionId returns the ManagedSubscriptionId field if non-nil, zero value otherwise.

### GetManagedSubscriptionIdOk

`func (o *AddVnetPeeringModel) GetManagedSubscriptionIdOk() (*string, bool)`

GetManagedSubscriptionIdOk returns a tuple with the ManagedSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedSubscriptionId

`func (o *AddVnetPeeringModel) SetManagedSubscriptionId(v string)`

SetManagedSubscriptionId sets ManagedSubscriptionId field to given value.

### HasManagedSubscriptionId

`func (o *AddVnetPeeringModel) HasManagedSubscriptionId() bool`

HasManagedSubscriptionId returns a boolean if a field has been set.

### SetManagedSubscriptionIdNil

`func (o *AddVnetPeeringModel) SetManagedSubscriptionIdNil(b bool)`

 SetManagedSubscriptionIdNil sets the value for ManagedSubscriptionId to be an explicit nil

### UnsetManagedSubscriptionId
`func (o *AddVnetPeeringModel) UnsetManagedSubscriptionId()`

UnsetManagedSubscriptionId ensures that no value is present for ManagedSubscriptionId, not even an explicit nil
### GetDnsServers

`func (o *AddVnetPeeringModel) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *AddVnetPeeringModel) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *AddVnetPeeringModel) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *AddVnetPeeringModel) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### SetDnsServersNil

`func (o *AddVnetPeeringModel) SetDnsServersNil(b bool)`

 SetDnsServersNil sets the value for DnsServers to be an explicit nil

### UnsetDnsServers
`func (o *AddVnetPeeringModel) UnsetDnsServers()`

UnsetDnsServers ensures that no value is present for DnsServers, not even an explicit nil
### GetDisableRoutePropagation

`func (o *AddVnetPeeringModel) GetDisableRoutePropagation() bool`

GetDisableRoutePropagation returns the DisableRoutePropagation field if non-nil, zero value otherwise.

### GetDisableRoutePropagationOk

`func (o *AddVnetPeeringModel) GetDisableRoutePropagationOk() (*bool, bool)`

GetDisableRoutePropagationOk returns a tuple with the DisableRoutePropagation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRoutePropagation

`func (o *AddVnetPeeringModel) SetDisableRoutePropagation(v bool)`

SetDisableRoutePropagation sets DisableRoutePropagation field to given value.

### HasDisableRoutePropagation

`func (o *AddVnetPeeringModel) HasDisableRoutePropagation() bool`

HasDisableRoutePropagation returns a boolean if a field has been set.

### GetNatGatewayConfig

`func (o *AddVnetPeeringModel) GetNatGatewayConfig() NatGatewayModel`

GetNatGatewayConfig returns the NatGatewayConfig field if non-nil, zero value otherwise.

### GetNatGatewayConfigOk

`func (o *AddVnetPeeringModel) GetNatGatewayConfigOk() (*NatGatewayModel, bool)`

GetNatGatewayConfigOk returns a tuple with the NatGatewayConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatGatewayConfig

`func (o *AddVnetPeeringModel) SetNatGatewayConfig(v NatGatewayModel)`

SetNatGatewayConfig sets NatGatewayConfig field to given value.

### HasNatGatewayConfig

`func (o *AddVnetPeeringModel) HasNatGatewayConfig() bool`

HasNatGatewayConfig returns a boolean if a field has been set.

### SetNatGatewayConfigNil

`func (o *AddVnetPeeringModel) SetNatGatewayConfigNil(b bool)`

 SetNatGatewayConfigNil sets the value for NatGatewayConfig to be an explicit nil

### UnsetNatGatewayConfig
`func (o *AddVnetPeeringModel) UnsetNatGatewayConfig()`

UnsetNatGatewayConfig ensures that no value is present for NatGatewayConfig, not even an explicit nil
### GetVirtualHubId

`func (o *AddVnetPeeringModel) GetVirtualHubId() string`

GetVirtualHubId returns the VirtualHubId field if non-nil, zero value otherwise.

### GetVirtualHubIdOk

`func (o *AddVnetPeeringModel) GetVirtualHubIdOk() (*string, bool)`

GetVirtualHubIdOk returns a tuple with the VirtualHubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualHubId

`func (o *AddVnetPeeringModel) SetVirtualHubId(v string)`

SetVirtualHubId sets VirtualHubId field to given value.

### HasVirtualHubId

`func (o *AddVnetPeeringModel) HasVirtualHubId() bool`

HasVirtualHubId returns a boolean if a field has been set.

### SetVirtualHubIdNil

`func (o *AddVnetPeeringModel) SetVirtualHubIdNil(b bool)`

 SetVirtualHubIdNil sets the value for VirtualHubId to be an explicit nil

### UnsetVirtualHubId
`func (o *AddVnetPeeringModel) UnsetVirtualHubId()`

UnsetVirtualHubId ensures that no value is present for VirtualHubId, not even an explicit nil
### GetVirtualHubResourceGroup

`func (o *AddVnetPeeringModel) GetVirtualHubResourceGroup() string`

GetVirtualHubResourceGroup returns the VirtualHubResourceGroup field if non-nil, zero value otherwise.

### GetVirtualHubResourceGroupOk

`func (o *AddVnetPeeringModel) GetVirtualHubResourceGroupOk() (*string, bool)`

GetVirtualHubResourceGroupOk returns a tuple with the VirtualHubResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualHubResourceGroup

`func (o *AddVnetPeeringModel) SetVirtualHubResourceGroup(v string)`

SetVirtualHubResourceGroup sets VirtualHubResourceGroup field to given value.

### HasVirtualHubResourceGroup

`func (o *AddVnetPeeringModel) HasVirtualHubResourceGroup() bool`

HasVirtualHubResourceGroup returns a boolean if a field has been set.

### SetVirtualHubResourceGroupNil

`func (o *AddVnetPeeringModel) SetVirtualHubResourceGroupNil(b bool)`

 SetVirtualHubResourceGroupNil sets the value for VirtualHubResourceGroup to be an explicit nil

### UnsetVirtualHubResourceGroup
`func (o *AddVnetPeeringModel) UnsetVirtualHubResourceGroup()`

UnsetVirtualHubResourceGroup ensures that no value is present for VirtualHubResourceGroup, not even an explicit nil
### GetVirtualHubName

`func (o *AddVnetPeeringModel) GetVirtualHubName() string`

GetVirtualHubName returns the VirtualHubName field if non-nil, zero value otherwise.

### GetVirtualHubNameOk

`func (o *AddVnetPeeringModel) GetVirtualHubNameOk() (*string, bool)`

GetVirtualHubNameOk returns a tuple with the VirtualHubName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualHubName

`func (o *AddVnetPeeringModel) SetVirtualHubName(v string)`

SetVirtualHubName sets VirtualHubName field to given value.

### HasVirtualHubName

`func (o *AddVnetPeeringModel) HasVirtualHubName() bool`

HasVirtualHubName returns a boolean if a field has been set.

### SetVirtualHubNameNil

`func (o *AddVnetPeeringModel) SetVirtualHubNameNil(b bool)`

 SetVirtualHubNameNil sets the value for VirtualHubName to be an explicit nil

### UnsetVirtualHubName
`func (o *AddVnetPeeringModel) UnsetVirtualHubName()`

UnsetVirtualHubName ensures that no value is present for VirtualHubName, not even an explicit nil
### GetEnableInternetSecurity

`func (o *AddVnetPeeringModel) GetEnableInternetSecurity() bool`

GetEnableInternetSecurity returns the EnableInternetSecurity field if non-nil, zero value otherwise.

### GetEnableInternetSecurityOk

`func (o *AddVnetPeeringModel) GetEnableInternetSecurityOk() (*bool, bool)`

GetEnableInternetSecurityOk returns a tuple with the EnableInternetSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInternetSecurity

`func (o *AddVnetPeeringModel) SetEnableInternetSecurity(v bool)`

SetEnableInternetSecurity sets EnableInternetSecurity field to given value.

### HasEnableInternetSecurity

`func (o *AddVnetPeeringModel) HasEnableInternetSecurity() bool`

HasEnableInternetSecurity returns a boolean if a field has been set.

### GetRoutingConfiguration

`func (o *AddVnetPeeringModel) GetRoutingConfiguration() string`

GetRoutingConfiguration returns the RoutingConfiguration field if non-nil, zero value otherwise.

### GetRoutingConfigurationOk

`func (o *AddVnetPeeringModel) GetRoutingConfigurationOk() (*string, bool)`

GetRoutingConfigurationOk returns a tuple with the RoutingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingConfiguration

`func (o *AddVnetPeeringModel) SetRoutingConfiguration(v string)`

SetRoutingConfiguration sets RoutingConfiguration field to given value.

### HasRoutingConfiguration

`func (o *AddVnetPeeringModel) HasRoutingConfiguration() bool`

HasRoutingConfiguration returns a boolean if a field has been set.

### SetRoutingConfigurationNil

`func (o *AddVnetPeeringModel) SetRoutingConfigurationNil(b bool)`

 SetRoutingConfigurationNil sets the value for RoutingConfiguration to be an explicit nil

### UnsetRoutingConfiguration
`func (o *AddVnetPeeringModel) UnsetRoutingConfiguration()`

UnsetRoutingConfiguration ensures that no value is present for RoutingConfiguration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


