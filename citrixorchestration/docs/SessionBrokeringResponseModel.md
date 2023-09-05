# SessionBrokeringResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutonomouslyBrokered** | **bool** | Indicates whether this is an HDX session established by direct connection without being brokered. | 
**DurationMilliseconds** | Pointer to **int32** | Time taken to broker the session. | [optional] 
**Time** | Pointer to **string** | Time at which the session was brokered. | [optional] 
**FormattedTime** | Pointer to **string** | Formatted time at which the session was brokered. RFC 3339 compatible format. | [optional] 
**UserName** | Pointer to **string** | The user name of the brokering user. | [optional] 
**UserSid** | Pointer to **string** | The SID of the brokering user. | [optional] 

## Methods

### NewSessionBrokeringResponseModel

`func NewSessionBrokeringResponseModel(autonomouslyBrokered bool, ) *SessionBrokeringResponseModel`

NewSessionBrokeringResponseModel instantiates a new SessionBrokeringResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionBrokeringResponseModelWithDefaults

`func NewSessionBrokeringResponseModelWithDefaults() *SessionBrokeringResponseModel`

NewSessionBrokeringResponseModelWithDefaults instantiates a new SessionBrokeringResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutonomouslyBrokered

`func (o *SessionBrokeringResponseModel) GetAutonomouslyBrokered() bool`

GetAutonomouslyBrokered returns the AutonomouslyBrokered field if non-nil, zero value otherwise.

### GetAutonomouslyBrokeredOk

`func (o *SessionBrokeringResponseModel) GetAutonomouslyBrokeredOk() (*bool, bool)`

GetAutonomouslyBrokeredOk returns a tuple with the AutonomouslyBrokered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutonomouslyBrokered

`func (o *SessionBrokeringResponseModel) SetAutonomouslyBrokered(v bool)`

SetAutonomouslyBrokered sets AutonomouslyBrokered field to given value.


### GetDurationMilliseconds

`func (o *SessionBrokeringResponseModel) GetDurationMilliseconds() int32`

GetDurationMilliseconds returns the DurationMilliseconds field if non-nil, zero value otherwise.

### GetDurationMillisecondsOk

`func (o *SessionBrokeringResponseModel) GetDurationMillisecondsOk() (*int32, bool)`

GetDurationMillisecondsOk returns a tuple with the DurationMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMilliseconds

`func (o *SessionBrokeringResponseModel) SetDurationMilliseconds(v int32)`

SetDurationMilliseconds sets DurationMilliseconds field to given value.

### HasDurationMilliseconds

`func (o *SessionBrokeringResponseModel) HasDurationMilliseconds() bool`

HasDurationMilliseconds returns a boolean if a field has been set.

### GetTime

`func (o *SessionBrokeringResponseModel) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *SessionBrokeringResponseModel) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *SessionBrokeringResponseModel) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *SessionBrokeringResponseModel) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetFormattedTime

`func (o *SessionBrokeringResponseModel) GetFormattedTime() string`

GetFormattedTime returns the FormattedTime field if non-nil, zero value otherwise.

### GetFormattedTimeOk

`func (o *SessionBrokeringResponseModel) GetFormattedTimeOk() (*string, bool)`

GetFormattedTimeOk returns a tuple with the FormattedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTime

`func (o *SessionBrokeringResponseModel) SetFormattedTime(v string)`

SetFormattedTime sets FormattedTime field to given value.

### HasFormattedTime

`func (o *SessionBrokeringResponseModel) HasFormattedTime() bool`

HasFormattedTime returns a boolean if a field has been set.

### GetUserName

`func (o *SessionBrokeringResponseModel) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SessionBrokeringResponseModel) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SessionBrokeringResponseModel) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SessionBrokeringResponseModel) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUserSid

`func (o *SessionBrokeringResponseModel) GetUserSid() string`

GetUserSid returns the UserSid field if non-nil, zero value otherwise.

### GetUserSidOk

`func (o *SessionBrokeringResponseModel) GetUserSidOk() (*string, bool)`

GetUserSidOk returns a tuple with the UserSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSid

`func (o *SessionBrokeringResponseModel) SetUserSid(v string)`

SetUserSid sets UserSid field to given value.

### HasUserSid

`func (o *SessionBrokeringResponseModel) HasUserSid() bool`

HasUserSid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


