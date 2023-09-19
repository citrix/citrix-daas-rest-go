# NetworkMapResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceName** | Pointer to **NullableString** | The name for the network.  This is for display only and may not be populated. | [optional] 
**DeviceId** | **string** | Network device ID.  Zero indicates the primary network device. | 
**Network** | [**HypervisorResourceRefResponseModel**](HypervisorResourceRefResponseModel.md) |  | 
**IsCardEnabled** | Pointer to **bool** | Gets or sets a value indicating whether IsCardEnabled. | [optional] 

## Methods

### NewNetworkMapResponseModel

`func NewNetworkMapResponseModel(deviceId string, network HypervisorResourceRefResponseModel, ) *NetworkMapResponseModel`

NewNetworkMapResponseModel instantiates a new NetworkMapResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkMapResponseModelWithDefaults

`func NewNetworkMapResponseModelWithDefaults() *NetworkMapResponseModel`

NewNetworkMapResponseModelWithDefaults instantiates a new NetworkMapResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceName

`func (o *NetworkMapResponseModel) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *NetworkMapResponseModel) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *NetworkMapResponseModel) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *NetworkMapResponseModel) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceNameNil

`func (o *NetworkMapResponseModel) SetDeviceNameNil(b bool)`

 SetDeviceNameNil sets the value for DeviceName to be an explicit nil

### UnsetDeviceName
`func (o *NetworkMapResponseModel) UnsetDeviceName()`

UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
### GetDeviceId

`func (o *NetworkMapResponseModel) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *NetworkMapResponseModel) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *NetworkMapResponseModel) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetNetwork

`func (o *NetworkMapResponseModel) GetNetwork() HypervisorResourceRefResponseModel`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkMapResponseModel) GetNetworkOk() (*HypervisorResourceRefResponseModel, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkMapResponseModel) SetNetwork(v HypervisorResourceRefResponseModel)`

SetNetwork sets Network field to given value.


### GetIsCardEnabled

`func (o *NetworkMapResponseModel) GetIsCardEnabled() bool`

GetIsCardEnabled returns the IsCardEnabled field if non-nil, zero value otherwise.

### GetIsCardEnabledOk

`func (o *NetworkMapResponseModel) GetIsCardEnabledOk() (*bool, bool)`

GetIsCardEnabledOk returns a tuple with the IsCardEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCardEnabled

`func (o *NetworkMapResponseModel) SetIsCardEnabled(v bool)`

SetIsCardEnabled sets IsCardEnabled field to given value.

### HasIsCardEnabled

`func (o *NetworkMapResponseModel) HasIsCardEnabled() bool`

HasIsCardEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


