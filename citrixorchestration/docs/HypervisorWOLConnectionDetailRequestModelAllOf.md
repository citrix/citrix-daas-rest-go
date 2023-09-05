# HypervisorWOLConnectionDetailRequestModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SccmWakeUpProxy** | Pointer to **bool** | Indicates whether Microsoft System Center Configuration Manager 2012 SP1 Wake-up Proxy is used for power management.  Optional; default is &#x60;false&#x60;. | [optional] 
**WakeOnLanPackets** | Pointer to [**WakeOnLanTransmission**](WakeOnLanTransmission.md) |  | [optional] 

## Methods

### NewHypervisorWOLConnectionDetailRequestModelAllOf

`func NewHypervisorWOLConnectionDetailRequestModelAllOf() *HypervisorWOLConnectionDetailRequestModelAllOf`

NewHypervisorWOLConnectionDetailRequestModelAllOf instantiates a new HypervisorWOLConnectionDetailRequestModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorWOLConnectionDetailRequestModelAllOfWithDefaults

`func NewHypervisorWOLConnectionDetailRequestModelAllOfWithDefaults() *HypervisorWOLConnectionDetailRequestModelAllOf`

NewHypervisorWOLConnectionDetailRequestModelAllOfWithDefaults instantiates a new HypervisorWOLConnectionDetailRequestModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSccmWakeUpProxy

`func (o *HypervisorWOLConnectionDetailRequestModelAllOf) GetSccmWakeUpProxy() bool`

GetSccmWakeUpProxy returns the SccmWakeUpProxy field if non-nil, zero value otherwise.

### GetSccmWakeUpProxyOk

`func (o *HypervisorWOLConnectionDetailRequestModelAllOf) GetSccmWakeUpProxyOk() (*bool, bool)`

GetSccmWakeUpProxyOk returns a tuple with the SccmWakeUpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSccmWakeUpProxy

`func (o *HypervisorWOLConnectionDetailRequestModelAllOf) SetSccmWakeUpProxy(v bool)`

SetSccmWakeUpProxy sets SccmWakeUpProxy field to given value.

### HasSccmWakeUpProxy

`func (o *HypervisorWOLConnectionDetailRequestModelAllOf) HasSccmWakeUpProxy() bool`

HasSccmWakeUpProxy returns a boolean if a field has been set.

### GetWakeOnLanPackets

`func (o *HypervisorWOLConnectionDetailRequestModelAllOf) GetWakeOnLanPackets() WakeOnLanTransmission`

GetWakeOnLanPackets returns the WakeOnLanPackets field if non-nil, zero value otherwise.

### GetWakeOnLanPacketsOk

`func (o *HypervisorWOLConnectionDetailRequestModelAllOf) GetWakeOnLanPacketsOk() (*WakeOnLanTransmission, bool)`

GetWakeOnLanPacketsOk returns a tuple with the WakeOnLanPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWakeOnLanPackets

`func (o *HypervisorWOLConnectionDetailRequestModelAllOf) SetWakeOnLanPackets(v WakeOnLanTransmission)`

SetWakeOnLanPackets sets WakeOnLanPackets field to given value.

### HasWakeOnLanPackets

`func (o *HypervisorWOLConnectionDetailRequestModelAllOf) HasWakeOnLanPackets() bool`

HasWakeOnLanPackets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


