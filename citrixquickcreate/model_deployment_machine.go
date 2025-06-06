/*
Quick Create Service for DaaS Quick Deploy with AWS WorkSpaces Core - Tech Preview

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: DaaSTechPreview
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citrixquickcreate

import (
	"encoding/json"
)

// checks if the DeploymentMachine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentMachine{}

// DeploymentMachine Deployment
type DeploymentMachine struct {
	// The type of provider associated with the account
	AccountType AccountType `json:"accountType"`
	// Deployment Id
	DeploymentId NullableString `json:"deploymentId,omitempty"`
	// MachineId on the Broker
	BrokerMachineId NullableString `json:"brokerMachineId,omitempty"`
	// Task Type
	TaskType NullableTaskType `json:"taskType,omitempty"`
	// Task Id
	TaskId NullableString `json:"taskId,omitempty"`
	// Hosting Connection Id
	ConnectionId NullableString `json:"connectionId,omitempty"`
	// Image Id
	ImageId NullableString `json:"imageId,omitempty"`
	// Machine Id
	MachineId NullableString `json:"machineId,omitempty"`
	// Name of the machine
	MachineName NullableString `json:"machineName,omitempty"`
	// Registration state of the machine
	RegistrationState NullableRegistrationState `json:"registrationState,omitempty"`
	// State of active session on machine
	SessionState NullableSessionState `json:"sessionState,omitempty"`
	// Count of active session on machine
	SessionCount NullableInt32 `json:"sessionCount,omitempty"`
	// Indicates if the machine is in Maintenance Mode
	MaintenanceMode NullableBool `json:"maintenanceMode,omitempty"`
	// List of users that are associated with the machine
	AssociatedUsers []string `json:"associatedUsers,omitempty"`
	// Id for the machine catalog of the machine, could be null if machine not in machine catalog
	BrokerMachineCatalogId NullableString `json:"brokerMachineCatalogId,omitempty"`
	// Id for the delivery group of the machine, could be null if machine not in delivery group
	BrokerDeliveryGroupId NullableString `json:"brokerDeliveryGroupId,omitempty"`
}

type _DeploymentMachine DeploymentMachine

// NewDeploymentMachine instantiates a new DeploymentMachine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentMachine(accountType AccountType) *DeploymentMachine {
	this := DeploymentMachine{}
	this.AccountType = accountType
	return &this
}

// NewDeploymentMachineWithDefaults instantiates a new DeploymentMachine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentMachineWithDefaults() *DeploymentMachine {
	this := DeploymentMachine{}
	return &this
}

// GetAccountType returns the AccountType field value
func (o *DeploymentMachine) GetAccountType() AccountType {
	if o == nil {
		var ret AccountType
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *DeploymentMachine) GetAccountTypeOk() (*AccountType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *DeploymentMachine) SetAccountType(v AccountType) {
	o.AccountType = v
}

// GetDeploymentId returns the DeploymentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentMachine) GetDeploymentId() string {
	if o == nil || IsNil(o.DeploymentId.Get()) {
		var ret string
		return ret
	}
	return *o.DeploymentId.Get()
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentMachine) GetDeploymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeploymentId.Get(), o.DeploymentId.IsSet()
}

// HasDeploymentId returns a boolean if a field has been set.
func (o *DeploymentMachine) HasDeploymentId() bool {
	if o != nil && o.DeploymentId.IsSet() {
		return true
	}

	return false
}

// SetDeploymentId gets a reference to the given NullableString and assigns it to the DeploymentId field.
func (o *DeploymentMachine) SetDeploymentId(v string) {
	o.DeploymentId.Set(&v)
}

// SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil
func (o *DeploymentMachine) SetDeploymentIdNil() {
	o.DeploymentId.Set(nil)
}

// UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
func (o *DeploymentMachine) UnsetDeploymentId() {
	o.DeploymentId.Unset()
}

// GetBrokerMachineId returns the BrokerMachineId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentMachine) GetBrokerMachineId() string {
	if o == nil || IsNil(o.BrokerMachineId.Get()) {
		var ret string
		return ret
	}
	return *o.BrokerMachineId.Get()
}

// GetBrokerMachineIdOk returns a tuple with the BrokerMachineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentMachine) GetBrokerMachineIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrokerMachineId.Get(), o.BrokerMachineId.IsSet()
}

// HasBrokerMachineId returns a boolean if a field has been set.
func (o *DeploymentMachine) HasBrokerMachineId() bool {
	if o != nil && o.BrokerMachineId.IsSet() {
		return true
	}

	return false
}

// SetBrokerMachineId gets a reference to the given NullableString and assigns it to the BrokerMachineId field.
func (o *DeploymentMachine) SetBrokerMachineId(v string) {
	o.BrokerMachineId.Set(&v)
}

// SetBrokerMachineIdNil sets the value for BrokerMachineId to be an explicit nil
func (o *DeploymentMachine) SetBrokerMachineIdNil() {
	o.BrokerMachineId.Set(nil)
}

// UnsetBrokerMachineId ensures that no value is present for BrokerMachineId, not even an explicit nil
func (o *DeploymentMachine) UnsetBrokerMachineId() {
	o.BrokerMachineId.Unset()
}

// GetTaskType returns the TaskType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentMachine) GetTaskType() TaskType {
	if o == nil || IsNil(o.TaskType.Get()) {
		var ret TaskType
		return ret
	}
	return *o.TaskType.Get()
}

// GetTaskTypeOk returns a tuple with the TaskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentMachine) GetTaskTypeOk() (*TaskType, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskType.Get(), o.TaskType.IsSet()
}

// HasTaskType returns a boolean if a field has been set.
func (o *DeploymentMachine) HasTaskType() bool {
	if o != nil && o.TaskType.IsSet() {
		return true
	}

	return false
}

// SetTaskType gets a reference to the given NullableTaskType and assigns it to the TaskType field.
func (o *DeploymentMachine) SetTaskType(v TaskType) {
	o.TaskType.Set(&v)
}

// SetTaskTypeNil sets the value for TaskType to be an explicit nil
func (o *DeploymentMachine) SetTaskTypeNil() {
	o.TaskType.Set(nil)
}

// UnsetTaskType ensures that no value is present for TaskType, not even an explicit nil
func (o *DeploymentMachine) UnsetTaskType() {
	o.TaskType.Unset()
}

// GetTaskId returns the TaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentMachine) GetTaskId() string {
	if o == nil || IsNil(o.TaskId.Get()) {
		var ret string
		return ret
	}
	return *o.TaskId.Get()
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentMachine) GetTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskId.Get(), o.TaskId.IsSet()
}

// HasTaskId returns a boolean if a field has been set.
func (o *DeploymentMachine) HasTaskId() bool {
	if o != nil && o.TaskId.IsSet() {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given NullableString and assigns it to the TaskId field.
func (o *DeploymentMachine) SetTaskId(v string) {
	o.TaskId.Set(&v)
}

// SetTaskIdNil sets the value for TaskId to be an explicit nil
func (o *DeploymentMachine) SetTaskIdNil() {
	o.TaskId.Set(nil)
}

// UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
func (o *DeploymentMachine) UnsetTaskId() {
	o.TaskId.Unset()
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentMachine) GetConnectionId() string {
	if o == nil || IsNil(o.ConnectionId.Get()) {
		var ret string
		return ret
	}
	return *o.ConnectionId.Get()
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentMachine) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectionId.Get(), o.ConnectionId.IsSet()
}

// HasConnectionId returns a boolean if a field has been set.
func (o *DeploymentMachine) HasConnectionId() bool {
	if o != nil && o.ConnectionId.IsSet() {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given NullableString and assigns it to the ConnectionId field.
func (o *DeploymentMachine) SetConnectionId(v string) {
	o.ConnectionId.Set(&v)
}

// SetConnectionIdNil sets the value for ConnectionId to be an explicit nil
func (o *DeploymentMachine) SetConnectionIdNil() {
	o.ConnectionId.Set(nil)
}

// UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
func (o *DeploymentMachine) UnsetConnectionId() {
	o.ConnectionId.Unset()
}

// GetImageId returns the ImageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentMachine) GetImageId() string {
	if o == nil || IsNil(o.ImageId.Get()) {
		var ret string
		return ret
	}
	return *o.ImageId.Get()
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentMachine) GetImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageId.Get(), o.ImageId.IsSet()
}

// HasImageId returns a boolean if a field has been set.
func (o *DeploymentMachine) HasImageId() bool {
	if o != nil && o.ImageId.IsSet() {
		return true
	}

	return false
}

// SetImageId gets a reference to the given NullableString and assigns it to the ImageId field.
func (o *DeploymentMachine) SetImageId(v string) {
	o.ImageId.Set(&v)
}

// SetImageIdNil sets the value for ImageId to be an explicit nil
func (o *DeploymentMachine) SetImageIdNil() {
	o.ImageId.Set(nil)
}

// UnsetImageId ensures that no value is present for ImageId, not even an explicit nil
func (o *DeploymentMachine) UnsetImageId() {
	o.ImageId.Unset()
}

// GetMachineId returns the MachineId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentMachine) GetMachineId() string {
	if o == nil || IsNil(o.MachineId.Get()) {
		var ret string
		return ret
	}
	return *o.MachineId.Get()
}

// GetMachineIdOk returns a tuple with the MachineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentMachine) GetMachineIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MachineId.Get(), o.MachineId.IsSet()
}

// HasMachineId returns a boolean if a field has been set.
func (o *DeploymentMachine) HasMachineId() bool {
	if o != nil && o.MachineId.IsSet() {
		return true
	}

	return false
}

// SetMachineId gets a reference to the given NullableString and assigns it to the MachineId field.
func (o *DeploymentMachine) SetMachineId(v string) {
	o.MachineId.Set(&v)
}

// SetMachineIdNil sets the value for MachineId to be an explicit nil
func (o *DeploymentMachine) SetMachineIdNil() {
	o.MachineId.Set(nil)
}

// UnsetMachineId ensures that no value is present for MachineId, not even an explicit nil
func (o *DeploymentMachine) UnsetMachineId() {
	o.MachineId.Unset()
}

// GetMachineName returns the MachineName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentMachine) GetMachineName() string {
	if o == nil || IsNil(o.MachineName.Get()) {
		var ret string
		return ret
	}
	return *o.MachineName.Get()
}

// GetMachineNameOk returns a tuple with the MachineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentMachine) GetMachineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MachineName.Get(), o.MachineName.IsSet()
}

// HasMachineName returns a boolean if a field has been set.
func (o *DeploymentMachine) HasMachineName() bool {
	if o != nil && o.MachineName.IsSet() {
		return true
	}

	return false
}

// SetMachineName gets a reference to the given NullableString and assigns it to the MachineName field.
func (o *DeploymentMachine) SetMachineName(v string) {
	o.MachineName.Set(&v)
}

// SetMachineNameNil sets the value for MachineName to be an explicit nil
func (o *DeploymentMachine) SetMachineNameNil() {
	o.MachineName.Set(nil)
}

// UnsetMachineName ensures that no value is present for MachineName, not even an explicit nil
func (o *DeploymentMachine) UnsetMachineName() {
	o.MachineName.Unset()
}

// GetRegistrationState returns the RegistrationState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentMachine) GetRegistrationState() RegistrationState {
	if o == nil || IsNil(o.RegistrationState.Get()) {
		var ret RegistrationState
		return ret
	}
	return *o.RegistrationState.Get()
}

// GetRegistrationStateOk returns a tuple with the RegistrationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentMachine) GetRegistrationStateOk() (*RegistrationState, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegistrationState.Get(), o.RegistrationState.IsSet()
}

// HasRegistrationState returns a boolean if a field has been set.
func (o *DeploymentMachine) HasRegistrationState() bool {
	if o != nil && o.RegistrationState.IsSet() {
		return true
	}

	return false
}

// SetRegistrationState gets a reference to the given NullableRegistrationState and assigns it to the RegistrationState field.
func (o *DeploymentMachine) SetRegistrationState(v RegistrationState) {
	o.RegistrationState.Set(&v)
}

// SetRegistrationStateNil sets the value for RegistrationState to be an explicit nil
func (o *DeploymentMachine) SetRegistrationStateNil() {
	o.RegistrationState.Set(nil)
}

// UnsetRegistrationState ensures that no value is present for RegistrationState, not even an explicit nil
func (o *DeploymentMachine) UnsetRegistrationState() {
	o.RegistrationState.Unset()
}

// GetSessionState returns the SessionState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentMachine) GetSessionState() SessionState {
	if o == nil || IsNil(o.SessionState.Get()) {
		var ret SessionState
		return ret
	}
	return *o.SessionState.Get()
}

// GetSessionStateOk returns a tuple with the SessionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentMachine) GetSessionStateOk() (*SessionState, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionState.Get(), o.SessionState.IsSet()
}

// HasSessionState returns a boolean if a field has been set.
func (o *DeploymentMachine) HasSessionState() bool {
	if o != nil && o.SessionState.IsSet() {
		return true
	}

	return false
}

// SetSessionState gets a reference to the given NullableSessionState and assigns it to the SessionState field.
func (o *DeploymentMachine) SetSessionState(v SessionState) {
	o.SessionState.Set(&v)
}

// SetSessionStateNil sets the value for SessionState to be an explicit nil
func (o *DeploymentMachine) SetSessionStateNil() {
	o.SessionState.Set(nil)
}

// UnsetSessionState ensures that no value is present for SessionState, not even an explicit nil
func (o *DeploymentMachine) UnsetSessionState() {
	o.SessionState.Unset()
}

// GetSessionCount returns the SessionCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentMachine) GetSessionCount() int32 {
	if o == nil || IsNil(o.SessionCount.Get()) {
		var ret int32
		return ret
	}
	return *o.SessionCount.Get()
}

// GetSessionCountOk returns a tuple with the SessionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentMachine) GetSessionCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionCount.Get(), o.SessionCount.IsSet()
}

// HasSessionCount returns a boolean if a field has been set.
func (o *DeploymentMachine) HasSessionCount() bool {
	if o != nil && o.SessionCount.IsSet() {
		return true
	}

	return false
}

// SetSessionCount gets a reference to the given NullableInt32 and assigns it to the SessionCount field.
func (o *DeploymentMachine) SetSessionCount(v int32) {
	o.SessionCount.Set(&v)
}

// SetSessionCountNil sets the value for SessionCount to be an explicit nil
func (o *DeploymentMachine) SetSessionCountNil() {
	o.SessionCount.Set(nil)
}

// UnsetSessionCount ensures that no value is present for SessionCount, not even an explicit nil
func (o *DeploymentMachine) UnsetSessionCount() {
	o.SessionCount.Unset()
}

// GetMaintenanceMode returns the MaintenanceMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentMachine) GetMaintenanceMode() bool {
	if o == nil || IsNil(o.MaintenanceMode.Get()) {
		var ret bool
		return ret
	}
	return *o.MaintenanceMode.Get()
}

// GetMaintenanceModeOk returns a tuple with the MaintenanceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentMachine) GetMaintenanceModeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaintenanceMode.Get(), o.MaintenanceMode.IsSet()
}

// HasMaintenanceMode returns a boolean if a field has been set.
func (o *DeploymentMachine) HasMaintenanceMode() bool {
	if o != nil && o.MaintenanceMode.IsSet() {
		return true
	}

	return false
}

// SetMaintenanceMode gets a reference to the given NullableBool and assigns it to the MaintenanceMode field.
func (o *DeploymentMachine) SetMaintenanceMode(v bool) {
	o.MaintenanceMode.Set(&v)
}

// SetMaintenanceModeNil sets the value for MaintenanceMode to be an explicit nil
func (o *DeploymentMachine) SetMaintenanceModeNil() {
	o.MaintenanceMode.Set(nil)
}

// UnsetMaintenanceMode ensures that no value is present for MaintenanceMode, not even an explicit nil
func (o *DeploymentMachine) UnsetMaintenanceMode() {
	o.MaintenanceMode.Unset()
}

// GetAssociatedUsers returns the AssociatedUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentMachine) GetAssociatedUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AssociatedUsers
}

// GetAssociatedUsersOk returns a tuple with the AssociatedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentMachine) GetAssociatedUsersOk() ([]string, bool) {
	if o == nil || IsNil(o.AssociatedUsers) {
		return nil, false
	}
	return o.AssociatedUsers, true
}

// HasAssociatedUsers returns a boolean if a field has been set.
func (o *DeploymentMachine) HasAssociatedUsers() bool {
	if o != nil && !IsNil(o.AssociatedUsers) {
		return true
	}

	return false
}

// SetAssociatedUsers gets a reference to the given []string and assigns it to the AssociatedUsers field.
func (o *DeploymentMachine) SetAssociatedUsers(v []string) {
	o.AssociatedUsers = v
}

// GetBrokerMachineCatalogId returns the BrokerMachineCatalogId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentMachine) GetBrokerMachineCatalogId() string {
	if o == nil || IsNil(o.BrokerMachineCatalogId.Get()) {
		var ret string
		return ret
	}
	return *o.BrokerMachineCatalogId.Get()
}

// GetBrokerMachineCatalogIdOk returns a tuple with the BrokerMachineCatalogId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentMachine) GetBrokerMachineCatalogIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrokerMachineCatalogId.Get(), o.BrokerMachineCatalogId.IsSet()
}

// HasBrokerMachineCatalogId returns a boolean if a field has been set.
func (o *DeploymentMachine) HasBrokerMachineCatalogId() bool {
	if o != nil && o.BrokerMachineCatalogId.IsSet() {
		return true
	}

	return false
}

// SetBrokerMachineCatalogId gets a reference to the given NullableString and assigns it to the BrokerMachineCatalogId field.
func (o *DeploymentMachine) SetBrokerMachineCatalogId(v string) {
	o.BrokerMachineCatalogId.Set(&v)
}

// SetBrokerMachineCatalogIdNil sets the value for BrokerMachineCatalogId to be an explicit nil
func (o *DeploymentMachine) SetBrokerMachineCatalogIdNil() {
	o.BrokerMachineCatalogId.Set(nil)
}

// UnsetBrokerMachineCatalogId ensures that no value is present for BrokerMachineCatalogId, not even an explicit nil
func (o *DeploymentMachine) UnsetBrokerMachineCatalogId() {
	o.BrokerMachineCatalogId.Unset()
}

// GetBrokerDeliveryGroupId returns the BrokerDeliveryGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeploymentMachine) GetBrokerDeliveryGroupId() string {
	if o == nil || IsNil(o.BrokerDeliveryGroupId.Get()) {
		var ret string
		return ret
	}
	return *o.BrokerDeliveryGroupId.Get()
}

// GetBrokerDeliveryGroupIdOk returns a tuple with the BrokerDeliveryGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeploymentMachine) GetBrokerDeliveryGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrokerDeliveryGroupId.Get(), o.BrokerDeliveryGroupId.IsSet()
}

// HasBrokerDeliveryGroupId returns a boolean if a field has been set.
func (o *DeploymentMachine) HasBrokerDeliveryGroupId() bool {
	if o != nil && o.BrokerDeliveryGroupId.IsSet() {
		return true
	}

	return false
}

// SetBrokerDeliveryGroupId gets a reference to the given NullableString and assigns it to the BrokerDeliveryGroupId field.
func (o *DeploymentMachine) SetBrokerDeliveryGroupId(v string) {
	o.BrokerDeliveryGroupId.Set(&v)
}

// SetBrokerDeliveryGroupIdNil sets the value for BrokerDeliveryGroupId to be an explicit nil
func (o *DeploymentMachine) SetBrokerDeliveryGroupIdNil() {
	o.BrokerDeliveryGroupId.Set(nil)
}

// UnsetBrokerDeliveryGroupId ensures that no value is present for BrokerDeliveryGroupId, not even an explicit nil
func (o *DeploymentMachine) UnsetBrokerDeliveryGroupId() {
	o.BrokerDeliveryGroupId.Unset()
}

func (o DeploymentMachine) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentMachine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountType"] = o.AccountType
	if o.DeploymentId.IsSet() {
		toSerialize["deploymentId"] = o.DeploymentId.Get()
	}
	if o.BrokerMachineId.IsSet() {
		toSerialize["brokerMachineId"] = o.BrokerMachineId.Get()
	}
	if o.TaskType.IsSet() {
		toSerialize["taskType"] = o.TaskType.Get()
	}
	if o.TaskId.IsSet() {
		toSerialize["taskId"] = o.TaskId.Get()
	}
	if o.ConnectionId.IsSet() {
		toSerialize["connectionId"] = o.ConnectionId.Get()
	}
	if o.ImageId.IsSet() {
		toSerialize["imageId"] = o.ImageId.Get()
	}
	if o.MachineId.IsSet() {
		toSerialize["machineId"] = o.MachineId.Get()
	}
	if o.MachineName.IsSet() {
		toSerialize["machineName"] = o.MachineName.Get()
	}
	if o.RegistrationState.IsSet() {
		toSerialize["registrationState"] = o.RegistrationState.Get()
	}
	if o.SessionState.IsSet() {
		toSerialize["sessionState"] = o.SessionState.Get()
	}
	if o.SessionCount.IsSet() {
		toSerialize["sessionCount"] = o.SessionCount.Get()
	}
	if o.MaintenanceMode.IsSet() {
		toSerialize["maintenanceMode"] = o.MaintenanceMode.Get()
	}
	if o.AssociatedUsers != nil {
		toSerialize["associatedUsers"] = o.AssociatedUsers
	}
	if o.BrokerMachineCatalogId.IsSet() {
		toSerialize["brokerMachineCatalogId"] = o.BrokerMachineCatalogId.Get()
	}
	if o.BrokerDeliveryGroupId.IsSet() {
		toSerialize["brokerDeliveryGroupId"] = o.BrokerDeliveryGroupId.Get()
	}
	return toSerialize, nil
}

type NullableDeploymentMachine struct {
	value *DeploymentMachine
	isSet bool
}

func (v NullableDeploymentMachine) Get() *DeploymentMachine {
	return v.value
}

func (v *NullableDeploymentMachine) Set(val *DeploymentMachine) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentMachine) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentMachine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentMachine(val *DeploymentMachine) *NullableDeploymentMachine {
	return &NullableDeploymentMachine{value: val, isSet: true}
}

func (v NullableDeploymentMachine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentMachine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
