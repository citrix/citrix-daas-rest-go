# AddMachineToMachineCatalogRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MachineName** | Pointer to **NullableString** | Specify the name of the machine to create (in the form &#39;domain\\machine&#39;).  A SID can also be specified. | [optional] 
**AssignedClientName** | Pointer to **NullableString** | The client name to which this machine will be assigned.  Machines can be assigned to multiple users, a single client IP address, or a single client name, but only to one of these categories at a time. | [optional] 
**AssignedIPAddress** | Pointer to **NullableString** | The client IP address to which this machine will be assigned.  Machines can be assigned to multiple users, a single client IP address, or a single client name, but only to one of these categories at a time. | [optional] 
**AssignedUsers** | Pointer to **[]string** | The user(s) to whom this machine will be assigned.  Machines can be assigned to multiple users, a single client IP address, or a single client name, but only to one of these categories at a time. | [optional] 
**HostedMachineId** | Pointer to **NullableString** | The unique ID by which the hypervisor recognizes the machine. Omit this for machines that are not power-managed. | [optional] 
**HypervisorConnection** | Pointer to **NullableString** | Hypervisor connection to use for power management of the machine. | [optional] 
**InMaintenanceMode** | Pointer to **NullableBool** | Specifies whether the machine is initially in maintenance mode.  A machine in maintenance mode is not available for new sessions, and for managed machines all automatic power management is disabled. Optional; default is &#x60;false&#x60;. | [optional] [default to false]
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of machine. | [optional] 

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

### SetMachineNameNil

`func (o *AddMachineToMachineCatalogRequestModel) SetMachineNameNil(b bool)`

 SetMachineNameNil sets the value for MachineName to be an explicit nil

### UnsetMachineName
`func (o *AddMachineToMachineCatalogRequestModel) UnsetMachineName()`

UnsetMachineName ensures that no value is present for MachineName, not even an explicit nil
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

### SetAssignedClientNameNil

`func (o *AddMachineToMachineCatalogRequestModel) SetAssignedClientNameNil(b bool)`

 SetAssignedClientNameNil sets the value for AssignedClientName to be an explicit nil

### UnsetAssignedClientName
`func (o *AddMachineToMachineCatalogRequestModel) UnsetAssignedClientName()`

UnsetAssignedClientName ensures that no value is present for AssignedClientName, not even an explicit nil
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

### SetAssignedIPAddressNil

`func (o *AddMachineToMachineCatalogRequestModel) SetAssignedIPAddressNil(b bool)`

 SetAssignedIPAddressNil sets the value for AssignedIPAddress to be an explicit nil

### UnsetAssignedIPAddress
`func (o *AddMachineToMachineCatalogRequestModel) UnsetAssignedIPAddress()`

UnsetAssignedIPAddress ensures that no value is present for AssignedIPAddress, not even an explicit nil
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

### SetAssignedUsersNil

`func (o *AddMachineToMachineCatalogRequestModel) SetAssignedUsersNil(b bool)`

 SetAssignedUsersNil sets the value for AssignedUsers to be an explicit nil

### UnsetAssignedUsers
`func (o *AddMachineToMachineCatalogRequestModel) UnsetAssignedUsers()`

UnsetAssignedUsers ensures that no value is present for AssignedUsers, not even an explicit nil
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

### SetHostedMachineIdNil

`func (o *AddMachineToMachineCatalogRequestModel) SetHostedMachineIdNil(b bool)`

 SetHostedMachineIdNil sets the value for HostedMachineId to be an explicit nil

### UnsetHostedMachineId
`func (o *AddMachineToMachineCatalogRequestModel) UnsetHostedMachineId()`

UnsetHostedMachineId ensures that no value is present for HostedMachineId, not even an explicit nil
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

### SetHypervisorConnectionNil

`func (o *AddMachineToMachineCatalogRequestModel) SetHypervisorConnectionNil(b bool)`

 SetHypervisorConnectionNil sets the value for HypervisorConnection to be an explicit nil

### UnsetHypervisorConnection
`func (o *AddMachineToMachineCatalogRequestModel) UnsetHypervisorConnection()`

UnsetHypervisorConnection ensures that no value is present for HypervisorConnection, not even an explicit nil
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

### SetInMaintenanceModeNil

`func (o *AddMachineToMachineCatalogRequestModel) SetInMaintenanceModeNil(b bool)`

 SetInMaintenanceModeNil sets the value for InMaintenanceMode to be an explicit nil

### UnsetInMaintenanceMode
`func (o *AddMachineToMachineCatalogRequestModel) UnsetInMaintenanceMode()`

UnsetInMaintenanceMode ensures that no value is present for InMaintenanceMode, not even an explicit nil
### GetMetadata

`func (o *AddMachineToMachineCatalogRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AddMachineToMachineCatalogRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AddMachineToMachineCatalogRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AddMachineToMachineCatalogRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *AddMachineToMachineCatalogRequestModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *AddMachineToMachineCatalogRequestModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


