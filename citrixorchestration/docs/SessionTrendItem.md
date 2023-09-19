# SessionTrendItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **NullableString** | The date and time for this SessionTrendItem. | [optional] 
**ConcurrentSessions** | Pointer to **NullableInt32** | ActiveSessions | [optional] 
**ConnectedSessions** | Pointer to **NullableInt32** | ConnectedSessions | [optional] 
**DisconnectedSessions** | Pointer to **NullableInt32** | DisconnectedSessions | [optional] 

## Methods

### NewSessionTrendItem

`func NewSessionTrendItem() *SessionTrendItem`

NewSessionTrendItem instantiates a new SessionTrendItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionTrendItemWithDefaults

`func NewSessionTrendItemWithDefaults() *SessionTrendItem`

NewSessionTrendItemWithDefaults instantiates a new SessionTrendItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *SessionTrendItem) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *SessionTrendItem) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *SessionTrendItem) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *SessionTrendItem) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *SessionTrendItem) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *SessionTrendItem) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetConcurrentSessions

`func (o *SessionTrendItem) GetConcurrentSessions() int32`

GetConcurrentSessions returns the ConcurrentSessions field if non-nil, zero value otherwise.

### GetConcurrentSessionsOk

`func (o *SessionTrendItem) GetConcurrentSessionsOk() (*int32, bool)`

GetConcurrentSessionsOk returns a tuple with the ConcurrentSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentSessions

`func (o *SessionTrendItem) SetConcurrentSessions(v int32)`

SetConcurrentSessions sets ConcurrentSessions field to given value.

### HasConcurrentSessions

`func (o *SessionTrendItem) HasConcurrentSessions() bool`

HasConcurrentSessions returns a boolean if a field has been set.

### SetConcurrentSessionsNil

`func (o *SessionTrendItem) SetConcurrentSessionsNil(b bool)`

 SetConcurrentSessionsNil sets the value for ConcurrentSessions to be an explicit nil

### UnsetConcurrentSessions
`func (o *SessionTrendItem) UnsetConcurrentSessions()`

UnsetConcurrentSessions ensures that no value is present for ConcurrentSessions, not even an explicit nil
### GetConnectedSessions

`func (o *SessionTrendItem) GetConnectedSessions() int32`

GetConnectedSessions returns the ConnectedSessions field if non-nil, zero value otherwise.

### GetConnectedSessionsOk

`func (o *SessionTrendItem) GetConnectedSessionsOk() (*int32, bool)`

GetConnectedSessionsOk returns a tuple with the ConnectedSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedSessions

`func (o *SessionTrendItem) SetConnectedSessions(v int32)`

SetConnectedSessions sets ConnectedSessions field to given value.

### HasConnectedSessions

`func (o *SessionTrendItem) HasConnectedSessions() bool`

HasConnectedSessions returns a boolean if a field has been set.

### SetConnectedSessionsNil

`func (o *SessionTrendItem) SetConnectedSessionsNil(b bool)`

 SetConnectedSessionsNil sets the value for ConnectedSessions to be an explicit nil

### UnsetConnectedSessions
`func (o *SessionTrendItem) UnsetConnectedSessions()`

UnsetConnectedSessions ensures that no value is present for ConnectedSessions, not even an explicit nil
### GetDisconnectedSessions

`func (o *SessionTrendItem) GetDisconnectedSessions() int32`

GetDisconnectedSessions returns the DisconnectedSessions field if non-nil, zero value otherwise.

### GetDisconnectedSessionsOk

`func (o *SessionTrendItem) GetDisconnectedSessionsOk() (*int32, bool)`

GetDisconnectedSessionsOk returns a tuple with the DisconnectedSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedSessions

`func (o *SessionTrendItem) SetDisconnectedSessions(v int32)`

SetDisconnectedSessions sets DisconnectedSessions field to given value.

### HasDisconnectedSessions

`func (o *SessionTrendItem) HasDisconnectedSessions() bool`

HasDisconnectedSessions returns a boolean if a field has been set.

### SetDisconnectedSessionsNil

`func (o *SessionTrendItem) SetDisconnectedSessionsNil(b bool)`

 SetDisconnectedSessionsNil sets the value for DisconnectedSessions to be an explicit nil

### UnsetDisconnectedSessions
`func (o *SessionTrendItem) UnsetDisconnectedSessions()`

UnsetDisconnectedSessions ensures that no value is present for DisconnectedSessions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


