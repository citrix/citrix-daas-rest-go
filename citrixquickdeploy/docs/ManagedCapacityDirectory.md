# ManagedCapacityDirectory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectoryId** | **string** |  | 
**DirectoryName** | **string** |  | 
**AzurePoolName** | Pointer to **string** |  | [optional] 
**CustomerDeleted** | Pointer to **bool** |  | [optional] 
**Subscriptions** | Pointer to [**[]ManagedCapacitySubscription**](ManagedCapacitySubscription.md) |  | [optional] 
**AssignedCustomer** | Pointer to [**ManagedCapacityAssignedCustomer**](ManagedCapacityAssignedCustomer.md) |  | [optional] 

## Methods

### NewManagedCapacityDirectory

`func NewManagedCapacityDirectory(directoryId string, directoryName string, ) *ManagedCapacityDirectory`

NewManagedCapacityDirectory instantiates a new ManagedCapacityDirectory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedCapacityDirectoryWithDefaults

`func NewManagedCapacityDirectoryWithDefaults() *ManagedCapacityDirectory`

NewManagedCapacityDirectoryWithDefaults instantiates a new ManagedCapacityDirectory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectoryId

`func (o *ManagedCapacityDirectory) GetDirectoryId() string`

GetDirectoryId returns the DirectoryId field if non-nil, zero value otherwise.

### GetDirectoryIdOk

`func (o *ManagedCapacityDirectory) GetDirectoryIdOk() (*string, bool)`

GetDirectoryIdOk returns a tuple with the DirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryId

`func (o *ManagedCapacityDirectory) SetDirectoryId(v string)`

SetDirectoryId sets DirectoryId field to given value.


### GetDirectoryName

`func (o *ManagedCapacityDirectory) GetDirectoryName() string`

GetDirectoryName returns the DirectoryName field if non-nil, zero value otherwise.

### GetDirectoryNameOk

`func (o *ManagedCapacityDirectory) GetDirectoryNameOk() (*string, bool)`

GetDirectoryNameOk returns a tuple with the DirectoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryName

`func (o *ManagedCapacityDirectory) SetDirectoryName(v string)`

SetDirectoryName sets DirectoryName field to given value.


### GetAzurePoolName

`func (o *ManagedCapacityDirectory) GetAzurePoolName() string`

GetAzurePoolName returns the AzurePoolName field if non-nil, zero value otherwise.

### GetAzurePoolNameOk

`func (o *ManagedCapacityDirectory) GetAzurePoolNameOk() (*string, bool)`

GetAzurePoolNameOk returns a tuple with the AzurePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzurePoolName

`func (o *ManagedCapacityDirectory) SetAzurePoolName(v string)`

SetAzurePoolName sets AzurePoolName field to given value.

### HasAzurePoolName

`func (o *ManagedCapacityDirectory) HasAzurePoolName() bool`

HasAzurePoolName returns a boolean if a field has been set.

### GetCustomerDeleted

`func (o *ManagedCapacityDirectory) GetCustomerDeleted() bool`

GetCustomerDeleted returns the CustomerDeleted field if non-nil, zero value otherwise.

### GetCustomerDeletedOk

`func (o *ManagedCapacityDirectory) GetCustomerDeletedOk() (*bool, bool)`

GetCustomerDeletedOk returns a tuple with the CustomerDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDeleted

`func (o *ManagedCapacityDirectory) SetCustomerDeleted(v bool)`

SetCustomerDeleted sets CustomerDeleted field to given value.

### HasCustomerDeleted

`func (o *ManagedCapacityDirectory) HasCustomerDeleted() bool`

HasCustomerDeleted returns a boolean if a field has been set.

### GetSubscriptions

`func (o *ManagedCapacityDirectory) GetSubscriptions() []ManagedCapacitySubscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *ManagedCapacityDirectory) GetSubscriptionsOk() (*[]ManagedCapacitySubscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *ManagedCapacityDirectory) SetSubscriptions(v []ManagedCapacitySubscription)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *ManagedCapacityDirectory) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetAssignedCustomer

`func (o *ManagedCapacityDirectory) GetAssignedCustomer() ManagedCapacityAssignedCustomer`

GetAssignedCustomer returns the AssignedCustomer field if non-nil, zero value otherwise.

### GetAssignedCustomerOk

`func (o *ManagedCapacityDirectory) GetAssignedCustomerOk() (*ManagedCapacityAssignedCustomer, bool)`

GetAssignedCustomerOk returns a tuple with the AssignedCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedCustomer

`func (o *ManagedCapacityDirectory) SetAssignedCustomer(v ManagedCapacityAssignedCustomer)`

SetAssignedCustomer sets AssignedCustomer field to given value.

### HasAssignedCustomer

`func (o *ManagedCapacityDirectory) HasAssignedCustomer() bool`

HasAssignedCustomer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


