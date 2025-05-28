# ConnectionLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrokeringTime** | Pointer to **time.Time** |  | [optional] 
**BrokeringUserName** | Pointer to **string** |  | [optional] 
**BrokeringUserUPN** | Pointer to **string** |  | [optional] 
**ConnectionFailureReason** | Pointer to [**ConnectionLogConnectionFailureReason**](ConnectionLogConnectionFailureReason.md) |  | [optional] 
**Disconnected** | Pointer to **bool** |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**EstablishmentTime** | Pointer to **time.Time** |  | [optional] 
**MachineDNSName** | Pointer to **string** |  | [optional] 
**MachineName** | Pointer to **string** |  | [optional] 
**MachineUid** | Pointer to **int32** |  | [optional] 
**Uid** | Pointer to **int64** |  | [optional] 

## Methods

### NewConnectionLog

`func NewConnectionLog() *ConnectionLog`

NewConnectionLog instantiates a new ConnectionLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionLogWithDefaults

`func NewConnectionLogWithDefaults() *ConnectionLog`

NewConnectionLogWithDefaults instantiates a new ConnectionLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrokeringTime

`func (o *ConnectionLog) GetBrokeringTime() time.Time`

GetBrokeringTime returns the BrokeringTime field if non-nil, zero value otherwise.

### GetBrokeringTimeOk

`func (o *ConnectionLog) GetBrokeringTimeOk() (*time.Time, bool)`

GetBrokeringTimeOk returns a tuple with the BrokeringTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokeringTime

`func (o *ConnectionLog) SetBrokeringTime(v time.Time)`

SetBrokeringTime sets BrokeringTime field to given value.

### HasBrokeringTime

`func (o *ConnectionLog) HasBrokeringTime() bool`

HasBrokeringTime returns a boolean if a field has been set.

### GetBrokeringUserName

`func (o *ConnectionLog) GetBrokeringUserName() string`

GetBrokeringUserName returns the BrokeringUserName field if non-nil, zero value otherwise.

### GetBrokeringUserNameOk

`func (o *ConnectionLog) GetBrokeringUserNameOk() (*string, bool)`

GetBrokeringUserNameOk returns a tuple with the BrokeringUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokeringUserName

`func (o *ConnectionLog) SetBrokeringUserName(v string)`

SetBrokeringUserName sets BrokeringUserName field to given value.

### HasBrokeringUserName

`func (o *ConnectionLog) HasBrokeringUserName() bool`

HasBrokeringUserName returns a boolean if a field has been set.

### GetBrokeringUserUPN

`func (o *ConnectionLog) GetBrokeringUserUPN() string`

GetBrokeringUserUPN returns the BrokeringUserUPN field if non-nil, zero value otherwise.

### GetBrokeringUserUPNOk

`func (o *ConnectionLog) GetBrokeringUserUPNOk() (*string, bool)`

GetBrokeringUserUPNOk returns a tuple with the BrokeringUserUPN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokeringUserUPN

`func (o *ConnectionLog) SetBrokeringUserUPN(v string)`

SetBrokeringUserUPN sets BrokeringUserUPN field to given value.

### HasBrokeringUserUPN

`func (o *ConnectionLog) HasBrokeringUserUPN() bool`

HasBrokeringUserUPN returns a boolean if a field has been set.

### GetConnectionFailureReason

`func (o *ConnectionLog) GetConnectionFailureReason() ConnectionLogConnectionFailureReason`

GetConnectionFailureReason returns the ConnectionFailureReason field if non-nil, zero value otherwise.

### GetConnectionFailureReasonOk

`func (o *ConnectionLog) GetConnectionFailureReasonOk() (*ConnectionLogConnectionFailureReason, bool)`

GetConnectionFailureReasonOk returns a tuple with the ConnectionFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionFailureReason

`func (o *ConnectionLog) SetConnectionFailureReason(v ConnectionLogConnectionFailureReason)`

SetConnectionFailureReason sets ConnectionFailureReason field to given value.

### HasConnectionFailureReason

`func (o *ConnectionLog) HasConnectionFailureReason() bool`

HasConnectionFailureReason returns a boolean if a field has been set.

### GetDisconnected

`func (o *ConnectionLog) GetDisconnected() bool`

GetDisconnected returns the Disconnected field if non-nil, zero value otherwise.

### GetDisconnectedOk

`func (o *ConnectionLog) GetDisconnectedOk() (*bool, bool)`

GetDisconnectedOk returns a tuple with the Disconnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnected

`func (o *ConnectionLog) SetDisconnected(v bool)`

SetDisconnected sets Disconnected field to given value.

### HasDisconnected

`func (o *ConnectionLog) HasDisconnected() bool`

HasDisconnected returns a boolean if a field has been set.

### GetEndTime

`func (o *ConnectionLog) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ConnectionLog) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ConnectionLog) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ConnectionLog) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEstablishmentTime

`func (o *ConnectionLog) GetEstablishmentTime() time.Time`

GetEstablishmentTime returns the EstablishmentTime field if non-nil, zero value otherwise.

### GetEstablishmentTimeOk

`func (o *ConnectionLog) GetEstablishmentTimeOk() (*time.Time, bool)`

GetEstablishmentTimeOk returns a tuple with the EstablishmentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstablishmentTime

`func (o *ConnectionLog) SetEstablishmentTime(v time.Time)`

SetEstablishmentTime sets EstablishmentTime field to given value.

### HasEstablishmentTime

`func (o *ConnectionLog) HasEstablishmentTime() bool`

HasEstablishmentTime returns a boolean if a field has been set.

### GetMachineDNSName

`func (o *ConnectionLog) GetMachineDNSName() string`

GetMachineDNSName returns the MachineDNSName field if non-nil, zero value otherwise.

### GetMachineDNSNameOk

`func (o *ConnectionLog) GetMachineDNSNameOk() (*string, bool)`

GetMachineDNSNameOk returns a tuple with the MachineDNSName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineDNSName

`func (o *ConnectionLog) SetMachineDNSName(v string)`

SetMachineDNSName sets MachineDNSName field to given value.

### HasMachineDNSName

`func (o *ConnectionLog) HasMachineDNSName() bool`

HasMachineDNSName returns a boolean if a field has been set.

### GetMachineName

`func (o *ConnectionLog) GetMachineName() string`

GetMachineName returns the MachineName field if non-nil, zero value otherwise.

### GetMachineNameOk

`func (o *ConnectionLog) GetMachineNameOk() (*string, bool)`

GetMachineNameOk returns a tuple with the MachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineName

`func (o *ConnectionLog) SetMachineName(v string)`

SetMachineName sets MachineName field to given value.

### HasMachineName

`func (o *ConnectionLog) HasMachineName() bool`

HasMachineName returns a boolean if a field has been set.

### GetMachineUid

`func (o *ConnectionLog) GetMachineUid() int32`

GetMachineUid returns the MachineUid field if non-nil, zero value otherwise.

### GetMachineUidOk

`func (o *ConnectionLog) GetMachineUidOk() (*int32, bool)`

GetMachineUidOk returns a tuple with the MachineUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineUid

`func (o *ConnectionLog) SetMachineUid(v int32)`

SetMachineUid sets MachineUid field to given value.

### HasMachineUid

`func (o *ConnectionLog) HasMachineUid() bool`

HasMachineUid returns a boolean if a field has been set.

### GetUid

`func (o *ConnectionLog) GetUid() int64`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ConnectionLog) GetUidOk() (*int64, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ConnectionLog) SetUid(v int64)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ConnectionLog) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


