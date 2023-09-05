# SessionConnectionResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectedViaHostName** | Pointer to **string** | The host name of the incoming connection. This is usually a gateway, router or client. | [optional] 
**ConnectedViaIP** | Pointer to **string** | The IP address of the incoming connection This is usually a gateway, router or client. | [optional] 
**ConnectionMode** | Pointer to [**ConnectionMode**](ConnectionMode.md) |  | [optional] 
**LaunchedViaHostName** | Pointer to **string** | The host name of the StoreFront server used to launch the session. | [optional] 
**LaunchedViaIP** | Pointer to **string** | The IP address of the StoreFront server used to launch the session. | [optional] 
**Protocol** | Pointer to [**ProtocolType**](ProtocolType.md) |  | [optional] 
**SecureIcaActive** | Pointer to **bool** | Indicates whether SecureICA is active on the session. | [optional] 
**SmartAccessTags** | Pointer to [**[]SmartAccessTagResponseModel**](SmartAccessTagResponseModel.md) | The Smart Access tags for this session. | [optional] 

## Methods

### NewSessionConnectionResponseModel

`func NewSessionConnectionResponseModel() *SessionConnectionResponseModel`

NewSessionConnectionResponseModel instantiates a new SessionConnectionResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionConnectionResponseModelWithDefaults

`func NewSessionConnectionResponseModelWithDefaults() *SessionConnectionResponseModel`

NewSessionConnectionResponseModelWithDefaults instantiates a new SessionConnectionResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectedViaHostName

`func (o *SessionConnectionResponseModel) GetConnectedViaHostName() string`

GetConnectedViaHostName returns the ConnectedViaHostName field if non-nil, zero value otherwise.

### GetConnectedViaHostNameOk

`func (o *SessionConnectionResponseModel) GetConnectedViaHostNameOk() (*string, bool)`

GetConnectedViaHostNameOk returns a tuple with the ConnectedViaHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedViaHostName

`func (o *SessionConnectionResponseModel) SetConnectedViaHostName(v string)`

SetConnectedViaHostName sets ConnectedViaHostName field to given value.

### HasConnectedViaHostName

`func (o *SessionConnectionResponseModel) HasConnectedViaHostName() bool`

HasConnectedViaHostName returns a boolean if a field has been set.

### GetConnectedViaIP

`func (o *SessionConnectionResponseModel) GetConnectedViaIP() string`

GetConnectedViaIP returns the ConnectedViaIP field if non-nil, zero value otherwise.

### GetConnectedViaIPOk

`func (o *SessionConnectionResponseModel) GetConnectedViaIPOk() (*string, bool)`

GetConnectedViaIPOk returns a tuple with the ConnectedViaIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedViaIP

`func (o *SessionConnectionResponseModel) SetConnectedViaIP(v string)`

SetConnectedViaIP sets ConnectedViaIP field to given value.

### HasConnectedViaIP

`func (o *SessionConnectionResponseModel) HasConnectedViaIP() bool`

HasConnectedViaIP returns a boolean if a field has been set.

### GetConnectionMode

`func (o *SessionConnectionResponseModel) GetConnectionMode() ConnectionMode`

GetConnectionMode returns the ConnectionMode field if non-nil, zero value otherwise.

### GetConnectionModeOk

`func (o *SessionConnectionResponseModel) GetConnectionModeOk() (*ConnectionMode, bool)`

GetConnectionModeOk returns a tuple with the ConnectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionMode

`func (o *SessionConnectionResponseModel) SetConnectionMode(v ConnectionMode)`

SetConnectionMode sets ConnectionMode field to given value.

### HasConnectionMode

`func (o *SessionConnectionResponseModel) HasConnectionMode() bool`

HasConnectionMode returns a boolean if a field has been set.

### GetLaunchedViaHostName

`func (o *SessionConnectionResponseModel) GetLaunchedViaHostName() string`

GetLaunchedViaHostName returns the LaunchedViaHostName field if non-nil, zero value otherwise.

### GetLaunchedViaHostNameOk

`func (o *SessionConnectionResponseModel) GetLaunchedViaHostNameOk() (*string, bool)`

GetLaunchedViaHostNameOk returns a tuple with the LaunchedViaHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchedViaHostName

`func (o *SessionConnectionResponseModel) SetLaunchedViaHostName(v string)`

SetLaunchedViaHostName sets LaunchedViaHostName field to given value.

### HasLaunchedViaHostName

`func (o *SessionConnectionResponseModel) HasLaunchedViaHostName() bool`

HasLaunchedViaHostName returns a boolean if a field has been set.

### GetLaunchedViaIP

`func (o *SessionConnectionResponseModel) GetLaunchedViaIP() string`

GetLaunchedViaIP returns the LaunchedViaIP field if non-nil, zero value otherwise.

### GetLaunchedViaIPOk

`func (o *SessionConnectionResponseModel) GetLaunchedViaIPOk() (*string, bool)`

GetLaunchedViaIPOk returns a tuple with the LaunchedViaIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchedViaIP

`func (o *SessionConnectionResponseModel) SetLaunchedViaIP(v string)`

SetLaunchedViaIP sets LaunchedViaIP field to given value.

### HasLaunchedViaIP

`func (o *SessionConnectionResponseModel) HasLaunchedViaIP() bool`

HasLaunchedViaIP returns a boolean if a field has been set.

### GetProtocol

`func (o *SessionConnectionResponseModel) GetProtocol() ProtocolType`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SessionConnectionResponseModel) GetProtocolOk() (*ProtocolType, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SessionConnectionResponseModel) SetProtocol(v ProtocolType)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SessionConnectionResponseModel) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSecureIcaActive

`func (o *SessionConnectionResponseModel) GetSecureIcaActive() bool`

GetSecureIcaActive returns the SecureIcaActive field if non-nil, zero value otherwise.

### GetSecureIcaActiveOk

`func (o *SessionConnectionResponseModel) GetSecureIcaActiveOk() (*bool, bool)`

GetSecureIcaActiveOk returns a tuple with the SecureIcaActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureIcaActive

`func (o *SessionConnectionResponseModel) SetSecureIcaActive(v bool)`

SetSecureIcaActive sets SecureIcaActive field to given value.

### HasSecureIcaActive

`func (o *SessionConnectionResponseModel) HasSecureIcaActive() bool`

HasSecureIcaActive returns a boolean if a field has been set.

### GetSmartAccessTags

`func (o *SessionConnectionResponseModel) GetSmartAccessTags() []SmartAccessTagResponseModel`

GetSmartAccessTags returns the SmartAccessTags field if non-nil, zero value otherwise.

### GetSmartAccessTagsOk

`func (o *SessionConnectionResponseModel) GetSmartAccessTagsOk() (*[]SmartAccessTagResponseModel, bool)`

GetSmartAccessTagsOk returns a tuple with the SmartAccessTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartAccessTags

`func (o *SessionConnectionResponseModel) SetSmartAccessTags(v []SmartAccessTagResponseModel)`

SetSmartAccessTags sets SmartAccessTags field to given value.

### HasSmartAccessTags

`func (o *SessionConnectionResponseModel) HasSmartAccessTags() bool`

HasSmartAccessTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


