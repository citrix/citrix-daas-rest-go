# DeliveryControllerResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The sid of the delivery controller | [optional] 
**Id** | Pointer to **string** | The id(uuid) of the delivery controller | [optional] 
**MachineName** | Pointer to **string** | The machine name of the delivery controller | [optional] 
**DnsName** | Pointer to **string** | The dns name of the delivery controller | [optional] 
**ServiceHostId** | Pointer to **string** | The service host id of the delivery controller | [optional] 
**ControllerVersion** | Pointer to **string** | The controller version of the delivery controller | [optional] 
**RegisteredDesktops** | Pointer to **int32** | The number of registered desktops in the delivery controller | [optional] 
**LastActivityTime** | Pointer to **string** | The LastActivityTime of the delivery controller | [optional] 
**LastStartTime** | Pointer to **string** | The LastStartTime of the delivery controller | [optional] 
**OSVersion** | Pointer to **string** | The OSVersion of the delivery controller | [optional] 
**OSType** | Pointer to **string** | The OSType of the delivery controller | [optional] 
**ServiceStatuses** | Pointer to [**[]DeliveryControllerServiceStatusResponseModel**](DeliveryControllerServiceStatusResponseModel.md) | The service statuses of the delivery controller. | [optional] 
**ControllerState** | Pointer to [**DeliveryControllerState**](DeliveryControllerState.md) |  | [optional] 
**ZoneId** | Pointer to **string** | The ID of the zone this zonable item belongs to | [optional] 
**ZonableItemId** | Pointer to **string** | The Id of this zonable Item. This is a generic name to access non-generic properties | [optional] 
**ZonableItemType** | Pointer to [**ZonableItemType**](ZonableItemType.md) |  | [optional] 
**ZoneName** | Pointer to **string** | The Zone&#39;s Name | [optional] 
**ZonableItemName** | Pointer to **string** | Tthe name of this zonable Item | [optional] 
**ZonableItemDescription** | Pointer to **string** | The description of this zonable Item | [optional] 

## Methods

### NewDeliveryControllerResponseModel

`func NewDeliveryControllerResponseModel() *DeliveryControllerResponseModel`

NewDeliveryControllerResponseModel instantiates a new DeliveryControllerResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryControllerResponseModelWithDefaults

`func NewDeliveryControllerResponseModelWithDefaults() *DeliveryControllerResponseModel`

NewDeliveryControllerResponseModelWithDefaults instantiates a new DeliveryControllerResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSid

`func (o *DeliveryControllerResponseModel) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *DeliveryControllerResponseModel) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *DeliveryControllerResponseModel) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *DeliveryControllerResponseModel) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetId

`func (o *DeliveryControllerResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeliveryControllerResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeliveryControllerResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeliveryControllerResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMachineName

`func (o *DeliveryControllerResponseModel) GetMachineName() string`

GetMachineName returns the MachineName field if non-nil, zero value otherwise.

### GetMachineNameOk

`func (o *DeliveryControllerResponseModel) GetMachineNameOk() (*string, bool)`

GetMachineNameOk returns a tuple with the MachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineName

`func (o *DeliveryControllerResponseModel) SetMachineName(v string)`

SetMachineName sets MachineName field to given value.

### HasMachineName

`func (o *DeliveryControllerResponseModel) HasMachineName() bool`

HasMachineName returns a boolean if a field has been set.

### GetDnsName

`func (o *DeliveryControllerResponseModel) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *DeliveryControllerResponseModel) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *DeliveryControllerResponseModel) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *DeliveryControllerResponseModel) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetServiceHostId

`func (o *DeliveryControllerResponseModel) GetServiceHostId() string`

GetServiceHostId returns the ServiceHostId field if non-nil, zero value otherwise.

### GetServiceHostIdOk

`func (o *DeliveryControllerResponseModel) GetServiceHostIdOk() (*string, bool)`

GetServiceHostIdOk returns a tuple with the ServiceHostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHostId

`func (o *DeliveryControllerResponseModel) SetServiceHostId(v string)`

SetServiceHostId sets ServiceHostId field to given value.

### HasServiceHostId

`func (o *DeliveryControllerResponseModel) HasServiceHostId() bool`

HasServiceHostId returns a boolean if a field has been set.

### GetControllerVersion

`func (o *DeliveryControllerResponseModel) GetControllerVersion() string`

GetControllerVersion returns the ControllerVersion field if non-nil, zero value otherwise.

### GetControllerVersionOk

`func (o *DeliveryControllerResponseModel) GetControllerVersionOk() (*string, bool)`

GetControllerVersionOk returns a tuple with the ControllerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerVersion

`func (o *DeliveryControllerResponseModel) SetControllerVersion(v string)`

SetControllerVersion sets ControllerVersion field to given value.

### HasControllerVersion

`func (o *DeliveryControllerResponseModel) HasControllerVersion() bool`

HasControllerVersion returns a boolean if a field has been set.

### GetRegisteredDesktops

`func (o *DeliveryControllerResponseModel) GetRegisteredDesktops() int32`

GetRegisteredDesktops returns the RegisteredDesktops field if non-nil, zero value otherwise.

### GetRegisteredDesktopsOk

`func (o *DeliveryControllerResponseModel) GetRegisteredDesktopsOk() (*int32, bool)`

GetRegisteredDesktopsOk returns a tuple with the RegisteredDesktops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDesktops

`func (o *DeliveryControllerResponseModel) SetRegisteredDesktops(v int32)`

SetRegisteredDesktops sets RegisteredDesktops field to given value.

### HasRegisteredDesktops

`func (o *DeliveryControllerResponseModel) HasRegisteredDesktops() bool`

HasRegisteredDesktops returns a boolean if a field has been set.

### GetLastActivityTime

`func (o *DeliveryControllerResponseModel) GetLastActivityTime() string`

GetLastActivityTime returns the LastActivityTime field if non-nil, zero value otherwise.

### GetLastActivityTimeOk

`func (o *DeliveryControllerResponseModel) GetLastActivityTimeOk() (*string, bool)`

GetLastActivityTimeOk returns a tuple with the LastActivityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityTime

`func (o *DeliveryControllerResponseModel) SetLastActivityTime(v string)`

SetLastActivityTime sets LastActivityTime field to given value.

### HasLastActivityTime

`func (o *DeliveryControllerResponseModel) HasLastActivityTime() bool`

HasLastActivityTime returns a boolean if a field has been set.

### GetLastStartTime

`func (o *DeliveryControllerResponseModel) GetLastStartTime() string`

GetLastStartTime returns the LastStartTime field if non-nil, zero value otherwise.

### GetLastStartTimeOk

`func (o *DeliveryControllerResponseModel) GetLastStartTimeOk() (*string, bool)`

GetLastStartTimeOk returns a tuple with the LastStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStartTime

`func (o *DeliveryControllerResponseModel) SetLastStartTime(v string)`

SetLastStartTime sets LastStartTime field to given value.

### HasLastStartTime

`func (o *DeliveryControllerResponseModel) HasLastStartTime() bool`

HasLastStartTime returns a boolean if a field has been set.

### GetOSVersion

`func (o *DeliveryControllerResponseModel) GetOSVersion() string`

GetOSVersion returns the OSVersion field if non-nil, zero value otherwise.

### GetOSVersionOk

`func (o *DeliveryControllerResponseModel) GetOSVersionOk() (*string, bool)`

GetOSVersionOk returns a tuple with the OSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOSVersion

`func (o *DeliveryControllerResponseModel) SetOSVersion(v string)`

SetOSVersion sets OSVersion field to given value.

### HasOSVersion

`func (o *DeliveryControllerResponseModel) HasOSVersion() bool`

HasOSVersion returns a boolean if a field has been set.

### GetOSType

`func (o *DeliveryControllerResponseModel) GetOSType() string`

GetOSType returns the OSType field if non-nil, zero value otherwise.

### GetOSTypeOk

`func (o *DeliveryControllerResponseModel) GetOSTypeOk() (*string, bool)`

GetOSTypeOk returns a tuple with the OSType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOSType

`func (o *DeliveryControllerResponseModel) SetOSType(v string)`

SetOSType sets OSType field to given value.

### HasOSType

`func (o *DeliveryControllerResponseModel) HasOSType() bool`

HasOSType returns a boolean if a field has been set.

### GetServiceStatuses

`func (o *DeliveryControllerResponseModel) GetServiceStatuses() []DeliveryControllerServiceStatusResponseModel`

GetServiceStatuses returns the ServiceStatuses field if non-nil, zero value otherwise.

### GetServiceStatusesOk

`func (o *DeliveryControllerResponseModel) GetServiceStatusesOk() (*[]DeliveryControllerServiceStatusResponseModel, bool)`

GetServiceStatusesOk returns a tuple with the ServiceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatuses

`func (o *DeliveryControllerResponseModel) SetServiceStatuses(v []DeliveryControllerServiceStatusResponseModel)`

SetServiceStatuses sets ServiceStatuses field to given value.

### HasServiceStatuses

`func (o *DeliveryControllerResponseModel) HasServiceStatuses() bool`

HasServiceStatuses returns a boolean if a field has been set.

### GetControllerState

`func (o *DeliveryControllerResponseModel) GetControllerState() DeliveryControllerState`

GetControllerState returns the ControllerState field if non-nil, zero value otherwise.

### GetControllerStateOk

`func (o *DeliveryControllerResponseModel) GetControllerStateOk() (*DeliveryControllerState, bool)`

GetControllerStateOk returns a tuple with the ControllerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerState

`func (o *DeliveryControllerResponseModel) SetControllerState(v DeliveryControllerState)`

SetControllerState sets ControllerState field to given value.

### HasControllerState

`func (o *DeliveryControllerResponseModel) HasControllerState() bool`

HasControllerState returns a boolean if a field has been set.

### GetZoneId

`func (o *DeliveryControllerResponseModel) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *DeliveryControllerResponseModel) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *DeliveryControllerResponseModel) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *DeliveryControllerResponseModel) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetZonableItemId

`func (o *DeliveryControllerResponseModel) GetZonableItemId() string`

GetZonableItemId returns the ZonableItemId field if non-nil, zero value otherwise.

### GetZonableItemIdOk

`func (o *DeliveryControllerResponseModel) GetZonableItemIdOk() (*string, bool)`

GetZonableItemIdOk returns a tuple with the ZonableItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonableItemId

`func (o *DeliveryControllerResponseModel) SetZonableItemId(v string)`

SetZonableItemId sets ZonableItemId field to given value.

### HasZonableItemId

`func (o *DeliveryControllerResponseModel) HasZonableItemId() bool`

HasZonableItemId returns a boolean if a field has been set.

### GetZonableItemType

`func (o *DeliveryControllerResponseModel) GetZonableItemType() ZonableItemType`

GetZonableItemType returns the ZonableItemType field if non-nil, zero value otherwise.

### GetZonableItemTypeOk

`func (o *DeliveryControllerResponseModel) GetZonableItemTypeOk() (*ZonableItemType, bool)`

GetZonableItemTypeOk returns a tuple with the ZonableItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonableItemType

`func (o *DeliveryControllerResponseModel) SetZonableItemType(v ZonableItemType)`

SetZonableItemType sets ZonableItemType field to given value.

### HasZonableItemType

`func (o *DeliveryControllerResponseModel) HasZonableItemType() bool`

HasZonableItemType returns a boolean if a field has been set.

### GetZoneName

`func (o *DeliveryControllerResponseModel) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *DeliveryControllerResponseModel) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *DeliveryControllerResponseModel) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *DeliveryControllerResponseModel) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### GetZonableItemName

`func (o *DeliveryControllerResponseModel) GetZonableItemName() string`

GetZonableItemName returns the ZonableItemName field if non-nil, zero value otherwise.

### GetZonableItemNameOk

`func (o *DeliveryControllerResponseModel) GetZonableItemNameOk() (*string, bool)`

GetZonableItemNameOk returns a tuple with the ZonableItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonableItemName

`func (o *DeliveryControllerResponseModel) SetZonableItemName(v string)`

SetZonableItemName sets ZonableItemName field to given value.

### HasZonableItemName

`func (o *DeliveryControllerResponseModel) HasZonableItemName() bool`

HasZonableItemName returns a boolean if a field has been set.

### GetZonableItemDescription

`func (o *DeliveryControllerResponseModel) GetZonableItemDescription() string`

GetZonableItemDescription returns the ZonableItemDescription field if non-nil, zero value otherwise.

### GetZonableItemDescriptionOk

`func (o *DeliveryControllerResponseModel) GetZonableItemDescriptionOk() (*string, bool)`

GetZonableItemDescriptionOk returns a tuple with the ZonableItemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonableItemDescription

`func (o *DeliveryControllerResponseModel) SetZonableItemDescription(v string)`

SetZonableItemDescription sets ZonableItemDescription field to given value.

### HasZonableItemDescription

`func (o *DeliveryControllerResponseModel) HasZonableItemDescription() bool`

HasZonableItemDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


