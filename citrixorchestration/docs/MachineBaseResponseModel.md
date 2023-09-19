# MachineBaseResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of machine. Used to be: DesktopUid (and wasn&#39;t globally unique) OR UUID, depending on context Needs to be globally unique Might be constructed from site ID + internal Uid?  or use uuid | 
**MachineCatalog** | Pointer to [**RefResponseModel**](RefResponseModel.md) |  | [optional] 
**Name** | Pointer to **NullableString** | DNS host name of the machine. Used to be: MachineName | [optional] 

## Methods

### NewMachineBaseResponseModel

`func NewMachineBaseResponseModel(id string, ) *MachineBaseResponseModel`

NewMachineBaseResponseModel instantiates a new MachineBaseResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineBaseResponseModelWithDefaults

`func NewMachineBaseResponseModelWithDefaults() *MachineBaseResponseModel`

NewMachineBaseResponseModelWithDefaults instantiates a new MachineBaseResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MachineBaseResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MachineBaseResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MachineBaseResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetMachineCatalog

`func (o *MachineBaseResponseModel) GetMachineCatalog() RefResponseModel`

GetMachineCatalog returns the MachineCatalog field if non-nil, zero value otherwise.

### GetMachineCatalogOk

`func (o *MachineBaseResponseModel) GetMachineCatalogOk() (*RefResponseModel, bool)`

GetMachineCatalogOk returns a tuple with the MachineCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineCatalog

`func (o *MachineBaseResponseModel) SetMachineCatalog(v RefResponseModel)`

SetMachineCatalog sets MachineCatalog field to given value.

### HasMachineCatalog

`func (o *MachineBaseResponseModel) HasMachineCatalog() bool`

HasMachineCatalog returns a boolean if a field has been set.

### GetName

`func (o *MachineBaseResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MachineBaseResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MachineBaseResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MachineBaseResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MachineBaseResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MachineBaseResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


