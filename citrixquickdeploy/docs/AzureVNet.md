# AzureVNet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** | ID of the Azure Subscription where VNet is configured | 
**ResourceGroup** | **string** | Name of the Resource Group the VNet is associated with | 
**Name** | **string** | Name of the VNet | 
**Location** | Pointer to **string** | Azure region where the VNet is located | [optional] 
**Subnets** | Pointer to [**[]AzureSubnet**](AzureSubnet.md) | Subnets that have been configured for this VNet | [optional] 
**DnsServers** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAzureVNet

`func NewAzureVNet(subscriptionId string, resourceGroup string, name string, ) *AzureVNet`

NewAzureVNet instantiates a new AzureVNet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureVNetWithDefaults

`func NewAzureVNetWithDefaults() *AzureVNet`

NewAzureVNetWithDefaults instantiates a new AzureVNet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *AzureVNet) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureVNet) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureVNet) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetResourceGroup

`func (o *AzureVNet) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *AzureVNet) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *AzureVNet) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.


### GetName

`func (o *AzureVNet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureVNet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureVNet) SetName(v string)`

SetName sets Name field to given value.


### GetLocation

`func (o *AzureVNet) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AzureVNet) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AzureVNet) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AzureVNet) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetSubnets

`func (o *AzureVNet) GetSubnets() []AzureSubnet`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *AzureVNet) GetSubnetsOk() (*[]AzureSubnet, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *AzureVNet) SetSubnets(v []AzureSubnet)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *AzureVNet) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetDnsServers

`func (o *AzureVNet) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *AzureVNet) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *AzureVNet) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *AzureVNet) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


