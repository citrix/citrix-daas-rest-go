# AssignMachineToUserRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Machine** | **string** | The machine to assign. May be an Id or name. | 
**Users** | Pointer to **[]string** | The user(s) to whom the machine should be assigned. Caller can specify SamName, UPN, or SID for each user. The system should look up the entity to determine the other AD properties (such as DisplayName and IsGroup properties). Groups should be rejected. If the caller passes an empty list, the machine is de-allocated if it was previously allocated. | [optional] 

## Methods

### NewAssignMachineToUserRequestModel

`func NewAssignMachineToUserRequestModel(machine string, ) *AssignMachineToUserRequestModel`

NewAssignMachineToUserRequestModel instantiates a new AssignMachineToUserRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignMachineToUserRequestModelWithDefaults

`func NewAssignMachineToUserRequestModelWithDefaults() *AssignMachineToUserRequestModel`

NewAssignMachineToUserRequestModelWithDefaults instantiates a new AssignMachineToUserRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachine

`func (o *AssignMachineToUserRequestModel) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *AssignMachineToUserRequestModel) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *AssignMachineToUserRequestModel) SetMachine(v string)`

SetMachine sets Machine field to given value.


### GetUsers

`func (o *AssignMachineToUserRequestModel) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AssignMachineToUserRequestModel) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AssignMachineToUserRequestModel) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *AssignMachineToUserRequestModel) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *AssignMachineToUserRequestModel) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *AssignMachineToUserRequestModel) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


