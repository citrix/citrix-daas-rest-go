# AzureVpnConnectionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**AzureVpnConnectionState**](AzureVpnConnectionState.md) | State of the vpn connection | [optional] 
**CspCustomer** | Pointer to **string** | Tenant customer the connection is associated with | [optional] 
**Region** | Pointer to **string** | Name of the azure region where the connection will be created | [optional] 
**VdaSubnet** | Pointer to [**ConnectionSubnet**](ConnectionSubnet.md) | Details of the vda subnet | [optional] 
**GatewaySubnet** | Pointer to [**ConnectionSubnet**](ConnectionSubnet.md) | Details of the subnet the gateway uses to communicate | [optional] 
**LocalGatewayIP** | Pointer to **string** | IP address of the on-prem gateway | [optional] 
**DnsServers** | Pointer to **[]string** | List of DNS Servers | [optional] 
**LocalAddresses** | Pointer to [**[]ConnectionSubnet**](ConnectionSubnet.md) | List of addresses that will be accessible behind the gateway | [optional] 
**ConnectionStatus** | Pointer to **string** | Status of the VPN connection | [optional] 
**GatewayIps** | Pointer to **[]string** | IP address of the gateway for this connection | [optional] 
**NumberAddressesInUse** | Pointer to **int32** | Number of addresses in use on the VDA subnet | [optional] 
**SubscriptionId** | Pointer to **string** | ID of the Managed Subscription the Connection is associated with. | [optional] 
**SharedKey** | Pointer to **string** | The pre shared key that was used for the connection (obfuscated) | [optional] 
**GatewaySku** | Pointer to **string** | The SKU type of the gateway | [optional] 
**ActiveActive** | Pointer to **bool** | Specifies if connection is active-standby or active-active | [optional] 
**TransactionId** | Pointer to **string** | The transaction ID associated with the vpn connection | [optional] 
**Error** | Pointer to **string** | Error that occurred while creating vpn connection | [optional] 
**Catalogs** | Pointer to **[]string** | List of associated catalogs | [optional] 
**Images** | Pointer to **[]string** | List of associated images | [optional] 
**Bastions** | Pointer to **[]string** | List of associated bastions | [optional] 
**Id** | Pointer to **string** | ID of the connection | [optional] 
**Type** | Pointer to [**OnPremConnectionType**](OnPremConnectionType.md) | The type of connection | [optional] 
**Name** | Pointer to **string** | Name of the connection | [optional] 
**StartedAt** | Pointer to **time.Time** | The datetime when the job started | [optional] 
**EstimatedTimeInMinute** | Pointer to **int32** | Estimated total time for the job to finish | [optional] 

## Methods

### NewAzureVpnConnectionModel

`func NewAzureVpnConnectionModel() *AzureVpnConnectionModel`

NewAzureVpnConnectionModel instantiates a new AzureVpnConnectionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureVpnConnectionModelWithDefaults

`func NewAzureVpnConnectionModelWithDefaults() *AzureVpnConnectionModel`

NewAzureVpnConnectionModelWithDefaults instantiates a new AzureVpnConnectionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *AzureVpnConnectionModel) GetState() AzureVpnConnectionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AzureVpnConnectionModel) GetStateOk() (*AzureVpnConnectionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AzureVpnConnectionModel) SetState(v AzureVpnConnectionState)`

SetState sets State field to given value.

### HasState

`func (o *AzureVpnConnectionModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCspCustomer

`func (o *AzureVpnConnectionModel) GetCspCustomer() string`

GetCspCustomer returns the CspCustomer field if non-nil, zero value otherwise.

### GetCspCustomerOk

`func (o *AzureVpnConnectionModel) GetCspCustomerOk() (*string, bool)`

GetCspCustomerOk returns a tuple with the CspCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomer

`func (o *AzureVpnConnectionModel) SetCspCustomer(v string)`

SetCspCustomer sets CspCustomer field to given value.

### HasCspCustomer

`func (o *AzureVpnConnectionModel) HasCspCustomer() bool`

HasCspCustomer returns a boolean if a field has been set.

### GetRegion

`func (o *AzureVpnConnectionModel) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AzureVpnConnectionModel) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AzureVpnConnectionModel) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AzureVpnConnectionModel) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetVdaSubnet

`func (o *AzureVpnConnectionModel) GetVdaSubnet() ConnectionSubnet`

GetVdaSubnet returns the VdaSubnet field if non-nil, zero value otherwise.

### GetVdaSubnetOk

`func (o *AzureVpnConnectionModel) GetVdaSubnetOk() (*ConnectionSubnet, bool)`

GetVdaSubnetOk returns a tuple with the VdaSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaSubnet

`func (o *AzureVpnConnectionModel) SetVdaSubnet(v ConnectionSubnet)`

SetVdaSubnet sets VdaSubnet field to given value.

### HasVdaSubnet

`func (o *AzureVpnConnectionModel) HasVdaSubnet() bool`

HasVdaSubnet returns a boolean if a field has been set.

### GetGatewaySubnet

`func (o *AzureVpnConnectionModel) GetGatewaySubnet() ConnectionSubnet`

GetGatewaySubnet returns the GatewaySubnet field if non-nil, zero value otherwise.

### GetGatewaySubnetOk

`func (o *AzureVpnConnectionModel) GetGatewaySubnetOk() (*ConnectionSubnet, bool)`

GetGatewaySubnetOk returns a tuple with the GatewaySubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySubnet

`func (o *AzureVpnConnectionModel) SetGatewaySubnet(v ConnectionSubnet)`

SetGatewaySubnet sets GatewaySubnet field to given value.

### HasGatewaySubnet

`func (o *AzureVpnConnectionModel) HasGatewaySubnet() bool`

HasGatewaySubnet returns a boolean if a field has been set.

### GetLocalGatewayIP

`func (o *AzureVpnConnectionModel) GetLocalGatewayIP() string`

GetLocalGatewayIP returns the LocalGatewayIP field if non-nil, zero value otherwise.

### GetLocalGatewayIPOk

`func (o *AzureVpnConnectionModel) GetLocalGatewayIPOk() (*string, bool)`

GetLocalGatewayIPOk returns a tuple with the LocalGatewayIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalGatewayIP

`func (o *AzureVpnConnectionModel) SetLocalGatewayIP(v string)`

SetLocalGatewayIP sets LocalGatewayIP field to given value.

### HasLocalGatewayIP

`func (o *AzureVpnConnectionModel) HasLocalGatewayIP() bool`

HasLocalGatewayIP returns a boolean if a field has been set.

### GetDnsServers

`func (o *AzureVpnConnectionModel) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *AzureVpnConnectionModel) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *AzureVpnConnectionModel) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *AzureVpnConnectionModel) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetLocalAddresses

`func (o *AzureVpnConnectionModel) GetLocalAddresses() []ConnectionSubnet`

GetLocalAddresses returns the LocalAddresses field if non-nil, zero value otherwise.

### GetLocalAddressesOk

`func (o *AzureVpnConnectionModel) GetLocalAddressesOk() (*[]ConnectionSubnet, bool)`

GetLocalAddressesOk returns a tuple with the LocalAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddresses

`func (o *AzureVpnConnectionModel) SetLocalAddresses(v []ConnectionSubnet)`

SetLocalAddresses sets LocalAddresses field to given value.

### HasLocalAddresses

`func (o *AzureVpnConnectionModel) HasLocalAddresses() bool`

HasLocalAddresses returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *AzureVpnConnectionModel) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *AzureVpnConnectionModel) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *AzureVpnConnectionModel) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *AzureVpnConnectionModel) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetGatewayIps

`func (o *AzureVpnConnectionModel) GetGatewayIps() []string`

GetGatewayIps returns the GatewayIps field if non-nil, zero value otherwise.

### GetGatewayIpsOk

`func (o *AzureVpnConnectionModel) GetGatewayIpsOk() (*[]string, bool)`

GetGatewayIpsOk returns a tuple with the GatewayIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIps

`func (o *AzureVpnConnectionModel) SetGatewayIps(v []string)`

SetGatewayIps sets GatewayIps field to given value.

### HasGatewayIps

`func (o *AzureVpnConnectionModel) HasGatewayIps() bool`

HasGatewayIps returns a boolean if a field has been set.

### GetNumberAddressesInUse

`func (o *AzureVpnConnectionModel) GetNumberAddressesInUse() int32`

GetNumberAddressesInUse returns the NumberAddressesInUse field if non-nil, zero value otherwise.

### GetNumberAddressesInUseOk

`func (o *AzureVpnConnectionModel) GetNumberAddressesInUseOk() (*int32, bool)`

GetNumberAddressesInUseOk returns a tuple with the NumberAddressesInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberAddressesInUse

`func (o *AzureVpnConnectionModel) SetNumberAddressesInUse(v int32)`

SetNumberAddressesInUse sets NumberAddressesInUse field to given value.

### HasNumberAddressesInUse

`func (o *AzureVpnConnectionModel) HasNumberAddressesInUse() bool`

HasNumberAddressesInUse returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *AzureVpnConnectionModel) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureVpnConnectionModel) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureVpnConnectionModel) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AzureVpnConnectionModel) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSharedKey

`func (o *AzureVpnConnectionModel) GetSharedKey() string`

GetSharedKey returns the SharedKey field if non-nil, zero value otherwise.

### GetSharedKeyOk

`func (o *AzureVpnConnectionModel) GetSharedKeyOk() (*string, bool)`

GetSharedKeyOk returns a tuple with the SharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedKey

`func (o *AzureVpnConnectionModel) SetSharedKey(v string)`

SetSharedKey sets SharedKey field to given value.

### HasSharedKey

`func (o *AzureVpnConnectionModel) HasSharedKey() bool`

HasSharedKey returns a boolean if a field has been set.

### GetGatewaySku

`func (o *AzureVpnConnectionModel) GetGatewaySku() string`

GetGatewaySku returns the GatewaySku field if non-nil, zero value otherwise.

### GetGatewaySkuOk

`func (o *AzureVpnConnectionModel) GetGatewaySkuOk() (*string, bool)`

GetGatewaySkuOk returns a tuple with the GatewaySku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySku

`func (o *AzureVpnConnectionModel) SetGatewaySku(v string)`

SetGatewaySku sets GatewaySku field to given value.

### HasGatewaySku

`func (o *AzureVpnConnectionModel) HasGatewaySku() bool`

HasGatewaySku returns a boolean if a field has been set.

### GetActiveActive

`func (o *AzureVpnConnectionModel) GetActiveActive() bool`

GetActiveActive returns the ActiveActive field if non-nil, zero value otherwise.

### GetActiveActiveOk

`func (o *AzureVpnConnectionModel) GetActiveActiveOk() (*bool, bool)`

GetActiveActiveOk returns a tuple with the ActiveActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveActive

`func (o *AzureVpnConnectionModel) SetActiveActive(v bool)`

SetActiveActive sets ActiveActive field to given value.

### HasActiveActive

`func (o *AzureVpnConnectionModel) HasActiveActive() bool`

HasActiveActive returns a boolean if a field has been set.

### GetTransactionId

`func (o *AzureVpnConnectionModel) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *AzureVpnConnectionModel) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *AzureVpnConnectionModel) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *AzureVpnConnectionModel) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetError

`func (o *AzureVpnConnectionModel) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AzureVpnConnectionModel) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AzureVpnConnectionModel) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AzureVpnConnectionModel) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCatalogs

`func (o *AzureVpnConnectionModel) GetCatalogs() []string`

GetCatalogs returns the Catalogs field if non-nil, zero value otherwise.

### GetCatalogsOk

`func (o *AzureVpnConnectionModel) GetCatalogsOk() (*[]string, bool)`

GetCatalogsOk returns a tuple with the Catalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogs

`func (o *AzureVpnConnectionModel) SetCatalogs(v []string)`

SetCatalogs sets Catalogs field to given value.

### HasCatalogs

`func (o *AzureVpnConnectionModel) HasCatalogs() bool`

HasCatalogs returns a boolean if a field has been set.

### GetImages

`func (o *AzureVpnConnectionModel) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *AzureVpnConnectionModel) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *AzureVpnConnectionModel) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *AzureVpnConnectionModel) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetBastions

`func (o *AzureVpnConnectionModel) GetBastions() []string`

GetBastions returns the Bastions field if non-nil, zero value otherwise.

### GetBastionsOk

`func (o *AzureVpnConnectionModel) GetBastionsOk() (*[]string, bool)`

GetBastionsOk returns a tuple with the Bastions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastions

`func (o *AzureVpnConnectionModel) SetBastions(v []string)`

SetBastions sets Bastions field to given value.

### HasBastions

`func (o *AzureVpnConnectionModel) HasBastions() bool`

HasBastions returns a boolean if a field has been set.

### GetId

`func (o *AzureVpnConnectionModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureVpnConnectionModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureVpnConnectionModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureVpnConnectionModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AzureVpnConnectionModel) GetType() OnPremConnectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AzureVpnConnectionModel) GetTypeOk() (*OnPremConnectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AzureVpnConnectionModel) SetType(v OnPremConnectionType)`

SetType sets Type field to given value.

### HasType

`func (o *AzureVpnConnectionModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *AzureVpnConnectionModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureVpnConnectionModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureVpnConnectionModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureVpnConnectionModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartedAt

`func (o *AzureVpnConnectionModel) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *AzureVpnConnectionModel) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *AzureVpnConnectionModel) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *AzureVpnConnectionModel) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetEstimatedTimeInMinute

`func (o *AzureVpnConnectionModel) GetEstimatedTimeInMinute() int32`

GetEstimatedTimeInMinute returns the EstimatedTimeInMinute field if non-nil, zero value otherwise.

### GetEstimatedTimeInMinuteOk

`func (o *AzureVpnConnectionModel) GetEstimatedTimeInMinuteOk() (*int32, bool)`

GetEstimatedTimeInMinuteOk returns a tuple with the EstimatedTimeInMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTimeInMinute

`func (o *AzureVpnConnectionModel) SetEstimatedTimeInMinute(v int32)`

SetEstimatedTimeInMinute sets EstimatedTimeInMinute field to given value.

### HasEstimatedTimeInMinute

`func (o *AzureVpnConnectionModel) HasEstimatedTimeInMinute() bool`

HasEstimatedTimeInMinute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


