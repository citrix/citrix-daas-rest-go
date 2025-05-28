# AzureVirtualNetworkGatewayMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageBandwidth** | Pointer to **float64** | The average Gateway S2S bandwidth in bytes/second | [optional] [readonly] 
**TunnelAverageBandwidth** | Pointer to **float64** | The average tunnel bandwidth in bytes/second | [optional] [readonly] 
**TunnelIngressBytes** | Pointer to **float64** | The average tunnel ingress bytes | [optional] [readonly] 
**TunnelEgressBytes** | Pointer to **float64** | The average tunnel egress bytes | [optional] [readonly] 
**TunnelEgressPackets** | Pointer to **float64** | The average tunnel egress packets | [optional] [readonly] 
**TunnelIngressPackets** | Pointer to **float64** | The average tunnel ingress packets | [optional] [readonly] 

## Methods

### NewAzureVirtualNetworkGatewayMetrics

`func NewAzureVirtualNetworkGatewayMetrics() *AzureVirtualNetworkGatewayMetrics`

NewAzureVirtualNetworkGatewayMetrics instantiates a new AzureVirtualNetworkGatewayMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureVirtualNetworkGatewayMetricsWithDefaults

`func NewAzureVirtualNetworkGatewayMetricsWithDefaults() *AzureVirtualNetworkGatewayMetrics`

NewAzureVirtualNetworkGatewayMetricsWithDefaults instantiates a new AzureVirtualNetworkGatewayMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageBandwidth

`func (o *AzureVirtualNetworkGatewayMetrics) GetAverageBandwidth() float64`

GetAverageBandwidth returns the AverageBandwidth field if non-nil, zero value otherwise.

### GetAverageBandwidthOk

`func (o *AzureVirtualNetworkGatewayMetrics) GetAverageBandwidthOk() (*float64, bool)`

GetAverageBandwidthOk returns a tuple with the AverageBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageBandwidth

`func (o *AzureVirtualNetworkGatewayMetrics) SetAverageBandwidth(v float64)`

SetAverageBandwidth sets AverageBandwidth field to given value.

### HasAverageBandwidth

`func (o *AzureVirtualNetworkGatewayMetrics) HasAverageBandwidth() bool`

HasAverageBandwidth returns a boolean if a field has been set.

### GetTunnelAverageBandwidth

`func (o *AzureVirtualNetworkGatewayMetrics) GetTunnelAverageBandwidth() float64`

GetTunnelAverageBandwidth returns the TunnelAverageBandwidth field if non-nil, zero value otherwise.

### GetTunnelAverageBandwidthOk

`func (o *AzureVirtualNetworkGatewayMetrics) GetTunnelAverageBandwidthOk() (*float64, bool)`

GetTunnelAverageBandwidthOk returns a tuple with the TunnelAverageBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelAverageBandwidth

`func (o *AzureVirtualNetworkGatewayMetrics) SetTunnelAverageBandwidth(v float64)`

SetTunnelAverageBandwidth sets TunnelAverageBandwidth field to given value.

### HasTunnelAverageBandwidth

`func (o *AzureVirtualNetworkGatewayMetrics) HasTunnelAverageBandwidth() bool`

HasTunnelAverageBandwidth returns a boolean if a field has been set.

### GetTunnelIngressBytes

`func (o *AzureVirtualNetworkGatewayMetrics) GetTunnelIngressBytes() float64`

GetTunnelIngressBytes returns the TunnelIngressBytes field if non-nil, zero value otherwise.

### GetTunnelIngressBytesOk

`func (o *AzureVirtualNetworkGatewayMetrics) GetTunnelIngressBytesOk() (*float64, bool)`

GetTunnelIngressBytesOk returns a tuple with the TunnelIngressBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelIngressBytes

`func (o *AzureVirtualNetworkGatewayMetrics) SetTunnelIngressBytes(v float64)`

SetTunnelIngressBytes sets TunnelIngressBytes field to given value.

### HasTunnelIngressBytes

`func (o *AzureVirtualNetworkGatewayMetrics) HasTunnelIngressBytes() bool`

HasTunnelIngressBytes returns a boolean if a field has been set.

### GetTunnelEgressBytes

`func (o *AzureVirtualNetworkGatewayMetrics) GetTunnelEgressBytes() float64`

GetTunnelEgressBytes returns the TunnelEgressBytes field if non-nil, zero value otherwise.

### GetTunnelEgressBytesOk

`func (o *AzureVirtualNetworkGatewayMetrics) GetTunnelEgressBytesOk() (*float64, bool)`

GetTunnelEgressBytesOk returns a tuple with the TunnelEgressBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelEgressBytes

`func (o *AzureVirtualNetworkGatewayMetrics) SetTunnelEgressBytes(v float64)`

SetTunnelEgressBytes sets TunnelEgressBytes field to given value.

### HasTunnelEgressBytes

`func (o *AzureVirtualNetworkGatewayMetrics) HasTunnelEgressBytes() bool`

HasTunnelEgressBytes returns a boolean if a field has been set.

### GetTunnelEgressPackets

`func (o *AzureVirtualNetworkGatewayMetrics) GetTunnelEgressPackets() float64`

GetTunnelEgressPackets returns the TunnelEgressPackets field if non-nil, zero value otherwise.

### GetTunnelEgressPacketsOk

`func (o *AzureVirtualNetworkGatewayMetrics) GetTunnelEgressPacketsOk() (*float64, bool)`

GetTunnelEgressPacketsOk returns a tuple with the TunnelEgressPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelEgressPackets

`func (o *AzureVirtualNetworkGatewayMetrics) SetTunnelEgressPackets(v float64)`

SetTunnelEgressPackets sets TunnelEgressPackets field to given value.

### HasTunnelEgressPackets

`func (o *AzureVirtualNetworkGatewayMetrics) HasTunnelEgressPackets() bool`

HasTunnelEgressPackets returns a boolean if a field has been set.

### GetTunnelIngressPackets

`func (o *AzureVirtualNetworkGatewayMetrics) GetTunnelIngressPackets() float64`

GetTunnelIngressPackets returns the TunnelIngressPackets field if non-nil, zero value otherwise.

### GetTunnelIngressPacketsOk

`func (o *AzureVirtualNetworkGatewayMetrics) GetTunnelIngressPacketsOk() (*float64, bool)`

GetTunnelIngressPacketsOk returns a tuple with the TunnelIngressPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelIngressPackets

`func (o *AzureVirtualNetworkGatewayMetrics) SetTunnelIngressPackets(v float64)`

SetTunnelIngressPackets sets TunnelIngressPackets field to given value.

### HasTunnelIngressPackets

`func (o *AzureVirtualNetworkGatewayMetrics) HasTunnelIngressPackets() bool`

HasTunnelIngressPackets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


