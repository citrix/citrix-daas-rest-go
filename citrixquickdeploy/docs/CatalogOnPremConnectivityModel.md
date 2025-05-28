# CatalogOnPremConnectivityModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VnetPeeringId** | Pointer to **string** | ID of the Vnet Peering that will be used for on-prem connections | [optional] 
**VpnConnectionId** | Pointer to **string** | ID of the VPN Connection that will be used by this catalog | [optional] 

## Methods

### NewCatalogOnPremConnectivityModel

`func NewCatalogOnPremConnectivityModel() *CatalogOnPremConnectivityModel`

NewCatalogOnPremConnectivityModel instantiates a new CatalogOnPremConnectivityModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogOnPremConnectivityModelWithDefaults

`func NewCatalogOnPremConnectivityModelWithDefaults() *CatalogOnPremConnectivityModel`

NewCatalogOnPremConnectivityModelWithDefaults instantiates a new CatalogOnPremConnectivityModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVnetPeeringId

`func (o *CatalogOnPremConnectivityModel) GetVnetPeeringId() string`

GetVnetPeeringId returns the VnetPeeringId field if non-nil, zero value otherwise.

### GetVnetPeeringIdOk

`func (o *CatalogOnPremConnectivityModel) GetVnetPeeringIdOk() (*string, bool)`

GetVnetPeeringIdOk returns a tuple with the VnetPeeringId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetPeeringId

`func (o *CatalogOnPremConnectivityModel) SetVnetPeeringId(v string)`

SetVnetPeeringId sets VnetPeeringId field to given value.

### HasVnetPeeringId

`func (o *CatalogOnPremConnectivityModel) HasVnetPeeringId() bool`

HasVnetPeeringId returns a boolean if a field has been set.

### GetVpnConnectionId

`func (o *CatalogOnPremConnectivityModel) GetVpnConnectionId() string`

GetVpnConnectionId returns the VpnConnectionId field if non-nil, zero value otherwise.

### GetVpnConnectionIdOk

`func (o *CatalogOnPremConnectivityModel) GetVpnConnectionIdOk() (*string, bool)`

GetVpnConnectionIdOk returns a tuple with the VpnConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConnectionId

`func (o *CatalogOnPremConnectivityModel) SetVpnConnectionId(v string)`

SetVpnConnectionId sets VpnConnectionId field to given value.

### HasVpnConnectionId

`func (o *CatalogOnPremConnectivityModel) HasVpnConnectionId() bool`

HasVpnConnectionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


