# AddAzureVpnConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CspCustomer** | Pointer to **NullableString** | Tenant customer the connection is associated with | [optional] 
**GatewaySku** | Pointer to **NullableString** | Sku type to provision | [optional] 
**GatewayGeneration** | Pointer to **NullableString** | Generation of VPN gateway | [optional] 
**Name** | **string** | Name to assign the connection | 
**Region** | **string** | Name of the azure region where the connection will be created | 
**VdaSubnet** | [**ConnectionSubnet**](ConnectionSubnet.md) | Details of the vda subnet | 
**GatewaySubnet** | [**ConnectionSubnet**](ConnectionSubnet.md) | Details of the subnet the gateway uses to communicate | 
**GatewayIP** | **string** | IP address of the gateway to connect to | 
**DnsServers** | Pointer to **[]string** | List of DNS Servers | [optional] 
**LocalAddresses** | Pointer to [**[]ConnectionSubnet**](ConnectionSubnet.md) | List of addresses that will be accessible behind the gateway | [optional] 
**SharedKey** | Pointer to **NullableString** | Pre-shared key that will be used to configure the IPSec tunnel | [optional] 
**ManagedSubscriptionId** | Pointer to **NullableString** | ID of the Managed Azure Subscription to create the connection in. | [optional] 
**Routes** | Pointer to [**[]AzureRoute**](AzureRoute.md) | Routes to be added to the vnet | [optional] 

## Methods

### NewAddAzureVpnConnection

`func NewAddAzureVpnConnection(name string, region string, vdaSubnet ConnectionSubnet, gatewaySubnet ConnectionSubnet, gatewayIP string, ) *AddAzureVpnConnection`

NewAddAzureVpnConnection instantiates a new AddAzureVpnConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAzureVpnConnectionWithDefaults

`func NewAddAzureVpnConnectionWithDefaults() *AddAzureVpnConnection`

NewAddAzureVpnConnectionWithDefaults instantiates a new AddAzureVpnConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCspCustomer

`func (o *AddAzureVpnConnection) GetCspCustomer() string`

GetCspCustomer returns the CspCustomer field if non-nil, zero value otherwise.

### GetCspCustomerOk

`func (o *AddAzureVpnConnection) GetCspCustomerOk() (*string, bool)`

GetCspCustomerOk returns a tuple with the CspCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspCustomer

`func (o *AddAzureVpnConnection) SetCspCustomer(v string)`

SetCspCustomer sets CspCustomer field to given value.

### HasCspCustomer

`func (o *AddAzureVpnConnection) HasCspCustomer() bool`

HasCspCustomer returns a boolean if a field has been set.

### SetCspCustomerNil

`func (o *AddAzureVpnConnection) SetCspCustomerNil(b bool)`

 SetCspCustomerNil sets the value for CspCustomer to be an explicit nil

### UnsetCspCustomer
`func (o *AddAzureVpnConnection) UnsetCspCustomer()`

UnsetCspCustomer ensures that no value is present for CspCustomer, not even an explicit nil
### GetGatewaySku

`func (o *AddAzureVpnConnection) GetGatewaySku() string`

GetGatewaySku returns the GatewaySku field if non-nil, zero value otherwise.

### GetGatewaySkuOk

`func (o *AddAzureVpnConnection) GetGatewaySkuOk() (*string, bool)`

GetGatewaySkuOk returns a tuple with the GatewaySku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySku

`func (o *AddAzureVpnConnection) SetGatewaySku(v string)`

SetGatewaySku sets GatewaySku field to given value.

### HasGatewaySku

`func (o *AddAzureVpnConnection) HasGatewaySku() bool`

HasGatewaySku returns a boolean if a field has been set.

### SetGatewaySkuNil

`func (o *AddAzureVpnConnection) SetGatewaySkuNil(b bool)`

 SetGatewaySkuNil sets the value for GatewaySku to be an explicit nil

### UnsetGatewaySku
`func (o *AddAzureVpnConnection) UnsetGatewaySku()`

UnsetGatewaySku ensures that no value is present for GatewaySku, not even an explicit nil
### GetGatewayGeneration

`func (o *AddAzureVpnConnection) GetGatewayGeneration() string`

GetGatewayGeneration returns the GatewayGeneration field if non-nil, zero value otherwise.

### GetGatewayGenerationOk

`func (o *AddAzureVpnConnection) GetGatewayGenerationOk() (*string, bool)`

GetGatewayGenerationOk returns a tuple with the GatewayGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayGeneration

`func (o *AddAzureVpnConnection) SetGatewayGeneration(v string)`

SetGatewayGeneration sets GatewayGeneration field to given value.

### HasGatewayGeneration

`func (o *AddAzureVpnConnection) HasGatewayGeneration() bool`

HasGatewayGeneration returns a boolean if a field has been set.

### SetGatewayGenerationNil

`func (o *AddAzureVpnConnection) SetGatewayGenerationNil(b bool)`

 SetGatewayGenerationNil sets the value for GatewayGeneration to be an explicit nil

### UnsetGatewayGeneration
`func (o *AddAzureVpnConnection) UnsetGatewayGeneration()`

UnsetGatewayGeneration ensures that no value is present for GatewayGeneration, not even an explicit nil
### GetName

`func (o *AddAzureVpnConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddAzureVpnConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddAzureVpnConnection) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *AddAzureVpnConnection) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AddAzureVpnConnection) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AddAzureVpnConnection) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetVdaSubnet

`func (o *AddAzureVpnConnection) GetVdaSubnet() ConnectionSubnet`

GetVdaSubnet returns the VdaSubnet field if non-nil, zero value otherwise.

### GetVdaSubnetOk

`func (o *AddAzureVpnConnection) GetVdaSubnetOk() (*ConnectionSubnet, bool)`

GetVdaSubnetOk returns a tuple with the VdaSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdaSubnet

`func (o *AddAzureVpnConnection) SetVdaSubnet(v ConnectionSubnet)`

SetVdaSubnet sets VdaSubnet field to given value.


### GetGatewaySubnet

`func (o *AddAzureVpnConnection) GetGatewaySubnet() ConnectionSubnet`

GetGatewaySubnet returns the GatewaySubnet field if non-nil, zero value otherwise.

### GetGatewaySubnetOk

`func (o *AddAzureVpnConnection) GetGatewaySubnetOk() (*ConnectionSubnet, bool)`

GetGatewaySubnetOk returns a tuple with the GatewaySubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySubnet

`func (o *AddAzureVpnConnection) SetGatewaySubnet(v ConnectionSubnet)`

SetGatewaySubnet sets GatewaySubnet field to given value.


### GetGatewayIP

`func (o *AddAzureVpnConnection) GetGatewayIP() string`

GetGatewayIP returns the GatewayIP field if non-nil, zero value otherwise.

### GetGatewayIPOk

`func (o *AddAzureVpnConnection) GetGatewayIPOk() (*string, bool)`

GetGatewayIPOk returns a tuple with the GatewayIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIP

`func (o *AddAzureVpnConnection) SetGatewayIP(v string)`

SetGatewayIP sets GatewayIP field to given value.


### GetDnsServers

`func (o *AddAzureVpnConnection) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *AddAzureVpnConnection) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *AddAzureVpnConnection) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *AddAzureVpnConnection) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### SetDnsServersNil

`func (o *AddAzureVpnConnection) SetDnsServersNil(b bool)`

 SetDnsServersNil sets the value for DnsServers to be an explicit nil

### UnsetDnsServers
`func (o *AddAzureVpnConnection) UnsetDnsServers()`

UnsetDnsServers ensures that no value is present for DnsServers, not even an explicit nil
### GetLocalAddresses

`func (o *AddAzureVpnConnection) GetLocalAddresses() []ConnectionSubnet`

GetLocalAddresses returns the LocalAddresses field if non-nil, zero value otherwise.

### GetLocalAddressesOk

`func (o *AddAzureVpnConnection) GetLocalAddressesOk() (*[]ConnectionSubnet, bool)`

GetLocalAddressesOk returns a tuple with the LocalAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddresses

`func (o *AddAzureVpnConnection) SetLocalAddresses(v []ConnectionSubnet)`

SetLocalAddresses sets LocalAddresses field to given value.

### HasLocalAddresses

`func (o *AddAzureVpnConnection) HasLocalAddresses() bool`

HasLocalAddresses returns a boolean if a field has been set.

### SetLocalAddressesNil

`func (o *AddAzureVpnConnection) SetLocalAddressesNil(b bool)`

 SetLocalAddressesNil sets the value for LocalAddresses to be an explicit nil

### UnsetLocalAddresses
`func (o *AddAzureVpnConnection) UnsetLocalAddresses()`

UnsetLocalAddresses ensures that no value is present for LocalAddresses, not even an explicit nil
### GetSharedKey

`func (o *AddAzureVpnConnection) GetSharedKey() string`

GetSharedKey returns the SharedKey field if non-nil, zero value otherwise.

### GetSharedKeyOk

`func (o *AddAzureVpnConnection) GetSharedKeyOk() (*string, bool)`

GetSharedKeyOk returns a tuple with the SharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedKey

`func (o *AddAzureVpnConnection) SetSharedKey(v string)`

SetSharedKey sets SharedKey field to given value.

### HasSharedKey

`func (o *AddAzureVpnConnection) HasSharedKey() bool`

HasSharedKey returns a boolean if a field has been set.

### SetSharedKeyNil

`func (o *AddAzureVpnConnection) SetSharedKeyNil(b bool)`

 SetSharedKeyNil sets the value for SharedKey to be an explicit nil

### UnsetSharedKey
`func (o *AddAzureVpnConnection) UnsetSharedKey()`

UnsetSharedKey ensures that no value is present for SharedKey, not even an explicit nil
### GetManagedSubscriptionId

`func (o *AddAzureVpnConnection) GetManagedSubscriptionId() string`

GetManagedSubscriptionId returns the ManagedSubscriptionId field if non-nil, zero value otherwise.

### GetManagedSubscriptionIdOk

`func (o *AddAzureVpnConnection) GetManagedSubscriptionIdOk() (*string, bool)`

GetManagedSubscriptionIdOk returns a tuple with the ManagedSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedSubscriptionId

`func (o *AddAzureVpnConnection) SetManagedSubscriptionId(v string)`

SetManagedSubscriptionId sets ManagedSubscriptionId field to given value.

### HasManagedSubscriptionId

`func (o *AddAzureVpnConnection) HasManagedSubscriptionId() bool`

HasManagedSubscriptionId returns a boolean if a field has been set.

### SetManagedSubscriptionIdNil

`func (o *AddAzureVpnConnection) SetManagedSubscriptionIdNil(b bool)`

 SetManagedSubscriptionIdNil sets the value for ManagedSubscriptionId to be an explicit nil

### UnsetManagedSubscriptionId
`func (o *AddAzureVpnConnection) UnsetManagedSubscriptionId()`

UnsetManagedSubscriptionId ensures that no value is present for ManagedSubscriptionId, not even an explicit nil
### GetRoutes

`func (o *AddAzureVpnConnection) GetRoutes() []AzureRoute`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *AddAzureVpnConnection) GetRoutesOk() (*[]AzureRoute, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *AddAzureVpnConnection) SetRoutes(v []AzureRoute)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *AddAzureVpnConnection) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### SetRoutesNil

`func (o *AddAzureVpnConnection) SetRoutesNil(b bool)`

 SetRoutesNil sets the value for Routes to be an explicit nil

### UnsetRoutes
`func (o *AddAzureVpnConnection) UnsetRoutes()`

UnsetRoutes ensures that no value is present for Routes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


