# AddCatalogMachineAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MachineName** | Pointer to **string** | Name of the machine (in UPN format) to be assigned to a catalog. | [optional] 
**AssignedUsers** | Pointer to **[]string** | List of users to assign to the machine. | [optional] 

## Methods

### NewAddCatalogMachineAssignment

`func NewAddCatalogMachineAssignment() *AddCatalogMachineAssignment`

NewAddCatalogMachineAssignment instantiates a new AddCatalogMachineAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCatalogMachineAssignmentWithDefaults

`func NewAddCatalogMachineAssignmentWithDefaults() *AddCatalogMachineAssignment`

NewAddCatalogMachineAssignmentWithDefaults instantiates a new AddCatalogMachineAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachineName

`func (o *AddCatalogMachineAssignment) GetMachineName() string`

GetMachineName returns the MachineName field if non-nil, zero value otherwise.

### GetMachineNameOk

`func (o *AddCatalogMachineAssignment) GetMachineNameOk() (*string, bool)`

GetMachineNameOk returns a tuple with the MachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineName

`func (o *AddCatalogMachineAssignment) SetMachineName(v string)`

SetMachineName sets MachineName field to given value.

### HasMachineName

`func (o *AddCatalogMachineAssignment) HasMachineName() bool`

HasMachineName returns a boolean if a field has been set.

### GetAssignedUsers

`func (o *AddCatalogMachineAssignment) GetAssignedUsers() []string`

GetAssignedUsers returns the AssignedUsers field if non-nil, zero value otherwise.

### GetAssignedUsersOk

`func (o *AddCatalogMachineAssignment) GetAssignedUsersOk() (*[]string, bool)`

GetAssignedUsersOk returns a tuple with the AssignedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUsers

`func (o *AddCatalogMachineAssignment) SetAssignedUsers(v []string)`

SetAssignedUsers sets AssignedUsers field to given value.

### HasAssignedUsers

`func (o *AddCatalogMachineAssignment) HasAssignedUsers() bool`

HasAssignedUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


