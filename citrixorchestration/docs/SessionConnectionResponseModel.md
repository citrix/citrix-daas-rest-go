# SessionConnectionResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectedViaHostName** | Pointer to **NullableString** | The host name of the incoming connection. This is usually a gateway, router or client. | [optional] 
**ConnectedViaIP** | Pointer to **NullableString** | The IP address of the incoming connection This is usually a gateway, router or client. | [optional] 
**ConnectionMode** | Pointer to [**ConnectionMode**](ConnectionMode.md) |  | [optional] 
**LaunchedViaHostName** | Pointer to **NullableString** | The host name of the StoreFront server used to launch the session. | [optional] 
**LaunchedViaIP** | Pointer to **NullableString** | The IP address of the StoreFront server used to launch the session. | [optional] 
**Protocol** | Pointer to [**ProtocolType**](ProtocolType.md) |  | [optional] 
**SecureIcaActive** | Pointer to **NullableBool** | Indicates whether SecureICA is active on the session. | [optional] 
**SmartAccessTags** | Pointer to [**[]SmartAccessTagResponseModel**](SmartAccessTagResponseModel.md) | The Smart Access tags for this session. | [optional] 
**LaunchedViaPublishedName** | Pointer to **NullableString** | The published name of the StoreFront server used to launch the session. | [optional] 

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

### SetConnectedViaHostNameNil

`func (o *SessionConnectionResponseModel) SetConnectedViaHostNameNil(b bool)`

 SetConnectedViaHostNameNil sets the value for ConnectedViaHostName to be an explicit nil

### UnsetConnectedViaHostName
`func (o *SessionConnectionResponseModel) UnsetConnectedViaHostName()`

UnsetConnectedViaHostName ensures that no value is present for ConnectedViaHostName, not even an explicit nil
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

### SetConnectedViaIPNil

`func (o *SessionConnectionResponseModel) SetConnectedViaIPNil(b bool)`

 SetConnectedViaIPNil sets the value for ConnectedViaIP to be an explicit nil

### UnsetConnectedViaIP
`func (o *SessionConnectionResponseModel) UnsetConnectedViaIP()`

UnsetConnectedViaIP ensures that no value is present for ConnectedViaIP, not even an explicit nil
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

### SetLaunchedViaHostNameNil

`func (o *SessionConnectionResponseModel) SetLaunchedViaHostNameNil(b bool)`

 SetLaunchedViaHostNameNil sets the value for LaunchedViaHostName to be an explicit nil

### UnsetLaunchedViaHostName
`func (o *SessionConnectionResponseModel) UnsetLaunchedViaHostName()`

UnsetLaunchedViaHostName ensures that no value is present for LaunchedViaHostName, not even an explicit nil
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

### SetLaunchedViaIPNil

`func (o *SessionConnectionResponseModel) SetLaunchedViaIPNil(b bool)`

 SetLaunchedViaIPNil sets the value for LaunchedViaIP to be an explicit nil

### UnsetLaunchedViaIP
`func (o *SessionConnectionResponseModel) UnsetLaunchedViaIP()`

UnsetLaunchedViaIP ensures that no value is present for LaunchedViaIP, not even an explicit nil
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

### SetSecureIcaActiveNil

`func (o *SessionConnectionResponseModel) SetSecureIcaActiveNil(b bool)`

 SetSecureIcaActiveNil sets the value for SecureIcaActive to be an explicit nil

### UnsetSecureIcaActive
`func (o *SessionConnectionResponseModel) UnsetSecureIcaActive()`

UnsetSecureIcaActive ensures that no value is present for SecureIcaActive, not even an explicit nil
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

### SetSmartAccessTagsNil

`func (o *SessionConnectionResponseModel) SetSmartAccessTagsNil(b bool)`

 SetSmartAccessTagsNil sets the value for SmartAccessTags to be an explicit nil

### UnsetSmartAccessTags
`func (o *SessionConnectionResponseModel) UnsetSmartAccessTags()`

UnsetSmartAccessTags ensures that no value is present for SmartAccessTags, not even an explicit nil
### GetLaunchedViaPublishedName

`func (o *SessionConnectionResponseModel) GetLaunchedViaPublishedName() string`

GetLaunchedViaPublishedName returns the LaunchedViaPublishedName field if non-nil, zero value otherwise.

### GetLaunchedViaPublishedNameOk

`func (o *SessionConnectionResponseModel) GetLaunchedViaPublishedNameOk() (*string, bool)`

GetLaunchedViaPublishedNameOk returns a tuple with the LaunchedViaPublishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchedViaPublishedName

`func (o *SessionConnectionResponseModel) SetLaunchedViaPublishedName(v string)`

SetLaunchedViaPublishedName sets LaunchedViaPublishedName field to given value.

### HasLaunchedViaPublishedName

`func (o *SessionConnectionResponseModel) HasLaunchedViaPublishedName() bool`

HasLaunchedViaPublishedName returns a boolean if a field has been set.

### SetLaunchedViaPublishedNameNil

`func (o *SessionConnectionResponseModel) SetLaunchedViaPublishedNameNil(b bool)`

 SetLaunchedViaPublishedNameNil sets the value for LaunchedViaPublishedName to be an explicit nil

### UnsetLaunchedViaPublishedName
`func (o *SessionConnectionResponseModel) UnsetLaunchedViaPublishedName()`

UnsetLaunchedViaPublishedName ensures that no value is present for LaunchedViaPublishedName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


