# VnetPeeringModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**VnetPeeringState**](VnetPeeringState.md) | State of the peering | [optional] 
**Error** | Pointer to **string** | Error that occured while peering the vnet | [optional] 
**TransactionId** | Pointer to **string** | ID of the transaction that performed pairing job | [optional] 
**Region** | Pointer to **string** | Azure region where the vnet in the managed subscription is located | [optional] 
**SourceVnetRegion** | Pointer to **string** | Azure region where the originating vnet is located | [optional] 
**IpAddressSpace** | Pointer to **[]string** | The IP Ranges allocated to the peering | [optional] [readonly] 
**NumberOfAddressesInUse** | Pointer to **int32** | Number of IP Addresses that are currently in use in the subnet | [optional] [readonly] 
**IsPeeringActive** | Pointer to **bool** | Indicates if the Peering is active | [optional] 
**CitrixManaged** | Pointer to **bool** | Indicates if the Peering is managed by Citrix | [optional] 
**PeeredVnetId** | Pointer to **string** | Resource ID of the Peered VNET | [optional] 
**SubscriptionId** | Pointer to **string** | ID of the Azure Subscription that the citrix managed vnet is associated with | [optional] [readonly] 
**DnsServers** | Pointer to **[]string** | IP address of the DNS server of this VNet | [optional] 
**CspCustomer** | Pointer to **string** | Indicates that partner-tenant relationship exists if not null | [optional] 
**Catalogs** | Pointer to **[]string** | List of associated catalogs | [optional] 
**Images** | Pointer to **[]string** | List of associated images | [optional] 
**Bastions** | Pointer to **[]string** | List of associated bastions | [optional] 
**UseGateway** | Pointer to **bool** | Indicates if the peered vnet will be using the customers gateway | [optional] 
**DisableRoutePropagation** | Pointer to **bool** | Indicates if route propagation in the route table should be disabled (option is valid only if customer is using gateway). | [optional] 
**NatGatewayConfig** | Pointer to [**NatGatewayModelOverview**](NatGatewayModelOverview.md) | The NAT gateway config | [optional] 
**Id** | Pointer to **string** | ID of the connection | [optional] 
**Type** | Pointer to [**OnPremConnectionType**](OnPremConnectionType.md) | The type of connection | [optional] 
**Name** | Pointer to **string** | Name of the connection | [optional] 
**StartedAt** | Pointer to **time.Time** | The datetime when the job started | [optional] 
**EstimatedTimeInMinute** | Pointer to **int32** | Estimated total time for the job to finish | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


