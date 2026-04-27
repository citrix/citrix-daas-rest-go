# VnetPeeringModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**VnetPeeringState**](VnetPeeringState.md) | State of the peering | [optional] 
**Error** | Pointer to **NullableString** | Error that occured while peering the vnet | [optional] 
**TransactionId** | Pointer to **NullableString** | ID of the transaction that performed pairing job | [optional] 
**Region** | Pointer to **NullableString** | Azure region where the vnet in the managed subscription is located | [optional] 
**SourceVnetRegion** | Pointer to **NullableString** | Azure region where the originating vnet is located | [optional] 
**IpAddressSpace** | Pointer to **[]string** | The IP Ranges allocated to the peering | [optional] [readonly] 
**NumberOfAddressesInUse** | Pointer to **int32** | Number of IP Addresses that are currently in use in the subnet | [optional] [readonly] 
**IsPeeringActive** | Pointer to **bool** | Indicates if the Peering is active | [optional] 
**CitrixManaged** | Pointer to **bool** | Indicates if the Peering is managed by Citrix | [optional] 
**PeeredVnetId** | Pointer to **NullableString** | Resource ID of the Peered VNET | [optional] 
**SubscriptionId** | Pointer to **NullableString** | ID of the Azure Subscription that the citrix managed vnet is associated with | [optional] [readonly] 
**DnsServers** | Pointer to **[]string** | IP address of the DNS server of this VNet | [optional] 
**CspCustomer** | Pointer to **NullableString** | Indicates that partner-tenant relationship exists if not null | [optional] 
**Catalogs** | Pointer to **[]string** | List of associated catalogs | [optional] 
**Images** | Pointer to **[]string** | List of associated images | [optional] 
**Bastions** | Pointer to **[]string** | List of associated bastions | [optional] 
**UseGateway** | Pointer to **NullableBool** | Indicates if the peered vnet will be using the customers gateway | [optional] 
**DisableRoutePropagation** | Pointer to **NullableBool** | Indicates if route propagation in the route table should be disabled (option is valid only if customer is using gateway). | [optional] 
**NatGatewayConfig** | Pointer to [**NullableNatGatewayModelOverview**](NatGatewayModelOverview.md) | The NAT gateway config | [optional] 
**RemoteTenantId** | Pointer to **NullableString** | The TenantId of the VNet we are peered into | [optional] 
**ResourceId** | Pointer to **NullableString** | The Azure Resource ID of the VNet Peering | [optional] 
**VirtualHubId** | Pointer to **NullableString** | Azure Resource ID of the Virtual Hub (for VWAN connections).  If populated, this is a VWAN connection instead of traditional VNet peering. | [optional] 
**VirtualHubResourceGroup** | Pointer to **NullableString** | Resource group containing the Virtual Hub.  Only used when VirtualHubId is populated. | [optional] 
**VirtualHubName** | Pointer to **NullableString** | Name of the Virtual Hub.  Only used when VirtualHubId is populated. | [optional] 
**EnableInternetSecurity** | Pointer to **NullableBool** | Enable internet security (firewall) for the VWAN connection.  Only used when VirtualHubId is populated. | [optional] 
**RoutingConfiguration** | Pointer to **NullableString** | JSON configuration for routing in the VWAN connection.  Only used when VirtualHubId is populated. | [optional] 
**Id** | Pointer to **NullableString** | ID of the connection | [optional] 
**Type** | Pointer to [**OnPremConnectionType**](OnPremConnectionType.md) | The type of connection | [optional] 
**Name** | Pointer to **NullableString** | Name of the connection | [optional] 
**StartedAt** | Pointer to **NullableTime** | The datetime when the job started | [optional] 
**EstimatedTimeInMinute** | Pointer to **NullableInt32** | Estimated total time for the job to finish | [optional] 

## Methods

### NewVnetPeeringModel

`func NewVnetPeeringModel() *VnetPeeringModel`

NewVnetPeeringModel instantiates a new VnetPeeringModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnetPeeringModelWithDefaults

`func NewVnetPeeringModelWithDefaults() *VnetPeeringModel`

NewVnetPeeringModelWithDefaults instantiates a new VnetPeeringModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *VnetPeeringModel) GetState() VnetPeeringState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VnetPeeringModel) GetStateOk() (*VnetPeeringState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VnetPeeringModel) SetState(v VnetPeeringState)`

SetState sets State field to given value.

### HasState

`func (o *VnetPeeringModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetError

`func (o *VnetPeeringModel) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *VnetPeeringModel) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *VnetPeeringModel) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *VnetPeeringModel) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *VnetPeeringModel) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *VnetPeeringModel) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetTransactionId

`func (o *VnetPeeringModel) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *VnetPeeringModel) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *VnetPeeringModel) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *VnetPeeringModel) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *VnetPeeringModel) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *VnetPeeringModel) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetRegion

`func (o *VnetPeeringModel) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *VnetPeeringModel) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *VnetPeeringModel) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *VnetPeeringModel) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *VnetPeeringModel) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *VnetPeeringModel) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetSourceVnetRegion

`func (o *VnetPeeringModel) GetSourceVnetRegion() string`

GetSourceVnetRegion returns the SourceVnetRegion field if non-nil, zero value otherwise.

### GetSourceVnetRegionOk

`func (o *VnetPeeringModel) GetSourceVnetRegionOk() (*string, bool)`

GetSourceVnetRegionOk returns a tuple with the SourceVnetRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVnetRegion

`func (o *VnetPeeringModel) SetSourceVnetRegion(v string)`

SetSourceVnetRegion sets SourceVnetRegion field to given value.

### HasSourceVnetRegion

`func (o *VnetPeeringModel) HasSourceVnetRegion() bool`

HasSourceVnetRegion returns a boolean if a field has been set.

### SetSourceVnetRegionNil

`func (o *VnetPeeringModel) SetSourceVnetRegionNil(b bool)`

 SetSourceVnetRegionNil sets the value for SourceVnetRegion to be an explicit nil

### UnsetSourceVnetRegion
`func (o *VnetPeeringModel) UnsetSourceVnetRegion()`

UnsetSourceVnetRegion ensures that no value is present for SourceVnetRegion, not even an explicit nil
### GetIpAddressSpace

`func (o *VnetPeeringModel) GetIpAddressSpace() []string`

GetIpAddressSpace returns the IpAddressSpace field if non-nil, zero value otherwise.

### GetIpAddressSpaceOk

`func (o *VnetPeeringModel) GetIpAddressSpaceOk() (*[]string, bool)`

GetIpAddressSpaceOk returns a tuple with the IpAddressSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressSpace

`func (o *VnetPeeringModel) SetIpAddressSpace(v []string)`

SetIpAddressSpace sets IpAddressSpace field to given value.

### HasIpAddressSpace

`func (o *VnetPeeringModel) HasIpAddressSpace() bool`

HasIpAddressSpace returns a boolean if a field has been set.

### SetIpAddressSpaceNil

`func (o *VnetPeeringModel) SetIpAddressSpaceNil(b bool)`

 SetIpAddressSpaceNil sets the value for IpAddressSpace to be an explicit nil

### UnsetIpAddressSpace
`func (o *VnetPeeringModel) UnsetIpAddressSpace()`

UnsetIpAddressSpace ensures that no value is present for IpAddressSpace, not even an explicit nil
### GetNumberOfAddressesInUse

`func (o *VnetPeeringModel) GetNumberOfAddressesInUse() int32`

GetNumberOfAddressesInUse returns the NumberOfAddressesInUse field if non-nil, zero value otherwise.

### GetNumberOfAddressesInUseOk

`func (o *VnetPeeringModel) GetNumberOfAddressesInUseOk() (*int32, bool)`

GetNumberOfAddressesInUseOk returns a tuple with the NumberOfAddressesInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAddressesInUse

`func (o *VnetPeeringModel) SetNumberOfAddressesInUse(v int32)`

SetNumberOfAddressesInUse sets NumberOfAddressesInUse field to given value.

### HasNumberOfAddressesInUse

`func (o *VnetPeeringModel) HasNumberOfAddressesInUse() bool`

HasNumberOfAddressesInUse returns a boolean if a field has been set.

### GetIsPeeringActive

`func (o *VnetPeeringModel) GetIsPeeringActive() bool`

GetIsPeeringActive returns the IsPeeringActive field if non-nil, zero value otherwise.

### GetIsPeeringActiveOk

`func (o *VnetPeeringModel) GetIsPeeringActiveOk() (*bool, bool)`

GetIsPeeringActiveOk returns a tuple with the IsPeeringActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPeeringActive

`func (o *VnetPeeringModel) SetIsPeeringActive(v bool)`

SetIsPeeringActive sets IsPeeringActive field to given value.

### HasIsPeeringActive

`func (o *VnetPeeringModel) HasIsPeeringActive() bool`

HasIsPeeringActive returns a boolean if a field has been set.

### GetCitrixManaged

`func (o *VnetPeeringModel) GetCitrixManaged() bool`

GetCitrixManaged returns the CitrixManaged field if non-nil, zero value otherwise.

### GetCitrixManagedOk

`func (o *VnetPeeringModel) GetCitrixManagedOk() (*bool, bool)`

GetCitrixManagedOk returns a tuple with the CitrixManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitrixManaged

`func (o *VnetPeeringModel) SetCitrixManaged(v bool)`

SetCitrixManaged sets CitrixManaged field to given value.

### HasCitrixManaged

`func (o *VnetPeeringModel) HasCitrixManaged() bool`

HasCitrixManaged returns a boolean if a field has been set.

### GetPeeredVnetId

`func (o *VnetPeeringModel) GetPeeredVnetId() string`

GetPeeredVnetId returns the PeeredVnetId field if non-nil, zero value otherwise.

### GetPeeredVnetIdOk

`func (o *VnetPeeringModel) GetPeeredVnetIdOk() (*string, bool)`

GetPeeredVnetIdOk returns a tuple with the PeeredVnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeeredVnetId

`func (o *VnetPeeringModel) SetPeeredVnetId(v string)`

SetPeeredVnetId sets PeeredVnetId field to given value.

### HasPeeredVnetId

`func (o *VnetPeeringModel) HasPeeredVnetId() bool`

HasPeeredVnetId returns a boolean if a field has been set.

### SetPeeredVnetIdNil

`func (o *VnetPeeringModel) SetPeeredVnetIdNil(b bool)`

 SetPeeredVnetIdNil sets the value for PeeredVnetId to be an explicit nil

### UnsetPeeredVnetId
`func (o *VnetPeeringModel) UnsetPeeredVnetId()`

UnsetPeeredVnetId ensures that no value is present for PeeredVnetId, not even an explicit nil
### GetSubscriptionId

`func (o *VnetPeeringModel) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *VnetPeeringModel) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *VnetPeeringModel) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *VnetPeeringModel) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *VnetPeeringModel) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *VnetPeeringModel) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetDnsServers

`func (o *VnetPeeringModel) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *VnetPeeringModel) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *VnetPeeringModel) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *VnetPeeringModel) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### SetDnsServersNil

`func (o *VnetPeeringModel) SetDnsServersNil(b bool)`

 SetDnsServersNil sets the value for DnsServers to be an explicit nil

### UnsetDnsServers
`func (o *VnetPeeringModel) UnsetDnsServers()`

UnsetDnsServers ensures that no value is present for DnsServers, not even an explicit nil
### GetCspCustomer

`func (o *VnetPeeringModel) GetCspCustomer() string`

GetCspCustomer returns the CspCustomer field if non-nil, zero value otherwise.

### GetCspCustomerOk

`func (o *VnetPeeringModel) GetCspCustomerOk() (*string, bool)`

GetCspCustomerOk returns a tuple with the CspCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomer

`func (o *VnetPeeringModel) SetCspCustomer(v string)`

SetCspCustomer sets CspCustomer field to given value.

### HasCspCustomer

`func (o *VnetPeeringModel) HasCspCustomer() bool`

HasCspCustomer returns a boolean if a field has been set.

### SetCspCustomerNil

`func (o *VnetPeeringModel) SetCspCustomerNil(b bool)`

 SetCspCustomerNil sets the value for CspCustomer to be an explicit nil

### UnsetCspCustomer
`func (o *VnetPeeringModel) UnsetCspCustomer()`

UnsetCspCustomer ensures that no value is present for CspCustomer, not even an explicit nil
### GetCatalogs

`func (o *VnetPeeringModel) GetCatalogs() []string`

GetCatalogs returns the Catalogs field if non-nil, zero value otherwise.

### GetCatalogsOk

`func (o *VnetPeeringModel) GetCatalogsOk() (*[]string, bool)`

GetCatalogsOk returns a tuple with the Catalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogs

`func (o *VnetPeeringModel) SetCatalogs(v []string)`

SetCatalogs sets Catalogs field to given value.

### HasCatalogs

`func (o *VnetPeeringModel) HasCatalogs() bool`

HasCatalogs returns a boolean if a field has been set.

### SetCatalogsNil

`func (o *VnetPeeringModel) SetCatalogsNil(b bool)`

 SetCatalogsNil sets the value for Catalogs to be an explicit nil

### UnsetCatalogs
`func (o *VnetPeeringModel) UnsetCatalogs()`

UnsetCatalogs ensures that no value is present for Catalogs, not even an explicit nil
### GetImages

`func (o *VnetPeeringModel) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *VnetPeeringModel) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *VnetPeeringModel) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *VnetPeeringModel) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *VnetPeeringModel) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *VnetPeeringModel) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetBastions

`func (o *VnetPeeringModel) GetBastions() []string`

GetBastions returns the Bastions field if non-nil, zero value otherwise.

### GetBastionsOk

`func (o *VnetPeeringModel) GetBastionsOk() (*[]string, bool)`

GetBastionsOk returns a tuple with the Bastions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastions

`func (o *VnetPeeringModel) SetBastions(v []string)`

SetBastions sets Bastions field to given value.

### HasBastions

`func (o *VnetPeeringModel) HasBastions() bool`

HasBastions returns a boolean if a field has been set.

### SetBastionsNil

`func (o *VnetPeeringModel) SetBastionsNil(b bool)`

 SetBastionsNil sets the value for Bastions to be an explicit nil

### UnsetBastions
`func (o *VnetPeeringModel) UnsetBastions()`

UnsetBastions ensures that no value is present for Bastions, not even an explicit nil
### GetUseGateway

`func (o *VnetPeeringModel) GetUseGateway() bool`

GetUseGateway returns the UseGateway field if non-nil, zero value otherwise.

### GetUseGatewayOk

`func (o *VnetPeeringModel) GetUseGatewayOk() (*bool, bool)`

GetUseGatewayOk returns a tuple with the UseGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGateway

`func (o *VnetPeeringModel) SetUseGateway(v bool)`

SetUseGateway sets UseGateway field to given value.

### HasUseGateway

`func (o *VnetPeeringModel) HasUseGateway() bool`

HasUseGateway returns a boolean if a field has been set.

### SetUseGatewayNil

`func (o *VnetPeeringModel) SetUseGatewayNil(b bool)`

 SetUseGatewayNil sets the value for UseGateway to be an explicit nil

### UnsetUseGateway
`func (o *VnetPeeringModel) UnsetUseGateway()`

UnsetUseGateway ensures that no value is present for UseGateway, not even an explicit nil
### GetDisableRoutePropagation

`func (o *VnetPeeringModel) GetDisableRoutePropagation() bool`

GetDisableRoutePropagation returns the DisableRoutePropagation field if non-nil, zero value otherwise.

### GetDisableRoutePropagationOk

`func (o *VnetPeeringModel) GetDisableRoutePropagationOk() (*bool, bool)`

GetDisableRoutePropagationOk returns a tuple with the DisableRoutePropagation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRoutePropagation

`func (o *VnetPeeringModel) SetDisableRoutePropagation(v bool)`

SetDisableRoutePropagation sets DisableRoutePropagation field to given value.

### HasDisableRoutePropagation

`func (o *VnetPeeringModel) HasDisableRoutePropagation() bool`

HasDisableRoutePropagation returns a boolean if a field has been set.

### SetDisableRoutePropagationNil

`func (o *VnetPeeringModel) SetDisableRoutePropagationNil(b bool)`

 SetDisableRoutePropagationNil sets the value for DisableRoutePropagation to be an explicit nil

### UnsetDisableRoutePropagation
`func (o *VnetPeeringModel) UnsetDisableRoutePropagation()`

UnsetDisableRoutePropagation ensures that no value is present for DisableRoutePropagation, not even an explicit nil
### GetNatGatewayConfig

`func (o *VnetPeeringModel) GetNatGatewayConfig() NatGatewayModelOverview`

GetNatGatewayConfig returns the NatGatewayConfig field if non-nil, zero value otherwise.

### GetNatGatewayConfigOk

`func (o *VnetPeeringModel) GetNatGatewayConfigOk() (*NatGatewayModelOverview, bool)`

GetNatGatewayConfigOk returns a tuple with the NatGatewayConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatGatewayConfig

`func (o *VnetPeeringModel) SetNatGatewayConfig(v NatGatewayModelOverview)`

SetNatGatewayConfig sets NatGatewayConfig field to given value.

### HasNatGatewayConfig

`func (o *VnetPeeringModel) HasNatGatewayConfig() bool`

HasNatGatewayConfig returns a boolean if a field has been set.

### SetNatGatewayConfigNil

`func (o *VnetPeeringModel) SetNatGatewayConfigNil(b bool)`

 SetNatGatewayConfigNil sets the value for NatGatewayConfig to be an explicit nil

### UnsetNatGatewayConfig
`func (o *VnetPeeringModel) UnsetNatGatewayConfig()`

UnsetNatGatewayConfig ensures that no value is present for NatGatewayConfig, not even an explicit nil
### GetRemoteTenantId

`func (o *VnetPeeringModel) GetRemoteTenantId() string`

GetRemoteTenantId returns the RemoteTenantId field if non-nil, zero value otherwise.

### GetRemoteTenantIdOk

`func (o *VnetPeeringModel) GetRemoteTenantIdOk() (*string, bool)`

GetRemoteTenantIdOk returns a tuple with the RemoteTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTenantId

`func (o *VnetPeeringModel) SetRemoteTenantId(v string)`

SetRemoteTenantId sets RemoteTenantId field to given value.

### HasRemoteTenantId

`func (o *VnetPeeringModel) HasRemoteTenantId() bool`

HasRemoteTenantId returns a boolean if a field has been set.

### SetRemoteTenantIdNil

`func (o *VnetPeeringModel) SetRemoteTenantIdNil(b bool)`

 SetRemoteTenantIdNil sets the value for RemoteTenantId to be an explicit nil

### UnsetRemoteTenantId
`func (o *VnetPeeringModel) UnsetRemoteTenantId()`

UnsetRemoteTenantId ensures that no value is present for RemoteTenantId, not even an explicit nil
### GetResourceId

`func (o *VnetPeeringModel) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *VnetPeeringModel) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *VnetPeeringModel) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *VnetPeeringModel) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *VnetPeeringModel) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *VnetPeeringModel) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetVirtualHubId

`func (o *VnetPeeringModel) GetVirtualHubId() string`

GetVirtualHubId returns the VirtualHubId field if non-nil, zero value otherwise.

### GetVirtualHubIdOk

`func (o *VnetPeeringModel) GetVirtualHubIdOk() (*string, bool)`

GetVirtualHubIdOk returns a tuple with the VirtualHubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualHubId

`func (o *VnetPeeringModel) SetVirtualHubId(v string)`

SetVirtualHubId sets VirtualHubId field to given value.

### HasVirtualHubId

`func (o *VnetPeeringModel) HasVirtualHubId() bool`

HasVirtualHubId returns a boolean if a field has been set.

### SetVirtualHubIdNil

`func (o *VnetPeeringModel) SetVirtualHubIdNil(b bool)`

 SetVirtualHubIdNil sets the value for VirtualHubId to be an explicit nil

### UnsetVirtualHubId
`func (o *VnetPeeringModel) UnsetVirtualHubId()`

UnsetVirtualHubId ensures that no value is present for VirtualHubId, not even an explicit nil
### GetVirtualHubResourceGroup

`func (o *VnetPeeringModel) GetVirtualHubResourceGroup() string`

GetVirtualHubResourceGroup returns the VirtualHubResourceGroup field if non-nil, zero value otherwise.

### GetVirtualHubResourceGroupOk

`func (o *VnetPeeringModel) GetVirtualHubResourceGroupOk() (*string, bool)`

GetVirtualHubResourceGroupOk returns a tuple with the VirtualHubResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualHubResourceGroup

`func (o *VnetPeeringModel) SetVirtualHubResourceGroup(v string)`

SetVirtualHubResourceGroup sets VirtualHubResourceGroup field to given value.

### HasVirtualHubResourceGroup

`func (o *VnetPeeringModel) HasVirtualHubResourceGroup() bool`

HasVirtualHubResourceGroup returns a boolean if a field has been set.

### SetVirtualHubResourceGroupNil

`func (o *VnetPeeringModel) SetVirtualHubResourceGroupNil(b bool)`

 SetVirtualHubResourceGroupNil sets the value for VirtualHubResourceGroup to be an explicit nil

### UnsetVirtualHubResourceGroup
`func (o *VnetPeeringModel) UnsetVirtualHubResourceGroup()`

UnsetVirtualHubResourceGroup ensures that no value is present for VirtualHubResourceGroup, not even an explicit nil
### GetVirtualHubName

`func (o *VnetPeeringModel) GetVirtualHubName() string`

GetVirtualHubName returns the VirtualHubName field if non-nil, zero value otherwise.

### GetVirtualHubNameOk

`func (o *VnetPeeringModel) GetVirtualHubNameOk() (*string, bool)`

GetVirtualHubNameOk returns a tuple with the VirtualHubName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualHubName

`func (o *VnetPeeringModel) SetVirtualHubName(v string)`

SetVirtualHubName sets VirtualHubName field to given value.

### HasVirtualHubName

`func (o *VnetPeeringModel) HasVirtualHubName() bool`

HasVirtualHubName returns a boolean if a field has been set.

### SetVirtualHubNameNil

`func (o *VnetPeeringModel) SetVirtualHubNameNil(b bool)`

 SetVirtualHubNameNil sets the value for VirtualHubName to be an explicit nil

### UnsetVirtualHubName
`func (o *VnetPeeringModel) UnsetVirtualHubName()`

UnsetVirtualHubName ensures that no value is present for VirtualHubName, not even an explicit nil
### GetEnableInternetSecurity

`func (o *VnetPeeringModel) GetEnableInternetSecurity() bool`

GetEnableInternetSecurity returns the EnableInternetSecurity field if non-nil, zero value otherwise.

### GetEnableInternetSecurityOk

`func (o *VnetPeeringModel) GetEnableInternetSecurityOk() (*bool, bool)`

GetEnableInternetSecurityOk returns a tuple with the EnableInternetSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInternetSecurity

`func (o *VnetPeeringModel) SetEnableInternetSecurity(v bool)`

SetEnableInternetSecurity sets EnableInternetSecurity field to given value.

### HasEnableInternetSecurity

`func (o *VnetPeeringModel) HasEnableInternetSecurity() bool`

HasEnableInternetSecurity returns a boolean if a field has been set.

### SetEnableInternetSecurityNil

`func (o *VnetPeeringModel) SetEnableInternetSecurityNil(b bool)`

 SetEnableInternetSecurityNil sets the value for EnableInternetSecurity to be an explicit nil

### UnsetEnableInternetSecurity
`func (o *VnetPeeringModel) UnsetEnableInternetSecurity()`

UnsetEnableInternetSecurity ensures that no value is present for EnableInternetSecurity, not even an explicit nil
### GetRoutingConfiguration

`func (o *VnetPeeringModel) GetRoutingConfiguration() string`

GetRoutingConfiguration returns the RoutingConfiguration field if non-nil, zero value otherwise.

### GetRoutingConfigurationOk

`func (o *VnetPeeringModel) GetRoutingConfigurationOk() (*string, bool)`

GetRoutingConfigurationOk returns a tuple with the RoutingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingConfiguration

`func (o *VnetPeeringModel) SetRoutingConfiguration(v string)`

SetRoutingConfiguration sets RoutingConfiguration field to given value.

### HasRoutingConfiguration

`func (o *VnetPeeringModel) HasRoutingConfiguration() bool`

HasRoutingConfiguration returns a boolean if a field has been set.

### SetRoutingConfigurationNil

`func (o *VnetPeeringModel) SetRoutingConfigurationNil(b bool)`

 SetRoutingConfigurationNil sets the value for RoutingConfiguration to be an explicit nil

### UnsetRoutingConfiguration
`func (o *VnetPeeringModel) UnsetRoutingConfiguration()`

UnsetRoutingConfiguration ensures that no value is present for RoutingConfiguration, not even an explicit nil
### GetId

`func (o *VnetPeeringModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VnetPeeringModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VnetPeeringModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VnetPeeringModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *VnetPeeringModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VnetPeeringModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *VnetPeeringModel) GetType() OnPremConnectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VnetPeeringModel) GetTypeOk() (*OnPremConnectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VnetPeeringModel) SetType(v OnPremConnectionType)`

SetType sets Type field to given value.

### HasType

`func (o *VnetPeeringModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *VnetPeeringModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VnetPeeringModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VnetPeeringModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VnetPeeringModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VnetPeeringModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VnetPeeringModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStartedAt

`func (o *VnetPeeringModel) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *VnetPeeringModel) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *VnetPeeringModel) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *VnetPeeringModel) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *VnetPeeringModel) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *VnetPeeringModel) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetEstimatedTimeInMinute

`func (o *VnetPeeringModel) GetEstimatedTimeInMinute() int32`

GetEstimatedTimeInMinute returns the EstimatedTimeInMinute field if non-nil, zero value otherwise.

### GetEstimatedTimeInMinuteOk

`func (o *VnetPeeringModel) GetEstimatedTimeInMinuteOk() (*int32, bool)`

GetEstimatedTimeInMinuteOk returns a tuple with the EstimatedTimeInMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTimeInMinute

`func (o *VnetPeeringModel) SetEstimatedTimeInMinute(v int32)`

SetEstimatedTimeInMinute sets EstimatedTimeInMinute field to given value.

### HasEstimatedTimeInMinute

`func (o *VnetPeeringModel) HasEstimatedTimeInMinute() bool`

HasEstimatedTimeInMinute returns a boolean if a field has been set.

### SetEstimatedTimeInMinuteNil

`func (o *VnetPeeringModel) SetEstimatedTimeInMinuteNil(b bool)`

 SetEstimatedTimeInMinuteNil sets the value for EstimatedTimeInMinute to be an explicit nil

### UnsetEstimatedTimeInMinute
`func (o *VnetPeeringModel) UnsetEstimatedTimeInMinute()`

UnsetEstimatedTimeInMinute ensures that no value is present for EstimatedTimeInMinute, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


