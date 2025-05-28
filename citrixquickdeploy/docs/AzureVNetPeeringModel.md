# AzureVNetPeeringModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeeringName** | Pointer to **string** | A unique name for this vnet peering | [optional] 
**CitrixVNet** | Pointer to [**AzureVNet**](AzureVNet.md) | The first VNet, the one owned by Citrix | [optional] 
**CustomerVNet** | Pointer to [**AzureVNet**](AzureVNet.md) | The second VNet which the customers owns | [optional] 
**UseCustomerGateway** | Pointer to **bool** | Indicates if the vpn peering will use the customer&#39;s gateway | [optional] 
**AllowForwardedTraffic** | Pointer to **bool** | Indicates if traffic forwarded from another peer should be allowed into the citrix created vnet  This does not apply to traffic from a gateway on a customer side  When using this feature, a route table entry will need to be added | [optional] 

## Methods

### NewAzureVNetPeeringModel

`func NewAzureVNetPeeringModel() *AzureVNetPeeringModel`

NewAzureVNetPeeringModel instantiates a new AzureVNetPeeringModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureVNetPeeringModelWithDefaults

`func NewAzureVNetPeeringModelWithDefaults() *AzureVNetPeeringModel`

NewAzureVNetPeeringModelWithDefaults instantiates a new AzureVNetPeeringModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeeringName

`func (o *AzureVNetPeeringModel) GetPeeringName() string`

GetPeeringName returns the PeeringName field if non-nil, zero value otherwise.

### GetPeeringNameOk

`func (o *AzureVNetPeeringModel) GetPeeringNameOk() (*string, bool)`

GetPeeringNameOk returns a tuple with the PeeringName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeeringName

`func (o *AzureVNetPeeringModel) SetPeeringName(v string)`

SetPeeringName sets PeeringName field to given value.

### HasPeeringName

`func (o *AzureVNetPeeringModel) HasPeeringName() bool`

HasPeeringName returns a boolean if a field has been set.

### GetCitrixVNet

`func (o *AzureVNetPeeringModel) GetCitrixVNet() AzureVNet`

GetCitrixVNet returns the CitrixVNet field if non-nil, zero value otherwise.

### GetCitrixVNetOk

`func (o *AzureVNetPeeringModel) GetCitrixVNetOk() (*AzureVNet, bool)`

GetCitrixVNetOk returns a tuple with the CitrixVNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitrixVNet

`func (o *AzureVNetPeeringModel) SetCitrixVNet(v AzureVNet)`

SetCitrixVNet sets CitrixVNet field to given value.

### HasCitrixVNet

`func (o *AzureVNetPeeringModel) HasCitrixVNet() bool`

HasCitrixVNet returns a boolean if a field has been set.

### GetCustomerVNet

`func (o *AzureVNetPeeringModel) GetCustomerVNet() AzureVNet`

GetCustomerVNet returns the CustomerVNet field if non-nil, zero value otherwise.

### GetCustomerVNetOk

`func (o *AzureVNetPeeringModel) GetCustomerVNetOk() (*AzureVNet, bool)`

GetCustomerVNetOk returns a tuple with the CustomerVNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerVNet

`func (o *AzureVNetPeeringModel) SetCustomerVNet(v AzureVNet)`

SetCustomerVNet sets CustomerVNet field to given value.

### HasCustomerVNet

`func (o *AzureVNetPeeringModel) HasCustomerVNet() bool`

HasCustomerVNet returns a boolean if a field has been set.

### GetUseCustomerGateway

`func (o *AzureVNetPeeringModel) GetUseCustomerGateway() bool`

GetUseCustomerGateway returns the UseCustomerGateway field if non-nil, zero value otherwise.

### GetUseCustomerGatewayOk

`func (o *AzureVNetPeeringModel) GetUseCustomerGatewayOk() (*bool, bool)`

GetUseCustomerGatewayOk returns a tuple with the UseCustomerGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCustomerGateway

`func (o *AzureVNetPeeringModel) SetUseCustomerGateway(v bool)`

SetUseCustomerGateway sets UseCustomerGateway field to given value.

### HasUseCustomerGateway

`func (o *AzureVNetPeeringModel) HasUseCustomerGateway() bool`

HasUseCustomerGateway returns a boolean if a field has been set.

### GetAllowForwardedTraffic

`func (o *AzureVNetPeeringModel) GetAllowForwardedTraffic() bool`

GetAllowForwardedTraffic returns the AllowForwardedTraffic field if non-nil, zero value otherwise.

### GetAllowForwardedTrafficOk

`func (o *AzureVNetPeeringModel) GetAllowForwardedTrafficOk() (*bool, bool)`

GetAllowForwardedTrafficOk returns a tuple with the AllowForwardedTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowForwardedTraffic

`func (o *AzureVNetPeeringModel) SetAllowForwardedTraffic(v bool)`

SetAllowForwardedTraffic sets AllowForwardedTraffic field to given value.

### HasAllowForwardedTraffic

`func (o *AzureVNetPeeringModel) HasAllowForwardedTraffic() bool`

HasAllowForwardedTraffic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


