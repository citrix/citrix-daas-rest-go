# PvsDeviceResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Machine name. | 
**Sid** | **string** | Machine&#39;s Active Directory SID. | 
**PvsSite** | [**RefResponseModel**](RefResponseModel.md) |  | 
**PvsCollection** | [**RefResponseModel**](RefResponseModel.md) |  | 
**HypervisorConnection** | Pointer to **NullableString** | Hypervisor connection associated with the machine.  Will be set if the caller requested &#x60;resolveHypervisor&#x3D;true&#x60; and the machine is located on one of the hypervisors connected to the site. | [optional] 
**HostedMachineId** | Pointer to **NullableString** | The unique ID by which the hypervisor recognizes the machine.  Will be set if the caller requested &#x60;resolveHypervisor&#x3D;true&#x60; and the machine is located on one of the hypervisors connected to the site. | [optional] 

## Methods

### NewPvsDeviceResponseModel

`func NewPvsDeviceResponseModel(name string, sid string, pvsSite RefResponseModel, pvsCollection RefResponseModel, ) *PvsDeviceResponseModel`

NewPvsDeviceResponseModel instantiates a new PvsDeviceResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPvsDeviceResponseModelWithDefaults

`func NewPvsDeviceResponseModelWithDefaults() *PvsDeviceResponseModel`

NewPvsDeviceResponseModelWithDefaults instantiates a new PvsDeviceResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PvsDeviceResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PvsDeviceResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PvsDeviceResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetSid

`func (o *PvsDeviceResponseModel) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *PvsDeviceResponseModel) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *PvsDeviceResponseModel) SetSid(v string)`

SetSid sets Sid field to given value.


### GetPvsSite

`func (o *PvsDeviceResponseModel) GetPvsSite() RefResponseModel`

GetPvsSite returns the PvsSite field if non-nil, zero value otherwise.

### GetPvsSiteOk

`func (o *PvsDeviceResponseModel) GetPvsSiteOk() (*RefResponseModel, bool)`

GetPvsSiteOk returns a tuple with the PvsSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsSite

`func (o *PvsDeviceResponseModel) SetPvsSite(v RefResponseModel)`

SetPvsSite sets PvsSite field to given value.


### GetPvsCollection

`func (o *PvsDeviceResponseModel) GetPvsCollection() RefResponseModel`

GetPvsCollection returns the PvsCollection field if non-nil, zero value otherwise.

### GetPvsCollectionOk

`func (o *PvsDeviceResponseModel) GetPvsCollectionOk() (*RefResponseModel, bool)`

GetPvsCollectionOk returns a tuple with the PvsCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsCollection

`func (o *PvsDeviceResponseModel) SetPvsCollection(v RefResponseModel)`

SetPvsCollection sets PvsCollection field to given value.


### GetHypervisorConnection

`func (o *PvsDeviceResponseModel) GetHypervisorConnection() string`

GetHypervisorConnection returns the HypervisorConnection field if non-nil, zero value otherwise.

### GetHypervisorConnectionOk

`func (o *PvsDeviceResponseModel) GetHypervisorConnectionOk() (*string, bool)`

GetHypervisorConnectionOk returns a tuple with the HypervisorConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorConnection

`func (o *PvsDeviceResponseModel) SetHypervisorConnection(v string)`

SetHypervisorConnection sets HypervisorConnection field to given value.

### HasHypervisorConnection

`func (o *PvsDeviceResponseModel) HasHypervisorConnection() bool`

HasHypervisorConnection returns a boolean if a field has been set.

### SetHypervisorConnectionNil

`func (o *PvsDeviceResponseModel) SetHypervisorConnectionNil(b bool)`

 SetHypervisorConnectionNil sets the value for HypervisorConnection to be an explicit nil

### UnsetHypervisorConnection
`func (o *PvsDeviceResponseModel) UnsetHypervisorConnection()`

UnsetHypervisorConnection ensures that no value is present for HypervisorConnection, not even an explicit nil
### GetHostedMachineId

`func (o *PvsDeviceResponseModel) GetHostedMachineId() string`

GetHostedMachineId returns the HostedMachineId field if non-nil, zero value otherwise.

### GetHostedMachineIdOk

`func (o *PvsDeviceResponseModel) GetHostedMachineIdOk() (*string, bool)`

GetHostedMachineIdOk returns a tuple with the HostedMachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedMachineId

`func (o *PvsDeviceResponseModel) SetHostedMachineId(v string)`

SetHostedMachineId sets HostedMachineId field to given value.

### HasHostedMachineId

`func (o *PvsDeviceResponseModel) HasHostedMachineId() bool`

HasHostedMachineId returns a boolean if a field has been set.

### SetHostedMachineIdNil

`func (o *PvsDeviceResponseModel) SetHostedMachineIdNil(b bool)`

 SetHostedMachineIdNil sets the value for HostedMachineId to be an explicit nil

### UnsetHostedMachineId
`func (o *PvsDeviceResponseModel) UnsetHostedMachineId()`

UnsetHostedMachineId ensures that no value is present for HostedMachineId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


