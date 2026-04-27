# AddVWanConnectionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name to assign the VWAN connection | 
**Subnet** | **string** | Subnet to assign to the Citrix-managed VNet that will be created | 
**SubnetMask** | **int32** | Mask for the created subnet in the managed VNet | 
**ManagedVnetRegion** | Pointer to **NullableString** | The region where the Citrix-managed VNet will be located.  If unspecified, the Virtual Hub&#39;s location is used. | [optional] 
**CspCustomerId** | Pointer to **NullableString** | Name of tenant customer ID if partner-tenant relationship exists, otherwise null | [optional] 
**CspSiteId** | Pointer to **NullableString** | Name of tenant site ID if partner-tenant relationship exists, otherwise null | [optional] 
**ManagedSubscriptionId** | Pointer to **NullableString** | ID of the managed subscription where the Citrix-managed VNet should be created. | [optional] 
**DnsServers** | Pointer to **[]string** | An array containing the DNS servers for the managed VNet | [optional] 
**Routes** | Pointer to [**[]AzureRoute**](AzureRoute.md) | Gets or sets a list of user-defined routes to add to the managed VNet subnet via a route table. | [optional] 
**DisableRoutePropagation** | Pointer to **bool** | Gets or sets a value indicating whether BGP route propagation is disabled on the route table  attached to the managed VNet subnet.  Only meaningful when Citrix.XenDesktop.Cloud.CatalogCommon.Models.Configuration.ManagedCapacity.AddVWanConnectionModel.Routes is non-empty. | [optional] 
**VirtualHubId** | **string** | Azure Resource ID of the Virtual Hub where the connection will be created.  Format: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Network/virtualHubs/{hubName} | 
**VirtualHubResourceGroup** | **string** | Resource group containing the Virtual Hub.  Must match the resource group in VirtualHubId. | 
**VirtualHubName** | **string** | Name of the Virtual Hub.  Must match the hub name in VirtualHubId. | 
**EnableInternetSecurity** | Pointer to **bool** | Enable internet security for the VWAN connection.  When true, forces all internet-bound traffic through Azure Firewall (if configured in the hub). | [optional] 
**RoutingConfiguration** | Pointer to **NullableString** | Routing configuration for the VWAN connection (JSON string).  Defines custom routing behavior for the connection. | [optional] 

## Methods

### NewAddVWanConnectionModel

`func NewAddVWanConnectionModel(name string, subnet string, subnetMask int32, virtualHubId string, virtualHubResourceGroup string, virtualHubName string, ) *AddVWanConnectionModel`

NewAddVWanConnectionModel instantiates a new AddVWanConnectionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVWanConnectionModelWithDefaults

`func NewAddVWanConnectionModelWithDefaults() *AddVWanConnectionModel`

NewAddVWanConnectionModelWithDefaults instantiates a new AddVWanConnectionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddVWanConnectionModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddVWanConnectionModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddVWanConnectionModel) SetName(v string)`

SetName sets Name field to given value.


### GetSubnet

`func (o *AddVWanConnectionModel) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *AddVWanConnectionModel) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *AddVWanConnectionModel) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetSubnetMask

`func (o *AddVWanConnectionModel) GetSubnetMask() int32`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *AddVWanConnectionModel) GetSubnetMaskOk() (*int32, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *AddVWanConnectionModel) SetSubnetMask(v int32)`

SetSubnetMask sets SubnetMask field to given value.


### GetManagedVnetRegion

`func (o *AddVWanConnectionModel) GetManagedVnetRegion() string`

GetManagedVnetRegion returns the ManagedVnetRegion field if non-nil, zero value otherwise.

### GetManagedVnetRegionOk

`func (o *AddVWanConnectionModel) GetManagedVnetRegionOk() (*string, bool)`

GetManagedVnetRegionOk returns a tuple with the ManagedVnetRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedVnetRegion

`func (o *AddVWanConnectionModel) SetManagedVnetRegion(v string)`

SetManagedVnetRegion sets ManagedVnetRegion field to given value.

### HasManagedVnetRegion

`func (o *AddVWanConnectionModel) HasManagedVnetRegion() bool`

HasManagedVnetRegion returns a boolean if a field has been set.

### SetManagedVnetRegionNil

`func (o *AddVWanConnectionModel) SetManagedVnetRegionNil(b bool)`

 SetManagedVnetRegionNil sets the value for ManagedVnetRegion to be an explicit nil

### UnsetManagedVnetRegion
`func (o *AddVWanConnectionModel) UnsetManagedVnetRegion()`

UnsetManagedVnetRegion ensures that no value is present for ManagedVnetRegion, not even an explicit nil
### GetCspCustomerId

`func (o *AddVWanConnectionModel) GetCspCustomerId() string`

GetCspCustomerId returns the CspCustomerId field if non-nil, zero value otherwise.

### GetCspCustomerIdOk

`func (o *AddVWanConnectionModel) GetCspCustomerIdOk() (*string, bool)`

GetCspCustomerIdOk returns a tuple with the CspCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomerId

`func (o *AddVWanConnectionModel) SetCspCustomerId(v string)`

SetCspCustomerId sets CspCustomerId field to given value.

### HasCspCustomerId

`func (o *AddVWanConnectionModel) HasCspCustomerId() bool`

HasCspCustomerId returns a boolean if a field has been set.

### SetCspCustomerIdNil

`func (o *AddVWanConnectionModel) SetCspCustomerIdNil(b bool)`

 SetCspCustomerIdNil sets the value for CspCustomerId to be an explicit nil

### UnsetCspCustomerId
`func (o *AddVWanConnectionModel) UnsetCspCustomerId()`

UnsetCspCustomerId ensures that no value is present for CspCustomerId, not even an explicit nil
### GetCspSiteId

`func (o *AddVWanConnectionModel) GetCspSiteId() string`

GetCspSiteId returns the CspSiteId field if non-nil, zero value otherwise.

### GetCspSiteIdOk

`func (o *AddVWanConnectionModel) GetCspSiteIdOk() (*string, bool)`

GetCspSiteIdOk returns a tuple with the CspSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspSiteId

`func (o *AddVWanConnectionModel) SetCspSiteId(v string)`

SetCspSiteId sets CspSiteId field to given value.

### HasCspSiteId

`func (o *AddVWanConnectionModel) HasCspSiteId() bool`

HasCspSiteId returns a boolean if a field has been set.

### SetCspSiteIdNil

`func (o *AddVWanConnectionModel) SetCspSiteIdNil(b bool)`

 SetCspSiteIdNil sets the value for CspSiteId to be an explicit nil

### UnsetCspSiteId
`func (o *AddVWanConnectionModel) UnsetCspSiteId()`

UnsetCspSiteId ensures that no value is present for CspSiteId, not even an explicit nil
### GetManagedSubscriptionId

`func (o *AddVWanConnectionModel) GetManagedSubscriptionId() string`

GetManagedSubscriptionId returns the ManagedSubscriptionId field if non-nil, zero value otherwise.

### GetManagedSubscriptionIdOk

`func (o *AddVWanConnectionModel) GetManagedSubscriptionIdOk() (*string, bool)`

GetManagedSubscriptionIdOk returns a tuple with the ManagedSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedSubscriptionId

`func (o *AddVWanConnectionModel) SetManagedSubscriptionId(v string)`

SetManagedSubscriptionId sets ManagedSubscriptionId field to given value.

### HasManagedSubscriptionId

`func (o *AddVWanConnectionModel) HasManagedSubscriptionId() bool`

HasManagedSubscriptionId returns a boolean if a field has been set.

### SetManagedSubscriptionIdNil

`func (o *AddVWanConnectionModel) SetManagedSubscriptionIdNil(b bool)`

 SetManagedSubscriptionIdNil sets the value for ManagedSubscriptionId to be an explicit nil

### UnsetManagedSubscriptionId
`func (o *AddVWanConnectionModel) UnsetManagedSubscriptionId()`

UnsetManagedSubscriptionId ensures that no value is present for ManagedSubscriptionId, not even an explicit nil
### GetDnsServers

`func (o *AddVWanConnectionModel) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *AddVWanConnectionModel) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *AddVWanConnectionModel) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *AddVWanConnectionModel) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### SetDnsServersNil

`func (o *AddVWanConnectionModel) SetDnsServersNil(b bool)`

 SetDnsServersNil sets the value for DnsServers to be an explicit nil

### UnsetDnsServers
`func (o *AddVWanConnectionModel) UnsetDnsServers()`

UnsetDnsServers ensures that no value is present for DnsServers, not even an explicit nil
### GetRoutes

`func (o *AddVWanConnectionModel) GetRoutes() []AzureRoute`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *AddVWanConnectionModel) GetRoutesOk() (*[]AzureRoute, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *AddVWanConnectionModel) SetRoutes(v []AzureRoute)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *AddVWanConnectionModel) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### SetRoutesNil

`func (o *AddVWanConnectionModel) SetRoutesNil(b bool)`

 SetRoutesNil sets the value for Routes to be an explicit nil

### UnsetRoutes
`func (o *AddVWanConnectionModel) UnsetRoutes()`

UnsetRoutes ensures that no value is present for Routes, not even an explicit nil
### GetDisableRoutePropagation

`func (o *AddVWanConnectionModel) GetDisableRoutePropagation() bool`

GetDisableRoutePropagation returns the DisableRoutePropagation field if non-nil, zero value otherwise.

### GetDisableRoutePropagationOk

`func (o *AddVWanConnectionModel) GetDisableRoutePropagationOk() (*bool, bool)`

GetDisableRoutePropagationOk returns a tuple with the DisableRoutePropagation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRoutePropagation

`func (o *AddVWanConnectionModel) SetDisableRoutePropagation(v bool)`

SetDisableRoutePropagation sets DisableRoutePropagation field to given value.

### HasDisableRoutePropagation

`func (o *AddVWanConnectionModel) HasDisableRoutePropagation() bool`

HasDisableRoutePropagation returns a boolean if a field has been set.

### GetVirtualHubId

`func (o *AddVWanConnectionModel) GetVirtualHubId() string`

GetVirtualHubId returns the VirtualHubId field if non-nil, zero value otherwise.

### GetVirtualHubIdOk

`func (o *AddVWanConnectionModel) GetVirtualHubIdOk() (*string, bool)`

GetVirtualHubIdOk returns a tuple with the VirtualHubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualHubId

`func (o *AddVWanConnectionModel) SetVirtualHubId(v string)`

SetVirtualHubId sets VirtualHubId field to given value.


### GetVirtualHubResourceGroup

`func (o *AddVWanConnectionModel) GetVirtualHubResourceGroup() string`

GetVirtualHubResourceGroup returns the VirtualHubResourceGroup field if non-nil, zero value otherwise.

### GetVirtualHubResourceGroupOk

`func (o *AddVWanConnectionModel) GetVirtualHubResourceGroupOk() (*string, bool)`

GetVirtualHubResourceGroupOk returns a tuple with the VirtualHubResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualHubResourceGroup

`func (o *AddVWanConnectionModel) SetVirtualHubResourceGroup(v string)`

SetVirtualHubResourceGroup sets VirtualHubResourceGroup field to given value.


### GetVirtualHubName

`func (o *AddVWanConnectionModel) GetVirtualHubName() string`

GetVirtualHubName returns the VirtualHubName field if non-nil, zero value otherwise.

### GetVirtualHubNameOk

`func (o *AddVWanConnectionModel) GetVirtualHubNameOk() (*string, bool)`

GetVirtualHubNameOk returns a tuple with the VirtualHubName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualHubName

`func (o *AddVWanConnectionModel) SetVirtualHubName(v string)`

SetVirtualHubName sets VirtualHubName field to given value.


### GetEnableInternetSecurity

`func (o *AddVWanConnectionModel) GetEnableInternetSecurity() bool`

GetEnableInternetSecurity returns the EnableInternetSecurity field if non-nil, zero value otherwise.

### GetEnableInternetSecurityOk

`func (o *AddVWanConnectionModel) GetEnableInternetSecurityOk() (*bool, bool)`

GetEnableInternetSecurityOk returns a tuple with the EnableInternetSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInternetSecurity

`func (o *AddVWanConnectionModel) SetEnableInternetSecurity(v bool)`

SetEnableInternetSecurity sets EnableInternetSecurity field to given value.

### HasEnableInternetSecurity

`func (o *AddVWanConnectionModel) HasEnableInternetSecurity() bool`

HasEnableInternetSecurity returns a boolean if a field has been set.

### GetRoutingConfiguration

`func (o *AddVWanConnectionModel) GetRoutingConfiguration() string`

GetRoutingConfiguration returns the RoutingConfiguration field if non-nil, zero value otherwise.

### GetRoutingConfigurationOk

`func (o *AddVWanConnectionModel) GetRoutingConfigurationOk() (*string, bool)`

GetRoutingConfigurationOk returns a tuple with the RoutingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingConfiguration

`func (o *AddVWanConnectionModel) SetRoutingConfiguration(v string)`

SetRoutingConfiguration sets RoutingConfiguration field to given value.

### HasRoutingConfiguration

`func (o *AddVWanConnectionModel) HasRoutingConfiguration() bool`

HasRoutingConfiguration returns a boolean if a field has been set.

### SetRoutingConfigurationNil

`func (o *AddVWanConnectionModel) SetRoutingConfigurationNil(b bool)`

 SetRoutingConfigurationNil sets the value for RoutingConfiguration to be an explicit nil

### UnsetRoutingConfiguration
`func (o *AddVWanConnectionModel) UnsetRoutingConfiguration()`

UnsetRoutingConfiguration ensures that no value is present for RoutingConfiguration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


