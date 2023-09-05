# NetworkMapRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceNameOrId** | Pointer to **string** | &#x60;DEPRECATED.  Use &lt;see cref&#x3D;&#39;NetworkDeviceNameOrId&#39;/&gt;.&#x60; The name or Id of the network.  Required. | [optional] 
**NetworkDeviceNameOrId** | Pointer to **string** | The name or Id of the network device. | [optional] 
**NetworkPath** | **string** | The path in the resource pool to the virtual network that the device should be attached to. This must be a path to a Network item in the resource pool to which the Machine Catalog is associated. | 

## Methods

### NewNetworkMapRequestModel

`func NewNetworkMapRequestModel(networkPath string, ) *NetworkMapRequestModel`

NewNetworkMapRequestModel instantiates a new NetworkMapRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkMapRequestModelWithDefaults

`func NewNetworkMapRequestModelWithDefaults() *NetworkMapRequestModel`

NewNetworkMapRequestModelWithDefaults instantiates a new NetworkMapRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceNameOrId

`func (o *NetworkMapRequestModel) GetDeviceNameOrId() string`

GetDeviceNameOrId returns the DeviceNameOrId field if non-nil, zero value otherwise.

### GetDeviceNameOrIdOk

`func (o *NetworkMapRequestModel) GetDeviceNameOrIdOk() (*string, bool)`

GetDeviceNameOrIdOk returns a tuple with the DeviceNameOrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceNameOrId

`func (o *NetworkMapRequestModel) SetDeviceNameOrId(v string)`

SetDeviceNameOrId sets DeviceNameOrId field to given value.

### HasDeviceNameOrId

`func (o *NetworkMapRequestModel) HasDeviceNameOrId() bool`

HasDeviceNameOrId returns a boolean if a field has been set.

### GetNetworkDeviceNameOrId

`func (o *NetworkMapRequestModel) GetNetworkDeviceNameOrId() string`

GetNetworkDeviceNameOrId returns the NetworkDeviceNameOrId field if non-nil, zero value otherwise.

### GetNetworkDeviceNameOrIdOk

`func (o *NetworkMapRequestModel) GetNetworkDeviceNameOrIdOk() (*string, bool)`

GetNetworkDeviceNameOrIdOk returns a tuple with the NetworkDeviceNameOrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceNameOrId

`func (o *NetworkMapRequestModel) SetNetworkDeviceNameOrId(v string)`

SetNetworkDeviceNameOrId sets NetworkDeviceNameOrId field to given value.

### HasNetworkDeviceNameOrId

`func (o *NetworkMapRequestModel) HasNetworkDeviceNameOrId() bool`

HasNetworkDeviceNameOrId returns a boolean if a field has been set.

### GetNetworkPath

`func (o *NetworkMapRequestModel) GetNetworkPath() string`

GetNetworkPath returns the NetworkPath field if non-nil, zero value otherwise.

### GetNetworkPathOk

`func (o *NetworkMapRequestModel) GetNetworkPathOk() (*string, bool)`

GetNetworkPathOk returns a tuple with the NetworkPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPath

`func (o *NetworkMapRequestModel) SetNetworkPath(v string)`

SetNetworkPath sets NetworkPath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


