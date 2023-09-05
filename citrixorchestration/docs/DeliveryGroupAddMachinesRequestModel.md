# DeliveryGroupAddMachinesRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MachineCatalog** | **string** | The machine catalog from which to add machines. May be an Id or name. | 
**Count** | Pointer to **int32** | The number of machines to assign from the machine catalog to the delivery group. | [optional] 
**AssignMachinesToUsers** | Pointer to [**[]AssignMachineToUserRequestModel**](AssignMachineToUserRequestModel.md) | List of specific machine assignments to users. | [optional] 

## Methods

### NewDeliveryGroupAddMachinesRequestModel

`func NewDeliveryGroupAddMachinesRequestModel(machineCatalog string, ) *DeliveryGroupAddMachinesRequestModel`

NewDeliveryGroupAddMachinesRequestModel instantiates a new DeliveryGroupAddMachinesRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryGroupAddMachinesRequestModelWithDefaults

`func NewDeliveryGroupAddMachinesRequestModelWithDefaults() *DeliveryGroupAddMachinesRequestModel`

NewDeliveryGroupAddMachinesRequestModelWithDefaults instantiates a new DeliveryGroupAddMachinesRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachineCatalog

`func (o *DeliveryGroupAddMachinesRequestModel) GetMachineCatalog() string`

GetMachineCatalog returns the MachineCatalog field if non-nil, zero value otherwise.

### GetMachineCatalogOk

`func (o *DeliveryGroupAddMachinesRequestModel) GetMachineCatalogOk() (*string, bool)`

GetMachineCatalogOk returns a tuple with the MachineCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineCatalog

`func (o *DeliveryGroupAddMachinesRequestModel) SetMachineCatalog(v string)`

SetMachineCatalog sets MachineCatalog field to given value.


### GetCount

`func (o *DeliveryGroupAddMachinesRequestModel) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DeliveryGroupAddMachinesRequestModel) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DeliveryGroupAddMachinesRequestModel) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *DeliveryGroupAddMachinesRequestModel) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetAssignMachinesToUsers

`func (o *DeliveryGroupAddMachinesRequestModel) GetAssignMachinesToUsers() []AssignMachineToUserRequestModel`

GetAssignMachinesToUsers returns the AssignMachinesToUsers field if non-nil, zero value otherwise.

### GetAssignMachinesToUsersOk

`func (o *DeliveryGroupAddMachinesRequestModel) GetAssignMachinesToUsersOk() (*[]AssignMachineToUserRequestModel, bool)`

GetAssignMachinesToUsersOk returns a tuple with the AssignMachinesToUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignMachinesToUsers

`func (o *DeliveryGroupAddMachinesRequestModel) SetAssignMachinesToUsers(v []AssignMachineToUserRequestModel)`

SetAssignMachinesToUsers sets AssignMachinesToUsers field to given value.

### HasAssignMachinesToUsers

`func (o *DeliveryGroupAddMachinesRequestModel) HasAssignMachinesToUsers() bool`

HasAssignMachinesToUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


