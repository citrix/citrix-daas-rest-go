# ConnectionLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrokeringTime** | Pointer to **NullableTime** |  | [optional] 
**BrokeringUserName** | Pointer to **NullableString** |  | [optional] 
**BrokeringUserUPN** | Pointer to **NullableString** |  | [optional] 
**ConnectionFailureReason** | Pointer to [**NullableConnectionLogConnectionFailureReason**](ConnectionLogConnectionFailureReason.md) |  | [optional] 
**Disconnected** | Pointer to **NullableBool** |  | [optional] 
**EndTime** | Pointer to **NullableTime** |  | [optional] 
**EstablishmentTime** | Pointer to **NullableTime** |  | [optional] 
**MachineDNSName** | Pointer to **NullableString** |  | [optional] 
**MachineName** | Pointer to **NullableString** |  | [optional] 
**MachineUid** | Pointer to **NullableInt32** |  | [optional] 
**Uid** | Pointer to **NullableInt64** |  | [optional] 

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

### SetBrokeringTimeNil

`func (o *ConnectionLog) SetBrokeringTimeNil(b bool)`

 SetBrokeringTimeNil sets the value for BrokeringTime to be an explicit nil

### UnsetBrokeringTime
`func (o *ConnectionLog) UnsetBrokeringTime()`

UnsetBrokeringTime ensures that no value is present for BrokeringTime, not even an explicit nil
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

### SetBrokeringUserNameNil

`func (o *ConnectionLog) SetBrokeringUserNameNil(b bool)`

 SetBrokeringUserNameNil sets the value for BrokeringUserName to be an explicit nil

### UnsetBrokeringUserName
`func (o *ConnectionLog) UnsetBrokeringUserName()`

UnsetBrokeringUserName ensures that no value is present for BrokeringUserName, not even an explicit nil
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

### SetBrokeringUserUPNNil

`func (o *ConnectionLog) SetBrokeringUserUPNNil(b bool)`

 SetBrokeringUserUPNNil sets the value for BrokeringUserUPN to be an explicit nil

### UnsetBrokeringUserUPN
`func (o *ConnectionLog) UnsetBrokeringUserUPN()`

UnsetBrokeringUserUPN ensures that no value is present for BrokeringUserUPN, not even an explicit nil
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

### SetConnectionFailureReasonNil

`func (o *ConnectionLog) SetConnectionFailureReasonNil(b bool)`

 SetConnectionFailureReasonNil sets the value for ConnectionFailureReason to be an explicit nil

### UnsetConnectionFailureReason
`func (o *ConnectionLog) UnsetConnectionFailureReason()`

UnsetConnectionFailureReason ensures that no value is present for ConnectionFailureReason, not even an explicit nil
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

### SetDisconnectedNil

`func (o *ConnectionLog) SetDisconnectedNil(b bool)`

 SetDisconnectedNil sets the value for Disconnected to be an explicit nil

### UnsetDisconnected
`func (o *ConnectionLog) UnsetDisconnected()`

UnsetDisconnected ensures that no value is present for Disconnected, not even an explicit nil
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

### SetEndTimeNil

`func (o *ConnectionLog) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *ConnectionLog) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
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

### SetEstablishmentTimeNil

`func (o *ConnectionLog) SetEstablishmentTimeNil(b bool)`

 SetEstablishmentTimeNil sets the value for EstablishmentTime to be an explicit nil

### UnsetEstablishmentTime
`func (o *ConnectionLog) UnsetEstablishmentTime()`

UnsetEstablishmentTime ensures that no value is present for EstablishmentTime, not even an explicit nil
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

### SetMachineDNSNameNil

`func (o *ConnectionLog) SetMachineDNSNameNil(b bool)`

 SetMachineDNSNameNil sets the value for MachineDNSName to be an explicit nil

### UnsetMachineDNSName
`func (o *ConnectionLog) UnsetMachineDNSName()`

UnsetMachineDNSName ensures that no value is present for MachineDNSName, not even an explicit nil
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

### SetMachineNameNil

`func (o *ConnectionLog) SetMachineNameNil(b bool)`

 SetMachineNameNil sets the value for MachineName to be an explicit nil

### UnsetMachineName
`func (o *ConnectionLog) UnsetMachineName()`

UnsetMachineName ensures that no value is present for MachineName, not even an explicit nil
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

### SetMachineUidNil

`func (o *ConnectionLog) SetMachineUidNil(b bool)`

 SetMachineUidNil sets the value for MachineUid to be an explicit nil

### UnsetMachineUid
`func (o *ConnectionLog) UnsetMachineUid()`

UnsetMachineUid ensures that no value is present for MachineUid, not even an explicit nil
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

### SetUidNil

`func (o *ConnectionLog) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *ConnectionLog) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


