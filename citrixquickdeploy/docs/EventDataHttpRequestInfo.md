# EventDataHttpRequestInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientRequestId** | Pointer to **string** |  | [optional] [readonly] 
**ClientIPAddress** | Pointer to [**IPAddress**](IPAddress.md) |  | [optional] 
**Method** | Pointer to **string** |  | [optional] [readonly] 
**Uri** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewEventDataHttpRequestInfo

`func NewEventDataHttpRequestInfo() *EventDataHttpRequestInfo`

NewEventDataHttpRequestInfo instantiates a new EventDataHttpRequestInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDataHttpRequestInfoWithDefaults

`func NewEventDataHttpRequestInfoWithDefaults() *EventDataHttpRequestInfo`

NewEventDataHttpRequestInfoWithDefaults instantiates a new EventDataHttpRequestInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientRequestId

`func (o *EventDataHttpRequestInfo) GetClientRequestId() string`

GetClientRequestId returns the ClientRequestId field if non-nil, zero value otherwise.

### GetClientRequestIdOk

`func (o *EventDataHttpRequestInfo) GetClientRequestIdOk() (*string, bool)`

GetClientRequestIdOk returns a tuple with the ClientRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRequestId

`func (o *EventDataHttpRequestInfo) SetClientRequestId(v string)`

SetClientRequestId sets ClientRequestId field to given value.

### HasClientRequestId

`func (o *EventDataHttpRequestInfo) HasClientRequestId() bool`

HasClientRequestId returns a boolean if a field has been set.

### GetClientIPAddress

`func (o *EventDataHttpRequestInfo) GetClientIPAddress() IPAddress`

GetClientIPAddress returns the ClientIPAddress field if non-nil, zero value otherwise.

### GetClientIPAddressOk

`func (o *EventDataHttpRequestInfo) GetClientIPAddressOk() (*IPAddress, bool)`

GetClientIPAddressOk returns a tuple with the ClientIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIPAddress

`func (o *EventDataHttpRequestInfo) SetClientIPAddress(v IPAddress)`

SetClientIPAddress sets ClientIPAddress field to given value.

### HasClientIPAddress

`func (o *EventDataHttpRequestInfo) HasClientIPAddress() bool`

HasClientIPAddress returns a boolean if a field has been set.

### GetMethod

`func (o *EventDataHttpRequestInfo) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *EventDataHttpRequestInfo) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *EventDataHttpRequestInfo) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *EventDataHttpRequestInfo) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetUri

`func (o *EventDataHttpRequestInfo) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *EventDataHttpRequestInfo) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *EventDataHttpRequestInfo) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *EventDataHttpRequestInfo) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


