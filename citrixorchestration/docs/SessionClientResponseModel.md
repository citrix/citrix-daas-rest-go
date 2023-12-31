# SessionClientResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **NullableString** | Unique identifier for the client device that has most recently been associated with the session. | [optional] 
**HardwareId** | Pointer to **NullableString** | Unique identifier for the client hardware that has been most recently associated with the session. | [optional] 
**IPAddress** | Pointer to **NullableString** | The IP address of the client connected to the session. | [optional] 
**Name** | Pointer to **NullableString** | The host name of the client connected to the session. | [optional] 
**Platform** | Pointer to **NullableString** | The name of client platform, as indicated by client product ID. | [optional] 
**ProductId** | Pointer to **NullableInt32** | The product ID of the client connected to the session. | [optional] 
**ReceiverIPAddress** | Pointer to **NullableString** | The IP address of the client as supplied by Receiver (for example, Receiver for Web) when the session was launched, or reconnected. | [optional] 
**ReceiverName** | Pointer to **NullableString** | The name of the client as supplied by Receiver (for example, Receiver for Web) when the session was launched, or reconnected. | [optional] 
**Version** | Pointer to **NullableString** | The version of the Citrix Receiver running on the client connected to the session. | [optional] 

## Methods

### NewSessionClientResponseModel

`func NewSessionClientResponseModel() *SessionClientResponseModel`

NewSessionClientResponseModel instantiates a new SessionClientResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionClientResponseModelWithDefaults

`func NewSessionClientResponseModelWithDefaults() *SessionClientResponseModel`

NewSessionClientResponseModelWithDefaults instantiates a new SessionClientResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *SessionClientResponseModel) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SessionClientResponseModel) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SessionClientResponseModel) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *SessionClientResponseModel) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *SessionClientResponseModel) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *SessionClientResponseModel) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetHardwareId

`func (o *SessionClientResponseModel) GetHardwareId() string`

GetHardwareId returns the HardwareId field if non-nil, zero value otherwise.

### GetHardwareIdOk

`func (o *SessionClientResponseModel) GetHardwareIdOk() (*string, bool)`

GetHardwareIdOk returns a tuple with the HardwareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareId

`func (o *SessionClientResponseModel) SetHardwareId(v string)`

SetHardwareId sets HardwareId field to given value.

### HasHardwareId

`func (o *SessionClientResponseModel) HasHardwareId() bool`

HasHardwareId returns a boolean if a field has been set.

### SetHardwareIdNil

`func (o *SessionClientResponseModel) SetHardwareIdNil(b bool)`

 SetHardwareIdNil sets the value for HardwareId to be an explicit nil

### UnsetHardwareId
`func (o *SessionClientResponseModel) UnsetHardwareId()`

UnsetHardwareId ensures that no value is present for HardwareId, not even an explicit nil
### GetIPAddress

`func (o *SessionClientResponseModel) GetIPAddress() string`

GetIPAddress returns the IPAddress field if non-nil, zero value otherwise.

### GetIPAddressOk

`func (o *SessionClientResponseModel) GetIPAddressOk() (*string, bool)`

GetIPAddressOk returns a tuple with the IPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPAddress

`func (o *SessionClientResponseModel) SetIPAddress(v string)`

SetIPAddress sets IPAddress field to given value.

### HasIPAddress

`func (o *SessionClientResponseModel) HasIPAddress() bool`

HasIPAddress returns a boolean if a field has been set.

### SetIPAddressNil

`func (o *SessionClientResponseModel) SetIPAddressNil(b bool)`

 SetIPAddressNil sets the value for IPAddress to be an explicit nil

### UnsetIPAddress
`func (o *SessionClientResponseModel) UnsetIPAddress()`

UnsetIPAddress ensures that no value is present for IPAddress, not even an explicit nil
### GetName

`func (o *SessionClientResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SessionClientResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SessionClientResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SessionClientResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SessionClientResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SessionClientResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPlatform

`func (o *SessionClientResponseModel) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *SessionClientResponseModel) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *SessionClientResponseModel) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *SessionClientResponseModel) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *SessionClientResponseModel) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *SessionClientResponseModel) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetProductId

`func (o *SessionClientResponseModel) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *SessionClientResponseModel) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *SessionClientResponseModel) SetProductId(v int32)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *SessionClientResponseModel) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *SessionClientResponseModel) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *SessionClientResponseModel) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetReceiverIPAddress

`func (o *SessionClientResponseModel) GetReceiverIPAddress() string`

GetReceiverIPAddress returns the ReceiverIPAddress field if non-nil, zero value otherwise.

### GetReceiverIPAddressOk

`func (o *SessionClientResponseModel) GetReceiverIPAddressOk() (*string, bool)`

GetReceiverIPAddressOk returns a tuple with the ReceiverIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverIPAddress

`func (o *SessionClientResponseModel) SetReceiverIPAddress(v string)`

SetReceiverIPAddress sets ReceiverIPAddress field to given value.

### HasReceiverIPAddress

`func (o *SessionClientResponseModel) HasReceiverIPAddress() bool`

HasReceiverIPAddress returns a boolean if a field has been set.

### SetReceiverIPAddressNil

`func (o *SessionClientResponseModel) SetReceiverIPAddressNil(b bool)`

 SetReceiverIPAddressNil sets the value for ReceiverIPAddress to be an explicit nil

### UnsetReceiverIPAddress
`func (o *SessionClientResponseModel) UnsetReceiverIPAddress()`

UnsetReceiverIPAddress ensures that no value is present for ReceiverIPAddress, not even an explicit nil
### GetReceiverName

`func (o *SessionClientResponseModel) GetReceiverName() string`

GetReceiverName returns the ReceiverName field if non-nil, zero value otherwise.

### GetReceiverNameOk

`func (o *SessionClientResponseModel) GetReceiverNameOk() (*string, bool)`

GetReceiverNameOk returns a tuple with the ReceiverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverName

`func (o *SessionClientResponseModel) SetReceiverName(v string)`

SetReceiverName sets ReceiverName field to given value.

### HasReceiverName

`func (o *SessionClientResponseModel) HasReceiverName() bool`

HasReceiverName returns a boolean if a field has been set.

### SetReceiverNameNil

`func (o *SessionClientResponseModel) SetReceiverNameNil(b bool)`

 SetReceiverNameNil sets the value for ReceiverName to be an explicit nil

### UnsetReceiverName
`func (o *SessionClientResponseModel) UnsetReceiverName()`

UnsetReceiverName ensures that no value is present for ReceiverName, not even an explicit nil
### GetVersion

`func (o *SessionClientResponseModel) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SessionClientResponseModel) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SessionClientResponseModel) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SessionClientResponseModel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *SessionClientResponseModel) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *SessionClientResponseModel) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


