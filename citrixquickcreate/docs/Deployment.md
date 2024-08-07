# Deployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | [**AccountType**](AccountType.md) |  | 
**DeploymentId** | Pointer to **NullableString** | Deployment Id | [optional] 
**DeploymentName** | Pointer to **NullableString** | Deployment Name | [optional] 
**AccountId** | Pointer to **NullableString** | Account Id | [optional] 
**AccountName** | Pointer to **NullableString** | Account name | [optional] 
**ConnectionId** | Pointer to **NullableString** | Connection Id | [optional] 
**ConnectionName** | Pointer to **NullableString** | Connection Name | [optional] 
**DeploymentState** | Pointer to [**NullableDeploymentState**](DeploymentState.md) |  | [optional] 
**UserCount** | Pointer to **int32** | The number of users in this deployment | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Error message associated with the deployment | [optional] 
**Warnings** | Pointer to [**[]DeploymentWarning**](DeploymentWarning.md) | Warnings and errors associated with the deployment | [optional] 
**ActiveTasks** | Pointer to [**[]GetTaskAsync200Response**](GetTaskAsync200Response.md) | Tasks currently being performed on the deployment | [optional] 
**BrokerMachineCatalogId** | Pointer to **NullableString** | Id for the machine catalog of the deployment, could be null | [optional] 
**BrokerDeliveryGroupId** | Pointer to **NullableString** | Id for the delivery group of the deployment, could be null | [optional] 
**CitrixManaged** | Pointer to **NullableBool** | Indicates whether the deployment is managed by Citrix | [optional] 
**ScaleSettings** | Pointer to [**NullableDeploymentScaleSettings**](DeploymentScaleSettings.md) |  | [optional] 

## Methods

### NewDeployment

`func NewDeployment(accountType AccountType, ) *Deployment`

NewDeployment instantiates a new Deployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentWithDefaults

`func NewDeploymentWithDefaults() *Deployment`

NewDeploymentWithDefaults instantiates a new Deployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *Deployment) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *Deployment) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *Deployment) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetDeploymentId

`func (o *Deployment) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *Deployment) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *Deployment) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *Deployment) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### SetDeploymentIdNil

`func (o *Deployment) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *Deployment) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
### GetDeploymentName

`func (o *Deployment) GetDeploymentName() string`

GetDeploymentName returns the DeploymentName field if non-nil, zero value otherwise.

### GetDeploymentNameOk

`func (o *Deployment) GetDeploymentNameOk() (*string, bool)`

GetDeploymentNameOk returns a tuple with the DeploymentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentName

`func (o *Deployment) SetDeploymentName(v string)`

SetDeploymentName sets DeploymentName field to given value.

### HasDeploymentName

`func (o *Deployment) HasDeploymentName() bool`

HasDeploymentName returns a boolean if a field has been set.

### SetDeploymentNameNil

`func (o *Deployment) SetDeploymentNameNil(b bool)`

 SetDeploymentNameNil sets the value for DeploymentName to be an explicit nil

### UnsetDeploymentName
`func (o *Deployment) UnsetDeploymentName()`

UnsetDeploymentName ensures that no value is present for DeploymentName, not even an explicit nil
### GetAccountId

`func (o *Deployment) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Deployment) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Deployment) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Deployment) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *Deployment) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *Deployment) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetAccountName

`func (o *Deployment) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *Deployment) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *Deployment) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *Deployment) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### SetAccountNameNil

`func (o *Deployment) SetAccountNameNil(b bool)`

 SetAccountNameNil sets the value for AccountName to be an explicit nil

### UnsetAccountName
`func (o *Deployment) UnsetAccountName()`

UnsetAccountName ensures that no value is present for AccountName, not even an explicit nil
### GetConnectionId

`func (o *Deployment) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *Deployment) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *Deployment) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *Deployment) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *Deployment) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *Deployment) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetConnectionName

`func (o *Deployment) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *Deployment) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *Deployment) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.

### HasConnectionName

`func (o *Deployment) HasConnectionName() bool`

HasConnectionName returns a boolean if a field has been set.

### SetConnectionNameNil

`func (o *Deployment) SetConnectionNameNil(b bool)`

 SetConnectionNameNil sets the value for ConnectionName to be an explicit nil

### UnsetConnectionName
`func (o *Deployment) UnsetConnectionName()`

UnsetConnectionName ensures that no value is present for ConnectionName, not even an explicit nil
### GetDeploymentState

`func (o *Deployment) GetDeploymentState() DeploymentState`

GetDeploymentState returns the DeploymentState field if non-nil, zero value otherwise.

### GetDeploymentStateOk

`func (o *Deployment) GetDeploymentStateOk() (*DeploymentState, bool)`

GetDeploymentStateOk returns a tuple with the DeploymentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentState

`func (o *Deployment) SetDeploymentState(v DeploymentState)`

SetDeploymentState sets DeploymentState field to given value.

### HasDeploymentState

`func (o *Deployment) HasDeploymentState() bool`

HasDeploymentState returns a boolean if a field has been set.

### SetDeploymentStateNil

`func (o *Deployment) SetDeploymentStateNil(b bool)`

 SetDeploymentStateNil sets the value for DeploymentState to be an explicit nil

### UnsetDeploymentState
`func (o *Deployment) UnsetDeploymentState()`

UnsetDeploymentState ensures that no value is present for DeploymentState, not even an explicit nil
### GetUserCount

`func (o *Deployment) GetUserCount() int32`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *Deployment) GetUserCountOk() (*int32, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *Deployment) SetUserCount(v int32)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *Deployment) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.

### GetErrorMessage

`func (o *Deployment) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Deployment) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Deployment) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Deployment) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *Deployment) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *Deployment) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetWarnings

`func (o *Deployment) GetWarnings() []DeploymentWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *Deployment) GetWarningsOk() (*[]DeploymentWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *Deployment) SetWarnings(v []DeploymentWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *Deployment) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *Deployment) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *Deployment) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
### GetActiveTasks

`func (o *Deployment) GetActiveTasks() []GetTaskAsync200Response`

GetActiveTasks returns the ActiveTasks field if non-nil, zero value otherwise.

### GetActiveTasksOk

`func (o *Deployment) GetActiveTasksOk() (*[]GetTaskAsync200Response, bool)`

GetActiveTasksOk returns a tuple with the ActiveTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTasks

`func (o *Deployment) SetActiveTasks(v []GetTaskAsync200Response)`

SetActiveTasks sets ActiveTasks field to given value.

### HasActiveTasks

`func (o *Deployment) HasActiveTasks() bool`

HasActiveTasks returns a boolean if a field has been set.

### SetActiveTasksNil

`func (o *Deployment) SetActiveTasksNil(b bool)`

 SetActiveTasksNil sets the value for ActiveTasks to be an explicit nil

### UnsetActiveTasks
`func (o *Deployment) UnsetActiveTasks()`

UnsetActiveTasks ensures that no value is present for ActiveTasks, not even an explicit nil
### GetBrokerMachineCatalogId

`func (o *Deployment) GetBrokerMachineCatalogId() string`

GetBrokerMachineCatalogId returns the BrokerMachineCatalogId field if non-nil, zero value otherwise.

### GetBrokerMachineCatalogIdOk

`func (o *Deployment) GetBrokerMachineCatalogIdOk() (*string, bool)`

GetBrokerMachineCatalogIdOk returns a tuple with the BrokerMachineCatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerMachineCatalogId

`func (o *Deployment) SetBrokerMachineCatalogId(v string)`

SetBrokerMachineCatalogId sets BrokerMachineCatalogId field to given value.

### HasBrokerMachineCatalogId

`func (o *Deployment) HasBrokerMachineCatalogId() bool`

HasBrokerMachineCatalogId returns a boolean if a field has been set.

### SetBrokerMachineCatalogIdNil

`func (o *Deployment) SetBrokerMachineCatalogIdNil(b bool)`

 SetBrokerMachineCatalogIdNil sets the value for BrokerMachineCatalogId to be an explicit nil

### UnsetBrokerMachineCatalogId
`func (o *Deployment) UnsetBrokerMachineCatalogId()`

UnsetBrokerMachineCatalogId ensures that no value is present for BrokerMachineCatalogId, not even an explicit nil
### GetBrokerDeliveryGroupId

`func (o *Deployment) GetBrokerDeliveryGroupId() string`

GetBrokerDeliveryGroupId returns the BrokerDeliveryGroupId field if non-nil, zero value otherwise.

### GetBrokerDeliveryGroupIdOk

`func (o *Deployment) GetBrokerDeliveryGroupIdOk() (*string, bool)`

GetBrokerDeliveryGroupIdOk returns a tuple with the BrokerDeliveryGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerDeliveryGroupId

`func (o *Deployment) SetBrokerDeliveryGroupId(v string)`

SetBrokerDeliveryGroupId sets BrokerDeliveryGroupId field to given value.

### HasBrokerDeliveryGroupId

`func (o *Deployment) HasBrokerDeliveryGroupId() bool`

HasBrokerDeliveryGroupId returns a boolean if a field has been set.

### SetBrokerDeliveryGroupIdNil

`func (o *Deployment) SetBrokerDeliveryGroupIdNil(b bool)`

 SetBrokerDeliveryGroupIdNil sets the value for BrokerDeliveryGroupId to be an explicit nil

### UnsetBrokerDeliveryGroupId
`func (o *Deployment) UnsetBrokerDeliveryGroupId()`

UnsetBrokerDeliveryGroupId ensures that no value is present for BrokerDeliveryGroupId, not even an explicit nil
### GetCitrixManaged

`func (o *Deployment) GetCitrixManaged() bool`

GetCitrixManaged returns the CitrixManaged field if non-nil, zero value otherwise.

### GetCitrixManagedOk

`func (o *Deployment) GetCitrixManagedOk() (*bool, bool)`

GetCitrixManagedOk returns a tuple with the CitrixManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitrixManaged

`func (o *Deployment) SetCitrixManaged(v bool)`

SetCitrixManaged sets CitrixManaged field to given value.

### HasCitrixManaged

`func (o *Deployment) HasCitrixManaged() bool`

HasCitrixManaged returns a boolean if a field has been set.

### SetCitrixManagedNil

`func (o *Deployment) SetCitrixManagedNil(b bool)`

 SetCitrixManagedNil sets the value for CitrixManaged to be an explicit nil

### UnsetCitrixManaged
`func (o *Deployment) UnsetCitrixManaged()`

UnsetCitrixManaged ensures that no value is present for CitrixManaged, not even an explicit nil
### GetScaleSettings

`func (o *Deployment) GetScaleSettings() DeploymentScaleSettings`

GetScaleSettings returns the ScaleSettings field if non-nil, zero value otherwise.

### GetScaleSettingsOk

`func (o *Deployment) GetScaleSettingsOk() (*DeploymentScaleSettings, bool)`

GetScaleSettingsOk returns a tuple with the ScaleSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleSettings

`func (o *Deployment) SetScaleSettings(v DeploymentScaleSettings)`

SetScaleSettings sets ScaleSettings field to given value.

### HasScaleSettings

`func (o *Deployment) HasScaleSettings() bool`

HasScaleSettings returns a boolean if a field has been set.

### SetScaleSettingsNil

`func (o *Deployment) SetScaleSettingsNil(b bool)`

 SetScaleSettingsNil sets the value for ScaleSettings to be an explicit nil

### UnsetScaleSettings
`func (o *Deployment) UnsetScaleSettings()`

UnsetScaleSettings ensures that no value is present for ScaleSettings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


