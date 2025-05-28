# DeploymentMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | [**AccountType**](AccountType.md) | The type of provider associated with the account | 
**DeploymentId** | Pointer to **NullableString** | Deployment Id | [optional] 
**BrokerMachineId** | Pointer to **NullableString** | MachineId on the Broker | [optional] 
**TaskType** | Pointer to [**NullableTaskType**](TaskType.md) | Task Type | [optional] 
**TaskId** | Pointer to **NullableString** | Task Id | [optional] 
**ConnectionId** | Pointer to **NullableString** | Hosting Connection Id | [optional] 
**ImageId** | Pointer to **NullableString** | Image Id | [optional] 
**MachineId** | Pointer to **NullableString** | Machine Id | [optional] 
**MachineName** | Pointer to **NullableString** | Name of the machine | [optional] 
**RegistrationState** | Pointer to [**NullableRegistrationState**](RegistrationState.md) | Registration state of the machine | [optional] 
**SessionState** | Pointer to [**NullableSessionState**](SessionState.md) | State of active session on machine | [optional] 
**SessionCount** | Pointer to **NullableInt32** | Count of active session on machine | [optional] 
**MaintenanceMode** | Pointer to **NullableBool** | Indicates if the machine is in Maintenance Mode | [optional] 
**AssociatedUsers** | Pointer to **[]string** | List of users that are associated with the machine | [optional] 
**BrokerMachineCatalogId** | Pointer to **NullableString** | Id for the machine catalog of the machine, could be null if machine not in machine catalog | [optional] 
**BrokerDeliveryGroupId** | Pointer to **NullableString** | Id for the delivery group of the machine, could be null if machine not in delivery group | [optional] 

## Methods

### NewDeploymentMachine

`func NewDeploymentMachine(accountType AccountType, ) *DeploymentMachine`

NewDeploymentMachine instantiates a new DeploymentMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentMachineWithDefaults

`func NewDeploymentMachineWithDefaults() *DeploymentMachine`

NewDeploymentMachineWithDefaults instantiates a new DeploymentMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *DeploymentMachine) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *DeploymentMachine) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *DeploymentMachine) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetDeploymentId

`func (o *DeploymentMachine) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *DeploymentMachine) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *DeploymentMachine) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *DeploymentMachine) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### SetDeploymentIdNil

`func (o *DeploymentMachine) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *DeploymentMachine) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
### GetBrokerMachineId

`func (o *DeploymentMachine) GetBrokerMachineId() string`

GetBrokerMachineId returns the BrokerMachineId field if non-nil, zero value otherwise.

### GetBrokerMachineIdOk

`func (o *DeploymentMachine) GetBrokerMachineIdOk() (*string, bool)`

GetBrokerMachineIdOk returns a tuple with the BrokerMachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerMachineId

`func (o *DeploymentMachine) SetBrokerMachineId(v string)`

SetBrokerMachineId sets BrokerMachineId field to given value.

### HasBrokerMachineId

`func (o *DeploymentMachine) HasBrokerMachineId() bool`

HasBrokerMachineId returns a boolean if a field has been set.

### SetBrokerMachineIdNil

`func (o *DeploymentMachine) SetBrokerMachineIdNil(b bool)`

 SetBrokerMachineIdNil sets the value for BrokerMachineId to be an explicit nil

### UnsetBrokerMachineId
`func (o *DeploymentMachine) UnsetBrokerMachineId()`

UnsetBrokerMachineId ensures that no value is present for BrokerMachineId, not even an explicit nil
### GetTaskType

`func (o *DeploymentMachine) GetTaskType() TaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *DeploymentMachine) GetTaskTypeOk() (*TaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *DeploymentMachine) SetTaskType(v TaskType)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *DeploymentMachine) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### SetTaskTypeNil

`func (o *DeploymentMachine) SetTaskTypeNil(b bool)`

 SetTaskTypeNil sets the value for TaskType to be an explicit nil

### UnsetTaskType
`func (o *DeploymentMachine) UnsetTaskType()`

UnsetTaskType ensures that no value is present for TaskType, not even an explicit nil
### GetTaskId

`func (o *DeploymentMachine) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *DeploymentMachine) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *DeploymentMachine) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *DeploymentMachine) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *DeploymentMachine) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *DeploymentMachine) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetConnectionId

`func (o *DeploymentMachine) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *DeploymentMachine) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *DeploymentMachine) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *DeploymentMachine) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *DeploymentMachine) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *DeploymentMachine) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetImageId

`func (o *DeploymentMachine) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *DeploymentMachine) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *DeploymentMachine) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *DeploymentMachine) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### SetImageIdNil

`func (o *DeploymentMachine) SetImageIdNil(b bool)`

 SetImageIdNil sets the value for ImageId to be an explicit nil

### UnsetImageId
`func (o *DeploymentMachine) UnsetImageId()`

UnsetImageId ensures that no value is present for ImageId, not even an explicit nil
### GetMachineId

`func (o *DeploymentMachine) GetMachineId() string`

GetMachineId returns the MachineId field if non-nil, zero value otherwise.

### GetMachineIdOk

`func (o *DeploymentMachine) GetMachineIdOk() (*string, bool)`

GetMachineIdOk returns a tuple with the MachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineId

`func (o *DeploymentMachine) SetMachineId(v string)`

SetMachineId sets MachineId field to given value.

### HasMachineId

`func (o *DeploymentMachine) HasMachineId() bool`

HasMachineId returns a boolean if a field has been set.

### SetMachineIdNil

`func (o *DeploymentMachine) SetMachineIdNil(b bool)`

 SetMachineIdNil sets the value for MachineId to be an explicit nil

### UnsetMachineId
`func (o *DeploymentMachine) UnsetMachineId()`

UnsetMachineId ensures that no value is present for MachineId, not even an explicit nil
### GetMachineName

`func (o *DeploymentMachine) GetMachineName() string`

GetMachineName returns the MachineName field if non-nil, zero value otherwise.

### GetMachineNameOk

`func (o *DeploymentMachine) GetMachineNameOk() (*string, bool)`

GetMachineNameOk returns a tuple with the MachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineName

`func (o *DeploymentMachine) SetMachineName(v string)`

SetMachineName sets MachineName field to given value.

### HasMachineName

`func (o *DeploymentMachine) HasMachineName() bool`

HasMachineName returns a boolean if a field has been set.

### SetMachineNameNil

`func (o *DeploymentMachine) SetMachineNameNil(b bool)`

 SetMachineNameNil sets the value for MachineName to be an explicit nil

### UnsetMachineName
`func (o *DeploymentMachine) UnsetMachineName()`

UnsetMachineName ensures that no value is present for MachineName, not even an explicit nil
### GetRegistrationState

`func (o *DeploymentMachine) GetRegistrationState() RegistrationState`

GetRegistrationState returns the RegistrationState field if non-nil, zero value otherwise.

### GetRegistrationStateOk

`func (o *DeploymentMachine) GetRegistrationStateOk() (*RegistrationState, bool)`

GetRegistrationStateOk returns a tuple with the RegistrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationState

`func (o *DeploymentMachine) SetRegistrationState(v RegistrationState)`

SetRegistrationState sets RegistrationState field to given value.

### HasRegistrationState

`func (o *DeploymentMachine) HasRegistrationState() bool`

HasRegistrationState returns a boolean if a field has been set.

### SetRegistrationStateNil

`func (o *DeploymentMachine) SetRegistrationStateNil(b bool)`

 SetRegistrationStateNil sets the value for RegistrationState to be an explicit nil

### UnsetRegistrationState
`func (o *DeploymentMachine) UnsetRegistrationState()`

UnsetRegistrationState ensures that no value is present for RegistrationState, not even an explicit nil
### GetSessionState

`func (o *DeploymentMachine) GetSessionState() SessionState`

GetSessionState returns the SessionState field if non-nil, zero value otherwise.

### GetSessionStateOk

`func (o *DeploymentMachine) GetSessionStateOk() (*SessionState, bool)`

GetSessionStateOk returns a tuple with the SessionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionState

`func (o *DeploymentMachine) SetSessionState(v SessionState)`

SetSessionState sets SessionState field to given value.

### HasSessionState

`func (o *DeploymentMachine) HasSessionState() bool`

HasSessionState returns a boolean if a field has been set.

### SetSessionStateNil

`func (o *DeploymentMachine) SetSessionStateNil(b bool)`

 SetSessionStateNil sets the value for SessionState to be an explicit nil

### UnsetSessionState
`func (o *DeploymentMachine) UnsetSessionState()`

UnsetSessionState ensures that no value is present for SessionState, not even an explicit nil
### GetSessionCount

`func (o *DeploymentMachine) GetSessionCount() int32`

GetSessionCount returns the SessionCount field if non-nil, zero value otherwise.

### GetSessionCountOk

`func (o *DeploymentMachine) GetSessionCountOk() (*int32, bool)`

GetSessionCountOk returns a tuple with the SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCount

`func (o *DeploymentMachine) SetSessionCount(v int32)`

SetSessionCount sets SessionCount field to given value.

### HasSessionCount

`func (o *DeploymentMachine) HasSessionCount() bool`

HasSessionCount returns a boolean if a field has been set.

### SetSessionCountNil

`func (o *DeploymentMachine) SetSessionCountNil(b bool)`

 SetSessionCountNil sets the value for SessionCount to be an explicit nil

### UnsetSessionCount
`func (o *DeploymentMachine) UnsetSessionCount()`

UnsetSessionCount ensures that no value is present for SessionCount, not even an explicit nil
### GetMaintenanceMode

`func (o *DeploymentMachine) GetMaintenanceMode() bool`

GetMaintenanceMode returns the MaintenanceMode field if non-nil, zero value otherwise.

### GetMaintenanceModeOk

`func (o *DeploymentMachine) GetMaintenanceModeOk() (*bool, bool)`

GetMaintenanceModeOk returns a tuple with the MaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceMode

`func (o *DeploymentMachine) SetMaintenanceMode(v bool)`

SetMaintenanceMode sets MaintenanceMode field to given value.

### HasMaintenanceMode

`func (o *DeploymentMachine) HasMaintenanceMode() bool`

HasMaintenanceMode returns a boolean if a field has been set.

### SetMaintenanceModeNil

`func (o *DeploymentMachine) SetMaintenanceModeNil(b bool)`

 SetMaintenanceModeNil sets the value for MaintenanceMode to be an explicit nil

### UnsetMaintenanceMode
`func (o *DeploymentMachine) UnsetMaintenanceMode()`

UnsetMaintenanceMode ensures that no value is present for MaintenanceMode, not even an explicit nil
### GetAssociatedUsers

`func (o *DeploymentMachine) GetAssociatedUsers() []string`

GetAssociatedUsers returns the AssociatedUsers field if non-nil, zero value otherwise.

### GetAssociatedUsersOk

`func (o *DeploymentMachine) GetAssociatedUsersOk() (*[]string, bool)`

GetAssociatedUsersOk returns a tuple with the AssociatedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedUsers

`func (o *DeploymentMachine) SetAssociatedUsers(v []string)`

SetAssociatedUsers sets AssociatedUsers field to given value.

### HasAssociatedUsers

`func (o *DeploymentMachine) HasAssociatedUsers() bool`

HasAssociatedUsers returns a boolean if a field has been set.

### SetAssociatedUsersNil

`func (o *DeploymentMachine) SetAssociatedUsersNil(b bool)`

 SetAssociatedUsersNil sets the value for AssociatedUsers to be an explicit nil

### UnsetAssociatedUsers
`func (o *DeploymentMachine) UnsetAssociatedUsers()`

UnsetAssociatedUsers ensures that no value is present for AssociatedUsers, not even an explicit nil
### GetBrokerMachineCatalogId

`func (o *DeploymentMachine) GetBrokerMachineCatalogId() string`

GetBrokerMachineCatalogId returns the BrokerMachineCatalogId field if non-nil, zero value otherwise.

### GetBrokerMachineCatalogIdOk

`func (o *DeploymentMachine) GetBrokerMachineCatalogIdOk() (*string, bool)`

GetBrokerMachineCatalogIdOk returns a tuple with the BrokerMachineCatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerMachineCatalogId

`func (o *DeploymentMachine) SetBrokerMachineCatalogId(v string)`

SetBrokerMachineCatalogId sets BrokerMachineCatalogId field to given value.

### HasBrokerMachineCatalogId

`func (o *DeploymentMachine) HasBrokerMachineCatalogId() bool`

HasBrokerMachineCatalogId returns a boolean if a field has been set.

### SetBrokerMachineCatalogIdNil

`func (o *DeploymentMachine) SetBrokerMachineCatalogIdNil(b bool)`

 SetBrokerMachineCatalogIdNil sets the value for BrokerMachineCatalogId to be an explicit nil

### UnsetBrokerMachineCatalogId
`func (o *DeploymentMachine) UnsetBrokerMachineCatalogId()`

UnsetBrokerMachineCatalogId ensures that no value is present for BrokerMachineCatalogId, not even an explicit nil
### GetBrokerDeliveryGroupId

`func (o *DeploymentMachine) GetBrokerDeliveryGroupId() string`

GetBrokerDeliveryGroupId returns the BrokerDeliveryGroupId field if non-nil, zero value otherwise.

### GetBrokerDeliveryGroupIdOk

`func (o *DeploymentMachine) GetBrokerDeliveryGroupIdOk() (*string, bool)`

GetBrokerDeliveryGroupIdOk returns a tuple with the BrokerDeliveryGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerDeliveryGroupId

`func (o *DeploymentMachine) SetBrokerDeliveryGroupId(v string)`

SetBrokerDeliveryGroupId sets BrokerDeliveryGroupId field to given value.

### HasBrokerDeliveryGroupId

`func (o *DeploymentMachine) HasBrokerDeliveryGroupId() bool`

HasBrokerDeliveryGroupId returns a boolean if a field has been set.

### SetBrokerDeliveryGroupIdNil

`func (o *DeploymentMachine) SetBrokerDeliveryGroupIdNil(b bool)`

 SetBrokerDeliveryGroupIdNil sets the value for BrokerDeliveryGroupId to be an explicit nil

### UnsetBrokerDeliveryGroupId
`func (o *DeploymentMachine) UnsetBrokerDeliveryGroupId()`

UnsetBrokerDeliveryGroupId ensures that no value is present for BrokerDeliveryGroupId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


