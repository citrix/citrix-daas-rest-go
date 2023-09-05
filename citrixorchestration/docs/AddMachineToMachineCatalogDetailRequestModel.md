# AddMachineToMachineCatalogDetailRequestModel

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
**PvsAddress** | Pointer to **string** | IP address of the PVS server to be used.  This only applies if the ProvisioningType is PVS. | [optional] 
**PvsDomain** | Pointer to **string** | The domain of the PVS server to be used. This only applies if the ProvisioningType is PVS. | [optional] 
**PvsCollectionIds** | Pointer to **[]string** | Collection IDs of PVS collections containing machines that should be added to the catalog.  This only applies if the ProvisioningType is PVS, and is required in that case.  Each item must be a valid PVS collection ID residing on the PVS server located at the specified . | [optional] 
**MachineAccountCreationRules** | Pointer to [**AddMachineToMachineCatalogDetailRequestModelAllOfMachineAccountCreationRules**](AddMachineToMachineCatalogDetailRequestModelAllOfMachineAccountCreationRules.md) |  | [optional] 
**AddAvailableMachineAccount** | Pointer to [**AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount**](AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount.md) |  | [optional] 

## Methods

### NewAddMachineToMachineCatalogDetailRequestModel

`func NewAddMachineToMachineCatalogDetailRequestModel() *AddMachineToMachineCatalogDetailRequestModel`

NewAddMachineToMachineCatalogDetailRequestModel instantiates a new AddMachineToMachineCatalogDetailRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddMachineToMachineCatalogDetailRequestModelWithDefaults

`func NewAddMachineToMachineCatalogDetailRequestModelWithDefaults() *AddMachineToMachineCatalogDetailRequestModel`

NewAddMachineToMachineCatalogDetailRequestModelWithDefaults instantiates a new AddMachineToMachineCatalogDetailRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachineName

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetMachineName() string`

GetMachineName returns the MachineName field if non-nil, zero value otherwise.

### GetMachineNameOk

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetMachineNameOk() (*string, bool)`

GetMachineNameOk returns a tuple with the MachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineName

`func (o *AddMachineToMachineCatalogDetailRequestModel) SetMachineName(v string)`

SetMachineName sets MachineName field to given value.

### HasMachineName

`func (o *AddMachineToMachineCatalogDetailRequestModel) HasMachineName() bool`

HasMachineName returns a boolean if a field has been set.

### GetAssignedClientName

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetAssignedClientName() string`

GetAssignedClientName returns the AssignedClientName field if non-nil, zero value otherwise.

### GetAssignedClientNameOk

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetAssignedClientNameOk() (*string, bool)`

GetAssignedClientNameOk returns a tuple with the AssignedClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedClientName

`func (o *AddMachineToMachineCatalogDetailRequestModel) SetAssignedClientName(v string)`

SetAssignedClientName sets AssignedClientName field to given value.

### HasAssignedClientName

`func (o *AddMachineToMachineCatalogDetailRequestModel) HasAssignedClientName() bool`

HasAssignedClientName returns a boolean if a field has been set.

### GetAssignedIPAddress

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetAssignedIPAddress() string`

GetAssignedIPAddress returns the AssignedIPAddress field if non-nil, zero value otherwise.

### GetAssignedIPAddressOk

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetAssignedIPAddressOk() (*string, bool)`

GetAssignedIPAddressOk returns a tuple with the AssignedIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedIPAddress

`func (o *AddMachineToMachineCatalogDetailRequestModel) SetAssignedIPAddress(v string)`

SetAssignedIPAddress sets AssignedIPAddress field to given value.

### HasAssignedIPAddress

`func (o *AddMachineToMachineCatalogDetailRequestModel) HasAssignedIPAddress() bool`

HasAssignedIPAddress returns a boolean if a field has been set.

### GetAssignedUsers

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetAssignedUsers() []string`

GetAssignedUsers returns the AssignedUsers field if non-nil, zero value otherwise.

### GetAssignedUsersOk

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetAssignedUsersOk() (*[]string, bool)`

GetAssignedUsersOk returns a tuple with the AssignedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUsers

`func (o *AddMachineToMachineCatalogDetailRequestModel) SetAssignedUsers(v []string)`

SetAssignedUsers sets AssignedUsers field to given value.

### HasAssignedUsers

`func (o *AddMachineToMachineCatalogDetailRequestModel) HasAssignedUsers() bool`

HasAssignedUsers returns a boolean if a field has been set.

### GetHostedMachineId

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetHostedMachineId() string`

GetHostedMachineId returns the HostedMachineId field if non-nil, zero value otherwise.

### GetHostedMachineIdOk

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetHostedMachineIdOk() (*string, bool)`

GetHostedMachineIdOk returns a tuple with the HostedMachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedMachineId

`func (o *AddMachineToMachineCatalogDetailRequestModel) SetHostedMachineId(v string)`

SetHostedMachineId sets HostedMachineId field to given value.

### HasHostedMachineId

`func (o *AddMachineToMachineCatalogDetailRequestModel) HasHostedMachineId() bool`

HasHostedMachineId returns a boolean if a field has been set.

### GetHypervisorConnection

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetHypervisorConnection() string`

GetHypervisorConnection returns the HypervisorConnection field if non-nil, zero value otherwise.

### GetHypervisorConnectionOk

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetHypervisorConnectionOk() (*string, bool)`

GetHypervisorConnectionOk returns a tuple with the HypervisorConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorConnection

`func (o *AddMachineToMachineCatalogDetailRequestModel) SetHypervisorConnection(v string)`

SetHypervisorConnection sets HypervisorConnection field to given value.

### HasHypervisorConnection

`func (o *AddMachineToMachineCatalogDetailRequestModel) HasHypervisorConnection() bool`

HasHypervisorConnection returns a boolean if a field has been set.

### GetInMaintenanceMode

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *AddMachineToMachineCatalogDetailRequestModel) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.

### HasInMaintenanceMode

`func (o *AddMachineToMachineCatalogDetailRequestModel) HasInMaintenanceMode() bool`

HasInMaintenanceMode returns a boolean if a field has been set.

### GetPvsAddress

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetPvsAddress() string`

GetPvsAddress returns the PvsAddress field if non-nil, zero value otherwise.

### GetPvsAddressOk

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetPvsAddressOk() (*string, bool)`

GetPvsAddressOk returns a tuple with the PvsAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsAddress

`func (o *AddMachineToMachineCatalogDetailRequestModel) SetPvsAddress(v string)`

SetPvsAddress sets PvsAddress field to given value.

### HasPvsAddress

`func (o *AddMachineToMachineCatalogDetailRequestModel) HasPvsAddress() bool`

HasPvsAddress returns a boolean if a field has been set.

### GetPvsDomain

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetPvsDomain() string`

GetPvsDomain returns the PvsDomain field if non-nil, zero value otherwise.

### GetPvsDomainOk

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetPvsDomainOk() (*string, bool)`

GetPvsDomainOk returns a tuple with the PvsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsDomain

`func (o *AddMachineToMachineCatalogDetailRequestModel) SetPvsDomain(v string)`

SetPvsDomain sets PvsDomain field to given value.

### HasPvsDomain

`func (o *AddMachineToMachineCatalogDetailRequestModel) HasPvsDomain() bool`

HasPvsDomain returns a boolean if a field has been set.

### GetPvsCollectionIds

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetPvsCollectionIds() []string`

GetPvsCollectionIds returns the PvsCollectionIds field if non-nil, zero value otherwise.

### GetPvsCollectionIdsOk

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetPvsCollectionIdsOk() (*[]string, bool)`

GetPvsCollectionIdsOk returns a tuple with the PvsCollectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsCollectionIds

`func (o *AddMachineToMachineCatalogDetailRequestModel) SetPvsCollectionIds(v []string)`

SetPvsCollectionIds sets PvsCollectionIds field to given value.

### HasPvsCollectionIds

`func (o *AddMachineToMachineCatalogDetailRequestModel) HasPvsCollectionIds() bool`

HasPvsCollectionIds returns a boolean if a field has been set.

### GetMachineAccountCreationRules

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetMachineAccountCreationRules() AddMachineToMachineCatalogDetailRequestModelAllOfMachineAccountCreationRules`

GetMachineAccountCreationRules returns the MachineAccountCreationRules field if non-nil, zero value otherwise.

### GetMachineAccountCreationRulesOk

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetMachineAccountCreationRulesOk() (*AddMachineToMachineCatalogDetailRequestModelAllOfMachineAccountCreationRules, bool)`

GetMachineAccountCreationRulesOk returns a tuple with the MachineAccountCreationRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineAccountCreationRules

`func (o *AddMachineToMachineCatalogDetailRequestModel) SetMachineAccountCreationRules(v AddMachineToMachineCatalogDetailRequestModelAllOfMachineAccountCreationRules)`

SetMachineAccountCreationRules sets MachineAccountCreationRules field to given value.

### HasMachineAccountCreationRules

`func (o *AddMachineToMachineCatalogDetailRequestModel) HasMachineAccountCreationRules() bool`

HasMachineAccountCreationRules returns a boolean if a field has been set.

### GetAddAvailableMachineAccount

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetAddAvailableMachineAccount() AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount`

GetAddAvailableMachineAccount returns the AddAvailableMachineAccount field if non-nil, zero value otherwise.

### GetAddAvailableMachineAccountOk

`func (o *AddMachineToMachineCatalogDetailRequestModel) GetAddAvailableMachineAccountOk() (*AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount, bool)`

GetAddAvailableMachineAccountOk returns a tuple with the AddAvailableMachineAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAvailableMachineAccount

`func (o *AddMachineToMachineCatalogDetailRequestModel) SetAddAvailableMachineAccount(v AddMachineToMachineCatalogDetailRequestModelAllOfAddAvailableMachineAccount)`

SetAddAvailableMachineAccount sets AddAvailableMachineAccount field to given value.

### HasAddAvailableMachineAccount

`func (o *AddMachineToMachineCatalogDetailRequestModel) HasAddAvailableMachineAccount() bool`

HasAddAvailableMachineAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


