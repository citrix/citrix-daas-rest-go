# SessionResponseModelClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** | Unique identifier for the client device that has most recently been associated with the session. | [optional] 
**HardwareId** | Pointer to **string** | Unique identifier for the client hardware that has been most recently associated with the session. | [optional] 
**IPAddress** | Pointer to **string** | The IP address of the client connected to the session. | [optional] 
**Name** | Pointer to **string** | The host name of the client connected to the session. | [optional] 
**Platform** | Pointer to **string** | The name of client platform, as indicated by client product ID. | [optional] 
**ProductId** | Pointer to **int32** | The product ID of the client connected to the session. | [optional] 
**ReceiverIPAddress** | Pointer to **string** | The IP address of the client as supplied by Receiver (for example, Receiver for Web) when the session was launched, or reconnected. | [optional] 
**ReceiverName** | Pointer to **string** | The name of the client as supplied by Receiver (for example, Receiver for Web) when the session was launched, or reconnected. | [optional] 
**Version** | Pointer to **string** | The version of the Citrix Receiver running on the client connected to the session. | [optional] 

## Methods

### NewSessionResponseModelClient

`func NewSessionResponseModelClient() *SessionResponseModelClient`

NewSessionResponseModelClient instantiates a new SessionResponseModelClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionResponseModelClientWithDefaults

`func NewSessionResponseModelClientWithDefaults() *SessionResponseModelClient`

NewSessionResponseModelClientWithDefaults instantiates a new SessionResponseModelClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *SessionResponseModelClient) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SessionResponseModelClient) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SessionResponseModelClient) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *SessionResponseModelClient) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetHardwareId

`func (o *SessionResponseModelClient) GetHardwareId() string`

GetHardwareId returns the HardwareId field if non-nil, zero value otherwise.

### GetHardwareIdOk

`func (o *SessionResponseModelClient) GetHardwareIdOk() (*string, bool)`

GetHardwareIdOk returns a tuple with the HardwareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareId

`func (o *SessionResponseModelClient) SetHardwareId(v string)`

SetHardwareId sets HardwareId field to given value.

### HasHardwareId

`func (o *SessionResponseModelClient) HasHardwareId() bool`

HasHardwareId returns a boolean if a field has been set.

### GetIPAddress

`func (o *SessionResponseModelClient) GetIPAddress() string`

GetIPAddress returns the IPAddress field if non-nil, zero value otherwise.

### GetIPAddressOk

`func (o *SessionResponseModelClient) GetIPAddressOk() (*string, bool)`

GetIPAddressOk returns a tuple with the IPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPAddress

`func (o *SessionResponseModelClient) SetIPAddress(v string)`

SetIPAddress sets IPAddress field to given value.

### HasIPAddress

`func (o *SessionResponseModelClient) HasIPAddress() bool`

HasIPAddress returns a boolean if a field has been set.

### GetName

`func (o *SessionResponseModelClient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SessionResponseModelClient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SessionResponseModelClient) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SessionResponseModelClient) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *SessionResponseModelClient) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *SessionResponseModelClient) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *SessionResponseModelClient) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *SessionResponseModelClient) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetProductId

`func (o *SessionResponseModelClient) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *SessionResponseModelClient) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *SessionResponseModelClient) SetProductId(v int32)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *SessionResponseModelClient) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetReceiverIPAddress

`func (o *SessionResponseModelClient) GetReceiverIPAddress() string`

GetReceiverIPAddress returns the ReceiverIPAddress field if non-nil, zero value otherwise.

### GetReceiverIPAddressOk

`func (o *SessionResponseModelClient) GetReceiverIPAddressOk() (*string, bool)`

GetReceiverIPAddressOk returns a tuple with the ReceiverIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverIPAddress

`func (o *SessionResponseModelClient) SetReceiverIPAddress(v string)`

SetReceiverIPAddress sets ReceiverIPAddress field to given value.

### HasReceiverIPAddress

`func (o *SessionResponseModelClient) HasReceiverIPAddress() bool`

HasReceiverIPAddress returns a boolean if a field has been set.

### GetReceiverName

`func (o *SessionResponseModelClient) GetReceiverName() string`

GetReceiverName returns the ReceiverName field if non-nil, zero value otherwise.

### GetReceiverNameOk

`func (o *SessionResponseModelClient) GetReceiverNameOk() (*string, bool)`

GetReceiverNameOk returns a tuple with the ReceiverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverName

`func (o *SessionResponseModelClient) SetReceiverName(v string)`

SetReceiverName sets ReceiverName field to given value.

### HasReceiverName

`func (o *SessionResponseModelClient) HasReceiverName() bool`

HasReceiverName returns a boolean if a field has been set.

### GetVersion

`func (o *SessionResponseModelClient) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SessionResponseModelClient) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SessionResponseModelClient) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SessionResponseModelClient) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


