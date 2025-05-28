# AzureResourceGroupStorageAccountsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceGroup** | **string** | Name of the resource group the storge accounts are attached to | 
**Items** | [**[]AzureStorageAccount**](AzureStorageAccount.md) | List of Storage Account that are configured for the specified resource group | 
**StorageAccounts** | Pointer to [**[]AzureStorageAccount**](AzureStorageAccount.md) | Alias of Items property for backward compatibility | [optional] 

## Methods

### NewAzureResourceGroupStorageAccountsModel

`func NewAzureResourceGroupStorageAccountsModel(resourceGroup string, items []AzureStorageAccount, ) *AzureResourceGroupStorageAccountsModel`

NewAzureResourceGroupStorageAccountsModel instantiates a new AzureResourceGroupStorageAccountsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureResourceGroupStorageAccountsModelWithDefaults

`func NewAzureResourceGroupStorageAccountsModelWithDefaults() *AzureResourceGroupStorageAccountsModel`

NewAzureResourceGroupStorageAccountsModelWithDefaults instantiates a new AzureResourceGroupStorageAccountsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceGroup

`func (o *AzureResourceGroupStorageAccountsModel) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *AzureResourceGroupStorageAccountsModel) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *AzureResourceGroupStorageAccountsModel) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.


### GetItems

`func (o *AzureResourceGroupStorageAccountsModel) GetItems() []AzureStorageAccount`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AzureResourceGroupStorageAccountsModel) GetItemsOk() (*[]AzureStorageAccount, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AzureResourceGroupStorageAccountsModel) SetItems(v []AzureStorageAccount)`

SetItems sets Items field to given value.


### GetStorageAccounts

`func (o *AzureResourceGroupStorageAccountsModel) GetStorageAccounts() []AzureStorageAccount`

GetStorageAccounts returns the StorageAccounts field if non-nil, zero value otherwise.

### GetStorageAccountsOk

`func (o *AzureResourceGroupStorageAccountsModel) GetStorageAccountsOk() (*[]AzureStorageAccount, bool)`

GetStorageAccountsOk returns a tuple with the StorageAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccounts

`func (o *AzureResourceGroupStorageAccountsModel) SetStorageAccounts(v []AzureStorageAccount)`

SetStorageAccounts sets StorageAccounts field to given value.

### HasStorageAccounts

`func (o *AzureResourceGroupStorageAccountsModel) HasStorageAccounts() bool`

HasStorageAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


