# UpdateAzureVpnConnectionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewaySku** | Pointer to **string** | Sku type for the virtual network gateway | [optional] 
**GatewayIP** | Pointer to **string** | IP address of the gateway to connect to | [optional] 
**LocalAddresses** | Pointer to [**[]ConnectionSubnet**](ConnectionSubnet.md) | List of addresses that will be accessible behind the gateway | [optional] 
**SharedKey** | Pointer to **string** | Pre-shared key that will be used to configure the IPSec tunnel | [optional] 
**DnsServers** | Pointer to **[]string** | List of DNS Servers | [optional] 

## Methods

### NewUpdateAzureVpnConnectionModel

`func NewUpdateAzureVpnConnectionModel() *UpdateAzureVpnConnectionModel`

NewUpdateAzureVpnConnectionModel instantiates a new UpdateAzureVpnConnectionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAzureVpnConnectionModelWithDefaults

`func NewUpdateAzureVpnConnectionModelWithDefaults() *UpdateAzureVpnConnectionModel`

NewUpdateAzureVpnConnectionModelWithDefaults instantiates a new UpdateAzureVpnConnectionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewaySku

`func (o *UpdateAzureVpnConnectionModel) GetGatewaySku() string`

GetGatewaySku returns the GatewaySku field if non-nil, zero value otherwise.

### GetGatewaySkuOk

`func (o *UpdateAzureVpnConnectionModel) GetGatewaySkuOk() (*string, bool)`

GetGatewaySkuOk returns a tuple with the GatewaySku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySku

`func (o *UpdateAzureVpnConnectionModel) SetGatewaySku(v string)`

SetGatewaySku sets GatewaySku field to given value.

### HasGatewaySku

`func (o *UpdateAzureVpnConnectionModel) HasGatewaySku() bool`

HasGatewaySku returns a boolean if a field has been set.

### GetGatewayIP

`func (o *UpdateAzureVpnConnectionModel) GetGatewayIP() string`

GetGatewayIP returns the GatewayIP field if non-nil, zero value otherwise.

### GetGatewayIPOk

`func (o *UpdateAzureVpnConnectionModel) GetGatewayIPOk() (*string, bool)`

GetGatewayIPOk returns a tuple with the GatewayIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIP

`func (o *UpdateAzureVpnConnectionModel) SetGatewayIP(v string)`

SetGatewayIP sets GatewayIP field to given value.

### HasGatewayIP

`func (o *UpdateAzureVpnConnectionModel) HasGatewayIP() bool`

HasGatewayIP returns a boolean if a field has been set.

### GetLocalAddresses

`func (o *UpdateAzureVpnConnectionModel) GetLocalAddresses() []ConnectionSubnet`

GetLocalAddresses returns the LocalAddresses field if non-nil, zero value otherwise.

### GetLocalAddressesOk

`func (o *UpdateAzureVpnConnectionModel) GetLocalAddressesOk() (*[]ConnectionSubnet, bool)`

GetLocalAddressesOk returns a tuple with the LocalAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddresses

`func (o *UpdateAzureVpnConnectionModel) SetLocalAddresses(v []ConnectionSubnet)`

SetLocalAddresses sets LocalAddresses field to given value.

### HasLocalAddresses

`func (o *UpdateAzureVpnConnectionModel) HasLocalAddresses() bool`

HasLocalAddresses returns a boolean if a field has been set.

### GetSharedKey

`func (o *UpdateAzureVpnConnectionModel) GetSharedKey() string`

GetSharedKey returns the SharedKey field if non-nil, zero value otherwise.

### GetSharedKeyOk

`func (o *UpdateAzureVpnConnectionModel) GetSharedKeyOk() (*string, bool)`

GetSharedKeyOk returns a tuple with the SharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedKey

`func (o *UpdateAzureVpnConnectionModel) SetSharedKey(v string)`

SetSharedKey sets SharedKey field to given value.

### HasSharedKey

`func (o *UpdateAzureVpnConnectionModel) HasSharedKey() bool`

HasSharedKey returns a boolean if a field has been set.

### GetDnsServers

`func (o *UpdateAzureVpnConnectionModel) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *UpdateAzureVpnConnectionModel) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *UpdateAzureVpnConnectionModel) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *UpdateAzureVpnConnectionModel) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


