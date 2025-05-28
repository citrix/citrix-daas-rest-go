# AlertContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorization** | Pointer to [**Authorization**](Authorization.md) |  | [optional] 
**Channels** | Pointer to **string** |  | [optional] 
**Claims** | Pointer to **string** |  | [optional] 
**Caller** | Pointer to **string** |  | [optional] 
**CorrelationId** | Pointer to **string** |  | [optional] 
**EventSource** | Pointer to **string** |  | [optional] 
**EventTimestamp** | Pointer to **time.Time** |  | [optional] 
**HttpRequest** | Pointer to **string** |  | [optional] 
**EventDataId** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**OperationName** | Pointer to **string** |  | [optional] 
**OperationId** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to [**Properties**](Properties.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SubStatus** | Pointer to **string** |  | [optional] 
**SubmissionTimestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAlertContext

`func NewAlertContext() *AlertContext`

NewAlertContext instantiates a new AlertContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertContextWithDefaults

`func NewAlertContextWithDefaults() *AlertContext`

NewAlertContextWithDefaults instantiates a new AlertContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorization

`func (o *AlertContext) GetAuthorization() Authorization`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *AlertContext) GetAuthorizationOk() (*Authorization, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *AlertContext) SetAuthorization(v Authorization)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *AlertContext) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.

### GetChannels

`func (o *AlertContext) GetChannels() string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *AlertContext) GetChannelsOk() (*string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *AlertContext) SetChannels(v string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *AlertContext) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetClaims

`func (o *AlertContext) GetClaims() string`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *AlertContext) GetClaimsOk() (*string, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *AlertContext) SetClaims(v string)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *AlertContext) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetCaller

`func (o *AlertContext) GetCaller() string`

GetCaller returns the Caller field if non-nil, zero value otherwise.

### GetCallerOk

`func (o *AlertContext) GetCallerOk() (*string, bool)`

GetCallerOk returns a tuple with the Caller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaller

`func (o *AlertContext) SetCaller(v string)`

SetCaller sets Caller field to given value.

### HasCaller

`func (o *AlertContext) HasCaller() bool`

HasCaller returns a boolean if a field has been set.

### GetCorrelationId

`func (o *AlertContext) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *AlertContext) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *AlertContext) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *AlertContext) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetEventSource

`func (o *AlertContext) GetEventSource() string`

GetEventSource returns the EventSource field if non-nil, zero value otherwise.

### GetEventSourceOk

`func (o *AlertContext) GetEventSourceOk() (*string, bool)`

GetEventSourceOk returns a tuple with the EventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSource

`func (o *AlertContext) SetEventSource(v string)`

SetEventSource sets EventSource field to given value.

### HasEventSource

`func (o *AlertContext) HasEventSource() bool`

HasEventSource returns a boolean if a field has been set.

### GetEventTimestamp

`func (o *AlertContext) GetEventTimestamp() time.Time`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *AlertContext) GetEventTimestampOk() (*time.Time, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *AlertContext) SetEventTimestamp(v time.Time)`

SetEventTimestamp sets EventTimestamp field to given value.

### HasEventTimestamp

`func (o *AlertContext) HasEventTimestamp() bool`

HasEventTimestamp returns a boolean if a field has been set.

### GetHttpRequest

`func (o *AlertContext) GetHttpRequest() string`

GetHttpRequest returns the HttpRequest field if non-nil, zero value otherwise.

### GetHttpRequestOk

`func (o *AlertContext) GetHttpRequestOk() (*string, bool)`

GetHttpRequestOk returns a tuple with the HttpRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRequest

`func (o *AlertContext) SetHttpRequest(v string)`

SetHttpRequest sets HttpRequest field to given value.

### HasHttpRequest

`func (o *AlertContext) HasHttpRequest() bool`

HasHttpRequest returns a boolean if a field has been set.

### GetEventDataId

`func (o *AlertContext) GetEventDataId() string`

GetEventDataId returns the EventDataId field if non-nil, zero value otherwise.

### GetEventDataIdOk

`func (o *AlertContext) GetEventDataIdOk() (*string, bool)`

GetEventDataIdOk returns a tuple with the EventDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDataId

`func (o *AlertContext) SetEventDataId(v string)`

SetEventDataId sets EventDataId field to given value.

### HasEventDataId

`func (o *AlertContext) HasEventDataId() bool`

HasEventDataId returns a boolean if a field has been set.

### GetLevel

`func (o *AlertContext) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AlertContext) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AlertContext) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *AlertContext) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetOperationName

`func (o *AlertContext) GetOperationName() string`

GetOperationName returns the OperationName field if non-nil, zero value otherwise.

### GetOperationNameOk

`func (o *AlertContext) GetOperationNameOk() (*string, bool)`

GetOperationNameOk returns a tuple with the OperationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationName

`func (o *AlertContext) SetOperationName(v string)`

SetOperationName sets OperationName field to given value.

### HasOperationName

`func (o *AlertContext) HasOperationName() bool`

HasOperationName returns a boolean if a field has been set.

### GetOperationId

`func (o *AlertContext) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *AlertContext) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *AlertContext) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.

### HasOperationId

`func (o *AlertContext) HasOperationId() bool`

HasOperationId returns a boolean if a field has been set.

### GetProperties

`func (o *AlertContext) GetProperties() Properties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AlertContext) GetPropertiesOk() (*Properties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AlertContext) SetProperties(v Properties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *AlertContext) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetStatus

`func (o *AlertContext) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AlertContext) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AlertContext) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AlertContext) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubStatus

`func (o *AlertContext) GetSubStatus() string`

GetSubStatus returns the SubStatus field if non-nil, zero value otherwise.

### GetSubStatusOk

`func (o *AlertContext) GetSubStatusOk() (*string, bool)`

GetSubStatusOk returns a tuple with the SubStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStatus

`func (o *AlertContext) SetSubStatus(v string)`

SetSubStatus sets SubStatus field to given value.

### HasSubStatus

`func (o *AlertContext) HasSubStatus() bool`

HasSubStatus returns a boolean if a field has been set.

### GetSubmissionTimestamp

`func (o *AlertContext) GetSubmissionTimestamp() time.Time`

GetSubmissionTimestamp returns the SubmissionTimestamp field if non-nil, zero value otherwise.

### GetSubmissionTimestampOk

`func (o *AlertContext) GetSubmissionTimestampOk() (*time.Time, bool)`

GetSubmissionTimestampOk returns a tuple with the SubmissionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionTimestamp

`func (o *AlertContext) SetSubmissionTimestamp(v time.Time)`

SetSubmissionTimestamp sets SubmissionTimestamp field to given value.

### HasSubmissionTimestamp

`func (o *AlertContext) HasSubmissionTimestamp() bool`

HasSubmissionTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


