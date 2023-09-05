# AddMachineToMachineCatalogRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MachineName** | Pointer to **string** | Specify the name of the machine to create (in the form &#39;domain\\machine&#39;).  A SID can also be specified. | [optional] 
**AssignedClientName** | Pointer to **string** | The client name to which this machine will be assigned.  Machines can be assigned to multiple users, a single client IP address, or a single client name, but only to one of these categories at a time. | [optional] 
**AssignedIPAddress** | Pointer to **string** | The client IP address to which this machine will be assigned.  Machines can be assigned to multiple users, a single client IP address, or a single client name, but only to one of these categories at a time. | [optional] 
**AssignedUsers** | Pointer to **[]string** | The user(s) to whom this machine will be assigned.  Machines can be assigned to multiple users, a single client IP address, or a single client name, but only to one of these categories at a time. | [optional] 
**HostedMachineId** | Pointer to **string** | The unique ID by which the hypervisor recognizes the machine. Omit this for machines that are not power-managed. | [optional] 
**HypervisorConnection** | Pointer to **string** | Hypervisor connection to use for power management of the machine. | [optional] 
**InMaintenanceMode** | Pointer to **bool** | Specifies whether the machine is initially in maintenance mode.  A machine in maintenance mode is not available for new sessions, and for managed machines all automatic power management is disabled. Optional; default is &#x60;false&#x60;. | [optional] [default to false]

## Methods

### NewAddMachineToMachineCatalogRequestModel

`func NewAddMachineToMachineCatalogRequestModel() *AddMachineToMachineCatalogRequestModel`

NewAddMachineToMachineCatalogRequestModel instantiates a new AddMachineToMachineCatalogRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddMachineToMachineCatalogRequestModelWithDefaults

`func NewAddMachineToMachineCatalogRequestModelWithDefaults() *AddMachineToMachineCatalogRequestModel`

NewAddMachineToMachineCatalogRequestModelWithDefaults instantiates a new AddMachineToMachineCatalogRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachineName

`func (o *AddMachineToMachineCatalogRequestModel) GetMachineName() string`

GetMachineName returns the MachineName field if non-nil, zero value otherwise.

### GetMachineNameOk

`func (o *AddMachineToMachineCatalogRequestModel) GetMachineNameOk() (*string, bool)`

GetMachineNameOk returns a tuple with the MachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineName

`func (o *AddMachineToMachineCatalogRequestModel) SetMachineName(v string)`

SetMachineName sets MachineName field to given value.

### HasMachineName

`func (o *AddMachineToMachineCatalogRequestModel) HasMachineName() bool`

HasMachineName returns a boolean if a field has been set.

### GetAssignedClientName

`func (o *AddMachineToMachineCatalogRequestModel) GetAssignedClientName() string`

GetAssignedClientName returns the AssignedClientName field if non-nil, zero value otherwise.

### GetAssignedClientNameOk

`func (o *AddMachineToMachineCatalogRequestModel) GetAssignedClientNameOk() (*string, bool)`

GetAssignedClientNameOk returns a tuple with the AssignedClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedClientName

`func (o *AddMachineToMachineCatalogRequestModel) SetAssignedClientName(v string)`

SetAssignedClientName sets AssignedClientName field to given value.

### HasAssignedClientName

`func (o *AddMachineToMachineCatalogRequestModel) HasAssignedClientName() bool`

HasAssignedClientName returns a boolean if a field has been set.

### GetAssignedIPAddress

`func (o *AddMachineToMachineCatalogRequestModel) GetAssignedIPAddress() string`

GetAssignedIPAddress returns the AssignedIPAddress field if non-nil, zero value otherwise.

### GetAssignedIPAddressOk

`func (o *AddMachineToMachineCatalogRequestModel) GetAssignedIPAddressOk() (*string, bool)`

GetAssignedIPAddressOk returns a tuple with the AssignedIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedIPAddress

`func (o *AddMachineToMachineCatalogRequestModel) SetAssignedIPAddress(v string)`

SetAssignedIPAddress sets AssignedIPAddress field to given value.

### HasAssignedIPAddress

`func (o *AddMachineToMachineCatalogRequestModel) HasAssignedIPAddress() bool`

HasAssignedIPAddress returns a boolean if a field has been set.

### GetAssignedUsers

`func (o *AddMachineToMachineCatalogRequestModel) GetAssignedUsers() []string`

GetAssignedUsers returns the AssignedUsers field if non-nil, zero value otherwise.

### GetAssignedUsersOk

`func (o *AddMachineToMachineCatalogRequestModel) GetAssignedUsersOk() (*[]string, bool)`

GetAssignedUsersOk returns a tuple with the AssignedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUsers

`func (o *AddMachineToMachineCatalogRequestModel) SetAssignedUsers(v []string)`

SetAssignedUsers sets AssignedUsers field to given value.

### HasAssignedUsers

`func (o *AddMachineToMachineCatalogRequestModel) HasAssignedUsers() bool`

HasAssignedUsers returns a boolean if a field has been set.

### GetHostedMachineId

`func (o *AddMachineToMachineCatalogRequestModel) GetHostedMachineId() string`

GetHostedMachineId returns the HostedMachineId field if non-nil, zero value otherwise.

### GetHostedMachineIdOk

`func (o *AddMachineToMachineCatalogRequestModel) GetHostedMachineIdOk() (*string, bool)`

GetHostedMachineIdOk returns a tuple with the HostedMachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedMachineId

`func (o *AddMachineToMachineCatalogRequestModel) SetHostedMachineId(v string)`

SetHostedMachineId sets HostedMachineId field to given value.

### HasHostedMachineId

`func (o *AddMachineToMachineCatalogRequestModel) HasHostedMachineId() bool`

HasHostedMachineId returns a boolean if a field has been set.

### GetHypervisorConnection

`func (o *AddMachineToMachineCatalogRequestModel) GetHypervisorConnection() string`

GetHypervisorConnection returns the HypervisorConnection field if non-nil, zero value otherwise.

### GetHypervisorConnectionOk

`func (o *AddMachineToMachineCatalogRequestModel) GetHypervisorConnectionOk() (*string, bool)`

GetHypervisorConnectionOk returns a tuple with the HypervisorConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorConnection

`func (o *AddMachineToMachineCatalogRequestModel) SetHypervisorConnection(v string)`

SetHypervisorConnection sets HypervisorConnection field to given value.

### HasHypervisorConnection

`func (o *AddMachineToMachineCatalogRequestModel) HasHypervisorConnection() bool`

HasHypervisorConnection returns a boolean if a field has been set.

### GetInMaintenanceMode

`func (o *AddMachineToMachineCatalogRequestModel) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *AddMachineToMachineCatalogRequestModel) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *AddMachineToMachineCatalogRequestModel) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.

### HasInMaintenanceMode

`func (o *AddMachineToMachineCatalogRequestModel) HasInMaintenanceMode() bool`

HasInMaintenanceMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


