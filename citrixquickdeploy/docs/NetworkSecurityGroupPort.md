# NetworkSecurityGroupPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortNumber** | Pointer to **NullableString** | Port Number | [optional] 
**Priority** | Pointer to **NullableInt32** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**Destination** | Pointer to **NullableString** |  | [optional] 
**Protocol** | Pointer to **NullableString** |  | [optional] 
**Action** | Pointer to **NullableString** |  | [optional] 
**Direction** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNetworkSecurityGroupPort

`func NewNetworkSecurityGroupPort() *NetworkSecurityGroupPort`

NewNetworkSecurityGroupPort instantiates a new NetworkSecurityGroupPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSecurityGroupPortWithDefaults

`func NewNetworkSecurityGroupPortWithDefaults() *NetworkSecurityGroupPort`

NewNetworkSecurityGroupPortWithDefaults instantiates a new NetworkSecurityGroupPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortNumber

`func (o *NetworkSecurityGroupPort) GetPortNumber() string`

GetPortNumber returns the PortNumber field if non-nil, zero value otherwise.

### GetPortNumberOk

`func (o *NetworkSecurityGroupPort) GetPortNumberOk() (*string, bool)`

GetPortNumberOk returns a tuple with the PortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNumber

`func (o *NetworkSecurityGroupPort) SetPortNumber(v string)`

SetPortNumber sets PortNumber field to given value.

### HasPortNumber

`func (o *NetworkSecurityGroupPort) HasPortNumber() bool`

HasPortNumber returns a boolean if a field has been set.

### SetPortNumberNil

`func (o *NetworkSecurityGroupPort) SetPortNumberNil(b bool)`

 SetPortNumberNil sets the value for PortNumber to be an explicit nil

### UnsetPortNumber
`func (o *NetworkSecurityGroupPort) UnsetPortNumber()`

UnsetPortNumber ensures that no value is present for PortNumber, not even an explicit nil
### GetPriority

`func (o *NetworkSecurityGroupPort) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NetworkSecurityGroupPort) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NetworkSecurityGroupPort) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NetworkSecurityGroupPort) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *NetworkSecurityGroupPort) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *NetworkSecurityGroupPort) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetSource

`func (o *NetworkSecurityGroupPort) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NetworkSecurityGroupPort) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NetworkSecurityGroupPort) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *NetworkSecurityGroupPort) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *NetworkSecurityGroupPort) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *NetworkSecurityGroupPort) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetDestination

`func (o *NetworkSecurityGroupPort) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *NetworkSecurityGroupPort) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *NetworkSecurityGroupPort) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *NetworkSecurityGroupPort) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### SetDestinationNil

`func (o *NetworkSecurityGroupPort) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *NetworkSecurityGroupPort) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetProtocol

`func (o *NetworkSecurityGroupPort) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NetworkSecurityGroupPort) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NetworkSecurityGroupPort) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *NetworkSecurityGroupPort) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *NetworkSecurityGroupPort) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *NetworkSecurityGroupPort) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetAction

`func (o *NetworkSecurityGroupPort) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *NetworkSecurityGroupPort) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *NetworkSecurityGroupPort) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *NetworkSecurityGroupPort) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *NetworkSecurityGroupPort) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *NetworkSecurityGroupPort) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetDirection

`func (o *NetworkSecurityGroupPort) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *NetworkSecurityGroupPort) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *NetworkSecurityGroupPort) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *NetworkSecurityGroupPort) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *NetworkSecurityGroupPort) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *NetworkSecurityGroupPort) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


