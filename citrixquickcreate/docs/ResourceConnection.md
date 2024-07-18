# ResourceConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | [**AccountType**](AccountType.md) |  | 
**AccountId** | Pointer to **NullableString** | The ID of the account | [optional] 
**ConnectionId** | **string** | The Id of the connection | 
**Name** | **string** | The name of the connection | 
**ConnectionState** | Pointer to [**NullableConnectionState**](ConnectionState.md) |  | [optional] 
**ZoneId** | Pointer to **NullableString** | Zone id | [optional] 
**ResourceLocationId** | Pointer to **NullableString** | Resource Location Id | [optional] 
**ActiveTasks** | Pointer to [**[]GetCustomerAccountTaskAsync200Response**](GetCustomerAccountTaskAsync200Response.md) | Tasks currently being performed on the connection | [optional] 
**CitrixManaged** | **bool** | Indicates whether the resource connection is Citrix managed | 
**AssociatedDeployments** | Pointer to [**[]AssociatedDeployment**](AssociatedDeployment.md) |  | [optional] 
**Warnings** | Pointer to [**[]ResourceConnectionWarning**](ResourceConnectionWarning.md) |  | [optional] 

## Methods

### NewResourceConnection

`func NewResourceConnection(accountType AccountType, connectionId string, name string, citrixManaged bool, ) *ResourceConnection`

NewResourceConnection instantiates a new ResourceConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceConnectionWithDefaults

`func NewResourceConnectionWithDefaults() *ResourceConnection`

NewResourceConnectionWithDefaults instantiates a new ResourceConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *ResourceConnection) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *ResourceConnection) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *ResourceConnection) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetAccountId

`func (o *ResourceConnection) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ResourceConnection) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ResourceConnection) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ResourceConnection) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *ResourceConnection) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *ResourceConnection) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetConnectionId

`func (o *ResourceConnection) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ResourceConnection) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ResourceConnection) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetName

`func (o *ResourceConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceConnection) SetName(v string)`

SetName sets Name field to given value.


### GetConnectionState

`func (o *ResourceConnection) GetConnectionState() ConnectionState`

GetConnectionState returns the ConnectionState field if non-nil, zero value otherwise.

### GetConnectionStateOk

`func (o *ResourceConnection) GetConnectionStateOk() (*ConnectionState, bool)`

GetConnectionStateOk returns a tuple with the ConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionState

`func (o *ResourceConnection) SetConnectionState(v ConnectionState)`

SetConnectionState sets ConnectionState field to given value.

### HasConnectionState

`func (o *ResourceConnection) HasConnectionState() bool`

HasConnectionState returns a boolean if a field has been set.

### SetConnectionStateNil

`func (o *ResourceConnection) SetConnectionStateNil(b bool)`

 SetConnectionStateNil sets the value for ConnectionState to be an explicit nil

### UnsetConnectionState
`func (o *ResourceConnection) UnsetConnectionState()`

UnsetConnectionState ensures that no value is present for ConnectionState, not even an explicit nil
### GetZoneId

`func (o *ResourceConnection) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ResourceConnection) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ResourceConnection) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ResourceConnection) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### SetZoneIdNil

`func (o *ResourceConnection) SetZoneIdNil(b bool)`

 SetZoneIdNil sets the value for ZoneId to be an explicit nil

### UnsetZoneId
`func (o *ResourceConnection) UnsetZoneId()`

UnsetZoneId ensures that no value is present for ZoneId, not even an explicit nil
### GetResourceLocationId

`func (o *ResourceConnection) GetResourceLocationId() string`

GetResourceLocationId returns the ResourceLocationId field if non-nil, zero value otherwise.

### GetResourceLocationIdOk

`func (o *ResourceConnection) GetResourceLocationIdOk() (*string, bool)`

GetResourceLocationIdOk returns a tuple with the ResourceLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLocationId

`func (o *ResourceConnection) SetResourceLocationId(v string)`

SetResourceLocationId sets ResourceLocationId field to given value.

### HasResourceLocationId

`func (o *ResourceConnection) HasResourceLocationId() bool`

HasResourceLocationId returns a boolean if a field has been set.

### SetResourceLocationIdNil

`func (o *ResourceConnection) SetResourceLocationIdNil(b bool)`

 SetResourceLocationIdNil sets the value for ResourceLocationId to be an explicit nil

### UnsetResourceLocationId
`func (o *ResourceConnection) UnsetResourceLocationId()`

UnsetResourceLocationId ensures that no value is present for ResourceLocationId, not even an explicit nil
### GetActiveTasks

`func (o *ResourceConnection) GetActiveTasks() []GetCustomerAccountTaskAsync200Response`

GetActiveTasks returns the ActiveTasks field if non-nil, zero value otherwise.

### GetActiveTasksOk

`func (o *ResourceConnection) GetActiveTasksOk() (*[]GetCustomerAccountTaskAsync200Response, bool)`

GetActiveTasksOk returns a tuple with the ActiveTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTasks

`func (o *ResourceConnection) SetActiveTasks(v []GetCustomerAccountTaskAsync200Response)`

SetActiveTasks sets ActiveTasks field to given value.

### HasActiveTasks

`func (o *ResourceConnection) HasActiveTasks() bool`

HasActiveTasks returns a boolean if a field has been set.

### SetActiveTasksNil

`func (o *ResourceConnection) SetActiveTasksNil(b bool)`

 SetActiveTasksNil sets the value for ActiveTasks to be an explicit nil

### UnsetActiveTasks
`func (o *ResourceConnection) UnsetActiveTasks()`

UnsetActiveTasks ensures that no value is present for ActiveTasks, not even an explicit nil
### GetCitrixManaged

`func (o *ResourceConnection) GetCitrixManaged() bool`

GetCitrixManaged returns the CitrixManaged field if non-nil, zero value otherwise.

### GetCitrixManagedOk

`func (o *ResourceConnection) GetCitrixManagedOk() (*bool, bool)`

GetCitrixManagedOk returns a tuple with the CitrixManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitrixManaged

`func (o *ResourceConnection) SetCitrixManaged(v bool)`

SetCitrixManaged sets CitrixManaged field to given value.


### GetAssociatedDeployments

`func (o *ResourceConnection) GetAssociatedDeployments() []AssociatedDeployment`

GetAssociatedDeployments returns the AssociatedDeployments field if non-nil, zero value otherwise.

### GetAssociatedDeploymentsOk

`func (o *ResourceConnection) GetAssociatedDeploymentsOk() (*[]AssociatedDeployment, bool)`

GetAssociatedDeploymentsOk returns a tuple with the AssociatedDeployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedDeployments

`func (o *ResourceConnection) SetAssociatedDeployments(v []AssociatedDeployment)`

SetAssociatedDeployments sets AssociatedDeployments field to given value.

### HasAssociatedDeployments

`func (o *ResourceConnection) HasAssociatedDeployments() bool`

HasAssociatedDeployments returns a boolean if a field has been set.

### SetAssociatedDeploymentsNil

`func (o *ResourceConnection) SetAssociatedDeploymentsNil(b bool)`

 SetAssociatedDeploymentsNil sets the value for AssociatedDeployments to be an explicit nil

### UnsetAssociatedDeployments
`func (o *ResourceConnection) UnsetAssociatedDeployments()`

UnsetAssociatedDeployments ensures that no value is present for AssociatedDeployments, not even an explicit nil
### GetWarnings

`func (o *ResourceConnection) GetWarnings() []ResourceConnectionWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ResourceConnection) GetWarningsOk() (*[]ResourceConnectionWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ResourceConnection) SetWarnings(v []ResourceConnectionWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ResourceConnection) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *ResourceConnection) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *ResourceConnection) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


