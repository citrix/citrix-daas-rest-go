# EditHypervisorWOLConnectionRequestModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SccmWakeUpProxy** | Pointer to **bool** | Specifies whether to use Microsoft System Center Configuration Manager 2012 SP1 Wake-up Proxy for power management.  Optional.  If not specified, will not be changed. | [optional] 
**WakeOnLanPackets** | Pointer to [**WakeOnLanTransmission**](WakeOnLanTransmission.md) |  | [optional] 

## Methods

### NewEditHypervisorWOLConnectionRequestModelAllOf

`func NewEditHypervisorWOLConnectionRequestModelAllOf() *EditHypervisorWOLConnectionRequestModelAllOf`

NewEditHypervisorWOLConnectionRequestModelAllOf instantiates a new EditHypervisorWOLConnectionRequestModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditHypervisorWOLConnectionRequestModelAllOfWithDefaults

`func NewEditHypervisorWOLConnectionRequestModelAllOfWithDefaults() *EditHypervisorWOLConnectionRequestModelAllOf`

NewEditHypervisorWOLConnectionRequestModelAllOfWithDefaults instantiates a new EditHypervisorWOLConnectionRequestModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSccmWakeUpProxy

`func (o *EditHypervisorWOLConnectionRequestModelAllOf) GetSccmWakeUpProxy() bool`

GetSccmWakeUpProxy returns the SccmWakeUpProxy field if non-nil, zero value otherwise.

### GetSccmWakeUpProxyOk

`func (o *EditHypervisorWOLConnectionRequestModelAllOf) GetSccmWakeUpProxyOk() (*bool, bool)`

GetSccmWakeUpProxyOk returns a tuple with the SccmWakeUpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSccmWakeUpProxy

`func (o *EditHypervisorWOLConnectionRequestModelAllOf) SetSccmWakeUpProxy(v bool)`

SetSccmWakeUpProxy sets SccmWakeUpProxy field to given value.

### HasSccmWakeUpProxy

`func (o *EditHypervisorWOLConnectionRequestModelAllOf) HasSccmWakeUpProxy() bool`

HasSccmWakeUpProxy returns a boolean if a field has been set.

### GetWakeOnLanPackets

`func (o *EditHypervisorWOLConnectionRequestModelAllOf) GetWakeOnLanPackets() WakeOnLanTransmission`

GetWakeOnLanPackets returns the WakeOnLanPackets field if non-nil, zero value otherwise.

### GetWakeOnLanPacketsOk

`func (o *EditHypervisorWOLConnectionRequestModelAllOf) GetWakeOnLanPacketsOk() (*WakeOnLanTransmission, bool)`

GetWakeOnLanPacketsOk returns a tuple with the WakeOnLanPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWakeOnLanPackets

`func (o *EditHypervisorWOLConnectionRequestModelAllOf) SetWakeOnLanPackets(v WakeOnLanTransmission)`

SetWakeOnLanPackets sets WakeOnLanPackets field to given value.

### HasWakeOnLanPackets

`func (o *EditHypervisorWOLConnectionRequestModelAllOf) HasWakeOnLanPackets() bool`

HasWakeOnLanPackets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


