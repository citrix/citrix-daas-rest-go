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

// checks if the Deployment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Deployment{}

// Deployment Deployment
type Deployment struct {
	// The type of provider associated with the account
	AccountType AccountType `json:"accountType"`
	// Deployment Id
	DeploymentId NullableString `json:"deploymentId,omitempty"`
	// Deployment Name
	DeploymentName NullableString `json:"deploymentName,omitempty"`
	// Account Id
	AccountId NullableString `json:"accountId,omitempty"`
	// Account name
	AccountName NullableString `json:"accountName,omitempty"`
	// Connection Id
	ConnectionId NullableString `json:"connectionId,omitempty"`
	// Connection Name
	ConnectionName NullableString `json:"connectionName,omitempty"`
	// Deployment State
	DeploymentState NullableDeploymentState `json:"deploymentState,omitempty"`
	// The number of users in this deployment
	UserCount *int32 `json:"userCount,omitempty"`
	// Error message associated with the deployment
	ErrorMessage NullableString `json:"errorMessage,omitempty"`
	// Warnings and errors associated with the deployment
	Warnings []DeploymentWarning `json:"warnings,omitempty"`
	// Tasks currently being performed on the deployment
	ActiveTasks []GetTaskAsync200Response `json:"activeTasks,omitempty"`
	// Id for the machine catalog of the deployment, could be null
	BrokerMachineCatalogId NullableString `json:"brokerMachineCatalogId,omitempty"`
	// Id for the delivery group of the deployment, could be null
	BrokerDeliveryGroupId NullableString `json:"brokerDeliveryGroupId,omitempty"`
	// Indicates whether the deployment is managed by Citrix
	CitrixManaged NullableBool `json:"citrixManaged,omitempty"`
	// Power management scale settings
	ScaleSettings NullableScaleSettings `json:"scaleSettings,omitempty"`
}

type _Deployment Deployment

// NewDeployment instantiates a new Deployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeployment(accountType AccountType) *Deployment {
	this := Deployment{}
	this.AccountType = accountType
	return &this
}

// NewDeploymentWithDefaults instantiates a new Deployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentWithDefaults() *Deployment {
	this := Deployment{}
	return &this
}

// GetAccountType returns the AccountType field value
func (o *Deployment) GetAccountType() AccountType {
	if o == nil {
		var ret AccountType
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *Deployment) GetAccountTypeOk() (*AccountType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *Deployment) SetAccountType(v AccountType) {
	o.AccountType = v
}

// GetDeploymentId returns the DeploymentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Deployment) GetDeploymentId() string {
	if o == nil || IsNil(o.DeploymentId.Get()) {
		var ret string
		return ret
	}
	return *o.DeploymentId.Get()
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Deployment) GetDeploymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeploymentId.Get(), o.DeploymentId.IsSet()
}

// HasDeploymentId returns a boolean if a field has been set.
func (o *Deployment) HasDeploymentId() bool {
	if o != nil && o.DeploymentId.IsSet() {
		return true
	}

	return false
}

// SetDeploymentId gets a reference to the given NullableString and assigns it to the DeploymentId field.
func (o *Deployment) SetDeploymentId(v string) {
	o.DeploymentId.Set(&v)
}

// SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil
func (o *Deployment) SetDeploymentIdNil() {
	o.DeploymentId.Set(nil)
}

// UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
func (o *Deployment) UnsetDeploymentId() {
	o.DeploymentId.Unset()
}

// GetDeploymentName returns the DeploymentName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Deployment) GetDeploymentName() string {
	if o == nil || IsNil(o.DeploymentName.Get()) {
		var ret string
		return ret
	}
	return *o.DeploymentName.Get()
}

// GetDeploymentNameOk returns a tuple with the DeploymentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Deployment) GetDeploymentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeploymentName.Get(), o.DeploymentName.IsSet()
}

// HasDeploymentName returns a boolean if a field has been set.
func (o *Deployment) HasDeploymentName() bool {
	if o != nil && o.DeploymentName.IsSet() {
		return true
	}

	return false
}

// SetDeploymentName gets a reference to the given NullableString and assigns it to the DeploymentName field.
func (o *Deployment) SetDeploymentName(v string) {
	o.DeploymentName.Set(&v)
}

// SetDeploymentNameNil sets the value for DeploymentName to be an explicit nil
func (o *Deployment) SetDeploymentNameNil() {
	o.DeploymentName.Set(nil)
}

// UnsetDeploymentName ensures that no value is present for DeploymentName, not even an explicit nil
func (o *Deployment) UnsetDeploymentName() {
	o.DeploymentName.Unset()
}

// GetAccountId returns the AccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Deployment) GetAccountId() string {
	if o == nil || IsNil(o.AccountId.Get()) {
		var ret string
		return ret
	}
	return *o.AccountId.Get()
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Deployment) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountId.Get(), o.AccountId.IsSet()
}

// HasAccountId returns a boolean if a field has been set.
func (o *Deployment) HasAccountId() bool {
	if o != nil && o.AccountId.IsSet() {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given NullableString and assigns it to the AccountId field.
func (o *Deployment) SetAccountId(v string) {
	o.AccountId.Set(&v)
}

// SetAccountIdNil sets the value for AccountId to be an explicit nil
func (o *Deployment) SetAccountIdNil() {
	o.AccountId.Set(nil)
}

// UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
func (o *Deployment) UnsetAccountId() {
	o.AccountId.Unset()
}

// GetAccountName returns the AccountName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Deployment) GetAccountName() string {
	if o == nil || IsNil(o.AccountName.Get()) {
		var ret string
		return ret
	}
	return *o.AccountName.Get()
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Deployment) GetAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountName.Get(), o.AccountName.IsSet()
}

// HasAccountName returns a boolean if a field has been set.
func (o *Deployment) HasAccountName() bool {
	if o != nil && o.AccountName.IsSet() {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given NullableString and assigns it to the AccountName field.
func (o *Deployment) SetAccountName(v string) {
	o.AccountName.Set(&v)
}

// SetAccountNameNil sets the value for AccountName to be an explicit nil
func (o *Deployment) SetAccountNameNil() {
	o.AccountName.Set(nil)
}

// UnsetAccountName ensures that no value is present for AccountName, not even an explicit nil
func (o *Deployment) UnsetAccountName() {
	o.AccountName.Unset()
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Deployment) GetConnectionId() string {
	if o == nil || IsNil(o.ConnectionId.Get()) {
		var ret string
		return ret
	}
	return *o.ConnectionId.Get()
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Deployment) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectionId.Get(), o.ConnectionId.IsSet()
}

// HasConnectionId returns a boolean if a field has been set.
func (o *Deployment) HasConnectionId() bool {
	if o != nil && o.ConnectionId.IsSet() {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given NullableString and assigns it to the ConnectionId field.
func (o *Deployment) SetConnectionId(v string) {
	o.ConnectionId.Set(&v)
}

// SetConnectionIdNil sets the value for ConnectionId to be an explicit nil
func (o *Deployment) SetConnectionIdNil() {
	o.ConnectionId.Set(nil)
}

// UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
func (o *Deployment) UnsetConnectionId() {
	o.ConnectionId.Unset()
}

// GetConnectionName returns the ConnectionName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Deployment) GetConnectionName() string {
	if o == nil || IsNil(o.ConnectionName.Get()) {
		var ret string
		return ret
	}
	return *o.ConnectionName.Get()
}

// GetConnectionNameOk returns a tuple with the ConnectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Deployment) GetConnectionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectionName.Get(), o.ConnectionName.IsSet()
}

// HasConnectionName returns a boolean if a field has been set.
func (o *Deployment) HasConnectionName() bool {
	if o != nil && o.ConnectionName.IsSet() {
		return true
	}

	return false
}

// SetConnectionName gets a reference to the given NullableString and assigns it to the ConnectionName field.
func (o *Deployment) SetConnectionName(v string) {
	o.ConnectionName.Set(&v)
}

// SetConnectionNameNil sets the value for ConnectionName to be an explicit nil
func (o *Deployment) SetConnectionNameNil() {
	o.ConnectionName.Set(nil)
}

// UnsetConnectionName ensures that no value is present for ConnectionName, not even an explicit nil
func (o *Deployment) UnsetConnectionName() {
	o.ConnectionName.Unset()
}

// GetDeploymentState returns the DeploymentState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Deployment) GetDeploymentState() DeploymentState {
	if o == nil || IsNil(o.DeploymentState.Get()) {
		var ret DeploymentState
		return ret
	}
	return *o.DeploymentState.Get()
}

// GetDeploymentStateOk returns a tuple with the DeploymentState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Deployment) GetDeploymentStateOk() (*DeploymentState, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeploymentState.Get(), o.DeploymentState.IsSet()
}

// HasDeploymentState returns a boolean if a field has been set.
func (o *Deployment) HasDeploymentState() bool {
	if o != nil && o.DeploymentState.IsSet() {
		return true
	}

	return false
}

// SetDeploymentState gets a reference to the given NullableDeploymentState and assigns it to the DeploymentState field.
func (o *Deployment) SetDeploymentState(v DeploymentState) {
	o.DeploymentState.Set(&v)
}

// SetDeploymentStateNil sets the value for DeploymentState to be an explicit nil
func (o *Deployment) SetDeploymentStateNil() {
	o.DeploymentState.Set(nil)
}

// UnsetDeploymentState ensures that no value is present for DeploymentState, not even an explicit nil
func (o *Deployment) UnsetDeploymentState() {
	o.DeploymentState.Unset()
}

// GetUserCount returns the UserCount field value if set, zero value otherwise.
func (o *Deployment) GetUserCount() int32 {
	if o == nil || IsNil(o.UserCount) {
		var ret int32
		return ret
	}
	return *o.UserCount
}

// GetUserCountOk returns a tuple with the UserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetUserCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UserCount) {
		return nil, false
	}
	return o.UserCount, true
}

// HasUserCount returns a boolean if a field has been set.
func (o *Deployment) HasUserCount() bool {
	if o != nil && !IsNil(o.UserCount) {
		return true
	}

	return false
}

// SetUserCount gets a reference to the given int32 and assigns it to the UserCount field.
func (o *Deployment) SetUserCount(v int32) {
	o.UserCount = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Deployment) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Deployment) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *Deployment) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *Deployment) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}

// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *Deployment) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *Deployment) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetWarnings returns the Warnings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Deployment) GetWarnings() []DeploymentWarning {
	if o == nil {
		var ret []DeploymentWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Deployment) GetWarningsOk() ([]DeploymentWarning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *Deployment) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []DeploymentWarning and assigns it to the Warnings field.
func (o *Deployment) SetWarnings(v []DeploymentWarning) {
	o.Warnings = v
}

// GetActiveTasks returns the ActiveTasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Deployment) GetActiveTasks() []GetTaskAsync200Response {
	if o == nil {
		var ret []GetTaskAsync200Response
		return ret
	}
	return o.ActiveTasks
}

// GetActiveTasksOk returns a tuple with the ActiveTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Deployment) GetActiveTasksOk() ([]GetTaskAsync200Response, bool) {
	if o == nil || IsNil(o.ActiveTasks) {
		return nil, false
	}
	return o.ActiveTasks, true
}

// HasActiveTasks returns a boolean if a field has been set.
func (o *Deployment) HasActiveTasks() bool {
	if o != nil && !IsNil(o.ActiveTasks) {
		return true
	}

	return false
}

// SetActiveTasks gets a reference to the given []GetTaskAsync200Response and assigns it to the ActiveTasks field.
func (o *Deployment) SetActiveTasks(v []GetTaskAsync200Response) {
	o.ActiveTasks = v
}

// GetBrokerMachineCatalogId returns the BrokerMachineCatalogId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Deployment) GetBrokerMachineCatalogId() string {
	if o == nil || IsNil(o.BrokerMachineCatalogId.Get()) {
		var ret string
		return ret
	}
	return *o.BrokerMachineCatalogId.Get()
}

// GetBrokerMachineCatalogIdOk returns a tuple with the BrokerMachineCatalogId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Deployment) GetBrokerMachineCatalogIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrokerMachineCatalogId.Get(), o.BrokerMachineCatalogId.IsSet()
}

// HasBrokerMachineCatalogId returns a boolean if a field has been set.
func (o *Deployment) HasBrokerMachineCatalogId() bool {
	if o != nil && o.BrokerMachineCatalogId.IsSet() {
		return true
	}

	return false
}

// SetBrokerMachineCatalogId gets a reference to the given NullableString and assigns it to the BrokerMachineCatalogId field.
func (o *Deployment) SetBrokerMachineCatalogId(v string) {
	o.BrokerMachineCatalogId.Set(&v)
}

// SetBrokerMachineCatalogIdNil sets the value for BrokerMachineCatalogId to be an explicit nil
func (o *Deployment) SetBrokerMachineCatalogIdNil() {
	o.BrokerMachineCatalogId.Set(nil)
}

// UnsetBrokerMachineCatalogId ensures that no value is present for BrokerMachineCatalogId, not even an explicit nil
func (o *Deployment) UnsetBrokerMachineCatalogId() {
	o.BrokerMachineCatalogId.Unset()
}

// GetBrokerDeliveryGroupId returns the BrokerDeliveryGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Deployment) GetBrokerDeliveryGroupId() string {
	if o == nil || IsNil(o.BrokerDeliveryGroupId.Get()) {
		var ret string
		return ret
	}
	return *o.BrokerDeliveryGroupId.Get()
}

// GetBrokerDeliveryGroupIdOk returns a tuple with the BrokerDeliveryGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Deployment) GetBrokerDeliveryGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrokerDeliveryGroupId.Get(), o.BrokerDeliveryGroupId.IsSet()
}

// HasBrokerDeliveryGroupId returns a boolean if a field has been set.
func (o *Deployment) HasBrokerDeliveryGroupId() bool {
	if o != nil && o.BrokerDeliveryGroupId.IsSet() {
		return true
	}

	return false
}

// SetBrokerDeliveryGroupId gets a reference to the given NullableString and assigns it to the BrokerDeliveryGroupId field.
func (o *Deployment) SetBrokerDeliveryGroupId(v string) {
	o.BrokerDeliveryGroupId.Set(&v)
}

// SetBrokerDeliveryGroupIdNil sets the value for BrokerDeliveryGroupId to be an explicit nil
func (o *Deployment) SetBrokerDeliveryGroupIdNil() {
	o.BrokerDeliveryGroupId.Set(nil)
}

// UnsetBrokerDeliveryGroupId ensures that no value is present for BrokerDeliveryGroupId, not even an explicit nil
func (o *Deployment) UnsetBrokerDeliveryGroupId() {
	o.BrokerDeliveryGroupId.Unset()
}

// GetCitrixManaged returns the CitrixManaged field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Deployment) GetCitrixManaged() bool {
	if o == nil || IsNil(o.CitrixManaged.Get()) {
		var ret bool
		return ret
	}
	return *o.CitrixManaged.Get()
}

// GetCitrixManagedOk returns a tuple with the CitrixManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Deployment) GetCitrixManagedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CitrixManaged.Get(), o.CitrixManaged.IsSet()
}

// HasCitrixManaged returns a boolean if a field has been set.
func (o *Deployment) HasCitrixManaged() bool {
	if o != nil && o.CitrixManaged.IsSet() {
		return true
	}

	return false
}

// SetCitrixManaged gets a reference to the given NullableBool and assigns it to the CitrixManaged field.
func (o *Deployment) SetCitrixManaged(v bool) {
	o.CitrixManaged.Set(&v)
}

// SetCitrixManagedNil sets the value for CitrixManaged to be an explicit nil
func (o *Deployment) SetCitrixManagedNil() {
	o.CitrixManaged.Set(nil)
}

// UnsetCitrixManaged ensures that no value is present for CitrixManaged, not even an explicit nil
func (o *Deployment) UnsetCitrixManaged() {
	o.CitrixManaged.Unset()
}

// GetScaleSettings returns the ScaleSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Deployment) GetScaleSettings() ScaleSettings {
	if o == nil || IsNil(o.ScaleSettings.Get()) {
		var ret ScaleSettings
		return ret
	}
	return *o.ScaleSettings.Get()
}

// GetScaleSettingsOk returns a tuple with the ScaleSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Deployment) GetScaleSettingsOk() (*ScaleSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScaleSettings.Get(), o.ScaleSettings.IsSet()
}

// HasScaleSettings returns a boolean if a field has been set.
func (o *Deployment) HasScaleSettings() bool {
	if o != nil && o.ScaleSettings.IsSet() {
		return true
	}

	return false
}

// SetScaleSettings gets a reference to the given NullableScaleSettings and assigns it to the ScaleSettings field.
func (o *Deployment) SetScaleSettings(v ScaleSettings) {
	o.ScaleSettings.Set(&v)
}

// SetScaleSettingsNil sets the value for ScaleSettings to be an explicit nil
func (o *Deployment) SetScaleSettingsNil() {
	o.ScaleSettings.Set(nil)
}

// UnsetScaleSettings ensures that no value is present for ScaleSettings, not even an explicit nil
func (o *Deployment) UnsetScaleSettings() {
	o.ScaleSettings.Unset()
}

func (o Deployment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Deployment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountType"] = o.AccountType
	if o.DeploymentId.IsSet() {
		toSerialize["deploymentId"] = o.DeploymentId.Get()
	}
	if o.DeploymentName.IsSet() {
		toSerialize["deploymentName"] = o.DeploymentName.Get()
	}
	if o.AccountId.IsSet() {
		toSerialize["accountId"] = o.AccountId.Get()
	}
	if o.AccountName.IsSet() {
		toSerialize["accountName"] = o.AccountName.Get()
	}
	if o.ConnectionId.IsSet() {
		toSerialize["connectionId"] = o.ConnectionId.Get()
	}
	if o.ConnectionName.IsSet() {
		toSerialize["connectionName"] = o.ConnectionName.Get()
	}
	if o.DeploymentState.IsSet() {
		toSerialize["deploymentState"] = o.DeploymentState.Get()
	}
	if !IsNil(o.UserCount) {
		toSerialize["userCount"] = o.UserCount
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["errorMessage"] = o.ErrorMessage.Get()
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	if o.ActiveTasks != nil {
		toSerialize["activeTasks"] = o.ActiveTasks
	}
	if o.BrokerMachineCatalogId.IsSet() {
		toSerialize["brokerMachineCatalogId"] = o.BrokerMachineCatalogId.Get()
	}
	if o.BrokerDeliveryGroupId.IsSet() {
		toSerialize["brokerDeliveryGroupId"] = o.BrokerDeliveryGroupId.Get()
	}
	if o.CitrixManaged.IsSet() {
		toSerialize["citrixManaged"] = o.CitrixManaged.Get()
	}
	if o.ScaleSettings.IsSet() {
		toSerialize["scaleSettings"] = o.ScaleSettings.Get()
	}
	return toSerialize, nil
}

type NullableDeployment struct {
	value *Deployment
	isSet bool
}

func (v NullableDeployment) Get() *Deployment {
	return v.value
}

func (v *NullableDeployment) Set(val *Deployment) {
	v.value = val
	v.isSet = true
}

func (v NullableDeployment) IsSet() bool {
	return v.isSet
}

func (v *NullableDeployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeployment(val *Deployment) *NullableDeployment {
	return &NullableDeployment{value: val, isSet: true}
}

func (v NullableDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
