# UpdateMachineRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedClientName** | Pointer to **NullableString** | The client name to which this machine will be assigned.  Machines can be assigned to multiple users, a single client IP address, or a single client name, but only to one of these categories at a time. | [optional] 
**AssignedIPAddress** | Pointer to **NullableString** | The client IP address to which this machine will be assigned.  Machines can be assigned to multiple users, a single client IP address, or a single client name, but only to one of these categories at a time. | [optional] 
**AssignedUsers** | Pointer to **[]string** | The user(s) to whom this machine will be assigned.  Machines can be assigned to multiple users, a single client IP address, or a single client name, but only to one of these categories at a time. | [optional] 
**HostedMachineId** | Pointer to **NullableString** | The unique ID by which the hypervisor recognizes the machine. Omit this for machines that are not power-managed. | [optional] 
**HypervisorConnection** | Pointer to **NullableString** | Hypervisor connection to use for power management of the machine. | [optional] 
**InMaintenanceMode** | Pointer to **NullableBool** | Specifies whether the machine is initially in maintenance mode.  A machine in maintenance mode is not available for new sessions, and for managed machines all automatic power management is disabled. If &#x60;null&#x60;, will not be changed. | [optional] 
**PublishedName** | Pointer to **NullableString** | Customized name of the machine that is displayed in StoreFront, if the machine has been published. It can be set only for private desktops. If &#x60;null&#x60;, will not be changed. If empty string (&#x60;\&quot;\&quot;&#x60;), the machine will be unassigned from any published name. | [optional] 
**Metadata** | Pointer to [**[]NameValueStringPairModel**](NameValueStringPairModel.md) | The metadata of machine. Set the value of the NameValueStringPairModel is null or empty will be remove this metadata. Not existing Name and Value NameValueStringPairModel object will be added. The same Name but different value object will be updated. | [optional] 

## Methods

### NewUpdateMachineRequestModel

`func NewUpdateMachineRequestModel() *UpdateMachineRequestModel`

NewUpdateMachineRequestModel instantiates a new UpdateMachineRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMachineRequestModelWithDefaults

`func NewUpdateMachineRequestModelWithDefaults() *UpdateMachineRequestModel`

NewUpdateMachineRequestModelWithDefaults instantiates a new UpdateMachineRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedClientName

`func (o *UpdateMachineRequestModel) GetAssignedClientName() string`

GetAssignedClientName returns the AssignedClientName field if non-nil, zero value otherwise.

### GetAssignedClientNameOk

`func (o *UpdateMachineRequestModel) GetAssignedClientNameOk() (*string, bool)`

GetAssignedClientNameOk returns a tuple with the AssignedClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedClientName

`func (o *UpdateMachineRequestModel) SetAssignedClientName(v string)`

SetAssignedClientName sets AssignedClientName field to given value.

### HasAssignedClientName

`func (o *UpdateMachineRequestModel) HasAssignedClientName() bool`

HasAssignedClientName returns a boolean if a field has been set.

### SetAssignedClientNameNil

`func (o *UpdateMachineRequestModel) SetAssignedClientNameNil(b bool)`

 SetAssignedClientNameNil sets the value for AssignedClientName to be an explicit nil

### UnsetAssignedClientName
`func (o *UpdateMachineRequestModel) UnsetAssignedClientName()`

UnsetAssignedClientName ensures that no value is present for AssignedClientName, not even an explicit nil
### GetAssignedIPAddress

`func (o *UpdateMachineRequestModel) GetAssignedIPAddress() string`

GetAssignedIPAddress returns the AssignedIPAddress field if non-nil, zero value otherwise.

### GetAssignedIPAddressOk

`func (o *UpdateMachineRequestModel) GetAssignedIPAddressOk() (*string, bool)`

GetAssignedIPAddressOk returns a tuple with the AssignedIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedIPAddress

`func (o *UpdateMachineRequestModel) SetAssignedIPAddress(v string)`

SetAssignedIPAddress sets AssignedIPAddress field to given value.

### HasAssignedIPAddress

`func (o *UpdateMachineRequestModel) HasAssignedIPAddress() bool`

HasAssignedIPAddress returns a boolean if a field has been set.

### SetAssignedIPAddressNil

`func (o *UpdateMachineRequestModel) SetAssignedIPAddressNil(b bool)`

 SetAssignedIPAddressNil sets the value for AssignedIPAddress to be an explicit nil

### UnsetAssignedIPAddress
`func (o *UpdateMachineRequestModel) UnsetAssignedIPAddress()`

UnsetAssignedIPAddress ensures that no value is present for AssignedIPAddress, not even an explicit nil
### GetAssignedUsers

`func (o *UpdateMachineRequestModel) GetAssignedUsers() []string`

GetAssignedUsers returns the AssignedUsers field if non-nil, zero value otherwise.

### GetAssignedUsersOk

`func (o *UpdateMachineRequestModel) GetAssignedUsersOk() (*[]string, bool)`

GetAssignedUsersOk returns a tuple with the AssignedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUsers

`func (o *UpdateMachineRequestModel) SetAssignedUsers(v []string)`

SetAssignedUsers sets AssignedUsers field to given value.

### HasAssignedUsers

`func (o *UpdateMachineRequestModel) HasAssignedUsers() bool`

HasAssignedUsers returns a boolean if a field has been set.

### SetAssignedUsersNil

`func (o *UpdateMachineRequestModel) SetAssignedUsersNil(b bool)`

 SetAssignedUsersNil sets the value for AssignedUsers to be an explicit nil

### UnsetAssignedUsers
`func (o *UpdateMachineRequestModel) UnsetAssignedUsers()`

UnsetAssignedUsers ensures that no value is present for AssignedUsers, not even an explicit nil
### GetHostedMachineId

`func (o *UpdateMachineRequestModel) GetHostedMachineId() string`

GetHostedMachineId returns the HostedMachineId field if non-nil, zero value otherwise.

### GetHostedMachineIdOk

`func (o *UpdateMachineRequestModel) GetHostedMachineIdOk() (*string, bool)`

GetHostedMachineIdOk returns a tuple with the HostedMachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedMachineId

`func (o *UpdateMachineRequestModel) SetHostedMachineId(v string)`

SetHostedMachineId sets HostedMachineId field to given value.

### HasHostedMachineId

`func (o *UpdateMachineRequestModel) HasHostedMachineId() bool`

HasHostedMachineId returns a boolean if a field has been set.

### SetHostedMachineIdNil

`func (o *UpdateMachineRequestModel) SetHostedMachineIdNil(b bool)`

 SetHostedMachineIdNil sets the value for HostedMachineId to be an explicit nil

### UnsetHostedMachineId
`func (o *UpdateMachineRequestModel) UnsetHostedMachineId()`

UnsetHostedMachineId ensures that no value is present for HostedMachineId, not even an explicit nil
### GetHypervisorConnection

`func (o *UpdateMachineRequestModel) GetHypervisorConnection() string`

GetHypervisorConnection returns the HypervisorConnection field if non-nil, zero value otherwise.

### GetHypervisorConnectionOk

`func (o *UpdateMachineRequestModel) GetHypervisorConnectionOk() (*string, bool)`

GetHypervisorConnectionOk returns a tuple with the HypervisorConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorConnection

`func (o *UpdateMachineRequestModel) SetHypervisorConnection(v string)`

SetHypervisorConnection sets HypervisorConnection field to given value.

### HasHypervisorConnection

`func (o *UpdateMachineRequestModel) HasHypervisorConnection() bool`

HasHypervisorConnection returns a boolean if a field has been set.

### SetHypervisorConnectionNil

`func (o *UpdateMachineRequestModel) SetHypervisorConnectionNil(b bool)`

 SetHypervisorConnectionNil sets the value for HypervisorConnection to be an explicit nil

### UnsetHypervisorConnection
`func (o *UpdateMachineRequestModel) UnsetHypervisorConnection()`

UnsetHypervisorConnection ensures that no value is present for HypervisorConnection, not even an explicit nil
### GetInMaintenanceMode

`func (o *UpdateMachineRequestModel) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *UpdateMachineRequestModel) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *UpdateMachineRequestModel) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.

### HasInMaintenanceMode

`func (o *UpdateMachineRequestModel) HasInMaintenanceMode() bool`

HasInMaintenanceMode returns a boolean if a field has been set.

### SetInMaintenanceModeNil

`func (o *UpdateMachineRequestModel) SetInMaintenanceModeNil(b bool)`

 SetInMaintenanceModeNil sets the value for InMaintenanceMode to be an explicit nil

### UnsetInMaintenanceMode
`func (o *UpdateMachineRequestModel) UnsetInMaintenanceMode()`

UnsetInMaintenanceMode ensures that no value is present for InMaintenanceMode, not even an explicit nil
### GetPublishedName

`func (o *UpdateMachineRequestModel) GetPublishedName() string`

GetPublishedName returns the PublishedName field if non-nil, zero value otherwise.

### GetPublishedNameOk

`func (o *UpdateMachineRequestModel) GetPublishedNameOk() (*string, bool)`

GetPublishedNameOk returns a tuple with the PublishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedName

`func (o *UpdateMachineRequestModel) SetPublishedName(v string)`

SetPublishedName sets PublishedName field to given value.

### HasPublishedName

`func (o *UpdateMachineRequestModel) HasPublishedName() bool`

HasPublishedName returns a boolean if a field has been set.

### SetPublishedNameNil

`func (o *UpdateMachineRequestModel) SetPublishedNameNil(b bool)`

 SetPublishedNameNil sets the value for PublishedName to be an explicit nil

### UnsetPublishedName
`func (o *UpdateMachineRequestModel) UnsetPublishedName()`

UnsetPublishedName ensures that no value is present for PublishedName, not even an explicit nil
### GetMetadata

`func (o *UpdateMachineRequestModel) GetMetadata() []NameValueStringPairModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateMachineRequestModel) GetMetadataOk() (*[]NameValueStringPairModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateMachineRequestModel) SetMetadata(v []NameValueStringPairModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateMachineRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *UpdateMachineRequestModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *UpdateMachineRequestModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


