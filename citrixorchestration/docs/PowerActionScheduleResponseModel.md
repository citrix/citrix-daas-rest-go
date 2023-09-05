# PowerActionScheduleResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**SupportedPowerAction**](SupportedPowerAction.md) |  | [optional] 
**ActionDueTime** | Pointer to **time.Time** |  | [optional] 
**DNSName** | Pointer to **string** |  | [optional] 
**HostedMachineName** | Pointer to **string** |  | [optional] 
**HypervisorConnectionUid** | Pointer to **int32** |  | [optional] 
**MachineName** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **int64** |  | [optional] 

## Methods

### NewPowerActionScheduleResponseModel

`func NewPowerActionScheduleResponseModel() *PowerActionScheduleResponseModel`

NewPowerActionScheduleResponseModel instantiates a new PowerActionScheduleResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerActionScheduleResponseModelWithDefaults

`func NewPowerActionScheduleResponseModelWithDefaults() *PowerActionScheduleResponseModel`

NewPowerActionScheduleResponseModelWithDefaults instantiates a new PowerActionScheduleResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PowerActionScheduleResponseModel) GetAction() SupportedPowerAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PowerActionScheduleResponseModel) GetActionOk() (*SupportedPowerAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PowerActionScheduleResponseModel) SetAction(v SupportedPowerAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *PowerActionScheduleResponseModel) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActionDueTime

`func (o *PowerActionScheduleResponseModel) GetActionDueTime() time.Time`

GetActionDueTime returns the ActionDueTime field if non-nil, zero value otherwise.

### GetActionDueTimeOk

`func (o *PowerActionScheduleResponseModel) GetActionDueTimeOk() (*time.Time, bool)`

GetActionDueTimeOk returns a tuple with the ActionDueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDueTime

`func (o *PowerActionScheduleResponseModel) SetActionDueTime(v time.Time)`

SetActionDueTime sets ActionDueTime field to given value.

### HasActionDueTime

`func (o *PowerActionScheduleResponseModel) HasActionDueTime() bool`

HasActionDueTime returns a boolean if a field has been set.

### GetDNSName

`func (o *PowerActionScheduleResponseModel) GetDNSName() string`

GetDNSName returns the DNSName field if non-nil, zero value otherwise.

### GetDNSNameOk

`func (o *PowerActionScheduleResponseModel) GetDNSNameOk() (*string, bool)`

GetDNSNameOk returns a tuple with the DNSName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDNSName

`func (o *PowerActionScheduleResponseModel) SetDNSName(v string)`

SetDNSName sets DNSName field to given value.

### HasDNSName

`func (o *PowerActionScheduleResponseModel) HasDNSName() bool`

HasDNSName returns a boolean if a field has been set.

### GetHostedMachineName

`func (o *PowerActionScheduleResponseModel) GetHostedMachineName() string`

GetHostedMachineName returns the HostedMachineName field if non-nil, zero value otherwise.

### GetHostedMachineNameOk

`func (o *PowerActionScheduleResponseModel) GetHostedMachineNameOk() (*string, bool)`

GetHostedMachineNameOk returns a tuple with the HostedMachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedMachineName

`func (o *PowerActionScheduleResponseModel) SetHostedMachineName(v string)`

SetHostedMachineName sets HostedMachineName field to given value.

### HasHostedMachineName

`func (o *PowerActionScheduleResponseModel) HasHostedMachineName() bool`

HasHostedMachineName returns a boolean if a field has been set.

### GetHypervisorConnectionUid

`func (o *PowerActionScheduleResponseModel) GetHypervisorConnectionUid() int32`

GetHypervisorConnectionUid returns the HypervisorConnectionUid field if non-nil, zero value otherwise.

### GetHypervisorConnectionUidOk

`func (o *PowerActionScheduleResponseModel) GetHypervisorConnectionUidOk() (*int32, bool)`

GetHypervisorConnectionUidOk returns a tuple with the HypervisorConnectionUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorConnectionUid

`func (o *PowerActionScheduleResponseModel) SetHypervisorConnectionUid(v int32)`

SetHypervisorConnectionUid sets HypervisorConnectionUid field to given value.

### HasHypervisorConnectionUid

`func (o *PowerActionScheduleResponseModel) HasHypervisorConnectionUid() bool`

HasHypervisorConnectionUid returns a boolean if a field has been set.

### GetMachineName

`func (o *PowerActionScheduleResponseModel) GetMachineName() string`

GetMachineName returns the MachineName field if non-nil, zero value otherwise.

### GetMachineNameOk

`func (o *PowerActionScheduleResponseModel) GetMachineNameOk() (*string, bool)`

GetMachineNameOk returns a tuple with the MachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineName

`func (o *PowerActionScheduleResponseModel) SetMachineName(v string)`

SetMachineName sets MachineName field to given value.

### HasMachineName

`func (o *PowerActionScheduleResponseModel) HasMachineName() bool`

HasMachineName returns a boolean if a field has been set.

### GetUid

`func (o *PowerActionScheduleResponseModel) GetUid() int64`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *PowerActionScheduleResponseModel) GetUidOk() (*int64, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *PowerActionScheduleResponseModel) SetUid(v int64)`

SetUid sets Uid field to given value.

### HasUid

`func (o *PowerActionScheduleResponseModel) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


