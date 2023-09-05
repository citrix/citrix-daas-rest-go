# MachineHostingResponseModelHypervisorConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the object. Needs to be globally unique This is likely to contain a guid but the property type should be string for future flexibility | [optional] 
**Uid** | Pointer to **int32** | DEPRECATED. Use Id. | [optional] 
**Name** | Pointer to **string** | Name of the object. | [optional] 

## Methods

### NewMachineHostingResponseModelHypervisorConnection

`func NewMachineHostingResponseModelHypervisorConnection() *MachineHostingResponseModelHypervisorConnection`

NewMachineHostingResponseModelHypervisorConnection instantiates a new MachineHostingResponseModelHypervisorConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineHostingResponseModelHypervisorConnectionWithDefaults

`func NewMachineHostingResponseModelHypervisorConnectionWithDefaults() *MachineHostingResponseModelHypervisorConnection`

NewMachineHostingResponseModelHypervisorConnectionWithDefaults instantiates a new MachineHostingResponseModelHypervisorConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MachineHostingResponseModelHypervisorConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MachineHostingResponseModelHypervisorConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MachineHostingResponseModelHypervisorConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MachineHostingResponseModelHypervisorConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *MachineHostingResponseModelHypervisorConnection) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *MachineHostingResponseModelHypervisorConnection) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *MachineHostingResponseModelHypervisorConnection) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *MachineHostingResponseModelHypervisorConnection) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetName

`func (o *MachineHostingResponseModelHypervisorConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MachineHostingResponseModelHypervisorConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MachineHostingResponseModelHypervisorConnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MachineHostingResponseModelHypervisorConnection) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


